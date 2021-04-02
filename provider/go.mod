module github.com/pulumi/pulumi-random/provider/v4

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0-20210325101119-03683ad99004
	github.com/pulumi/pulumi/sdk/v3 v3.0.0-beta.1
	github.com/terraform-providers/terraform-provider-random/shim v0.0.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/terraform-providers/terraform-provider-random/shim => ./shim
)
