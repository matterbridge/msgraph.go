// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UnifiedRoleAssignment undocumented
type UnifiedRoleAssignment struct {
	// Entity is the base model of UnifiedRoleAssignment
	Entity
	// PrincipalID undocumented
	PrincipalID *string `json:"principalId,omitempty"`
	// ResourceScope undocumented
	ResourceScope *string `json:"resourceScope,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
}