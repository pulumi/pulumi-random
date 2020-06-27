module github.com/pulumi/pulumi-random/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.5.2-0.20200625200836-ce2050eb7e10
	github.com/pulumi/pulumi/sdk/v2 v2.5.0
	github.com/terraform-providers/terraform-provider-random v0.0.0-20190925210718-83518d96ae4f
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
replace github.com/pulumi/pulumi-terraform-bridge/v2 => ../../pulumi-terraform-bridge
replace github.com/pulumi/pulumi/pkg/v2 => ../../pulumi/pkg

