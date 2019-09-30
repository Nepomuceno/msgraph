# PlannerPlanDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedWith** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**CategoryDescriptions** | Pointer to [**AnyOfmicrosoftGraphPlannerCategoryDescriptions**](anyOf&lt;microsoft.graph.plannerCategoryDescriptions&gt;.md) |  | [optional] 

## Methods

### GetSharedWith

`func (o *PlannerPlanDetails) GetSharedWith() AnyOfobject`

GetSharedWith returns the SharedWith field if non-nil, zero value otherwise.

### GetSharedWithOk

`func (o *PlannerPlanDetails) GetSharedWithOk() (AnyOfobject, bool)`

GetSharedWithOk returns a tuple with the SharedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharedWith

`func (o *PlannerPlanDetails) HasSharedWith() bool`

HasSharedWith returns a boolean if a field has been set.

### SetSharedWith

`func (o *PlannerPlanDetails) SetSharedWith(v AnyOfobject)`

SetSharedWith gets a reference to the given AnyOfobject and assigns it to the SharedWith field.

### SetSharedWithExplicitNull

`func (o *PlannerPlanDetails) SetSharedWithExplicitNull(b bool)`

SetSharedWithExplicitNull (un)sets SharedWith to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharedWith value is set to nil even if false is passed
### GetCategoryDescriptions

`func (o *PlannerPlanDetails) GetCategoryDescriptions() AnyOfmicrosoftGraphPlannerCategoryDescriptions`

GetCategoryDescriptions returns the CategoryDescriptions field if non-nil, zero value otherwise.

### GetCategoryDescriptionsOk

`func (o *PlannerPlanDetails) GetCategoryDescriptionsOk() (AnyOfmicrosoftGraphPlannerCategoryDescriptions, bool)`

GetCategoryDescriptionsOk returns a tuple with the CategoryDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategoryDescriptions

`func (o *PlannerPlanDetails) HasCategoryDescriptions() bool`

HasCategoryDescriptions returns a boolean if a field has been set.

### SetCategoryDescriptions

`func (o *PlannerPlanDetails) SetCategoryDescriptions(v AnyOfmicrosoftGraphPlannerCategoryDescriptions)`

SetCategoryDescriptions gets a reference to the given AnyOfmicrosoftGraphPlannerCategoryDescriptions and assigns it to the CategoryDescriptions field.

### SetCategoryDescriptionsExplicitNull

`func (o *PlannerPlanDetails) SetCategoryDescriptionsExplicitNull(b bool)`

SetCategoryDescriptionsExplicitNull (un)sets CategoryDescriptions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CategoryDescriptions value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


