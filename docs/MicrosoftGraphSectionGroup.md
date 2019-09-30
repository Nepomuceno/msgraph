# MicrosoftGraphSectionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Self** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**SectionsUrl** | Pointer to **string** |  | [optional] 
**SectionGroupsUrl** | Pointer to **string** |  | [optional] 
**ParentNotebook** | Pointer to [**AnyOfmicrosoftGraphNotebook**](anyOf&lt;microsoft.graph.notebook&gt;.md) |  | [optional] 
**ParentSectionGroup** | Pointer to [**AnyOfmicrosoftGraphSectionGroup**](anyOf&lt;microsoft.graph.sectionGroup&gt;.md) |  | [optional] 
**Sections** | Pointer to [**[]MicrosoftGraphOnenoteSection**](microsoft.graph.onenoteSection.md) |  | [optional] 
**SectionGroups** | Pointer to [**[]MicrosoftGraphSectionGroup**](microsoft.graph.sectionGroup.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphSectionGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSectionGroup) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphSectionGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphSectionGroup) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetSelf

`func (o *MicrosoftGraphSectionGroup) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MicrosoftGraphSectionGroup) GetSelfOk() (string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSelf

`func (o *MicrosoftGraphSectionGroup) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### SetSelf

`func (o *MicrosoftGraphSectionGroup) SetSelf(v string)`

SetSelf gets a reference to the given string and assigns it to the Self field.

### SetSelfExplicitNull

`func (o *MicrosoftGraphSectionGroup) SetSelfExplicitNull(b bool)`

SetSelfExplicitNull (un)sets Self to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Self value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphSectionGroup) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphSectionGroup) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphSectionGroup) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphSectionGroup) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphSectionGroup) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphSectionGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphSectionGroup) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphSectionGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphSectionGroup) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphSectionGroup) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetCreatedBy

`func (o *MicrosoftGraphSectionGroup) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphSectionGroup) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphSectionGroup) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphSectionGroup) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphSectionGroup) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetLastModifiedBy

`func (o *MicrosoftGraphSectionGroup) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphSectionGroup) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphSectionGroup) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphSectionGroup) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphSectionGroup) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphSectionGroup) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphSectionGroup) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphSectionGroup) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphSectionGroup) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphSectionGroup) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetSectionsUrl

`func (o *MicrosoftGraphSectionGroup) GetSectionsUrl() string`

GetSectionsUrl returns the SectionsUrl field if non-nil, zero value otherwise.

### GetSectionsUrlOk

`func (o *MicrosoftGraphSectionGroup) GetSectionsUrlOk() (string, bool)`

GetSectionsUrlOk returns a tuple with the SectionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionsUrl

`func (o *MicrosoftGraphSectionGroup) HasSectionsUrl() bool`

HasSectionsUrl returns a boolean if a field has been set.

### SetSectionsUrl

`func (o *MicrosoftGraphSectionGroup) SetSectionsUrl(v string)`

SetSectionsUrl gets a reference to the given string and assigns it to the SectionsUrl field.

### SetSectionsUrlExplicitNull

`func (o *MicrosoftGraphSectionGroup) SetSectionsUrlExplicitNull(b bool)`

SetSectionsUrlExplicitNull (un)sets SectionsUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SectionsUrl value is set to nil even if false is passed
### GetSectionGroupsUrl

`func (o *MicrosoftGraphSectionGroup) GetSectionGroupsUrl() string`

GetSectionGroupsUrl returns the SectionGroupsUrl field if non-nil, zero value otherwise.

### GetSectionGroupsUrlOk

`func (o *MicrosoftGraphSectionGroup) GetSectionGroupsUrlOk() (string, bool)`

GetSectionGroupsUrlOk returns a tuple with the SectionGroupsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionGroupsUrl

`func (o *MicrosoftGraphSectionGroup) HasSectionGroupsUrl() bool`

HasSectionGroupsUrl returns a boolean if a field has been set.

### SetSectionGroupsUrl

`func (o *MicrosoftGraphSectionGroup) SetSectionGroupsUrl(v string)`

SetSectionGroupsUrl gets a reference to the given string and assigns it to the SectionGroupsUrl field.

