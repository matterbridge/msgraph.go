// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ManagedAllDeviceCertificateState undocumented
type ManagedAllDeviceCertificateState struct {
	// Entity is the base model of ManagedAllDeviceCertificateState
	Entity
	// CertificateRevokeStatus Revoke status
	CertificateRevokeStatus *CertificateRevocationStatus `json:"certificateRevokeStatus,omitempty"`
	// ManagedDeviceDisplayName Device display name
	ManagedDeviceDisplayName *string `json:"managedDeviceDisplayName,omitempty"`
	// UserPrincipalName User principal name
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// CertificateExpirationDateTime Certificate expiry date
	CertificateExpirationDateTime *time.Time `json:"certificateExpirationDateTime,omitempty"`
	// CertificateIssuerName Issuer
	CertificateIssuerName *string `json:"certificateIssuerName,omitempty"`
	// CertificateThumbprint Thumbprint
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty"`
	// CertificateSerialNumber Serial number
	CertificateSerialNumber *string `json:"certificateSerialNumber,omitempty"`
	// CertificateSubjectName Certificate subject name
	CertificateSubjectName *string `json:"certificateSubjectName,omitempty"`
	// CertificateKeyUsages Key Usage
	CertificateKeyUsages *int `json:"certificateKeyUsages,omitempty"`
	// CertificateExtendedKeyUsages Enhanced Key Usage
	CertificateExtendedKeyUsages *string `json:"certificateExtendedKeyUsages,omitempty"`
	// CertificateIssuanceDateTime Issuance date
	CertificateIssuanceDateTime *time.Time `json:"certificateIssuanceDateTime,omitempty"`
}
