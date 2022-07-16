package cloudtruth

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataCloudTruthUser() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth user data source",
		ReadContext: dataCloudTruthUserRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The CloudTruth user's name, one of 'name' or 'email' is required",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"email": {
				Description: "The user's email, one of 'name' or 'email' is required\"",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"role": {
				Description: "The user's role",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The user's type: interactive or service",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

// look up users by email or name
func dataCloudTruthUserRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "dataCloudTruthUserRead")
	c := meta.(*cloudTruthClient)
	nameOrEmail := d.Get("name")
	if nameOrEmail == "" {
		nameOrEmail = d.Get("email")
	}
	user, err := c.lookupUser(ctx, nameOrEmail.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(user.GetId())
	err = d.Set("role", user.GetRole())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("type", user.GetType())
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
