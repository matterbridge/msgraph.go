// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// InvitedUserMessageInfo undocumented
type InvitedUserMessageInfo struct {
	// Object is the base model of InvitedUserMessageInfo
	Object
	// CcRecipients undocumented
	CcRecipients []Recipient `json:"ccRecipients,omitempty"`
	// MessageLanguage undocumented
	MessageLanguage *string `json:"messageLanguage,omitempty"`
	// CustomizedMessageBody undocumented
	CustomizedMessageBody *string `json:"customizedMessageBody,omitempty"`
}
