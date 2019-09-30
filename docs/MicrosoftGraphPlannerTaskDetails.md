# MicrosoftGraphPlannerTaskDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PreviewType** | Pointer to [**AnyOfmicrosoftGraphPlannerPreviewType**](anyOf&lt;microsoft.graph.plannerPreviewType&gt;.md) |  | [optional] 
**References** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Checklist** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphPlannerTaskDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPlannerTaskDetails) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphPlannerTaskDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphPlannerTaskDetails) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDescription

`func (o *MicrosoftGraphPlannerTaskDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphPlannerTaskDetails) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphPlannerTaskDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphPlannerTaskDetails) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphPlannerTaskDetails) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPreviewType

`func (o *MicrosoftGraphPlannerTaskDetails) GetPreviewType() AnyOfmicrosoftGraphPlannerPreviewType`

GetPreviewType returns the PreviewType field if non-nil, zero value otherwise.

### GetPreviewTypeOk

`func (o *MicrosoftGraphPlannerTaskDetails) GetPreviewTypeOk() (AnyOfmicrosoftGraphPlannerPreviewType, bool)`

GetPreviewTypeOk returns a tuple with the PreviewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreviewType

`func (o *MicrosoftGraphPlannerTaskDetails) HasPreviewType() bool`

HasPreviewType returns a boolean if a field has been set.

### SetPreviewType

`func (o *MicrosoftGraphPlannerTaskDetails) SetPreviewType(v AnyOfmicrosoftGraphPlannerPreviewType)`

SetPreviewType gets a reference to the given AnyOfmicrosoftGraphPlannerPreviewType and assigns it to the PreviewType field.

### SetPreviewTypeExplicitNull

`func (o *MicrosoftGraphPlannerTaskDetails) SetPreviewTypeExplicitNull(b bool)`

SetPreviewTypeExplicitNull (un)sets PreviewType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreviewType value is set to nil even if false is passed
### GetReferences

`func (o *MicrosoftGraphPlannerTaskDetails) GetReferences() AnyOfobject`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *MicrosoftGraphPlannerTaskDetails) GetReferencesOk() (AnyOfobject, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferences

`func (o *MicrosoftGraphPlannerTaskDetails) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### SetReferences

`func (o *MicrosoftGraphPlannerTaskDetails) SetReferences(v AnyOfobject)`

SetReferences gets a reference to the given AnyOfobject and assigns it to the References field.

### SetReferencesExplicitNull

`func (o *MicrosoftGraphPlannerTaskDetails) SetReferencesExplicitNull(b bool)`

SetReferencesExplicitNull (un)sets References to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The References value is set to nil even if false is passed
### GetChecklist

`func (o *MicrosoftGraphPlannerTaskDetails) GetChecklist() AnyOfobject`

GetChecklist returns the Checklist field if non-nil, zero value otherwise.

### GetChecklistOk

`func (o *MicrosoftGraphPlannerTaskDetails) GetChecklistOk() (AnyOfobject, bool)`

GetChecklistOk returns a tuple with the Checklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChecklist

`func (o *MicrosoftGraphPlannerTaskDetails) HasChecklist() bool`

HasChecklist returns a boolean if a field has been set.

### SetChecklist

`func (o *MicrosoftGraphPlannerTaskDetails) SetChecklist(v AnyOfobject)`

SetChecklist gets a reference to the given AnyOfobject and assigns it to the Checklist field.

### SetChecklistExplicitNull

`func (o *MicrosoftGraphPlannerTaskDetails) SetChecklistExplicitNull(b bool)`

SetChecklistExplicitNull (un)sets Checklist to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Checklist value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


