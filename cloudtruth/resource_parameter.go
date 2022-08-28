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
)

var intAndStringRuleTypes = []string{"min", "max"}
var stringRuletypes = []string{"regex"}

func resourceParameter() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth Parameter and environment specific value (defaulting to the 'default' environment)",

		CreateContext: resourceParameterCreate,
		ReadContext:   resourceParameterRead,
		UpdateContext: resourceParameterUpdate,
		DeleteContext: resourceParameterDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the CloudTruth Parameter, unique per project",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the CloudTruth Parameter",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project": {
				Description: "The CloudTruth project where the Parameter is defined",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"secret": {
				Description: "Whether or not the Parameter is a secret, defaults to false (non-secret)",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"type": {
				Description: "The parameter type, can be a builtin (string|int|boolean) or another custom type",
				Type:        schema.TypeString,
				Optional:    true, // Optional except when one or more rules are specified
				Default:     "string",
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

func resourceParameterCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "entering resourceParameterCreate")
	defer tflog.Debug(ctx, "exiting resourceParameterCreate")
	paramName := d.Get("name").(string)
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}
	paramTypeName := d.Get("type").(string)
	var paramType *cloudtruthapi.ParameterType
	if paramTypeName != "" {
		paramType = c.lookupType(ctx, paramTypeName)
		if paramType == nil {
			return diag.FromErr(fmt.Errorf("unknown parameter type %s", paramTypeName))
		}
	}
	var rules map[string]any
	if rules, err = validateAndFetchRules(ctx, c, d, paramType); err != nil {
		return diag.FromErr(err)
	}

	var paramID string
	var paramCreateResp *cloudtruthapi.Parameter
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		paramCreate := paramCreateConfig(d, paramTypeName)
		paramCreateResp, r, err = c.openAPIClient.ProjectsApi.ProjectsParametersCreate(ctx,
			*projID).ParameterCreate(*paramCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceParameterCreate: error creating parameter %s", paramName), r, err)
		}
		paramID = paramCreateResp.GetId()
		return nil
	})

	if retryError != nil {
		return diag.FromErr(retryError)
	}
	baseParamType := getBaseParamType(ctx, paramType, c)
	for ruleName, ruleVal := range rules {
		ruleID, err := addRuleToParam(ctx, c, paramID, *projID, baseParamType, ruleName, ruleVal)
		if err != nil {
			return diag.FromErr(err)
		}
		// Compute the rule ID
		err = d.Set(fmt.Sprintf("%s_id", ruleName), ruleID)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId(paramID)
	return nil
}

// boolean base types cannot have rules
// integer base types can have up to two rules: min and max
// string base types can have up to three rules: min_len, max_len and regex but the cloudtruth_parameter resource
// specifies max_len and min_len as 'max' and 'min' to minimize the number of fields
func validateAndFetchRules(ctx context.Context, c *cloudTruthClient, d *schema.ResourceData,
	paramType *cloudtruthapi.ParameterType) (map[string]any, error) {
	baseParamTypeName := getBaseParamType(ctx, paramType, c)
	outRules := map[string]any{}
	if baseParamTypeName == "boolean" {
		return nil, fmt.Errorf("the base type 'boolean' does not support rules")
	}
	for _, r := range append(intAndStringRuleTypes, stringRuletypes...) {
		if val, ok := d.GetOk(r); ok {
			if baseParamTypeName == "integer" && r == "regex" {
				return nil, fmt.Errorf("the base type '%s' does not support the %s rule type", baseParamTypeName, r)
			}
			outRules[r] = val
		}
	}
	return outRules, nil
}

func getBaseParamType(ctx context.Context, paramType *cloudtruthapi.ParameterType, c *cloudTruthClient) string {
	var baseParamTypeName string
	baseParamType := paramType
	for {
		if baseParamType.ParentName.IsSet() && baseParamType.Parent.Get() != nil {
			baseParamType = c.lookupType(ctx, baseParamType.GetParent())
		} else {
			baseParamTypeName = baseParamType.GetName()
			break
		}
	}
	return baseParamTypeName
}

