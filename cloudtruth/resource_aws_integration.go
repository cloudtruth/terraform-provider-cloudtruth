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

func resourceAWSIntegration() *schema.Resource {
	return &schema.Resource{
		Description: `A CloudTruth AWS integration.`,

		CreateContext: resourceAWSIntegrationCreate,
		ReadContext:   resourceAWSIntegrationRead,
		UpdateContext: resourceAWSIntegrationUpdate,
		DeleteContext: resourceAWSIntegrationDelete,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "The ID of the source AWS account",
				Type:        schema.TypeString,
				Required:    true,
			},
			"role": {
				Description: "The name of the role which CloudTruth will assume in the AWS account",
				Type:        schema.TypeString,
				Required:    true,
			},
			"writable": {
				Description: "Whether or not the CloudTruth integration can write to the AWS service(s), defaults to false",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"kms_key_id": {
				Description: "The name of the role which CloudTruth will assume in the AWS account",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"aws_enabled_regions": {
				Description: "The AWS regions where the integration will be used, at lease one region must be specified",
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"aws_enabled_services": {
				Description: `The AWS services which the integration will use, one or more of ssm|secretsmanager|s3, 
at least one service must be specified`,
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceAWSIntegrationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSIntegrationCreate")
	defer tflog.Debug(ctx, "exiting resourceAWSIntegrationCreate")
	c := meta.(*cloudTruthClient)
	intCreate := cloudtruthapi.NewAwsIntegrationCreateWithDefaults()
	accountID := d.Get("account_id").(string)
	intCreate.SetAwsAccountId(accountID)
	role := d.Get("role").(string)
	intCreate.SetAwsRoleName(role)
	rawRegions := d.Get("aws_enabled_regions").([]interface{})
	regions := make([]cloudtruthapi.AwsRegionEnum, len(rawRegions))
	for i, v := range rawRegions {
		region, err := cloudtruthapi.NewAwsRegionEnumFromValue(v.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		regions[i] = *region
	}
	intCreate.SetAwsEnabledRegions(regions)
	rawServices := d.Get("aws_enabled_services").([]interface{})
	services := make([]cloudtruthapi.AwsServiceEnum, len(rawServices))
	for i, v := range rawServices {
		service, err := cloudtruthapi.NewAwsServiceEnumFromValue(v.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		services[i] = *service
	}
	intCreate.SetAwsEnabledServices(services)

	// Optional properties
	kmsKeyID := d.Get("kms_key_id").(string)
	if kmsKeyID != "" {
		intCreate.SetAwsKmsKeyId(kmsKeyID)
	}
	intCreate.SetWritable(d.Get("writable").(bool))

	var integration *cloudtruthapi.AwsIntegration
	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var r *http.Response
		var err error
		integration, r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsCreate(ctx).AwsIntegrationCreate(*intCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAWSIntegrationCreate: error creating integration %s@%s", role, accountID), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	d.SetId(integration.GetId())
	return nil
}

func resourceAWSIntegrationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSIntegrationRead")
	defer tflog.Debug(ctx, "exiting resourceAWSIntegrationRead")
	// todo
	return nil
}

func resourceAWSIntegrationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSIntegrationUpdate")
	defer tflog.Debug(ctx, "exiting resourceAWSIntegrationUpdate")
	// todo
	return resourceAWSIntegrationRead(ctx, d, meta)
}

func resourceAWSIntegrationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSIntegrationDelete")
	defer tflog.Debug(ctx, "exiting resourceAWSIntegrationDelete")
	c := meta.(*cloudTruthClient)
	accountID := d.Get("account_id").(string)
	role := d.Get("role").(string)
	intID := d.Id()

	retryError := resource.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.IntegrationsApi.IntegrationsAwsDestroy(ctx, intID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAWSIntegrationDelete: error destroying AWS integration %s@%s", role, accountID), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	return nil
}
