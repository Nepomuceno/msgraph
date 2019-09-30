# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountEnabled** | Pointer to **bool** |  | [optional] 
**AgeGroup** | Pointer to **string** |  | [optional] 
**AssignedLicenses** | Pointer to [**[]MicrosoftGraphAssignedLicense**](microsoft.graph.assignedLicense.md) |  | [optional] 
**AssignedPlans** | Pointer to [**[]MicrosoftGraphAssignedPlan**](microsoft.graph.assignedPlan.md) |  | [optional] 
**BusinessPhones** | Pointer to **[]string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**ConsentProvidedForMinor** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EmployeeId** | Pointer to **string** |  | [optional] 
**FaxNumber** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**ImAddresses** | Pointer to **[]string** |  | [optional] 
**IsResourceAccount** | Pointer to **bool** |  | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**LegalAgeGroupClassification** | Pointer to **string** |  | [optional] 
**LicenseAssignmentStates** | Pointer to [**[]AnyOfmicrosoftGraphLicenseAssignmentState**](anyOf&lt;microsoft.graph.licenseAssignmentState&gt;.md) |  | [optional] 
**Mail** | Pointer to **string** |  | [optional] 
**MailNickname** | Pointer to **string** |  | [optional] 
**MobilePhone** | Pointer to **string** |  | [optional] 
**OnPremisesDistinguishedName** | Pointer to **string** |  | [optional] 
**OnPremisesExtensionAttributes** | Pointer to [**AnyOfmicrosoftGraphOnPremisesExtensionAttributes**](anyOf&lt;microsoft.graph.onPremisesExtensionAttributes&gt;.md) |  | [optional] 
**OnPremisesImmutableId** | Pointer to **string** |  | [optional] 
**OnPremisesLastSyncDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**OnPremisesProvisioningErrors** | Pointer to [**[]AnyOfmicrosoftGraphOnPremisesProvisioningError**](anyOf&lt;microsoft.graph.onPremisesProvisioningError&gt;.md) |  | [optional] 
**OnPremisesSecurityIdentifier** | Pointer to **string** |  | [optional] 
**OnPremisesSyncEnabled** | Pointer to **bool** |  | [optional] 
**OnPremisesDomainName** | Pointer to **string** |  | [optional] 
**OnPremisesSamAccountName** | Pointer to **string** |  | [optional] 
**OnPremisesUserPrincipalName** | Pointer to **string** |  | [optional] 
**OtherMails** | Pointer to **[]string** |  | [optional] 
**PasswordPolicies** | Pointer to **string** |  | [optional] 
**PasswordProfile** | Pointer to [**AnyOfmicrosoftGraphPasswordProfile**](anyOf&lt;microsoft.graph.passwordProfile&gt;.md) |  | [optional] 
**OfficeLocation** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**PreferredLanguage** | Pointer to **string** |  | [optional] 
**ProvisionedPlans** | Pointer to [**[]MicrosoftGraphProvisionedPlan**](microsoft.graph.provisionedPlan.md) |  | [optional] 
**ProxyAddresses** | Pointer to **[]string** |  | [optional] 
**ShowInAddressList** | Pointer to **bool** |  | [optional] 
**SignInSessionsValidFromDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StreetAddress** | Pointer to **string** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 
**UsageLocation** | Pointer to **string** |  | [optional] 
**UserPrincipalName** | Pointer to **string** |  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**MailboxSettings** | Pointer to [**AnyOfmicrosoftGraphMailboxSettings**](anyOf&lt;microsoft.graph.mailboxSettings&gt;.md) |  | [optional] 
**DeviceEnrollmentLimit** | Pointer to **int32** | The limit on the maximum number of devices that the user is permitted to enroll. Allowed values are 5 or 1000. | [optional] 
**AboutMe** | Pointer to **string** |  | [optional] 
**Birthday** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**HireDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Interests** | Pointer to **[]string** |  | [optional] 
**MySite** | Pointer to **string** |  | [optional] 
**PastProjects** | Pointer to **[]string** |  | [optional] 
**PreferredName** | Pointer to **string** |  | [optional] 
**Responsibilities** | Pointer to **[]string** |  | [optional] 
**Schools** | Pointer to **[]string** |  | [optional] 
**Skills** | Pointer to **[]string** |  | [optional] 
**OwnedDevices** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**RegisteredDevices** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**Manager** | Pointer to [**AnyOfmicrosoftGraphDirectoryObject**](anyOf&lt;microsoft.graph.directoryObject&gt;.md) |  | [optional] 
**DirectReports** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**MemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**CreatedObjects** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**OwnedObjects** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**LicenseDetails** | Pointer to [**[]MicrosoftGraphLicenseDetails**](microsoft.graph.licenseDetails.md) |  | [optional] 
**TransitiveMemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 
**Outlook** | Pointer to [**AnyOfmicrosoftGraphOutlookUser**](anyOf&lt;microsoft.graph.outlookUser&gt;.md) |  | [optional] 
**Messages** | Pointer to [**[]MicrosoftGraphMessage**](microsoft.graph.message.md) |  | [optional] 
**MailFolders** | Pointer to [**[]MicrosoftGraphMailFolder**](microsoft.graph.mailFolder.md) |  | [optional] 
**Calendar** | Pointer to [**AnyOfmicrosoftGraphCalendar**](anyOf&lt;microsoft.graph.calendar&gt;.md) |  | [optional] 
**Calendars** | Pointer to [**[]MicrosoftGraphCalendar**](microsoft.graph.calendar.md) |  | [optional] 
**CalendarGroups** | Pointer to [**[]MicrosoftGraphCalendarGroup**](microsoft.graph.calendarGroup.md) |  | [optional] 
**CalendarView** | Pointer to [**[]MicrosoftGraphEvent**](microsoft.graph.event.md) |  | [optional] 
**Events** | Pointer to [**[]MicrosoftGraphEvent**](microsoft.graph.event.md) |  | [optional] 
**People** | Pointer to [**[]MicrosoftGraphPerson**](microsoft.graph.person.md) |  | [optional] 
**Contacts** | Pointer to [**[]MicrosoftGraphContact**](microsoft.graph.contact.md) |  | [optional] 
**ContactFolders** | Pointer to [**[]MicrosoftGraphContactFolder**](microsoft.graph.contactFolder.md) |  | [optional] 
**InferenceClassification** | Pointer to [**AnyOfmicrosoftGraphInferenceClassification**](anyOf&lt;microsoft.graph.inferenceClassification&gt;.md) |  | [optional] 
**Photo** | Pointer to [**AnyOfmicrosoftGraphProfilePhoto**](anyOf&lt;microsoft.graph.profilePhoto&gt;.md) |  | [optional] 
**Photos** | Pointer to [**[]MicrosoftGraphProfilePhoto**](microsoft.graph.profilePhoto.md) |  | [optional] 
**Drive** | Pointer to [**AnyOfmicrosoftGraphDrive**](anyOf&lt;microsoft.graph.drive&gt;.md) |  | [optional] 
**Drives** | Pointer to [**[]MicrosoftGraphDrive**](microsoft.graph.drive.md) |  | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](microsoft.graph.extension.md) |  | [optional] 
**ManagedDevices** | Pointer to [**[]MicrosoftGraphManagedDevice**](microsoft.graph.managedDevice.md) |  | [optional] 
**ManagedAppRegistrations** | Pointer to [**[]MicrosoftGraphManagedAppRegistration**](microsoft.graph.managedAppRegistration.md) |  | [optional] 
**DeviceManagementTroubleshootingEvents** | Pointer to [**[]MicrosoftGraphDeviceManagementTroubleshootingEvent**](microsoft.graph.deviceManagementTroubleshootingEvent.md) |  | [optional] 
**Planner** | Pointer to [**AnyOfmicrosoftGraphPlannerUser**](anyOf&lt;microsoft.graph.plannerUser&gt;.md) |  | [optional] 
**Insights** | Pointer to [**AnyOfmicrosoftGraphOfficeGraphInsights**](anyOf&lt;microsoft.graph.officeGraphInsights&gt;.md) |  | [optional] 
**Settings** | Pointer to [**AnyOfmicrosoftGraphUserSettings**](anyOf&lt;microsoft.graph.userSettings&gt;.md) |  | [optional] 
**Onenote** | Pointer to [**AnyOfmicrosoftGraphOnenote**](anyOf&lt;microsoft.graph.onenote&gt;.md) |  | [optional] 
**Activities** | Pointer to [**[]MicrosoftGraphUserActivity**](microsoft.graph.userActivity.md) |  | [optional] 
**JoinedTeams** | Pointer to [**[]MicrosoftGraphGroup**](microsoft.graph.group.md) |  | [optional] 

## Methods

### GetAccountEnabled

`func (o *User) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *User) GetAccountEnabledOk() (bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountEnabled

`func (o *User) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### SetAccountEnabled

