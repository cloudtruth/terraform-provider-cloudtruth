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

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Description: "A CloudTruth User",

		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"email": {
				Description: "The user's email address, used for the initial invite",
				Type:        schema.TypeString,
				Required:    true,
			},
			"role": {
				Description: "The user's access role in the target CloudTruth organization",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	c := meta.(*cloudTruthClient)
	tflog.Debug(ctx, "resourceUserCreate")
	email := d.Get("email").(string)
	role := d.Get("role").(string)

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		roleEnum, _ := cloudtruthapi.NewRoleEnumFromValue(role)
		nullableRollEnum := cloudtruthapi.NewNullableRoleEnum(roleEnum)
		invitationCreate := cloudtruthapi.NewInvitationCreate(email, *nullableRollEnum)
		_, r, err = c.openAPIClient.InvitationsApi.InvitationsCreate(context.Background()).InvitationCreate(*invitationCreate).Execute()
		if err != nil {
			outErr := fmt.Errorf("resourceUserCreate: error inviting user %s: %w", email, err)
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

	d.SetId(email)
	return nil
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceUserRead")

	return nil
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceUserUpdate")
	return nil
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "resourceUserDelete")

	return nil
}
