# ManagedDeviceMobileAppConfigurationUserStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserDisplayName** | Pointer to **string** | User name of the DevicePolicyStatus. | [optional] 
**DevicesCount** | Pointer to **int32** | Devices count for that user. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | Compliance status of the policy report. | [optional] 
**LastReportedDateTime** | Pointer to [**time.Time**](time.Time.md) | Last modified date time of the policy report. | [optional] 
**UserPrincipalName** | Pointer to **string** | UserPrincipalName. | [optional] 

## Methods

### GetUserDisplayName

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) GetUserDisplayNameOk() (string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserDisplayName

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayName

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) SetUserDisplayName(v string)`

SetUserDisplayName gets a reference to the given string and assigns it to the UserDisplayName field.

### SetUserDisplayNameExplicitNull

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) SetUserDisplayNameExplicitNull(b bool)`

SetUserDisplayNameExplicitNull (un)sets UserDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserDisplayName value is set to nil even if false is passed
### GetDevicesCount

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) GetDevicesCount() int32`

GetDevicesCount returns the DevicesCount field if non-nil, zero value otherwise.

### GetDevicesCountOk

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) GetDevicesCountOk() (int32, bool)`

GetDevicesCountOk returns a tuple with the DevicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDevicesCount

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) HasDevicesCount() bool`

HasDevicesCount returns a boolean if a field has been set.

### SetDevicesCount

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) SetDevicesCount(v int32)`

SetDevicesCount gets a reference to the given int32 and assigns it to the DevicesCount field.

### GetStatus

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) GetStatus() AnyOfmicrosoftGraphComplianceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) GetStatusOk() (AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) SetStatus(v AnyOfmicrosoftGraphComplianceStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the Status field.

### GetLastReportedDateTime

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) GetLastReportedDateTime() time.Time`

GetLastReportedDateTime returns the LastReportedDateTime field if non-nil, zero value otherwise.

### GetLastReportedDateTimeOk

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) GetLastReportedDateTimeOk() (time.Time, bool)`

GetLastReportedDateTimeOk returns a tuple with the LastReportedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastReportedDateTime

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) HasLastReportedDateTime() bool`

HasLastReportedDateTime returns a boolean if a field has been set.

### SetLastReportedDateTime

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) SetLastReportedDateTime(v time.Time)`

SetLastReportedDateTime gets a reference to the given time.Time and assigns it to the LastReportedDateTime field.

### GetUserPrincipalName

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *ManagedDeviceMobileAppConfigurationUserStatus) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


