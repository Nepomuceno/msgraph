# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentFolderId** | Pointer to **string** |  | [optional] 
**Birthday** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**FileAs** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**Initials** | Pointer to **string** |  | [optional] 
**MiddleName** | Pointer to **string** |  | [optional] 
**NickName** | Pointer to **string** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**YomiGivenName** | Pointer to **string** |  | [optional] 
**YomiSurname** | Pointer to **string** |  | [optional] 
**YomiCompanyName** | Pointer to **string** |  | [optional] 
**Generation** | Pointer to **string** |  | [optional] 
**EmailAddresses** | Pointer to [**[]AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) |  | [optional] 
**ImAddresses** | Pointer to **[]string** |  | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**OfficeLocation** | Pointer to **string** |  | [optional] 
**Profession** | Pointer to **string** |  | [optional] 
**BusinessHomePage** | Pointer to **string** |  | [optional] 
**AssistantName** | Pointer to **string** |  | [optional] 
**Manager** | Pointer to **string** |  | [optional] 
**HomePhones** | Pointer to **[]string** |  | [optional] 
**MobilePhone** | Pointer to **string** |  | [optional] 
**BusinessPhones** | Pointer to **[]string** |  | [optional] 
**HomeAddress** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) |  | [optional] 
**BusinessAddress** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) |  | [optional] 
**OtherAddress** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) |  | [optional] 
**SpouseName** | Pointer to **string** |  | [optional] 
**PersonalNotes** | Pointer to **string** |  | [optional] 
**Children** | Pointer to **[]string** |  | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md) |  | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md) |  | [optional] 
**Photo** | Pointer to [**AnyOfmicrosoftGraphProfilePhoto**](anyOf&lt;microsoft.graph.profilePhoto&gt;.md) |  | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](microsoft.graph.extension.md) |  | [optional] 

## Methods

### GetParentFolderId

`func (o *Contact) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *Contact) GetParentFolderIdOk() (string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentFolderId

`func (o *Contact) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderId

`func (o *Contact) SetParentFolderId(v string)`

SetParentFolderId gets a reference to the given string and assigns it to the ParentFolderId field.

### SetParentFolderIdExplicitNull

`func (o *Contact) SetParentFolderIdExplicitNull(b bool)`

SetParentFolderIdExplicitNull (un)sets ParentFolderId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentFolderId value is set to nil even if false is passed
### GetBirthday

`func (o *Contact) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *Contact) GetBirthdayOk() (time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBirthday

`func (o *Contact) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthday

`func (o *Contact) SetBirthday(v time.Time)`

SetBirthday gets a reference to the given time.Time and assigns it to the Birthday field.

### SetBirthdayExplicitNull

`func (o *Contact) SetBirthdayExplicitNull(b bool)`

SetBirthdayExplicitNull (un)sets Birthday to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Birthday value is set to nil even if false is passed
### GetFileAs

`func (o *Contact) GetFileAs() string`

GetFileAs returns the FileAs field if non-nil, zero value otherwise.

### GetFileAsOk

`func (o *Contact) GetFileAsOk() (string, bool)`

GetFileAsOk returns a tuple with the FileAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileAs

`func (o *Contact) HasFileAs() bool`

HasFileAs returns a boolean if a field has been set.

### SetFileAs

`func (o *Contact) SetFileAs(v string)`

SetFileAs gets a reference to the given string and assigns it to the FileAs field.

### SetFileAsExplicitNull

`func (o *Contact) SetFileAsExplicitNull(b bool)`

SetFileAsExplicitNull (un)sets FileAs to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileAs value is set to nil even if false is passed
### GetDisplayName

