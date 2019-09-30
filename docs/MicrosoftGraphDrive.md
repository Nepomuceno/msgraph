# MicrosoftGraphDrive

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
**DriveType** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**Quota** | Pointer to [**AnyOfmicrosoftGraphQuota**](anyOf&lt;microsoft.graph.quota&gt;.md) |  | [optional] 
**SharePointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) |  | [optional] 
**System** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Items** | Pointer to [**[]MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md) |  | [optional] 
**List** | Pointer to [**AnyOfmicrosoftGraphList**](anyOf&lt;microsoft.graph.list&gt;.md) |  | [optional] 
**Root** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) |  | [optional] 
**Special** | Pointer to [**[]MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDrive) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDrive) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDrive) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDrive) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedBy

`func (o *MicrosoftGraphDrive) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphDrive) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphDrive) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphDrive) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphDrive) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphDrive) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphDrive) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphDrive) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphDrive) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphDrive) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDrive) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphDrive) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphDrive) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphDrive) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetETag

`func (o *MicrosoftGraphDrive) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MicrosoftGraphDrive) GetETagOk() (string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasETag

`func (o *MicrosoftGraphDrive) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETag

`func (o *MicrosoftGraphDrive) SetETag(v string)`

SetETag gets a reference to the given string and assigns it to the ETag field.

### SetETagExplicitNull

`func (o *MicrosoftGraphDrive) SetETagExplicitNull(b bool)`

SetETagExplicitNull (un)sets ETag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ETag value is set to nil even if false is passed
### GetLastModifiedBy

`func (o *MicrosoftGraphDrive) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphDrive) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphDrive) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphDrive) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphDrive) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphDrive) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphDrive) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphDrive) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphDrive) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetName

`func (o *MicrosoftGraphDrive) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphDrive) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphDrive) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphDrive) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphDrive) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetParentReference

`func (o *MicrosoftGraphDrive) GetParentReference() AnyOfmicrosoftGraphItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MicrosoftGraphDrive) GetParentReferenceOk() (AnyOfmicrosoftGraphItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentReference

`func (o *MicrosoftGraphDrive) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### SetParentReference

`func (o *MicrosoftGraphDrive) SetParentReference(v AnyOfmicrosoftGraphItemReference)`

SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.

### SetParentReferenceExplicitNull

`func (o *MicrosoftGraphDrive) SetParentReferenceExplicitNull(b bool)`

SetParentReferenceExplicitNull (un)sets ParentReference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentReference value is set to nil even if false is passed
### GetWebUrl

`func (o *MicrosoftGraphDrive) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphDrive) GetWebUrlOk() (string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebUrl

`func (o *MicrosoftGraphDrive) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrl

`func (o *MicrosoftGraphDrive) SetWebUrl(v string)`

SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.

### SetWebUrlExplicitNull

`func (o *MicrosoftGraphDrive) SetWebUrlExplicitNull(b bool)`

SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebUrl value is set to nil even if false is passed
### GetCreatedByUser

`func (o *MicrosoftGraphDrive) GetCreatedByUser() AnyOfmicrosoftGraphUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *MicrosoftGraphDrive) GetCreatedByUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedByUser

`func (o *MicrosoftGraphDrive) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUser

`func (o *MicrosoftGraphDrive) SetCreatedByUser(v AnyOfmicrosoftGraphUser)`

SetCreatedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the CreatedByUser field.

### SetCreatedByUserExplicitNull

`func (o *MicrosoftGraphDrive) SetCreatedByUserExplicitNull(b bool)`

SetCreatedByUserExplicitNull (un)sets CreatedByUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedByUser value is set to nil even if false is passed
### GetLastModifiedByUser

`func (o *MicrosoftGraphDrive) GetLastModifiedByUser() AnyOfmicrosoftGraphUser`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *MicrosoftGraphDrive) GetLastModifiedByUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedByUser

`func (o *MicrosoftGraphDrive) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.

### SetLastModifiedByUser

`func (o *MicrosoftGraphDrive) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser)`

SetLastModifiedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the LastModifiedByUser field.

### SetLastModifiedByUserExplicitNull

`func (o *MicrosoftGraphDrive) SetLastModifiedByUserExplicitNull(b bool)`

SetLastModifiedByUserExplicitNull (un)sets LastModifiedByUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedByUser value is set to nil even if false is passed
### GetDriveType

`func (o *MicrosoftGraphDrive) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *MicrosoftGraphDrive) GetDriveTypeOk() (string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDriveType

`func (o *MicrosoftGraphDrive) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### SetDriveType

`func (o *MicrosoftGraphDrive) SetDriveType(v string)`

SetDriveType gets a reference to the given string and assigns it to the DriveType field.

### SetDriveTypeExplicitNull

`func (o *MicrosoftGraphDrive) SetDriveTypeExplicitNull(b bool)`

SetDriveTypeExplicitNull (un)sets DriveType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DriveType value is set to nil even if false is passed
### GetOwner

