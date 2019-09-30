# PlannerTaskDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**PreviewType** | Pointer to [**AnyOfmicrosoftGraphPlannerPreviewType**](anyOf&lt;microsoft.graph.plannerPreviewType&gt;.md) |  | [optional] 
**References** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Checklist** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 

## Methods

### GetDescription

`func (o *PlannerTaskDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlannerTaskDetails) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *PlannerTaskDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *PlannerTaskDetails) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *PlannerTaskDetails) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPreviewType

`func (o *PlannerTaskDetails) GetPreviewType() AnyOfmicrosoftGraphPlannerPreviewType`

GetPreviewType returns the PreviewType field if non-nil, zero value otherwise.

### GetPreviewTypeOk

`func (o *PlannerTaskDetails) GetPreviewTypeOk() (AnyOfmicrosoftGraphPlannerPreviewType, bool)`

GetPreviewTypeOk returns a tuple with the PreviewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreviewType

`func (o *PlannerTaskDetails) HasPreviewType() bool`

HasPreviewType returns a boolean if a field has been set.

### SetPreviewType

`func (o *PlannerTaskDetails) SetPreviewType(v AnyOfmicrosoftGraphPlannerPreviewType)`

SetPreviewType gets a reference to the given AnyOfmicrosoftGraphPlannerPreviewType and assigns it to the PreviewType field.

### SetPreviewTypeExplicitNull

`func (o *PlannerTaskDetails) SetPreviewTypeExplicitNull(b bool)`

SetPreviewTypeExplicitNull (un)sets PreviewType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreviewType value is set to nil even if false is passed
### GetReferences

`func (o *PlannerTaskDetails) GetReferences() AnyOfobject`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *PlannerTaskDetails) GetReferencesOk() (AnyOfobject, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferences

`func (o *PlannerTaskDetails) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### SetReferences

`func (o *PlannerTaskDetails) SetReferences(v AnyOfobject)`

SetReferences gets a reference to the given AnyOfobject and assigns it to the References field.

### SetReferencesExplicitNull

`func (o *PlannerTaskDetails) SetReferencesExplicitNull(b bool)`

SetReferencesExplicitNull (un)sets References to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The References value is set to nil even if false is passed
### GetChecklist

`func (o *PlannerTaskDetails) GetChecklist() AnyOfobject`

GetChecklist returns the Checklist field if non-nil, zero value otherwise.

### GetChecklistOk

`func (o *PlannerTaskDetails) GetChecklistOk() (AnyOfobject, bool)`

GetChecklistOk returns a tuple with the Checklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChecklist

`func (o *PlannerTaskDetails) HasChecklist() bool`

HasChecklist returns a boolean if a field has been set.

### SetChecklist

`func (o *PlannerTaskDetails) SetChecklist(v AnyOfobject)`

SetChecklist gets a reference to the given AnyOfobject and assigns it to the Checklist field.

### SetChecklistExplicitNull

`func (o *PlannerTaskDetails) SetChecklistExplicitNull(b bool)`

SetChecklistExplicitNull (un)sets Checklist to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Checklist value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


