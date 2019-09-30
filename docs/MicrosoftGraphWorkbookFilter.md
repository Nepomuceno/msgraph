# MicrosoftGraphWorkbookFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Criteria** | Pointer to [**AnyOfmicrosoftGraphWorkbookFilterCriteria**](anyOf&lt;microsoft.graph.workbookFilterCriteria&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWorkbookFilter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookFilter) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookFilter) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCriteria

`func (o *MicrosoftGraphWorkbookFilter) GetCriteria() AnyOfmicrosoftGraphWorkbookFilterCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *MicrosoftGraphWorkbookFilter) GetCriteriaOk() (AnyOfmicrosoftGraphWorkbookFilterCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCriteria

`func (o *MicrosoftGraphWorkbookFilter) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### SetCriteria

`func (o *MicrosoftGraphWorkbookFilter) SetCriteria(v AnyOfmicrosoftGraphWorkbookFilterCriteria)`

SetCriteria gets a reference to the given AnyOfmicrosoftGraphWorkbookFilterCriteria and assigns it to the Criteria field.

### SetCriteriaExplicitNull

`func (o *MicrosoftGraphWorkbookFilter) SetCriteriaExplicitNull(b bool)`

SetCriteriaExplicitNull (un)sets Criteria to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Criteria value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


