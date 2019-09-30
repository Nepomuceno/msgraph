# WindowsInformationProtectionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevokeOnMdmHandoffDisabled** | Pointer to **bool** | New property in RS2, pending documentation | [optional] 
**MdmEnrollmentUrl** | Pointer to **string** | Enrollment url for the MDM | [optional] 
**WindowsHelloForBusinessBlocked** | Pointer to **bool** | Boolean value that sets Windows Hello for Business as a method for signing into Windows. | [optional] 
**PinMinimumLength** | Pointer to **int32** | Integer value that sets the minimum number of characters required for the PIN. Default value is 4. The lowest number you can configure for this policy setting is 4. The largest number you can configure must be less than the number configured in the Maximum PIN length policy setting or the number 127, whichever is the lowest. | [optional] 
**PinUppercaseLetters** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements**](anyOf&lt;microsoft.graph.windowsInformationProtectionPinCharacterRequirements&gt;.md) | Integer value that configures the use of uppercase letters in the Windows Hello for Business PIN. Default is NotAllow. | [optional] 
**PinLowercaseLetters** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements**](anyOf&lt;microsoft.graph.windowsInformationProtectionPinCharacterRequirements&gt;.md) | Integer value that configures the use of lowercase letters in the Windows Hello for Business PIN. Default is NotAllow. | [optional] 
**PinSpecialCharacters** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements**](anyOf&lt;microsoft.graph.windowsInformationProtectionPinCharacterRequirements&gt;.md) | Integer value that configures the use of special characters in the Windows Hello for Business PIN. Valid special characters for Windows Hello for Business PIN gestures include: ! \&quot; # $ % &amp; &#39; ( ) * + , - . / : ; &lt; &#x3D; &gt; ? @ [ \\ ] ^ _ &#x60; { | } ~. Default is NotAllow. | [optional] 
**PinExpirationDays** | Pointer to **int32** | Integer value specifies the period of time (in days) that a PIN can be used before the system requires the user to change it. The largest number you can configure for this policy setting is 730. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then the user&#39;s PIN will never expire. This node was added in Windows 10, version 1511. Default is 0. | [optional] 
**NumberOfPastPinsRemembered** | Pointer to **int32** | Integer value that specifies the number of past PINs that can be associated to a user account that can&#39;t be reused. The largest number you can configure for this policy setting is 50. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then storage of previous PINs is not required. This node was added in Windows 10, version 1511. Default is 0. | [optional] 
**PasswordMaximumAttemptCount** | Pointer to **int32** | The number of authentication failures allowed before the device will be wiped. A value of 0 disables device wipe functionality. Range is an integer X where 4 &lt;&#x3D; X &lt;&#x3D; 16 for desktop and 0 &lt;&#x3D; X &lt;&#x3D; 999 for mobile devices. | [optional] 
**MinutesOfInactivityBeforeDeviceLock** | Pointer to **int32** | Specifies the maximum amount of time (in minutes) allowed after the device is idle that will cause the device to become PIN or password locked.   Range is an integer X where 0 &lt;&#x3D; X &lt;&#x3D; 999. | [optional] 
**DaysWithoutContactBeforeUnenroll** | Pointer to **int32** | Offline interval before app data is wiped (days)  | [optional] 

## Methods

### GetRevokeOnMdmHandoffDisabled

`func (o *WindowsInformationProtectionPolicy) GetRevokeOnMdmHandoffDisabled() bool`

GetRevokeOnMdmHandoffDisabled returns the RevokeOnMdmHandoffDisabled field if non-nil, zero value otherwise.

### GetRevokeOnMdmHandoffDisabledOk

`func (o *WindowsInformationProtectionPolicy) GetRevokeOnMdmHandoffDisabledOk() (bool, bool)`

GetRevokeOnMdmHandoffDisabledOk returns a tuple with the RevokeOnMdmHandoffDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRevokeOnMdmHandoffDisabled

