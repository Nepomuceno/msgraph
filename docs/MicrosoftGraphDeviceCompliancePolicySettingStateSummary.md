# MicrosoftGraphDeviceCompliancePolicySettingStateSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Setting** | Pointer to **string** | The setting class name and property name. | [optional] 
**SettingName** | Pointer to **string** | Name of the setting. | [optional] 
**PlatformType** | Pointer to [**AnyOfmicrosoftGraphPolicyPlatformType**](anyOf&lt;microsoft.graph.policyPlatformType&gt;.md) | Setting platform | [optional] 
**UnknownDeviceCount** | Pointer to **int32** | Number of unknown devices | [optional] 
**NotApplicableDeviceCount** | Pointer to **int32** | Number of not applicable devices | [optional] 
**CompliantDeviceCount** | Pointer to **int32** | Number of compliant devices | [optional] 
**RemediatedDeviceCount** | Pointer to **int32** | Number of remediated devices | [optional] 
**NonCompliantDeviceCount** | Pointer to **int32** | Number of NonCompliant devices | [optional] 
**ErrorDeviceCount** | Pointer to **int32** | Number of error devices | [optional] 
**ConflictDeviceCount** | Pointer to **int32** | Number of conflict devices | [optional] 
**DeviceComplianceSettingStates** | Pointer to [**[]MicrosoftGraphDeviceComplianceSettingState**](microsoft.graph.deviceComplianceSettingState.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetSetting

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetSettingOk() (string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSetting

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### SetSetting

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetSetting(v string)`

SetSetting gets a reference to the given string and assigns it to the Setting field.

### SetSettingExplicitNull

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetSettingExplicitNull(b bool)`

SetSettingExplicitNull (un)sets Setting to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Setting value is set to nil even if false is passed
### GetSettingName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetSettingNameOk() (string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetSettingName(v string)`

SetSettingName gets a reference to the given string and assigns it to the SettingName field.

### SetSettingNameExplicitNull

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetSettingNameExplicitNull(b bool)`

SetSettingNameExplicitNull (un)sets SettingName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingName value is set to nil even if false is passed
### GetPlatformType

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetPlatformType() AnyOfmicrosoftGraphPolicyPlatformType`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetPlatformTypeOk() (AnyOfmicrosoftGraphPolicyPlatformType, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlatformType

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### SetPlatformType

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetPlatformType(v AnyOfmicrosoftGraphPolicyPlatformType)`

SetPlatformType gets a reference to the given AnyOfmicrosoftGraphPolicyPlatformType and assigns it to the PlatformType field.

### GetUnknownDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetUnknownDeviceCountOk() (int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnknownDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.

### SetUnknownDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount gets a reference to the given int32 and assigns it to the UnknownDeviceCount field.

### GetNotApplicableDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetNotApplicableDeviceCountOk() (int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotApplicableDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### SetNotApplicableDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount gets a reference to the given int32 and assigns it to the NotApplicableDeviceCount field.

### GetCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetCompliantDeviceCountOk() (int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### SetCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount gets a reference to the given int32 and assigns it to the CompliantDeviceCount field.

### GetRemediatedDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetRemediatedDeviceCountOk() (int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemediatedDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### SetRemediatedDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount gets a reference to the given int32 and assigns it to the RemediatedDeviceCount field.

### GetNonCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetNonCompliantDeviceCountOk() (int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNonCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### SetNonCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount gets a reference to the given int32 and assigns it to the NonCompliantDeviceCount field.

### GetErrorDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetErrorDeviceCountOk() (int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### SetErrorDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount gets a reference to the given int32 and assigns it to the ErrorDeviceCount field.

### GetConflictDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetConflictDeviceCountOk() (int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConflictDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### SetConflictDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount gets a reference to the given int32 and assigns it to the ConflictDeviceCount field.

### GetDeviceComplianceSettingStates

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetDeviceComplianceSettingStates() []MicrosoftGraphDeviceComplianceSettingState`

GetDeviceComplianceSettingStates returns the DeviceComplianceSettingStates field if non-nil, zero value otherwise.

### GetDeviceComplianceSettingStatesOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetDeviceComplianceSettingStatesOk() ([]MicrosoftGraphDeviceComplianceSettingState, bool)`

GetDeviceComplianceSettingStatesOk returns a tuple with the DeviceComplianceSettingStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceComplianceSettingStates

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasDeviceComplianceSettingStates() bool`

HasDeviceComplianceSettingStates returns a boolean if a field has been set.

### SetDeviceComplianceSettingStates

`func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetDeviceComplianceSettingStates(v []MicrosoftGraphDeviceComplianceSettingState)`

SetDeviceComplianceSettingStates gets a reference to the given []MicrosoftGraphDeviceComplianceSettingState and assigns it to the DeviceComplianceSettingStates field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


