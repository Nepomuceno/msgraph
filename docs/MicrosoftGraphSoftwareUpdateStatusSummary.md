# MicrosoftGraphSoftwareUpdateStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetCompliantDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetCompliantDeviceCountOk() (int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### SetCompliantDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount gets a reference to the given int32 and assigns it to the CompliantDeviceCount field.

### GetNonCompliantDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNonCompliantDeviceCountOk() (int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNonCompliantDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### SetNonCompliantDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount gets a reference to the given int32 and assigns it to the NonCompliantDeviceCount field.

### GetRemediatedDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetRemediatedDeviceCountOk() (int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemediatedDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### SetRemediatedDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount gets a reference to the given int32 and assigns it to the RemediatedDeviceCount field.

### GetErrorDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetErrorDeviceCountOk() (int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### SetErrorDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount gets a reference to the given int32 and assigns it to the ErrorDeviceCount field.

### GetUnknownDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetUnknownDeviceCountOk() (int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnknownDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.

### SetUnknownDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount gets a reference to the given int32 and assigns it to the UnknownDeviceCount field.

### GetConflictDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetConflictDeviceCountOk() (int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConflictDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### SetConflictDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount gets a reference to the given int32 and assigns it to the ConflictDeviceCount field.

### GetNotApplicableDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNotApplicableDeviceCountOk() (int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotApplicableDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### SetNotApplicableDeviceCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount gets a reference to the given int32 and assigns it to the NotApplicableDeviceCount field.

### GetCompliantUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetCompliantUserCount() int32`

GetCompliantUserCount returns the CompliantUserCount field if non-nil, zero value otherwise.

### GetCompliantUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetCompliantUserCountOk() (int32, bool)`

GetCompliantUserCountOk returns a tuple with the CompliantUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasCompliantUserCount() bool`

HasCompliantUserCount returns a boolean if a field has been set.

### SetCompliantUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetCompliantUserCount(v int32)`

SetCompliantUserCount gets a reference to the given int32 and assigns it to the CompliantUserCount field.

### GetNonCompliantUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNonCompliantUserCount() int32`

GetNonCompliantUserCount returns the NonCompliantUserCount field if non-nil, zero value otherwise.

### GetNonCompliantUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNonCompliantUserCountOk() (int32, bool)`

GetNonCompliantUserCountOk returns a tuple with the NonCompliantUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNonCompliantUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasNonCompliantUserCount() bool`

HasNonCompliantUserCount returns a boolean if a field has been set.

### SetNonCompliantUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetNonCompliantUserCount(v int32)`

SetNonCompliantUserCount gets a reference to the given int32 and assigns it to the NonCompliantUserCount field.

### GetRemediatedUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetRemediatedUserCount() int32`

GetRemediatedUserCount returns the RemediatedUserCount field if non-nil, zero value otherwise.

### GetRemediatedUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetRemediatedUserCountOk() (int32, bool)`

GetRemediatedUserCountOk returns a tuple with the RemediatedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemediatedUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasRemediatedUserCount() bool`

HasRemediatedUserCount returns a boolean if a field has been set.

### SetRemediatedUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetRemediatedUserCount(v int32)`

SetRemediatedUserCount gets a reference to the given int32 and assigns it to the RemediatedUserCount field.

### GetErrorUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetErrorUserCount() int32`

GetErrorUserCount returns the ErrorUserCount field if non-nil, zero value otherwise.

### GetErrorUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetErrorUserCountOk() (int32, bool)`

GetErrorUserCountOk returns a tuple with the ErrorUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasErrorUserCount() bool`

HasErrorUserCount returns a boolean if a field has been set.

### SetErrorUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetErrorUserCount(v int32)`

SetErrorUserCount gets a reference to the given int32 and assigns it to the ErrorUserCount field.

### GetUnknownUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetUnknownUserCount() int32`

GetUnknownUserCount returns the UnknownUserCount field if non-nil, zero value otherwise.

### GetUnknownUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetUnknownUserCountOk() (int32, bool)`

GetUnknownUserCountOk returns a tuple with the UnknownUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnknownUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasUnknownUserCount() bool`

HasUnknownUserCount returns a boolean if a field has been set.

### SetUnknownUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetUnknownUserCount(v int32)`

SetUnknownUserCount gets a reference to the given int32 and assigns it to the UnknownUserCount field.

### GetConflictUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetConflictUserCount() int32`

GetConflictUserCount returns the ConflictUserCount field if non-nil, zero value otherwise.

### GetConflictUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetConflictUserCountOk() (int32, bool)`

GetConflictUserCountOk returns a tuple with the ConflictUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConflictUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasConflictUserCount() bool`

HasConflictUserCount returns a boolean if a field has been set.

### SetConflictUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetConflictUserCount(v int32)`

SetConflictUserCount gets a reference to the given int32 and assigns it to the ConflictUserCount field.

### GetNotApplicableUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNotApplicableUserCount() int32`

GetNotApplicableUserCount returns the NotApplicableUserCount field if non-nil, zero value otherwise.

### GetNotApplicableUserCountOk

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) GetNotApplicableUserCountOk() (int32, bool)`

GetNotApplicableUserCountOk returns a tuple with the NotApplicableUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotApplicableUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) HasNotApplicableUserCount() bool`

HasNotApplicableUserCount returns a boolean if a field has been set.

### SetNotApplicableUserCount

`func (o *MicrosoftGraphSoftwareUpdateStatusSummary) SetNotApplicableUserCount(v int32)`

SetNotApplicableUserCount gets a reference to the given int32 and assigns it to the NotApplicableUserCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


