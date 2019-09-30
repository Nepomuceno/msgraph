# MicrosoftGraphList

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
**List** | Pointer to [**AnyOfmicrosoftGraphListInfo**](anyOf&lt;microsoft.graph.listInfo&gt;.md) |  | [optional] 
**SharepointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) |  | [optional] 
**System** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Columns** | Pointer to [**[]MicrosoftGraphColumnDefinition**](microsoft.graph.columnDefinition.md) |  | [optional] 
**ContentTypes** | Pointer to [**[]MicrosoftGraphContentType**](microsoft.graph.contentType.md) |  | [optional] 
**Drive** | Pointer to [**AnyOfmicrosoftGraphDrive**](anyOf&lt;microsoft.graph.drive&gt;.md) |  | [optional] 
**Items** | Pointer to [**[]MicrosoftGraphListItem**](microsoft.graph.listItem.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphList) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphList) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphList) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedBy

`func (o *MicrosoftGraphList) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphList) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphList) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphList) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphList) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphList) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphList) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphList) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphList) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphList) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphList) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphList) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetETag

`func (o *MicrosoftGraphList) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MicrosoftGraphList) GetETagOk() (string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasETag

`func (o *MicrosoftGraphList) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETag

`func (o *MicrosoftGraphList) SetETag(v string)`

SetETag gets a reference to the given string and assigns it to the ETag field.

### SetETagExplicitNull

`func (o *MicrosoftGraphList) SetETagExplicitNull(b bool)`

SetETagExplicitNull (un)sets ETag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ETag value is set to nil even if false is passed
### GetLastModifiedBy

`func (o *MicrosoftGraphList) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphList) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphList) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphList) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphList) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphList) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphList) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphList) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphList) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetName

`func (o *MicrosoftGraphList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphList) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphList) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphList) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphList) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetParentReference

`func (o *MicrosoftGraphList) GetParentReference() AnyOfmicrosoftGraphItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MicrosoftGraphList) GetParentReferenceOk() (AnyOfmicrosoftGraphItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentReference

`func (o *MicrosoftGraphList) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### SetParentReference

`func (o *MicrosoftGraphList) SetParentReference(v AnyOfmicrosoftGraphItemReference)`

SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.

### SetParentReferenceExplicitNull

`func (o *MicrosoftGraphList) SetParentReferenceExplicitNull(b bool)`

SetParentReferenceExplicitNull (un)sets ParentReference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentReference value is set to nil even if false is passed
### GetWebUrl

`func (o *MicrosoftGraphList) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphList) GetWebUrlOk() (string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebUrl

`func (o *MicrosoftGraphList) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrl

`func (o *MicrosoftGraphList) SetWebUrl(v string)`

SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.

### SetWebUrlExplicitNull

`func (o *MicrosoftGraphList) SetWebUrlExplicitNull(b bool)`

SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebUrl value is set to nil even if false is passed
### GetCreatedByUser

`func (o *MicrosoftGraphList) GetCreatedByUser() AnyOfmicrosoftGraphUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *MicrosoftGraphList) GetCreatedByUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedByUser

`func (o *MicrosoftGraphList) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUser

`func (o *MicrosoftGraphList) SetCreatedByUser(v AnyOfmicrosoftGraphUser)`

SetCreatedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the CreatedByUser field.

### SetCreatedByUserExplicitNull

`func (o *MicrosoftGraphList) SetCreatedByUserExplicitNull(b bool)`

SetCreatedByUserExplicitNull (un)sets CreatedByUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedByUser value is set to nil even if false is passed
### GetLastModifiedByUser

`func (o *MicrosoftGraphList) GetLastModifiedByUser() AnyOfmicrosoftGraphUser`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *MicrosoftGraphList) GetLastModifiedByUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedByUser

`func (o *MicrosoftGraphList) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.

### SetLastModifiedByUser

`func (o *MicrosoftGraphList) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser)`

SetLastModifiedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the LastModifiedByUser field.

### SetLastModifiedByUserExplicitNull

