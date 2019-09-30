# MicrosoftGraphUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DeletedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUser) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphUser) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphUser) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDeletedDateTime

`func (o *MicrosoftGraphUser) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphUser) GetDeletedDateTimeOk() (time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeletedDateTime

`func (o *MicrosoftGraphUser) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphUser) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime gets a reference to the given time.Time and assigns it to the DeletedDateTime field.

### SetDeletedDateTimeExplicitNull

`func (o *MicrosoftGraphUser) SetDeletedDateTimeExplicitNull(b bool)`

SetDeletedDateTimeExplicitNull (un)sets DeletedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeletedDateTime value is set to nil even if false is passed
### GetAccountEnabled

`func (o *MicrosoftGraphUser) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *MicrosoftGraphUser) GetAccountEnabledOk() (bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountEnabled

`func (o *MicrosoftGraphUser) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### SetAccountEnabled

`func (o *MicrosoftGraphUser) SetAccountEnabled(v bool)`

SetAccountEnabled gets a reference to the given bool and assigns it to the AccountEnabled field.

### SetAccountEnabledExplicitNull

`func (o *MicrosoftGraphUser) SetAccountEnabledExplicitNull(b bool)`

SetAccountEnabledExplicitNull (un)sets AccountEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AccountEnabled value is set to nil even if false is passed
### GetAgeGroup

`func (o *MicrosoftGraphUser) GetAgeGroup() string`

GetAgeGroup returns the AgeGroup field if non-nil, zero value otherwise.

### GetAgeGroupOk

`func (o *MicrosoftGraphUser) GetAgeGroupOk() (string, bool)`

GetAgeGroupOk returns a tuple with the AgeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAgeGroup

`func (o *MicrosoftGraphUser) HasAgeGroup() bool`

HasAgeGroup returns a boolean if a field has been set.

### SetAgeGroup

`func (o *MicrosoftGraphUser) SetAgeGroup(v string)`

SetAgeGroup gets a reference to the given string and assigns it to the AgeGroup field.

### SetAgeGroupExplicitNull

`func (o *MicrosoftGraphUser) SetAgeGroupExplicitNull(b bool)`

SetAgeGroupExplicitNull (un)sets AgeGroup to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AgeGroup value is set to nil even if false is passed
### GetAssignedLicenses

`func (o *MicrosoftGraphUser) GetAssignedLicenses() []MicrosoftGraphAssignedLicense`

GetAssignedLicenses returns the AssignedLicenses field if non-nil, zero value otherwise.

### GetAssignedLicensesOk

`func (o *MicrosoftGraphUser) GetAssignedLicensesOk() ([]MicrosoftGraphAssignedLicense, bool)`

GetAssignedLicensesOk returns a tuple with the AssignedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedLicenses

`func (o *MicrosoftGraphUser) HasAssignedLicenses() bool`

HasAssignedLicenses returns a boolean if a field has been set.

### SetAssignedLicenses

`func (o *MicrosoftGraphUser) SetAssignedLicenses(v []MicrosoftGraphAssignedLicense)`

SetAssignedLicenses gets a reference to the given []MicrosoftGraphAssignedLicense and assigns it to the AssignedLicenses field.

### GetAssignedPlans

`func (o *MicrosoftGraphUser) GetAssignedPlans() []MicrosoftGraphAssignedPlan`

GetAssignedPlans returns the AssignedPlans field if non-nil, zero value otherwise.

### GetAssignedPlansOk

`func (o *MicrosoftGraphUser) GetAssignedPlansOk() ([]MicrosoftGraphAssignedPlan, bool)`

GetAssignedPlansOk returns a tuple with the AssignedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedPlans

`func (o *MicrosoftGraphUser) HasAssignedPlans() bool`

HasAssignedPlans returns a boolean if a field has been set.

### SetAssignedPlans

`func (o *MicrosoftGraphUser) SetAssignedPlans(v []MicrosoftGraphAssignedPlan)`

SetAssignedPlans gets a reference to the given []MicrosoftGraphAssignedPlan and assigns it to the AssignedPlans field.

### GetBusinessPhones

`func (o *MicrosoftGraphUser) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *MicrosoftGraphUser) GetBusinessPhonesOk() ([]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessPhones

`func (o *MicrosoftGraphUser) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### SetBusinessPhones

`func (o *MicrosoftGraphUser) SetBusinessPhones(v []string)`

SetBusinessPhones gets a reference to the given []string and assigns it to the BusinessPhones field.

### GetCity

`func (o *MicrosoftGraphUser) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MicrosoftGraphUser) GetCityOk() (string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCity

`func (o *MicrosoftGraphUser) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCity

`func (o *MicrosoftGraphUser) SetCity(v string)`

SetCity gets a reference to the given string and assigns it to the City field.

### SetCityExplicitNull

`func (o *MicrosoftGraphUser) SetCityExplicitNull(b bool)`

SetCityExplicitNull (un)sets City to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The City value is set to nil even if false is passed
### GetCompanyName

`func (o *MicrosoftGraphUser) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *MicrosoftGraphUser) GetCompanyNameOk() (string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompanyName

`func (o *MicrosoftGraphUser) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyName

`func (o *MicrosoftGraphUser) SetCompanyName(v string)`

SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.

### SetCompanyNameExplicitNull

`func (o *MicrosoftGraphUser) SetCompanyNameExplicitNull(b bool)`

SetCompanyNameExplicitNull (un)sets CompanyName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CompanyName value is set to nil even if false is passed
### GetConsentProvidedForMinor

`func (o *MicrosoftGraphUser) GetConsentProvidedForMinor() string`

GetConsentProvidedForMinor returns the ConsentProvidedForMinor field if non-nil, zero value otherwise.

### GetConsentProvidedForMinorOk

`func (o *MicrosoftGraphUser) GetConsentProvidedForMinorOk() (string, bool)`

GetConsentProvidedForMinorOk returns a tuple with the ConsentProvidedForMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConsentProvidedForMinor

`func (o *MicrosoftGraphUser) HasConsentProvidedForMinor() bool`

HasConsentProvidedForMinor returns a boolean if a field has been set.

### SetConsentProvidedForMinor

`func (o *MicrosoftGraphUser) SetConsentProvidedForMinor(v string)`

SetConsentProvidedForMinor gets a reference to the given string and assigns it to the ConsentProvidedForMinor field.

### SetConsentProvidedForMinorExplicitNull

`func (o *MicrosoftGraphUser) SetConsentProvidedForMinorExplicitNull(b bool)`

SetConsentProvidedForMinorExplicitNull (un)sets ConsentProvidedForMinor to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConsentProvidedForMinor value is set to nil even if false is passed
### GetCountry

`func (o *MicrosoftGraphUser) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MicrosoftGraphUser) GetCountryOk() (string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCountry

`func (o *MicrosoftGraphUser) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountry

`func (o *MicrosoftGraphUser) SetCountry(v string)`

SetCountry gets a reference to the given string and assigns it to the Country field.

### SetCountryExplicitNull

`func (o *MicrosoftGraphUser) SetCountryExplicitNull(b bool)`

SetCountryExplicitNull (un)sets Country to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Country value is set to nil even if false is passed
### GetDepartment

`func (o *MicrosoftGraphUser) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *MicrosoftGraphUser) GetDepartmentOk() (string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDepartment

`func (o *MicrosoftGraphUser) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartment

`func (o *MicrosoftGraphUser) SetDepartment(v string)`

SetDepartment gets a reference to the given string and assigns it to the Department field.

### SetDepartmentExplicitNull

`func (o *MicrosoftGraphUser) SetDepartmentExplicitNull(b bool)`

SetDepartmentExplicitNull (un)sets Department to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Department value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphUser) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphUser) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphUser) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetEmployeeId

