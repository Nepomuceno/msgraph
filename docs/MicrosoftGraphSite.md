# MicrosoftGraphSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentReference** | Pointer to [**AnyOfmicrosoftGraphItemReference**](anyOf&lt;microsoft.graph.itemReference&gt;.md) |  | [optional] 
**WebUrl** | Pointer to **string** |  | [optional] 
**CreatedByUser** | Pointer to [**AnyOfmicrosoftGraphUser**](anyOf&lt;microsoft.graph.user&gt;.md) |  | [optional] 
**LastModifiedByUser** | Pointer to [**AnyOfmicrosoftGraphUser**](anyOf&lt;microsoft.graph.user&gt;.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Root** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**SharepointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) |  | [optional] 
**SiteCollection** | Pointer to [**AnyOfmicrosoftGraphSiteCollection**](anyOf&lt;microsoft.graph.siteCollection&gt;.md) |  | [optional] 
**Analytics** | Pointer to [**AnyOfmicrosoftGraphItemAnalytics**](anyOf&lt;microsoft.graph.itemAnalytics&gt;.md) |  | [optional] 
**Columns** | Pointer to [**[]MicrosoftGraphColumnDefinition**](microsoft.graph.columnDefinition.md) |  | [optional] 
**ContentTypes** | Pointer to [**[]MicrosoftGraphContentType**](microsoft.graph.contentType.md) |  | [optional] 
**Drive** | Pointer to [**AnyOfmicrosoftGraphDrive**](anyOf&lt;microsoft.graph.drive&gt;.md) |  | [optional] 
**Drives** | Pointer to [**[]MicrosoftGraphDrive**](microsoft.graph.drive.md) |  | [optional] 
**Items** | Pointer to [**[]MicrosoftGraphBaseItem**](microsoft.graph.baseItem.md) |  | [optional] 
**Lists** | Pointer to [**[]MicrosoftGraphList**](microsoft.graph.list.md) |  | [optional] 
**Sites** | Pointer to [**[]MicrosoftGraphSite**](microsoft.graph.site.md) |  | [optional] 
**Onenote** | Pointer to [**AnyOfmicrosoftGraphOnenote**](anyOf&lt;microsoft.graph.onenote&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphSite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSite) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphSite) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphSite) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedBy

`func (o *MicrosoftGraphSite) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphSite) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphSite) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphSite) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphSite) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphSite) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphSite) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphSite) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphSite) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphSite) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphSite) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphSite) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphSite) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphSite) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetETag

`func (o *MicrosoftGraphSite) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MicrosoftGraphSite) GetETagOk() (string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasETag

`func (o *MicrosoftGraphSite) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETag

`func (o *MicrosoftGraphSite) SetETag(v string)`

SetETag gets a reference to the given string and assigns it to the ETag field.

### SetETagExplicitNull

`func (o *MicrosoftGraphSite) SetETagExplicitNull(b bool)`

SetETagExplicitNull (un)sets ETag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ETag value is set to nil even if false is passed
### GetLastModifiedBy

`func (o *MicrosoftGraphSite) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphSite) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphSite) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphSite) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphSite) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphSite) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphSite) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphSite) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphSite) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetName

`func (o *MicrosoftGraphSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphSite) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphSite) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphSite) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphSite) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetParentReference

`func (o *MicrosoftGraphSite) GetParentReference() AnyOfmicrosoftGraphItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MicrosoftGraphSite) GetParentReferenceOk() (AnyOfmicrosoftGraphItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentReference

`func (o *MicrosoftGraphSite) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### SetParentReference

`func (o *MicrosoftGraphSite) SetParentReference(v AnyOfmicrosoftGraphItemReference)`

SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.

### SetParentReferenceExplicitNull

`func (o *MicrosoftGraphSite) SetParentReferenceExplicitNull(b bool)`

SetParentReferenceExplicitNull (un)sets ParentReference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentReference value is set to nil even if false is passed
### GetWebUrl

`func (o *MicrosoftGraphSite) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphSite) GetWebUrlOk() (string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebUrl

`func (o *MicrosoftGraphSite) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrl

`func (o *MicrosoftGraphSite) SetWebUrl(v string)`

SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.

### SetWebUrlExplicitNull

`func (o *MicrosoftGraphSite) SetWebUrlExplicitNull(b bool)`

SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebUrl value is set to nil even if false is passed
### GetCreatedByUser

`func (o *MicrosoftGraphSite) GetCreatedByUser() AnyOfmicrosoftGraphUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *MicrosoftGraphSite) GetCreatedByUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedByUser

`func (o *MicrosoftGraphSite) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUser

`func (o *MicrosoftGraphSite) SetCreatedByUser(v AnyOfmicrosoftGraphUser)`

SetCreatedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the CreatedByUser field.

### SetCreatedByUserExplicitNull

`func (o *MicrosoftGraphSite) SetCreatedByUserExplicitNull(b bool)`

SetCreatedByUserExplicitNull (un)sets CreatedByUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedByUser value is set to nil even if false is passed
### GetLastModifiedByUser

`func (o *MicrosoftGraphSite) GetLastModifiedByUser() AnyOfmicrosoftGraphUser`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *MicrosoftGraphSite) GetLastModifiedByUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedByUser

`func (o *MicrosoftGraphSite) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.

### SetLastModifiedByUser

`func (o *MicrosoftGraphSite) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser)`

SetLastModifiedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the LastModifiedByUser field.

### SetLastModifiedByUserExplicitNull

`func (o *MicrosoftGraphSite) SetLastModifiedByUserExplicitNull(b bool)`

SetLastModifiedByUserExplicitNull (un)sets LastModifiedByUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedByUser value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphSite) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphSite) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphSite) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphSite) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphSite) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetRoot

`func (o *MicrosoftGraphSite) GetRoot() AnyOfobject`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *MicrosoftGraphSite) GetRootOk() (AnyOfobject, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoot

`func (o *MicrosoftGraphSite) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRoot

`func (o *MicrosoftGraphSite) SetRoot(v AnyOfobject)`

SetRoot gets a reference to the given AnyOfobject and assigns it to the Root field.

### SetRootExplicitNull

`func (o *MicrosoftGraphSite) SetRootExplicitNull(b bool)`

SetRootExplicitNull (un)sets Root to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Root value is set to nil even if false is passed
### GetSharepointIds

`func (o *MicrosoftGraphSite) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *MicrosoftGraphSite) GetSharepointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharepointIds

`func (o *MicrosoftGraphSite) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIds

`func (o *MicrosoftGraphSite) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.

### SetSharepointIdsExplicitNull

`func (o *MicrosoftGraphSite) SetSharepointIdsExplicitNull(b bool)`

SetSharepointIdsExplicitNull (un)sets SharepointIds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharepointIds value is set to nil even if false is passed
### GetSiteCollection

`func (o *MicrosoftGraphSite) GetSiteCollection() AnyOfmicrosoftGraphSiteCollection`

GetSiteCollection returns the SiteCollection field if non-nil, zero value otherwise.

### GetSiteCollectionOk

`func (o *MicrosoftGraphSite) GetSiteCollectionOk() (AnyOfmicrosoftGraphSiteCollection, bool)`

GetSiteCollectionOk returns a tuple with the SiteCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiteCollection

`func (o *MicrosoftGraphSite) HasSiteCollection() bool`

HasSiteCollection returns a boolean if a field has been set.

### SetSiteCollection

`func (o *MicrosoftGraphSite) SetSiteCollection(v AnyOfmicrosoftGraphSiteCollection)`

SetSiteCollection gets a reference to the given AnyOfmicrosoftGraphSiteCollection and assigns it to the SiteCollection field.

### SetSiteCollectionExplicitNull

`func (o *MicrosoftGraphSite) SetSiteCollectionExplicitNull(b bool)`

SetSiteCollectionExplicitNull (un)sets SiteCollection to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SiteCollection value is set to nil even if false is passed
### GetAnalytics

`func (o *MicrosoftGraphSite) GetAnalytics() AnyOfmicrosoftGraphItemAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *MicrosoftGraphSite) GetAnalyticsOk() (AnyOfmicrosoftGraphItemAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAnalytics

`func (o *MicrosoftGraphSite) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### SetAnalytics

`func (o *MicrosoftGraphSite) SetAnalytics(v AnyOfmicrosoftGraphItemAnalytics)`

SetAnalytics gets a reference to the given AnyOfmicrosoftGraphItemAnalytics and assigns it to the Analytics field.

### SetAnalyticsExplicitNull

`func (o *MicrosoftGraphSite) SetAnalyticsExplicitNull(b bool)`

SetAnalyticsExplicitNull (un)sets Analytics to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Analytics value is set to nil even if false is passed
### GetColumns

