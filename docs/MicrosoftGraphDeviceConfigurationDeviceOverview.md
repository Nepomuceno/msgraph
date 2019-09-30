# MicrosoftGraphDeviceConfigurationDeviceOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PendingCount** | Pointer to **int32** | Number of pending devices | [optional] 
**NotApplicableCount** | Pointer to **int32** | Number of not applicable devices | [optional] 
**SuccessCount** | Pointer to **int32** | Number of succeeded devices | [optional] 
**ErrorCount** | Pointer to **int32** | Number of error devices | [optional] 
**FailedCount** | Pointer to **int32** | Number of failed devices | [optional] 
**LastUpdateDateTime** | Pointer to [**time.Time**](time.Time.md) | Last update time | [optional] 
**ConfigurationVersion** | Pointer to **int32** | Version of the policy for that overview | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetPendingCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetPendingCountOk() (int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPendingCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.

### SetPendingCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) SetPendingCount(v int32)`

SetPendingCount gets a reference to the given int32 and assigns it to the PendingCount field.

### GetNotApplicableCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetNotApplicableCount() int32`

GetNotApplicableCount returns the NotApplicableCount field if non-nil, zero value otherwise.

### GetNotApplicableCountOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetNotApplicableCountOk() (int32, bool)`

GetNotApplicableCountOk returns a tuple with the NotApplicableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotApplicableCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) HasNotApplicableCount() bool`

HasNotApplicableCount returns a boolean if a field has been set.

### SetNotApplicableCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) SetNotApplicableCount(v int32)`

SetNotApplicableCount gets a reference to the given int32 and assigns it to the NotApplicableCount field.

### GetSuccessCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetSuccessCountOk() (int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSuccessCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### SetSuccessCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) SetSuccessCount(v int32)`

SetSuccessCount gets a reference to the given int32 and assigns it to the SuccessCount field.

### GetErrorCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetErrorCountOk() (int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### SetErrorCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) SetErrorCount(v int32)`

SetErrorCount gets a reference to the given int32 and assigns it to the ErrorCount field.

### GetFailedCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetFailedCountOk() (int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailedCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### SetFailedCount

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) SetFailedCount(v int32)`

SetFailedCount gets a reference to the given int32 and assigns it to the FailedCount field.

### GetLastUpdateDateTime

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetLastUpdateDateTime() time.Time`

GetLastUpdateDateTime returns the LastUpdateDateTime field if non-nil, zero value otherwise.

### GetLastUpdateDateTimeOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetLastUpdateDateTimeOk() (time.Time, bool)`

GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdateDateTime

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) HasLastUpdateDateTime() bool`

HasLastUpdateDateTime returns a boolean if a field has been set.

### SetLastUpdateDateTime

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) SetLastUpdateDateTime(v time.Time)`

SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.

### GetConfigurationVersion

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetConfigurationVersion() int32`

GetConfigurationVersion returns the ConfigurationVersion field if non-nil, zero value otherwise.

### GetConfigurationVersionOk

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) GetConfigurationVersionOk() (int32, bool)`

GetConfigurationVersionOk returns a tuple with the ConfigurationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfigurationVersion

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) HasConfigurationVersion() bool`

HasConfigurationVersion returns a boolean if a field has been set.

### SetConfigurationVersion

`func (o *MicrosoftGraphDeviceConfigurationDeviceOverview) SetConfigurationVersion(v int32)`

SetConfigurationVersion gets a reference to the given int32 and assigns it to the ConfigurationVersion field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


