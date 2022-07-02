package cloudtruth

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				Description: "The tag's value",
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
	project := d.Get("project").(string)
	envID, _, err := c.lookupEnvProj(ctx, environment, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTagRead: %w", err))
	}
	name := d.Get("name").(string)

	resp, r, err := c.openAPIClient.EnvironmentsApi.EnvironmentsTagsList(ctx, *envID).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTagRead: error looking up tag %s: %+v", name, r))
	}
	if resp.GetCount() != 1 {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTagRead: expected 1 value for tag %s, found %d instead",
			name, resp.GetCount()))
	}
	// We know there is only one tag at this point
	// There should only ever be one tag per environment per project
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

/*
// Return a map of tag names -> tag values
func dataCloudTruthTags() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth Tag data source",
		ReadContext: dataCloudTruthTagsRead,
		Schema: map[string]*schema.Schema{
			"project": {
				Description: "The CloudTruth Project",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environment": {
				Description: "The CloudTruth Environment",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"as_of": {
				Description: "Filter for all tag values defined 'as of' the specified RFC3333 date, mutually exclusive with 'tag'",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tag": {
				Description: "Filter for tag values matching a specific tag, mutually exclusive with 'as_of'",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tag_values": {
				Description: "The computed values for all CloudTruth Tags matching the filter(s) and environment",
				Type:        schema.TypeMap,
				Computed:    true,
			},
		},
	}
}

func dataCloudTruthTagsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "dataCloudTruthTagsRead")
	environment := d.Get("environment").(string)
	project := d.Get("project").(string)
	envID, projID, err := c.lookupEnvProj(ctx, environment, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTagsRead: %w", err))
	}

	// Handle as_of and tag filters
	tagListRequest := c.openAPIClient.ProjectsApi.ProjectsTagsList(context.Background(),
		*projID).Environment(environment)
	asOf := d.Get("as_of").(string)
	tag := d.Get("tag").(string)
	if asOf != "" {
		if tag != "" {
			return diag.Errorf("dataCloudTruthTagsRead: 'as_of' and 'tag' cannot both be specified as tag filters")
		}
		asOfTime, err := datetime.Parse(asOf, time.UTC)
		if err != nil {
			return diag.FromErr(fmt.Errorf("dataCloudTruthTagsRead: %w", err))
		}
		tagListRequest = paramListRequest.AsOf(asOfTime)
	}
	if tag != "" {
		tagListRequest = paramListRequest.Tag(tag)
	}

	resp, r, err := tagListRequest.Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTagsRead: error looking up tags in the %s environment in the %s project: %+v",
			environment, project, r))
	}

	results := resp.GetResults()
	tagsMap := make(map[string]any)
	var pageNum int32 = 1
	for results != nil {
		for _, res := range results {
			tagName := res.GetName()
			// There should only be one value but the result is a map, so we iterate over it
			for _, v := range res.GetValues() {
				tagEnvValue := v.GetValue()
				// We need to exclude tags that do not have values set in the target environment
				if tagEnvValue != "" {
					tagsMap[paramName] = paramEnvValue
				}
			}
		}
		if resp.GetNext() != "" {
			pageNum++
			tagListRequest = c.openAPIClient.ProjectsApi.ProjectsTagsList(context.Background(),
				*projID).Environment(environment).Page(pageNum)
			resp, r, err = tagListRequest.Execute()
			if err != nil {
				return diag.FromErr(fmt.Errorf("dataCloudTruthTagsRead: error looking up tags in the %s environment in the %s project: %+v",
					environment, project, r))
			}
			results = resp.GetResults()
		} else {
			break
		}
	}

	err = d.Set("tag_values", tagsMap)
	if err != nil {
		return diag.FromErr(fmt.Errorf("dataCloudTruthTagsRead: %w", err))
	}

	// We use a composite ID - <PROJECT_ID>:<ENVIRONMENT_ID>
	// The results represent a set of possibly filtered tags in a given environment and project
	// Ideally, these will be shared across all TF resources in a given state
	// However, we do need to be mindful that a user could pull in the set of tags more than once
	// within the same state, with possible variations on filters, if that turns out to be an issue
	// then we may need to rethink the ID strategy here
	d.SetId(fmt.Sprintf("%s:%s", *projID, *envID))
	return nil
}
*/