`func (o *MicrosoftGraphSite) GetColumns() []MicrosoftGraphColumnDefinition`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *MicrosoftGraphSite) GetColumnsOk() ([]MicrosoftGraphColumnDefinition, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumns

`func (o *MicrosoftGraphSite) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumns

`func (o *MicrosoftGraphSite) SetColumns(v []MicrosoftGraphColumnDefinition)`

SetColumns gets a reference to the given []MicrosoftGraphColumnDefinition and assigns it to the Columns field.

### GetContentTypes

`func (o *MicrosoftGraphSite) GetContentTypes() []MicrosoftGraphContentType`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *MicrosoftGraphSite) GetContentTypesOk() ([]MicrosoftGraphContentType, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentTypes

`func (o *MicrosoftGraphSite) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### SetContentTypes

`func (o *MicrosoftGraphSite) SetContentTypes(v []MicrosoftGraphContentType)`

SetContentTypes gets a reference to the given []MicrosoftGraphContentType and assigns it to the ContentTypes field.

### GetDrive

`func (o *MicrosoftGraphSite) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *MicrosoftGraphSite) GetDriveOk() (AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrive

`func (o *MicrosoftGraphSite) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDrive

`func (o *MicrosoftGraphSite) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive gets a reference to the given AnyOfmicrosoftGraphDrive and assigns it to the Drive field.

### SetDriveExplicitNull

`func (o *MicrosoftGraphSite) SetDriveExplicitNull(b bool)`

SetDriveExplicitNull (un)sets Drive to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Drive value is set to nil even if false is passed
### GetDrives

`func (o *MicrosoftGraphSite) GetDrives() []MicrosoftGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *MicrosoftGraphSite) GetDrivesOk() ([]MicrosoftGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrives

`func (o *MicrosoftGraphSite) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### SetDrives

`func (o *MicrosoftGraphSite) SetDrives(v []MicrosoftGraphDrive)`

SetDrives gets a reference to the given []MicrosoftGraphDrive and assigns it to the Drives field.

### GetItems

`func (o *MicrosoftGraphSite) GetItems() []MicrosoftGraphBaseItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MicrosoftGraphSite) GetItemsOk() ([]MicrosoftGraphBaseItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItems

`func (o *MicrosoftGraphSite) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItems

`func (o *MicrosoftGraphSite) SetItems(v []MicrosoftGraphBaseItem)`

SetItems gets a reference to the given []MicrosoftGraphBaseItem and assigns it to the Items field.

### GetLists

`func (o *MicrosoftGraphSite) GetLists() []MicrosoftGraphList`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *MicrosoftGraphSite) GetListsOk() ([]MicrosoftGraphList, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLists

`func (o *MicrosoftGraphSite) HasLists() bool`

HasLists returns a boolean if a field has been set.

### SetLists

`func (o *MicrosoftGraphSite) SetLists(v []MicrosoftGraphList)`

SetLists gets a reference to the given []MicrosoftGraphList and assigns it to the Lists field.

### GetSites

`func (o *MicrosoftGraphSite) GetSites() []MicrosoftGraphSite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *MicrosoftGraphSite) GetSitesOk() ([]MicrosoftGraphSite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSites

`func (o *MicrosoftGraphSite) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSites

`func (o *MicrosoftGraphSite) SetSites(v []MicrosoftGraphSite)`

SetSites gets a reference to the given []MicrosoftGraphSite and assigns it to the Sites field.

### GetOnenote

`func (o *MicrosoftGraphSite) GetOnenote() AnyOfmicrosoftGraphOnenote`

GetOnenote returns the Onenote field if non-nil, zero value otherwise.

### GetOnenoteOk

`func (o *MicrosoftGraphSite) GetOnenoteOk() (AnyOfmicrosoftGraphOnenote, bool)`

GetOnenoteOk returns a tuple with the Onenote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnenote

`func (o *MicrosoftGraphSite) HasOnenote() bool`

HasOnenote returns a boolean if a field has been set.

### SetOnenote

`func (o *MicrosoftGraphSite) SetOnenote(v AnyOfmicrosoftGraphOnenote)`

SetOnenote gets a reference to the given AnyOfmicrosoftGraphOnenote and assigns it to the Onenote field.

### SetOnenoteExplicitNull

`func (o *MicrosoftGraphSite) SetOnenoteExplicitNull(b bool)`

SetOnenoteExplicitNull (un)sets Onenote to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Onenote value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


