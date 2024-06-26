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

type FirewallRulesetInitParameters struct {

	// A note that can be used to annotate the ruleset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the ruleset.
	// **Modifying this attribute will force creation of a new resource.**
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Rules []map[string]*string `json:"rules,omitempty" tf:"rules,omitempty"`
}

type FirewallRulesetObservation struct {

	// The ID of the account where the ruleset is being created.
	// The account identifier to target for the resource.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// A note that can be used to annotate the ruleset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the ruleset.
	// **Modifying this attribute will force creation of a new resource.**
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Rules []map[string]*string `json:"rules,omitempty" tf:"rules,omitempty"`
}

type FirewallRulesetParameters struct {

	// The ID of the account where the ruleset is being created.
	// The account identifier to target for the resource.
	// +crossplane:generate:reference:type=github.com/milkpirate/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// A note that can be used to annotate the ruleset.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the ruleset.
	// **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Rules []map[string]*string `json:"rules,omitempty" tf:"rules,omitempty"`
}

// FirewallRulesetSpec defines the desired state of FirewallRuleset
type FirewallRulesetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallRulesetParameters `json:"forProvider"`
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
	InitProvider FirewallRulesetInitParameters `json:"initProvider,omitempty"`
}

// FirewallRulesetStatus defines the observed state of FirewallRuleset.
type FirewallRulesetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallRulesetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRuleset is the Schema for the FirewallRulesets API. Provides the ability to manage a Magic Firewall Ruleset and it's firewall rules which are used with Magic Transit.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type FirewallRuleset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   FirewallRulesetSpec   `json:"spec"`
	Status FirewallRulesetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRulesetList contains a list of FirewallRulesets
type FirewallRulesetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallRuleset `json:"items"`
}

// Repository type metadata.
var (
	FirewallRuleset_Kind             = "FirewallRuleset"
	FirewallRuleset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallRuleset_Kind}.String()
	FirewallRuleset_KindAPIVersion   = FirewallRuleset_Kind + "." + CRDGroupVersion.String()
	FirewallRuleset_GroupVersionKind = CRDGroupVersion.WithKind(FirewallRuleset_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallRuleset{}, &FirewallRulesetList{})
}
