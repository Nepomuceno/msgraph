# MicrosoftGraphDeviceConfigurationUserStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserDisplayName** | Pointer to **string** | User name of the DevicePolicyStatus. | [optional] 
**DevicesCount** | Pointer to **int32** | Devices count for that user. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | Compliance status of the policy report. | [optional] 
**LastReportedDateTime** | Pointer to [**time.Time**](time.Time.md) | Last modified date time of the policy report. | [optional] 
**UserPrincipalName** | Pointer to **string** | UserPrincipalName. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetUserDisplayName

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetUserDisplayNameOk() (string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserDisplayName

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayName

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetUserDisplayName(v string)`

SetUserDisplayName gets a reference to the given string and assigns it to the UserDisplayName field.

### SetUserDisplayNameExplicitNull

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetUserDisplayNameExplicitNull(b bool)`

SetUserDisplayNameExplicitNull (un)sets UserDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserDisplayName value is set to nil even if false is passed
### GetDevicesCount

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetDevicesCount() int32`

GetDevicesCount returns the DevicesCount field if non-nil, zero value otherwise.

### GetDevicesCountOk

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetDevicesCountOk() (int32, bool)`

GetDevicesCountOk returns a tuple with the DevicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDevicesCount

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) HasDevicesCount() bool`

HasDevicesCount returns a boolean if a field has been set.

### SetDevicesCount

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetDevicesCount(v int32)`

SetDevicesCount gets a reference to the given int32 and assigns it to the DevicesCount field.

### GetStatus

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetStatus() AnyOfmicrosoftGraphComplianceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetStatusOk() (AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetStatus(v AnyOfmicrosoftGraphComplianceStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the Status field.

### GetLastReportedDateTime

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetLastReportedDateTime() time.Time`

GetLastReportedDateTime returns the LastReportedDateTime field if non-nil, zero value otherwise.

### GetLastReportedDateTimeOk

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetLastReportedDateTimeOk() (time.Time, bool)`

GetLastReportedDateTimeOk returns a tuple with the LastReportedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastReportedDateTime

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) HasLastReportedDateTime() bool`

HasLastReportedDateTime returns a boolean if a field has been set.

### SetLastReportedDateTime

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetLastReportedDateTime(v time.Time)`

SetLastReportedDateTime gets a reference to the given time.Time and assigns it to the LastReportedDateTime field.

### GetUserPrincipalName

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphDeviceConfigurationUserStatus) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


