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

	// (Boolean) Whether to enable the activity log.
	// Whether to enable the activity log.
	ActivityLogEnabled *bool `json:"activityLogEnabled,omitempty" tf:"activity_log_enabled,omitempty"`

	// (Block List, Max: 1) Configuration block for antivirus traffic scanning. (see below for nested schema)
	// Configuration block for antivirus traffic scanning.
	Antivirus []AntivirusInitParameters `json:"antivirus,omitempty" tf:"antivirus,omitempty"`

	// (Block List, Max: 1) Configuration for a custom block page. (see below for nested schema)
	// Configuration for a custom block page.
	BlockPage []BlockPageInitParameters `json:"blockPage,omitempty" tf:"block_page,omitempty"`

	// (Block List, Max: 1) Configuration for body scanning. (see below for nested schema)
	// Configuration for body scanning.
	BodyScanning []BodyScanningInitParameters `json:"bodyScanning,omitempty" tf:"body_scanning,omitempty"`

	// PKI. (see below for nested schema)
	// Configuration for custom certificates / BYO-PKI.
	CustomCertificate []CustomCertificateInitParameters `json:"customCertificate,omitempty" tf:"custom_certificate,omitempty"`

	// mail matching. (see below for nested schema)
	// Configuration for extended e-mail matching.
	ExtendedEmailMatching []ExtendedEmailMatchingInitParameters `json:"extendedEmailMatching,omitempty" tf:"extended_email_matching,omitempty"`

	// (Block List, Max: 1) Configure compliance with Federal Information Processing Standards. (see below for nested schema)
	// Configure compliance with Federal Information Processing Standards.
	Fips []FipsInitParameters `json:"fips,omitempty" tf:"fips,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	Logging []LoggingInitParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// identity onramp for Browser Isolation. Defaults to false.
	// Enable non-identity onramp for Browser Isolation. Defaults to `false`.
	NonIdentityBrowserIsolationEnabled *bool `json:"nonIdentityBrowserIsolationEnabled,omitempty" tf:"non_identity_browser_isolation_enabled,omitempty"`

	// (Block List, Max: 1) Configuration for DLP Payload Logging. (see below for nested schema)
	// Configuration for DLP Payload Logging.
	PayloadLog []PayloadLogInitParameters `json:"payloadLog,omitempty" tf:"payload_log,omitempty"`

	// (Boolean) Indicator that protocol detection is enabled.
	// Indicator that protocol detection is enabled.
	ProtocolDetectionEnabled *bool `json:"protocolDetectionEnabled,omitempty" tf:"protocol_detection_enabled,omitempty"`

	// (Block List, Max: 1) Configuration block for specifying which protocols are proxied. (see below for nested schema)
	// Configuration block for specifying which protocols are proxied.
	Proxy []ProxyInitParameters `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// (Block List, Max: 1) Configuration for SSH Session Logging. (see below for nested schema)
	// Configuration for SSH Session Logging.
	SSHSessionLog []SSHSessionLogInitParameters `json:"sshSessionLog,omitempty" tf:"ssh_session_log,omitempty"`

	// (Boolean) Indicator that decryption of TLS traffic is enabled.
	// Indicator that decryption of TLS traffic is enabled.
	TLSDecryptEnabled *bool `json:"tlsDecryptEnabled,omitempty" tf:"tls_decrypt_enabled,omitempty"`

	// (Boolean) Safely browse websites in Browser Isolation through a URL. Defaults to false.
	// Safely browse websites in Browser Isolation through a URL. Defaults to `false`.
	URLBrowserIsolationEnabled *bool `json:"urlBrowserIsolationEnabled,omitempty" tf:"url_browser_isolation_enabled,omitempty"`
}

