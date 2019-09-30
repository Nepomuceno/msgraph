# MicrosoftGraphRemoteLockActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionName** | Pointer to **string** | Action name | [optional] 
**ActionState** | Pointer to [**AnyOfmicrosoftGraphActionState**](anyOf&lt;microsoft.graph.actionState&gt;.md) | State of the action | [optional] 
**StartDateTime** | Pointer to [**time.Time**](time.Time.md) | Time the action was initiated | [optional] 
**LastUpdatedDateTime** | Pointer to [**time.Time**](time.Time.md) | Time the action state was last updated | [optional] 
**UnlockPin** | Pointer to **string** | Pin to unlock the client | [optional] 

## Methods

### GetActionName

`func (o *MicrosoftGraphRemoteLockActionResult) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *MicrosoftGraphRemoteLockActionResult) GetActionNameOk() (string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionName

`func (o *MicrosoftGraphRemoteLockActionResult) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### SetActionName

`func (o *MicrosoftGraphRemoteLockActionResult) SetActionName(v string)`

SetActionName gets a reference to the given string and assigns it to the ActionName field.

### SetActionNameExplicitNull

`func (o *MicrosoftGraphRemoteLockActionResult) SetActionNameExplicitNull(b bool)`

SetActionNameExplicitNull (un)sets ActionName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActionName value is set to nil even if false is passed
### GetActionState

`func (o *MicrosoftGraphRemoteLockActionResult) GetActionState() AnyOfmicrosoftGraphActionState`

GetActionState returns the ActionState field if non-nil, zero value otherwise.

### GetActionStateOk

`func (o *MicrosoftGraphRemoteLockActionResult) GetActionStateOk() (AnyOfmicrosoftGraphActionState, bool)`

GetActionStateOk returns a tuple with the ActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionState

`func (o *MicrosoftGraphRemoteLockActionResult) HasActionState() bool`

HasActionState returns a boolean if a field has been set.

### SetActionState

`func (o *MicrosoftGraphRemoteLockActionResult) SetActionState(v AnyOfmicrosoftGraphActionState)`

SetActionState gets a reference to the given AnyOfmicrosoftGraphActionState and assigns it to the ActionState field.

### GetStartDateTime

`func (o *MicrosoftGraphRemoteLockActionResult) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphRemoteLockActionResult) GetStartDateTimeOk() (time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDateTime

`func (o *MicrosoftGraphRemoteLockActionResult) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTime

`func (o *MicrosoftGraphRemoteLockActionResult) SetStartDateTime(v time.Time)`

SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.

### GetLastUpdatedDateTime

`func (o *MicrosoftGraphRemoteLockActionResult) GetLastUpdatedDateTime() time.Time`

GetLastUpdatedDateTime returns the LastUpdatedDateTime field if non-nil, zero value otherwise.

### GetLastUpdatedDateTimeOk

`func (o *MicrosoftGraphRemoteLockActionResult) GetLastUpdatedDateTimeOk() (time.Time, bool)`

GetLastUpdatedDateTimeOk returns a tuple with the LastUpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdatedDateTime

`func (o *MicrosoftGraphRemoteLockActionResult) HasLastUpdatedDateTime() bool`

HasLastUpdatedDateTime returns a boolean if a field has been set.

### SetLastUpdatedDateTime

`func (o *MicrosoftGraphRemoteLockActionResult) SetLastUpdatedDateTime(v time.Time)`

SetLastUpdatedDateTime gets a reference to the given time.Time and assigns it to the LastUpdatedDateTime field.

### GetUnlockPin

`func (o *MicrosoftGraphRemoteLockActionResult) GetUnlockPin() string`

GetUnlockPin returns the UnlockPin field if non-nil, zero value otherwise.

### GetUnlockPinOk

`func (o *MicrosoftGraphRemoteLockActionResult) GetUnlockPinOk() (string, bool)`

GetUnlockPinOk returns a tuple with the UnlockPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnlockPin

`func (o *MicrosoftGraphRemoteLockActionResult) HasUnlockPin() bool`

HasUnlockPin returns a boolean if a field has been set.

### SetUnlockPin

`func (o *MicrosoftGraphRemoteLockActionResult) SetUnlockPin(v string)`

SetUnlockPin gets a reference to the given string and assigns it to the UnlockPin field.

### SetUnlockPinExplicitNull

`func (o *MicrosoftGraphRemoteLockActionResult) SetUnlockPinExplicitNull(b bool)`

SetUnlockPinExplicitNull (un)sets UnlockPin to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UnlockPin value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


