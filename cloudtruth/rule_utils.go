package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Utility functions for parameters and custom types

func getBaseParamType(ctx context.Context, paramType *cloudtruthapi.ParameterType, c *cloudTruthClient) string {
	var baseParamTypeName string
	baseParamType := paramType
	for {
		if baseParamType.Parent.Get() != nil && baseParamType.ParentName.IsSet() {
			baseParamType = c.lookupType(ctx, baseParamType.GetParentName())
		} else {
			baseParamTypeName = baseParamType.GetName()
			break
		}
	}
	return baseParamTypeName
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
