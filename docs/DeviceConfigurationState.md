# DeviceConfigurationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingStates** | Pointer to [**[]AnyOfmicrosoftGraphDeviceConfigurationSettingState**](anyOf&lt;microsoft.graph.deviceConfigurationSettingState&gt;.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The name of the policy for this policyBase | [optional] 
**Version** | Pointer to **int32** | The version of the policy | [optional] 
**PlatformType** | Pointer to [**AnyOfmicrosoftGraphPolicyPlatformType**](anyOf&lt;microsoft.graph.policyPlatformType&gt;.md) | Platform type that the policy applies to | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | The compliance state of the policy | [optional] 
**SettingCount** | Pointer to **int32** | Count of how many setting a policy holds | [optional] 

## Methods

### GetSettingStates

`func (o *DeviceConfigurationState) GetSettingStates() []AnyOfmicrosoftGraphDeviceConfigurationSettingState`

GetSettingStates returns the SettingStates field if non-nil, zero value otherwise.

### GetSettingStatesOk

`func (o *DeviceConfigurationState) GetSettingStatesOk() ([]AnyOfmicrosoftGraphDeviceConfigurationSettingState, bool)`

GetSettingStatesOk returns a tuple with the SettingStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingStates

`func (o *DeviceConfigurationState) HasSettingStates() bool`

HasSettingStates returns a boolean if a field has been set.

### SetSettingStates

`func (o *DeviceConfigurationState) SetSettingStates(v []AnyOfmicrosoftGraphDeviceConfigurationSettingState)`

SetSettingStates gets a reference to the given []AnyOfmicrosoftGraphDeviceConfigurationSettingState and assigns it to the SettingStates field.

### GetDisplayName

`func (o *DeviceConfigurationState) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceConfigurationState) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *DeviceConfigurationState) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *DeviceConfigurationState) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *DeviceConfigurationState) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetVersion

`func (o *DeviceConfigurationState) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceConfigurationState) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *DeviceConfigurationState) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *DeviceConfigurationState) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetPlatformType

`func (o *DeviceConfigurationState) GetPlatformType() AnyOfmicrosoftGraphPolicyPlatformType`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *DeviceConfigurationState) GetPlatformTypeOk() (AnyOfmicrosoftGraphPolicyPlatformType, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlatformType

`func (o *DeviceConfigurationState) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### SetPlatformType

`func (o *DeviceConfigurationState) SetPlatformType(v AnyOfmicrosoftGraphPolicyPlatformType)`

SetPlatformType gets a reference to the given AnyOfmicrosoftGraphPolicyPlatformType and assigns it to the PlatformType field.

### GetState

`func (o *DeviceConfigurationState) GetState() AnyOfmicrosoftGraphComplianceStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeviceConfigurationState) GetStateOk() (AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *DeviceConfigurationState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *DeviceConfigurationState) SetState(v AnyOfmicrosoftGraphComplianceStatus)`

SetState gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the State field.

### GetSettingCount

`func (o *DeviceConfigurationState) GetSettingCount() int32`

GetSettingCount returns the SettingCount field if non-nil, zero value otherwise.

### GetSettingCountOk

`func (o *DeviceConfigurationState) GetSettingCountOk() (int32, bool)`

GetSettingCountOk returns a tuple with the SettingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingCount

`func (o *DeviceConfigurationState) HasSettingCount() bool`

HasSettingCount returns a boolean if a field has been set.

### SetSettingCount

`func (o *DeviceConfigurationState) SetSettingCount(v int32)`

SetSettingCount gets a reference to the given int32 and assigns it to the SettingCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