type AccountObservation struct {

	// (String) The account identifier to target for the resource.
	// The account identifier to target for the resource.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Boolean) Whether to enable the activity log.
	// Whether to enable the activity log.
	ActivityLogEnabled *bool `json:"activityLogEnabled,omitempty" tf:"activity_log_enabled,omitempty"`

	// (Block List, Max: 1) Configuration block for antivirus traffic scanning. (see below for nested schema)
	// Configuration block for antivirus traffic scanning.
	Antivirus []AntivirusObservation `json:"antivirus,omitempty" tf:"antivirus,omitempty"`

	// (Block List, Max: 1) Configuration for a custom block page. (see below for nested schema)
	// Configuration for a custom block page.
	BlockPage []BlockPageObservation `json:"blockPage,omitempty" tf:"block_page,omitempty"`

	// (Block List, Max: 1) Configuration for body scanning. (see below for nested schema)
	// Configuration for body scanning.
	BodyScanning []BodyScanningObservation `json:"bodyScanning,omitempty" tf:"body_scanning,omitempty"`

	// PKI. (see below for nested schema)
	// Configuration for custom certificates / BYO-PKI.
	CustomCertificate []CustomCertificateObservation `json:"customCertificate,omitempty" tf:"custom_certificate,omitempty"`

	// mail matching. (see below for nested schema)
	// Configuration for extended e-mail matching.
	ExtendedEmailMatching []ExtendedEmailMatchingObservation `json:"extendedEmailMatching,omitempty" tf:"extended_email_matching,omitempty"`

	// (Block List, Max: 1) Configure compliance with Federal Information Processing Standards. (see below for nested schema)
	// Configure compliance with Federal Information Processing Standards.
	Fips []FipsObservation `json:"fips,omitempty" tf:"fips,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	Logging []LoggingObservation `json:"logging,omitempty" tf:"logging,omitempty"`

	// identity onramp for Browser Isolation. Defaults to false.
	// Enable non-identity onramp for Browser Isolation. Defaults to `false`.
	NonIdentityBrowserIsolationEnabled *bool `json:"nonIdentityBrowserIsolationEnabled,omitempty" tf:"non_identity_browser_isolation_enabled,omitempty"`

	// (Block List, Max: 1) Configuration for DLP Payload Logging. (see below for nested schema)
	// Configuration for DLP Payload Logging.
	PayloadLog []PayloadLogObservation `json:"payloadLog,omitempty" tf:"payload_log,omitempty"`

	// (Boolean) Indicator that protocol detection is enabled.
	// Indicator that protocol detection is enabled.
	ProtocolDetectionEnabled *bool `json:"protocolDetectionEnabled,omitempty" tf:"protocol_detection_enabled,omitempty"`

	// (Block List, Max: 1) Configuration block for specifying which protocols are proxied. (see below for nested schema)
	// Configuration block for specifying which protocols are proxied.
	Proxy []ProxyObservation `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// (Block List, Max: 1) Configuration for SSH Session Logging. (see below for nested schema)
	// Configuration for SSH Session Logging.
	SSHSessionLog []SSHSessionLogObservation `json:"sshSessionLog,omitempty" tf:"ssh_session_log,omitempty"`

	// (Boolean) Indicator that decryption of TLS traffic is enabled.
	// Indicator that decryption of TLS traffic is enabled.
	TLSDecryptEnabled *bool `json:"tlsDecryptEnabled,omitempty" tf:"tls_decrypt_enabled,omitempty"`

	// (Boolean) Safely browse websites in Browser Isolation through a URL. Defaults to false.
	// Safely browse websites in Browser Isolation through a URL. Defaults to `false`.
	URLBrowserIsolationEnabled *bool `json:"urlBrowserIsolationEnabled,omitempty" tf:"url_browser_isolation_enabled,omitempty"`
}