`func (o *WindowsInformationProtectionPolicy) HasRevokeOnMdmHandoffDisabled() bool`

HasRevokeOnMdmHandoffDisabled returns a boolean if a field has been set.

### SetRevokeOnMdmHandoffDisabled

`func (o *WindowsInformationProtectionPolicy) SetRevokeOnMdmHandoffDisabled(v bool)`

SetRevokeOnMdmHandoffDisabled gets a reference to the given bool and assigns it to the RevokeOnMdmHandoffDisabled field.

### GetMdmEnrollmentUrl

`func (o *WindowsInformationProtectionPolicy) GetMdmEnrollmentUrl() string`

GetMdmEnrollmentUrl returns the MdmEnrollmentUrl field if non-nil, zero value otherwise.

### GetMdmEnrollmentUrlOk

`func (o *WindowsInformationProtectionPolicy) GetMdmEnrollmentUrlOk() (string, bool)`

GetMdmEnrollmentUrlOk returns a tuple with the MdmEnrollmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMdmEnrollmentUrl

`func (o *WindowsInformationProtectionPolicy) HasMdmEnrollmentUrl() bool`

HasMdmEnrollmentUrl returns a boolean if a field has been set.

### SetMdmEnrollmentUrl

`func (o *WindowsInformationProtectionPolicy) SetMdmEnrollmentUrl(v string)`

SetMdmEnrollmentUrl gets a reference to the given string and assigns it to the MdmEnrollmentUrl field.

### SetMdmEnrollmentUrlExplicitNull

`func (o *WindowsInformationProtectionPolicy) SetMdmEnrollmentUrlExplicitNull(b bool)`

SetMdmEnrollmentUrlExplicitNull (un)sets MdmEnrollmentUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MdmEnrollmentUrl value is set to nil even if false is passed
### GetWindowsHelloForBusinessBlocked

`func (o *WindowsInformationProtectionPolicy) GetWindowsHelloForBusinessBlocked() bool`

GetWindowsHelloForBusinessBlocked returns the WindowsHelloForBusinessBlocked field if non-nil, zero value otherwise.

### GetWindowsHelloForBusinessBlockedOk

`func (o *WindowsInformationProtectionPolicy) GetWindowsHelloForBusinessBlockedOk() (bool, bool)`

GetWindowsHelloForBusinessBlockedOk returns a tuple with the WindowsHelloForBusinessBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsHelloForBusinessBlocked

`func (o *WindowsInformationProtectionPolicy) HasWindowsHelloForBusinessBlocked() bool`

HasWindowsHelloForBusinessBlocked returns a boolean if a field has been set.

### SetWindowsHelloForBusinessBlocked

`func (o *WindowsInformationProtectionPolicy) SetWindowsHelloForBusinessBlocked(v bool)`

SetWindowsHelloForBusinessBlocked gets a reference to the given bool and assigns it to the WindowsHelloForBusinessBlocked field.

### GetPinMinimumLength

`func (o *WindowsInformationProtectionPolicy) GetPinMinimumLength() int32`

GetPinMinimumLength returns the PinMinimumLength field if non-nil, zero value otherwise.

### GetPinMinimumLengthOk

`func (o *WindowsInformationProtectionPolicy) GetPinMinimumLengthOk() (int32, bool)`

GetPinMinimumLengthOk returns a tuple with the PinMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinMinimumLength

`func (o *WindowsInformationProtectionPolicy) HasPinMinimumLength() bool`

HasPinMinimumLength returns a boolean if a field has been set.

### SetPinMinimumLength

`func (o *WindowsInformationProtectionPolicy) SetPinMinimumLength(v int32)`

SetPinMinimumLength gets a reference to the given int32 and assigns it to the PinMinimumLength field.

### GetPinUppercaseLetters

`func (o *WindowsInformationProtectionPolicy) GetPinUppercaseLetters() AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements`

GetPinUppercaseLetters returns the PinUppercaseLetters field if non-nil, zero value otherwise.

