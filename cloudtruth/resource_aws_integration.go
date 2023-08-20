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
	"sort"
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
				Description: "The name of the role which CloudTruth will assume in the target AWS account",
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
				Description: `The ID of the KMS key which CloudTruth will use to decrypt content in the target AWS account (optional and
needed only when the content is encryped with a non-default KMS key)`,
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"aws_enabled_regions": {
				Description: "The AWS regions where the integration will be used, at lease one region must be specified",
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"aws_enabled_services": {
				Description: `The AWS services which the integration will use, one or more of ssm|secretsmanager|s3, 
at least one service must be specified`,
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"external_id": {
				Description: `The generated external ID for the AWS integration, needed for CloudTruth to assume the specified role
in the target AWS account`,
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func getIntegrationRegions(rawRegions []interface{}) ([]cloudtruthapi.AwsRegionEnum, error) {
	regions := make([]cloudtruthapi.AwsRegionEnum, len(rawRegions))
	for i, v := range rawRegions {
		region, err := cloudtruthapi.NewAwsRegionEnumFromValue(v.(string))
		if err != nil {
			return nil, err
		}
		regions[i] = *region
	}
	return regions, nil
}

func getIntegrationServices(rawServices []interface{}) ([]cloudtruthapi.AwsServiceEnum, error) {
	services := make([]cloudtruthapi.AwsServiceEnum, len(rawServices))
	for i, v := range rawServices {
		service, err := cloudtruthapi.NewAwsServiceEnumFromValue(v.(string))
		if err != nil {
			return nil, err
		}
		services[i] = *service
	}
	return services, nil
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
	rawRegions := d.Get("aws_enabled_regions").(*schema.Set)
	regions, err := getIntegrationRegions(rawRegions.List())
	if err != nil {
		return diag.FromErr(err)
	}
	intCreate.SetAwsEnabledRegions(regions)
	rawServices := d.Get("aws_enabled_services").(*schema.Set)
	services, err := getIntegrationServices(rawServices.List())
	if err != nil {
		return diag.FromErr(err)
	}
	intCreate.SetAwsEnabledServices(services)

	// Optional properties
	kmsKeyID := d.Get("kms_key_id").(string)
	if kmsKeyID != "" {
		intCreate.SetAwsKmsKeyId(kmsKeyID)
	}
	intCreate.SetWritable(d.Get("writable").(bool))

	var integration *cloudtruthapi.AwsIntegration
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
		var r *http.Response
		var err error
		integration, r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAwsCreate(ctx).AwsIntegrationCreate(*intCreate).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAWSIntegrationCreate: error creating integration %s@%s", role, accountID), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}
	err = d.Set("external_id", integration.GetAwsExternalId())
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(integration.GetId())
	return nil
}

func resourceAWSIntegrationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSIntegrationRead")
	defer tflog.Debug(ctx, "exiting resourceAWSIntegrationRead")
	c := meta.(*cloudTruthClient)
	accountID := d.Get("account_id").(string)
	role := d.Get("role").(string)
	integrationID := d.Id()

	var integration *cloudtruthapi.AwsIntegration
	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutRead), func() *retry.RetryError {
		var r *http.Response
		var err error
		integration, r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAwsRetrieve(ctx, integrationID).Execute()
		if err != nil {
			return handleAPIError(fmt.Sprintf("resourceAWSIntegrationRead: error reading AWS integration %s@%s",
				role, accountID), r, err)
		}
		return nil
	})
	if retryError != nil {
		return diag.FromErr(retryError)
	}

	// Read & set all properties from the deployed resource
	err := d.Set("writable", integration.GetWritable())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("kms_key_id", integration.GetAwsKmsKeyId())
	if err != nil {
		return diag.FromErr(err)
	}
	regions := integration.GetAwsEnabledRegions()
	sort.Slice(regions, func(i, j int) bool {
		return regions[i] < regions[j]
	})
	err = d.Set("aws_enabled_regions", regions)
	if err != nil {
		return diag.FromErr(err)
	}
	services := integration.GetAwsEnabledServices()
	sort.Slice(services, func(i, j int) bool {
		return services[i] < services[j]
	})
	err = d.Set("aws_enabled_services", services)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(integrationID)
	return nil
}

func resourceAWSIntegrationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSIntegrationUpdate")
	defer tflog.Debug(ctx, "exiting resourceAWSIntegrationUpdate")
	c := meta.(*cloudTruthClient)
	integrationID := d.Id()
	accountID := d.Get("account_id").(string)
	role := d.Get("role").(string)
	patchedAwsIntegration := cloudtruthapi.PatchedAwsIntegration{}
	hasChange := false
	props := map[string]func(v string){
		"account_id": patchedAwsIntegration.SetAwsAccountId,
		"role":       patchedAwsIntegration.SetAwsRoleName,
		"kms_key_id": patchedAwsIntegration.SetAwsKmsKeyId,
	}
	for prop := range props {
		if d.HasChange(prop) {
			props[prop](d.Get(prop).(string))
			hasChange = true
		}
	}
	if d.HasChange("writable") {
		patchedAwsIntegration.SetWritable(d.Get("writable").(bool))
		hasChange = true
	}

	if d.HasChange("writable") {
		patchedAwsIntegration.SetWritable(d.Get("writable").(bool))
		hasChange = true
	}

	if d.HasChange("aws_enabled_regions") {
		x, y := d.GetChange("aws_enabled_regions")
		fmt.Print(x)
		fmt.Print(y)

		rawRegions := d.Get("aws_enabled_regions").(*schema.Set)
		regions, err := getIntegrationRegions(rawRegions.List())
		if err != nil {
			return diag.FromErr(err)
		}
		patchedAwsIntegration.SetAwsEnabledRegions(regions)
		hasChange = true
	}
	if d.HasChange("aws_enabled_service") {
		rawServices := d.Get("aws_enabled_service").(*schema.Set)
		services, err := getIntegrationServices(rawServices.List())
		if err != nil {
			return diag.FromErr(err)
		}
		patchedAwsIntegration.SetAwsEnabledServices(services)
		hasChange = true
	}

	if hasChange {
		retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutUpdate), func() *retry.RetryError {
			var r *http.Response
			var err error
			_, r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAwsPartialUpdate(ctx, integrationID).PatchedAwsIntegration(patchedAwsIntegration).Execute()
			if err != nil {
				return handleAPIError(fmt.Sprintf("resourceAWSIntegrationUpdate: error updating AWS integration %s@%s",
					role, accountID), r, err)
			}
			return nil
		})
		if retryError != nil {
			return diag.FromErr(retryError)
		}
	}
	d.SetId(integrationID)
	return resourceAWSIntegrationRead(ctx, d, meta)
}

func resourceAWSIntegrationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "entering resourceAWSIntegrationDelete")
	defer tflog.Debug(ctx, "exiting resourceAWSIntegrationDelete")
	c := meta.(*cloudTruthClient)
	accountID := d.Get("account_id").(string)
	role := d.Get("role").(string)
	intID := d.Id()

	retryError := retry.RetryContext(ctx, d.Timeout(schema.TimeoutDelete), func() *retry.RetryError {
		var r *http.Response
		var err error
		r, err = c.openAPIClient.IntegrationsAPI.IntegrationsAwsDestroy(ctx, intID).Execute()
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
