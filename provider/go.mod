module github.com/pulumi/pulumi-random/provider/v2

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.3.0
	github.com/pulumi/pulumi/pkg/v2 v2.1.1-0.20200505193935-7b446d6c55ad // indirect
	github.com/pulumi/pulumi/sdk/v2 v2.1.1-0.20200505193935-7b446d6c55ad
	github.com/terraform-providers/terraform-provider-random v0.0.0-20190925210718-83518d96ae4f
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