`func (o *Contact) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Contact) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *Contact) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *Contact) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *Contact) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetGivenName

`func (o *Contact) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *Contact) GetGivenNameOk() (string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGivenName

`func (o *Contact) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenName

`func (o *Contact) SetGivenName(v string)`

SetGivenName gets a reference to the given string and assigns it to the GivenName field.

### SetGivenNameExplicitNull

`func (o *Contact) SetGivenNameExplicitNull(b bool)`

SetGivenNameExplicitNull (un)sets GivenName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GivenName value is set to nil even if false is passed
### GetInitials

`func (o *Contact) GetInitials() string`

GetInitials returns the Initials field if non-nil, zero value otherwise.

### GetInitialsOk

`func (o *Contact) GetInitialsOk() (string, bool)`

GetInitialsOk returns a tuple with the Initials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInitials

`func (o *Contact) HasInitials() bool`

HasInitials returns a boolean if a field has been set.

### SetInitials

`func (o *Contact) SetInitials(v string)`

SetInitials gets a reference to the given string and assigns it to the Initials field.

### SetInitialsExplicitNull

`func (o *Contact) SetInitialsExplicitNull(b bool)`

SetInitialsExplicitNull (un)sets Initials to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Initials value is set to nil even if false is passed
### GetMiddleName

`func (o *Contact) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *Contact) GetMiddleNameOk() (string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMiddleName

`func (o *Contact) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleName

`func (o *Contact) SetMiddleName(v string)`

SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.

### SetMiddleNameExplicitNull

`func (o *Contact) SetMiddleNameExplicitNull(b bool)`

SetMiddleNameExplicitNull (un)sets MiddleName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MiddleName value is set to nil even if false is passed
### GetNickName

`func (o *Contact) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *Contact) GetNickNameOk() (string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNickName

`func (o *Contact) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickName

`func (o *Contact) SetNickName(v string)`

SetNickName gets a reference to the given string and assigns it to the NickName field.

### SetNickNameExplicitNull

`func (o *Contact) SetNickNameExplicitNull(b bool)`

SetNickNameExplicitNull (un)sets NickName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NickName value is set to nil even if false is passed
### GetSurname

`func (o *Contact) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *Contact) GetSurnameOk() (string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSurname

`func (o *Contact) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurname

`func (o *Contact) SetSurname(v string)`

SetSurname gets a reference to the given string and assigns it to the Surname field.

### SetSurnameExplicitNull

`func (o *Contact) SetSurnameExplicitNull(b bool)`

SetSurnameExplicitNull (un)sets Surname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Surname value is set to nil even if false is passed
### GetTitle

`func (o *Contact) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Contact) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *Contact) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *Contact) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### SetTitleExplicitNull

`func (o *Contact) SetTitleExplicitNull(b bool)`

SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Title value is set to nil even if false is passed
### GetYomiGivenName

`func (o *Contact) GetYomiGivenName() string`

GetYomiGivenName returns the YomiGivenName field if non-nil, zero value otherwise.

### GetYomiGivenNameOk

`func (o *Contact) GetYomiGivenNameOk() (string, bool)`

GetYomiGivenNameOk returns a tuple with the YomiGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasYomiGivenName

`func (o *Contact) HasYomiGivenName() bool`

HasYomiGivenName returns a boolean if a field has been set.

### SetYomiGivenName

`func (o *Contact) SetYomiGivenName(v string)`

SetYomiGivenName gets a reference to the given string and assigns it to the YomiGivenName field.

### SetYomiGivenNameExplicitNull

`func (o *Contact) SetYomiGivenNameExplicitNull(b bool)`

SetYomiGivenNameExplicitNull (un)sets YomiGivenName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The YomiGivenName value is set to nil even if false is passed
### GetYomiSurname

`func (o *Contact) GetYomiSurname() string`

GetYomiSurname returns the YomiSurname field if non-nil, zero value otherwise.

### GetYomiSurnameOk

`func (o *Contact) GetYomiSurnameOk() (string, bool)`

GetYomiSurnameOk returns a tuple with the YomiSurname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasYomiSurname

`func (o *Contact) HasYomiSurname() bool`

HasYomiSurname returns a boolean if a field has been set.

### SetYomiSurname

`func (o *Contact) SetYomiSurname(v string)`

SetYomiSurname gets a reference to the given string and assigns it to the YomiSurname field.

### SetYomiSurnameExplicitNull

`func (o *Contact) SetYomiSurnameExplicitNull(b bool)`

SetYomiSurnameExplicitNull (un)sets YomiSurname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The YomiSurname value is set to nil even if false is passed
### GetYomiCompanyName

`func (o *Contact) GetYomiCompanyName() string`

GetYomiCompanyName returns the YomiCompanyName field if non-nil, zero value otherwise.

### GetYomiCompanyNameOk

`func (o *Contact) GetYomiCompanyNameOk() (string, bool)`

GetYomiCompanyNameOk returns a tuple with the YomiCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasYomiCompanyName

`func (o *Contact) HasYomiCompanyName() bool`