### GetPinUppercaseLettersOk

`func (o *WindowsInformationProtectionPolicy) GetPinUppercaseLettersOk() (AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, bool)`

GetPinUppercaseLettersOk returns a tuple with the PinUppercaseLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinUppercaseLetters

`func (o *WindowsInformationProtectionPolicy) HasPinUppercaseLetters() bool`

HasPinUppercaseLetters returns a boolean if a field has been set.

### SetPinUppercaseLetters

`func (o *WindowsInformationProtectionPolicy) SetPinUppercaseLetters(v AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements)`

SetPinUppercaseLetters gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements and assigns it to the PinUppercaseLetters field.

### GetPinLowercaseLetters

`func (o *WindowsInformationProtectionPolicy) GetPinLowercaseLetters() AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements`

GetPinLowercaseLetters returns the PinLowercaseLetters field if non-nil, zero value otherwise.

### GetPinLowercaseLettersOk

`func (o *WindowsInformationProtectionPolicy) GetPinLowercaseLettersOk() (AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, bool)`

GetPinLowercaseLettersOk returns a tuple with the PinLowercaseLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinLowercaseLetters

`func (o *WindowsInformationProtectionPolicy) HasPinLowercaseLetters() bool`

HasPinLowercaseLetters returns a boolean if a field has been set.

### SetPinLowercaseLetters

`func (o *WindowsInformationProtectionPolicy) SetPinLowercaseLetters(v AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements)`

SetPinLowercaseLetters gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements and assigns it to the PinLowercaseLetters field.

### GetPinSpecialCharacters

`func (o *WindowsInformationProtectionPolicy) GetPinSpecialCharacters() AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements`

GetPinSpecialCharacters returns the PinSpecialCharacters field if non-nil, zero value otherwise.

### GetPinSpecialCharactersOk

`func (o *WindowsInformationProtectionPolicy) GetPinSpecialCharactersOk() (AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, bool)`

GetPinSpecialCharactersOk returns a tuple with the PinSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinSpecialCharacters

`func (o *WindowsInformationProtectionPolicy) HasPinSpecialCharacters() bool`

HasPinSpecialCharacters returns a boolean if a field has been set.

### SetPinSpecialCharacters

`func (o *WindowsInformationProtectionPolicy) SetPinSpecialCharacters(v AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements)`

SetPinSpecialCharacters gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements and assigns it to the PinSpecialCharacters field.

### GetPinExpirationDays

`func (o *WindowsInformationProtectionPolicy) GetPinExpirationDays() int32`

GetPinExpirationDays returns the PinExpirationDays field if non-nil, zero value otherwise.

### GetPinExpirationDaysOk

`func (o *WindowsInformationProtectionPolicy) GetPinExpirationDaysOk() (int32, bool)`

GetPinExpirationDaysOk returns a tuple with the PinExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinExpirationDays

`func (o *WindowsInformationProtectionPolicy) HasPinExpirationDays() bool`

HasPinExpirationDays returns a boolean if a field has been set.

### SetPinExpirationDays

`func (o *WindowsInformationProtectionPolicy) SetPinExpirationDays(v int32)`

SetPinExpirationDays gets a reference to the given int32 and assigns it to the PinExpirationDays field.

### GetNumberOfPastPinsRemembered

`func (o *WindowsInformationProtectionPolicy) GetNumberOfPastPinsRemembered() int32`

GetNumberOfPastPinsRemembered returns the NumberOfPastPinsRemembered field if non-nil, zero value otherwise.

### GetNumberOfPastPinsRememberedOk

`func (o *WindowsInformationProtectionPolicy) GetNumberOfPastPinsRememberedOk() (int32, bool)`

GetNumberOfPastPinsRememberedOk returns a tuple with the NumberOfPastPinsRemembered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberOfPastPinsRemembered

`func (o *WindowsInformationProtectionPolicy) HasNumberOfPastPinsRemembered() bool`

