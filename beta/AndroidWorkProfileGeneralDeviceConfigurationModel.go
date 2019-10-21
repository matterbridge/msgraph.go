// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidWorkProfileGeneralDeviceConfiguration Android Work Profile general device configuration.
type AndroidWorkProfileGeneralDeviceConfiguration struct {
	// DeviceConfiguration is the base model of AndroidWorkProfileGeneralDeviceConfiguration
	DeviceConfiguration
	// PasswordBlockFingerprintUnlock Indicates whether or not to block fingerprint unlock.
	PasswordBlockFingerprintUnlock *bool `json:"passwordBlockFingerprintUnlock,omitempty"`
	// PasswordBlockTrustAgents Indicates whether or not to block Smart Lock and other trust agents.
	PasswordBlockTrustAgents *bool `json:"passwordBlockTrustAgents,omitempty"`
	// PasswordExpirationDays Number of days before the password expires. Valid values 1 to 365
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumLength Minimum length of passwords. Valid values 4 to 16
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeScreenTimeout Minutes of inactivity before the screen times out.
	PasswordMinutesOfInactivityBeforeScreenTimeout *int `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	// PasswordPreviousPasswordBlockCount Number of previous passwords to block. Valid values 0 to 24
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordSignInFailureCountBeforeFactoryReset Number of sign in failures allowed before factory reset. Valid values 1 to 16
	PasswordSignInFailureCountBeforeFactoryReset *int `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	// PasswordRequiredType Type of password that is required.
	PasswordRequiredType *AndroidWorkProfileRequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// WorkProfileDataSharingType Type of data sharing that is allowed.
	WorkProfileDataSharingType *AndroidWorkProfileCrossProfileDataSharingType `json:"workProfileDataSharingType,omitempty"`
	// WorkProfileBlockNotificationsWhileDeviceLocked Indicates whether or not to block notifications while device locked.
	WorkProfileBlockNotificationsWhileDeviceLocked *bool `json:"workProfileBlockNotificationsWhileDeviceLocked,omitempty"`
	// WorkProfileBlockAddingAccounts Block users from adding/removing accounts in work profile.
	WorkProfileBlockAddingAccounts *bool `json:"workProfileBlockAddingAccounts,omitempty"`
	// WorkProfileBluetoothEnableContactSharing Allow bluetooth devices to access enterprise contacts.
	WorkProfileBluetoothEnableContactSharing *bool `json:"workProfileBluetoothEnableContactSharing,omitempty"`
	// WorkProfileBlockScreenCapture Block screen capture in work profile.
	WorkProfileBlockScreenCapture *bool `json:"workProfileBlockScreenCapture,omitempty"`
	// WorkProfileBlockCrossProfileCallerID Block display work profile caller ID in personal profile.
	WorkProfileBlockCrossProfileCallerID *bool `json:"workProfileBlockCrossProfileCallerId,omitempty"`
	// WorkProfileBlockCamera Block work profile camera.
	WorkProfileBlockCamera *bool `json:"workProfileBlockCamera,omitempty"`
	// WorkProfileBlockCrossProfileContactsSearch Block work profile contacts availability in personal profile.
	WorkProfileBlockCrossProfileContactsSearch *bool `json:"workProfileBlockCrossProfileContactsSearch,omitempty"`
	// WorkProfileBlockCrossProfileCopyPaste Boolean that indicates if the setting disallow cross profile copy/paste is enabled.
	WorkProfileBlockCrossProfileCopyPaste *bool `json:"workProfileBlockCrossProfileCopyPaste,omitempty"`
	// WorkProfileDefaultAppPermissionPolicy Type of password that is required.
	WorkProfileDefaultAppPermissionPolicy *AndroidWorkProfileDefaultAppPermissionPolicyType `json:"workProfileDefaultAppPermissionPolicy,omitempty"`
	// WorkProfilePasswordBlockFingerprintUnlock Indicates whether or not to block fingerprint unlock for work profile.
	WorkProfilePasswordBlockFingerprintUnlock *bool `json:"workProfilePasswordBlockFingerprintUnlock,omitempty"`
	// WorkProfilePasswordBlockTrustAgents Indicates whether or not to block Smart Lock and other trust agents for work profile.
	WorkProfilePasswordBlockTrustAgents *bool `json:"workProfilePasswordBlockTrustAgents,omitempty"`
	// WorkProfilePasswordExpirationDays Number of days before the work profile password expires. Valid values 1 to 365
	WorkProfilePasswordExpirationDays *int `json:"workProfilePasswordExpirationDays,omitempty"`
	// WorkProfilePasswordMinimumLength Minimum length of work profile password. Valid values 4 to 16
	WorkProfilePasswordMinimumLength *int `json:"workProfilePasswordMinimumLength,omitempty"`
	// WorkProfilePasswordMinNumericCharacters Minimum # of numeric characters required in work profile password. Valid values 1 to 10
	WorkProfilePasswordMinNumericCharacters *int `json:"workProfilePasswordMinNumericCharacters,omitempty"`
	// WorkProfilePasswordMinNonLetterCharacters Minimum # of non-letter characters required in work profile password. Valid values 1 to 10
	WorkProfilePasswordMinNonLetterCharacters *int `json:"workProfilePasswordMinNonLetterCharacters,omitempty"`
	// WorkProfilePasswordMinLetterCharacters Minimum # of letter characters required in work profile password. Valid values 1 to 10
	WorkProfilePasswordMinLetterCharacters *int `json:"workProfilePasswordMinLetterCharacters,omitempty"`
	// WorkProfilePasswordMinLowerCaseCharacters Minimum # of lower-case characters required in work profile password. Valid values 1 to 10
	WorkProfilePasswordMinLowerCaseCharacters *int `json:"workProfilePasswordMinLowerCaseCharacters,omitempty"`
	// WorkProfilePasswordMinUpperCaseCharacters Minimum # of upper-case characters required in work profile password. Valid values 1 to 10
	WorkProfilePasswordMinUpperCaseCharacters *int `json:"workProfilePasswordMinUpperCaseCharacters,omitempty"`
	// WorkProfilePasswordMinSymbolCharacters Minimum # of symbols required in work profile password. Valid values 1 to 10
	WorkProfilePasswordMinSymbolCharacters *int `json:"workProfilePasswordMinSymbolCharacters,omitempty"`
	// WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout Minutes of inactivity before the screen times out.
	WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout *int `json:"workProfilePasswordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	// WorkProfilePasswordPreviousPasswordBlockCount Number of previous work profile passwords to block. Valid values 0 to 24
	WorkProfilePasswordPreviousPasswordBlockCount *int `json:"workProfilePasswordPreviousPasswordBlockCount,omitempty"`
	// WorkProfilePasswordSignInFailureCountBeforeFactoryReset Number of sign in failures allowed before work profile is removed and all corporate data deleted. Valid values 1 to 16
	WorkProfilePasswordSignInFailureCountBeforeFactoryReset *int `json:"workProfilePasswordSignInFailureCountBeforeFactoryReset,omitempty"`
	// WorkProfilePasswordRequiredType Type of work profile password that is required.
	WorkProfilePasswordRequiredType *AndroidWorkProfileRequiredPasswordType `json:"workProfilePasswordRequiredType,omitempty"`
	// WorkProfileRequirePassword Password is required or not for work profile
	WorkProfileRequirePassword *bool `json:"workProfileRequirePassword,omitempty"`
	// SecurityRequireVerifyApps Require the Android Verify apps feature is turned on.
	SecurityRequireVerifyApps *bool `json:"securityRequireVerifyApps,omitempty"`
	// VpnAlwaysOnPackageIdentifier Enable lockdown mode for always-on VPN.
	VpnAlwaysOnPackageIdentifier *string `json:"vpnAlwaysOnPackageIdentifier,omitempty"`
	// VpnEnableAlwaysOnLockdownMode Enable lockdown mode for always-on VPN.
	VpnEnableAlwaysOnLockdownMode *bool `json:"vpnEnableAlwaysOnLockdownMode,omitempty"`
	// WorkProfileAllowWidgets Allow widgets from work profile apps.
	WorkProfileAllowWidgets *bool `json:"workProfileAllowWidgets,omitempty"`
}
