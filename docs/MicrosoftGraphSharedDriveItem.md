# MicrosoftGraphSharedDriveItem

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
**Owner** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**DriveItem** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) |  | [optional] 
**Items** | Pointer to [**[]MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md) |  | [optional] 
**List** | Pointer to [**AnyOfmicrosoftGraphList**](anyOf&lt;microsoft.graph.list&gt;.md) |  | [optional] 
**ListItem** | Pointer to [**AnyOfmicrosoftGraphListItem**](anyOf&lt;microsoft.graph.listItem&gt;.md) |  | [optional] 
**Root** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) |  | [optional] 
**Site** | Pointer to [**AnyOfmicrosoftGraphSite**](anyOf&lt;microsoft.graph.site&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphSharedDriveItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSharedDriveItem) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphSharedDriveItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphSharedDriveItem) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedBy

`func (o *MicrosoftGraphSharedDriveItem) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphSharedDriveItem) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphSharedDriveItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphSharedDriveItem) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphSharedDriveItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphSharedDriveItem) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphSharedDriveItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphSharedDriveItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphSharedDriveItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphSharedDriveItem) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphSharedDriveItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphSharedDriveItem) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetETag

`func (o *MicrosoftGraphSharedDriveItem) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MicrosoftGraphSharedDriveItem) GetETagOk() (string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasETag

`func (o *MicrosoftGraphSharedDriveItem) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETag

`func (o *MicrosoftGraphSharedDriveItem) SetETag(v string)`

SetETag gets a reference to the given string and assigns it to the ETag field.

### SetETagExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetETagExplicitNull(b bool)`

SetETagExplicitNull (un)sets ETag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ETag value is set to nil even if false is passed
### GetLastModifiedBy

`func (o *MicrosoftGraphSharedDriveItem) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphSharedDriveItem) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphSharedDriveItem) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphSharedDriveItem) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphSharedDriveItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphSharedDriveItem) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphSharedDriveItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphSharedDriveItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetName

`func (o *MicrosoftGraphSharedDriveItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphSharedDriveItem) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphSharedDriveItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphSharedDriveItem) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetParentReference

`func (o *MicrosoftGraphSharedDriveItem) GetParentReference() AnyOfmicrosoftGraphItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MicrosoftGraphSharedDriveItem) GetParentReferenceOk() (AnyOfmicrosoftGraphItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentReference

`func (o *MicrosoftGraphSharedDriveItem) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### SetParentReference

`func (o *MicrosoftGraphSharedDriveItem) SetParentReference(v AnyOfmicrosoftGraphItemReference)`

SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.

### SetParentReferenceExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetParentReferenceExplicitNull(b bool)`

SetParentReferenceExplicitNull (un)sets ParentReference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentReference value is set to nil even if false is passed
### GetWebUrl

`func (o *MicrosoftGraphSharedDriveItem) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphSharedDriveItem) GetWebUrlOk() (string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebUrl

`func (o *MicrosoftGraphSharedDriveItem) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrl

`func (o *MicrosoftGraphSharedDriveItem) SetWebUrl(v string)`

SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.

### SetWebUrlExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetWebUrlExplicitNull(b bool)`

SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebUrl value is set to nil even if false is passed
### GetCreatedByUser

`func (o *MicrosoftGraphSharedDriveItem) GetCreatedByUser() AnyOfmicrosoftGraphUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *MicrosoftGraphSharedDriveItem) GetCreatedByUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedByUser

`func (o *MicrosoftGraphSharedDriveItem) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUser

`func (o *MicrosoftGraphSharedDriveItem) SetCreatedByUser(v AnyOfmicrosoftGraphUser)`

SetCreatedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the CreatedByUser field.

### SetCreatedByUserExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetCreatedByUserExplicitNull(b bool)`

SetCreatedByUserExplicitNull (un)sets CreatedByUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedByUser value is set to nil even if false is passed
### GetLastModifiedByUser

`func (o *MicrosoftGraphSharedDriveItem) GetLastModifiedByUser() AnyOfmicrosoftGraphUser`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *MicrosoftGraphSharedDriveItem) GetLastModifiedByUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedByUser

`func (o *MicrosoftGraphSharedDriveItem) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.

### SetLastModifiedByUser

`func (o *MicrosoftGraphSharedDriveItem) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser)`

SetLastModifiedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the LastModifiedByUser field.

### SetLastModifiedByUserExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetLastModifiedByUserExplicitNull(b bool)`

SetLastModifiedByUserExplicitNull (un)sets LastModifiedByUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedByUser value is set to nil even if false is passed
### GetOwner

