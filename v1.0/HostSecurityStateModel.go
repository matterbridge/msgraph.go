// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// HostSecurityState undocumented
type HostSecurityState struct {
	// Object is the base model of HostSecurityState
	Object
	// Fqdn undocumented
	Fqdn *string `json:"fqdn,omitempty"`
	// IsAzureAdJoined undocumented
	IsAzureAdJoined *bool `json:"isAzureAdJoined,omitempty"`
	// IsAzureAdRegistered undocumented
	IsAzureAdRegistered *bool `json:"isAzureAdRegistered,omitempty"`
	// IsHybridAzureDomainJoined undocumented
	IsHybridAzureDomainJoined *bool `json:"isHybridAzureDomainJoined,omitempty"`
	// NetBiosName undocumented
	NetBiosName *string `json:"netBiosName,omitempty"`
	// Os undocumented
	Os *string `json:"os,omitempty"`
	// PrivateIPAddress undocumented
	PrivateIPAddress *string `json:"privateIpAddress,omitempty"`
	// PublicIPAddress undocumented
	PublicIPAddress *string `json:"publicIpAddress,omitempty"`
	// RiskScore undocumented
	RiskScore *string `json:"riskScore,omitempty"`
}
