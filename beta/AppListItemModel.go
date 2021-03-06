// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AppListItem undocumented
type AppListItem struct {
	// Object is the base model of AppListItem
	Object
	// Name The application name
	Name *string `json:"name,omitempty"`
	// Publisher The publisher of the application
	Publisher *string `json:"publisher,omitempty"`
	// AppStoreURL The Store URL of the application
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
	// AppID The application or bundle identifier of the application
	AppID *string `json:"appId,omitempty"`
}
