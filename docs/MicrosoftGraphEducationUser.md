# MicrosoftGraphEducationUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphEducationUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEducationUser) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphEducationUser) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphEducationUser) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetPrimaryRole

`func (o *MicrosoftGraphEducationUser) GetPrimaryRole() AnyOfmicrosoftGraphEducationUserRole`

GetPrimaryRole returns the PrimaryRole field if non-nil, zero value otherwise.

### GetPrimaryRoleOk

`func (o *MicrosoftGraphEducationUser) GetPrimaryRoleOk() (AnyOfmicrosoftGraphEducationUserRole, bool)`

GetPrimaryRoleOk returns a tuple with the PrimaryRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrimaryRole

`func (o *MicrosoftGraphEducationUser) HasPrimaryRole() bool`

HasPrimaryRole returns a boolean if a field has been set.

### SetPrimaryRole

`func (o *MicrosoftGraphEducationUser) SetPrimaryRole(v AnyOfmicrosoftGraphEducationUserRole)`

SetPrimaryRole gets a reference to the given AnyOfmicrosoftGraphEducationUserRole and assigns it to the PrimaryRole field.

### GetMiddleName

`func (o *MicrosoftGraphEducationUser) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *MicrosoftGraphEducationUser) GetMiddleNameOk() (string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMiddleName

`func (o *MicrosoftGraphEducationUser) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleName

`func (o *MicrosoftGraphEducationUser) SetMiddleName(v string)`

SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.

### SetMiddleNameExplicitNull

`func (o *MicrosoftGraphEducationUser) SetMiddleNameExplicitNull(b bool)`

SetMiddleNameExplicitNull (un)sets MiddleName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MiddleName value is set to nil even if false is passed
### GetExternalSource

