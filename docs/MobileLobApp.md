# MobileLobApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommittedContentVersion** | Pointer to **string** | The internal committed content version. | [optional] 
**FileName** | Pointer to **string** | The name of the main Lob application file. | [optional] 
**Size** | Pointer to **int64** | The total size, including all uploaded files. | [optional] 
**ContentVersions** | Pointer to [**[]MicrosoftGraphMobileAppContent**](microsoft.graph.mobileAppContent.md) |  | [optional] 

## Methods

### GetCommittedContentVersion

`func (o *MobileLobApp) GetCommittedContentVersion() string`

GetCommittedContentVersion returns the CommittedContentVersion field if non-nil, zero value otherwise.

### GetCommittedContentVersionOk

`func (o *MobileLobApp) GetCommittedContentVersionOk() (string, bool)`

GetCommittedContentVersionOk returns a tuple with the CommittedContentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCommittedContentVersion

`func (o *MobileLobApp) HasCommittedContentVersion() bool`

HasCommittedContentVersion returns a boolean if a field has been set.

### SetCommittedContentVersion

`func (o *MobileLobApp) SetCommittedContentVersion(v string)`

SetCommittedContentVersion gets a reference to the given string and assigns it to the CommittedContentVersion field.

### SetCommittedContentVersionExplicitNull

`func (o *MobileLobApp) SetCommittedContentVersionExplicitNull(b bool)`

SetCommittedContentVersionExplicitNull (un)sets CommittedContentVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CommittedContentVersion value is set to nil even if false is passed
### GetFileName

`func (o *MobileLobApp) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MobileLobApp) GetFileNameOk() (string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileName

`func (o *MobileLobApp) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileName

`func (o *MobileLobApp) SetFileName(v string)`

SetFileName gets a reference to the given string and assigns it to the FileName field.

### SetFileNameExplicitNull

`func (o *MobileLobApp) SetFileNameExplicitNull(b bool)`

SetFileNameExplicitNull (un)sets FileName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileName value is set to nil even if false is passed
### GetSize

`func (o *MobileLobApp) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MobileLobApp) GetSizeOk() (int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *MobileLobApp) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *MobileLobApp) SetSize(v int64)`

SetSize gets a reference to the given int64 and assigns it to the Size field.

### GetContentVersions

`func (o *MobileLobApp) GetContentVersions() []MicrosoftGraphMobileAppContent`

GetContentVersions returns the ContentVersions field if non-nil, zero value otherwise.

### GetContentVersionsOk

`func (o *MobileLobApp) GetContentVersionsOk() ([]MicrosoftGraphMobileAppContent, bool)`

GetContentVersionsOk returns a tuple with the ContentVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentVersions

`func (o *MobileLobApp) HasContentVersions() bool`

HasContentVersions returns a boolean if a field has been set.

### SetContentVersions

`func (o *MobileLobApp) SetContentVersions(v []MicrosoftGraphMobileAppContent)`

SetContentVersions gets a reference to the given []MicrosoftGraphMobileAppContent and assigns it to the ContentVersions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


