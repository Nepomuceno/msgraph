# PlannerBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**OrderHint** | Pointer to **string** |  | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPlannerTask**](microsoft.graph.plannerTask.md) |  | [optional] 

## Methods

### GetName

`func (o *PlannerBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlannerBucket) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *PlannerBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *PlannerBucket) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetPlanId

`func (o *PlannerBucket) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlannerBucket) GetPlanIdOk() (string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlanId

`func (o *PlannerBucket) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanId

`func (o *PlannerBucket) SetPlanId(v string)`

SetPlanId gets a reference to the given string and assigns it to the PlanId field.

### SetPlanIdExplicitNull

`func (o *PlannerBucket) SetPlanIdExplicitNull(b bool)`

SetPlanIdExplicitNull (un)sets PlanId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PlanId value is set to nil even if false is passed
### GetOrderHint

`func (o *PlannerBucket) GetOrderHint() string`

GetOrderHint returns the OrderHint field if non-nil, zero value otherwise.

### GetOrderHintOk

`func (o *PlannerBucket) GetOrderHintOk() (string, bool)`

GetOrderHintOk returns a tuple with the OrderHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrderHint

`func (o *PlannerBucket) HasOrderHint() bool`

HasOrderHint returns a boolean if a field has been set.

### SetOrderHint

`func (o *PlannerBucket) SetOrderHint(v string)`

SetOrderHint gets a reference to the given string and assigns it to the OrderHint field.

### SetOrderHintExplicitNull

`func (o *PlannerBucket) SetOrderHintExplicitNull(b bool)`

SetOrderHintExplicitNull (un)sets OrderHint to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OrderHint value is set to nil even if false is passed
### GetTasks

`func (o *PlannerBucket) GetTasks() []MicrosoftGraphPlannerTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *PlannerBucket) GetTasksOk() ([]MicrosoftGraphPlannerTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTasks

`func (o *PlannerBucket) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasks

`func (o *PlannerBucket) SetTasks(v []MicrosoftGraphPlannerTask)`

SetTasks gets a reference to the given []MicrosoftGraphPlannerTask and assigns it to the Tasks field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


