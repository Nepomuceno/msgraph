# MicrosoftGraphDeviceEnrollmentLimitConfiguration

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
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPriority

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetPriorityOk() (int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPriority

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriority

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) SetPriority(v int32)`

SetPriority gets a reference to the given int32 and assigns it to the Priority field.

### GetCreatedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetVersion

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetAssignments() []MicrosoftGraphEnrollmentConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetAssignmentsOk() ([]MicrosoftGraphEnrollmentConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) SetAssignments(v []MicrosoftGraphEnrollmentConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphEnrollmentConfigurationAssignment and assigns it to the Assignments field.

### GetLimit

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) GetLimitOk() (int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimit

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimit

`func (o *MicrosoftGraphDeviceEnrollmentLimitConfiguration) SetLimit(v int32)`

SetLimit gets a reference to the given int32 and assigns it to the Limit field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


