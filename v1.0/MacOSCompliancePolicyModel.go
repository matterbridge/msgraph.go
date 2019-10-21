// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MacOSCompliancePolicy This class contains compliance settings for Mac OS.
type MacOSCompliancePolicy struct {
	// DeviceCompliancePolicy is the base model of MacOSCompliancePolicy
	DeviceCompliancePolicy
	// PasswordRequired Whether or not to require a password.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	// PasswordBlockSimple Indicates whether or not to block simple passwords.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`
	// PasswordExpirationDays Number of days before the password expires. Valid values 1 to 65535
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumLength Minimum length of password. Valid values 4 to 14
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeLock Minutes of inactivity before a password is required.
	PasswordMinutesOfInactivityBeforeLock *int `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	// PasswordPreviousPasswordBlockCount Number of previous passwords to block. Valid values 1 to 24
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordMinimumCharacterSetCount The number of character sets required in the password.
	PasswordMinimumCharacterSetCount *int `json:"passwordMinimumCharacterSetCount,omitempty"`
	// PasswordRequiredType The required password type.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// OsMinimumVersion Minimum MacOS version.
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	// OsMaximumVersion Maximum MacOS version.
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	// SystemIntegrityProtectionEnabled Require that devices have enabled system integrity protection.
	SystemIntegrityProtectionEnabled *bool `json:"systemIntegrityProtectionEnabled,omitempty"`
	// DeviceThreatProtectionEnabled Require that devices have enabled device threat protection.
	DeviceThreatProtectionEnabled *bool `json:"deviceThreatProtectionEnabled,omitempty"`
	// DeviceThreatProtectionRequiredSecurityLevel Require Mobile Threat Protection minimum risk level to report noncompliance.
	DeviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`
	// StorageRequireEncryption Require encryption on Mac OS devices.
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`
	// FirewallEnabled Whether the firewall should be enabled or not.
	FirewallEnabled *bool `json:"firewallEnabled,omitempty"`
	// FirewallBlockAllIncoming Corresponds to the “Block all incoming connections” option.
	FirewallBlockAllIncoming *bool `json:"firewallBlockAllIncoming,omitempty"`
	// FirewallEnableStealthMode Corresponds to “Enable stealth mode.”
	FirewallEnableStealthMode *bool `json:"firewallEnableStealthMode,omitempty"`
}
