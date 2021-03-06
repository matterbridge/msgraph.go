// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedAppStatus Represents app protection and configuration status for the organization.
type ManagedAppStatus struct {
	// Entity is the base model of ManagedAppStatus
	Entity
	// DisplayName Friendly name of the status report.
	DisplayName *string `json:"displayName,omitempty"`
	// Version Version of the entity.
	Version *string `json:"version,omitempty"`
}
