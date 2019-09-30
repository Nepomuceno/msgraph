# MicrosoftGraphDeviceCompliancePolicyState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SettingStates** | Pointer to [**[]AnyOfmicrosoftGraphDeviceCompliancePolicySettingState**](anyOf&lt;microsoft.graph.deviceCompliancePolicySettingState&gt;.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The name of the policy for this policyBase | [optional] 
**Version** | Pointer to **int32** | The version of the policy | [optional] 
**PlatformType** | Pointer to [**AnyOfmicrosoftGraphPolicyPlatformType**](anyOf&lt;microsoft.graph.policyPlatformType&gt;.md) | Platform type that the policy applies to | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | The compliance state of the policy | [optional] 
**SettingCount** | Pointer to **int32** | Count of how many setting a policy holds | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetSettingStates

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetSettingStates() []AnyOfmicrosoftGraphDeviceCompliancePolicySettingState`

GetSettingStates returns the SettingStates field if non-nil, zero value otherwise.

### GetSettingStatesOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetSettingStatesOk() ([]AnyOfmicrosoftGraphDeviceCompliancePolicySettingState, bool)`

GetSettingStatesOk returns a tuple with the SettingStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingStates

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasSettingStates() bool`

HasSettingStates returns a boolean if a field has been set.

### SetSettingStates

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetSettingStates(v []AnyOfmicrosoftGraphDeviceCompliancePolicySettingState)`

SetSettingStates gets a reference to the given []AnyOfmicrosoftGraphDeviceCompliancePolicySettingState and assigns it to the SettingStates field.

### GetDisplayName

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetVersion

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetPlatformType

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetPlatformType() AnyOfmicrosoftGraphPolicyPlatformType`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetPlatformTypeOk() (AnyOfmicrosoftGraphPolicyPlatformType, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlatformType

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### SetPlatformType

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetPlatformType(v AnyOfmicrosoftGraphPolicyPlatformType)`

SetPlatformType gets a reference to the given AnyOfmicrosoftGraphPolicyPlatformType and assigns it to the PlatformType field.

### GetState

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetState() AnyOfmicrosoftGraphComplianceStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetStateOk() (AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetState(v AnyOfmicrosoftGraphComplianceStatus)`

SetState gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the State field.

### GetSettingCount

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetSettingCount() int32`

GetSettingCount returns the SettingCount field if non-nil, zero value otherwise.

### GetSettingCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyState) GetSettingCountOk() (int32, bool)`

GetSettingCountOk returns a tuple with the SettingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingCount

`func (o *MicrosoftGraphDeviceCompliancePolicyState) HasSettingCount() bool`

HasSettingCount returns a boolean if a field has been set.

### SetSettingCount

`func (o *MicrosoftGraphDeviceCompliancePolicyState) SetSettingCount(v int32)`

SetSettingCount gets a reference to the given int32 and assigns it to the SettingCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


