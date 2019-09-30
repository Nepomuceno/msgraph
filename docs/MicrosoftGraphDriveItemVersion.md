# MicrosoftGraphDriveItemVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Publication** | Pointer to [**AnyOfmicrosoftGraphPublicationFacet**](anyOf&lt;microsoft.graph.publicationFacet&gt;.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDriveItemVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDriveItemVersion) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDriveItemVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDriveItemVersion) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedBy

`func (o *MicrosoftGraphDriveItemVersion) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphDriveItemVersion) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphDriveItemVersion) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphDriveItemVersion) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphDriveItemVersion) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphDriveItemVersion) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphDriveItemVersion) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphDriveItemVersion) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphDriveItemVersion) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphDriveItemVersion) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetPublication

`func (o *MicrosoftGraphDriveItemVersion) GetPublication() AnyOfmicrosoftGraphPublicationFacet`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *MicrosoftGraphDriveItemVersion) GetPublicationOk() (AnyOfmicrosoftGraphPublicationFacet, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublication

`func (o *MicrosoftGraphDriveItemVersion) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublication

`func (o *MicrosoftGraphDriveItemVersion) SetPublication(v AnyOfmicrosoftGraphPublicationFacet)`

SetPublication gets a reference to the given AnyOfmicrosoftGraphPublicationFacet and assigns it to the Publication field.

### SetPublicationExplicitNull

`func (o *MicrosoftGraphDriveItemVersion) SetPublicationExplicitNull(b bool)`

SetPublicationExplicitNull (un)sets Publication to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publication value is set to nil even if false is passed
### GetContent

`func (o *MicrosoftGraphDriveItemVersion) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphDriveItemVersion) GetContentOk() (string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContent

`func (o *MicrosoftGraphDriveItemVersion) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContent

`func (o *MicrosoftGraphDriveItemVersion) SetContent(v string)`

SetContent gets a reference to the given string and assigns it to the Content field.

### SetContentExplicitNull

`func (o *MicrosoftGraphDriveItemVersion) SetContentExplicitNull(b bool)`

SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Content value is set to nil even if false is passed
### GetSize

`func (o *MicrosoftGraphDriveItemVersion) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphDriveItemVersion) GetSizeOk() (int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *MicrosoftGraphDriveItemVersion) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *MicrosoftGraphDriveItemVersion) SetSize(v int64)`

SetSize gets a reference to the given int64 and assigns it to the Size field.

### SetSizeExplicitNull

`func (o *MicrosoftGraphDriveItemVersion) SetSizeExplicitNull(b bool)`

SetSizeExplicitNull (un)sets Size to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Size value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


