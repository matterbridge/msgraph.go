// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PlannerPlanDetails undocumented
type PlannerPlanDetails struct {
	// Entity is the base model of PlannerPlanDetails
	Entity
	// SharedWith undocumented
	SharedWith *PlannerUserIDs `json:"sharedWith,omitempty"`
	// CategoryDescriptions undocumented
	CategoryDescriptions *PlannerCategoryDescriptions `json:"categoryDescriptions,omitempty"`
}
