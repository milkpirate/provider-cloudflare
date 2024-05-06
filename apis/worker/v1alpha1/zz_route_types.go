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

type RouteInitParameters struct {

	// (String) The route pattern to associate the Worker with.
	// The [route pattern](https://developers.cloudflare.com/workers/about/routes/) to associate the Worker with.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type RouteObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The route pattern to associate the Worker with.
	// The [route pattern](https://developers.cloudflare.com/workers/about/routes/) to associate the Worker with.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// (String) Worker script name to invoke for requests that match the route pattern.
	// Worker script name to invoke for requests that match the route pattern.
	ScriptName *string `json:"scriptName,omitempty" tf:"script_name,omitempty"`

	// (String) The zone identifier to target for the resource. Modifying this attribute will force creation of a new resource.
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type RouteParameters struct {

	// (String) The route pattern to associate the Worker with.
	// The [route pattern](https://developers.cloudflare.com/workers/about/routes/) to associate the Worker with.
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// (String) Worker script name to invoke for requests that match the route pattern.
	// Worker script name to invoke for requests that match the route pattern.
	// +crossplane:generate:reference:type=Script
	// +kubebuilder:validation:Optional
	ScriptName *string `json:"scriptName,omitempty" tf:"script_name,omitempty"`

	// Reference to a Script to populate scriptName.
	// +kubebuilder:validation:Optional
	ScriptNameRef *v1.Reference `json:"scriptNameRef,omitempty" tf:"-"`

	// Selector for a Script to populate scriptName.
	// +kubebuilder:validation:Optional
	ScriptNameSelector *v1.Selector `json:"scriptNameSelector,omitempty" tf:"-"`

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

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteParameters `json:"forProvider"`
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
	InitProvider RouteInitParameters `json:"initProvider,omitempty"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route is the Schema for the Routes API. Provides a Cloudflare worker route resource. A route will also require a cloudflare_worker_script.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.pattern) || (has(self.initProvider) && has(self.initProvider.pattern))",message="spec.forProvider.pattern is a required parameter"
	Spec   RouteSpec   `json:"spec"`
	Status RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	Route_Kind             = "Route"
	Route_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Route_Kind}.String()
	Route_KindAPIVersion   = Route_Kind + "." + CRDGroupVersion.String()
	Route_GroupVersionKind = CRDGroupVersion.WithKind(Route_Kind)
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
