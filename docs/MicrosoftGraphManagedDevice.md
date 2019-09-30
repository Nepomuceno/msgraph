# MicrosoftGraphManagedDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** | Unique Identifier for the user associated with the device | [optional] 
**DeviceName** | Pointer to **string** | Name of the device | [optional] 
**ManagedDeviceOwnerType** | Pointer to [**AnyOfmicrosoftGraphManagedDeviceOwnerType**](anyOf&lt;microsoft.graph.managedDeviceOwnerType&gt;.md) | Ownership of the device. Can be &#39;company&#39; or &#39;personal&#39; | [optional] 
**DeviceActionResults** | Pointer to [**[]AnyOfmicrosoftGraphDeviceActionResult**](anyOf&lt;microsoft.graph.deviceActionResult&gt;.md) | List of ComplexType deviceActionResult objects. | [optional] 
**EnrolledDateTime** | Pointer to [**time.Time**](time.Time.md) | Enrollment time of the device. | [optional] 
**LastSyncDateTime** | Pointer to [**time.Time**](time.Time.md) | The date and time that the device last completed a successful sync with Intune. | [optional] 
**OperatingSystem** | Pointer to **string** | Operating system of the device. Windows, iOS, etc. | [optional] 
**ComplianceState** | Pointer to [**AnyOfmicrosoftGraphComplianceState**](anyOf&lt;microsoft.graph.complianceState&gt;.md) | Compliance state of the device. | [optional] 
**JailBroken** | Pointer to **string** | whether the device is jail broken or rooted. | [optional] 
**ManagementAgent** | Pointer to [**AnyOfmicrosoftGraphManagementAgentType**](anyOf&lt;microsoft.graph.managementAgentType&gt;.md) | Management channel of the device. Intune, EAS, etc. | [optional] 
**OsVersion** | Pointer to **string** | Operating system version of the device. | [optional] 
**EasActivated** | Pointer to **bool** | Whether the device is Exchange ActiveSync activated. | [optional] 
**EasDeviceId** | Pointer to **string** | Exchange ActiveSync Id of the device. | [optional] 
**EasActivationDateTime** | Pointer to [**time.Time**](time.Time.md) | Exchange ActivationSync activation time of the device. | [optional] 
**AzureADRegistered** | Pointer to **bool** | Whether the device is Azure Active Directory registered. | [optional] 
**DeviceEnrollmentType** | Pointer to [**AnyOfmicrosoftGraphDeviceEnrollmentType**](anyOf&lt;microsoft.graph.deviceEnrollmentType&gt;.md) | Enrollment type of the device. | [optional] 
**ActivationLockBypassCode** | Pointer to **string** | Code that allows the Activation Lock on a device to be bypassed. | [optional] 
**EmailAddress** | Pointer to **string** | Email(s) for the user associated with the device | [optional] 
**AzureADDeviceId** | Pointer to **string** | The unique identifier for the Azure Active Directory device. Read only. | [optional] 
**DeviceRegistrationState** | Pointer to [**AnyOfmicrosoftGraphDeviceRegistrationState**](anyOf&lt;microsoft.graph.deviceRegistrationState&gt;.md) | Device registration state. | [optional] 
**DeviceCategoryDisplayName** | Pointer to **string** | Device category display name | [optional] 
**IsSupervised** | Pointer to **bool** | Device supervised status | [optional] 
**ExchangeLastSuccessfulSyncDateTime** | Pointer to [**time.Time**](time.Time.md) | Last time the device contacted Exchange. | [optional] 
**ExchangeAccessState** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementExchangeAccessState**](anyOf&lt;microsoft.graph.deviceManagementExchangeAccessState&gt;.md) | The Access State of the device in Exchange. | [optional] 
**ExchangeAccessStateReason** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason**](anyOf&lt;microsoft.graph.deviceManagementExchangeAccessStateReason&gt;.md) | The reason for the device&#39;s access state in Exchange. | [optional] 
**RemoteAssistanceSessionUrl** | Pointer to **string** | Url that allows a Remote Assistance session to be established with the device. | [optional] 
**RemoteAssistanceSessionErrorDetails** | Pointer to **string** | An error string that identifies issues when creating Remote Assistance session objects. | [optional] 
**IsEncrypted** | Pointer to **bool** | Device encryption status | [optional] 
**UserPrincipalName** | Pointer to **string** | Device user principal name | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**Manufacturer** | Pointer to **string** | Manufacturer of the device | [optional] 
**Imei** | Pointer to **string** | IMEI | [optional] 
**ComplianceGracePeriodExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) | The DateTime when device compliance grace period expires | [optional] 
**SerialNumber** | Pointer to **string** | SerialNumber | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number of the device | [optional] 
**AndroidSecurityPatchLevel** | Pointer to **string** | Android security patch level | [optional] 
**UserDisplayName** | Pointer to **string** | User display name | [optional] 
**ConfigurationManagerClientEnabledFeatures** | Pointer to [**AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures**](anyOf&lt;microsoft.graph.configurationManagerClientEnabledFeatures&gt;.md) | ConfigrMgr client enabled features | [optional] 
**WiFiMacAddress** | Pointer to **string** | Wi-Fi MAC | [optional] 
**DeviceHealthAttestationState** | Pointer to [**AnyOfmicrosoftGraphDeviceHealthAttestationState**](anyOf&lt;microsoft.graph.deviceHealthAttestationState&gt;.md) | The device health attestation state. | [optional] 
**SubscriberCarrier** | Pointer to **string** | Subscriber Carrier | [optional] 
**Meid** | Pointer to **string** | MEID | [optional] 
**TotalStorageSpaceInBytes** | Pointer to **int64** | Total Storage in Bytes | [optional] 
**FreeStorageSpaceInBytes** | Pointer to **int64** | Free Storage in Bytes | [optional] 
**ManagedDeviceName** | Pointer to **string** | Automatically generated name to identify a device. Can be overwritten to a user friendly name. | [optional] 
**PartnerReportedThreatState** | Pointer to [**AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState**](anyOf&lt;microsoft.graph.managedDevicePartnerReportedHealthState&gt;.md) | Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. | [optional] 
**DeviceConfigurationStates** | Pointer to [**[]MicrosoftGraphDeviceConfigurationState**](microsoft.graph.deviceConfigurationState.md) |  | [optional] 
**DeviceCompliancePolicyStates** | Pointer to [**[]MicrosoftGraphDeviceCompliancePolicyState**](microsoft.graph.deviceCompliancePolicyState.md) |  | [optional] 
**DeviceCategory** | Pointer to [**AnyOfmicrosoftGraphDeviceCategory**](anyOf&lt;microsoft.graph.deviceCategory&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphManagedDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedDevice) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphManagedDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphManagedDevice) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetUserId

