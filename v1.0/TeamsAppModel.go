// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TeamsApp undocumented
type TeamsApp struct {
	// Entity is the base model of TeamsApp
	Entity
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// DistributionMethod undocumented
	DistributionMethod *TeamsAppDistributionMethod `json:"distributionMethod,omitempty"`
	// AppDefinitions undocumented
	AppDefinitions []TeamsAppDefinition `json:"appDefinitions,omitempty"`
}
