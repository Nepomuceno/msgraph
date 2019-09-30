# SettingStateDeviceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingName** | Pointer to **string** | Name of the setting | [optional] 
**InstancePath** | Pointer to **string** | Name of the InstancePath for the setting | [optional] 
**UnknownDeviceCount** | Pointer to **int32** | Device Unkown count for the setting | [optional] 
**NotApplicableDeviceCount** | Pointer to **int32** | Device Not Applicable count for the setting | [optional] 
**CompliantDeviceCount** | Pointer to **int32** | Device Compliant count for the setting | [optional] 
**RemediatedDeviceCount** | Pointer to **int32** | Device Compliant count for the setting | [optional] 
**NonCompliantDeviceCount** | Pointer to **int32** | Device NonCompliant count for the setting | [optional] 
**ErrorDeviceCount** | Pointer to **int32** | Device error count for the setting | [optional] 
**ConflictDeviceCount** | Pointer to **int32** | Device conflict error count for the setting | [optional] 

## Methods

### GetSettingName

`func (o *SettingStateDeviceSummary) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *SettingStateDeviceSummary) GetSettingNameOk() (string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingName

`func (o *SettingStateDeviceSummary) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingName

`func (o *SettingStateDeviceSummary) SetSettingName(v string)`

SetSettingName gets a reference to the given string and assigns it to the SettingName field.

### SetSettingNameExplicitNull

`func (o *SettingStateDeviceSummary) SetSettingNameExplicitNull(b bool)`

SetSettingNameExplicitNull (un)sets SettingName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingName value is set to nil even if false is passed
### GetInstancePath

`func (o *SettingStateDeviceSummary) GetInstancePath() string`

GetInstancePath returns the InstancePath field if non-nil, zero value otherwise.

### GetInstancePathOk

`func (o *SettingStateDeviceSummary) GetInstancePathOk() (string, bool)`

GetInstancePathOk returns a tuple with the InstancePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstancePath

`func (o *SettingStateDeviceSummary) HasInstancePath() bool`

HasInstancePath returns a boolean if a field has been set.

### SetInstancePath

`func (o *SettingStateDeviceSummary) SetInstancePath(v string)`

SetInstancePath gets a reference to the given string and assigns it to the InstancePath field.

### SetInstancePathExplicitNull

`func (o *SettingStateDeviceSummary) SetInstancePathExplicitNull(b bool)`

SetInstancePathExplicitNull (un)sets InstancePath to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InstancePath value is set to nil even if false is passed
### GetUnknownDeviceCount

`func (o *SettingStateDeviceSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *SettingStateDeviceSummary) GetUnknownDeviceCountOk() (int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnknownDeviceCount

`func (o *SettingStateDeviceSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.

### SetUnknownDeviceCount

`func (o *SettingStateDeviceSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount gets a reference to the given int32 and assigns it to the UnknownDeviceCount field.

### GetNotApplicableDeviceCount

`func (o *SettingStateDeviceSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *SettingStateDeviceSummary) GetNotApplicableDeviceCountOk() (int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotApplicableDeviceCount

`func (o *SettingStateDeviceSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### SetNotApplicableDeviceCount

`func (o *SettingStateDeviceSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount gets a reference to the given int32 and assigns it to the NotApplicableDeviceCount field.

### GetCompliantDeviceCount

`func (o *SettingStateDeviceSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *SettingStateDeviceSummary) GetCompliantDeviceCountOk() (int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantDeviceCount

`func (o *SettingStateDeviceSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### SetCompliantDeviceCount

`func (o *SettingStateDeviceSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount gets a reference to the given int32 and assigns it to the CompliantDeviceCount field.

### GetRemediatedDeviceCount

`func (o *SettingStateDeviceSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *SettingStateDeviceSummary) GetRemediatedDeviceCountOk() (int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemediatedDeviceCount

`func (o *SettingStateDeviceSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### SetRemediatedDeviceCount

`func (o *SettingStateDeviceSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount gets a reference to the given int32 and assigns it to the RemediatedDeviceCount field.

### GetNonCompliantDeviceCount

`func (o *SettingStateDeviceSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *SettingStateDeviceSummary) GetNonCompliantDeviceCountOk() (int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNonCompliantDeviceCount

`func (o *SettingStateDeviceSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### SetNonCompliantDeviceCount

`func (o *SettingStateDeviceSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount gets a reference to the given int32 and assigns it to the NonCompliantDeviceCount field.

### GetErrorDeviceCount

`func (o *SettingStateDeviceSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *SettingStateDeviceSummary) GetErrorDeviceCountOk() (int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorDeviceCount

`func (o *SettingStateDeviceSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### SetErrorDeviceCount

`func (o *SettingStateDeviceSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount gets a reference to the given int32 and assigns it to the ErrorDeviceCount field.

### GetConflictDeviceCount

`func (o *SettingStateDeviceSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *SettingStateDeviceSummary) GetConflictDeviceCountOk() (int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConflictDeviceCount

`func (o *SettingStateDeviceSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### SetConflictDeviceCount

`func (o *SettingStateDeviceSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount gets a reference to the given int32 and assigns it to the ConflictDeviceCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


