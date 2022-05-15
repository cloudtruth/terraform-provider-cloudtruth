package cloudtruth

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceParameter() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth parameter data source",
		ReadContext: dataSourceParameterRead,
		Schema: map[string]*schema.Schema{
			"env": {
				Description: "The CloudTruth environment",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceParameterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	idFromAPI := "my-id"
	d.SetId(idFromAPI)
	return diag.Errorf("not implemented")
}
