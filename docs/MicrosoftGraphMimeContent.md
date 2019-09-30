# MicrosoftGraphMimeContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Indicates the content mime type. | [optional] 
**Value** | Pointer to **string** | The byte array that contains the actual content. | [optional] 

## Methods

### GetType

`func (o *MicrosoftGraphMimeContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphMimeContent) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *MicrosoftGraphMimeContent) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *MicrosoftGraphMimeContent) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### SetTypeExplicitNull

`func (o *MicrosoftGraphMimeContent) SetTypeExplicitNull(b bool)`

SetTypeExplicitNull (un)sets Type to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Type value is set to nil even if false is passed
### GetValue

`func (o *MicrosoftGraphMimeContent) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphMimeContent) GetValueOk() (string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *MicrosoftGraphMimeContent) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *MicrosoftGraphMimeContent) SetValue(v string)`

SetValue gets a reference to the given string and assigns it to the Value field.

### SetValueExplicitNull

`func (o *MicrosoftGraphMimeContent) SetValueExplicitNull(b bool)`

SetValueExplicitNull (un)sets Value to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Value value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


