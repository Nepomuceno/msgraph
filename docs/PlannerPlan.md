# PlannerPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPlannerTask**](microsoft.graph.plannerTask.md) |  | [optional] 
**Buckets** | Pointer to [**[]MicrosoftGraphPlannerBucket**](microsoft.graph.plannerBucket.md) |  | [optional] 
**Details** | Pointer to [**AnyOfmicrosoftGraphPlannerPlanDetails**](anyOf&lt;microsoft.graph.plannerPlanDetails&gt;.md) |  | [optional] 

## Methods

### GetCreatedBy

`func (o *PlannerPlan) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PlannerPlan) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *PlannerPlan) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *PlannerPlan) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *PlannerPlan) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *PlannerPlan) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *PlannerPlan) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *PlannerPlan) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *PlannerPlan) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *PlannerPlan) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetOwner

`func (o *PlannerPlan) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PlannerPlan) GetOwnerOk() (string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *PlannerPlan) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *PlannerPlan) SetOwner(v string)`

SetOwner gets a reference to the given string and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *PlannerPlan) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetTitle

`func (o *PlannerPlan) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlannerPlan) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *PlannerPlan) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *PlannerPlan) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetTasks

`func (o *PlannerPlan) GetTasks() []MicrosoftGraphPlannerTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *PlannerPlan) GetTasksOk() ([]MicrosoftGraphPlannerTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTasks

`func (o *PlannerPlan) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasks

`func (o *PlannerPlan) SetTasks(v []MicrosoftGraphPlannerTask)`

SetTasks gets a reference to the given []MicrosoftGraphPlannerTask and assigns it to the Tasks field.

### GetBuckets

`func (o *PlannerPlan) GetBuckets() []MicrosoftGraphPlannerBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *PlannerPlan) GetBucketsOk() ([]MicrosoftGraphPlannerBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBuckets

`func (o *PlannerPlan) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### SetBuckets

`func (o *PlannerPlan) SetBuckets(v []MicrosoftGraphPlannerBucket)`

SetBuckets gets a reference to the given []MicrosoftGraphPlannerBucket and assigns it to the Buckets field.

### GetDetails

`func (o *PlannerPlan) GetDetails() AnyOfmicrosoftGraphPlannerPlanDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PlannerPlan) GetDetailsOk() (AnyOfmicrosoftGraphPlannerPlanDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetails

`func (o *PlannerPlan) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetails

`func (o *PlannerPlan) SetDetails(v AnyOfmicrosoftGraphPlannerPlanDetails)`

SetDetails gets a reference to the given AnyOfmicrosoftGraphPlannerPlanDetails and assigns it to the Details field.

### SetDetailsExplicitNull

`func (o *PlannerPlan) SetDetailsExplicitNull(b bool)`

SetDetailsExplicitNull (un)sets Details to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Details value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


