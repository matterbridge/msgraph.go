// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// LocalSecurityOptionsMinimumSessionSecurity undocumented
type LocalSecurityOptionsMinimumSessionSecurity int

const (
	// LocalSecurityOptionsMinimumSessionSecurityVNone undocumented
	LocalSecurityOptionsMinimumSessionSecurityVNone LocalSecurityOptionsMinimumSessionSecurity = 0
	// LocalSecurityOptionsMinimumSessionSecurityVRequireNtmlV2SessionSecurity undocumented
	LocalSecurityOptionsMinimumSessionSecurityVRequireNtmlV2SessionSecurity LocalSecurityOptionsMinimumSessionSecurity = 1
	// LocalSecurityOptionsMinimumSessionSecurityVRequire128BitEncryption undocumented
	LocalSecurityOptionsMinimumSessionSecurityVRequire128BitEncryption LocalSecurityOptionsMinimumSessionSecurity = 2
	// LocalSecurityOptionsMinimumSessionSecurityVNtlmV2And128BitEncryption undocumented
	LocalSecurityOptionsMinimumSessionSecurityVNtlmV2And128BitEncryption LocalSecurityOptionsMinimumSessionSecurity = 3
)

// LocalSecurityOptionsMinimumSessionSecurityPNone returns a pointer to LocalSecurityOptionsMinimumSessionSecurityVNone
func LocalSecurityOptionsMinimumSessionSecurityPNone() *LocalSecurityOptionsMinimumSessionSecurity {
	v := LocalSecurityOptionsMinimumSessionSecurityVNone
	return &v
}

// LocalSecurityOptionsMinimumSessionSecurityPRequireNtmlV2SessionSecurity returns a pointer to LocalSecurityOptionsMinimumSessionSecurityVRequireNtmlV2SessionSecurity
func LocalSecurityOptionsMinimumSessionSecurityPRequireNtmlV2SessionSecurity() *LocalSecurityOptionsMinimumSessionSecurity {
	v := LocalSecurityOptionsMinimumSessionSecurityVRequireNtmlV2SessionSecurity
	return &v
}

// LocalSecurityOptionsMinimumSessionSecurityPRequire128BitEncryption returns a pointer to LocalSecurityOptionsMinimumSessionSecurityVRequire128BitEncryption
func LocalSecurityOptionsMinimumSessionSecurityPRequire128BitEncryption() *LocalSecurityOptionsMinimumSessionSecurity {
	v := LocalSecurityOptionsMinimumSessionSecurityVRequire128BitEncryption
	return &v
}

// LocalSecurityOptionsMinimumSessionSecurityPNtlmV2And128BitEncryption returns a pointer to LocalSecurityOptionsMinimumSessionSecurityVNtlmV2And128BitEncryption
func LocalSecurityOptionsMinimumSessionSecurityPNtlmV2And128BitEncryption() *LocalSecurityOptionsMinimumSessionSecurity {
	v := LocalSecurityOptionsMinimumSessionSecurityVNtlmV2And128BitEncryption
	return &v
}
