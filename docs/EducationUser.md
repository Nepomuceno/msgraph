# EducationUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryRole** | Pointer to [**AnyOfmicrosoftGraphEducationUserRole**](anyOf&lt;microsoft.graph.educationUserRole&gt;.md) |  | [optional] 
**MiddleName** | Pointer to **string** |  | [optional] 
**ExternalSource** | Pointer to [**AnyOfmicrosoftGraphEducationExternalSource**](anyOf&lt;microsoft.graph.educationExternalSource&gt;.md) |  | [optional] 
**ResidenceAddress** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) |  | [optional] 
**MailingAddress** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) |  | [optional] 
**Student** | Pointer to [**AnyOfmicrosoftGraphEducationStudent**](anyOf&lt;microsoft.graph.educationStudent&gt;.md) |  | [optional] 
**Teacher** | Pointer to [**AnyOfmicrosoftGraphEducationTeacher**](anyOf&lt;microsoft.graph.educationTeacher&gt;.md) |  | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**AccountEnabled** | Pointer to **bool** |  | [optional] 
**AssignedLicenses** | Pointer to [**[]MicrosoftGraphAssignedLicense**](microsoft.graph.assignedLicense.md) |  | [optional] 
**AssignedPlans** | Pointer to [**[]MicrosoftGraphAssignedPlan**](microsoft.graph.assignedPlan.md) |  | [optional] 
**BusinessPhones** | Pointer to **[]string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**Mail** | Pointer to **string** |  | [optional] 
**MailNickname** | Pointer to **string** |  | [optional] 
**MobilePhone** | Pointer to **string** |  | [optional] 
**PasswordPolicies** | Pointer to **string** |  | [optional] 
**PasswordProfile** | Pointer to [**AnyOfmicrosoftGraphPasswordProfile**](anyOf&lt;microsoft.graph.passwordProfile&gt;.md) |  | [optional] 
**OfficeLocation** | Pointer to **string** |  | [optional] 
**PreferredLanguage** | Pointer to **string** |  | [optional] 
**ProvisionedPlans** | Pointer to [**[]MicrosoftGraphProvisionedPlan**](microsoft.graph.provisionedPlan.md) |  | [optional] 
**RefreshTokensValidFromDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ShowInAddressList** | Pointer to **bool** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 
**UsageLocation** | Pointer to **string** |  | [optional] 
**UserPrincipalName** | Pointer to **string** |  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**Schools** | Pointer to [**[]MicrosoftGraphEducationSchool**](microsoft.graph.educationSchool.md) |  | [optional] 
**Classes** | Pointer to [**[]MicrosoftGraphEducationClass**](microsoft.graph.educationClass.md) |  | [optional] 
**User** | Pointer to [**AnyOfmicrosoftGraphUser**](anyOf&lt;microsoft.graph.user&gt;.md) |  | [optional] 

## Methods

### GetPrimaryRole

`func (o *EducationUser) GetPrimaryRole() AnyOfmicrosoftGraphEducationUserRole`

GetPrimaryRole returns the PrimaryRole field if non-nil, zero value otherwise.

### GetPrimaryRoleOk

`func (o *EducationUser) GetPrimaryRoleOk() (AnyOfmicrosoftGraphEducationUserRole, bool)`

GetPrimaryRoleOk returns a tuple with the PrimaryRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrimaryRole

`func (o *EducationUser) HasPrimaryRole() bool`

HasPrimaryRole returns a boolean if a field has been set.

### SetPrimaryRole

`func (o *EducationUser) SetPrimaryRole(v AnyOfmicrosoftGraphEducationUserRole)`

SetPrimaryRole gets a reference to the given AnyOfmicrosoftGraphEducationUserRole and assigns it to the PrimaryRole field.

### GetMiddleName

`func (o *EducationUser) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *EducationUser) GetMiddleNameOk() (string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMiddleName

`func (o *EducationUser) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleName

`func (o *EducationUser) SetMiddleName(v string)`

SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.

### SetMiddleNameExplicitNull

`func (o *EducationUser) SetMiddleNameExplicitNull(b bool)`

SetMiddleNameExplicitNull (un)sets MiddleName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MiddleName value is set to nil even if false is passed
### GetExternalSource

`func (o *EducationUser) GetExternalSource() AnyOfmicrosoftGraphEducationExternalSource`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *EducationUser) GetExternalSourceOk() (AnyOfmicrosoftGraphEducationExternalSource, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalSource

`func (o *EducationUser) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSource

`func (o *EducationUser) SetExternalSource(v AnyOfmicrosoftGraphEducationExternalSource)`

SetExternalSource gets a reference to the given AnyOfmicrosoftGraphEducationExternalSource and assigns it to the ExternalSource field.

### SetExternalSourceExplicitNull

`func (o *EducationUser) SetExternalSourceExplicitNull(b bool)`

SetExternalSourceExplicitNull (un)sets ExternalSource to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalSource value is set to nil even if false is passed
### GetResidenceAddress

`func (o *EducationUser) GetResidenceAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetResidenceAddress returns the ResidenceAddress field if non-nil, zero value otherwise.

### GetResidenceAddressOk

`func (o *EducationUser) GetResidenceAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetResidenceAddressOk returns a tuple with the ResidenceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResidenceAddress

`func (o *EducationUser) HasResidenceAddress() bool`

HasResidenceAddress returns a boolean if a field has been set.

### SetResidenceAddress

`func (o *EducationUser) SetResidenceAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetResidenceAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the ResidenceAddress field.

### SetResidenceAddressExplicitNull

`func (o *EducationUser) SetResidenceAddressExplicitNull(b bool)`

SetResidenceAddressExplicitNull (un)sets ResidenceAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResidenceAddress value is set to nil even if false is passed
### GetMailingAddress

`func (o *EducationUser) GetMailingAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetMailingAddress returns the MailingAddress field if non-nil, zero value otherwise.

### GetMailingAddressOk

`func (o *EducationUser) GetMailingAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetMailingAddressOk returns a tuple with the MailingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailingAddress

`func (o *EducationUser) HasMailingAddress() bool`

HasMailingAddress returns a boolean if a field has been set.

### SetMailingAddress

`func (o *EducationUser) SetMailingAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetMailingAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the MailingAddress field.

### SetMailingAddressExplicitNull

`func (o *EducationUser) SetMailingAddressExplicitNull(b bool)`

SetMailingAddressExplicitNull (un)sets MailingAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailingAddress value is set to nil even if false is passed
### GetStudent

`func (o *EducationUser) GetStudent() AnyOfmicrosoftGraphEducationStudent`

GetStudent returns the Student field if non-nil, zero value otherwise.

### GetStudentOk

`func (o *EducationUser) GetStudentOk() (AnyOfmicrosoftGraphEducationStudent, bool)`

GetStudentOk returns a tuple with the Student field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStudent

`func (o *EducationUser) HasStudent() bool`

HasStudent returns a boolean if a field has been set.

### SetStudent

`func (o *EducationUser) SetStudent(v AnyOfmicrosoftGraphEducationStudent)`

SetStudent gets a reference to the given AnyOfmicrosoftGraphEducationStudent and assigns it to the Student field.

### SetStudentExplicitNull

`func (o *EducationUser) SetStudentExplicitNull(b bool)`

SetStudentExplicitNull (un)sets Student to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Student value is set to nil even if false is passed
### GetTeacher

`func (o *EducationUser) GetTeacher() AnyOfmicrosoftGraphEducationTeacher`

GetTeacher returns the Teacher field if non-nil, zero value otherwise.

### GetTeacherOk

`func (o *EducationUser) GetTeacherOk() (AnyOfmicrosoftGraphEducationTeacher, bool)`

GetTeacherOk returns a tuple with the Teacher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeacher

`func (o *EducationUser) HasTeacher() bool`

HasTeacher returns a boolean if a field has been set.

### SetTeacher

`func (o *EducationUser) SetTeacher(v AnyOfmicrosoftGraphEducationTeacher)`

SetTeacher gets a reference to the given AnyOfmicrosoftGraphEducationTeacher and assigns it to the Teacher field.

### SetTeacherExplicitNull

`func (o *EducationUser) SetTeacherExplicitNull(b bool)`

SetTeacherExplicitNull (un)sets Teacher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Teacher value is set to nil even if false is passed
### GetCreatedBy

`func (o *EducationUser) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EducationUser) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *EducationUser) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *EducationUser) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *EducationUser) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetAccountEnabled

`func (o *EducationUser) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *EducationUser) GetAccountEnabledOk() (bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountEnabled

`func (o *EducationUser) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### SetAccountEnabled

`func (o *EducationUser) SetAccountEnabled(v bool)`

SetAccountEnabled gets a reference to the given bool and assigns it to the AccountEnabled field.

### SetAccountEnabledExplicitNull

`func (o *EducationUser) SetAccountEnabledExplicitNull(b bool)`

SetAccountEnabledExplicitNull (un)sets AccountEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AccountEnabled value is set to nil even if false is passed
### GetAssignedLicenses

`func (o *EducationUser) GetAssignedLicenses() []MicrosoftGraphAssignedLicense`

GetAssignedLicenses returns the AssignedLicenses field if non-nil, zero value otherwise.

### GetAssignedLicensesOk

`func (o *EducationUser) GetAssignedLicensesOk() ([]MicrosoftGraphAssignedLicense, bool)`

GetAssignedLicensesOk returns a tuple with the AssignedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedLicenses

`func (o *EducationUser) HasAssignedLicenses() bool`

HasAssignedLicenses returns a boolean if a field has been set.

### SetAssignedLicenses

`func (o *EducationUser) SetAssignedLicenses(v []MicrosoftGraphAssignedLicense)`

SetAssignedLicenses gets a reference to the given []MicrosoftGraphAssignedLicense and assigns it to the AssignedLicenses field.

### GetAssignedPlans

`func (o *EducationUser) GetAssignedPlans() []MicrosoftGraphAssignedPlan`

GetAssignedPlans returns the AssignedPlans field if non-nil, zero value otherwise.

### GetAssignedPlansOk

`func (o *EducationUser) GetAssignedPlansOk() ([]MicrosoftGraphAssignedPlan, bool)`

GetAssignedPlansOk returns a tuple with the AssignedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedPlans

`func (o *EducationUser) HasAssignedPlans() bool`

HasAssignedPlans returns a boolean if a field has been set.

### SetAssignedPlans

`func (o *EducationUser) SetAssignedPlans(v []MicrosoftGraphAssignedPlan)`

SetAssignedPlans gets a reference to the given []MicrosoftGraphAssignedPlan and assigns it to the AssignedPlans field.

### GetBusinessPhones

`func (o *EducationUser) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *EducationUser) GetBusinessPhonesOk() ([]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessPhones

`func (o *EducationUser) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### SetBusinessPhones

`func (o *EducationUser) SetBusinessPhones(v []string)`

SetBusinessPhones gets a reference to the given []string and assigns it to the BusinessPhones field.

### GetDepartment

`func (o *EducationUser) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *EducationUser) GetDepartmentOk() (string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDepartment

`func (o *EducationUser) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartment

`func (o *EducationUser) SetDepartment(v string)`

SetDepartment gets a reference to the given string and assigns it to the Department field.

### SetDepartmentExplicitNull

`func (o *EducationUser) SetDepartmentExplicitNull(b bool)`

SetDepartmentExplicitNull (un)sets Department to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Department value is set to nil even if false is passed
### GetDisplayName

