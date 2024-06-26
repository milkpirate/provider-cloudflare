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

type CustomPagesInitParameters struct {

	// (String) The id of the forbidden page.
	// The id of the forbidden page.
	Forbidden *string `json:"forbidden,omitempty" tf:"forbidden,omitempty"`

	// (String) The id of the identity denied page.
	// The id of the identity denied page.
	IdentityDenied *string `json:"identityDenied,omitempty" tf:"identity_denied,omitempty"`
}

type CustomPagesObservation struct {

	// (String) The id of the forbidden page.
	// The id of the forbidden page.
	Forbidden *string `json:"forbidden,omitempty" tf:"forbidden,omitempty"`

	// (String) The id of the identity denied page.
	// The id of the identity denied page.
	IdentityDenied *string `json:"identityDenied,omitempty" tf:"identity_denied,omitempty"`
}

type CustomPagesParameters struct {

	// (String) The id of the forbidden page.
	// The id of the forbidden page.
	// +kubebuilder:validation:Optional
	Forbidden *string `json:"forbidden,omitempty" tf:"forbidden,omitempty"`

	// (String) The id of the identity denied page.
	// The id of the identity denied page.
	// +kubebuilder:validation:Optional
	IdentityDenied *string `json:"identityDenied,omitempty" tf:"identity_denied,omitempty"`
}

type LoginDesignInitParameters struct {

	// (String) The background color on the login page.
	// The background color on the login page.
	BackgroundColor *string `json:"backgroundColor,omitempty" tf:"background_color,omitempty"`

	// (String) The text at the bottom of the login page.
	// The text at the bottom of the login page.
	FooterText *string `json:"footerText,omitempty" tf:"footer_text,omitempty"`

	// (String) The text at the top of the login page.
	// The text at the top of the login page.
	HeaderText *string `json:"headerText,omitempty" tf:"header_text,omitempty"`

	// (String) The URL of the logo on the login page.
	// The URL of the logo on the login page.
	LogoPath *string `json:"logoPath,omitempty" tf:"logo_path,omitempty"`

	// (String) The text color on the login page.
	// The text color on the login page.
	TextColor *string `json:"textColor,omitempty" tf:"text_color,omitempty"`
}

type LoginDesignObservation struct {

	// (String) The background color on the login page.
	// The background color on the login page.
	BackgroundColor *string `json:"backgroundColor,omitempty" tf:"background_color,omitempty"`

	// (String) The text at the bottom of the login page.
	// The text at the bottom of the login page.
	FooterText *string `json:"footerText,omitempty" tf:"footer_text,omitempty"`

	// (String) The text at the top of the login page.
	// The text at the top of the login page.
	HeaderText *string `json:"headerText,omitempty" tf:"header_text,omitempty"`

	// (String) The URL of the logo on the login page.
	// The URL of the logo on the login page.
	LogoPath *string `json:"logoPath,omitempty" tf:"logo_path,omitempty"`

	// (String) The text color on the login page.
	// The text color on the login page.
	TextColor *string `json:"textColor,omitempty" tf:"text_color,omitempty"`
}

type LoginDesignParameters struct {

	// (String) The background color on the login page.
	// The background color on the login page.
	// +kubebuilder:validation:Optional
	BackgroundColor *string `json:"backgroundColor,omitempty" tf:"background_color,omitempty"`

	// (String) The text at the bottom of the login page.
	// The text at the bottom of the login page.
	// +kubebuilder:validation:Optional
	FooterText *string `json:"footerText,omitempty" tf:"footer_text,omitempty"`

	// (String) The text at the top of the login page.
	// The text at the top of the login page.
	// +kubebuilder:validation:Optional
	HeaderText *string `json:"headerText,omitempty" tf:"header_text,omitempty"`

	// (String) The URL of the logo on the login page.
	// The URL of the logo on the login page.
	// +kubebuilder:validation:Optional
	LogoPath *string `json:"logoPath,omitempty" tf:"logo_path,omitempty"`

	// (String) The text color on the login page.
	// The text color on the login page.
	// +kubebuilder:validation:Optional
	TextColor *string `json:"textColor,omitempty" tf:"text_color,omitempty"`
}

