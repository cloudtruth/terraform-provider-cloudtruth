package cloudtruth

/*
import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

func resourceTag() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "A CloudTruth Tag and environment specific value (defaulting to the 'default' environment)",

		CreateContext: resourceTagCreate,
		ReadContext:   resourceTagRead,
		UpdateContext: resourceTagUpdate,
		DeleteContext: resourceTagDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the CloudTruth Tag, unique per project",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the CloudTruth Tag",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project": {
				Description: "The CloudTruth project where the Tag is defined",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"environment": {
				Description: "The CloudTruth environment where the Tag's value is defined. Defaults to the 'default' environment",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
			},
			"secret": {
				Description: "Whether or not the Tag is a secret, defaults to false (non-secret)",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"value": {
				Description: "The value of the CloudTruth Tag, specific to an environment (can be overridden/inherited relative to other environments)",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				// With external tags, the value is computed but should not be used to diff/detect changes
				// It is only useful in read operations. A change to the location and/or filter field would indicate a resource change
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					external := d.Get("external").(bool)
					return external
				},
			},
			"dynamic": {
				Description: "Whether or not to evaluate/interpolate the Tag's value (incompatible with secret tags)",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"external": {
				Description: "Whether or not the value is external, defaults to false",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"location": {
				Description: "The location of the secret value, required for external tags",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"filter": {
				Description: "An optional filter (path/query), optional and used only with external tags",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
		},
	}
}

// In Terraform terms, a 'resource_tag' represents a CloudTruth Tag AND its environment specific value
// Therefore, when this provider destroys a tag resource, it only removes the per-environment value unless
// the tag is only defined in one environment, in which case it is destroyed entirely
func resourceTagCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceTagCreate")
	environment := d.Get("environment").(string)
	project := d.Get("project").(string)
	envID, projID, err := c.lookupEnvProj(ctx, environment, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTagCreate: %w", err))
	}

	tagName := d.Get("name").(string)
	var tagID, valueID string
	lookupResp, r, err := c.openAPIClient.ProjectsApi.ProjectsTagsList(context.Background(),
		*projID).Environment(*envID).Name(tagName).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTagCreate: error looking up tag %s: %+v", tagName, r))
	}

	if lookupResp.GetCount() == 1 { // Named tag already exists
		tagID = lookupResp.GetResults()[0].GetId()
	} else {
		tagCreate := paramCreateConfig(d)
		tagCreateResp, _, err := c.openAPIClient.ProjectsApi.ProjectsTagsCreate(context.Background(),
			*projID).TagCreate(*tagCreate).Execute()
		if err != nil {
			return diag.FromErr(fmt.Errorf("resourceTagCreate: %w", err))
		}
		tagID = paramCreateResp.GetId()
	}
	valueCreate, err := valueCreateConfig(*envID, d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTagCreate: %w", err))
	}
	valueResp, _, err := c.openAPIClient.ProjectsApi.ProjectsTagsValuesCreate(context.Background(),
		tagID, *projID).ValueCreate(*valueCreate).Execute()
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTagCreate: %w", err))
	}
	valueID = valueResp.GetId()

	// Then add its value in the specified environment with composite ID - <PARAMETER_ID>:<PARAMETER_VALUE_ID>
	internalID := fmt.Sprintf("%s:%s", tagID, valueID)
	d.SetId(internalID)
	return nil
}

func tagCreateConfig(d *schema.ResourceData) *cloudtruthapi.TagCreate {
	tagName := d.Get("name").(string)
	tagCreate := cloudtruthapi.NewTagCreate(paramName)
	tagDesc := d.Get("description").(string)
	if tagDesc != "" {
		tagCreate.SetDescription(paramDesc)
	}
	isSecret := d.Get("secret").(bool)
	tagCreate.SetSecret(isSecret)
	return tagCreate
}

func valueCreateConfig(envID string, d *schema.ResourceData) (*cloudtruthapi.ValueCreate, error) {
	name := d.Get("name").(string)
	value := d.Get("value").(string)
	dynamic := d.Get("dynamic").(bool)
	valueCreate := cloudtruthapi.NewValueCreate(value)
}

func resourceTagRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTagRead")
	return dataCloudTruthTagRead(ctx, d, meta)
}

func resourceTagUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceTagUpdate")
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTagUpdate: %w", err))
	}
	return resourceTagRead(ctx, d, meta)
}

// If a tag is defined in the target environment only, we destroy it
// If it is defined in one or more other environments, we only destroy the value
// defined in the target environment
func resourceTagDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceTagDelete")
	c := meta.(*cloudTruthClient)
	environment := d.Get("environment").(string)
	project := d.Get("project").(string)
	projID, err := c.lookupProject(ctx, project)
	if err != nil {
		return diag.FromErr(fmt.Errorf("resourceTagDelete: %w", err))
	}

	return nil
}
*/
