/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/milkpirate/provider-cloudflare/config/access"
	"github.com/milkpirate/provider-cloudflare/config/account"
	"github.com/milkpirate/provider-cloudflare/config/apishield"
	"github.com/milkpirate/provider-cloudflare/config/argo"
	"github.com/milkpirate/provider-cloudflare/config/authenticatedoriginpulls"
	"github.com/milkpirate/provider-cloudflare/config/byoip"
	"github.com/milkpirate/provider-cloudflare/config/certificate"
	"github.com/milkpirate/provider-cloudflare/config/custom"
	"github.com/milkpirate/provider-cloudflare/config/customhostname"
	"github.com/milkpirate/provider-cloudflare/config/dlp"
	"github.com/milkpirate/provider-cloudflare/config/dns"
	"github.com/milkpirate/provider-cloudflare/config/emailrouting"
	"github.com/milkpirate/provider-cloudflare/config/filters"
	"github.com/milkpirate/provider-cloudflare/config/firewall"
	"github.com/milkpirate/provider-cloudflare/config/lists"
	"github.com/milkpirate/provider-cloudflare/config/loadbalancer"
	"github.com/milkpirate/provider-cloudflare/config/logpush"
	"github.com/milkpirate/provider-cloudflare/config/magic"
	"github.com/milkpirate/provider-cloudflare/config/notification"
	"github.com/milkpirate/provider-cloudflare/config/originca"
	"github.com/milkpirate/provider-cloudflare/config/page"
	"github.com/milkpirate/provider-cloudflare/config/pages"
	"github.com/milkpirate/provider-cloudflare/config/ruleset"
	"github.com/milkpirate/provider-cloudflare/config/spectrum"
	"github.com/milkpirate/provider-cloudflare/config/teams"
	"github.com/milkpirate/provider-cloudflare/config/waf"
	"github.com/milkpirate/provider-cloudflare/config/waitingroom"
	"github.com/milkpirate/provider-cloudflare/config/warp"
	"github.com/milkpirate/provider-cloudflare/config/web3"
	"github.com/milkpirate/provider-cloudflare/config/worker"
	"github.com/milkpirate/provider-cloudflare/config/zone"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "cloudflare"
	modulePath     = "github.com/milkpirate/provider-cloudflare"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		access.Configure,
		account.Configure,
		apishield.Configure,
		argo.Configure,
		authenticatedoriginpulls.Configure,
		byoip.Configure,
		certificate.Configure,
		custom.Configure,
		customhostname.Configure,
		dlp.Configure,
		dns.Configure,
		emailrouting.Configure,
		filters.Configure,
		firewall.Configure,
		lists.Configure,
		loadbalancer.Configure,
		logpush.Configure,
		magic.Configure,
		notification.Configure,
		originca.Configure,
		page.Configure,
		pages.Configure,
		ruleset.Configure,
		spectrum.Configure,
		teams.Configure,
		waf.Configure,
		waitingroom.Configure,
		warp.Configure,
		web3.Configure,
		worker.Configure,
		zone.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
