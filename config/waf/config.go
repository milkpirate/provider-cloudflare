package waf

import "github.com/crossplane/upjet/pkg/config"

// Configure adds configurations for waf group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_waf_group", func(r *config.Resource) {
		r.ShortGroup = "waf"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/milkpirate/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_waf_override", func(r *config.Resource) {
		r.ShortGroup = "waf"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/milkpirate/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_waf_package", func(r *config.Resource) {
		r.ShortGroup = "waf"
		r.Kind = "WAFPackage"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/milkpirate/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
	p.AddResourceConfigurator("cloudflare_waf_rule", func(r *config.Resource) {
		r.ShortGroup = "waf"
		r.References["zone_id"] = config.Reference{
			Type: "github.com/milkpirate/provider-cloudflare/apis/zone/v1alpha1.Zone",
		}
	})
}
