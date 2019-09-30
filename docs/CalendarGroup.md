# CalendarGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ClassId** | Pointer to **string** |  | [optional] 
**ChangeKey** | Pointer to **string** |  | [optional] 
**Calendars** | Pointer to [**[]MicrosoftGraphCalendar**](microsoft.graph.calendar.md) |  | [optional] 

## Methods

### GetName

`func (o *CalendarGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CalendarGroup) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CalendarGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CalendarGroup) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *CalendarGroup) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetClassId

`func (o *CalendarGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CalendarGroup) GetClassIdOk() (string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClassId

`func (o *CalendarGroup) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### SetClassId

`func (o *CalendarGroup) SetClassId(v string)`

SetClassId gets a reference to the given string and assigns it to the ClassId field.

### SetClassIdExplicitNull

`func (o *CalendarGroup) SetClassIdExplicitNull(b bool)`

SetClassIdExplicitNull (un)sets ClassId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ClassId value is set to nil even if false is passed
### GetChangeKey

`func (o *CalendarGroup) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *CalendarGroup) GetChangeKeyOk() (string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangeKey

`func (o *CalendarGroup) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKey

`func (o *CalendarGroup) SetChangeKey(v string)`

SetChangeKey gets a reference to the given string and assigns it to the ChangeKey field.

### SetChangeKeyExplicitNull

`func (o *CalendarGroup) SetChangeKeyExplicitNull(b bool)`

SetChangeKeyExplicitNull (un)sets ChangeKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChangeKey value is set to nil even if false is passed
### GetCalendars

`func (o *CalendarGroup) GetCalendars() []MicrosoftGraphCalendar`

GetCalendars returns the Calendars field if non-nil, zero value otherwise.

### GetCalendarsOk

`func (o *CalendarGroup) GetCalendarsOk() ([]MicrosoftGraphCalendar, bool)`

GetCalendarsOk returns a tuple with the Calendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendars

`func (o *CalendarGroup) HasCalendars() bool`

HasCalendars returns a boolean if a field has been set.

### SetCalendars

`func (o *CalendarGroup) SetCalendars(v []MicrosoftGraphCalendar)`

SetCalendars gets a reference to the given []MicrosoftGraphCalendar and assigns it to the Calendars field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