`func (o *MicrosoftGraphEducationUser) GetExternalSource() AnyOfmicrosoftGraphEducationExternalSource`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *MicrosoftGraphEducationUser) GetExternalSourceOk() (AnyOfmicrosoftGraphEducationExternalSource, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalSource

`func (o *MicrosoftGraphEducationUser) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSource

`func (o *MicrosoftGraphEducationUser) SetExternalSource(v AnyOfmicrosoftGraphEducationExternalSource)`

SetExternalSource gets a reference to the given AnyOfmicrosoftGraphEducationExternalSource and assigns it to the ExternalSource field.

### SetExternalSourceExplicitNull

`func (o *MicrosoftGraphEducationUser) SetExternalSourceExplicitNull(b bool)`

SetExternalSourceExplicitNull (un)sets ExternalSource to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalSource value is set to nil even if false is passed
### GetResidenceAddress

`func (o *MicrosoftGraphEducationUser) GetResidenceAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetResidenceAddress returns the ResidenceAddress field if non-nil, zero value otherwise.

### GetResidenceAddressOk

`func (o *MicrosoftGraphEducationUser) GetResidenceAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetResidenceAddressOk returns a tuple with the ResidenceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResidenceAddress

`func (o *MicrosoftGraphEducationUser) HasResidenceAddress() bool`

HasResidenceAddress returns a boolean if a field has been set.

### SetResidenceAddress

`func (o *MicrosoftGraphEducationUser) SetResidenceAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetResidenceAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the ResidenceAddress field.

### SetResidenceAddressExplicitNull

`func (o *MicrosoftGraphEducationUser) SetResidenceAddressExplicitNull(b bool)`

SetResidenceAddressExplicitNull (un)sets ResidenceAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResidenceAddress value is set to nil even if false is passed
### GetMailingAddress

`func (o *MicrosoftGraphEducationUser) GetMailingAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetMailingAddress returns the MailingAddress field if non-nil, zero value otherwise.

### GetMailingAddressOk

`func (o *MicrosoftGraphEducationUser) GetMailingAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetMailingAddressOk returns a tuple with the MailingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailingAddress

`func (o *MicrosoftGraphEducationUser) HasMailingAddress() bool`

HasMailingAddress returns a boolean if a field has been set.

### SetMailingAddress

`func (o *MicrosoftGraphEducationUser) SetMailingAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetMailingAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the MailingAddress field.

### SetMailingAddressExplicitNull

`func (o *MicrosoftGraphEducationUser) SetMailingAddressExplicitNull(b bool)`

SetMailingAddressExplicitNull (un)sets MailingAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailingAddress value is set to nil even if false is passed
### GetStudent

`func (o *MicrosoftGraphEducationUser) GetStudent() AnyOfmicrosoftGraphEducationStudent`

GetStudent returns the Student field if non-nil, zero value otherwise.

### GetStudentOk

`func (o *MicrosoftGraphEducationUser) GetStudentOk() (AnyOfmicrosoftGraphEducationStudent, bool)`

GetStudentOk returns a tuple with the Student field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStudent

`func (o *MicrosoftGraphEducationUser) HasStudent() bool`

HasStudent returns a boolean if a field has been set.

### SetStudent

`func (o *MicrosoftGraphEducationUser) SetStudent(v AnyOfmicrosoftGraphEducationStudent)`

SetStudent gets a reference to the given AnyOfmicrosoftGraphEducationStudent and assigns it to the Student field.

### SetStudentExplicitNull

`func (o *MicrosoftGraphEducationUser) SetStudentExplicitNull(b bool)`

SetStudentExplicitNull (un)sets Student to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Student value is set to nil even if false is passed
### GetTeacher

`func (o *MicrosoftGraphEducationUser) GetTeacher() AnyOfmicrosoftGraphEducationTeacher`

GetTeacher returns the Teacher field if non-nil, zero value otherwise.

### GetTeacherOk

`func (o *MicrosoftGraphEducationUser) GetTeacherOk() (AnyOfmicrosoftGraphEducationTeacher, bool)`

GetTeacherOk returns a tuple with the Teacher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeacher

`func (o *MicrosoftGraphEducationUser) HasTeacher() bool`

HasTeacher returns a boolean if a field has been set.

### SetTeacher

`func (o *MicrosoftGraphEducationUser) SetTeacher(v AnyOfmicrosoftGraphEducationTeacher)`

SetTeacher gets a reference to the given AnyOfmicrosoftGraphEducationTeacher and assigns it to the Teacher field.

### SetTeacherExplicitNull

`func (o *MicrosoftGraphEducationUser) SetTeacherExplicitNull(b bool)`

SetTeacherExplicitNull (un)sets Teacher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Teacher value is set to nil even if false is passed
### GetCreatedBy

`func (o *MicrosoftGraphEducationUser) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphEducationUser) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphEducationUser) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphEducationUser) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphEducationUser) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetAccountEnabled

`func (o *MicrosoftGraphEducationUser) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *MicrosoftGraphEducationUser) GetAccountEnabledOk() (bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountEnabled

`func (o *MicrosoftGraphEducationUser) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### SetAccountEnabled

`func (o *MicrosoftGraphEducationUser) SetAccountEnabled(v bool)`

SetAccountEnabled gets a reference to the given bool and assigns it to the AccountEnabled field.

### SetAccountEnabledExplicitNull

`func (o *MicrosoftGraphEducationUser) SetAccountEnabledExplicitNull(b bool)`

SetAccountEnabledExplicitNull (un)sets AccountEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AccountEnabled value is set to nil even if false is passed
### GetAssignedLicenses

`func (o *MicrosoftGraphEducationUser) GetAssignedLicenses() []MicrosoftGraphAssignedLicense`

GetAssignedLicenses returns the AssignedLicenses field if non-nil, zero value otherwise.

### GetAssignedLicensesOk

`func (o *MicrosoftGraphEducationUser) GetAssignedLicensesOk() ([]MicrosoftGraphAssignedLicense, bool)`

GetAssignedLicensesOk returns a tuple with the AssignedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedLicenses

`func (o *MicrosoftGraphEducationUser) HasAssignedLicenses() bool`

HasAssignedLicenses returns a boolean if a field has been set.

