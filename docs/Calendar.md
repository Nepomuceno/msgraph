# Calendar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetName

`func (o *Calendar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Calendar) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Calendar) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Calendar) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *Calendar) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetColor

`func (o *Calendar) GetColor() AnyOfmicrosoftGraphCalendarColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Calendar) GetColorOk() (AnyOfmicrosoftGraphCalendarColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColor

`func (o *Calendar) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColor

`func (o *Calendar) SetColor(v AnyOfmicrosoftGraphCalendarColor)`

SetColor gets a reference to the given AnyOfmicrosoftGraphCalendarColor and assigns it to the Color field.

### SetColorExplicitNull

`func (o *Calendar) SetColorExplicitNull(b bool)`

SetColorExplicitNull (un)sets Color to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Color value is set to nil even if false is passed
### GetChangeKey

`func (o *Calendar) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *Calendar) GetChangeKeyOk() (string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangeKey

`func (o *Calendar) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKey

`func (o *Calendar) SetChangeKey(v string)`

SetChangeKey gets a reference to the given string and assigns it to the ChangeKey field.

### SetChangeKeyExplicitNull

`func (o *Calendar) SetChangeKeyExplicitNull(b bool)`

SetChangeKeyExplicitNull (un)sets ChangeKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChangeKey value is set to nil even if false is passed
### GetCanShare

`func (o *Calendar) GetCanShare() bool`

GetCanShare returns the CanShare field if non-nil, zero value otherwise.

### GetCanShareOk

`func (o *Calendar) GetCanShareOk() (bool, bool)`

GetCanShareOk returns a tuple with the CanShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanShare

`func (o *Calendar) HasCanShare() bool`

HasCanShare returns a boolean if a field has been set.

### SetCanShare

`func (o *Calendar) SetCanShare(v bool)`

SetCanShare gets a reference to the given bool and assigns it to the CanShare field.

### SetCanShareExplicitNull

`func (o *Calendar) SetCanShareExplicitNull(b bool)`

SetCanShareExplicitNull (un)sets CanShare to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CanShare value is set to nil even if false is passed
### GetCanViewPrivateItems

`func (o *Calendar) GetCanViewPrivateItems() bool`

GetCanViewPrivateItems returns the CanViewPrivateItems field if non-nil, zero value otherwise.

### GetCanViewPrivateItemsOk

`func (o *Calendar) GetCanViewPrivateItemsOk() (bool, bool)`

GetCanViewPrivateItemsOk returns a tuple with the CanViewPrivateItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanViewPrivateItems

`func (o *Calendar) HasCanViewPrivateItems() bool`

HasCanViewPrivateItems returns a boolean if a field has been set.

### SetCanViewPrivateItems

`func (o *Calendar) SetCanViewPrivateItems(v bool)`

SetCanViewPrivateItems gets a reference to the given bool and assigns it to the CanViewPrivateItems field.

### SetCanViewPrivateItemsExplicitNull

`func (o *Calendar) SetCanViewPrivateItemsExplicitNull(b bool)`

SetCanViewPrivateItemsExplicitNull (un)sets CanViewPrivateItems to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CanViewPrivateItems value is set to nil even if false is passed
### GetCanEdit

`func (o *Calendar) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *Calendar) GetCanEditOk() (bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanEdit

`func (o *Calendar) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### SetCanEdit

`func (o *Calendar) SetCanEdit(v bool)`

SetCanEdit gets a reference to the given bool and assigns it to the CanEdit field.

### SetCanEditExplicitNull

`func (o *Calendar) SetCanEditExplicitNull(b bool)`

SetCanEditExplicitNull (un)sets CanEdit to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CanEdit value is set to nil even if false is passed
### GetOwner

`func (o *Calendar) GetOwner() AnyOfmicrosoftGraphEmailAddress`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Calendar) GetOwnerOk() (AnyOfmicrosoftGraphEmailAddress, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *Calendar) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *Calendar) SetOwner(v AnyOfmicrosoftGraphEmailAddress)`

SetOwner gets a reference to the given AnyOfmicrosoftGraphEmailAddress and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *Calendar) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetSingleValueExtendedProperties

`func (o *Calendar) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *Calendar) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *Calendar) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *Calendar) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *Calendar) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *Calendar) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *Calendar) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *Calendar) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetEvents

`func (o *Calendar) GetEvents() []MicrosoftGraphEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Calendar) GetEventsOk() ([]MicrosoftGraphEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvents

`func (o *Calendar) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEvents

`func (o *Calendar) SetEvents(v []MicrosoftGraphEvent)`

SetEvents gets a reference to the given []MicrosoftGraphEvent and assigns it to the Events field.

### GetCalendarView

`func (o *Calendar) GetCalendarView() []MicrosoftGraphEvent`

GetCalendarView returns the CalendarView field if non-nil, zero value otherwise.

### GetCalendarViewOk

`func (o *Calendar) GetCalendarViewOk() ([]MicrosoftGraphEvent, bool)`

GetCalendarViewOk returns a tuple with the CalendarView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendarView

`func (o *Calendar) HasCalendarView() bool`

HasCalendarView returns a boolean if a field has been set.

### SetCalendarView

`func (o *Calendar) SetCalendarView(v []MicrosoftGraphEvent)`

SetCalendarView gets a reference to the given []MicrosoftGraphEvent and assigns it to the CalendarView field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


