module github.com/pulumi/pulumi-random/provider/v4

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.8.2-0.20210930033847-5bba386a0e79
	github.com/pulumi/pulumi/sdk/v3 v3.13.2
	github.com/terraform-providers/terraform-provider-random/shim v0.0.0
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/terraform-providers/terraform-provider-random/shim => ./shim
)
