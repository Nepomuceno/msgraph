# MicrosoftGraphDeviceHealthAttestationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateDateTime** | Pointer to **string** | The Timestamp of the last update. | [optional] 
**ContentNamespaceUrl** | Pointer to **string** | The DHA report version. (Namespace version) | [optional] 
**DeviceHealthAttestationStatus** | Pointer to **string** | The DHA report version. (Namespace version) | [optional] 
**ContentVersion** | Pointer to **string** | The HealthAttestation state schema version | [optional] 
**IssuedDateTime** | Pointer to [**time.Time**](time.Time.md) | The DateTime when device was evaluated or issued to MDM | [optional] 
**AttestationIdentityKey** | Pointer to **string** | TWhen an Attestation Identity Key (AIK) is present on a device, it indicates that the device has an endorsement key (EK) certificate. | [optional] 
**ResetCount** | Pointer to **int64** | The number of times a PC device has hibernated or resumed | [optional] 
**RestartCount** | Pointer to **int64** | The number of times a PC device has rebooted | [optional] 
**DataExcutionPolicy** | Pointer to **string** | DEP Policy defines a set of hardware and software technologies that perform additional checks on memory  | [optional] 
**BitLockerStatus** | Pointer to **string** | On or Off of BitLocker Drive Encryption | [optional] 
**BootManagerVersion** | Pointer to **string** | The version of the Boot Manager | [optional] 
**CodeIntegrityCheckVersion** | Pointer to **string** | The version of the Boot Manager | [optional] 
**SecureBoot** | Pointer to **string** | When Secure Boot is enabled, the core components must have the correct cryptographic signatures | [optional] 
**BootDebugging** | Pointer to **string** | When bootDebugging is enabled, the device is used in development and testing | [optional] 
**OperatingSystemKernelDebugging** | Pointer to **string** | When operatingSystemKernelDebugging is enabled, the device is used in development and testing | [optional] 
**CodeIntegrity** | Pointer to **string** |  When code integrity is enabled, code execution is restricted to integrity verified code | [optional] 
**TestSigning** | Pointer to **string** | When test signing is allowed, the device does not enforce signature validation during boot | [optional] 
**SafeMode** | Pointer to **string** | Safe mode is a troubleshooting option for Windows that starts your computer in a limited state | [optional] 
**WindowsPE** | Pointer to **string** | Operating system running with limited services that is used to prepare a computer for Windows | [optional] 
**EarlyLaunchAntiMalwareDriverProtection** | Pointer to **string** | ELAM provides protection for the computers in your network when they start up | [optional] 
**VirtualSecureMode** | Pointer to **string** | VSM is a container that protects high value assets from a compromised kernel | [optional] 
**PcrHashAlgorithm** | Pointer to **string** | Informational attribute that identifies the HASH algorithm that was used by TPM | [optional] 
**BootAppSecurityVersion** | Pointer to **string** | The security version number of the Boot Application | [optional] 
**BootManagerSecurityVersion** | Pointer to **string** | The security version number of the Boot Application | [optional] 
**TpmVersion** | Pointer to **string** | The security version number of the Boot Application | [optional] 
**Pcr0** | Pointer to **string** | The measurement that is captured in PCR[0] | [optional] 
**SecureBootConfigurationPolicyFingerPrint** | Pointer to **string** | Fingerprint of the Custom Secure Boot Configuration Policy | [optional] 
**CodeIntegrityPolicy** | Pointer to **string** | The Code Integrity policy that is controlling the security of the boot environment | [optional] 
**BootRevisionListInfo** | Pointer to **string** | The Boot Revision List that was loaded during initial boot on the attested device | [optional] 
**OperatingSystemRevListInfo** | Pointer to **string** | The Operating System Revision List that was loaded during initial boot on the attested device | [optional] 
**HealthStatusMismatchInfo** | Pointer to **string** | This attribute appears if DHA-Service detects an integrity issue | [optional] 
**HealthAttestationSupportedStatus** | Pointer to **string** | This attribute indicates if DHA is supported for the device | [optional] 

