// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ThreatCategory undocumented
type ThreatCategory int

const (
	// ThreatCategoryVSpam undocumented
	ThreatCategoryVSpam ThreatCategory = 1
	// ThreatCategoryVPhishing undocumented
	ThreatCategoryVPhishing ThreatCategory = 2
	// ThreatCategoryVMalware undocumented
	ThreatCategoryVMalware ThreatCategory = 3
	// ThreatCategoryVUnknownFutureValue undocumented
	ThreatCategoryVUnknownFutureValue ThreatCategory = 4
)

// ThreatCategoryPSpam returns a pointer to ThreatCategoryVSpam
func ThreatCategoryPSpam() *ThreatCategory {
	v := ThreatCategoryVSpam
	return &v
}

// ThreatCategoryPPhishing returns a pointer to ThreatCategoryVPhishing
func ThreatCategoryPPhishing() *ThreatCategory {
	v := ThreatCategoryVPhishing
	return &v
}

// ThreatCategoryPMalware returns a pointer to ThreatCategoryVMalware
func ThreatCategoryPMalware() *ThreatCategory {
	v := ThreatCategoryVMalware
	return &v
}

// ThreatCategoryPUnknownFutureValue returns a pointer to ThreatCategoryVUnknownFutureValue
func ThreatCategoryPUnknownFutureValue() *ThreatCategory {
	v := ThreatCategoryVUnknownFutureValue
	return &v
}
