package cloudtruth

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var ruleSchema = map[string]*schema.Schema{
	"type": {
		Description: "The type of rule, for strings: max_len|min_len|regex. For integers: max|min. Booleans do not support rules",
		Type:        schema.TypeString,
		Required:    true,
	},
	"constraint": {
		Description: "The limiting value of the rule as a string e.g. '1' for max_len|min_len|max|min rule types or '.*' for a regex rule type",
		Type:        schema.TypeString,
		Required:    true,
	},
}
