//go:build acceptance
// +build acceptance

package acceptance

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/terraform-provider-stevan/internal/provider"
	"testing"
)

const (
	providerConfig = `
		provider "stevan" {
        host = "host"
        auth_token = "auth_token"
}

	`
)

var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"stevan": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)

func TestAccstevanApiProductResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig +
					`
resource "stevan_api_product" "example" {
    name = "My Name"

    description = "MyDescription"

    labels = "labels"

}

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("stevan_api_product.example", "name", "name"),
					resource.TestCheckResourceAttr("stevan_api_product.example", "description", "description"),
					resource.TestCheckResourceAttr("stevan_api_product.example", "labels", "labels"),
				),
			},
		},
	})
}
