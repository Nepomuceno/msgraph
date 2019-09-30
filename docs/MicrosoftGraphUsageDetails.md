# MicrosoftGraphUsageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastAccessedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### GetLastAccessedDateTime

`func (o *MicrosoftGraphUsageDetails) GetLastAccessedDateTime() time.Time`

GetLastAccessedDateTime returns the LastAccessedDateTime field if non-nil, zero value otherwise.

### GetLastAccessedDateTimeOk

`func (o *MicrosoftGraphUsageDetails) GetLastAccessedDateTimeOk() (time.Time, bool)`

GetLastAccessedDateTimeOk returns a tuple with the LastAccessedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastAccessedDateTime

`func (o *MicrosoftGraphUsageDetails) HasLastAccessedDateTime() bool`

HasLastAccessedDateTime returns a boolean if a field has been set.

### SetLastAccessedDateTime

`func (o *MicrosoftGraphUsageDetails) SetLastAccessedDateTime(v time.Time)`

SetLastAccessedDateTime gets a reference to the given time.Time and assigns it to the LastAccessedDateTime field.

### SetLastAccessedDateTimeExplicitNull

`func (o *MicrosoftGraphUsageDetails) SetLastAccessedDateTimeExplicitNull(b bool)`

SetLastAccessedDateTimeExplicitNull (un)sets LastAccessedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastAccessedDateTime value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphUsageDetails) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphUsageDetails) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphUsageDetails) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphUsageDetails) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphUsageDetails) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


