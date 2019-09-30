# IosMobileAppConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodedSettingXml** | Pointer to **string** | mdm app configuration Base64 binary. | [optional] 
**Settings** | Pointer to [**[]AnyOfmicrosoftGraphAppConfigurationSettingItem**](anyOf&lt;microsoft.graph.appConfigurationSettingItem&gt;.md) | app configuration setting items. | [optional] 

## Methods

### GetEncodedSettingXml

`func (o *IosMobileAppConfiguration) GetEncodedSettingXml() string`

GetEncodedSettingXml returns the EncodedSettingXml field if non-nil, zero value otherwise.

### GetEncodedSettingXmlOk

`func (o *IosMobileAppConfiguration) GetEncodedSettingXmlOk() (string, bool)`

GetEncodedSettingXmlOk returns a tuple with the EncodedSettingXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEncodedSettingXml

`func (o *IosMobileAppConfiguration) HasEncodedSettingXml() bool`

HasEncodedSettingXml returns a boolean if a field has been set.

### SetEncodedSettingXml

`func (o *IosMobileAppConfiguration) SetEncodedSettingXml(v string)`

SetEncodedSettingXml gets a reference to the given string and assigns it to the EncodedSettingXml field.

### SetEncodedSettingXmlExplicitNull

`func (o *IosMobileAppConfiguration) SetEncodedSettingXmlExplicitNull(b bool)`

SetEncodedSettingXmlExplicitNull (un)sets EncodedSettingXml to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EncodedSettingXml value is set to nil even if false is passed
### GetSettings

`func (o *IosMobileAppConfiguration) GetSettings() []AnyOfmicrosoftGraphAppConfigurationSettingItem`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *IosMobileAppConfiguration) GetSettingsOk() ([]AnyOfmicrosoftGraphAppConfigurationSettingItem, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettings

`func (o *IosMobileAppConfiguration) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettings

`func (o *IosMobileAppConfiguration) SetSettings(v []AnyOfmicrosoftGraphAppConfigurationSettingItem)`

SetSettings gets a reference to the given []AnyOfmicrosoftGraphAppConfigurationSettingItem and assigns it to the Settings field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


