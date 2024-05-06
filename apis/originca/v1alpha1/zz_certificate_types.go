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

	// encoded. Modifying this attribute will force creation of a new resource.
	// The Certificate Signing Request. Must be newline-encoded. **Modifying this attribute will force creation of a new resource.**
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`

	// (Set of String) A list of hostnames or wildcard names bound to the certificate. Modifying this attribute will force creation of a new resource.
	// A list of hostnames or wildcard names bound to the certificate. **Modifying this attribute will force creation of a new resource.**
	Hostnames []*string `json:"hostnames,omitempty" tf:"hostnames,omitempty"`

	MinDaysForRenewal *float64 `json:"minDaysForRenewal,omitempty" tf:"min_days_for_renewal,omitempty"`

	// rsa, origin-ecc, keyless-certificate. Modifying this attribute will force creation of a new resource.
	// The signature type desired on the certificate. Available values: `origin-rsa`, `origin-ecc`, `keyless-certificate`. **Modifying this attribute will force creation of a new resource.**
	RequestType *string `json:"requestType,omitempty" tf:"request_type,omitempty"`

	// (Number) The number of days for which the certificate should be valid. Available values: 7, 30, 90, 365, 730, 1095, 5475. Modifying this attribute will force creation of a new resource.
	// The number of days for which the certificate should be valid. Available values: `7`, `30`, `90`, `365`, `730`, `1095`, `5475`. **Modifying this attribute will force creation of a new resource.**
	RequestedValidity *float64 `json:"requestedValidity,omitempty" tf:"requested_validity,omitempty"`
}

type CertificateObservation struct {

	// (String) The Origin CA certificate.
	// The Origin CA certificate.
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// encoded. Modifying this attribute will force creation of a new resource.
	// The Certificate Signing Request. Must be newline-encoded. **Modifying this attribute will force creation of a new resource.**
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`

	// (String) The datetime when the certificate will expire.
	// The datetime when the certificate will expire.
	ExpiresOn *string `json:"expiresOn,omitempty" tf:"expires_on,omitempty"`

	// (Set of String) A list of hostnames or wildcard names bound to the certificate. Modifying this attribute will force creation of a new resource.
	// A list of hostnames or wildcard names bound to the certificate. **Modifying this attribute will force creation of a new resource.**
	Hostnames []*string `json:"hostnames,omitempty" tf:"hostnames,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MinDaysForRenewal *float64 `json:"minDaysForRenewal,omitempty" tf:"min_days_for_renewal,omitempty"`

	// rsa, origin-ecc, keyless-certificate. Modifying this attribute will force creation of a new resource.
	// The signature type desired on the certificate. Available values: `origin-rsa`, `origin-ecc`, `keyless-certificate`. **Modifying this attribute will force creation of a new resource.**
	RequestType *string `json:"requestType,omitempty" tf:"request_type,omitempty"`

	// (Number) The number of days for which the certificate should be valid. Available values: 7, 30, 90, 365, 730, 1095, 5475. Modifying this attribute will force creation of a new resource.
	// The number of days for which the certificate should be valid. Available values: `7`, `30`, `90`, `365`, `730`, `1095`, `5475`. **Modifying this attribute will force creation of a new resource.**
	RequestedValidity *float64 `json:"requestedValidity,omitempty" tf:"requested_validity,omitempty"`
}

type CertificateParameters struct {

	// encoded. Modifying this attribute will force creation of a new resource.
	// The Certificate Signing Request. Must be newline-encoded. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Csr *string `json:"csr,omitempty" tf:"csr,omitempty"`

	// (Set of String) A list of hostnames or wildcard names bound to the certificate. Modifying this attribute will force creation of a new resource.
	// A list of hostnames or wildcard names bound to the certificate. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	Hostnames []*string `json:"hostnames,omitempty" tf:"hostnames,omitempty"`

	// +kubebuilder:validation:Optional
	MinDaysForRenewal *float64 `json:"minDaysForRenewal,omitempty" tf:"min_days_for_renewal,omitempty"`

	// rsa, origin-ecc, keyless-certificate. Modifying this attribute will force creation of a new resource.
	// The signature type desired on the certificate. Available values: `origin-rsa`, `origin-ecc`, `keyless-certificate`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	RequestType *string `json:"requestType,omitempty" tf:"request_type,omitempty"`

	// (Number) The number of days for which the certificate should be valid. Available values: 7, 30, 90, 365, 730, 1095, 5475. Modifying this attribute will force creation of a new resource.
	// The number of days for which the certificate should be valid. Available values: `7`, `30`, `90`, `365`, `730`, `1095`, `5475`. **Modifying this attribute will force creation of a new resource.**
	// +kubebuilder:validation:Optional
	RequestedValidity *float64 `json:"requestedValidity,omitempty" tf:"requested_validity,omitempty"`
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

// Certificate is the Schema for the Certificates API. Provides a Cloudflare Origin CA certificate used to protect traffic to your origin without involving a third party Certificate Authority.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.csr) || (has(self.initProvider) && has(self.initProvider.csr))",message="spec.forProvider.csr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostnames) || (has(self.initProvider) && has(self.initProvider.hostnames))",message="spec.forProvider.hostnames is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.requestType) || (has(self.initProvider) && has(self.initProvider.requestType))",message="spec.forProvider.requestType is a required parameter"
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