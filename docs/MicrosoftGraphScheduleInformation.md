# MicrosoftGraphScheduleInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleId** | Pointer to **string** |  | [optional] 
**ScheduleItems** | Pointer to [**[]AnyOfmicrosoftGraphScheduleItem**](anyOf&lt;microsoft.graph.scheduleItem&gt;.md) |  | [optional] 
**AvailabilityView** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**AnyOfmicrosoftGraphFreeBusyError**](anyOf&lt;microsoft.graph.freeBusyError&gt;.md) |  | [optional] 
**WorkingHours** | Pointer to [**AnyOfmicrosoftGraphWorkingHours**](anyOf&lt;microsoft.graph.workingHours&gt;.md) |  | [optional] 

## Methods

### GetScheduleId

`func (o *MicrosoftGraphScheduleInformation) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *MicrosoftGraphScheduleInformation) GetScheduleIdOk() (string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduleId

`func (o *MicrosoftGraphScheduleInformation) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleId

`func (o *MicrosoftGraphScheduleInformation) SetScheduleId(v string)`

SetScheduleId gets a reference to the given string and assigns it to the ScheduleId field.

### SetScheduleIdExplicitNull

`func (o *MicrosoftGraphScheduleInformation) SetScheduleIdExplicitNull(b bool)`

SetScheduleIdExplicitNull (un)sets ScheduleId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ScheduleId value is set to nil even if false is passed
### GetScheduleItems

`func (o *MicrosoftGraphScheduleInformation) GetScheduleItems() []AnyOfmicrosoftGraphScheduleItem`

GetScheduleItems returns the ScheduleItems field if non-nil, zero value otherwise.

### GetScheduleItemsOk

`func (o *MicrosoftGraphScheduleInformation) GetScheduleItemsOk() ([]AnyOfmicrosoftGraphScheduleItem, bool)`

GetScheduleItemsOk returns a tuple with the ScheduleItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduleItems

`func (o *MicrosoftGraphScheduleInformation) HasScheduleItems() bool`

HasScheduleItems returns a boolean if a field has been set.

### SetScheduleItems

`func (o *MicrosoftGraphScheduleInformation) SetScheduleItems(v []AnyOfmicrosoftGraphScheduleItem)`

SetScheduleItems gets a reference to the given []AnyOfmicrosoftGraphScheduleItem and assigns it to the ScheduleItems field.

### GetAvailabilityView

`func (o *MicrosoftGraphScheduleInformation) GetAvailabilityView() string`

GetAvailabilityView returns the AvailabilityView field if non-nil, zero value otherwise.

### GetAvailabilityViewOk

`func (o *MicrosoftGraphScheduleInformation) GetAvailabilityViewOk() (string, bool)`

GetAvailabilityViewOk returns a tuple with the AvailabilityView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAvailabilityView

`func (o *MicrosoftGraphScheduleInformation) HasAvailabilityView() bool`

HasAvailabilityView returns a boolean if a field has been set.

### SetAvailabilityView

`func (o *MicrosoftGraphScheduleInformation) SetAvailabilityView(v string)`

SetAvailabilityView gets a reference to the given string and assigns it to the AvailabilityView field.

### SetAvailabilityViewExplicitNull

`func (o *MicrosoftGraphScheduleInformation) SetAvailabilityViewExplicitNull(b bool)`

SetAvailabilityViewExplicitNull (un)sets AvailabilityView to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AvailabilityView value is set to nil even if false is passed
### GetError

`func (o *MicrosoftGraphScheduleInformation) GetError() AnyOfmicrosoftGraphFreeBusyError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphScheduleInformation) GetErrorOk() (AnyOfmicrosoftGraphFreeBusyError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *MicrosoftGraphScheduleInformation) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *MicrosoftGraphScheduleInformation) SetError(v AnyOfmicrosoftGraphFreeBusyError)`

SetError gets a reference to the given AnyOfmicrosoftGraphFreeBusyError and assigns it to the Error field.

### SetErrorExplicitNull

`func (o *MicrosoftGraphScheduleInformation) SetErrorExplicitNull(b bool)`

SetErrorExplicitNull (un)sets Error to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Error value is set to nil even if false is passed
### GetWorkingHours

`func (o *MicrosoftGraphScheduleInformation) GetWorkingHours() AnyOfmicrosoftGraphWorkingHours`

GetWorkingHours returns the WorkingHours field if non-nil, zero value otherwise.

### GetWorkingHoursOk

`func (o *MicrosoftGraphScheduleInformation) GetWorkingHoursOk() (AnyOfmicrosoftGraphWorkingHours, bool)`

GetWorkingHoursOk returns a tuple with the WorkingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkingHours

`func (o *MicrosoftGraphScheduleInformation) HasWorkingHours() bool`

HasWorkingHours returns a boolean if a field has been set.

### SetWorkingHours

`func (o *MicrosoftGraphScheduleInformation) SetWorkingHours(v AnyOfmicrosoftGraphWorkingHours)`

SetWorkingHours gets a reference to the given AnyOfmicrosoftGraphWorkingHours and assigns it to the WorkingHours field.

### SetWorkingHoursExplicitNull

`func (o *MicrosoftGraphScheduleInformation) SetWorkingHoursExplicitNull(b bool)`

SetWorkingHoursExplicitNull (un)sets WorkingHours to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkingHours value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


