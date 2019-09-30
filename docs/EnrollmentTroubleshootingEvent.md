# EnrollmentTroubleshootingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagedDeviceIdentifier** | Pointer to **string** | Device identifier created or collected by Intune. | [optional] 
**OperatingSystem** | Pointer to **string** | Operating System. | [optional] 
**OsVersion** | Pointer to **string** | OS Version. | [optional] 
**UserId** | Pointer to **string** | Identifier for the user that tried to enroll the device. | [optional] 
**DeviceId** | Pointer to **string** | Azure AD device identifier. | [optional] 
**EnrollmentType** | Pointer to [**AnyOfmicrosoftGraphDeviceEnrollmentType**](anyOf&lt;microsoft.graph.deviceEnrollmentType&gt;.md) | Type of the enrollment. | [optional] 
**FailureCategory** | Pointer to [**AnyOfmicrosoftGraphDeviceEnrollmentFailureReason**](anyOf&lt;microsoft.graph.deviceEnrollmentFailureReason&gt;.md) | Highlevel failure category. | [optional] 
**FailureReason** | Pointer to **string** | Detailed failure reason. | [optional] 

## Methods

### GetManagedDeviceIdentifier

`func (o *EnrollmentTroubleshootingEvent) GetManagedDeviceIdentifier() string`

GetManagedDeviceIdentifier returns the ManagedDeviceIdentifier field if non-nil, zero value otherwise.

### GetManagedDeviceIdentifierOk

`func (o *EnrollmentTroubleshootingEvent) GetManagedDeviceIdentifierOk() (string, bool)`

GetManagedDeviceIdentifierOk returns a tuple with the ManagedDeviceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedDeviceIdentifier

`func (o *EnrollmentTroubleshootingEvent) HasManagedDeviceIdentifier() bool`

HasManagedDeviceIdentifier returns a boolean if a field has been set.

### SetManagedDeviceIdentifier

`func (o *EnrollmentTroubleshootingEvent) SetManagedDeviceIdentifier(v string)`

SetManagedDeviceIdentifier gets a reference to the given string and assigns it to the ManagedDeviceIdentifier field.

### SetManagedDeviceIdentifierExplicitNull

`func (o *EnrollmentTroubleshootingEvent) SetManagedDeviceIdentifierExplicitNull(b bool)`

SetManagedDeviceIdentifierExplicitNull (un)sets ManagedDeviceIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ManagedDeviceIdentifier value is set to nil even if false is passed
### GetOperatingSystem

`func (o *EnrollmentTroubleshootingEvent) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *EnrollmentTroubleshootingEvent) GetOperatingSystemOk() (string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperatingSystem

`func (o *EnrollmentTroubleshootingEvent) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystem

`func (o *EnrollmentTroubleshootingEvent) SetOperatingSystem(v string)`

SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.

### SetOperatingSystemExplicitNull

`func (o *EnrollmentTroubleshootingEvent) SetOperatingSystemExplicitNull(b bool)`

SetOperatingSystemExplicitNull (un)sets OperatingSystem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperatingSystem value is set to nil even if false is passed
### GetOsVersion

`func (o *EnrollmentTroubleshootingEvent) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *EnrollmentTroubleshootingEvent) GetOsVersionOk() (string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsVersion

`func (o *EnrollmentTroubleshootingEvent) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersion

`func (o *EnrollmentTroubleshootingEvent) SetOsVersion(v string)`

SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.

### SetOsVersionExplicitNull

`func (o *EnrollmentTroubleshootingEvent) SetOsVersionExplicitNull(b bool)`

SetOsVersionExplicitNull (un)sets OsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsVersion value is set to nil even if false is passed
### GetUserId

`func (o *EnrollmentTroubleshootingEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EnrollmentTroubleshootingEvent) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *EnrollmentTroubleshootingEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *EnrollmentTroubleshootingEvent) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### SetUserIdExplicitNull

`func (o *EnrollmentTroubleshootingEvent) SetUserIdExplicitNull(b bool)`

SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserId value is set to nil even if false is passed
### GetDeviceId

`func (o *EnrollmentTroubleshootingEvent) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *EnrollmentTroubleshootingEvent) GetDeviceIdOk() (string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceId

`func (o *EnrollmentTroubleshootingEvent) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceId

`func (o *EnrollmentTroubleshootingEvent) SetDeviceId(v string)`

SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.

### SetDeviceIdExplicitNull

`func (o *EnrollmentTroubleshootingEvent) SetDeviceIdExplicitNull(b bool)`

SetDeviceIdExplicitNull (un)sets DeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceId value is set to nil even if false is passed
### GetEnrollmentType

`func (o *EnrollmentTroubleshootingEvent) GetEnrollmentType() AnyOfmicrosoftGraphDeviceEnrollmentType`

GetEnrollmentType returns the EnrollmentType field if non-nil, zero value otherwise.

### GetEnrollmentTypeOk

`func (o *EnrollmentTroubleshootingEvent) GetEnrollmentTypeOk() (AnyOfmicrosoftGraphDeviceEnrollmentType, bool)`

GetEnrollmentTypeOk returns a tuple with the EnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnrollmentType

`func (o *EnrollmentTroubleshootingEvent) HasEnrollmentType() bool`

HasEnrollmentType returns a boolean if a field has been set.

### SetEnrollmentType

`func (o *EnrollmentTroubleshootingEvent) SetEnrollmentType(v AnyOfmicrosoftGraphDeviceEnrollmentType)`

SetEnrollmentType gets a reference to the given AnyOfmicrosoftGraphDeviceEnrollmentType and assigns it to the EnrollmentType field.

### GetFailureCategory

`func (o *EnrollmentTroubleshootingEvent) GetFailureCategory() AnyOfmicrosoftGraphDeviceEnrollmentFailureReason`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *EnrollmentTroubleshootingEvent) GetFailureCategoryOk() (AnyOfmicrosoftGraphDeviceEnrollmentFailureReason, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailureCategory

`func (o *EnrollmentTroubleshootingEvent) HasFailureCategory() bool`

HasFailureCategory returns a boolean if a field has been set.

### SetFailureCategory

`func (o *EnrollmentTroubleshootingEvent) SetFailureCategory(v AnyOfmicrosoftGraphDeviceEnrollmentFailureReason)`

SetFailureCategory gets a reference to the given AnyOfmicrosoftGraphDeviceEnrollmentFailureReason and assigns it to the FailureCategory field.

### GetFailureReason

`func (o *EnrollmentTroubleshootingEvent) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *EnrollmentTroubleshootingEvent) GetFailureReasonOk() (string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailureReason

`func (o *EnrollmentTroubleshootingEvent) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReason

`func (o *EnrollmentTroubleshootingEvent) SetFailureReason(v string)`

SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.

### SetFailureReasonExplicitNull

`func (o *EnrollmentTroubleshootingEvent) SetFailureReasonExplicitNull(b bool)`

SetFailureReasonExplicitNull (un)sets FailureReason to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FailureReason value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


