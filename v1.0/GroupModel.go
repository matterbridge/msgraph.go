// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Group undocumented
type Group struct {
	// DirectoryObject is the base model of Group
	DirectoryObject
	// AssignedLicenses undocumented
	AssignedLicenses []AssignedLicense `json:"assignedLicenses,omitempty"`
	// Classification undocumented
	Classification *string `json:"classification,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// HasMembersWithLicenseErrors undocumented
	HasMembersWithLicenseErrors *bool `json:"hasMembersWithLicenseErrors,omitempty"`
	// GroupTypes undocumented
	GroupTypes []string `json:"groupTypes,omitempty"`
	// LicenseProcessingState undocumented
	LicenseProcessingState *LicenseProcessingState `json:"licenseProcessingState,omitempty"`
	// Mail undocumented
	Mail *string `json:"mail,omitempty"`
	// MailEnabled undocumented
	MailEnabled *bool `json:"mailEnabled,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// OnPremisesLastSyncDateTime undocumented
	OnPremisesLastSyncDateTime *time.Time `json:"onPremisesLastSyncDateTime,omitempty"`
	// OnPremisesProvisioningErrors undocumented
	OnPremisesProvisioningErrors []OnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`
	// OnPremisesSecurityIdentifier undocumented
	OnPremisesSecurityIdentifier *string `json:"onPremisesSecurityIdentifier,omitempty"`
	// OnPremisesSyncEnabled undocumented
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`
	// PreferredDataLocation undocumented
	PreferredDataLocation *string `json:"preferredDataLocation,omitempty"`
	// ProxyAddresses undocumented
	ProxyAddresses []string `json:"proxyAddresses,omitempty"`
	// RenewedDateTime undocumented
	RenewedDateTime *time.Time `json:"renewedDateTime,omitempty"`
	// SecurityEnabled undocumented
	SecurityEnabled *bool `json:"securityEnabled,omitempty"`
	// Visibility undocumented
	Visibility *string `json:"visibility,omitempty"`
	// AllowExternalSenders undocumented
	AllowExternalSenders *bool `json:"allowExternalSenders,omitempty"`
	// AutoSubscribeNewMembers undocumented
	AutoSubscribeNewMembers *bool `json:"autoSubscribeNewMembers,omitempty"`
	// IsSubscribedByMail undocumented
	IsSubscribedByMail *bool `json:"isSubscribedByMail,omitempty"`
	// UnseenCount undocumented
	UnseenCount *int `json:"unseenCount,omitempty"`
	// IsArchived undocumented
	IsArchived *bool `json:"isArchived,omitempty"`
	// Members undocumented
	Members []DirectoryObject `json:"members,omitempty"`
	// MemberOf undocumented
	MemberOf []DirectoryObject `json:"memberOf,omitempty"`
	// MembersWithLicenseErrors undocumented
	MembersWithLicenseErrors []DirectoryObject `json:"membersWithLicenseErrors,omitempty"`
	// TransitiveMembers undocumented
	TransitiveMembers []DirectoryObject `json:"transitiveMembers,omitempty"`
	// TransitiveMemberOf undocumented
	TransitiveMemberOf []DirectoryObject `json:"transitiveMemberOf,omitempty"`
	// CreatedOnBehalfOf undocumented
	CreatedOnBehalfOf *DirectoryObject `json:"createdOnBehalfOf,omitempty"`
	// Owners undocumented
	Owners []DirectoryObject `json:"owners,omitempty"`
	// Settings undocumented
	Settings []GroupSetting `json:"settings,omitempty"`
	// Conversations undocumented
	Conversations []Conversation `json:"conversations,omitempty"`
	// Photos undocumented
	Photos []ProfilePhoto `json:"photos,omitempty"`
	// AcceptedSenders undocumented
	AcceptedSenders []DirectoryObject `json:"acceptedSenders,omitempty"`
	// RejectedSenders undocumented
	RejectedSenders []DirectoryObject `json:"rejectedSenders,omitempty"`
	// Threads undocumented
	Threads []ConversationThread `json:"threads,omitempty"`
	// Calendar undocumented
	Calendar *Calendar `json:"calendar,omitempty"`
	// CalendarView undocumented
	CalendarView []Event `json:"calendarView,omitempty"`
	// Events undocumented
	Events []Event `json:"events,omitempty"`
	// Photo undocumented
	Photo *ProfilePhoto `json:"photo,omitempty"`
	// Drive undocumented
	Drive *Drive `json:"drive,omitempty"`
	// Drives undocumented
	Drives []Drive `json:"drives,omitempty"`
	// Sites undocumented
	Sites []Site `json:"sites,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
	// GroupLifecyclePolicies undocumented
	GroupLifecyclePolicies []GroupLifecyclePolicy `json:"groupLifecyclePolicies,omitempty"`
	// Planner undocumented
	Planner *PlannerGroup `json:"planner,omitempty"`
	// Onenote undocumented
	Onenote *Onenote `json:"onenote,omitempty"`
	// Team undocumented
	Team *Team `json:"team,omitempty"`
}
