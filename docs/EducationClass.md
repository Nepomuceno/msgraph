# EducationClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**MailNickname** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**ClassCode** | Pointer to **string** |  | [optional] 
**ExternalName** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalSource** | Pointer to [**AnyOfmicrosoftGraphEducationExternalSource**](anyOf&lt;microsoft.graph.educationExternalSource&gt;.md) |  | [optional] 
**Term** | Pointer to [**AnyOfmicrosoftGraphEducationTerm**](anyOf&lt;microsoft.graph.educationTerm&gt;.md) |  | [optional] 
**Schools** | Pointer to [**[]MicrosoftGraphEducationSchool**](microsoft.graph.educationSchool.md) |  | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphEducationUser**](microsoft.graph.educationUser.md) |  | [optional] 
**Teachers** | Pointer to [**[]MicrosoftGraphEducationUser**](microsoft.graph.educationUser.md) |  | [optional] 
**Group** | Pointer to [**AnyOfmicrosoftGraphGroup**](anyOf&lt;microsoft.graph.group&gt;.md) |  | [optional] 

## Methods

### GetDisplayName

`func (o *EducationClass) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EducationClass) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *EducationClass) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *EducationClass) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetMailNickname

`func (o *EducationClass) GetMailNickname() string`

GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.

### GetMailNicknameOk

`func (o *EducationClass) GetMailNicknameOk() (string, bool)`

GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailNickname

`func (o *EducationClass) HasMailNickname() bool`

HasMailNickname returns a boolean if a field has been set.

### SetMailNickname

`func (o *EducationClass) SetMailNickname(v string)`

SetMailNickname gets a reference to the given string and assigns it to the MailNickname field.

### GetDescription

`func (o *EducationClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EducationClass) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *EducationClass) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *EducationClass) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *EducationClass) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetCreatedBy

`func (o *EducationClass) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EducationClass) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *EducationClass) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *EducationClass) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *EducationClass) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetClassCode

`func (o *EducationClass) GetClassCode() string`

GetClassCode returns the ClassCode field if non-nil, zero value otherwise.

### GetClassCodeOk

`func (o *EducationClass) GetClassCodeOk() (string, bool)`

GetClassCodeOk returns a tuple with the ClassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClassCode

`func (o *EducationClass) HasClassCode() bool`

HasClassCode returns a boolean if a field has been set.

### SetClassCode

`func (o *EducationClass) SetClassCode(v string)`

SetClassCode gets a reference to the given string and assigns it to the ClassCode field.

### SetClassCodeExplicitNull

`func (o *EducationClass) SetClassCodeExplicitNull(b bool)`

SetClassCodeExplicitNull (un)sets ClassCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ClassCode value is set to nil even if false is passed
### GetExternalName

`func (o *EducationClass) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *EducationClass) GetExternalNameOk() (string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalName

`func (o *EducationClass) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### SetExternalName

`func (o *EducationClass) SetExternalName(v string)`

SetExternalName gets a reference to the given string and assigns it to the ExternalName field.

### SetExternalNameExplicitNull

`func (o *EducationClass) SetExternalNameExplicitNull(b bool)`

SetExternalNameExplicitNull (un)sets ExternalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalName value is set to nil even if false is passed
### GetExternalId

`func (o *EducationClass) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EducationClass) GetExternalIdOk() (string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalId

`func (o *EducationClass) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalId

`func (o *EducationClass) SetExternalId(v string)`

SetExternalId gets a reference to the given string and assigns it to the ExternalId field.

### SetExternalIdExplicitNull

`func (o *EducationClass) SetExternalIdExplicitNull(b bool)`

SetExternalIdExplicitNull (un)sets ExternalId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalId value is set to nil even if false is passed
### GetExternalSource

`func (o *EducationClass) GetExternalSource() AnyOfmicrosoftGraphEducationExternalSource`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *EducationClass) GetExternalSourceOk() (AnyOfmicrosoftGraphEducationExternalSource, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalSource

`func (o *EducationClass) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSource

`func (o *EducationClass) SetExternalSource(v AnyOfmicrosoftGraphEducationExternalSource)`

SetExternalSource gets a reference to the given AnyOfmicrosoftGraphEducationExternalSource and assigns it to the ExternalSource field.

### SetExternalSourceExplicitNull

`func (o *EducationClass) SetExternalSourceExplicitNull(b bool)`

SetExternalSourceExplicitNull (un)sets ExternalSource to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalSource value is set to nil even if false is passed
### GetTerm

`func (o *EducationClass) GetTerm() AnyOfmicrosoftGraphEducationTerm`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *EducationClass) GetTermOk() (AnyOfmicrosoftGraphEducationTerm, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTerm

`func (o *EducationClass) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### SetTerm

`func (o *EducationClass) SetTerm(v AnyOfmicrosoftGraphEducationTerm)`

SetTerm gets a reference to the given AnyOfmicrosoftGraphEducationTerm and assigns it to the Term field.

### SetTermExplicitNull

`func (o *EducationClass) SetTermExplicitNull(b bool)`

SetTermExplicitNull (un)sets Term to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Term value is set to nil even if false is passed
### GetSchools

`func (o *EducationClass) GetSchools() []MicrosoftGraphEducationSchool`

GetSchools returns the Schools field if non-nil, zero value otherwise.

### GetSchoolsOk

`func (o *EducationClass) GetSchoolsOk() ([]MicrosoftGraphEducationSchool, bool)`

GetSchoolsOk returns a tuple with the Schools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchools

`func (o *EducationClass) HasSchools() bool`

HasSchools returns a boolean if a field has been set.

### SetSchools

`func (o *EducationClass) SetSchools(v []MicrosoftGraphEducationSchool)`

SetSchools gets a reference to the given []MicrosoftGraphEducationSchool and assigns it to the Schools field.

### GetMembers

`func (o *EducationClass) GetMembers() []MicrosoftGraphEducationUser`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *EducationClass) GetMembersOk() ([]MicrosoftGraphEducationUser, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *EducationClass) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *EducationClass) SetMembers(v []MicrosoftGraphEducationUser)`

SetMembers gets a reference to the given []MicrosoftGraphEducationUser and assigns it to the Members field.

### GetTeachers

`func (o *EducationClass) GetTeachers() []MicrosoftGraphEducationUser`

GetTeachers returns the Teachers field if non-nil, zero value otherwise.

### GetTeachersOk

`func (o *EducationClass) GetTeachersOk() ([]MicrosoftGraphEducationUser, bool)`

GetTeachersOk returns a tuple with the Teachers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeachers

`func (o *EducationClass) HasTeachers() bool`

HasTeachers returns a boolean if a field has been set.

### SetTeachers

`func (o *EducationClass) SetTeachers(v []MicrosoftGraphEducationUser)`

SetTeachers gets a reference to the given []MicrosoftGraphEducationUser and assigns it to the Teachers field.

### GetGroup

`func (o *EducationClass) GetGroup() AnyOfmicrosoftGraphGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *EducationClass) GetGroupOk() (AnyOfmicrosoftGraphGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroup

`func (o *EducationClass) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroup

`func (o *EducationClass) SetGroup(v AnyOfmicrosoftGraphGroup)`

SetGroup gets a reference to the given AnyOfmicrosoftGraphGroup and assigns it to the Group field.

### SetGroupExplicitNull

`func (o *EducationClass) SetGroupExplicitNull(b bool)`

SetGroupExplicitNull (un)sets Group to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Group value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


