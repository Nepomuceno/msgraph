# MicrosoftGraphFileAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**IsInline** | Pointer to **bool** |  | [optional] 
**ContentId** | Pointer to **string** |  | [optional] 
**ContentLocation** | Pointer to **string** |  | [optional] 
**ContentBytes** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphFileAttachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphFileAttachment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphFileAttachment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphFileAttachment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphFileAttachment) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphFileAttachment) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphFileAttachment) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphFileAttachment) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphFileAttachment) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetName

`func (o *MicrosoftGraphFileAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphFileAttachment) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphFileAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphFileAttachment) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphFileAttachment) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetContentType

`func (o *MicrosoftGraphFileAttachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MicrosoftGraphFileAttachment) GetContentTypeOk() (string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentType

`func (o *MicrosoftGraphFileAttachment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentType

`func (o *MicrosoftGraphFileAttachment) SetContentType(v string)`

SetContentType gets a reference to the given string and assigns it to the ContentType field.

### SetContentTypeExplicitNull

`func (o *MicrosoftGraphFileAttachment) SetContentTypeExplicitNull(b bool)`

SetContentTypeExplicitNull (un)sets ContentType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentType value is set to nil even if false is passed
### GetSize

`func (o *MicrosoftGraphFileAttachment) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphFileAttachment) GetSizeOk() (int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *MicrosoftGraphFileAttachment) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *MicrosoftGraphFileAttachment) SetSize(v int32)`

SetSize gets a reference to the given int32 and assigns it to the Size field.

### GetIsInline

`func (o *MicrosoftGraphFileAttachment) GetIsInline() bool`

GetIsInline returns the IsInline field if non-nil, zero value otherwise.

### GetIsInlineOk

`func (o *MicrosoftGraphFileAttachment) GetIsInlineOk() (bool, bool)`

GetIsInlineOk returns a tuple with the IsInline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsInline

`func (o *MicrosoftGraphFileAttachment) HasIsInline() bool`

HasIsInline returns a boolean if a field has been set.

### SetIsInline

`func (o *MicrosoftGraphFileAttachment) SetIsInline(v bool)`

SetIsInline gets a reference to the given bool and assigns it to the IsInline field.

### GetContentId

`func (o *MicrosoftGraphFileAttachment) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *MicrosoftGraphFileAttachment) GetContentIdOk() (string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentId

`func (o *MicrosoftGraphFileAttachment) HasContentId() bool`

HasContentId returns a boolean if a field has been set.

### SetContentId

`func (o *MicrosoftGraphFileAttachment) SetContentId(v string)`

SetContentId gets a reference to the given string and assigns it to the ContentId field.

### SetContentIdExplicitNull

`func (o *MicrosoftGraphFileAttachment) SetContentIdExplicitNull(b bool)`

SetContentIdExplicitNull (un)sets ContentId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentId value is set to nil even if false is passed
### GetContentLocation

`func (o *MicrosoftGraphFileAttachment) GetContentLocation() string`

GetContentLocation returns the ContentLocation field if non-nil, zero value otherwise.

### GetContentLocationOk

`func (o *MicrosoftGraphFileAttachment) GetContentLocationOk() (string, bool)`

GetContentLocationOk returns a tuple with the ContentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentLocation

`func (o *MicrosoftGraphFileAttachment) HasContentLocation() bool`

HasContentLocation returns a boolean if a field has been set.

### SetContentLocation

`func (o *MicrosoftGraphFileAttachment) SetContentLocation(v string)`

SetContentLocation gets a reference to the given string and assigns it to the ContentLocation field.

### SetContentLocationExplicitNull

`func (o *MicrosoftGraphFileAttachment) SetContentLocationExplicitNull(b bool)`

SetContentLocationExplicitNull (un)sets ContentLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentLocation value is set to nil even if false is passed
### GetContentBytes

`func (o *MicrosoftGraphFileAttachment) GetContentBytes() string`

GetContentBytes returns the ContentBytes field if non-nil, zero value otherwise.

### GetContentBytesOk

`func (o *MicrosoftGraphFileAttachment) GetContentBytesOk() (string, bool)`

GetContentBytesOk returns a tuple with the ContentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentBytes

`func (o *MicrosoftGraphFileAttachment) HasContentBytes() bool`

HasContentBytes returns a boolean if a field has been set.

### SetContentBytes

`func (o *MicrosoftGraphFileAttachment) SetContentBytes(v string)`

SetContentBytes gets a reference to the given string and assigns it to the ContentBytes field.

### SetContentBytesExplicitNull

`func (o *MicrosoftGraphFileAttachment) SetContentBytesExplicitNull(b bool)`

SetContentBytesExplicitNull (un)sets ContentBytes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentBytes value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


