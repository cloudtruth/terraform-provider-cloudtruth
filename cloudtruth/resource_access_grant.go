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
	"strings"
)

func resourceAccessGrant() *schema.Resource {
	return &schema.Resource{
		Description: `A grant which assigns an access role to a principal (a user or group) for a specific environment
or project. In order to use access grants, you must have a CloudTruth Business Subscription. To assign viewer,
contributor or admin level access grants to users and groups, you must create an owner grant specifying at least one
interactive user as the grant owner first. After that you can create additional grants with non-owner roles to other
users and groups.

Your provider API key must have organization OWNER or ADMIN access to create, update and delete access grants.
`,

		CreateContext: resourceAccessGrantCreate,
		ReadContext:   resourceAccessGrantRead,
		UpdateContext: resourceAccessGrantUpdate,
		DeleteContext: resourceAccessGrantDelete,

		Schema: map[string]*schema.Schema{
			"role": {
				Description: `The grant role: one of OWNER, ADMIN, CONTRIB, VIEWER. You must create an owner access
grant to the target resource before creating any non-owner grants. If you are creating owner and non-owner grants at
the same time, use the 'depends_on' directive (with the non-owner grant depending on the owner grant) to ensure that the
non-owner grant(s) the owner grant is created before the non-owner grant(s) is created.`,
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": {
				Description: "The target environment to which the role will provide access, mutually exclusive with 'project'",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project": {
				Description: "The target project to which the role will provide access, mutually exclusive with 'environment'",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"user": {
				Description: "The display name or email of the user who will be granted the role, mutually exclusive with 'group'.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"group": {
				Description: "The group which will be granted the role, mutually exclusive with 'user'",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"principal_id": {
				Description: "The internal ID of the user or group granted the role",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func parsePrincipalInput(ctx context.Context, d *schema.ResourceData, meta any) (*string, error) {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "parsePrincipalInput")
	var principalName, principalURL string
	if _, ok := d.GetOk("user"); ok {
		principalName = d.Get("user").(string)
		user, err := c.lookupUser(ctx, principalName)
		if err != nil {
			return nil, err
		}
		principalURL = fmt.Sprintf("https://api.cloudtruth.io/api/v1/users/%s/", user.GetId())
	}
	if _, ok := d.GetOk("group"); ok {
		if principalName != "" {
			return nil, fmt.Errorf("parsePrincipalInput: cannot specity both users and groups with an access grant")
		}
		principalName = d.Get("group").(string)
		group, err := c.lookupGroup(ctx, principalName)
		if err != nil {
			return nil, err
		}
		principalURL = fmt.Sprintf("https://api.cloudtruth.io/api/v1/groups/%s/", group.GetId())
	}
	return &principalURL, nil
}

func parseScopeInput(ctx context.Context, d *schema.ResourceData, meta any) (*string, error) {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "parseScopeInput")
	var scopeName, scopeURL string
	if _, ok := d.GetOk("environment"); ok {
		scopeName = d.Get("environment").(string)
		envID, err := c.lookupEnvironment(ctx, scopeName)
		if err != nil {
			return nil, err
		}
		scopeURL = fmt.Sprintf("https://api.cloudtruth.io/api/v1/environments/%s/", *envID)
	}
	if _, ok := d.GetOk("project"); ok {
		if scopeName != "" {
			return nil, fmt.Errorf("resourceAccessGrantCreate: cannot specity both project and environment names with an access grant")
		}
		scopeName = d.Get("project").(string)
		projectID, err := c.lookupProject(ctx, scopeName)
		if err != nil {
			return nil, err
		}
		scopeURL = fmt.Sprintf("https://api.cloudtruth.io/api/v1/projects/%s/", *projectID)
	}
	return &scopeURL, nil
}

func resourceAccessGrantCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceAccessGrantCreate")
	principalURL, err := parsePrincipalInput(ctx, d, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	scopeURL, err := parseScopeInput(ctx, d, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	role, err := cloudtruthapi.NewRoleEnumFromValue(d.Get("role").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	grantCreate := cloudtruthapi.NewGrantWithDefaults()
	grantCreate.SetPrincipal(*principalURL)
	grantCreate.SetScope(*scopeURL)
	grantCreate.SetRole(*role)

	var grant *cloudtruthapi.Grant
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		grant, r, err = c.openAPIClient.GrantsApi.GrantsCreate(ctx).Grant(*grantCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAccessGrantCreate: error creating grant for principal %s and scope %s", *principalURL, *scopeURL), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	err = d.Set("role", grant.GetRole())
	if err != nil {
		return diag.FromErr(err)
	}
	principalID, err := getPrincipalID(ctx, grant.GetPrincipal())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("principal_id", principalID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(grant.GetId())
	return nil
}

func resourceAccessGrantRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceAccessGrantRead")
	c := meta.(*cloudTruthClient)
	grantID := d.Id()

	var grant *cloudtruthapi.Grant
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		var err error
		grant, r, err = c.openAPIClient.GrantsApi.GrantsRetrieve(ctx, grantID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAccessGrantRead: error reading grant %s", grantID), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	err := d.Set("role", grant.GetRole())
	if err != nil {
		return diag.FromErr(err)
	}
	principalID, err := getPrincipalID(ctx, grant.GetPrincipal())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("principal_id", principalID)
	if err != nil {
		return diag.FromErr(err)
	}
	// intentionally omitting environment and project, they are not settable in the UI
	d.SetId(grant.GetId())
	return nil
}

func resourceAccessGrantUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceAccessGrantUpdate")
	c := meta.(*cloudTruthClient)
	patchedGrant := cloudtruthapi.PatchedGrant{}
	hasChange := false
	if d.HasChanges("user", "group") {
		principalURL, err := parsePrincipalInput(ctx, d, meta)
		if err != nil {
			return diag.FromErr(err)
		}
		principalID, err := getPrincipalID(ctx, *principalURL)
		if err != nil {
			return diag.FromErr(err)
		}
		curPrincipalID := d.Get("principal_id")
		if curPrincipalID != principalID {
			patchedGrant.SetPrincipal(*principalURL)
			hasChange = true
		}
	}
	if d.HasChange("role") {
		role, err := cloudtruthapi.NewRoleEnumFromValue(d.Get("role").(string))
		if err != nil {
			return diag.FromErr(err)
		}
		patchedGrant.SetRole(*role)
		hasChange = true
	}

	if hasChange {
		retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.GrantsApi.GrantsPartialUpdate(ctx, d.Id()).PatchedGrant(patchedGrant).Execute()
			if err != nil {
				return handleAPIError(fmt.Sprintf("resourceAccessGrantUpdate: error updating grant with ID %s", d.Id()), r, err)
			}
			return nil
		})
		if retryError != nil {
			return diag.FromErr(retryError)
		}
	}
	return resourceAccessGrantRead(ctx, d, meta)
}

func resourceAccessGrantDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceAccessGrantDelete")
	c := meta.(*cloudTruthClient)
	grantID := d.Id()

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.GrantsApi.GrantsDestroy(ctx, grantID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAccessGrantDelete: error destroying grant %s", d.Id()), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return nil
}

func getPrincipalID(ctx context.Context, principalURL string) (*string, error) {
	tflog.Debug(ctx, "getUserName")
	// Sample principal URLs
	// https://api.cloudtruth.io/api/v1/users/USER_ID/
	// https://api.cloudtruth.io/api/v1/groups/GROUP_ID/
	urlSegments := strings.Split(principalURL, "/")
	if len(urlSegments) < 7 {
		return nil, fmt.Errorf("did not find the principal ID in the URL %s", principalURL)
	} else {
		id := urlSegments[6]
		return &id, nil
	}
}
