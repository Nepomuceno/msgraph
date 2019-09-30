# MicrosoftGraphUploadSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**NextExpectedRanges** | Pointer to **[]string** |  | [optional] 
**UploadUrl** | Pointer to **string** |  | [optional] 

## Methods

### GetExpirationDateTime

`func (o *MicrosoftGraphUploadSession) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphUploadSession) GetExpirationDateTimeOk() (time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpirationDateTime

`func (o *MicrosoftGraphUploadSession) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphUploadSession) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.

### SetExpirationDateTimeExplicitNull

`func (o *MicrosoftGraphUploadSession) SetExpirationDateTimeExplicitNull(b bool)`

SetExpirationDateTimeExplicitNull (un)sets ExpirationDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExpirationDateTime value is set to nil even if false is passed
### GetNextExpectedRanges

`func (o *MicrosoftGraphUploadSession) GetNextExpectedRanges() []string`

GetNextExpectedRanges returns the NextExpectedRanges field if non-nil, zero value otherwise.

### GetNextExpectedRangesOk

`func (o *MicrosoftGraphUploadSession) GetNextExpectedRangesOk() ([]string, bool)`

GetNextExpectedRangesOk returns a tuple with the NextExpectedRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextExpectedRanges

`func (o *MicrosoftGraphUploadSession) HasNextExpectedRanges() bool`

HasNextExpectedRanges returns a boolean if a field has been set.

### SetNextExpectedRanges

`func (o *MicrosoftGraphUploadSession) SetNextExpectedRanges(v []string)`

SetNextExpectedRanges gets a reference to the given []string and assigns it to the NextExpectedRanges field.

### GetUploadUrl

`func (o *MicrosoftGraphUploadSession) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *MicrosoftGraphUploadSession) GetUploadUrlOk() (string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUploadUrl

`func (o *MicrosoftGraphUploadSession) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### SetUploadUrl

`func (o *MicrosoftGraphUploadSession) SetUploadUrl(v string)`

SetUploadUrl gets a reference to the given string and assigns it to the UploadUrl field.

### SetUploadUrlExplicitNull

`func (o *MicrosoftGraphUploadSession) SetUploadUrlExplicitNull(b bool)`

SetUploadUrlExplicitNull (un)sets UploadUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UploadUrl value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


