package page

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for page group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_page_rule", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "github.com/milkpirate/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
