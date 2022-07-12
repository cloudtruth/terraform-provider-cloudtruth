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
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth User",

		CreateContext: resourceGroupCreate,
		ReadContext:   resourceGroupRead,
		UpdateContext: resourceGroupUpdate,
		DeleteContext: resourceGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The user's email address, used for the initial invite",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The user's access role in the target CloudTruth organization",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"users": {
				Description: "The user's access role in the target CloudTruth organization",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
		},
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceGroupCreate")
	groupName := d.Get("name").(string)
	desc := d.Get("description").(string)

	userURIs := []string{}
	groupCreate := cloudtruthapi.NewGroupWithDefaults()
	groupCreate.SetName(groupName)
	if desc != "" {
		groupCreate.SetDescription(desc)
	}
	if _, ok := d.GetOk("users"); ok {
		users := d.Get("users").([]interface{})
		userURIs = make([]string, len(users))
		/* todo: need a user cache here, we can't rely on user names, we need to lookup IDs
		for i, v := range users {
			userName := fmt.Sprint(v)
			userURIs[i], err = fetchUserURI(userName)
		} */
	}
	groupCreate.SetUsers(userURIs)

	var group *cloudtruthapi.Group
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		group, r, err = c.openAPIClient.GroupsApi.GroupsCreate(ctx).Group(*groupCreate).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceGroupCreate: error creating group %s: %w", groupName, err)
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

	d.SetId(group.Id)
	return nil
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceGroupRead")
	c := meta.(*cloudTruthClient)
	groupName := d.Get("name").(string)

	var resp *cloudtruthapi.PaginatedGroupList
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *resource.RetryError {
		var r *http.Response
		var err error
		resp, r, err = c.openAPIClient.GroupsApi.GroupsList(ctx).Name(groupName).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceGroupRead: error reading group %s: %w", groupName, err)
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

	res := resp.GetResults()
	if len(res) != 1 {
		return diag.FromErr(fmt.Errorf("resourceGroupRead: found %d groups, expcted to find 1", len(res)))
	}
	d.SetId(resp.GetResults()[0].GetId())
	return nil
}

func resourceGroupUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceGroupUpdate")
	c := meta.(*cloudTruthClient)
	groupName := d.Get("name")
	groupID := d.Id()
	patchedGroup := cloudtruthapi.PatchedGroup{}
	hasChange := false
	if d.HasChange("users") {
		patchedGroup.SetUsers(d.Get("users").([]string))
		hasChange = true
	}
	if d.HasChange("description") {
		patchedGroup.SetDescription(d.Get("description").(string))
		hasChange = true
	}
	if hasChange {
		retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.GroupsApi.GroupsPartialUpdate(ctx, groupID).PatchedGroup(patchedGroup).Execute()
			if err != nil {
				outErr := fmt.Errorf("resourceGroupUpdate: error updating group %s: %w", groupName, err)
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
	}
	return resourceGroupRead(ctx, d, meta)
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceGroupDelete")
	c := meta.(*cloudTruthClient)
	groupName := d.Get("name").(string)
	groupID := d.Id()

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.GroupsApi.GroupsDestroy(ctx, groupID).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceGroupDelete: error destroying group %s: %w", groupName, err)
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
	return nil
}
