// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SensitivityPolicySettings undocumented
type SensitivityPolicySettings struct {
	// Entity is the base model of SensitivityPolicySettings
	Entity
	// IsMandatory undocumented
	IsMandatory *bool `json:"isMandatory,omitempty"`
	// HelpWebURL undocumented
	HelpWebURL *string `json:"helpWebUrl,omitempty"`
	// DowngradeSensitivityRequiresJustification undocumented
	DowngradeSensitivityRequiresJustification *bool `json:"downgradeSensitivityRequiresJustification,omitempty"`
}
