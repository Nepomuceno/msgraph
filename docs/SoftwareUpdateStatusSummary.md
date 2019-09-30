# SoftwareUpdateStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the policy. | [optional] 
**CompliantDeviceCount** | Pointer to **int32** | Number of compliant devices. | [optional] 
**NonCompliantDeviceCount** | Pointer to **int32** | Number of non compliant devices. | [optional] 
**RemediatedDeviceCount** | Pointer to **int32** | Number of remediated devices. | [optional] 
**ErrorDeviceCount** | Pointer to **int32** | Number of devices had error. | [optional] 
**UnknownDeviceCount** | Pointer to **int32** | Number of unknown devices. | [optional] 
**ConflictDeviceCount** | Pointer to **int32** | Number of conflict devices. | [optional] 
**NotApplicableDeviceCount** | Pointer to **int32** | Number of not applicable devices. | [optional] 
**CompliantUserCount** | Pointer to **int32** | Number of compliant users. | [optional] 
**NonCompliantUserCount** | Pointer to **int32** | Number of non compliant users. | [optional] 
**RemediatedUserCount** | Pointer to **int32** | Number of remediated users. | [optional] 
**ErrorUserCount** | Pointer to **int32** | Number of users had error. | [optional] 
**UnknownUserCount** | Pointer to **int32** | Number of unknown users. | [optional] 
**ConflictUserCount** | Pointer to **int32** | Number of conflict users. | [optional] 
**NotApplicableUserCount** | Pointer to **int32** | Number of not applicable users. | [optional] 

## Methods

### GetDisplayName

`func (o *SoftwareUpdateStatusSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SoftwareUpdateStatusSummary) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *SoftwareUpdateStatusSummary) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *SoftwareUpdateStatusSummary) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *SoftwareUpdateStatusSummary) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetCompliantDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetCompliantDeviceCountOk() (int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### SetCompliantDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount gets a reference to the given int32 and assigns it to the CompliantDeviceCount field.

### GetNonCompliantDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetNonCompliantDeviceCountOk() (int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNonCompliantDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### SetNonCompliantDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount gets a reference to the given int32 and assigns it to the NonCompliantDeviceCount field.

### GetRemediatedDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetRemediatedDeviceCountOk() (int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemediatedDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### SetRemediatedDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount gets a reference to the given int32 and assigns it to the RemediatedDeviceCount field.

### GetErrorDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetErrorDeviceCountOk() (int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### SetErrorDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount gets a reference to the given int32 and assigns it to the ErrorDeviceCount field.

### GetUnknownDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetUnknownDeviceCountOk() (int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnknownDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.

### SetUnknownDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount gets a reference to the given int32 and assigns it to the UnknownDeviceCount field.

### GetConflictDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetConflictDeviceCountOk() (int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConflictDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### SetConflictDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount gets a reference to the given int32 and assigns it to the ConflictDeviceCount field.

### GetNotApplicableDeviceCount

`func (o *SoftwareUpdateStatusSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *SoftwareUpdateStatusSummary) GetNotApplicableDeviceCountOk() (int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotApplicableDeviceCount

`func (o *SoftwareUpdateStatusSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### SetNotApplicableDeviceCount

`func (o *SoftwareUpdateStatusSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount gets a reference to the given int32 and assigns it to the NotApplicableDeviceCount field.

### GetCompliantUserCount

`func (o *SoftwareUpdateStatusSummary) GetCompliantUserCount() int32`

GetCompliantUserCount returns the CompliantUserCount field if non-nil, zero value otherwise.

### GetCompliantUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetCompliantUserCountOk() (int32, bool)`

GetCompliantUserCountOk returns a tuple with the CompliantUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantUserCount

`func (o *SoftwareUpdateStatusSummary) HasCompliantUserCount() bool`

HasCompliantUserCount returns a boolean if a field has been set.

### SetCompliantUserCount

`func (o *SoftwareUpdateStatusSummary) SetCompliantUserCount(v int32)`

SetCompliantUserCount gets a reference to the given int32 and assigns it to the CompliantUserCount field.

### GetNonCompliantUserCount

`func (o *SoftwareUpdateStatusSummary) GetNonCompliantUserCount() int32`

GetNonCompliantUserCount returns the NonCompliantUserCount field if non-nil, zero value otherwise.

### GetNonCompliantUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetNonCompliantUserCountOk() (int32, bool)`

GetNonCompliantUserCountOk returns a tuple with the NonCompliantUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNonCompliantUserCount

`func (o *SoftwareUpdateStatusSummary) HasNonCompliantUserCount() bool`

HasNonCompliantUserCount returns a boolean if a field has been set.

### SetNonCompliantUserCount

`func (o *SoftwareUpdateStatusSummary) SetNonCompliantUserCount(v int32)`

SetNonCompliantUserCount gets a reference to the given int32 and assigns it to the NonCompliantUserCount field.

### GetRemediatedUserCount

`func (o *SoftwareUpdateStatusSummary) GetRemediatedUserCount() int32`

GetRemediatedUserCount returns the RemediatedUserCount field if non-nil, zero value otherwise.

### GetRemediatedUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetRemediatedUserCountOk() (int32, bool)`

GetRemediatedUserCountOk returns a tuple with the RemediatedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemediatedUserCount

`func (o *SoftwareUpdateStatusSummary) HasRemediatedUserCount() bool`

HasRemediatedUserCount returns a boolean if a field has been set.

### SetRemediatedUserCount

`func (o *SoftwareUpdateStatusSummary) SetRemediatedUserCount(v int32)`

SetRemediatedUserCount gets a reference to the given int32 and assigns it to the RemediatedUserCount field.

### GetErrorUserCount

`func (o *SoftwareUpdateStatusSummary) GetErrorUserCount() int32`

GetErrorUserCount returns the ErrorUserCount field if non-nil, zero value otherwise.

### GetErrorUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetErrorUserCountOk() (int32, bool)`

GetErrorUserCountOk returns a tuple with the ErrorUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorUserCount

`func (o *SoftwareUpdateStatusSummary) HasErrorUserCount() bool`

HasErrorUserCount returns a boolean if a field has been set.

### SetErrorUserCount

`func (o *SoftwareUpdateStatusSummary) SetErrorUserCount(v int32)`

SetErrorUserCount gets a reference to the given int32 and assigns it to the ErrorUserCount field.

### GetUnknownUserCount

`func (o *SoftwareUpdateStatusSummary) GetUnknownUserCount() int32`

GetUnknownUserCount returns the UnknownUserCount field if non-nil, zero value otherwise.

### GetUnknownUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetUnknownUserCountOk() (int32, bool)`

GetUnknownUserCountOk returns a tuple with the UnknownUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnknownUserCount

`func (o *SoftwareUpdateStatusSummary) HasUnknownUserCount() bool`

HasUnknownUserCount returns a boolean if a field has been set.

### SetUnknownUserCount

`func (o *SoftwareUpdateStatusSummary) SetUnknownUserCount(v int32)`

SetUnknownUserCount gets a reference to the given int32 and assigns it to the UnknownUserCount field.

### GetConflictUserCount

`func (o *SoftwareUpdateStatusSummary) GetConflictUserCount() int32`

GetConflictUserCount returns the ConflictUserCount field if non-nil, zero value otherwise.

### GetConflictUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetConflictUserCountOk() (int32, bool)`

GetConflictUserCountOk returns a tuple with the ConflictUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConflictUserCount

`func (o *SoftwareUpdateStatusSummary) HasConflictUserCount() bool`

HasConflictUserCount returns a boolean if a field has been set.

### SetConflictUserCount

`func (o *SoftwareUpdateStatusSummary) SetConflictUserCount(v int32)`

SetConflictUserCount gets a reference to the given int32 and assigns it to the ConflictUserCount field.

### GetNotApplicableUserCount

`func (o *SoftwareUpdateStatusSummary) GetNotApplicableUserCount() int32`

GetNotApplicableUserCount returns the NotApplicableUserCount field if non-nil, zero value otherwise.

### GetNotApplicableUserCountOk

`func (o *SoftwareUpdateStatusSummary) GetNotApplicableUserCountOk() (int32, bool)`

GetNotApplicableUserCountOk returns a tuple with the NotApplicableUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotApplicableUserCount

`func (o *SoftwareUpdateStatusSummary) HasNotApplicableUserCount() bool`

HasNotApplicableUserCount returns a boolean if a field has been set.

### SetNotApplicableUserCount

`func (o *SoftwareUpdateStatusSummary) SetNotApplicableUserCount(v int32)`

SetNotApplicableUserCount gets a reference to the given int32 and assigns it to the NotApplicableUserCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