`func (o *User) SetAccountEnabled(v bool)`

SetAccountEnabled gets a reference to the given bool and assigns it to the AccountEnabled field.

### SetAccountEnabledExplicitNull

`func (o *User) SetAccountEnabledExplicitNull(b bool)`

SetAccountEnabledExplicitNull (un)sets AccountEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AccountEnabled value is set to nil even if false is passed
### GetAgeGroup

`func (o *User) GetAgeGroup() string`

GetAgeGroup returns the AgeGroup field if non-nil, zero value otherwise.

### GetAgeGroupOk

`func (o *User) GetAgeGroupOk() (string, bool)`

GetAgeGroupOk returns a tuple with the AgeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAgeGroup

`func (o *User) HasAgeGroup() bool`

HasAgeGroup returns a boolean if a field has been set.

### SetAgeGroup

`func (o *User) SetAgeGroup(v string)`

SetAgeGroup gets a reference to the given string and assigns it to the AgeGroup field.

### SetAgeGroupExplicitNull

`func (o *User) SetAgeGroupExplicitNull(b bool)`

SetAgeGroupExplicitNull (un)sets AgeGroup to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AgeGroup value is set to nil even if false is passed
### GetAssignedLicenses

`func (o *User) GetAssignedLicenses() []MicrosoftGraphAssignedLicense`

GetAssignedLicenses returns the AssignedLicenses field if non-nil, zero value otherwise.

### GetAssignedLicensesOk

`func (o *User) GetAssignedLicensesOk() ([]MicrosoftGraphAssignedLicense, bool)`

GetAssignedLicensesOk returns a tuple with the AssignedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedLicenses

`func (o *User) HasAssignedLicenses() bool`

HasAssignedLicenses returns a boolean if a field has been set.

### SetAssignedLicenses

`func (o *User) SetAssignedLicenses(v []MicrosoftGraphAssignedLicense)`

SetAssignedLicenses gets a reference to the given []MicrosoftGraphAssignedLicense and assigns it to the AssignedLicenses field.

### GetAssignedPlans

`func (o *User) GetAssignedPlans() []MicrosoftGraphAssignedPlan`

GetAssignedPlans returns the AssignedPlans field if non-nil, zero value otherwise.

### GetAssignedPlansOk

`func (o *User) GetAssignedPlansOk() ([]MicrosoftGraphAssignedPlan, bool)`

GetAssignedPlansOk returns a tuple with the AssignedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedPlans

`func (o *User) HasAssignedPlans() bool`

HasAssignedPlans returns a boolean if a field has been set.

### SetAssignedPlans

`func (o *User) SetAssignedPlans(v []MicrosoftGraphAssignedPlan)`

SetAssignedPlans gets a reference to the given []MicrosoftGraphAssignedPlan and assigns it to the AssignedPlans field.

### GetBusinessPhones

`func (o *User) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *User) GetBusinessPhonesOk() ([]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessPhones

`func (o *User) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### SetBusinessPhones

`func (o *User) SetBusinessPhones(v []string)`

SetBusinessPhones gets a reference to the given []string and assigns it to the BusinessPhones field.

### GetCity

`func (o *User) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *User) GetCityOk() (string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCity

`func (o *User) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCity

`func (o *User) SetCity(v string)`

SetCity gets a reference to the given string and assigns it to the City field.

### SetCityExplicitNull

`func (o *User) SetCityExplicitNull(b bool)`

SetCityExplicitNull (un)sets City to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The City value is set to nil even if false is passed
### GetCompanyName

`func (o *User) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *User) GetCompanyNameOk() (string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompanyName

`func (o *User) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyName

`func (o *User) SetCompanyName(v string)`

SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.

### SetCompanyNameExplicitNull

`func (o *User) SetCompanyNameExplicitNull(b bool)`

SetCompanyNameExplicitNull (un)sets CompanyName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CompanyName value is set to nil even if false is passed
### GetConsentProvidedForMinor

`func (o *User) GetConsentProvidedForMinor() string`

GetConsentProvidedForMinor returns the ConsentProvidedForMinor field if non-nil, zero value otherwise.

### GetConsentProvidedForMinorOk

`func (o *User) GetConsentProvidedForMinorOk() (string, bool)`

GetConsentProvidedForMinorOk returns a tuple with the ConsentProvidedForMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConsentProvidedForMinor

`func (o *User) HasConsentProvidedForMinor() bool`

HasConsentProvidedForMinor returns a boolean if a field has been set.

### SetConsentProvidedForMinor

`func (o *User) SetConsentProvidedForMinor(v string)`

SetConsentProvidedForMinor gets a reference to the given string and assigns it to the ConsentProvidedForMinor field.

### SetConsentProvidedForMinorExplicitNull

`func (o *User) SetConsentProvidedForMinorExplicitNull(b bool)`

SetConsentProvidedForMinorExplicitNull (un)sets ConsentProvidedForMinor to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConsentProvidedForMinor value is set to nil even if false is passed
### GetCountry

`func (o *User) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *User) GetCountryOk() (string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCountry

`func (o *User) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountry

`func (o *User) SetCountry(v string)`

SetCountry gets a reference to the given string and assigns it to the Country field.

### SetCountryExplicitNull

`func (o *User) SetCountryExplicitNull(b bool)`

SetCountryExplicitNull (un)sets Country to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Country value is set to nil even if false is passed
### GetDepartment

`func (o *User) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *User) GetDepartmentOk() (string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDepartment

`func (o *User) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartment

`func (o *User) SetDepartment(v string)`

SetDepartment gets a reference to the given string and assigns it to the Department field.

### SetDepartmentExplicitNull

`func (o *User) SetDepartmentExplicitNull(b bool)`

SetDepartmentExplicitNull (un)sets Department to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Department value is set to nil even if false is passed
### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *User) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetEmployeeId

