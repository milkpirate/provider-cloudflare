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

type AccountInitParameters struct {

	// (Boolean) Whether 2FA is enforced on the account. Defaults to false.
	// Whether 2FA is enforced on the account. Defaults to `false`.
	EnforceTwofactor *bool `json:"enforceTwofactor,omitempty" tf:"enforce_twofactor,omitempty"`

	// (String) The name of the account that is displayed in the Cloudflare dashboard.
	// The name of the account that is displayed in the Cloudflare dashboard.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Account type. Available values: enterprise, standard. Defaults to standard. Modifying this attribute will force creation of a new resource.
	// Account type. Available values: `enterprise`, `standard`. Defaults to `standard`. **Modifying this attribute will force creation of a new resource.**
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AccountObservation struct {

	// (Boolean) Whether 2FA is enforced on the account. Defaults to false.
	// Whether 2FA is enforced on the account. Defaults to `false`.
	EnforceTwofactor *bool `json:"enforceTwofactor,omitempty" tf:"enforce_twofactor,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the account that is displayed in the Cloudflare dashboard.
	// The name of the account that is displayed in the Cloudflare dashboard.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Account type. Available values: enterprise, standard. Defaults to standard. Modifying this attribute will force creation of a new resource.
	// Account type. Available values: `enterprise`, `standard`. Defaults to `standard`. **Modifying this attribute will force creation of a new resource.**
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AccountParameters struct {

	// (Boolean) Whether 2FA is enforced on the account. Defaults to false.
	// Whether 2FA is enforced on the account. Defaults to `false`.
	// +kubebuilder:validation:Optional
	EnforceTwofactor *bool `json:"enforceTwofactor,omitempty" tf:"enforce_twofactor,omitempty"`

	// (String) The name of the account that is displayed in the Cloudflare dashboard.
	// The name of the account that is displayed in the Cloudflare dashboard.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Account type. Available values: enterprise, standard. Defaults to standard. Modifying this attribute will force creation of a new resource.
	// Account type. Available values: `enterprise`, `standard`. Defaults to `standard`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountParameters `json:"forProvider"`
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
	InitProvider AccountInitParameters `json:"initProvider,omitempty"`
}

// AccountStatus defines the observed state of Account.
type AccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Account is the Schema for the Accounts API. Provides a Cloudflare Account resource. Account is the basic resource for working with Cloudflare zones, teams and users.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AccountSpec   `json:"spec"`
	Status AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountList contains a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

// Repository type metadata.
var (
	Account_Kind             = "Account"
	Account_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Account_Kind}.String()
	Account_KindAPIVersion   = Account_Kind + "." + CRDGroupVersion.String()
	Account_GroupVersionKind = CRDGroupVersion.WithKind(Account_Kind)
)

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
