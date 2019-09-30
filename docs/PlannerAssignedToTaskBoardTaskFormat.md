# PlannerAssignedToTaskBoardTaskFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnassignedOrderHint** | Pointer to **string** |  | [optional] 
**OrderHintsByAssignee** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 

## Methods

### GetUnassignedOrderHint

`func (o *PlannerAssignedToTaskBoardTaskFormat) GetUnassignedOrderHint() string`

GetUnassignedOrderHint returns the UnassignedOrderHint field if non-nil, zero value otherwise.

### GetUnassignedOrderHintOk

`func (o *PlannerAssignedToTaskBoardTaskFormat) GetUnassignedOrderHintOk() (string, bool)`

GetUnassignedOrderHintOk returns a tuple with the UnassignedOrderHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnassignedOrderHint

`func (o *PlannerAssignedToTaskBoardTaskFormat) HasUnassignedOrderHint() bool`

HasUnassignedOrderHint returns a boolean if a field has been set.

### SetUnassignedOrderHint

`func (o *PlannerAssignedToTaskBoardTaskFormat) SetUnassignedOrderHint(v string)`

SetUnassignedOrderHint gets a reference to the given string and assigns it to the UnassignedOrderHint field.

### SetUnassignedOrderHintExplicitNull

`func (o *PlannerAssignedToTaskBoardTaskFormat) SetUnassignedOrderHintExplicitNull(b bool)`

SetUnassignedOrderHintExplicitNull (un)sets UnassignedOrderHint to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UnassignedOrderHint value is set to nil even if false is passed
### GetOrderHintsByAssignee

`func (o *PlannerAssignedToTaskBoardTaskFormat) GetOrderHintsByAssignee() AnyOfobject`

GetOrderHintsByAssignee returns the OrderHintsByAssignee field if non-nil, zero value otherwise.

### GetOrderHintsByAssigneeOk

`func (o *PlannerAssignedToTaskBoardTaskFormat) GetOrderHintsByAssigneeOk() (AnyOfobject, bool)`

GetOrderHintsByAssigneeOk returns a tuple with the OrderHintsByAssignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrderHintsByAssignee

`func (o *PlannerAssignedToTaskBoardTaskFormat) HasOrderHintsByAssignee() bool`

HasOrderHintsByAssignee returns a boolean if a field has been set.

### SetOrderHintsByAssignee

`func (o *PlannerAssignedToTaskBoardTaskFormat) SetOrderHintsByAssignee(v AnyOfobject)`

SetOrderHintsByAssignee gets a reference to the given AnyOfobject and assigns it to the OrderHintsByAssignee field.

### SetOrderHintsByAssigneeExplicitNull

`func (o *PlannerAssignedToTaskBoardTaskFormat) SetOrderHintsByAssigneeExplicitNull(b bool)`

SetOrderHintsByAssigneeExplicitNull (un)sets OrderHintsByAssignee to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OrderHintsByAssignee value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