`func (o *User) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *User) GetEmployeeIdOk() (string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmployeeId

`func (o *User) HasEmployeeId() bool`

HasEmployeeId returns a boolean if a field has been set.

### SetEmployeeId

`func (o *User) SetEmployeeId(v string)`

SetEmployeeId gets a reference to the given string and assigns it to the EmployeeId field.

### SetEmployeeIdExplicitNull

`func (o *User) SetEmployeeIdExplicitNull(b bool)`

SetEmployeeIdExplicitNull (un)sets EmployeeId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EmployeeId value is set to nil even if false is passed
### GetFaxNumber

`func (o *User) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *User) GetFaxNumberOk() (string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFaxNumber

`func (o *User) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

### SetFaxNumber

`func (o *User) SetFaxNumber(v string)`

SetFaxNumber gets a reference to the given string and assigns it to the FaxNumber field.

### SetFaxNumberExplicitNull

`func (o *User) SetFaxNumberExplicitNull(b bool)`

SetFaxNumberExplicitNull (un)sets FaxNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FaxNumber value is set to nil even if false is passed
### GetGivenName

`func (o *User) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *User) GetGivenNameOk() (string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGivenName

`func (o *User) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenName

`func (o *User) SetGivenName(v string)`

SetGivenName gets a reference to the given string and assigns it to the GivenName field.

### SetGivenNameExplicitNull

`func (o *User) SetGivenNameExplicitNull(b bool)`

SetGivenNameExplicitNull (un)sets GivenName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GivenName value is set to nil even if false is passed
### GetImAddresses

`func (o *User) GetImAddresses() []string`

GetImAddresses returns the ImAddresses field if non-nil, zero value otherwise.

### GetImAddressesOk

`func (o *User) GetImAddressesOk() ([]string, bool)`

GetImAddressesOk returns a tuple with the ImAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImAddresses

`func (o *User) HasImAddresses() bool`

HasImAddresses returns a boolean if a field has been set.

### SetImAddresses

`func (o *User) SetImAddresses(v []string)`

SetImAddresses gets a reference to the given []string and assigns it to the ImAddresses field.

### GetIsResourceAccount

`func (o *User) GetIsResourceAccount() bool`

GetIsResourceAccount returns the IsResourceAccount field if non-nil, zero value otherwise.

### GetIsResourceAccountOk

`func (o *User) GetIsResourceAccountOk() (bool, bool)`

GetIsResourceAccountOk returns a tuple with the IsResourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsResourceAccount

`func (o *User) HasIsResourceAccount() bool`

HasIsResourceAccount returns a boolean if a field has been set.

### SetIsResourceAccount

`func (o *User) SetIsResourceAccount(v bool)`

SetIsResourceAccount gets a reference to the given bool and assigns it to the IsResourceAccount field.

### SetIsResourceAccountExplicitNull

`func (o *User) SetIsResourceAccountExplicitNull(b bool)`

SetIsResourceAccountExplicitNull (un)sets IsResourceAccount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsResourceAccount value is set to nil even if false is passed
### GetJobTitle

`func (o *User) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *User) GetJobTitleOk() (string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJobTitle

`func (o *User) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitle

`func (o *User) SetJobTitle(v string)`

SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.

### SetJobTitleExplicitNull

`func (o *User) SetJobTitleExplicitNull(b bool)`

SetJobTitleExplicitNull (un)sets JobTitle to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The JobTitle value is set to nil even if false is passed
### GetLegalAgeGroupClassification

`func (o *User) GetLegalAgeGroupClassification() string`

GetLegalAgeGroupClassification returns the LegalAgeGroupClassification field if non-nil, zero value otherwise.

### GetLegalAgeGroupClassificationOk

`func (o *User) GetLegalAgeGroupClassificationOk() (string, bool)`

GetLegalAgeGroupClassificationOk returns a tuple with the LegalAgeGroupClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLegalAgeGroupClassification

`func (o *User) HasLegalAgeGroupClassification() bool`

HasLegalAgeGroupClassification returns a boolean if a field has been set.

### SetLegalAgeGroupClassification

`func (o *User) SetLegalAgeGroupClassification(v string)`

SetLegalAgeGroupClassification gets a reference to the given string and assigns it to the LegalAgeGroupClassification field.

### SetLegalAgeGroupClassificationExplicitNull

`func (o *User) SetLegalAgeGroupClassificationExplicitNull(b bool)`

SetLegalAgeGroupClassificationExplicitNull (un)sets LegalAgeGroupClassification to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LegalAgeGroupClassification value is set to nil even if false is passed
### GetLicenseAssignmentStates

`func (o *User) GetLicenseAssignmentStates() []AnyOfmicrosoftGraphLicenseAssignmentState`

GetLicenseAssignmentStates returns the LicenseAssignmentStates field if non-nil, zero value otherwise.

### GetLicenseAssignmentStatesOk

`func (o *User) GetLicenseAssignmentStatesOk() ([]AnyOfmicrosoftGraphLicenseAssignmentState, bool)`

GetLicenseAssignmentStatesOk returns a tuple with the LicenseAssignmentStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicenseAssignmentStates

`func (o *User) HasLicenseAssignmentStates() bool`

HasLicenseAssignmentStates returns a boolean if a field has been set.

### SetLicenseAssignmentStates

`func (o *User) SetLicenseAssignmentStates(v []AnyOfmicrosoftGraphLicenseAssignmentState)`

SetLicenseAssignmentStates gets a reference to the given []AnyOfmicrosoftGraphLicenseAssignmentState and assigns it to the LicenseAssignmentStates field.

### GetMail

`func (o *User) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *User) GetMailOk() (string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMail

`func (o *User) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMail

`func (o *User) SetMail(v string)`

SetMail gets a reference to the given string and assigns it to the Mail field.

### SetMailExplicitNull

`func (o *User) SetMailExplicitNull(b bool)`

SetMailExplicitNull (un)sets Mail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Mail value is set to nil even if false is passed
### GetMailNickname

`func (o *User) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *User) GetMailNicknameOk() (string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailNickname

`func (o *User) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNickname

`func (o *User) SetMailNickname(v string)`

SetMailNickname gets a reference to the given string and assigns it to the MailNickname field.

### SetMailNicknameExplicitNull

`func (o *User) SetMailNicknameExplicitNull(b bool)`

SetMailNicknameExplicitNull (un)sets MailNickname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailNickname value is set to nil even if false is passed
### GetMobilePhone

`func (o *User) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *User) GetMobilePhoneOk() (string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobilePhone

`func (o *User) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhone

`func (o *User) SetMobilePhone(v string)`

SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.

### SetMobilePhoneExplicitNull

`func (o *User) SetMobilePhoneExplicitNull(b bool)`

SetMobilePhoneExplicitNull (un)sets MobilePhone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MobilePhone value is set to nil even if false is passed
### GetOnPremisesDistinguishedName

`func (o *User) GetOnPremisesDistinguishedName() string`

GetOnPremisesDistinguishedName returns the OnPremisesDistinguishedName field if non-nil, zero value otherwise.

### GetOnPremisesDistinguishedNameOk

`func (o *User) GetOnPremisesDistinguishedNameOk() (string, bool)`

GetOnPremisesDistinguishedNameOk returns a tuple with the OnPremisesDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesDistinguishedName

`func (o *User) HasOnPremisesDistinguishedName() bool`

HasOnPremisesDistinguishedName returns a boolean if a field has been set.

### SetOnPremisesDistinguishedName

`func (o *User) SetOnPremisesDistinguishedName(v string)`

SetOnPremisesDistinguishedName gets a reference to the given string and assigns it to the OnPremisesDistinguishedName field.

### SetOnPremisesDistinguishedNameExplicitNull

`func (o *User) SetOnPremisesDistinguishedNameExplicitNull(b bool)`

SetOnPremisesDistinguishedNameExplicitNull (un)sets OnPremisesDistinguishedName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesDistinguishedName value is set to nil even if false is passed
### GetOnPremisesExtensionAttributes

`func (o *User) GetOnPremisesExtensionAttributes() AnyOfmicrosoftGraphOnPremisesExtensionAttributes`

GetOnPremisesExtensionAttributes returns the OnPremisesExtensionAttributes field if non-nil, zero value otherwise.

### GetOnPremisesExtensionAttributesOk

`func (o *User) GetOnPremisesExtensionAttributesOk() (AnyOfmicrosoftGraphOnPremisesExtensionAttributes, bool)`

GetOnPremisesExtensionAttributesOk returns a tuple with the OnPremisesExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesExtensionAttributes

`func (o *User) HasOnPremisesExtensionAttributes() bool`

HasOnPremisesExtensionAttributes returns a boolean if a field has been set.

### SetOnPremisesExtensionAttributes

`func (o *User) SetOnPremisesExtensionAttributes(v AnyOfmicrosoftGraphOnPremisesExtensionAttributes)`

SetOnPremisesExtensionAttributes gets a reference to the given AnyOfmicrosoftGraphOnPremisesExtensionAttributes and assigns it to the OnPremisesExtensionAttributes field.

### SetOnPremisesExtensionAttributesExplicitNull

`func (o *User) SetOnPremisesExtensionAttributesExplicitNull(b bool)`

SetOnPremisesExtensionAttributesExplicitNull (un)sets OnPremisesExtensionAttributes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesExtensionAttributes value is set to nil even if false is passed
### GetOnPremisesImmutableId

`func (o *User) GetOnPremisesImmutableId() string`

GetOnPremisesImmutableId returns the OnPremisesImmutableId field if non-nil, zero value otherwise.

### GetOnPremisesImmutableIdOk

`func (o *User) GetOnPremisesImmutableIdOk() (string, bool)`

GetOnPremisesImmutableIdOk returns a tuple with the OnPremisesImmutableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesImmutableId

`func (o *User) HasOnPremisesImmutableId() bool`

HasOnPremisesImmutableId returns a boolean if a field has been set.

### SetOnPremisesImmutableId

`func (o *User) SetOnPremisesImmutableId(v string)`

SetOnPremisesImmutableId gets a reference to the given string and assigns it to the OnPremisesImmutableId field.

### SetOnPremisesImmutableIdExplicitNull

`func (o *User) SetOnPremisesImmutableIdExplicitNull(b bool)`

SetOnPremisesImmutableIdExplicitNull (un)sets OnPremisesImmutableId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesImmutableId value is set to nil even if false is passed
### GetOnPremisesLastSyncDateTime

`func (o *User) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *User) GetOnPremisesLastSyncDateTimeOk() (time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesLastSyncDateTime

`func (o *User) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTime

`func (o *User) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime gets a reference to the given time.Time and assigns it to the OnPremisesLastSyncDateTime field.

### SetOnPremisesLastSyncDateTimeExplicitNull

`func (o *User) SetOnPremisesLastSyncDateTimeExplicitNull(b bool)`

SetOnPremisesLastSyncDateTimeExplicitNull (un)sets OnPremisesLastSyncDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesLastSyncDateTime value is set to nil even if false is passed
### GetOnPremisesProvisioningErrors

`func (o *User) GetOnPremisesProvisioningErrors() []AnyOfmicrosoftGraphOnPremisesProvisioningError`

GetOnPremisesProvisioningErrors returns the OnPremisesProvisioningErrors field if non-nil, zero value otherwise.

### GetOnPremisesProvisioningErrorsOk

`func (o *User) GetOnPremisesProvisioningErrorsOk() ([]AnyOfmicrosoftGraphOnPremisesProvisioningError, bool)`

GetOnPremisesProvisioningErrorsOk returns a tuple with the OnPremisesProvisioningErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesProvisioningErrors

`func (o *User) HasOnPremisesProvisioningErrors() bool`

HasOnPremisesProvisioningErrors returns a boolean if a field has been set.

### SetOnPremisesProvisioningErrors

`func (o *User) SetOnPremisesProvisioningErrors(v []AnyOfmicrosoftGraphOnPremisesProvisioningError)`

SetOnPremisesProvisioningErrors gets a reference to the given []AnyOfmicrosoftGraphOnPremisesProvisioningError and assigns it to the OnPremisesProvisioningErrors field.

### GetOnPremisesSecurityIdentifier

`func (o *User) GetOnPremisesSecurityIdentifier() string`

GetOnPremisesSecurityIdentifier returns the OnPremisesSecurityIdentifier field if non-nil, zero value otherwise.

### GetOnPremisesSecurityIdentifierOk

`func (o *User) GetOnPremisesSecurityIdentifierOk() (string, bool)`

GetOnPremisesSecurityIdentifierOk returns a tuple with the OnPremisesSecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSecurityIdentifier

`func (o *User) HasOnPremisesSecurityIdentifier() bool`

HasOnPremisesSecurityIdentifier returns a boolean if a field has been set.

### SetOnPremisesSecurityIdentifier

`func (o *User) SetOnPremisesSecurityIdentifier(v string)`

SetOnPremisesSecurityIdentifier gets a reference to the given string and assigns it to the OnPremisesSecurityIdentifier field.

### SetOnPremisesSecurityIdentifierExplicitNull

`func (o *User) SetOnPremisesSecurityIdentifierExplicitNull(b bool)`

SetOnPremisesSecurityIdentifierExplicitNull (un)sets OnPremisesSecurityIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSecurityIdentifier value is set to nil even if false is passed
### GetOnPremisesSyncEnabled

`func (o *User) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *User) GetOnPremisesSyncEnabledOk() (bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSyncEnabled

`func (o *User) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabled

`func (o *User) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled gets a reference to the given bool and assigns it to the OnPremisesSyncEnabled field.

### SetOnPremisesSyncEnabledExplicitNull

`func (o *User) SetOnPremisesSyncEnabledExplicitNull(b bool)`

SetOnPremisesSyncEnabledExplicitNull (un)sets OnPremisesSyncEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSyncEnabled value is set to nil even if false is passed
### GetOnPremisesDomainName

`func (o *User) GetOnPremisesDomainName() string`

GetOnPremisesDomainName returns the OnPremisesDomainName field if non-nil, zero value otherwise.

### GetOnPremisesDomainNameOk

`func (o *User) GetOnPremisesDomainNameOk() (string, bool)`

GetOnPremisesDomainNameOk returns a tuple with the OnPremisesDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesDomainName

`func (o *User) HasOnPremisesDomainName() bool`

HasOnPremisesDomainName returns a boolean if a field has been set.

### SetOnPremisesDomainName

`func (o *User) SetOnPremisesDomainName(v string)`

SetOnPremisesDomainName gets a reference to the given string and assigns it to the OnPremisesDomainName field.

### SetOnPremisesDomainNameExplicitNull

`func (o *User) SetOnPremisesDomainNameExplicitNull(b bool)`

SetOnPremisesDomainNameExplicitNull (un)sets OnPremisesDomainName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesDomainName value is set to nil even if false is passed
### GetOnPremisesSamAccountName

`func (o *User) GetOnPremisesSamAccountName() string`

GetOnPremisesSamAccountName returns the OnPremisesSamAccountName field if non-nil, zero value otherwise.

### GetOnPremisesSamAccountNameOk

`func (o *User) GetOnPremisesSamAccountNameOk() (string, bool)`

GetOnPremisesSamAccountNameOk returns a tuple with the OnPremisesSamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSamAccountName

`func (o *User) HasOnPremisesSamAccountName() bool`

HasOnPremisesSamAccountName returns a boolean if a field has been set.

### SetOnPremisesSamAccountName

`func (o *User) SetOnPremisesSamAccountName(v string)`

SetOnPremisesSamAccountName gets a reference to the given string and assigns it to the OnPremisesSamAccountName field.

### SetOnPremisesSamAccountNameExplicitNull

`func (o *User) SetOnPremisesSamAccountNameExplicitNull(b bool)`

SetOnPremisesSamAccountNameExplicitNull (un)sets OnPremisesSamAccountName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSamAccountName value is set to nil even if false is passed
### GetOnPremisesUserPrincipalName

`func (o *User) GetOnPremisesUserPrincipalName() string`

GetOnPremisesUserPrincipalName returns the OnPremisesUserPrincipalName field if non-nil, zero value otherwise.

### GetOnPremisesUserPrincipalNameOk

`func (o *User) GetOnPremisesUserPrincipalNameOk() (string, bool)`

GetOnPremisesUserPrincipalNameOk returns a tuple with the OnPremisesUserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesUserPrincipalName

`func (o *User) HasOnPremisesUserPrincipalName() bool`

HasOnPremisesUserPrincipalName returns a boolean if a field has been set.

### SetOnPremisesUserPrincipalName

`func (o *User) SetOnPremisesUserPrincipalName(v string)`

SetOnPremisesUserPrincipalName gets a reference to the given string and assigns it to the OnPremisesUserPrincipalName field.

### SetOnPremisesUserPrincipalNameExplicitNull

`func (o *User) SetOnPremisesUserPrincipalNameExplicitNull(b bool)`

SetOnPremisesUserPrincipalNameExplicitNull (un)sets OnPremisesUserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesUserPrincipalName value is set to nil even if false is passed
### GetOtherMails

`func (o *User) GetOtherMails() []string`

GetOtherMails returns the OtherMails field if non-nil, zero value otherwise.

### GetOtherMailsOk

`func (o *User) GetOtherMailsOk() ([]string, bool)`

GetOtherMailsOk returns a tuple with the OtherMails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOtherMails

`func (o *User) HasOtherMails() bool`

HasOtherMails returns a boolean if a field has been set.

### SetOtherMails

`func (o *User) SetOtherMails(v []string)`

SetOtherMails gets a reference to the given []string and assigns it to the OtherMails field.

### GetPasswordPolicies

`func (o *User) GetPasswordPolicies() string`

GetPasswordPolicies returns the PasswordPolicies field if non-nil, zero value otherwise.

### GetPasswordPoliciesOk

`func (o *User) GetPasswordPoliciesOk() (string, bool)`

GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPolicies

`func (o *User) HasPasswordPolicies() bool`

HasPasswordPolicies returns a boolean if a field has been set.

### SetPasswordPolicies

`func (o *User) SetPasswordPolicies(v string)`

SetPasswordPolicies gets a reference to the given string and assigns it to the PasswordPolicies field.

### SetPasswordPoliciesExplicitNull

`func (o *User) SetPasswordPoliciesExplicitNull(b bool)`

SetPasswordPoliciesExplicitNull (un)sets PasswordPolicies to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPolicies value is set to nil even if false is passed
### GetPasswordProfile

`func (o *User) GetPasswordProfile() AnyOfmicrosoftGraphPasswordProfile`

GetPasswordProfile returns the PasswordProfile field if non-nil, zero value otherwise.

### GetPasswordProfileOk

`func (o *User) GetPasswordProfileOk() (AnyOfmicrosoftGraphPasswordProfile, bool)`

GetPasswordProfileOk returns a tuple with the PasswordProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordProfile

`func (o *User) HasPasswordProfile() bool`

HasPasswordProfile returns a boolean if a field has been set.

### SetPasswordProfile

`func (o *User) SetPasswordProfile(v AnyOfmicrosoftGraphPasswordProfile)`

SetPasswordProfile gets a reference to the given AnyOfmicrosoftGraphPasswordProfile and assigns it to the PasswordProfile field.

### SetPasswordProfileExplicitNull

`func (o *User) SetPasswordProfileExplicitNull(b bool)`

SetPasswordProfileExplicitNull (un)sets PasswordProfile to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordProfile value is set to nil even if false is passed
### GetOfficeLocation

`func (o *User) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *User) GetOfficeLocationOk() (string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOfficeLocation

`func (o *User) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocation

`func (o *User) SetOfficeLocation(v string)`

SetOfficeLocation gets a reference to the given string and assigns it to the OfficeLocation field.

### SetOfficeLocationExplicitNull

`func (o *User) SetOfficeLocationExplicitNull(b bool)`

SetOfficeLocationExplicitNull (un)sets OfficeLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OfficeLocation value is set to nil even if false is passed
### GetPostalCode

`func (o *User) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *User) GetPostalCodeOk() (string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPostalCode

`func (o *User) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCode

`func (o *User) SetPostalCode(v string)`

SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.

### SetPostalCodeExplicitNull

`func (o *User) SetPostalCodeExplicitNull(b bool)`

SetPostalCodeExplicitNull (un)sets PostalCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PostalCode value is set to nil even if false is passed
### GetPreferredLanguage

`func (o *User) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *User) GetPreferredLanguageOk() (string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreferredLanguage

`func (o *User) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguage

`func (o *User) SetPreferredLanguage(v string)`

SetPreferredLanguage gets a reference to the given string and assigns it to the PreferredLanguage field.

### SetPreferredLanguageExplicitNull

`func (o *User) SetPreferredLanguageExplicitNull(b bool)`

SetPreferredLanguageExplicitNull (un)sets PreferredLanguage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreferredLanguage value is set to nil even if false is passed
### GetProvisionedPlans

`func (o *User) GetProvisionedPlans() []MicrosoftGraphProvisionedPlan`

GetProvisionedPlans returns the ProvisionedPlans field if non-nil, zero value otherwise.

### GetProvisionedPlansOk

`func (o *User) GetProvisionedPlansOk() ([]MicrosoftGraphProvisionedPlan, bool)`

GetProvisionedPlansOk returns a tuple with the ProvisionedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProvisionedPlans

`func (o *User) HasProvisionedPlans() bool`

HasProvisionedPlans returns a boolean if a field has been set.

### SetProvisionedPlans

`func (o *User) SetProvisionedPlans(v []MicrosoftGraphProvisionedPlan)`

SetProvisionedPlans gets a reference to the given []MicrosoftGraphProvisionedPlan and assigns it to the ProvisionedPlans field.

### GetProxyAddresses

`func (o *User) GetProxyAddresses() []string`

GetProxyAddresses returns the ProxyAddresses field if non-nil, zero value otherwise.

### GetProxyAddressesOk

`func (o *User) GetProxyAddressesOk() ([]string, bool)`

GetProxyAddressesOk returns a tuple with the ProxyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProxyAddresses

`func (o *User) HasProxyAddresses() bool`

HasProxyAddresses returns a boolean if a field has been set.

### SetProxyAddresses

`func (o *User) SetProxyAddresses(v []string)`

SetProxyAddresses gets a reference to the given []string and assigns it to the ProxyAddresses field.

### GetShowInAddressList

`func (o *User) GetShowInAddressList() bool`

GetShowInAddressList returns the ShowInAddressList field if non-nil, zero value otherwise.

### GetShowInAddressListOk

`func (o *User) GetShowInAddressListOk() (bool, bool)`

GetShowInAddressListOk returns a tuple with the ShowInAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowInAddressList

`func (o *User) HasShowInAddressList() bool`

HasShowInAddressList returns a boolean if a field has been set.

### SetShowInAddressList

`func (o *User) SetShowInAddressList(v bool)`

SetShowInAddressList gets a reference to the given bool and assigns it to the ShowInAddressList field.

### SetShowInAddressListExplicitNull

`func (o *User) SetShowInAddressListExplicitNull(b bool)`

SetShowInAddressListExplicitNull (un)sets ShowInAddressList to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ShowInAddressList value is set to nil even if false is passed
### GetSignInSessionsValidFromDateTime

`func (o *User) GetSignInSessionsValidFromDateTime() time.Time`

GetSignInSessionsValidFromDateTime returns the SignInSessionsValidFromDateTime field if non-nil, zero value otherwise.

### GetSignInSessionsValidFromDateTimeOk

`func (o *User) GetSignInSessionsValidFromDateTimeOk() (time.Time, bool)`

GetSignInSessionsValidFromDateTimeOk returns a tuple with the SignInSessionsValidFromDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSignInSessionsValidFromDateTime

`func (o *User) HasSignInSessionsValidFromDateTime() bool`

HasSignInSessionsValidFromDateTime returns a boolean if a field has been set.

### SetSignInSessionsValidFromDateTime

`func (o *User) SetSignInSessionsValidFromDateTime(v time.Time)`

SetSignInSessionsValidFromDateTime gets a reference to the given time.Time and assigns it to the SignInSessionsValidFromDateTime field.

### SetSignInSessionsValidFromDateTimeExplicitNull

`func (o *User) SetSignInSessionsValidFromDateTimeExplicitNull(b bool)`

SetSignInSessionsValidFromDateTimeExplicitNull (un)sets SignInSessionsValidFromDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SignInSessionsValidFromDateTime value is set to nil even if false is passed
### GetState

`func (o *User) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *User) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *User) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### SetStateExplicitNull

