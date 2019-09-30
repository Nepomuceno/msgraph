# RemoteLockActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnlockPin** | Pointer to **string** | Pin to unlock the client | [optional] 

## Methods

### GetUnlockPin

`func (o *RemoteLockActionResult) GetUnlockPin() string`

GetUnlockPin returns the UnlockPin field if non-nil, zero value otherwise.

### GetUnlockPinOk

`func (o *RemoteLockActionResult) GetUnlockPinOk() (string, bool)`

GetUnlockPinOk returns a tuple with the UnlockPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnlockPin

`func (o *RemoteLockActionResult) HasUnlockPin() bool`

HasUnlockPin returns a boolean if a field has been set.

### SetUnlockPin

`func (o *RemoteLockActionResult) SetUnlockPin(v string)`

SetUnlockPin gets a reference to the given string and assigns it to the UnlockPin field.

### SetUnlockPinExplicitNull

`func (o *RemoteLockActionResult) SetUnlockPinExplicitNull(b bool)`

SetUnlockPinExplicitNull (un)sets UnlockPin to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UnlockPin value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


