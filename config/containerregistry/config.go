package containerregistry

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-azure/apis/rconfig"
	"github.com/upbound/provider-azure/config/common"
)

// Configure configures containerregistry group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_container_registry", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"encryption"},
		}
	})

	p.AddResourceConfigurator("azurerm_container_registry_token", func(r *config.Resource) {
		r.References["scope_map_id"] = config.Reference{
			Type:      "ScopeMap",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_container_connected_registry", func(r *config.Resource) {
		r.References["container_registry_id"] = config.Reference{
			Type:      "Registry",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["sync_token_id"] = config.Reference{
			Type:      "Token",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.ExternalName.IdentifierFields = common.RemoveIndex(r.ExternalName.IdentifierFields, "container_registry_id")
	})

	p.AddResourceConfigurator("azurerm_container_registry_webhook", func(r *config.Resource) {
		r.ExternalName.IdentifierFields = common.RemoveIndex(r.ExternalName.IdentifierFields, "registry_name")
	})
}
