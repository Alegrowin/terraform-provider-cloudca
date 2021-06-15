
package cloudca

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudCAResourceEnvironment(t *testing.T) {
	t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudCAResourceEnvironment,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"cloudca_environment.environment", "service_code", regexp.MustCompile("^compute-qc")),
				),
			},
		},
	})
}

const testAccCloudCAResourceEnvironment = `
resource "cloudca_environment" "environment" {
	service_code      = "compute-qc"
	organization_code = "test"
	name              = "production"
	description       = "Environment for production workloads"
	admin_role        = ["pat"]
	read_only_role    = ["franz","bob"]
}
`