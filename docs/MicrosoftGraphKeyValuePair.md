# MicrosoftGraphKeyValuePair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for this key-value pair | [optional] 
**Value** | Pointer to **string** | Value for this key-value pair | [optional] 

## Methods

### GetName

`func (o *MicrosoftGraphKeyValuePair) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphKeyValuePair) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphKeyValuePair) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphKeyValuePair) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetValue

`func (o *MicrosoftGraphKeyValuePair) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphKeyValuePair) GetValueOk() (string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *MicrosoftGraphKeyValuePair) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *MicrosoftGraphKeyValuePair) SetValue(v string)`

SetValue gets a reference to the given string and assigns it to the Value field.

### SetValueExplicitNull

`func (o *MicrosoftGraphKeyValuePair) SetValueExplicitNull(b bool)`

SetValueExplicitNull (un)sets Value to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Value value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


