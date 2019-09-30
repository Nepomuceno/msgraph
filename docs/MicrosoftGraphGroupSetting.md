# MicrosoftGraphGroupSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] 
**Values** | Pointer to [**[]MicrosoftGraphSettingValue**](microsoft.graph.settingValue.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphGroupSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphGroupSetting) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphGroupSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphGroupSetting) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphGroupSetting) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphGroupSetting) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphGroupSetting) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphGroupSetting) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphGroupSetting) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetTemplateId

`func (o *MicrosoftGraphGroupSetting) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *MicrosoftGraphGroupSetting) GetTemplateIdOk() (string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTemplateId

`func (o *MicrosoftGraphGroupSetting) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateId

`func (o *MicrosoftGraphGroupSetting) SetTemplateId(v string)`

SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.

### SetTemplateIdExplicitNull

`func (o *MicrosoftGraphGroupSetting) SetTemplateIdExplicitNull(b bool)`

SetTemplateIdExplicitNull (un)sets TemplateId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TemplateId value is set to nil even if false is passed
### GetValues

`func (o *MicrosoftGraphGroupSetting) GetValues() []MicrosoftGraphSettingValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MicrosoftGraphGroupSetting) GetValuesOk() ([]MicrosoftGraphSettingValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *MicrosoftGraphGroupSetting) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *MicrosoftGraphGroupSetting) SetValues(v []MicrosoftGraphSettingValue)`

SetValues gets a reference to the given []MicrosoftGraphSettingValue and assigns it to the Values field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


