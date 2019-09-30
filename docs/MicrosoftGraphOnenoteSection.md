# MicrosoftGraphOnenoteSection

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
**Links** | Pointer to [**AnyOfmicrosoftGraphSectionLinks**](anyOf&lt;microsoft.graph.sectionLinks&gt;.md) |  | [optional] 
**PagesUrl** | Pointer to **string** |  | [optional] 
**ParentNotebook** | Pointer to [**AnyOfmicrosoftGraphNotebook**](anyOf&lt;microsoft.graph.notebook&gt;.md) |  | [optional] 
**ParentSectionGroup** | Pointer to [**AnyOfmicrosoftGraphSectionGroup**](anyOf&lt;microsoft.graph.sectionGroup&gt;.md) |  | [optional] 
**Pages** | Pointer to [**[]MicrosoftGraphOnenotePage**](microsoft.graph.onenotePage.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphOnenoteSection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOnenoteSection) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphOnenoteSection) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphOnenoteSection) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetSelf

`func (o *MicrosoftGraphOnenoteSection) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MicrosoftGraphOnenoteSection) GetSelfOk() (string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSelf

`func (o *MicrosoftGraphOnenoteSection) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### SetSelf

`func (o *MicrosoftGraphOnenoteSection) SetSelf(v string)`

SetSelf gets a reference to the given string and assigns it to the Self field.

### SetSelfExplicitNull

`func (o *MicrosoftGraphOnenoteSection) SetSelfExplicitNull(b bool)`

SetSelfExplicitNull (un)sets Self to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Self value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphOnenoteSection) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOnenoteSection) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphOnenoteSection) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOnenoteSection) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphOnenoteSection) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphOnenoteSection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphOnenoteSection) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphOnenoteSection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphOnenoteSection) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphOnenoteSection) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetCreatedBy

`func (o *MicrosoftGraphOnenoteSection) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphOnenoteSection) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphOnenoteSection) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphOnenoteSection) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphOnenoteSection) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetLastModifiedBy

`func (o *MicrosoftGraphOnenoteSection) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphOnenoteSection) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphOnenoteSection) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphOnenoteSection) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphOnenoteSection) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphOnenoteSection) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphOnenoteSection) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphOnenoteSection) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphOnenoteSection) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphOnenoteSection) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetIsDefault

`func (o *MicrosoftGraphOnenoteSection) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphOnenoteSection) GetIsDefaultOk() (bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDefault

`func (o *MicrosoftGraphOnenoteSection) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefault

`func (o *MicrosoftGraphOnenoteSection) SetIsDefault(v bool)`

SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.

### SetIsDefaultExplicitNull

`func (o *MicrosoftGraphOnenoteSection) SetIsDefaultExplicitNull(b bool)`

SetIsDefaultExplicitNull (un)sets IsDefault to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsDefault value is set to nil even if false is passed
### GetLinks

`func (o *MicrosoftGraphOnenoteSection) GetLinks() AnyOfmicrosoftGraphSectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MicrosoftGraphOnenoteSection) GetLinksOk() (AnyOfmicrosoftGraphSectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinks

`func (o *MicrosoftGraphOnenoteSection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinks

`func (o *MicrosoftGraphOnenoteSection) SetLinks(v AnyOfmicrosoftGraphSectionLinks)`

SetLinks gets a reference to the given AnyOfmicrosoftGraphSectionLinks and assigns it to the Links field.

### SetLinksExplicitNull

`func (o *MicrosoftGraphOnenoteSection) SetLinksExplicitNull(b bool)`

SetLinksExplicitNull (un)sets Links to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Links value is set to nil even if false is passed
### GetPagesUrl

`func (o *MicrosoftGraphOnenoteSection) GetPagesUrl() string`

GetPagesUrl returns the PagesUrl field if non-nil, zero value otherwise.

### GetPagesUrlOk

`func (o *MicrosoftGraphOnenoteSection) GetPagesUrlOk() (string, bool)`

GetPagesUrlOk returns a tuple with the PagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPagesUrl

`func (o *MicrosoftGraphOnenoteSection) HasPagesUrl() bool`

HasPagesUrl returns a boolean if a field has been set.

### SetPagesUrl

`func (o *MicrosoftGraphOnenoteSection) SetPagesUrl(v string)`

SetPagesUrl gets a reference to the given string and assigns it to the PagesUrl field.

### SetPagesUrlExplicitNull

`func (o *MicrosoftGraphOnenoteSection) SetPagesUrlExplicitNull(b bool)`

SetPagesUrlExplicitNull (un)sets PagesUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PagesUrl value is set to nil even if false is passed
### GetParentNotebook

`func (o *MicrosoftGraphOnenoteSection) GetParentNotebook() AnyOfmicrosoftGraphNotebook`

GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.

### GetParentNotebookOk

`func (o *MicrosoftGraphOnenoteSection) GetParentNotebookOk() (AnyOfmicrosoftGraphNotebook, bool)`

GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentNotebook

`func (o *MicrosoftGraphOnenoteSection) HasParentNotebook() bool`

HasParentNotebook returns a boolean if a field has been set.

### SetParentNotebook

`func (o *MicrosoftGraphOnenoteSection) SetParentNotebook(v AnyOfmicrosoftGraphNotebook)`

SetParentNotebook gets a reference to the given AnyOfmicrosoftGraphNotebook and assigns it to the ParentNotebook field.

### SetParentNotebookExplicitNull

`func (o *MicrosoftGraphOnenoteSection) SetParentNotebookExplicitNull(b bool)`

SetParentNotebookExplicitNull (un)sets ParentNotebook to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentNotebook value is set to nil even if false is passed
### GetParentSectionGroup

`func (o *MicrosoftGraphOnenoteSection) GetParentSectionGroup() AnyOfmicrosoftGraphSectionGroup`

GetParentSectionGroup returns the ParentSectionGroup field if non-nil, zero value otherwise.

### GetParentSectionGroupOk

`func (o *MicrosoftGraphOnenoteSection) GetParentSectionGroupOk() (AnyOfmicrosoftGraphSectionGroup, bool)`

GetParentSectionGroupOk returns a tuple with the ParentSectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentSectionGroup

`func (o *MicrosoftGraphOnenoteSection) HasParentSectionGroup() bool`

HasParentSectionGroup returns a boolean if a field has been set.

### SetParentSectionGroup

`func (o *MicrosoftGraphOnenoteSection) SetParentSectionGroup(v AnyOfmicrosoftGraphSectionGroup)`

SetParentSectionGroup gets a reference to the given AnyOfmicrosoftGraphSectionGroup and assigns it to the ParentSectionGroup field.

### SetParentSectionGroupExplicitNull

`func (o *MicrosoftGraphOnenoteSection) SetParentSectionGroupExplicitNull(b bool)`

SetParentSectionGroupExplicitNull (un)sets ParentSectionGroup to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentSectionGroup value is set to nil even if false is passed
### GetPages

`func (o *MicrosoftGraphOnenoteSection) GetPages() []MicrosoftGraphOnenotePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *MicrosoftGraphOnenoteSection) GetPagesOk() ([]MicrosoftGraphOnenotePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPages

`func (o *MicrosoftGraphOnenoteSection) HasPages() bool`

HasPages returns a boolean if a field has been set.

### SetPages

`func (o *MicrosoftGraphOnenoteSection) SetPages(v []MicrosoftGraphOnenotePage)`

SetPages gets a reference to the given []MicrosoftGraphOnenotePage and assigns it to the Pages field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


