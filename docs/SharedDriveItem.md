# SharedDriveItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**DriveItem** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) |  | [optional] 
**Items** | Pointer to [**[]MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md) |  | [optional] 
**List** | Pointer to [**AnyOfmicrosoftGraphList**](anyOf&lt;microsoft.graph.list&gt;.md) |  | [optional] 
**ListItem** | Pointer to [**AnyOfmicrosoftGraphListItem**](anyOf&lt;microsoft.graph.listItem&gt;.md) |  | [optional] 
**Root** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) |  | [optional] 
**Site** | Pointer to [**AnyOfmicrosoftGraphSite**](anyOf&lt;microsoft.graph.site&gt;.md) |  | [optional] 

## Methods

### GetOwner

`func (o *SharedDriveItem) GetOwner() AnyOfmicrosoftGraphIdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SharedDriveItem) GetOwnerOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *SharedDriveItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *SharedDriveItem) SetOwner(v AnyOfmicrosoftGraphIdentitySet)`

SetOwner gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *SharedDriveItem) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetDriveItem

`func (o *SharedDriveItem) GetDriveItem() AnyOfmicrosoftGraphDriveItem`

GetDriveItem returns the DriveItem field if non-nil, zero value otherwise.

### GetDriveItemOk

`func (o *SharedDriveItem) GetDriveItemOk() (AnyOfmicrosoftGraphDriveItem, bool)`

GetDriveItemOk returns a tuple with the DriveItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDriveItem

`func (o *SharedDriveItem) HasDriveItem() bool`

HasDriveItem returns a boolean if a field has been set.

### SetDriveItem

`func (o *SharedDriveItem) SetDriveItem(v AnyOfmicrosoftGraphDriveItem)`

SetDriveItem gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the DriveItem field.

### SetDriveItemExplicitNull

`func (o *SharedDriveItem) SetDriveItemExplicitNull(b bool)`

SetDriveItemExplicitNull (un)sets DriveItem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DriveItem value is set to nil even if false is passed
### GetItems

`func (o *SharedDriveItem) GetItems() []MicrosoftGraphDriveItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SharedDriveItem) GetItemsOk() ([]MicrosoftGraphDriveItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItems

`func (o *SharedDriveItem) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItems

`func (o *SharedDriveItem) SetItems(v []MicrosoftGraphDriveItem)`

SetItems gets a reference to the given []MicrosoftGraphDriveItem and assigns it to the Items field.

### GetList

`func (o *SharedDriveItem) GetList() AnyOfmicrosoftGraphList`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *SharedDriveItem) GetListOk() (AnyOfmicrosoftGraphList, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasList

`func (o *SharedDriveItem) HasList() bool`

HasList returns a boolean if a field has been set.

### SetList

`func (o *SharedDriveItem) SetList(v AnyOfmicrosoftGraphList)`

SetList gets a reference to the given AnyOfmicrosoftGraphList and assigns it to the List field.

### SetListExplicitNull

`func (o *SharedDriveItem) SetListExplicitNull(b bool)`

SetListExplicitNull (un)sets List to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The List value is set to nil even if false is passed
### GetListItem

`func (o *SharedDriveItem) GetListItem() AnyOfmicrosoftGraphListItem`

GetListItem returns the ListItem field if non-nil, zero value otherwise.

### GetListItemOk

`func (o *SharedDriveItem) GetListItemOk() (AnyOfmicrosoftGraphListItem, bool)`

GetListItemOk returns a tuple with the ListItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListItem

`func (o *SharedDriveItem) HasListItem() bool`

HasListItem returns a boolean if a field has been set.

### SetListItem

`func (o *SharedDriveItem) SetListItem(v AnyOfmicrosoftGraphListItem)`

SetListItem gets a reference to the given AnyOfmicrosoftGraphListItem and assigns it to the ListItem field.

### SetListItemExplicitNull

`func (o *SharedDriveItem) SetListItemExplicitNull(b bool)`

SetListItemExplicitNull (un)sets ListItem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ListItem value is set to nil even if false is passed
### GetRoot

`func (o *SharedDriveItem) GetRoot() AnyOfmicrosoftGraphDriveItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *SharedDriveItem) GetRootOk() (AnyOfmicrosoftGraphDriveItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoot

`func (o *SharedDriveItem) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRoot

`func (o *SharedDriveItem) SetRoot(v AnyOfmicrosoftGraphDriveItem)`

SetRoot gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the Root field.

### SetRootExplicitNull

`func (o *SharedDriveItem) SetRootExplicitNull(b bool)`

SetRootExplicitNull (un)sets Root to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Root value is set to nil even if false is passed
### GetSite

`func (o *SharedDriveItem) GetSite() AnyOfmicrosoftGraphSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *SharedDriveItem) GetSiteOk() (AnyOfmicrosoftGraphSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSite

`func (o *SharedDriveItem) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSite

`func (o *SharedDriveItem) SetSite(v AnyOfmicrosoftGraphSite)`

SetSite gets a reference to the given AnyOfmicrosoftGraphSite and assigns it to the Site field.

### SetSiteExplicitNull

`func (o *SharedDriveItem) SetSiteExplicitNull(b bool)`

SetSiteExplicitNull (un)sets Site to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Site value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