## Methods

### GetLastUpdateDateTime

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetLastUpdateDateTime() string`

GetLastUpdateDateTime returns the LastUpdateDateTime field if non-nil, zero value otherwise.

### GetLastUpdateDateTimeOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetLastUpdateDateTimeOk() (string, bool)`

GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastUpdateDateTime

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasLastUpdateDateTime() bool`

HasLastUpdateDateTime returns a boolean if a field has been set.

### SetLastUpdateDateTime

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetLastUpdateDateTime(v string)`

SetLastUpdateDateTime gets a reference to the given string and assigns it to the LastUpdateDateTime field.

### SetLastUpdateDateTimeExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetLastUpdateDateTimeExplicitNull(b bool)`

SetLastUpdateDateTimeExplicitNull (un)sets LastUpdateDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastUpdateDateTime value is set to nil even if false is passed
### GetContentNamespaceUrl

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetContentNamespaceUrl() string`

GetContentNamespaceUrl returns the ContentNamespaceUrl field if non-nil, zero value otherwise.

### GetContentNamespaceUrlOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetContentNamespaceUrlOk() (string, bool)`

GetContentNamespaceUrlOk returns a tuple with the ContentNamespaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentNamespaceUrl

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasContentNamespaceUrl() bool`

HasContentNamespaceUrl returns a boolean if a field has been set.

### SetContentNamespaceUrl

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetContentNamespaceUrl(v string)`

SetContentNamespaceUrl gets a reference to the given string and assigns it to the ContentNamespaceUrl field.

### SetContentNamespaceUrlExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetContentNamespaceUrlExplicitNull(b bool)`

SetContentNamespaceUrlExplicitNull (un)sets ContentNamespaceUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentNamespaceUrl value is set to nil even if false is passed
### GetDeviceHealthAttestationStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetDeviceHealthAttestationStatus() string`

GetDeviceHealthAttestationStatus returns the DeviceHealthAttestationStatus field if non-nil, zero value otherwise.

### GetDeviceHealthAttestationStatusOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetDeviceHealthAttestationStatusOk() (string, bool)`

GetDeviceHealthAttestationStatusOk returns a tuple with the DeviceHealthAttestationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceHealthAttestationStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasDeviceHealthAttestationStatus() bool`

HasDeviceHealthAttestationStatus returns a boolean if a field has been set.

### SetDeviceHealthAttestationStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetDeviceHealthAttestationStatus(v string)`

SetDeviceHealthAttestationStatus gets a reference to the given string and assigns it to the DeviceHealthAttestationStatus field.

### SetDeviceHealthAttestationStatusExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetDeviceHealthAttestationStatusExplicitNull(b bool)`

SetDeviceHealthAttestationStatusExplicitNull (un)sets DeviceHealthAttestationStatus to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceHealthAttestationStatus value is set to nil even if false is passed
### GetContentVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetContentVersion() string`

GetContentVersion returns the ContentVersion field if non-nil, zero value otherwise.

### GetContentVersionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetContentVersionOk() (string, bool)`

GetContentVersionOk returns a tuple with the ContentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasContentVersion() bool`

HasContentVersion returns a boolean if a field has been set.

### SetContentVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetContentVersion(v string)`

SetContentVersion gets a reference to the given string and assigns it to the ContentVersion field.

### SetContentVersionExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetContentVersionExplicitNull(b bool)`

SetContentVersionExplicitNull (un)sets ContentVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentVersion value is set to nil even if false is passed
### GetIssuedDateTime

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetIssuedDateTime() time.Time`

GetIssuedDateTime returns the IssuedDateTime field if non-nil, zero value otherwise.

### GetIssuedDateTimeOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetIssuedDateTimeOk() (time.Time, bool)`

GetIssuedDateTimeOk returns a tuple with the IssuedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIssuedDateTime

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasIssuedDateTime() bool`

HasIssuedDateTime returns a boolean if a field has been set.

### SetIssuedDateTime

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetIssuedDateTime(v time.Time)`

