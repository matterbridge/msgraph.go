// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidWorkProfileCustomConfiguration Android Work Profile custom configuration
type AndroidWorkProfileCustomConfiguration struct {
	// DeviceConfiguration is the base model of AndroidWorkProfileCustomConfiguration
	DeviceConfiguration
	// OMASettings OMA settings. This collection can contain a maximum of 500 elements.
	OMASettings []OMASetting `json:"omaSettings,omitempty"`
}
