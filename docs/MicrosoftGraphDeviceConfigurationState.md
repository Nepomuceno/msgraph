# MicrosoftGraphDeviceConfigurationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SettingStates** | Pointer to [**[]AnyOfmicrosoftGraphDeviceConfigurationSettingState**](anyOf&lt;microsoft.graph.deviceConfigurationSettingState&gt;.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The name of the policy for this policyBase | [optional] 
**Version** | Pointer to **int32** | The version of the policy | [optional] 
**PlatformType** | Pointer to [**AnyOfmicrosoftGraphPolicyPlatformType**](anyOf&lt;microsoft.graph.policyPlatformType&gt;.md) | Platform type that the policy applies to | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | The compliance state of the policy | [optional] 
**SettingCount** | Pointer to **int32** | Count of how many setting a policy holds | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceConfigurationState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceConfigurationState) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceConfigurationState) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceConfigurationState) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetSettingStates

`func (o *MicrosoftGraphDeviceConfigurationState) GetSettingStates() []AnyOfmicrosoftGraphDeviceConfigurationSettingState`

GetSettingStates returns the SettingStates field if non-nil, zero value otherwise.

### GetSettingStatesOk

`func (o *MicrosoftGraphDeviceConfigurationState) GetSettingStatesOk() ([]AnyOfmicrosoftGraphDeviceConfigurationSettingState, bool)`

GetSettingStatesOk returns a tuple with the SettingStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingStates

`func (o *MicrosoftGraphDeviceConfigurationState) HasSettingStates() bool`

HasSettingStates returns a boolean if a field has been set.

### SetSettingStates

`func (o *MicrosoftGraphDeviceConfigurationState) SetSettingStates(v []AnyOfmicrosoftGraphDeviceConfigurationSettingState)`

SetSettingStates gets a reference to the given []AnyOfmicrosoftGraphDeviceConfigurationSettingState and assigns it to the SettingStates field.

### GetDisplayName

`func (o *MicrosoftGraphDeviceConfigurationState) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDeviceConfigurationState) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphDeviceConfigurationState) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphDeviceConfigurationState) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphDeviceConfigurationState) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetVersion

`func (o *MicrosoftGraphDeviceConfigurationState) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphDeviceConfigurationState) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphDeviceConfigurationState) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphDeviceConfigurationState) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetPlatformType

`func (o *MicrosoftGraphDeviceConfigurationState) GetPlatformType() AnyOfmicrosoftGraphPolicyPlatformType`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *MicrosoftGraphDeviceConfigurationState) GetPlatformTypeOk() (AnyOfmicrosoftGraphPolicyPlatformType, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlatformType

`func (o *MicrosoftGraphDeviceConfigurationState) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### SetPlatformType

`func (o *MicrosoftGraphDeviceConfigurationState) SetPlatformType(v AnyOfmicrosoftGraphPolicyPlatformType)`

SetPlatformType gets a reference to the given AnyOfmicrosoftGraphPolicyPlatformType and assigns it to the PlatformType field.

### GetState

`func (o *MicrosoftGraphDeviceConfigurationState) GetState() AnyOfmicrosoftGraphComplianceStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphDeviceConfigurationState) GetStateOk() (AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *MicrosoftGraphDeviceConfigurationState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *MicrosoftGraphDeviceConfigurationState) SetState(v AnyOfmicrosoftGraphComplianceStatus)`

SetState gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the State field.

### GetSettingCount

`func (o *MicrosoftGraphDeviceConfigurationState) GetSettingCount() int32`

GetSettingCount returns the SettingCount field if non-nil, zero value otherwise.

### GetSettingCountOk

`func (o *MicrosoftGraphDeviceConfigurationState) GetSettingCountOk() (int32, bool)`

GetSettingCountOk returns a tuple with the SettingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingCount

`func (o *MicrosoftGraphDeviceConfigurationState) HasSettingCount() bool`

HasSettingCount returns a boolean if a field has been set.

### SetSettingCount

`func (o *MicrosoftGraphDeviceConfigurationState) SetSettingCount(v int32)`

SetSettingCount gets a reference to the given int32 and assigns it to the SettingCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


