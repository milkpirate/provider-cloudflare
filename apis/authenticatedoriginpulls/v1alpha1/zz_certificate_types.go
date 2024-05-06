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

type CertificateInitParameters struct {

	// (String) The public client certificate. Modifying this attribute will force creation of a new resource.
	// The public client certificate. **Modifying this attribute will force creation of a new resource.**
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// zone, per-hostname. Modifying this attribute will force creation of a new resource.
	// The form of Authenticated Origin Pulls to upload the certificate to. Available values: `per-zone`, `per-hostname`. **Modifying this attribute will force creation of a new resource.**
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CertificateObservation struct {

	// (String) The public client certificate. Modifying this attribute will force creation of a new resource.
	// The public client certificate. **Modifying this attribute will force creation of a new resource.**
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// (String) Modifying this attribute will force creation of a new resource.
	// **Modifying this attribute will force creation of a new resource.**
	ExpiresOn *string `json:"expiresOn,omitempty" tf:"expires_on,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Modifying this attribute will force creation of a new resource.
	// **Modifying this attribute will force creation of a new resource.**
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// (String) Modifying this attribute will force creation of a new resource.
	// **Modifying this attribute will force creation of a new resource.**
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	// (String) Modifying this attribute will force creation of a new resource.
	// **Modifying this attribute will force creation of a new resource.**
	Signature *string `json:"signature,omitempty" tf:"signature,omitempty"`

	// (String) Modifying this attribute will force creation of a new resource.
	// **Modifying this attribute will force creation of a new resource.**
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// zone, per-hostname. Modifying this attribute will force creation of a new resource.
	// The form of Authenticated Origin Pulls to upload the certificate to. Available values: `per-zone`, `per-hostname`. **Modifying this attribute will force creation of a new resource.**
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) Modifying this attribute will force creation of a new resource.
	// **Modifying this attribute will force creation of a new resource.**
	UploadedOn *string `json:"uploadedOn,omitempty" tf:"uploaded_on,omitempty"`

	// (String) The zone identifier to target for the resource. Modifying this attribute will force creation of a new resource.
	// The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type CertificateParameters struct {

	// (String) The public client certificate. Modifying this attribute will force creation of a new resource.
	// The public client certificate. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// (String, Sensitive) The private key of the client certificate. Modifying this attribute will force creation of a new resource.
	// The private key of the client certificate. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// zone, per-hostname. Modifying this attribute will force creation of a new resource.
	// The form of Authenticated Origin Pulls to upload the certificate to. Available values: `per-zone`, `per-hostname`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

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

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
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
	InitProvider CertificateInitParameters `json:"initProvider,omitempty"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Certificate is the Schema for the Certificates API. Provides a Cloudflare Authenticated Origin Pulls certificate resource. An uploaded client certificate is required to use Per-Zone or Per-Hostname Authenticated Origin Pulls.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.certificate) || (has(self.initProvider) && has(self.initProvider.certificate))",message="spec.forProvider.certificate is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privateKeySecretRef)",message="spec.forProvider.privateKeySecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   CertificateSpec   `json:"spec"`
	Status CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}