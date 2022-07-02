package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

const testAccEpochTimeTag = `
data "cloudtruth_tag" "epoch_time" {
  project     = "%s"
  environment = "%s"
  name        = "%s"
}
`

func TestDataSourceTag(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccEpochTimeTag, accTestProject, defaultEnv, epochTimeTag),
				Check: resource.ComposeTestCheckFunc(
					// epoch time tag
					resource.TestCheckResourceAttr("data.cloudtruth_tag.epoch_time", "environment", defaultEnv),
					resource.TestCheckResourceAttr("data.cloudtruth_tag.epoch_time", "project", accTestProject),
					resource.TestCheckResourceAttr("data.cloudtruth_tag.epoch_time", "name", epochTimeTag),
					resource.TestCheckResourceAttr("data.cloudtruth_tag.epoch_time", "timestamp", epochTimeTagTimestamp),
				),
			},
		},
	})
}
