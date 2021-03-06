// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// FileSecurityProfile undocumented
type FileSecurityProfile struct {
	// Entity is the base model of FileSecurityProfile
	Entity
	// ActivityGroupNames undocumented
	ActivityGroupNames []string `json:"activityGroupNames,omitempty"`
	// AzureSubscriptionID undocumented
	AzureSubscriptionID *string `json:"azureSubscriptionId,omitempty"`
	// AzureTenantID undocumented
	AzureTenantID *string `json:"azureTenantId,omitempty"`
	// CertificateThumbprint undocumented
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty"`
	// Extensions undocumented
	Extensions []string `json:"extensions,omitempty"`
	// FileType undocumented
	FileType *string `json:"fileType,omitempty"`
	// FirstSeenDateTime undocumented
	FirstSeenDateTime *time.Time `json:"firstSeenDateTime,omitempty"`
	// Hashes undocumented
	Hashes []FileHash `json:"hashes,omitempty"`
	// LastSeenDateTime undocumented
	LastSeenDateTime *time.Time `json:"lastSeenDateTime,omitempty"`
	// MalwareStates undocumented
	MalwareStates []MalwareState `json:"malwareStates,omitempty"`
	// Names undocumented
	Names []string `json:"names,omitempty"`
	// RiskScore undocumented
	RiskScore *string `json:"riskScore,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
	// Tags undocumented
	Tags []string `json:"tags,omitempty"`
	// VendorInformation undocumented
	VendorInformation *SecurityVendorInformation `json:"vendorInformation,omitempty"`
	// VulnerabilityStates undocumented
	VulnerabilityStates []VulnerabilityState `json:"vulnerabilityStates,omitempty"`
}
