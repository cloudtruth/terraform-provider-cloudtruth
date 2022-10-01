package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"strings"
)

const ruleOperationRetries = 5

func resourceType() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth Type, a user-defined type which can be used to create CloudTruth Parameters",

		CreateContext: resourceTypeCreate,
		ReadContext:   resourceTypeRead,
		UpdateContext: resourceTypeUpdate,
		DeleteContext: resourceTypeDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the CloudTruth Type, unique per project",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "A description of the CloudTruth Type",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: "The base type of this custom type, can be a builtin (string|int|boolean) or another custom type",
				Type:        schema.TypeString,
				Required:    true,
			},
			"min": {
				Description: `A rule constraint: the minimum value for integer types, the minimum length for string types.
This value is specified as a string because we need a reliable way to distinguish between default/zero values and unset values.  We use
the empty string value for that purpose.`,
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"min_id": {
				Description: "The internal ID of the min rule",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"max": {
				Description: `A rule constraint: the maximum value for integer types, the maximum length for string types.
This value is specified as a string because we need a reliable way to distinguish between default/zero values and unset values.  We use
the empty string value for that purpose.`,
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"max_id": {
				Description: "The internal ID of the max rule",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"regex": {
				Description: "A CloudTruth rule constraint: the regular expression a string type must match, only valid with string types",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"regex_id": {
				Description: "The internal ID of the regex rule",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceTypeCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceTypeCreate")
	defer tflog.Debug(ctx, "exiting resourceTypeCreate")
	c := meta.(*cloudTruthClient)
	typeName := d.Get("name").(string)
	typeCreate := cloudtruthapi.NewParameterTypeCreate(typeName)
	typeDesc := d.Get("description").(string)
	if typeDesc != "" {
		typeCreate.SetDescription(typeDesc)
	}
	// concreteType is the underlying type (builtin or custom) of the cloudtruth_type resource
	concreteType := c.lookupType(ctx, d.Get("type").(string))
	baseParamType := getBaseParamType(ctx, concreteType, c)

	var rules map[string]any
	var err error
	if rules, err = validateAndFetchRules(ctx, c, d, concreteType); err != nil {
		return diag.FromErr(err)
	}

	typeCreate.SetParent(concreteType.GetUrl())
	var typeID string
	var parameterType *cloudtruthapi.ParameterType
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		parameterType, r, err = c.openAPIClient.TypesApi.TypesCreate(ctx).ParameterTypeCreate(*typeCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceTypeCreate: error creating type %s", typeName), r, err)
		}
		typeID = parameterType.GetId()
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	for ruleName, ruleVal := range rules {
		ruleID, err := addRuleToType(ctx, c, typeID, baseParamType, ruleName, ruleVal.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		// Compute the rule ID
		err = d.Set(fmt.Sprintf("%s_id", ruleName), ruleID)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	c.types[typeName] = *parameterType
	d.SetId(typeID)
	return nil
}

func addRuleToType(ctx context.Context, c *cloudTruthClient, typeID string, baseParamType, ruleName, ruleVal string) (*string, error) {
	tflog.Debug(ctx, "entering addRuleToType")
	defer tflog.Debug(ctx, "exiting addRuleToType")
	var apiError error
	if baseParamType == "string" {
		if ruleName == "max" || ruleName == "min" {
			// the API string rule names are min_len and max_len not max/min which are only for integers
			// in the provider we use max/min for strings too, to avoid having 4 properties: max/min/max_len/min_len
			ruleName = fmt.Sprintf("%s_len", ruleName)
		}
	}
	var createTypeRule cloudtruthapi.ParameterTypeRuleCreate
	typeEnum, err := cloudtruthapi.NewParameterRuleTypeEnumFromValue(ruleName)
	if err != nil {
		return nil, err
	}

	createTypeRule.SetType(*typeEnum)
	createTypeRule.SetConstraint(ruleVal)

	var typeRule *cloudtruthapi.ParameterTypeRule
	var r *http.Response
	retryCount := 0
	for retryCount < ruleOperationRetries {
		typeRule, r, err = c.openAPIClient.TypesApi.TypesRulesCreate(ctx, typeID).ParameterTypeRuleCreate(createTypeRule).Execute()
		if r.StatusCode >= 500 {
			apiError = err
			retryCount++
		} else if r.StatusCode >= 400 {
			clientErr := err.(*cloudtruthapi.GenericOpenAPIError)
			return nil, fmt.Errorf("client %d error: %s", r.StatusCode, clientErr.Body())
		} else {
			apiError = nil
			break
		}
	}
	if apiError != nil {
		return nil, apiError
	}
	typeRuleID := typeRule.GetId()
	return &typeRuleID, nil
}

func resourceTypeRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceTypeRead")
	defer tflog.Debug(ctx, "exiting resourceTypeCreate")
	c := meta.(*cloudTruthClient)
	typeName := d.Get("name").(string)
	typeID := d.Id()

	var parameterType *cloudtruthapi.ParameterType
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		parameterType, r, err = c.openAPIClient.TypesApi.TypesRetrieve(ctx, typeID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceTypeRead: error reading type %s", typeName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	err := d.Set("name", parameterType.GetName())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", parameterType.GetDescription())
	if err != nil {
		return diag.FromErr(err)
	}
	for _, rule := range parameterType.Rules {
		ruleName := string(rule.GetType())
		if strings.Contains(ruleName, "_len") {
			ruleName = ruleName[:3] // e.g. max_len -> max
		}
		err = d.Set(ruleName, rule.GetConstraint())
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(parameterType.GetId())
	return nil
}

func updateTypeRules(ctx context.Context, paramID, typeName string, d *schema.ResourceData,
	c *cloudTruthClient) (*http.Response, error) {
	tflog.Debug(ctx, "entering updateTypeRules")
	defer tflog.Debug(ctx, "exiting updateTypeRules")
	paramType := c.lookupType(ctx, typeName)
	baseParamType := getBaseParamType(ctx, paramType, c)
	rules := intAndStringRuleTypes
	if baseParamType == "string" { // regex types are valid with string types but not int types
		rules = append(rules, stringRuleTypes...)
	}
	var updatedRules []string
	for _, rule := range rules {
		if d.HasChange(rule) {
			updatedRules = append(updatedRules, rule)
		}
	}

	// We return the last response (and stop if we receive an error) however, we could consider returning all of the http responses
	var r *http.Response
	var err error
	if len(updatedRules) > 0 {
		for _, updatedRule := range updatedRules {
			ruleIDProperty := fmt.Sprintf("%s_id", updatedRule)
			ruleID := d.Get(ruleIDProperty).(string)
			newRuleVal := d.Get(updatedRule).(string)
			if newRuleVal == "" { // A value of "" == time for deletion
				r, err = deleteTypeRule(ctx, paramID, updatedRule, ruleID, c)
				if err == nil { // Don't set the property to "" unless the delete succeeds
					err = d.Set(ruleIDProperty, "")
				}
			} else {
				r, err = updateTypeRule(ctx, paramID, updatedRule, ruleID, baseParamType, d, c)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return r, nil
}

// Type rules are NOT associated with a CloudTruth project, unlike parameter rules which are project specific
func updateTypeRule(ctx context.Context, paramID, ruleName, ruleID, baseParamType string, d *schema.ResourceData, c *cloudTruthClient) (*http.Response, error) {
	tflog.Debug(ctx, "entering updateTypeRule")
	defer tflog.Debug(ctx, "exiting updateTypeRule")
	_, newVal := d.GetChange(ruleName)
	ruleVal := newVal.(string)
	retryCount := 0
	var apiError error
	var ruleUpdateRequest cloudtruthapi.ApiTypesRulesUpdateRequest

	// The provider uses max/min for strings instead of max_len/min_len
	var providerRuleName = ruleName
	if baseParamType == "string" {
		if strings.Contains(providerRuleName, "max") || strings.Contains(providerRuleName, "min") {
			providerRuleName = providerRuleName + "_len"
		}
	}
	ruleType, err := cloudtruthapi.NewParameterRuleTypeEnumFromValue(providerRuleName)
	if err != nil {
		return nil, err
	}
	paramTypeRule := cloudtruthapi.NewParameterTypeRuleWithDefaults()
	paramTypeRule.SetId(ruleID)
	paramTypeRule.SetConstraint(ruleVal)
	paramTypeRule.SetType(*ruleType)

	var r *http.Response
	for retryCount < ruleOperationRetries {
		ruleUpdateRequest = c.openAPIClient.TypesApi.TypesRulesUpdate(ctx, ruleID, paramID).ParameterTypeRule(*paramTypeRule)
		_, r, err = ruleUpdateRequest.Execute()
		if r.StatusCode >= 500 {
			tflog.Debug(ctx, fmt.Sprintf("updateTypeRule: error updating rule %s: %+v", ruleName, err))
			apiError = err
			retryCount++
		} else {
			apiError = nil
			break
		}
	}
	if apiError != nil {
		return nil, apiError
	}
	return r, nil
}

func deleteTypeRule(ctx context.Context, paramID, ruleName, ruleID string, c *cloudTruthClient) (*http.Response, error) {
	tflog.Debug(ctx, "entering deleteTypeRule")
	defer tflog.Debug(ctx, "exiting deleteTypeRule")
	retryCount := 0
	var ruleDestroyRequest cloudtruthapi.ApiTypesRulesDestroyRequest
	var r *http.Response
	var apiError, err error

	for retryCount < ruleOperationRetries {
		ruleDestroyRequest = c.openAPIClient.TypesApi.TypesRulesDestroy(ctx, ruleID, paramID)
		r, err = ruleDestroyRequest.Execute()
		if r.StatusCode >= 500 {
			tflog.Debug(ctx, fmt.Sprintf("deleteTypeRule: error deleting rule %s: %+v", ruleName, err))
			apiError = err
			retryCount++
		} else {
			apiError = nil
			break
		}
	}
	if apiError != nil {
		return nil, apiError
	}
	return r, nil
}

func resourceTypeUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceTypeUpdate")
	defer tflog.Debug(ctx, "exiting resourceTypeUpdate")
	c := meta.(*cloudTruthClient)
	typeName := d.Get("name").(string)
	hasChange := false
	patchedTypeUpdate := cloudtruthapi.NewPatchedParameterType()
	if d.HasChange("description") {
		typeDesc := d.Get("description").(string)
		patchedTypeUpdate.SetDescription(typeDesc)
		hasChange = true
	}

	typeID := d.Id()
	if hasChange {
		retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.TypesApi.TypesPartialUpdate(ctx, typeID).PatchedParameterType(*patchedTypeUpdate).Execute()
			if err != nil {
				return handleAPIError(fmt.Sprintf("resourceTypeUpdate: error updating type %s", typeName), r, err)
			}
			return nil
		})
		if retryError != nil {
			return diag.FromErr(retryError)
		}
	}

	// Rule changes
	if d.HasChange("max") || d.HasChange("min") || d.HasChange("regex") {
		retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
			r, err := updateTypeRules(ctx, typeID, typeName, d, c)
			if err != nil {
				return handleAPIError(fmt.Sprintf("resourceParameterUpdate: error updating rules for parameter %s", typeName), r, err)
			}
			return nil
		})
		if retryError != nil {
			return diag.FromErr(retryError)
		}
	}
	return resourceTypeRead(ctx, d, meta)
}

func resourceTypeDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceTypeDelete")
	defer tflog.Debug(ctx, "exiting resourceTypeDelete")
	c := meta.(*cloudTruthClient)
	typeName := d.Get("name").(string)
	typeID := d.Id()

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.TypesApi.TypesDestroy(ctx, typeID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceTypeDelete: error deleting type %s", typeName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	d.SetId(typeID)
	delete(c.types, typeName)
	return nil
}