`func (o *MicrosoftGraphUser) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *MicrosoftGraphUser) GetEmployeeIdOk() (string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmployeeId

`func (o *MicrosoftGraphUser) HasEmployeeId() bool`

HasEmployeeId returns a boolean if a field has been set.

### SetEmployeeId

`func (o *MicrosoftGraphUser) SetEmployeeId(v string)`

SetEmployeeId gets a reference to the given string and assigns it to the EmployeeId field.

### SetEmployeeIdExplicitNull

`func (o *MicrosoftGraphUser) SetEmployeeIdExplicitNull(b bool)`

SetEmployeeIdExplicitNull (un)sets EmployeeId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EmployeeId value is set to nil even if false is passed
### GetFaxNumber

`func (o *MicrosoftGraphUser) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *MicrosoftGraphUser) GetFaxNumberOk() (string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFaxNumber

`func (o *MicrosoftGraphUser) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

### SetFaxNumber

`func (o *MicrosoftGraphUser) SetFaxNumber(v string)`

SetFaxNumber gets a reference to the given string and assigns it to the FaxNumber field.

### SetFaxNumberExplicitNull

`func (o *MicrosoftGraphUser) SetFaxNumberExplicitNull(b bool)`

SetFaxNumberExplicitNull (un)sets FaxNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FaxNumber value is set to nil even if false is passed
### GetGivenName

`func (o *MicrosoftGraphUser) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MicrosoftGraphUser) GetGivenNameOk() (string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGivenName

`func (o *MicrosoftGraphUser) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenName

`func (o *MicrosoftGraphUser) SetGivenName(v string)`

SetGivenName gets a reference to the given string and assigns it to the GivenName field.

### SetGivenNameExplicitNull

`func (o *MicrosoftGraphUser) SetGivenNameExplicitNull(b bool)`

SetGivenNameExplicitNull (un)sets GivenName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GivenName value is set to nil even if false is passed
### GetImAddresses

`func (o *MicrosoftGraphUser) GetImAddresses() []string`

GetImAddresses returns the ImAddresses field if non-nil, zero value otherwise.

### GetImAddressesOk

`func (o *MicrosoftGraphUser) GetImAddressesOk() ([]string, bool)`

GetImAddressesOk returns a tuple with the ImAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImAddresses

`func (o *MicrosoftGraphUser) HasImAddresses() bool`

HasImAddresses returns a boolean if a field has been set.

### SetImAddresses

`func (o *MicrosoftGraphUser) SetImAddresses(v []string)`

SetImAddresses gets a reference to the given []string and assigns it to the ImAddresses field.

### GetIsResourceAccount

`func (o *MicrosoftGraphUser) GetIsResourceAccount() bool`

GetIsResourceAccount returns the IsResourceAccount field if non-nil, zero value otherwise.

### GetIsResourceAccountOk

`func (o *MicrosoftGraphUser) GetIsResourceAccountOk() (bool, bool)`

GetIsResourceAccountOk returns a tuple with the IsResourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsResourceAccount

`func (o *MicrosoftGraphUser) HasIsResourceAccount() bool`

HasIsResourceAccount returns a boolean if a field has been set.

### SetIsResourceAccount

`func (o *MicrosoftGraphUser) SetIsResourceAccount(v bool)`

SetIsResourceAccount gets a reference to the given bool and assigns it to the IsResourceAccount field.

### SetIsResourceAccountExplicitNull

`func (o *MicrosoftGraphUser) SetIsResourceAccountExplicitNull(b bool)`

SetIsResourceAccountExplicitNull (un)sets IsResourceAccount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsResourceAccount value is set to nil even if false is passed
### GetJobTitle

`func (o *MicrosoftGraphUser) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *MicrosoftGraphUser) GetJobTitleOk() (string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJobTitle

`func (o *MicrosoftGraphUser) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitle

`func (o *MicrosoftGraphUser) SetJobTitle(v string)`

SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.

### SetJobTitleExplicitNull

`func (o *MicrosoftGraphUser) SetJobTitleExplicitNull(b bool)`

SetJobTitleExplicitNull (un)sets JobTitle to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The JobTitle value is set to nil even if false is passed
### GetLegalAgeGroupClassification

`func (o *MicrosoftGraphUser) GetLegalAgeGroupClassification() string`

GetLegalAgeGroupClassification returns the LegalAgeGroupClassification field if non-nil, zero value otherwise.

### GetLegalAgeGroupClassificationOk

`func (o *MicrosoftGraphUser) GetLegalAgeGroupClassificationOk() (string, bool)`

GetLegalAgeGroupClassificationOk returns a tuple with the LegalAgeGroupClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLegalAgeGroupClassification

`func (o *MicrosoftGraphUser) HasLegalAgeGroupClassification() bool`

HasLegalAgeGroupClassification returns a boolean if a field has been set.

### SetLegalAgeGroupClassification

`func (o *MicrosoftGraphUser) SetLegalAgeGroupClassification(v string)`

SetLegalAgeGroupClassification gets a reference to the given string and assigns it to the LegalAgeGroupClassification field.

### SetLegalAgeGroupClassificationExplicitNull

`func (o *MicrosoftGraphUser) SetLegalAgeGroupClassificationExplicitNull(b bool)`

SetLegalAgeGroupClassificationExplicitNull (un)sets LegalAgeGroupClassification to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LegalAgeGroupClassification value is set to nil even if false is passed
### GetLicenseAssignmentStates

`func (o *MicrosoftGraphUser) GetLicenseAssignmentStates() []AnyOfmicrosoftGraphLicenseAssignmentState`

GetLicenseAssignmentStates returns the LicenseAssignmentStates field if non-nil, zero value otherwise.

### GetLicenseAssignmentStatesOk

`func (o *MicrosoftGraphUser) GetLicenseAssignmentStatesOk() ([]AnyOfmicrosoftGraphLicenseAssignmentState, bool)`

GetLicenseAssignmentStatesOk returns a tuple with the LicenseAssignmentStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicenseAssignmentStates

`func (o *MicrosoftGraphUser) HasLicenseAssignmentStates() bool`

HasLicenseAssignmentStates returns a boolean if a field has been set.

### SetLicenseAssignmentStates

`func (o *MicrosoftGraphUser) SetLicenseAssignmentStates(v []AnyOfmicrosoftGraphLicenseAssignmentState)`

SetLicenseAssignmentStates gets a reference to the given []AnyOfmicrosoftGraphLicenseAssignmentState and assigns it to the LicenseAssignmentStates field.

### GetMail

`func (o *MicrosoftGraphUser) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *MicrosoftGraphUser) GetMailOk() (string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMail

`func (o *MicrosoftGraphUser) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMail

`func (o *MicrosoftGraphUser) SetMail(v string)`

SetMail gets a reference to the given string and assigns it to the Mail field.

### SetMailExplicitNull

`func (o *MicrosoftGraphUser) SetMailExplicitNull(b bool)`

SetMailExplicitNull (un)sets Mail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Mail value is set to nil even if false is passed
### GetMailNickname

`func (o *MicrosoftGraphUser) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *MicrosoftGraphUser) GetMailNicknameOk() (string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailNickname

`func (o *MicrosoftGraphUser) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNickname

`func (o *MicrosoftGraphUser) SetMailNickname(v string)`

SetMailNickname gets a reference to the given string and assigns it to the MailNickname field.

### SetMailNicknameExplicitNull

`func (o *MicrosoftGraphUser) SetMailNicknameExplicitNull(b bool)`

SetMailNicknameExplicitNull (un)sets MailNickname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailNickname value is set to nil even if false is passed
### GetMobilePhone

`func (o *MicrosoftGraphUser) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *MicrosoftGraphUser) GetMobilePhoneOk() (string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobilePhone

`func (o *MicrosoftGraphUser) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhone

`func (o *MicrosoftGraphUser) SetMobilePhone(v string)`

SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.

### SetMobilePhoneExplicitNull

`func (o *MicrosoftGraphUser) SetMobilePhoneExplicitNull(b bool)`

SetMobilePhoneExplicitNull (un)sets MobilePhone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MobilePhone value is set to nil even if false is passed
### GetOnPremisesDistinguishedName

`func (o *MicrosoftGraphUser) GetOnPremisesDistinguishedName() string`

GetOnPremisesDistinguishedName returns the OnPremisesDistinguishedName field if non-nil, zero value otherwise.

### GetOnPremisesDistinguishedNameOk

`func (o *MicrosoftGraphUser) GetOnPremisesDistinguishedNameOk() (string, bool)`

GetOnPremisesDistinguishedNameOk returns a tuple with the OnPremisesDistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesDistinguishedName

`func (o *MicrosoftGraphUser) HasOnPremisesDistinguishedName() bool`

HasOnPremisesDistinguishedName returns a boolean if a field has been set.

### SetOnPremisesDistinguishedName

`func (o *MicrosoftGraphUser) SetOnPremisesDistinguishedName(v string)`

SetOnPremisesDistinguishedName gets a reference to the given string and assigns it to the OnPremisesDistinguishedName field.

### SetOnPremisesDistinguishedNameExplicitNull

`func (o *MicrosoftGraphUser) SetOnPremisesDistinguishedNameExplicitNull(b bool)`

SetOnPremisesDistinguishedNameExplicitNull (un)sets OnPremisesDistinguishedName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesDistinguishedName value is set to nil even if false is passed
### GetOnPremisesExtensionAttributes

`func (o *MicrosoftGraphUser) GetOnPremisesExtensionAttributes() AnyOfmicrosoftGraphOnPremisesExtensionAttributes`

GetOnPremisesExtensionAttributes returns the OnPremisesExtensionAttributes field if non-nil, zero value otherwise.

### GetOnPremisesExtensionAttributesOk

`func (o *MicrosoftGraphUser) GetOnPremisesExtensionAttributesOk() (AnyOfmicrosoftGraphOnPremisesExtensionAttributes, bool)`

GetOnPremisesExtensionAttributesOk returns a tuple with the OnPremisesExtensionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesExtensionAttributes

`func (o *MicrosoftGraphUser) HasOnPremisesExtensionAttributes() bool`

HasOnPremisesExtensionAttributes returns a boolean if a field has been set.

### SetOnPremisesExtensionAttributes

`func (o *MicrosoftGraphUser) SetOnPremisesExtensionAttributes(v AnyOfmicrosoftGraphOnPremisesExtensionAttributes)`

SetOnPremisesExtensionAttributes gets a reference to the given AnyOfmicrosoftGraphOnPremisesExtensionAttributes and assigns it to the OnPremisesExtensionAttributes field.

### SetOnPremisesExtensionAttributesExplicitNull

`func (o *MicrosoftGraphUser) SetOnPremisesExtensionAttributesExplicitNull(b bool)`

SetOnPremisesExtensionAttributesExplicitNull (un)sets OnPremisesExtensionAttributes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesExtensionAttributes value is set to nil even if false is passed
### GetOnPremisesImmutableId

`func (o *MicrosoftGraphUser) GetOnPremisesImmutableId() string`

GetOnPremisesImmutableId returns the OnPremisesImmutableId field if non-nil, zero value otherwise.

### GetOnPremisesImmutableIdOk

`func (o *MicrosoftGraphUser) GetOnPremisesImmutableIdOk() (string, bool)`

GetOnPremisesImmutableIdOk returns a tuple with the OnPremisesImmutableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesImmutableId

`func (o *MicrosoftGraphUser) HasOnPremisesImmutableId() bool`

HasOnPremisesImmutableId returns a boolean if a field has been set.

### SetOnPremisesImmutableId

`func (o *MicrosoftGraphUser) SetOnPremisesImmutableId(v string)`

SetOnPremisesImmutableId gets a reference to the given string and assigns it to the OnPremisesImmutableId field.

### SetOnPremisesImmutableIdExplicitNull

`func (o *MicrosoftGraphUser) SetOnPremisesImmutableIdExplicitNull(b bool)`

SetOnPremisesImmutableIdExplicitNull (un)sets OnPremisesImmutableId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesImmutableId value is set to nil even if false is passed
### GetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphUser) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *MicrosoftGraphUser) GetOnPremisesLastSyncDateTimeOk() (time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphUser) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphUser) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime gets a reference to the given time.Time and assigns it to the OnPremisesLastSyncDateTime field.

### SetOnPremisesLastSyncDateTimeExplicitNull

`func (o *MicrosoftGraphUser) SetOnPremisesLastSyncDateTimeExplicitNull(b bool)`

SetOnPremisesLastSyncDateTimeExplicitNull (un)sets OnPremisesLastSyncDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesLastSyncDateTime value is set to nil even if false is passed
### GetOnPremisesProvisioningErrors

`func (o *MicrosoftGraphUser) GetOnPremisesProvisioningErrors() []AnyOfmicrosoftGraphOnPremisesProvisioningError`

GetOnPremisesProvisioningErrors returns the OnPremisesProvisioningErrors field if non-nil, zero value otherwise.

### GetOnPremisesProvisioningErrorsOk

`func (o *MicrosoftGraphUser) GetOnPremisesProvisioningErrorsOk() ([]AnyOfmicrosoftGraphOnPremisesProvisioningError, bool)`

GetOnPremisesProvisioningErrorsOk returns a tuple with the OnPremisesProvisioningErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesProvisioningErrors

`func (o *MicrosoftGraphUser) HasOnPremisesProvisioningErrors() bool`

HasOnPremisesProvisioningErrors returns a boolean if a field has been set.

### SetOnPremisesProvisioningErrors

`func (o *MicrosoftGraphUser) SetOnPremisesProvisioningErrors(v []AnyOfmicrosoftGraphOnPremisesProvisioningError)`

SetOnPremisesProvisioningErrors gets a reference to the given []AnyOfmicrosoftGraphOnPremisesProvisioningError and assigns it to the OnPremisesProvisioningErrors field.

### GetOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphUser) GetOnPremisesSecurityIdentifier() string`

GetOnPremisesSecurityIdentifier returns the OnPremisesSecurityIdentifier field if non-nil, zero value otherwise.

### GetOnPremisesSecurityIdentifierOk

`func (o *MicrosoftGraphUser) GetOnPremisesSecurityIdentifierOk() (string, bool)`

GetOnPremisesSecurityIdentifierOk returns a tuple with the OnPremisesSecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphUser) HasOnPremisesSecurityIdentifier() bool`

HasOnPremisesSecurityIdentifier returns a boolean if a field has been set.

### SetOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphUser) SetOnPremisesSecurityIdentifier(v string)`

