# Site

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetDisplayName

`func (o *Site) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Site) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *Site) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *Site) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *Site) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetRoot

`func (o *Site) GetRoot() AnyOfobject`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *Site) GetRootOk() (AnyOfobject, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoot

`func (o *Site) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRoot

`func (o *Site) SetRoot(v AnyOfobject)`

SetRoot gets a reference to the given AnyOfobject and assigns it to the Root field.

### SetRootExplicitNull

`func (o *Site) SetRootExplicitNull(b bool)`

SetRootExplicitNull (un)sets Root to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Root value is set to nil even if false is passed
### GetSharepointIds

`func (o *Site) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *Site) GetSharepointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharepointIds

`func (o *Site) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIds

`func (o *Site) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.

### SetSharepointIdsExplicitNull

`func (o *Site) SetSharepointIdsExplicitNull(b bool)`

SetSharepointIdsExplicitNull (un)sets SharepointIds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharepointIds value is set to nil even if false is passed
### GetSiteCollection

`func (o *Site) GetSiteCollection() AnyOfmicrosoftGraphSiteCollection`

GetSiteCollection returns the SiteCollection field if non-nil, zero value otherwise.

### GetSiteCollectionOk

`func (o *Site) GetSiteCollectionOk() (AnyOfmicrosoftGraphSiteCollection, bool)`

GetSiteCollectionOk returns a tuple with the SiteCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiteCollection

`func (o *Site) HasSiteCollection() bool`

HasSiteCollection returns a boolean if a field has been set.

### SetSiteCollection

`func (o *Site) SetSiteCollection(v AnyOfmicrosoftGraphSiteCollection)`

SetSiteCollection gets a reference to the given AnyOfmicrosoftGraphSiteCollection and assigns it to the SiteCollection field.

### SetSiteCollectionExplicitNull

`func (o *Site) SetSiteCollectionExplicitNull(b bool)`

SetSiteCollectionExplicitNull (un)sets SiteCollection to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SiteCollection value is set to nil even if false is passed
### GetAnalytics

`func (o *Site) GetAnalytics() AnyOfmicrosoftGraphItemAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *Site) GetAnalyticsOk() (AnyOfmicrosoftGraphItemAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAnalytics

`func (o *Site) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### SetAnalytics

`func (o *Site) SetAnalytics(v AnyOfmicrosoftGraphItemAnalytics)`

SetAnalytics gets a reference to the given AnyOfmicrosoftGraphItemAnalytics and assigns it to the Analytics field.

### SetAnalyticsExplicitNull

`func (o *Site) SetAnalyticsExplicitNull(b bool)`

SetAnalyticsExplicitNull (un)sets Analytics to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Analytics value is set to nil even if false is passed
### GetColumns

`func (o *Site) GetColumns() []MicrosoftGraphColumnDefinition`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Site) GetColumnsOk() ([]MicrosoftGraphColumnDefinition, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumns

`func (o *Site) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumns

`func (o *Site) SetColumns(v []MicrosoftGraphColumnDefinition)`

SetColumns gets a reference to the given []MicrosoftGraphColumnDefinition and assigns it to the Columns field.

### GetContentTypes

`func (o *Site) GetContentTypes() []MicrosoftGraphContentType`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *Site) GetContentTypesOk() ([]MicrosoftGraphContentType, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentTypes

`func (o *Site) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### SetContentTypes

`func (o *Site) SetContentTypes(v []MicrosoftGraphContentType)`

SetContentTypes gets a reference to the given []MicrosoftGraphContentType and assigns it to the ContentTypes field.

### GetDrive

`func (o *Site) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *Site) GetDriveOk() (AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrive

`func (o *Site) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDrive

`func (o *Site) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive gets a reference to the given AnyOfmicrosoftGraphDrive and assigns it to the Drive field.

### SetDriveExplicitNull

`func (o *Site) SetDriveExplicitNull(b bool)`

SetDriveExplicitNull (un)sets Drive to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Drive value is set to nil even if false is passed
### GetDrives

`func (o *Site) GetDrives() []MicrosoftGraphDrive`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *Site) GetDrivesOk() ([]MicrosoftGraphDrive, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrives

`func (o *Site) HasDrives() bool`

HasDrives returns a boolean if a field has been set.

### SetDrives

`func (o *Site) SetDrives(v []MicrosoftGraphDrive)`

SetDrives gets a reference to the given []MicrosoftGraphDrive and assigns it to the Drives field.

### GetItems

`func (o *Site) GetItems() []MicrosoftGraphBaseItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Site) GetItemsOk() ([]MicrosoftGraphBaseItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItems

`func (o *Site) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItems

`func (o *Site) SetItems(v []MicrosoftGraphBaseItem)`

SetItems gets a reference to the given []MicrosoftGraphBaseItem and assigns it to the Items field.

### GetLists

`func (o *Site) GetLists() []MicrosoftGraphList`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *Site) GetListsOk() ([]MicrosoftGraphList, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLists

`func (o *Site) HasLists() bool`

HasLists returns a boolean if a field has been set.

### SetLists

`func (o *Site) SetLists(v []MicrosoftGraphList)`

SetLists gets a reference to the given []MicrosoftGraphList and assigns it to the Lists field.

### GetSites

`func (o *Site) GetSites() []MicrosoftGraphSite`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *Site) GetSitesOk() ([]MicrosoftGraphSite, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSites

`func (o *Site) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSites

`func (o *Site) SetSites(v []MicrosoftGraphSite)`

SetSites gets a reference to the given []MicrosoftGraphSite and assigns it to the Sites field.

### GetOnenote

`func (o *Site) GetOnenote() AnyOfmicrosoftGraphOnenote`

GetOnenote returns the Onenote field if non-nil, zero value otherwise.

### GetOnenoteOk

`func (o *Site) GetOnenoteOk() (AnyOfmicrosoftGraphOnenote, bool)`

GetOnenoteOk returns a tuple with the Onenote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnenote

`func (o *Site) HasOnenote() bool`

HasOnenote returns a boolean if a field has been set.

### SetOnenote

`func (o *Site) SetOnenote(v AnyOfmicrosoftGraphOnenote)`

SetOnenote gets a reference to the given AnyOfmicrosoftGraphOnenote and assigns it to the Onenote field.

### SetOnenoteExplicitNull

`func (o *Site) SetOnenoteExplicitNull(b bool)`

SetOnenoteExplicitNull (un)sets Onenote to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Onenote value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


