// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// SharePointActivityUserDetail undocumented
type SharePointActivityUserDetail struct {
	// Entity is the base model of SharePointActivityUserDetail
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// DeletedDate undocumented
	DeletedDate *time.Time `json:"deletedDate,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *time.Time `json:"lastActivityDate,omitempty"`
	// ViewedOrEditedFileCount undocumented
	ViewedOrEditedFileCount *int `json:"viewedOrEditedFileCount,omitempty"`
	// SyncedFileCount undocumented
	SyncedFileCount *int `json:"syncedFileCount,omitempty"`
	// SharedInternallyFileCount undocumented
	SharedInternallyFileCount *int `json:"sharedInternallyFileCount,omitempty"`
	// SharedExternallyFileCount undocumented
	SharedExternallyFileCount *int `json:"sharedExternallyFileCount,omitempty"`
	// VisitedPageCount undocumented
	VisitedPageCount *int `json:"visitedPageCount,omitempty"`
	// AssignedProducts undocumented
	AssignedProducts []string `json:"assignedProducts,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