func addRuleToParam(ctx context.Context, c *cloudTruthClient, paramID, projectID, baseParamType, ruleName string, ruleVal any) (*string, error) {
	tflog.Debug(ctx, "entering addRuleToType")
	defer tflog.Debug(ctx, "exiting addRuleToType")
	retryCount := 0
	var apiError error
	var createTypeRule cloudtruthapi.ParameterRuleCreate
	if baseParamType == "string" {
		if ruleName == "max" || ruleName == "min" {
			ruleName = fmt.Sprintf("%s_len", ruleName) // the API string rule names are min_len and max_len not max/min
		}
	}
	typeEnum, err := cloudtruthapi.NewParameterRuleTypeEnumFromValue(ruleName)
	if err != nil {
		return nil, err
	}
	createTypeRule.SetType(*typeEnum)
	createTypeRule.SetConstraint(ruleVal.(string))

	var paramRule *cloudtruthapi.ParameterRule
	var r *http.Response
	for retryCount < ruleOperationRetries {
		paramRule, r, err = c.openAPIClient.ProjectsApi.ProjectsParametersRulesCreate(ctx, paramID, projectID).ParameterRuleCreate(createTypeRule).Execute()
		if r.StatusCode >= 500 {
			tflog.Debug(ctx, fmt.Sprintf("addRuleToType: %s", err))
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
	ruleID := paramRule.GetId()
	return &ruleID, nil
}

func paramCreateConfig(d *schema.ResourceData, paramType string) *cloudtruthapi.ParameterCreate {
	paramName := d.Get("name").(string)
	paramCreate := cloudtruthapi.NewParameterCreate(paramName)
	paramDesc := d.Get("description").(string)
	if paramDesc != "" {
		paramCreate.SetDescription(paramDesc)
	}
	paramCreate.SetType(paramType)
	isSecret := d.Get("secret").(bool)
	paramCreate.SetSecret(isSecret)
	return paramCreate
}

func resourceParameterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceParameterRead")
	defer tflog.Debug(ctx, "exiting resourceParameterRead")
	c := meta.(*cloudTruthClient)
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(err)
	}
	paramName := d.Get("name").(string)

	var resp *cloudtruthapi.PaginatedParameterList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		resp, r, err = c.openAPIClient.ProjectsApi.ProjectsParametersList(ctx, *projID).Name(paramName).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceParameterRead: error looking up parameter %s", paramName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	if resp.GetCount() != 1 {
		return diag.FromErr(fmt.Errorf("resourceParameterRead: expected 1 value for parameter %s, found %d instead",
			paramName, resp.GetCount()))
	}
	d.SetId(resp.GetResults()[0].GetId())
	return nil
}

