// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OfficeSuiteInstallProgressDisplayLevel undocumented
type OfficeSuiteInstallProgressDisplayLevel int

const (
	// OfficeSuiteInstallProgressDisplayLevelVNone undocumented
	OfficeSuiteInstallProgressDisplayLevelVNone OfficeSuiteInstallProgressDisplayLevel = 0
	// OfficeSuiteInstallProgressDisplayLevelVFull undocumented
	OfficeSuiteInstallProgressDisplayLevelVFull OfficeSuiteInstallProgressDisplayLevel = 1
)

// OfficeSuiteInstallProgressDisplayLevelPNone returns a pointer to OfficeSuiteInstallProgressDisplayLevelVNone
func OfficeSuiteInstallProgressDisplayLevelPNone() *OfficeSuiteInstallProgressDisplayLevel {
	v := OfficeSuiteInstallProgressDisplayLevelVNone
	return &v
}

// OfficeSuiteInstallProgressDisplayLevelPFull returns a pointer to OfficeSuiteInstallProgressDisplayLevelVFull
func OfficeSuiteInstallProgressDisplayLevelPFull() *OfficeSuiteInstallProgressDisplayLevel {
	v := OfficeSuiteInstallProgressDisplayLevelVFull
	return &v
}