HasNumberOfPastPinsRemembered returns a boolean if a field has been set.

### SetNumberOfPastPinsRemembered

`func (o *WindowsInformationProtectionPolicy) SetNumberOfPastPinsRemembered(v int32)`

SetNumberOfPastPinsRemembered gets a reference to the given int32 and assigns it to the NumberOfPastPinsRemembered field.

### GetPasswordMaximumAttemptCount

`func (o *WindowsInformationProtectionPolicy) GetPasswordMaximumAttemptCount() int32`

GetPasswordMaximumAttemptCount returns the PasswordMaximumAttemptCount field if non-nil, zero value otherwise.

### GetPasswordMaximumAttemptCountOk

`func (o *WindowsInformationProtectionPolicy) GetPasswordMaximumAttemptCountOk() (int32, bool)`

GetPasswordMaximumAttemptCountOk returns a tuple with the PasswordMaximumAttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMaximumAttemptCount

`func (o *WindowsInformationProtectionPolicy) HasPasswordMaximumAttemptCount() bool`

HasPasswordMaximumAttemptCount returns a boolean if a field has been set.

### SetPasswordMaximumAttemptCount

`func (o *WindowsInformationProtectionPolicy) SetPasswordMaximumAttemptCount(v int32)`

SetPasswordMaximumAttemptCount gets a reference to the given int32 and assigns it to the PasswordMaximumAttemptCount field.

### GetMinutesOfInactivityBeforeDeviceLock

`func (o *WindowsInformationProtectionPolicy) GetMinutesOfInactivityBeforeDeviceLock() int32`

GetMinutesOfInactivityBeforeDeviceLock returns the MinutesOfInactivityBeforeDeviceLock field if non-nil, zero value otherwise.

### GetMinutesOfInactivityBeforeDeviceLockOk

`func (o *WindowsInformationProtectionPolicy) GetMinutesOfInactivityBeforeDeviceLockOk() (int32, bool)`

GetMinutesOfInactivityBeforeDeviceLockOk returns a tuple with the MinutesOfInactivityBeforeDeviceLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinutesOfInactivityBeforeDeviceLock

`func (o *WindowsInformationProtectionPolicy) HasMinutesOfInactivityBeforeDeviceLock() bool`

HasMinutesOfInactivityBeforeDeviceLock returns a boolean if a field has been set.

### SetMinutesOfInactivityBeforeDeviceLock

`func (o *WindowsInformationProtectionPolicy) SetMinutesOfInactivityBeforeDeviceLock(v int32)`

SetMinutesOfInactivityBeforeDeviceLock gets a reference to the given int32 and assigns it to the MinutesOfInactivityBeforeDeviceLock field.

### GetDaysWithoutContactBeforeUnenroll

`func (o *WindowsInformationProtectionPolicy) GetDaysWithoutContactBeforeUnenroll() int32`

GetDaysWithoutContactBeforeUnenroll returns the DaysWithoutContactBeforeUnenroll field if non-nil, zero value otherwise.

### GetDaysWithoutContactBeforeUnenrollOk

`func (o *WindowsInformationProtectionPolicy) GetDaysWithoutContactBeforeUnenrollOk() (int32, bool)`

GetDaysWithoutContactBeforeUnenrollOk returns a tuple with the DaysWithoutContactBeforeUnenroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDaysWithoutContactBeforeUnenroll

`func (o *WindowsInformationProtectionPolicy) HasDaysWithoutContactBeforeUnenroll() bool`

HasDaysWithoutContactBeforeUnenroll returns a boolean if a field has been set.

### SetDaysWithoutContactBeforeUnenroll

`func (o *WindowsInformationProtectionPolicy) SetDaysWithoutContactBeforeUnenroll(v int32)`

SetDaysWithoutContactBeforeUnenroll gets a reference to the given int32 and assigns it to the DaysWithoutContactBeforeUnenroll field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


