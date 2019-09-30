# MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UnassignedOrderHint** | Pointer to **string** |  | [optional] 
**OrderHintsByAssignee** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetUnassignedOrderHint

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) GetUnassignedOrderHint() string`

GetUnassignedOrderHint returns the UnassignedOrderHint field if non-nil, zero value otherwise.

### GetUnassignedOrderHintOk

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) GetUnassignedOrderHintOk() (string, bool)`

GetUnassignedOrderHintOk returns a tuple with the UnassignedOrderHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnassignedOrderHint

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) HasUnassignedOrderHint() bool`

HasUnassignedOrderHint returns a boolean if a field has been set.

### SetUnassignedOrderHint

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) SetUnassignedOrderHint(v string)`

SetUnassignedOrderHint gets a reference to the given string and assigns it to the UnassignedOrderHint field.

### SetUnassignedOrderHintExplicitNull

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) SetUnassignedOrderHintExplicitNull(b bool)`

SetUnassignedOrderHintExplicitNull (un)sets UnassignedOrderHint to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UnassignedOrderHint value is set to nil even if false is passed
### GetOrderHintsByAssignee

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) GetOrderHintsByAssignee() AnyOfobject`

GetOrderHintsByAssignee returns the OrderHintsByAssignee field if non-nil, zero value otherwise.

### GetOrderHintsByAssigneeOk

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) GetOrderHintsByAssigneeOk() (AnyOfobject, bool)`

GetOrderHintsByAssigneeOk returns a tuple with the OrderHintsByAssignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrderHintsByAssignee

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) HasOrderHintsByAssignee() bool`

HasOrderHintsByAssignee returns a boolean if a field has been set.

### SetOrderHintsByAssignee

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) SetOrderHintsByAssignee(v AnyOfobject)`

SetOrderHintsByAssignee gets a reference to the given AnyOfobject and assigns it to the OrderHintsByAssignee field.

### SetOrderHintsByAssigneeExplicitNull

`func (o *MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat) SetOrderHintsByAssigneeExplicitNull(b bool)`

SetOrderHintsByAssigneeExplicitNull (un)sets OrderHintsByAssignee to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OrderHintsByAssignee value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


