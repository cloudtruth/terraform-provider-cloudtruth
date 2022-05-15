package cloudtruth

import (
	"testing"
	//"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProject(t *testing.T) {
	t.Skip("resource not yet implemented, remove this once you add your own code")
}

const testAccResourceProject = `
resource "project_resource" "foo" {
  sample_attribute = "bar"
}
`
