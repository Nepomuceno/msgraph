# MicrosoftGraphEducationStudent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraduationYear** | Pointer to **string** |  | [optional] 
**Grade** | Pointer to **string** |  | [optional] 
**BirthDate** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to [**AnyOfmicrosoftGraphEducationGender**](anyOf&lt;microsoft.graph.educationGender&gt;.md) |  | [optional] 
**StudentNumber** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 

## Methods

### GetGraduationYear

`func (o *MicrosoftGraphEducationStudent) GetGraduationYear() string`

GetGraduationYear returns the GraduationYear field if non-nil, zero value otherwise.

### GetGraduationYearOk

`func (o *MicrosoftGraphEducationStudent) GetGraduationYearOk() (string, bool)`

GetGraduationYearOk returns a tuple with the GraduationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGraduationYear

`func (o *MicrosoftGraphEducationStudent) HasGraduationYear() bool`

HasGraduationYear returns a boolean if a field has been set.

### SetGraduationYear

`func (o *MicrosoftGraphEducationStudent) SetGraduationYear(v string)`

SetGraduationYear gets a reference to the given string and assigns it to the GraduationYear field.

### SetGraduationYearExplicitNull

`func (o *MicrosoftGraphEducationStudent) SetGraduationYearExplicitNull(b bool)`

SetGraduationYearExplicitNull (un)sets GraduationYear to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GraduationYear value is set to nil even if false is passed
### GetGrade

`func (o *MicrosoftGraphEducationStudent) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *MicrosoftGraphEducationStudent) GetGradeOk() (string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGrade

`func (o *MicrosoftGraphEducationStudent) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### SetGrade

`func (o *MicrosoftGraphEducationStudent) SetGrade(v string)`

SetGrade gets a reference to the given string and assigns it to the Grade field.

### SetGradeExplicitNull

`func (o *MicrosoftGraphEducationStudent) SetGradeExplicitNull(b bool)`

SetGradeExplicitNull (un)sets Grade to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Grade value is set to nil even if false is passed
### GetBirthDate

`func (o *MicrosoftGraphEducationStudent) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *MicrosoftGraphEducationStudent) GetBirthDateOk() (string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBirthDate

`func (o *MicrosoftGraphEducationStudent) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDate

`func (o *MicrosoftGraphEducationStudent) SetBirthDate(v string)`

SetBirthDate gets a reference to the given string and assigns it to the BirthDate field.

### SetBirthDateExplicitNull

`func (o *MicrosoftGraphEducationStudent) SetBirthDateExplicitNull(b bool)`

SetBirthDateExplicitNull (un)sets BirthDate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BirthDate value is set to nil even if false is passed
### GetGender

`func (o *MicrosoftGraphEducationStudent) GetGender() AnyOfmicrosoftGraphEducationGender`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *MicrosoftGraphEducationStudent) GetGenderOk() (AnyOfmicrosoftGraphEducationGender, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGender

`func (o *MicrosoftGraphEducationStudent) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGender

`func (o *MicrosoftGraphEducationStudent) SetGender(v AnyOfmicrosoftGraphEducationGender)`

SetGender gets a reference to the given AnyOfmicrosoftGraphEducationGender and assigns it to the Gender field.

### SetGenderExplicitNull

`func (o *MicrosoftGraphEducationStudent) SetGenderExplicitNull(b bool)`

SetGenderExplicitNull (un)sets Gender to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Gender value is set to nil even if false is passed
### GetStudentNumber

`func (o *MicrosoftGraphEducationStudent) GetStudentNumber() string`

GetStudentNumber returns the StudentNumber field if non-nil, zero value otherwise.

### GetStudentNumberOk

`func (o *MicrosoftGraphEducationStudent) GetStudentNumberOk() (string, bool)`

GetStudentNumberOk returns a tuple with the StudentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStudentNumber

`func (o *MicrosoftGraphEducationStudent) HasStudentNumber() bool`

HasStudentNumber returns a boolean if a field has been set.

### SetStudentNumber

`func (o *MicrosoftGraphEducationStudent) SetStudentNumber(v string)`

SetStudentNumber gets a reference to the given string and assigns it to the StudentNumber field.

### SetStudentNumberExplicitNull

`func (o *MicrosoftGraphEducationStudent) SetStudentNumberExplicitNull(b bool)`

SetStudentNumberExplicitNull (un)sets StudentNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StudentNumber value is set to nil even if false is passed
### GetExternalId

`func (o *MicrosoftGraphEducationStudent) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *MicrosoftGraphEducationStudent) GetExternalIdOk() (string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalId

`func (o *MicrosoftGraphEducationStudent) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalId

`func (o *MicrosoftGraphEducationStudent) SetExternalId(v string)`

SetExternalId gets a reference to the given string and assigns it to the ExternalId field.

### SetExternalIdExplicitNull

`func (o *MicrosoftGraphEducationStudent) SetExternalIdExplicitNull(b bool)`

SetExternalIdExplicitNull (un)sets ExternalId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


