// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SharingLink undocumented
type SharingLink struct {
	// Object is the base model of SharingLink
	Object
	// Application undocumented
	Application *Identity `json:"application,omitempty"`
	// Scope undocumented
	Scope *string `json:"scope,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}
