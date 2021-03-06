// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OMASettingBase64 undocumented
type OMASettingBase64 struct {
	// OMASetting is the base model of OMASettingBase64
	OMASetting
	// FileName File name associated with the Value property (*.cer | *.crt | *.p7b | *.bin).
	FileName *string `json:"fileName,omitempty"`
	// Value Value. (Base64 encoded string)
	Value *string `json:"value,omitempty"`
}
