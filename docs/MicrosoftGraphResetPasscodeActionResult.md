# MicrosoftGraphResetPasscodeActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionName** | Pointer to **string** | Action name | [optional] 
**ActionState** | Pointer to [**AnyOfmicrosoftGraphActionState**](anyOf&lt;microsoft.graph.actionState&gt;.md) | State of the action | [optional] 
**StartDateTime** | Pointer to [**time.Time**](time.Time.md) | Time the action was initiated | [optional] 
**LastUpdatedDateTime** | Pointer to [**time.Time**](time.Time.md) | Time the action state was last updated | [optional] 
**Passcode** | Pointer to **string** | Newly generated passcode for the device  | [optional] 

## Methods

### GetActionName

`func (o *MicrosoftGraphResetPasscodeActionResult) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *MicrosoftGraphResetPasscodeActionResult) GetActionNameOk() (string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionName

`func (o *MicrosoftGraphResetPasscodeActionResult) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### SetActionName

`func (o *MicrosoftGraphResetPasscodeActionResult) SetActionName(v string)`

SetActionName gets a reference to the given string and assigns it to the ActionName field.

### SetActionNameExplicitNull

`func (o *MicrosoftGraphResetPasscodeActionResult) SetActionNameExplicitNull(b bool)`

SetActionNameExplicitNull (un)sets ActionName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActionName value is set to nil even if false is passed
### GetActionState

`func (o *MicrosoftGraphResetPasscodeActionResult) GetActionState() AnyOfmicrosoftGraphActionState`

GetActionState returns the ActionState field if non-nil, zero value otherwise.

### GetActionStateOk

`func (o *MicrosoftGraphResetPasscodeActionResult) GetActionStateOk() (AnyOfmicrosoftGraphActionState, bool)`

GetActionStateOk returns a tuple with the ActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionState

`func (o *MicrosoftGraphResetPasscodeActionResult) HasActionState() bool`

HasActionState returns a boolean if a field has been set.

### SetActionState

`func (o *MicrosoftGraphResetPasscodeActionResult) SetActionState(v AnyOfmicrosoftGraphActionState)`

SetActionState gets a reference to the given AnyOfmicrosoftGraphActionState and assigns it to the ActionState field.

### GetStartDateTime

`func (o *MicrosoftGraphResetPasscodeActionResult) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphResetPasscodeActionResult) GetStartDateTimeOk() (time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDateTime

`func (o *MicrosoftGraphResetPasscodeActionResult) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTime

`func (o *MicrosoftGraphResetPasscodeActionResult) SetStartDateTime(v time.Time)`

SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.

### GetLastUpdatedDateTime

`func (o *MicrosoftGraphResetPasscodeActionResult) GetLastUpdatedDateTime() time.Time`

GetLastUpdatedDateTime returns the LastUpdatedDateTime field if non-nil, zero value otherwise.

### GetLastUpdatedDateTimeOk

`func (o *MicrosoftGraphResetPasscodeActionResult) GetLastUpdatedDateTimeOk() (time.Time, bool)`

GetLastUpdatedDateTimeOk returns a tuple with the LastUpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdatedDateTime

`func (o *MicrosoftGraphResetPasscodeActionResult) HasLastUpdatedDateTime() bool`

HasLastUpdatedDateTime returns a boolean if a field has been set.

### SetLastUpdatedDateTime

`func (o *MicrosoftGraphResetPasscodeActionResult) SetLastUpdatedDateTime(v time.Time)`

SetLastUpdatedDateTime gets a reference to the given time.Time and assigns it to the LastUpdatedDateTime field.

### GetPasscode

`func (o *MicrosoftGraphResetPasscodeActionResult) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *MicrosoftGraphResetPasscodeActionResult) GetPasscodeOk() (string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasscode

`func (o *MicrosoftGraphResetPasscodeActionResult) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.

### SetPasscode

`func (o *MicrosoftGraphResetPasscodeActionResult) SetPasscode(v string)`

SetPasscode gets a reference to the given string and assigns it to the Passcode field.

### SetPasscodeExplicitNull

`func (o *MicrosoftGraphResetPasscodeActionResult) SetPasscodeExplicitNull(b bool)`

SetPasscodeExplicitNull (un)sets Passcode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Passcode value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


