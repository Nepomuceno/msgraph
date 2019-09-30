# MicrosoftGraphRemoteItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**File** | Pointer to [**AnyOfmicrosoftGraphFile**](anyOf&lt;microsoft.graph.file&gt;.md) |  | [optional] 
**FileSystemInfo** | Pointer to [**AnyOfmicrosoftGraphFileSystemInfo**](anyOf&lt;microsoft.graph.fileSystemInfo&gt;.md) |  | [optional] 
**Folder** | Pointer to [**AnyOfmicrosoftGraphFolder**](anyOf&lt;microsoft.graph.folder&gt;.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Package** | Pointer to [**AnyOfmicrosoftGraphPackage**](anyOf&lt;microsoft.graph.package&gt;.md) |  | [optional] 
**ParentReference** | Pointer to [**AnyOfmicrosoftGraphItemReference**](anyOf&lt;microsoft.graph.itemReference&gt;.md) |  | [optional] 
**Shared** | Pointer to [**AnyOfmicrosoftGraphShared**](anyOf&lt;microsoft.graph.shared&gt;.md) |  | [optional] 
**SharepointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**SpecialFolder** | Pointer to [**AnyOfmicrosoftGraphSpecialFolder**](anyOf&lt;microsoft.graph.specialFolder&gt;.md) |  | [optional] 
**WebDavUrl** | Pointer to **string** |  | [optional] 
**WebUrl** | Pointer to **string** |  | [optional] 

## Methods

### GetCreatedBy

`func (o *MicrosoftGraphRemoteItem) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphRemoteItem) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphRemoteItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphRemoteItem) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphRemoteItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphRemoteItem) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphRemoteItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphRemoteItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetFile

`func (o *MicrosoftGraphRemoteItem) GetFile() AnyOfmicrosoftGraphFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *MicrosoftGraphRemoteItem) GetFileOk() (AnyOfmicrosoftGraphFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFile

`func (o *MicrosoftGraphRemoteItem) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFile

`func (o *MicrosoftGraphRemoteItem) SetFile(v AnyOfmicrosoftGraphFile)`

SetFile gets a reference to the given AnyOfmicrosoftGraphFile and assigns it to the File field.

### SetFileExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetFileExplicitNull(b bool)`

SetFileExplicitNull (un)sets File to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The File value is set to nil even if false is passed
### GetFileSystemInfo

`func (o *MicrosoftGraphRemoteItem) GetFileSystemInfo() AnyOfmicrosoftGraphFileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *MicrosoftGraphRemoteItem) GetFileSystemInfoOk() (AnyOfmicrosoftGraphFileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileSystemInfo

`func (o *MicrosoftGraphRemoteItem) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### SetFileSystemInfo

`func (o *MicrosoftGraphRemoteItem) SetFileSystemInfo(v AnyOfmicrosoftGraphFileSystemInfo)`

SetFileSystemInfo gets a reference to the given AnyOfmicrosoftGraphFileSystemInfo and assigns it to the FileSystemInfo field.

### SetFileSystemInfoExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetFileSystemInfoExplicitNull(b bool)`

SetFileSystemInfoExplicitNull (un)sets FileSystemInfo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileSystemInfo value is set to nil even if false is passed
### GetFolder

`func (o *MicrosoftGraphRemoteItem) GetFolder() AnyOfmicrosoftGraphFolder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *MicrosoftGraphRemoteItem) GetFolderOk() (AnyOfmicrosoftGraphFolder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFolder

`func (o *MicrosoftGraphRemoteItem) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolder

`func (o *MicrosoftGraphRemoteItem) SetFolder(v AnyOfmicrosoftGraphFolder)`

SetFolder gets a reference to the given AnyOfmicrosoftGraphFolder and assigns it to the Folder field.

### SetFolderExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetFolderExplicitNull(b bool)`

SetFolderExplicitNull (un)sets Folder to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Folder value is set to nil even if false is passed
### GetId

`func (o *MicrosoftGraphRemoteItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphRemoteItem) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphRemoteItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphRemoteItem) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### SetIdExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetIdExplicitNull(b bool)`

SetIdExplicitNull (un)sets Id to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Id value is set to nil even if false is passed
### GetLastModifiedBy

`func (o *MicrosoftGraphRemoteItem) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphRemoteItem) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphRemoteItem) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphRemoteItem) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphRemoteItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphRemoteItem) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphRemoteItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphRemoteItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetName

`func (o *MicrosoftGraphRemoteItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphRemoteItem) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphRemoteItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphRemoteItem) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetPackage

`func (o *MicrosoftGraphRemoteItem) GetPackage() AnyOfmicrosoftGraphPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *MicrosoftGraphRemoteItem) GetPackageOk() (AnyOfmicrosoftGraphPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPackage

`func (o *MicrosoftGraphRemoteItem) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### SetPackage

`func (o *MicrosoftGraphRemoteItem) SetPackage(v AnyOfmicrosoftGraphPackage)`

SetPackage gets a reference to the given AnyOfmicrosoftGraphPackage and assigns it to the Package field.

### SetPackageExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetPackageExplicitNull(b bool)`

SetPackageExplicitNull (un)sets Package to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Package value is set to nil even if false is passed
### GetParentReference

`func (o *MicrosoftGraphRemoteItem) GetParentReference() AnyOfmicrosoftGraphItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MicrosoftGraphRemoteItem) GetParentReferenceOk() (AnyOfmicrosoftGraphItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentReference

`func (o *MicrosoftGraphRemoteItem) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### SetParentReference

`func (o *MicrosoftGraphRemoteItem) SetParentReference(v AnyOfmicrosoftGraphItemReference)`

SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.

### SetParentReferenceExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetParentReferenceExplicitNull(b bool)`

SetParentReferenceExplicitNull (un)sets ParentReference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentReference value is set to nil even if false is passed
### GetShared

`func (o *MicrosoftGraphRemoteItem) GetShared() AnyOfmicrosoftGraphShared`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *MicrosoftGraphRemoteItem) GetSharedOk() (AnyOfmicrosoftGraphShared, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShared

`func (o *MicrosoftGraphRemoteItem) HasShared() bool`

HasShared returns a boolean if a field has been set.

### SetShared

`func (o *MicrosoftGraphRemoteItem) SetShared(v AnyOfmicrosoftGraphShared)`

SetShared gets a reference to the given AnyOfmicrosoftGraphShared and assigns it to the Shared field.

### SetSharedExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetSharedExplicitNull(b bool)`

SetSharedExplicitNull (un)sets Shared to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Shared value is set to nil even if false is passed
### GetSharepointIds

`func (o *MicrosoftGraphRemoteItem) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *MicrosoftGraphRemoteItem) GetSharepointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharepointIds

`func (o *MicrosoftGraphRemoteItem) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIds

`func (o *MicrosoftGraphRemoteItem) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.

### SetSharepointIdsExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetSharepointIdsExplicitNull(b bool)`

SetSharepointIdsExplicitNull (un)sets SharepointIds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharepointIds value is set to nil even if false is passed
### GetSize

`func (o *MicrosoftGraphRemoteItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphRemoteItem) GetSizeOk() (int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *MicrosoftGraphRemoteItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *MicrosoftGraphRemoteItem) SetSize(v int64)`

SetSize gets a reference to the given int64 and assigns it to the Size field.

### SetSizeExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetSizeExplicitNull(b bool)`

SetSizeExplicitNull (un)sets Size to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Size value is set to nil even if false is passed
### GetSpecialFolder

`func (o *MicrosoftGraphRemoteItem) GetSpecialFolder() AnyOfmicrosoftGraphSpecialFolder`

GetSpecialFolder returns the SpecialFolder field if non-nil, zero value otherwise.

### GetSpecialFolderOk

`func (o *MicrosoftGraphRemoteItem) GetSpecialFolderOk() (AnyOfmicrosoftGraphSpecialFolder, bool)`

GetSpecialFolderOk returns a tuple with the SpecialFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecialFolder

`func (o *MicrosoftGraphRemoteItem) HasSpecialFolder() bool`

HasSpecialFolder returns a boolean if a field has been set.

### SetSpecialFolder

`func (o *MicrosoftGraphRemoteItem) SetSpecialFolder(v AnyOfmicrosoftGraphSpecialFolder)`

SetSpecialFolder gets a reference to the given AnyOfmicrosoftGraphSpecialFolder and assigns it to the SpecialFolder field.

### SetSpecialFolderExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetSpecialFolderExplicitNull(b bool)`

SetSpecialFolderExplicitNull (un)sets SpecialFolder to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SpecialFolder value is set to nil even if false is passed
### GetWebDavUrl

`func (o *MicrosoftGraphRemoteItem) GetWebDavUrl() string`

GetWebDavUrl returns the WebDavUrl field if non-nil, zero value otherwise.

### GetWebDavUrlOk

`func (o *MicrosoftGraphRemoteItem) GetWebDavUrlOk() (string, bool)`

GetWebDavUrlOk returns a tuple with the WebDavUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebDavUrl

`func (o *MicrosoftGraphRemoteItem) HasWebDavUrl() bool`

HasWebDavUrl returns a boolean if a field has been set.

### SetWebDavUrl

`func (o *MicrosoftGraphRemoteItem) SetWebDavUrl(v string)`

SetWebDavUrl gets a reference to the given string and assigns it to the WebDavUrl field.

### SetWebDavUrlExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetWebDavUrlExplicitNull(b bool)`

SetWebDavUrlExplicitNull (un)sets WebDavUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebDavUrl value is set to nil even if false is passed
### GetWebUrl

`func (o *MicrosoftGraphRemoteItem) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphRemoteItem) GetWebUrlOk() (string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebUrl

`func (o *MicrosoftGraphRemoteItem) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrl

`func (o *MicrosoftGraphRemoteItem) SetWebUrl(v string)`

SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.

### SetWebUrlExplicitNull

`func (o *MicrosoftGraphRemoteItem) SetWebUrlExplicitNull(b bool)`

SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebUrl value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


