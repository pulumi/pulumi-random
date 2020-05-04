module github.com/pulumi/pulumi-random/provider/v2

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.0.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.2.0
	github.com/pulumi/pulumi/sdk/v2 v2.0.1-0.20200424001829-090f390d7b1a
	github.com/terraform-providers/terraform-provider-random v0.0.0-20190925210718-83518d96ae4f
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
