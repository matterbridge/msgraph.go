// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OutlookGeoCoordinates undocumented
type OutlookGeoCoordinates struct {
	// Object is the base model of OutlookGeoCoordinates
	Object
	// Latitude undocumented
	Latitude *float64 `json:"latitude,omitempty"`
	// Longitude undocumented
	Longitude *float64 `json:"longitude,omitempty"`
	// Accuracy undocumented
	Accuracy *float64 `json:"accuracy,omitempty"`
	// Altitude undocumented
	Altitude *float64 `json:"altitude,omitempty"`
	// AltitudeAccuracy undocumented
	AltitudeAccuracy *float64 `json:"altitudeAccuracy,omitempty"`
}