SetIssuedDateTime gets a reference to the given time.Time and assigns it to the IssuedDateTime field.

### GetAttestationIdentityKey

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetAttestationIdentityKey() string`

GetAttestationIdentityKey returns the AttestationIdentityKey field if non-nil, zero value otherwise.

### GetAttestationIdentityKeyOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetAttestationIdentityKeyOk() (string, bool)`

GetAttestationIdentityKeyOk returns a tuple with the AttestationIdentityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttestationIdentityKey

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasAttestationIdentityKey() bool`

HasAttestationIdentityKey returns a boolean if a field has been set.

### SetAttestationIdentityKey

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetAttestationIdentityKey(v string)`

SetAttestationIdentityKey gets a reference to the given string and assigns it to the AttestationIdentityKey field.

### SetAttestationIdentityKeyExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetAttestationIdentityKeyExplicitNull(b bool)`

SetAttestationIdentityKeyExplicitNull (un)sets AttestationIdentityKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AttestationIdentityKey value is set to nil even if false is passed
### GetResetCount

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetResetCount() int64`

GetResetCount returns the ResetCount field if non-nil, zero value otherwise.

### GetResetCountOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetResetCountOk() (int64, bool)`

GetResetCountOk returns a tuple with the ResetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResetCount

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasResetCount() bool`

HasResetCount returns a boolean if a field has been set.

### SetResetCount

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetResetCount(v int64)`

SetResetCount gets a reference to the given int64 and assigns it to the ResetCount field.

### GetRestartCount

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetRestartCount() int64`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetRestartCountOk() (int64, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRestartCount

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasRestartCount() bool`

HasRestartCount returns a boolean if a field has been set.

### SetRestartCount

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetRestartCount(v int64)`

SetRestartCount gets a reference to the given int64 and assigns it to the RestartCount field.

### GetDataExcutionPolicy

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetDataExcutionPolicy() string`

GetDataExcutionPolicy returns the DataExcutionPolicy field if non-nil, zero value otherwise.

### GetDataExcutionPolicyOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetDataExcutionPolicyOk() (string, bool)`

GetDataExcutionPolicyOk returns a tuple with the DataExcutionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDataExcutionPolicy

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasDataExcutionPolicy() bool`

HasDataExcutionPolicy returns a boolean if a field has been set.

### SetDataExcutionPolicy

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetDataExcutionPolicy(v string)`

SetDataExcutionPolicy gets a reference to the given string and assigns it to the DataExcutionPolicy field.

### SetDataExcutionPolicyExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetDataExcutionPolicyExplicitNull(b bool)`

SetDataExcutionPolicyExplicitNull (un)sets DataExcutionPolicy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DataExcutionPolicy value is set to nil even if false is passed
### GetBitLockerStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBitLockerStatus() string`

GetBitLockerStatus returns the BitLockerStatus field if non-nil, zero value otherwise.

### GetBitLockerStatusOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBitLockerStatusOk() (string, bool)`

GetBitLockerStatusOk returns a tuple with the BitLockerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBitLockerStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasBitLockerStatus() bool`

HasBitLockerStatus returns a boolean if a field has been set.

### SetBitLockerStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBitLockerStatus(v string)`

SetBitLockerStatus gets a reference to the given string and assigns it to the BitLockerStatus field.

### SetBitLockerStatusExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBitLockerStatusExplicitNull(b bool)`

SetBitLockerStatusExplicitNull (un)sets BitLockerStatus to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BitLockerStatus value is set to nil even if false is passed
### GetBootManagerVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootManagerVersion() string`

GetBootManagerVersion returns the BootManagerVersion field if non-nil, zero value otherwise.

### GetBootManagerVersionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootManagerVersionOk() (string, bool)`

GetBootManagerVersionOk returns a tuple with the BootManagerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBootManagerVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasBootManagerVersion() bool`

HasBootManagerVersion returns a boolean if a field has been set.

### SetBootManagerVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootManagerVersion(v string)`

