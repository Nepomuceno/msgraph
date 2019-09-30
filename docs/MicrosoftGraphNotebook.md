# MicrosoftGraphNotebook

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
**IsDefault** | Pointer to **bool** |  | [optional] 
**UserRole** | Pointer to [**AnyOfmicrosoftGraphOnenoteUserRole**](anyOf&lt;microsoft.graph.onenoteUserRole&gt;.md) |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**SectionsUrl** | Pointer to **string** |  | [optional] 
**SectionGroupsUrl** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphNotebookLinks**](anyOf&lt;microsoft.graph.notebookLinks&gt;.md) |  | [optional] 
**Sections** | Pointer to [**[]MicrosoftGraphOnenoteSection**](microsoft.graph.onenoteSection.md) |  | [optional] 
**SectionGroups** | Pointer to [**[]MicrosoftGraphSectionGroup**](microsoft.graph.sectionGroup.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphNotebook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphNotebook) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphNotebook) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphNotebook) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetSelf

`func (o *MicrosoftGraphNotebook) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MicrosoftGraphNotebook) GetSelfOk() (string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSelf

`func (o *MicrosoftGraphNotebook) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### SetSelf

`func (o *MicrosoftGraphNotebook) SetSelf(v string)`

SetSelf gets a reference to the given string and assigns it to the Self field.

### SetSelfExplicitNull

`func (o *MicrosoftGraphNotebook) SetSelfExplicitNull(b bool)`

SetSelfExplicitNull (un)sets Self to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Self value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphNotebook) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphNotebook) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphNotebook) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphNotebook) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphNotebook) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphNotebook) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphNotebook) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphNotebook) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphNotebook) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphNotebook) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetCreatedBy

`func (o *MicrosoftGraphNotebook) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphNotebook) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphNotebook) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphNotebook) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphNotebook) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetLastModifiedBy

`func (o *MicrosoftGraphNotebook) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphNotebook) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphNotebook) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphNotebook) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphNotebook) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphNotebook) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphNotebook) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphNotebook) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphNotebook) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphNotebook) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetIsDefault

`func (o *MicrosoftGraphNotebook) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphNotebook) GetIsDefaultOk() (bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDefault

`func (o *MicrosoftGraphNotebook) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefault

`func (o *MicrosoftGraphNotebook) SetIsDefault(v bool)`

SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.

### SetIsDefaultExplicitNull

`func (o *MicrosoftGraphNotebook) SetIsDefaultExplicitNull(b bool)`

SetIsDefaultExplicitNull (un)sets IsDefault to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsDefault value is set to nil even if false is passed
### GetUserRole

