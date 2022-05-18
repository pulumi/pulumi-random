module github.com/pulumi/pulumi-random/provider/v4

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.22.1
	github.com/pulumi/pulumi/sdk/v3 v3.32.1
	github.com/terraform-providers/terraform-provider-random/shim v0.0.0
)

replace (
	github.com/hashicorp/terraform-exec => github.com/hashicorp/terraform-exec v0.15.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211230170131-3a7c83bfab87
	// Temporary: Replace when pulumi-yaml > 0.3.0:
	github.com/pulumi/pulumi/pkg/v3 => github.com/pulumi/pulumi/pkg/v3 v3.32.1
	github.com/terraform-providers/terraform-provider-random => github.com/pulumi/terraform-provider-random v1.3.2-0.20220518210607-abc35fdeceb1
	github.com/terraform-providers/terraform-provider-random/shim => ./shim
)
