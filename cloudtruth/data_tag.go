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
				Computed:    true,
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

	var resp *cloudtruthapi.PaginatedTagList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		resp, r, err = c.openAPIClient.EnvironmentsApi.EnvironmentsTagsList(ctx, *envID).Name(name).Execute()
		if err != nil {
			outErr := fmt.Errorf("dataCloudTruthTagRead: error looking up tag %s: %w", name, err)
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
		return diag.FromErr(fmt.Errorf("dataCloudTruthTagRead: expected 1 value for tag %s, found %d instead",
			name, resp.GetCount()))
	}
	// There should only ever be one tag per environment
	results := resp.GetResults()
	tag := results[0]
	ts := tag.Timestamp.Format(time.RFC3339)
	tagID := tag.GetId()
	err = d.Set("timestamp", ts)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTagRead: %w", err))
	}
	d.SetId(tagID)
	return nil
}
