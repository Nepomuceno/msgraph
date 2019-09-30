# IosUpdateDeviceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallStatus** | Pointer to [**AnyOfmicrosoftGraphIosUpdatesInstallStatus**](anyOf&lt;microsoft.graph.iosUpdatesInstallStatus&gt;.md) | The installation status of the policy report. | [optional] 
**OsVersion** | Pointer to **string** | The device version that is being reported. | [optional] 
**DeviceId** | Pointer to **string** | The device id that is being reported. | [optional] 
**UserId** | Pointer to **string** | The User id that is being reported. | [optional] 
**DeviceDisplayName** | Pointer to **string** | Device name of the DevicePolicyStatus. | [optional] 
**UserName** | Pointer to **string** | The User Name that is being reported | [optional] 
**DeviceModel** | Pointer to **string** | The device model that is being reported | [optional] 
**ComplianceGracePeriodExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) | The DateTime when device compliance grace period expires | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | Compliance status of the policy report. | [optional] 
**LastReportedDateTime** | Pointer to [**time.Time**](time.Time.md) | Last modified date time of the policy report. | [optional] 
**UserPrincipalName** | Pointer to **string** | UserPrincipalName. | [optional] 

## Methods

### GetInstallStatus

`func (o *IosUpdateDeviceStatus) GetInstallStatus() AnyOfmicrosoftGraphIosUpdatesInstallStatus`

GetInstallStatus returns the InstallStatus field if non-nil, zero value otherwise.

### GetInstallStatusOk

`func (o *IosUpdateDeviceStatus) GetInstallStatusOk() (AnyOfmicrosoftGraphIosUpdatesInstallStatus, bool)`

GetInstallStatusOk returns a tuple with the InstallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstallStatus

`func (o *IosUpdateDeviceStatus) HasInstallStatus() bool`

HasInstallStatus returns a boolean if a field has been set.

### SetInstallStatus

`func (o *IosUpdateDeviceStatus) SetInstallStatus(v AnyOfmicrosoftGraphIosUpdatesInstallStatus)`

SetInstallStatus gets a reference to the given AnyOfmicrosoftGraphIosUpdatesInstallStatus and assigns it to the InstallStatus field.

### GetOsVersion

`func (o *IosUpdateDeviceStatus) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *IosUpdateDeviceStatus) GetOsVersionOk() (string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsVersion

`func (o *IosUpdateDeviceStatus) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersion

`func (o *IosUpdateDeviceStatus) SetOsVersion(v string)`

SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.

### SetOsVersionExplicitNull

`func (o *IosUpdateDeviceStatus) SetOsVersionExplicitNull(b bool)`

SetOsVersionExplicitNull (un)sets OsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsVersion value is set to nil even if false is passed
### GetDeviceId

`func (o *IosUpdateDeviceStatus) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *IosUpdateDeviceStatus) GetDeviceIdOk() (string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceId

`func (o *IosUpdateDeviceStatus) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceId

`func (o *IosUpdateDeviceStatus) SetDeviceId(v string)`

SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.

### SetDeviceIdExplicitNull

`func (o *IosUpdateDeviceStatus) SetDeviceIdExplicitNull(b bool)`

SetDeviceIdExplicitNull (un)sets DeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceId value is set to nil even if false is passed
### GetUserId

