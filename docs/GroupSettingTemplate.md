# GroupSettingTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Values** | Pointer to [**[]MicrosoftGraphSettingTemplateValue**](microsoft.graph.settingTemplateValue.md) |  | [optional] 

## Methods

### GetDisplayName

`func (o *GroupSettingTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GroupSettingTemplate) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *GroupSettingTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *GroupSettingTemplate) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *GroupSettingTemplate) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *GroupSettingTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupSettingTemplate) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *GroupSettingTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *GroupSettingTemplate) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *GroupSettingTemplate) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetValues

`func (o *GroupSettingTemplate) GetValues() []MicrosoftGraphSettingTemplateValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GroupSettingTemplate) GetValuesOk() ([]MicrosoftGraphSettingTemplateValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *GroupSettingTemplate) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *GroupSettingTemplate) SetValues(v []MicrosoftGraphSettingTemplateValue)`

SetValues gets a reference to the given []MicrosoftGraphSettingTemplateValue and assigns it to the Values field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


