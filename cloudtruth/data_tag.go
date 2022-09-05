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
	"time"
)

func dataCloudTruthTag() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth tag data source",
		ReadContext: dataCloudTruthTagRead,
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
				Description: "The name of the tag",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The tag's description",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"timestamp": {
				Description: "The tag's RFC 3339 timestamp ",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataCloudTruthTagRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "dataCloudTruthTagRead")
	environment := d.Get("environment").(string)
	envID, err := c.lookupEnvironment(ctx, environment)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTagRead: %w", err))
	}
	name := d.Get("name").(string)
	tagID := d.Id()

	var tag *cloudtruthapi.Tag
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		tag, r, err = c.openAPIClient.EnvironmentsApi.EnvironmentsTagsRetrieve(ctx, *envID, tagID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("dataCloudTruthTagRead: error reading tag %s", name), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	ts := tag.Timestamp.Format(time.RFC3339)
	err = d.Set("timestamp", ts)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("name", name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", d.Get("description").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(tagID)
	return nil
}
