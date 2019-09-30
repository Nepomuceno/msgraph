# DriveItemVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 

## Methods

### GetContent

`func (o *DriveItemVersion) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DriveItemVersion) GetContentOk() (string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContent

`func (o *DriveItemVersion) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContent

`func (o *DriveItemVersion) SetContent(v string)`

SetContent gets a reference to the given string and assigns it to the Content field.

### SetContentExplicitNull

`func (o *DriveItemVersion) SetContentExplicitNull(b bool)`

SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Content value is set to nil even if false is passed
### GetSize

`func (o *DriveItemVersion) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DriveItemVersion) GetSizeOk() (int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *DriveItemVersion) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *DriveItemVersion) SetSize(v int64)`

SetSize gets a reference to the given int64 and assigns it to the Size field.

### SetSizeExplicitNull

`func (o *DriveItemVersion) SetSizeExplicitNull(b bool)`

SetSizeExplicitNull (un)sets Size to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Size value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

