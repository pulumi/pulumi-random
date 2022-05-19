module github.com/terraform-providers/terraform-provider-random/shim

go 1.15

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.16.0
	github.com/terraform-providers/terraform-provider-random v1.3.2-0.20210112153945-304bbf724bde
)

replace github.com/terraform-providers/terraform-provider-random => github.com/pulumi/terraform-provider-random v1.3.2-0.20220518210607-abc35fdeceb1