type AccountParameters struct {

	// (String) The account identifier to target for the resource.
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

	// (Boolean) Whether to enable the activity log.
	// Whether to enable the activity log.
	// +kubebuilder:validation:Optional
	ActivityLogEnabled *bool `json:"activityLogEnabled,omitempty" tf:"activity_log_enabled,omitempty"`

	// (Block List, Max: 1) Configuration block for antivirus traffic scanning. (see below for nested schema)
	// Configuration block for antivirus traffic scanning.
	// +kubebuilder:validation:Optional
	Antivirus []AntivirusParameters `json:"antivirus,omitempty" tf:"antivirus,omitempty"`

	// (Block List, Max: 1) Configuration for a custom block page. (see below for nested schema)
	// Configuration for a custom block page.
	// +kubebuilder:validation:Optional
	BlockPage []BlockPageParameters `json:"blockPage,omitempty" tf:"block_page,omitempty"`

	// (Block List, Max: 1) Configuration for body scanning. (see below for nested schema)
	// Configuration for body scanning.
	// +kubebuilder:validation:Optional
	BodyScanning []BodyScanningParameters `json:"bodyScanning,omitempty" tf:"body_scanning,omitempty"`

	// PKI. (see below for nested schema)
	// Configuration for custom certificates / BYO-PKI.
	// +kubebuilder:validation:Optional
	CustomCertificate []CustomCertificateParameters `json:"customCertificate,omitempty" tf:"custom_certificate,omitempty"`

	// mail matching. (see below for nested schema)
	// Configuration for extended e-mail matching.
	// +kubebuilder:validation:Optional
	ExtendedEmailMatching []ExtendedEmailMatchingParameters `json:"extendedEmailMatching,omitempty" tf:"extended_email_matching,omitempty"`

	// (Block List, Max: 1) Configure compliance with Federal Information Processing Standards. (see below for nested schema)
	// Configure compliance with Federal Information Processing Standards.
	// +kubebuilder:validation:Optional
	Fips []FipsParameters `json:"fips,omitempty" tf:"fips,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// identity onramp for Browser Isolation. Defaults to false.
	// Enable non-identity onramp for Browser Isolation. Defaults to `false`.
	// +kubebuilder:validation:Optional
	NonIdentityBrowserIsolationEnabled *bool `json:"nonIdentityBrowserIsolationEnabled,omitempty" tf:"non_identity_browser_isolation_enabled,omitempty"`

	// (Block List, Max: 1) Configuration for DLP Payload Logging. (see below for nested schema)
	// Configuration for DLP Payload Logging.
	// +kubebuilder:validation:Optional
	PayloadLog []PayloadLogParameters `json:"payloadLog,omitempty" tf:"payload_log,omitempty"`

	// (Boolean) Indicator that protocol detection is enabled.
	// Indicator that protocol detection is enabled.
	// +kubebuilder:validation:Optional
	ProtocolDetectionEnabled *bool `json:"protocolDetectionEnabled,omitempty" tf:"protocol_detection_enabled,omitempty"`

	// (Block List, Max: 1) Configuration block for specifying which protocols are proxied. (see below for nested schema)
	// Configuration block for specifying which protocols are proxied.
	// +kubebuilder:validation:Optional
	Proxy []ProxyParameters `json:"proxy,omitempty" tf:"proxy,omitempty"`

	// (Block List, Max: 1) Configuration for SSH Session Logging. (see below for nested schema)
	// Configuration for SSH Session Logging.
	// +kubebuilder:validation:Optional
	SSHSessionLog []SSHSessionLogParameters `json:"sshSessionLog,omitempty" tf:"ssh_session_log,omitempty"`

	// (Boolean) Indicator that decryption of TLS traffic is enabled.
	// Indicator that decryption of TLS traffic is enabled.
	// +kubebuilder:validation:Optional
	TLSDecryptEnabled *bool `json:"tlsDecryptEnabled,omitempty" tf:"tls_decrypt_enabled,omitempty"`

	// (Boolean) Safely browse websites in Browser Isolation through a URL. Defaults to false.
	// Safely browse websites in Browser Isolation through a URL. Defaults to `false`.
	// +kubebuilder:validation:Optional
	URLBrowserIsolationEnabled *bool `json:"urlBrowserIsolationEnabled,omitempty" tf:"url_browser_isolation_enabled,omitempty"`
}

type AntivirusInitParameters struct {

	// (Boolean) Scan on file download.
	// Scan on file download.
	EnabledDownloadPhase *bool `json:"enabledDownloadPhase,omitempty" tf:"enabled_download_phase,omitempty"`

	// (Boolean) Scan on file upload.
	// Scan on file upload.
	EnabledUploadPhase *bool `json:"enabledUploadPhase,omitempty" tf:"enabled_upload_phase,omitempty"`

	// (Boolean) Block requests for files that cannot be scanned.
	// Block requests for files that cannot be scanned.
	FailClosed *bool `json:"failClosed,omitempty" tf:"fail_closed,omitempty"`

	// (Block List, Max: 1) Set notifications for antivirus. (see below for nested schema)
	// Set notifications for antivirus.
	NotificationSettings []NotificationSettingsInitParameters `json:"notificationSettings,omitempty" tf:"notification_settings,omitempty"`
}

type AntivirusObservation struct {

	// (Boolean) Scan on file download.
	// Scan on file download.
	EnabledDownloadPhase *bool `json:"enabledDownloadPhase,omitempty" tf:"enabled_download_phase,omitempty"`

	// (Boolean) Scan on file upload.
	// Scan on file upload.
	EnabledUploadPhase *bool `json:"enabledUploadPhase,omitempty" tf:"enabled_upload_phase,omitempty"`

	// (Boolean) Block requests for files that cannot be scanned.
	// Block requests for files that cannot be scanned.
	FailClosed *bool `json:"failClosed,omitempty" tf:"fail_closed,omitempty"`

	// (Block List, Max: 1) Set notifications for antivirus. (see below for nested schema)
	// Set notifications for antivirus.
	NotificationSettings []NotificationSettingsObservation `json:"notificationSettings,omitempty" tf:"notification_settings,omitempty"`
}

type AntivirusParameters struct {

	// (Boolean) Scan on file download.
	// Scan on file download.
	// +kubebuilder:validation:Optional
	EnabledDownloadPhase *bool `json:"enabledDownloadPhase" tf:"enabled_download_phase,omitempty"`

	// (Boolean) Scan on file upload.
	// Scan on file upload.
	// +kubebuilder:validation:Optional
	EnabledUploadPhase *bool `json:"enabledUploadPhase" tf:"enabled_upload_phase,omitempty"`

	// (Boolean) Block requests for files that cannot be scanned.
	// Block requests for files that cannot be scanned.
	// +kubebuilder:validation:Optional
	FailClosed *bool `json:"failClosed" tf:"fail_closed,omitempty"`

	// (Block List, Max: 1) Set notifications for antivirus. (see below for nested schema)
	// Set notifications for antivirus.
	// +kubebuilder:validation:Optional
	NotificationSettings []NotificationSettingsParameters `json:"notificationSettings,omitempty" tf:"notification_settings,omitempty"`
}

type BlockPageInitParameters struct {

	// (String) Hex code of block page background color.
	// Hex code of block page background color.
	BackgroundColor *string `json:"backgroundColor,omitempty" tf:"background_color,omitempty"`

	// (Boolean) Indicator of enablement.
	// Indicator of enablement.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Block page footer text.
	// Block page footer text.
	FooterText *string `json:"footerText,omitempty" tf:"footer_text,omitempty"`

	// (String) Block page header text.
	// Block page header text.
	HeaderText *string `json:"headerText,omitempty" tf:"header_text,omitempty"`

	// (String) URL of block page logo.
	// URL of block page logo.
	LogoPath *string `json:"logoPath,omitempty" tf:"logo_path,omitempty"`

	// (String) Admin email for users to contact.
	// Admin email for users to contact.
	MailtoAddress *string `json:"mailtoAddress,omitempty" tf:"mailto_address,omitempty"`

	// (String) Subject line for emails created from block page.
	// Subject line for emails created from block page.
	MailtoSubject *string `json:"mailtoSubject,omitempty" tf:"mailto_subject,omitempty"`

	// (String) Name of block page configuration.
	// Name of block page configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BlockPageObservation struct {

	// (String) Hex code of block page background color.
	// Hex code of block page background color.
	BackgroundColor *string `json:"backgroundColor,omitempty" tf:"background_color,omitempty"`

	// (Boolean) Indicator of enablement.
	// Indicator of enablement.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Block page footer text.
	// Block page footer text.
	FooterText *string `json:"footerText,omitempty" tf:"footer_text,omitempty"`

	// (String) Block page header text.
	// Block page header text.
	HeaderText *string `json:"headerText,omitempty" tf:"header_text,omitempty"`

	// (String) URL of block page logo.
	// URL of block page logo.
	LogoPath *string `json:"logoPath,omitempty" tf:"logo_path,omitempty"`

	// (String) Admin email for users to contact.
	// Admin email for users to contact.
	MailtoAddress *string `json:"mailtoAddress,omitempty" tf:"mailto_address,omitempty"`

	// (String) Subject line for emails created from block page.
	// Subject line for emails created from block page.
	MailtoSubject *string `json:"mailtoSubject,omitempty" tf:"mailto_subject,omitempty"`

	// (String) Name of block page configuration.
	// Name of block page configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BlockPageParameters struct {

	// (String) Hex code of block page background color.
	// Hex code of block page background color.
	// +kubebuilder:validation:Optional
	BackgroundColor *string `json:"backgroundColor,omitempty" tf:"background_color,omitempty"`

	// (Boolean) Indicator of enablement.
	// Indicator of enablement.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Block page footer text.
	// Block page footer text.
	// +kubebuilder:validation:Optional
	FooterText *string `json:"footerText,omitempty" tf:"footer_text,omitempty"`

	// (String) Block page header text.
	// Block page header text.
	// +kubebuilder:validation:Optional
	HeaderText *string `json:"headerText,omitempty" tf:"header_text,omitempty"`

	// (String) URL of block page logo.
	// URL of block page logo.
	// +kubebuilder:validation:Optional
	LogoPath *string `json:"logoPath,omitempty" tf:"logo_path,omitempty"`

	// (String) Admin email for users to contact.
	// Admin email for users to contact.
	// +kubebuilder:validation:Optional
	MailtoAddress *string `json:"mailtoAddress,omitempty" tf:"mailto_address,omitempty"`

	// (String) Subject line for emails created from block page.
	// Subject line for emails created from block page.
	// +kubebuilder:validation:Optional
	MailtoSubject *string `json:"mailtoSubject,omitempty" tf:"mailto_subject,omitempty"`

	// (String) Name of block page configuration.
	// Name of block page configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BodyScanningInitParameters struct {

	// (String) Body scanning inspection mode. Available values: deep, shallow.
	// Body scanning inspection mode. Available values: `deep`, `shallow`.
	InspectionMode *string `json:"inspectionMode,omitempty" tf:"inspection_mode,omitempty"`
}

type BodyScanningObservation struct {

	// (String) Body scanning inspection mode. Available values: deep, shallow.
	// Body scanning inspection mode. Available values: `deep`, `shallow`.
	InspectionMode *string `json:"inspectionMode,omitempty" tf:"inspection_mode,omitempty"`
}

type BodyScanningParameters struct {

	// (String) Body scanning inspection mode. Available values: deep, shallow.
	// Body scanning inspection mode. Available values: `deep`, `shallow`.
	// +kubebuilder:validation:Optional
	InspectionMode *string `json:"inspectionMode" tf:"inspection_mode,omitempty"`
}

type CustomCertificateInitParameters struct {

	// (Boolean) Enable notification settings.
	// Whether TLS encryption should use a custom certificate.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// ID of custom certificate.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CustomCertificateObservation struct {

	// (Boolean) Enable notification settings.
	// Whether TLS encryption should use a custom certificate.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// ID of custom certificate.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type CustomCertificateParameters struct {

	// (Boolean) Enable notification settings.
	// Whether TLS encryption should use a custom certificate.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// ID of custom certificate.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DNSInitParameters struct {

	// (Boolean) Whether to log all activity.
	// Whether to log all activity.
	LogAll *bool `json:"logAll,omitempty" tf:"log_all,omitempty"`

	// (Boolean)
	LogBlocks *bool `json:"logBlocks,omitempty" tf:"log_blocks,omitempty"`
}

type DNSObservation struct {

	// (Boolean) Whether to log all activity.
	// Whether to log all activity.
	LogAll *bool `json:"logAll,omitempty" tf:"log_all,omitempty"`

	// (Boolean)
	LogBlocks *bool `json:"logBlocks,omitempty" tf:"log_blocks,omitempty"`
}

type DNSParameters struct {

	// (Boolean) Whether to log all activity.
	// Whether to log all activity.
	// +kubebuilder:validation:Optional
	LogAll *bool `json:"logAll" tf:"log_all,omitempty"`

	// (Boolean)
	// +kubebuilder:validation:Optional
	LogBlocks *bool `json:"logBlocks" tf:"log_blocks,omitempty"`
}

type ExtendedEmailMatchingInitParameters struct {

	// (Boolean) Enable notification settings.
	// Whether e-mails should be matched on all variants of user emails (with + or . modifiers) in Firewall policies.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ExtendedEmailMatchingObservation struct {

	// (Boolean) Enable notification settings.
	// Whether e-mails should be matched on all variants of user emails (with + or . modifiers) in Firewall policies.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ExtendedEmailMatchingParameters struct {

	// (Boolean) Enable notification settings.
	// Whether e-mails should be matched on all variants of user emails (with + or . modifiers) in Firewall policies.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type FipsInitParameters struct {

	// compliant TLS configuration.
	// Only allow FIPS-compliant TLS configuration.
	TLS *bool `json:"tls,omitempty" tf:"tls,omitempty"`
}

type FipsObservation struct {

	// compliant TLS configuration.
	// Only allow FIPS-compliant TLS configuration.
	TLS *bool `json:"tls,omitempty" tf:"tls,omitempty"`
}

type FipsParameters struct {

	// compliant TLS configuration.
	// Only allow FIPS-compliant TLS configuration.
	// +kubebuilder:validation:Optional
	TLS *bool `json:"tls,omitempty" tf:"tls,omitempty"`
}

type HTTPInitParameters struct {

	// (Boolean) Whether to log all activity.
	// Whether to log all activity.
	LogAll *bool `json:"logAll,omitempty" tf:"log_all,omitempty"`

	// (Boolean)
	LogBlocks *bool `json:"logBlocks,omitempty" tf:"log_blocks,omitempty"`
}

type HTTPObservation struct {

	// (Boolean) Whether to log all activity.
	// Whether to log all activity.
	LogAll *bool `json:"logAll,omitempty" tf:"log_all,omitempty"`

	// (Boolean)
	LogBlocks *bool `json:"logBlocks,omitempty" tf:"log_blocks,omitempty"`
}

type HTTPParameters struct {

	// (Boolean) Whether to log all activity.
	// Whether to log all activity.
	// +kubebuilder:validation:Optional
	LogAll *bool `json:"logAll" tf:"log_all,omitempty"`

	// (Boolean)
	// +kubebuilder:validation:Optional
	LogBlocks *bool `json:"logBlocks" tf:"log_blocks,omitempty"`
}

type L4InitParameters struct {

	// (Boolean) Whether to log all activity.
	// Whether to log all activity.
	LogAll *bool `json:"logAll,omitempty" tf:"log_all,omitempty"`

	// (Boolean)
	LogBlocks *bool `json:"logBlocks,omitempty" tf:"log_blocks,omitempty"`
}

type L4Observation struct {

	// (Boolean) Whether to log all activity.
	// Whether to log all activity.
	LogAll *bool `json:"logAll,omitempty" tf:"log_all,omitempty"`

	// (Boolean)
	LogBlocks *bool `json:"logBlocks,omitempty" tf:"log_blocks,omitempty"`
}

type L4Parameters struct {

	// (Boolean) Whether to log all activity.
	// Whether to log all activity.
	// +kubebuilder:validation:Optional
	LogAll *bool `json:"logAll" tf:"log_all,omitempty"`

	// (Boolean)
	// +kubebuilder:validation:Optional
	LogBlocks *bool `json:"logBlocks" tf:"log_blocks,omitempty"`
}

type LoggingInitParameters struct {

	// (Boolean) Redact personally identifiable information from activity logging (PII fields are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	// Redact personally identifiable information from activity logging (PII fields are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii *bool `json:"redactPii,omitempty" tf:"redact_pii,omitempty"`

	// (Block List, Min: 1, Max: 1) Represents whether all requests are logged or only the blocked requests are slogged in DNS, HTTP and L4 filters. (see below for nested schema)
	// Represents whether all requests are logged or only the blocked requests are slogged in DNS, HTTP and L4 filters.
	SettingsByRuleType []SettingsByRuleTypeInitParameters `json:"settingsByRuleType,omitempty" tf:"settings_by_rule_type,omitempty"`
}

type LoggingObservation struct {

	// (Boolean) Redact personally identifiable information from activity logging (PII fields are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	// Redact personally identifiable information from activity logging (PII fields are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii *bool `json:"redactPii,omitempty" tf:"redact_pii,omitempty"`

	// (Block List, Min: 1, Max: 1) Represents whether all requests are logged or only the blocked requests are slogged in DNS, HTTP and L4 filters. (see below for nested schema)
	// Represents whether all requests are logged or only the blocked requests are slogged in DNS, HTTP and L4 filters.
	SettingsByRuleType []SettingsByRuleTypeObservation `json:"settingsByRuleType,omitempty" tf:"settings_by_rule_type,omitempty"`
}

type LoggingParameters struct {

	// (Boolean) Redact personally identifiable information from activity logging (PII fields are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	// Redact personally identifiable information from activity logging (PII fields are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	// +kubebuilder:validation:Optional
	RedactPii *bool `json:"redactPii" tf:"redact_pii,omitempty"`

	// (Block List, Min: 1, Max: 1) Represents whether all requests are logged or only the blocked requests are slogged in DNS, HTTP and L4 filters. (see below for nested schema)
	// Represents whether all requests are logged or only the blocked requests are slogged in DNS, HTTP and L4 filters.
	// +kubebuilder:validation:Optional
	SettingsByRuleType []SettingsByRuleTypeParameters `json:"settingsByRuleType" tf:"settings_by_rule_type,omitempty"`
}

type NotificationSettingsInitParameters struct {

	// (Boolean) Enable notification settings.
	// Enable notification settings.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Notification content.
	// Notification content.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Support URL to show in the notification.
	// Support URL to show in the notification.
	SupportURL *string `json:"supportUrl,omitempty" tf:"support_url,omitempty"`
}

type NotificationSettingsObservation struct {

	// (Boolean) Enable notification settings.
	// Enable notification settings.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Notification content.
	// Notification content.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Support URL to show in the notification.
	// Support URL to show in the notification.
	SupportURL *string `json:"supportUrl,omitempty" tf:"support_url,omitempty"`
}

type NotificationSettingsParameters struct {

	// (Boolean) Enable notification settings.
	// Enable notification settings.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Notification content.
	// Notification content.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Support URL to show in the notification.
	// Support URL to show in the notification.
	// +kubebuilder:validation:Optional
	SupportURL *string `json:"supportUrl,omitempty" tf:"support_url,omitempty"`
}

type PayloadLogInitParameters struct {

	// (String) Public key used to encrypt matched payloads.
	// Public key used to encrypt matched payloads.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type PayloadLogObservation struct {

	// (String) Public key used to encrypt matched payloads.
	// Public key used to encrypt matched payloads.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type PayloadLogParameters struct {

	// (String) Public key used to encrypt matched payloads.
	// Public key used to encrypt matched payloads.
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey" tf:"public_key,omitempty"`
}

