// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// OMASettingDateTime undocumented
type OMASettingDateTime struct {
	// OMASetting is the base model of OMASettingDateTime
	OMASetting
	// Value Value.
	Value *time.Time `json:"value,omitempty"`
}
