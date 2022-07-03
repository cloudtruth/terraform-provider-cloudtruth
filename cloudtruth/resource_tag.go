package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/nav-inc/datetime"
	"time"
)

func resourceTag() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "A CloudTruth Tag - unique per environment, defaulting to the 'default' environment",

		CreateContext: resourceTagCreate,
		ReadContext:   resourceTagRead,
		UpdateContext: resourceTagUpdate,
		DeleteContext: resourceTagDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the CloudTruth Tag, unique per project, per environment",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the CloudTruth Tag",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environment": {
				Description: "The CloudTruth environment where the Tag's value is defined. Defaults to the 'default' environment",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
			},
			"timestamp": {
				Description:  "The tag's RFC 3339 timestamp e.g. 2022-07-03T21:44:00.171Z",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},
		},
	}
}

func resourceTagCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceTagCreate")
	environment := d.Get("environment").(string)
	envID, err := c.lookupEnvironment(ctx, environment)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTagCreate: %w", err))
	}

	tagName := d.Get("name").(string)
	var tagID string
	tagCreate := cloudtruthapi.NewTagCreate(tagName)
	desc := d.Get("description").(string)
	if desc != "" {
		tagCreate.SetDescription(desc)
	}
	timestamp := d.Get("timestamp").(string)
	tsTime, err := datetime.Parse(timestamp, time.UTC)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTagCreate: %w", err))
	}
	tagCreate.SetTimestamp(tsTime)
	tagCreateResp, _, err := c.openAPIClient.EnvironmentsApi.EnvironmentsTagsCreate(ctx, *envID).TagCreate(*tagCreate).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTagCreate: %w", err))
	}
	tagID = tagCreateResp.GetId()
	d.SetId(tagID)
	return nil
}

func resourceTagRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTagRead")
	return dataCloudTruthTagRead(ctx, d, meta)
}

func resourceTagUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceTagUpdate")
	environment := d.Get("environment").(string)
	envID, err := c.lookupEnvironment(ctx, environment)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTagUpdate: %w", err))
	}
	hasChange := false
	patchedTagUpdate := cloudtruthapi.NewPatchedTagUpdate()
	if d.HasChange("description") {
		tagDesc := d.Get("description").(string)
		patchedTagUpdate.SetDescription(tagDesc)
		hasChange = true
	}
	if d.HasChange("timestamp") {
		timestamp := d.Get("timestamp").(string)
		tsTime, err := datetime.Parse(timestamp, time.UTC)
		if err != nil {
			return diag.FromErr(fmt.Errorf("resourceTagCreate: %w", err))
		}
		patchedTagUpdate.SetTimestamp(tsTime)
		hasChange = true
	}
	tagID := d.Id()
	if hasChange {
		_, _, err = c.openAPIClient.EnvironmentsApi.EnvironmentsTagsPartialUpdate(ctx, *envID, tagID).PatchedTagUpdate(*patchedTagUpdate).Execute()
		if err != nil {
			return diag.FromErr(fmt.Errorf("resourceTagUpdate: %w", err))
		}
	}
	d.SetId(tagID)
	return resourceTagRead(ctx, d, meta)
}

func resourceTagDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceTagDelete")
	environment := d.Get("environment").(string)
	envID, err := c.lookupEnvironment(ctx, environment)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTagDelete: %w", err))
	}

	tagID := d.Id()
	_, err = c.openAPIClient.EnvironmentsApi.EnvironmentsTagsDestroy(ctx, *envID, tagID).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTagDelete: %w", err))
	}
	d.SetId(tagID)
	return nil
}