`func (o *User) SetStateExplicitNull(b bool)`

SetStateExplicitNull (un)sets State to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The State value is set to nil even if false is passed
### GetStreetAddress

`func (o *User) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *User) GetStreetAddressOk() (string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStreetAddress

`func (o *User) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddress

`func (o *User) SetStreetAddress(v string)`

SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.

### SetStreetAddressExplicitNull

`func (o *User) SetStreetAddressExplicitNull(b bool)`

SetStreetAddressExplicitNull (un)sets StreetAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StreetAddress value is set to nil even if false is passed
### GetSurname

`func (o *User) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *User) GetSurnameOk() (string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSurname

`func (o *User) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurname

`func (o *User) SetSurname(v string)`

SetSurname gets a reference to the given string and assigns it to the Surname field.

### SetSurnameExplicitNull

`func (o *User) SetSurnameExplicitNull(b bool)`

SetSurnameExplicitNull (un)sets Surname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Surname value is set to nil even if false is passed
### GetUsageLocation

`func (o *User) GetUsageLocation() string`

GetUsageLocation returns the UsageLocation field if non-nil, zero value otherwise.

### GetUsageLocationOk

`func (o *User) GetUsageLocationOk() (string, bool)`

GetUsageLocationOk returns a tuple with the UsageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageLocation

