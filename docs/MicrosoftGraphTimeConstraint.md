# MicrosoftGraphTimeConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityDomain** | Pointer to [**AnyOfmicrosoftGraphActivityDomain**](anyOf&lt;microsoft.graph.activityDomain&gt;.md) |  | [optional] 
**TimeSlots** | Pointer to [**[]AnyOfmicrosoftGraphTimeSlot**](anyOf&lt;microsoft.graph.timeSlot&gt;.md) |  | [optional] 

## Methods

### GetActivityDomain

`func (o *MicrosoftGraphTimeConstraint) GetActivityDomain() AnyOfmicrosoftGraphActivityDomain`

GetActivityDomain returns the ActivityDomain field if non-nil, zero value otherwise.

### GetActivityDomainOk

`func (o *MicrosoftGraphTimeConstraint) GetActivityDomainOk() (AnyOfmicrosoftGraphActivityDomain, bool)`

GetActivityDomainOk returns a tuple with the ActivityDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivityDomain

`func (o *MicrosoftGraphTimeConstraint) HasActivityDomain() bool`

HasActivityDomain returns a boolean if a field has been set.

### SetActivityDomain

`func (o *MicrosoftGraphTimeConstraint) SetActivityDomain(v AnyOfmicrosoftGraphActivityDomain)`

SetActivityDomain gets a reference to the given AnyOfmicrosoftGraphActivityDomain and assigns it to the ActivityDomain field.

### SetActivityDomainExplicitNull

`func (o *MicrosoftGraphTimeConstraint) SetActivityDomainExplicitNull(b bool)`

SetActivityDomainExplicitNull (un)sets ActivityDomain to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActivityDomain value is set to nil even if false is passed
### GetTimeSlots

`func (o *MicrosoftGraphTimeConstraint) GetTimeSlots() []AnyOfmicrosoftGraphTimeSlot`

GetTimeSlots returns the TimeSlots field if non-nil, zero value otherwise.

### GetTimeSlotsOk

`func (o *MicrosoftGraphTimeConstraint) GetTimeSlotsOk() ([]AnyOfmicrosoftGraphTimeSlot, bool)`

GetTimeSlotsOk returns a tuple with the TimeSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeSlots

`func (o *MicrosoftGraphTimeConstraint) HasTimeSlots() bool`

HasTimeSlots returns a boolean if a field has been set.

### SetTimeSlots

`func (o *MicrosoftGraphTimeConstraint) SetTimeSlots(v []AnyOfmicrosoftGraphTimeSlot)`

SetTimeSlots gets a reference to the given []AnyOfmicrosoftGraphTimeSlot and assigns it to the TimeSlots field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


