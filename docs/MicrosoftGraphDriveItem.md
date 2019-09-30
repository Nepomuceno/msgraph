# MicrosoftGraphDriveItem

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
**Audio** | Pointer to [**AnyOfmicrosoftGraphAudio**](anyOf&lt;microsoft.graph.audio&gt;.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**CTag** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to [**AnyOfmicrosoftGraphDeleted**](anyOf&lt;microsoft.graph.deleted&gt;.md) |  | [optional] 
**File** | Pointer to [**AnyOfmicrosoftGraphFile**](anyOf&lt;microsoft.graph.file&gt;.md) |  | [optional] 
**FileSystemInfo** | Pointer to [**AnyOfmicrosoftGraphFileSystemInfo**](anyOf&lt;microsoft.graph.fileSystemInfo&gt;.md) |  | [optional] 
**Folder** | Pointer to [**AnyOfmicrosoftGraphFolder**](anyOf&lt;microsoft.graph.folder&gt;.md) |  | [optional] 
**Image** | Pointer to [**AnyOfmicrosoftGraphImage**](anyOf&lt;microsoft.graph.image&gt;.md) |  | [optional] 
**Location** | Pointer to [**AnyOfmicrosoftGraphGeoCoordinates**](anyOf&lt;microsoft.graph.geoCoordinates&gt;.md) |  | [optional] 
**Package** | Pointer to [**AnyOfmicrosoftGraphPackage**](anyOf&lt;microsoft.graph.package&gt;.md) |  | [optional] 
**Photo** | Pointer to [**AnyOfmicrosoftGraphPhoto**](anyOf&lt;microsoft.graph.photo&gt;.md) |  | [optional] 
**Publication** | Pointer to [**AnyOfmicrosoftGraphPublicationFacet**](anyOf&lt;microsoft.graph.publicationFacet&gt;.md) |  | [optional] 
**RemoteItem** | Pointer to [**AnyOfmicrosoftGraphRemoteItem**](anyOf&lt;microsoft.graph.remoteItem&gt;.md) |  | [optional] 
**Root** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**SearchResult** | Pointer to [**AnyOfmicrosoftGraphSearchResult**](anyOf&lt;microsoft.graph.searchResult&gt;.md) |  | [optional] 
**Shared** | Pointer to [**AnyOfmicrosoftGraphShared**](anyOf&lt;microsoft.graph.shared&gt;.md) |  | [optional] 
**SharepointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**SpecialFolder** | Pointer to [**AnyOfmicrosoftGraphSpecialFolder**](anyOf&lt;microsoft.graph.specialFolder&gt;.md) |  | [optional] 
**Video** | Pointer to [**AnyOfmicrosoftGraphVideo**](anyOf&lt;microsoft.graph.video&gt;.md) |  | [optional] 
**WebDavUrl** | Pointer to **string** |  | [optional] 
**Workbook** | Pointer to [**AnyOfmicrosoftGraphWorkbook**](anyOf&lt;microsoft.graph.workbook&gt;.md) |  | [optional] 
**Analytics** | Pointer to [**AnyOfmicrosoftGraphItemAnalytics**](anyOf&lt;microsoft.graph.itemAnalytics&gt;.md) |  | [optional] 
**Children** | Pointer to [**[]MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md) |  | [optional] 
**ListItem** | Pointer to [**AnyOfmicrosoftGraphListItem**](anyOf&lt;microsoft.graph.listItem&gt;.md) |  | [optional] 
**Permissions** | Pointer to [**[]MicrosoftGraphPermission**](microsoft.graph.permission.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]MicrosoftGraphSubscription**](microsoft.graph.subscription.md) |  | [optional] 
**Thumbnails** | Pointer to [**[]MicrosoftGraphThumbnailSet**](microsoft.graph.thumbnailSet.md) |  | [optional] 
**Versions** | Pointer to [**[]MicrosoftGraphDriveItemVersion**](microsoft.graph.driveItemVersion.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDriveItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDriveItem) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDriveItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDriveItem) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedBy

`func (o *MicrosoftGraphDriveItem) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphDriveItem) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphDriveItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphDriveItem) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphDriveItem) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphDriveItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphDriveItem) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphDriveItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphDriveItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphDriveItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDriveItem) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphDriveItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphDriveItem) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphDriveItem) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetETag

`func (o *MicrosoftGraphDriveItem) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MicrosoftGraphDriveItem) GetETagOk() (string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasETag

`func (o *MicrosoftGraphDriveItem) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETag

`func (o *MicrosoftGraphDriveItem) SetETag(v string)`

SetETag gets a reference to the given string and assigns it to the ETag field.

### SetETagExplicitNull

`func (o *MicrosoftGraphDriveItem) SetETagExplicitNull(b bool)`

SetETagExplicitNull (un)sets ETag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ETag value is set to nil even if false is passed
### GetLastModifiedBy

`func (o *MicrosoftGraphDriveItem) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphDriveItem) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphDriveItem) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphDriveItem) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphDriveItem) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphDriveItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphDriveItem) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphDriveItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphDriveItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetName

`func (o *MicrosoftGraphDriveItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphDriveItem) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphDriveItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphDriveItem) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphDriveItem) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetParentReference

`func (o *MicrosoftGraphDriveItem) GetParentReference() AnyOfmicrosoftGraphItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MicrosoftGraphDriveItem) GetParentReferenceOk() (AnyOfmicrosoftGraphItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentReference

`func (o *MicrosoftGraphDriveItem) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### SetParentReference

`func (o *MicrosoftGraphDriveItem) SetParentReference(v AnyOfmicrosoftGraphItemReference)`

SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.

### SetParentReferenceExplicitNull

`func (o *MicrosoftGraphDriveItem) SetParentReferenceExplicitNull(b bool)`

SetParentReferenceExplicitNull (un)sets ParentReference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentReference value is set to nil even if false is passed
### GetWebUrl

`func (o *MicrosoftGraphDriveItem) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphDriveItem) GetWebUrlOk() (string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebUrl

`func (o *MicrosoftGraphDriveItem) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrl

`func (o *MicrosoftGraphDriveItem) SetWebUrl(v string)`

SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.

### SetWebUrlExplicitNull

`func (o *MicrosoftGraphDriveItem) SetWebUrlExplicitNull(b bool)`

SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebUrl value is set to nil even if false is passed
### GetCreatedByUser

`func (o *MicrosoftGraphDriveItem) GetCreatedByUser() AnyOfmicrosoftGraphUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *MicrosoftGraphDriveItem) GetCreatedByUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedByUser

`func (o *MicrosoftGraphDriveItem) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUser

`func (o *MicrosoftGraphDriveItem) SetCreatedByUser(v AnyOfmicrosoftGraphUser)`

SetCreatedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the CreatedByUser field.

### SetCreatedByUserExplicitNull

`func (o *MicrosoftGraphDriveItem) SetCreatedByUserExplicitNull(b bool)`

SetCreatedByUserExplicitNull (un)sets CreatedByUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedByUser value is set to nil even if false is passed
### GetLastModifiedByUser

`func (o *MicrosoftGraphDriveItem) GetLastModifiedByUser() AnyOfmicrosoftGraphUser`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *MicrosoftGraphDriveItem) GetLastModifiedByUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedByUser

`func (o *MicrosoftGraphDriveItem) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.

### SetLastModifiedByUser

`func (o *MicrosoftGraphDriveItem) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser)`

SetLastModifiedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the LastModifiedByUser field.

### SetLastModifiedByUserExplicitNull

`func (o *MicrosoftGraphDriveItem) SetLastModifiedByUserExplicitNull(b bool)`

SetLastModifiedByUserExplicitNull (un)sets LastModifiedByUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedByUser value is set to nil even if false is passed
### GetAudio

`func (o *MicrosoftGraphDriveItem) GetAudio() AnyOfmicrosoftGraphAudio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *MicrosoftGraphDriveItem) GetAudioOk() (AnyOfmicrosoftGraphAudio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudio

`func (o *MicrosoftGraphDriveItem) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### SetAudio

`func (o *MicrosoftGraphDriveItem) SetAudio(v AnyOfmicrosoftGraphAudio)`

SetAudio gets a reference to the given AnyOfmicrosoftGraphAudio and assigns it to the Audio field.

### SetAudioExplicitNull

`func (o *MicrosoftGraphDriveItem) SetAudioExplicitNull(b bool)`

SetAudioExplicitNull (un)sets Audio to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Audio value is set to nil even if false is passed
### GetContent

`func (o *MicrosoftGraphDriveItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphDriveItem) GetContentOk() (string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContent

`func (o *MicrosoftGraphDriveItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContent

`func (o *MicrosoftGraphDriveItem) SetContent(v string)`

SetContent gets a reference to the given string and assigns it to the Content field.

### SetContentExplicitNull

`func (o *MicrosoftGraphDriveItem) SetContentExplicitNull(b bool)`

SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Content value is set to nil even if false is passed
### GetCTag

`func (o *MicrosoftGraphDriveItem) GetCTag() string`

GetCTag returns the CTag field if non-nil, zero value otherwise.

### GetCTagOk

`func (o *MicrosoftGraphDriveItem) GetCTagOk() (string, bool)`

GetCTagOk returns a tuple with the CTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCTag

`func (o *MicrosoftGraphDriveItem) HasCTag() bool`

HasCTag returns a boolean if a field has been set.

### SetCTag

`func (o *MicrosoftGraphDriveItem) SetCTag(v string)`

SetCTag gets a reference to the given string and assigns it to the CTag field.

### SetCTagExplicitNull

`func (o *MicrosoftGraphDriveItem) SetCTagExplicitNull(b bool)`

SetCTagExplicitNull (un)sets CTag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CTag value is set to nil even if false is passed
### GetDeleted

`func (o *MicrosoftGraphDriveItem) GetDeleted() AnyOfmicrosoftGraphDeleted`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *MicrosoftGraphDriveItem) GetDeletedOk() (AnyOfmicrosoftGraphDeleted, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleted

`func (o *MicrosoftGraphDriveItem) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeleted

`func (o *MicrosoftGraphDriveItem) SetDeleted(v AnyOfmicrosoftGraphDeleted)`

SetDeleted gets a reference to the given AnyOfmicrosoftGraphDeleted and assigns it to the Deleted field.

### SetDeletedExplicitNull

`func (o *MicrosoftGraphDriveItem) SetDeletedExplicitNull(b bool)`

SetDeletedExplicitNull (un)sets Deleted to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Deleted value is set to nil even if false is passed
### GetFile

`func (o *MicrosoftGraphDriveItem) GetFile() AnyOfmicrosoftGraphFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *MicrosoftGraphDriveItem) GetFileOk() (AnyOfmicrosoftGraphFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFile

`func (o *MicrosoftGraphDriveItem) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFile

`func (o *MicrosoftGraphDriveItem) SetFile(v AnyOfmicrosoftGraphFile)`

SetFile gets a reference to the given AnyOfmicrosoftGraphFile and assigns it to the File field.

### SetFileExplicitNull

`func (o *MicrosoftGraphDriveItem) SetFileExplicitNull(b bool)`

SetFileExplicitNull (un)sets File to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The File value is set to nil even if false is passed
### GetFileSystemInfo

`func (o *MicrosoftGraphDriveItem) GetFileSystemInfo() AnyOfmicrosoftGraphFileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *MicrosoftGraphDriveItem) GetFileSystemInfoOk() (AnyOfmicrosoftGraphFileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileSystemInfo

`func (o *MicrosoftGraphDriveItem) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### SetFileSystemInfo

`func (o *MicrosoftGraphDriveItem) SetFileSystemInfo(v AnyOfmicrosoftGraphFileSystemInfo)`

SetFileSystemInfo gets a reference to the given AnyOfmicrosoftGraphFileSystemInfo and assigns it to the FileSystemInfo field.

### SetFileSystemInfoExplicitNull

`func (o *MicrosoftGraphDriveItem) SetFileSystemInfoExplicitNull(b bool)`

SetFileSystemInfoExplicitNull (un)sets FileSystemInfo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileSystemInfo value is set to nil even if false is passed
### GetFolder

`func (o *MicrosoftGraphDriveItem) GetFolder() AnyOfmicrosoftGraphFolder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *MicrosoftGraphDriveItem) GetFolderOk() (AnyOfmicrosoftGraphFolder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFolder

`func (o *MicrosoftGraphDriveItem) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolder

`func (o *MicrosoftGraphDriveItem) SetFolder(v AnyOfmicrosoftGraphFolder)`

SetFolder gets a reference to the given AnyOfmicrosoftGraphFolder and assigns it to the Folder field.

### SetFolderExplicitNull

`func (o *MicrosoftGraphDriveItem) SetFolderExplicitNull(b bool)`

SetFolderExplicitNull (un)sets Folder to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Folder value is set to nil even if false is passed
### GetImage

`func (o *MicrosoftGraphDriveItem) GetImage() AnyOfmicrosoftGraphImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *MicrosoftGraphDriveItem) GetImageOk() (AnyOfmicrosoftGraphImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImage

`func (o *MicrosoftGraphDriveItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImage

`func (o *MicrosoftGraphDriveItem) SetImage(v AnyOfmicrosoftGraphImage)`

SetImage gets a reference to the given AnyOfmicrosoftGraphImage and assigns it to the Image field.

### SetImageExplicitNull

`func (o *MicrosoftGraphDriveItem) SetImageExplicitNull(b bool)`

SetImageExplicitNull (un)sets Image to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Image value is set to nil even if false is passed
### GetLocation

`func (o *MicrosoftGraphDriveItem) GetLocation() AnyOfmicrosoftGraphGeoCoordinates`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphDriveItem) GetLocationOk() (AnyOfmicrosoftGraphGeoCoordinates, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocation

`func (o *MicrosoftGraphDriveItem) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocation

`func (o *MicrosoftGraphDriveItem) SetLocation(v AnyOfmicrosoftGraphGeoCoordinates)`

SetLocation gets a reference to the given AnyOfmicrosoftGraphGeoCoordinates and assigns it to the Location field.

### SetLocationExplicitNull

`func (o *MicrosoftGraphDriveItem) SetLocationExplicitNull(b bool)`

SetLocationExplicitNull (un)sets Location to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Location value is set to nil even if false is passed
### GetPackage

`func (o *MicrosoftGraphDriveItem) GetPackage() AnyOfmicrosoftGraphPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *MicrosoftGraphDriveItem) GetPackageOk() (AnyOfmicrosoftGraphPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPackage

`func (o *MicrosoftGraphDriveItem) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### SetPackage

`func (o *MicrosoftGraphDriveItem) SetPackage(v AnyOfmicrosoftGraphPackage)`

SetPackage gets a reference to the given AnyOfmicrosoftGraphPackage and assigns it to the Package field.

### SetPackageExplicitNull

`func (o *MicrosoftGraphDriveItem) SetPackageExplicitNull(b bool)`

SetPackageExplicitNull (un)sets Package to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Package value is set to nil even if false is passed
### GetPhoto

`func (o *MicrosoftGraphDriveItem) GetPhoto() AnyOfmicrosoftGraphPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MicrosoftGraphDriveItem) GetPhotoOk() (AnyOfmicrosoftGraphPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhoto

`func (o *MicrosoftGraphDriveItem) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhoto

`func (o *MicrosoftGraphDriveItem) SetPhoto(v AnyOfmicrosoftGraphPhoto)`

SetPhoto gets a reference to the given AnyOfmicrosoftGraphPhoto and assigns it to the Photo field.

### SetPhotoExplicitNull

`func (o *MicrosoftGraphDriveItem) SetPhotoExplicitNull(b bool)`

SetPhotoExplicitNull (un)sets Photo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Photo value is set to nil even if false is passed
### GetPublication

`func (o *MicrosoftGraphDriveItem) GetPublication() AnyOfmicrosoftGraphPublicationFacet`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *MicrosoftGraphDriveItem) GetPublicationOk() (AnyOfmicrosoftGraphPublicationFacet, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublication

`func (o *MicrosoftGraphDriveItem) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublication

`func (o *MicrosoftGraphDriveItem) SetPublication(v AnyOfmicrosoftGraphPublicationFacet)`

SetPublication gets a reference to the given AnyOfmicrosoftGraphPublicationFacet and assigns it to the Publication field.

### SetPublicationExplicitNull

`func (o *MicrosoftGraphDriveItem) SetPublicationExplicitNull(b bool)`

SetPublicationExplicitNull (un)sets Publication to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publication value is set to nil even if false is passed
### GetRemoteItem

`func (o *MicrosoftGraphDriveItem) GetRemoteItem() AnyOfmicrosoftGraphRemoteItem`

GetRemoteItem returns the RemoteItem field if non-nil, zero value otherwise.

### GetRemoteItemOk

`func (o *MicrosoftGraphDriveItem) GetRemoteItemOk() (AnyOfmicrosoftGraphRemoteItem, bool)`

GetRemoteItemOk returns a tuple with the RemoteItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemoteItem

`func (o *MicrosoftGraphDriveItem) HasRemoteItem() bool`

HasRemoteItem returns a boolean if a field has been set.

### SetRemoteItem

`func (o *MicrosoftGraphDriveItem) SetRemoteItem(v AnyOfmicrosoftGraphRemoteItem)`

SetRemoteItem gets a reference to the given AnyOfmicrosoftGraphRemoteItem and assigns it to the RemoteItem field.

### SetRemoteItemExplicitNull

`func (o *MicrosoftGraphDriveItem) SetRemoteItemExplicitNull(b bool)`

SetRemoteItemExplicitNull (un)sets RemoteItem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RemoteItem value is set to nil even if false is passed
### GetRoot

`func (o *MicrosoftGraphDriveItem) GetRoot() AnyOfobject`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *MicrosoftGraphDriveItem) GetRootOk() (AnyOfobject, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoot

`func (o *MicrosoftGraphDriveItem) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRoot

`func (o *MicrosoftGraphDriveItem) SetRoot(v AnyOfobject)`

SetRoot gets a reference to the given AnyOfobject and assigns it to the Root field.

### SetRootExplicitNull

`func (o *MicrosoftGraphDriveItem) SetRootExplicitNull(b bool)`

SetRootExplicitNull (un)sets Root to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Root value is set to nil even if false is passed
### GetSearchResult

`func (o *MicrosoftGraphDriveItem) GetSearchResult() AnyOfmicrosoftGraphSearchResult`

GetSearchResult returns the SearchResult field if non-nil, zero value otherwise.

### GetSearchResultOk

`func (o *MicrosoftGraphDriveItem) GetSearchResultOk() (AnyOfmicrosoftGraphSearchResult, bool)`

GetSearchResultOk returns a tuple with the SearchResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchResult

`func (o *MicrosoftGraphDriveItem) HasSearchResult() bool`

HasSearchResult returns a boolean if a field has been set.

### SetSearchResult

`func (o *MicrosoftGraphDriveItem) SetSearchResult(v AnyOfmicrosoftGraphSearchResult)`

SetSearchResult gets a reference to the given AnyOfmicrosoftGraphSearchResult and assigns it to the SearchResult field.

### SetSearchResultExplicitNull

`func (o *MicrosoftGraphDriveItem) SetSearchResultExplicitNull(b bool)`

SetSearchResultExplicitNull (un)sets SearchResult to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SearchResult value is set to nil even if false is passed
### GetShared

`func (o *MicrosoftGraphDriveItem) GetShared() AnyOfmicrosoftGraphShared`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *MicrosoftGraphDriveItem) GetSharedOk() (AnyOfmicrosoftGraphShared, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShared

`func (o *MicrosoftGraphDriveItem) HasShared() bool`

HasShared returns a boolean if a field has been set.

### SetShared

`func (o *MicrosoftGraphDriveItem) SetShared(v AnyOfmicrosoftGraphShared)`

SetShared gets a reference to the given AnyOfmicrosoftGraphShared and assigns it to the Shared field.

### SetSharedExplicitNull

`func (o *MicrosoftGraphDriveItem) SetSharedExplicitNull(b bool)`

SetSharedExplicitNull (un)sets Shared to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Shared value is set to nil even if false is passed
### GetSharepointIds

`func (o *MicrosoftGraphDriveItem) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *MicrosoftGraphDriveItem) GetSharepointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharepointIds

`func (o *MicrosoftGraphDriveItem) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIds

`func (o *MicrosoftGraphDriveItem) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.

### SetSharepointIdsExplicitNull

`func (o *MicrosoftGraphDriveItem) SetSharepointIdsExplicitNull(b bool)`

SetSharepointIdsExplicitNull (un)sets SharepointIds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharepointIds value is set to nil even if false is passed
### GetSize

`func (o *MicrosoftGraphDriveItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphDriveItem) GetSizeOk() (int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *MicrosoftGraphDriveItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *MicrosoftGraphDriveItem) SetSize(v int64)`

SetSize gets a reference to the given int64 and assigns it to the Size field.

### SetSizeExplicitNull

`func (o *MicrosoftGraphDriveItem) SetSizeExplicitNull(b bool)`

SetSizeExplicitNull (un)sets Size to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Size value is set to nil even if false is passed
### GetSpecialFolder

`func (o *MicrosoftGraphDriveItem) GetSpecialFolder() AnyOfmicrosoftGraphSpecialFolder`

GetSpecialFolder returns the SpecialFolder field if non-nil, zero value otherwise.

### GetSpecialFolderOk

`func (o *MicrosoftGraphDriveItem) GetSpecialFolderOk() (AnyOfmicrosoftGraphSpecialFolder, bool)`

GetSpecialFolderOk returns a tuple with the SpecialFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecialFolder

`func (o *MicrosoftGraphDriveItem) HasSpecialFolder() bool`

HasSpecialFolder returns a boolean if a field has been set.

### SetSpecialFolder

`func (o *MicrosoftGraphDriveItem) SetSpecialFolder(v AnyOfmicrosoftGraphSpecialFolder)`

SetSpecialFolder gets a reference to the given AnyOfmicrosoftGraphSpecialFolder and assigns it to the SpecialFolder field.

### SetSpecialFolderExplicitNull

`func (o *MicrosoftGraphDriveItem) SetSpecialFolderExplicitNull(b bool)`

SetSpecialFolderExplicitNull (un)sets SpecialFolder to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SpecialFolder value is set to nil even if false is passed
### GetVideo

`func (o *MicrosoftGraphDriveItem) GetVideo() AnyOfmicrosoftGraphVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *MicrosoftGraphDriveItem) GetVideoOk() (AnyOfmicrosoftGraphVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVideo

`func (o *MicrosoftGraphDriveItem) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### SetVideo

`func (o *MicrosoftGraphDriveItem) SetVideo(v AnyOfmicrosoftGraphVideo)`

SetVideo gets a reference to the given AnyOfmicrosoftGraphVideo and assigns it to the Video field.

### SetVideoExplicitNull

`func (o *MicrosoftGraphDriveItem) SetVideoExplicitNull(b bool)`

SetVideoExplicitNull (un)sets Video to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Video value is set to nil even if false is passed
### GetWebDavUrl

`func (o *MicrosoftGraphDriveItem) GetWebDavUrl() string`

GetWebDavUrl returns the WebDavUrl field if non-nil, zero value otherwise.

### GetWebDavUrlOk

`func (o *MicrosoftGraphDriveItem) GetWebDavUrlOk() (string, bool)`

GetWebDavUrlOk returns a tuple with the WebDavUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebDavUrl

`func (o *MicrosoftGraphDriveItem) HasWebDavUrl() bool`

HasWebDavUrl returns a boolean if a field has been set.

### SetWebDavUrl

`func (o *MicrosoftGraphDriveItem) SetWebDavUrl(v string)`

SetWebDavUrl gets a reference to the given string and assigns it to the WebDavUrl field.

### SetWebDavUrlExplicitNull

`func (o *MicrosoftGraphDriveItem) SetWebDavUrlExplicitNull(b bool)`

SetWebDavUrlExplicitNull (un)sets WebDavUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebDavUrl value is set to nil even if false is passed
### GetWorkbook

`func (o *MicrosoftGraphDriveItem) GetWorkbook() AnyOfmicrosoftGraphWorkbook`

GetWorkbook returns the Workbook field if non-nil, zero value otherwise.

### GetWorkbookOk

`func (o *MicrosoftGraphDriveItem) GetWorkbookOk() (AnyOfmicrosoftGraphWorkbook, bool)`

GetWorkbookOk returns a tuple with the Workbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkbook

`func (o *MicrosoftGraphDriveItem) HasWorkbook() bool`

HasWorkbook returns a boolean if a field has been set.

### SetWorkbook

`func (o *MicrosoftGraphDriveItem) SetWorkbook(v AnyOfmicrosoftGraphWorkbook)`

SetWorkbook gets a reference to the given AnyOfmicrosoftGraphWorkbook and assigns it to the Workbook field.

### SetWorkbookExplicitNull

`func (o *MicrosoftGraphDriveItem) SetWorkbookExplicitNull(b bool)`

SetWorkbookExplicitNull (un)sets Workbook to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Workbook value is set to nil even if false is passed
### GetAnalytics

`func (o *MicrosoftGraphDriveItem) GetAnalytics() AnyOfmicrosoftGraphItemAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *MicrosoftGraphDriveItem) GetAnalyticsOk() (AnyOfmicrosoftGraphItemAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAnalytics

`func (o *MicrosoftGraphDriveItem) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### SetAnalytics

`func (o *MicrosoftGraphDriveItem) SetAnalytics(v AnyOfmicrosoftGraphItemAnalytics)`

SetAnalytics gets a reference to the given AnyOfmicrosoftGraphItemAnalytics and assigns it to the Analytics field.

### SetAnalyticsExplicitNull

`func (o *MicrosoftGraphDriveItem) SetAnalyticsExplicitNull(b bool)`

SetAnalyticsExplicitNull (un)sets Analytics to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Analytics value is set to nil even if false is passed
### GetChildren

`func (o *MicrosoftGraphDriveItem) GetChildren() []MicrosoftGraphDriveItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MicrosoftGraphDriveItem) GetChildrenOk() ([]MicrosoftGraphDriveItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildren

`func (o *MicrosoftGraphDriveItem) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildren

`func (o *MicrosoftGraphDriveItem) SetChildren(v []MicrosoftGraphDriveItem)`

SetChildren gets a reference to the given []MicrosoftGraphDriveItem and assigns it to the Children field.

### GetListItem

`func (o *MicrosoftGraphDriveItem) GetListItem() AnyOfmicrosoftGraphListItem`

GetListItem returns the ListItem field if non-nil, zero value otherwise.

### GetListItemOk

`func (o *MicrosoftGraphDriveItem) GetListItemOk() (AnyOfmicrosoftGraphListItem, bool)`

GetListItemOk returns a tuple with the ListItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListItem

`func (o *MicrosoftGraphDriveItem) HasListItem() bool`

HasListItem returns a boolean if a field has been set.

### SetListItem

`func (o *MicrosoftGraphDriveItem) SetListItem(v AnyOfmicrosoftGraphListItem)`

SetListItem gets a reference to the given AnyOfmicrosoftGraphListItem and assigns it to the ListItem field.

### SetListItemExplicitNull

`func (o *MicrosoftGraphDriveItem) SetListItemExplicitNull(b bool)`

SetListItemExplicitNull (un)sets ListItem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ListItem value is set to nil even if false is passed
### GetPermissions

`func (o *MicrosoftGraphDriveItem) GetPermissions() []MicrosoftGraphPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MicrosoftGraphDriveItem) GetPermissionsOk() ([]MicrosoftGraphPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissions

`func (o *MicrosoftGraphDriveItem) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissions

`func (o *MicrosoftGraphDriveItem) SetPermissions(v []MicrosoftGraphPermission)`

SetPermissions gets a reference to the given []MicrosoftGraphPermission and assigns it to the Permissions field.

### GetSubscriptions

`func (o *MicrosoftGraphDriveItem) GetSubscriptions() []MicrosoftGraphSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *MicrosoftGraphDriveItem) GetSubscriptionsOk() ([]MicrosoftGraphSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscriptions

`func (o *MicrosoftGraphDriveItem) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### SetSubscriptions

`func (o *MicrosoftGraphDriveItem) SetSubscriptions(v []MicrosoftGraphSubscription)`

SetSubscriptions gets a reference to the given []MicrosoftGraphSubscription and assigns it to the Subscriptions field.

### GetThumbnails

`func (o *MicrosoftGraphDriveItem) GetThumbnails() []MicrosoftGraphThumbnailSet`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *MicrosoftGraphDriveItem) GetThumbnailsOk() ([]MicrosoftGraphThumbnailSet, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThumbnails

`func (o *MicrosoftGraphDriveItem) HasThumbnails() bool`

HasThumbnails returns a boolean if a field has been set.

### SetThumbnails

`func (o *MicrosoftGraphDriveItem) SetThumbnails(v []MicrosoftGraphThumbnailSet)`

SetThumbnails gets a reference to the given []MicrosoftGraphThumbnailSet and assigns it to the Thumbnails field.

### GetVersions

`func (o *MicrosoftGraphDriveItem) GetVersions() []MicrosoftGraphDriveItemVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *MicrosoftGraphDriveItem) GetVersionsOk() ([]MicrosoftGraphDriveItemVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersions

`func (o *MicrosoftGraphDriveItem) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### SetVersions

`func (o *MicrosoftGraphDriveItem) SetVersions(v []MicrosoftGraphDriveItemVersion)`

SetVersions gets a reference to the given []MicrosoftGraphDriveItemVersion and assigns it to the Versions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


