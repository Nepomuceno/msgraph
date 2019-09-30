# AndroidCompliancePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordRequired** | Pointer to **bool** | Require a password to unlock device. | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | Minimum password length. Valid values 4 to 16 | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphAndroidRequiredPasswordType**](anyOf&lt;microsoft.graph.androidRequiredPasswordType&gt;.md) | Type of characters in password | [optional] 
**PasswordMinutesOfInactivityBeforeLock** | Pointer to **int32** | Minutes of inactivity before a password is required. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | Number of days before the password expires. Valid values 1 to 365 | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | Number of previous passwords to block. Valid values 1 to 24 | [optional] 
**SecurityPreventInstallAppsFromUnknownSources** | Pointer to **bool** | Require that devices disallow installation of apps from unknown sources. | [optional] 
**SecurityDisableUsbDebugging** | Pointer to **bool** | Disable USB debugging on Android devices. | [optional] 
**SecurityRequireVerifyApps** | Pointer to **bool** | Require the Android Verify apps feature is turned on. | [optional] 
**DeviceThreatProtectionEnabled** | Pointer to **bool** | Require that devices have enabled device threat protection. | [optional] 
**DeviceThreatProtectionRequiredSecurityLevel** | Pointer to [**AnyOfmicrosoftGraphDeviceThreatProtectionLevel**](anyOf&lt;microsoft.graph.deviceThreatProtectionLevel&gt;.md) | Require Mobile Threat Protection minimum risk level to report noncompliance. | [optional] 
**SecurityBlockJailbrokenDevices** | Pointer to **bool** | Devices must not be jailbroken or rooted. | [optional] 
**OsMinimumVersion** | Pointer to **string** | Minimum Android version. | [optional] 
**OsMaximumVersion** | Pointer to **string** | Maximum Android version. | [optional] 
**MinAndroidSecurityPatchLevel** | Pointer to **string** | Minimum Android security patch level. | [optional] 
**StorageRequireEncryption** | Pointer to **bool** | Require encryption on Android devices. | [optional] 
**SecurityRequireSafetyNetAttestationBasicIntegrity** | Pointer to **bool** | Require the device to pass the SafetyNet basic integrity check. | [optional] 
**SecurityRequireSafetyNetAttestationCertifiedDevice** | Pointer to **bool** | Require the device to pass the SafetyNet certified device check. | [optional] 
**SecurityRequireGooglePlayServices** | Pointer to **bool** | Require Google Play Services to be installed and enabled on the device. | [optional] 
**SecurityRequireUpToDateSecurityProviders** | Pointer to **bool** | Require the device to have up to date security providers. The device will require Google Play Services to be enabled and up to date. | [optional] 
**SecurityRequireCompanyPortalAppIntegrity** | Pointer to **bool** | Require the device to pass the Company Portal client app runtime integrity check. | [optional] 

## Methods

### GetPasswordRequired

