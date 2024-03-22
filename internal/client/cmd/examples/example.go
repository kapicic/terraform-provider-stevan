package main

import (
	"fmt"

	"github.com/stevan-sdk/pkg/client"
)

func main() {
	client := client.NewClient("http://example.com", "TOKEN")

	fmt.Printf("%#v", client)
}
