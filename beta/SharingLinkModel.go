// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SharingLink undocumented
type SharingLink struct {
	// Object is the base model of SharingLink
	Object
	// Application undocumented
	Application *Identity `json:"application,omitempty"`
	// PreventsDownload undocumented
	PreventsDownload *bool `json:"preventsDownload,omitempty"`
	// ConfiguratorURL undocumented
	ConfiguratorURL *string `json:"configuratorUrl,omitempty"`
	// Scope undocumented
	Scope *string `json:"scope,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// WebHTML undocumented
	WebHTML *string `json:"webHtml,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}
