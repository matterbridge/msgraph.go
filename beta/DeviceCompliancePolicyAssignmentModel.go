// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceCompliancePolicyAssignment Device compliance policy assignment.
type DeviceCompliancePolicyAssignment struct {
	// Entity is the base model of DeviceCompliancePolicyAssignment
	Entity
	// Target Target for the compliance policy assignment.
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}
