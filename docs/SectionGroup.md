# SectionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionsUrl** | Pointer to **string** |  | [optional] 
**SectionGroupsUrl** | Pointer to **string** |  | [optional] 
**ParentNotebook** | Pointer to [**AnyOfmicrosoftGraphNotebook**](anyOf&lt;microsoft.graph.notebook&gt;.md) |  | [optional] 
**ParentSectionGroup** | Pointer to [**AnyOfmicrosoftGraphSectionGroup**](anyOf&lt;microsoft.graph.sectionGroup&gt;.md) |  | [optional] 
**Sections** | Pointer to [**[]MicrosoftGraphOnenoteSection**](microsoft.graph.onenoteSection.md) |  | [optional] 
**SectionGroups** | Pointer to [**[]MicrosoftGraphSectionGroup**](microsoft.graph.sectionGroup.md) |  | [optional] 

## Methods

### GetSectionsUrl

`func (o *SectionGroup) GetSectionsUrl() string`

GetSectionsUrl returns the SectionsUrl field if non-nil, zero value otherwise.

### GetSectionsUrlOk

`func (o *SectionGroup) GetSectionsUrlOk() (string, bool)`

GetSectionsUrlOk returns a tuple with the SectionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionsUrl

`func (o *SectionGroup) HasSectionsUrl() bool`

HasSectionsUrl returns a boolean if a field has been set.

### SetSectionsUrl

`func (o *SectionGroup) SetSectionsUrl(v string)`

SetSectionsUrl gets a reference to the given string and assigns it to the SectionsUrl field.

### SetSectionsUrlExplicitNull

`func (o *SectionGroup) SetSectionsUrlExplicitNull(b bool)`

SetSectionsUrlExplicitNull (un)sets SectionsUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SectionsUrl value is set to nil even if false is passed
### GetSectionGroupsUrl

`func (o *SectionGroup) GetSectionGroupsUrl() string`

GetSectionGroupsUrl returns the SectionGroupsUrl field if non-nil, zero value otherwise.

### GetSectionGroupsUrlOk

`func (o *SectionGroup) GetSectionGroupsUrlOk() (string, bool)`

GetSectionGroupsUrlOk returns a tuple with the SectionGroupsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionGroupsUrl

`func (o *SectionGroup) HasSectionGroupsUrl() bool`

HasSectionGroupsUrl returns a boolean if a field has been set.

### SetSectionGroupsUrl

`func (o *SectionGroup) SetSectionGroupsUrl(v string)`

SetSectionGroupsUrl gets a reference to the given string and assigns it to the SectionGroupsUrl field.

### SetSectionGroupsUrlExplicitNull

`func (o *SectionGroup) SetSectionGroupsUrlExplicitNull(b bool)`

SetSectionGroupsUrlExplicitNull (un)sets SectionGroupsUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SectionGroupsUrl value is set to nil even if false is passed
### GetParentNotebook

`func (o *SectionGroup) GetParentNotebook() AnyOfmicrosoftGraphNotebook`

GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.

### GetParentNotebookOk

`func (o *SectionGroup) GetParentNotebookOk() (AnyOfmicrosoftGraphNotebook, bool)`

GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentNotebook

`func (o *SectionGroup) HasParentNotebook() bool`

HasParentNotebook returns a boolean if a field has been set.

### SetParentNotebook

`func (o *SectionGroup) SetParentNotebook(v AnyOfmicrosoftGraphNotebook)`

SetParentNotebook gets a reference to the given AnyOfmicrosoftGraphNotebook and assigns it to the ParentNotebook field.

### SetParentNotebookExplicitNull

`func (o *SectionGroup) SetParentNotebookExplicitNull(b bool)`

SetParentNotebookExplicitNull (un)sets ParentNotebook to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentNotebook value is set to nil even if false is passed
### GetParentSectionGroup

`func (o *SectionGroup) GetParentSectionGroup() AnyOfmicrosoftGraphSectionGroup`

GetParentSectionGroup returns the ParentSectionGroup field if non-nil, zero value otherwise.

### GetParentSectionGroupOk

`func (o *SectionGroup) GetParentSectionGroupOk() (AnyOfmicrosoftGraphSectionGroup, bool)`

GetParentSectionGroupOk returns a tuple with the ParentSectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentSectionGroup

`func (o *SectionGroup) HasParentSectionGroup() bool`

HasParentSectionGroup returns a boolean if a field has been set.

### SetParentSectionGroup

`func (o *SectionGroup) SetParentSectionGroup(v AnyOfmicrosoftGraphSectionGroup)`

SetParentSectionGroup gets a reference to the given AnyOfmicrosoftGraphSectionGroup and assigns it to the ParentSectionGroup field.

### SetParentSectionGroupExplicitNull

`func (o *SectionGroup) SetParentSectionGroupExplicitNull(b bool)`

SetParentSectionGroupExplicitNull (un)sets ParentSectionGroup to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentSectionGroup value is set to nil even if false is passed
### GetSections

`func (o *SectionGroup) GetSections() []MicrosoftGraphOnenoteSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SectionGroup) GetSectionsOk() ([]MicrosoftGraphOnenoteSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSections

`func (o *SectionGroup) HasSections() bool`

HasSections returns a boolean if a field has been set.

### SetSections

`func (o *SectionGroup) SetSections(v []MicrosoftGraphOnenoteSection)`

SetSections gets a reference to the given []MicrosoftGraphOnenoteSection and assigns it to the Sections field.

### GetSectionGroups

`func (o *SectionGroup) GetSectionGroups() []MicrosoftGraphSectionGroup`

GetSectionGroups returns the SectionGroups field if non-nil, zero value otherwise.

### GetSectionGroupsOk

`func (o *SectionGroup) GetSectionGroupsOk() ([]MicrosoftGraphSectionGroup, bool)`

GetSectionGroupsOk returns a tuple with the SectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionGroups

`func (o *SectionGroup) HasSectionGroups() bool`

HasSectionGroups returns a boolean if a field has been set.

### SetSectionGroups

`func (o *SectionGroup) SetSectionGroups(v []MicrosoftGraphSectionGroup)`

SetSectionGroups gets a reference to the given []MicrosoftGraphSectionGroup and assigns it to the SectionGroups field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


