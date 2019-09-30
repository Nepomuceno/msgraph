# MicrosoftGraphPlannerPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPlannerTask**](microsoft.graph.plannerTask.md) |  | [optional] 
**Buckets** | Pointer to [**[]MicrosoftGraphPlannerBucket**](microsoft.graph.plannerBucket.md) |  | [optional] 
**Details** | Pointer to [**AnyOfmicrosoftGraphPlannerPlanDetails**](anyOf&lt;microsoft.graph.plannerPlanDetails&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphPlannerPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPlannerPlan) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphPlannerPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphPlannerPlan) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedBy

`func (o *MicrosoftGraphPlannerPlan) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphPlannerPlan) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphPlannerPlan) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphPlannerPlan) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphPlannerPlan) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphPlannerPlan) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphPlannerPlan) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphPlannerPlan) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphPlannerPlan) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphPlannerPlan) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetOwner

`func (o *MicrosoftGraphPlannerPlan) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphPlannerPlan) GetOwnerOk() (string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *MicrosoftGraphPlannerPlan) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *MicrosoftGraphPlannerPlan) SetOwner(v string)`

SetOwner gets a reference to the given string and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *MicrosoftGraphPlannerPlan) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetTitle

`func (o *MicrosoftGraphPlannerPlan) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphPlannerPlan) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *MicrosoftGraphPlannerPlan) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *MicrosoftGraphPlannerPlan) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetTasks

`func (o *MicrosoftGraphPlannerPlan) GetTasks() []MicrosoftGraphPlannerTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *MicrosoftGraphPlannerPlan) GetTasksOk() ([]MicrosoftGraphPlannerTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTasks

`func (o *MicrosoftGraphPlannerPlan) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasks

`func (o *MicrosoftGraphPlannerPlan) SetTasks(v []MicrosoftGraphPlannerTask)`

SetTasks gets a reference to the given []MicrosoftGraphPlannerTask and assigns it to the Tasks field.

### GetBuckets

`func (o *MicrosoftGraphPlannerPlan) GetBuckets() []MicrosoftGraphPlannerBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *MicrosoftGraphPlannerPlan) GetBucketsOk() ([]MicrosoftGraphPlannerBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBuckets

`func (o *MicrosoftGraphPlannerPlan) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### SetBuckets

`func (o *MicrosoftGraphPlannerPlan) SetBuckets(v []MicrosoftGraphPlannerBucket)`

SetBuckets gets a reference to the given []MicrosoftGraphPlannerBucket and assigns it to the Buckets field.

### GetDetails

`func (o *MicrosoftGraphPlannerPlan) GetDetails() AnyOfmicrosoftGraphPlannerPlanDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphPlannerPlan) GetDetailsOk() (AnyOfmicrosoftGraphPlannerPlanDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetails

`func (o *MicrosoftGraphPlannerPlan) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetails

`func (o *MicrosoftGraphPlannerPlan) SetDetails(v AnyOfmicrosoftGraphPlannerPlanDetails)`

SetDetails gets a reference to the given AnyOfmicrosoftGraphPlannerPlanDetails and assigns it to the Details field.

### SetDetailsExplicitNull

`func (o *MicrosoftGraphPlannerPlan) SetDetailsExplicitNull(b bool)`

SetDetailsExplicitNull (un)sets Details to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Details value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


