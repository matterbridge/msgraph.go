// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AuthenticationRequirementPolicy undocumented
type AuthenticationRequirementPolicy struct {
	// Object is the base model of AuthenticationRequirementPolicy
	Object
	// RequirementProvider undocumented
	RequirementProvider *RequirementProvider `json:"requirementProvider,omitempty"`
	// Detail undocumented
	Detail *string `json:"detail,omitempty"`
}
