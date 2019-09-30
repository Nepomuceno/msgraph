# DeviceCompliancePolicyState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingStates** | Pointer to [**[]AnyOfmicrosoftGraphDeviceCompliancePolicySettingState**](anyOf&lt;microsoft.graph.deviceCompliancePolicySettingState&gt;.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The name of the policy for this policyBase | [optional] 
**Version** | Pointer to **int32** | The version of the policy | [optional] 
**PlatformType** | Pointer to [**AnyOfmicrosoftGraphPolicyPlatformType**](anyOf&lt;microsoft.graph.policyPlatformType&gt;.md) | Platform type that the policy applies to | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | The compliance state of the policy | [optional] 
**SettingCount** | Pointer to **int32** | Count of how many setting a policy holds | [optional] 

## Methods

### GetSettingStates

`func (o *DeviceCompliancePolicyState) GetSettingStates() []AnyOfmicrosoftGraphDeviceCompliancePolicySettingState`

GetSettingStates returns the SettingStates field if non-nil, zero value otherwise.

### GetSettingStatesOk

`func (o *DeviceCompliancePolicyState) GetSettingStatesOk() ([]AnyOfmicrosoftGraphDeviceCompliancePolicySettingState, bool)`

GetSettingStatesOk returns a tuple with the SettingStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingStates

`func (o *DeviceCompliancePolicyState) HasSettingStates() bool`

HasSettingStates returns a boolean if a field has been set.

### SetSettingStates

`func (o *DeviceCompliancePolicyState) SetSettingStates(v []AnyOfmicrosoftGraphDeviceCompliancePolicySettingState)`

SetSettingStates gets a reference to the given []AnyOfmicrosoftGraphDeviceCompliancePolicySettingState and assigns it to the SettingStates field.

### GetDisplayName

`func (o *DeviceCompliancePolicyState) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceCompliancePolicyState) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *DeviceCompliancePolicyState) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *DeviceCompliancePolicyState) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *DeviceCompliancePolicyState) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetVersion

`func (o *DeviceCompliancePolicyState) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceCompliancePolicyState) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *DeviceCompliancePolicyState) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *DeviceCompliancePolicyState) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetPlatformType

`func (o *DeviceCompliancePolicyState) GetPlatformType() AnyOfmicrosoftGraphPolicyPlatformType`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *DeviceCompliancePolicyState) GetPlatformTypeOk() (AnyOfmicrosoftGraphPolicyPlatformType, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlatformType

`func (o *DeviceCompliancePolicyState) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### SetPlatformType

`func (o *DeviceCompliancePolicyState) SetPlatformType(v AnyOfmicrosoftGraphPolicyPlatformType)`

SetPlatformType gets a reference to the given AnyOfmicrosoftGraphPolicyPlatformType and assigns it to the PlatformType field.

### GetState

`func (o *DeviceCompliancePolicyState) GetState() AnyOfmicrosoftGraphComplianceStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeviceCompliancePolicyState) GetStateOk() (AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *DeviceCompliancePolicyState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *DeviceCompliancePolicyState) SetState(v AnyOfmicrosoftGraphComplianceStatus)`

SetState gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the State field.

### GetSettingCount

`func (o *DeviceCompliancePolicyState) GetSettingCount() int32`

GetSettingCount returns the SettingCount field if non-nil, zero value otherwise.

### GetSettingCountOk

`func (o *DeviceCompliancePolicyState) GetSettingCountOk() (int32, bool)`

GetSettingCountOk returns a tuple with the SettingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingCount

`func (o *DeviceCompliancePolicyState) HasSettingCount() bool`

HasSettingCount returns a boolean if a field has been set.

### SetSettingCount

`func (o *DeviceCompliancePolicyState) SetSettingCount(v int32)`

SetSettingCount gets a reference to the given int32 and assigns it to the SettingCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


