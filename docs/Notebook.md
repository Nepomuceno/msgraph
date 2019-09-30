# Notebook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **bool** |  | [optional] 
**UserRole** | Pointer to [**AnyOfmicrosoftGraphOnenoteUserRole**](anyOf&lt;microsoft.graph.onenoteUserRole&gt;.md) |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**SectionsUrl** | Pointer to **string** |  | [optional] 
**SectionGroupsUrl** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphNotebookLinks**](anyOf&lt;microsoft.graph.notebookLinks&gt;.md) |  | [optional] 
**Sections** | Pointer to [**[]MicrosoftGraphOnenoteSection**](microsoft.graph.onenoteSection.md) |  | [optional] 
**SectionGroups** | Pointer to [**[]MicrosoftGraphSectionGroup**](microsoft.graph.sectionGroup.md) |  | [optional] 

## Methods

### GetIsDefault

`func (o *Notebook) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Notebook) GetIsDefaultOk() (bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDefault

`func (o *Notebook) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefault

`func (o *Notebook) SetIsDefault(v bool)`

SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.

### SetIsDefaultExplicitNull

`func (o *Notebook) SetIsDefaultExplicitNull(b bool)`

SetIsDefaultExplicitNull (un)sets IsDefault to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsDefault value is set to nil even if false is passed
### GetUserRole

`func (o *Notebook) GetUserRole() AnyOfmicrosoftGraphOnenoteUserRole`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *Notebook) GetUserRoleOk() (AnyOfmicrosoftGraphOnenoteUserRole, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserRole

`func (o *Notebook) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### SetUserRole

`func (o *Notebook) SetUserRole(v AnyOfmicrosoftGraphOnenoteUserRole)`

SetUserRole gets a reference to the given AnyOfmicrosoftGraphOnenoteUserRole and assigns it to the UserRole field.

### SetUserRoleExplicitNull

`func (o *Notebook) SetUserRoleExplicitNull(b bool)`

SetUserRoleExplicitNull (un)sets UserRole to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserRole value is set to nil even if false is passed
### GetIsShared

`func (o *Notebook) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *Notebook) GetIsSharedOk() (bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsShared

`func (o *Notebook) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### SetIsShared

`func (o *Notebook) SetIsShared(v bool)`

SetIsShared gets a reference to the given bool and assigns it to the IsShared field.

### SetIsSharedExplicitNull

`func (o *Notebook) SetIsSharedExplicitNull(b bool)`

SetIsSharedExplicitNull (un)sets IsShared to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsShared value is set to nil even if false is passed
### GetSectionsUrl

`func (o *Notebook) GetSectionsUrl() string`

GetSectionsUrl returns the SectionsUrl field if non-nil, zero value otherwise.

### GetSectionsUrlOk

`func (o *Notebook) GetSectionsUrlOk() (string, bool)`

GetSectionsUrlOk returns a tuple with the SectionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionsUrl

`func (o *Notebook) HasSectionsUrl() bool`

HasSectionsUrl returns a boolean if a field has been set.

### SetSectionsUrl

`func (o *Notebook) SetSectionsUrl(v string)`

SetSectionsUrl gets a reference to the given string and assigns it to the SectionsUrl field.

### SetSectionsUrlExplicitNull

`func (o *Notebook) SetSectionsUrlExplicitNull(b bool)`

SetSectionsUrlExplicitNull (un)sets SectionsUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SectionsUrl value is set to nil even if false is passed
### GetSectionGroupsUrl

`func (o *Notebook) GetSectionGroupsUrl() string`

GetSectionGroupsUrl returns the SectionGroupsUrl field if non-nil, zero value otherwise.

### GetSectionGroupsUrlOk

`func (o *Notebook) GetSectionGroupsUrlOk() (string, bool)`

GetSectionGroupsUrlOk returns a tuple with the SectionGroupsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionGroupsUrl

`func (o *Notebook) HasSectionGroupsUrl() bool`

HasSectionGroupsUrl returns a boolean if a field has been set.

### SetSectionGroupsUrl

`func (o *Notebook) SetSectionGroupsUrl(v string)`

SetSectionGroupsUrl gets a reference to the given string and assigns it to the SectionGroupsUrl field.

### SetSectionGroupsUrlExplicitNull

`func (o *Notebook) SetSectionGroupsUrlExplicitNull(b bool)`

SetSectionGroupsUrlExplicitNull (un)sets SectionGroupsUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SectionGroupsUrl value is set to nil even if false is passed
### GetLinks

`func (o *Notebook) GetLinks() AnyOfmicrosoftGraphNotebookLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Notebook) GetLinksOk() (AnyOfmicrosoftGraphNotebookLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinks

`func (o *Notebook) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinks

`func (o *Notebook) SetLinks(v AnyOfmicrosoftGraphNotebookLinks)`

SetLinks gets a reference to the given AnyOfmicrosoftGraphNotebookLinks and assigns it to the Links field.

### SetLinksExplicitNull

`func (o *Notebook) SetLinksExplicitNull(b bool)`

SetLinksExplicitNull (un)sets Links to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Links value is set to nil even if false is passed
### GetSections

`func (o *Notebook) GetSections() []MicrosoftGraphOnenoteSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *Notebook) GetSectionsOk() ([]MicrosoftGraphOnenoteSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSections

`func (o *Notebook) HasSections() bool`

HasSections returns a boolean if a field has been set.

### SetSections

`func (o *Notebook) SetSections(v []MicrosoftGraphOnenoteSection)`

SetSections gets a reference to the given []MicrosoftGraphOnenoteSection and assigns it to the Sections field.

### GetSectionGroups

`func (o *Notebook) GetSectionGroups() []MicrosoftGraphSectionGroup`

GetSectionGroups returns the SectionGroups field if non-nil, zero value otherwise.

### GetSectionGroupsOk

`func (o *Notebook) GetSectionGroupsOk() ([]MicrosoftGraphSectionGroup, bool)`

GetSectionGroupsOk returns a tuple with the SectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionGroups

`func (o *Notebook) HasSectionGroups() bool`

HasSectionGroups returns a boolean if a field has been set.

### SetSectionGroups

`func (o *Notebook) SetSectionGroups(v []MicrosoftGraphSectionGroup)`

SetSectionGroups gets a reference to the given []MicrosoftGraphSectionGroup and assigns it to the SectionGroups field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