SetBootManagerVersion gets a reference to the given string and assigns it to the BootManagerVersion field.

### SetBootManagerVersionExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootManagerVersionExplicitNull(b bool)`

SetBootManagerVersionExplicitNull (un)sets BootManagerVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BootManagerVersion value is set to nil even if false is passed
### GetCodeIntegrityCheckVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetCodeIntegrityCheckVersion() string`

GetCodeIntegrityCheckVersion returns the CodeIntegrityCheckVersion field if non-nil, zero value otherwise.

### GetCodeIntegrityCheckVersionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetCodeIntegrityCheckVersionOk() (string, bool)`

GetCodeIntegrityCheckVersionOk returns a tuple with the CodeIntegrityCheckVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCodeIntegrityCheckVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasCodeIntegrityCheckVersion() bool`

HasCodeIntegrityCheckVersion returns a boolean if a field has been set.

### SetCodeIntegrityCheckVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetCodeIntegrityCheckVersion(v string)`

SetCodeIntegrityCheckVersion gets a reference to the given string and assigns it to the CodeIntegrityCheckVersion field.

### SetCodeIntegrityCheckVersionExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetCodeIntegrityCheckVersionExplicitNull(b bool)`

SetCodeIntegrityCheckVersionExplicitNull (un)sets CodeIntegrityCheckVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CodeIntegrityCheckVersion value is set to nil even if false is passed
### GetSecureBoot

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetSecureBoot() string`

GetSecureBoot returns the SecureBoot field if non-nil, zero value otherwise.

### GetSecureBootOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetSecureBootOk() (string, bool)`

GetSecureBootOk returns a tuple with the SecureBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecureBoot

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasSecureBoot() bool`

HasSecureBoot returns a boolean if a field has been set.

### SetSecureBoot

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetSecureBoot(v string)`

SetSecureBoot gets a reference to the given string and assigns it to the SecureBoot field.

### SetSecureBootExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetSecureBootExplicitNull(b bool)`

SetSecureBootExplicitNull (un)sets SecureBoot to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SecureBoot value is set to nil even if false is passed
### GetBootDebugging

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootDebugging() string`

GetBootDebugging returns the BootDebugging field if non-nil, zero value otherwise.

### GetBootDebuggingOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootDebuggingOk() (string, bool)`

GetBootDebuggingOk returns a tuple with the BootDebugging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBootDebugging

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasBootDebugging() bool`

HasBootDebugging returns a boolean if a field has been set.

### SetBootDebugging

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootDebugging(v string)`

SetBootDebugging gets a reference to the given string and assigns it to the BootDebugging field.

### SetBootDebuggingExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootDebuggingExplicitNull(b bool)`

SetBootDebuggingExplicitNull (un)sets BootDebugging to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BootDebugging value is set to nil even if false is passed
### GetOperatingSystemKernelDebugging

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetOperatingSystemKernelDebugging() string`

GetOperatingSystemKernelDebugging returns the OperatingSystemKernelDebugging field if non-nil, zero value otherwise.

### GetOperatingSystemKernelDebuggingOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetOperatingSystemKernelDebuggingOk() (string, bool)`

GetOperatingSystemKernelDebuggingOk returns a tuple with the OperatingSystemKernelDebugging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperatingSystemKernelDebugging

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasOperatingSystemKernelDebugging() bool`

HasOperatingSystemKernelDebugging returns a boolean if a field has been set.

### SetOperatingSystemKernelDebugging

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetOperatingSystemKernelDebugging(v string)`

SetOperatingSystemKernelDebugging gets a reference to the given string and assigns it to the OperatingSystemKernelDebugging field.

### SetOperatingSystemKernelDebuggingExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetOperatingSystemKernelDebuggingExplicitNull(b bool)`

SetOperatingSystemKernelDebuggingExplicitNull (un)sets OperatingSystemKernelDebugging to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperatingSystemKernelDebugging value is set to nil even if false is passed
### GetCodeIntegrity

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetCodeIntegrity() string`

GetCodeIntegrity returns the CodeIntegrity field if non-nil, zero value otherwise.

### GetCodeIntegrityOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetCodeIntegrityOk() (string, bool)`

