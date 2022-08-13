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
			"rule": {
				Description: `The rule(s) describing allowable values for a parameter of this type. Add separate blocks per rule. 
Note that string types support max_len|min_len|regex rules, integers support min|max rules and booleans don't support any rules.'`,
				Type:     schema.TypeList,
				Optional: true, // Optional and disallowed with boolean types
				Elem: &schema.Resource{
					Schema: ruleSchema,
				},
			},
			"min_rule_id": {
				Description: "The internal ID of the min/min_len rule if it exists, only available with string/integer types",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"max_rule_id": {
				Description: "The internal ID of the max/max_len rule if it exists, only available with string/integer types",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"regex_rule_id": {
				Description: "The internal ID of the regex rule if it exists, only available with string types",
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
	if err = validateRules(ctx, c, d, paramName, paramType); err != nil {
		return diag.FromErr(err)
	}

	if _, ok := d.GetOk("rule"); ok {
		rules := d.Get("rule").([]any)
		for _, r := range rules {
			ruleMap := r.(map[string]any)
			ruleID, err := addRuleToParam(ctx, c, paramID, *projID, ruleMap)
			if err != nil {
				return diag.FromErr(err)
			}
			ruleType := ruleMap["type"].(string)
			if strings.Contains(ruleType, "min") {
				err = d.Set("min_rule_id", ruleID)
			} else if strings.Contains(ruleType, "max") {
				err = d.Set("max_rule_id", ruleID)
			} else {
				err = d.Set("regex_rule_id", ruleID)
			}
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if retryError != nil {
		return diag.FromErr(retryError)
	}
	d.SetId(paramID)
	return nil
}

// boolean base types cannot have rules
// integer base types can have up to two rules: min and max
// string base types can have up to three rules: min_len, max_len and regex
func validateRules(ctx context.Context, c *cloudTruthClient, d *schema.ResourceData,
	paramName string, paramType *cloudtruthapi.ParameterType) error {
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
	rules := d.Get("rule").([]any)
	switch baseParamTypeName {
	case "boolean":
		return fmt.Errorf("the base type of the %s parameter is boolean, it does not support rules", paramName)
	case "integer":
		if len(rules) > 2 {
			return fmt.Errorf("the base type of the %s parameter is integer, it accepts no more than two rules", paramName)
		}
	case "string":
		if len(rules) > 3 {
			return fmt.Errorf("the base type of the %s parameter is string, it accepts no more than three rules", paramName)
		}
	}
	return nil
}

func addRuleToParam(ctx context.Context, c *cloudTruthClient, paramID, projectID string, rule map[string]any) (*string, error) {
	tflog.Debug(ctx, "addRuleToType")
	retryCount := 0
	var apiError error
	var createTypeRule cloudtruthapi.ParameterRuleCreate
	typeEnum, err := cloudtruthapi.NewParameterRuleTypeEnumFromValue(rule["type"].(string))
	if err != nil {
		return nil, err
	}
	createTypeRule.SetType(*typeEnum)
	createTypeRule.SetConstraint(rule["constraint"].(string))
	var paramRule *cloudtruthapi.ParameterRule
	var r *http.Response
	for retryCount < addRuleRetries {
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
	tflog.Debug(ctx, "updateParameter")
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
	/*if d.HasChange("rule") {
		// Update rules separately, do not set hasParamChange to true
		updateRules(ctx, paramID, projID, d, c)
	}*/
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

/*
func updateRules(ctx context.Context, paramID, projID string, d *schema.ResourceData, c *cloudTruthClient) (*http.Response, error) {
	tflog.Debug(ctx, "updateRules")
	rules := d.Get("rule").([]any)
	//retryCount := 0
	var apiError error
	for _, r := range rules {
		// add comprehensive diff/update here
		fmt.Sprintf("%+v", r)
	}
	old, new := d.GetChange("rule")
	var r *http.Response
	/*for retryCount < addRuleRetries {
		_, r, err := c.openAPIClient.ProjectsApi.ProjectsParametersRulesPartialUpdateExecute()
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
	return r, nil
}
*/

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
			if r != nil && r.StatusCode >= http.StatusInternalServerError {
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
