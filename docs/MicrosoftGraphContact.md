# MicrosoftGraphContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ChangeKey** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphContact) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphContact) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphContact) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphContact) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphContact) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphContact) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphContact) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphContact) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphContact) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphContact) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphContact) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphContact) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphContact) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetChangeKey

`func (o *MicrosoftGraphContact) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphContact) GetChangeKeyOk() (string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangeKey

`func (o *MicrosoftGraphContact) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKey

`func (o *MicrosoftGraphContact) SetChangeKey(v string)`

SetChangeKey gets a reference to the given string and assigns it to the ChangeKey field.

### SetChangeKeyExplicitNull

`func (o *MicrosoftGraphContact) SetChangeKeyExplicitNull(b bool)`

SetChangeKeyExplicitNull (un)sets ChangeKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChangeKey value is set to nil even if false is passed
### GetCategories

`func (o *MicrosoftGraphContact) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphContact) GetCategoriesOk() ([]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *MicrosoftGraphContact) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *MicrosoftGraphContact) SetCategories(v []string)`

SetCategories gets a reference to the given []string and assigns it to the Categories field.

### GetParentFolderId

`func (o *MicrosoftGraphContact) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MicrosoftGraphContact) GetParentFolderIdOk() (string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentFolderId

`func (o *MicrosoftGraphContact) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderId

`func (o *MicrosoftGraphContact) SetParentFolderId(v string)`

SetParentFolderId gets a reference to the given string and assigns it to the ParentFolderId field.

### SetParentFolderIdExplicitNull

`func (o *MicrosoftGraphContact) SetParentFolderIdExplicitNull(b bool)`

SetParentFolderIdExplicitNull (un)sets ParentFolderId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentFolderId value is set to nil even if false is passed
### GetBirthday

`func (o *MicrosoftGraphContact) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *MicrosoftGraphContact) GetBirthdayOk() (time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBirthday

`func (o *MicrosoftGraphContact) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthday

`func (o *MicrosoftGraphContact) SetBirthday(v time.Time)`

SetBirthday gets a reference to the given time.Time and assigns it to the Birthday field.

### SetBirthdayExplicitNull

`func (o *MicrosoftGraphContact) SetBirthdayExplicitNull(b bool)`

SetBirthdayExplicitNull (un)sets Birthday to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Birthday value is set to nil even if false is passed
### GetFileAs

`func (o *MicrosoftGraphContact) GetFileAs() string`

GetFileAs returns the FileAs field if non-nil, zero value otherwise.

### GetFileAsOk

`func (o *MicrosoftGraphContact) GetFileAsOk() (string, bool)`

GetFileAsOk returns a tuple with the FileAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileAs

`func (o *MicrosoftGraphContact) HasFileAs() bool`

HasFileAs returns a boolean if a field has been set.

### SetFileAs

`func (o *MicrosoftGraphContact) SetFileAs(v string)`

SetFileAs gets a reference to the given string and assigns it to the FileAs field.

### SetFileAsExplicitNull

`func (o *MicrosoftGraphContact) SetFileAsExplicitNull(b bool)`

SetFileAsExplicitNull (un)sets FileAs to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileAs value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphContact) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphContact) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphContact) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphContact) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphContact) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetGivenName

`func (o *MicrosoftGraphContact) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MicrosoftGraphContact) GetGivenNameOk() (string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGivenName

`func (o *MicrosoftGraphContact) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenName

`func (o *MicrosoftGraphContact) SetGivenName(v string)`

SetGivenName gets a reference to the given string and assigns it to the GivenName field.

### SetGivenNameExplicitNull

`func (o *MicrosoftGraphContact) SetGivenNameExplicitNull(b bool)`

SetGivenNameExplicitNull (un)sets GivenName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GivenName value is set to nil even if false is passed
### GetInitials

`func (o *MicrosoftGraphContact) GetInitials() string`

GetInitials returns the Initials field if non-nil, zero value otherwise.

### GetInitialsOk

`func (o *MicrosoftGraphContact) GetInitialsOk() (string, bool)`

