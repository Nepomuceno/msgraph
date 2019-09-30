# MicrosoftGraphScoredEmailAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**RelevanceScore** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**SelectionLikelihood** | Pointer to [**AnyOfmicrosoftGraphSelectionLikelihoodInfo**](anyOf&lt;microsoft.graph.selectionLikelihoodInfo&gt;.md) |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 

## Methods

### GetAddress

`func (o *MicrosoftGraphScoredEmailAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphScoredEmailAddress) GetAddressOk() (string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddress

`func (o *MicrosoftGraphScoredEmailAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddress

`func (o *MicrosoftGraphScoredEmailAddress) SetAddress(v string)`

SetAddress gets a reference to the given string and assigns it to the Address field.

### SetAddressExplicitNull

`func (o *MicrosoftGraphScoredEmailAddress) SetAddressExplicitNull(b bool)`

SetAddressExplicitNull (un)sets Address to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Address value is set to nil even if false is passed
### GetRelevanceScore

`func (o *MicrosoftGraphScoredEmailAddress) GetRelevanceScore() AnyOfnumberstringstring`

GetRelevanceScore returns the RelevanceScore field if non-nil, zero value otherwise.

### GetRelevanceScoreOk

`func (o *MicrosoftGraphScoredEmailAddress) GetRelevanceScoreOk() (AnyOfnumberstringstring, bool)`

GetRelevanceScoreOk returns a tuple with the RelevanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRelevanceScore

`func (o *MicrosoftGraphScoredEmailAddress) HasRelevanceScore() bool`

HasRelevanceScore returns a boolean if a field has been set.

### SetRelevanceScore

`func (o *MicrosoftGraphScoredEmailAddress) SetRelevanceScore(v AnyOfnumberstringstring)`

SetRelevanceScore gets a reference to the given AnyOfnumberstringstring and assigns it to the RelevanceScore field.

### SetRelevanceScoreExplicitNull

`func (o *MicrosoftGraphScoredEmailAddress) SetRelevanceScoreExplicitNull(b bool)`

SetRelevanceScoreExplicitNull (un)sets RelevanceScore to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RelevanceScore value is set to nil even if false is passed
### GetSelectionLikelihood

`func (o *MicrosoftGraphScoredEmailAddress) GetSelectionLikelihood() AnyOfmicrosoftGraphSelectionLikelihoodInfo`

GetSelectionLikelihood returns the SelectionLikelihood field if non-nil, zero value otherwise.

### GetSelectionLikelihoodOk

`func (o *MicrosoftGraphScoredEmailAddress) GetSelectionLikelihoodOk() (AnyOfmicrosoftGraphSelectionLikelihoodInfo, bool)`

GetSelectionLikelihoodOk returns a tuple with the SelectionLikelihood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSelectionLikelihood

`func (o *MicrosoftGraphScoredEmailAddress) HasSelectionLikelihood() bool`

HasSelectionLikelihood returns a boolean if a field has been set.

### SetSelectionLikelihood

`func (o *MicrosoftGraphScoredEmailAddress) SetSelectionLikelihood(v AnyOfmicrosoftGraphSelectionLikelihoodInfo)`

SetSelectionLikelihood gets a reference to the given AnyOfmicrosoftGraphSelectionLikelihoodInfo and assigns it to the SelectionLikelihood field.

### SetSelectionLikelihoodExplicitNull

`func (o *MicrosoftGraphScoredEmailAddress) SetSelectionLikelihoodExplicitNull(b bool)`

SetSelectionLikelihoodExplicitNull (un)sets SelectionLikelihood to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SelectionLikelihood value is set to nil even if false is passed
### GetItemId

`func (o *MicrosoftGraphScoredEmailAddress) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *MicrosoftGraphScoredEmailAddress) GetItemIdOk() (string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItemId

`func (o *MicrosoftGraphScoredEmailAddress) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemId

`func (o *MicrosoftGraphScoredEmailAddress) SetItemId(v string)`

SetItemId gets a reference to the given string and assigns it to the ItemId field.

### SetItemIdExplicitNull

`func (o *MicrosoftGraphScoredEmailAddress) SetItemIdExplicitNull(b bool)`

SetItemIdExplicitNull (un)sets ItemId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ItemId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


