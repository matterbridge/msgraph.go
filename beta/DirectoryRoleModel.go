// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DirectoryRole undocumented
type DirectoryRole struct {
	// DirectoryObject is the base model of DirectoryRole
	DirectoryObject
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// RoleTemplateID undocumented
	RoleTemplateID *string `json:"roleTemplateId,omitempty"`
	// Members undocumented
	Members []DirectoryObject `json:"members,omitempty"`
	// ScopedMembers undocumented
	ScopedMembers []ScopedRoleMembership `json:"scopedMembers,omitempty"`
}
