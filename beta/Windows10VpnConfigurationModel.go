// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Windows10VpnConfiguration By providing the configurations in this profile you can instruct the Windows 10 device (desktop or mobile) to connect to desired VPN endpoint. By specifying the authentication method and security types expected by VPN endpoint you can make the VPN connection seamless for end user.
type Windows10VpnConfiguration struct {
	// WindowsVpnConfiguration is the base model of Windows10VpnConfiguration
	WindowsVpnConfiguration
	// ProfileTarget Profile target type.
	ProfileTarget *Windows10VpnProfileTarget `json:"profileTarget,omitempty"`
	// ConnectionType Connection type.
	ConnectionType *Windows10VpnConnectionType `json:"connectionType,omitempty"`
	// EnableSplitTunneling Enable split tunneling.
	EnableSplitTunneling *bool `json:"enableSplitTunneling,omitempty"`
	// EnableAlwaysOn Enable Always On mode.
	EnableAlwaysOn *bool `json:"enableAlwaysOn,omitempty"`
	// EnableDeviceTunnel Enable device tunnel.
	EnableDeviceTunnel *bool `json:"enableDeviceTunnel,omitempty"`
	// EnableDNSRegistration Enable IP address registration with internal DNS.
	EnableDNSRegistration *bool `json:"enableDnsRegistration,omitempty"`
	// DNSSuffixes Specify DNS suffixes to add to the DNS search list to properly route short names.
	DNSSuffixes []string `json:"dnsSuffixes,omitempty"`
	// AuthenticationMethod Authentication method.
	AuthenticationMethod *Windows10VpnAuthenticationMethod `json:"authenticationMethod,omitempty"`
	// RememberUserCredentials Remember user credentials.
	RememberUserCredentials *bool `json:"rememberUserCredentials,omitempty"`
	// EnableConditionalAccess Enable conditional access.
	EnableConditionalAccess *bool `json:"enableConditionalAccess,omitempty"`
	// EnableSingleSignOnWithAlternateCertificate Enable single sign-on (SSO) with alternate certificate.
	EnableSingleSignOnWithAlternateCertificate *bool `json:"enableSingleSignOnWithAlternateCertificate,omitempty"`
	// SingleSignOnEku Single sign-on Extended Key Usage (EKU).
	SingleSignOnEku *ExtendedKeyUsage `json:"singleSignOnEku,omitempty"`
	// SingleSignOnIssuerHash Single sign-on issuer hash.
	SingleSignOnIssuerHash *string `json:"singleSignOnIssuerHash,omitempty"`
	// EapXML Extensible Authentication Protocol (EAP) XML. (UTF8 encoded byte array)
	EapXML *Binary `json:"eapXml,omitempty"`
	// ProxyServer Proxy Server.
	ProxyServer *Windows10VpnProxyServer `json:"proxyServer,omitempty"`
	// AssociatedApps Associated Apps. This collection can contain a maximum of 10000 elements.
	AssociatedApps []Windows10AssociatedApps `json:"associatedApps,omitempty"`
	// OnlyAssociatedAppsCanUseConnection Only associated Apps can use connection (per-app VPN).
	OnlyAssociatedAppsCanUseConnection *bool `json:"onlyAssociatedAppsCanUseConnection,omitempty"`
	// WindowsInformationProtectionDomain Windows Information Protection (WIP) domain to associate with this connection.
	WindowsInformationProtectionDomain *string `json:"windowsInformationProtectionDomain,omitempty"`
	// TrafficRules Traffic rules. This collection can contain a maximum of 1000 elements.
	TrafficRules []VpnTrafficRule `json:"trafficRules,omitempty"`
	// Routes Routes (optional for third-party providers). This collection can contain a maximum of 1000 elements.
	Routes []VpnRoute `json:"routes,omitempty"`
	// DNSRules DNS rules. This collection can contain a maximum of 1000 elements.
	DNSRules []VpnDNSRule `json:"dnsRules,omitempty"`
	// TrustedNetworkDomains Trusted Network Domains
	TrustedNetworkDomains []string `json:"trustedNetworkDomains,omitempty"`
	// IdentityCertificate undocumented
	IdentityCertificate *WindowsCertificateProfileBase `json:"identityCertificate,omitempty"`
}
