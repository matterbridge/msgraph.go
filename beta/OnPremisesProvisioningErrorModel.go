// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// OnPremisesProvisioningError undocumented
type OnPremisesProvisioningError struct {
	// Object is the base model of OnPremisesProvisioningError
	Object
	// Value undocumented
	Value *string `json:"value,omitempty"`
	// Category undocumented
	Category *string `json:"category,omitempty"`
	// PropertyCausingError undocumented
	PropertyCausingError *string `json:"propertyCausingError,omitempty"`
	// OccurredDateTime undocumented
	OccurredDateTime *time.Time `json:"occurredDateTime,omitempty"`
}
