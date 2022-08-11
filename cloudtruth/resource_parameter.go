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

var ruleTypes = [5]string{"max", "min", "max_len", "min_len", "regex"}

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
			"max": {
				Description: "A CloudTruth rule constraint: the maximum value for an integer type, only valid with integer types",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"max_id": {
				Description: "The internal ID of the max rule",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"min": {
				Description: "A CloudTruth rule constraint: the minimum value for an integer type, only valid with integer types",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"min_id": {
				Description: "The internal ID of the min rule",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"max_len": {
				Description: "A CloudTruth rule constraint: the maximum length for a string type, only valid with string types",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"max_len_id": {
				Description: "The internal ID of the max_len rule",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"min_len": {
				Description: "A CloudTruth rule constraint: the minimum length for a string type, only valid with string types",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"min_len_id": {
				Description: "The internal ID of the min_len rule",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"regex": {
				Description: "A CloudTruth rule constraint: the regular expression a string type must match, only valid with string types",
				Type:        schema.TypeString,
				Optional:    true,
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
	tflog.Debug(ctx, "resourceParameterCreate")
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
		paramCreate := paramCreateConfig(d)
		paramCreateResp, r, err = c.openAPIClient.ProjectsApi.ProjectsParametersCreate(ctx,
			*projID).ParameterCreate(*paramCreate).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceParameterCreate: error creating parameter %s: %w", paramName, err)
			if r.StatusCode >= http.StatusInternalServerError {
				return resource.RetryableError(outErr)
			} else {
				return resource.NonRetryableError(outErr)
			}
		}
		paramID = paramCreateResp.GetId()
		return nil
	})

	if retryError != nil {
		return diag.FromErr(retryError)
	}
	for k, v := range rules {
		ruleID, err := addRuleToParam(ctx, c, paramID, *projID, k, v)
		if err != nil {
			return diag.FromErr(err)
		}
		// Compute the rule ID
		err = d.Set(fmt.Sprintf("%s_id", k), ruleID)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId(paramID)
	return nil
}

// boolean base types cannot have rules
// integer base types can have up to two rules: min and max
// string base types can have up to three rules: min_len, max_len and regex
func validateAndFetchRules(ctx context.Context, c *cloudTruthClient, d *schema.ResourceData,
	paramType *cloudtruthapi.ParameterType) (map[string]any, error) {
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
	outRules := map[string]any{}
	intRules := []string{"max", "min"}
	for _, r := range intRules {
		if val, ok := d.GetOk(r); ok {
			if baseParamTypeName != "integer" {
				return nil, fmt.Errorf("the base type %s does not support the %s rule type", baseParamTypeName, r)
			}
			outRules[r] = val
		}
	}
	stringRules := []string{"max_len", "min_len", "regex"}
	for _, r := range stringRules {
		if val, ok := d.GetOk(r); ok {
			if baseParamTypeName != "string" {
				return nil, fmt.Errorf("the base type %s does not support the %s rule type", baseParamTypeName, r)
			} else {
				outRules[r] = val
			}
		}
	}
	return outRules, nil
}

func addRuleToParam(ctx context.Context, c *cloudTruthClient, paramID, projectID string, ruleName string, ruleVal any) (*string, error) {
	tflog.Debug(ctx, "addRuleToType")
	retryCount := 0
	var apiError error
	var createTypeRule cloudtruthapi.ParameterRuleCreate
	typeEnum, err := cloudtruthapi.NewParameterRuleTypeEnumFromValue(ruleName)
	if err != nil {
		return nil, err
	}
	createTypeRule.SetType(*typeEnum)
	if ruleName == "regex" {
		createTypeRule.SetConstraint(ruleVal.(string))
	} else {
		createTypeRule.SetConstraint(fmt.Sprintf("%d", ruleVal))
	}

	var paramRule *cloudtruthapi.ParameterRule
	var r *http.Response
	for retryCount < addRuleRetries {
		paramRule, r, err = c.openAPIClient.ProjectsApi.ProjectsParametersRulesCreate(ctx, paramID, projectID).ParameterRuleCreate(createTypeRule).Execute()
		if r.StatusCode >= 500 {
			tflog.Debug(ctx, fmt.Sprintf("addRuleToType: %s", err))
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
	ruleID := paramRule.GetId()
	return &ruleID, nil
}

func paramCreateConfig(d *schema.ResourceData) *cloudtruthapi.ParameterCreate {
	paramName := d.Get("name").(string)
	paramCreate := cloudtruthapi.NewParameterCreate(paramName)
	paramDesc := d.Get("description").(string)
	if paramDesc != "" {
		paramCreate.SetDescription(paramDesc)
	}
	isSecret := d.Get("secret").(bool)
	paramCreate.SetSecret(isSecret)
	return paramCreate
}

func resourceParameterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceParameterRead")
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
			outErr := fmt.Errorf("resourceParameterRead: error looking up parameter %s: %w", paramName, err)
			if r.StatusCode >= http.StatusInternalServerError {
				return resource.RetryableError(outErr)
			} else {
				return resource.NonRetryableError(outErr)
			}
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

func updateParameter(ctx context.Context, paramID, projID string, d *schema.ResourceData, c *cloudTruthClient) (*http.Response, error) {
	patchedParam := cloudtruthapi.PatchedParameter{}
	hasParamChange := false
	var updatedRules []string
	for _, rule := range ruleTypes {
		if d.HasChange(rule) {
			updatedRules = append(updatedRules, rule)
		}
	}
	if len(updatedRules) > 0 {
		for _, updatedRule := range updatedRules {
			ruleID := d.Get(fmt.Sprintf("%s_id", updatedRule)).(string)
			//todo: return this?
			_, err := updateRule(ctx, paramID, projID, updatedRule, ruleID, d, c)
			if err != nil {
				return nil, err
			}
		}
	}

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
		_, r, err = c.openAPIClient.ProjectsApi.ProjectsParametersPartialUpdate(ctx, paramID, projID).
			PatchedParameter(patchedParam).Execute()
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// todo:
// handle the update constraint case first
// then the delete constraint case
func updateRule(ctx context.Context, paramID, projID, ruleName, ruleID string, d *schema.ResourceData, c *cloudTruthClient) (*http.Response, error) {
	_, newVal := d.GetChange(ruleName)
	// regex rule values are strings, all other integers
	var ruleVal string
	if ruleName == "regex" {
		ruleVal = newVal.(string)
	} else {
		ruleVal = fmt.Sprintf("%d", newVal.(int))
	}
	retryCount := 0
	var apiError error
	var ruleUpdateRequest cloudtruthapi.ApiProjectsParametersRulesUpdateRequest

	// NewParameterRule(url string, id string, parameter string, type_ ParameterRuleTypeEnum, constraint string, createdAt time.Time, modifiedAt time.Time)
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
	for retryCount < addRuleRetries {
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

func resourceParameterUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceParameterUpdate")
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceParameterUpdate: %w", err))
	}
	paramName := d.Get("name").(string)

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		r, err := updateParameter(ctx, d.Id(), *projID, d, c)
		if err != nil {
			outErr := fmt.Errorf("resourceParameterUpdate: error updating parameter level config for parameter %s: %w", paramName, err)
			if r.StatusCode >= http.StatusInternalServerError {
				return resource.RetryableError(outErr)
			} else {
				return resource.NonRetryableError(outErr)
			}
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return resourceParameterRead(ctx, d, meta)
}

func resourceParameterDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceParameterDelete")
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
			outErr := fmt.Errorf("resourceParameterDelete: error deleting parameter %s: %w", paramName, err)
			if r.StatusCode >= http.StatusInternalServerError {
				return resource.RetryableError(outErr)
			} else {
				return resource.NonRetryableError(outErr)
			}
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return nil
}