type OrganizationInitParameters struct {

	// (Boolean) When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
	// When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp *bool `json:"allowAuthenticateViaWarp,omitempty" tf:"allow_authenticate_via_warp,omitempty"`

	// (String) The unique subdomain assigned to your Zero Trust organization.
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain *string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`

	// (Boolean) When set to true, users skip the identity provider selection step during login.
	// When set to true, users skip the identity provider selection step during login.
	AutoRedirectToIdentity *bool `json:"autoRedirectToIdentity,omitempty" tf:"auto_redirect_to_identity,omitempty"`

	// (Block List) Custom pages for your Zero Trust organization. (see below for nested schema)
	// Custom pages for your Zero Trust organization.
	CustomPages []CustomPagesInitParameters `json:"customPages,omitempty" tf:"custom_pages,omitempty"`

	// (Boolean) When set to true, this will disable all editing of Access resources via the Zero Trust Dashboard.
	// When set to true, this will disable all editing of Access resources via the Zero Trust Dashboard.
	IsUIReadOnly *bool `json:"isUiReadOnly,omitempty" tf:"is_ui_read_only,omitempty"`

	// (Block List) (see below for nested schema)
	LoginDesign []LoginDesignInitParameters `json:"loginDesign,omitempty" tf:"login_design,omitempty"`

	// (String) The name of your Zero Trust organization.
	// The name of your Zero Trust organization.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// authorise. Must be in the format 48h or 2h45m.
	// How often a user will be forced to re-authorise. Must be in the format `48h` or `2h45m`.
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// (String) A description of the reason why the UI read only field is being toggled.
	// A description of the reason why the UI read only field is being toggled.
	UIReadOnlyToggleReason *string `json:"uiReadOnlyToggleReason,omitempty" tf:"ui_read_only_toggle_reason,omitempty"`

	// (String) The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count. Must be in the format 300ms or 2h45m.
	// The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count. Must be in the format `300ms` or `2h45m`.
	UserSeatExpirationInactiveTime *string `json:"userSeatExpirationInactiveTime,omitempty" tf:"user_seat_expiration_inactive_time,omitempty"`

	// (String) The amount of time that tokens issued for applications will be valid. Must be in the format 30m or 2h45m. Valid time units are: m, h.
	// The amount of time that tokens issued for applications will be valid. Must be in the format 30m or 2h45m. Valid time units are: m, h.
	WarpAuthSessionDuration *string `json:"warpAuthSessionDuration,omitempty" tf:"warp_auth_session_duration,omitempty"`
}

type OrganizationObservation struct {

	// (String) The account identifier to target for the resource. Conflicts with zone_id.
	// The account identifier to target for the resource. Conflicts with `zone_id`.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Boolean) When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
	// When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp *bool `json:"allowAuthenticateViaWarp,omitempty" tf:"allow_authenticate_via_warp,omitempty"`

	// (String) The unique subdomain assigned to your Zero Trust organization.
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain *string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`

	// (Boolean) When set to true, users skip the identity provider selection step during login.
	// When set to true, users skip the identity provider selection step during login.
	AutoRedirectToIdentity *bool `json:"autoRedirectToIdentity,omitempty" tf:"auto_redirect_to_identity,omitempty"`

	// (Block List) Custom pages for your Zero Trust organization. (see below for nested schema)
	// Custom pages for your Zero Trust organization.
	CustomPages []CustomPagesObservation `json:"customPages,omitempty" tf:"custom_pages,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) When set to true, this will disable all editing of Access resources via the Zero Trust Dashboard.
	// When set to true, this will disable all editing of Access resources via the Zero Trust Dashboard.
	IsUIReadOnly *bool `json:"isUiReadOnly,omitempty" tf:"is_ui_read_only,omitempty"`

	// (Block List) (see below for nested schema)
	LoginDesign []LoginDesignObservation `json:"loginDesign,omitempty" tf:"login_design,omitempty"`

	// (String) The name of your Zero Trust organization.
	// The name of your Zero Trust organization.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// authorise. Must be in the format 48h or 2h45m.
	// How often a user will be forced to re-authorise. Must be in the format `48h` or `2h45m`.
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// (String) A description of the reason why the UI read only field is being toggled.
	// A description of the reason why the UI read only field is being toggled.
	UIReadOnlyToggleReason *string `json:"uiReadOnlyToggleReason,omitempty" tf:"ui_read_only_toggle_reason,omitempty"`

	// (String) The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count. Must be in the format 300ms or 2h45m.
	// The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count. Must be in the format `300ms` or `2h45m`.
	UserSeatExpirationInactiveTime *string `json:"userSeatExpirationInactiveTime,omitempty" tf:"user_seat_expiration_inactive_time,omitempty"`

	// (String) The amount of time that tokens issued for applications will be valid. Must be in the format 30m or 2h45m. Valid time units are: m, h.
	// The amount of time that tokens issued for applications will be valid. Must be in the format 30m or 2h45m. Valid time units are: m, h.
	WarpAuthSessionDuration *string `json:"warpAuthSessionDuration,omitempty" tf:"warp_auth_session_duration,omitempty"`

	// (String) The zone identifier to target for the resource. Conflicts with account_id.
	// The zone identifier to target for the resource. Conflicts with `account_id`.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type OrganizationParameters struct {

	// (String) The account identifier to target for the resource. Conflicts with zone_id.
	// The account identifier to target for the resource. Conflicts with `zone_id`.
	// +crossplane:generate:reference:type=github.com/milkpirate/provider-cloudflare/apis/account/v1alpha1.Account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Reference to a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDRef *v1.Reference `json:"accountIdRef,omitempty" tf:"-"`

	// Selector for a Account in account to populate accountId.
	// +kubebuilder:validation:Optional
	AccountIDSelector *v1.Selector `json:"accountIdSelector,omitempty" tf:"-"`

	// (Boolean) When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
	// When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
	// +kubebuilder:validation:Optional
	AllowAuthenticateViaWarp *bool `json:"allowAuthenticateViaWarp,omitempty" tf:"allow_authenticate_via_warp,omitempty"`

	// (String) The unique subdomain assigned to your Zero Trust organization.
	// The unique subdomain assigned to your Zero Trust organization.
	// +kubebuilder:validation:Optional
	AuthDomain *string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`

	// (Boolean) When set to true, users skip the identity provider selection step during login.
	// When set to true, users skip the identity provider selection step during login.
	// +kubebuilder:validation:Optional
	AutoRedirectToIdentity *bool `json:"autoRedirectToIdentity,omitempty" tf:"auto_redirect_to_identity,omitempty"`

	// (Block List) Custom pages for your Zero Trust organization. (see below for nested schema)
	// Custom pages for your Zero Trust organization.
	// +kubebuilder:validation:Optional
	CustomPages []CustomPagesParameters `json:"customPages,omitempty" tf:"custom_pages,omitempty"`

	// (Boolean) When set to true, this will disable all editing of Access resources via the Zero Trust Dashboard.
	// When set to true, this will disable all editing of Access resources via the Zero Trust Dashboard.
	// +kubebuilder:validation:Optional
	IsUIReadOnly *bool `json:"isUiReadOnly,omitempty" tf:"is_ui_read_only,omitempty"`

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	LoginDesign []LoginDesignParameters `json:"loginDesign,omitempty" tf:"login_design,omitempty"`

	// (String) The name of your Zero Trust organization.
	// The name of your Zero Trust organization.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// authorise. Must be in the format 48h or 2h45m.
	// How often a user will be forced to re-authorise. Must be in the format `48h` or `2h45m`.
	// +kubebuilder:validation:Optional
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// (String) A description of the reason why the UI read only field is being toggled.
	// A description of the reason why the UI read only field is being toggled.
	// +kubebuilder:validation:Optional
	UIReadOnlyToggleReason *string `json:"uiReadOnlyToggleReason,omitempty" tf:"ui_read_only_toggle_reason,omitempty"`

	// (String) The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count. Must be in the format 300ms or 2h45m.
	// The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count. Must be in the format `300ms` or `2h45m`.
	// +kubebuilder:validation:Optional
	UserSeatExpirationInactiveTime *string `json:"userSeatExpirationInactiveTime,omitempty" tf:"user_seat_expiration_inactive_time,omitempty"`

	// (String) The amount of time that tokens issued for applications will be valid. Must be in the format 30m or 2h45m. Valid time units are: m, h.
	// The amount of time that tokens issued for applications will be valid. Must be in the format 30m or 2h45m. Valid time units are: m, h.
	// +kubebuilder:validation:Optional
	WarpAuthSessionDuration *string `json:"warpAuthSessionDuration,omitempty" tf:"warp_auth_session_duration,omitempty"`

	// (String) The zone identifier to target for the resource. Conflicts with account_id.
	// The zone identifier to target for the resource. Conflicts with `account_id`.
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

// OrganizationSpec defines the desired state of Organization
type OrganizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationParameters `json:"forProvider"`
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
	InitProvider OrganizationInitParameters `json:"initProvider,omitempty"`
}

// OrganizationStatus defines the observed state of Organization.
type OrganizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Organization is the Schema for the Organizations API. A Zero Trust organization defines the user login experience.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Organization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authDomain) || (has(self.initProvider) && has(self.initProvider.authDomain))",message="spec.forProvider.authDomain is a required parameter"
	Spec   OrganizationSpec   `json:"spec"`
	Status OrganizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationList contains a list of Organizations
type OrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Organization `json:"items"`
}

// Repository type metadata.
var (
	Organization_Kind             = "Organization"
	Organization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Organization_Kind}.String()
	Organization_KindAPIVersion   = Organization_Kind + "." + CRDGroupVersion.String()
	Organization_GroupVersionKind = CRDGroupVersion.WithKind(Organization_Kind)
)

func init() {
	SchemeBuilder.Register(&Organization{}, &OrganizationList{})
}