### SetAssignedLicenses

`func (o *MicrosoftGraphEducationUser) SetAssignedLicenses(v []MicrosoftGraphAssignedLicense)`

SetAssignedLicenses gets a reference to the given []MicrosoftGraphAssignedLicense and assigns it to the AssignedLicenses field.

### GetAssignedPlans

`func (o *MicrosoftGraphEducationUser) GetAssignedPlans() []MicrosoftGraphAssignedPlan`

GetAssignedPlans returns the AssignedPlans field if non-nil, zero value otherwise.

### GetAssignedPlansOk

`func (o *MicrosoftGraphEducationUser) GetAssignedPlansOk() ([]MicrosoftGraphAssignedPlan, bool)`

GetAssignedPlansOk returns a tuple with the AssignedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedPlans

`func (o *MicrosoftGraphEducationUser) HasAssignedPlans() bool`

HasAssignedPlans returns a boolean if a field has been set.

### SetAssignedPlans

`func (o *MicrosoftGraphEducationUser) SetAssignedPlans(v []MicrosoftGraphAssignedPlan)`

SetAssignedPlans gets a reference to the given []MicrosoftGraphAssignedPlan and assigns it to the AssignedPlans field.

### GetBusinessPhones

`func (o *MicrosoftGraphEducationUser) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *MicrosoftGraphEducationUser) GetBusinessPhonesOk() ([]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessPhones

`func (o *MicrosoftGraphEducationUser) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### SetBusinessPhones

`func (o *MicrosoftGraphEducationUser) SetBusinessPhones(v []string)`

SetBusinessPhones gets a reference to the given []string and assigns it to the BusinessPhones field.

### GetDepartment

`func (o *MicrosoftGraphEducationUser) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *MicrosoftGraphEducationUser) GetDepartmentOk() (string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDepartment

`func (o *MicrosoftGraphEducationUser) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartment

`func (o *MicrosoftGraphEducationUser) SetDepartment(v string)`

SetDepartment gets a reference to the given string and assigns it to the Department field.

### SetDepartmentExplicitNull

`func (o *MicrosoftGraphEducationUser) SetDepartmentExplicitNull(b bool)`

SetDepartmentExplicitNull (un)sets Department to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Department value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphEducationUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphEducationUser) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphEducationUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphEducationUser) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphEducationUser) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetGivenName

`func (o *MicrosoftGraphEducationUser) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MicrosoftGraphEducationUser) GetGivenNameOk() (string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGivenName

`func (o *MicrosoftGraphEducationUser) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenName

`func (o *MicrosoftGraphEducationUser) SetGivenName(v string)`

SetGivenName gets a reference to the given string and assigns it to the GivenName field.

### SetGivenNameExplicitNull

`func (o *MicrosoftGraphEducationUser) SetGivenNameExplicitNull(b bool)`

SetGivenNameExplicitNull (un)sets GivenName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GivenName value is set to nil even if false is passed
### GetMail

