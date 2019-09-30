# MicrosoftGraphEducationSchool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalSource** | Pointer to [**AnyOfmicrosoftGraphEducationExternalSource**](anyOf&lt;microsoft.graph.educationExternalSource&gt;.md) |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphEducationSchool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEducationSchool) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphEducationSchool) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphEducationSchool) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphEducationSchool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphEducationSchool) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphEducationSchool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphEducationSchool) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphEducationSchool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphEducationSchool) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphEducationSchool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphEducationSchool) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetExternalSource

`func (o *MicrosoftGraphEducationSchool) GetExternalSource() AnyOfmicrosoftGraphEducationExternalSource`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *MicrosoftGraphEducationSchool) GetExternalSourceOk() (AnyOfmicrosoftGraphEducationExternalSource, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalSource

`func (o *MicrosoftGraphEducationSchool) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSource

`func (o *MicrosoftGraphEducationSchool) SetExternalSource(v AnyOfmicrosoftGraphEducationExternalSource)`

SetExternalSource gets a reference to the given AnyOfmicrosoftGraphEducationExternalSource and assigns it to the ExternalSource field.

### SetExternalSourceExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetExternalSourceExplicitNull(b bool)`

SetExternalSourceExplicitNull (un)sets ExternalSource to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalSource value is set to nil even if false is passed
### GetPrincipalEmail

`func (o *MicrosoftGraphEducationSchool) GetPrincipalEmail() string`

GetPrincipalEmail returns the PrincipalEmail field if non-nil, zero value otherwise.

### GetPrincipalEmailOk

`func (o *MicrosoftGraphEducationSchool) GetPrincipalEmailOk() (string, bool)`

GetPrincipalEmailOk returns a tuple with the PrincipalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrincipalEmail

`func (o *MicrosoftGraphEducationSchool) HasPrincipalEmail() bool`

HasPrincipalEmail returns a boolean if a field has been set.

### SetPrincipalEmail

`func (o *MicrosoftGraphEducationSchool) SetPrincipalEmail(v string)`

SetPrincipalEmail gets a reference to the given string and assigns it to the PrincipalEmail field.

### SetPrincipalEmailExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetPrincipalEmailExplicitNull(b bool)`

SetPrincipalEmailExplicitNull (un)sets PrincipalEmail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrincipalEmail value is set to nil even if false is passed
### GetPrincipalName

`func (o *MicrosoftGraphEducationSchool) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *MicrosoftGraphEducationSchool) GetPrincipalNameOk() (string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrincipalName

`func (o *MicrosoftGraphEducationSchool) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalName

`func (o *MicrosoftGraphEducationSchool) SetPrincipalName(v string)`

SetPrincipalName gets a reference to the given string and assigns it to the PrincipalName field.

### SetPrincipalNameExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetPrincipalNameExplicitNull(b bool)`

SetPrincipalNameExplicitNull (un)sets PrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrincipalName value is set to nil even if false is passed
### GetExternalPrincipalId

`func (o *MicrosoftGraphEducationSchool) GetExternalPrincipalId() string`

GetExternalPrincipalId returns the ExternalPrincipalId field if non-nil, zero value otherwise.

### GetExternalPrincipalIdOk

`func (o *MicrosoftGraphEducationSchool) GetExternalPrincipalIdOk() (string, bool)`

GetExternalPrincipalIdOk returns a tuple with the ExternalPrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalPrincipalId

`func (o *MicrosoftGraphEducationSchool) HasExternalPrincipalId() bool`

HasExternalPrincipalId returns a boolean if a field has been set.

### SetExternalPrincipalId

`func (o *MicrosoftGraphEducationSchool) SetExternalPrincipalId(v string)`

SetExternalPrincipalId gets a reference to the given string and assigns it to the ExternalPrincipalId field.

### SetExternalPrincipalIdExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetExternalPrincipalIdExplicitNull(b bool)`

SetExternalPrincipalIdExplicitNull (un)sets ExternalPrincipalId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalPrincipalId value is set to nil even if false is passed
### GetLowestGrade

`func (o *MicrosoftGraphEducationSchool) GetLowestGrade() string`

GetLowestGrade returns the LowestGrade field if non-nil, zero value otherwise.

### GetLowestGradeOk

`func (o *MicrosoftGraphEducationSchool) GetLowestGradeOk() (string, bool)`

GetLowestGradeOk returns a tuple with the LowestGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLowestGrade

`func (o *MicrosoftGraphEducationSchool) HasLowestGrade() bool`

HasLowestGrade returns a boolean if a field has been set.

### SetLowestGrade

`func (o *MicrosoftGraphEducationSchool) SetLowestGrade(v string)`

SetLowestGrade gets a reference to the given string and assigns it to the LowestGrade field.

### SetLowestGradeExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetLowestGradeExplicitNull(b bool)`

SetLowestGradeExplicitNull (un)sets LowestGrade to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LowestGrade value is set to nil even if false is passed
### GetHighestGrade

`func (o *MicrosoftGraphEducationSchool) GetHighestGrade() string`

GetHighestGrade returns the HighestGrade field if non-nil, zero value otherwise.

### GetHighestGradeOk

`func (o *MicrosoftGraphEducationSchool) GetHighestGradeOk() (string, bool)`

GetHighestGradeOk returns a tuple with the HighestGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHighestGrade

`func (o *MicrosoftGraphEducationSchool) HasHighestGrade() bool`

HasHighestGrade returns a boolean if a field has been set.

### SetHighestGrade

`func (o *MicrosoftGraphEducationSchool) SetHighestGrade(v string)`

SetHighestGrade gets a reference to the given string and assigns it to the HighestGrade field.

### SetHighestGradeExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetHighestGradeExplicitNull(b bool)`

SetHighestGradeExplicitNull (un)sets HighestGrade to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HighestGrade value is set to nil even if false is passed
### GetSchoolNumber

`func (o *MicrosoftGraphEducationSchool) GetSchoolNumber() string`

GetSchoolNumber returns the SchoolNumber field if non-nil, zero value otherwise.

### GetSchoolNumberOk

`func (o *MicrosoftGraphEducationSchool) GetSchoolNumberOk() (string, bool)`

GetSchoolNumberOk returns a tuple with the SchoolNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchoolNumber

`func (o *MicrosoftGraphEducationSchool) HasSchoolNumber() bool`

HasSchoolNumber returns a boolean if a field has been set.

### SetSchoolNumber

`func (o *MicrosoftGraphEducationSchool) SetSchoolNumber(v string)`

SetSchoolNumber gets a reference to the given string and assigns it to the SchoolNumber field.

### SetSchoolNumberExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetSchoolNumberExplicitNull(b bool)`

SetSchoolNumberExplicitNull (un)sets SchoolNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SchoolNumber value is set to nil even if false is passed
### GetExternalId

`func (o *MicrosoftGraphEducationSchool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *MicrosoftGraphEducationSchool) GetExternalIdOk() (string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalId

`func (o *MicrosoftGraphEducationSchool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalId

`func (o *MicrosoftGraphEducationSchool) SetExternalId(v string)`

SetExternalId gets a reference to the given string and assigns it to the ExternalId field.

### SetExternalIdExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetExternalIdExplicitNull(b bool)`

SetExternalIdExplicitNull (un)sets ExternalId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalId value is set to nil even if false is passed
### GetPhone

`func (o *MicrosoftGraphEducationSchool) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MicrosoftGraphEducationSchool) GetPhoneOk() (string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhone

`func (o *MicrosoftGraphEducationSchool) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhone

`func (o *MicrosoftGraphEducationSchool) SetPhone(v string)`

SetPhone gets a reference to the given string and assigns it to the Phone field.

### SetPhoneExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetPhoneExplicitNull(b bool)`

SetPhoneExplicitNull (un)sets Phone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Phone value is set to nil even if false is passed
### GetFax

`func (o *MicrosoftGraphEducationSchool) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *MicrosoftGraphEducationSchool) GetFaxOk() (string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFax

`func (o *MicrosoftGraphEducationSchool) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFax

`func (o *MicrosoftGraphEducationSchool) SetFax(v string)`

SetFax gets a reference to the given string and assigns it to the Fax field.

### SetFaxExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetFaxExplicitNull(b bool)`

SetFaxExplicitNull (un)sets Fax to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Fax value is set to nil even if false is passed
### GetCreatedBy

`func (o *MicrosoftGraphEducationSchool) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphEducationSchool) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphEducationSchool) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphEducationSchool) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetAddress

`func (o *MicrosoftGraphEducationSchool) GetAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphEducationSchool) GetAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddress

`func (o *MicrosoftGraphEducationSchool) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddress

`func (o *MicrosoftGraphEducationSchool) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the Address field.

### SetAddressExplicitNull

`func (o *MicrosoftGraphEducationSchool) SetAddressExplicitNull(b bool)`

SetAddressExplicitNull (un)sets Address to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Address value is set to nil even if false is passed
### GetClasses

`func (o *MicrosoftGraphEducationSchool) GetClasses() []MicrosoftGraphEducationClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *MicrosoftGraphEducationSchool) GetClassesOk() ([]MicrosoftGraphEducationClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClasses

`func (o *MicrosoftGraphEducationSchool) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### SetClasses

`func (o *MicrosoftGraphEducationSchool) SetClasses(v []MicrosoftGraphEducationClass)`

SetClasses gets a reference to the given []MicrosoftGraphEducationClass and assigns it to the Classes field.

### GetUsers

`func (o *MicrosoftGraphEducationSchool) GetUsers() []MicrosoftGraphEducationUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *MicrosoftGraphEducationSchool) GetUsersOk() ([]MicrosoftGraphEducationUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsers

`func (o *MicrosoftGraphEducationSchool) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsers

`func (o *MicrosoftGraphEducationSchool) SetUsers(v []MicrosoftGraphEducationUser)`

SetUsers gets a reference to the given []MicrosoftGraphEducationUser and assigns it to the Users field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


