# EducationOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalSource** | Pointer to [**AnyOfmicrosoftGraphEducationExternalSource**](anyOf&lt;microsoft.graph.educationExternalSource&gt;.md) |  | [optional] 

## Methods

### GetDisplayName

`func (o *EducationOrganization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EducationOrganization) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *EducationOrganization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *EducationOrganization) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *EducationOrganization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EducationOrganization) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *EducationOrganization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *EducationOrganization) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *EducationOrganization) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetExternalSource

`func (o *EducationOrganization) GetExternalSource() AnyOfmicrosoftGraphEducationExternalSource`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *EducationOrganization) GetExternalSourceOk() (AnyOfmicrosoftGraphEducationExternalSource, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalSource

`func (o *EducationOrganization) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSource

`func (o *EducationOrganization) SetExternalSource(v AnyOfmicrosoftGraphEducationExternalSource)`

SetExternalSource gets a reference to the given AnyOfmicrosoftGraphEducationExternalSource and assigns it to the ExternalSource field.

### SetExternalSourceExplicitNull

`func (o *EducationOrganization) SetExternalSourceExplicitNull(b bool)`

SetExternalSourceExplicitNull (un)sets ExternalSource to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalSource value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


