// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AttendeeAvailability undocumented
type AttendeeAvailability struct {
	// Object is the base model of AttendeeAvailability
	Object
	// Attendee undocumented
	Attendee *AttendeeBase `json:"attendee,omitempty"`
	// Availability undocumented
	Availability *FreeBusyStatus `json:"availability,omitempty"`
}