GetCodeIntegrityOk returns a tuple with the CodeIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCodeIntegrity

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasCodeIntegrity() bool`

HasCodeIntegrity returns a boolean if a field has been set.

### SetCodeIntegrity

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetCodeIntegrity(v string)`

SetCodeIntegrity gets a reference to the given string and assigns it to the CodeIntegrity field.

### SetCodeIntegrityExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetCodeIntegrityExplicitNull(b bool)`

SetCodeIntegrityExplicitNull (un)sets CodeIntegrity to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CodeIntegrity value is set to nil even if false is passed
### GetTestSigning

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetTestSigning() string`

GetTestSigning returns the TestSigning field if non-nil, zero value otherwise.

### GetTestSigningOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetTestSigningOk() (string, bool)`

GetTestSigningOk returns a tuple with the TestSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTestSigning

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasTestSigning() bool`

HasTestSigning returns a boolean if a field has been set.

### SetTestSigning

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetTestSigning(v string)`

SetTestSigning gets a reference to the given string and assigns it to the TestSigning field.

### SetTestSigningExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetTestSigningExplicitNull(b bool)`

SetTestSigningExplicitNull (un)sets TestSigning to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TestSigning value is set to nil even if false is passed
### GetSafeMode

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetSafeMode() string`

GetSafeMode returns the SafeMode field if non-nil, zero value otherwise.

### GetSafeModeOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetSafeModeOk() (string, bool)`

GetSafeModeOk returns a tuple with the SafeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSafeMode

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasSafeMode() bool`

HasSafeMode returns a boolean if a field has been set.

### SetSafeMode

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetSafeMode(v string)`

SetSafeMode gets a reference to the given string and assigns it to the SafeMode field.

### SetSafeModeExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetSafeModeExplicitNull(b bool)`

SetSafeModeExplicitNull (un)sets SafeMode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SafeMode value is set to nil even if false is passed
### GetWindowsPE

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetWindowsPE() string`

GetWindowsPE returns the WindowsPE field if non-nil, zero value otherwise.

### GetWindowsPEOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetWindowsPEOk() (string, bool)`

GetWindowsPEOk returns a tuple with the WindowsPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsPE

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasWindowsPE() bool`

HasWindowsPE returns a boolean if a field has been set.

### SetWindowsPE

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetWindowsPE(v string)`

SetWindowsPE gets a reference to the given string and assigns it to the WindowsPE field.

### SetWindowsPEExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetWindowsPEExplicitNull(b bool)`

SetWindowsPEExplicitNull (un)sets WindowsPE to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WindowsPE value is set to nil even if false is passed
### GetEarlyLaunchAntiMalwareDriverProtection

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetEarlyLaunchAntiMalwareDriverProtection() string`

GetEarlyLaunchAntiMalwareDriverProtection returns the EarlyLaunchAntiMalwareDriverProtection field if non-nil, zero value otherwise.

### GetEarlyLaunchAntiMalwareDriverProtectionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetEarlyLaunchAntiMalwareDriverProtectionOk() (string, bool)`

GetEarlyLaunchAntiMalwareDriverProtectionOk returns a tuple with the EarlyLaunchAntiMalwareDriverProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEarlyLaunchAntiMalwareDriverProtection

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasEarlyLaunchAntiMalwareDriverProtection() bool`

HasEarlyLaunchAntiMalwareDriverProtection returns a boolean if a field has been set.

### SetEarlyLaunchAntiMalwareDriverProtection

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetEarlyLaunchAntiMalwareDriverProtection(v string)`

SetEarlyLaunchAntiMalwareDriverProtection gets a reference to the given string and assigns it to the EarlyLaunchAntiMalwareDriverProtection field.

### SetEarlyLaunchAntiMalwareDriverProtectionExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetEarlyLaunchAntiMalwareDriverProtectionExplicitNull(b bool)`

