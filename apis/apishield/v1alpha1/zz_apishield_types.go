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

type APIShieldInitParameters struct {

	// ids can be computed in a privacy-preserving manner. (see below for nested schema)
	// Characteristics define properties across which auth-ids can be computed in a privacy-preserving manner.
	AuthIDCharacteristics []AuthIDCharacteristicsInitParameters `json:"authIdCharacteristics,omitempty" tf:"auth_id_characteristics,omitempty"`
}

type APIShieldObservation struct {

	// ids can be computed in a privacy-preserving manner. (see below for nested schema)
	// Characteristics define properties across which auth-ids can be computed in a privacy-preserving manner.
	AuthIDCharacteristics []AuthIDCharacteristicsObservation `json:"authIdCharacteristics,omitempty" tf:"auth_id_characteristics,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The zone identifier to target for the resource. Modifying this attribute will force creation of a new resource.
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type APIShieldParameters struct {

	// ids can be computed in a privacy-preserving manner. (see below for nested schema)
	// Characteristics define properties across which auth-ids can be computed in a privacy-preserving manner.
	// +kubebuilder:validation:Optional
	AuthIDCharacteristics []AuthIDCharacteristicsParameters `json:"authIdCharacteristics,omitempty" tf:"auth_id_characteristics,omitempty"`

	// (String) The zone identifier to target for the resource. Modifying this attribute will force creation of a new resource.
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
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

type AuthIDCharacteristicsInitParameters struct {

	// (String) The name of the characteristic.
	// The name of the characteristic.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The type of characteristic. Available values: header, cookie.
	// The type of characteristic. Available values: `header`, `cookie`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AuthIDCharacteristicsObservation struct {

	// (String) The name of the characteristic.
	// The name of the characteristic.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The type of characteristic. Available values: header, cookie.
	// The type of characteristic. Available values: `header`, `cookie`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AuthIDCharacteristicsParameters struct {

	// (String) The name of the characteristic.
	// The name of the characteristic.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The type of characteristic. Available values: header, cookie.
	// The type of characteristic. Available values: `header`, `cookie`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// APIShieldSpec defines the desired state of APIShield
type APIShieldSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIShieldParameters `json:"forProvider"`
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
	InitProvider APIShieldInitParameters `json:"initProvider,omitempty"`
}

// APIShieldStatus defines the observed state of APIShield.
type APIShieldStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIShieldObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIShield is the Schema for the APIShields API. Provides a resource to manage API Shield configurations.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type APIShield struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIShieldSpec   `json:"spec"`
	Status            APIShieldStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIShieldList contains a list of APIShields
type APIShieldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIShield `json:"items"`
}

// Repository type metadata.
var (
	APIShield_Kind             = "APIShield"
	APIShield_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIShield_Kind}.String()
	APIShield_KindAPIVersion   = APIShield_Kind + "." + CRDGroupVersion.String()
	APIShield_GroupVersionKind = CRDGroupVersion.WithKind(APIShield_Kind)
)

func init() {
	SchemeBuilder.Register(&APIShield{}, &APIShieldList{})
}
