module github.com/pulumi/pulumi-random/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.11.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.9.2
	github.com/pulumi/pulumi/sdk/v2 v2.10.2-0.20200916204740-92a7d717a4d1
	github.com/pulumi/tf2pulumi v0.9.0 // indirect
	github.com/terraform-providers/terraform-provider-random v0.0.0-20190925210718-83518d96ae4f
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/terraform-providers/terraform-provider-random => github.com/pulumi/terraform-provider-random v1.3.2-0.20200707155416-47754b56d29d
)
