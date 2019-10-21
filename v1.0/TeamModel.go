// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Team undocumented
type Team struct {
	// Entity is the base model of Team
	Entity
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// MemberSettings undocumented
	MemberSettings *TeamMemberSettings `json:"memberSettings,omitempty"`
	// GuestSettings undocumented
	GuestSettings *TeamGuestSettings `json:"guestSettings,omitempty"`
	// MessagingSettings undocumented
	MessagingSettings *TeamMessagingSettings `json:"messagingSettings,omitempty"`
	// FunSettings undocumented
	FunSettings *TeamFunSettings `json:"funSettings,omitempty"`
	// IsArchived undocumented
	IsArchived *bool `json:"isArchived,omitempty"`
	// Channels undocumented
	Channels []Channel `json:"channels,omitempty"`
	// InstalledApps undocumented
	InstalledApps []TeamsAppInstallation `json:"installedApps,omitempty"`
	// Operations undocumented
	Operations []TeamsAsyncOperation `json:"operations,omitempty"`
}
