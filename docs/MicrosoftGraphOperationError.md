# MicrosoftGraphOperationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### GetCode

`func (o *MicrosoftGraphOperationError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MicrosoftGraphOperationError) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *MicrosoftGraphOperationError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *MicrosoftGraphOperationError) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### SetCodeExplicitNull

`func (o *MicrosoftGraphOperationError) SetCodeExplicitNull(b bool)`

SetCodeExplicitNull (un)sets Code to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Code value is set to nil even if false is passed
### GetMessage

`func (o *MicrosoftGraphOperationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MicrosoftGraphOperationError) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *MicrosoftGraphOperationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *MicrosoftGraphOperationError) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### SetMessageExplicitNull

`func (o *MicrosoftGraphOperationError) SetMessageExplicitNull(b bool)`

SetMessageExplicitNull (un)sets Message to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Message value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