`func (o *User) HasUsageLocation() bool`

HasUsageLocation returns a boolean if a field has been set.

### SetUsageLocation

`func (o *User) SetUsageLocation(v string)`

SetUsageLocation gets a reference to the given string and assigns it to the UsageLocation field.

### SetUsageLocationExplicitNull

`func (o *User) SetUsageLocationExplicitNull(b bool)`

SetUsageLocationExplicitNull (un)sets UsageLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UsageLocation value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *User) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *User) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *User) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *User) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *User) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetUserType

`func (o *User) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *User) GetUserTypeOk() (string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserType

`func (o *User) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserType

`func (o *User) SetUserType(v string)`

SetUserType gets a reference to the given string and assigns it to the UserType field.

### SetUserTypeExplicitNull

`func (o *User) SetUserTypeExplicitNull(b bool)`

SetUserTypeExplicitNull (un)sets UserType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserType value is set to nil even if false is passed
### GetMailboxSettings

`func (o *User) GetMailboxSettings() AnyOfmicrosoftGraphMailboxSettings`

GetMailboxSettings returns the MailboxSettings field if non-nil, zero value otherwise.

### GetMailboxSettingsOk

`func (o *User) GetMailboxSettingsOk() (AnyOfmicrosoftGraphMailboxSettings, bool)`

GetMailboxSettingsOk returns a tuple with the MailboxSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailboxSettings

`func (o *User) HasMailboxSettings() bool`

HasMailboxSettings returns a boolean if a field has been set.

### SetMailboxSettings

`func (o *User) SetMailboxSettings(v AnyOfmicrosoftGraphMailboxSettings)`

SetMailboxSettings gets a reference to the given AnyOfmicrosoftGraphMailboxSettings and assigns it to the MailboxSettings field.

### SetMailboxSettingsExplicitNull

`func (o *User) SetMailboxSettingsExplicitNull(b bool)`

SetMailboxSettingsExplicitNull (un)sets MailboxSettings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailboxSettings value is set to nil even if false is passed
### GetDeviceEnrollmentLimit

`func (o *User) GetDeviceEnrollmentLimit() int32`

GetDeviceEnrollmentLimit returns the DeviceEnrollmentLimit field if non-nil, zero value otherwise.

### GetDeviceEnrollmentLimitOk

`func (o *User) GetDeviceEnrollmentLimitOk() (int32, bool)`

GetDeviceEnrollmentLimitOk returns a tuple with the DeviceEnrollmentLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceEnrollmentLimit

`func (o *User) HasDeviceEnrollmentLimit() bool`

HasDeviceEnrollmentLimit returns a boolean if a field has been set.

### SetDeviceEnrollmentLimit

`func (o *User) SetDeviceEnrollmentLimit(v int32)`

SetDeviceEnrollmentLimit gets a reference to the given int32 and assigns it to the DeviceEnrollmentLimit field.

### GetAboutMe

`func (o *User) GetAboutMe() string`

GetAboutMe returns the AboutMe field if non-nil, zero value otherwise.

### GetAboutMeOk

`func (o *User) GetAboutMeOk() (string, bool)`

GetAboutMeOk returns a tuple with the AboutMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAboutMe

`func (o *User) HasAboutMe() bool`

HasAboutMe returns a boolean if a field has been set.

### SetAboutMe

`func (o *User) SetAboutMe(v string)`

SetAboutMe gets a reference to the given string and assigns it to the AboutMe field.

### SetAboutMeExplicitNull

