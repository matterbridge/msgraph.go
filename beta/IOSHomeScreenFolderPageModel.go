// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IOSHomeScreenFolderPage undocumented
type IOSHomeScreenFolderPage struct {
	// Object is the base model of IOSHomeScreenFolderPage
	Object
	// DisplayName Name of the folder page
	DisplayName *string `json:"displayName,omitempty"`
	// Apps A list of apps to appear on a page within a folder. This collection can contain a maximum of 500 elements.
	Apps []IOSHomeScreenApp `json:"apps,omitempty"`
}