### SetSectionGroupsUrlExplicitNull

`func (o *MicrosoftGraphSectionGroup) SetSectionGroupsUrlExplicitNull(b bool)`

SetSectionGroupsUrlExplicitNull (un)sets SectionGroupsUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SectionGroupsUrl value is set to nil even if false is passed
### GetParentNotebook

`func (o *MicrosoftGraphSectionGroup) GetParentNotebook() AnyOfmicrosoftGraphNotebook`

GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.

### GetParentNotebookOk

`func (o *MicrosoftGraphSectionGroup) GetParentNotebookOk() (AnyOfmicrosoftGraphNotebook, bool)`

GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentNotebook

`func (o *MicrosoftGraphSectionGroup) HasParentNotebook() bool`

HasParentNotebook returns a boolean if a field has been set.

### SetParentNotebook

`func (o *MicrosoftGraphSectionGroup) SetParentNotebook(v AnyOfmicrosoftGraphNotebook)`

SetParentNotebook gets a reference to the given AnyOfmicrosoftGraphNotebook and assigns it to the ParentNotebook field.

### SetParentNotebookExplicitNull

`func (o *MicrosoftGraphSectionGroup) SetParentNotebookExplicitNull(b bool)`

SetParentNotebookExplicitNull (un)sets ParentNotebook to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentNotebook value is set to nil even if false is passed
### GetParentSectionGroup

`func (o *MicrosoftGraphSectionGroup) GetParentSectionGroup() AnyOfmicrosoftGraphSectionGroup`

GetParentSectionGroup returns the ParentSectionGroup field if non-nil, zero value otherwise.

### GetParentSectionGroupOk

`func (o *MicrosoftGraphSectionGroup) GetParentSectionGroupOk() (AnyOfmicrosoftGraphSectionGroup, bool)`

GetParentSectionGroupOk returns a tuple with the ParentSectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentSectionGroup

`func (o *MicrosoftGraphSectionGroup) HasParentSectionGroup() bool`

HasParentSectionGroup returns a boolean if a field has been set.

### SetParentSectionGroup

`func (o *MicrosoftGraphSectionGroup) SetParentSectionGroup(v AnyOfmicrosoftGraphSectionGroup)`

SetParentSectionGroup gets a reference to the given AnyOfmicrosoftGraphSectionGroup and assigns it to the ParentSectionGroup field.

### SetParentSectionGroupExplicitNull

`func (o *MicrosoftGraphSectionGroup) SetParentSectionGroupExplicitNull(b bool)`

SetParentSectionGroupExplicitNull (un)sets ParentSectionGroup to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentSectionGroup value is set to nil even if false is passed
### GetSections

`func (o *MicrosoftGraphSectionGroup) GetSections() []MicrosoftGraphOnenoteSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *MicrosoftGraphSectionGroup) GetSectionsOk() ([]MicrosoftGraphOnenoteSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSections

`func (o *MicrosoftGraphSectionGroup) HasSections() bool`

HasSections returns a boolean if a field has been set.

### SetSections

`func (o *MicrosoftGraphSectionGroup) SetSections(v []MicrosoftGraphOnenoteSection)`

SetSections gets a reference to the given []MicrosoftGraphOnenoteSection and assigns it to the Sections field.

### GetSectionGroups

`func (o *MicrosoftGraphSectionGroup) GetSectionGroups() []MicrosoftGraphSectionGroup`

GetSectionGroups returns the SectionGroups field if non-nil, zero value otherwise.

### GetSectionGroupsOk

`func (o *MicrosoftGraphSectionGroup) GetSectionGroupsOk() ([]MicrosoftGraphSectionGroup, bool)`

GetSectionGroupsOk returns a tuple with the SectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionGroups

`func (o *MicrosoftGraphSectionGroup) HasSectionGroups() bool`

HasSectionGroups returns a boolean if a field has been set.

### SetSectionGroups

`func (o *MicrosoftGraphSectionGroup) SetSectionGroups(v []MicrosoftGraphSectionGroup)`

SetSectionGroups gets a reference to the given []MicrosoftGraphSectionGroup and assigns it to the SectionGroups field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


