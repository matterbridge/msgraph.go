// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ChatActivityStatistics undocumented
type ChatActivityStatistics struct {
	// ActivityStatistics is the base model of ChatActivityStatistics
	ActivityStatistics
	// AfterHours undocumented
	AfterHours *time.Duration `json:"afterHours,omitempty"`
}