`func (o *MicrosoftGraphManagedDevice) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphManagedDevice) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *MicrosoftGraphManagedDevice) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *MicrosoftGraphManagedDevice) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### SetUserIdExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetUserIdExplicitNull(b bool)`

SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserId value is set to nil even if false is passed
### GetDeviceName

`func (o *MicrosoftGraphManagedDevice) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MicrosoftGraphManagedDevice) GetDeviceNameOk() (string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceName

`func (o *MicrosoftGraphManagedDevice) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceName

`func (o *MicrosoftGraphManagedDevice) SetDeviceName(v string)`

SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.

### SetDeviceNameExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetDeviceNameExplicitNull(b bool)`

SetDeviceNameExplicitNull (un)sets DeviceName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceName value is set to nil even if false is passed
### GetManagedDeviceOwnerType

`func (o *MicrosoftGraphManagedDevice) GetManagedDeviceOwnerType() AnyOfmicrosoftGraphManagedDeviceOwnerType`

GetManagedDeviceOwnerType returns the ManagedDeviceOwnerType field if non-nil, zero value otherwise.

### GetManagedDeviceOwnerTypeOk

`func (o *MicrosoftGraphManagedDevice) GetManagedDeviceOwnerTypeOk() (AnyOfmicrosoftGraphManagedDeviceOwnerType, bool)`

GetManagedDeviceOwnerTypeOk returns a tuple with the ManagedDeviceOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedDeviceOwnerType

`func (o *MicrosoftGraphManagedDevice) HasManagedDeviceOwnerType() bool`

HasManagedDeviceOwnerType returns a boolean if a field has been set.

### SetManagedDeviceOwnerType

`func (o *MicrosoftGraphManagedDevice) SetManagedDeviceOwnerType(v AnyOfmicrosoftGraphManagedDeviceOwnerType)`

SetManagedDeviceOwnerType gets a reference to the given AnyOfmicrosoftGraphManagedDeviceOwnerType and assigns it to the ManagedDeviceOwnerType field.

### GetDeviceActionResults

`func (o *MicrosoftGraphManagedDevice) GetDeviceActionResults() []AnyOfmicrosoftGraphDeviceActionResult`

GetDeviceActionResults returns the DeviceActionResults field if non-nil, zero value otherwise.

### GetDeviceActionResultsOk

`func (o *MicrosoftGraphManagedDevice) GetDeviceActionResultsOk() ([]AnyOfmicrosoftGraphDeviceActionResult, bool)`

GetDeviceActionResultsOk returns a tuple with the DeviceActionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceActionResults

`func (o *MicrosoftGraphManagedDevice) HasDeviceActionResults() bool`

HasDeviceActionResults returns a boolean if a field has been set.

### SetDeviceActionResults

`func (o *MicrosoftGraphManagedDevice) SetDeviceActionResults(v []AnyOfmicrosoftGraphDeviceActionResult)`

SetDeviceActionResults gets a reference to the given []AnyOfmicrosoftGraphDeviceActionResult and assigns it to the DeviceActionResults field.

### GetEnrolledDateTime

`func (o *MicrosoftGraphManagedDevice) GetEnrolledDateTime() time.Time`

GetEnrolledDateTime returns the EnrolledDateTime field if non-nil, zero value otherwise.

### GetEnrolledDateTimeOk

`func (o *MicrosoftGraphManagedDevice) GetEnrolledDateTimeOk() (time.Time, bool)`

GetEnrolledDateTimeOk returns a tuple with the EnrolledDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnrolledDateTime

`func (o *MicrosoftGraphManagedDevice) HasEnrolledDateTime() bool`

HasEnrolledDateTime returns a boolean if a field has been set.

### SetEnrolledDateTime

`func (o *MicrosoftGraphManagedDevice) SetEnrolledDateTime(v time.Time)`

SetEnrolledDateTime gets a reference to the given time.Time and assigns it to the EnrolledDateTime field.

### GetLastSyncDateTime

`func (o *MicrosoftGraphManagedDevice) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *MicrosoftGraphManagedDevice) GetLastSyncDateTimeOk() (time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSyncDateTime

`func (o *MicrosoftGraphManagedDevice) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### SetLastSyncDateTime

