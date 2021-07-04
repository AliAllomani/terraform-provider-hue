package hue

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccHueLight_basic(t *testing.T) {

	lightIndex := 1
	resourceName := "hue_light.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckHueLightDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHueLightConfig(lightIndex),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "state.0.on", "true"),
				),
			},
			// {
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func testAccCheckHueLightDestroy(s *terraform.State) error {

	for rsName, rs := range s.RootModule().Resources {
		if rs.Type != "hue_light" {
			continue
		}

		return fmt.Errorf("The light '%s' with ID '%s' was not deleted properly", rsName, rs.Primary.ID)
	}

	return nil
}

func testAccHueLightConfig(lightIndex int) string {
	return fmt.Sprintf(`

resource "hue_light" "test" {
  light_index = %d

  state {
    on = true
  }
}
`, lightIndex)
}
