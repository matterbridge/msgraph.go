// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceDetail undocumented
type DeviceDetail struct {
	// Object is the base model of DeviceDetail
	Object
	// DeviceID undocumented
	DeviceID *string `json:"deviceId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// OperatingSystem undocumented
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// Browser undocumented
	Browser *string `json:"browser,omitempty"`
	// BrowserID undocumented
	BrowserID *string `json:"browserId,omitempty"`
	// IsCompliant undocumented
	IsCompliant *bool `json:"isCompliant,omitempty"`
	// IsManaged undocumented
	IsManaged *bool `json:"isManaged,omitempty"`
	// TrustType undocumented
	TrustType *string `json:"trustType,omitempty"`
}
