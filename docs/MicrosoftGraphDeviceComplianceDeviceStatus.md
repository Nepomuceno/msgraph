# MicrosoftGraphDeviceComplianceDeviceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DeviceDisplayName** | Pointer to **string** | Device name of the DevicePolicyStatus. | [optional] 
**UserName** | Pointer to **string** | The User Name that is being reported | [optional] 
**DeviceModel** | Pointer to **string** | The device model that is being reported | [optional] 
**ComplianceGracePeriodExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) | The DateTime when device compliance grace period expires | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | Compliance status of the policy report. | [optional] 
**LastReportedDateTime** | Pointer to [**time.Time**](time.Time.md) | Last modified date time of the policy report. | [optional] 
**UserPrincipalName** | Pointer to **string** | UserPrincipalName. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDeviceDisplayName

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetDeviceDisplayNameOk() (string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceDisplayName

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) HasDeviceDisplayName() bool`

HasDeviceDisplayName returns a boolean if a field has been set.

### SetDeviceDisplayName

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) SetDeviceDisplayName(v string)`

SetDeviceDisplayName gets a reference to the given string and assigns it to the DeviceDisplayName field.

### SetDeviceDisplayNameExplicitNull

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) SetDeviceDisplayNameExplicitNull(b bool)`

SetDeviceDisplayNameExplicitNull (un)sets DeviceDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceDisplayName value is set to nil even if false is passed
### GetUserName

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetUserNameOk() (string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserName

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserName

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) SetUserName(v string)`

SetUserName gets a reference to the given string and assigns it to the UserName field.

### SetUserNameExplicitNull

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) SetUserNameExplicitNull(b bool)`

SetUserNameExplicitNull (un)sets UserName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserName value is set to nil even if false is passed
### GetDeviceModel

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetDeviceModelOk() (string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceModel

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### SetDeviceModel

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) SetDeviceModel(v string)`

SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.

### SetDeviceModelExplicitNull

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) SetDeviceModelExplicitNull(b bool)`

SetDeviceModelExplicitNull (un)sets DeviceModel to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceModel value is set to nil even if false is passed
### GetComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetComplianceGracePeriodExpirationDateTime() time.Time`

GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceGracePeriodExpirationDateTimeOk

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetComplianceGracePeriodExpirationDateTimeOk() (time.Time, bool)`

GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) HasComplianceGracePeriodExpirationDateTime() bool`

HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.

### SetComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) SetComplianceGracePeriodExpirationDateTime(v time.Time)`

SetComplianceGracePeriodExpirationDateTime gets a reference to the given time.Time and assigns it to the ComplianceGracePeriodExpirationDateTime field.

### GetStatus

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetStatus() AnyOfmicrosoftGraphComplianceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetStatusOk() (AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) SetStatus(v AnyOfmicrosoftGraphComplianceStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the Status field.

### GetLastReportedDateTime

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetLastReportedDateTime() time.Time`

GetLastReportedDateTime returns the LastReportedDateTime field if non-nil, zero value otherwise.

### GetLastReportedDateTimeOk

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetLastReportedDateTimeOk() (time.Time, bool)`

GetLastReportedDateTimeOk returns a tuple with the LastReportedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastReportedDateTime

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) HasLastReportedDateTime() bool`

HasLastReportedDateTime returns a boolean if a field has been set.

### SetLastReportedDateTime

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) SetLastReportedDateTime(v time.Time)`

SetLastReportedDateTime gets a reference to the given time.Time and assigns it to the LastReportedDateTime field.

### GetUserPrincipalName

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphDeviceComplianceDeviceStatus) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


