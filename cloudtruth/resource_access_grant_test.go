package cloudtruth

/*
import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"math/rand"
	"testing"
	"time"
)


func TestAccResourceAccessGrantGroup(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	groupName := fmt.Sprintf("TestGroup-%d", rand.Intn(100000))
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			},
		},
	})
}

func TestAccResourceAccessGrantUser(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	groupName := fmt.Sprintf("TestGroup-%d", rand.Intn(100000))
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
		},
	})
}

func testAccResourceGrantUser(name, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_group" "basic" {
  		name = "%s"
  		description  = "%s"
	}
	`, name, desc)
}

func testAccResourceGrantGroup(name, desc string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_group" "basic" {
  		name = "%s"
  		description  = "%s"
	}
	`, name, desc)
}
*/
