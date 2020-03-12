module github.com/pulumi/pulumi-random

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.0.0
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.12.1
	github.com/pulumi/pulumi-terraform-bridge v1.8.2
	github.com/terraform-providers/terraform-provider-random v0.0.0-20190925210718-83518d96ae4f
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
