// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsPhone81ImportedPFXCertificateProfile Windows 8.1 Phone and Mobile PFX Import certificate profile
type WindowsPhone81ImportedPFXCertificateProfile struct {
	// WindowsCertificateProfileBase is the base model of WindowsPhone81ImportedPFXCertificateProfile
	WindowsCertificateProfileBase
	// IntendedPurpose undocumented
	IntendedPurpose *IntendedPurpose `json:"intendedPurpose,omitempty"`
	// ManagedDeviceCertificateStates undocumented
	ManagedDeviceCertificateStates []ManagedDeviceCertificateState `json:"managedDeviceCertificateStates,omitempty"`
}
