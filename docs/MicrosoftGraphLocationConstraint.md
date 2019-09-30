# MicrosoftGraphLocationConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locations** | Pointer to [**[]AnyOfmicrosoftGraphLocationConstraintItem**](anyOf&lt;microsoft.graph.locationConstraintItem&gt;.md) |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 
**SuggestLocation** | Pointer to **bool** |  | [optional] 

## Methods

### GetLocations

`func (o *MicrosoftGraphLocationConstraint) GetLocations() []AnyOfmicrosoftGraphLocationConstraintItem`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *MicrosoftGraphLocationConstraint) GetLocationsOk() ([]AnyOfmicrosoftGraphLocationConstraintItem, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocations

`func (o *MicrosoftGraphLocationConstraint) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocations

`func (o *MicrosoftGraphLocationConstraint) SetLocations(v []AnyOfmicrosoftGraphLocationConstraintItem)`

SetLocations gets a reference to the given []AnyOfmicrosoftGraphLocationConstraintItem and assigns it to the Locations field.

### GetIsRequired

`func (o *MicrosoftGraphLocationConstraint) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *MicrosoftGraphLocationConstraint) GetIsRequiredOk() (bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsRequired

`func (o *MicrosoftGraphLocationConstraint) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### SetIsRequired

`func (o *MicrosoftGraphLocationConstraint) SetIsRequired(v bool)`

SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.

### SetIsRequiredExplicitNull

`func (o *MicrosoftGraphLocationConstraint) SetIsRequiredExplicitNull(b bool)`

SetIsRequiredExplicitNull (un)sets IsRequired to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsRequired value is set to nil even if false is passed
### GetSuggestLocation

`func (o *MicrosoftGraphLocationConstraint) GetSuggestLocation() bool`

GetSuggestLocation returns the SuggestLocation field if non-nil, zero value otherwise.

### GetSuggestLocationOk

`func (o *MicrosoftGraphLocationConstraint) GetSuggestLocationOk() (bool, bool)`

GetSuggestLocationOk returns a tuple with the SuggestLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSuggestLocation

`func (o *MicrosoftGraphLocationConstraint) HasSuggestLocation() bool`

HasSuggestLocation returns a boolean if a field has been set.

### SetSuggestLocation

`func (o *MicrosoftGraphLocationConstraint) SetSuggestLocation(v bool)`

SetSuggestLocation gets a reference to the given bool and assigns it to the SuggestLocation field.

### SetSuggestLocationExplicitNull

`func (o *MicrosoftGraphLocationConstraint) SetSuggestLocationExplicitNull(b bool)`

SetSuggestLocationExplicitNull (un)sets SuggestLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SuggestLocation value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


