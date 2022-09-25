package cloudtruth

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"sort"
)

var userSchema = map[string]*schema.Schema{
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
	"organization": {
		Description: "The name of the CloudTruth organization containing the user",
		Type:        schema.TypeString,
		Computed:    true,
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
}

func dataCloudTruthUser() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth user data source",
		ReadContext: dataCloudTruthUserRead,
		Schema:      userSchema,
	}
}

// look up users by email or name
func dataCloudTruthUserRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering dataCloudTruthUserRead")
	defer tflog.Debug(ctx, "exiting dataCloudTruthUserRead")
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
	err = d.Set("organization", user.GetOrganizationName())
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func dataCloudTruthUsers() *schema.Resource {
	return &schema.Resource{
		Description: "A data source representing all user accounts in a CloudTruth organization, sorted by name and optionally filtered by type (interactive/service/all)",
		ReadContext: dataCloudTruthUsersRead,
		Schema: map[string]*schema.Schema{
			"users": {
				Description: "The CloudTruth user's name, one of 'name' or 'email' is required",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: userSchema,
				},
			},
			"organization": {
				Description: "The CloudTruth organization",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The type of user account to return: interactive, service or all, default: interactive",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "interactive",
			},
		},
	}
}

func dataCloudTruthUsersRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering dataCloudTruthUsersRead")
	defer tflog.Debug(ctx, "exiting dataCloudTruthUsersRead")
	c := meta.(*cloudTruthClient)
	typeFilter := d.Get("type").(string)
	dataUsers := make(map[string]interface{})
	var org string
	for _, user := range c.users {
		if (typeFilter == *user.Type) || (typeFilter == "all") {
			// Since we cache dataUsers by name and email, use the user object's name as the new map key
			// to dedupe (double write)
			dataUser := make(map[string]interface{})
			name := user.GetName()
			dataUser["name"] = name
			if user.GetEmail() != "" {
				dataUser["email"] = user.GetEmail()
			}
			dataUser["role"] = user.GetRole()
			dataUser["type"] = user.GetType()
			dataUsers[name] = dataUser
			if org == "" {
				org = user.GetOrganizationName()
			}
		}
	}

	userList := make([]interface{}, 0, len(dataUsers))
	for _, v := range dataUsers {
		userList = append(userList, v)
	}
	// sort the list for stable plans (and tests)
	sort.Slice(userList, func(i, j int) bool {
		user1 := userList[i].(map[string]interface{})
		user2 := userList[j].(map[string]interface{})
		return user1["name"].(string) < user2["name"].(string)
	})

	err := d.Set("users", userList)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("type", typeFilter)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("organization", org)
	if err != nil {
		return diag.FromErr(err)
	}

	// Use the org name + filter as the ID of the cloudtruth_users data source
	d.SetId(fmt.Sprintf("%s:%s", org, typeFilter))
	return nil
}
