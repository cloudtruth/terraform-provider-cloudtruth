package cloudtruth

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataCloudTruthParameter() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth parameter data source",
		ReadContext: dataCloudTruthParameterRead,
		Schema: map[string]*schema.Schema{
			"env": {
				Description: "The CloudTruth environment",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project": {
				Description: "The CloudTruth project",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The name of the secret",
				Type:        schema.TypeString,
				Required:    true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

//<BaseURL>/projects/<project>/parameters/?evaluate=true&mask_secrets=false&name=<name>&wrap=false&environment=<environment>'
func dataCloudTruthParameterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "dataCloudTruthParameterRead")
	// todo: build URL here, possibly add builder type
	resp, err := c.Get(c.config.BaseURL)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("%v+", resp))
	env := url.QueryEscape(d.Get("enviroment").(string))
	name := url.QueryEscape(d.Get("name").(string))
	project := url.QueryEscape(d.Get("project").(string))
	tflog.Debug(ctx, env+name+project)
	d.SetId("temporary id")
	return nil
}
