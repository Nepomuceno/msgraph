# MicrosoftGraphTrending

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**ResourceVisualization** | Pointer to [**AnyOfmicrosoftGraphResourceVisualization**](anyOf&lt;microsoft.graph.resourceVisualization&gt;.md) |  | [optional] 
**ResourceReference** | Pointer to [**AnyOfmicrosoftGraphResourceReference**](anyOf&lt;microsoft.graph.resourceReference&gt;.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphEntity**](anyOf&lt;microsoft.graph.entity&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphTrending) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTrending) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphTrending) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphTrending) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetWeight

`func (o *MicrosoftGraphTrending) GetWeight() AnyOfnumberstringstring`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *MicrosoftGraphTrending) GetWeightOk() (AnyOfnumberstringstring, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWeight

`func (o *MicrosoftGraphTrending) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeight

`func (o *MicrosoftGraphTrending) SetWeight(v AnyOfnumberstringstring)`

SetWeight gets a reference to the given AnyOfnumberstringstring and assigns it to the Weight field.

### GetResourceVisualization

`func (o *MicrosoftGraphTrending) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization`

GetResourceVisualization returns the ResourceVisualization field if non-nil, zero value otherwise.

### GetResourceVisualizationOk

`func (o *MicrosoftGraphTrending) GetResourceVisualizationOk() (AnyOfmicrosoftGraphResourceVisualization, bool)`

GetResourceVisualizationOk returns a tuple with the ResourceVisualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceVisualization

`func (o *MicrosoftGraphTrending) HasResourceVisualization() bool`

HasResourceVisualization returns a boolean if a field has been set.

### SetResourceVisualization

`func (o *MicrosoftGraphTrending) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization)`

SetResourceVisualization gets a reference to the given AnyOfmicrosoftGraphResourceVisualization and assigns it to the ResourceVisualization field.

### SetResourceVisualizationExplicitNull

`func (o *MicrosoftGraphTrending) SetResourceVisualizationExplicitNull(b bool)`

SetResourceVisualizationExplicitNull (un)sets ResourceVisualization to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceVisualization value is set to nil even if false is passed
### GetResourceReference

`func (o *MicrosoftGraphTrending) GetResourceReference() AnyOfmicrosoftGraphResourceReference`

GetResourceReference returns the ResourceReference field if non-nil, zero value otherwise.

### GetResourceReferenceOk

`func (o *MicrosoftGraphTrending) GetResourceReferenceOk() (AnyOfmicrosoftGraphResourceReference, bool)`

GetResourceReferenceOk returns a tuple with the ResourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceReference

`func (o *MicrosoftGraphTrending) HasResourceReference() bool`

HasResourceReference returns a boolean if a field has been set.

### SetResourceReference

`func (o *MicrosoftGraphTrending) SetResourceReference(v AnyOfmicrosoftGraphResourceReference)`

SetResourceReference gets a reference to the given AnyOfmicrosoftGraphResourceReference and assigns it to the ResourceReference field.

### SetResourceReferenceExplicitNull

`func (o *MicrosoftGraphTrending) SetResourceReferenceExplicitNull(b bool)`

SetResourceReferenceExplicitNull (un)sets ResourceReference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceReference value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphTrending) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphTrending) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphTrending) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphTrending) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphTrending) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetResource

`func (o *MicrosoftGraphTrending) GetResource() AnyOfmicrosoftGraphEntity`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphTrending) GetResourceOk() (AnyOfmicrosoftGraphEntity, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResource

`func (o *MicrosoftGraphTrending) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResource

`func (o *MicrosoftGraphTrending) SetResource(v AnyOfmicrosoftGraphEntity)`

SetResource gets a reference to the given AnyOfmicrosoftGraphEntity and assigns it to the Resource field.

### SetResourceExplicitNull

`func (o *MicrosoftGraphTrending) SetResourceExplicitNull(b bool)`

SetResourceExplicitNull (un)sets Resource to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Resource value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