`func (o *MicrosoftGraphNotebook) GetUserRole() AnyOfmicrosoftGraphOnenoteUserRole`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *MicrosoftGraphNotebook) GetUserRoleOk() (AnyOfmicrosoftGraphOnenoteUserRole, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserRole

`func (o *MicrosoftGraphNotebook) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### SetUserRole

`func (o *MicrosoftGraphNotebook) SetUserRole(v AnyOfmicrosoftGraphOnenoteUserRole)`

SetUserRole gets a reference to the given AnyOfmicrosoftGraphOnenoteUserRole and assigns it to the UserRole field.

### SetUserRoleExplicitNull

`func (o *MicrosoftGraphNotebook) SetUserRoleExplicitNull(b bool)`

SetUserRoleExplicitNull (un)sets UserRole to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserRole value is set to nil even if false is passed
### GetIsShared

`func (o *MicrosoftGraphNotebook) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *MicrosoftGraphNotebook) GetIsSharedOk() (bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsShared

`func (o *MicrosoftGraphNotebook) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### SetIsShared

`func (o *MicrosoftGraphNotebook) SetIsShared(v bool)`

SetIsShared gets a reference to the given bool and assigns it to the IsShared field.

### SetIsSharedExplicitNull

`func (o *MicrosoftGraphNotebook) SetIsSharedExplicitNull(b bool)`

SetIsSharedExplicitNull (un)sets IsShared to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsShared value is set to nil even if false is passed
### GetSectionsUrl

`func (o *MicrosoftGraphNotebook) GetSectionsUrl() string`

GetSectionsUrl returns the SectionsUrl field if non-nil, zero value otherwise.

### GetSectionsUrlOk

`func (o *MicrosoftGraphNotebook) GetSectionsUrlOk() (string, bool)`

GetSectionsUrlOk returns a tuple with the SectionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionsUrl

`func (o *MicrosoftGraphNotebook) HasSectionsUrl() bool`

HasSectionsUrl returns a boolean if a field has been set.

### SetSectionsUrl

`func (o *MicrosoftGraphNotebook) SetSectionsUrl(v string)`

SetSectionsUrl gets a reference to the given string and assigns it to the SectionsUrl field.

### SetSectionsUrlExplicitNull

`func (o *MicrosoftGraphNotebook) SetSectionsUrlExplicitNull(b bool)`

SetSectionsUrlExplicitNull (un)sets SectionsUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SectionsUrl value is set to nil even if false is passed
### GetSectionGroupsUrl

`func (o *MicrosoftGraphNotebook) GetSectionGroupsUrl() string`

GetSectionGroupsUrl returns the SectionGroupsUrl field if non-nil, zero value otherwise.

### GetSectionGroupsUrlOk

`func (o *MicrosoftGraphNotebook) GetSectionGroupsUrlOk() (string, bool)`

GetSectionGroupsUrlOk returns a tuple with the SectionGroupsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionGroupsUrl

`func (o *MicrosoftGraphNotebook) HasSectionGroupsUrl() bool`

HasSectionGroupsUrl returns a boolean if a field has been set.

### SetSectionGroupsUrl

`func (o *MicrosoftGraphNotebook) SetSectionGroupsUrl(v string)`

SetSectionGroupsUrl gets a reference to the given string and assigns it to the SectionGroupsUrl field.

### SetSectionGroupsUrlExplicitNull

`func (o *MicrosoftGraphNotebook) SetSectionGroupsUrlExplicitNull(b bool)`

SetSectionGroupsUrlExplicitNull (un)sets SectionGroupsUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SectionGroupsUrl value is set to nil even if false is passed
### GetLinks

`func (o *MicrosoftGraphNotebook) GetLinks() AnyOfmicrosoftGraphNotebookLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MicrosoftGraphNotebook) GetLinksOk() (AnyOfmicrosoftGraphNotebookLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinks

`func (o *MicrosoftGraphNotebook) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinks

`func (o *MicrosoftGraphNotebook) SetLinks(v AnyOfmicrosoftGraphNotebookLinks)`

SetLinks gets a reference to the given AnyOfmicrosoftGraphNotebookLinks and assigns it to the Links field.

### SetLinksExplicitNull

`func (o *MicrosoftGraphNotebook) SetLinksExplicitNull(b bool)`

SetLinksExplicitNull (un)sets Links to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Links value is set to nil even if false is passed
### GetSections

`func (o *MicrosoftGraphNotebook) GetSections() []MicrosoftGraphOnenoteSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *MicrosoftGraphNotebook) GetSectionsOk() ([]MicrosoftGraphOnenoteSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSections

`func (o *MicrosoftGraphNotebook) HasSections() bool`

HasSections returns a boolean if a field has been set.

### SetSections

`func (o *MicrosoftGraphNotebook) SetSections(v []MicrosoftGraphOnenoteSection)`

SetSections gets a reference to the given []MicrosoftGraphOnenoteSection and assigns it to the Sections field.

### GetSectionGroups

`func (o *MicrosoftGraphNotebook) GetSectionGroups() []MicrosoftGraphSectionGroup`

GetSectionGroups returns the SectionGroups field if non-nil, zero value otherwise.

### GetSectionGroupsOk

`func (o *MicrosoftGraphNotebook) GetSectionGroupsOk() ([]MicrosoftGraphSectionGroup, bool)`

GetSectionGroupsOk returns a tuple with the SectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionGroups

`func (o *MicrosoftGraphNotebook) HasSectionGroups() bool`

HasSectionGroups returns a boolean if a field has been set.

### SetSectionGroups

`func (o *MicrosoftGraphNotebook) SetSectionGroups(v []MicrosoftGraphSectionGroup)`

SetSectionGroups gets a reference to the given []MicrosoftGraphSectionGroup and assigns it to the SectionGroups field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


