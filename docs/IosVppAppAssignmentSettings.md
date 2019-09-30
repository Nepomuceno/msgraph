# IosVppAppAssignmentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseDeviceLicensing** | Pointer to **bool** | Whether or not to use device licensing. | [optional] 
**VpnConfigurationId** | Pointer to **string** | The VPN Configuration Id to apply for this app. | [optional] 

## Methods

### GetUseDeviceLicensing

`func (o *IosVppAppAssignmentSettings) GetUseDeviceLicensing() bool`

GetUseDeviceLicensing returns the UseDeviceLicensing field if non-nil, zero value otherwise.

### GetUseDeviceLicensingOk

`func (o *IosVppAppAssignmentSettings) GetUseDeviceLicensingOk() (bool, bool)`

GetUseDeviceLicensingOk returns a tuple with the UseDeviceLicensing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUseDeviceLicensing

`func (o *IosVppAppAssignmentSettings) HasUseDeviceLicensing() bool`

HasUseDeviceLicensing returns a boolean if a field has been set.

### SetUseDeviceLicensing

`func (o *IosVppAppAssignmentSettings) SetUseDeviceLicensing(v bool)`

SetUseDeviceLicensing gets a reference to the given bool and assigns it to the UseDeviceLicensing field.

### GetVpnConfigurationId

`func (o *IosVppAppAssignmentSettings) GetVpnConfigurationId() string`

GetVpnConfigurationId returns the VpnConfigurationId field if non-nil, zero value otherwise.

### GetVpnConfigurationIdOk

`func (o *IosVppAppAssignmentSettings) GetVpnConfigurationIdOk() (string, bool)`

GetVpnConfigurationIdOk returns a tuple with the VpnConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVpnConfigurationId

`func (o *IosVppAppAssignmentSettings) HasVpnConfigurationId() bool`

HasVpnConfigurationId returns a boolean if a field has been set.

### SetVpnConfigurationId

`func (o *IosVppAppAssignmentSettings) SetVpnConfigurationId(v string)`

SetVpnConfigurationId gets a reference to the given string and assigns it to the VpnConfigurationId field.

### SetVpnConfigurationIdExplicitNull

`func (o *IosVppAppAssignmentSettings) SetVpnConfigurationIdExplicitNull(b bool)`

SetVpnConfigurationIdExplicitNull (un)sets VpnConfigurationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VpnConfigurationId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