SetEarlyLaunchAntiMalwareDriverProtectionExplicitNull (un)sets EarlyLaunchAntiMalwareDriverProtection to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EarlyLaunchAntiMalwareDriverProtection value is set to nil even if false is passed
### GetVirtualSecureMode

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetVirtualSecureMode() string`

GetVirtualSecureMode returns the VirtualSecureMode field if non-nil, zero value otherwise.

### GetVirtualSecureModeOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetVirtualSecureModeOk() (string, bool)`

GetVirtualSecureModeOk returns a tuple with the VirtualSecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualSecureMode

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasVirtualSecureMode() bool`

HasVirtualSecureMode returns a boolean if a field has been set.

### SetVirtualSecureMode

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetVirtualSecureMode(v string)`

SetVirtualSecureMode gets a reference to the given string and assigns it to the VirtualSecureMode field.

### SetVirtualSecureModeExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetVirtualSecureModeExplicitNull(b bool)`

SetVirtualSecureModeExplicitNull (un)sets VirtualSecureMode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VirtualSecureMode value is set to nil even if false is passed
### GetPcrHashAlgorithm

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetPcrHashAlgorithm() string`

GetPcrHashAlgorithm returns the PcrHashAlgorithm field if non-nil, zero value otherwise.

### GetPcrHashAlgorithmOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetPcrHashAlgorithmOk() (string, bool)`

GetPcrHashAlgorithmOk returns a tuple with the PcrHashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPcrHashAlgorithm

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasPcrHashAlgorithm() bool`

HasPcrHashAlgorithm returns a boolean if a field has been set.

### SetPcrHashAlgorithm

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetPcrHashAlgorithm(v string)`

SetPcrHashAlgorithm gets a reference to the given string and assigns it to the PcrHashAlgorithm field.

### SetPcrHashAlgorithmExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetPcrHashAlgorithmExplicitNull(b bool)`

SetPcrHashAlgorithmExplicitNull (un)sets PcrHashAlgorithm to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PcrHashAlgorithm value is set to nil even if false is passed
### GetBootAppSecurityVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootAppSecurityVersion() string`

GetBootAppSecurityVersion returns the BootAppSecurityVersion field if non-nil, zero value otherwise.

### GetBootAppSecurityVersionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootAppSecurityVersionOk() (string, bool)`

GetBootAppSecurityVersionOk returns a tuple with the BootAppSecurityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBootAppSecurityVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasBootAppSecurityVersion() bool`

HasBootAppSecurityVersion returns a boolean if a field has been set.

### SetBootAppSecurityVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootAppSecurityVersion(v string)`

SetBootAppSecurityVersion gets a reference to the given string and assigns it to the BootAppSecurityVersion field.

### SetBootAppSecurityVersionExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootAppSecurityVersionExplicitNull(b bool)`

SetBootAppSecurityVersionExplicitNull (un)sets BootAppSecurityVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BootAppSecurityVersion value is set to nil even if false is passed
### GetBootManagerSecurityVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootManagerSecurityVersion() string`

GetBootManagerSecurityVersion returns the BootManagerSecurityVersion field if non-nil, zero value otherwise.

### GetBootManagerSecurityVersionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootManagerSecurityVersionOk() (string, bool)`

GetBootManagerSecurityVersionOk returns a tuple with the BootManagerSecurityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBootManagerSecurityVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasBootManagerSecurityVersion() bool`

HasBootManagerSecurityVersion returns a boolean if a field has been set.

### SetBootManagerSecurityVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootManagerSecurityVersion(v string)`

SetBootManagerSecurityVersion gets a reference to the given string and assigns it to the BootManagerSecurityVersion field.

### SetBootManagerSecurityVersionExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootManagerSecurityVersionExplicitNull(b bool)`

SetBootManagerSecurityVersionExplicitNull (un)sets BootManagerSecurityVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BootManagerSecurityVersion value is set to nil even if false is passed
### GetTpmVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetTpmVersion() string`

