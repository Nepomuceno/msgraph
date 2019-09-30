# MicrosoftGraphEnrollmentTroubleshootingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**EventDateTime** | Pointer to [**time.Time**](time.Time.md) | Time when the event occurred . | [optional] 
**CorrelationId** | Pointer to **string** | Id used for tracing the failure in the service. | [optional] 
**ManagedDeviceIdentifier** | Pointer to **string** | Device identifier created or collected by Intune. | [optional] 
**OperatingSystem** | Pointer to **string** | Operating System. | [optional] 
**OsVersion** | Pointer to **string** | OS Version. | [optional] 
**UserId** | Pointer to **string** | Identifier for the user that tried to enroll the device. | [optional] 
**DeviceId** | Pointer to **string** | Azure AD device identifier. | [optional] 
**EnrollmentType** | Pointer to [**AnyOfmicrosoftGraphDeviceEnrollmentType**](anyOf&lt;microsoft.graph.deviceEnrollmentType&gt;.md) | Type of the enrollment. | [optional] 
**FailureCategory** | Pointer to [**AnyOfmicrosoftGraphDeviceEnrollmentFailureReason**](anyOf&lt;microsoft.graph.deviceEnrollmentFailureReason&gt;.md) | Highlevel failure category. | [optional] 
**FailureReason** | Pointer to **string** | Detailed failure reason. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetEventDateTime

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetEventDateTimeOk() (time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventDateTime

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) HasEventDateTime() bool`

HasEventDateTime returns a boolean if a field has been set.

### SetEventDateTime

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetEventDateTime(v time.Time)`

SetEventDateTime gets a reference to the given time.Time and assigns it to the EventDateTime field.

### GetCorrelationId

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetCorrelationIdOk() (string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCorrelationId

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationId

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetCorrelationId(v string)`

SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.

### SetCorrelationIdExplicitNull

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetCorrelationIdExplicitNull(b bool)`

SetCorrelationIdExplicitNull (un)sets CorrelationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CorrelationId value is set to nil even if false is passed
### GetManagedDeviceIdentifier

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetManagedDeviceIdentifier() string`

GetManagedDeviceIdentifier returns the ManagedDeviceIdentifier field if non-nil, zero value otherwise.

### GetManagedDeviceIdentifierOk

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetManagedDeviceIdentifierOk() (string, bool)`

GetManagedDeviceIdentifierOk returns a tuple with the ManagedDeviceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedDeviceIdentifier

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) HasManagedDeviceIdentifier() bool`

HasManagedDeviceIdentifier returns a boolean if a field has been set.

### SetManagedDeviceIdentifier

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetManagedDeviceIdentifier(v string)`

SetManagedDeviceIdentifier gets a reference to the given string and assigns it to the ManagedDeviceIdentifier field.

### SetManagedDeviceIdentifierExplicitNull

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetManagedDeviceIdentifierExplicitNull(b bool)`

SetManagedDeviceIdentifierExplicitNull (un)sets ManagedDeviceIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ManagedDeviceIdentifier value is set to nil even if false is passed
### GetOperatingSystem

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetOperatingSystemOk() (string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperatingSystem

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystem

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetOperatingSystem(v string)`

SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.

### SetOperatingSystemExplicitNull

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetOperatingSystemExplicitNull(b bool)`

SetOperatingSystemExplicitNull (un)sets OperatingSystem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperatingSystem value is set to nil even if false is passed
### GetOsVersion

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetOsVersionOk() (string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsVersion

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersion

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetOsVersion(v string)`

SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.

### SetOsVersionExplicitNull

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetOsVersionExplicitNull(b bool)`

SetOsVersionExplicitNull (un)sets OsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsVersion value is set to nil even if false is passed
### GetUserId

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### SetUserIdExplicitNull

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetUserIdExplicitNull(b bool)`

SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserId value is set to nil even if false is passed
### GetDeviceId

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetDeviceIdOk() (string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceId

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceId

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetDeviceId(v string)`

SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.

### SetDeviceIdExplicitNull

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetDeviceIdExplicitNull(b bool)`

SetDeviceIdExplicitNull (un)sets DeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceId value is set to nil even if false is passed
### GetEnrollmentType

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetEnrollmentType() AnyOfmicrosoftGraphDeviceEnrollmentType`

GetEnrollmentType returns the EnrollmentType field if non-nil, zero value otherwise.

### GetEnrollmentTypeOk

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetEnrollmentTypeOk() (AnyOfmicrosoftGraphDeviceEnrollmentType, bool)`

GetEnrollmentTypeOk returns a tuple with the EnrollmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnrollmentType

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) HasEnrollmentType() bool`

HasEnrollmentType returns a boolean if a field has been set.

### SetEnrollmentType

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetEnrollmentType(v AnyOfmicrosoftGraphDeviceEnrollmentType)`

SetEnrollmentType gets a reference to the given AnyOfmicrosoftGraphDeviceEnrollmentType and assigns it to the EnrollmentType field.

### GetFailureCategory

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetFailureCategory() AnyOfmicrosoftGraphDeviceEnrollmentFailureReason`

GetFailureCategory returns the FailureCategory field if non-nil, zero value otherwise.

### GetFailureCategoryOk

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetFailureCategoryOk() (AnyOfmicrosoftGraphDeviceEnrollmentFailureReason, bool)`

GetFailureCategoryOk returns a tuple with the FailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailureCategory

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) HasFailureCategory() bool`

HasFailureCategory returns a boolean if a field has been set.

### SetFailureCategory

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetFailureCategory(v AnyOfmicrosoftGraphDeviceEnrollmentFailureReason)`

SetFailureCategory gets a reference to the given AnyOfmicrosoftGraphDeviceEnrollmentFailureReason and assigns it to the FailureCategory field.

### GetFailureReason

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) GetFailureReasonOk() (string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailureReason

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReason

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetFailureReason(v string)`

SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.

### SetFailureReasonExplicitNull

`func (o *MicrosoftGraphEnrollmentTroubleshootingEvent) SetFailureReasonExplicitNull(b bool)`

SetFailureReasonExplicitNull (un)sets FailureReason to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FailureReason value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


