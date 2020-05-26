module github.com/pulumi/pulumi-random/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.3.3
	github.com/pulumi/pulumi/pkg/v2 v2.2.1 // indirect
	github.com/pulumi/pulumi/sdk/v2 v2.2.1
	github.com/terraform-providers/terraform-provider-random v0.0.0-20190925210718-83518d96ae4f
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
