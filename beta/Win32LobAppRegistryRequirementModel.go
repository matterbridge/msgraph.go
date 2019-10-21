// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Win32LobAppRegistryRequirement undocumented
type Win32LobAppRegistryRequirement struct {
	// Win32LobAppRequirement is the base model of Win32LobAppRegistryRequirement
	Win32LobAppRequirement
	// Check32BitOn64System A value indicating whether this registry path is for checking 32-bit app on 64-bit system
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`
	// KeyPath The registry key path to detect Win32 Line of Business (LoB) app
	KeyPath *string `json:"keyPath,omitempty"`
	// ValueName The registry value name
	ValueName *string `json:"valueName,omitempty"`
	// DetectionType The registry data detection type
	DetectionType *Win32LobAppRegistryDetectionType `json:"detectionType,omitempty"`
}
