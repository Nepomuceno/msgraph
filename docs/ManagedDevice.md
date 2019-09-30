# ManagedDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetUserId

`func (o *ManagedDevice) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ManagedDevice) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *ManagedDevice) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *ManagedDevice) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### SetUserIdExplicitNull

`func (o *ManagedDevice) SetUserIdExplicitNull(b bool)`

SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserId value is set to nil even if false is passed
### GetDeviceName

`func (o *ManagedDevice) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ManagedDevice) GetDeviceNameOk() (string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceName

`func (o *ManagedDevice) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceName

`func (o *ManagedDevice) SetDeviceName(v string)`

SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.

### SetDeviceNameExplicitNull

`func (o *ManagedDevice) SetDeviceNameExplicitNull(b bool)`

SetDeviceNameExplicitNull (un)sets DeviceName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceName value is set to nil even if false is passed
### GetManagedDeviceOwnerType

`func (o *ManagedDevice) GetManagedDeviceOwnerType() AnyOfmicrosoftGraphManagedDeviceOwnerType`

GetManagedDeviceOwnerType returns the ManagedDeviceOwnerType field if non-nil, zero value otherwise.

### GetManagedDeviceOwnerTypeOk

`func (o *ManagedDevice) GetManagedDeviceOwnerTypeOk() (AnyOfmicrosoftGraphManagedDeviceOwnerType, bool)`

GetManagedDeviceOwnerTypeOk returns a tuple with the ManagedDeviceOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedDeviceOwnerType

`func (o *ManagedDevice) HasManagedDeviceOwnerType() bool`

HasManagedDeviceOwnerType returns a boolean if a field has been set.

### SetManagedDeviceOwnerType

`func (o *ManagedDevice) SetManagedDeviceOwnerType(v AnyOfmicrosoftGraphManagedDeviceOwnerType)`

SetManagedDeviceOwnerType gets a reference to the given AnyOfmicrosoftGraphManagedDeviceOwnerType and assigns it to the ManagedDeviceOwnerType field.

### GetDeviceActionResults

`func (o *ManagedDevice) GetDeviceActionResults() []AnyOfmicrosoftGraphDeviceActionResult`

GetDeviceActionResults returns the DeviceActionResults field if non-nil, zero value otherwise.

### GetDeviceActionResultsOk

`func (o *ManagedDevice) GetDeviceActionResultsOk() ([]AnyOfmicrosoftGraphDeviceActionResult, bool)`

GetDeviceActionResultsOk returns a tuple with the DeviceActionResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceActionResults

`func (o *ManagedDevice) HasDeviceActionResults() bool`

HasDeviceActionResults returns a boolean if a field has been set.

### SetDeviceActionResults

`func (o *ManagedDevice) SetDeviceActionResults(v []AnyOfmicrosoftGraphDeviceActionResult)`

SetDeviceActionResults gets a reference to the given []AnyOfmicrosoftGraphDeviceActionResult and assigns it to the DeviceActionResults field.

### GetEnrolledDateTime

`func (o *ManagedDevice) GetEnrolledDateTime() time.Time`

GetEnrolledDateTime returns the EnrolledDateTime field if non-nil, zero value otherwise.

### GetEnrolledDateTimeOk

`func (o *ManagedDevice) GetEnrolledDateTimeOk() (time.Time, bool)`

GetEnrolledDateTimeOk returns a tuple with the EnrolledDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnrolledDateTime

`func (o *ManagedDevice) HasEnrolledDateTime() bool`

HasEnrolledDateTime returns a boolean if a field has been set.

### SetEnrolledDateTime

`func (o *ManagedDevice) SetEnrolledDateTime(v time.Time)`

SetEnrolledDateTime gets a reference to the given time.Time and assigns it to the EnrolledDateTime field.

### GetLastSyncDateTime

`func (o *ManagedDevice) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *ManagedDevice) GetLastSyncDateTimeOk() (time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSyncDateTime

`func (o *ManagedDevice) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### SetLastSyncDateTime

`func (o *ManagedDevice) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.

### GetOperatingSystem

`func (o *ManagedDevice) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ManagedDevice) GetOperatingSystemOk() (string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperatingSystem

`func (o *ManagedDevice) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystem

`func (o *ManagedDevice) SetOperatingSystem(v string)`

SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.

### SetOperatingSystemExplicitNull

`func (o *ManagedDevice) SetOperatingSystemExplicitNull(b bool)`

SetOperatingSystemExplicitNull (un)sets OperatingSystem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperatingSystem value is set to nil even if false is passed
### GetComplianceState

`func (o *ManagedDevice) GetComplianceState() AnyOfmicrosoftGraphComplianceState`

GetComplianceState returns the ComplianceState field if non-nil, zero value otherwise.

### GetComplianceStateOk

`func (o *ManagedDevice) GetComplianceStateOk() (AnyOfmicrosoftGraphComplianceState, bool)`

GetComplianceStateOk returns a tuple with the ComplianceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplianceState

`func (o *ManagedDevice) HasComplianceState() bool`

HasComplianceState returns a boolean if a field has been set.

### SetComplianceState

`func (o *ManagedDevice) SetComplianceState(v AnyOfmicrosoftGraphComplianceState)`

SetComplianceState gets a reference to the given AnyOfmicrosoftGraphComplianceState and assigns it to the ComplianceState field.

### GetJailBroken

`func (o *ManagedDevice) GetJailBroken() string`

GetJailBroken returns the JailBroken field if non-nil, zero value otherwise.

### GetJailBrokenOk

`func (o *ManagedDevice) GetJailBrokenOk() (string, bool)`

GetJailBrokenOk returns a tuple with the JailBroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJailBroken

`func (o *ManagedDevice) HasJailBroken() bool`

HasJailBroken returns a boolean if a field has been set.

### SetJailBroken

`func (o *ManagedDevice) SetJailBroken(v string)`

SetJailBroken gets a reference to the given string and assigns it to the JailBroken field.

### SetJailBrokenExplicitNull

`func (o *ManagedDevice) SetJailBrokenExplicitNull(b bool)`

SetJailBrokenExplicitNull (un)sets JailBroken to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The JailBroken value is set to nil even if false is passed
### GetManagementAgent

`func (o *ManagedDevice) GetManagementAgent() AnyOfmicrosoftGraphManagementAgentType`

GetManagementAgent returns the ManagementAgent field if non-nil, zero value otherwise.

### GetManagementAgentOk

`func (o *ManagedDevice) GetManagementAgentOk() (AnyOfmicrosoftGraphManagementAgentType, bool)`

GetManagementAgentOk returns a tuple with the ManagementAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagementAgent

`func (o *ManagedDevice) HasManagementAgent() bool`

HasManagementAgent returns a boolean if a field has been set.

### SetManagementAgent

`func (o *ManagedDevice) SetManagementAgent(v AnyOfmicrosoftGraphManagementAgentType)`

SetManagementAgent gets a reference to the given AnyOfmicrosoftGraphManagementAgentType and assigns it to the ManagementAgent field.

### GetOsVersion

`func (o *ManagedDevice) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ManagedDevice) GetOsVersionOk() (string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsVersion

`func (o *ManagedDevice) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersion

`func (o *ManagedDevice) SetOsVersion(v string)`

SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.

### SetOsVersionExplicitNull

`func (o *ManagedDevice) SetOsVersionExplicitNull(b bool)`

SetOsVersionExplicitNull (un)sets OsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsVersion value is set to nil even if false is passed
### GetEasActivated

`func (o *ManagedDevice) GetEasActivated() bool`

GetEasActivated returns the EasActivated field if non-nil, zero value otherwise.

### GetEasActivatedOk

`func (o *ManagedDevice) GetEasActivatedOk() (bool, bool)`

GetEasActivatedOk returns a tuple with the EasActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEasActivated

`func (o *ManagedDevice) HasEasActivated() bool`

HasEasActivated returns a boolean if a field has been set.

### SetEasActivated

`func (o *ManagedDevice) SetEasActivated(v bool)`

SetEasActivated gets a reference to the given bool and assigns it to the EasActivated field.

### GetEasDeviceId

`func (o *ManagedDevice) GetEasDeviceId() string`

GetEasDeviceId returns the EasDeviceId field if non-nil, zero value otherwise.

### GetEasDeviceIdOk

`func (o *ManagedDevice) GetEasDeviceIdOk() (string, bool)`

GetEasDeviceIdOk returns a tuple with the EasDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEasDeviceId

`func (o *ManagedDevice) HasEasDeviceId() bool`

HasEasDeviceId returns a boolean if a field has been set.

### SetEasDeviceId

`func (o *ManagedDevice) SetEasDeviceId(v string)`

SetEasDeviceId gets a reference to the given string and assigns it to the EasDeviceId field.

### SetEasDeviceIdExplicitNull

`func (o *ManagedDevice) SetEasDeviceIdExplicitNull(b bool)`

SetEasDeviceIdExplicitNull (un)sets EasDeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EasDeviceId value is set to nil even if false is passed
### GetEasActivationDateTime

`func (o *ManagedDevice) GetEasActivationDateTime() time.Time`

GetEasActivationDateTime returns the EasActivationDateTime field if non-nil, zero value otherwise.

### GetEasActivationDateTimeOk

`func (o *ManagedDevice) GetEasActivationDateTimeOk() (time.Time, bool)`

GetEasActivationDateTimeOk returns a tuple with the EasActivationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEasActivationDateTime

`func (o *ManagedDevice) HasEasActivationDateTime() bool`

HasEasActivationDateTime returns a boolean if a field has been set.

### SetEasActivationDateTime

`func (o *ManagedDevice) SetEasActivationDateTime(v time.Time)`

SetEasActivationDateTime gets a reference to the given time.Time and assigns it to the EasActivationDateTime field.

### GetAzureADRegistered

`func (o *ManagedDevice) GetAzureADRegistered() bool`

GetAzureADRegistered returns the AzureADRegistered field if non-nil, zero value otherwise.

### GetAzureADRegisteredOk

`func (o *ManagedDevice) GetAzureADRegisteredOk() (bool, bool)`

GetAzureADRegisteredOk returns a tuple with the AzureADRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureADRegistered

`func (o *ManagedDevice) HasAzureADRegistered() bool`

HasAzureADRegistered returns a boolean if a field has been set.

### SetAzureADRegistered

`func (o *ManagedDevice) SetAzureADRegistered(v bool)`

SetAzureADRegistered gets a reference to the given bool and assigns it to the AzureADRegistered field.

### SetAzureADRegisteredExplicitNull

`func (o *ManagedDevice) SetAzureADRegisteredExplicitNull(b bool)`

SetAzureADRegisteredExplicitNull (un)sets AzureADRegistered to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AzureADRegistered value is set to nil even if false is passed
### GetDeviceEnrollmentType

`func (o *ManagedDevice) GetDeviceEnrollmentType() AnyOfmicrosoftGraphDeviceEnrollmentType`

GetDeviceEnrollmentType returns the DeviceEnrollmentType field if non-nil, zero value otherwise.

### GetDeviceEnrollmentTypeOk

`func (o *ManagedDevice) GetDeviceEnrollmentTypeOk() (AnyOfmicrosoftGraphDeviceEnrollmentType, bool)`

GetDeviceEnrollmentTypeOk returns a tuple with the DeviceEnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceEnrollmentType

`func (o *ManagedDevice) HasDeviceEnrollmentType() bool`

HasDeviceEnrollmentType returns a boolean if a field has been set.

### SetDeviceEnrollmentType

`func (o *ManagedDevice) SetDeviceEnrollmentType(v AnyOfmicrosoftGraphDeviceEnrollmentType)`

SetDeviceEnrollmentType gets a reference to the given AnyOfmicrosoftGraphDeviceEnrollmentType and assigns it to the DeviceEnrollmentType field.

### GetActivationLockBypassCode

`func (o *ManagedDevice) GetActivationLockBypassCode() string`

GetActivationLockBypassCode returns the ActivationLockBypassCode field if non-nil, zero value otherwise.

### GetActivationLockBypassCodeOk

`func (o *ManagedDevice) GetActivationLockBypassCodeOk() (string, bool)`

GetActivationLockBypassCodeOk returns a tuple with the ActivationLockBypassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivationLockBypassCode

`func (o *ManagedDevice) HasActivationLockBypassCode() bool`

HasActivationLockBypassCode returns a boolean if a field has been set.

### SetActivationLockBypassCode

`func (o *ManagedDevice) SetActivationLockBypassCode(v string)`

SetActivationLockBypassCode gets a reference to the given string and assigns it to the ActivationLockBypassCode field.

### SetActivationLockBypassCodeExplicitNull

`func (o *ManagedDevice) SetActivationLockBypassCodeExplicitNull(b bool)`

SetActivationLockBypassCodeExplicitNull (un)sets ActivationLockBypassCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActivationLockBypassCode value is set to nil even if false is passed
### GetEmailAddress

`func (o *ManagedDevice) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ManagedDevice) GetEmailAddressOk() (string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailAddress

`func (o *ManagedDevice) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddress

`func (o *ManagedDevice) SetEmailAddress(v string)`

SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.

### SetEmailAddressExplicitNull

`func (o *ManagedDevice) SetEmailAddressExplicitNull(b bool)`

SetEmailAddressExplicitNull (un)sets EmailAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EmailAddress value is set to nil even if false is passed
### GetAzureADDeviceId

`func (o *ManagedDevice) GetAzureADDeviceId() string`

GetAzureADDeviceId returns the AzureADDeviceId field if non-nil, zero value otherwise.

### GetAzureADDeviceIdOk

`func (o *ManagedDevice) GetAzureADDeviceIdOk() (string, bool)`

GetAzureADDeviceIdOk returns a tuple with the AzureADDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureADDeviceId

`func (o *ManagedDevice) HasAzureADDeviceId() bool`

HasAzureADDeviceId returns a boolean if a field has been set.

### SetAzureADDeviceId

`func (o *ManagedDevice) SetAzureADDeviceId(v string)`

SetAzureADDeviceId gets a reference to the given string and assigns it to the AzureADDeviceId field.

### SetAzureADDeviceIdExplicitNull

`func (o *ManagedDevice) SetAzureADDeviceIdExplicitNull(b bool)`

SetAzureADDeviceIdExplicitNull (un)sets AzureADDeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AzureADDeviceId value is set to nil even if false is passed
### GetDeviceRegistrationState

`func (o *ManagedDevice) GetDeviceRegistrationState() AnyOfmicrosoftGraphDeviceRegistrationState`

GetDeviceRegistrationState returns the DeviceRegistrationState field if non-nil, zero value otherwise.

### GetDeviceRegistrationStateOk

`func (o *ManagedDevice) GetDeviceRegistrationStateOk() (AnyOfmicrosoftGraphDeviceRegistrationState, bool)`

GetDeviceRegistrationStateOk returns a tuple with the DeviceRegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceRegistrationState

`func (o *ManagedDevice) HasDeviceRegistrationState() bool`

HasDeviceRegistrationState returns a boolean if a field has been set.

### SetDeviceRegistrationState

`func (o *ManagedDevice) SetDeviceRegistrationState(v AnyOfmicrosoftGraphDeviceRegistrationState)`

SetDeviceRegistrationState gets a reference to the given AnyOfmicrosoftGraphDeviceRegistrationState and assigns it to the DeviceRegistrationState field.

### GetDeviceCategoryDisplayName

`func (o *ManagedDevice) GetDeviceCategoryDisplayName() string`

GetDeviceCategoryDisplayName returns the DeviceCategoryDisplayName field if non-nil, zero value otherwise.

### GetDeviceCategoryDisplayNameOk

`func (o *ManagedDevice) GetDeviceCategoryDisplayNameOk() (string, bool)`

GetDeviceCategoryDisplayNameOk returns a tuple with the DeviceCategoryDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCategoryDisplayName

`func (o *ManagedDevice) HasDeviceCategoryDisplayName() bool`

HasDeviceCategoryDisplayName returns a boolean if a field has been set.

### SetDeviceCategoryDisplayName

`func (o *ManagedDevice) SetDeviceCategoryDisplayName(v string)`

SetDeviceCategoryDisplayName gets a reference to the given string and assigns it to the DeviceCategoryDisplayName field.

### SetDeviceCategoryDisplayNameExplicitNull

`func (o *ManagedDevice) SetDeviceCategoryDisplayNameExplicitNull(b bool)`

SetDeviceCategoryDisplayNameExplicitNull (un)sets DeviceCategoryDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceCategoryDisplayName value is set to nil even if false is passed
### GetIsSupervised

`func (o *ManagedDevice) GetIsSupervised() bool`

GetIsSupervised returns the IsSupervised field if non-nil, zero value otherwise.

### GetIsSupervisedOk

`func (o *ManagedDevice) GetIsSupervisedOk() (bool, bool)`

GetIsSupervisedOk returns a tuple with the IsSupervised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsSupervised

`func (o *ManagedDevice) HasIsSupervised() bool`

HasIsSupervised returns a boolean if a field has been set.

### SetIsSupervised

`func (o *ManagedDevice) SetIsSupervised(v bool)`

SetIsSupervised gets a reference to the given bool and assigns it to the IsSupervised field.

### GetExchangeLastSuccessfulSyncDateTime

`func (o *ManagedDevice) GetExchangeLastSuccessfulSyncDateTime() time.Time`

GetExchangeLastSuccessfulSyncDateTime returns the ExchangeLastSuccessfulSyncDateTime field if non-nil, zero value otherwise.

### GetExchangeLastSuccessfulSyncDateTimeOk

`func (o *ManagedDevice) GetExchangeLastSuccessfulSyncDateTimeOk() (time.Time, bool)`

GetExchangeLastSuccessfulSyncDateTimeOk returns a tuple with the ExchangeLastSuccessfulSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeLastSuccessfulSyncDateTime

`func (o *ManagedDevice) HasExchangeLastSuccessfulSyncDateTime() bool`

HasExchangeLastSuccessfulSyncDateTime returns a boolean if a field has been set.

### SetExchangeLastSuccessfulSyncDateTime

`func (o *ManagedDevice) SetExchangeLastSuccessfulSyncDateTime(v time.Time)`

SetExchangeLastSuccessfulSyncDateTime gets a reference to the given time.Time and assigns it to the ExchangeLastSuccessfulSyncDateTime field.

### GetExchangeAccessState

`func (o *ManagedDevice) GetExchangeAccessState() AnyOfmicrosoftGraphDeviceManagementExchangeAccessState`

GetExchangeAccessState returns the ExchangeAccessState field if non-nil, zero value otherwise.

### GetExchangeAccessStateOk

`func (o *ManagedDevice) GetExchangeAccessStateOk() (AnyOfmicrosoftGraphDeviceManagementExchangeAccessState, bool)`

GetExchangeAccessStateOk returns a tuple with the ExchangeAccessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeAccessState

`func (o *ManagedDevice) HasExchangeAccessState() bool`

HasExchangeAccessState returns a boolean if a field has been set.

### SetExchangeAccessState

`func (o *ManagedDevice) SetExchangeAccessState(v AnyOfmicrosoftGraphDeviceManagementExchangeAccessState)`

SetExchangeAccessState gets a reference to the given AnyOfmicrosoftGraphDeviceManagementExchangeAccessState and assigns it to the ExchangeAccessState field.

### GetExchangeAccessStateReason

`func (o *ManagedDevice) GetExchangeAccessStateReason() AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason`

GetExchangeAccessStateReason returns the ExchangeAccessStateReason field if non-nil, zero value otherwise.

### GetExchangeAccessStateReasonOk

`func (o *ManagedDevice) GetExchangeAccessStateReasonOk() (AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason, bool)`

GetExchangeAccessStateReasonOk returns a tuple with the ExchangeAccessStateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeAccessStateReason

`func (o *ManagedDevice) HasExchangeAccessStateReason() bool`

HasExchangeAccessStateReason returns a boolean if a field has been set.

### SetExchangeAccessStateReason

`func (o *ManagedDevice) SetExchangeAccessStateReason(v AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason)`

SetExchangeAccessStateReason gets a reference to the given AnyOfmicrosoftGraphDeviceManagementExchangeAccessStateReason and assigns it to the ExchangeAccessStateReason field.

### GetRemoteAssistanceSessionUrl

`func (o *ManagedDevice) GetRemoteAssistanceSessionUrl() string`

GetRemoteAssistanceSessionUrl returns the RemoteAssistanceSessionUrl field if non-nil, zero value otherwise.

### GetRemoteAssistanceSessionUrlOk

`func (o *ManagedDevice) GetRemoteAssistanceSessionUrlOk() (string, bool)`

GetRemoteAssistanceSessionUrlOk returns a tuple with the RemoteAssistanceSessionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemoteAssistanceSessionUrl

`func (o *ManagedDevice) HasRemoteAssistanceSessionUrl() bool`

HasRemoteAssistanceSessionUrl returns a boolean if a field has been set.

### SetRemoteAssistanceSessionUrl

`func (o *ManagedDevice) SetRemoteAssistanceSessionUrl(v string)`

SetRemoteAssistanceSessionUrl gets a reference to the given string and assigns it to the RemoteAssistanceSessionUrl field.

### SetRemoteAssistanceSessionUrlExplicitNull

`func (o *ManagedDevice) SetRemoteAssistanceSessionUrlExplicitNull(b bool)`

SetRemoteAssistanceSessionUrlExplicitNull (un)sets RemoteAssistanceSessionUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RemoteAssistanceSessionUrl value is set to nil even if false is passed
### GetRemoteAssistanceSessionErrorDetails

`func (o *ManagedDevice) GetRemoteAssistanceSessionErrorDetails() string`

GetRemoteAssistanceSessionErrorDetails returns the RemoteAssistanceSessionErrorDetails field if non-nil, zero value otherwise.

### GetRemoteAssistanceSessionErrorDetailsOk

`func (o *ManagedDevice) GetRemoteAssistanceSessionErrorDetailsOk() (string, bool)`

GetRemoteAssistanceSessionErrorDetailsOk returns a tuple with the RemoteAssistanceSessionErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemoteAssistanceSessionErrorDetails

`func (o *ManagedDevice) HasRemoteAssistanceSessionErrorDetails() bool`

HasRemoteAssistanceSessionErrorDetails returns a boolean if a field has been set.

### SetRemoteAssistanceSessionErrorDetails

`func (o *ManagedDevice) SetRemoteAssistanceSessionErrorDetails(v string)`

SetRemoteAssistanceSessionErrorDetails gets a reference to the given string and assigns it to the RemoteAssistanceSessionErrorDetails field.

### SetRemoteAssistanceSessionErrorDetailsExplicitNull

`func (o *ManagedDevice) SetRemoteAssistanceSessionErrorDetailsExplicitNull(b bool)`

SetRemoteAssistanceSessionErrorDetailsExplicitNull (un)sets RemoteAssistanceSessionErrorDetails to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RemoteAssistanceSessionErrorDetails value is set to nil even if false is passed
### GetIsEncrypted

`func (o *ManagedDevice) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *ManagedDevice) GetIsEncryptedOk() (bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEncrypted

`func (o *ManagedDevice) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### SetIsEncrypted

`func (o *ManagedDevice) SetIsEncrypted(v bool)`

SetIsEncrypted gets a reference to the given bool and assigns it to the IsEncrypted field.

### GetUserPrincipalName

`func (o *ManagedDevice) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *ManagedDevice) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *ManagedDevice) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *ManagedDevice) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *ManagedDevice) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetModel

`func (o *ManagedDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ManagedDevice) GetModelOk() (string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModel

`func (o *ManagedDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModel

`func (o *ManagedDevice) SetModel(v string)`

SetModel gets a reference to the given string and assigns it to the Model field.

### SetModelExplicitNull

`func (o *ManagedDevice) SetModelExplicitNull(b bool)`

SetModelExplicitNull (un)sets Model to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Model value is set to nil even if false is passed
### GetManufacturer

`func (o *ManagedDevice) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ManagedDevice) GetManufacturerOk() (string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManufacturer

`func (o *ManagedDevice) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturer

`func (o *ManagedDevice) SetManufacturer(v string)`

SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.

### SetManufacturerExplicitNull

`func (o *ManagedDevice) SetManufacturerExplicitNull(b bool)`

SetManufacturerExplicitNull (un)sets Manufacturer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Manufacturer value is set to nil even if false is passed
### GetImei

`func (o *ManagedDevice) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *ManagedDevice) GetImeiOk() (string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImei

`func (o *ManagedDevice) HasImei() bool`

HasImei returns a boolean if a field has been set.

### SetImei

`func (o *ManagedDevice) SetImei(v string)`

SetImei gets a reference to the given string and assigns it to the Imei field.

### SetImeiExplicitNull

`func (o *ManagedDevice) SetImeiExplicitNull(b bool)`

SetImeiExplicitNull (un)sets Imei to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Imei value is set to nil even if false is passed
### GetComplianceGracePeriodExpirationDateTime

`func (o *ManagedDevice) GetComplianceGracePeriodExpirationDateTime() time.Time`

GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceGracePeriodExpirationDateTimeOk

`func (o *ManagedDevice) GetComplianceGracePeriodExpirationDateTimeOk() (time.Time, bool)`

GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplianceGracePeriodExpirationDateTime

`func (o *ManagedDevice) HasComplianceGracePeriodExpirationDateTime() bool`

HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.

### SetComplianceGracePeriodExpirationDateTime

`func (o *ManagedDevice) SetComplianceGracePeriodExpirationDateTime(v time.Time)`

SetComplianceGracePeriodExpirationDateTime gets a reference to the given time.Time and assigns it to the ComplianceGracePeriodExpirationDateTime field.

### GetSerialNumber

`func (o *ManagedDevice) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ManagedDevice) GetSerialNumberOk() (string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSerialNumber

`func (o *ManagedDevice) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumber

`func (o *ManagedDevice) SetSerialNumber(v string)`

SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.

### SetSerialNumberExplicitNull

`func (o *ManagedDevice) SetSerialNumberExplicitNull(b bool)`

SetSerialNumberExplicitNull (un)sets SerialNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SerialNumber value is set to nil even if false is passed
### GetPhoneNumber

`func (o *ManagedDevice) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ManagedDevice) GetPhoneNumberOk() (string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhoneNumber

`func (o *ManagedDevice) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumber

`func (o *ManagedDevice) SetPhoneNumber(v string)`

SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.

### SetPhoneNumberExplicitNull

`func (o *ManagedDevice) SetPhoneNumberExplicitNull(b bool)`

SetPhoneNumberExplicitNull (un)sets PhoneNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PhoneNumber value is set to nil even if false is passed
### GetAndroidSecurityPatchLevel

`func (o *ManagedDevice) GetAndroidSecurityPatchLevel() string`

GetAndroidSecurityPatchLevel returns the AndroidSecurityPatchLevel field if non-nil, zero value otherwise.

### GetAndroidSecurityPatchLevelOk

`func (o *ManagedDevice) GetAndroidSecurityPatchLevelOk() (string, bool)`

GetAndroidSecurityPatchLevelOk returns a tuple with the AndroidSecurityPatchLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAndroidSecurityPatchLevel

`func (o *ManagedDevice) HasAndroidSecurityPatchLevel() bool`

HasAndroidSecurityPatchLevel returns a boolean if a field has been set.

### SetAndroidSecurityPatchLevel

`func (o *ManagedDevice) SetAndroidSecurityPatchLevel(v string)`

SetAndroidSecurityPatchLevel gets a reference to the given string and assigns it to the AndroidSecurityPatchLevel field.

### SetAndroidSecurityPatchLevelExplicitNull

`func (o *ManagedDevice) SetAndroidSecurityPatchLevelExplicitNull(b bool)`

SetAndroidSecurityPatchLevelExplicitNull (un)sets AndroidSecurityPatchLevel to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AndroidSecurityPatchLevel value is set to nil even if false is passed
### GetUserDisplayName

`func (o *ManagedDevice) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *ManagedDevice) GetUserDisplayNameOk() (string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserDisplayName

`func (o *ManagedDevice) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayName

`func (o *ManagedDevice) SetUserDisplayName(v string)`

SetUserDisplayName gets a reference to the given string and assigns it to the UserDisplayName field.

### SetUserDisplayNameExplicitNull

`func (o *ManagedDevice) SetUserDisplayNameExplicitNull(b bool)`

SetUserDisplayNameExplicitNull (un)sets UserDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserDisplayName value is set to nil even if false is passed
### GetConfigurationManagerClientEnabledFeatures

`func (o *ManagedDevice) GetConfigurationManagerClientEnabledFeatures() AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures`

GetConfigurationManagerClientEnabledFeatures returns the ConfigurationManagerClientEnabledFeatures field if non-nil, zero value otherwise.

### GetConfigurationManagerClientEnabledFeaturesOk

`func (o *ManagedDevice) GetConfigurationManagerClientEnabledFeaturesOk() (AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures, bool)`

GetConfigurationManagerClientEnabledFeaturesOk returns a tuple with the ConfigurationManagerClientEnabledFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfigurationManagerClientEnabledFeatures

`func (o *ManagedDevice) HasConfigurationManagerClientEnabledFeatures() bool`

HasConfigurationManagerClientEnabledFeatures returns a boolean if a field has been set.

### SetConfigurationManagerClientEnabledFeatures

`func (o *ManagedDevice) SetConfigurationManagerClientEnabledFeatures(v AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures)`

SetConfigurationManagerClientEnabledFeatures gets a reference to the given AnyOfmicrosoftGraphConfigurationManagerClientEnabledFeatures and assigns it to the ConfigurationManagerClientEnabledFeatures field.

### SetConfigurationManagerClientEnabledFeaturesExplicitNull

`func (o *ManagedDevice) SetConfigurationManagerClientEnabledFeaturesExplicitNull(b bool)`

SetConfigurationManagerClientEnabledFeaturesExplicitNull (un)sets ConfigurationManagerClientEnabledFeatures to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConfigurationManagerClientEnabledFeatures value is set to nil even if false is passed
### GetWiFiMacAddress

`func (o *ManagedDevice) GetWiFiMacAddress() string`

GetWiFiMacAddress returns the WiFiMacAddress field if non-nil, zero value otherwise.

### GetWiFiMacAddressOk

`func (o *ManagedDevice) GetWiFiMacAddressOk() (string, bool)`

GetWiFiMacAddressOk returns a tuple with the WiFiMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWiFiMacAddress

`func (o *ManagedDevice) HasWiFiMacAddress() bool`

HasWiFiMacAddress returns a boolean if a field has been set.

### SetWiFiMacAddress

`func (o *ManagedDevice) SetWiFiMacAddress(v string)`

SetWiFiMacAddress gets a reference to the given string and assigns it to the WiFiMacAddress field.

### SetWiFiMacAddressExplicitNull

`func (o *ManagedDevice) SetWiFiMacAddressExplicitNull(b bool)`

SetWiFiMacAddressExplicitNull (un)sets WiFiMacAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WiFiMacAddress value is set to nil even if false is passed
### GetDeviceHealthAttestationState

`func (o *ManagedDevice) GetDeviceHealthAttestationState() AnyOfmicrosoftGraphDeviceHealthAttestationState`

GetDeviceHealthAttestationState returns the DeviceHealthAttestationState field if non-nil, zero value otherwise.

### GetDeviceHealthAttestationStateOk

`func (o *ManagedDevice) GetDeviceHealthAttestationStateOk() (AnyOfmicrosoftGraphDeviceHealthAttestationState, bool)`

GetDeviceHealthAttestationStateOk returns a tuple with the DeviceHealthAttestationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceHealthAttestationState

`func (o *ManagedDevice) HasDeviceHealthAttestationState() bool`

HasDeviceHealthAttestationState returns a boolean if a field has been set.

### SetDeviceHealthAttestationState

`func (o *ManagedDevice) SetDeviceHealthAttestationState(v AnyOfmicrosoftGraphDeviceHealthAttestationState)`

SetDeviceHealthAttestationState gets a reference to the given AnyOfmicrosoftGraphDeviceHealthAttestationState and assigns it to the DeviceHealthAttestationState field.

### SetDeviceHealthAttestationStateExplicitNull

`func (o *ManagedDevice) SetDeviceHealthAttestationStateExplicitNull(b bool)`

SetDeviceHealthAttestationStateExplicitNull (un)sets DeviceHealthAttestationState to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceHealthAttestationState value is set to nil even if false is passed
### GetSubscriberCarrier

`func (o *ManagedDevice) GetSubscriberCarrier() string`

GetSubscriberCarrier returns the SubscriberCarrier field if non-nil, zero value otherwise.

### GetSubscriberCarrierOk

`func (o *ManagedDevice) GetSubscriberCarrierOk() (string, bool)`

GetSubscriberCarrierOk returns a tuple with the SubscriberCarrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscriberCarrier

`func (o *ManagedDevice) HasSubscriberCarrier() bool`

HasSubscriberCarrier returns a boolean if a field has been set.

### SetSubscriberCarrier

`func (o *ManagedDevice) SetSubscriberCarrier(v string)`

SetSubscriberCarrier gets a reference to the given string and assigns it to the SubscriberCarrier field.

### SetSubscriberCarrierExplicitNull

`func (o *ManagedDevice) SetSubscriberCarrierExplicitNull(b bool)`

SetSubscriberCarrierExplicitNull (un)sets SubscriberCarrier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SubscriberCarrier value is set to nil even if false is passed
### GetMeid

`func (o *ManagedDevice) GetMeid() string`

GetMeid returns the Meid field if non-nil, zero value otherwise.

### GetMeidOk

`func (o *ManagedDevice) GetMeidOk() (string, bool)`

GetMeidOk returns a tuple with the Meid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMeid

`func (o *ManagedDevice) HasMeid() bool`

HasMeid returns a boolean if a field has been set.

### SetMeid

`func (o *ManagedDevice) SetMeid(v string)`

SetMeid gets a reference to the given string and assigns it to the Meid field.

### SetMeidExplicitNull

`func (o *ManagedDevice) SetMeidExplicitNull(b bool)`

SetMeidExplicitNull (un)sets Meid to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Meid value is set to nil even if false is passed
### GetTotalStorageSpaceInBytes

`func (o *ManagedDevice) GetTotalStorageSpaceInBytes() int64`

GetTotalStorageSpaceInBytes returns the TotalStorageSpaceInBytes field if non-nil, zero value otherwise.

### GetTotalStorageSpaceInBytesOk

`func (o *ManagedDevice) GetTotalStorageSpaceInBytesOk() (int64, bool)`

GetTotalStorageSpaceInBytesOk returns a tuple with the TotalStorageSpaceInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalStorageSpaceInBytes

`func (o *ManagedDevice) HasTotalStorageSpaceInBytes() bool`

HasTotalStorageSpaceInBytes returns a boolean if a field has been set.

### SetTotalStorageSpaceInBytes

`func (o *ManagedDevice) SetTotalStorageSpaceInBytes(v int64)`

SetTotalStorageSpaceInBytes gets a reference to the given int64 and assigns it to the TotalStorageSpaceInBytes field.

### GetFreeStorageSpaceInBytes

`func (o *ManagedDevice) GetFreeStorageSpaceInBytes() int64`

GetFreeStorageSpaceInBytes returns the FreeStorageSpaceInBytes field if non-nil, zero value otherwise.

### GetFreeStorageSpaceInBytesOk

`func (o *ManagedDevice) GetFreeStorageSpaceInBytesOk() (int64, bool)`

GetFreeStorageSpaceInBytesOk returns a tuple with the FreeStorageSpaceInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFreeStorageSpaceInBytes

`func (o *ManagedDevice) HasFreeStorageSpaceInBytes() bool`

HasFreeStorageSpaceInBytes returns a boolean if a field has been set.

### SetFreeStorageSpaceInBytes

`func (o *ManagedDevice) SetFreeStorageSpaceInBytes(v int64)`

SetFreeStorageSpaceInBytes gets a reference to the given int64 and assigns it to the FreeStorageSpaceInBytes field.

### GetManagedDeviceName

`func (o *ManagedDevice) GetManagedDeviceName() string`

GetManagedDeviceName returns the ManagedDeviceName field if non-nil, zero value otherwise.

### GetManagedDeviceNameOk

`func (o *ManagedDevice) GetManagedDeviceNameOk() (string, bool)`

GetManagedDeviceNameOk returns a tuple with the ManagedDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedDeviceName

`func (o *ManagedDevice) HasManagedDeviceName() bool`

HasManagedDeviceName returns a boolean if a field has been set.

### SetManagedDeviceName

`func (o *ManagedDevice) SetManagedDeviceName(v string)`

SetManagedDeviceName gets a reference to the given string and assigns it to the ManagedDeviceName field.

### SetManagedDeviceNameExplicitNull

`func (o *ManagedDevice) SetManagedDeviceNameExplicitNull(b bool)`

SetManagedDeviceNameExplicitNull (un)sets ManagedDeviceName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ManagedDeviceName value is set to nil even if false is passed
### GetPartnerReportedThreatState

`func (o *ManagedDevice) GetPartnerReportedThreatState() AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState`

GetPartnerReportedThreatState returns the PartnerReportedThreatState field if non-nil, zero value otherwise.

### GetPartnerReportedThreatStateOk

`func (o *ManagedDevice) GetPartnerReportedThreatStateOk() (AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState, bool)`

GetPartnerReportedThreatStateOk returns a tuple with the PartnerReportedThreatState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPartnerReportedThreatState

`func (o *ManagedDevice) HasPartnerReportedThreatState() bool`

HasPartnerReportedThreatState returns a boolean if a field has been set.

### SetPartnerReportedThreatState

`func (o *ManagedDevice) SetPartnerReportedThreatState(v AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState)`

SetPartnerReportedThreatState gets a reference to the given AnyOfmicrosoftGraphManagedDevicePartnerReportedHealthState and assigns it to the PartnerReportedThreatState field.

### GetDeviceConfigurationStates

`func (o *ManagedDevice) GetDeviceConfigurationStates() []MicrosoftGraphDeviceConfigurationState`

GetDeviceConfigurationStates returns the DeviceConfigurationStates field if non-nil, zero value otherwise.

### GetDeviceConfigurationStatesOk

`func (o *ManagedDevice) GetDeviceConfigurationStatesOk() ([]MicrosoftGraphDeviceConfigurationState, bool)`

GetDeviceConfigurationStatesOk returns a tuple with the DeviceConfigurationStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceConfigurationStates

`func (o *ManagedDevice) HasDeviceConfigurationStates() bool`

HasDeviceConfigurationStates returns a boolean if a field has been set.

### SetDeviceConfigurationStates

`func (o *ManagedDevice) SetDeviceConfigurationStates(v []MicrosoftGraphDeviceConfigurationState)`

SetDeviceConfigurationStates gets a reference to the given []MicrosoftGraphDeviceConfigurationState and assigns it to the DeviceConfigurationStates field.

### GetDeviceCompliancePolicyStates

`func (o *ManagedDevice) GetDeviceCompliancePolicyStates() []MicrosoftGraphDeviceCompliancePolicyState`

GetDeviceCompliancePolicyStates returns the DeviceCompliancePolicyStates field if non-nil, zero value otherwise.

### GetDeviceCompliancePolicyStatesOk

`func (o *ManagedDevice) GetDeviceCompliancePolicyStatesOk() ([]MicrosoftGraphDeviceCompliancePolicyState, bool)`

GetDeviceCompliancePolicyStatesOk returns a tuple with the DeviceCompliancePolicyStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCompliancePolicyStates

`func (o *ManagedDevice) HasDeviceCompliancePolicyStates() bool`

HasDeviceCompliancePolicyStates returns a boolean if a field has been set.

### SetDeviceCompliancePolicyStates

`func (o *ManagedDevice) SetDeviceCompliancePolicyStates(v []MicrosoftGraphDeviceCompliancePolicyState)`

SetDeviceCompliancePolicyStates gets a reference to the given []MicrosoftGraphDeviceCompliancePolicyState and assigns it to the DeviceCompliancePolicyStates field.

### GetDeviceCategory

`func (o *ManagedDevice) GetDeviceCategory() AnyOfmicrosoftGraphDeviceCategory`

GetDeviceCategory returns the DeviceCategory field if non-nil, zero value otherwise.

### GetDeviceCategoryOk

`func (o *ManagedDevice) GetDeviceCategoryOk() (AnyOfmicrosoftGraphDeviceCategory, bool)`

GetDeviceCategoryOk returns a tuple with the DeviceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCategory

`func (o *ManagedDevice) HasDeviceCategory() bool`

HasDeviceCategory returns a boolean if a field has been set.

### SetDeviceCategory

`func (o *ManagedDevice) SetDeviceCategory(v AnyOfmicrosoftGraphDeviceCategory)`

SetDeviceCategory gets a reference to the given AnyOfmicrosoftGraphDeviceCategory and assigns it to the DeviceCategory field.

### SetDeviceCategoryExplicitNull

`func (o *ManagedDevice) SetDeviceCategoryExplicitNull(b bool)`

SetDeviceCategoryExplicitNull (un)sets DeviceCategory to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceCategory value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