`func (o *User) SetAboutMeExplicitNull(b bool)`

SetAboutMeExplicitNull (un)sets AboutMe to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AboutMe value is set to nil even if false is passed
### GetBirthday

`func (o *User) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *User) GetBirthdayOk() (time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBirthday

`func (o *User) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthday

`func (o *User) SetBirthday(v time.Time)`

SetBirthday gets a reference to the given time.Time and assigns it to the Birthday field.

### GetHireDate

`func (o *User) GetHireDate() time.Time`

GetHireDate returns the HireDate field if non-nil, zero value otherwise.

### GetHireDateOk

`func (o *User) GetHireDateOk() (time.Time, bool)`

GetHireDateOk returns a tuple with the HireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHireDate

`func (o *User) HasHireDate() bool`

HasHireDate returns a boolean if a field has been set.

### SetHireDate

`func (o *User) SetHireDate(v time.Time)`

SetHireDate gets a reference to the given time.Time and assigns it to the HireDate field.

### GetInterests

`func (o *User) GetInterests() []string`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *User) GetInterestsOk() ([]string, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInterests

`func (o *User) HasInterests() bool`

HasInterests returns a boolean if a field has been set.

### SetInterests

`func (o *User) SetInterests(v []string)`

SetInterests gets a reference to the given []string and assigns it to the Interests field.

### GetMySite

`func (o *User) GetMySite() string`

GetMySite returns the MySite field if non-nil, zero value otherwise.

### GetMySiteOk

`func (o *User) GetMySiteOk() (string, bool)`

GetMySiteOk returns a tuple with the MySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMySite

`func (o *User) HasMySite() bool`

HasMySite returns a boolean if a field has been set.

### SetMySite

`func (o *User) SetMySite(v string)`

SetMySite gets a reference to the given string and assigns it to the MySite field.

### SetMySiteExplicitNull

`func (o *User) SetMySiteExplicitNull(b bool)`

SetMySiteExplicitNull (un)sets MySite to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MySite value is set to nil even if false is passed
### GetPastProjects

`func (o *User) GetPastProjects() []string`

GetPastProjects returns the PastProjects field if non-nil, zero value otherwise.

### GetPastProjectsOk

`func (o *User) GetPastProjectsOk() ([]string, bool)`

GetPastProjectsOk returns a tuple with the PastProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPastProjects

`func (o *User) HasPastProjects() bool`

HasPastProjects returns a boolean if a field has been set.

### SetPastProjects

`func (o *User) SetPastProjects(v []string)`

SetPastProjects gets a reference to the given []string and assigns it to the PastProjects field.

### GetPreferredName

`func (o *User) GetPreferredName() string`

GetPreferredName returns the PreferredName field if non-nil, zero value otherwise.

### GetPreferredNameOk

`func (o *User) GetPreferredNameOk() (string, bool)`

GetPreferredNameOk returns a tuple with the PreferredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreferredName

`func (o *User) HasPreferredName() bool`

HasPreferredName returns a boolean if a field has been set.

### SetPreferredName

`func (o *User) SetPreferredName(v string)`

SetPreferredName gets a reference to the given string and assigns it to the PreferredName field.

### SetPreferredNameExplicitNull

`func (o *User) SetPreferredNameExplicitNull(b bool)`

SetPreferredNameExplicitNull (un)sets PreferredName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreferredName value is set to nil even if false is passed
### GetResponsibilities

`func (o *User) GetResponsibilities() []string`

GetResponsibilities returns the Responsibilities field if non-nil, zero value otherwise.

### GetResponsibilitiesOk

`func (o *User) GetResponsibilitiesOk() ([]string, bool)`

GetResponsibilitiesOk returns a tuple with the Responsibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponsibilities

`func (o *User) HasResponsibilities() bool`

HasResponsibilities returns a boolean if a field has been set.

### SetResponsibilities

`func (o *User) SetResponsibilities(v []string)`

SetResponsibilities gets a reference to the given []string and assigns it to the Responsibilities field.

### GetSchools

`func (o *User) GetSchools() []string`

GetSchools returns the Schools field if non-nil, zero value otherwise.

### GetSchoolsOk

`func (o *User) GetSchoolsOk() ([]string, bool)`

GetSchoolsOk returns a tuple with the Schools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchools

`func (o *User) HasSchools() bool`

HasSchools returns a boolean if a field has been set.

### SetSchools

`func (o *User) SetSchools(v []string)`

SetSchools gets a reference to the given []string and assigns it to the Schools field.

### GetSkills

`func (o *User) GetSkills() []string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *User) GetSkillsOk() ([]string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSkills

`func (o *User) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### SetSkills

`func (o *User) SetSkills(v []string)`

SetSkills gets a reference to the given []string and assigns it to the Skills field.

### GetOwnedDevices

`func (o *User) GetOwnedDevices() []MicrosoftGraphDirectoryObject`

GetOwnedDevices returns the OwnedDevices field if non-nil, zero value otherwise.

### GetOwnedDevicesOk

`func (o *User) GetOwnedDevicesOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetOwnedDevicesOk returns a tuple with the OwnedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwnedDevices

`func (o *User) HasOwnedDevices() bool`

HasOwnedDevices returns a boolean if a field has been set.

### SetOwnedDevices

`func (o *User) SetOwnedDevices(v []MicrosoftGraphDirectoryObject)`

SetOwnedDevices gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the OwnedDevices field.

### GetRegisteredDevices

`func (o *User) GetRegisteredDevices() []MicrosoftGraphDirectoryObject`

GetRegisteredDevices returns the RegisteredDevices field if non-nil, zero value otherwise.

### GetRegisteredDevicesOk

`func (o *User) GetRegisteredDevicesOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetRegisteredDevicesOk returns a tuple with the RegisteredDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRegisteredDevices

`func (o *User) HasRegisteredDevices() bool`

HasRegisteredDevices returns a boolean if a field has been set.

### SetRegisteredDevices

`func (o *User) SetRegisteredDevices(v []MicrosoftGraphDirectoryObject)`

SetRegisteredDevices gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the RegisteredDevices field.

### GetManager

`func (o *User) GetManager() AnyOfmicrosoftGraphDirectoryObject`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *User) GetManagerOk() (AnyOfmicrosoftGraphDirectoryObject, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManager

`func (o *User) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManager

`func (o *User) SetManager(v AnyOfmicrosoftGraphDirectoryObject)`

SetManager gets a reference to the given AnyOfmicrosoftGraphDirectoryObject and assigns it to the Manager field.

### SetManagerExplicitNull

`func (o *User) SetManagerExplicitNull(b bool)`

SetManagerExplicitNull (un)sets Manager to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Manager value is set to nil even if false is passed
### GetDirectReports

`func (o *User) GetDirectReports() []MicrosoftGraphDirectoryObject`

GetDirectReports returns the DirectReports field if non-nil, zero value otherwise.

### GetDirectReportsOk

`func (o *User) GetDirectReportsOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetDirectReportsOk returns a tuple with the DirectReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirectReports

`func (o *User) HasDirectReports() bool`

HasDirectReports returns a boolean if a field has been set.

### SetDirectReports

`func (o *User) SetDirectReports(v []MicrosoftGraphDirectoryObject)`

SetDirectReports gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the DirectReports field.

### GetMemberOf

`func (o *User) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *User) GetMemberOfOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMemberOf

`func (o *User) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### SetMemberOf

`func (o *User) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the MemberOf field.

### GetCreatedObjects

`func (o *User) GetCreatedObjects() []MicrosoftGraphDirectoryObject`

GetCreatedObjects returns the CreatedObjects field if non-nil, zero value otherwise.

### GetCreatedObjectsOk

`func (o *User) GetCreatedObjectsOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetCreatedObjectsOk returns a tuple with the CreatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedObjects

`func (o *User) HasCreatedObjects() bool`

HasCreatedObjects returns a boolean if a field has been set.

### SetCreatedObjects

`func (o *User) SetCreatedObjects(v []MicrosoftGraphDirectoryObject)`

SetCreatedObjects gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the CreatedObjects field.

### GetOwnedObjects

`func (o *User) GetOwnedObjects() []MicrosoftGraphDirectoryObject`

GetOwnedObjects returns the OwnedObjects field if non-nil, zero value otherwise.

### GetOwnedObjectsOk

`func (o *User) GetOwnedObjectsOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetOwnedObjectsOk returns a tuple with the OwnedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwnedObjects

`func (o *User) HasOwnedObjects() bool`

HasOwnedObjects returns a boolean if a field has been set.

### SetOwnedObjects

`func (o *User) SetOwnedObjects(v []MicrosoftGraphDirectoryObject)`

SetOwnedObjects gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the OwnedObjects field.

### GetLicenseDetails

`func (o *User) GetLicenseDetails() []MicrosoftGraphLicenseDetails`

GetLicenseDetails returns the LicenseDetails field if non-nil, zero value otherwise.

### GetLicenseDetailsOk

`func (o *User) GetLicenseDetailsOk() ([]MicrosoftGraphLicenseDetails, bool)`

GetLicenseDetailsOk returns a tuple with the LicenseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicenseDetails

`func (o *User) HasLicenseDetails() bool`

HasLicenseDetails returns a boolean if a field has been set.

### SetLicenseDetails

`func (o *User) SetLicenseDetails(v []MicrosoftGraphLicenseDetails)`

SetLicenseDetails gets a reference to the given []MicrosoftGraphLicenseDetails and assigns it to the LicenseDetails field.

### GetTransitiveMemberOf

`func (o *User) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *User) GetTransitiveMemberOfOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransitiveMemberOf

