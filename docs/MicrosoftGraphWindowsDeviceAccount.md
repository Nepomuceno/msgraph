# MicrosoftGraphWindowsDeviceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 

## Methods

### GetPassword

`func (o *MicrosoftGraphWindowsDeviceAccount) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MicrosoftGraphWindowsDeviceAccount) GetPasswordOk() (string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPassword

`func (o *MicrosoftGraphWindowsDeviceAccount) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPassword

`func (o *MicrosoftGraphWindowsDeviceAccount) SetPassword(v string)`

SetPassword gets a reference to the given string and assigns it to the Password field.

### SetPasswordExplicitNull

`func (o *MicrosoftGraphWindowsDeviceAccount) SetPasswordExplicitNull(b bool)`

SetPasswordExplicitNull (un)sets Password to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Password value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


