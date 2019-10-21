// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Call undocumented
type Call struct {
	// Entity is the base model of Call
	Entity
	// State undocumented
	State *CallState `json:"state,omitempty"`
	// MediaState undocumented
	MediaState *CallMediaState `json:"mediaState,omitempty"`
	// ResultInfo undocumented
	ResultInfo *ResultInfo `json:"resultInfo,omitempty"`
	// TerminationReason undocumented
	TerminationReason *string `json:"terminationReason,omitempty"`
	// Direction undocumented
	Direction *CallDirection `json:"direction,omitempty"`
	// RingingTimeoutInSeconds undocumented
	RingingTimeoutInSeconds *int `json:"ringingTimeoutInSeconds,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// CallbackURI undocumented
	CallbackURI *string `json:"callbackUri,omitempty"`
	// CallRoutes undocumented
	CallRoutes []CallRoute `json:"callRoutes,omitempty"`
	// Source undocumented
	Source *ParticipantInfo `json:"source,omitempty"`
	// Targets undocumented
	Targets []ParticipantInfo `json:"targets,omitempty"`
	// AnsweredBy undocumented
	AnsweredBy *ParticipantInfo `json:"answeredBy,omitempty"`
	// RequestedModalities undocumented
	RequestedModalities []Modality `json:"requestedModalities,omitempty"`
	// ActiveModalities undocumented
	ActiveModalities []Modality `json:"activeModalities,omitempty"`
	// MediaConfig undocumented
	MediaConfig *MediaConfig `json:"mediaConfig,omitempty"`
	// ChatInfo undocumented
	ChatInfo *ChatInfo `json:"chatInfo,omitempty"`
	// MeetingInfo undocumented
	MeetingInfo *MeetingInfo `json:"meetingInfo,omitempty"`
	// MeetingCapability undocumented
	MeetingCapability *MeetingCapability `json:"meetingCapability,omitempty"`
	// RoutingPolicies undocumented
	RoutingPolicies []RoutingPolicy `json:"routingPolicies,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
	// MyParticipantID undocumented
	MyParticipantID *string `json:"myParticipantId,omitempty"`
	// ToneInfo undocumented
	ToneInfo *ToneInfo `json:"toneInfo,omitempty"`
	// Participants undocumented
	Participants []Participant `json:"participants,omitempty"`
	// AudioRoutingGroups undocumented
	AudioRoutingGroups []AudioRoutingGroup `json:"audioRoutingGroups,omitempty"`
	// Operations undocumented
	Operations []CommsOperation `json:"operations,omitempty"`
}