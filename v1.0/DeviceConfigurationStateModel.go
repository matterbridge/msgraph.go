// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceConfigurationState Device Configuration State for a given device.
type DeviceConfigurationState struct {
	// Entity is the base model of DeviceConfigurationState
	Entity
	// SettingStates undocumented
	SettingStates []DeviceConfigurationSettingState `json:"settingStates,omitempty"`
	// DisplayName The name of the policy for this policyBase
	DisplayName *string `json:"displayName,omitempty"`
	// Version The version of the policy
	Version *int `json:"version,omitempty"`
	// PlatformType Platform type that the policy applies to
	PlatformType *PolicyPlatformType `json:"platformType,omitempty"`
	// State The compliance state of the policy
	State *ComplianceStatus `json:"state,omitempty"`
	// SettingCount Count of how many setting a policy holds
	SettingCount *int `json:"settingCount,omitempty"`
}