`func (o *MicrosoftGraphDrive) GetOwner() AnyOfmicrosoftGraphIdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphDrive) GetOwnerOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *MicrosoftGraphDrive) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *MicrosoftGraphDrive) SetOwner(v AnyOfmicrosoftGraphIdentitySet)`

SetOwner gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *MicrosoftGraphDrive) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetQuota

`func (o *MicrosoftGraphDrive) GetQuota() AnyOfmicrosoftGraphQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *MicrosoftGraphDrive) GetQuotaOk() (AnyOfmicrosoftGraphQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuota

`func (o *MicrosoftGraphDrive) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuota

`func (o *MicrosoftGraphDrive) SetQuota(v AnyOfmicrosoftGraphQuota)`

SetQuota gets a reference to the given AnyOfmicrosoftGraphQuota and assigns it to the Quota field.

### SetQuotaExplicitNull

`func (o *MicrosoftGraphDrive) SetQuotaExplicitNull(b bool)`

SetQuotaExplicitNull (un)sets Quota to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Quota value is set to nil even if false is passed
### GetSharePointIds

`func (o *MicrosoftGraphDrive) GetSharePointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharePointIds returns the SharePointIds field if non-nil, zero value otherwise.

### GetSharePointIdsOk

`func (o *MicrosoftGraphDrive) GetSharePointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharePointIdsOk returns a tuple with the SharePointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharePointIds

`func (o *MicrosoftGraphDrive) HasSharePointIds() bool`

HasSharePointIds returns a boolean if a field has been set.

### SetSharePointIds

`func (o *MicrosoftGraphDrive) SetSharePointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharePointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharePointIds field.

### SetSharePointIdsExplicitNull

`func (o *MicrosoftGraphDrive) SetSharePointIdsExplicitNull(b bool)`

SetSharePointIdsExplicitNull (un)sets SharePointIds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharePointIds value is set to nil even if false is passed
### GetSystem

`func (o *MicrosoftGraphDrive) GetSystem() AnyOfobject`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *MicrosoftGraphDrive) GetSystemOk() (AnyOfobject, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSystem

`func (o *MicrosoftGraphDrive) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystem

`func (o *MicrosoftGraphDrive) SetSystem(v AnyOfobject)`

SetSystem gets a reference to the given AnyOfobject and assigns it to the System field.

### SetSystemExplicitNull

`func (o *MicrosoftGraphDrive) SetSystemExplicitNull(b bool)`

SetSystemExplicitNull (un)sets System to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The System value is set to nil even if false is passed
### GetItems

`func (o *MicrosoftGraphDrive) GetItems() []MicrosoftGraphDriveItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MicrosoftGraphDrive) GetItemsOk() ([]MicrosoftGraphDriveItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItems

`func (o *MicrosoftGraphDrive) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItems

`func (o *MicrosoftGraphDrive) SetItems(v []MicrosoftGraphDriveItem)`

SetItems gets a reference to the given []MicrosoftGraphDriveItem and assigns it to the Items field.

### GetList

`func (o *MicrosoftGraphDrive) GetList() AnyOfmicrosoftGraphList`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *MicrosoftGraphDrive) GetListOk() (AnyOfmicrosoftGraphList, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasList

`func (o *MicrosoftGraphDrive) HasList() bool`

HasList returns a boolean if a field has been set.

### SetList

`func (o *MicrosoftGraphDrive) SetList(v AnyOfmicrosoftGraphList)`

SetList gets a reference to the given AnyOfmicrosoftGraphList and assigns it to the List field.

### SetListExplicitNull

`func (o *MicrosoftGraphDrive) SetListExplicitNull(b bool)`

SetListExplicitNull (un)sets List to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The List value is set to nil even if false is passed
### GetRoot

`func (o *MicrosoftGraphDrive) GetRoot() AnyOfmicrosoftGraphDriveItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *MicrosoftGraphDrive) GetRootOk() (AnyOfmicrosoftGraphDriveItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoot

`func (o *MicrosoftGraphDrive) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRoot

`func (o *MicrosoftGraphDrive) SetRoot(v AnyOfmicrosoftGraphDriveItem)`

SetRoot gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the Root field.

### SetRootExplicitNull

`func (o *MicrosoftGraphDrive) SetRootExplicitNull(b bool)`

SetRootExplicitNull (un)sets Root to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Root value is set to nil even if false is passed
### GetSpecial

`func (o *MicrosoftGraphDrive) GetSpecial() []MicrosoftGraphDriveItem`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *MicrosoftGraphDrive) GetSpecialOk() ([]MicrosoftGraphDriveItem, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecial

`func (o *MicrosoftGraphDrive) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.

### SetSpecial

`func (o *MicrosoftGraphDrive) SetSpecial(v []MicrosoftGraphDriveItem)`

SetSpecial gets a reference to the given []MicrosoftGraphDriveItem and assigns it to the Special field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


