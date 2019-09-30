# DeviceComplianceScheduledActionForRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleName** | Pointer to **string** | Name of the rule which this scheduled action applies to. | [optional] 
**ScheduledActionConfigurations** | Pointer to [**[]MicrosoftGraphDeviceComplianceActionItem**](microsoft.graph.deviceComplianceActionItem.md) |  | [optional] 

## Methods

### GetRuleName

`func (o *DeviceComplianceScheduledActionForRule) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *DeviceComplianceScheduledActionForRule) GetRuleNameOk() (string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleName

`func (o *DeviceComplianceScheduledActionForRule) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### SetRuleName

`func (o *DeviceComplianceScheduledActionForRule) SetRuleName(v string)`

SetRuleName gets a reference to the given string and assigns it to the RuleName field.

### SetRuleNameExplicitNull

`func (o *DeviceComplianceScheduledActionForRule) SetRuleNameExplicitNull(b bool)`

SetRuleNameExplicitNull (un)sets RuleName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RuleName value is set to nil even if false is passed
### GetScheduledActionConfigurations

`func (o *DeviceComplianceScheduledActionForRule) GetScheduledActionConfigurations() []MicrosoftGraphDeviceComplianceActionItem`

GetScheduledActionConfigurations returns the ScheduledActionConfigurations field if non-nil, zero value otherwise.

### GetScheduledActionConfigurationsOk

`func (o *DeviceComplianceScheduledActionForRule) GetScheduledActionConfigurationsOk() ([]MicrosoftGraphDeviceComplianceActionItem, bool)`

GetScheduledActionConfigurationsOk returns a tuple with the ScheduledActionConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledActionConfigurations

`func (o *DeviceComplianceScheduledActionForRule) HasScheduledActionConfigurations() bool`

HasScheduledActionConfigurations returns a boolean if a field has been set.

### SetScheduledActionConfigurations

`func (o *DeviceComplianceScheduledActionForRule) SetScheduledActionConfigurations(v []MicrosoftGraphDeviceComplianceActionItem)`

SetScheduledActionConfigurations gets a reference to the given []MicrosoftGraphDeviceComplianceActionItem and assigns it to the ScheduledActionConfigurations field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


