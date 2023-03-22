package cloudtruth

import (
	"context"
	"fmt"
	"github.com/cloudtruth/terraform-provider-cloudtruth/pkg/cloudtruthapi"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Description: `A CloudTruth Group for managing access to a set of users.

Your provider API key must have organization OWNER or ADMIN access to create, update and delete groups.
`,

		CreateContext: resourceGroupCreate,
		ReadContext:   resourceGroupRead,
		UpdateContext: resourceGroupUpdate,
		DeleteContext: resourceGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the group",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "A description of the group",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"users": {
				Description: "The CloudTruth users who are members of the group",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"user_uris": {
				Description: "The URIs of the CloudTruth users who are members of the group",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
		},
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceGroupCreate")
	defer tflog.Debug(ctx, "exiting resourceGroupCreate")
	c := meta.(*cloudTruthClient)
	groupName := d.Get("name").(string)
	desc := d.Get("description").(string)

	userURIs := make([]string, 0)
	groupCreate := cloudtruthapi.NewGroupWithDefaults()
	groupCreate.SetName(groupName)
	if desc != "" {
		groupCreate.SetDescription(desc)
	}
	if _, ok := d.GetOk("users"); ok {
		users := d.Get("users").(*schema.Set)
		for _, v := range users.List() {
			userName := fmt.Sprint(v)
			user, err := c.lookupUser(ctx, userName)
			userURIs = append(userURIs, user.GetUrl())
			if err != nil {
				return diag.FromErr(err)
			} else if user == nil {
				return diag.FromErr(fmt.Errorf("resourceGroupCreate: failed to find user %s", userName))
			}
		}
	}
	groupCreate.SetUsers(userURIs)

	var group *cloudtruthapi.Group
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
		var r *http.Response
		var err error
		group, r, err = c.openAPIClient.GroupsApi.GroupsCreate(ctx).Group(*groupCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceGroupCreate: error creating group %s", groupName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	err := d.Set("user_uris", userURIs)
	if err != nil {
		return diag.FromErr(err)
	}
	c.groups[groupName] = *group
	d.SetId(group.Id)
	return nil
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceGroupRead")
	defer tflog.Debug(ctx, "exiting resourceGroupRead")
	c := meta.(*cloudTruthClient)
	groupName := d.Get("name").(string)
	groupID := d.Id()

	var group *cloudtruthapi.Group
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		var r *http.Response
		var err error
		group, r, err = c.openAPIClient.GroupsApi.GroupsRetrieve(ctx, groupID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceGroupRead: error reading group %s", groupName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	err := d.Set("name", group.GetName())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", group.GetDescription())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("user_uris", group.GetUsers())
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(group.GetId())
	return nil
}

func resourceGroupUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceGroupUpdate")
	defer tflog.Debug(ctx, "exiting resourceGroupUpdate")
	c := meta.(*cloudTruthClient)
	groupName := d.Get("name")
	groupID := d.Id()
	patchedGroup := cloudtruthapi.PatchedGroup{}
	hasChange := false
	// users indicates an intentional change, user_uris alone indicates drift
	if d.HasChange("users") || d.HasChange("user_uris") {
		userURIs := make([]string, 0)
		users := d.Get("users").(*schema.Set)
		for _, v := range users.List() {
			userName := fmt.Sprint(v)
			user, err := c.lookupUser(ctx, userName)
			userURIs = append(userURIs, user.GetUrl())
			if err != nil {
				return diag.FromErr(err)
			} else if user == nil {
				return diag.FromErr(fmt.Errorf("resourceGroupUpdate: failed to find user %s", userName))
			}
		}
		patchedGroup.SetUsers(userURIs)
		hasChange = true
	}
	if d.HasChange("name") {
		patchedGroup.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("description") {
		patchedGroup.SetDescription(d.Get("description").(string))
		hasChange = true
	}
	if hasChange {
		retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *retry.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.GroupsApi.GroupsPartialUpdate(ctx, groupID).PatchedGroup(patchedGroup).Execute()
			if err != nil {
				return handleAPIError(fmt.Sprintf("resourceGroupUpdate: error updating group %s", groupName), r, err)
			}
			return nil
		})
		if retryError != nil {
			return diag.FromErr(retryError)
		}
	}
	return resourceGroupRead(ctx, d, meta)
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceGroupDelete")
	defer tflog.Debug(ctx, "exiting resourceGroupDelete")
	c := meta.(*cloudTruthClient)
	groupName := d.Get("name").(string)
	groupID := d.Id()

	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *retry.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.GroupsApi.GroupsDestroy(ctx, groupID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceGroupDelete: error destroying group %s", groupName), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	delete(c.groups, groupName)
	return nil
}
