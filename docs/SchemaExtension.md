# SchemaExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**TargetTypes** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to [**[]MicrosoftGraphExtensionSchemaProperty**](microsoft.graph.extensionSchemaProperty.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 

## Methods

### GetDescription

`func (o *SchemaExtension) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaExtension) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *SchemaExtension) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *SchemaExtension) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *SchemaExtension) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetTargetTypes

`func (o *SchemaExtension) GetTargetTypes() []string`

GetTargetTypes returns the TargetTypes field if non-nil, zero value otherwise.

### GetTargetTypesOk

`func (o *SchemaExtension) GetTargetTypesOk() ([]string, bool)`

GetTargetTypesOk returns a tuple with the TargetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetTypes

`func (o *SchemaExtension) HasTargetTypes() bool`

HasTargetTypes returns a boolean if a field has been set.

### SetTargetTypes

`func (o *SchemaExtension) SetTargetTypes(v []string)`

SetTargetTypes gets a reference to the given []string and assigns it to the TargetTypes field.

### GetProperties

`func (o *SchemaExtension) GetProperties() []MicrosoftGraphExtensionSchemaProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SchemaExtension) GetPropertiesOk() ([]MicrosoftGraphExtensionSchemaProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProperties

`func (o *SchemaExtension) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetProperties

`func (o *SchemaExtension) SetProperties(v []MicrosoftGraphExtensionSchemaProperty)`

SetProperties gets a reference to the given []MicrosoftGraphExtensionSchemaProperty and assigns it to the Properties field.

### GetStatus

`func (o *SchemaExtension) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SchemaExtension) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *SchemaExtension) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *SchemaExtension) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetOwner

`func (o *SchemaExtension) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SchemaExtension) GetOwnerOk() (string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *SchemaExtension) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *SchemaExtension) SetOwner(v string)`

SetOwner gets a reference to the given string and assigns it to the Owner field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