`func (o *MicrosoftGraphSharedDriveItem) GetOwner() AnyOfmicrosoftGraphIdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphSharedDriveItem) GetOwnerOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *MicrosoftGraphSharedDriveItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *MicrosoftGraphSharedDriveItem) SetOwner(v AnyOfmicrosoftGraphIdentitySet)`

SetOwner gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetDriveItem

`func (o *MicrosoftGraphSharedDriveItem) GetDriveItem() AnyOfmicrosoftGraphDriveItem`

GetDriveItem returns the DriveItem field if non-nil, zero value otherwise.

### GetDriveItemOk

`func (o *MicrosoftGraphSharedDriveItem) GetDriveItemOk() (AnyOfmicrosoftGraphDriveItem, bool)`

GetDriveItemOk returns a tuple with the DriveItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDriveItem

`func (o *MicrosoftGraphSharedDriveItem) HasDriveItem() bool`

HasDriveItem returns a boolean if a field has been set.

### SetDriveItem

`func (o *MicrosoftGraphSharedDriveItem) SetDriveItem(v AnyOfmicrosoftGraphDriveItem)`

SetDriveItem gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the DriveItem field.

### SetDriveItemExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetDriveItemExplicitNull(b bool)`

SetDriveItemExplicitNull (un)sets DriveItem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DriveItem value is set to nil even if false is passed
### GetItems

`func (o *MicrosoftGraphSharedDriveItem) GetItems() []MicrosoftGraphDriveItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MicrosoftGraphSharedDriveItem) GetItemsOk() ([]MicrosoftGraphDriveItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItems

`func (o *MicrosoftGraphSharedDriveItem) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItems

`func (o *MicrosoftGraphSharedDriveItem) SetItems(v []MicrosoftGraphDriveItem)`

SetItems gets a reference to the given []MicrosoftGraphDriveItem and assigns it to the Items field.

### GetList

`func (o *MicrosoftGraphSharedDriveItem) GetList() AnyOfmicrosoftGraphList`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *MicrosoftGraphSharedDriveItem) GetListOk() (AnyOfmicrosoftGraphList, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasList

`func (o *MicrosoftGraphSharedDriveItem) HasList() bool`

HasList returns a boolean if a field has been set.

### SetList

`func (o *MicrosoftGraphSharedDriveItem) SetList(v AnyOfmicrosoftGraphList)`

SetList gets a reference to the given AnyOfmicrosoftGraphList and assigns it to the List field.

### SetListExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetListExplicitNull(b bool)`

SetListExplicitNull (un)sets List to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The List value is set to nil even if false is passed
### GetListItem

`func (o *MicrosoftGraphSharedDriveItem) GetListItem() AnyOfmicrosoftGraphListItem`

GetListItem returns the ListItem field if non-nil, zero value otherwise.

### GetListItemOk

`func (o *MicrosoftGraphSharedDriveItem) GetListItemOk() (AnyOfmicrosoftGraphListItem, bool)`

GetListItemOk returns a tuple with the ListItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListItem

`func (o *MicrosoftGraphSharedDriveItem) HasListItem() bool`

HasListItem returns a boolean if a field has been set.

### SetListItem

`func (o *MicrosoftGraphSharedDriveItem) SetListItem(v AnyOfmicrosoftGraphListItem)`

SetListItem gets a reference to the given AnyOfmicrosoftGraphListItem and assigns it to the ListItem field.

### SetListItemExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetListItemExplicitNull(b bool)`

SetListItemExplicitNull (un)sets ListItem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ListItem value is set to nil even if false is passed
### GetRoot

`func (o *MicrosoftGraphSharedDriveItem) GetRoot() AnyOfmicrosoftGraphDriveItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *MicrosoftGraphSharedDriveItem) GetRootOk() (AnyOfmicrosoftGraphDriveItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoot

`func (o *MicrosoftGraphSharedDriveItem) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRoot

`func (o *MicrosoftGraphSharedDriveItem) SetRoot(v AnyOfmicrosoftGraphDriveItem)`

SetRoot gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the Root field.

### SetRootExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetRootExplicitNull(b bool)`

SetRootExplicitNull (un)sets Root to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Root value is set to nil even if false is passed
### GetSite

`func (o *MicrosoftGraphSharedDriveItem) GetSite() AnyOfmicrosoftGraphSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MicrosoftGraphSharedDriveItem) GetSiteOk() (AnyOfmicrosoftGraphSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSite

`func (o *MicrosoftGraphSharedDriveItem) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSite

`func (o *MicrosoftGraphSharedDriveItem) SetSite(v AnyOfmicrosoftGraphSite)`

SetSite gets a reference to the given AnyOfmicrosoftGraphSite and assigns it to the Site field.

### SetSiteExplicitNull

`func (o *MicrosoftGraphSharedDriveItem) SetSiteExplicitNull(b bool)`

SetSiteExplicitNull (un)sets Site to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Site value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


