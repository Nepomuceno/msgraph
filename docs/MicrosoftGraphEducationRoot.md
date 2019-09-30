# MicrosoftGraphEducationRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Classes** | Pointer to [**[]MicrosoftGraphEducationClass**](microsoft.graph.educationClass.md) |  | [optional] 
**Schools** | Pointer to [**[]MicrosoftGraphEducationSchool**](microsoft.graph.educationSchool.md) |  | [optional] 
**Users** | Pointer to [**[]MicrosoftGraphEducationUser**](microsoft.graph.educationUser.md) |  | [optional] 
**Me** | Pointer to [**AnyOfmicrosoftGraphEducationUser**](anyOf&lt;microsoft.graph.educationUser&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphEducationRoot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEducationRoot) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphEducationRoot) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphEducationRoot) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetClasses

`func (o *MicrosoftGraphEducationRoot) GetClasses() []MicrosoftGraphEducationClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *MicrosoftGraphEducationRoot) GetClassesOk() ([]MicrosoftGraphEducationClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClasses

`func (o *MicrosoftGraphEducationRoot) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### SetClasses

`func (o *MicrosoftGraphEducationRoot) SetClasses(v []MicrosoftGraphEducationClass)`

SetClasses gets a reference to the given []MicrosoftGraphEducationClass and assigns it to the Classes field.

### GetSchools

`func (o *MicrosoftGraphEducationRoot) GetSchools() []MicrosoftGraphEducationSchool`

GetSchools returns the Schools field if non-nil, zero value otherwise.

### GetSchoolsOk

`func (o *MicrosoftGraphEducationRoot) GetSchoolsOk() ([]MicrosoftGraphEducationSchool, bool)`

GetSchoolsOk returns a tuple with the Schools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchools

`func (o *MicrosoftGraphEducationRoot) HasSchools() bool`

HasSchools returns a boolean if a field has been set.

### SetSchools

`func (o *MicrosoftGraphEducationRoot) SetSchools(v []MicrosoftGraphEducationSchool)`

SetSchools gets a reference to the given []MicrosoftGraphEducationSchool and assigns it to the Schools field.

### GetUsers

`func (o *MicrosoftGraphEducationRoot) GetUsers() []MicrosoftGraphEducationUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *MicrosoftGraphEducationRoot) GetUsersOk() ([]MicrosoftGraphEducationUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsers

`func (o *MicrosoftGraphEducationRoot) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsers

`func (o *MicrosoftGraphEducationRoot) SetUsers(v []MicrosoftGraphEducationUser)`

SetUsers gets a reference to the given []MicrosoftGraphEducationUser and assigns it to the Users field.

### GetMe

`func (o *MicrosoftGraphEducationRoot) GetMe() AnyOfmicrosoftGraphEducationUser`

GetMe returns the Me field if non-nil, zero value otherwise.

### GetMeOk

`func (o *MicrosoftGraphEducationRoot) GetMeOk() (AnyOfmicrosoftGraphEducationUser, bool)`

GetMeOk returns a tuple with the Me field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMe

`func (o *MicrosoftGraphEducationRoot) HasMe() bool`

HasMe returns a boolean if a field has been set.

### SetMe

`func (o *MicrosoftGraphEducationRoot) SetMe(v AnyOfmicrosoftGraphEducationUser)`

SetMe gets a reference to the given AnyOfmicrosoftGraphEducationUser and assigns it to the Me field.

### SetMeExplicitNull

`func (o *MicrosoftGraphEducationRoot) SetMeExplicitNull(b bool)`

SetMeExplicitNull (un)sets Me to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Me value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


