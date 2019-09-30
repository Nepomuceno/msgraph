# DeviceEnrollmentWindowsHelloForBusinessConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PinMinimumLength** | Pointer to **int32** |  | [optional] 
**PinMaximumLength** | Pointer to **int32** |  | [optional] 
**PinUppercaseCharactersUsage** | Pointer to [**AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage**](anyOf&lt;microsoft.graph.windowsHelloForBusinessPinUsage&gt;.md) |  | [optional] 
**PinLowercaseCharactersUsage** | Pointer to [**AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage**](anyOf&lt;microsoft.graph.windowsHelloForBusinessPinUsage&gt;.md) |  | [optional] 
**PinSpecialCharactersUsage** | Pointer to [**AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage**](anyOf&lt;microsoft.graph.windowsHelloForBusinessPinUsage&gt;.md) |  | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphEnablement**](anyOf&lt;microsoft.graph.enablement&gt;.md) |  | [optional] 
**SecurityDeviceRequired** | Pointer to **bool** |  | [optional] 
**UnlockWithBiometricsEnabled** | Pointer to **bool** |  | [optional] 
**RemotePassportEnabled** | Pointer to **bool** |  | [optional] 
**PinPreviousBlockCount** | Pointer to **int32** |  | [optional] 
**PinExpirationInDays** | Pointer to **int32** |  | [optional] 
**EnhancedBiometricsState** | Pointer to [**AnyOfmicrosoftGraphEnablement**](anyOf&lt;microsoft.graph.enablement&gt;.md) |  | [optional] 

## Methods

### GetPinMinimumLength

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinMinimumLength() int32`

GetPinMinimumLength returns the PinMinimumLength field if non-nil, zero value otherwise.

### GetPinMinimumLengthOk

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinMinimumLengthOk() (int32, bool)`

GetPinMinimumLengthOk returns a tuple with the PinMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinMinimumLength

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinMinimumLength() bool`

HasPinMinimumLength returns a boolean if a field has been set.

### SetPinMinimumLength

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinMinimumLength(v int32)`

SetPinMinimumLength gets a reference to the given int32 and assigns it to the PinMinimumLength field.

### GetPinMaximumLength

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinMaximumLength() int32`

GetPinMaximumLength returns the PinMaximumLength field if non-nil, zero value otherwise.

### GetPinMaximumLengthOk

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinMaximumLengthOk() (int32, bool)`

GetPinMaximumLengthOk returns a tuple with the PinMaximumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinMaximumLength

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinMaximumLength() bool`

HasPinMaximumLength returns a boolean if a field has been set.

### SetPinMaximumLength

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinMaximumLength(v int32)`

SetPinMaximumLength gets a reference to the given int32 and assigns it to the PinMaximumLength field.

### GetPinUppercaseCharactersUsage

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinUppercaseCharactersUsage() AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage`

GetPinUppercaseCharactersUsage returns the PinUppercaseCharactersUsage field if non-nil, zero value otherwise.

### GetPinUppercaseCharactersUsageOk

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinUppercaseCharactersUsageOk() (AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage, bool)`

GetPinUppercaseCharactersUsageOk returns a tuple with the PinUppercaseCharactersUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinUppercaseCharactersUsage

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinUppercaseCharactersUsage() bool`

HasPinUppercaseCharactersUsage returns a boolean if a field has been set.

### SetPinUppercaseCharactersUsage

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinUppercaseCharactersUsage(v AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage)`

SetPinUppercaseCharactersUsage gets a reference to the given AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage and assigns it to the PinUppercaseCharactersUsage field.

### GetPinLowercaseCharactersUsage

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinLowercaseCharactersUsage() AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage`

GetPinLowercaseCharactersUsage returns the PinLowercaseCharactersUsage field if non-nil, zero value otherwise.

### GetPinLowercaseCharactersUsageOk

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinLowercaseCharactersUsageOk() (AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage, bool)`

GetPinLowercaseCharactersUsageOk returns a tuple with the PinLowercaseCharactersUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinLowercaseCharactersUsage

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinLowercaseCharactersUsage() bool`

HasPinLowercaseCharactersUsage returns a boolean if a field has been set.

### SetPinLowercaseCharactersUsage

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinLowercaseCharactersUsage(v AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage)`

SetPinLowercaseCharactersUsage gets a reference to the given AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage and assigns it to the PinLowercaseCharactersUsage field.

### GetPinSpecialCharactersUsage

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinSpecialCharactersUsage() AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage`

GetPinSpecialCharactersUsage returns the PinSpecialCharactersUsage field if non-nil, zero value otherwise.

### GetPinSpecialCharactersUsageOk

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinSpecialCharactersUsageOk() (AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage, bool)`

GetPinSpecialCharactersUsageOk returns a tuple with the PinSpecialCharactersUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinSpecialCharactersUsage

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinSpecialCharactersUsage() bool`

HasPinSpecialCharactersUsage returns a boolean if a field has been set.

### SetPinSpecialCharactersUsage

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinSpecialCharactersUsage(v AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage)`

