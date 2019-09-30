# MicrosoftGraphIncompleteData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MissingDataBeforeDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**WasThrottled** | Pointer to **bool** |  | [optional] 

## Methods

### GetMissingDataBeforeDateTime

`func (o *MicrosoftGraphIncompleteData) GetMissingDataBeforeDateTime() time.Time`

GetMissingDataBeforeDateTime returns the MissingDataBeforeDateTime field if non-nil, zero value otherwise.

### GetMissingDataBeforeDateTimeOk

`func (o *MicrosoftGraphIncompleteData) GetMissingDataBeforeDateTimeOk() (time.Time, bool)`

GetMissingDataBeforeDateTimeOk returns a tuple with the MissingDataBeforeDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMissingDataBeforeDateTime

`func (o *MicrosoftGraphIncompleteData) HasMissingDataBeforeDateTime() bool`

HasMissingDataBeforeDateTime returns a boolean if a field has been set.

### SetMissingDataBeforeDateTime

`func (o *MicrosoftGraphIncompleteData) SetMissingDataBeforeDateTime(v time.Time)`

SetMissingDataBeforeDateTime gets a reference to the given time.Time and assigns it to the MissingDataBeforeDateTime field.

### SetMissingDataBeforeDateTimeExplicitNull

`func (o *MicrosoftGraphIncompleteData) SetMissingDataBeforeDateTimeExplicitNull(b bool)`

SetMissingDataBeforeDateTimeExplicitNull (un)sets MissingDataBeforeDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MissingDataBeforeDateTime value is set to nil even if false is passed
### GetWasThrottled

`func (o *MicrosoftGraphIncompleteData) GetWasThrottled() bool`

GetWasThrottled returns the WasThrottled field if non-nil, zero value otherwise.

### GetWasThrottledOk

`func (o *MicrosoftGraphIncompleteData) GetWasThrottledOk() (bool, bool)`

GetWasThrottledOk returns a tuple with the WasThrottled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWasThrottled

`func (o *MicrosoftGraphIncompleteData) HasWasThrottled() bool`

HasWasThrottled returns a boolean if a field has been set.

### SetWasThrottled

`func (o *MicrosoftGraphIncompleteData) SetWasThrottled(v bool)`

SetWasThrottled gets a reference to the given bool and assigns it to the WasThrottled field.

### SetWasThrottledExplicitNull

`func (o *MicrosoftGraphIncompleteData) SetWasThrottledExplicitNull(b bool)`

SetWasThrottledExplicitNull (un)sets WasThrottled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WasThrottled value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


