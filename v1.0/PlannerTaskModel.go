// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// PlannerTask undocumented
type PlannerTask struct {
	// Entity is the base model of PlannerTask
	Entity
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// PlanID undocumented
	PlanID *string `json:"planId,omitempty"`
	// BucketID undocumented
	BucketID *string `json:"bucketId,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// OrderHint undocumented
	OrderHint *string `json:"orderHint,omitempty"`
	// AssigneePriority undocumented
	AssigneePriority *string `json:"assigneePriority,omitempty"`
	// PercentComplete undocumented
	PercentComplete *int `json:"percentComplete,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DueDateTime undocumented
	DueDateTime *time.Time `json:"dueDateTime,omitempty"`
	// HasDescription undocumented
	HasDescription *bool `json:"hasDescription,omitempty"`
	// PreviewType undocumented
	PreviewType *PlannerPreviewType `json:"previewType,omitempty"`
	// CompletedDateTime undocumented
	CompletedDateTime *time.Time `json:"completedDateTime,omitempty"`
	// CompletedBy undocumented
	CompletedBy *IdentitySet `json:"completedBy,omitempty"`
	// ReferenceCount undocumented
	ReferenceCount *int `json:"referenceCount,omitempty"`
	// ChecklistItemCount undocumented
	ChecklistItemCount *int `json:"checklistItemCount,omitempty"`
	// ActiveChecklistItemCount undocumented
	ActiveChecklistItemCount *int `json:"activeChecklistItemCount,omitempty"`
	// AppliedCategories undocumented
	AppliedCategories *PlannerAppliedCategories `json:"appliedCategories,omitempty"`
	// Assignments undocumented
	Assignments *PlannerAssignments `json:"assignments,omitempty"`
	// ConversationThreadID undocumented
	ConversationThreadID *string `json:"conversationThreadId,omitempty"`
	// Details undocumented
	Details *PlannerTaskDetails `json:"details,omitempty"`
	// AssignedToTaskBoardFormat undocumented
	AssignedToTaskBoardFormat *PlannerAssignedToTaskBoardTaskFormat `json:"assignedToTaskBoardFormat,omitempty"`
	// ProgressTaskBoardFormat undocumented
	ProgressTaskBoardFormat *PlannerProgressTaskBoardTaskFormat `json:"progressTaskBoardFormat,omitempty"`
	// BucketTaskBoardFormat undocumented
	BucketTaskBoardFormat *PlannerBucketTaskBoardTaskFormat `json:"bucketTaskBoardFormat,omitempty"`
}