GetInitialsOk returns a tuple with the Initials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInitials

`func (o *MicrosoftGraphContact) HasInitials() bool`

HasInitials returns a boolean if a field has been set.

### SetInitials

`func (o *MicrosoftGraphContact) SetInitials(v string)`

SetInitials gets a reference to the given string and assigns it to the Initials field.

### SetInitialsExplicitNull

`func (o *MicrosoftGraphContact) SetInitialsExplicitNull(b bool)`

SetInitialsExplicitNull (un)sets Initials to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Initials value is set to nil even if false is passed
### GetMiddleName

`func (o *MicrosoftGraphContact) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *MicrosoftGraphContact) GetMiddleNameOk() (string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMiddleName

`func (o *MicrosoftGraphContact) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleName

`func (o *MicrosoftGraphContact) SetMiddleName(v string)`

SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.

### SetMiddleNameExplicitNull

`func (o *MicrosoftGraphContact) SetMiddleNameExplicitNull(b bool)`

SetMiddleNameExplicitNull (un)sets MiddleName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MiddleName value is set to nil even if false is passed
### GetNickName

`func (o *MicrosoftGraphContact) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *MicrosoftGraphContact) GetNickNameOk() (string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNickName

`func (o *MicrosoftGraphContact) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickName

`func (o *MicrosoftGraphContact) SetNickName(v string)`

SetNickName gets a reference to the given string and assigns it to the NickName field.

### SetNickNameExplicitNull

`func (o *MicrosoftGraphContact) SetNickNameExplicitNull(b bool)`

SetNickNameExplicitNull (un)sets NickName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NickName value is set to nil even if false is passed
### GetSurname

`func (o *MicrosoftGraphContact) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *MicrosoftGraphContact) GetSurnameOk() (string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSurname

`func (o *MicrosoftGraphContact) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurname

`func (o *MicrosoftGraphContact) SetSurname(v string)`

SetSurname gets a reference to the given string and assigns it to the Surname field.

### SetSurnameExplicitNull

`func (o *MicrosoftGraphContact) SetSurnameExplicitNull(b bool)`

SetSurnameExplicitNull (un)sets Surname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Surname value is set to nil even if false is passed
### GetTitle

`func (o *MicrosoftGraphContact) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphContact) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *MicrosoftGraphContact) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *MicrosoftGraphContact) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### SetTitleExplicitNull

`func (o *MicrosoftGraphContact) SetTitleExplicitNull(b bool)`

SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Title value is set to nil even if false is passed
### GetYomiGivenName

`func (o *MicrosoftGraphContact) GetYomiGivenName() string`

GetYomiGivenName returns the YomiGivenName field if non-nil, zero value otherwise.

### GetYomiGivenNameOk

`func (o *MicrosoftGraphContact) GetYomiGivenNameOk() (string, bool)`

GetYomiGivenNameOk returns a tuple with the YomiGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasYomiGivenName

`func (o *MicrosoftGraphContact) HasYomiGivenName() bool`

HasYomiGivenName returns a boolean if a field has been set.

### SetYomiGivenName

`func (o *MicrosoftGraphContact) SetYomiGivenName(v string)`

SetYomiGivenName gets a reference to the given string and assigns it to the YomiGivenName field.

### SetYomiGivenNameExplicitNull

`func (o *MicrosoftGraphContact) SetYomiGivenNameExplicitNull(b bool)`

SetYomiGivenNameExplicitNull (un)sets YomiGivenName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The YomiGivenName value is set to nil even if false is passed
### GetYomiSurname

`func (o *MicrosoftGraphContact) GetYomiSurname() string`

GetYomiSurname returns the YomiSurname field if non-nil, zero value otherwise.

### GetYomiSurnameOk

`func (o *MicrosoftGraphContact) GetYomiSurnameOk() (string, bool)`

GetYomiSurnameOk returns a tuple with the YomiSurname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasYomiSurname

`func (o *MicrosoftGraphContact) HasYomiSurname() bool`

HasYomiSurname returns a boolean if a field has been set.

### SetYomiSurname

`func (o *MicrosoftGraphContact) SetYomiSurname(v string)`

SetYomiSurname gets a reference to the given string and assigns it to the YomiSurname field.

### SetYomiSurnameExplicitNull

`func (o *MicrosoftGraphContact) SetYomiSurnameExplicitNull(b bool)`

SetYomiSurnameExplicitNull (un)sets YomiSurname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The YomiSurname value is set to nil even if false is passed
### GetYomiCompanyName

`func (o *MicrosoftGraphContact) GetYomiCompanyName() string`

GetYomiCompanyName returns the YomiCompanyName field if non-nil, zero value otherwise.

### GetYomiCompanyNameOk

`func (o *MicrosoftGraphContact) GetYomiCompanyNameOk() (string, bool)`

GetYomiCompanyNameOk returns a tuple with the YomiCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasYomiCompanyName

`func (o *MicrosoftGraphContact) HasYomiCompanyName() bool`

HasYomiCompanyName returns a boolean if a field has been set.

### SetYomiCompanyName

`func (o *MicrosoftGraphContact) SetYomiCompanyName(v string)`

SetYomiCompanyName gets a reference to the given string and assigns it to the YomiCompanyName field.

### SetYomiCompanyNameExplicitNull

`func (o *MicrosoftGraphContact) SetYomiCompanyNameExplicitNull(b bool)`

SetYomiCompanyNameExplicitNull (un)sets YomiCompanyName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The YomiCompanyName value is set to nil even if false is passed
### GetGeneration

`func (o *MicrosoftGraphContact) GetGeneration() string`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *MicrosoftGraphContact) GetGenerationOk() (string, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGeneration

`func (o *MicrosoftGraphContact) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### SetGeneration

