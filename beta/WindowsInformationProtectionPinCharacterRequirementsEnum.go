// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsInformationProtectionPinCharacterRequirements undocumented
type WindowsInformationProtectionPinCharacterRequirements int

const (
	// WindowsInformationProtectionPinCharacterRequirementsVNotAllow undocumented
	WindowsInformationProtectionPinCharacterRequirementsVNotAllow WindowsInformationProtectionPinCharacterRequirements = 0
	// WindowsInformationProtectionPinCharacterRequirementsVRequireAtLeastOne undocumented
	WindowsInformationProtectionPinCharacterRequirementsVRequireAtLeastOne WindowsInformationProtectionPinCharacterRequirements = 1
	// WindowsInformationProtectionPinCharacterRequirementsVAllow undocumented
	WindowsInformationProtectionPinCharacterRequirementsVAllow WindowsInformationProtectionPinCharacterRequirements = 2
)

// WindowsInformationProtectionPinCharacterRequirementsPNotAllow returns a pointer to WindowsInformationProtectionPinCharacterRequirementsVNotAllow
func WindowsInformationProtectionPinCharacterRequirementsPNotAllow() *WindowsInformationProtectionPinCharacterRequirements {
	v := WindowsInformationProtectionPinCharacterRequirementsVNotAllow
	return &v
}

// WindowsInformationProtectionPinCharacterRequirementsPRequireAtLeastOne returns a pointer to WindowsInformationProtectionPinCharacterRequirementsVRequireAtLeastOne
func WindowsInformationProtectionPinCharacterRequirementsPRequireAtLeastOne() *WindowsInformationProtectionPinCharacterRequirements {
	v := WindowsInformationProtectionPinCharacterRequirementsVRequireAtLeastOne
	return &v
}

// WindowsInformationProtectionPinCharacterRequirementsPAllow returns a pointer to WindowsInformationProtectionPinCharacterRequirementsVAllow
func WindowsInformationProtectionPinCharacterRequirementsPAllow() *WindowsInformationProtectionPinCharacterRequirements {
	v := WindowsInformationProtectionPinCharacterRequirementsVAllow
	return &v
}
