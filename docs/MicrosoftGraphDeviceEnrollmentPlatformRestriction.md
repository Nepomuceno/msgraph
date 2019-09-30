# MicrosoftGraphDeviceEnrollmentPlatformRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformBlocked** | Pointer to **bool** | Block the platform from enrolling | [optional] 
**PersonalDeviceEnrollmentBlocked** | Pointer to **bool** | Block personally owned devices from enrolling | [optional] 
**OsMinimumVersion** | Pointer to **string** | Min OS version supported | [optional] 
**OsMaximumVersion** | Pointer to **string** | Max OS version supported | [optional] 

## Methods

### GetPlatformBlocked

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetPlatformBlocked() bool`

GetPlatformBlocked returns the PlatformBlocked field if non-nil, zero value otherwise.

### GetPlatformBlockedOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetPlatformBlockedOk() (bool, bool)`

GetPlatformBlockedOk returns a tuple with the PlatformBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlatformBlocked

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) HasPlatformBlocked() bool`

HasPlatformBlocked returns a boolean if a field has been set.

### SetPlatformBlocked

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) SetPlatformBlocked(v bool)`

SetPlatformBlocked gets a reference to the given bool and assigns it to the PlatformBlocked field.

### GetPersonalDeviceEnrollmentBlocked

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetPersonalDeviceEnrollmentBlocked() bool`

GetPersonalDeviceEnrollmentBlocked returns the PersonalDeviceEnrollmentBlocked field if non-nil, zero value otherwise.

### GetPersonalDeviceEnrollmentBlockedOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetPersonalDeviceEnrollmentBlockedOk() (bool, bool)`

GetPersonalDeviceEnrollmentBlockedOk returns a tuple with the PersonalDeviceEnrollmentBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPersonalDeviceEnrollmentBlocked

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) HasPersonalDeviceEnrollmentBlocked() bool`

HasPersonalDeviceEnrollmentBlocked returns a boolean if a field has been set.

### SetPersonalDeviceEnrollmentBlocked

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) SetPersonalDeviceEnrollmentBlocked(v bool)`

SetPersonalDeviceEnrollmentBlocked gets a reference to the given bool and assigns it to the PersonalDeviceEnrollmentBlocked field.

### GetOsMinimumVersion

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetOsMinimumVersion() string`

GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.

### GetOsMinimumVersionOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetOsMinimumVersionOk() (string, bool)`

GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMinimumVersion

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) HasOsMinimumVersion() bool`

HasOsMinimumVersion returns a boolean if a field has been set.

### SetOsMinimumVersion

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) SetOsMinimumVersion(v string)`

SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.

### SetOsMinimumVersionExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) SetOsMinimumVersionExplicitNull(b bool)`

SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMinimumVersion value is set to nil even if false is passed
### GetOsMaximumVersion

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetOsMaximumVersion() string`

GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.

### GetOsMaximumVersionOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetOsMaximumVersionOk() (string, bool)`

GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsMaximumVersion

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) HasOsMaximumVersion() bool`

HasOsMaximumVersion returns a boolean if a field has been set.

### SetOsMaximumVersion

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) SetOsMaximumVersion(v string)`

SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.

### SetOsMaximumVersionExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) SetOsMaximumVersionExplicitNull(b bool)`

SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsMaximumVersion value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


