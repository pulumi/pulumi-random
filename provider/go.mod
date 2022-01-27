module github.com/pulumi/pulumi-random/provider/v4

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.11.0
	github.com/pulumi/pulumi/sdk/v3 v3.22.0
	github.com/terraform-providers/terraform-provider-random/shim v0.0.0
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e

	github.com/pulumi/pulumi-terraform-bridge/v3 => github.com/pulumi/pulumi-terraform-bridge/v3 v3.16.2-0.20220126223139-b865ae4ef63f
	github.com/terraform-providers/terraform-provider-random => github.com/pulumi/terraform-provider-random v1.3.2-0.20210512110019-3feaba7fd924
	github.com/terraform-providers/terraform-provider-random/shim => ./shim
)