type ProxyInitParameters struct {

	// (Boolean) Whether root ca is enabled account wide for ZT clients.
	// Whether root ca is enabled account wide for ZT clients.
	RootCA *bool `json:"rootCa,omitempty" tf:"root_ca,omitempty"`

	// (Boolean) Whether gateway proxy is enabled on gateway devices for TCP traffic.
	// Whether gateway proxy is enabled on gateway devices for TCP traffic.
	TCP *bool `json:"tcp,omitempty" tf:"tcp,omitempty"`

	// (Boolean) Whether gateway proxy is enabled on gateway devices for UDP traffic.
	// Whether gateway proxy is enabled on gateway devices for UDP traffic.
	UDP *bool `json:"udp,omitempty" tf:"udp,omitempty"`
}

type ProxyObservation struct {

	// (Boolean) Whether root ca is enabled account wide for ZT clients.
	// Whether root ca is enabled account wide for ZT clients.
	RootCA *bool `json:"rootCa,omitempty" tf:"root_ca,omitempty"`

	// (Boolean) Whether gateway proxy is enabled on gateway devices for TCP traffic.
	// Whether gateway proxy is enabled on gateway devices for TCP traffic.
	TCP *bool `json:"tcp,omitempty" tf:"tcp,omitempty"`

	// (Boolean) Whether gateway proxy is enabled on gateway devices for UDP traffic.
	// Whether gateway proxy is enabled on gateway devices for UDP traffic.
	UDP *bool `json:"udp,omitempty" tf:"udp,omitempty"`
}

