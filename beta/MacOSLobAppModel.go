// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MacOSLobApp Contains properties and inherited properties for the MacOS LOB App.
type MacOSLobApp struct {
	// MobileLobApp is the base model of MacOSLobApp
	MobileLobApp
	// BundleID The bundle id.
	BundleID *string `json:"bundleId,omitempty"`
	// MinimumSupportedOperatingSystem The value for the minimum applicable operating system.
	MinimumSupportedOperatingSystem *MacOSMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
	// BuildNumber The build number of MacOS Line of Business (LoB) app.
	BuildNumber *string `json:"buildNumber,omitempty"`
	// VersionNumber The version number of MacOS Line of Business (LoB) app.
	VersionNumber *string `json:"versionNumber,omitempty"`
	// ChildApps The app list in this bundle package
	ChildApps []MacOSLobChildApp `json:"childApps,omitempty"`
	// IdentityVersion The identity version.
	IdentityVersion *string `json:"identityVersion,omitempty"`
	// Md5HashChunkSize The chunk size for MD5 hash
	Md5HashChunkSize *int `json:"md5HashChunkSize,omitempty"`
	// Md5Hash The MD5 hash codes
	Md5Hash []string `json:"md5Hash,omitempty"`
	// IgnoreVersionDetection A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for macOS Line of Business (LoB) apps that use a self update feature.
	IgnoreVersionDetection *bool `json:"ignoreVersionDetection,omitempty"`
}
