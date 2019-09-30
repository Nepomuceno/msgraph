# MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PendingCount** | Pointer to **int32** | Number of pending Users | [optional] 
**NotApplicableCount** | Pointer to **int32** | Number of not applicable users | [optional] 
**SuccessCount** | Pointer to **int32** | Number of succeeded Users | [optional] 
**ErrorCount** | Pointer to **int32** | Number of error Users | [optional] 
**FailedCount** | Pointer to **int32** | Number of failed Users | [optional] 
**LastUpdateDateTime** | Pointer to [**time.Time**](time.Time.md) | Last update time | [optional] 
**ConfigurationVersion** | Pointer to **int32** | Version of the policy for that overview | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetPendingCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetPendingCountOk() (int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPendingCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.

### SetPendingCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) SetPendingCount(v int32)`

SetPendingCount gets a reference to the given int32 and assigns it to the PendingCount field.

### GetNotApplicableCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetNotApplicableCount() int32`

GetNotApplicableCount returns the NotApplicableCount field if non-nil, zero value otherwise.

### GetNotApplicableCountOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetNotApplicableCountOk() (int32, bool)`

GetNotApplicableCountOk returns a tuple with the NotApplicableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotApplicableCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) HasNotApplicableCount() bool`

HasNotApplicableCount returns a boolean if a field has been set.

### SetNotApplicableCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) SetNotApplicableCount(v int32)`

SetNotApplicableCount gets a reference to the given int32 and assigns it to the NotApplicableCount field.

### GetSuccessCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetSuccessCountOk() (int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSuccessCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### SetSuccessCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) SetSuccessCount(v int32)`

SetSuccessCount gets a reference to the given int32 and assigns it to the SuccessCount field.

### GetErrorCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetErrorCountOk() (int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### SetErrorCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) SetErrorCount(v int32)`

SetErrorCount gets a reference to the given int32 and assigns it to the ErrorCount field.

### GetFailedCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetFailedCountOk() (int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailedCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### SetFailedCount

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) SetFailedCount(v int32)`

SetFailedCount gets a reference to the given int32 and assigns it to the FailedCount field.

### GetLastUpdateDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetLastUpdateDateTime() time.Time`

GetLastUpdateDateTime returns the LastUpdateDateTime field if non-nil, zero value otherwise.

### GetLastUpdateDateTimeOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetLastUpdateDateTimeOk() (time.Time, bool)`

GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdateDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) HasLastUpdateDateTime() bool`

HasLastUpdateDateTime returns a boolean if a field has been set.

### SetLastUpdateDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) SetLastUpdateDateTime(v time.Time)`

SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.

### GetConfigurationVersion

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetConfigurationVersion() int32`

GetConfigurationVersion returns the ConfigurationVersion field if non-nil, zero value otherwise.

### GetConfigurationVersionOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) GetConfigurationVersionOk() (int32, bool)`

GetConfigurationVersionOk returns a tuple with the ConfigurationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfigurationVersion

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) HasConfigurationVersion() bool`

HasConfigurationVersion returns a boolean if a field has been set.

### SetConfigurationVersion

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary) SetConfigurationVersion(v int32)`

SetConfigurationVersion gets a reference to the given int32 and assigns it to the ConfigurationVersion field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


