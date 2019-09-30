# MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphEnrollmentConfigurationAssignment**](microsoft.graph.enrollmentConfigurationAssignment.md) |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPriority

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPriorityOk() (int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPriority

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriority

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPriority(v int32)`

SetPriority gets a reference to the given int32 and assigns it to the Priority field.

### GetCreatedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetVersion

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetAssignments() []MicrosoftGraphEnrollmentConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetAssignmentsOk() ([]MicrosoftGraphEnrollmentConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetAssignments(v []MicrosoftGraphEnrollmentConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphEnrollmentConfigurationAssignment and assigns it to the Assignments field.

### GetPinMinimumLength

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinMinimumLength() int32`

GetPinMinimumLength returns the PinMinimumLength field if non-nil, zero value otherwise.

### GetPinMinimumLengthOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinMinimumLengthOk() (int32, bool)`

GetPinMinimumLengthOk returns a tuple with the PinMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinMinimumLength

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinMinimumLength() bool`

HasPinMinimumLength returns a boolean if a field has been set.

### SetPinMinimumLength

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinMinimumLength(v int32)`

SetPinMinimumLength gets a reference to the given int32 and assigns it to the PinMinimumLength field.

### GetPinMaximumLength

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinMaximumLength() int32`

GetPinMaximumLength returns the PinMaximumLength field if non-nil, zero value otherwise.

### GetPinMaximumLengthOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinMaximumLengthOk() (int32, bool)`

GetPinMaximumLengthOk returns a tuple with the PinMaximumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinMaximumLength

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinMaximumLength() bool`

HasPinMaximumLength returns a boolean if a field has been set.

### SetPinMaximumLength

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinMaximumLength(v int32)`

SetPinMaximumLength gets a reference to the given int32 and assigns it to the PinMaximumLength field.

### GetPinUppercaseCharactersUsage

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinUppercaseCharactersUsage() AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage`

GetPinUppercaseCharactersUsage returns the PinUppercaseCharactersUsage field if non-nil, zero value otherwise.

### GetPinUppercaseCharactersUsageOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinUppercaseCharactersUsageOk() (AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage, bool)`

GetPinUppercaseCharactersUsageOk returns a tuple with the PinUppercaseCharactersUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinUppercaseCharactersUsage

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinUppercaseCharactersUsage() bool`

HasPinUppercaseCharactersUsage returns a boolean if a field has been set.

### SetPinUppercaseCharactersUsage

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinUppercaseCharactersUsage(v AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage)`

SetPinUppercaseCharactersUsage gets a reference to the given AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage and assigns it to the PinUppercaseCharactersUsage field.

### GetPinLowercaseCharactersUsage

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinLowercaseCharactersUsage() AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage`

GetPinLowercaseCharactersUsage returns the PinLowercaseCharactersUsage field if non-nil, zero value otherwise.

### GetPinLowercaseCharactersUsageOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinLowercaseCharactersUsageOk() (AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage, bool)`

GetPinLowercaseCharactersUsageOk returns a tuple with the PinLowercaseCharactersUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinLowercaseCharactersUsage

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinLowercaseCharactersUsage() bool`

HasPinLowercaseCharactersUsage returns a boolean if a field has been set.

### SetPinLowercaseCharactersUsage

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinLowercaseCharactersUsage(v AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage)`

SetPinLowercaseCharactersUsage gets a reference to the given AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage and assigns it to the PinLowercaseCharactersUsage field.

### GetPinSpecialCharactersUsage

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinSpecialCharactersUsage() AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage`

GetPinSpecialCharactersUsage returns the PinSpecialCharactersUsage field if non-nil, zero value otherwise.

### GetPinSpecialCharactersUsageOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinSpecialCharactersUsageOk() (AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage, bool)`

GetPinSpecialCharactersUsageOk returns a tuple with the PinSpecialCharactersUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinSpecialCharactersUsage

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinSpecialCharactersUsage() bool`

HasPinSpecialCharactersUsage returns a boolean if a field has been set.

### SetPinSpecialCharactersUsage

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinSpecialCharactersUsage(v AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage)`

SetPinSpecialCharactersUsage gets a reference to the given AnyOfmicrosoftGraphWindowsHelloForBusinessPinUsage and assigns it to the PinSpecialCharactersUsage field.

### GetState

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetState() AnyOfmicrosoftGraphEnablement`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetStateOk() (AnyOfmicrosoftGraphEnablement, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetState(v AnyOfmicrosoftGraphEnablement)`

SetState gets a reference to the given AnyOfmicrosoftGraphEnablement and assigns it to the State field.

### GetSecurityDeviceRequired

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetSecurityDeviceRequired() bool`

GetSecurityDeviceRequired returns the SecurityDeviceRequired field if non-nil, zero value otherwise.

### GetSecurityDeviceRequiredOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetSecurityDeviceRequiredOk() (bool, bool)`

GetSecurityDeviceRequiredOk returns a tuple with the SecurityDeviceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityDeviceRequired

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasSecurityDeviceRequired() bool`

HasSecurityDeviceRequired returns a boolean if a field has been set.

### SetSecurityDeviceRequired

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetSecurityDeviceRequired(v bool)`

SetSecurityDeviceRequired gets a reference to the given bool and assigns it to the SecurityDeviceRequired field.

### GetUnlockWithBiometricsEnabled

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetUnlockWithBiometricsEnabled() bool`

GetUnlockWithBiometricsEnabled returns the UnlockWithBiometricsEnabled field if non-nil, zero value otherwise.

### GetUnlockWithBiometricsEnabledOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetUnlockWithBiometricsEnabledOk() (bool, bool)`

GetUnlockWithBiometricsEnabledOk returns a tuple with the UnlockWithBiometricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnlockWithBiometricsEnabled

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasUnlockWithBiometricsEnabled() bool`

HasUnlockWithBiometricsEnabled returns a boolean if a field has been set.

### SetUnlockWithBiometricsEnabled

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetUnlockWithBiometricsEnabled(v bool)`

SetUnlockWithBiometricsEnabled gets a reference to the given bool and assigns it to the UnlockWithBiometricsEnabled field.

### GetRemotePassportEnabled

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetRemotePassportEnabled() bool`

GetRemotePassportEnabled returns the RemotePassportEnabled field if non-nil, zero value otherwise.

### GetRemotePassportEnabledOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetRemotePassportEnabledOk() (bool, bool)`

GetRemotePassportEnabledOk returns a tuple with the RemotePassportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemotePassportEnabled

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasRemotePassportEnabled() bool`

HasRemotePassportEnabled returns a boolean if a field has been set.

### SetRemotePassportEnabled

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetRemotePassportEnabled(v bool)`

SetRemotePassportEnabled gets a reference to the given bool and assigns it to the RemotePassportEnabled field.

### GetPinPreviousBlockCount

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinPreviousBlockCount() int32`

GetPinPreviousBlockCount returns the PinPreviousBlockCount field if non-nil, zero value otherwise.

### GetPinPreviousBlockCountOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinPreviousBlockCountOk() (int32, bool)`

GetPinPreviousBlockCountOk returns a tuple with the PinPreviousBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinPreviousBlockCount

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinPreviousBlockCount() bool`

HasPinPreviousBlockCount returns a boolean if a field has been set.

### SetPinPreviousBlockCount

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinPreviousBlockCount(v int32)`

SetPinPreviousBlockCount gets a reference to the given int32 and assigns it to the PinPreviousBlockCount field.

### GetPinExpirationInDays

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinExpirationInDays() int32`

GetPinExpirationInDays returns the PinExpirationInDays field if non-nil, zero value otherwise.

### GetPinExpirationInDaysOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetPinExpirationInDaysOk() (int32, bool)`

GetPinExpirationInDaysOk returns a tuple with the PinExpirationInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinExpirationInDays

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasPinExpirationInDays() bool`

HasPinExpirationInDays returns a boolean if a field has been set.

### SetPinExpirationInDays

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetPinExpirationInDays(v int32)`

SetPinExpirationInDays gets a reference to the given int32 and assigns it to the PinExpirationInDays field.

### GetEnhancedBiometricsState

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetEnhancedBiometricsState() AnyOfmicrosoftGraphEnablement`

GetEnhancedBiometricsState returns the EnhancedBiometricsState field if non-nil, zero value otherwise.

### GetEnhancedBiometricsStateOk

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) GetEnhancedBiometricsStateOk() (AnyOfmicrosoftGraphEnablement, bool)`

GetEnhancedBiometricsStateOk returns a tuple with the EnhancedBiometricsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnhancedBiometricsState

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) HasEnhancedBiometricsState() bool`

HasEnhancedBiometricsState returns a boolean if a field has been set.

### SetEnhancedBiometricsState

`func (o *MicrosoftGraphDeviceEnrollmentWindowsHelloForBusinessConfiguration) SetEnhancedBiometricsState(v AnyOfmicrosoftGraphEnablement)`

SetEnhancedBiometricsState gets a reference to the given AnyOfmicrosoftGraphEnablement and assigns it to the EnhancedBiometricsState field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


