// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EventMessageRequestObject undocumented
type EventMessageRequestObject struct {
	// EventMessage is the base model of EventMessageRequestObject
	EventMessage
	// PreviousLocation undocumented
	PreviousLocation *Location `json:"previousLocation,omitempty"`
	// PreviousStartDateTime undocumented
	PreviousStartDateTime *DateTimeTimeZone `json:"previousStartDateTime,omitempty"`
	// PreviousEndDateTime undocumented
	PreviousEndDateTime *DateTimeTimeZone `json:"previousEndDateTime,omitempty"`
	// ResponseRequested undocumented
	ResponseRequested *bool `json:"responseRequested,omitempty"`
}
