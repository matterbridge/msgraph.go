// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// TimeOffRequestObject undocumented
type TimeOffRequestObject struct {
	// ScheduleChangeRequestObject is the base model of TimeOffRequestObject
	ScheduleChangeRequestObject
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// TimeOffReasonID undocumented
	TimeOffReasonID *string `json:"timeOffReasonId,omitempty"`
}
