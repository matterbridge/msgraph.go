// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TargetResource undocumented
type TargetResource struct {
	// Object is the base model of TargetResource
	Object
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// GroupType undocumented
	GroupType *GroupType `json:"groupType,omitempty"`
	// ModifiedProperties undocumented
	ModifiedProperties []ModifiedProperty `json:"modifiedProperties,omitempty"`
}