`func (o *MicrosoftGraphManagedDevice) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.

### GetOperatingSystem

`func (o *MicrosoftGraphManagedDevice) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *MicrosoftGraphManagedDevice) GetOperatingSystemOk() (string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperatingSystem

`func (o *MicrosoftGraphManagedDevice) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystem

`func (o *MicrosoftGraphManagedDevice) SetOperatingSystem(v string)`

SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.

### SetOperatingSystemExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetOperatingSystemExplicitNull(b bool)`

SetOperatingSystemExplicitNull (un)sets OperatingSystem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperatingSystem value is set to nil even if false is passed
### GetComplianceState

`func (o *MicrosoftGraphManagedDevice) GetComplianceState() AnyOfmicrosoftGraphComplianceState`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *MicrosoftGraphManagedDevice) GetComplianceStateOk() (AnyOfmicrosoftGraphComplianceState, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplianceState

`func (o *MicrosoftGraphManagedDevice) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.

### SetComplianceState

`func (o *MicrosoftGraphManagedDevice) SetComplianceState(v AnyOfmicrosoftGraphComplianceState)`

SetComplianceState gets a reference to the given AnyOfmicrosoftGraphComplianceState and assigns it to the ComplianceState field.

### GetJailBroken

`func (o *MicrosoftGraphManagedDevice) GetJailBroken() string`

GetJailBroken returns the JailBroken field if non-nil, zero value otherwise.

### GetJailBrokenOk

`func (o *MicrosoftGraphManagedDevice) GetJailBrokenOk() (string, bool)`

GetJailBrokenOk returns a tuple with the JailBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJailBroken

`func (o *MicrosoftGraphManagedDevice) HasJailBroken() bool`

HasJailBroken returns a boolean if a field has been set.

### SetJailBroken

`func (o *MicrosoftGraphManagedDevice) SetJailBroken(v string)`

SetJailBroken gets a reference to the given string and assigns it to the JailBroken field.

### SetJailBrokenExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetJailBrokenExplicitNull(b bool)`

SetJailBrokenExplicitNull (un)sets JailBroken to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The JailBroken value is set to nil even if false is passed
### GetManagementAgent

`func (o *MicrosoftGraphManagedDevice) GetManagementAgent() AnyOfmicrosoftGraphManagementAgentType`

GetManagementAgent returns the ManagementAgent field if non-nil, zero value otherwise.

### GetManagementAgentOk

`func (o *MicrosoftGraphManagedDevice) GetManagementAgentOk() (AnyOfmicrosoftGraphManagementAgentType, bool)`

GetManagementAgentOk returns a tuple with the ManagementAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagementAgent

`func (o *MicrosoftGraphManagedDevice) HasManagementAgent() bool`

HasManagementAgent returns a boolean if a field has been set.

### SetManagementAgent

`func (o *MicrosoftGraphManagedDevice) SetManagementAgent(v AnyOfmicrosoftGraphManagementAgentType)`

SetManagementAgent gets a reference to the given AnyOfmicrosoftGraphManagementAgentType and assigns it to the ManagementAgent field.

### GetOsVersion

`func (o *MicrosoftGraphManagedDevice) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *MicrosoftGraphManagedDevice) GetOsVersionOk() (string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsVersion

`func (o *MicrosoftGraphManagedDevice) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersion

`func (o *MicrosoftGraphManagedDevice) SetOsVersion(v string)`

SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.

### SetOsVersionExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetOsVersionExplicitNull(b bool)`

SetOsVersionExplicitNull (un)sets OsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsVersion value is set to nil even if false is passed
### GetEasActivated

`func (o *MicrosoftGraphManagedDevice) GetEasActivated() bool`

GetEasActivated returns the EasActivated field if non-nil, zero value otherwise.

### GetEasActivatedOk

`func (o *MicrosoftGraphManagedDevice) GetEasActivatedOk() (bool, bool)`

GetEasActivatedOk returns a tuple with the EasActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEasActivated

`func (o *MicrosoftGraphManagedDevice) HasEasActivated() bool`

HasEasActivated returns a boolean if a field has been set.

### SetEasActivated

`func (o *MicrosoftGraphManagedDevice) SetEasActivated(v bool)`

SetEasActivated gets a reference to the given bool and assigns it to the EasActivated field.

### GetEasDeviceId

`func (o *MicrosoftGraphManagedDevice) GetEasDeviceId() string`

GetEasDeviceId returns the EasDeviceId field if non-nil, zero value otherwise.

### GetEasDeviceIdOk

`func (o *MicrosoftGraphManagedDevice) GetEasDeviceIdOk() (string, bool)`

GetEasDeviceIdOk returns a tuple with the EasDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEasDeviceId

`func (o *MicrosoftGraphManagedDevice) HasEasDeviceId() bool`

HasEasDeviceId returns a boolean if a field has been set.

### SetEasDeviceId

`func (o *MicrosoftGraphManagedDevice) SetEasDeviceId(v string)`

SetEasDeviceId gets a reference to the given string and assigns it to the EasDeviceId field.

### SetEasDeviceIdExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetEasDeviceIdExplicitNull(b bool)`

SetEasDeviceIdExplicitNull (un)sets EasDeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EasDeviceId value is set to nil even if false is passed
### GetEasActivationDateTime

`func (o *MicrosoftGraphManagedDevice) GetEasActivationDateTime() time.Time`

GetEasActivationDateTime returns the EasActivationDateTime field if non-nil, zero value otherwise.

### GetEasActivationDateTimeOk

`func (o *MicrosoftGraphManagedDevice) GetEasActivationDateTimeOk() (time.Time, bool)`

GetEasActivationDateTimeOk returns a tuple with the EasActivationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEasActivationDateTime

`func (o *MicrosoftGraphManagedDevice) HasEasActivationDateTime() bool`

HasEasActivationDateTime returns a boolean if a field has been set.

### SetEasActivationDateTime

`func (o *MicrosoftGraphManagedDevice) SetEasActivationDateTime(v time.Time)`

SetEasActivationDateTime gets a reference to the given time.Time and assigns it to the EasActivationDateTime field.

### GetAzureADRegistered

`func (o *MicrosoftGraphManagedDevice) GetAzureADRegistered() bool`

GetAzureADRegistered returns the AzureADRegistered field if non-nil, zero value otherwise.

### GetAzureADRegisteredOk

`func (o *MicrosoftGraphManagedDevice) GetAzureADRegisteredOk() (bool, bool)`

GetAzureADRegisteredOk returns a tuple with the AzureADRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureADRegistered

`func (o *MicrosoftGraphManagedDevice) HasAzureADRegistered() bool`

HasAzureADRegistered returns a boolean if a field has been set.

### SetAzureADRegistered

`func (o *MicrosoftGraphManagedDevice) SetAzureADRegistered(v bool)`

SetAzureADRegistered gets a reference to the given bool and assigns it to the AzureADRegistered field.

### SetAzureADRegisteredExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetAzureADRegisteredExplicitNull(b bool)`

SetAzureADRegisteredExplicitNull (un)sets AzureADRegistered to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AzureADRegistered value is set to nil even if false is passed
### GetDeviceEnrollmentType

`func (o *MicrosoftGraphManagedDevice) GetDeviceEnrollmentType() AnyOfmicrosoftGraphDeviceEnrollmentType`

GetDeviceEnrollmentType returns the DeviceEnrollmentType field if non-nil, zero value otherwise.

### GetDeviceEnrollmentTypeOk

`func (o *MicrosoftGraphManagedDevice) GetDeviceEnrollmentTypeOk() (AnyOfmicrosoftGraphDeviceEnrollmentType, bool)`

GetDeviceEnrollmentTypeOk returns a tuple with the DeviceEnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceEnrollmentType

`func (o *MicrosoftGraphManagedDevice) HasDeviceEnrollmentType() bool`

HasDeviceEnrollmentType returns a boolean if a field has been set.

### SetDeviceEnrollmentType

`func (o *MicrosoftGraphManagedDevice) SetDeviceEnrollmentType(v AnyOfmicrosoftGraphDeviceEnrollmentType)`

SetDeviceEnrollmentType gets a reference to the given AnyOfmicrosoftGraphDeviceEnrollmentType and assigns it to the DeviceEnrollmentType field.

### GetActivationLockBypassCode

`func (o *MicrosoftGraphManagedDevice) GetActivationLockBypassCode() string`

GetActivationLockBypassCode returns the ActivationLockBypassCode field if non-nil, zero value otherwise.

### GetActivationLockBypassCodeOk

`func (o *MicrosoftGraphManagedDevice) GetActivationLockBypassCodeOk() (string, bool)`

GetActivationLockBypassCodeOk returns a tuple with the ActivationLockBypassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivationLockBypassCode

`func (o *MicrosoftGraphManagedDevice) HasActivationLockBypassCode() bool`

HasActivationLockBypassCode returns a boolean if a field has been set.

### SetActivationLockBypassCode

`func (o *MicrosoftGraphManagedDevice) SetActivationLockBypassCode(v string)`

SetActivationLockBypassCode gets a reference to the given string and assigns it to the ActivationLockBypassCode field.

### SetActivationLockBypassCodeExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetActivationLockBypassCodeExplicitNull(b bool)`

SetActivationLockBypassCodeExplicitNull (un)sets ActivationLockBypassCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActivationLockBypassCode value is set to nil even if false is passed
### GetEmailAddress

`func (o *MicrosoftGraphManagedDevice) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MicrosoftGraphManagedDevice) GetEmailAddressOk() (string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailAddress

`func (o *MicrosoftGraphManagedDevice) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddress

`func (o *MicrosoftGraphManagedDevice) SetEmailAddress(v string)`

SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.

### SetEmailAddressExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetEmailAddressExplicitNull(b bool)`

SetEmailAddressExplicitNull (un)sets EmailAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EmailAddress value is set to nil even if false is passed
### GetAzureADDeviceId

`func (o *MicrosoftGraphManagedDevice) GetAzureADDeviceId() string`

GetAzureADDeviceId returns the AzureADDeviceId field if non-nil, zero value otherwise.

### GetAzureADDeviceIdOk

`func (o *MicrosoftGraphManagedDevice) GetAzureADDeviceIdOk() (string, bool)`

GetAzureADDeviceIdOk returns a tuple with the AzureADDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureADDeviceId

`func (o *MicrosoftGraphManagedDevice) HasAzureADDeviceId() bool`

HasAzureADDeviceId returns a boolean if a field has been set.

### SetAzureADDeviceId

`func (o *MicrosoftGraphManagedDevice) SetAzureADDeviceId(v string)`

SetAzureADDeviceId gets a reference to the given string and assigns it to the AzureADDeviceId field.

### SetAzureADDeviceIdExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetAzureADDeviceIdExplicitNull(b bool)`

SetAzureADDeviceIdExplicitNull (un)sets AzureADDeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AzureADDeviceId value is set to nil even if false is passed
### GetDeviceRegistrationState

`func (o *MicrosoftGraphManagedDevice) GetDeviceRegistrationState() AnyOfmicrosoftGraphDeviceRegistrationState`

GetDeviceRegistrationState returns the DeviceRegistrationState field if non-nil, zero value otherwise.

### GetDeviceRegistrationStateOk

`func (o *MicrosoftGraphManagedDevice) GetDeviceRegistrationStateOk() (AnyOfmicrosoftGraphDeviceRegistrationState, bool)`

GetDeviceRegistrationStateOk returns a tuple with the DeviceRegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceRegistrationState

`func (o *MicrosoftGraphManagedDevice) HasDeviceRegistrationState() bool`

HasDeviceRegistrationState returns a boolean if a field has been set.

### SetDeviceRegistrationState

`func (o *MicrosoftGraphManagedDevice) SetDeviceRegistrationState(v AnyOfmicrosoftGraphDeviceRegistrationState)`

SetDeviceRegistrationState gets a reference to the given AnyOfmicrosoftGraphDeviceRegistrationState and assigns it to the DeviceRegistrationState field.

### GetDeviceCategoryDisplayName

`func (o *MicrosoftGraphManagedDevice) GetDeviceCategoryDisplayName() string`

GetDeviceCategoryDisplayName returns the DeviceCategoryDisplayName field if non-nil, zero value otherwise.

### GetDeviceCategoryDisplayNameOk

`func (o *MicrosoftGraphManagedDevice) GetDeviceCategoryDisplayNameOk() (string, bool)`

GetDeviceCategoryDisplayNameOk returns a tuple with the DeviceCategoryDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCategoryDisplayName

`func (o *MicrosoftGraphManagedDevice) HasDeviceCategoryDisplayName() bool`

HasDeviceCategoryDisplayName returns a boolean if a field has been set.

### SetDeviceCategoryDisplayName

`func (o *MicrosoftGraphManagedDevice) SetDeviceCategoryDisplayName(v string)`

SetDeviceCategoryDisplayName gets a reference to the given string and assigns it to the DeviceCategoryDisplayName field.

### SetDeviceCategoryDisplayNameExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetDeviceCategoryDisplayNameExplicitNull(b bool)`

SetDeviceCategoryDisplayNameExplicitNull (un)sets DeviceCategoryDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceCategoryDisplayName value is set to nil even if false is passed
### GetIsSupervised

`func (o *MicrosoftGraphManagedDevice) GetIsSupervised() bool`

GetIsSupervised returns the IsSupervised field if non-nil, zero value otherwise.

### GetIsSupervisedOk

`func (o *MicrosoftGraphManagedDevice) GetIsSupervisedOk() (bool, bool)`

GetIsSupervisedOk returns a tuple with the IsSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsSupervised

`func (o *MicrosoftGraphManagedDevice) HasIsSupervised() bool`

HasIsSupervised returns a boolean if a field has been set.

### SetIsSupervised

`func (o *MicrosoftGraphManagedDevice) SetIsSupervised(v bool)`

SetIsSupervised gets a reference to the given bool and assigns it to the IsSupervised field.

### GetExchangeLastSuccessfulSyncDateTime

`func (o *MicrosoftGraphManagedDevice) GetExchangeLastSuccessfulSyncDateTime() time.Time`

GetExchangeLastSuccessfulSyncDateTime returns the ExchangeLastSuccessfulSyncDateTime field if non-nil, zero value otherwise.

### GetExchangeLastSuccessfulSyncDateTimeOk

`func (o *MicrosoftGraphManagedDevice) GetExchangeLastSuccessfulSyncDateTimeOk() (time.Time, bool)`

GetExchangeLastSuccessfulSyncDateTimeOk returns a tuple with the ExchangeLastSuccessfulSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeLastSuccessfulSyncDateTime

`func (o *MicrosoftGraphManagedDevice) HasExchangeLastSuccessfulSyncDateTime() bool`

HasExchangeLastSuccessfulSyncDateTime returns a boolean if a field has been set.

### SetExchangeLastSuccessfulSyncDateTime

`func (o *MicrosoftGraphManagedDevice) SetExchangeLastSuccessfulSyncDateTime(v time.Time)`

SetExchangeLastSuccessfulSyncDateTime gets a reference to the given time.Time and assigns it to the ExchangeLastSuccessfulSyncDateTime field.

### GetExchangeAccessState

`func (o *MicrosoftGraphManagedDevice) GetExchangeAccessState() AnyOfmicrosoftGraphDeviceManagementExchangeAccessState`

GetExchangeAccessState returns the ExchangeAccessState field if non-nil, zero value otherwise.

### GetExchangeAccessStateOk

`func (o *MicrosoftGraphManagedDevice) GetExchangeAccessStateOk() (AnyOfmicrosoftGraphDeviceManagementExchangeAccessState, bool)`

GetExchangeAccessStateOk returns a tuple with the ExchangeAccessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeAccessState

`func (o *MicrosoftGraphManagedDevice) HasExchangeAccessState() bool`

HasExchangeAccessState returns a boolean if a field has been set.

### SetExchangeAccessState

`func (o *MicrosoftGraphManagedDevice) SetExchangeAccessState(v AnyOfmicrosoftGraphDeviceManagementExchangeAccessState)`

SetExchangeAccessState gets a reference to the given AnyOfmicrosoftGraphDeviceManagementExchangeAccessState and assigns it to the ExchangeAccessState field.

### GetExchangeAccessStateReason

`func (o *MicrosoftGraphManagedDevice) GetExchangeAccessStateReason() AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason`

GetExchangeAccessStateReason returns the ExchangeAccessStateReason field if non-nil, zero value otherwise.

### GetExchangeAccessStateReasonOk

`func (o *MicrosoftGraphManagedDevice) GetExchangeAccessStateReasonOk() (AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason, bool)`

GetExchangeAccessStateReasonOk returns a tuple with the ExchangeAccessStateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeAccessStateReason

`func (o *MicrosoftGraphManagedDevice) HasExchangeAccessStateReason() bool`

HasExchangeAccessStateReason returns a boolean if a field has been set.

### SetExchangeAccessStateReason

`func (o *MicrosoftGraphManagedDevice) SetExchangeAccessStateReason(v AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason)`

SetExchangeAccessStateReason gets a reference to the given AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason and assigns it to the ExchangeAccessStateReason field.

### GetRemoteAssistanceSessionUrl

`func (o *MicrosoftGraphManagedDevice) GetRemoteAssistanceSessionUrl() string`

GetRemoteAssistanceSessionUrl returns the RemoteAssistanceSessionUrl field if non-nil, zero value otherwise.

### GetRemoteAssistanceSessionUrlOk

`func (o *MicrosoftGraphManagedDevice) GetRemoteAssistanceSessionUrlOk() (string, bool)`

GetRemoteAssistanceSessionUrlOk returns a tuple with the RemoteAssistanceSessionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemoteAssistanceSessionUrl

`func (o *MicrosoftGraphManagedDevice) HasRemoteAssistanceSessionUrl() bool`

HasRemoteAssistanceSessionUrl returns a boolean if a field has been set.

### SetRemoteAssistanceSessionUrl

`func (o *MicrosoftGraphManagedDevice) SetRemoteAssistanceSessionUrl(v string)`

SetRemoteAssistanceSessionUrl gets a reference to the given string and assigns it to the RemoteAssistanceSessionUrl field.

### SetRemoteAssistanceSessionUrlExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetRemoteAssistanceSessionUrlExplicitNull(b bool)`

SetRemoteAssistanceSessionUrlExplicitNull (un)sets RemoteAssistanceSessionUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RemoteAssistanceSessionUrl value is set to nil even if false is passed
### GetRemoteAssistanceSessionErrorDetails

`func (o *MicrosoftGraphManagedDevice) GetRemoteAssistanceSessionErrorDetails() string`

GetRemoteAssistanceSessionErrorDetails returns the RemoteAssistanceSessionErrorDetails field if non-nil, zero value otherwise.

### GetRemoteAssistanceSessionErrorDetailsOk

`func (o *MicrosoftGraphManagedDevice) GetRemoteAssistanceSessionErrorDetailsOk() (string, bool)`

GetRemoteAssistanceSessionErrorDetailsOk returns a tuple with the RemoteAssistanceSessionErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemoteAssistanceSessionErrorDetails

`func (o *MicrosoftGraphManagedDevice) HasRemoteAssistanceSessionErrorDetails() bool`

HasRemoteAssistanceSessionErrorDetails returns a boolean if a field has been set.

### SetRemoteAssistanceSessionErrorDetails

`func (o *MicrosoftGraphManagedDevice) SetRemoteAssistanceSessionErrorDetails(v string)`

SetRemoteAssistanceSessionErrorDetails gets a reference to the given string and assigns it to the RemoteAssistanceSessionErrorDetails field.

### SetRemoteAssistanceSessionErrorDetailsExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetRemoteAssistanceSessionErrorDetailsExplicitNull(b bool)`

SetRemoteAssistanceSessionErrorDetailsExplicitNull (un)sets RemoteAssistanceSessionErrorDetails to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RemoteAssistanceSessionErrorDetails value is set to nil even if false is passed
### GetIsEncrypted

`func (o *MicrosoftGraphManagedDevice) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *MicrosoftGraphManagedDevice) GetIsEncryptedOk() (bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEncrypted

`func (o *MicrosoftGraphManagedDevice) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### SetIsEncrypted

`func (o *MicrosoftGraphManagedDevice) SetIsEncrypted(v bool)`

SetIsEncrypted gets a reference to the given bool and assigns it to the IsEncrypted field.

### GetUserPrincipalName

`func (o *MicrosoftGraphManagedDevice) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphManagedDevice) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *MicrosoftGraphManagedDevice) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphManagedDevice) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetModel

`func (o *MicrosoftGraphManagedDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MicrosoftGraphManagedDevice) GetModelOk() (string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModel

`func (o *MicrosoftGraphManagedDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModel

`func (o *MicrosoftGraphManagedDevice) SetModel(v string)`

SetModel gets a reference to the given string and assigns it to the Model field.

### SetModelExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetModelExplicitNull(b bool)`

SetModelExplicitNull (un)sets Model to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Model value is set to nil even if false is passed
### GetManufacturer

`func (o *MicrosoftGraphManagedDevice) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *MicrosoftGraphManagedDevice) GetManufacturerOk() (string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManufacturer

`func (o *MicrosoftGraphManagedDevice) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturer

`func (o *MicrosoftGraphManagedDevice) SetManufacturer(v string)`

SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.

### SetManufacturerExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetManufacturerExplicitNull(b bool)`

SetManufacturerExplicitNull (un)sets Manufacturer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Manufacturer value is set to nil even if false is passed
### GetImei

`func (o *MicrosoftGraphManagedDevice) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *MicrosoftGraphManagedDevice) GetImeiOk() (string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImei

`func (o *MicrosoftGraphManagedDevice) HasImei() bool`

HasImei returns a boolean if a field has been set.

### SetImei

`func (o *MicrosoftGraphManagedDevice) SetImei(v string)`

SetImei gets a reference to the given string and assigns it to the Imei field.

### SetImeiExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetImeiExplicitNull(b bool)`

SetImeiExplicitNull (un)sets Imei to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Imei value is set to nil even if false is passed
### GetComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphManagedDevice) GetComplianceGracePeriodExpirationDateTime() time.Time`

GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceGracePeriodExpirationDateTimeOk

`func (o *MicrosoftGraphManagedDevice) GetComplianceGracePeriodExpirationDateTimeOk() (time.Time, bool)`

GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphManagedDevice) HasComplianceGracePeriodExpirationDateTime() bool`

HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.

### SetComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphManagedDevice) SetComplianceGracePeriodExpirationDateTime(v time.Time)`

SetComplianceGracePeriodExpirationDateTime gets a reference to the given time.Time and assigns it to the ComplianceGracePeriodExpirationDateTime field.

### GetSerialNumber

`func (o *MicrosoftGraphManagedDevice) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MicrosoftGraphManagedDevice) GetSerialNumberOk() (string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSerialNumber

`func (o *MicrosoftGraphManagedDevice) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumber

`func (o *MicrosoftGraphManagedDevice) SetSerialNumber(v string)`

SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.

### SetSerialNumberExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetSerialNumberExplicitNull(b bool)`

SetSerialNumberExplicitNull (un)sets SerialNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SerialNumber value is set to nil even if false is passed
### GetPhoneNumber

`func (o *MicrosoftGraphManagedDevice) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *MicrosoftGraphManagedDevice) GetPhoneNumberOk() (string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhoneNumber

`func (o *MicrosoftGraphManagedDevice) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumber

`func (o *MicrosoftGraphManagedDevice) SetPhoneNumber(v string)`

SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.

### SetPhoneNumberExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetPhoneNumberExplicitNull(b bool)`

SetPhoneNumberExplicitNull (un)sets PhoneNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PhoneNumber value is set to nil even if false is passed
### GetAndroidSecurityPatchLevel

`func (o *MicrosoftGraphManagedDevice) GetAndroidSecurityPatchLevel() string`

GetAndroidSecurityPatchLevel returns the AndroidSecurityPatchLevel field if non-nil, zero value otherwise.

### GetAndroidSecurityPatchLevelOk

`func (o *MicrosoftGraphManagedDevice) GetAndroidSecurityPatchLevelOk() (string, bool)`

GetAndroidSecurityPatchLevelOk returns a tuple with the AndroidSecurityPatchLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAndroidSecurityPatchLevel

`func (o *MicrosoftGraphManagedDevice) HasAndroidSecurityPatchLevel() bool`

HasAndroidSecurityPatchLevel returns a boolean if a field has been set.

### SetAndroidSecurityPatchLevel

`func (o *MicrosoftGraphManagedDevice) SetAndroidSecurityPatchLevel(v string)`

SetAndroidSecurityPatchLevel gets a reference to the given string and assigns it to the AndroidSecurityPatchLevel field.

### SetAndroidSecurityPatchLevelExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetAndroidSecurityPatchLevelExplicitNull(b bool)`

SetAndroidSecurityPatchLevelExplicitNull (un)sets AndroidSecurityPatchLevel to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AndroidSecurityPatchLevel value is set to nil even if false is passed
### GetUserDisplayName

`func (o *MicrosoftGraphManagedDevice) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *MicrosoftGraphManagedDevice) GetUserDisplayNameOk() (string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserDisplayName

`func (o *MicrosoftGraphManagedDevice) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayName

`func (o *MicrosoftGraphManagedDevice) SetUserDisplayName(v string)`

SetUserDisplayName gets a reference to the given string and assigns it to the UserDisplayName field.

### SetUserDisplayNameExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetUserDisplayNameExplicitNull(b bool)`

SetUserDisplayNameExplicitNull (un)sets UserDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserDisplayName value is set to nil even if false is passed
### GetConfigurationManagerClientEnabledFeatures

`func (o *MicrosoftGraphManagedDevice) GetConfigurationManagerClientEnabledFeatures() AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures`

GetConfigurationManagerClientEnabledFeatures returns the ConfigurationManagerClientEnabledFeatures field if non-nil, zero value otherwise.

### GetConfigurationManagerClientEnabledFeaturesOk

`func (o *MicrosoftGraphManagedDevice) GetConfigurationManagerClientEnabledFeaturesOk() (AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures, bool)`

GetConfigurationManagerClientEnabledFeaturesOk returns a tuple with the ConfigurationManagerClientEnabledFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfigurationManagerClientEnabledFeatures

`func (o *MicrosoftGraphManagedDevice) HasConfigurationManagerClientEnabledFeatures() bool`

HasConfigurationManagerClientEnabledFeatures returns a boolean if a field has been set.

### SetConfigurationManagerClientEnabledFeatures

`func (o *MicrosoftGraphManagedDevice) SetConfigurationManagerClientEnabledFeatures(v AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures)`

SetConfigurationManagerClientEnabledFeatures gets a reference to the given AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures and assigns it to the ConfigurationManagerClientEnabledFeatures field.

### SetConfigurationManagerClientEnabledFeaturesExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetConfigurationManagerClientEnabledFeaturesExplicitNull(b bool)`

SetConfigurationManagerClientEnabledFeaturesExplicitNull (un)sets ConfigurationManagerClientEnabledFeatures to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConfigurationManagerClientEnabledFeatures value is set to nil even if false is passed
### GetWiFiMacAddress

`func (o *MicrosoftGraphManagedDevice) GetWiFiMacAddress() string`

GetWiFiMacAddress returns the WiFiMacAddress field if non-nil, zero value otherwise.

### GetWiFiMacAddressOk

`func (o *MicrosoftGraphManagedDevice) GetWiFiMacAddressOk() (string, bool)`

GetWiFiMacAddressOk returns a tuple with the WiFiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiMacAddress

`func (o *MicrosoftGraphManagedDevice) HasWiFiMacAddress() bool`

HasWiFiMacAddress returns a boolean if a field has been set.

### SetWiFiMacAddress

`func (o *MicrosoftGraphManagedDevice) SetWiFiMacAddress(v string)`

SetWiFiMacAddress gets a reference to the given string and assigns it to the WiFiMacAddress field.

### SetWiFiMacAddressExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetWiFiMacAddressExplicitNull(b bool)`

SetWiFiMacAddressExplicitNull (un)sets WiFiMacAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WiFiMacAddress value is set to nil even if false is passed
### GetDeviceHealthAttestationState

`func (o *MicrosoftGraphManagedDevice) GetDeviceHealthAttestationState() AnyOfmicrosoftGraphDeviceHealthAttestationState`

GetDeviceHealthAttestationState returns the DeviceHealthAttestationState field if non-nil, zero value otherwise.

### GetDeviceHealthAttestationStateOk

`func (o *MicrosoftGraphManagedDevice) GetDeviceHealthAttestationStateOk() (AnyOfmicrosoftGraphDeviceHealthAttestationState, bool)`

GetDeviceHealthAttestationStateOk returns a tuple with the DeviceHealthAttestationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceHealthAttestationState

`func (o *MicrosoftGraphManagedDevice) HasDeviceHealthAttestationState() bool`

HasDeviceHealthAttestationState returns a boolean if a field has been set.

### SetDeviceHealthAttestationState

`func (o *MicrosoftGraphManagedDevice) SetDeviceHealthAttestationState(v AnyOfmicrosoftGraphDeviceHealthAttestationState)`

SetDeviceHealthAttestationState gets a reference to the given AnyOfmicrosoftGraphDeviceHealthAttestationState and assigns it to the DeviceHealthAttestationState field.

### SetDeviceHealthAttestationStateExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetDeviceHealthAttestationStateExplicitNull(b bool)`

SetDeviceHealthAttestationStateExplicitNull (un)sets DeviceHealthAttestationState to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceHealthAttestationState value is set to nil even if false is passed
### GetSubscriberCarrier

`func (o *MicrosoftGraphManagedDevice) GetSubscriberCarrier() string`

GetSubscriberCarrier returns the SubscriberCarrier field if non-nil, zero value otherwise.

### GetSubscriberCarrierOk

`func (o *MicrosoftGraphManagedDevice) GetSubscriberCarrierOk() (string, bool)`

GetSubscriberCarrierOk returns a tuple with the SubscriberCarrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscriberCarrier

`func (o *MicrosoftGraphManagedDevice) HasSubscriberCarrier() bool`

HasSubscriberCarrier returns a boolean if a field has been set.

### SetSubscriberCarrier

`func (o *MicrosoftGraphManagedDevice) SetSubscriberCarrier(v string)`

SetSubscriberCarrier gets a reference to the given string and assigns it to the SubscriberCarrier field.

### SetSubscriberCarrierExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetSubscriberCarrierExplicitNull(b bool)`

SetSubscriberCarrierExplicitNull (un)sets SubscriberCarrier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SubscriberCarrier value is set to nil even if false is passed
### GetMeid

`func (o *MicrosoftGraphManagedDevice) GetMeid() string`

GetMeid returns the Meid field if non-nil, zero value otherwise.

### GetMeidOk

`func (o *MicrosoftGraphManagedDevice) GetMeidOk() (string, bool)`

GetMeidOk returns a tuple with the Meid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMeid

`func (o *MicrosoftGraphManagedDevice) HasMeid() bool`

HasMeid returns a boolean if a field has been set.

### SetMeid

`func (o *MicrosoftGraphManagedDevice) SetMeid(v string)`

SetMeid gets a reference to the given string and assigns it to the Meid field.

### SetMeidExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetMeidExplicitNull(b bool)`

SetMeidExplicitNull (un)sets Meid to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Meid value is set to nil even if false is passed
### GetTotalStorageSpaceInBytes

`func (o *MicrosoftGraphManagedDevice) GetTotalStorageSpaceInBytes() int64`

GetTotalStorageSpaceInBytes returns the TotalStorageSpaceInBytes field if non-nil, zero value otherwise.

### GetTotalStorageSpaceInBytesOk

`func (o *MicrosoftGraphManagedDevice) GetTotalStorageSpaceInBytesOk() (int64, bool)`

GetTotalStorageSpaceInBytesOk returns a tuple with the TotalStorageSpaceInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalStorageSpaceInBytes

`func (o *MicrosoftGraphManagedDevice) HasTotalStorageSpaceInBytes() bool`

HasTotalStorageSpaceInBytes returns a boolean if a field has been set.

### SetTotalStorageSpaceInBytes

`func (o *MicrosoftGraphManagedDevice) SetTotalStorageSpaceInBytes(v int64)`

SetTotalStorageSpaceInBytes gets a reference to the given int64 and assigns it to the TotalStorageSpaceInBytes field.

### GetFreeStorageSpaceInBytes

`func (o *MicrosoftGraphManagedDevice) GetFreeStorageSpaceInBytes() int64`

GetFreeStorageSpaceInBytes returns the FreeStorageSpaceInBytes field if non-nil, zero value otherwise.

### GetFreeStorageSpaceInBytesOk

`func (o *MicrosoftGraphManagedDevice) GetFreeStorageSpaceInBytesOk() (int64, bool)`

GetFreeStorageSpaceInBytesOk returns a tuple with the FreeStorageSpaceInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFreeStorageSpaceInBytes

`func (o *MicrosoftGraphManagedDevice) HasFreeStorageSpaceInBytes() bool`

HasFreeStorageSpaceInBytes returns a boolean if a field has been set.

### SetFreeStorageSpaceInBytes

`func (o *MicrosoftGraphManagedDevice) SetFreeStorageSpaceInBytes(v int64)`

SetFreeStorageSpaceInBytes gets a reference to the given int64 and assigns it to the FreeStorageSpaceInBytes field.

### GetManagedDeviceName

`func (o *MicrosoftGraphManagedDevice) GetManagedDeviceName() string`

GetManagedDeviceName returns the ManagedDeviceName field if non-nil, zero value otherwise.

### GetManagedDeviceNameOk

`func (o *MicrosoftGraphManagedDevice) GetManagedDeviceNameOk() (string, bool)`

GetManagedDeviceNameOk returns a tuple with the ManagedDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedDeviceName

`func (o *MicrosoftGraphManagedDevice) HasManagedDeviceName() bool`

HasManagedDeviceName returns a boolean if a field has been set.

### SetManagedDeviceName

`func (o *MicrosoftGraphManagedDevice) SetManagedDeviceName(v string)`

SetManagedDeviceName gets a reference to the given string and assigns it to the ManagedDeviceName field.

### SetManagedDeviceNameExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetManagedDeviceNameExplicitNull(b bool)`

SetManagedDeviceNameExplicitNull (un)sets ManagedDeviceName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ManagedDeviceName value is set to nil even if false is passed
### GetPartnerReportedThreatState

`func (o *MicrosoftGraphManagedDevice) GetPartnerReportedThreatState() AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState`

GetPartnerReportedThreatState returns the PartnerReportedThreatState field if non-nil, zero value otherwise.

### GetPartnerReportedThreatStateOk

`func (o *MicrosoftGraphManagedDevice) GetPartnerReportedThreatStateOk() (AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState, bool)`

GetPartnerReportedThreatStateOk returns a tuple with the PartnerReportedThreatState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPartnerReportedThreatState

`func (o *MicrosoftGraphManagedDevice) HasPartnerReportedThreatState() bool`

HasPartnerReportedThreatState returns a boolean if a field has been set.

### SetPartnerReportedThreatState

`func (o *MicrosoftGraphManagedDevice) SetPartnerReportedThreatState(v AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState)`

SetPartnerReportedThreatState gets a reference to the given AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState and assigns it to the PartnerReportedThreatState field.

### GetDeviceConfigurationStates

`func (o *MicrosoftGraphManagedDevice) GetDeviceConfigurationStates() []MicrosoftGraphDeviceConfigurationState`

GetDeviceConfigurationStates returns the DeviceConfigurationStates field if non-nil, zero value otherwise.

### GetDeviceConfigurationStatesOk

`func (o *MicrosoftGraphManagedDevice) GetDeviceConfigurationStatesOk() ([]MicrosoftGraphDeviceConfigurationState, bool)`

GetDeviceConfigurationStatesOk returns a tuple with the DeviceConfigurationStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceConfigurationStates

`func (o *MicrosoftGraphManagedDevice) HasDeviceConfigurationStates() bool`

HasDeviceConfigurationStates returns a boolean if a field has been set.

### SetDeviceConfigurationStates

`func (o *MicrosoftGraphManagedDevice) SetDeviceConfigurationStates(v []MicrosoftGraphDeviceConfigurationState)`

SetDeviceConfigurationStates gets a reference to the given []MicrosoftGraphDeviceConfigurationState and assigns it to the DeviceConfigurationStates field.

### GetDeviceCompliancePolicyStates

`func (o *MicrosoftGraphManagedDevice) GetDeviceCompliancePolicyStates() []MicrosoftGraphDeviceCompliancePolicyState`

GetDeviceCompliancePolicyStates returns the DeviceCompliancePolicyStates field if non-nil, zero value otherwise.

### GetDeviceCompliancePolicyStatesOk

`func (o *MicrosoftGraphManagedDevice) GetDeviceCompliancePolicyStatesOk() ([]MicrosoftGraphDeviceCompliancePolicyState, bool)`

GetDeviceCompliancePolicyStatesOk returns a tuple with the DeviceCompliancePolicyStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCompliancePolicyStates

`func (o *MicrosoftGraphManagedDevice) HasDeviceCompliancePolicyStates() bool`

HasDeviceCompliancePolicyStates returns a boolean if a field has been set.

### SetDeviceCompliancePolicyStates

`func (o *MicrosoftGraphManagedDevice) SetDeviceCompliancePolicyStates(v []MicrosoftGraphDeviceCompliancePolicyState)`

SetDeviceCompliancePolicyStates gets a reference to the given []MicrosoftGraphDeviceCompliancePolicyState and assigns it to the DeviceCompliancePolicyStates field.

### GetDeviceCategory

`func (o *MicrosoftGraphManagedDevice) GetDeviceCategory() AnyOfmicrosoftGraphDeviceCategory`

GetDeviceCategory returns the DeviceCategory field if non-nil, zero value otherwise.

### GetDeviceCategoryOk

`func (o *MicrosoftGraphManagedDevice) GetDeviceCategoryOk() (AnyOfmicrosoftGraphDeviceCategory, bool)`

GetDeviceCategoryOk returns a tuple with the DeviceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCategory

`func (o *MicrosoftGraphManagedDevice) HasDeviceCategory() bool`

HasDeviceCategory returns a boolean if a field has been set.

### SetDeviceCategory

`func (o *MicrosoftGraphManagedDevice) SetDeviceCategory(v AnyOfmicrosoftGraphDeviceCategory)`

SetDeviceCategory gets a reference to the given AnyOfmicrosoftGraphDeviceCategory and assigns it to the DeviceCategory field.

### SetDeviceCategoryExplicitNull

`func (o *MicrosoftGraphManagedDevice) SetDeviceCategoryExplicitNull(b bool)`

SetDeviceCategoryExplicitNull (un)sets DeviceCategory to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceCategory value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