`func (o *User) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.

### SetTransitiveMemberOf

`func (o *User) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the TransitiveMemberOf field.

### GetOutlook

`func (o *User) GetOutlook() AnyOfmicrosoftGraphOutlookUser`

GetOutlook returns the Outlook field if non-nil, zero value otherwise.

### GetOutlookOk

`func (o *User) GetOutlookOk() (AnyOfmicrosoftGraphOutlookUser, bool)`

GetOutlookOk returns a tuple with the Outlook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutlook

`func (o *User) HasOutlook() bool`

HasOutlook returns a boolean if a field has been set.

### SetOutlook

`func (o *User) SetOutlook(v AnyOfmicrosoftGraphOutlookUser)`

SetOutlook gets a reference to the given AnyOfmicrosoftGraphOutlookUser and assigns it to the Outlook field.

### SetOutlookExplicitNull

`func (o *User) SetOutlookExplicitNull(b bool)`

SetOutlookExplicitNull (un)sets Outlook to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Outlook value is set to nil even if false is passed
### GetMessages

`func (o *User) GetMessages() []MicrosoftGraphMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *User) GetMessagesOk() ([]MicrosoftGraphMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessages

`func (o *User) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessages

`func (o *User) SetMessages(v []MicrosoftGraphMessage)`

SetMessages gets a reference to the given []MicrosoftGraphMessage and assigns it to the Messages field.

### GetMailFolders

`func (o *User) GetMailFolders() []MicrosoftGraphMailFolder`

GetMailFolders returns the MailFolders field if non-nil, zero value otherwise.

### GetMailFoldersOk

`func (o *User) GetMailFoldersOk() ([]MicrosoftGraphMailFolder, bool)`

GetMailFoldersOk returns a tuple with the MailFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailFolders

`func (o *User) HasMailFolders() bool`

HasMailFolders returns a boolean if a field has been set.

### SetMailFolders

`func (o *User) SetMailFolders(v []MicrosoftGraphMailFolder)`

SetMailFolders gets a reference to the given []MicrosoftGraphMailFolder and assigns it to the MailFolders field.

### GetCalendar