SetPinSpecialCharactersUsage gets a reference to the given AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage and assigns it to the PinSpecialCharactersUsage field.

### GetState

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetState() AnyOfmicrosoftGraphEnablement`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetStateOk() (AnyOfmicrosoftGraphEnablement, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetState(v AnyOfmicrosoftGraphEnablement)`

SetState gets a reference to the given AnyOfmicrosoftGraphEnablement and assigns it to the State field.

### GetSecurityDeviceRequired

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetSecurityDeviceRequired() bool`

GetSecurityDeviceRequired returns the SecurityDeviceRequired field if non-nil, zero value otherwise.

### GetSecurityDeviceRequiredOk

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetSecurityDeviceRequiredOk() (bool, bool)`

GetSecurityDeviceRequiredOk returns a tuple with the SecurityDeviceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityDeviceRequired

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) HasSecurityDeviceRequired() bool`

HasSecurityDeviceRequired returns a boolean if a field has been set.

### SetSecurityDeviceRequired

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetSecurityDeviceRequired(v bool)`

SetSecurityDeviceRequired gets a reference to the given bool and assigns it to the SecurityDeviceRequired field.

### GetUnlockWithBiometricsEnabled

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetUnlockWithBiometricsEnabled() bool`

GetUnlockWithBiometricsEnabled returns the UnlockWithBiometricsEnabled field if non-nil, zero value otherwise.

### GetUnlockWithBiometricsEnabledOk

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetUnlockWithBiometricsEnabledOk() (bool, bool)`

GetUnlockWithBiometricsEnabledOk returns a tuple with the UnlockWithBiometricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnlockWithBiometricsEnabled

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) HasUnlockWithBiometricsEnabled() bool`

HasUnlockWithBiometricsEnabled returns a boolean if a field has been set.

### SetUnlockWithBiometricsEnabled

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetUnlockWithBiometricsEnabled(v bool)`

SetUnlockWithBiometricsEnabled gets a reference to the given bool and assigns it to the UnlockWithBiometricsEnabled field.

### GetRemotePassportEnabled

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetRemotePassportEnabled() bool`

GetRemotePassportEnabled returns the RemotePassportEnabled field if non-nil, zero value otherwise.

### GetRemotePassportEnabledOk

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetRemotePassportEnabledOk() (bool, bool)`

GetRemotePassportEnabledOk returns a tuple with the RemotePassportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemotePassportEnabled

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) HasRemotePassportEnabled() bool`

HasRemotePassportEnabled returns a boolean if a field has been set.

### SetRemotePassportEnabled

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetRemotePassportEnabled(v bool)`

SetRemotePassportEnabled gets a reference to the given bool and assigns it to the RemotePassportEnabled field.

### GetPinPreviousBlockCount

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinPreviousBlockCount() int32`

GetPinPreviousBlockCount returns the PinPreviousBlockCount field if non-nil, zero value otherwise.

### GetPinPreviousBlockCountOk

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinPreviousBlockCountOk() (int32, bool)`

GetPinPreviousBlockCountOk returns a tuple with the PinPreviousBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinPreviousBlockCount

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinPreviousBlockCount() bool`

HasPinPreviousBlockCount returns a boolean if a field has been set.

### SetPinPreviousBlockCount

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinPreviousBlockCount(v int32)`

SetPinPreviousBlockCount gets a reference to the given int32 and assigns it to the PinPreviousBlockCount field.

### GetPinExpirationInDays

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinExpirationInDays() int32`

GetPinExpirationInDays returns the PinExpirationInDays field if non-nil, zero value otherwise.

### GetPinExpirationInDaysOk

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinExpirationInDaysOk() (int32, bool)`

GetPinExpirationInDaysOk returns a tuple with the PinExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinExpirationInDays

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinExpirationInDays() bool`

HasPinExpirationInDays returns a boolean if a field has been set.

### SetPinExpirationInDays

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinExpirationInDays(v int32)`

SetPinExpirationInDays gets a reference to the given int32 and assigns it to the PinExpirationInDays field.

### GetEnhancedBiometricsState

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetEnhancedBiometricsState() AnyOfmicrosoftGraphEnablement`

GetEnhancedBiometricsState returns the EnhancedBiometricsState field if non-nil, zero value otherwise.

### GetEnhancedBiometricsStateOk

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) GetEnhancedBiometricsStateOk() (AnyOfmicrosoftGraphEnablement, bool)`

GetEnhancedBiometricsStateOk returns a tuple with the EnhancedBiometricsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnhancedBiometricsState

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) HasEnhancedBiometricsState() bool`

HasEnhancedBiometricsState returns a boolean if a field has been set.

### SetEnhancedBiometricsState

`func (o *DeviceEnrollmentWindowsHelloForBusinessConfiguration) SetEnhancedBiometricsState(v AnyOfmicrosoftGraphEnablement)`

SetEnhancedBiometricsState gets a reference to the given AnyOfmicrosoftGraphEnablement and assigns it to the EnhancedBiometricsState field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


