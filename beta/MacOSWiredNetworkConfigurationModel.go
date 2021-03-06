// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MacOSWiredNetworkConfiguration MacOS wired network configuration profile.
type MacOSWiredNetworkConfiguration struct {
	// DeviceConfiguration is the base model of MacOSWiredNetworkConfiguration
	DeviceConfiguration
	// NetworkName Network Name
	NetworkName *string `json:"networkName,omitempty"`
	// NetworkInterface Network interface.
	NetworkInterface *WiredNetworkInterface `json:"networkInterface,omitempty"`
	// EapType Extensible Authentication Protocol (EAP). Indicates the type of EAP protocol set on the wired network.
	EapType *EapType `json:"eapType,omitempty"`
	// EapFastConfiguration EAP-FAST Configuration Option when EAP-FAST is the selected EAP Type.
	EapFastConfiguration *EapFastConfiguration `json:"eapFastConfiguration,omitempty"`
	// TrustedServerCertificateNames Trusted server certificate names when EAP Type is configured to EAP-TLS/TTLS/FAST or PEAP. This is the common name used in the certificates issued by your trusted certificate authority (CA). If you provide this information, you can bypass the dynamic trust dialog that is displayed on end users devices when they connect to this wired network.
	TrustedServerCertificateNames []string `json:"trustedServerCertificateNames,omitempty"`
	// AuthenticationMethod Authentication Method when EAP Type is configured to PEAP or EAP-TTLS.
	AuthenticationMethod *WiFiAuthenticationMethod `json:"authenticationMethod,omitempty"`
	// NonEapAuthenticationMethodForEapTtls Non-EAP Method for Authentication (Inner Identity) when EAP Type is EAP-TTLS and Authenticationmethod is Username and Password.
	NonEapAuthenticationMethodForEapTtls *NonEapAuthenticationMethodForEapTtlsType `json:"nonEapAuthenticationMethodForEapTtls,omitempty"`
	// EnableOuterIdentityPrivacy Enable identity privacy (Outer Identity) when EAP Type is configured to EAP-TTLS, EAP-FAST or PEAP. This property masks usernames with the text you enter. For example, if you use 'anonymous', each user that authenticates with this wired network using their real username is displayed as 'anonymous'.
	EnableOuterIdentityPrivacy *string `json:"enableOuterIdentityPrivacy,omitempty"`
	// RootCertificateForServerValidation undocumented
	RootCertificateForServerValidation *MacOSTrustedRootCertificate `json:"rootCertificateForServerValidation,omitempty"`
	// IdentityCertificateForClientAuthentication undocumented
	IdentityCertificateForClientAuthentication *MacOSCertificateProfileBase `json:"identityCertificateForClientAuthentication,omitempty"`
}