`func (o *MicrosoftGraphList) SetLastModifiedByUserExplicitNull(b bool)`

SetLastModifiedByUserExplicitNull (un)sets LastModifiedByUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedByUser value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphList) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphList) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphList) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphList) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphList) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetList

`func (o *MicrosoftGraphList) GetList() AnyOfmicrosoftGraphListInfo`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *MicrosoftGraphList) GetListOk() (AnyOfmicrosoftGraphListInfo, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasList

`func (o *MicrosoftGraphList) HasList() bool`

HasList returns a boolean if a field has been set.

### SetList

`func (o *MicrosoftGraphList) SetList(v AnyOfmicrosoftGraphListInfo)`

SetList gets a reference to the given AnyOfmicrosoftGraphListInfo and assigns it to the List field.

### SetListExplicitNull

`func (o *MicrosoftGraphList) SetListExplicitNull(b bool)`

SetListExplicitNull (un)sets List to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The List value is set to nil even if false is passed
### GetSharepointIds

`func (o *MicrosoftGraphList) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *MicrosoftGraphList) GetSharepointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharepointIds

`func (o *MicrosoftGraphList) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIds

`func (o *MicrosoftGraphList) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.

### SetSharepointIdsExplicitNull

`func (o *MicrosoftGraphList) SetSharepointIdsExplicitNull(b bool)`

SetSharepointIdsExplicitNull (un)sets SharepointIds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharepointIds value is set to nil even if false is passed
### GetSystem

`func (o *MicrosoftGraphList) GetSystem() AnyOfobject`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *MicrosoftGraphList) GetSystemOk() (AnyOfobject, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSystem

`func (o *MicrosoftGraphList) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystem

`func (o *MicrosoftGraphList) SetSystem(v AnyOfobject)`

SetSystem gets a reference to the given AnyOfobject and assigns it to the System field.

### SetSystemExplicitNull

`func (o *MicrosoftGraphList) SetSystemExplicitNull(b bool)`

SetSystemExplicitNull (un)sets System to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The System value is set to nil even if false is passed
### GetColumns

`func (o *MicrosoftGraphList) GetColumns() []MicrosoftGraphColumnDefinition`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *MicrosoftGraphList) GetColumnsOk() ([]MicrosoftGraphColumnDefinition, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumns

`func (o *MicrosoftGraphList) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumns

`func (o *MicrosoftGraphList) SetColumns(v []MicrosoftGraphColumnDefinition)`

SetColumns gets a reference to the given []MicrosoftGraphColumnDefinition and assigns it to the Columns field.

### GetContentTypes

`func (o *MicrosoftGraphList) GetContentTypes() []MicrosoftGraphContentType`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *MicrosoftGraphList) GetContentTypesOk() ([]MicrosoftGraphContentType, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentTypes

`func (o *MicrosoftGraphList) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### SetContentTypes

`func (o *MicrosoftGraphList) SetContentTypes(v []MicrosoftGraphContentType)`

SetContentTypes gets a reference to the given []MicrosoftGraphContentType and assigns it to the ContentTypes field.

### GetDrive

`func (o *MicrosoftGraphList) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *MicrosoftGraphList) GetDriveOk() (AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrive

`func (o *MicrosoftGraphList) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDrive

`func (o *MicrosoftGraphList) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive gets a reference to the given AnyOfmicrosoftGraphDrive and assigns it to the Drive field.

### SetDriveExplicitNull

`func (o *MicrosoftGraphList) SetDriveExplicitNull(b bool)`

SetDriveExplicitNull (un)sets Drive to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Drive value is set to nil even if false is passed
### GetItems

`func (o *MicrosoftGraphList) GetItems() []MicrosoftGraphListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MicrosoftGraphList) GetItemsOk() ([]MicrosoftGraphListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItems

`func (o *MicrosoftGraphList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItems

`func (o *MicrosoftGraphList) SetItems(v []MicrosoftGraphListItem)`

SetItems gets a reference to the given []MicrosoftGraphListItem and assigns it to the Items field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


