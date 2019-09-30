# MicrosoftGraphLocateDeviceActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionName** | Pointer to **string** | Action name | [optional] 
**ActionState** | Pointer to [**AnyOfmicrosoftGraphActionState**](anyOf&lt;microsoft.graph.actionState&gt;.md) | State of the action | [optional] 
**StartDateTime** | Pointer to [**time.Time**](time.Time.md) | Time the action was initiated | [optional] 
**LastUpdatedDateTime** | Pointer to [**time.Time**](time.Time.md) | Time the action state was last updated | [optional] 
**DeviceLocation** | Pointer to [**AnyOfmicrosoftGraphDeviceGeoLocation**](anyOf&lt;microsoft.graph.deviceGeoLocation&gt;.md) | device location | [optional] 

## Methods

### GetActionName

`func (o *MicrosoftGraphLocateDeviceActionResult) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *MicrosoftGraphLocateDeviceActionResult) GetActionNameOk() (string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionName

`func (o *MicrosoftGraphLocateDeviceActionResult) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### SetActionName

`func (o *MicrosoftGraphLocateDeviceActionResult) SetActionName(v string)`

SetActionName gets a reference to the given string and assigns it to the ActionName field.

### SetActionNameExplicitNull

`func (o *MicrosoftGraphLocateDeviceActionResult) SetActionNameExplicitNull(b bool)`

SetActionNameExplicitNull (un)sets ActionName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActionName value is set to nil even if false is passed
### GetActionState

`func (o *MicrosoftGraphLocateDeviceActionResult) GetActionState() AnyOfmicrosoftGraphActionState`

GetActionState returns the ActionState field if non-nil, zero value otherwise.

### GetActionStateOk

`func (o *MicrosoftGraphLocateDeviceActionResult) GetActionStateOk() (AnyOfmicrosoftGraphActionState, bool)`

GetActionStateOk returns a tuple with the ActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionState

`func (o *MicrosoftGraphLocateDeviceActionResult) HasActionState() bool`

HasActionState returns a boolean if a field has been set.

### SetActionState

`func (o *MicrosoftGraphLocateDeviceActionResult) SetActionState(v AnyOfmicrosoftGraphActionState)`

SetActionState gets a reference to the given AnyOfmicrosoftGraphActionState and assigns it to the ActionState field.

### GetStartDateTime

`func (o *MicrosoftGraphLocateDeviceActionResult) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphLocateDeviceActionResult) GetStartDateTimeOk() (time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDateTime

`func (o *MicrosoftGraphLocateDeviceActionResult) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTime

`func (o *MicrosoftGraphLocateDeviceActionResult) SetStartDateTime(v time.Time)`

SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.

### GetLastUpdatedDateTime

`func (o *MicrosoftGraphLocateDeviceActionResult) GetLastUpdatedDateTime() time.Time`

GetLastUpdatedDateTime returns the LastUpdatedDateTime field if non-nil, zero value otherwise.

### GetLastUpdatedDateTimeOk

`func (o *MicrosoftGraphLocateDeviceActionResult) GetLastUpdatedDateTimeOk() (time.Time, bool)`

GetLastUpdatedDateTimeOk returns a tuple with the LastUpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdatedDateTime

`func (o *MicrosoftGraphLocateDeviceActionResult) HasLastUpdatedDateTime() bool`

HasLastUpdatedDateTime returns a boolean if a field has been set.

### SetLastUpdatedDateTime

`func (o *MicrosoftGraphLocateDeviceActionResult) SetLastUpdatedDateTime(v time.Time)`

SetLastUpdatedDateTime gets a reference to the given time.Time and assigns it to the LastUpdatedDateTime field.

### GetDeviceLocation

`func (o *MicrosoftGraphLocateDeviceActionResult) GetDeviceLocation() AnyOfmicrosoftGraphDeviceGeoLocation`

GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.

### GetDeviceLocationOk

`func (o *MicrosoftGraphLocateDeviceActionResult) GetDeviceLocationOk() (AnyOfmicrosoftGraphDeviceGeoLocation, bool)`

GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceLocation

`func (o *MicrosoftGraphLocateDeviceActionResult) HasDeviceLocation() bool`

HasDeviceLocation returns a boolean if a field has been set.

### SetDeviceLocation

`func (o *MicrosoftGraphLocateDeviceActionResult) SetDeviceLocation(v AnyOfmicrosoftGraphDeviceGeoLocation)`

SetDeviceLocation gets a reference to the given AnyOfmicrosoftGraphDeviceGeoLocation and assigns it to the DeviceLocation field.

### SetDeviceLocationExplicitNull

`func (o *MicrosoftGraphLocateDeviceActionResult) SetDeviceLocationExplicitNull(b bool)`

SetDeviceLocationExplicitNull (un)sets DeviceLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceLocation value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


