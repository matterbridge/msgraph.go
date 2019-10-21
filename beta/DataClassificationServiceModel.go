// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DataClassificationService undocumented
type DataClassificationService struct {
	// Entity is the base model of DataClassificationService
	Entity
	// ExactMatchDataStores undocumented
	ExactMatchDataStores []ExactMatchDataStore `json:"exactMatchDataStores,omitempty"`
	// SensitiveTypes undocumented
	SensitiveTypes []SensitiveType `json:"sensitiveTypes,omitempty"`
	// Jobs undocumented
	Jobs []JobResponseBase `json:"jobs,omitempty"`
	// ClassifyFileJobs undocumented
	ClassifyFileJobs []JobResponseBase `json:"classifyFileJobs,omitempty"`
	// ClassifyTextJobs undocumented
	ClassifyTextJobs []JobResponseBase `json:"classifyTextJobs,omitempty"`
	// EvaluateLabelJobs undocumented
	EvaluateLabelJobs []JobResponseBase `json:"evaluateLabelJobs,omitempty"`
	// ClassifyText undocumented
	ClassifyText []TextClassificationRequestObject `json:"classifyText,omitempty"`
	// ClassifyFile undocumented
	ClassifyFile []FileClassificationRequestObject `json:"classifyFile,omitempty"`
	// SensitivityLabels undocumented
	SensitivityLabels []SensitivityLabel `json:"sensitivityLabels,omitempty"`
	// ExactMatchUploadAgents undocumented
	ExactMatchUploadAgents []ExactMatchUploadAgent `json:"exactMatchUploadAgents,omitempty"`
}
