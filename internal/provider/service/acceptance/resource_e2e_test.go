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

func TestAccstevanServiceResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig +
					`
resource "stevan_get_service_200_response" "example" {
    connect_timeout = "412598"

    created_at = "93089629"

    enabled = true

    host = "labore tempor proident cillum"

    id = "Duis pariatur Excepteur"

    name = "consectetur sunt esse"

    path = "laborum sunt reprehenderit"

    port = "-12117843"

    protocol = "mollit in officia"

    read_timeout = "21222558"

    retries = "-84754244"

    updated_at = "-75092970"

    write_timeout = "-99302911"

    ca_certificates = [
        "ca_certificates"
    ]

    client_certificate = {
                id = "in commodo aliquip"
}


    tags = [
        "tags"
    ]

    tls_verify = true

    tls_verify_depth = "-14601984"

    url = "ut ex esse"

    runtime_group_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"

    service_id = "velit culpa ad consectetur nisi"

}

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "connect_timeout", 1234),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "created_at", 1234),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "enabled", false),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "host", "host"),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "id", "id"),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "name", "name"),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "path", "path"),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "port", 1234),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "protocol", "protocol"),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "read_timeout", 1234),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "retries", 1234),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "updated_at", 1234),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "write_timeout", 1234),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "ca_certificates.0", "ca_certificates"),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "client_certificate.id", "id"),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "tags.0", "tags"),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "tls_verify", false),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "tls_verify_depth", 1234),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "url", "url"),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "runtime_group_id", "runtime_group_id"),
					resource.TestCheckResourceAttr("stevan_get_service_200_response.example", "service_id", "service_id"),
				),
			},
		},
	})
}