func updateParameterRules(ctx context.Context, paramID, projID string, paramTypeName string, d *schema.ResourceData,
	c *cloudTruthClient) (*http.Response, error) {
	tflog.Debug(ctx, "entering updateParameterRules")
	defer tflog.Debug(ctx, "exiting updateParameterRules")
	paramType := c.lookupType(ctx, paramTypeName)
	baseParamType := getBaseParamType(ctx, paramType, c)
	rules := intAndStringRuleTypes
	if baseParamType == "string" {
		rules = append(rules, stringRuletypes...)
	}
	var updatedRules []string
	for _, rule := range rules {
		if d.HasChange(rule) {
			updatedRules = append(updatedRules, rule)
		}
	}

	// We return the last response (and stop if we receive an error) however, we could consider returning all of the responses
	var r *http.Response
	var err error
	if len(updatedRules) > 0 {
		for _, updatedRule := range updatedRules {
			ruleIDProperty := fmt.Sprintf("%s_id", updatedRule)
			ruleID := d.Get(ruleIDProperty).(string)
			newRuleVal := d.Get(updatedRule).(string)
			if newRuleVal == "" { // A value of "" == time for deletion
				r, err = deleteRule(ctx, paramID, projID, updatedRule, ruleID, c)
				if err == nil { // Don't set the property to "" unless the delete succeeds
					err = d.Set(ruleIDProperty, "")
				}
			} else {
				r, err = updateRule(ctx, paramID, projID, updatedRule, ruleID, d, c)
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return r, nil
}

func updateParameter(ctx context.Context, paramID, projID string, paramTypeName string, d *schema.ResourceData,
	c *cloudTruthClient) (*http.Response, error) {
	tflog.Debug(ctx, "entering updateParameter")
	defer tflog.Debug(ctx, "exiting updateParameter")
	patchedParam := cloudtruthapi.PatchedParameter{}
	hasParamChange := false
	if d.HasChange("description") {
		paramDesc := d.Get("description").(string)
		patchedParam.SetDescription(paramDesc)
		hasParamChange = true
	}
	if d.HasChange("secret") {
		paramIsSecret := d.Get("secret").(bool)
		patchedParam.SetSecret(paramIsSecret)
		hasParamChange = true
	}
	var r *http.Response
	var err error
	if hasParamChange {
		// The caller handles retries
		_, r, err = c.openAPIClient.ProjectsApi.ProjectsParametersPartialUpdate(ctx, paramID, projID).
			PatchedParameter(patchedParam).Execute()
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

func updateRule(ctx context.Context, paramID, projID, ruleName, ruleID string, d *schema.ResourceData, c *cloudTruthClient) (*http.Response, error) {
	tflog.Debug(ctx, "entering updateRule")
	defer tflog.Debug(ctx, "exiting updateRule")
	_, newVal := d.GetChange(ruleName)
	ruleVal := newVal.(string)
	retryCount := 0
	var apiError error
	var ruleUpdateRequest cloudtruthapi.ApiProjectsParametersRulesUpdateRequest
	ruleType, err := cloudtruthapi.NewParameterRuleTypeEnumFromValue(ruleName)
	if err != nil {
		return nil, err
	}
	paramRule := cloudtruthapi.NewParameterRuleWithDefaults()
	paramRule.SetId(ruleID)
	paramRule.SetConstraint(ruleVal)
	paramRule.SetType(*ruleType)
	paramRule.SetParameter(paramID)
	paramRule.SetType(*ruleType)

	var r *http.Response
	// todo: (non)retryable errors here?
	for retryCount < ruleOperationRetries {
		ruleUpdateRequest = c.openAPIClient.ProjectsApi.ProjectsParametersRulesUpdate(ctx, ruleID, paramID, projID).ParameterRule(*paramRule)
		_, r, err = ruleUpdateRequest.Execute()
		if r.StatusCode >= 500 {
			tflog.Debug(ctx, fmt.Sprintf("updateRule: error updating rule %s: %+v", ruleName, err))
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

func deleteRule(ctx context.Context, paramID, projID, ruleName, ruleID string, c *cloudTruthClient) (*http.Response, error) {
	tflog.Debug(ctx, "entering deleteRule")
	defer tflog.Debug(ctx, "exiting updateRule")
	retryCount := 0
	var ruleDestroyRequest cloudtruthapi.ApiProjectsParametersRulesDestroyRequest
	var r *http.Response
	var apiError, err error

	// todo: (non)retryable errors here?
	for retryCount < ruleOperationRetries {
		ruleDestroyRequest = c.openAPIClient.ProjectsApi.ProjectsParametersRulesDestroy(ctx, ruleID, paramID, projID)
		r, err = ruleDestroyRequest.Execute()
		if r.StatusCode >= 500 {
			tflog.Debug(ctx, fmt.Sprintf("deleteRule: error deleting rule %s: %+v", ruleName, err))
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

func resourceParameterUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceParameterUpdate")
	defer tflog.Debug(ctx, "exiting resourceParameterUpdate")
	c := meta.(*cloudTruthClient)
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterUpdate: %w", err))
	}
	paramName, paramType := d.Get("name").(string), d.Get("type").(string)

	// Top level property changes except for rules
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		r, err := updateParameter(ctx, d.Id(), *projID, paramType, d, c)
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceParameterUpdate: error updating parameter level config for parameter %s", paramName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	// Rule changes
	retryError = resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		r, err := updateParameterRules(ctx, d.Id(), *projID, paramType, d, c)
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceParameterUpdate: error updating rules for parameter %s", paramName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return resourceParameterRead(ctx, d, meta)
}

func resourceParameterDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceParameterDelete")
	defer tflog.Debug(ctx, "exiting resourceParameterDelete")
	c := meta.(*cloudTruthClient)
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterDelete: %w", err))
	}
	paramName := d.Get("name").(string)

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.ProjectsApi.ProjectsParametersDestroy(ctx, d.Id(), *projID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceParameterDelete: error deleting parameter %s", paramName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return nil
}
