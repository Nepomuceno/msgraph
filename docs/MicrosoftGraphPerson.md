# MicrosoftGraphPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 
**Birthday** | Pointer to **string** |  | [optional] 
**PersonNotes** | Pointer to **string** |  | [optional] 
**IsFavorite** | Pointer to **bool** |  | [optional] 
**ScoredEmailAddresses** | Pointer to [**[]AnyOfmicrosoftGraphScoredEmailAddress**](anyOf&lt;microsoft.graph.scoredEmailAddress&gt;.md) |  | [optional] 
**Phones** | Pointer to [**[]AnyOfmicrosoftGraphPhone**](anyOf&lt;microsoft.graph.phone&gt;.md) |  | [optional] 
**PostalAddresses** | Pointer to [**[]AnyOfmicrosoftGraphLocation**](anyOf&lt;microsoft.graph.location&gt;.md) |  | [optional] 
**Websites** | Pointer to [**[]AnyOfmicrosoftGraphWebsite**](anyOf&lt;microsoft.graph.website&gt;.md) |  | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**YomiCompany** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**OfficeLocation** | Pointer to **string** |  | [optional] 
**Profession** | Pointer to **string** |  | [optional] 
**PersonType** | Pointer to [**AnyOfmicrosoftGraphPersonType**](anyOf&lt;microsoft.graph.personType&gt;.md) |  | [optional] 
**UserPrincipalName** | Pointer to **string** |  | [optional] 
**ImAddress** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphPerson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPerson) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphPerson) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphPerson) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphPerson) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphPerson) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphPerson) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphPerson) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphPerson) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetGivenName

`func (o *MicrosoftGraphPerson) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *MicrosoftGraphPerson) GetGivenNameOk() (string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGivenName

`func (o *MicrosoftGraphPerson) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### SetGivenName

`func (o *MicrosoftGraphPerson) SetGivenName(v string)`

SetGivenName gets a reference to the given string and assigns it to the GivenName field.

### SetGivenNameExplicitNull

`func (o *MicrosoftGraphPerson) SetGivenNameExplicitNull(b bool)`

SetGivenNameExplicitNull (un)sets GivenName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GivenName value is set to nil even if false is passed
### GetSurname

`func (o *MicrosoftGraphPerson) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *MicrosoftGraphPerson) GetSurnameOk() (string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSurname

`func (o *MicrosoftGraphPerson) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurname

`func (o *MicrosoftGraphPerson) SetSurname(v string)`

SetSurname gets a reference to the given string and assigns it to the Surname field.

### SetSurnameExplicitNull

`func (o *MicrosoftGraphPerson) SetSurnameExplicitNull(b bool)`

SetSurnameExplicitNull (un)sets Surname to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Surname value is set to nil even if false is passed
### GetBirthday

`func (o *MicrosoftGraphPerson) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *MicrosoftGraphPerson) GetBirthdayOk() (string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBirthday

`func (o *MicrosoftGraphPerson) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthday

`func (o *MicrosoftGraphPerson) SetBirthday(v string)`

SetBirthday gets a reference to the given string and assigns it to the Birthday field.

### SetBirthdayExplicitNull

`func (o *MicrosoftGraphPerson) SetBirthdayExplicitNull(b bool)`

SetBirthdayExplicitNull (un)sets Birthday to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Birthday value is set to nil even if false is passed
### GetPersonNotes

`func (o *MicrosoftGraphPerson) GetPersonNotes() string`

GetPersonNotes returns the PersonNotes field if non-nil, zero value otherwise.

### GetPersonNotesOk

`func (o *MicrosoftGraphPerson) GetPersonNotesOk() (string, bool)`

GetPersonNotesOk returns a tuple with the PersonNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPersonNotes

`func (o *MicrosoftGraphPerson) HasPersonNotes() bool`

HasPersonNotes returns a boolean if a field has been set.

### SetPersonNotes

`func (o *MicrosoftGraphPerson) SetPersonNotes(v string)`

SetPersonNotes gets a reference to the given string and assigns it to the PersonNotes field.

### SetPersonNotesExplicitNull

`func (o *MicrosoftGraphPerson) SetPersonNotesExplicitNull(b bool)`

SetPersonNotesExplicitNull (un)sets PersonNotes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PersonNotes value is set to nil even if false is passed
### GetIsFavorite

`func (o *MicrosoftGraphPerson) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *MicrosoftGraphPerson) GetIsFavoriteOk() (bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsFavorite

`func (o *MicrosoftGraphPerson) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavorite

`func (o *MicrosoftGraphPerson) SetIsFavorite(v bool)`

SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.

### SetIsFavoriteExplicitNull

`func (o *MicrosoftGraphPerson) SetIsFavoriteExplicitNull(b bool)`

