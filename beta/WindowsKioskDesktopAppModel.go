// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsKioskDesktopApp undocumented
type WindowsKioskDesktopApp struct {
	// WindowsKioskAppBase is the base model of WindowsKioskDesktopApp
	WindowsKioskAppBase
	// Path Define the path of a desktop app
	Path *string `json:"path,omitempty"`
	// DesktopApplicationID Define the DesktopApplicationID of the app
	DesktopApplicationID *string `json:"desktopApplicationId,omitempty"`
	// DesktopApplicationLinkPath Define the DesktopApplicationLinkPath of the app
	DesktopApplicationLinkPath *string `json:"desktopApplicationLinkPath,omitempty"`
}