SetOnPremisesSecurityIdentifier gets a reference to the given string and assigns it to the OnPremisesSecurityIdentifier field.

### SetOnPremisesSecurityIdentifierExplicitNull

`func (o *MicrosoftGraphUser) SetOnPremisesSecurityIdentifierExplicitNull(b bool)`

SetOnPremisesSecurityIdentifierExplicitNull (un)sets OnPremisesSecurityIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSecurityIdentifier value is set to nil even if false is passed
### GetOnPremisesSyncEnabled

`func (o *MicrosoftGraphUser) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *MicrosoftGraphUser) GetOnPremisesSyncEnabledOk() (bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSyncEnabled

`func (o *MicrosoftGraphUser) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabled

`func (o *MicrosoftGraphUser) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled gets a reference to the given bool and assigns it to the OnPremisesSyncEnabled field.

### SetOnPremisesSyncEnabledExplicitNull

`func (o *MicrosoftGraphUser) SetOnPremisesSyncEnabledExplicitNull(b bool)`

SetOnPremisesSyncEnabledExplicitNull (un)sets OnPremisesSyncEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSyncEnabled value is set to nil even if false is passed
### GetOnPremisesDomainName

`func (o *MicrosoftGraphUser) GetOnPremisesDomainName() string`

GetOnPremisesDomainName returns the OnPremisesDomainName field if non-nil, zero value otherwise.

### GetOnPremisesDomainNameOk

`func (o *MicrosoftGraphUser) GetOnPremisesDomainNameOk() (string, bool)`

GetOnPremisesDomainNameOk returns a tuple with the OnPremisesDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesDomainName

`func (o *MicrosoftGraphUser) HasOnPremisesDomainName() bool`

HasOnPremisesDomainName returns a boolean if a field has been set.

### SetOnPremisesDomainName

`func (o *MicrosoftGraphUser) SetOnPremisesDomainName(v string)`

SetOnPremisesDomainName gets a reference to the given string and assigns it to the OnPremisesDomainName field.

### SetOnPremisesDomainNameExplicitNull

`func (o *MicrosoftGraphUser) SetOnPremisesDomainNameExplicitNull(b bool)`

SetOnPremisesDomainNameExplicitNull (un)sets OnPremisesDomainName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesDomainName value is set to nil even if false is passed
### GetOnPremisesSamAccountName

`func (o *MicrosoftGraphUser) GetOnPremisesSamAccountName() string`

GetOnPremisesSamAccountName returns the OnPremisesSamAccountName field if non-nil, zero value otherwise.

### GetOnPremisesSamAccountNameOk

`func (o *MicrosoftGraphUser) GetOnPremisesSamAccountNameOk() (string, bool)`

GetOnPremisesSamAccountNameOk returns a tuple with the OnPremisesSamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSamAccountName

`func (o *MicrosoftGraphUser) HasOnPremisesSamAccountName() bool`

HasOnPremisesSamAccountName returns a boolean if a field has been set.

### SetOnPremisesSamAccountName

`func (o *MicrosoftGraphUser) SetOnPremisesSamAccountName(v string)`

SetOnPremisesSamAccountName gets a reference to the given string and assigns it to the OnPremisesSamAccountName field.

### SetOnPremisesSamAccountNameExplicitNull

`func (o *MicrosoftGraphUser) SetOnPremisesSamAccountNameExplicitNull(b bool)`

SetOnPremisesSamAccountNameExplicitNull (un)sets OnPremisesSamAccountName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSamAccountName value is set to nil even if false is passed
### GetOnPremisesUserPrincipalName

`func (o *MicrosoftGraphUser) GetOnPremisesUserPrincipalName() string`

GetOnPremisesUserPrincipalName returns the OnPremisesUserPrincipalName field if non-nil, zero value otherwise.

### GetOnPremisesUserPrincipalNameOk

`func (o *MicrosoftGraphUser) GetOnPremisesUserPrincipalNameOk() (string, bool)`

GetOnPremisesUserPrincipalNameOk returns a tuple with the OnPremisesUserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesUserPrincipalName

`func (o *MicrosoftGraphUser) HasOnPremisesUserPrincipalName() bool`

HasOnPremisesUserPrincipalName returns a boolean if a field has been set.

### SetOnPremisesUserPrincipalName

`func (o *MicrosoftGraphUser) SetOnPremisesUserPrincipalName(v string)`

SetOnPremisesUserPrincipalName gets a reference to the given string and assigns it to the OnPremisesUserPrincipalName field.

### SetOnPremisesUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphUser) SetOnPremisesUserPrincipalNameExplicitNull(b bool)`

SetOnPremisesUserPrincipalNameExplicitNull (un)sets OnPremisesUserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesUserPrincipalName value is set to nil even if false is passed
### GetOtherMails

`func (o *MicrosoftGraphUser) GetOtherMails() []string`

GetOtherMails returns the OtherMails field if non-nil, zero value otherwise.

### GetOtherMailsOk

