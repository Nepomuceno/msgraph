# MicrosoftGraphDeviceComplianceScheduledActionForRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**RuleName** | Pointer to **string** | Name of the rule which this scheduled action applies to. | [optional] 
**ScheduledActionConfigurations** | Pointer to [**[]MicrosoftGraphDeviceComplianceActionItem**](microsoft.graph.deviceComplianceActionItem.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetRuleName

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) GetRuleNameOk() (string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRuleName

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### SetRuleName

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) SetRuleName(v string)`

SetRuleName gets a reference to the given string and assigns it to the RuleName field.

### SetRuleNameExplicitNull

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) SetRuleNameExplicitNull(b bool)`

SetRuleNameExplicitNull (un)sets RuleName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RuleName value is set to nil even if false is passed
### GetScheduledActionConfigurations

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) GetScheduledActionConfigurations() []MicrosoftGraphDeviceComplianceActionItem`

GetScheduledActionConfigurations returns the ScheduledActionConfigurations field if non-nil, zero value otherwise.

### GetScheduledActionConfigurationsOk

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) GetScheduledActionConfigurationsOk() ([]MicrosoftGraphDeviceComplianceActionItem, bool)`

GetScheduledActionConfigurationsOk returns a tuple with the ScheduledActionConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledActionConfigurations

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) HasScheduledActionConfigurations() bool`

HasScheduledActionConfigurations returns a boolean if a field has been set.

### SetScheduledActionConfigurations

`func (o *MicrosoftGraphDeviceComplianceScheduledActionForRule) SetScheduledActionConfigurations(v []MicrosoftGraphDeviceComplianceActionItem)`

SetScheduledActionConfigurations gets a reference to the given []MicrosoftGraphDeviceComplianceActionItem and assigns it to the ScheduledActionConfigurations field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


