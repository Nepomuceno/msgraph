# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**IsInline** | Pointer to **bool** |  | [optional] 

## Methods

### GetLastModifiedDateTime

`func (o *Attachment) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *Attachment) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *Attachment) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *Attachment) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *Attachment) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetName

`func (o *Attachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Attachment) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Attachment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Attachment) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *Attachment) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetContentType

`func (o *Attachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Attachment) GetContentTypeOk() (string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentType

`func (o *Attachment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentType

`func (o *Attachment) SetContentType(v string)`

SetContentType gets a reference to the given string and assigns it to the ContentType field.

### SetContentTypeExplicitNull

`func (o *Attachment) SetContentTypeExplicitNull(b bool)`

SetContentTypeExplicitNull (un)sets ContentType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentType value is set to nil even if false is passed
### GetSize

`func (o *Attachment) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Attachment) GetSizeOk() (int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *Attachment) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *Attachment) SetSize(v int32)`

SetSize gets a reference to the given int32 and assigns it to the Size field.

### GetIsInline

`func (o *Attachment) GetIsInline() bool`

GetIsInline returns the IsInline field if non-nil, zero value otherwise.

### GetIsInlineOk

`func (o *Attachment) GetIsInlineOk() (bool, bool)`

GetIsInlineOk returns a tuple with the IsInline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsInline

`func (o *Attachment) HasIsInline() bool`

HasIsInline returns a boolean if a field has been set.

### SetIsInline

`func (o *Attachment) SetIsInline(v bool)`

SetIsInline gets a reference to the given bool and assigns it to the IsInline field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