`func (o *EducationUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EducationUser) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *EducationUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *EducationUser) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *EducationUser) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetGivenName

`func (o *EducationUser) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *EducationUser) GetGivenNameOk() (string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGivenName

`func (o *EducationUser) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenName

`func (o *EducationUser) SetGivenName(v string)`

SetGivenName gets a reference to the given string and assigns it to the GivenName field.

### SetGivenNameExplicitNull

`func (o *EducationUser) SetGivenNameExplicitNull(b bool)`

SetGivenNameExplicitNull (un)sets GivenName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GivenName value is set to nil even if false is passed
### GetMail

`func (o *EducationUser) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *EducationUser) GetMailOk() (string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMail

`func (o *EducationUser) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMail

`func (o *EducationUser) SetMail(v string)`

SetMail gets a reference to the given string and assigns it to the Mail field.

### SetMailExplicitNull

`func (o *EducationUser) SetMailExplicitNull(b bool)`

SetMailExplicitNull (un)sets Mail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Mail value is set to nil even if false is passed
### GetMailNickname

`func (o *EducationUser) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *EducationUser) GetMailNicknameOk() (string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailNickname

`func (o *EducationUser) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNickname

`func (o *EducationUser) SetMailNickname(v string)`

SetMailNickname gets a reference to the given string and assigns it to the MailNickname field.

### SetMailNicknameExplicitNull

`func (o *EducationUser) SetMailNicknameExplicitNull(b bool)`

SetMailNicknameExplicitNull (un)sets MailNickname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailNickname value is set to nil even if false is passed
### GetMobilePhone

`func (o *EducationUser) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *EducationUser) GetMobilePhoneOk() (string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobilePhone

`func (o *EducationUser) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhone

`func (o *EducationUser) SetMobilePhone(v string)`

SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.

### SetMobilePhoneExplicitNull

`func (o *EducationUser) SetMobilePhoneExplicitNull(b bool)`

SetMobilePhoneExplicitNull (un)sets MobilePhone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MobilePhone value is set to nil even if false is passed
### GetPasswordPolicies

`func (o *EducationUser) GetPasswordPolicies() string`

GetPasswordPolicies returns the PasswordPolicies field if non-nil, zero value otherwise.

### GetPasswordPoliciesOk

`func (o *EducationUser) GetPasswordPoliciesOk() (string, bool)`

GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPolicies

`func (o *EducationUser) HasPasswordPolicies() bool`

HasPasswordPolicies returns a boolean if a field has been set.

### SetPasswordPolicies

`func (o *EducationUser) SetPasswordPolicies(v string)`

SetPasswordPolicies gets a reference to the given string and assigns it to the PasswordPolicies field.

### SetPasswordPoliciesExplicitNull

`func (o *EducationUser) SetPasswordPoliciesExplicitNull(b bool)`

SetPasswordPoliciesExplicitNull (un)sets PasswordPolicies to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPolicies value is set to nil even if false is passed
### GetPasswordProfile

`func (o *EducationUser) GetPasswordProfile() AnyOfmicrosoftGraphPasswordProfile`

GetPasswordProfile returns the PasswordProfile field if non-nil, zero value otherwise.

### GetPasswordProfileOk

`func (o *EducationUser) GetPasswordProfileOk() (AnyOfmicrosoftGraphPasswordProfile, bool)`

GetPasswordProfileOk returns a tuple with the PasswordProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordProfile

`func (o *EducationUser) HasPasswordProfile() bool`

HasPasswordProfile returns a boolean if a field has been set.

### SetPasswordProfile

`func (o *EducationUser) SetPasswordProfile(v AnyOfmicrosoftGraphPasswordProfile)`

SetPasswordProfile gets a reference to the given AnyOfmicrosoftGraphPasswordProfile and assigns it to the PasswordProfile field.

### SetPasswordProfileExplicitNull

`func (o *EducationUser) SetPasswordProfileExplicitNull(b bool)`

SetPasswordProfileExplicitNull (un)sets PasswordProfile to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordProfile value is set to nil even if false is passed
### GetOfficeLocation

`func (o *EducationUser) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *EducationUser) GetOfficeLocationOk() (string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOfficeLocation

`func (o *EducationUser) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocation

`func (o *EducationUser) SetOfficeLocation(v string)`

SetOfficeLocation gets a reference to the given string and assigns it to the OfficeLocation field.

### SetOfficeLocationExplicitNull

`func (o *EducationUser) SetOfficeLocationExplicitNull(b bool)`

SetOfficeLocationExplicitNull (un)sets OfficeLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OfficeLocation value is set to nil even if false is passed
### GetPreferredLanguage

`func (o *EducationUser) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *EducationUser) GetPreferredLanguageOk() (string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreferredLanguage

`func (o *EducationUser) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguage

`func (o *EducationUser) SetPreferredLanguage(v string)`

SetPreferredLanguage gets a reference to the given string and assigns it to the PreferredLanguage field.

### SetPreferredLanguageExplicitNull

`func (o *EducationUser) SetPreferredLanguageExplicitNull(b bool)`

SetPreferredLanguageExplicitNull (un)sets PreferredLanguage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreferredLanguage value is set to nil even if false is passed
### GetProvisionedPlans

`func (o *EducationUser) GetProvisionedPlans() []MicrosoftGraphProvisionedPlan`

GetProvisionedPlans returns the ProvisionedPlans field if non-nil, zero value otherwise.

### GetProvisionedPlansOk

`func (o *EducationUser) GetProvisionedPlansOk() ([]MicrosoftGraphProvisionedPlan, bool)`

GetProvisionedPlansOk returns a tuple with the ProvisionedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProvisionedPlans

`func (o *EducationUser) HasProvisionedPlans() bool`

HasProvisionedPlans returns a boolean if a field has been set.

### SetProvisionedPlans

`func (o *EducationUser) SetProvisionedPlans(v []MicrosoftGraphProvisionedPlan)`

SetProvisionedPlans gets a reference to the given []MicrosoftGraphProvisionedPlan and assigns it to the ProvisionedPlans field.

### GetRefreshTokensValidFromDateTime

`func (o *EducationUser) GetRefreshTokensValidFromDateTime() time.Time`

GetRefreshTokensValidFromDateTime returns the RefreshTokensValidFromDateTime field if non-nil, zero value otherwise.

### GetRefreshTokensValidFromDateTimeOk

`func (o *EducationUser) GetRefreshTokensValidFromDateTimeOk() (time.Time, bool)`

GetRefreshTokensValidFromDateTimeOk returns a tuple with the RefreshTokensValidFromDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRefreshTokensValidFromDateTime

`func (o *EducationUser) HasRefreshTokensValidFromDateTime() bool`

HasRefreshTokensValidFromDateTime returns a boolean if a field has been set.

### SetRefreshTokensValidFromDateTime

`func (o *EducationUser) SetRefreshTokensValidFromDateTime(v time.Time)`

SetRefreshTokensValidFromDateTime gets a reference to the given time.Time and assigns it to the RefreshTokensValidFromDateTime field.

### SetRefreshTokensValidFromDateTimeExplicitNull

`func (o *EducationUser) SetRefreshTokensValidFromDateTimeExplicitNull(b bool)`

SetRefreshTokensValidFromDateTimeExplicitNull (un)sets RefreshTokensValidFromDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RefreshTokensValidFromDateTime value is set to nil even if false is passed
### GetShowInAddressList

`func (o *EducationUser) GetShowInAddressList() bool`

GetShowInAddressList returns the ShowInAddressList field if non-nil, zero value otherwise.

### GetShowInAddressListOk

`func (o *EducationUser) GetShowInAddressListOk() (bool, bool)`

GetShowInAddressListOk returns a tuple with the ShowInAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowInAddressList

`func (o *EducationUser) HasShowInAddressList() bool`

HasShowInAddressList returns a boolean if a field has been set.

### SetShowInAddressList

`func (o *EducationUser) SetShowInAddressList(v bool)`

SetShowInAddressList gets a reference to the given bool and assigns it to the ShowInAddressList field.

### SetShowInAddressListExplicitNull

`func (o *EducationUser) SetShowInAddressListExplicitNull(b bool)`

SetShowInAddressListExplicitNull (un)sets ShowInAddressList to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ShowInAddressList value is set to nil even if false is passed
### GetSurname

`func (o *EducationUser) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *EducationUser) GetSurnameOk() (string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSurname

`func (o *EducationUser) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurname

`func (o *EducationUser) SetSurname(v string)`

SetSurname gets a reference to the given string and assigns it to the Surname field.

### SetSurnameExplicitNull

`func (o *EducationUser) SetSurnameExplicitNull(b bool)`

SetSurnameExplicitNull (un)sets Surname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Surname value is set to nil even if false is passed
### GetUsageLocation

`func (o *EducationUser) GetUsageLocation() string`

GetUsageLocation returns the UsageLocation field if non-nil, zero value otherwise.

### GetUsageLocationOk

`func (o *EducationUser) GetUsageLocationOk() (string, bool)`

GetUsageLocationOk returns a tuple with the UsageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageLocation

`func (o *EducationUser) HasUsageLocation() bool`

HasUsageLocation returns a boolean if a field has been set.

### SetUsageLocation

`func (o *EducationUser) SetUsageLocation(v string)`

SetUsageLocation gets a reference to the given string and assigns it to the UsageLocation field.

### SetUsageLocationExplicitNull

`func (o *EducationUser) SetUsageLocationExplicitNull(b bool)`

SetUsageLocationExplicitNull (un)sets UsageLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UsageLocation value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *EducationUser) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *EducationUser) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *EducationUser) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *EducationUser) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *EducationUser) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetUserType

`func (o *EducationUser) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *EducationUser) GetUserTypeOk() (string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserType

`func (o *EducationUser) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserType

`func (o *EducationUser) SetUserType(v string)`

SetUserType gets a reference to the given string and assigns it to the UserType field.

### SetUserTypeExplicitNull

`func (o *EducationUser) SetUserTypeExplicitNull(b bool)`

SetUserTypeExplicitNull (un)sets UserType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserType value is set to nil even if false is passed
### GetSchools

`func (o *EducationUser) GetSchools() []MicrosoftGraphEducationSchool`

GetSchools returns the Schools field if non-nil, zero value otherwise.

### GetSchoolsOk

`func (o *EducationUser) GetSchoolsOk() ([]MicrosoftGraphEducationSchool, bool)`

GetSchoolsOk returns a tuple with the Schools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchools

`func (o *EducationUser) HasSchools() bool`

HasSchools returns a boolean if a field has been set.

### SetSchools

`func (o *EducationUser) SetSchools(v []MicrosoftGraphEducationSchool)`

SetSchools gets a reference to the given []MicrosoftGraphEducationSchool and assigns it to the Schools field.

### GetClasses

`func (o *EducationUser) GetClasses() []MicrosoftGraphEducationClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *EducationUser) GetClassesOk() ([]MicrosoftGraphEducationClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClasses

`func (o *EducationUser) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### SetClasses

`func (o *EducationUser) SetClasses(v []MicrosoftGraphEducationClass)`

SetClasses gets a reference to the given []MicrosoftGraphEducationClass and assigns it to the Classes field.

### GetUser

`func (o *EducationUser) GetUser() AnyOfmicrosoftGraphUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EducationUser) GetUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUser

`func (o *EducationUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUser

`func (o *EducationUser) SetUser(v AnyOfmicrosoftGraphUser)`

SetUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the User field.

### SetUserExplicitNull

`func (o *EducationUser) SetUserExplicitNull(b bool)`

SetUserExplicitNull (un)sets User to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The User value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


