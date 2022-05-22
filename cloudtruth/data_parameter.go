package cloudtruth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

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
				Description: "The name of the parameter",
				Type:        schema.TypeString,
				Required:    true,
			},
			"value": &schema.Schema{
				Description: "The parameter's value",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataCloudTruthParameterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "dataCloudTruthParameterRead")
	// todo: add support for provider level env/project lookup
	env := url.QueryEscape(d.Get("env").(string))
	project := url.QueryEscape(d.Get("project").(string))
	parameterName := url.QueryEscape(d.Get("name").(string))

}
