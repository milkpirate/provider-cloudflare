// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/milkpirate/provider-cloudflare/internal/controller/access/application"
	cacertificate "github.com/milkpirate/provider-cloudflare/internal/controller/access/cacertificate"
	group "github.com/milkpirate/provider-cloudflare/internal/controller/access/group"
	identityprovider "github.com/milkpirate/provider-cloudflare/internal/controller/access/identityprovider"
	keysconfiguration "github.com/milkpirate/provider-cloudflare/internal/controller/access/keysconfiguration"
	mutualtlscertificate "github.com/milkpirate/provider-cloudflare/internal/controller/access/mutualtlscertificate"
	organization "github.com/milkpirate/provider-cloudflare/internal/controller/access/organization"
	policy "github.com/milkpirate/provider-cloudflare/internal/controller/access/policy"
	rule "github.com/milkpirate/provider-cloudflare/internal/controller/access/rule"
	servicetoken "github.com/milkpirate/provider-cloudflare/internal/controller/access/servicetoken"
	account "github.com/milkpirate/provider-cloudflare/internal/controller/account/account"
	apitoken "github.com/milkpirate/provider-cloudflare/internal/controller/account/apitoken"
	member "github.com/milkpirate/provider-cloudflare/internal/controller/account/member"
	apishield "github.com/milkpirate/provider-cloudflare/internal/controller/apishield/apishield"
	argo "github.com/milkpirate/provider-cloudflare/internal/controller/argo/argo"
	tunnel "github.com/milkpirate/provider-cloudflare/internal/controller/argo/tunnel"
	tunnelconfig "github.com/milkpirate/provider-cloudflare/internal/controller/argo/tunnelconfig"
	tunnelroute "github.com/milkpirate/provider-cloudflare/internal/controller/argo/tunnelroute"
	tunnelvirtualnetwork "github.com/milkpirate/provider-cloudflare/internal/controller/argo/tunnelvirtualnetwork"
	authenticatedoriginspulls "github.com/milkpirate/provider-cloudflare/internal/controller/authenticatedoriginpulls/authenticatedoriginspulls"
	certificate "github.com/milkpirate/provider-cloudflare/internal/controller/authenticatedoriginpulls/certificate"
	ipprefix "github.com/milkpirate/provider-cloudflare/internal/controller/byoip/ipprefix"
	pack "github.com/milkpirate/provider-cloudflare/internal/controller/certificate/pack"
	pages "github.com/milkpirate/provider-cloudflare/internal/controller/custom/pages"
	ssl "github.com/milkpirate/provider-cloudflare/internal/controller/custom/ssl"
	fallbackorigin "github.com/milkpirate/provider-cloudflare/internal/controller/customhostname/fallbackorigin"
	hostname "github.com/milkpirate/provider-cloudflare/internal/controller/customhostname/hostname"
	profile "github.com/milkpirate/provider-cloudflare/internal/controller/dlp/profile"
	record "github.com/milkpirate/provider-cloudflare/internal/controller/dns/record"
	address "github.com/milkpirate/provider-cloudflare/internal/controller/emailrouting/address"
	catchall "github.com/milkpirate/provider-cloudflare/internal/controller/emailrouting/catchall"
	ruleemailrouting "github.com/milkpirate/provider-cloudflare/internal/controller/emailrouting/rule"
	settings "github.com/milkpirate/provider-cloudflare/internal/controller/emailrouting/settings"
	filter "github.com/milkpirate/provider-cloudflare/internal/controller/filters/filter"
	rulefirewall "github.com/milkpirate/provider-cloudflare/internal/controller/firewall/rule"
	list "github.com/milkpirate/provider-cloudflare/internal/controller/lists/list"
	loadbalancer "github.com/milkpirate/provider-cloudflare/internal/controller/loadbalancer/loadbalancer"
	monitor "github.com/milkpirate/provider-cloudflare/internal/controller/loadbalancer/monitor"
	pool "github.com/milkpirate/provider-cloudflare/internal/controller/loadbalancer/pool"
	job "github.com/milkpirate/provider-cloudflare/internal/controller/logpush/job"
	ownershipchallenge "github.com/milkpirate/provider-cloudflare/internal/controller/logpush/ownershipchallenge"
	firewallruleset "github.com/milkpirate/provider-cloudflare/internal/controller/magic/firewallruleset"
	gretunnel "github.com/milkpirate/provider-cloudflare/internal/controller/magic/gretunnel"
	ipsectunnel "github.com/milkpirate/provider-cloudflare/internal/controller/magic/ipsectunnel"
	staticroute "github.com/milkpirate/provider-cloudflare/internal/controller/magic/staticroute"
	policynotification "github.com/milkpirate/provider-cloudflare/internal/controller/notification/policy"
	policywebhooks "github.com/milkpirate/provider-cloudflare/internal/controller/notification/policywebhooks"
	certificateoriginca "github.com/milkpirate/provider-cloudflare/internal/controller/originca/certificate"
	rulepage "github.com/milkpirate/provider-cloudflare/internal/controller/page/rule"
	domain "github.com/milkpirate/provider-cloudflare/internal/controller/pages/domain"
	project "github.com/milkpirate/provider-cloudflare/internal/controller/pages/project"
	providerconfig "github.com/milkpirate/provider-cloudflare/internal/controller/providerconfig"
	ruleset "github.com/milkpirate/provider-cloudflare/internal/controller/ruleset/ruleset"
	applicationspectrum "github.com/milkpirate/provider-cloudflare/internal/controller/spectrum/application"
	accountteams "github.com/milkpirate/provider-cloudflare/internal/controller/teams/account"
	listteams "github.com/milkpirate/provider-cloudflare/internal/controller/teams/list"
	location "github.com/milkpirate/provider-cloudflare/internal/controller/teams/location"
	proxyendpoint "github.com/milkpirate/provider-cloudflare/internal/controller/teams/proxyendpoint"
	ruleteams "github.com/milkpirate/provider-cloudflare/internal/controller/teams/rule"
	event "github.com/milkpirate/provider-cloudflare/internal/controller/waitingroom/event"
	room "github.com/milkpirate/provider-cloudflare/internal/controller/waitingroom/room"
	rules "github.com/milkpirate/provider-cloudflare/internal/controller/waitingroom/rules"
	devicepolicycertificates "github.com/milkpirate/provider-cloudflare/internal/controller/warp/devicepolicycertificates"
	devicepostureintegration "github.com/milkpirate/provider-cloudflare/internal/controller/warp/devicepostureintegration"
	deviceposturerule "github.com/milkpirate/provider-cloudflare/internal/controller/warp/deviceposturerule"
	devicesettingspolicy "github.com/milkpirate/provider-cloudflare/internal/controller/warp/devicesettingspolicy"
	fallbackdomain "github.com/milkpirate/provider-cloudflare/internal/controller/warp/fallbackdomain"
	splittunnel "github.com/milkpirate/provider-cloudflare/internal/controller/warp/splittunnel"
	hostnameweb3 "github.com/milkpirate/provider-cloudflare/internal/controller/web3/hostname"
	crontrigger "github.com/milkpirate/provider-cloudflare/internal/controller/worker/crontrigger"
	kv "github.com/milkpirate/provider-cloudflare/internal/controller/worker/kv"
	kvnamespace "github.com/milkpirate/provider-cloudflare/internal/controller/worker/kvnamespace"
	route "github.com/milkpirate/provider-cloudflare/internal/controller/worker/route"
	script "github.com/milkpirate/provider-cloudflare/internal/controller/worker/script"
	dnssec "github.com/milkpirate/provider-cloudflare/internal/controller/zone/dnssec"
	healthcheck "github.com/milkpirate/provider-cloudflare/internal/controller/zone/healthcheck"
	logpullretention "github.com/milkpirate/provider-cloudflare/internal/controller/zone/logpullretention"
	managedheaders "github.com/milkpirate/provider-cloudflare/internal/controller/zone/managedheaders"
	ratelimit "github.com/milkpirate/provider-cloudflare/internal/controller/zone/ratelimit"
	settingsoverride "github.com/milkpirate/provider-cloudflare/internal/controller/zone/settingsoverride"
	tieredcache "github.com/milkpirate/provider-cloudflare/internal/controller/zone/tieredcache"
	totaltls "github.com/milkpirate/provider-cloudflare/internal/controller/zone/totaltls"
	urlnormalizationsettings "github.com/milkpirate/provider-cloudflare/internal/controller/zone/urlnormalizationsettings"
	useragentblockingrule "github.com/milkpirate/provider-cloudflare/internal/controller/zone/useragentblockingrule"
	zone "github.com/milkpirate/provider-cloudflare/internal/controller/zone/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		cacertificate.Setup,
		group.Setup,
		identityprovider.Setup,
		keysconfiguration.Setup,
		mutualtlscertificate.Setup,
		organization.Setup,
		policy.Setup,
		rule.Setup,
		servicetoken.Setup,
		account.Setup,
		apitoken.Setup,
		member.Setup,
		apishield.Setup,
		argo.Setup,
		tunnel.Setup,
		tunnelconfig.Setup,
		tunnelroute.Setup,
		tunnelvirtualnetwork.Setup,
		authenticatedoriginspulls.Setup,
		certificate.Setup,
		ipprefix.Setup,
		pack.Setup,
		pages.Setup,
		ssl.Setup,
		fallbackorigin.Setup,
		hostname.Setup,
		profile.Setup,
		record.Setup,
		address.Setup,
		catchall.Setup,
		ruleemailrouting.Setup,
		settings.Setup,
		filter.Setup,
		rulefirewall.Setup,
		list.Setup,
		loadbalancer.Setup,
		monitor.Setup,
		pool.Setup,
		job.Setup,
		ownershipchallenge.Setup,
		firewallruleset.Setup,
		gretunnel.Setup,
		ipsectunnel.Setup,
		staticroute.Setup,
		policynotification.Setup,
		policywebhooks.Setup,
		certificateoriginca.Setup,
		rulepage.Setup,
		domain.Setup,
		project.Setup,
		providerconfig.Setup,
		ruleset.Setup,
		applicationspectrum.Setup,
		accountteams.Setup,
		listteams.Setup,
		location.Setup,
		proxyendpoint.Setup,
		ruleteams.Setup,
		event.Setup,
		room.Setup,
		rules.Setup,
		devicepolicycertificates.Setup,
		devicepostureintegration.Setup,
		deviceposturerule.Setup,
		devicesettingspolicy.Setup,
		fallbackdomain.Setup,
		splittunnel.Setup,
		hostnameweb3.Setup,
		crontrigger.Setup,
		kv.Setup,
		kvnamespace.Setup,
		route.Setup,
		script.Setup,
		dnssec.Setup,
		healthcheck.Setup,
		logpullretention.Setup,
		managedheaders.Setup,
		ratelimit.Setup,
		settingsoverride.Setup,
		tieredcache.Setup,
		totaltls.Setup,
		urlnormalizationsettings.Setup,
		useragentblockingrule.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
