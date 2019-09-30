# MicrosoftGraphPlannerAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**AssignedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**OrderHint** | Pointer to **string** |  | [optional] 

## Methods

### GetAssignedBy

`func (o *MicrosoftGraphPlannerAssignment) GetAssignedBy() AnyOfmicrosoftGraphIdentitySet`

GetAssignedBy returns the AssignedBy field if non-nil, zero value otherwise.

### GetAssignedByOk

`func (o *MicrosoftGraphPlannerAssignment) GetAssignedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetAssignedByOk returns a tuple with the AssignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedBy

`func (o *MicrosoftGraphPlannerAssignment) HasAssignedBy() bool`

HasAssignedBy returns a boolean if a field has been set.

### SetAssignedBy

`func (o *MicrosoftGraphPlannerAssignment) SetAssignedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetAssignedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the AssignedBy field.

### SetAssignedByExplicitNull

`func (o *MicrosoftGraphPlannerAssignment) SetAssignedByExplicitNull(b bool)`

SetAssignedByExplicitNull (un)sets AssignedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AssignedBy value is set to nil even if false is passed
### GetAssignedDateTime

`func (o *MicrosoftGraphPlannerAssignment) GetAssignedDateTime() time.Time`

GetAssignedDateTime returns the AssignedDateTime field if non-nil, zero value otherwise.

### GetAssignedDateTimeOk

`func (o *MicrosoftGraphPlannerAssignment) GetAssignedDateTimeOk() (time.Time, bool)`

GetAssignedDateTimeOk returns a tuple with the AssignedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedDateTime

`func (o *MicrosoftGraphPlannerAssignment) HasAssignedDateTime() bool`

HasAssignedDateTime returns a boolean if a field has been set.

### SetAssignedDateTime

`func (o *MicrosoftGraphPlannerAssignment) SetAssignedDateTime(v time.Time)`

SetAssignedDateTime gets a reference to the given time.Time and assigns it to the AssignedDateTime field.

### SetAssignedDateTimeExplicitNull

`func (o *MicrosoftGraphPlannerAssignment) SetAssignedDateTimeExplicitNull(b bool)`

SetAssignedDateTimeExplicitNull (un)sets AssignedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AssignedDateTime value is set to nil even if false is passed
### GetOrderHint

`func (o *MicrosoftGraphPlannerAssignment) GetOrderHint() string`

GetOrderHint returns the OrderHint field if non-nil, zero value otherwise.

### GetOrderHintOk

`func (o *MicrosoftGraphPlannerAssignment) GetOrderHintOk() (string, bool)`

GetOrderHintOk returns a tuple with the OrderHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrderHint

`func (o *MicrosoftGraphPlannerAssignment) HasOrderHint() bool`

HasOrderHint returns a boolean if a field has been set.

### SetOrderHint

`func (o *MicrosoftGraphPlannerAssignment) SetOrderHint(v string)`

SetOrderHint gets a reference to the given string and assigns it to the OrderHint field.

### SetOrderHintExplicitNull

`func (o *MicrosoftGraphPlannerAssignment) SetOrderHintExplicitNull(b bool)`

SetOrderHintExplicitNull (un)sets OrderHint to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OrderHint value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


