module github.com/pulumi/pulumi-random

require (
	github.com/Azure/go-autorest/autorest v0.9.0 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.1.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.2.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.1.0 // indirect
	github.com/hashicorp/terraform v0.12.6
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.1.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190828172748-3f206601e7a1
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-random v0.0.0-20190808172436-0ce299bbf59e
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7 // indirect
)

replace (
	github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v31.1.0+incompatible
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

go 1.13