`func (o *MicrosoftGraphUser) GetOtherMailsOk() ([]string, bool)`

GetOtherMailsOk returns a tuple with the OtherMails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOtherMails

`func (o *MicrosoftGraphUser) HasOtherMails() bool`

HasOtherMails returns a boolean if a field has been set.

### SetOtherMails

`func (o *MicrosoftGraphUser) SetOtherMails(v []string)`

SetOtherMails gets a reference to the given []string and assigns it to the OtherMails field.

### GetPasswordPolicies

`func (o *MicrosoftGraphUser) GetPasswordPolicies() string`

GetPasswordPolicies returns the PasswordPolicies field if non-nil, zero value otherwise.

### GetPasswordPoliciesOk

`func (o *MicrosoftGraphUser) GetPasswordPoliciesOk() (string, bool)`

GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPolicies

`func (o *MicrosoftGraphUser) HasPasswordPolicies() bool`

HasPasswordPolicies returns a boolean if a field has been set.

### SetPasswordPolicies

`func (o *MicrosoftGraphUser) SetPasswordPolicies(v string)`

SetPasswordPolicies gets a reference to the given string and assigns it to the PasswordPolicies field.

### SetPasswordPoliciesExplicitNull

`func (o *MicrosoftGraphUser) SetPasswordPoliciesExplicitNull(b bool)`

SetPasswordPoliciesExplicitNull (un)sets PasswordPolicies to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPolicies value is set to nil even if false is passed
### GetPasswordProfile

`func (o *MicrosoftGraphUser) GetPasswordProfile() AnyOfmicrosoftGraphPasswordProfile`

GetPasswordProfile returns the PasswordProfile field if non-nil, zero value otherwise.

### GetPasswordProfileOk

`func (o *MicrosoftGraphUser) GetPasswordProfileOk() (AnyOfmicrosoftGraphPasswordProfile, bool)`

GetPasswordProfileOk returns a tuple with the PasswordProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordProfile

`func (o *MicrosoftGraphUser) HasPasswordProfile() bool`

HasPasswordProfile returns a boolean if a field has been set.

### SetPasswordProfile

`func (o *MicrosoftGraphUser) SetPasswordProfile(v AnyOfmicrosoftGraphPasswordProfile)`

SetPasswordProfile gets a reference to the given AnyOfmicrosoftGraphPasswordProfile and assigns it to the PasswordProfile field.

### SetPasswordProfileExplicitNull

`func (o *MicrosoftGraphUser) SetPasswordProfileExplicitNull(b bool)`

SetPasswordProfileExplicitNull (un)sets PasswordProfile to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordProfile value is set to nil even if false is passed
### GetOfficeLocation

`func (o *MicrosoftGraphUser) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *MicrosoftGraphUser) GetOfficeLocationOk() (string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOfficeLocation

`func (o *MicrosoftGraphUser) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocation

`func (o *MicrosoftGraphUser) SetOfficeLocation(v string)`

SetOfficeLocation gets a reference to the given string and assigns it to the OfficeLocation field.

### SetOfficeLocationExplicitNull

`func (o *MicrosoftGraphUser) SetOfficeLocationExplicitNull(b bool)`

SetOfficeLocationExplicitNull (un)sets OfficeLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OfficeLocation value is set to nil even if false is passed
### GetPostalCode

`func (o *MicrosoftGraphUser) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *MicrosoftGraphUser) GetPostalCodeOk() (string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPostalCode

`func (o *MicrosoftGraphUser) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCode

`func (o *MicrosoftGraphUser) SetPostalCode(v string)`

SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.

### SetPostalCodeExplicitNull

`func (o *MicrosoftGraphUser) SetPostalCodeExplicitNull(b bool)`

SetPostalCodeExplicitNull (un)sets PostalCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PostalCode value is set to nil even if false is passed
### GetPreferredLanguage

`func (o *MicrosoftGraphUser) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *MicrosoftGraphUser) GetPreferredLanguageOk() (string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreferredLanguage

`func (o *MicrosoftGraphUser) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguage

`func (o *MicrosoftGraphUser) SetPreferredLanguage(v string)`

SetPreferredLanguage gets a reference to the given string and assigns it to the PreferredLanguage field.

### SetPreferredLanguageExplicitNull

`func (o *MicrosoftGraphUser) SetPreferredLanguageExplicitNull(b bool)`

SetPreferredLanguageExplicitNull (un)sets PreferredLanguage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreferredLanguage value is set to nil even if false is passed
### GetProvisionedPlans

`func (o *MicrosoftGraphUser) GetProvisionedPlans() []MicrosoftGraphProvisionedPlan`

GetProvisionedPlans returns the ProvisionedPlans field if non-nil, zero value otherwise.

### GetProvisionedPlansOk

`func (o *MicrosoftGraphUser) GetProvisionedPlansOk() ([]MicrosoftGraphProvisionedPlan, bool)`

GetProvisionedPlansOk returns a tuple with the ProvisionedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProvisionedPlans

`func (o *MicrosoftGraphUser) HasProvisionedPlans() bool`

HasProvisionedPlans returns a boolean if a field has been set.

### SetProvisionedPlans

`func (o *MicrosoftGraphUser) SetProvisionedPlans(v []MicrosoftGraphProvisionedPlan)`

SetProvisionedPlans gets a reference to the given []MicrosoftGraphProvisionedPlan and assigns it to the ProvisionedPlans field.

### GetProxyAddresses

`func (o *MicrosoftGraphUser) GetProxyAddresses() []string`

GetProxyAddresses returns the ProxyAddresses field if non-nil, zero value otherwise.

### GetProxyAddressesOk

`func (o *MicrosoftGraphUser) GetProxyAddressesOk() ([]string, bool)`

GetProxyAddressesOk returns a tuple with the ProxyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProxyAddresses

`func (o *MicrosoftGraphUser) HasProxyAddresses() bool`

HasProxyAddresses returns a boolean if a field has been set.

### SetProxyAddresses

`func (o *MicrosoftGraphUser) SetProxyAddresses(v []string)`

SetProxyAddresses gets a reference to the given []string and assigns it to the ProxyAddresses field.

### GetShowInAddressList

`func (o *MicrosoftGraphUser) GetShowInAddressList() bool`

GetShowInAddressList returns the ShowInAddressList field if non-nil, zero value otherwise.

### GetShowInAddressListOk

`func (o *MicrosoftGraphUser) GetShowInAddressListOk() (bool, bool)`

GetShowInAddressListOk returns a tuple with the ShowInAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowInAddressList

`func (o *MicrosoftGraphUser) HasShowInAddressList() bool`

HasShowInAddressList returns a boolean if a field has been set.

### SetShowInAddressList

`func (o *MicrosoftGraphUser) SetShowInAddressList(v bool)`

SetShowInAddressList gets a reference to the given bool and assigns it to the ShowInAddressList field.

### SetShowInAddressListExplicitNull

`func (o *MicrosoftGraphUser) SetShowInAddressListExplicitNull(b bool)`

SetShowInAddressListExplicitNull (un)sets ShowInAddressList to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ShowInAddressList value is set to nil even if false is passed
### GetSignInSessionsValidFromDateTime

`func (o *MicrosoftGraphUser) GetSignInSessionsValidFromDateTime() time.Time`

GetSignInSessionsValidFromDateTime returns the SignInSessionsValidFromDateTime field if non-nil, zero value otherwise.

### GetSignInSessionsValidFromDateTimeOk

`func (o *MicrosoftGraphUser) GetSignInSessionsValidFromDateTimeOk() (time.Time, bool)`

GetSignInSessionsValidFromDateTimeOk returns a tuple with the SignInSessionsValidFromDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSignInSessionsValidFromDateTime

`func (o *MicrosoftGraphUser) HasSignInSessionsValidFromDateTime() bool`

HasSignInSessionsValidFromDateTime returns a boolean if a field has been set.

### SetSignInSessionsValidFromDateTime

`func (o *MicrosoftGraphUser) SetSignInSessionsValidFromDateTime(v time.Time)`

SetSignInSessionsValidFromDateTime gets a reference to the given time.Time and assigns it to the SignInSessionsValidFromDateTime field.

### SetSignInSessionsValidFromDateTimeExplicitNull

`func (o *MicrosoftGraphUser) SetSignInSessionsValidFromDateTimeExplicitNull(b bool)`

SetSignInSessionsValidFromDateTimeExplicitNull (un)sets SignInSessionsValidFromDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SignInSessionsValidFromDateTime value is set to nil even if false is passed
### GetState