`func (o *User) GetCalendar() AnyOfmicrosoftGraphCalendar`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *User) GetCalendarOk() (AnyOfmicrosoftGraphCalendar, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendar

`func (o *User) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### SetCalendar

`func (o *User) SetCalendar(v AnyOfmicrosoftGraphCalendar)`

SetCalendar gets a reference to the given AnyOfmicrosoftGraphCalendar and assigns it to the Calendar field.

### SetCalendarExplicitNull

`func (o *User) SetCalendarExplicitNull(b bool)`

SetCalendarExplicitNull (un)sets Calendar to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Calendar value is set to nil even if false is passed
### GetCalendars

`func (o *User) GetCalendars() []MicrosoftGraphCalendar`

GetCalendars returns the Calendars field if non-nil, zero value otherwise.

### GetCalendarsOk

`func (o *User) GetCalendarsOk() ([]MicrosoftGraphCalendar, bool)`

GetCalendarsOk returns a tuple with the Calendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendars

`func (o *User) HasCalendars() bool`

HasCalendars returns a boolean if a field has been set.

### SetCalendars

`func (o *User) SetCalendars(v []MicrosoftGraphCalendar)`

SetCalendars gets a reference to the given []MicrosoftGraphCalendar and assigns it to the Calendars field.

### GetCalendarGroups

`func (o *User) GetCalendarGroups() []MicrosoftGraphCalendarGroup`

GetCalendarGroups returns the CalendarGroups field if non-nil, zero value otherwise.

### GetCalendarGroupsOk

`func (o *User) GetCalendarGroupsOk() ([]MicrosoftGraphCalendarGroup, bool)`

GetCalendarGroupsOk returns a tuple with the CalendarGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendarGroups

`func (o *User) HasCalendarGroups() bool`

HasCalendarGroups returns a boolean if a field has been set.

### SetCalendarGroups

`func (o *User) SetCalendarGroups(v []MicrosoftGraphCalendarGroup)`

SetCalendarGroups gets a reference to the given []MicrosoftGraphCalendarGroup and assigns it to the CalendarGroups field.

### GetCalendarView

`func (o *User) GetCalendarView() []MicrosoftGraphEvent`

GetCalendarView returns the CalendarView field if non-nil, zero value otherwise.

### GetCalendarViewOk

`func (o *User) GetCalendarViewOk() ([]MicrosoftGraphEvent, bool)`

GetCalendarViewOk returns a tuple with the CalendarView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendarView

`func (o *User) HasCalendarView() bool`

HasCalendarView returns a boolean if a field has been set.

### SetCalendarView

`func (o *User) SetCalendarView(v []MicrosoftGraphEvent)`

SetCalendarView gets a reference to the given []MicrosoftGraphEvent and assigns it to the CalendarView field.

### GetEvents

`func (o *User) GetEvents() []MicrosoftGraphEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *User) GetEventsOk() ([]MicrosoftGraphEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvents

`func (o *User) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEvents

`func (o *User) SetEvents(v []MicrosoftGraphEvent)`

SetEvents gets a reference to the given []MicrosoftGraphEvent and assigns it to the Events field.

### GetPeople

`func (o *User) GetPeople() []MicrosoftGraphPerson`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *User) GetPeopleOk() ([]MicrosoftGraphPerson, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeople

`func (o *User) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### SetPeople

`func (o *User) SetPeople(v []MicrosoftGraphPerson)`

SetPeople gets a reference to the given []MicrosoftGraphPerson and assigns it to the People field.

### GetContacts

`func (o *User) GetContacts() []MicrosoftGraphContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *User) GetContactsOk() ([]MicrosoftGraphContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContacts

`func (o *User) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContacts

`func (o *User) SetContacts(v []MicrosoftGraphContact)`

SetContacts gets a reference to the given []MicrosoftGraphContact and assigns it to the Contacts field.

### GetContactFolders

`func (o *User) GetContactFolders() []MicrosoftGraphContactFolder`

GetContactFolders returns the ContactFolders field if non-nil, zero value otherwise.

### GetContactFoldersOk

`func (o *User) GetContactFoldersOk() ([]MicrosoftGraphContactFolder, bool)`

GetContactFoldersOk returns a tuple with the ContactFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContactFolders

`func (o *User) HasContactFolders() bool`

HasContactFolders returns a boolean if a field has been set.

### SetContactFolders

`func (o *User) SetContactFolders(v []MicrosoftGraphContactFolder)`

SetContactFolders gets a reference to the given []MicrosoftGraphContactFolder and assigns it to the ContactFolders field.

### GetInferenceClassification

`func (o *User) GetInferenceClassification() AnyOfmicrosoftGraphInferenceClassification`

GetInferenceClassification returns the InferenceClassification field if non-nil, zero value otherwise.

### GetInferenceClassificationOk

`func (o *User) GetInferenceClassificationOk() (AnyOfmicrosoftGraphInferenceClassification, bool)`

GetInferenceClassificationOk returns a tuple with the InferenceClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInferenceClassification

`func (o *User) HasInferenceClassification() bool`

HasInferenceClassification returns a boolean if a field has been set.

### SetInferenceClassification

`func (o *User) SetInferenceClassification(v AnyOfmicrosoftGraphInferenceClassification)`

SetInferenceClassification gets a reference to the given AnyOfmicrosoftGraphInferenceClassification and assigns it to the InferenceClassification field.

### SetInferenceClassificationExplicitNull

`func (o *User) SetInferenceClassificationExplicitNull(b bool)`

SetInferenceClassificationExplicitNull (un)sets InferenceClassification to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InferenceClassification value is set to nil even if false is passed
### GetPhoto

`func (o *User) GetPhoto() AnyOfmicrosoftGraphProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *User) GetPhotoOk() (AnyOfmicrosoftGraphProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhoto

`func (o *User) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhoto

`func (o *User) SetPhoto(v AnyOfmicrosoftGraphProfilePhoto)`

SetPhoto gets a reference to the given AnyOfmicrosoftGraphProfilePhoto and assigns it to the Photo field.

### SetPhotoExplicitNull

`func (o *User) SetPhotoExplicitNull(b bool)`

SetPhotoExplicitNull (un)sets Photo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Photo value is set to nil even if false is passed
### GetPhotos

`func (o *User) GetPhotos() []MicrosoftGraphProfilePhoto`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *User) GetPhotosOk() ([]MicrosoftGraphProfilePhoto, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhotos

`func (o *User) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### SetPhotos

`func (o *User) SetPhotos(v []MicrosoftGraphProfilePhoto)`

SetPhotos gets a reference to the given []MicrosoftGraphProfilePhoto and assigns it to the Photos field.

### GetDrive

`func (o *User) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *User) GetDriveOk() (AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrive

`func (o *User) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDrive

`func (o *User) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive gets a reference to the given AnyOfmicrosoftGraphDrive and assigns it to the Drive field.

### SetDriveExplicitNull

`func (o *User) SetDriveExplicitNull(b bool)`

SetDriveExplicitNull (un)sets Drive to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Drive value is set to nil even if false is passed
### GetDrives

`func (o *User) GetDrives() []MicrosoftGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *User) GetDrivesOk() ([]MicrosoftGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrives

`func (o *User) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### SetDrives

`func (o *User) SetDrives(v []MicrosoftGraphDrive)`

SetDrives gets a reference to the given []MicrosoftGraphDrive and assigns it to the Drives field.

### GetExtensions

`func (o *User) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *User) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *User) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *User) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.

### GetManagedDevices

`func (o *User) GetManagedDevices() []MicrosoftGraphManagedDevice`

GetManagedDevices returns the ManagedDevices field if non-nil, zero value otherwise.

### GetManagedDevicesOk

`func (o *User) GetManagedDevicesOk() ([]MicrosoftGraphManagedDevice, bool)`

GetManagedDevicesOk returns a tuple with the ManagedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedDevices

`func (o *User) HasManagedDevices() bool`

HasManagedDevices returns a boolean if a field has been set.

### SetManagedDevices

`func (o *User) SetManagedDevices(v []MicrosoftGraphManagedDevice)`

SetManagedDevices gets a reference to the given []MicrosoftGraphManagedDevice and assigns it to the ManagedDevices field.

### GetManagedAppRegistrations

`func (o *User) GetManagedAppRegistrations() []MicrosoftGraphManagedAppRegistration`

GetManagedAppRegistrations returns the ManagedAppRegistrations field if non-nil, zero value otherwise.

### GetManagedAppRegistrationsOk

`func (o *User) GetManagedAppRegistrationsOk() ([]MicrosoftGraphManagedAppRegistration, bool)`

GetManagedAppRegistrationsOk returns a tuple with the ManagedAppRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedAppRegistrations

`func (o *User) HasManagedAppRegistrations() bool`

HasManagedAppRegistrations returns a boolean if a field has been set.

### SetManagedAppRegistrations

`func (o *User) SetManagedAppRegistrations(v []MicrosoftGraphManagedAppRegistration)`

SetManagedAppRegistrations gets a reference to the given []MicrosoftGraphManagedAppRegistration and assigns it to the ManagedAppRegistrations field.

### GetDeviceManagementTroubleshootingEvents

`func (o *User) GetDeviceManagementTroubleshootingEvents() []MicrosoftGraphDeviceManagementTroubleshootingEvent`

GetDeviceManagementTroubleshootingEvents returns the DeviceManagementTroubleshootingEvents field if non-nil, zero value otherwise.

### GetDeviceManagementTroubleshootingEventsOk

`func (o *User) GetDeviceManagementTroubleshootingEventsOk() ([]MicrosoftGraphDeviceManagementTroubleshootingEvent, bool)`

GetDeviceManagementTroubleshootingEventsOk returns a tuple with the DeviceManagementTroubleshootingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceManagementTroubleshootingEvents

`func (o *User) HasDeviceManagementTroubleshootingEvents() bool`

HasDeviceManagementTroubleshootingEvents returns a boolean if a field has been set.

### SetDeviceManagementTroubleshootingEvents

`func (o *User) SetDeviceManagementTroubleshootingEvents(v []MicrosoftGraphDeviceManagementTroubleshootingEvent)`

SetDeviceManagementTroubleshootingEvents gets a reference to the given []MicrosoftGraphDeviceManagementTroubleshootingEvent and assigns it to the DeviceManagementTroubleshootingEvents field.

### GetPlanner

`func (o *User) GetPlanner() AnyOfmicrosoftGraphPlannerUser`

GetPlanner returns the Planner field if non-nil, zero value otherwise.

### GetPlannerOk

`func (o *User) GetPlannerOk() (AnyOfmicrosoftGraphPlannerUser, bool)`

GetPlannerOk returns a tuple with the Planner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlanner

`func (o *User) HasPlanner() bool`

HasPlanner returns a boolean if a field has been set.

### SetPlanner

`func (o *User) SetPlanner(v AnyOfmicrosoftGraphPlannerUser)`

SetPlanner gets a reference to the given AnyOfmicrosoftGraphPlannerUser and assigns it to the Planner field.

### SetPlannerExplicitNull

`func (o *User) SetPlannerExplicitNull(b bool)`

SetPlannerExplicitNull (un)sets Planner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Planner value is set to nil even if false is passed
### GetInsights

`func (o *User) GetInsights() AnyOfmicrosoftGraphOfficeGraphInsights`

GetInsights returns the Insights field if non-nil, zero value otherwise.

### GetInsightsOk

`func (o *User) GetInsightsOk() (AnyOfmicrosoftGraphOfficeGraphInsights, bool)`

GetInsightsOk returns a tuple with the Insights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInsights

`func (o *User) HasInsights() bool`

HasInsights returns a boolean if a field has been set.

### SetInsights

`func (o *User) SetInsights(v AnyOfmicrosoftGraphOfficeGraphInsights)`

SetInsights gets a reference to the given AnyOfmicrosoftGraphOfficeGraphInsights and assigns it to the Insights field.

### SetInsightsExplicitNull

`func (o *User) SetInsightsExplicitNull(b bool)`

SetInsightsExplicitNull (un)sets Insights to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Insights value is set to nil even if false is passed
### GetSettings

`func (o *User) GetSettings() AnyOfmicrosoftGraphUserSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *User) GetSettingsOk() (AnyOfmicrosoftGraphUserSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettings

`func (o *User) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettings

`func (o *User) SetSettings(v AnyOfmicrosoftGraphUserSettings)`

SetSettings gets a reference to the given AnyOfmicrosoftGraphUserSettings and assigns it to the Settings field.

### SetSettingsExplicitNull

`func (o *User) SetSettingsExplicitNull(b bool)`

SetSettingsExplicitNull (un)sets Settings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Settings value is set to nil even if false is passed
### GetOnenote

`func (o *User) GetOnenote() AnyOfmicrosoftGraphOnenote`

GetOnenote returns the Onenote field if non-nil, zero value otherwise.

### GetOnenoteOk

`func (o *User) GetOnenoteOk() (AnyOfmicrosoftGraphOnenote, bool)`

GetOnenoteOk returns a tuple with the Onenote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnenote

`func (o *User) HasOnenote() bool`

HasOnenote returns a boolean if a field has been set.

### SetOnenote

`func (o *User) SetOnenote(v AnyOfmicrosoftGraphOnenote)`

SetOnenote gets a reference to the given AnyOfmicrosoftGraphOnenote and assigns it to the Onenote field.

### SetOnenoteExplicitNull

`func (o *User) SetOnenoteExplicitNull(b bool)`

SetOnenoteExplicitNull (un)sets Onenote to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Onenote value is set to nil even if false is passed
### GetActivities

`func (o *User) GetActivities() []MicrosoftGraphUserActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *User) GetActivitiesOk() ([]MicrosoftGraphUserActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivities

`func (o *User) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### SetActivities

`func (o *User) SetActivities(v []MicrosoftGraphUserActivity)`

SetActivities gets a reference to the given []MicrosoftGraphUserActivity and assigns it to the Activities field.

### GetJoinedTeams

`func (o *User) GetJoinedTeams() []MicrosoftGraphGroup`

GetJoinedTeams returns the JoinedTeams field if non-nil, zero value otherwise.

### GetJoinedTeamsOk

`func (o *User) GetJoinedTeamsOk() ([]MicrosoftGraphGroup, bool)`

GetJoinedTeamsOk returns a tuple with the JoinedTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJoinedTeams

`func (o *User) HasJoinedTeams() bool`

HasJoinedTeams returns a boolean if a field has been set.

### SetJoinedTeams

`func (o *User) SetJoinedTeams(v []MicrosoftGraphGroup)`

SetJoinedTeams gets a reference to the given []MicrosoftGraphGroup and assigns it to the JoinedTeams field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


