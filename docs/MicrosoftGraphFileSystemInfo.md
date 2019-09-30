# MicrosoftGraphFileSystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastAccessedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### GetCreatedDateTime

`func (o *MicrosoftGraphFileSystemInfo) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphFileSystemInfo) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphFileSystemInfo) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphFileSystemInfo) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphFileSystemInfo) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetLastAccessedDateTime

`func (o *MicrosoftGraphFileSystemInfo) GetLastAccessedDateTime() time.Time`

GetLastAccessedDateTime returns the LastAccessedDateTime field if non-nil, zero value otherwise.

### GetLastAccessedDateTimeOk

`func (o *MicrosoftGraphFileSystemInfo) GetLastAccessedDateTimeOk() (time.Time, bool)`

GetLastAccessedDateTimeOk returns a tuple with the LastAccessedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastAccessedDateTime

`func (o *MicrosoftGraphFileSystemInfo) HasLastAccessedDateTime() bool`

HasLastAccessedDateTime returns a boolean if a field has been set.

### SetLastAccessedDateTime

`func (o *MicrosoftGraphFileSystemInfo) SetLastAccessedDateTime(v time.Time)`

SetLastAccessedDateTime gets a reference to the given time.Time and assigns it to the LastAccessedDateTime field.

### SetLastAccessedDateTimeExplicitNull

`func (o *MicrosoftGraphFileSystemInfo) SetLastAccessedDateTimeExplicitNull(b bool)`

SetLastAccessedDateTimeExplicitNull (un)sets LastAccessedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastAccessedDateTime value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphFileSystemInfo) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphFileSystemInfo) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphFileSystemInfo) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphFileSystemInfo) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphFileSystemInfo) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