type ProxyParameters struct {

	// (Boolean) Whether root ca is enabled account wide for ZT clients.
	// Whether root ca is enabled account wide for ZT clients.
	// +kubebuilder:validation:Optional
	RootCA *bool `json:"rootCa" tf:"root_ca,omitempty"`

	// (Boolean) Whether gateway proxy is enabled on gateway devices for TCP traffic.
	// Whether gateway proxy is enabled on gateway devices for TCP traffic.
	// +kubebuilder:validation:Optional
	TCP *bool `json:"tcp" tf:"tcp,omitempty"`

	// (Boolean) Whether gateway proxy is enabled on gateway devices for UDP traffic.
	// Whether gateway proxy is enabled on gateway devices for UDP traffic.
	// +kubebuilder:validation:Optional
	UDP *bool `json:"udp" tf:"udp,omitempty"`
}

type SSHSessionLogInitParameters struct {

	// (String) Public key used to encrypt matched payloads.
	// Public key used to encrypt ssh session.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type SSHSessionLogObservation struct {

	// (String) Public key used to encrypt matched payloads.
	// Public key used to encrypt ssh session.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type SSHSessionLogParameters struct {

	// (String) Public key used to encrypt matched payloads.
	// Public key used to encrypt ssh session.
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey" tf:"public_key,omitempty"`
}

