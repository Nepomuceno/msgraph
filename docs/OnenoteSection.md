# OnenoteSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphSectionLinks**](anyOf&lt;microsoft.graph.sectionLinks&gt;.md) |  | [optional] 
**PagesUrl** | Pointer to **string** |  | [optional] 
**ParentNotebook** | Pointer to [**AnyOfmicrosoftGraphNotebook**](anyOf&lt;microsoft.graph.notebook&gt;.md) |  | [optional] 
**ParentSectionGroup** | Pointer to [**AnyOfmicrosoftGraphSectionGroup**](anyOf&lt;microsoft.graph.sectionGroup&gt;.md) |  | [optional] 
**Pages** | Pointer to [**[]MicrosoftGraphOnenotePage**](microsoft.graph.onenotePage.md) |  | [optional] 

## Methods

### GetIsDefault

`func (o *OnenoteSection) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *OnenoteSection) GetIsDefaultOk() (bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDefault

`func (o *OnenoteSection) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefault

`func (o *OnenoteSection) SetIsDefault(v bool)`

SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.

### SetIsDefaultExplicitNull

`func (o *OnenoteSection) SetIsDefaultExplicitNull(b bool)`

SetIsDefaultExplicitNull (un)sets IsDefault to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsDefault value is set to nil even if false is passed
### GetLinks

`func (o *OnenoteSection) GetLinks() AnyOfmicrosoftGraphSectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OnenoteSection) GetLinksOk() (AnyOfmicrosoftGraphSectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinks

`func (o *OnenoteSection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinks

`func (o *OnenoteSection) SetLinks(v AnyOfmicrosoftGraphSectionLinks)`

SetLinks gets a reference to the given AnyOfmicrosoftGraphSectionLinks and assigns it to the Links field.

### SetLinksExplicitNull

`func (o *OnenoteSection) SetLinksExplicitNull(b bool)`

SetLinksExplicitNull (un)sets Links to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Links value is set to nil even if false is passed
### GetPagesUrl

`func (o *OnenoteSection) GetPagesUrl() string`

GetPagesUrl returns the PagesUrl field if non-nil, zero value otherwise.

### GetPagesUrlOk

`func (o *OnenoteSection) GetPagesUrlOk() (string, bool)`

GetPagesUrlOk returns a tuple with the PagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPagesUrl

`func (o *OnenoteSection) HasPagesUrl() bool`

HasPagesUrl returns a boolean if a field has been set.

### SetPagesUrl

`func (o *OnenoteSection) SetPagesUrl(v string)`

SetPagesUrl gets a reference to the given string and assigns it to the PagesUrl field.

### SetPagesUrlExplicitNull

`func (o *OnenoteSection) SetPagesUrlExplicitNull(b bool)`

SetPagesUrlExplicitNull (un)sets PagesUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PagesUrl value is set to nil even if false is passed
### GetParentNotebook

`func (o *OnenoteSection) GetParentNotebook() AnyOfmicrosoftGraphNotebook`

GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.

### GetParentNotebookOk

`func (o *OnenoteSection) GetParentNotebookOk() (AnyOfmicrosoftGraphNotebook, bool)`

GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentNotebook

`func (o *OnenoteSection) HasParentNotebook() bool`

HasParentNotebook returns a boolean if a field has been set.

### SetParentNotebook

`func (o *OnenoteSection) SetParentNotebook(v AnyOfmicrosoftGraphNotebook)`

SetParentNotebook gets a reference to the given AnyOfmicrosoftGraphNotebook and assigns it to the ParentNotebook field.

### SetParentNotebookExplicitNull

`func (o *OnenoteSection) SetParentNotebookExplicitNull(b bool)`

SetParentNotebookExplicitNull (un)sets ParentNotebook to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentNotebook value is set to nil even if false is passed
### GetParentSectionGroup

`func (o *OnenoteSection) GetParentSectionGroup() AnyOfmicrosoftGraphSectionGroup`

GetParentSectionGroup returns the ParentSectionGroup field if non-nil, zero value otherwise.

### GetParentSectionGroupOk

`func (o *OnenoteSection) GetParentSectionGroupOk() (AnyOfmicrosoftGraphSectionGroup, bool)`

GetParentSectionGroupOk returns a tuple with the ParentSectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentSectionGroup

`func (o *OnenoteSection) HasParentSectionGroup() bool`

HasParentSectionGroup returns a boolean if a field has been set.

### SetParentSectionGroup

`func (o *OnenoteSection) SetParentSectionGroup(v AnyOfmicrosoftGraphSectionGroup)`

SetParentSectionGroup gets a reference to the given AnyOfmicrosoftGraphSectionGroup and assigns it to the ParentSectionGroup field.

### SetParentSectionGroupExplicitNull

`func (o *OnenoteSection) SetParentSectionGroupExplicitNull(b bool)`

SetParentSectionGroupExplicitNull (un)sets ParentSectionGroup to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentSectionGroup value is set to nil even if false is passed
### GetPages

`func (o *OnenoteSection) GetPages() []MicrosoftGraphOnenotePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *OnenoteSection) GetPagesOk() ([]MicrosoftGraphOnenotePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPages

`func (o *OnenoteSection) HasPages() bool`

HasPages returns a boolean if a field has been set.

### SetPages

`func (o *OnenoteSection) SetPages(v []MicrosoftGraphOnenotePage)`

SetPages gets a reference to the given []MicrosoftGraphOnenotePage and assigns it to the Pages field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


