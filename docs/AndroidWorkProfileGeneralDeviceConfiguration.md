# AndroidWorkProfileGeneralDeviceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordBlockFingerprintUnlock** | Pointer to **bool** | Indicates whether or not to block fingerprint unlock. | [optional] 
**PasswordBlockTrustAgents** | Pointer to **bool** | Indicates whether or not to block Smart Lock and other trust agents. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | Number of days before the password expires. Valid values 1 to 365 | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | Minimum length of passwords. Valid values 4 to 16 | [optional] 
**PasswordMinutesOfInactivityBeforeScreenTimeout** | Pointer to **int32** | Minutes of inactivity before the screen times out. | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | Number of previous passwords to block. Valid values 0 to 24 | [optional] 
**PasswordSignInFailureCountBeforeFactoryReset** | Pointer to **int32** | Number of sign in failures allowed before factory reset. Valid values 1 to 16 | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType**](anyOf&lt;microsoft.graph.androidWorkProfileRequiredPasswordType&gt;.md) | Type of password that is required. | [optional] 
**WorkProfileDataSharingType** | Pointer to [**AnyOfmicrosoftGraphAndroidWorkProfileCrossProfileDataSharingType**](anyOf&lt;microsoft.graph.androidWorkProfileCrossProfileDataSharingType&gt;.md) | Type of data sharing that is allowed. | [optional] 
**WorkProfileBlockNotificationsWhileDeviceLocked** | Pointer to **bool** | Indicates whether or not to block notifications while device locked. | [optional] 
**WorkProfileBlockAddingAccounts** | Pointer to **bool** | Block users from adding/removing accounts in work profile. | [optional] 
**WorkProfileBluetoothEnableContactSharing** | Pointer to **bool** | Allow bluetooth devices to access enterprise contacts. | [optional] 
**WorkProfileBlockScreenCapture** | Pointer to **bool** | Block screen capture in work profile. | [optional] 
**WorkProfileBlockCrossProfileCallerId** | Pointer to **bool** | Block display work profile caller ID in personal profile. | [optional] 
**WorkProfileBlockCamera** | Pointer to **bool** | Block work profile camera. | [optional] 
**WorkProfileBlockCrossProfileContactsSearch** | Pointer to **bool** | Block work profile contacts availability in personal profile. | [optional] 
**WorkProfileBlockCrossProfileCopyPaste** | Pointer to **bool** | Boolean that indicates if the setting disallow cross profile copy/paste is enabled. | [optional] 
**WorkProfileDefaultAppPermissionPolicy** | Pointer to [**AnyOfmicrosoftGraphAndroidWorkProfileDefaultAppPermissionPolicyType**](anyOf&lt;microsoft.graph.androidWorkProfileDefaultAppPermissionPolicyType&gt;.md) | Type of password that is required. | [optional] 
**WorkProfilePasswordBlockFingerprintUnlock** | Pointer to **bool** | Indicates whether or not to block fingerprint unlock for work profile. | [optional] 
**WorkProfilePasswordBlockTrustAgents** | Pointer to **bool** | Indicates whether or not to block Smart Lock and other trust agents for work profile. | [optional] 
**WorkProfilePasswordExpirationDays** | Pointer to **int32** | Number of days before the work profile password expires. Valid values 1 to 365 | [optional] 
**WorkProfilePasswordMinimumLength** | Pointer to **int32** | Minimum length of work profile password. Valid values 4 to 16 | [optional] 
**WorkProfilePasswordMinNumericCharacters** | Pointer to **int32** | Minimum # of numeric characters required in work profile password. Valid values 1 to 10 | [optional] 
**WorkProfilePasswordMinNonLetterCharacters** | Pointer to **int32** | Minimum # of non-letter characters required in work profile password. Valid values 1 to 10 | [optional] 
**WorkProfilePasswordMinLetterCharacters** | Pointer to **int32** | Minimum # of letter characters required in work profile password. Valid values 1 to 10 | [optional] 
**WorkProfilePasswordMinLowerCaseCharacters** | Pointer to **int32** | Minimum # of lower-case characters required in work profile password. Valid values 1 to 10 | [optional] 
**WorkProfilePasswordMinUpperCaseCharacters** | Pointer to **int32** | Minimum # of upper-case characters required in work profile password. Valid values 1 to 10 | [optional] 
**WorkProfilePasswordMinSymbolCharacters** | Pointer to **int32** | Minimum # of symbols required in work profile password. Valid values 1 to 10 | [optional] 
**WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout** | Pointer to **int32** | Minutes of inactivity before the screen times out. | [optional] 
**WorkProfilePasswordPreviousPasswordBlockCount** | Pointer to **int32** | Number of previous work profile passwords to block. Valid values 0 to 24 | [optional] 
**WorkProfilePasswordSignInFailureCountBeforeFactoryReset** | Pointer to **int32** | Number of sign in failures allowed before work profile is removed and all corporate data deleted. Valid values 1 to 16 | [optional] 
**WorkProfilePasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType**](anyOf&lt;microsoft.graph.androidWorkProfileRequiredPasswordType&gt;.md) | Type of work profile password that is required. | [optional] 
**WorkProfileRequirePassword** | Pointer to **bool** | Password is required or not for work profile | [optional] 
**SecurityRequireVerifyApps** | Pointer to **bool** | Require the Android Verify apps feature is turned on. | [optional] 

