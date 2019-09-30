# MicrosoftGraphCalendar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Color** | Pointer to [**AnyOfmicrosoftGraphCalendarColor**](anyOf&lt;microsoft.graph.calendarColor&gt;.md) |  | [optional] 
**ChangeKey** | Pointer to **string** |  | [optional] 
**CanShare** | Pointer to **bool** |  | [optional] 
**CanViewPrivateItems** | Pointer to **bool** |  | [optional] 
**CanEdit** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) |  | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md) |  | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md) |  | [optional] 
**Events** | Pointer to [**[]MicrosoftGraphEvent**](microsoft.graph.event.md) |  | [optional] 
**CalendarView** | Pointer to [**[]MicrosoftGraphEvent**](microsoft.graph.event.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphCalendar) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphCalendar) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphCalendar) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphCalendar) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetName

`func (o *MicrosoftGraphCalendar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphCalendar) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphCalendar) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphCalendar) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphCalendar) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetColor

`func (o *MicrosoftGraphCalendar) GetColor() AnyOfmicrosoftGraphCalendarColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *MicrosoftGraphCalendar) GetColorOk() (AnyOfmicrosoftGraphCalendarColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColor

`func (o *MicrosoftGraphCalendar) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColor

`func (o *MicrosoftGraphCalendar) SetColor(v AnyOfmicrosoftGraphCalendarColor)`

SetColor gets a reference to the given AnyOfmicrosoftGraphCalendarColor and assigns it to the Color field.

### SetColorExplicitNull

`func (o *MicrosoftGraphCalendar) SetColorExplicitNull(b bool)`

SetColorExplicitNull (un)sets Color to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Color value is set to nil even if false is passed
### GetChangeKey

`func (o *MicrosoftGraphCalendar) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphCalendar) GetChangeKeyOk() (string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangeKey

`func (o *MicrosoftGraphCalendar) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKey

`func (o *MicrosoftGraphCalendar) SetChangeKey(v string)`

SetChangeKey gets a reference to the given string and assigns it to the ChangeKey field.

### SetChangeKeyExplicitNull

`func (o *MicrosoftGraphCalendar) SetChangeKeyExplicitNull(b bool)`

SetChangeKeyExplicitNull (un)sets ChangeKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChangeKey value is set to nil even if false is passed
### GetCanShare

`func (o *MicrosoftGraphCalendar) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *MicrosoftGraphCalendar) GetCanShareOk() (bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanShare

`func (o *MicrosoftGraphCalendar) HasCanShare() bool`

HasCanShare returns a boolean if a field has been set.

### SetCanShare

`func (o *MicrosoftGraphCalendar) SetCanShare(v bool)`

SetCanShare gets a reference to the given bool and assigns it to the CanShare field.

### SetCanShareExplicitNull

`func (o *MicrosoftGraphCalendar) SetCanShareExplicitNull(b bool)`

SetCanShareExplicitNull (un)sets CanShare to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CanShare value is set to nil even if false is passed
### GetCanViewPrivateItems

`func (o *MicrosoftGraphCalendar) GetCanViewPrivateItems() bool`

GetCanViewPrivateItems returns the CanViewPrivateItems field if non-nil, zero value otherwise.

### GetCanViewPrivateItemsOk

`func (o *MicrosoftGraphCalendar) GetCanViewPrivateItemsOk() (bool, bool)`

GetCanViewPrivateItemsOk returns a tuple with the CanViewPrivateItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanViewPrivateItems

`func (o *MicrosoftGraphCalendar) HasCanViewPrivateItems() bool`

HasCanViewPrivateItems returns a boolean if a field has been set.

### SetCanViewPrivateItems

`func (o *MicrosoftGraphCalendar) SetCanViewPrivateItems(v bool)`

SetCanViewPrivateItems gets a reference to the given bool and assigns it to the CanViewPrivateItems field.

### SetCanViewPrivateItemsExplicitNull

`func (o *MicrosoftGraphCalendar) SetCanViewPrivateItemsExplicitNull(b bool)`

SetCanViewPrivateItemsExplicitNull (un)sets CanViewPrivateItems to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CanViewPrivateItems value is set to nil even if false is passed
### GetCanEdit

`func (o *MicrosoftGraphCalendar) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *MicrosoftGraphCalendar) GetCanEditOk() (bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanEdit

`func (o *MicrosoftGraphCalendar) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### SetCanEdit

`func (o *MicrosoftGraphCalendar) SetCanEdit(v bool)`

SetCanEdit gets a reference to the given bool and assigns it to the CanEdit field.

### SetCanEditExplicitNull

`func (o *MicrosoftGraphCalendar) SetCanEditExplicitNull(b bool)`

SetCanEditExplicitNull (un)sets CanEdit to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CanEdit value is set to nil even if false is passed
### GetOwner

`func (o *MicrosoftGraphCalendar) GetOwner() AnyOfmicrosoftGraphEmailAddress`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphCalendar) GetOwnerOk() (AnyOfmicrosoftGraphEmailAddress, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *MicrosoftGraphCalendar) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *MicrosoftGraphCalendar) SetOwner(v AnyOfmicrosoftGraphEmailAddress)`

SetOwner gets a reference to the given AnyOfmicrosoftGraphEmailAddress and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *MicrosoftGraphCalendar) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphCalendar) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphCalendar) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphCalendar) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphCalendar) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphCalendar) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphCalendar) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphCalendar) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphCalendar) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetEvents

`func (o *MicrosoftGraphCalendar) GetEvents() []MicrosoftGraphEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *MicrosoftGraphCalendar) GetEventsOk() ([]MicrosoftGraphEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvents

`func (o *MicrosoftGraphCalendar) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEvents

`func (o *MicrosoftGraphCalendar) SetEvents(v []MicrosoftGraphEvent)`

SetEvents gets a reference to the given []MicrosoftGraphEvent and assigns it to the Events field.

### GetCalendarView

`func (o *MicrosoftGraphCalendar) GetCalendarView() []MicrosoftGraphEvent`

GetCalendarView returns the CalendarView field if non-nil, zero value otherwise.

### GetCalendarViewOk

`func (o *MicrosoftGraphCalendar) GetCalendarViewOk() ([]MicrosoftGraphEvent, bool)`

GetCalendarViewOk returns a tuple with the CalendarView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendarView

`func (o *MicrosoftGraphCalendar) HasCalendarView() bool`

HasCalendarView returns a boolean if a field has been set.

### SetCalendarView

`func (o *MicrosoftGraphCalendar) SetCalendarView(v []MicrosoftGraphEvent)`

SetCalendarView gets a reference to the given []MicrosoftGraphEvent and assigns it to the CalendarView field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


