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

func TestAccstevanApiProductVersionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig +
					`
resource "stevan_api_product_version" "example" {
    name = "FirstVersion"

    gateway_service = {
                id = "09b4786a-3e48-4631-8f6b-62d1d8e1a7f3"
                control_plane_id = "e4d9ebb1-26b4-426a-b00e-cb67044f3baf"
}


    publish_status = "publish_status"

    deprecated = false

    notify = true

    api_product_id = "d32d905a-ed33-46a3-a093-d8f536af9a8a"

}

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("stevan_api_product_version.example", "name", "name"),
					resource.TestCheckResourceAttr("stevan_api_product_version.example", "gateway_service.id", "id"),
					resource.TestCheckResourceAttr("stevan_api_product_version.example", "gateway_service.control_plane_id", "control_plane_id"),
					resource.TestCheckResourceAttr("stevan_api_product_version.example", "publish_status", "publish_status"),
					resource.TestCheckResourceAttr("stevan_api_product_version.example", "deprecated", false),
					resource.TestCheckResourceAttr("stevan_api_product_version.example", "notify", false),
					resource.TestCheckResourceAttr("stevan_api_product_version.example", "api_product_id", "api_product_id"),
				),
			},
		},
	})
}
