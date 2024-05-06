// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConfigurationInitParameters struct {

	// (String) The request property to target. Available values: ip, ip6, ip_range, asn, country. Modifying this attribute will force creation of a new resource.
	// The request property to target. Available values: `ip`, `ip6`, `ip_range`, `asn`, `country`. **Modifying this attribute will force creation of a new resource.**
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// (String) The value to target. Depends on target's type. Modifying this attribute will force creation of a new resource.
	// The value to target. Depends on target's type. **Modifying this attribute will force creation of a new resource.**
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConfigurationObservation struct {

	// (String) The request property to target. Available values: ip, ip6, ip_range, asn, country. Modifying this attribute will force creation of a new resource.
	// The request property to target. Available values: `ip`, `ip6`, `ip_range`, `asn`, `country`. **Modifying this attribute will force creation of a new resource.**
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// (String) The value to target. Depends on target's type. Modifying this attribute will force creation of a new resource.
	// The value to target. Depends on target's type. **Modifying this attribute will force creation of a new resource.**
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConfigurationParameters struct {

	// (String) The request property to target. Available values: ip, ip6, ip_range, asn, country. Modifying this attribute will force creation of a new resource.
	// The request property to target. Available values: `ip`, `ip6`, `ip_range`, `asn`, `country`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Target *string `json:"target" tf:"target,omitempty"`

	// (String) The value to target. Depends on target's type. Modifying this attribute will force creation of a new resource.
	// The value to target. Depends on target's type. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type RuleInitParameters struct {

	// (Block List, Min: 1, Max: 1) Rule configuration to apply to a matched request. Modifying this attribute will force creation of a new resource. (see below for nested schema)
	// Rule configuration to apply to a matched request. **Modifying this attribute will force creation of a new resource.**
	Configuration []ConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// (String) The action to apply to a matched request. Available values: block, challenge, whitelist, js_challenge, managed_challenge.
	// The action to apply to a matched request. Available values: `block`, `challenge`, `whitelist`, `js_challenge`, `managed_challenge`.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// (String) A personal note about the rule. Typically used as a reminder or explanation for the rule.
	// A personal note about the rule. Typically used as a reminder or explanation for the rule.
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`
}

type RuleObservation struct {

	// (String) The account identifier to target for the resource. Must provide only one of account_id, zone_id. Modifying this attribute will force creation of a new resource.
	// The account identifier to target for the resource. Must provide only one of `account_id`, `zone_id`. **Modifying this attribute will force creation of a new resource.**
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Block List, Min: 1, Max: 1) Rule configuration to apply to a matched request. Modifying this attribute will force creation of a new resource. (see below for nested schema)
	// Rule configuration to apply to a matched request. **Modifying this attribute will force creation of a new resource.**
	Configuration []ConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The action to apply to a matched request. Available values: block, challenge, whitelist, js_challenge, managed_challenge.
	// The action to apply to a matched request. Available values: `block`, `challenge`, `whitelist`, `js_challenge`, `managed_challenge`.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// (String) A personal note about the rule. Typically used as a reminder or explanation for the rule.
	// A personal note about the rule. Typically used as a reminder or explanation for the rule.
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// (String) The zone identifier to target for the resource. Must provide only one of account_id, zone_id. Modifying this attribute will force creation of a new resource.
	// The zone identifier to target for the resource. Must provide only one of `account_id`, `zone_id`. **Modifying this attribute will force creation of a new resource.**
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type RuleParameters struct {

	// (String) The account identifier to target for the resource. Must provide only one of account_id, zone_id. Modifying this attribute will force creation of a new resource.
	// The account identifier to target for the resource. Must provide only one of `account_id`, `zone_id`. **Modifying this attribute will force creation of a new resource.**
	// +crossplane:generate:reference:type=github.com/milkpirate/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// (Block List, Min: 1, Max: 1) Rule configuration to apply to a matched request. Modifying this attribute will force creation of a new resource. (see below for nested schema)
	// Rule configuration to apply to a matched request. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// (String) The action to apply to a matched request. Available values: block, challenge, whitelist, js_challenge, managed_challenge.
	// The action to apply to a matched request. Available values: `block`, `challenge`, `whitelist`, `js_challenge`, `managed_challenge`.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// (String) A personal note about the rule. Typically used as a reminder or explanation for the rule.
	// A personal note about the rule. Typically used as a reminder or explanation for the rule.
	// +kubebuilder:validation:Optional
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// (String) The zone identifier to target for the resource. Must provide only one of account_id, zone_id. Modifying this attribute will force creation of a new resource.
	// The zone identifier to target for the resource. Must provide only one of `account_id`, `zone_id`. **Modifying this attribute will force creation of a new resource.**
	// +crossplane:generate:reference:type=github.com/milkpirate/provider-cloudflare/apis/zone/v1alpha1.Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone in zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// RuleSpec defines the desired state of Rule
type RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RuleInitParameters `json:"initProvider,omitempty"`
}

// RuleStatus defines the observed state of Rule.
type RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Rule is the Schema for the Rules API. Provides a Cloudflare IP Firewall Access Rule resource. Access control can be applied on basis of IP addresses, IP ranges, AS numbers or countries.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configuration) || (has(self.initProvider) && has(self.initProvider.configuration))",message="spec.forProvider.configuration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mode) || (has(self.initProvider) && has(self.initProvider.mode))",message="spec.forProvider.mode is a required parameter"
	Spec   RuleSpec   `json:"spec"`
	Status RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleList contains a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

// Repository type metadata.
var (
	Rule_Kind             = "Rule"
	Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rule_Kind}.String()
	Rule_KindAPIVersion   = Rule_Kind + "." + CRDGroupVersion.String()
	Rule_GroupVersionKind = CRDGroupVersion.WithKind(Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}