GetTpmVersion returns the TpmVersion field if non-nil, zero value otherwise.

### GetTpmVersionOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetTpmVersionOk() (string, bool)`

GetTpmVersionOk returns a tuple with the TpmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTpmVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasTpmVersion() bool`

HasTpmVersion returns a boolean if a field has been set.

### SetTpmVersion

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetTpmVersion(v string)`

SetTpmVersion gets a reference to the given string and assigns it to the TpmVersion field.

### SetTpmVersionExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetTpmVersionExplicitNull(b bool)`

SetTpmVersionExplicitNull (un)sets TpmVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TpmVersion value is set to nil even if false is passed
### GetPcr0

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetPcr0() string`

GetPcr0 returns the Pcr0 field if non-nil, zero value otherwise.

### GetPcr0Ok

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetPcr0Ok() (string, bool)`

GetPcr0Ok returns a tuple with the Pcr0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPcr0

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasPcr0() bool`

HasPcr0 returns a boolean if a field has been set.

### SetPcr0

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetPcr0(v string)`

SetPcr0 gets a reference to the given string and assigns it to the Pcr0 field.

### SetPcr0ExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetPcr0ExplicitNull(b bool)`

SetPcr0ExplicitNull (un)sets Pcr0 to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Pcr0 value is set to nil even if false is passed
### GetSecureBootConfigurationPolicyFingerPrint

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetSecureBootConfigurationPolicyFingerPrint() string`

GetSecureBootConfigurationPolicyFingerPrint returns the SecureBootConfigurationPolicyFingerPrint field if non-nil, zero value otherwise.

### GetSecureBootConfigurationPolicyFingerPrintOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetSecureBootConfigurationPolicyFingerPrintOk() (string, bool)`

GetSecureBootConfigurationPolicyFingerPrintOk returns a tuple with the SecureBootConfigurationPolicyFingerPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecureBootConfigurationPolicyFingerPrint

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasSecureBootConfigurationPolicyFingerPrint() bool`

HasSecureBootConfigurationPolicyFingerPrint returns a boolean if a field has been set.

### SetSecureBootConfigurationPolicyFingerPrint

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetSecureBootConfigurationPolicyFingerPrint(v string)`

SetSecureBootConfigurationPolicyFingerPrint gets a reference to the given string and assigns it to the SecureBootConfigurationPolicyFingerPrint field.

### SetSecureBootConfigurationPolicyFingerPrintExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetSecureBootConfigurationPolicyFingerPrintExplicitNull(b bool)`

SetSecureBootConfigurationPolicyFingerPrintExplicitNull (un)sets SecureBootConfigurationPolicyFingerPrint to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SecureBootConfigurationPolicyFingerPrint value is set to nil even if false is passed
### GetCodeIntegrityPolicy

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetCodeIntegrityPolicy() string`

GetCodeIntegrityPolicy returns the CodeIntegrityPolicy field if non-nil, zero value otherwise.

### GetCodeIntegrityPolicyOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetCodeIntegrityPolicyOk() (string, bool)`

GetCodeIntegrityPolicyOk returns a tuple with the CodeIntegrityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCodeIntegrityPolicy

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasCodeIntegrityPolicy() bool`

HasCodeIntegrityPolicy returns a boolean if a field has been set.

### SetCodeIntegrityPolicy

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetCodeIntegrityPolicy(v string)`

SetCodeIntegrityPolicy gets a reference to the given string and assigns it to the CodeIntegrityPolicy field.

### SetCodeIntegrityPolicyExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetCodeIntegrityPolicyExplicitNull(b bool)`

SetCodeIntegrityPolicyExplicitNull (un)sets CodeIntegrityPolicy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CodeIntegrityPolicy value is set to nil even if false is passed
### GetBootRevisionListInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootRevisionListInfo() string`

GetBootRevisionListInfo returns the BootRevisionListInfo field if non-nil, zero value otherwise.

### GetBootRevisionListInfoOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetBootRevisionListInfoOk() (string, bool)`

