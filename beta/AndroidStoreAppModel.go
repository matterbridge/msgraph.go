// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidStoreApp Contains properties and inherited properties for Android store apps.
type AndroidStoreApp struct {
	// MobileApp is the base model of AndroidStoreApp
	MobileApp
	// PackageID The package identifier.
	PackageID *string `json:"packageId,omitempty"`
	// AppIdentifier The Identity Name.
	AppIdentifier *string `json:"appIdentifier,omitempty"`
	// AppStoreURL The Android app store URL.
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
	// MinimumSupportedOperatingSystem The value for the minimum applicable operating system.
	MinimumSupportedOperatingSystem *AndroidMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
}
