package cloudtruth

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
)

func paramOrTemplateStateIDFunc(resourceName, projectName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		paramID := s.RootModule().Resources[resourceName].Primary.ID
		return fmt.Sprintf("%s.%s", projectName, paramID), nil
	}
}

func parseProjectAndID(ctx context.Context, c *cloudTruthClient, projParamID string) (*string, *string, error) {
	tflog.Debug(ctx, "entering parseProjectAndID")
	defer tflog.Debug(ctx, "exiting parseProjectAndID")

	projDotParam := strings.Split(projParamID, ".")
	if len(projDotParam) != 2 {
		return nil, nil, fmt.Errorf("invalid import ID format: %s, you must use the format 'PROJECT_NAME.PARAMETER_ID' or 'PROJECT_NAME.TEMPLATE_ID'", projParamID)
	}
	projName, paramOrTemplateID := projDotParam[0], projDotParam[1]
	projID, err := c.lookupProject(ctx, projName)
	if err != nil {
		return nil, nil, err
	}
	return projID, &paramOrTemplateID, nil
}

/*
We cannot import Parameters or Templates solely by ID, the project must also be specified, therefore we
this function instead of using schema.ImportStatePassthroughContext. We support this in the following formats:
PROJECT_NAME.PARAMETER_ID or PROJECT_NAME.TEMPLATE_ID
*/
func paramOrTemplateImportHelper(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	tflog.Debug(ctx, "entering paramOrTemplateImportHelper")
	defer tflog.Debug(ctx, "exiting paramOrTemplateImportHelper")
	c := meta.(*cloudTruthClient)

	projID, paramOrTemplateID, err := parseProjectAndID(ctx, c, d.Id())
	if err != nil {
		return nil, err
	}
	projName, err := c.lookupProject(ctx, *projID) // We have the project ID, look up its name
	if err != nil {
		return nil, err
	}
	err = d.Set("project", *projName)
	if err != nil {
		return nil, err
	}
	d.SetId(*paramOrTemplateID)

	return []*schema.ResourceData{d}, nil
}
