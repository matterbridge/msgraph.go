// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserFlowType undocumented
type UserFlowType int

const (
	// UserFlowTypeVSignUp undocumented
	UserFlowTypeVSignUp UserFlowType = 1
	// UserFlowTypeVSignIn undocumented
	UserFlowTypeVSignIn UserFlowType = 2
	// UserFlowTypeVSignUpOrSignIn undocumented
	UserFlowTypeVSignUpOrSignIn UserFlowType = 3
	// UserFlowTypeVPasswordReset undocumented
	UserFlowTypeVPasswordReset UserFlowType = 4
	// UserFlowTypeVProfileUpdate undocumented
	UserFlowTypeVProfileUpdate UserFlowType = 5
	// UserFlowTypeVResourceOwner undocumented
	UserFlowTypeVResourceOwner UserFlowType = 6
	// UserFlowTypeVUnknownFutureValue undocumented
	UserFlowTypeVUnknownFutureValue UserFlowType = 7
)

// UserFlowTypePSignUp returns a pointer to UserFlowTypeVSignUp
func UserFlowTypePSignUp() *UserFlowType {
	v := UserFlowTypeVSignUp
	return &v
}

// UserFlowTypePSignIn returns a pointer to UserFlowTypeVSignIn
func UserFlowTypePSignIn() *UserFlowType {
	v := UserFlowTypeVSignIn
	return &v
}

// UserFlowTypePSignUpOrSignIn returns a pointer to UserFlowTypeVSignUpOrSignIn
func UserFlowTypePSignUpOrSignIn() *UserFlowType {
	v := UserFlowTypeVSignUpOrSignIn
	return &v
}

// UserFlowTypePPasswordReset returns a pointer to UserFlowTypeVPasswordReset
func UserFlowTypePPasswordReset() *UserFlowType {
	v := UserFlowTypeVPasswordReset
	return &v
}

// UserFlowTypePProfileUpdate returns a pointer to UserFlowTypeVProfileUpdate
func UserFlowTypePProfileUpdate() *UserFlowType {
	v := UserFlowTypeVProfileUpdate
	return &v
}

// UserFlowTypePResourceOwner returns a pointer to UserFlowTypeVResourceOwner
func UserFlowTypePResourceOwner() *UserFlowType {
	v := UserFlowTypeVResourceOwner
	return &v
}

// UserFlowTypePUnknownFutureValue returns a pointer to UserFlowTypeVUnknownFutureValue
func UserFlowTypePUnknownFutureValue() *UserFlowType {
	v := UserFlowTypeVUnknownFutureValue
	return &v
}
