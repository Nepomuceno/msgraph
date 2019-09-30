# MicrosoftGraphWindowsDefenderScanActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionName** | Pointer to **string** | Action name | [optional] 
**ActionState** | Pointer to [**AnyOfmicrosoftGraphActionState**](anyOf&lt;microsoft.graph.actionState&gt;.md) | State of the action | [optional] 
**StartDateTime** | Pointer to [**time.Time**](time.Time.md) | Time the action was initiated | [optional] 
**LastUpdatedDateTime** | Pointer to [**time.Time**](time.Time.md) | Time the action state was last updated | [optional] 
**ScanType** | Pointer to **string** | Scan type either full scan or quick scan | [optional] 

## Methods

### GetActionName

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) GetActionNameOk() (string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionName

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### SetActionName

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) SetActionName(v string)`

SetActionName gets a reference to the given string and assigns it to the ActionName field.

### SetActionNameExplicitNull

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) SetActionNameExplicitNull(b bool)`

SetActionNameExplicitNull (un)sets ActionName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActionName value is set to nil even if false is passed
### GetActionState

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) GetActionState() AnyOfmicrosoftGraphActionState`

GetActionState returns the ActionState field if non-nil, zero value otherwise.

### GetActionStateOk

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) GetActionStateOk() (AnyOfmicrosoftGraphActionState, bool)`

GetActionStateOk returns a tuple with the ActionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionState

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) HasActionState() bool`

HasActionState returns a boolean if a field has been set.

### SetActionState

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) SetActionState(v AnyOfmicrosoftGraphActionState)`

SetActionState gets a reference to the given AnyOfmicrosoftGraphActionState and assigns it to the ActionState field.

### GetStartDateTime

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) GetStartDateTimeOk() (time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDateTime

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTime

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) SetStartDateTime(v time.Time)`

SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.

### GetLastUpdatedDateTime

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) GetLastUpdatedDateTime() time.Time`

GetLastUpdatedDateTime returns the LastUpdatedDateTime field if non-nil, zero value otherwise.

### GetLastUpdatedDateTimeOk

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) GetLastUpdatedDateTimeOk() (time.Time, bool)`

GetLastUpdatedDateTimeOk returns a tuple with the LastUpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdatedDateTime

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) HasLastUpdatedDateTime() bool`

HasLastUpdatedDateTime returns a boolean if a field has been set.

### SetLastUpdatedDateTime

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) SetLastUpdatedDateTime(v time.Time)`

SetLastUpdatedDateTime gets a reference to the given time.Time and assigns it to the LastUpdatedDateTime field.

### GetScanType

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) GetScanTypeOk() (string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScanType

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### SetScanType

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) SetScanType(v string)`

SetScanType gets a reference to the given string and assigns it to the ScanType field.

### SetScanTypeExplicitNull

`func (o *MicrosoftGraphWindowsDefenderScanActionResult) SetScanTypeExplicitNull(b bool)`

SetScanTypeExplicitNull (un)sets ScanType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ScanType value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