HasYomiCompanyName returns a boolean if a field has been set.

### SetYomiCompanyName

`func (o *Contact) SetYomiCompanyName(v string)`

SetYomiCompanyName gets a reference to the given string and assigns it to the YomiCompanyName field.

### SetYomiCompanyNameExplicitNull

`func (o *Contact) SetYomiCompanyNameExplicitNull(b bool)`

SetYomiCompanyNameExplicitNull (un)sets YomiCompanyName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The YomiCompanyName value is set to nil even if false is passed
### GetGeneration

`func (o *Contact) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *Contact) GetGenerationOk() (string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGeneration

`func (o *Contact) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### SetGeneration

`func (o *Contact) SetGeneration(v string)`

SetGeneration gets a reference to the given string and assigns it to the Generation field.

### SetGenerationExplicitNull

`func (o *Contact) SetGenerationExplicitNull(b bool)`

SetGenerationExplicitNull (un)sets Generation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Generation value is set to nil even if false is passed
### GetEmailAddresses

`func (o *Contact) GetEmailAddresses() []AnyOfmicrosoftGraphEmailAddress`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *Contact) GetEmailAddressesOk() ([]AnyOfmicrosoftGraphEmailAddress, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailAddresses

`func (o *Contact) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### SetEmailAddresses

`func (o *Contact) SetEmailAddresses(v []AnyOfmicrosoftGraphEmailAddress)`

SetEmailAddresses gets a reference to the given []AnyOfmicrosoftGraphEmailAddress and assigns it to the EmailAddresses field.

### GetImAddresses

`func (o *Contact) GetImAddresses() []string`

GetImAddresses returns the ImAddresses field if non-nil, zero value otherwise.

### GetImAddressesOk

`func (o *Contact) GetImAddressesOk() ([]string, bool)`

GetImAddressesOk returns a tuple with the ImAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImAddresses

`func (o *Contact) HasImAddresses() bool`

HasImAddresses returns a boolean if a field has been set.

### SetImAddresses

`func (o *Contact) SetImAddresses(v []string)`

SetImAddresses gets a reference to the given []string and assigns it to the ImAddresses field.

### GetJobTitle

`func (o *Contact) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Contact) GetJobTitleOk() (string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJobTitle

`func (o *Contact) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitle

`func (o *Contact) SetJobTitle(v string)`

SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.

### SetJobTitleExplicitNull

`func (o *Contact) SetJobTitleExplicitNull(b bool)`

SetJobTitleExplicitNull (un)sets JobTitle to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The JobTitle value is set to nil even if false is passed
### GetCompanyName

`func (o *Contact) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Contact) GetCompanyNameOk() (string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompanyName

`func (o *Contact) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyName

`func (o *Contact) SetCompanyName(v string)`

SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.

### SetCompanyNameExplicitNull

`func (o *Contact) SetCompanyNameExplicitNull(b bool)`

SetCompanyNameExplicitNull (un)sets CompanyName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CompanyName value is set to nil even if false is passed
### GetDepartment

`func (o *Contact) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Contact) GetDepartmentOk() (string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDepartment

`func (o *Contact) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartment

`func (o *Contact) SetDepartment(v string)`

SetDepartment gets a reference to the given string and assigns it to the Department field.

### SetDepartmentExplicitNull

`func (o *Contact) SetDepartmentExplicitNull(b bool)`

SetDepartmentExplicitNull (un)sets Department to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Department value is set to nil even if false is passed
### GetOfficeLocation

`func (o *Contact) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *Contact) GetOfficeLocationOk() (string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOfficeLocation

`func (o *Contact) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocation

`func (o *Contact) SetOfficeLocation(v string)`

SetOfficeLocation gets a reference to the given string and assigns it to the OfficeLocation field.

### SetOfficeLocationExplicitNull

`func (o *Contact) SetOfficeLocationExplicitNull(b bool)`

SetOfficeLocationExplicitNull (un)sets OfficeLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OfficeLocation value is set to nil even if false is passed
### GetProfession

`func (o *Contact) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *Contact) GetProfessionOk() (string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfession

`func (o *Contact) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### SetProfession

`func (o *Contact) SetProfession(v string)`

SetProfession gets a reference to the given string and assigns it to the Profession field.

### SetProfessionExplicitNull

`func (o *Contact) SetProfessionExplicitNull(b bool)`

SetProfessionExplicitNull (un)sets Profession to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Profession value is set to nil even if false is passed
### GetBusinessHomePage

`func (o *Contact) GetBusinessHomePage() string`

GetBusinessHomePage returns the BusinessHomePage field if non-nil, zero value otherwise.

### GetBusinessHomePageOk

`func (o *Contact) GetBusinessHomePageOk() (string, bool)`

GetBusinessHomePageOk returns a tuple with the BusinessHomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessHomePage

`func (o *Contact) HasBusinessHomePage() bool`

HasBusinessHomePage returns a boolean if a field has been set.

### SetBusinessHomePage

`func (o *Contact) SetBusinessHomePage(v string)`

SetBusinessHomePage gets a reference to the given string and assigns it to the BusinessHomePage field.

### SetBusinessHomePageExplicitNull

`func (o *Contact) SetBusinessHomePageExplicitNull(b bool)`

SetBusinessHomePageExplicitNull (un)sets BusinessHomePage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BusinessHomePage value is set to nil even if false is passed
### GetAssistantName

`func (o *Contact) GetAssistantName() string`

GetAssistantName returns the AssistantName field if non-nil, zero value otherwise.

### GetAssistantNameOk

`func (o *Contact) GetAssistantNameOk() (string, bool)`

GetAssistantNameOk returns a tuple with the AssistantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssistantName

`func (o *Contact) HasAssistantName() bool`

HasAssistantName returns a boolean if a field has been set.

### SetAssistantName

`func (o *Contact) SetAssistantName(v string)`

SetAssistantName gets a reference to the given string and assigns it to the AssistantName field.

### SetAssistantNameExplicitNull

`func (o *Contact) SetAssistantNameExplicitNull(b bool)`

SetAssistantNameExplicitNull (un)sets AssistantName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AssistantName value is set to nil even if false is passed
### GetManager

`func (o *Contact) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *Contact) GetManagerOk() (string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManager

`func (o *Contact) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManager

`func (o *Contact) SetManager(v string)`

SetManager gets a reference to the given string and assigns it to the Manager field.

### SetManagerExplicitNull

`func (o *Contact) SetManagerExplicitNull(b bool)`

SetManagerExplicitNull (un)sets Manager to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Manager value is set to nil even if false is passed
### GetHomePhones

`func (o *Contact) GetHomePhones() []string`

GetHomePhones returns the HomePhones field if non-nil, zero value otherwise.

### GetHomePhonesOk

`func (o *Contact) GetHomePhonesOk() ([]string, bool)`

GetHomePhonesOk returns a tuple with the HomePhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHomePhones

`func (o *Contact) HasHomePhones() bool`

HasHomePhones returns a boolean if a field has been set.

### SetHomePhones

`func (o *Contact) SetHomePhones(v []string)`

SetHomePhones gets a reference to the given []string and assigns it to the HomePhones field.

### GetMobilePhone

`func (o *Contact) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *Contact) GetMobilePhoneOk() (string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobilePhone

`func (o *Contact) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhone

`func (o *Contact) SetMobilePhone(v string)`

SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.

### SetMobilePhoneExplicitNull

`func (o *Contact) SetMobilePhoneExplicitNull(b bool)`

SetMobilePhoneExplicitNull (un)sets MobilePhone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MobilePhone value is set to nil even if false is passed
### GetBusinessPhones

`func (o *Contact) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *Contact) GetBusinessPhonesOk() ([]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessPhones

`func (o *Contact) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### SetBusinessPhones

`func (o *Contact) SetBusinessPhones(v []string)`

SetBusinessPhones gets a reference to the given []string and assigns it to the BusinessPhones field.

### GetHomeAddress

`func (o *Contact) GetHomeAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetHomeAddress returns the HomeAddress field if non-nil, zero value otherwise.

### GetHomeAddressOk

`func (o *Contact) GetHomeAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetHomeAddressOk returns a tuple with the HomeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHomeAddress

`func (o *Contact) HasHomeAddress() bool`

HasHomeAddress returns a boolean if a field has been set.

### SetHomeAddress

`func (o *Contact) SetHomeAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetHomeAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the HomeAddress field.

### SetHomeAddressExplicitNull

`func (o *Contact) SetHomeAddressExplicitNull(b bool)`

SetHomeAddressExplicitNull (un)sets HomeAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HomeAddress value is set to nil even if false is passed
### GetBusinessAddress

`func (o *Contact) GetBusinessAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetBusinessAddress returns the BusinessAddress field if non-nil, zero value otherwise.

### GetBusinessAddressOk

`func (o *Contact) GetBusinessAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetBusinessAddressOk returns a tuple with the BusinessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessAddress

`func (o *Contact) HasBusinessAddress() bool`

HasBusinessAddress returns a boolean if a field has been set.

### SetBusinessAddress

`func (o *Contact) SetBusinessAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetBusinessAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the BusinessAddress field.

### SetBusinessAddressExplicitNull

`func (o *Contact) SetBusinessAddressExplicitNull(b bool)`

SetBusinessAddressExplicitNull (un)sets BusinessAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BusinessAddress value is set to nil even if false is passed
### GetOtherAddress

`func (o *Contact) GetOtherAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetOtherAddress returns the OtherAddress field if non-nil, zero value otherwise.

### GetOtherAddressOk

`func (o *Contact) GetOtherAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetOtherAddressOk returns a tuple with the OtherAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOtherAddress

`func (o *Contact) HasOtherAddress() bool`

HasOtherAddress returns a boolean if a field has been set.

### SetOtherAddress

`func (o *Contact) SetOtherAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetOtherAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the OtherAddress field.

### SetOtherAddressExplicitNull

`func (o *Contact) SetOtherAddressExplicitNull(b bool)`

SetOtherAddressExplicitNull (un)sets OtherAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OtherAddress value is set to nil even if false is passed
### GetSpouseName

`func (o *Contact) GetSpouseName() string`

GetSpouseName returns the SpouseName field if non-nil, zero value otherwise.

### GetSpouseNameOk

`func (o *Contact) GetSpouseNameOk() (string, bool)`

GetSpouseNameOk returns a tuple with the SpouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpouseName

`func (o *Contact) HasSpouseName() bool`

HasSpouseName returns a boolean if a field has been set.

### SetSpouseName

`func (o *Contact) SetSpouseName(v string)`

SetSpouseName gets a reference to the given string and assigns it to the SpouseName field.

### SetSpouseNameExplicitNull

`func (o *Contact) SetSpouseNameExplicitNull(b bool)`

SetSpouseNameExplicitNull (un)sets SpouseName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SpouseName value is set to nil even if false is passed
### GetPersonalNotes

`func (o *Contact) GetPersonalNotes() string`

GetPersonalNotes returns the PersonalNotes field if non-nil, zero value otherwise.

### GetPersonalNotesOk

`func (o *Contact) GetPersonalNotesOk() (string, bool)`

GetPersonalNotesOk returns a tuple with the PersonalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPersonalNotes

`func (o *Contact) HasPersonalNotes() bool`

HasPersonalNotes returns a boolean if a field has been set.

### SetPersonalNotes

`func (o *Contact) SetPersonalNotes(v string)`

SetPersonalNotes gets a reference to the given string and assigns it to the PersonalNotes field.

### SetPersonalNotesExplicitNull

`func (o *Contact) SetPersonalNotesExplicitNull(b bool)`

SetPersonalNotesExplicitNull (un)sets PersonalNotes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PersonalNotes value is set to nil even if false is passed
### GetChildren

`func (o *Contact) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Contact) GetChildrenOk() ([]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildren

`func (o *Contact) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildren

`func (o *Contact) SetChildren(v []string)`

SetChildren gets a reference to the given []string and assigns it to the Children field.

### GetSingleValueExtendedProperties

`func (o *Contact) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *Contact) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *Contact) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *Contact) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *Contact) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *Contact) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *Contact) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *Contact) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetPhoto

`func (o *Contact) GetPhoto() AnyOfmicrosoftGraphProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *Contact) GetPhotoOk() (AnyOfmicrosoftGraphProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhoto

`func (o *Contact) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhoto

`func (o *Contact) SetPhoto(v AnyOfmicrosoftGraphProfilePhoto)`

SetPhoto gets a reference to the given AnyOfmicrosoftGraphProfilePhoto and assigns it to the Photo field.

### SetPhotoExplicitNull

`func (o *Contact) SetPhotoExplicitNull(b bool)`

SetPhotoExplicitNull (un)sets Photo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Photo value is set to nil even if false is passed
### GetExtensions

`func (o *Contact) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Contact) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *Contact) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *Contact) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


