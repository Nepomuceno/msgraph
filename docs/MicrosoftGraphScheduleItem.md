# MicrosoftGraphScheduleItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**End** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**IsPrivate** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphFreeBusyStatus**](anyOf&lt;microsoft.graph.freeBusyStatus&gt;.md) |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 

## Methods

### GetStart

`func (o *MicrosoftGraphScheduleItem) GetStart() AnyOfmicrosoftGraphDateTimeTimeZone`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MicrosoftGraphScheduleItem) GetStartOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *MicrosoftGraphScheduleItem) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *MicrosoftGraphScheduleItem) SetStart(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetStart gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the Start field.

### SetStartExplicitNull

`func (o *MicrosoftGraphScheduleItem) SetStartExplicitNull(b bool)`

SetStartExplicitNull (un)sets Start to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Start value is set to nil even if false is passed
### GetEnd

`func (o *MicrosoftGraphScheduleItem) GetEnd() AnyOfmicrosoftGraphDateTimeTimeZone`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MicrosoftGraphScheduleItem) GetEndOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *MicrosoftGraphScheduleItem) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *MicrosoftGraphScheduleItem) SetEnd(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetEnd gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the End field.

### SetEndExplicitNull

`func (o *MicrosoftGraphScheduleItem) SetEndExplicitNull(b bool)`

SetEndExplicitNull (un)sets End to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The End value is set to nil even if false is passed
### GetIsPrivate

`func (o *MicrosoftGraphScheduleItem) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *MicrosoftGraphScheduleItem) GetIsPrivateOk() (bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsPrivate

`func (o *MicrosoftGraphScheduleItem) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### SetIsPrivate

`func (o *MicrosoftGraphScheduleItem) SetIsPrivate(v bool)`

SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.

### SetIsPrivateExplicitNull

`func (o *MicrosoftGraphScheduleItem) SetIsPrivateExplicitNull(b bool)`

SetIsPrivateExplicitNull (un)sets IsPrivate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsPrivate value is set to nil even if false is passed
### GetStatus

`func (o *MicrosoftGraphScheduleItem) GetStatus() AnyOfmicrosoftGraphFreeBusyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphScheduleItem) GetStatusOk() (AnyOfmicrosoftGraphFreeBusyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphScheduleItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphScheduleItem) SetStatus(v AnyOfmicrosoftGraphFreeBusyStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphFreeBusyStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphScheduleItem) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetSubject

`func (o *MicrosoftGraphScheduleItem) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MicrosoftGraphScheduleItem) GetSubjectOk() (string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubject

`func (o *MicrosoftGraphScheduleItem) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubject

`func (o *MicrosoftGraphScheduleItem) SetSubject(v string)`

SetSubject gets a reference to the given string and assigns it to the Subject field.

### SetSubjectExplicitNull

`func (o *MicrosoftGraphScheduleItem) SetSubjectExplicitNull(b bool)`

SetSubjectExplicitNull (un)sets Subject to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Subject value is set to nil even if false is passed
### GetLocation

`func (o *MicrosoftGraphScheduleItem) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphScheduleItem) GetLocationOk() (string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocation

`func (o *MicrosoftGraphScheduleItem) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocation

`func (o *MicrosoftGraphScheduleItem) SetLocation(v string)`

SetLocation gets a reference to the given string and assigns it to the Location field.

### SetLocationExplicitNull

`func (o *MicrosoftGraphScheduleItem) SetLocationExplicitNull(b bool)`

SetLocationExplicitNull (un)sets Location to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Location value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