type SettingsByRuleTypeInitParameters struct {

	// (Block List, Min: 1, Max: 1) Logging configuration for DNS requests. (see below for nested schema)
	// Logging configuration for DNS requests.
	DNS []DNSInitParameters `json:"dns,omitempty" tf:"dns,omitempty"`

	// (Block List, Min: 1, Max: 1) Logging configuration for HTTP requests. (see below for nested schema)
	// Logging configuration for HTTP requests.
	HTTP []HTTPInitParameters `json:"http,omitempty" tf:"http,omitempty"`

	// (Block List, Min: 1, Max: 1) Logging configuration for layer 4 requests. (see below for nested schema)
	// Logging configuration for layer 4 requests.
	L4 []L4InitParameters `json:"l4,omitempty" tf:"l4,omitempty"`
}

type SettingsByRuleTypeObservation struct {

	// (Block List, Min: 1, Max: 1) Logging configuration for DNS requests. (see below for nested schema)
	// Logging configuration for DNS requests.
	DNS []DNSObservation `json:"dns,omitempty" tf:"dns,omitempty"`

	// (Block List, Min: 1, Max: 1) Logging configuration for HTTP requests. (see below for nested schema)
	// Logging configuration for HTTP requests.
	HTTP []HTTPObservation `json:"http,omitempty" tf:"http,omitempty"`

	// (Block List, Min: 1, Max: 1) Logging configuration for layer 4 requests. (see below for nested schema)
	// Logging configuration for layer 4 requests.
	L4 []L4Observation `json:"l4,omitempty" tf:"l4,omitempty"`
}

type SettingsByRuleTypeParameters struct {

	// (Block List, Min: 1, Max: 1) Logging configuration for DNS requests. (see below for nested schema)
	// Logging configuration for DNS requests.
	// +kubebuilder:validation:Optional
	DNS []DNSParameters `json:"dns" tf:"dns,omitempty"`

	// (Block List, Min: 1, Max: 1) Logging configuration for HTTP requests. (see below for nested schema)
	// Logging configuration for HTTP requests.
	// +kubebuilder:validation:Optional
	HTTP []HTTPParameters `json:"http" tf:"http,omitempty"`

	// (Block List, Min: 1, Max: 1) Logging configuration for layer 4 requests. (see below for nested schema)
	// Logging configuration for layer 4 requests.
	// +kubebuilder:validation:Optional
	L4 []L4Parameters `json:"l4" tf:"l4,omitempty"`
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

// Account is the Schema for the Accounts API. Provides a Cloudflare Teams Account resource. The Teams Account resource defines configuration for secure web gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountSpec   `json:"spec"`
	Status            AccountStatus `json:"status,omitempty"`
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