GetBootRevisionListInfoOk returns a tuple with the BootRevisionListInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBootRevisionListInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasBootRevisionListInfo() bool`

HasBootRevisionListInfo returns a boolean if a field has been set.

### SetBootRevisionListInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootRevisionListInfo(v string)`

SetBootRevisionListInfo gets a reference to the given string and assigns it to the BootRevisionListInfo field.

### SetBootRevisionListInfoExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetBootRevisionListInfoExplicitNull(b bool)`

SetBootRevisionListInfoExplicitNull (un)sets BootRevisionListInfo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BootRevisionListInfo value is set to nil even if false is passed
### GetOperatingSystemRevListInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetOperatingSystemRevListInfo() string`

GetOperatingSystemRevListInfo returns the OperatingSystemRevListInfo field if non-nil, zero value otherwise.

### GetOperatingSystemRevListInfoOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetOperatingSystemRevListInfoOk() (string, bool)`

GetOperatingSystemRevListInfoOk returns a tuple with the OperatingSystemRevListInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperatingSystemRevListInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasOperatingSystemRevListInfo() bool`

HasOperatingSystemRevListInfo returns a boolean if a field has been set.

### SetOperatingSystemRevListInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetOperatingSystemRevListInfo(v string)`

SetOperatingSystemRevListInfo gets a reference to the given string and assigns it to the OperatingSystemRevListInfo field.

### SetOperatingSystemRevListInfoExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetOperatingSystemRevListInfoExplicitNull(b bool)`

SetOperatingSystemRevListInfoExplicitNull (un)sets OperatingSystemRevListInfo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperatingSystemRevListInfo value is set to nil even if false is passed
### GetHealthStatusMismatchInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetHealthStatusMismatchInfo() string`

GetHealthStatusMismatchInfo returns the HealthStatusMismatchInfo field if non-nil, zero value otherwise.

### GetHealthStatusMismatchInfoOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetHealthStatusMismatchInfoOk() (string, bool)`

GetHealthStatusMismatchInfoOk returns a tuple with the HealthStatusMismatchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHealthStatusMismatchInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasHealthStatusMismatchInfo() bool`

HasHealthStatusMismatchInfo returns a boolean if a field has been set.

### SetHealthStatusMismatchInfo

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetHealthStatusMismatchInfo(v string)`

SetHealthStatusMismatchInfo gets a reference to the given string and assigns it to the HealthStatusMismatchInfo field.

### SetHealthStatusMismatchInfoExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetHealthStatusMismatchInfoExplicitNull(b bool)`

SetHealthStatusMismatchInfoExplicitNull (un)sets HealthStatusMismatchInfo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HealthStatusMismatchInfo value is set to nil even if false is passed
### GetHealthAttestationSupportedStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetHealthAttestationSupportedStatus() string`

GetHealthAttestationSupportedStatus returns the HealthAttestationSupportedStatus field if non-nil, zero value otherwise.

### GetHealthAttestationSupportedStatusOk

`func (o *MicrosoftGraphDeviceHealthAttestationState) GetHealthAttestationSupportedStatusOk() (string, bool)`

GetHealthAttestationSupportedStatusOk returns a tuple with the HealthAttestationSupportedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHealthAttestationSupportedStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) HasHealthAttestationSupportedStatus() bool`

HasHealthAttestationSupportedStatus returns a boolean if a field has been set.

### SetHealthAttestationSupportedStatus

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetHealthAttestationSupportedStatus(v string)`

SetHealthAttestationSupportedStatus gets a reference to the given string and assigns it to the HealthAttestationSupportedStatus field.

### SetHealthAttestationSupportedStatusExplicitNull

`func (o *MicrosoftGraphDeviceHealthAttestationState) SetHealthAttestationSupportedStatusExplicitNull(b bool)`

SetHealthAttestationSupportedStatusExplicitNull (un)sets HealthAttestationSupportedStatus to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HealthAttestationSupportedStatus value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