## Methods

### GetPasswordBlockFingerprintUnlock

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlock() bool`

GetPasswordBlockFingerprintUnlock returns the PasswordBlockFingerprintUnlock field if non-nil, zero value otherwise.

### GetPasswordBlockFingerprintUnlockOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlockOk() (bool, bool)`

GetPasswordBlockFingerprintUnlockOk returns a tuple with the PasswordBlockFingerprintUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockFingerprintUnlock

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasPasswordBlockFingerprintUnlock() bool`

HasPasswordBlockFingerprintUnlock returns a boolean if a field has been set.

### SetPasswordBlockFingerprintUnlock

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordBlockFingerprintUnlock(v bool)`

SetPasswordBlockFingerprintUnlock gets a reference to the given bool and assigns it to the PasswordBlockFingerprintUnlock field.

### GetPasswordBlockTrustAgents

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockTrustAgents() bool`

GetPasswordBlockTrustAgents returns the PasswordBlockTrustAgents field if non-nil, zero value otherwise.

### GetPasswordBlockTrustAgentsOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordBlockTrustAgentsOk() (bool, bool)`

GetPasswordBlockTrustAgentsOk returns a tuple with the PasswordBlockTrustAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockTrustAgents

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasPasswordBlockTrustAgents() bool`

HasPasswordBlockTrustAgents returns a boolean if a field has been set.

### SetPasswordBlockTrustAgents

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordBlockTrustAgents(v bool)`

SetPasswordBlockTrustAgents gets a reference to the given bool and assigns it to the PasswordBlockTrustAgents field.

### GetPasswordExpirationDays

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasswordMinutesOfInactivityBeforeScreenTimeout returns the PasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasPasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordSignInFailureCountBeforeFactoryReset

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset() int32`

GetPasswordSignInFailureCountBeforeFactoryReset returns the PasswordSignInFailureCountBeforeFactoryReset field if non-nil, zero value otherwise.

### GetPasswordSignInFailureCountBeforeFactoryResetOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryResetOk() (int32, bool)`

GetPasswordSignInFailureCountBeforeFactoryResetOk returns a tuple with the PasswordSignInFailureCountBeforeFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordSignInFailureCountBeforeFactoryReset

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasPasswordSignInFailureCountBeforeFactoryReset() bool`

HasPasswordSignInFailureCountBeforeFactoryReset returns a boolean if a field has been set.

### SetPasswordSignInFailureCountBeforeFactoryReset

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(v int32)`

SetPasswordSignInFailureCountBeforeFactoryReset gets a reference to the given int32 and assigns it to the PasswordSignInFailureCountBeforeFactoryReset field.

### SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull(b bool)`

SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull (un)sets PasswordSignInFailureCountBeforeFactoryReset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordSignInFailureCountBeforeFactoryReset value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordRequiredType() AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetPasswordRequiredType(v AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetWorkProfileDataSharingType

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileDataSharingType() AnyOfmicrosoftGraphAndroidWorkProfileCrossProfileDataSharingType`

GetWorkProfileDataSharingType returns the WorkProfileDataSharingType field if non-nil, zero value otherwise.

### GetWorkProfileDataSharingTypeOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileDataSharingTypeOk() (AnyOfmicrosoftGraphAndroidWorkProfileCrossProfileDataSharingType, bool)`

GetWorkProfileDataSharingTypeOk returns a tuple with the WorkProfileDataSharingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileDataSharingType

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileDataSharingType() bool`

HasWorkProfileDataSharingType returns a boolean if a field has been set.

### SetWorkProfileDataSharingType

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileDataSharingType(v AnyOfmicrosoftGraphAndroidWorkProfileCrossProfileDataSharingType)`

SetWorkProfileDataSharingType gets a reference to the given AnyOfmicrosoftGraphAndroidWorkProfileCrossProfileDataSharingType and assigns it to the WorkProfileDataSharingType field.

### GetWorkProfileBlockNotificationsWhileDeviceLocked

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockNotificationsWhileDeviceLocked() bool`

GetWorkProfileBlockNotificationsWhileDeviceLocked returns the WorkProfileBlockNotificationsWhileDeviceLocked field if non-nil, zero value otherwise.

### GetWorkProfileBlockNotificationsWhileDeviceLockedOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockNotificationsWhileDeviceLockedOk() (bool, bool)`

GetWorkProfileBlockNotificationsWhileDeviceLockedOk returns a tuple with the WorkProfileBlockNotificationsWhileDeviceLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockNotificationsWhileDeviceLocked

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockNotificationsWhileDeviceLocked() bool`

HasWorkProfileBlockNotificationsWhileDeviceLocked returns a boolean if a field has been set.

### SetWorkProfileBlockNotificationsWhileDeviceLocked

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockNotificationsWhileDeviceLocked(v bool)`

SetWorkProfileBlockNotificationsWhileDeviceLocked gets a reference to the given bool and assigns it to the WorkProfileBlockNotificationsWhileDeviceLocked field.

### GetWorkProfileBlockAddingAccounts

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockAddingAccounts() bool`

GetWorkProfileBlockAddingAccounts returns the WorkProfileBlockAddingAccounts field if non-nil, zero value otherwise.

### GetWorkProfileBlockAddingAccountsOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockAddingAccountsOk() (bool, bool)`

GetWorkProfileBlockAddingAccountsOk returns a tuple with the WorkProfileBlockAddingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockAddingAccounts

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockAddingAccounts() bool`

HasWorkProfileBlockAddingAccounts returns a boolean if a field has been set.

### SetWorkProfileBlockAddingAccounts

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockAddingAccounts(v bool)`

SetWorkProfileBlockAddingAccounts gets a reference to the given bool and assigns it to the WorkProfileBlockAddingAccounts field.

### GetWorkProfileBluetoothEnableContactSharing

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBluetoothEnableContactSharing() bool`

GetWorkProfileBluetoothEnableContactSharing returns the WorkProfileBluetoothEnableContactSharing field if non-nil, zero value otherwise.

### GetWorkProfileBluetoothEnableContactSharingOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBluetoothEnableContactSharingOk() (bool, bool)`

GetWorkProfileBluetoothEnableContactSharingOk returns a tuple with the WorkProfileBluetoothEnableContactSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBluetoothEnableContactSharing

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBluetoothEnableContactSharing() bool`

HasWorkProfileBluetoothEnableContactSharing returns a boolean if a field has been set.

### SetWorkProfileBluetoothEnableContactSharing

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBluetoothEnableContactSharing(v bool)`

SetWorkProfileBluetoothEnableContactSharing gets a reference to the given bool and assigns it to the WorkProfileBluetoothEnableContactSharing field.

### GetWorkProfileBlockScreenCapture

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockScreenCapture() bool`

GetWorkProfileBlockScreenCapture returns the WorkProfileBlockScreenCapture field if non-nil, zero value otherwise.

### GetWorkProfileBlockScreenCaptureOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockScreenCaptureOk() (bool, bool)`

GetWorkProfileBlockScreenCaptureOk returns a tuple with the WorkProfileBlockScreenCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockScreenCapture

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockScreenCapture() bool`

HasWorkProfileBlockScreenCapture returns a boolean if a field has been set.

### SetWorkProfileBlockScreenCapture

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockScreenCapture(v bool)`

SetWorkProfileBlockScreenCapture gets a reference to the given bool and assigns it to the WorkProfileBlockScreenCapture field.

### GetWorkProfileBlockCrossProfileCallerId

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCallerId() bool`

GetWorkProfileBlockCrossProfileCallerId returns the WorkProfileBlockCrossProfileCallerId field if non-nil, zero value otherwise.

### GetWorkProfileBlockCrossProfileCallerIdOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCallerIdOk() (bool, bool)`

GetWorkProfileBlockCrossProfileCallerIdOk returns a tuple with the WorkProfileBlockCrossProfileCallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockCrossProfileCallerId

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockCrossProfileCallerId() bool`

HasWorkProfileBlockCrossProfileCallerId returns a boolean if a field has been set.

### SetWorkProfileBlockCrossProfileCallerId

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileCallerId(v bool)`

SetWorkProfileBlockCrossProfileCallerId gets a reference to the given bool and assigns it to the WorkProfileBlockCrossProfileCallerId field.

### GetWorkProfileBlockCamera

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCamera() bool`

GetWorkProfileBlockCamera returns the WorkProfileBlockCamera field if non-nil, zero value otherwise.

### GetWorkProfileBlockCameraOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCameraOk() (bool, bool)`

GetWorkProfileBlockCameraOk returns a tuple with the WorkProfileBlockCamera field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockCamera

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockCamera() bool`

HasWorkProfileBlockCamera returns a boolean if a field has been set.

### SetWorkProfileBlockCamera

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCamera(v bool)`

SetWorkProfileBlockCamera gets a reference to the given bool and assigns it to the WorkProfileBlockCamera field.

### GetWorkProfileBlockCrossProfileContactsSearch

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileContactsSearch() bool`

GetWorkProfileBlockCrossProfileContactsSearch returns the WorkProfileBlockCrossProfileContactsSearch field if non-nil, zero value otherwise.

### GetWorkProfileBlockCrossProfileContactsSearchOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileContactsSearchOk() (bool, bool)`

GetWorkProfileBlockCrossProfileContactsSearchOk returns a tuple with the WorkProfileBlockCrossProfileContactsSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockCrossProfileContactsSearch

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockCrossProfileContactsSearch() bool`

HasWorkProfileBlockCrossProfileContactsSearch returns a boolean if a field has been set.

### SetWorkProfileBlockCrossProfileContactsSearch

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileContactsSearch(v bool)`

SetWorkProfileBlockCrossProfileContactsSearch gets a reference to the given bool and assigns it to the WorkProfileBlockCrossProfileContactsSearch field.

### GetWorkProfileBlockCrossProfileCopyPaste

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCopyPaste() bool`

GetWorkProfileBlockCrossProfileCopyPaste returns the WorkProfileBlockCrossProfileCopyPaste field if non-nil, zero value otherwise.

### GetWorkProfileBlockCrossProfileCopyPasteOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileBlockCrossProfileCopyPasteOk() (bool, bool)`

GetWorkProfileBlockCrossProfileCopyPasteOk returns a tuple with the WorkProfileBlockCrossProfileCopyPaste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileBlockCrossProfileCopyPaste

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileBlockCrossProfileCopyPaste() bool`

HasWorkProfileBlockCrossProfileCopyPaste returns a boolean if a field has been set.

### SetWorkProfileBlockCrossProfileCopyPaste

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileBlockCrossProfileCopyPaste(v bool)`

SetWorkProfileBlockCrossProfileCopyPaste gets a reference to the given bool and assigns it to the WorkProfileBlockCrossProfileCopyPaste field.

### GetWorkProfileDefaultAppPermissionPolicy

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileDefaultAppPermissionPolicy() AnyOfmicrosoftGraphAndroidWorkProfileDefaultAppPermissionPolicyType`

GetWorkProfileDefaultAppPermissionPolicy returns the WorkProfileDefaultAppPermissionPolicy field if non-nil, zero value otherwise.

### GetWorkProfileDefaultAppPermissionPolicyOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileDefaultAppPermissionPolicyOk() (AnyOfmicrosoftGraphAndroidWorkProfileDefaultAppPermissionPolicyType, bool)`

GetWorkProfileDefaultAppPermissionPolicyOk returns a tuple with the WorkProfileDefaultAppPermissionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileDefaultAppPermissionPolicy

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileDefaultAppPermissionPolicy() bool`

HasWorkProfileDefaultAppPermissionPolicy returns a boolean if a field has been set.

### SetWorkProfileDefaultAppPermissionPolicy

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileDefaultAppPermissionPolicy(v AnyOfmicrosoftGraphAndroidWorkProfileDefaultAppPermissionPolicyType)`

SetWorkProfileDefaultAppPermissionPolicy gets a reference to the given AnyOfmicrosoftGraphAndroidWorkProfileDefaultAppPermissionPolicyType and assigns it to the WorkProfileDefaultAppPermissionPolicy field.

### GetWorkProfilePasswordBlockFingerprintUnlock

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockFingerprintUnlock() bool`

GetWorkProfilePasswordBlockFingerprintUnlock returns the WorkProfilePasswordBlockFingerprintUnlock field if non-nil, zero value otherwise.

### GetWorkProfilePasswordBlockFingerprintUnlockOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockFingerprintUnlockOk() (bool, bool)`

GetWorkProfilePasswordBlockFingerprintUnlockOk returns a tuple with the WorkProfilePasswordBlockFingerprintUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordBlockFingerprintUnlock

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordBlockFingerprintUnlock() bool`

HasWorkProfilePasswordBlockFingerprintUnlock returns a boolean if a field has been set.

### SetWorkProfilePasswordBlockFingerprintUnlock

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordBlockFingerprintUnlock(v bool)`

SetWorkProfilePasswordBlockFingerprintUnlock gets a reference to the given bool and assigns it to the WorkProfilePasswordBlockFingerprintUnlock field.

### GetWorkProfilePasswordBlockTrustAgents

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockTrustAgents() bool`

GetWorkProfilePasswordBlockTrustAgents returns the WorkProfilePasswordBlockTrustAgents field if non-nil, zero value otherwise.

### GetWorkProfilePasswordBlockTrustAgentsOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordBlockTrustAgentsOk() (bool, bool)`

GetWorkProfilePasswordBlockTrustAgentsOk returns a tuple with the WorkProfilePasswordBlockTrustAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordBlockTrustAgents

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordBlockTrustAgents() bool`

HasWorkProfilePasswordBlockTrustAgents returns a boolean if a field has been set.

### SetWorkProfilePasswordBlockTrustAgents

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordBlockTrustAgents(v bool)`

SetWorkProfilePasswordBlockTrustAgents gets a reference to the given bool and assigns it to the WorkProfilePasswordBlockTrustAgents field.

### GetWorkProfilePasswordExpirationDays

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordExpirationDays() int32`

GetWorkProfilePasswordExpirationDays returns the WorkProfilePasswordExpirationDays field if non-nil, zero value otherwise.

### GetWorkProfilePasswordExpirationDaysOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordExpirationDaysOk() (int32, bool)`

GetWorkProfilePasswordExpirationDaysOk returns a tuple with the WorkProfilePasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordExpirationDays

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordExpirationDays() bool`

HasWorkProfilePasswordExpirationDays returns a boolean if a field has been set.

### SetWorkProfilePasswordExpirationDays

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordExpirationDays(v int32)`

SetWorkProfilePasswordExpirationDays gets a reference to the given int32 and assigns it to the WorkProfilePasswordExpirationDays field.

### SetWorkProfilePasswordExpirationDaysExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordExpirationDaysExplicitNull(b bool)`

SetWorkProfilePasswordExpirationDaysExplicitNull (un)sets WorkProfilePasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordExpirationDays value is set to nil even if false is passed
### GetWorkProfilePasswordMinimumLength

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLength() int32`

GetWorkProfilePasswordMinimumLength returns the WorkProfilePasswordMinimumLength field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinimumLengthOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLengthOk() (int32, bool)`

GetWorkProfilePasswordMinimumLengthOk returns a tuple with the WorkProfilePasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinimumLength

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinimumLength() bool`

HasWorkProfilePasswordMinimumLength returns a boolean if a field has been set.

### SetWorkProfilePasswordMinimumLength

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLength(v int32)`

SetWorkProfilePasswordMinimumLength gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinimumLength field.

### SetWorkProfilePasswordMinimumLengthExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLengthExplicitNull(b bool)`

SetWorkProfilePasswordMinimumLengthExplicitNull (un)sets WorkProfilePasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinimumLength value is set to nil even if false is passed
### GetWorkProfilePasswordMinNumericCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinNumericCharacters() int32`

GetWorkProfilePasswordMinNumericCharacters returns the WorkProfilePasswordMinNumericCharacters field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinNumericCharactersOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinNumericCharactersOk() (int32, bool)`

GetWorkProfilePasswordMinNumericCharactersOk returns a tuple with the WorkProfilePasswordMinNumericCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinNumericCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinNumericCharacters() bool`

HasWorkProfilePasswordMinNumericCharacters returns a boolean if a field has been set.

### SetWorkProfilePasswordMinNumericCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinNumericCharacters(v int32)`

SetWorkProfilePasswordMinNumericCharacters gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinNumericCharacters field.

### SetWorkProfilePasswordMinNumericCharactersExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinNumericCharactersExplicitNull(b bool)`

SetWorkProfilePasswordMinNumericCharactersExplicitNull (un)sets WorkProfilePasswordMinNumericCharacters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinNumericCharacters value is set to nil even if false is passed
### GetWorkProfilePasswordMinNonLetterCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinNonLetterCharacters() int32`

GetWorkProfilePasswordMinNonLetterCharacters returns the WorkProfilePasswordMinNonLetterCharacters field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinNonLetterCharactersOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinNonLetterCharactersOk() (int32, bool)`

GetWorkProfilePasswordMinNonLetterCharactersOk returns a tuple with the WorkProfilePasswordMinNonLetterCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinNonLetterCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinNonLetterCharacters() bool`

HasWorkProfilePasswordMinNonLetterCharacters returns a boolean if a field has been set.

### SetWorkProfilePasswordMinNonLetterCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinNonLetterCharacters(v int32)`

SetWorkProfilePasswordMinNonLetterCharacters gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinNonLetterCharacters field.

### SetWorkProfilePasswordMinNonLetterCharactersExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinNonLetterCharactersExplicitNull(b bool)`

SetWorkProfilePasswordMinNonLetterCharactersExplicitNull (un)sets WorkProfilePasswordMinNonLetterCharacters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinNonLetterCharacters value is set to nil even if false is passed
### GetWorkProfilePasswordMinLetterCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinLetterCharacters() int32`

GetWorkProfilePasswordMinLetterCharacters returns the WorkProfilePasswordMinLetterCharacters field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinLetterCharactersOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinLetterCharactersOk() (int32, bool)`

GetWorkProfilePasswordMinLetterCharactersOk returns a tuple with the WorkProfilePasswordMinLetterCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinLetterCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinLetterCharacters() bool`

HasWorkProfilePasswordMinLetterCharacters returns a boolean if a field has been set.

### SetWorkProfilePasswordMinLetterCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinLetterCharacters(v int32)`

SetWorkProfilePasswordMinLetterCharacters gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinLetterCharacters field.

### SetWorkProfilePasswordMinLetterCharactersExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinLetterCharactersExplicitNull(b bool)`

SetWorkProfilePasswordMinLetterCharactersExplicitNull (un)sets WorkProfilePasswordMinLetterCharacters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinLetterCharacters value is set to nil even if false is passed
### GetWorkProfilePasswordMinLowerCaseCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinLowerCaseCharacters() int32`

GetWorkProfilePasswordMinLowerCaseCharacters returns the WorkProfilePasswordMinLowerCaseCharacters field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinLowerCaseCharactersOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinLowerCaseCharactersOk() (int32, bool)`

GetWorkProfilePasswordMinLowerCaseCharactersOk returns a tuple with the WorkProfilePasswordMinLowerCaseCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinLowerCaseCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinLowerCaseCharacters() bool`

HasWorkProfilePasswordMinLowerCaseCharacters returns a boolean if a field has been set.

### SetWorkProfilePasswordMinLowerCaseCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinLowerCaseCharacters(v int32)`

SetWorkProfilePasswordMinLowerCaseCharacters gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinLowerCaseCharacters field.

### SetWorkProfilePasswordMinLowerCaseCharactersExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinLowerCaseCharactersExplicitNull(b bool)`

SetWorkProfilePasswordMinLowerCaseCharactersExplicitNull (un)sets WorkProfilePasswordMinLowerCaseCharacters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinLowerCaseCharacters value is set to nil even if false is passed
### GetWorkProfilePasswordMinUpperCaseCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinUpperCaseCharacters() int32`

GetWorkProfilePasswordMinUpperCaseCharacters returns the WorkProfilePasswordMinUpperCaseCharacters field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinUpperCaseCharactersOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinUpperCaseCharactersOk() (int32, bool)`

GetWorkProfilePasswordMinUpperCaseCharactersOk returns a tuple with the WorkProfilePasswordMinUpperCaseCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinUpperCaseCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinUpperCaseCharacters() bool`

HasWorkProfilePasswordMinUpperCaseCharacters returns a boolean if a field has been set.

### SetWorkProfilePasswordMinUpperCaseCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinUpperCaseCharacters(v int32)`

SetWorkProfilePasswordMinUpperCaseCharacters gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinUpperCaseCharacters field.

### SetWorkProfilePasswordMinUpperCaseCharactersExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinUpperCaseCharactersExplicitNull(b bool)`

SetWorkProfilePasswordMinUpperCaseCharactersExplicitNull (un)sets WorkProfilePasswordMinUpperCaseCharacters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinUpperCaseCharacters value is set to nil even if false is passed
### GetWorkProfilePasswordMinSymbolCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinSymbolCharacters() int32`

GetWorkProfilePasswordMinSymbolCharacters returns the WorkProfilePasswordMinSymbolCharacters field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinSymbolCharactersOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinSymbolCharactersOk() (int32, bool)`

GetWorkProfilePasswordMinSymbolCharactersOk returns a tuple with the WorkProfilePasswordMinSymbolCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinSymbolCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinSymbolCharacters() bool`

HasWorkProfilePasswordMinSymbolCharacters returns a boolean if a field has been set.

### SetWorkProfilePasswordMinSymbolCharacters

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinSymbolCharacters(v int32)`

SetWorkProfilePasswordMinSymbolCharacters gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinSymbolCharacters field.

### SetWorkProfilePasswordMinSymbolCharactersExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinSymbolCharactersExplicitNull(b bool)`

SetWorkProfilePasswordMinSymbolCharactersExplicitNull (un)sets WorkProfilePasswordMinSymbolCharacters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinSymbolCharacters value is set to nil even if false is passed
### GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout returns the WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetWorkProfilePasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetWorkProfilePasswordPreviousPasswordBlockCount

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordPreviousPasswordBlockCount() int32`

GetWorkProfilePasswordPreviousPasswordBlockCount returns the WorkProfilePasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetWorkProfilePasswordPreviousPasswordBlockCountOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetWorkProfilePasswordPreviousPasswordBlockCountOk returns a tuple with the WorkProfilePasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordPreviousPasswordBlockCount

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordPreviousPasswordBlockCount() bool`

HasWorkProfilePasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetWorkProfilePasswordPreviousPasswordBlockCount

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordPreviousPasswordBlockCount(v int32)`

SetWorkProfilePasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the WorkProfilePasswordPreviousPasswordBlockCount field.

### SetWorkProfilePasswordPreviousPasswordBlockCountExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetWorkProfilePasswordPreviousPasswordBlockCountExplicitNull (un)sets WorkProfilePasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset() int32`

GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset returns the WorkProfilePasswordSignInFailureCountBeforeFactoryReset field if non-nil, zero value otherwise.

### GetWorkProfilePasswordSignInFailureCountBeforeFactoryResetOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordSignInFailureCountBeforeFactoryResetOk() (int32, bool)`

GetWorkProfilePasswordSignInFailureCountBeforeFactoryResetOk returns a tuple with the WorkProfilePasswordSignInFailureCountBeforeFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordSignInFailureCountBeforeFactoryReset

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordSignInFailureCountBeforeFactoryReset() bool`

HasWorkProfilePasswordSignInFailureCountBeforeFactoryReset returns a boolean if a field has been set.

### SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(v int32)`

SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset gets a reference to the given int32 and assigns it to the WorkProfilePasswordSignInFailureCountBeforeFactoryReset field.

### SetWorkProfilePasswordSignInFailureCountBeforeFactoryResetExplicitNull

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordSignInFailureCountBeforeFactoryResetExplicitNull(b bool)`

SetWorkProfilePasswordSignInFailureCountBeforeFactoryResetExplicitNull (un)sets WorkProfilePasswordSignInFailureCountBeforeFactoryReset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkProfilePasswordSignInFailureCountBeforeFactoryReset value is set to nil even if false is passed
### GetWorkProfilePasswordRequiredType

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordRequiredType() AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType`

GetWorkProfilePasswordRequiredType returns the WorkProfilePasswordRequiredType field if non-nil, zero value otherwise.

### GetWorkProfilePasswordRequiredTypeOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfilePasswordRequiredTypeOk() (AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType, bool)`

GetWorkProfilePasswordRequiredTypeOk returns a tuple with the WorkProfilePasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfilePasswordRequiredType

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfilePasswordRequiredType() bool`

HasWorkProfilePasswordRequiredType returns a boolean if a field has been set.

### SetWorkProfilePasswordRequiredType

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfilePasswordRequiredType(v AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType)`

SetWorkProfilePasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphAndroidWorkProfileRequiredPasswordType and assigns it to the WorkProfilePasswordRequiredType field.

### GetWorkProfileRequirePassword

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileRequirePassword() bool`

GetWorkProfileRequirePassword returns the WorkProfileRequirePassword field if non-nil, zero value otherwise.

### GetWorkProfileRequirePasswordOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetWorkProfileRequirePasswordOk() (bool, bool)`

GetWorkProfileRequirePasswordOk returns a tuple with the WorkProfileRequirePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkProfileRequirePassword

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasWorkProfileRequirePassword() bool`

HasWorkProfileRequirePassword returns a boolean if a field has been set.

### SetWorkProfileRequirePassword

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetWorkProfileRequirePassword(v bool)`

SetWorkProfileRequirePassword gets a reference to the given bool and assigns it to the WorkProfileRequirePassword field.

### GetSecurityRequireVerifyApps

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetSecurityRequireVerifyApps() bool`

GetSecurityRequireVerifyApps returns the SecurityRequireVerifyApps field if non-nil, zero value otherwise.

### GetSecurityRequireVerifyAppsOk

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) GetSecurityRequireVerifyAppsOk() (bool, bool)`

GetSecurityRequireVerifyAppsOk returns a tuple with the SecurityRequireVerifyApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityRequireVerifyApps

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) HasSecurityRequireVerifyApps() bool`

HasSecurityRequireVerifyApps returns a boolean if a field has been set.

### SetSecurityRequireVerifyApps

`func (o *AndroidWorkProfileGeneralDeviceConfiguration) SetSecurityRequireVerifyApps(v bool)`

SetSecurityRequireVerifyApps gets a reference to the given bool and assigns it to the SecurityRequireVerifyApps field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


