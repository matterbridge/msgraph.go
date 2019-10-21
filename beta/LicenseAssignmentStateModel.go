// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// LicenseAssignmentState undocumented
type LicenseAssignmentState struct {
	// Object is the base model of LicenseAssignmentState
	Object
	// SKUID undocumented
	SKUID *UUID `json:"skuId,omitempty"`
	// DisabledPlans undocumented
	DisabledPlans []UUID `json:"disabledPlans,omitempty"`
	// AssignedByGroup undocumented
	AssignedByGroup *string `json:"assignedByGroup,omitempty"`
	// State undocumented
	State *string `json:"state,omitempty"`
	// Error undocumented
	Error *string `json:"error,omitempty"`
}
