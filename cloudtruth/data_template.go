package cloudtruth

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataCloudTruthTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth template data source",
		ReadContext: dataCloudTruthTemplateRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Description: "The CloudTruth project",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environment": {
				Description: "The CloudTruth environment",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
			},
			"name": {
				Description: "The name of the template",
				Type:        schema.TypeString,
				Required:    true,
			},
			"value": {
				Description: "The template's value",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataCloudTruthTemplateRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}
