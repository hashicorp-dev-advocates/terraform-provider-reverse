// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccStringResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccStringResourceConfig("one"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"reverse_string.test",
						tfjsonpath.New("input"),
						knownvalue.StringExact("one"),
					),
					statecheck.ExpectKnownValue(
						"reverse_string.test",
						tfjsonpath.New("result"),
						knownvalue.StringExact("eno"),
					),
				},
			},
			// Update and Read testing
			{
				Config: testAccStringResourceConfig("two"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"reverse_string.test",
						tfjsonpath.New("input"),
						knownvalue.StringExact("two"),
					),
					statecheck.ExpectKnownValue(
						"reverse_string.test",
						tfjsonpath.New("result"),
						knownvalue.StringExact("owt"),
					),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccStringResourceConfig(input string) string {
	return fmt.Sprintf(`
resource "reverse_string" "test" {
  input = %[1]q
}
`, input)
}
