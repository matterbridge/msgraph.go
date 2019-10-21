// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RoleAssignment The Role Assignment resource. Role assignments tie together a role definition with members and scopes. There can be one or more role assignments per role. This applies to custom and built-in roles.
type RoleAssignment struct {
	// Entity is the base model of RoleAssignment
	Entity
	// DisplayName The display or friendly name of the role Assignment.
	DisplayName *string `json:"displayName,omitempty"`
	// Description Description of the Role Assignment.
	Description *string `json:"description,omitempty"`
	// ResourceScopes List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
	ResourceScopes []string `json:"resourceScopes,omitempty"`
	// RoleDefinition undocumented
	RoleDefinition *RoleDefinition `json:"roleDefinition,omitempty"`
}