`func (o *AndroidCompliancePolicy) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *AndroidCompliancePolicy) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *AndroidCompliancePolicy) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *AndroidCompliancePolicy) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetPasswordMinimumLength

`func (o *AndroidCompliancePolicy) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *AndroidCompliancePolicy) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *AndroidCompliancePolicy) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *AndroidCompliancePolicy) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *AndroidCompliancePolicy) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *AndroidCompliancePolicy) GetPasswordRequiredType() AnyOfmicrosoftGraphAndroidRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *AndroidCompliancePolicy) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphAndroidRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *AndroidCompliancePolicy) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *AndroidCompliancePolicy) SetPasswordRequiredType(v AnyOfmicrosoftGraphAndroidRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphAndroidRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordMinutesOfInactivityBeforeLock

`func (o *AndroidCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock() int32`

GetPasswordMinutesOfInactivityBeforeLock returns the PasswordMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeLockOk

`func (o *AndroidCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLockOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeLockOk returns a tuple with the PasswordMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeLock

`func (o *AndroidCompliancePolicy) HasPasswordMinutesOfInactivityBeforeLock() bool`

HasPasswordMinutesOfInactivityBeforeLock returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeLock

`func (o *AndroidCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(v int32)`

SetPasswordMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeLock field.

### SetPasswordMinutesOfInactivityBeforeLockExplicitNull

`func (o *AndroidCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLockExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeLockExplicitNull (un)sets PasswordMinutesOfInactivityBeforeLock to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeLock value is set to nil even if false is passed
### GetPasswordExpirationDays

`func (o *AndroidCompliancePolicy) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *AndroidCompliancePolicy) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *AndroidCompliancePolicy) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *AndroidCompliancePolicy) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *AndroidCompliancePolicy) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *AndroidCompliancePolicy) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *AndroidCompliancePolicy) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *AndroidCompliancePolicy) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *AndroidCompliancePolicy) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *AndroidCompliancePolicy) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetSecurityPreventInstallAppsFromUnknownSources

`func (o *AndroidCompliancePolicy) GetSecurityPreventInstallAppsFromUnknownSources() bool`

GetSecurityPreventInstallAppsFromUnknownSources returns the SecurityPreventInstallAppsFromUnknownSources field if non-nil, zero value otherwise.

### GetSecurityPreventInstallAppsFromUnknownSourcesOk

`func (o *AndroidCompliancePolicy) GetSecurityPreventInstallAppsFromUnknownSourcesOk() (bool, bool)`

GetSecurityPreventInstallAppsFromUnknownSourcesOk returns a tuple with the SecurityPreventInstallAppsFromUnknownSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityPreventInstallAppsFromUnknownSources

`func (o *AndroidCompliancePolicy) HasSecurityPreventInstallAppsFromUnknownSources() bool`

HasSecurityPreventInstallAppsFromUnknownSources returns a boolean if a field has been set.

### SetSecurityPreventInstallAppsFromUnknownSources

`func (o *AndroidCompliancePolicy) SetSecurityPreventInstallAppsFromUnknownSources(v bool)`

SetSecurityPreventInstallAppsFromUnknownSources gets a reference to the given bool and assigns it to the SecurityPreventInstallAppsFromUnknownSources field.

### GetSecurityDisableUsbDebugging

`func (o *AndroidCompliancePolicy) GetSecurityDisableUsbDebugging() bool`

GetSecurityDisableUsbDebugging returns the SecurityDisableUsbDebugging field if non-nil, zero value otherwise.

### GetSecurityDisableUsbDebuggingOk

`func (o *AndroidCompliancePolicy) GetSecurityDisableUsbDebuggingOk() (bool, bool)`

GetSecurityDisableUsbDebuggingOk returns a tuple with the SecurityDisableUsbDebugging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityDisableUsbDebugging

`func (o *AndroidCompliancePolicy) HasSecurityDisableUsbDebugging() bool`

HasSecurityDisableUsbDebugging returns a boolean if a field has been set.

### SetSecurityDisableUsbDebugging

`func (o *AndroidCompliancePolicy) SetSecurityDisableUsbDebugging(v bool)`

SetSecurityDisableUsbDebugging gets a reference to the given bool and assigns it to the SecurityDisableUsbDebugging field.

### GetSecurityRequireVerifyApps

`func (o *AndroidCompliancePolicy) GetSecurityRequireVerifyApps() bool`

GetSecurityRequireVerifyApps returns the SecurityRequireVerifyApps field if non-nil, zero value otherwise.

### GetSecurityRequireVerifyAppsOk

`func (o *AndroidCompliancePolicy) GetSecurityRequireVerifyAppsOk() (bool, bool)`

GetSecurityRequireVerifyAppsOk returns a tuple with the SecurityRequireVerifyApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireVerifyApps

`func (o *AndroidCompliancePolicy) HasSecurityRequireVerifyApps() bool`

HasSecurityRequireVerifyApps returns a boolean if a field has been set.

### SetSecurityRequireVerifyApps

`func (o *AndroidCompliancePolicy) SetSecurityRequireVerifyApps(v bool)`

SetSecurityRequireVerifyApps gets a reference to the given bool and assigns it to the SecurityRequireVerifyApps field.

### GetDeviceThreatProtectionEnabled

`func (o *AndroidCompliancePolicy) GetDeviceThreatProtectionEnabled() bool`

GetDeviceThreatProtectionEnabled returns the DeviceThreatProtectionEnabled field if non-nil, zero value otherwise.

### GetDeviceThreatProtectionEnabledOk

`func (o *AndroidCompliancePolicy) GetDeviceThreatProtectionEnabledOk() (bool, bool)`

GetDeviceThreatProtectionEnabledOk returns a tuple with the DeviceThreatProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceThreatProtectionEnabled

`func (o *AndroidCompliancePolicy) HasDeviceThreatProtectionEnabled() bool`

HasDeviceThreatProtectionEnabled returns a boolean if a field has been set.

### SetDeviceThreatProtectionEnabled

`func (o *AndroidCompliancePolicy) SetDeviceThreatProtectionEnabled(v bool)`

SetDeviceThreatProtectionEnabled gets a reference to the given bool and assigns it to the DeviceThreatProtectionEnabled field.

### GetDeviceThreatProtectionRequiredSecurityLevel

`func (o *AndroidCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel() AnyOfmicrosoftGraphDeviceThreatProtectionLevel`

GetDeviceThreatProtectionRequiredSecurityLevel returns the DeviceThreatProtectionRequiredSecurityLevel field if non-nil, zero value otherwise.

### GetDeviceThreatProtectionRequiredSecurityLevelOk

`func (o *AndroidCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevelOk() (AnyOfmicrosoftGraphDeviceThreatProtectionLevel, bool)`

GetDeviceThreatProtectionRequiredSecurityLevelOk returns a tuple with the DeviceThreatProtectionRequiredSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceThreatProtectionRequiredSecurityLevel

`func (o *AndroidCompliancePolicy) HasDeviceThreatProtectionRequiredSecurityLevel() bool`

HasDeviceThreatProtectionRequiredSecurityLevel returns a boolean if a field has been set.

### SetDeviceThreatProtectionRequiredSecurityLevel

`func (o *AndroidCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(v AnyOfmicrosoftGraphDeviceThreatProtectionLevel)`

SetDeviceThreatProtectionRequiredSecurityLevel gets a reference to the given AnyOfmicrosoftGraphDeviceThreatProtectionLevel and assigns it to the DeviceThreatProtectionRequiredSecurityLevel field.

### GetSecurityBlockJailbrokenDevices

`func (o *AndroidCompliancePolicy) GetSecurityBlockJailbrokenDevices() bool`

GetSecurityBlockJailbrokenDevices returns the SecurityBlockJailbrokenDevices field if non-nil, zero value otherwise.

### GetSecurityBlockJailbrokenDevicesOk

`func (o *AndroidCompliancePolicy) GetSecurityBlockJailbrokenDevicesOk() (bool, bool)`

GetSecurityBlockJailbrokenDevicesOk returns a tuple with the SecurityBlockJailbrokenDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityBlockJailbrokenDevices

`func (o *AndroidCompliancePolicy) HasSecurityBlockJailbrokenDevices() bool`

HasSecurityBlockJailbrokenDevices returns a boolean if a field has been set.

### SetSecurityBlockJailbrokenDevices

`func (o *AndroidCompliancePolicy) SetSecurityBlockJailbrokenDevices(v bool)`

SetSecurityBlockJailbrokenDevices gets a reference to the given bool and assigns it to the SecurityBlockJailbrokenDevices field.

### GetOsMinimumVersion

`func (o *AndroidCompliancePolicy) GetOsMinimumVersion() string`

GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.

### GetOsMinimumVersionOk

`func (o *AndroidCompliancePolicy) GetOsMinimumVersionOk() (string, bool)`

GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMinimumVersion

`func (o *AndroidCompliancePolicy) HasOsMinimumVersion() bool`

HasOsMinimumVersion returns a boolean if a field has been set.

### SetOsMinimumVersion

`func (o *AndroidCompliancePolicy) SetOsMinimumVersion(v string)`

SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.

### SetOsMinimumVersionExplicitNull

`func (o *AndroidCompliancePolicy) SetOsMinimumVersionExplicitNull(b bool)`

SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMinimumVersion value is set to nil even if false is passed
### GetOsMaximumVersion

`func (o *AndroidCompliancePolicy) GetOsMaximumVersion() string`

GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.

### GetOsMaximumVersionOk

`func (o *AndroidCompliancePolicy) GetOsMaximumVersionOk() (string, bool)`

GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMaximumVersion

`func (o *AndroidCompliancePolicy) HasOsMaximumVersion() bool`

HasOsMaximumVersion returns a boolean if a field has been set.

### SetOsMaximumVersion

`func (o *AndroidCompliancePolicy) SetOsMaximumVersion(v string)`

SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.

### SetOsMaximumVersionExplicitNull

`func (o *AndroidCompliancePolicy) SetOsMaximumVersionExplicitNull(b bool)`

SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMaximumVersion value is set to nil even if false is passed
### GetMinAndroidSecurityPatchLevel

`func (o *AndroidCompliancePolicy) GetMinAndroidSecurityPatchLevel() string`

GetMinAndroidSecurityPatchLevel returns the MinAndroidSecurityPatchLevel field if non-nil, zero value otherwise.

### GetMinAndroidSecurityPatchLevelOk

`func (o *AndroidCompliancePolicy) GetMinAndroidSecurityPatchLevelOk() (string, bool)`

GetMinAndroidSecurityPatchLevelOk returns a tuple with the MinAndroidSecurityPatchLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinAndroidSecurityPatchLevel

`func (o *AndroidCompliancePolicy) HasMinAndroidSecurityPatchLevel() bool`

HasMinAndroidSecurityPatchLevel returns a boolean if a field has been set.

### SetMinAndroidSecurityPatchLevel

`func (o *AndroidCompliancePolicy) SetMinAndroidSecurityPatchLevel(v string)`

SetMinAndroidSecurityPatchLevel gets a reference to the given string and assigns it to the MinAndroidSecurityPatchLevel field.

### SetMinAndroidSecurityPatchLevelExplicitNull

`func (o *AndroidCompliancePolicy) SetMinAndroidSecurityPatchLevelExplicitNull(b bool)`

SetMinAndroidSecurityPatchLevelExplicitNull (un)sets MinAndroidSecurityPatchLevel to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinAndroidSecurityPatchLevel value is set to nil even if false is passed
### GetStorageRequireEncryption

`func (o *AndroidCompliancePolicy) GetStorageRequireEncryption() bool`

GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.

### GetStorageRequireEncryptionOk

`func (o *AndroidCompliancePolicy) GetStorageRequireEncryptionOk() (bool, bool)`

GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireEncryption

`func (o *AndroidCompliancePolicy) HasStorageRequireEncryption() bool`

HasStorageRequireEncryption returns a boolean if a field has been set.

### SetStorageRequireEncryption

`func (o *AndroidCompliancePolicy) SetStorageRequireEncryption(v bool)`

SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.

### GetSecurityRequireSafetyNetAttestationBasicIntegrity

`func (o *AndroidCompliancePolicy) GetSecurityRequireSafetyNetAttestationBasicIntegrity() bool`

GetSecurityRequireSafetyNetAttestationBasicIntegrity returns the SecurityRequireSafetyNetAttestationBasicIntegrity field if non-nil, zero value otherwise.

### GetSecurityRequireSafetyNetAttestationBasicIntegrityOk

`func (o *AndroidCompliancePolicy) GetSecurityRequireSafetyNetAttestationBasicIntegrityOk() (bool, bool)`

GetSecurityRequireSafetyNetAttestationBasicIntegrityOk returns a tuple with the SecurityRequireSafetyNetAttestationBasicIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireSafetyNetAttestationBasicIntegrity

`func (o *AndroidCompliancePolicy) HasSecurityRequireSafetyNetAttestationBasicIntegrity() bool`

HasSecurityRequireSafetyNetAttestationBasicIntegrity returns a boolean if a field has been set.

### SetSecurityRequireSafetyNetAttestationBasicIntegrity

`func (o *AndroidCompliancePolicy) SetSecurityRequireSafetyNetAttestationBasicIntegrity(v bool)`

SetSecurityRequireSafetyNetAttestationBasicIntegrity gets a reference to the given bool and assigns it to the SecurityRequireSafetyNetAttestationBasicIntegrity field.

### GetSecurityRequireSafetyNetAttestationCertifiedDevice

`func (o *AndroidCompliancePolicy) GetSecurityRequireSafetyNetAttestationCertifiedDevice() bool`

GetSecurityRequireSafetyNetAttestationCertifiedDevice returns the SecurityRequireSafetyNetAttestationCertifiedDevice field if non-nil, zero value otherwise.

### GetSecurityRequireSafetyNetAttestationCertifiedDeviceOk

`func (o *AndroidCompliancePolicy) GetSecurityRequireSafetyNetAttestationCertifiedDeviceOk() (bool, bool)`

GetSecurityRequireSafetyNetAttestationCertifiedDeviceOk returns a tuple with the SecurityRequireSafetyNetAttestationCertifiedDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireSafetyNetAttestationCertifiedDevice

`func (o *AndroidCompliancePolicy) HasSecurityRequireSafetyNetAttestationCertifiedDevice() bool`

HasSecurityRequireSafetyNetAttestationCertifiedDevice returns a boolean if a field has been set.

### SetSecurityRequireSafetyNetAttestationCertifiedDevice

`func (o *AndroidCompliancePolicy) SetSecurityRequireSafetyNetAttestationCertifiedDevice(v bool)`

SetSecurityRequireSafetyNetAttestationCertifiedDevice gets a reference to the given bool and assigns it to the SecurityRequireSafetyNetAttestationCertifiedDevice field.

### GetSecurityRequireGooglePlayServices

`func (o *AndroidCompliancePolicy) GetSecurityRequireGooglePlayServices() bool`

GetSecurityRequireGooglePlayServices returns the SecurityRequireGooglePlayServices field if non-nil, zero value otherwise.

### GetSecurityRequireGooglePlayServicesOk

`func (o *AndroidCompliancePolicy) GetSecurityRequireGooglePlayServicesOk() (bool, bool)`

GetSecurityRequireGooglePlayServicesOk returns a tuple with the SecurityRequireGooglePlayServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireGooglePlayServices

`func (o *AndroidCompliancePolicy) HasSecurityRequireGooglePlayServices() bool`

HasSecurityRequireGooglePlayServices returns a boolean if a field has been set.

### SetSecurityRequireGooglePlayServices

`func (o *AndroidCompliancePolicy) SetSecurityRequireGooglePlayServices(v bool)`

SetSecurityRequireGooglePlayServices gets a reference to the given bool and assigns it to the SecurityRequireGooglePlayServices field.

### GetSecurityRequireUpToDateSecurityProviders

`func (o *AndroidCompliancePolicy) GetSecurityRequireUpToDateSecurityProviders() bool`

GetSecurityRequireUpToDateSecurityProviders returns the SecurityRequireUpToDateSecurityProviders field if non-nil, zero value otherwise.

### GetSecurityRequireUpToDateSecurityProvidersOk

`func (o *AndroidCompliancePolicy) GetSecurityRequireUpToDateSecurityProvidersOk() (bool, bool)`

GetSecurityRequireUpToDateSecurityProvidersOk returns a tuple with the SecurityRequireUpToDateSecurityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireUpToDateSecurityProviders

`func (o *AndroidCompliancePolicy) HasSecurityRequireUpToDateSecurityProviders() bool`

HasSecurityRequireUpToDateSecurityProviders returns a boolean if a field has been set.

### SetSecurityRequireUpToDateSecurityProviders

`func (o *AndroidCompliancePolicy) SetSecurityRequireUpToDateSecurityProviders(v bool)`

SetSecurityRequireUpToDateSecurityProviders gets a reference to the given bool and assigns it to the SecurityRequireUpToDateSecurityProviders field.

### GetSecurityRequireCompanyPortalAppIntegrity

`func (o *AndroidCompliancePolicy) GetSecurityRequireCompanyPortalAppIntegrity() bool`

GetSecurityRequireCompanyPortalAppIntegrity returns the SecurityRequireCompanyPortalAppIntegrity field if non-nil, zero value otherwise.

### GetSecurityRequireCompanyPortalAppIntegrityOk

`func (o *AndroidCompliancePolicy) GetSecurityRequireCompanyPortalAppIntegrityOk() (bool, bool)`

GetSecurityRequireCompanyPortalAppIntegrityOk returns a tuple with the SecurityRequireCompanyPortalAppIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireCompanyPortalAppIntegrity

`func (o *AndroidCompliancePolicy) HasSecurityRequireCompanyPortalAppIntegrity() bool`

HasSecurityRequireCompanyPortalAppIntegrity returns a boolean if a field has been set.

### SetSecurityRequireCompanyPortalAppIntegrity

`func (o *AndroidCompliancePolicy) SetSecurityRequireCompanyPortalAppIntegrity(v bool)`

SetSecurityRequireCompanyPortalAppIntegrity gets a reference to the given bool and assigns it to the SecurityRequireCompanyPortalAppIntegrity field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