`func (o *IosUpdateDeviceStatus) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *IosUpdateDeviceStatus) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *IosUpdateDeviceStatus) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *IosUpdateDeviceStatus) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### SetUserIdExplicitNull

`func (o *IosUpdateDeviceStatus) SetUserIdExplicitNull(b bool)`

SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserId value is set to nil even if false is passed
### GetDeviceDisplayName

`func (o *IosUpdateDeviceStatus) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *IosUpdateDeviceStatus) GetDeviceDisplayNameOk() (string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceDisplayName

`func (o *IosUpdateDeviceStatus) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### SetDeviceDisplayName

`func (o *IosUpdateDeviceStatus) SetDeviceDisplayName(v string)`

SetDeviceDisplayName gets a reference to the given string and assigns it to the DeviceDisplayName field.

### SetDeviceDisplayNameExplicitNull

`func (o *IosUpdateDeviceStatus) SetDeviceDisplayNameExplicitNull(b bool)`

SetDeviceDisplayNameExplicitNull (un)sets DeviceDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceDisplayName value is set to nil even if false is passed
### GetUserName

`func (o *IosUpdateDeviceStatus) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *IosUpdateDeviceStatus) GetUserNameOk() (string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserName

`func (o *IosUpdateDeviceStatus) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserName

`func (o *IosUpdateDeviceStatus) SetUserName(v string)`

SetUserName gets a reference to the given string and assigns it to the UserName field.

### SetUserNameExplicitNull

`func (o *IosUpdateDeviceStatus) SetUserNameExplicitNull(b bool)`

SetUserNameExplicitNull (un)sets UserName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserName value is set to nil even if false is passed
### GetDeviceModel

`func (o *IosUpdateDeviceStatus) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *IosUpdateDeviceStatus) GetDeviceModelOk() (string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceModel

`func (o *IosUpdateDeviceStatus) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### SetDeviceModel

`func (o *IosUpdateDeviceStatus) SetDeviceModel(v string)`

SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.

### SetDeviceModelExplicitNull

`func (o *IosUpdateDeviceStatus) SetDeviceModelExplicitNull(b bool)`

SetDeviceModelExplicitNull (un)sets DeviceModel to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceModel value is set to nil even if false is passed
### GetComplianceGracePeriodExpirationDateTime

`func (o *IosUpdateDeviceStatus) GetComplianceGracePeriodExpirationDateTime() time.Time`

GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceGracePeriodExpirationDateTimeOk

`func (o *IosUpdateDeviceStatus) GetComplianceGracePeriodExpirationDateTimeOk() (time.Time, bool)`

GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplianceGracePeriodExpirationDateTime

`func (o *IosUpdateDeviceStatus) HasComplianceGracePeriodExpirationDateTime() bool`

HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.

### SetComplianceGracePeriodExpirationDateTime

`func (o *IosUpdateDeviceStatus) SetComplianceGracePeriodExpirationDateTime(v time.Time)`

SetComplianceGracePeriodExpirationDateTime gets a reference to the given time.Time and assigns it to the ComplianceGracePeriodExpirationDateTime field.

### GetStatus

`func (o *IosUpdateDeviceStatus) GetStatus() AnyOfmicrosoftGraphComplianceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IosUpdateDeviceStatus) GetStatusOk() (AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *IosUpdateDeviceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *IosUpdateDeviceStatus) SetStatus(v AnyOfmicrosoftGraphComplianceStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the Status field.

### GetLastReportedDateTime

`func (o *IosUpdateDeviceStatus) GetLastReportedDateTime() time.Time`

GetLastReportedDateTime returns the LastReportedDateTime field if non-nil, zero value otherwise.

### GetLastReportedDateTimeOk

`func (o *IosUpdateDeviceStatus) GetLastReportedDateTimeOk() (time.Time, bool)`

GetLastReportedDateTimeOk returns a tuple with the LastReportedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastReportedDateTime

`func (o *IosUpdateDeviceStatus) HasLastReportedDateTime() bool`

HasLastReportedDateTime returns a boolean if a field has been set.

### SetLastReportedDateTime

`func (o *IosUpdateDeviceStatus) SetLastReportedDateTime(v time.Time)`

SetLastReportedDateTime gets a reference to the given time.Time and assigns it to the LastReportedDateTime field.

### GetUserPrincipalName

`func (o *IosUpdateDeviceStatus) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *IosUpdateDeviceStatus) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *IosUpdateDeviceStatus) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *IosUpdateDeviceStatus) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *IosUpdateDeviceStatus) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


