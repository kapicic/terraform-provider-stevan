# stevan Terraform Provider 1.0.0
The stevan platform API
This repository contains a Terraform provider that allows you to manage resources through the stevan API.

## Prerequisites

- [Go](https://golang.org/doc/install) >= 1.19

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0

- Access to the stevan API.

## Installing The Provider

1. Clone the repository:
```bash
git clone https://github.com/liblaber/terraform-provider-stevan.git
```

2. Navigate to the directory:
```bash
cd terraform-provider-stevan
```

3. Update module references:
```bash
go mod tidy
```

4. Build the provider:
```bash
go build -o terraform-provider-stevan
```

5. Move the provider to your plugins directory:
```bash
mkdir -p ~/.terraform.d/plugins/example.com/user/stevan/1.0.0/<distribution>
mv terraform-provider-stevan ~/.terraform.d/plugins/example.com/user/stevan/1.0.0/<distribution>
```

Note: The directory structure is important. The provider must be located at `~/.terraform.d/plugins/example.com/user/stevan/1.0.0/<distribution>/terraform-provider-stevan`
Also please change `example.com/user` and `<distribution>` to match your real values.
To get the <distribution> run `terraform version`, possible values: `linux_amd64`, `darwin_arm64`, `windows_amd64`, etc.

## Setting Up The Provider

1. Configure the provider:

In your Terraform configuration, reference the provider and supply the necessary credentials:

```hcl
provider "stevan" {
host = "https://localhost/"
api_token = "YOUR_API_TOKEN"
}
```

## Running The Provider

To plan and apply your Terraform configuration:

1. Initialize your configuration:

```bash
terraform init -plugin-dir=~/.terraform.d/plugins
```

2. Plan your changes:

```bash
terraform plan
```

3. Apply your configuration:

```bash
terraform apply
```

## Debugging

If you encounter any issues or unexpected behaviors, enable debug mode by setting the environment variable:

```bash
export TF_PROVIDER_DEBUG=true
```

Then, run your Terraform commands.

## Running Tests

1. Generate the docs:
```bash
go generate ./...
```

2. To execute the provider's tests, follow these steps:

**a. Run Unit Tests**:
```bash
make unit-test
```

**b. Run Acceptance Tests**:
```bash
make acceptance-test
```

## Publishing the Provider

1. Tag your release:

```bash
git tag v1.0.0
git push --tags
```

2. Build a release binary for your platform:

```bash
GOOS=linux GOARCH=amd64 go build -o terraform-provider-stevan
```

3. Upload the binary to the GitHub release or any other distribution method you prefer.

Note: For wide-reaching utility, consider registering your provider with the official Terraform provider registry once
it becomes popular within the community.
