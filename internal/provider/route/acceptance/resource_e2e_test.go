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

func TestAccstevanRouteResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig +
					`
resource "stevan_route" "example" {
    created_at = "79316364"

    headers = "headers"

    hosts = [
        "hosts"
    ]

    https_redirect_status_code = "426"

    id = "56c4566c-14cc-4132-9011-4139fcbbe50a"

    methods = [
        "methods"
    ]

    name = "deserunt dolor aliquip ut"

    path_handling = "v0"

    paths = [
        "paths"
    ]

    preserve_host = false

    protocols = [
        "protocols"
    ]

    regex_priority = "regex_priority"

    request_buffering = true

    response_buffering = true

    service = {
                id = "tempor mollit exercitation"
}


    snis = [
        "snis"
    ]

    strip_path = true

    tags = [
        "tags"
    ]

    updated_at = "-26100906"

    runtime_group_id = "9524ec7d-36d9-465d-a8c5-83a3c9390458"

    route_id = "e94215b0-9198-42ce-bf0b-98afff31b2a7"

}

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("stevan_route.example", "created_at", 1234),
					resource.TestCheckResourceAttr("stevan_route.example", "headers", "headers"),
					resource.TestCheckResourceAttr("stevan_route.example", "hosts.0", "hosts"),
					resource.TestCheckResourceAttr("stevan_route.example", "https_redirect_status_code", 1234),
					resource.TestCheckResourceAttr("stevan_route.example", "id", "id"),
					resource.TestCheckResourceAttr("stevan_route.example", "methods.0", "methods"),
					resource.TestCheckResourceAttr("stevan_route.example", "name", "name"),
					resource.TestCheckResourceAttr("stevan_route.example", "path_handling", "path_handling"),
					resource.TestCheckResourceAttr("stevan_route.example", "paths.0", "paths"),
					resource.TestCheckResourceAttr("stevan_route.example", "preserve_host", false),
					resource.TestCheckResourceAttr("stevan_route.example", "protocols.0", "protocols"),
					resource.TestCheckResourceAttr("stevan_route.example", "regex_priority", 1234),
					resource.TestCheckResourceAttr("stevan_route.example", "request_buffering", false),
					resource.TestCheckResourceAttr("stevan_route.example", "response_buffering", false),
					resource.TestCheckResourceAttr("stevan_route.example", "service.id", "id"),
					resource.TestCheckResourceAttr("stevan_route.example", "snis.0", "snis"),
					resource.TestCheckResourceAttr("stevan_route.example", "strip_path", false),
					resource.TestCheckResourceAttr("stevan_route.example", "tags.0", "tags"),
					resource.TestCheckResourceAttr("stevan_route.example", "updated_at", 1234),
					resource.TestCheckResourceAttr("stevan_route.example", "runtime_group_id", "runtime_group_id"),
					resource.TestCheckResourceAttr("stevan_route.example", "route_id", "route_id"),
				),
			},
		},
	})
}
