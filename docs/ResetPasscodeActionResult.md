# ResetPasscodeActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passcode** | Pointer to **string** | Newly generated passcode for the device  | [optional] 

## Methods

### GetPasscode

`func (o *ResetPasscodeActionResult) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *ResetPasscodeActionResult) GetPasscodeOk() (string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscode

`func (o *ResetPasscodeActionResult) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.

### SetPasscode

`func (o *ResetPasscodeActionResult) SetPasscode(v string)`

SetPasscode gets a reference to the given string and assigns it to the Passcode field.

### SetPasscodeExplicitNull

`func (o *ResetPasscodeActionResult) SetPasscodeExplicitNull(b bool)`

SetPasscodeExplicitNull (un)sets Passcode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Passcode value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


