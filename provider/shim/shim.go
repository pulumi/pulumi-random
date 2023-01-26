package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/terraform-providers/terraform-provider-random/internal/provider"
)

func NewProvider() tfpf.Provider {
	return provider.New()
}
