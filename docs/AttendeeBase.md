# AttendeeBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AnyOfmicrosoftGraphAttendeeType**](anyOf&lt;microsoft.graph.attendeeType&gt;.md) |  | [optional] 

## Methods

### GetType

`func (o *AttendeeBase) GetType() AnyOfmicrosoftGraphAttendeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttendeeBase) GetTypeOk() (AnyOfmicrosoftGraphAttendeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *AttendeeBase) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *AttendeeBase) SetType(v AnyOfmicrosoftGraphAttendeeType)`

SetType gets a reference to the given AnyOfmicrosoftGraphAttendeeType and assigns it to the Type field.

### SetTypeExplicitNull

`func (o *AttendeeBase) SetTypeExplicitNull(b bool)`

SetTypeExplicitNull (un)sets Type to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Type value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


