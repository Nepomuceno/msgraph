# MicrosoftGraphUsedInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LastUsed** | Pointer to [**AnyOfmicrosoftGraphUsageDetails**](anyOf&lt;microsoft.graph.usageDetails&gt;.md) |  | [optional] 
**ResourceVisualization** | Pointer to [**AnyOfmicrosoftGraphResourceVisualization**](anyOf&lt;microsoft.graph.resourceVisualization&gt;.md) |  | [optional] 
**ResourceReference** | Pointer to [**AnyOfmicrosoftGraphResourceReference**](anyOf&lt;microsoft.graph.resourceReference&gt;.md) |  | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphEntity**](anyOf&lt;microsoft.graph.entity&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphUsedInsight) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUsedInsight) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphUsedInsight) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphUsedInsight) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastUsed

`func (o *MicrosoftGraphUsedInsight) GetLastUsed() AnyOfmicrosoftGraphUsageDetails`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *MicrosoftGraphUsedInsight) GetLastUsedOk() (AnyOfmicrosoftGraphUsageDetails, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUsed

`func (o *MicrosoftGraphUsedInsight) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsed

`func (o *MicrosoftGraphUsedInsight) SetLastUsed(v AnyOfmicrosoftGraphUsageDetails)`

SetLastUsed gets a reference to the given AnyOfmicrosoftGraphUsageDetails and assigns it to the LastUsed field.

### SetLastUsedExplicitNull

`func (o *MicrosoftGraphUsedInsight) SetLastUsedExplicitNull(b bool)`

SetLastUsedExplicitNull (un)sets LastUsed to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastUsed value is set to nil even if false is passed
### GetResourceVisualization

`func (o *MicrosoftGraphUsedInsight) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization`

GetResourceVisualization returns the ResourceVisualization field if non-nil, zero value otherwise.

### GetResourceVisualizationOk

`func (o *MicrosoftGraphUsedInsight) GetResourceVisualizationOk() (AnyOfmicrosoftGraphResourceVisualization, bool)`

GetResourceVisualizationOk returns a tuple with the ResourceVisualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceVisualization

`func (o *MicrosoftGraphUsedInsight) HasResourceVisualization() bool`

HasResourceVisualization returns a boolean if a field has been set.

### SetResourceVisualization

`func (o *MicrosoftGraphUsedInsight) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization)`

SetResourceVisualization gets a reference to the given AnyOfmicrosoftGraphResourceVisualization and assigns it to the ResourceVisualization field.

### SetResourceVisualizationExplicitNull

`func (o *MicrosoftGraphUsedInsight) SetResourceVisualizationExplicitNull(b bool)`

SetResourceVisualizationExplicitNull (un)sets ResourceVisualization to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceVisualization value is set to nil even if false is passed
### GetResourceReference

`func (o *MicrosoftGraphUsedInsight) GetResourceReference() AnyOfmicrosoftGraphResourceReference`

GetResourceReference returns the ResourceReference field if non-nil, zero value otherwise.

### GetResourceReferenceOk

`func (o *MicrosoftGraphUsedInsight) GetResourceReferenceOk() (AnyOfmicrosoftGraphResourceReference, bool)`

GetResourceReferenceOk returns a tuple with the ResourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceReference

`func (o *MicrosoftGraphUsedInsight) HasResourceReference() bool`

HasResourceReference returns a boolean if a field has been set.

### SetResourceReference

`func (o *MicrosoftGraphUsedInsight) SetResourceReference(v AnyOfmicrosoftGraphResourceReference)`

SetResourceReference gets a reference to the given AnyOfmicrosoftGraphResourceReference and assigns it to the ResourceReference field.

### SetResourceReferenceExplicitNull

`func (o *MicrosoftGraphUsedInsight) SetResourceReferenceExplicitNull(b bool)`

SetResourceReferenceExplicitNull (un)sets ResourceReference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceReference value is set to nil even if false is passed
### GetResource

`func (o *MicrosoftGraphUsedInsight) GetResource() AnyOfmicrosoftGraphEntity`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphUsedInsight) GetResourceOk() (AnyOfmicrosoftGraphEntity, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResource

`func (o *MicrosoftGraphUsedInsight) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResource

`func (o *MicrosoftGraphUsedInsight) SetResource(v AnyOfmicrosoftGraphEntity)`

SetResource gets a reference to the given AnyOfmicrosoftGraphEntity and assigns it to the Resource field.

### SetResourceExplicitNull

`func (o *MicrosoftGraphUsedInsight) SetResourceExplicitNull(b bool)`

SetResourceExplicitNull (un)sets Resource to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Resource value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


