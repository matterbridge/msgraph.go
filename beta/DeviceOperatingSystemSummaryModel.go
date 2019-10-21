// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceOperatingSystemSummary undocumented
type DeviceOperatingSystemSummary struct {
	// Object is the base model of DeviceOperatingSystemSummary
	Object
	// AndroidCount Number of android device count.
	AndroidCount *int `json:"androidCount,omitempty"`
	// IOSCount Number of iOS device count.
	IOSCount *int `json:"iosCount,omitempty"`
	// MacOSCount Number of Mac OS X device count.
	MacOSCount *int `json:"macOSCount,omitempty"`
	// WindowsMobileCount Number of Windows mobile device count.
	WindowsMobileCount *int `json:"windowsMobileCount,omitempty"`
	// WindowsCount Number of Windows device count.
	WindowsCount *int `json:"windowsCount,omitempty"`
	// UnknownCount Number of unknown device count.
	UnknownCount *int `json:"unknownCount,omitempty"`
}
