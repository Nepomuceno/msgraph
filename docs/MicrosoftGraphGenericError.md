# MicrosoftGraphGenericError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### GetMessage

`func (o *MicrosoftGraphGenericError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MicrosoftGraphGenericError) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *MicrosoftGraphGenericError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *MicrosoftGraphGenericError) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### SetMessageExplicitNull

`func (o *MicrosoftGraphGenericError) SetMessageExplicitNull(b bool)`

SetMessageExplicitNull (un)sets Message to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Message value is set to nil even if false is passed
### GetCode

`func (o *MicrosoftGraphGenericError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MicrosoftGraphGenericError) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *MicrosoftGraphGenericError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *MicrosoftGraphGenericError) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### SetCodeExplicitNull

`func (o *MicrosoftGraphGenericError) SetCodeExplicitNull(b bool)`

SetCodeExplicitNull (un)sets Code to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Code value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


