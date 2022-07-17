package cloudtruth

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type user struct {
	name     string
	email    string
	role     string
	userType string
}

func genUserDataSource(sourceName string, u user) string {
	return fmt.Sprintf(`
data "cloudtruth_user" "%s" {
	name = "%s"
	email = "%s"
}
`, sourceName, u.name, u.email)
}

func genUsersDataSource(sourceName string, userType string) string {
	return fmt.Sprintf(`
data "cloudtruth_users" "%s" {
	type  = "%s"
}
`, sourceName, userType)
}

// todo: provide an way to set test input in the file (and others)
// so that this can run against any CloudTruth organization
var testUser1 = user{
	name:     "Matthew Conway",
	email:    "matt@cloudtruth.com",
	role:     "OWNER",
	userType: "interactive",
}

var testUser2 = user{
	name:     "Matthew Cummings",
	email:    "matthewcummings516@gmail.com",
	role:     "OWNER",
	userType: "interactive",
}

func TestDataSourceUser(t *testing.T) {
	sourceName := "test_user_1"
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: genUserDataSource(sourceName, testUser1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_user.%s", sourceName), "name", testUser1.name),
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_user.%s", sourceName), "role", testUser1.role),
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_user.%s", sourceName), "email", testUser1.email),
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_user.%s", sourceName), "type", testUser1.userType),
				),
			},
		},
	})
}

func TestDataSourceUsers(t *testing.T) {
	sourceName := "test_users"
	userType := "interactive"
	resource.Test(t, resource.TestCase{
		ProviderFactories:         testProviderFactories,
		PreCheck:                  func() { testAccPreCheck(t) },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: genUsersDataSource(sourceName, userType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_users.%s", sourceName), "type", userType),
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_users.%s", sourceName), "users.0.name", testUser1.name),
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_users.%s", sourceName), "users.0.role", testUser1.role),
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_users.%s", sourceName), "users.0.email", testUser1.email),
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_users.%s", sourceName), "users.0.type", testUser1.userType),
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_users.%s", sourceName), "users.1.name", testUser2.name),
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_users.%s", sourceName), "users.1.role", testUser2.role),
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_users.%s", sourceName), "users.1.email", testUser2.email),
					resource.TestCheckResourceAttr(fmt.Sprintf("data.cloudtruth_users.%s", sourceName), "users.1.type", testUser2.userType),
				),
			},
		},
	})
}
