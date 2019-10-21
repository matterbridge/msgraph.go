// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EditionUpgradeConfiguration Windows 10 Edition Upgrade configuration.
type EditionUpgradeConfiguration struct {
	// DeviceConfiguration is the base model of EditionUpgradeConfiguration
	DeviceConfiguration
	// LicenseType Edition Upgrade License Type.
	LicenseType *EditionUpgradeLicenseType `json:"licenseType,omitempty"`
	// TargetEdition Edition Upgrade Target Edition.
	TargetEdition *Windows10EditionType `json:"targetEdition,omitempty"`
	// License Edition Upgrade License File Content.
	License *string `json:"license,omitempty"`
	// ProductKey Edition Upgrade Product Key.
	ProductKey *string `json:"productKey,omitempty"`
}
