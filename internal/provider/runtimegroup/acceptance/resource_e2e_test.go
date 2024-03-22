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

func TestAccstevanRunTimeGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig +
					`
resource "stevan_runtime_group" "example" {
    name = "Test Runtime Group"

    description = "A test runtime group for exploration."

    labels = "labels"

    cluster_type = "cluster_type"

    auth_type = "auth_type"

}

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("stevan_runtime_group.example", "name", "name"),
					resource.TestCheckResourceAttr("stevan_runtime_group.example", "description", "description"),
					resource.TestCheckResourceAttr("stevan_runtime_group.example", "labels", "labels"),
					resource.TestCheckResourceAttr("stevan_runtime_group.example", "cluster_type", "cluster_type"),
					resource.TestCheckResourceAttr("stevan_runtime_group.example", "auth_type", "auth_type"),
				),
			},
		},
	})
}
