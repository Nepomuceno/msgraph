# MicrosoftGraphEducationOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalSource** | Pointer to [**AnyOfmicrosoftGraphEducationExternalSource**](anyOf&lt;microsoft.graph.educationExternalSource&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphEducationOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEducationOrganization) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphEducationOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphEducationOrganization) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphEducationOrganization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphEducationOrganization) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphEducationOrganization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphEducationOrganization) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphEducationOrganization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphEducationOrganization) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphEducationOrganization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphEducationOrganization) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphEducationOrganization) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetExternalSource

`func (o *MicrosoftGraphEducationOrganization) GetExternalSource() AnyOfmicrosoftGraphEducationExternalSource`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *MicrosoftGraphEducationOrganization) GetExternalSourceOk() (AnyOfmicrosoftGraphEducationExternalSource, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalSource

`func (o *MicrosoftGraphEducationOrganization) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSource

`func (o *MicrosoftGraphEducationOrganization) SetExternalSource(v AnyOfmicrosoftGraphEducationExternalSource)`

SetExternalSource gets a reference to the given AnyOfmicrosoftGraphEducationExternalSource and assigns it to the ExternalSource field.

### SetExternalSourceExplicitNull

`func (o *MicrosoftGraphEducationOrganization) SetExternalSourceExplicitNull(b bool)`

SetExternalSourceExplicitNull (un)sets ExternalSource to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalSource value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