`func (o *MicrosoftGraphUser) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphUser) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *MicrosoftGraphUser) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *MicrosoftGraphUser) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### SetStateExplicitNull

`func (o *MicrosoftGraphUser) SetStateExplicitNull(b bool)`

SetStateExplicitNull (un)sets State to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The State value is set to nil even if false is passed
### GetStreetAddress

`func (o *MicrosoftGraphUser) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *MicrosoftGraphUser) GetStreetAddressOk() (string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStreetAddress

`func (o *MicrosoftGraphUser) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddress

`func (o *MicrosoftGraphUser) SetStreetAddress(v string)`

SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.

### SetStreetAddressExplicitNull

`func (o *MicrosoftGraphUser) SetStreetAddressExplicitNull(b bool)`

SetStreetAddressExplicitNull (un)sets StreetAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StreetAddress value is set to nil even if false is passed
### GetSurname

`func (o *MicrosoftGraphUser) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *MicrosoftGraphUser) GetSurnameOk() (string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSurname

`func (o *MicrosoftGraphUser) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurname

`func (o *MicrosoftGraphUser) SetSurname(v string)`

SetSurname gets a reference to the given string and assigns it to the Surname field.

### SetSurnameExplicitNull

`func (o *MicrosoftGraphUser) SetSurnameExplicitNull(b bool)`

SetSurnameExplicitNull (un)sets Surname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Surname value is set to nil even if false is passed
### GetUsageLocation

`func (o *MicrosoftGraphUser) GetUsageLocation() string`

GetUsageLocation returns the UsageLocation field if non-nil, zero value otherwise.

### GetUsageLocationOk

`func (o *MicrosoftGraphUser) GetUsageLocationOk() (string, bool)`

GetUsageLocationOk returns a tuple with the UsageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageLocation

`func (o *MicrosoftGraphUser) HasUsageLocation() bool`

HasUsageLocation returns a boolean if a field has been set.

### SetUsageLocation

`func (o *MicrosoftGraphUser) SetUsageLocation(v string)`

SetUsageLocation gets a reference to the given string and assigns it to the UsageLocation field.

### SetUsageLocationExplicitNull

`func (o *MicrosoftGraphUser) SetUsageLocationExplicitNull(b bool)`

SetUsageLocationExplicitNull (un)sets UsageLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UsageLocation value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *MicrosoftGraphUser) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphUser) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *MicrosoftGraphUser) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphUser) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphUser) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetUserType

`func (o *MicrosoftGraphUser) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *MicrosoftGraphUser) GetUserTypeOk() (string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserType

`func (o *MicrosoftGraphUser) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserType

`func (o *MicrosoftGraphUser) SetUserType(v string)`

SetUserType gets a reference to the given string and assigns it to the UserType field.

### SetUserTypeExplicitNull

`func (o *MicrosoftGraphUser) SetUserTypeExplicitNull(b bool)`

SetUserTypeExplicitNull (un)sets UserType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserType value is set to nil even if false is passed
### GetMailboxSettings

`func (o *MicrosoftGraphUser) GetMailboxSettings() AnyOfmicrosoftGraphMailboxSettings`

GetMailboxSettings returns the MailboxSettings field if non-nil, zero value otherwise.

### GetMailboxSettingsOk

`func (o *MicrosoftGraphUser) GetMailboxSettingsOk() (AnyOfmicrosoftGraphMailboxSettings, bool)`

GetMailboxSettingsOk returns a tuple with the MailboxSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailboxSettings

`func (o *MicrosoftGraphUser) HasMailboxSettings() bool`

HasMailboxSettings returns a boolean if a field has been set.

### SetMailboxSettings

`func (o *MicrosoftGraphUser) SetMailboxSettings(v AnyOfmicrosoftGraphMailboxSettings)`

SetMailboxSettings gets a reference to the given AnyOfmicrosoftGraphMailboxSettings and assigns it to the MailboxSettings field.

### SetMailboxSettingsExplicitNull

`func (o *MicrosoftGraphUser) SetMailboxSettingsExplicitNull(b bool)`

SetMailboxSettingsExplicitNull (un)sets MailboxSettings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailboxSettings value is set to nil even if false is passed
### GetDeviceEnrollmentLimit

`func (o *MicrosoftGraphUser) GetDeviceEnrollmentLimit() int32`

GetDeviceEnrollmentLimit returns the DeviceEnrollmentLimit field if non-nil, zero value otherwise.

### GetDeviceEnrollmentLimitOk

`func (o *MicrosoftGraphUser) GetDeviceEnrollmentLimitOk() (int32, bool)`

GetDeviceEnrollmentLimitOk returns a tuple with the DeviceEnrollmentLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceEnrollmentLimit

`func (o *MicrosoftGraphUser) HasDeviceEnrollmentLimit() bool`

HasDeviceEnrollmentLimit returns a boolean if a field has been set.

### SetDeviceEnrollmentLimit

`func (o *MicrosoftGraphUser) SetDeviceEnrollmentLimit(v int32)`

SetDeviceEnrollmentLimit gets a reference to the given int32 and assigns it to the DeviceEnrollmentLimit field.

### GetAboutMe

`func (o *MicrosoftGraphUser) GetAboutMe() string`

GetAboutMe returns the AboutMe field if non-nil, zero value otherwise.

### GetAboutMeOk

`func (o *MicrosoftGraphUser) GetAboutMeOk() (string, bool)`

GetAboutMeOk returns a tuple with the AboutMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAboutMe

`func (o *MicrosoftGraphUser) HasAboutMe() bool`

HasAboutMe returns a boolean if a field has been set.

### SetAboutMe

`func (o *MicrosoftGraphUser) SetAboutMe(v string)`

SetAboutMe gets a reference to the given string and assigns it to the AboutMe field.

### SetAboutMeExplicitNull

`func (o *MicrosoftGraphUser) SetAboutMeExplicitNull(b bool)`

SetAboutMeExplicitNull (un)sets AboutMe to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AboutMe value is set to nil even if false is passed
### GetBirthday

`func (o *MicrosoftGraphUser) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *MicrosoftGraphUser) GetBirthdayOk() (time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBirthday

`func (o *MicrosoftGraphUser) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthday

`func (o *MicrosoftGraphUser) SetBirthday(v time.Time)`

SetBirthday gets a reference to the given time.Time and assigns it to the Birthday field.

### GetHireDate

`func (o *MicrosoftGraphUser) GetHireDate() time.Time`

GetHireDate returns the HireDate field if non-nil, zero value otherwise.

### GetHireDateOk

`func (o *MicrosoftGraphUser) GetHireDateOk() (time.Time, bool)`

GetHireDateOk returns a tuple with the HireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHireDate

`func (o *MicrosoftGraphUser) HasHireDate() bool`

HasHireDate returns a boolean if a field has been set.

### SetHireDate

`func (o *MicrosoftGraphUser) SetHireDate(v time.Time)`

SetHireDate gets a reference to the given time.Time and assigns it to the HireDate field.

### GetInterests

`func (o *MicrosoftGraphUser) GetInterests() []string`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *MicrosoftGraphUser) GetInterestsOk() ([]string, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInterests

`func (o *MicrosoftGraphUser) HasInterests() bool`

HasInterests returns a boolean if a field has been set.

### SetInterests

`func (o *MicrosoftGraphUser) SetInterests(v []string)`

SetInterests gets a reference to the given []string and assigns it to the Interests field.

### GetMySite

`func (o *MicrosoftGraphUser) GetMySite() string`

GetMySite returns the MySite field if non-nil, zero value otherwise.

### GetMySiteOk

`func (o *MicrosoftGraphUser) GetMySiteOk() (string, bool)`

GetMySiteOk returns a tuple with the MySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMySite

`func (o *MicrosoftGraphUser) HasMySite() bool`

HasMySite returns a boolean if a field has been set.

### SetMySite

`func (o *MicrosoftGraphUser) SetMySite(v string)`

SetMySite gets a reference to the given string and assigns it to the MySite field.

### SetMySiteExplicitNull

`func (o *MicrosoftGraphUser) SetMySiteExplicitNull(b bool)`

SetMySiteExplicitNull (un)sets MySite to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MySite value is set to nil even if false is passed
### GetPastProjects

`func (o *MicrosoftGraphUser) GetPastProjects() []string`

GetPastProjects returns the PastProjects field if non-nil, zero value otherwise.

### GetPastProjectsOk

`func (o *MicrosoftGraphUser) GetPastProjectsOk() ([]string, bool)`

GetPastProjectsOk returns a tuple with the PastProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPastProjects

`func (o *MicrosoftGraphUser) HasPastProjects() bool`

HasPastProjects returns a boolean if a field has been set.

### SetPastProjects

`func (o *MicrosoftGraphUser) SetPastProjects(v []string)`

SetPastProjects gets a reference to the given []string and assigns it to the PastProjects field.

### GetPreferredName

`func (o *MicrosoftGraphUser) GetPreferredName() string`

GetPreferredName returns the PreferredName field if non-nil, zero value otherwise.

### GetPreferredNameOk

`func (o *MicrosoftGraphUser) GetPreferredNameOk() (string, bool)`

GetPreferredNameOk returns a tuple with the PreferredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreferredName

`func (o *MicrosoftGraphUser) HasPreferredName() bool`

HasPreferredName returns a boolean if a field has been set.

### SetPreferredName

`func (o *MicrosoftGraphUser) SetPreferredName(v string)`

SetPreferredName gets a reference to the given string and assigns it to the PreferredName field.

### SetPreferredNameExplicitNull

`func (o *MicrosoftGraphUser) SetPreferredNameExplicitNull(b bool)`

SetPreferredNameExplicitNull (un)sets PreferredName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreferredName value is set to nil even if false is passed
### GetResponsibilities

`func (o *MicrosoftGraphUser) GetResponsibilities() []string`

GetResponsibilities returns the Responsibilities field if non-nil, zero value otherwise.

### GetResponsibilitiesOk

`func (o *MicrosoftGraphUser) GetResponsibilitiesOk() ([]string, bool)`

GetResponsibilitiesOk returns a tuple with the Responsibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponsibilities

`func (o *MicrosoftGraphUser) HasResponsibilities() bool`

HasResponsibilities returns a boolean if a field has been set.

### SetResponsibilities

`func (o *MicrosoftGraphUser) SetResponsibilities(v []string)`

SetResponsibilities gets a reference to the given []string and assigns it to the Responsibilities field.

### GetSchools

`func (o *MicrosoftGraphUser) GetSchools() []string`

GetSchools returns the Schools field if non-nil, zero value otherwise.

### GetSchoolsOk

`func (o *MicrosoftGraphUser) GetSchoolsOk() ([]string, bool)`

GetSchoolsOk returns a tuple with the Schools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchools

`func (o *MicrosoftGraphUser) HasSchools() bool`

HasSchools returns a boolean if a field has been set.

### SetSchools

`func (o *MicrosoftGraphUser) SetSchools(v []string)`

SetSchools gets a reference to the given []string and assigns it to the Schools field.

### GetSkills

`func (o *MicrosoftGraphUser) GetSkills() []string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *MicrosoftGraphUser) GetSkillsOk() ([]string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSkills

`func (o *MicrosoftGraphUser) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### SetSkills

`func (o *MicrosoftGraphUser) SetSkills(v []string)`

SetSkills gets a reference to the given []string and assigns it to the Skills field.

### GetOwnedDevices

`func (o *MicrosoftGraphUser) GetOwnedDevices() []MicrosoftGraphDirectoryObject`

GetOwnedDevices returns the OwnedDevices field if non-nil, zero value otherwise.

### GetOwnedDevicesOk

`func (o *MicrosoftGraphUser) GetOwnedDevicesOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetOwnedDevicesOk returns a tuple with the OwnedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwnedDevices

`func (o *MicrosoftGraphUser) HasOwnedDevices() bool`

HasOwnedDevices returns a boolean if a field has been set.

### SetOwnedDevices

`func (o *MicrosoftGraphUser) SetOwnedDevices(v []MicrosoftGraphDirectoryObject)`

SetOwnedDevices gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the OwnedDevices field.

### GetRegisteredDevices

`func (o *MicrosoftGraphUser) GetRegisteredDevices() []MicrosoftGraphDirectoryObject`

GetRegisteredDevices returns the RegisteredDevices field if non-nil, zero value otherwise.

### GetRegisteredDevicesOk

`func (o *MicrosoftGraphUser) GetRegisteredDevicesOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetRegisteredDevicesOk returns a tuple with the RegisteredDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRegisteredDevices

`func (o *MicrosoftGraphUser) HasRegisteredDevices() bool`

HasRegisteredDevices returns a boolean if a field has been set.

### SetRegisteredDevices

`func (o *MicrosoftGraphUser) SetRegisteredDevices(v []MicrosoftGraphDirectoryObject)`

SetRegisteredDevices gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the RegisteredDevices field.

### GetManager

`func (o *MicrosoftGraphUser) GetManager() AnyOfmicrosoftGraphDirectoryObject`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *MicrosoftGraphUser) GetManagerOk() (AnyOfmicrosoftGraphDirectoryObject, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManager

`func (o *MicrosoftGraphUser) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManager

`func (o *MicrosoftGraphUser) SetManager(v AnyOfmicrosoftGraphDirectoryObject)`

SetManager gets a reference to the given AnyOfmicrosoftGraphDirectoryObject and assigns it to the Manager field.

### SetManagerExplicitNull

`func (o *MicrosoftGraphUser) SetManagerExplicitNull(b bool)`

SetManagerExplicitNull (un)sets Manager to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Manager value is set to nil even if false is passed
### GetDirectReports

`func (o *MicrosoftGraphUser) GetDirectReports() []MicrosoftGraphDirectoryObject`

GetDirectReports returns the DirectReports field if non-nil, zero value otherwise.

### GetDirectReportsOk

`func (o *MicrosoftGraphUser) GetDirectReportsOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetDirectReportsOk returns a tuple with the DirectReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirectReports

`func (o *MicrosoftGraphUser) HasDirectReports() bool`

HasDirectReports returns a boolean if a field has been set.

### SetDirectReports

`func (o *MicrosoftGraphUser) SetDirectReports(v []MicrosoftGraphDirectoryObject)`

SetDirectReports gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the DirectReports field.

### GetMemberOf

`func (o *MicrosoftGraphUser) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *MicrosoftGraphUser) GetMemberOfOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMemberOf

`func (o *MicrosoftGraphUser) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### SetMemberOf

`func (o *MicrosoftGraphUser) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the MemberOf field.

### GetCreatedObjects

`func (o *MicrosoftGraphUser) GetCreatedObjects() []MicrosoftGraphDirectoryObject`

GetCreatedObjects returns the CreatedObjects field if non-nil, zero value otherwise.

### GetCreatedObjectsOk

`func (o *MicrosoftGraphUser) GetCreatedObjectsOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetCreatedObjectsOk returns a tuple with the CreatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedObjects

`func (o *MicrosoftGraphUser) HasCreatedObjects() bool`

HasCreatedObjects returns a boolean if a field has been set.

### SetCreatedObjects

`func (o *MicrosoftGraphUser) SetCreatedObjects(v []MicrosoftGraphDirectoryObject)`

SetCreatedObjects gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the CreatedObjects field.

### GetOwnedObjects

`func (o *MicrosoftGraphUser) GetOwnedObjects() []MicrosoftGraphDirectoryObject`

GetOwnedObjects returns the OwnedObjects field if non-nil, zero value otherwise.

### GetOwnedObjectsOk

`func (o *MicrosoftGraphUser) GetOwnedObjectsOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetOwnedObjectsOk returns a tuple with the OwnedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwnedObjects

`func (o *MicrosoftGraphUser) HasOwnedObjects() bool`

HasOwnedObjects returns a boolean if a field has been set.

### SetOwnedObjects

`func (o *MicrosoftGraphUser) SetOwnedObjects(v []MicrosoftGraphDirectoryObject)`

SetOwnedObjects gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the OwnedObjects field.

### GetLicenseDetails

`func (o *MicrosoftGraphUser) GetLicenseDetails() []MicrosoftGraphLicenseDetails`

GetLicenseDetails returns the LicenseDetails field if non-nil, zero value otherwise.

### GetLicenseDetailsOk

`func (o *MicrosoftGraphUser) GetLicenseDetailsOk() ([]MicrosoftGraphLicenseDetails, bool)`

GetLicenseDetailsOk returns a tuple with the LicenseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicenseDetails

`func (o *MicrosoftGraphUser) HasLicenseDetails() bool`

HasLicenseDetails returns a boolean if a field has been set.

### SetLicenseDetails

`func (o *MicrosoftGraphUser) SetLicenseDetails(v []MicrosoftGraphLicenseDetails)`

SetLicenseDetails gets a reference to the given []MicrosoftGraphLicenseDetails and assigns it to the LicenseDetails field.

### GetTransitiveMemberOf

`func (o *MicrosoftGraphUser) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *MicrosoftGraphUser) GetTransitiveMemberOfOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTransitiveMemberOf

`func (o *MicrosoftGraphUser) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.

### SetTransitiveMemberOf

