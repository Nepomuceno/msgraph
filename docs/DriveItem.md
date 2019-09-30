# DriveItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetAudio

`func (o *DriveItem) GetAudio() AnyOfmicrosoftGraphAudio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *DriveItem) GetAudioOk() (AnyOfmicrosoftGraphAudio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudio

`func (o *DriveItem) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### SetAudio

`func (o *DriveItem) SetAudio(v AnyOfmicrosoftGraphAudio)`

SetAudio gets a reference to the given AnyOfmicrosoftGraphAudio and assigns it to the Audio field.

### SetAudioExplicitNull

`func (o *DriveItem) SetAudioExplicitNull(b bool)`

SetAudioExplicitNull (un)sets Audio to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Audio value is set to nil even if false is passed
### GetContent

`func (o *DriveItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DriveItem) GetContentOk() (string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContent

`func (o *DriveItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContent

`func (o *DriveItem) SetContent(v string)`

SetContent gets a reference to the given string and assigns it to the Content field.

### SetContentExplicitNull

`func (o *DriveItem) SetContentExplicitNull(b bool)`

SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Content value is set to nil even if false is passed
### GetCTag

`func (o *DriveItem) GetCTag() string`

GetCTag returns the CTag field if non-nil, zero value otherwise.

### GetCTagOk

`func (o *DriveItem) GetCTagOk() (string, bool)`

GetCTagOk returns a tuple with the CTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCTag

`func (o *DriveItem) HasCTag() bool`

HasCTag returns a boolean if a field has been set.

### SetCTag

`func (o *DriveItem) SetCTag(v string)`

SetCTag gets a reference to the given string and assigns it to the CTag field.

### SetCTagExplicitNull

`func (o *DriveItem) SetCTagExplicitNull(b bool)`

SetCTagExplicitNull (un)sets CTag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CTag value is set to nil even if false is passed
### GetDeleted

`func (o *DriveItem) GetDeleted() AnyOfmicrosoftGraphDeleted`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DriveItem) GetDeletedOk() (AnyOfmicrosoftGraphDeleted, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleted

`func (o *DriveItem) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeleted

`func (o *DriveItem) SetDeleted(v AnyOfmicrosoftGraphDeleted)`

SetDeleted gets a reference to the given AnyOfmicrosoftGraphDeleted and assigns it to the Deleted field.

### SetDeletedExplicitNull

`func (o *DriveItem) SetDeletedExplicitNull(b bool)`

SetDeletedExplicitNull (un)sets Deleted to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Deleted value is set to nil even if false is passed
### GetFile

`func (o *DriveItem) GetFile() AnyOfmicrosoftGraphFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *DriveItem) GetFileOk() (AnyOfmicrosoftGraphFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFile

`func (o *DriveItem) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFile

`func (o *DriveItem) SetFile(v AnyOfmicrosoftGraphFile)`

SetFile gets a reference to the given AnyOfmicrosoftGraphFile and assigns it to the File field.

### SetFileExplicitNull

`func (o *DriveItem) SetFileExplicitNull(b bool)`

SetFileExplicitNull (un)sets File to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The File value is set to nil even if false is passed
### GetFileSystemInfo

`func (o *DriveItem) GetFileSystemInfo() AnyOfmicrosoftGraphFileSystemInfo`

GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.

### GetFileSystemInfoOk

`func (o *DriveItem) GetFileSystemInfoOk() (AnyOfmicrosoftGraphFileSystemInfo, bool)`

GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileSystemInfo

`func (o *DriveItem) HasFileSystemInfo() bool`

HasFileSystemInfo returns a boolean if a field has been set.

### SetFileSystemInfo

`func (o *DriveItem) SetFileSystemInfo(v AnyOfmicrosoftGraphFileSystemInfo)`

SetFileSystemInfo gets a reference to the given AnyOfmicrosoftGraphFileSystemInfo and assigns it to the FileSystemInfo field.

### SetFileSystemInfoExplicitNull

`func (o *DriveItem) SetFileSystemInfoExplicitNull(b bool)`

SetFileSystemInfoExplicitNull (un)sets FileSystemInfo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileSystemInfo value is set to nil even if false is passed
### GetFolder

`func (o *DriveItem) GetFolder() AnyOfmicrosoftGraphFolder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *DriveItem) GetFolderOk() (AnyOfmicrosoftGraphFolder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFolder

`func (o *DriveItem) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolder

`func (o *DriveItem) SetFolder(v AnyOfmicrosoftGraphFolder)`

SetFolder gets a reference to the given AnyOfmicrosoftGraphFolder and assigns it to the Folder field.

### SetFolderExplicitNull

`func (o *DriveItem) SetFolderExplicitNull(b bool)`

SetFolderExplicitNull (un)sets Folder to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Folder value is set to nil even if false is passed
### GetImage

`func (o *DriveItem) GetImage() AnyOfmicrosoftGraphImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DriveItem) GetImageOk() (AnyOfmicrosoftGraphImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImage

`func (o *DriveItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImage

`func (o *DriveItem) SetImage(v AnyOfmicrosoftGraphImage)`

SetImage gets a reference to the given AnyOfmicrosoftGraphImage and assigns it to the Image field.

### SetImageExplicitNull

`func (o *DriveItem) SetImageExplicitNull(b bool)`

SetImageExplicitNull (un)sets Image to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Image value is set to nil even if false is passed
### GetLocation

`func (o *DriveItem) GetLocation() AnyOfmicrosoftGraphGeoCoordinates`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DriveItem) GetLocationOk() (AnyOfmicrosoftGraphGeoCoordinates, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocation

`func (o *DriveItem) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocation

`func (o *DriveItem) SetLocation(v AnyOfmicrosoftGraphGeoCoordinates)`

SetLocation gets a reference to the given AnyOfmicrosoftGraphGeoCoordinates and assigns it to the Location field.

### SetLocationExplicitNull

`func (o *DriveItem) SetLocationExplicitNull(b bool)`

SetLocationExplicitNull (un)sets Location to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Location value is set to nil even if false is passed
### GetPackage

`func (o *DriveItem) GetPackage() AnyOfmicrosoftGraphPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *DriveItem) GetPackageOk() (AnyOfmicrosoftGraphPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPackage

`func (o *DriveItem) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### SetPackage

`func (o *DriveItem) SetPackage(v AnyOfmicrosoftGraphPackage)`

SetPackage gets a reference to the given AnyOfmicrosoftGraphPackage and assigns it to the Package field.

### SetPackageExplicitNull

`func (o *DriveItem) SetPackageExplicitNull(b bool)`

SetPackageExplicitNull (un)sets Package to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Package value is set to nil even if false is passed
### GetPhoto

`func (o *DriveItem) GetPhoto() AnyOfmicrosoftGraphPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *DriveItem) GetPhotoOk() (AnyOfmicrosoftGraphPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhoto

`func (o *DriveItem) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhoto

`func (o *DriveItem) SetPhoto(v AnyOfmicrosoftGraphPhoto)`

SetPhoto gets a reference to the given AnyOfmicrosoftGraphPhoto and assigns it to the Photo field.

### SetPhotoExplicitNull

`func (o *DriveItem) SetPhotoExplicitNull(b bool)`

SetPhotoExplicitNull (un)sets Photo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Photo value is set to nil even if false is passed
### GetPublication

`func (o *DriveItem) GetPublication() AnyOfmicrosoftGraphPublicationFacet`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *DriveItem) GetPublicationOk() (AnyOfmicrosoftGraphPublicationFacet, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublication

`func (o *DriveItem) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublication

`func (o *DriveItem) SetPublication(v AnyOfmicrosoftGraphPublicationFacet)`

SetPublication gets a reference to the given AnyOfmicrosoftGraphPublicationFacet and assigns it to the Publication field.

### SetPublicationExplicitNull

`func (o *DriveItem) SetPublicationExplicitNull(b bool)`

SetPublicationExplicitNull (un)sets Publication to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publication value is set to nil even if false is passed
### GetRemoteItem

`func (o *DriveItem) GetRemoteItem() AnyOfmicrosoftGraphRemoteItem`

GetRemoteItem returns the RemoteItem field if non-nil, zero value otherwise.

### GetRemoteItemOk

`func (o *DriveItem) GetRemoteItemOk() (AnyOfmicrosoftGraphRemoteItem, bool)`

GetRemoteItemOk returns a tuple with the RemoteItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemoteItem

`func (o *DriveItem) HasRemoteItem() bool`

HasRemoteItem returns a boolean if a field has been set.

### SetRemoteItem

`func (o *DriveItem) SetRemoteItem(v AnyOfmicrosoftGraphRemoteItem)`

SetRemoteItem gets a reference to the given AnyOfmicrosoftGraphRemoteItem and assigns it to the RemoteItem field.

### SetRemoteItemExplicitNull

`func (o *DriveItem) SetRemoteItemExplicitNull(b bool)`

SetRemoteItemExplicitNull (un)sets RemoteItem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RemoteItem value is set to nil even if false is passed
### GetRoot

`func (o *DriveItem) GetRoot() AnyOfobject`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *DriveItem) GetRootOk() (AnyOfobject, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoot

`func (o *DriveItem) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRoot

`func (o *DriveItem) SetRoot(v AnyOfobject)`

SetRoot gets a reference to the given AnyOfobject and assigns it to the Root field.

### SetRootExplicitNull

`func (o *DriveItem) SetRootExplicitNull(b bool)`

SetRootExplicitNull (un)sets Root to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Root value is set to nil even if false is passed
### GetSearchResult

`func (o *DriveItem) GetSearchResult() AnyOfmicrosoftGraphSearchResult`

GetSearchResult returns the SearchResult field if non-nil, zero value otherwise.

### GetSearchResultOk

`func (o *DriveItem) GetSearchResultOk() (AnyOfmicrosoftGraphSearchResult, bool)`

GetSearchResultOk returns a tuple with the SearchResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSearchResult

`func (o *DriveItem) HasSearchResult() bool`

HasSearchResult returns a boolean if a field has been set.

### SetSearchResult

`func (o *DriveItem) SetSearchResult(v AnyOfmicrosoftGraphSearchResult)`

SetSearchResult gets a reference to the given AnyOfmicrosoftGraphSearchResult and assigns it to the SearchResult field.

### SetSearchResultExplicitNull

`func (o *DriveItem) SetSearchResultExplicitNull(b bool)`

SetSearchResultExplicitNull (un)sets SearchResult to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SearchResult value is set to nil even if false is passed
### GetShared

`func (o *DriveItem) GetShared() AnyOfmicrosoftGraphShared`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *DriveItem) GetSharedOk() (AnyOfmicrosoftGraphShared, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShared

`func (o *DriveItem) HasShared() bool`

HasShared returns a boolean if a field has been set.

### SetShared

`func (o *DriveItem) SetShared(v AnyOfmicrosoftGraphShared)`

SetShared gets a reference to the given AnyOfmicrosoftGraphShared and assigns it to the Shared field.

### SetSharedExplicitNull

`func (o *DriveItem) SetSharedExplicitNull(b bool)`

SetSharedExplicitNull (un)sets Shared to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Shared value is set to nil even if false is passed
### GetSharepointIds

`func (o *DriveItem) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *DriveItem) GetSharepointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharepointIds

`func (o *DriveItem) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIds

`func (o *DriveItem) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.

### SetSharepointIdsExplicitNull

`func (o *DriveItem) SetSharepointIdsExplicitNull(b bool)`

SetSharepointIdsExplicitNull (un)sets SharepointIds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharepointIds value is set to nil even if false is passed
### GetSize

`func (o *DriveItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DriveItem) GetSizeOk() (int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *DriveItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *DriveItem) SetSize(v int64)`

SetSize gets a reference to the given int64 and assigns it to the Size field.

### SetSizeExplicitNull

`func (o *DriveItem) SetSizeExplicitNull(b bool)`

SetSizeExplicitNull (un)sets Size to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Size value is set to nil even if false is passed
### GetSpecialFolder

`func (o *DriveItem) GetSpecialFolder() AnyOfmicrosoftGraphSpecialFolder`

GetSpecialFolder returns the SpecialFolder field if non-nil, zero value otherwise.

### GetSpecialFolderOk

`func (o *DriveItem) GetSpecialFolderOk() (AnyOfmicrosoftGraphSpecialFolder, bool)`

GetSpecialFolderOk returns a tuple with the SpecialFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecialFolder

`func (o *DriveItem) HasSpecialFolder() bool`

HasSpecialFolder returns a boolean if a field has been set.

### SetSpecialFolder

`func (o *DriveItem) SetSpecialFolder(v AnyOfmicrosoftGraphSpecialFolder)`

SetSpecialFolder gets a reference to the given AnyOfmicrosoftGraphSpecialFolder and assigns it to the SpecialFolder field.

### SetSpecialFolderExplicitNull

`func (o *DriveItem) SetSpecialFolderExplicitNull(b bool)`

SetSpecialFolderExplicitNull (un)sets SpecialFolder to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SpecialFolder value is set to nil even if false is passed
### GetVideo

`func (o *DriveItem) GetVideo() AnyOfmicrosoftGraphVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *DriveItem) GetVideoOk() (AnyOfmicrosoftGraphVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVideo

`func (o *DriveItem) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### SetVideo

`func (o *DriveItem) SetVideo(v AnyOfmicrosoftGraphVideo)`

SetVideo gets a reference to the given AnyOfmicrosoftGraphVideo and assigns it to the Video field.

### SetVideoExplicitNull

`func (o *DriveItem) SetVideoExplicitNull(b bool)`

SetVideoExplicitNull (un)sets Video to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Video value is set to nil even if false is passed
### GetWebDavUrl

`func (o *DriveItem) GetWebDavUrl() string`

GetWebDavUrl returns the WebDavUrl field if non-nil, zero value otherwise.

### GetWebDavUrlOk

`func (o *DriveItem) GetWebDavUrlOk() (string, bool)`

GetWebDavUrlOk returns a tuple with the WebDavUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebDavUrl

`func (o *DriveItem) HasWebDavUrl() bool`

HasWebDavUrl returns a boolean if a field has been set.

### SetWebDavUrl

`func (o *DriveItem) SetWebDavUrl(v string)`

SetWebDavUrl gets a reference to the given string and assigns it to the WebDavUrl field.

### SetWebDavUrlExplicitNull

`func (o *DriveItem) SetWebDavUrlExplicitNull(b bool)`

SetWebDavUrlExplicitNull (un)sets WebDavUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebDavUrl value is set to nil even if false is passed
### GetWorkbook

`func (o *DriveItem) GetWorkbook() AnyOfmicrosoftGraphWorkbook`

GetWorkbook returns the Workbook field if non-nil, zero value otherwise.

### GetWorkbookOk

`func (o *DriveItem) GetWorkbookOk() (AnyOfmicrosoftGraphWorkbook, bool)`

GetWorkbookOk returns a tuple with the Workbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkbook

`func (o *DriveItem) HasWorkbook() bool`

HasWorkbook returns a boolean if a field has been set.

### SetWorkbook

`func (o *DriveItem) SetWorkbook(v AnyOfmicrosoftGraphWorkbook)`

SetWorkbook gets a reference to the given AnyOfmicrosoftGraphWorkbook and assigns it to the Workbook field.

### SetWorkbookExplicitNull

`func (o *DriveItem) SetWorkbookExplicitNull(b bool)`

SetWorkbookExplicitNull (un)sets Workbook to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Workbook value is set to nil even if false is passed
### GetAnalytics

`func (o *DriveItem) GetAnalytics() AnyOfmicrosoftGraphItemAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *DriveItem) GetAnalyticsOk() (AnyOfmicrosoftGraphItemAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAnalytics

`func (o *DriveItem) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### SetAnalytics

`func (o *DriveItem) SetAnalytics(v AnyOfmicrosoftGraphItemAnalytics)`

SetAnalytics gets a reference to the given AnyOfmicrosoftGraphItemAnalytics and assigns it to the Analytics field.

### SetAnalyticsExplicitNull

`func (o *DriveItem) SetAnalyticsExplicitNull(b bool)`

SetAnalyticsExplicitNull (un)sets Analytics to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Analytics value is set to nil even if false is passed
### GetChildren

`func (o *DriveItem) GetChildren() []MicrosoftGraphDriveItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *DriveItem) GetChildrenOk() ([]MicrosoftGraphDriveItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildren

`func (o *DriveItem) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### SetChildren

`func (o *DriveItem) SetChildren(v []MicrosoftGraphDriveItem)`

SetChildren gets a reference to the given []MicrosoftGraphDriveItem and assigns it to the Children field.

### GetListItem

`func (o *DriveItem) GetListItem() AnyOfmicrosoftGraphListItem`

GetListItem returns the ListItem field if non-nil, zero value otherwise.

### GetListItemOk

`func (o *DriveItem) GetListItemOk() (AnyOfmicrosoftGraphListItem, bool)`

GetListItemOk returns a tuple with the ListItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListItem

`func (o *DriveItem) HasListItem() bool`

HasListItem returns a boolean if a field has been set.

### SetListItem

`func (o *DriveItem) SetListItem(v AnyOfmicrosoftGraphListItem)`

SetListItem gets a reference to the given AnyOfmicrosoftGraphListItem and assigns it to the ListItem field.

### SetListItemExplicitNull

`func (o *DriveItem) SetListItemExplicitNull(b bool)`

SetListItemExplicitNull (un)sets ListItem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ListItem value is set to nil even if false is passed
### GetPermissions

`func (o *DriveItem) GetPermissions() []MicrosoftGraphPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DriveItem) GetPermissionsOk() ([]MicrosoftGraphPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissions

`func (o *DriveItem) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissions

`func (o *DriveItem) SetPermissions(v []MicrosoftGraphPermission)`

SetPermissions gets a reference to the given []MicrosoftGraphPermission and assigns it to the Permissions field.

### GetSubscriptions

`func (o *DriveItem) GetSubscriptions() []MicrosoftGraphSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *DriveItem) GetSubscriptionsOk() ([]MicrosoftGraphSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscriptions

`func (o *DriveItem) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### SetSubscriptions

`func (o *DriveItem) SetSubscriptions(v []MicrosoftGraphSubscription)`

SetSubscriptions gets a reference to the given []MicrosoftGraphSubscription and assigns it to the Subscriptions field.

### GetThumbnails

`func (o *DriveItem) GetThumbnails() []MicrosoftGraphThumbnailSet`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *DriveItem) GetThumbnailsOk() ([]MicrosoftGraphThumbnailSet, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThumbnails

`func (o *DriveItem) HasThumbnails() bool`

HasThumbnails returns a boolean if a field has been set.

### SetThumbnails

`func (o *DriveItem) SetThumbnails(v []MicrosoftGraphThumbnailSet)`

SetThumbnails gets a reference to the given []MicrosoftGraphThumbnailSet and assigns it to the Thumbnails field.

### GetVersions

`func (o *DriveItem) GetVersions() []MicrosoftGraphDriveItemVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *DriveItem) GetVersionsOk() ([]MicrosoftGraphDriveItemVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersions

`func (o *DriveItem) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### SetVersions

`func (o *DriveItem) SetVersions(v []MicrosoftGraphDriveItemVersion)`

SetVersions gets a reference to the given []MicrosoftGraphDriveItemVersion and assigns it to the Versions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