SetIsFavoriteExplicitNull (un)sets IsFavorite to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsFavorite value is set to nil even if false is passed
### GetScoredEmailAddresses

`func (o *MicrosoftGraphPerson) GetScoredEmailAddresses() []AnyOfmicrosoftGraphScoredEmailAddress`

GetScoredEmailAddresses returns the ScoredEmailAddresses field if non-nil, zero value otherwise.

### GetScoredEmailAddressesOk

`func (o *MicrosoftGraphPerson) GetScoredEmailAddressesOk() ([]AnyOfmicrosoftGraphScoredEmailAddress, bool)`

GetScoredEmailAddressesOk returns a tuple with the ScoredEmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScoredEmailAddresses

`func (o *MicrosoftGraphPerson) HasScoredEmailAddresses() bool`

HasScoredEmailAddresses returns a boolean if a field has been set.

### SetScoredEmailAddresses

`func (o *MicrosoftGraphPerson) SetScoredEmailAddresses(v []AnyOfmicrosoftGraphScoredEmailAddress)`

SetScoredEmailAddresses gets a reference to the given []AnyOfmicrosoftGraphScoredEmailAddress and assigns it to the ScoredEmailAddresses field.

### GetPhones

`func (o *MicrosoftGraphPerson) GetPhones() []AnyOfmicrosoftGraphPhone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *MicrosoftGraphPerson) GetPhonesOk() ([]AnyOfmicrosoftGraphPhone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhones

`func (o *MicrosoftGraphPerson) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### SetPhones

`func (o *MicrosoftGraphPerson) SetPhones(v []AnyOfmicrosoftGraphPhone)`

SetPhones gets a reference to the given []AnyOfmicrosoftGraphPhone and assigns it to the Phones field.

### GetPostalAddresses

`func (o *MicrosoftGraphPerson) GetPostalAddresses() []AnyOfmicrosoftGraphLocation`

GetPostalAddresses returns the PostalAddresses field if non-nil, zero value otherwise.

### GetPostalAddressesOk

`func (o *MicrosoftGraphPerson) GetPostalAddressesOk() ([]AnyOfmicrosoftGraphLocation, bool)`

GetPostalAddressesOk returns a tuple with the PostalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPostalAddresses

`func (o *MicrosoftGraphPerson) HasPostalAddresses() bool`

HasPostalAddresses returns a boolean if a field has been set.

### SetPostalAddresses

`func (o *MicrosoftGraphPerson) SetPostalAddresses(v []AnyOfmicrosoftGraphLocation)`

SetPostalAddresses gets a reference to the given []AnyOfmicrosoftGraphLocation and assigns it to the PostalAddresses field.

### GetWebsites

