# EducationSchool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalEmail** | Pointer to **string** |  | [optional] 
**PrincipalName** | Pointer to **string** |  | [optional] 
**ExternalPrincipalId** | Pointer to **string** |  | [optional] 
**LowestGrade** | Pointer to **string** |  | [optional] 
**HighestGrade** | Pointer to **string** |  | [optional] 
**SchoolNumber** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Fax** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**Address** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) |  | [optional] 
**Classes** | Pointer to [**[]MicrosoftGraphEducationClass**](microsoft.graph.educationClass.md) |  | [optional] 
**Users** | Pointer to [**[]MicrosoftGraphEducationUser**](microsoft.graph.educationUser.md) |  | [optional] 

## Methods

### GetPrincipalEmail

`func (o *EducationSchool) GetPrincipalEmail() string`

GetPrincipalEmail returns the PrincipalEmail field if non-nil, zero value otherwise.

### GetPrincipalEmailOk

`func (o *EducationSchool) GetPrincipalEmailOk() (string, bool)`

GetPrincipalEmailOk returns a tuple with the PrincipalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrincipalEmail

`func (o *EducationSchool) HasPrincipalEmail() bool`

HasPrincipalEmail returns a boolean if a field has been set.

### SetPrincipalEmail

`func (o *EducationSchool) SetPrincipalEmail(v string)`

SetPrincipalEmail gets a reference to the given string and assigns it to the PrincipalEmail field.

### SetPrincipalEmailExplicitNull

`func (o *EducationSchool) SetPrincipalEmailExplicitNull(b bool)`

SetPrincipalEmailExplicitNull (un)sets PrincipalEmail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrincipalEmail value is set to nil even if false is passed
### GetPrincipalName

`func (o *EducationSchool) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *EducationSchool) GetPrincipalNameOk() (string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrincipalName

`func (o *EducationSchool) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalName

`func (o *EducationSchool) SetPrincipalName(v string)`

SetPrincipalName gets a reference to the given string and assigns it to the PrincipalName field.

### SetPrincipalNameExplicitNull

`func (o *EducationSchool) SetPrincipalNameExplicitNull(b bool)`

SetPrincipalNameExplicitNull (un)sets PrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrincipalName value is set to nil even if false is passed
### GetExternalPrincipalId

`func (o *EducationSchool) GetExternalPrincipalId() string`

GetExternalPrincipalId returns the ExternalPrincipalId field if non-nil, zero value otherwise.

### GetExternalPrincipalIdOk

`func (o *EducationSchool) GetExternalPrincipalIdOk() (string, bool)`

GetExternalPrincipalIdOk returns a tuple with the ExternalPrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalPrincipalId

`func (o *EducationSchool) HasExternalPrincipalId() bool`

HasExternalPrincipalId returns a boolean if a field has been set.

### SetExternalPrincipalId

`func (o *EducationSchool) SetExternalPrincipalId(v string)`

SetExternalPrincipalId gets a reference to the given string and assigns it to the ExternalPrincipalId field.

### SetExternalPrincipalIdExplicitNull

`func (o *EducationSchool) SetExternalPrincipalIdExplicitNull(b bool)`

SetExternalPrincipalIdExplicitNull (un)sets ExternalPrincipalId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalPrincipalId value is set to nil even if false is passed
### GetLowestGrade

`func (o *EducationSchool) GetLowestGrade() string`

GetLowestGrade returns the LowestGrade field if non-nil, zero value otherwise.

### GetLowestGradeOk

`func (o *EducationSchool) GetLowestGradeOk() (string, bool)`

GetLowestGradeOk returns a tuple with the LowestGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLowestGrade

`func (o *EducationSchool) HasLowestGrade() bool`

HasLowestGrade returns a boolean if a field has been set.

### SetLowestGrade

`func (o *EducationSchool) SetLowestGrade(v string)`

SetLowestGrade gets a reference to the given string and assigns it to the LowestGrade field.

### SetLowestGradeExplicitNull

`func (o *EducationSchool) SetLowestGradeExplicitNull(b bool)`

SetLowestGradeExplicitNull (un)sets LowestGrade to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LowestGrade value is set to nil even if false is passed
### GetHighestGrade

`func (o *EducationSchool) GetHighestGrade() string`

GetHighestGrade returns the HighestGrade field if non-nil, zero value otherwise.

### GetHighestGradeOk

`func (o *EducationSchool) GetHighestGradeOk() (string, bool)`

