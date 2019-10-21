// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// User Represents an Azure Active Directory user object.
type User struct {
	// DirectoryObject is the base model of User
	DirectoryObject
	// AccountEnabled undocumented
	AccountEnabled *bool `json:"accountEnabled,omitempty"`
	// AgeGroup undocumented
	AgeGroup *string `json:"ageGroup,omitempty"`
	// AssignedLicenses undocumented
	AssignedLicenses []AssignedLicense `json:"assignedLicenses,omitempty"`
	// AssignedPlans undocumented
	AssignedPlans []AssignedPlan `json:"assignedPlans,omitempty"`
	// BusinessPhones undocumented
	BusinessPhones []string `json:"businessPhones,omitempty"`
	// City undocumented
	City *string `json:"city,omitempty"`
	// CompanyName undocumented
	CompanyName *string `json:"companyName,omitempty"`
	// ConsentProvidedForMinor undocumented
	ConsentProvidedForMinor *string `json:"consentProvidedForMinor,omitempty"`
	// Country undocumented
	Country *string `json:"country,omitempty"`
	// Department undocumented
	Department *string `json:"department,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EmployeeID undocumented
	EmployeeID *string `json:"employeeId,omitempty"`
	// FaxNumber undocumented
	FaxNumber *string `json:"faxNumber,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// ImAddresses undocumented
	ImAddresses []string `json:"imAddresses,omitempty"`
	// IsResourceAccount undocumented
	IsResourceAccount *bool `json:"isResourceAccount,omitempty"`
	// JobTitle undocumented
	JobTitle *string `json:"jobTitle,omitempty"`
	// LegalAgeGroupClassification undocumented
	LegalAgeGroupClassification *string `json:"legalAgeGroupClassification,omitempty"`
	// LicenseAssignmentStates undocumented
	LicenseAssignmentStates []LicenseAssignmentState `json:"licenseAssignmentStates,omitempty"`
	// Mail undocumented
	Mail *string `json:"mail,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// MobilePhone undocumented
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// OnPremisesDistinguishedName undocumented
	OnPremisesDistinguishedName *string `json:"onPremisesDistinguishedName,omitempty"`
	// OnPremisesExtensionAttributes undocumented
	OnPremisesExtensionAttributes *OnPremisesExtensionAttributes `json:"onPremisesExtensionAttributes,omitempty"`
	// OnPremisesImmutableID undocumented
	OnPremisesImmutableID *string `json:"onPremisesImmutableId,omitempty"`
	// OnPremisesLastSyncDateTime undocumented
	OnPremisesLastSyncDateTime *time.Time `json:"onPremisesLastSyncDateTime,omitempty"`
	// OnPremisesProvisioningErrors undocumented
	OnPremisesProvisioningErrors []OnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`
	// OnPremisesSecurityIdentifier undocumented
	OnPremisesSecurityIdentifier *string `json:"onPremisesSecurityIdentifier,omitempty"`
	// OnPremisesSyncEnabled undocumented
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`
	// OnPremisesDomainName undocumented
	OnPremisesDomainName *string `json:"onPremisesDomainName,omitempty"`
	// OnPremisesSamAccountName undocumented
	OnPremisesSamAccountName *string `json:"onPremisesSamAccountName,omitempty"`
	// OnPremisesUserPrincipalName undocumented
	OnPremisesUserPrincipalName *string `json:"onPremisesUserPrincipalName,omitempty"`
	// OtherMails undocumented
	OtherMails []string `json:"otherMails,omitempty"`
	// PasswordPolicies undocumented
	PasswordPolicies *string `json:"passwordPolicies,omitempty"`
	// PasswordProfile undocumented
	PasswordProfile *PasswordProfile `json:"passwordProfile,omitempty"`
	// OfficeLocation undocumented
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// PostalCode undocumented
	PostalCode *string `json:"postalCode,omitempty"`
	// PreferredLanguage undocumented
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// ProvisionedPlans undocumented
	ProvisionedPlans []ProvisionedPlan `json:"provisionedPlans,omitempty"`
	// ProxyAddresses undocumented
	ProxyAddresses []string `json:"proxyAddresses,omitempty"`
	// ShowInAddressList undocumented
	ShowInAddressList *bool `json:"showInAddressList,omitempty"`
	// SignInSessionsValidFromDateTime undocumented
	SignInSessionsValidFromDateTime *time.Time `json:"signInSessionsValidFromDateTime,omitempty"`
	// State undocumented
	State *string `json:"state,omitempty"`
	// StreetAddress undocumented
	StreetAddress *string `json:"streetAddress,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// UsageLocation undocumented
	UsageLocation *string `json:"usageLocation,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// UserType undocumented
	UserType *string `json:"userType,omitempty"`
	// MailboxSettings undocumented
	MailboxSettings *MailboxSettings `json:"mailboxSettings,omitempty"`
	// DeviceEnrollmentLimit The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000.
	DeviceEnrollmentLimit *int `json:"deviceEnrollmentLimit,omitempty"`
	// AboutMe undocumented
	AboutMe *string `json:"aboutMe,omitempty"`
	// Birthday undocumented
	Birthday *time.Time `json:"birthday,omitempty"`
	// HireDate undocumented
	HireDate *time.Time `json:"hireDate,omitempty"`
	// Interests undocumented
	Interests []string `json:"interests,omitempty"`
	// MySite undocumented
	MySite *string `json:"mySite,omitempty"`
	// PastProjects undocumented
	PastProjects []string `json:"pastProjects,omitempty"`
	// PreferredName undocumented
	PreferredName *string `json:"preferredName,omitempty"`
	// Responsibilities undocumented
	Responsibilities []string `json:"responsibilities,omitempty"`
	// Schools undocumented
	Schools []string `json:"schools,omitempty"`
	// Skills undocumented
	Skills []string `json:"skills,omitempty"`
	// OwnedDevices undocumented
	OwnedDevices []DirectoryObject `json:"ownedDevices,omitempty"`
	// RegisteredDevices undocumented
	RegisteredDevices []DirectoryObject `json:"registeredDevices,omitempty"`
	// Manager undocumented
	Manager *DirectoryObject `json:"manager,omitempty"`
	// DirectReports undocumented
	DirectReports []DirectoryObject `json:"directReports,omitempty"`
	// MemberOf undocumented
	MemberOf []DirectoryObject `json:"memberOf,omitempty"`
	// CreatedObjects undocumented
	CreatedObjects []DirectoryObject `json:"createdObjects,omitempty"`
	// OwnedObjects undocumented
	OwnedObjects []DirectoryObject `json:"ownedObjects,omitempty"`
	// LicenseDetails undocumented
	LicenseDetails []LicenseDetails `json:"licenseDetails,omitempty"`
	// TransitiveMemberOf undocumented
	TransitiveMemberOf []DirectoryObject `json:"transitiveMemberOf,omitempty"`
	// Outlook undocumented
	Outlook *OutlookUser `json:"outlook,omitempty"`
	// Messages undocumented
	Messages []Message `json:"messages,omitempty"`
	// MailFolders undocumented
	MailFolders []MailFolder `json:"mailFolders,omitempty"`
	// Calendar undocumented
	Calendar *Calendar `json:"calendar,omitempty"`
	// Calendars undocumented
	Calendars []Calendar `json:"calendars,omitempty"`
	// CalendarGroups undocumented
	CalendarGroups []CalendarGroup `json:"calendarGroups,omitempty"`
	// CalendarView undocumented
	CalendarView []Event `json:"calendarView,omitempty"`
	// Events undocumented
	Events []Event `json:"events,omitempty"`
	// People undocumented
	People []Person `json:"people,omitempty"`
	// Contacts undocumented
	Contacts []Contact `json:"contacts,omitempty"`
	// ContactFolders undocumented
	ContactFolders []ContactFolder `json:"contactFolders,omitempty"`
	// InferenceClassification undocumented
	InferenceClassification *InferenceClassification `json:"inferenceClassification,omitempty"`
	// Photo undocumented
	Photo *ProfilePhoto `json:"photo,omitempty"`
	// Photos undocumented
	Photos []ProfilePhoto `json:"photos,omitempty"`
	// Drive undocumented
	Drive *Drive `json:"drive,omitempty"`
	// Drives undocumented
	Drives []Drive `json:"drives,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
	// ManagedDevices undocumented
	ManagedDevices []ManagedDevice `json:"managedDevices,omitempty"`
	// ManagedAppRegistrations undocumented
	ManagedAppRegistrations []ManagedAppRegistration `json:"managedAppRegistrations,omitempty"`
	// DeviceManagementTroubleshootingEvents undocumented
	DeviceManagementTroubleshootingEvents []DeviceManagementTroubleshootingEvent `json:"deviceManagementTroubleshootingEvents,omitempty"`
	// Planner undocumented
	Planner *PlannerUser `json:"planner,omitempty"`
	// Insights undocumented
	Insights *OfficeGraphInsights `json:"insights,omitempty"`
	// Settings undocumented
	Settings *UserSettings `json:"settings,omitempty"`
	// Onenote undocumented
	Onenote *Onenote `json:"onenote,omitempty"`
	// Activities undocumented
	Activities []UserActivity `json:"activities,omitempty"`
	// OnlineMeetings undocumented
	OnlineMeetings []OnlineMeeting `json:"onlineMeetings,omitempty"`
	// JoinedTeams undocumented
	JoinedTeams []Group `json:"joinedTeams,omitempty"`
}