`func (o *MicrosoftGraphPerson) GetWebsites() []AnyOfmicrosoftGraphWebsite`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *MicrosoftGraphPerson) GetWebsitesOk() ([]AnyOfmicrosoftGraphWebsite, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebsites

`func (o *MicrosoftGraphPerson) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.

### SetWebsites

`func (o *MicrosoftGraphPerson) SetWebsites(v []AnyOfmicrosoftGraphWebsite)`

SetWebsites gets a reference to the given []AnyOfmicrosoftGraphWebsite and assigns it to the Websites field.

### GetJobTitle

`func (o *MicrosoftGraphPerson) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *MicrosoftGraphPerson) GetJobTitleOk() (string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJobTitle

`func (o *MicrosoftGraphPerson) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitle

`func (o *MicrosoftGraphPerson) SetJobTitle(v string)`

SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.

### SetJobTitleExplicitNull

`func (o *MicrosoftGraphPerson) SetJobTitleExplicitNull(b bool)`

SetJobTitleExplicitNull (un)sets JobTitle to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The JobTitle value is set to nil even if false is passed
### GetCompanyName

`func (o *MicrosoftGraphPerson) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *MicrosoftGraphPerson) GetCompanyNameOk() (string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompanyName

`func (o *MicrosoftGraphPerson) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyName

`func (o *MicrosoftGraphPerson) SetCompanyName(v string)`

SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.

### SetCompanyNameExplicitNull

`func (o *MicrosoftGraphPerson) SetCompanyNameExplicitNull(b bool)`

SetCompanyNameExplicitNull (un)sets CompanyName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CompanyName value is set to nil even if false is passed
### GetYomiCompany

`func (o *MicrosoftGraphPerson) GetYomiCompany() string`

GetYomiCompany returns the YomiCompany field if non-nil, zero value otherwise.

### GetYomiCompanyOk

`func (o *MicrosoftGraphPerson) GetYomiCompanyOk() (string, bool)`

GetYomiCompanyOk returns a tuple with the YomiCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasYomiCompany

`func (o *MicrosoftGraphPerson) HasYomiCompany() bool`

HasYomiCompany returns a boolean if a field has been set.

### SetYomiCompany

`func (o *MicrosoftGraphPerson) SetYomiCompany(v string)`

SetYomiCompany gets a reference to the given string and assigns it to the YomiCompany field.

### SetYomiCompanyExplicitNull

`func (o *MicrosoftGraphPerson) SetYomiCompanyExplicitNull(b bool)`

SetYomiCompanyExplicitNull (un)sets YomiCompany to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The YomiCompany value is set to nil even if false is passed
### GetDepartment

`func (o *MicrosoftGraphPerson) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *MicrosoftGraphPerson) GetDepartmentOk() (string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDepartment

`func (o *MicrosoftGraphPerson) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartment

`func (o *MicrosoftGraphPerson) SetDepartment(v string)`

SetDepartment gets a reference to the given string and assigns it to the Department field.

### SetDepartmentExplicitNull

`func (o *MicrosoftGraphPerson) SetDepartmentExplicitNull(b bool)`

SetDepartmentExplicitNull (un)sets Department to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Department value is set to nil even if false is passed
### GetOfficeLocation

`func (o *MicrosoftGraphPerson) GetOfficeLocation() string`

GetOfficeLocation returns the OfficeLocation field if non-nil, zero value otherwise.

### GetOfficeLocationOk

`func (o *MicrosoftGraphPerson) GetOfficeLocationOk() (string, bool)`

GetOfficeLocationOk returns a tuple with the OfficeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOfficeLocation

`func (o *MicrosoftGraphPerson) HasOfficeLocation() bool`

HasOfficeLocation returns a boolean if a field has been set.

### SetOfficeLocation

`func (o *MicrosoftGraphPerson) SetOfficeLocation(v string)`

SetOfficeLocation gets a reference to the given string and assigns it to the OfficeLocation field.

### SetOfficeLocationExplicitNull

`func (o *MicrosoftGraphPerson) SetOfficeLocationExplicitNull(b bool)`

SetOfficeLocationExplicitNull (un)sets OfficeLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OfficeLocation value is set to nil even if false is passed
### GetProfession

`func (o *MicrosoftGraphPerson) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *MicrosoftGraphPerson) GetProfessionOk() (string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfession

`func (o *MicrosoftGraphPerson) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### SetProfession

`func (o *MicrosoftGraphPerson) SetProfession(v string)`

SetProfession gets a reference to the given string and assigns it to the Profession field.

### SetProfessionExplicitNull

`func (o *MicrosoftGraphPerson) SetProfessionExplicitNull(b bool)`

SetProfessionExplicitNull (un)sets Profession to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Profession value is set to nil even if false is passed
### GetPersonType

`func (o *MicrosoftGraphPerson) GetPersonType() AnyOfmicrosoftGraphPersonType`

GetPersonType returns the PersonType field if non-nil, zero value otherwise.

### GetPersonTypeOk

`func (o *MicrosoftGraphPerson) GetPersonTypeOk() (AnyOfmicrosoftGraphPersonType, bool)`

GetPersonTypeOk returns a tuple with the PersonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPersonType

`func (o *MicrosoftGraphPerson) HasPersonType() bool`

HasPersonType returns a boolean if a field has been set.

### SetPersonType

`func (o *MicrosoftGraphPerson) SetPersonType(v AnyOfmicrosoftGraphPersonType)`

SetPersonType gets a reference to the given AnyOfmicrosoftGraphPersonType and assigns it to the PersonType field.

### SetPersonTypeExplicitNull

`func (o *MicrosoftGraphPerson) SetPersonTypeExplicitNull(b bool)`

SetPersonTypeExplicitNull (un)sets PersonType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PersonType value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *MicrosoftGraphPerson) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphPerson) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *MicrosoftGraphPerson) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphPerson) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphPerson) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetImAddress

`func (o *MicrosoftGraphPerson) GetImAddress() string`

GetImAddress returns the ImAddress field if non-nil, zero value otherwise.

### GetImAddressOk

`func (o *MicrosoftGraphPerson) GetImAddressOk() (string, bool)`

GetImAddressOk returns a tuple with the ImAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImAddress

`func (o *MicrosoftGraphPerson) HasImAddress() bool`

HasImAddress returns a boolean if a field has been set.

### SetImAddress

`func (o *MicrosoftGraphPerson) SetImAddress(v string)`

SetImAddress gets a reference to the given string and assigns it to the ImAddress field.

### SetImAddressExplicitNull

`func (o *MicrosoftGraphPerson) SetImAddressExplicitNull(b bool)`

SetImAddressExplicitNull (un)sets ImAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ImAddress value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


