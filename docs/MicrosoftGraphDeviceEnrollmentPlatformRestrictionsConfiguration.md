# MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration

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
**IosRestriction** | Pointer to [**AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction**](anyOf&lt;microsoft.graph.deviceEnrollmentPlatformRestriction&gt;.md) |  | [optional] 
**WindowsRestriction** | Pointer to [**AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction**](anyOf&lt;microsoft.graph.deviceEnrollmentPlatformRestriction&gt;.md) |  | [optional] 
**WindowsMobileRestriction** | Pointer to [**AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction**](anyOf&lt;microsoft.graph.deviceEnrollmentPlatformRestriction&gt;.md) |  | [optional] 
**AndroidRestriction** | Pointer to [**AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction**](anyOf&lt;microsoft.graph.deviceEnrollmentPlatformRestriction&gt;.md) |  | [optional] 
**MacOSRestriction** | Pointer to [**AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction**](anyOf&lt;microsoft.graph.deviceEnrollmentPlatformRestriction&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPriority

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetPriorityOk() (int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPriority

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriority

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetPriority(v int32)`

SetPriority gets a reference to the given int32 and assigns it to the Priority field.

### GetCreatedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetVersion

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetAssignments() []MicrosoftGraphEnrollmentConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetAssignmentsOk() ([]MicrosoftGraphEnrollmentConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetAssignments(v []MicrosoftGraphEnrollmentConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphEnrollmentConfigurationAssignment and assigns it to the Assignments field.

### GetIosRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetIosRestriction() AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction`

GetIosRestriction returns the IosRestriction field if non-nil, zero value otherwise.

### GetIosRestrictionOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetIosRestrictionOk() (AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction, bool)`

GetIosRestrictionOk returns a tuple with the IosRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIosRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasIosRestriction() bool`

HasIosRestriction returns a boolean if a field has been set.

### SetIosRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetIosRestriction(v AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction)`

SetIosRestriction gets a reference to the given AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction and assigns it to the IosRestriction field.

### SetIosRestrictionExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetIosRestrictionExplicitNull(b bool)`

SetIosRestrictionExplicitNull (un)sets IosRestriction to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IosRestriction value is set to nil even if false is passed
### GetWindowsRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetWindowsRestriction() AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction`

GetWindowsRestriction returns the WindowsRestriction field if non-nil, zero value otherwise.

### GetWindowsRestrictionOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetWindowsRestrictionOk() (AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction, bool)`

GetWindowsRestrictionOk returns a tuple with the WindowsRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasWindowsRestriction() bool`

HasWindowsRestriction returns a boolean if a field has been set.

### SetWindowsRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetWindowsRestriction(v AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction)`

SetWindowsRestriction gets a reference to the given AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction and assigns it to the WindowsRestriction field.

### SetWindowsRestrictionExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetWindowsRestrictionExplicitNull(b bool)`

SetWindowsRestrictionExplicitNull (un)sets WindowsRestriction to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WindowsRestriction value is set to nil even if false is passed
### GetWindowsMobileRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetWindowsMobileRestriction() AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction`

GetWindowsMobileRestriction returns the WindowsMobileRestriction field if non-nil, zero value otherwise.

### GetWindowsMobileRestrictionOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetWindowsMobileRestrictionOk() (AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction, bool)`

GetWindowsMobileRestrictionOk returns a tuple with the WindowsMobileRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsMobileRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasWindowsMobileRestriction() bool`

HasWindowsMobileRestriction returns a boolean if a field has been set.

### SetWindowsMobileRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetWindowsMobileRestriction(v AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction)`

SetWindowsMobileRestriction gets a reference to the given AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction and assigns it to the WindowsMobileRestriction field.

### SetWindowsMobileRestrictionExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetWindowsMobileRestrictionExplicitNull(b bool)`

SetWindowsMobileRestrictionExplicitNull (un)sets WindowsMobileRestriction to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WindowsMobileRestriction value is set to nil even if false is passed
### GetAndroidRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetAndroidRestriction() AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction`

GetAndroidRestriction returns the AndroidRestriction field if non-nil, zero value otherwise.

### GetAndroidRestrictionOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetAndroidRestrictionOk() (AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction, bool)`

GetAndroidRestrictionOk returns a tuple with the AndroidRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAndroidRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasAndroidRestriction() bool`

HasAndroidRestriction returns a boolean if a field has been set.

### SetAndroidRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetAndroidRestriction(v AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction)`

SetAndroidRestriction gets a reference to the given AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction and assigns it to the AndroidRestriction field.

### SetAndroidRestrictionExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetAndroidRestrictionExplicitNull(b bool)`

SetAndroidRestrictionExplicitNull (un)sets AndroidRestriction to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AndroidRestriction value is set to nil even if false is passed
### GetMacOSRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetMacOSRestriction() AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction`

GetMacOSRestriction returns the MacOSRestriction field if non-nil, zero value otherwise.

### GetMacOSRestrictionOk

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) GetMacOSRestrictionOk() (AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction, bool)`

GetMacOSRestrictionOk returns a tuple with the MacOSRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMacOSRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) HasMacOSRestriction() bool`

HasMacOSRestriction returns a boolean if a field has been set.

### SetMacOSRestriction

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetMacOSRestriction(v AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction)`

SetMacOSRestriction gets a reference to the given AnyOfmicrosoftGraphDeviceEnrollmentPlatformRestriction and assigns it to the MacOSRestriction field.

### SetMacOSRestrictionExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentPlatformRestrictionsConfiguration) SetMacOSRestrictionExplicitNull(b bool)`

SetMacOSRestrictionExplicitNull (un)sets MacOSRestriction to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MacOSRestriction value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


