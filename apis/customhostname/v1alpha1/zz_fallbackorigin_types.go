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

type FallbackOriginInitParameters struct {

	// (String) Hostname you intend to fallback requests to. Origin must be a proxied A/AAAA/CNAME DNS record within Clouldflare.
	// Hostname you intend to fallback requests to. Origin must be a proxied A/AAAA/CNAME DNS record within Clouldflare.
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`
}

type FallbackOriginObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Hostname you intend to fallback requests to. Origin must be a proxied A/AAAA/CNAME DNS record within Clouldflare.
	// Hostname you intend to fallback requests to. Origin must be a proxied A/AAAA/CNAME DNS record within Clouldflare.
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`

	// (String) Status of the fallback origin's activation.
	// Status of the fallback origin's activation.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (String) The zone identifier to target for the resource. Modifying this attribute will force creation of a new resource.
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type FallbackOriginParameters struct {

	// (String) Hostname you intend to fallback requests to. Origin must be a proxied A/AAAA/CNAME DNS record within Clouldflare.
	// Hostname you intend to fallback requests to. Origin must be a proxied A/AAAA/CNAME DNS record within Clouldflare.
	// +kubebuilder:validation:Optional
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`

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

// FallbackOriginSpec defines the desired state of FallbackOrigin
type FallbackOriginSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FallbackOriginParameters `json:"forProvider"`
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
	InitProvider FallbackOriginInitParameters `json:"initProvider,omitempty"`
}

// FallbackOriginStatus defines the observed state of FallbackOrigin.
type FallbackOriginStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FallbackOriginObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FallbackOrigin is the Schema for the FallbackOrigins API. Provides a Cloudflare custom hostname fallback origin resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type FallbackOrigin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.origin) || (has(self.initProvider) && has(self.initProvider.origin))",message="spec.forProvider.origin is a required parameter"
	Spec   FallbackOriginSpec   `json:"spec"`
	Status FallbackOriginStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FallbackOriginList contains a list of FallbackOrigins
type FallbackOriginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FallbackOrigin `json:"items"`
}

// Repository type metadata.
var (
	FallbackOrigin_Kind             = "FallbackOrigin"
	FallbackOrigin_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FallbackOrigin_Kind}.String()
	FallbackOrigin_KindAPIVersion   = FallbackOrigin_Kind + "." + CRDGroupVersion.String()
	FallbackOrigin_GroupVersionKind = CRDGroupVersion.WithKind(FallbackOrigin_Kind)
)

func init() {
	SchemeBuilder.Register(&FallbackOrigin{}, &FallbackOriginList{})
}