`func (o *MicrosoftGraphUser) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the TransitiveMemberOf field.

### GetOutlook

`func (o *MicrosoftGraphUser) GetOutlook() AnyOfmicrosoftGraphOutlookUser`

GetOutlook returns the Outlook field if non-nil, zero value otherwise.

### GetOutlookOk

`func (o *MicrosoftGraphUser) GetOutlookOk() (AnyOfmicrosoftGraphOutlookUser, bool)`

GetOutlookOk returns a tuple with the Outlook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutlook

`func (o *MicrosoftGraphUser) HasOutlook() bool`

HasOutlook returns a boolean if a field has been set.

### SetOutlook

`func (o *MicrosoftGraphUser) SetOutlook(v AnyOfmicrosoftGraphOutlookUser)`

SetOutlook gets a reference to the given AnyOfmicrosoftGraphOutlookUser and assigns it to the Outlook field.

### SetOutlookExplicitNull

`func (o *MicrosoftGraphUser) SetOutlookExplicitNull(b bool)`

SetOutlookExplicitNull (un)sets Outlook to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Outlook value is set to nil even if false is passed
### GetMessages

`func (o *MicrosoftGraphUser) GetMessages() []MicrosoftGraphMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MicrosoftGraphUser) GetMessagesOk() ([]MicrosoftGraphMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessages

`func (o *MicrosoftGraphUser) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessages

`func (o *MicrosoftGraphUser) SetMessages(v []MicrosoftGraphMessage)`

SetMessages gets a reference to the given []MicrosoftGraphMessage and assigns it to the Messages field.

### GetMailFolders

`func (o *MicrosoftGraphUser) GetMailFolders() []MicrosoftGraphMailFolder`

GetMailFolders returns the MailFolders field if non-nil, zero value otherwise.

### GetMailFoldersOk

`func (o *MicrosoftGraphUser) GetMailFoldersOk() ([]MicrosoftGraphMailFolder, bool)`

GetMailFoldersOk returns a tuple with the MailFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailFolders

`func (o *MicrosoftGraphUser) HasMailFolders() bool`

HasMailFolders returns a boolean if a field has been set.

### SetMailFolders

`func (o *MicrosoftGraphUser) SetMailFolders(v []MicrosoftGraphMailFolder)`

SetMailFolders gets a reference to the given []MicrosoftGraphMailFolder and assigns it to the MailFolders field.

### GetCalendar

`func (o *MicrosoftGraphUser) GetCalendar() AnyOfmicrosoftGraphCalendar`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *MicrosoftGraphUser) GetCalendarOk() (AnyOfmicrosoftGraphCalendar, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendar

`func (o *MicrosoftGraphUser) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### SetCalendar

`func (o *MicrosoftGraphUser) SetCalendar(v AnyOfmicrosoftGraphCalendar)`

SetCalendar gets a reference to the given AnyOfmicrosoftGraphCalendar and assigns it to the Calendar field.

### SetCalendarExplicitNull

`func (o *MicrosoftGraphUser) SetCalendarExplicitNull(b bool)`

SetCalendarExplicitNull (un)sets Calendar to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Calendar value is set to nil even if false is passed
### GetCalendars

`func (o *MicrosoftGraphUser) GetCalendars() []MicrosoftGraphCalendar`

GetCalendars returns the Calendars field if non-nil, zero value otherwise.

### GetCalendarsOk

`func (o *MicrosoftGraphUser) GetCalendarsOk() ([]MicrosoftGraphCalendar, bool)`

GetCalendarsOk returns a tuple with the Calendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendars

`func (o *MicrosoftGraphUser) HasCalendars() bool`

HasCalendars returns a boolean if a field has been set.

### SetCalendars

`func (o *MicrosoftGraphUser) SetCalendars(v []MicrosoftGraphCalendar)`

SetCalendars gets a reference to the given []MicrosoftGraphCalendar and assigns it to the Calendars field.

### GetCalendarGroups

`func (o *MicrosoftGraphUser) GetCalendarGroups() []MicrosoftGraphCalendarGroup`

GetCalendarGroups returns the CalendarGroups field if non-nil, zero value otherwise.

### GetCalendarGroupsOk

`func (o *MicrosoftGraphUser) GetCalendarGroupsOk() ([]MicrosoftGraphCalendarGroup, bool)`

GetCalendarGroupsOk returns a tuple with the CalendarGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendarGroups

`func (o *MicrosoftGraphUser) HasCalendarGroups() bool`

HasCalendarGroups returns a boolean if a field has been set.

### SetCalendarGroups

`func (o *MicrosoftGraphUser) SetCalendarGroups(v []MicrosoftGraphCalendarGroup)`

SetCalendarGroups gets a reference to the given []MicrosoftGraphCalendarGroup and assigns it to the CalendarGroups field.

### GetCalendarView

`func (o *MicrosoftGraphUser) GetCalendarView() []MicrosoftGraphEvent`

GetCalendarView returns the CalendarView field if non-nil, zero value otherwise.

### GetCalendarViewOk

`func (o *MicrosoftGraphUser) GetCalendarViewOk() ([]MicrosoftGraphEvent, bool)`

GetCalendarViewOk returns a tuple with the CalendarView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendarView

`func (o *MicrosoftGraphUser) HasCalendarView() bool`

HasCalendarView returns a boolean if a field has been set.

### SetCalendarView

`func (o *MicrosoftGraphUser) SetCalendarView(v []MicrosoftGraphEvent)`

SetCalendarView gets a reference to the given []MicrosoftGraphEvent and assigns it to the CalendarView field.

### GetEvents

`func (o *MicrosoftGraphUser) GetEvents() []MicrosoftGraphEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MicrosoftGraphUser) GetEventsOk() ([]MicrosoftGraphEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvents

`func (o *MicrosoftGraphUser) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEvents

`func (o *MicrosoftGraphUser) SetEvents(v []MicrosoftGraphEvent)`

SetEvents gets a reference to the given []MicrosoftGraphEvent and assigns it to the Events field.

### GetPeople

`func (o *MicrosoftGraphUser) GetPeople() []MicrosoftGraphPerson`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *MicrosoftGraphUser) GetPeopleOk() ([]MicrosoftGraphPerson, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeople

`func (o *MicrosoftGraphUser) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### SetPeople

`func (o *MicrosoftGraphUser) SetPeople(v []MicrosoftGraphPerson)`

SetPeople gets a reference to the given []MicrosoftGraphPerson and assigns it to the People field.

### GetContacts

`func (o *MicrosoftGraphUser) GetContacts() []MicrosoftGraphContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *MicrosoftGraphUser) GetContactsOk() ([]MicrosoftGraphContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContacts

`func (o *MicrosoftGraphUser) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContacts

`func (o *MicrosoftGraphUser) SetContacts(v []MicrosoftGraphContact)`

SetContacts gets a reference to the given []MicrosoftGraphContact and assigns it to the Contacts field.

### GetContactFolders

`func (o *MicrosoftGraphUser) GetContactFolders() []MicrosoftGraphContactFolder`

GetContactFolders returns the ContactFolders field if non-nil, zero value otherwise.

### GetContactFoldersOk

`func (o *MicrosoftGraphUser) GetContactFoldersOk() ([]MicrosoftGraphContactFolder, bool)`

GetContactFoldersOk returns a tuple with the ContactFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContactFolders

`func (o *MicrosoftGraphUser) HasContactFolders() bool`

HasContactFolders returns a boolean if a field has been set.

### SetContactFolders

`func (o *MicrosoftGraphUser) SetContactFolders(v []MicrosoftGraphContactFolder)`

SetContactFolders gets a reference to the given []MicrosoftGraphContactFolder and assigns it to the ContactFolders field.

### GetInferenceClassification

`func (o *MicrosoftGraphUser) GetInferenceClassification() AnyOfmicrosoftGraphInferenceClassification`

GetInferenceClassification returns the InferenceClassification field if non-nil, zero value otherwise.

### GetInferenceClassificationOk

`func (o *MicrosoftGraphUser) GetInferenceClassificationOk() (AnyOfmicrosoftGraphInferenceClassification, bool)`

GetInferenceClassificationOk returns a tuple with the InferenceClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInferenceClassification

`func (o *MicrosoftGraphUser) HasInferenceClassification() bool`

HasInferenceClassification returns a boolean if a field has been set.

### SetInferenceClassification

`func (o *MicrosoftGraphUser) SetInferenceClassification(v AnyOfmicrosoftGraphInferenceClassification)`

SetInferenceClassification gets a reference to the given AnyOfmicrosoftGraphInferenceClassification and assigns it to the InferenceClassification field.

### SetInferenceClassificationExplicitNull

`func (o *MicrosoftGraphUser) SetInferenceClassificationExplicitNull(b bool)`

SetInferenceClassificationExplicitNull (un)sets InferenceClassification to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InferenceClassification value is set to nil even if false is passed
### GetPhoto

`func (o *MicrosoftGraphUser) GetPhoto() AnyOfmicrosoftGraphProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MicrosoftGraphUser) GetPhotoOk() (AnyOfmicrosoftGraphProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhoto

`func (o *MicrosoftGraphUser) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhoto

`func (o *MicrosoftGraphUser) SetPhoto(v AnyOfmicrosoftGraphProfilePhoto)`

SetPhoto gets a reference to the given AnyOfmicrosoftGraphProfilePhoto and assigns it to the Photo field.

### SetPhotoExplicitNull

`func (o *MicrosoftGraphUser) SetPhotoExplicitNull(b bool)`

SetPhotoExplicitNull (un)sets Photo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Photo value is set to nil even if false is passed
### GetPhotos

`func (o *MicrosoftGraphUser) GetPhotos() []MicrosoftGraphProfilePhoto`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *MicrosoftGraphUser) GetPhotosOk() ([]MicrosoftGraphProfilePhoto, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhotos

`func (o *MicrosoftGraphUser) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### SetPhotos

`func (o *MicrosoftGraphUser) SetPhotos(v []MicrosoftGraphProfilePhoto)`

SetPhotos gets a reference to the given []MicrosoftGraphProfilePhoto and assigns it to the Photos field.

### GetDrive

`func (o *MicrosoftGraphUser) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *MicrosoftGraphUser) GetDriveOk() (AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrive

`func (o *MicrosoftGraphUser) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDrive

`func (o *MicrosoftGraphUser) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive gets a reference to the given AnyOfmicrosoftGraphDrive and assigns it to the Drive field.

### SetDriveExplicitNull

`func (o *MicrosoftGraphUser) SetDriveExplicitNull(b bool)`

SetDriveExplicitNull (un)sets Drive to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Drive value is set to nil even if false is passed
### GetDrives

`func (o *MicrosoftGraphUser) GetDrives() []MicrosoftGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *MicrosoftGraphUser) GetDrivesOk() ([]MicrosoftGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrives

`func (o *MicrosoftGraphUser) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### SetDrives

`func (o *MicrosoftGraphUser) SetDrives(v []MicrosoftGraphDrive)`

SetDrives gets a reference to the given []MicrosoftGraphDrive and assigns it to the Drives field.

### GetExtensions

`func (o *MicrosoftGraphUser) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphUser) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *MicrosoftGraphUser) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *MicrosoftGraphUser) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.