`func (o *MicrosoftGraphEducationUser) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *MicrosoftGraphEducationUser) GetMailOk() (string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMail

`func (o *MicrosoftGraphEducationUser) HasMail() bool`

HasMail returns a boolean if a field has been set.

### SetMail

`func (o *MicrosoftGraphEducationUser) SetMail(v string)`

SetMail gets a reference to the given string and assigns it to the Mail field.

### SetMailExplicitNull

`func (o *MicrosoftGraphEducationUser) SetMailExplicitNull(b bool)`

SetMailExplicitNull (un)sets Mail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Mail value is set to nil even if false is passed
### GetMailNickname

`func (o *MicrosoftGraphEducationUser) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *MicrosoftGraphEducationUser) GetMailNicknameOk() (string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailNickname

`func (o *MicrosoftGraphEducationUser) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNickname

`func (o *MicrosoftGraphEducationUser) SetMailNickname(v string)`

SetMailNickname gets a reference to the given string and assigns it to the MailNickname field.

### SetMailNicknameExplicitNull

`func (o *MicrosoftGraphEducationUser) SetMailNicknameExplicitNull(b bool)`

SetMailNicknameExplicitNull (un)sets MailNickname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailNickname value is set to nil even if false is passed
### GetMobilePhone

`func (o *MicrosoftGraphEducationUser) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *MicrosoftGraphEducationUser) GetMobilePhoneOk() (string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobilePhone

`func (o *MicrosoftGraphEducationUser) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhone

`func (o *MicrosoftGraphEducationUser) SetMobilePhone(v string)`

SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.

### SetMobilePhoneExplicitNull

`func (o *MicrosoftGraphEducationUser) SetMobilePhoneExplicitNull(b bool)`

SetMobilePhoneExplicitNull (un)sets MobilePhone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MobilePhone value is set to nil even if false is passed
### GetPasswordPolicies

`func (o *MicrosoftGraphEducationUser) GetPasswordPolicies() string`

GetPasswordPolicies returns the PasswordPolicies field if non-nil, zero value otherwise.

### GetPasswordPoliciesOk

`func (o *MicrosoftGraphEducationUser) GetPasswordPoliciesOk() (string, bool)`

GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPolicies

`func (o *MicrosoftGraphEducationUser) HasPasswordPolicies() bool`

HasPasswordPolicies returns a boolean if a field has been set.

### SetPasswordPolicies

`func (o *MicrosoftGraphEducationUser) SetPasswordPolicies(v string)`

SetPasswordPolicies gets a reference to the given string and assigns it to the PasswordPolicies field.

### SetPasswordPoliciesExplicitNull

`func (o *MicrosoftGraphEducationUser) SetPasswordPoliciesExplicitNull(b bool)`

SetPasswordPoliciesExplicitNull (un)sets PasswordPolicies to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPolicies value is set to nil even if false is passed
### GetPasswordProfile

`func (o *MicrosoftGraphEducationUser) GetPasswordProfile() AnyOfmicrosoftGraphPasswordProfile`

GetPasswordProfile returns the PasswordProfile field if non-nil, zero value otherwise.

### GetPasswordProfileOk

`func (o *MicrosoftGraphEducationUser) GetPasswordProfileOk() (AnyOfmicrosoftGraphPasswordProfile, bool)`

GetPasswordProfileOk returns a tuple with the PasswordProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordProfile

`func (o *MicrosoftGraphEducationUser) HasPasswordProfile() bool`

HasPasswordProfile returns a boolean if a field has been set.

### SetPasswordProfile

`func (o *MicrosoftGraphEducationUser) SetPasswordProfile(v AnyOfmicrosoftGraphPasswordProfile)`

SetPasswordProfile gets a reference to the given AnyOfmicrosoftGraphPasswordProfile and assigns it to the PasswordProfile field.

### SetPasswordProfileExplicitNull

`func (o *MicrosoftGraphEducationUser) SetPasswordProfileExplicitNull(b bool)`

SetPasswordProfileExplicitNull (un)sets PasswordProfile to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordProfile value is set to nil even if false is passed
### GetOfficeLocation

`func (o *MicrosoftGraphEducationUser) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *MicrosoftGraphEducationUser) GetOfficeLocationOk() (string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOfficeLocation

`func (o *MicrosoftGraphEducationUser) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocation

`func (o *MicrosoftGraphEducationUser) SetOfficeLocation(v string)`

SetOfficeLocation gets a reference to the given string and assigns it to the OfficeLocation field.

### SetOfficeLocationExplicitNull

`func (o *MicrosoftGraphEducationUser) SetOfficeLocationExplicitNull(b bool)`

SetOfficeLocationExplicitNull (un)sets OfficeLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OfficeLocation value is set to nil even if false is passed
### GetPreferredLanguage

`func (o *MicrosoftGraphEducationUser) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *MicrosoftGraphEducationUser) GetPreferredLanguageOk() (string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreferredLanguage

`func (o *MicrosoftGraphEducationUser) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguage

`func (o *MicrosoftGraphEducationUser) SetPreferredLanguage(v string)`

SetPreferredLanguage gets a reference to the given string and assigns it to the PreferredLanguage field.

### SetPreferredLanguageExplicitNull

`func (o *MicrosoftGraphEducationUser) SetPreferredLanguageExplicitNull(b bool)`

SetPreferredLanguageExplicitNull (un)sets PreferredLanguage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreferredLanguage value is set to nil even if false is passed
### GetProvisionedPlans

`func (o *MicrosoftGraphEducationUser) GetProvisionedPlans() []MicrosoftGraphProvisionedPlan`

GetProvisionedPlans returns the ProvisionedPlans field if non-nil, zero value otherwise.

### GetProvisionedPlansOk

`func (o *MicrosoftGraphEducationUser) GetProvisionedPlansOk() ([]MicrosoftGraphProvisionedPlan, bool)`

GetProvisionedPlansOk returns a tuple with the ProvisionedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProvisionedPlans

`func (o *MicrosoftGraphEducationUser) HasProvisionedPlans() bool`

HasProvisionedPlans returns a boolean if a field has been set.

### SetProvisionedPlans

`func (o *MicrosoftGraphEducationUser) SetProvisionedPlans(v []MicrosoftGraphProvisionedPlan)`

SetProvisionedPlans gets a reference to the given []MicrosoftGraphProvisionedPlan and assigns it to the ProvisionedPlans field.

### GetRefreshTokensValidFromDateTime

`func (o *MicrosoftGraphEducationUser) GetRefreshTokensValidFromDateTime() time.Time`

GetRefreshTokensValidFromDateTime returns the RefreshTokensValidFromDateTime field if non-nil, zero value otherwise.

### GetRefreshTokensValidFromDateTimeOk

`func (o *MicrosoftGraphEducationUser) GetRefreshTokensValidFromDateTimeOk() (time.Time, bool)`

GetRefreshTokensValidFromDateTimeOk returns a tuple with the RefreshTokensValidFromDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRefreshTokensValidFromDateTime

`func (o *MicrosoftGraphEducationUser) HasRefreshTokensValidFromDateTime() bool`

HasRefreshTokensValidFromDateTime returns a boolean if a field has been set.

### SetRefreshTokensValidFromDateTime

`func (o *MicrosoftGraphEducationUser) SetRefreshTokensValidFromDateTime(v time.Time)`

SetRefreshTokensValidFromDateTime gets a reference to the given time.Time and assigns it to the RefreshTokensValidFromDateTime field.

### SetRefreshTokensValidFromDateTimeExplicitNull

`func (o *MicrosoftGraphEducationUser) SetRefreshTokensValidFromDateTimeExplicitNull(b bool)`

SetRefreshTokensValidFromDateTimeExplicitNull (un)sets RefreshTokensValidFromDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RefreshTokensValidFromDateTime value is set to nil even if false is passed
### GetShowInAddressList

`func (o *MicrosoftGraphEducationUser) GetShowInAddressList() bool`

GetShowInAddressList returns the ShowInAddressList field if non-nil, zero value otherwise.

### GetShowInAddressListOk

`func (o *MicrosoftGraphEducationUser) GetShowInAddressListOk() (bool, bool)`

GetShowInAddressListOk returns a tuple with the ShowInAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowInAddressList

`func (o *MicrosoftGraphEducationUser) HasShowInAddressList() bool`

HasShowInAddressList returns a boolean if a field has been set.

### SetShowInAddressList

`func (o *MicrosoftGraphEducationUser) SetShowInAddressList(v bool)`

SetShowInAddressList gets a reference to the given bool and assigns it to the ShowInAddressList field.

### SetShowInAddressListExplicitNull

`func (o *MicrosoftGraphEducationUser) SetShowInAddressListExplicitNull(b bool)`

SetShowInAddressListExplicitNull (un)sets ShowInAddressList to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ShowInAddressList value is set to nil even if false is passed
### GetSurname

`func (o *MicrosoftGraphEducationUser) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *MicrosoftGraphEducationUser) GetSurnameOk() (string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSurname

`func (o *MicrosoftGraphEducationUser) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurname

`func (o *MicrosoftGraphEducationUser) SetSurname(v string)`

SetSurname gets a reference to the given string and assigns it to the Surname field.

### SetSurnameExplicitNull

`func (o *MicrosoftGraphEducationUser) SetSurnameExplicitNull(b bool)`

SetSurnameExplicitNull (un)sets Surname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Surname value is set to nil even if false is passed
### GetUsageLocation

`func (o *MicrosoftGraphEducationUser) GetUsageLocation() string`

GetUsageLocation returns the UsageLocation field if non-nil, zero value otherwise.

### GetUsageLocationOk

`func (o *MicrosoftGraphEducationUser) GetUsageLocationOk() (string, bool)`

GetUsageLocationOk returns a tuple with the UsageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageLocation

`func (o *MicrosoftGraphEducationUser) HasUsageLocation() bool`

HasUsageLocation returns a boolean if a field has been set.

### SetUsageLocation

`func (o *MicrosoftGraphEducationUser) SetUsageLocation(v string)`

SetUsageLocation gets a reference to the given string and assigns it to the UsageLocation field.

### SetUsageLocationExplicitNull

`func (o *MicrosoftGraphEducationUser) SetUsageLocationExplicitNull(b bool)`

SetUsageLocationExplicitNull (un)sets UsageLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UsageLocation value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *MicrosoftGraphEducationUser) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphEducationUser) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *MicrosoftGraphEducationUser) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphEducationUser) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphEducationUser) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetUserType

`func (o *MicrosoftGraphEducationUser) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *MicrosoftGraphEducationUser) GetUserTypeOk() (string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserType

`func (o *MicrosoftGraphEducationUser) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserType

`func (o *MicrosoftGraphEducationUser) SetUserType(v string)`

SetUserType gets a reference to the given string and assigns it to the UserType field.

### SetUserTypeExplicitNull

`func (o *MicrosoftGraphEducationUser) SetUserTypeExplicitNull(b bool)`

SetUserTypeExplicitNull (un)sets UserType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserType value is set to nil even if false is passed
### GetSchools

`func (o *MicrosoftGraphEducationUser) GetSchools() []MicrosoftGraphEducationSchool`

GetSchools returns the Schools field if non-nil, zero value otherwise.

### GetSchoolsOk

`func (o *MicrosoftGraphEducationUser) GetSchoolsOk() ([]MicrosoftGraphEducationSchool, bool)`

GetSchoolsOk returns a tuple with the Schools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchools

`func (o *MicrosoftGraphEducationUser) HasSchools() bool`

HasSchools returns a boolean if a field has been set.

### SetSchools

`func (o *MicrosoftGraphEducationUser) SetSchools(v []MicrosoftGraphEducationSchool)`

SetSchools gets a reference to the given []MicrosoftGraphEducationSchool and assigns it to the Schools field.

### GetClasses

`func (o *MicrosoftGraphEducationUser) GetClasses() []MicrosoftGraphEducationClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *MicrosoftGraphEducationUser) GetClassesOk() ([]MicrosoftGraphEducationClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClasses

`func (o *MicrosoftGraphEducationUser) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### SetClasses

`func (o *MicrosoftGraphEducationUser) SetClasses(v []MicrosoftGraphEducationClass)`

SetClasses gets a reference to the given []MicrosoftGraphEducationClass and assigns it to the Classes field.

### GetUser

`func (o *MicrosoftGraphEducationUser) GetUser() AnyOfmicrosoftGraphUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MicrosoftGraphEducationUser) GetUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUser

`func (o *MicrosoftGraphEducationUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUser

`func (o *MicrosoftGraphEducationUser) SetUser(v AnyOfmicrosoftGraphUser)`

SetUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the User field.

### SetUserExplicitNull

`func (o *MicrosoftGraphEducationUser) SetUserExplicitNull(b bool)`

SetUserExplicitNull (un)sets User to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The User value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