`func (o *MicrosoftGraphContact) SetGeneration(v string)`

SetGeneration gets a reference to the given string and assigns it to the Generation field.

### SetGenerationExplicitNull

`func (o *MicrosoftGraphContact) SetGenerationExplicitNull(b bool)`

SetGenerationExplicitNull (un)sets Generation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Generation value is set to nil even if false is passed
### GetEmailAddresses

`func (o *MicrosoftGraphContact) GetEmailAddresses() []AnyOfmicrosoftGraphEmailAddress`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *MicrosoftGraphContact) GetEmailAddressesOk() ([]AnyOfmicrosoftGraphEmailAddress, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailAddresses

`func (o *MicrosoftGraphContact) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### SetEmailAddresses

`func (o *MicrosoftGraphContact) SetEmailAddresses(v []AnyOfmicrosoftGraphEmailAddress)`

SetEmailAddresses gets a reference to the given []AnyOfmicrosoftGraphEmailAddress and assigns it to the EmailAddresses field.

### GetImAddresses

`func (o *MicrosoftGraphContact) GetImAddresses() []string`

GetImAddresses returns the ImAddresses field if non-nil, zero value otherwise.

### GetImAddressesOk

`func (o *MicrosoftGraphContact) GetImAddressesOk() ([]string, bool)`

GetImAddressesOk returns a tuple with the ImAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImAddresses

`func (o *MicrosoftGraphContact) HasImAddresses() bool`

HasImAddresses returns a boolean if a field has been set.

### SetImAddresses

`func (o *MicrosoftGraphContact) SetImAddresses(v []string)`

SetImAddresses gets a reference to the given []string and assigns it to the ImAddresses field.

### GetJobTitle

`func (o *MicrosoftGraphContact) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *MicrosoftGraphContact) GetJobTitleOk() (string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJobTitle

`func (o *MicrosoftGraphContact) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitle

`func (o *MicrosoftGraphContact) SetJobTitle(v string)`

SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.

### SetJobTitleExplicitNull

`func (o *MicrosoftGraphContact) SetJobTitleExplicitNull(b bool)`

SetJobTitleExplicitNull (un)sets JobTitle to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The JobTitle value is set to nil even if false is passed
### GetCompanyName

`func (o *MicrosoftGraphContact) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *MicrosoftGraphContact) GetCompanyNameOk() (string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompanyName

`func (o *MicrosoftGraphContact) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyName

`func (o *MicrosoftGraphContact) SetCompanyName(v string)`

SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.

### SetCompanyNameExplicitNull

`func (o *MicrosoftGraphContact) SetCompanyNameExplicitNull(b bool)`

SetCompanyNameExplicitNull (un)sets CompanyName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CompanyName value is set to nil even if false is passed
### GetDepartment

`func (o *MicrosoftGraphContact) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *MicrosoftGraphContact) GetDepartmentOk() (string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDepartment

`func (o *MicrosoftGraphContact) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartment

`func (o *MicrosoftGraphContact) SetDepartment(v string)`

SetDepartment gets a reference to the given string and assigns it to the Department field.

### SetDepartmentExplicitNull

`func (o *MicrosoftGraphContact) SetDepartmentExplicitNull(b bool)`

SetDepartmentExplicitNull (un)sets Department to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Department value is set to nil even if false is passed
### GetOfficeLocation

`func (o *MicrosoftGraphContact) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *MicrosoftGraphContact) GetOfficeLocationOk() (string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOfficeLocation

`func (o *MicrosoftGraphContact) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocation

`func (o *MicrosoftGraphContact) SetOfficeLocation(v string)`

SetOfficeLocation gets a reference to the given string and assigns it to the OfficeLocation field.

### SetOfficeLocationExplicitNull

`func (o *MicrosoftGraphContact) SetOfficeLocationExplicitNull(b bool)`

SetOfficeLocationExplicitNull (un)sets OfficeLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OfficeLocation value is set to nil even if false is passed
### GetProfession

`func (o *MicrosoftGraphContact) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *MicrosoftGraphContact) GetProfessionOk() (string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfession

`func (o *MicrosoftGraphContact) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### SetProfession

`func (o *MicrosoftGraphContact) SetProfession(v string)`

SetProfession gets a reference to the given string and assigns it to the Profession field.

### SetProfessionExplicitNull

`func (o *MicrosoftGraphContact) SetProfessionExplicitNull(b bool)`

SetProfessionExplicitNull (un)sets Profession to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Profession value is set to nil even if false is passed
### GetBusinessHomePage

`func (o *MicrosoftGraphContact) GetBusinessHomePage() string`

GetBusinessHomePage returns the BusinessHomePage field if non-nil, zero value otherwise.

### GetBusinessHomePageOk

`func (o *MicrosoftGraphContact) GetBusinessHomePageOk() (string, bool)`

GetBusinessHomePageOk returns a tuple with the BusinessHomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessHomePage

`func (o *MicrosoftGraphContact) HasBusinessHomePage() bool`

HasBusinessHomePage returns a boolean if a field has been set.

### SetBusinessHomePage

`func (o *MicrosoftGraphContact) SetBusinessHomePage(v string)`

SetBusinessHomePage gets a reference to the given string and assigns it to the BusinessHomePage field.

### SetBusinessHomePageExplicitNull

`func (o *MicrosoftGraphContact) SetBusinessHomePageExplicitNull(b bool)`

SetBusinessHomePageExplicitNull (un)sets BusinessHomePage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BusinessHomePage value is set to nil even if false is passed
### GetAssistantName

`func (o *MicrosoftGraphContact) GetAssistantName() string`

GetAssistantName returns the AssistantName field if non-nil, zero value otherwise.

### GetAssistantNameOk

`func (o *MicrosoftGraphContact) GetAssistantNameOk() (string, bool)`

GetAssistantNameOk returns a tuple with the AssistantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssistantName

`func (o *MicrosoftGraphContact) HasAssistantName() bool`

HasAssistantName returns a boolean if a field has been set.

### SetAssistantName

`func (o *MicrosoftGraphContact) SetAssistantName(v string)`

SetAssistantName gets a reference to the given string and assigns it to the AssistantName field.

### SetAssistantNameExplicitNull

`func (o *MicrosoftGraphContact) SetAssistantNameExplicitNull(b bool)`

SetAssistantNameExplicitNull (un)sets AssistantName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AssistantName value is set to nil even if false is passed
### GetManager

`func (o *MicrosoftGraphContact) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *MicrosoftGraphContact) GetManagerOk() (string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManager

`func (o *MicrosoftGraphContact) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManager

`func (o *MicrosoftGraphContact) SetManager(v string)`

SetManager gets a reference to the given string and assigns it to the Manager field.

### SetManagerExplicitNull

`func (o *MicrosoftGraphContact) SetManagerExplicitNull(b bool)`

SetManagerExplicitNull (un)sets Manager to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Manager value is set to nil even if false is passed
### GetHomePhones

`func (o *MicrosoftGraphContact) GetHomePhones() []string`

GetHomePhones returns the HomePhones field if non-nil, zero value otherwise.

### GetHomePhonesOk

`func (o *MicrosoftGraphContact) GetHomePhonesOk() ([]string, bool)`

GetHomePhonesOk returns a tuple with the HomePhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHomePhones

`func (o *MicrosoftGraphContact) HasHomePhones() bool`

HasHomePhones returns a boolean if a field has been set.

### SetHomePhones

`func (o *MicrosoftGraphContact) SetHomePhones(v []string)`

SetHomePhones gets a reference to the given []string and assigns it to the HomePhones field.

### GetMobilePhone

`func (o *MicrosoftGraphContact) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *MicrosoftGraphContact) GetMobilePhoneOk() (string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobilePhone

`func (o *MicrosoftGraphContact) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhone

`func (o *MicrosoftGraphContact) SetMobilePhone(v string)`

SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.

### SetMobilePhoneExplicitNull

`func (o *MicrosoftGraphContact) SetMobilePhoneExplicitNull(b bool)`

SetMobilePhoneExplicitNull (un)sets MobilePhone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MobilePhone value is set to nil even if false is passed
### GetBusinessPhones

`func (o *MicrosoftGraphContact) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *MicrosoftGraphContact) GetBusinessPhonesOk() ([]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessPhones

`func (o *MicrosoftGraphContact) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### SetBusinessPhones

`func (o *MicrosoftGraphContact) SetBusinessPhones(v []string)`

SetBusinessPhones gets a reference to the given []string and assigns it to the BusinessPhones field.

### GetHomeAddress

`func (o *MicrosoftGraphContact) GetHomeAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetHomeAddress returns the HomeAddress field if non-nil, zero value otherwise.

### GetHomeAddressOk

`func (o *MicrosoftGraphContact) GetHomeAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetHomeAddressOk returns a tuple with the HomeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHomeAddress

`func (o *MicrosoftGraphContact) HasHomeAddress() bool`

HasHomeAddress returns a boolean if a field has been set.

### SetHomeAddress

`func (o *MicrosoftGraphContact) SetHomeAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetHomeAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the HomeAddress field.

### SetHomeAddressExplicitNull

`func (o *MicrosoftGraphContact) SetHomeAddressExplicitNull(b bool)`

SetHomeAddressExplicitNull (un)sets HomeAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HomeAddress value is set to nil even if false is passed
### GetBusinessAddress

`func (o *MicrosoftGraphContact) GetBusinessAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetBusinessAddress returns the BusinessAddress field if non-nil, zero value otherwise.

### GetBusinessAddressOk

`func (o *MicrosoftGraphContact) GetBusinessAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetBusinessAddressOk returns a tuple with the BusinessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessAddress

`func (o *MicrosoftGraphContact) HasBusinessAddress() bool`

HasBusinessAddress returns a boolean if a field has been set.

### SetBusinessAddress

`func (o *MicrosoftGraphContact) SetBusinessAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetBusinessAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the BusinessAddress field.

### SetBusinessAddressExplicitNull

`func (o *MicrosoftGraphContact) SetBusinessAddressExplicitNull(b bool)`

SetBusinessAddressExplicitNull (un)sets BusinessAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BusinessAddress value is set to nil even if false is passed
### GetOtherAddress

`func (o *MicrosoftGraphContact) GetOtherAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetOtherAddress returns the OtherAddress field if non-nil, zero value otherwise.

### GetOtherAddressOk

`func (o *MicrosoftGraphContact) GetOtherAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetOtherAddressOk returns a tuple with the OtherAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOtherAddress

`func (o *MicrosoftGraphContact) HasOtherAddress() bool`

HasOtherAddress returns a boolean if a field has been set.

### SetOtherAddress

`func (o *MicrosoftGraphContact) SetOtherAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetOtherAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the OtherAddress field.

### SetOtherAddressExplicitNull

`func (o *MicrosoftGraphContact) SetOtherAddressExplicitNull(b bool)`

SetOtherAddressExplicitNull (un)sets OtherAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OtherAddress value is set to nil even if false is passed
### GetSpouseName

`func (o *MicrosoftGraphContact) GetSpouseName() string`

GetSpouseName returns the SpouseName field if non-nil, zero value otherwise.

### GetSpouseNameOk

`func (o *MicrosoftGraphContact) GetSpouseNameOk() (string, bool)`

GetSpouseNameOk returns a tuple with the SpouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpouseName

`func (o *MicrosoftGraphContact) HasSpouseName() bool`

HasSpouseName returns a boolean if a field has been set.

### SetSpouseName

`func (o *MicrosoftGraphContact) SetSpouseName(v string)`

SetSpouseName gets a reference to the given string and assigns it to the SpouseName field.

### SetSpouseNameExplicitNull

`func (o *MicrosoftGraphContact) SetSpouseNameExplicitNull(b bool)`

SetSpouseNameExplicitNull (un)sets SpouseName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SpouseName value is set to nil even if false is passed
### GetPersonalNotes

`func (o *MicrosoftGraphContact) GetPersonalNotes() string`

GetPersonalNotes returns the PersonalNotes field if non-nil, zero value otherwise.

### GetPersonalNotesOk

`func (o *MicrosoftGraphContact) GetPersonalNotesOk() (string, bool)`

GetPersonalNotesOk returns a tuple with the PersonalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPersonalNotes

`func (o *MicrosoftGraphContact) HasPersonalNotes() bool`

HasPersonalNotes returns a boolean if a field has been set.

### SetPersonalNotes

`func (o *MicrosoftGraphContact) SetPersonalNotes(v string)`

SetPersonalNotes gets a reference to the given string and assigns it to the PersonalNotes field.

### SetPersonalNotesExplicitNull

`func (o *MicrosoftGraphContact) SetPersonalNotesExplicitNull(b bool)`

SetPersonalNotesExplicitNull (un)sets PersonalNotes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PersonalNotes value is set to nil even if false is passed
### GetChildren

`func (o *MicrosoftGraphContact) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MicrosoftGraphContact) GetChildrenOk() ([]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildren

`func (o *MicrosoftGraphContact) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildren

`func (o *MicrosoftGraphContact) SetChildren(v []string)`

SetChildren gets a reference to the given []string and assigns it to the Children field.

### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphContact) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphContact) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphContact) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphContact) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphContact) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphContact) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphContact) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphContact) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetPhoto

`func (o *MicrosoftGraphContact) GetPhoto() AnyOfmicrosoftGraphProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MicrosoftGraphContact) GetPhotoOk() (AnyOfmicrosoftGraphProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhoto

`func (o *MicrosoftGraphContact) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhoto

`func (o *MicrosoftGraphContact) SetPhoto(v AnyOfmicrosoftGraphProfilePhoto)`

SetPhoto gets a reference to the given AnyOfmicrosoftGraphProfilePhoto and assigns it to the Photo field.

### SetPhotoExplicitNull

`func (o *MicrosoftGraphContact) SetPhotoExplicitNull(b bool)`

SetPhotoExplicitNull (un)sets Photo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Photo value is set to nil even if false is passed
### GetExtensions

`func (o *MicrosoftGraphContact) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphContact) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *MicrosoftGraphContact) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *MicrosoftGraphContact) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