### GetManagedDevices

`func (o *MicrosoftGraphUser) GetManagedDevices() []MicrosoftGraphManagedDevice`

GetManagedDevices returns the ManagedDevices field if non-nil, zero value otherwise.

### GetManagedDevicesOk

`func (o *MicrosoftGraphUser) GetManagedDevicesOk() ([]MicrosoftGraphManagedDevice, bool)`

GetManagedDevicesOk returns a tuple with the ManagedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedDevices

`func (o *MicrosoftGraphUser) HasManagedDevices() bool`

HasManagedDevices returns a boolean if a field has been set.

### SetManagedDevices

`func (o *MicrosoftGraphUser) SetManagedDevices(v []MicrosoftGraphManagedDevice)`

SetManagedDevices gets a reference to the given []MicrosoftGraphManagedDevice and assigns it to the ManagedDevices field.

### GetManagedAppRegistrations

`func (o *MicrosoftGraphUser) GetManagedAppRegistrations() []MicrosoftGraphManagedAppRegistration`

GetManagedAppRegistrations returns the ManagedAppRegistrations field if non-nil, zero value otherwise.

### GetManagedAppRegistrationsOk

`func (o *MicrosoftGraphUser) GetManagedAppRegistrationsOk() ([]MicrosoftGraphManagedAppRegistration, bool)`

GetManagedAppRegistrationsOk returns a tuple with the ManagedAppRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedAppRegistrations

`func (o *MicrosoftGraphUser) HasManagedAppRegistrations() bool`

HasManagedAppRegistrations returns a boolean if a field has been set.

### SetManagedAppRegistrations

`func (o *MicrosoftGraphUser) SetManagedAppRegistrations(v []MicrosoftGraphManagedAppRegistration)`

SetManagedAppRegistrations gets a reference to the given []MicrosoftGraphManagedAppRegistration and assigns it to the ManagedAppRegistrations field.

### GetDeviceManagementTroubleshootingEvents

`func (o *MicrosoftGraphUser) GetDeviceManagementTroubleshootingEvents() []MicrosoftGraphDeviceManagementTroubleshootingEvent`

GetDeviceManagementTroubleshootingEvents returns the DeviceManagementTroubleshootingEvents field if non-nil, zero value otherwise.

### GetDeviceManagementTroubleshootingEventsOk

`func (o *MicrosoftGraphUser) GetDeviceManagementTroubleshootingEventsOk() ([]MicrosoftGraphDeviceManagementTroubleshootingEvent, bool)`

GetDeviceManagementTroubleshootingEventsOk returns a tuple with the DeviceManagementTroubleshootingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceManagementTroubleshootingEvents

`func (o *MicrosoftGraphUser) HasDeviceManagementTroubleshootingEvents() bool`

HasDeviceManagementTroubleshootingEvents returns a boolean if a field has been set.

### SetDeviceManagementTroubleshootingEvents

`func (o *MicrosoftGraphUser) SetDeviceManagementTroubleshootingEvents(v []MicrosoftGraphDeviceManagementTroubleshootingEvent)`

SetDeviceManagementTroubleshootingEvents gets a reference to the given []MicrosoftGraphDeviceManagementTroubleshootingEvent and assigns it to the DeviceManagementTroubleshootingEvents field.

### GetPlanner

`func (o *MicrosoftGraphUser) GetPlanner() AnyOfmicrosoftGraphPlannerUser`

GetPlanner returns the Planner field if non-nil, zero value otherwise.

### GetPlannerOk

`func (o *MicrosoftGraphUser) GetPlannerOk() (AnyOfmicrosoftGraphPlannerUser, bool)`

GetPlannerOk returns a tuple with the Planner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlanner

`func (o *MicrosoftGraphUser) HasPlanner() bool`

HasPlanner returns a boolean if a field has been set.

### SetPlanner

`func (o *MicrosoftGraphUser) SetPlanner(v AnyOfmicrosoftGraphPlannerUser)`

SetPlanner gets a reference to the given AnyOfmicrosoftGraphPlannerUser and assigns it to the Planner field.

### SetPlannerExplicitNull

`func (o *MicrosoftGraphUser) SetPlannerExplicitNull(b bool)`

SetPlannerExplicitNull (un)sets Planner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Planner value is set to nil even if false is passed
### GetInsights

`func (o *MicrosoftGraphUser) GetInsights() AnyOfmicrosoftGraphOfficeGraphInsights`

GetInsights returns the Insights field if non-nil, zero value otherwise.

### GetInsightsOk

`func (o *MicrosoftGraphUser) GetInsightsOk() (AnyOfmicrosoftGraphOfficeGraphInsights, bool)`

GetInsightsOk returns a tuple with the Insights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInsights

`func (o *MicrosoftGraphUser) HasInsights() bool`

HasInsights returns a boolean if a field has been set.

### SetInsights

`func (o *MicrosoftGraphUser) SetInsights(v AnyOfmicrosoftGraphOfficeGraphInsights)`

SetInsights gets a reference to the given AnyOfmicrosoftGraphOfficeGraphInsights and assigns it to the Insights field.

### SetInsightsExplicitNull

`func (o *MicrosoftGraphUser) SetInsightsExplicitNull(b bool)`

SetInsightsExplicitNull (un)sets Insights to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Insights value is set to nil even if false is passed
### GetSettings

`func (o *MicrosoftGraphUser) GetSettings() AnyOfmicrosoftGraphUserSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MicrosoftGraphUser) GetSettingsOk() (AnyOfmicrosoftGraphUserSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettings

`func (o *MicrosoftGraphUser) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettings

`func (o *MicrosoftGraphUser) SetSettings(v AnyOfmicrosoftGraphUserSettings)`

SetSettings gets a reference to the given AnyOfmicrosoftGraphUserSettings and assigns it to the Settings field.

### SetSettingsExplicitNull

`func (o *MicrosoftGraphUser) SetSettingsExplicitNull(b bool)`

SetSettingsExplicitNull (un)sets Settings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Settings value is set to nil even if false is passed
### GetOnenote

`func (o *MicrosoftGraphUser) GetOnenote() AnyOfmicrosoftGraphOnenote`

GetOnenote returns the Onenote field if non-nil, zero value otherwise.

### GetOnenoteOk

`func (o *MicrosoftGraphUser) GetOnenoteOk() (AnyOfmicrosoftGraphOnenote, bool)`

GetOnenoteOk returns a tuple with the Onenote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnenote

`func (o *MicrosoftGraphUser) HasOnenote() bool`

HasOnenote returns a boolean if a field has been set.

### SetOnenote

`func (o *MicrosoftGraphUser) SetOnenote(v AnyOfmicrosoftGraphOnenote)`

SetOnenote gets a reference to the given AnyOfmicrosoftGraphOnenote and assigns it to the Onenote field.

### SetOnenoteExplicitNull

`func (o *MicrosoftGraphUser) SetOnenoteExplicitNull(b bool)`

SetOnenoteExplicitNull (un)sets Onenote to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Onenote value is set to nil even if false is passed
### GetActivities

`func (o *MicrosoftGraphUser) GetActivities() []MicrosoftGraphUserActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *MicrosoftGraphUser) GetActivitiesOk() ([]MicrosoftGraphUserActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivities

`func (o *MicrosoftGraphUser) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### SetActivities

`func (o *MicrosoftGraphUser) SetActivities(v []MicrosoftGraphUserActivity)`

SetActivities gets a reference to the given []MicrosoftGraphUserActivity and assigns it to the Activities field.

### GetJoinedTeams

`func (o *MicrosoftGraphUser) GetJoinedTeams() []MicrosoftGraphGroup`

GetJoinedTeams returns the JoinedTeams field if non-nil, zero value otherwise.

### GetJoinedTeamsOk

`func (o *MicrosoftGraphUser) GetJoinedTeamsOk() ([]MicrosoftGraphGroup, bool)`

GetJoinedTeamsOk returns a tuple with the JoinedTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJoinedTeams

`func (o *MicrosoftGraphUser) HasJoinedTeams() bool`

HasJoinedTeams returns a boolean if a field has been set.

### SetJoinedTeams

`func (o *MicrosoftGraphUser) SetJoinedTeams(v []MicrosoftGraphGroup)`

SetJoinedTeams gets a reference to the given []MicrosoftGraphGroup and assigns it to the JoinedTeams field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


