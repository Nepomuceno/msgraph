# MicrosoftGraphSharedInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LastShared** | Pointer to [**AnyOfmicrosoftGraphSharingDetail**](anyOf&lt;microsoft.graph.sharingDetail&gt;.md) |  | [optional] 
**SharingHistory** | Pointer to [**[]AnyOfmicrosoftGraphSharingDetail**](anyOf&lt;microsoft.graph.sharingDetail&gt;.md) |  | [optional] 
**ResourceVisualization** | Pointer to [**AnyOfmicrosoftGraphResourceVisualization**](anyOf&lt;microsoft.graph.resourceVisualization&gt;.md) |  | [optional] 
**ResourceReference** | Pointer to [**AnyOfmicrosoftGraphResourceReference**](anyOf&lt;microsoft.graph.resourceReference&gt;.md) |  | [optional] 
**LastSharedMethod** | Pointer to [**AnyOfmicrosoftGraphEntity**](anyOf&lt;microsoft.graph.entity&gt;.md) |  | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphEntity**](anyOf&lt;microsoft.graph.entity&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphSharedInsight) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSharedInsight) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphSharedInsight) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphSharedInsight) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastShared

`func (o *MicrosoftGraphSharedInsight) GetLastShared() AnyOfmicrosoftGraphSharingDetail`

GetLastShared returns the LastShared field if non-nil, zero value otherwise.

### GetLastSharedOk

`func (o *MicrosoftGraphSharedInsight) GetLastSharedOk() (AnyOfmicrosoftGraphSharingDetail, bool)`

GetLastSharedOk returns a tuple with the LastShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastShared

`func (o *MicrosoftGraphSharedInsight) HasLastShared() bool`

HasLastShared returns a boolean if a field has been set.

### SetLastShared

`func (o *MicrosoftGraphSharedInsight) SetLastShared(v AnyOfmicrosoftGraphSharingDetail)`

SetLastShared gets a reference to the given AnyOfmicrosoftGraphSharingDetail and assigns it to the LastShared field.

### SetLastSharedExplicitNull

`func (o *MicrosoftGraphSharedInsight) SetLastSharedExplicitNull(b bool)`

SetLastSharedExplicitNull (un)sets LastShared to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastShared value is set to nil even if false is passed
### GetSharingHistory

`func (o *MicrosoftGraphSharedInsight) GetSharingHistory() []AnyOfmicrosoftGraphSharingDetail`

GetSharingHistory returns the SharingHistory field if non-nil, zero value otherwise.

### GetSharingHistoryOk

`func (o *MicrosoftGraphSharedInsight) GetSharingHistoryOk() ([]AnyOfmicrosoftGraphSharingDetail, bool)`

GetSharingHistoryOk returns a tuple with the SharingHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharingHistory

`func (o *MicrosoftGraphSharedInsight) HasSharingHistory() bool`

HasSharingHistory returns a boolean if a field has been set.

### SetSharingHistory

`func (o *MicrosoftGraphSharedInsight) SetSharingHistory(v []AnyOfmicrosoftGraphSharingDetail)`

SetSharingHistory gets a reference to the given []AnyOfmicrosoftGraphSharingDetail and assigns it to the SharingHistory field.

### GetResourceVisualization

`func (o *MicrosoftGraphSharedInsight) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization`

GetResourceVisualization returns the ResourceVisualization field if non-nil, zero value otherwise.

### GetResourceVisualizationOk

`func (o *MicrosoftGraphSharedInsight) GetResourceVisualizationOk() (AnyOfmicrosoftGraphResourceVisualization, bool)`

GetResourceVisualizationOk returns a tuple with the ResourceVisualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceVisualization

`func (o *MicrosoftGraphSharedInsight) HasResourceVisualization() bool`

HasResourceVisualization returns a boolean if a field has been set.

### SetResourceVisualization

`func (o *MicrosoftGraphSharedInsight) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization)`

SetResourceVisualization gets a reference to the given AnyOfmicrosoftGraphResourceVisualization and assigns it to the ResourceVisualization field.

### SetResourceVisualizationExplicitNull

`func (o *MicrosoftGraphSharedInsight) SetResourceVisualizationExplicitNull(b bool)`

SetResourceVisualizationExplicitNull (un)sets ResourceVisualization to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceVisualization value is set to nil even if false is passed
### GetResourceReference

`func (o *MicrosoftGraphSharedInsight) GetResourceReference() AnyOfmicrosoftGraphResourceReference`

GetResourceReference returns the ResourceReference field if non-nil, zero value otherwise.

### GetResourceReferenceOk

`func (o *MicrosoftGraphSharedInsight) GetResourceReferenceOk() (AnyOfmicrosoftGraphResourceReference, bool)`

GetResourceReferenceOk returns a tuple with the ResourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceReference

`func (o *MicrosoftGraphSharedInsight) HasResourceReference() bool`

HasResourceReference returns a boolean if a field has been set.

### SetResourceReference

`func (o *MicrosoftGraphSharedInsight) SetResourceReference(v AnyOfmicrosoftGraphResourceReference)`

SetResourceReference gets a reference to the given AnyOfmicrosoftGraphResourceReference and assigns it to the ResourceReference field.

### SetResourceReferenceExplicitNull

`func (o *MicrosoftGraphSharedInsight) SetResourceReferenceExplicitNull(b bool)`

SetResourceReferenceExplicitNull (un)sets ResourceReference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceReference value is set to nil even if false is passed
### GetLastSharedMethod

`func (o *MicrosoftGraphSharedInsight) GetLastSharedMethod() AnyOfmicrosoftGraphEntity`

GetLastSharedMethod returns the LastSharedMethod field if non-nil, zero value otherwise.

### GetLastSharedMethodOk

`func (o *MicrosoftGraphSharedInsight) GetLastSharedMethodOk() (AnyOfmicrosoftGraphEntity, bool)`

GetLastSharedMethodOk returns a tuple with the LastSharedMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSharedMethod

`func (o *MicrosoftGraphSharedInsight) HasLastSharedMethod() bool`

HasLastSharedMethod returns a boolean if a field has been set.

### SetLastSharedMethod

`func (o *MicrosoftGraphSharedInsight) SetLastSharedMethod(v AnyOfmicrosoftGraphEntity)`

SetLastSharedMethod gets a reference to the given AnyOfmicrosoftGraphEntity and assigns it to the LastSharedMethod field.

### SetLastSharedMethodExplicitNull

`func (o *MicrosoftGraphSharedInsight) SetLastSharedMethodExplicitNull(b bool)`

SetLastSharedMethodExplicitNull (un)sets LastSharedMethod to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastSharedMethod value is set to nil even if false is passed
### GetResource

`func (o *MicrosoftGraphSharedInsight) GetResource() AnyOfmicrosoftGraphEntity`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphSharedInsight) GetResourceOk() (AnyOfmicrosoftGraphEntity, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResource

`func (o *MicrosoftGraphSharedInsight) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResource

`func (o *MicrosoftGraphSharedInsight) SetResource(v AnyOfmicrosoftGraphEntity)`

SetResource gets a reference to the given AnyOfmicrosoftGraphEntity and assigns it to the Resource field.

### SetResourceExplicitNull

`func (o *MicrosoftGraphSharedInsight) SetResourceExplicitNull(b bool)`

SetResourceExplicitNull (un)sets Resource to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Resource value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


