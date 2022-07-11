package cloudtruth

/*
import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"math/rand"
	"time"
)

DISABLE FOR NOW

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"math/rand"
	"testing"
	"time"
)

func TestAccResourceUserBasic(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	email := fmt.Sprintf("matthewcummings516+%d@gmail.com", rand.Intn(100000))
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviderFactories,
		PreCheck:          func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserCreateBasic(email, ownerRole),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloudtruth_user.basic", "email", email),
					resource.TestCheckResourceAttr("cloudtruth_user.basic", "role", ownerRole),
				),
			},
		},
	})
}

func testAccResourceUserCreateBasic(email, role string) string {
	return fmt.Sprintf(`
	resource "cloudtruth_user" "basic" {
  		email = "%s"
  		role  = "%s"
	}
	`, email, role)
}

*/