GetHighestGradeOk returns a tuple with the HighestGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHighestGrade

`func (o *EducationSchool) HasHighestGrade() bool`

HasHighestGrade returns a boolean if a field has been set.

### SetHighestGrade

`func (o *EducationSchool) SetHighestGrade(v string)`

SetHighestGrade gets a reference to the given string and assigns it to the HighestGrade field.

### SetHighestGradeExplicitNull

`func (o *EducationSchool) SetHighestGradeExplicitNull(b bool)`

SetHighestGradeExplicitNull (un)sets HighestGrade to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HighestGrade value is set to nil even if false is passed
### GetSchoolNumber

`func (o *EducationSchool) GetSchoolNumber() string`

GetSchoolNumber returns the SchoolNumber field if non-nil, zero value otherwise.

### GetSchoolNumberOk

`func (o *EducationSchool) GetSchoolNumberOk() (string, bool)`

GetSchoolNumberOk returns a tuple with the SchoolNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchoolNumber

`func (o *EducationSchool) HasSchoolNumber() bool`

HasSchoolNumber returns a boolean if a field has been set.

### SetSchoolNumber

`func (o *EducationSchool) SetSchoolNumber(v string)`

SetSchoolNumber gets a reference to the given string and assigns it to the SchoolNumber field.

### SetSchoolNumberExplicitNull

`func (o *EducationSchool) SetSchoolNumberExplicitNull(b bool)`

SetSchoolNumberExplicitNull (un)sets SchoolNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SchoolNumber value is set to nil even if false is passed
### GetExternalId

`func (o *EducationSchool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EducationSchool) GetExternalIdOk() (string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalId

`func (o *EducationSchool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalId

`func (o *EducationSchool) SetExternalId(v string)`

SetExternalId gets a reference to the given string and assigns it to the ExternalId field.

### SetExternalIdExplicitNull

`func (o *EducationSchool) SetExternalIdExplicitNull(b bool)`

SetExternalIdExplicitNull (un)sets ExternalId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalId value is set to nil even if false is passed
### GetPhone

`func (o *EducationSchool) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *EducationSchool) GetPhoneOk() (string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhone

`func (o *EducationSchool) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhone

`func (o *EducationSchool) SetPhone(v string)`

SetPhone gets a reference to the given string and assigns it to the Phone field.

### SetPhoneExplicitNull

`func (o *EducationSchool) SetPhoneExplicitNull(b bool)`

SetPhoneExplicitNull (un)sets Phone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Phone value is set to nil even if false is passed
### GetFax

`func (o *EducationSchool) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *EducationSchool) GetFaxOk() (string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFax

`func (o *EducationSchool) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFax

`func (o *EducationSchool) SetFax(v string)`

SetFax gets a reference to the given string and assigns it to the Fax field.

### SetFaxExplicitNull

`func (o *EducationSchool) SetFaxExplicitNull(b bool)`

SetFaxExplicitNull (un)sets Fax to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Fax value is set to nil even if false is passed
### GetCreatedBy

`func (o *EducationSchool) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EducationSchool) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *EducationSchool) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *EducationSchool) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *EducationSchool) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetAddress

`func (o *EducationSchool) GetAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EducationSchool) GetAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddress

`func (o *EducationSchool) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddress

`func (o *EducationSchool) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the Address field.

### SetAddressExplicitNull

`func (o *EducationSchool) SetAddressExplicitNull(b bool)`

SetAddressExplicitNull (un)sets Address to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Address value is set to nil even if false is passed
### GetClasses

`func (o *EducationSchool) GetClasses() []MicrosoftGraphEducationClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *EducationSchool) GetClassesOk() ([]MicrosoftGraphEducationClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClasses

`func (o *EducationSchool) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### SetClasses

`func (o *EducationSchool) SetClasses(v []MicrosoftGraphEducationClass)`

SetClasses gets a reference to the given []MicrosoftGraphEducationClass and assigns it to the Classes field.

### GetUsers

`func (o *EducationSchool) GetUsers() []MicrosoftGraphEducationUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *EducationSchool) GetUsersOk() ([]MicrosoftGraphEducationUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsers

`func (o *EducationSchool) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsers

`func (o *EducationSchool) SetUsers(v []MicrosoftGraphEducationUser)`

SetUsers gets a reference to the given []MicrosoftGraphEducationUser and assigns it to the Users field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


