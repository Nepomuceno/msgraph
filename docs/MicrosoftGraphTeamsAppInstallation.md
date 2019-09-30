# MicrosoftGraphTeamsAppInstallation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TeamsApp** | Pointer to [**AnyOfmicrosoftGraphTeamsApp**](anyOf&lt;microsoft.graph.teamsApp&gt;.md) |  | [optional] 
**TeamsAppDefinition** | Pointer to [**AnyOfmicrosoftGraphTeamsAppDefinition**](anyOf&lt;microsoft.graph.teamsAppDefinition&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphTeamsAppInstallation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTeamsAppInstallation) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphTeamsAppInstallation) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphTeamsAppInstallation) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetTeamsApp

`func (o *MicrosoftGraphTeamsAppInstallation) GetTeamsApp() AnyOfmicrosoftGraphTeamsApp`

GetTeamsApp returns the TeamsApp field if non-nil, zero value otherwise.

### GetTeamsAppOk

`func (o *MicrosoftGraphTeamsAppInstallation) GetTeamsAppOk() (AnyOfmicrosoftGraphTeamsApp, bool)`

GetTeamsAppOk returns a tuple with the TeamsApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeamsApp

`func (o *MicrosoftGraphTeamsAppInstallation) HasTeamsApp() bool`

HasTeamsApp returns a boolean if a field has been set.

### SetTeamsApp

`func (o *MicrosoftGraphTeamsAppInstallation) SetTeamsApp(v AnyOfmicrosoftGraphTeamsApp)`

SetTeamsApp gets a reference to the given AnyOfmicrosoftGraphTeamsApp and assigns it to the TeamsApp field.

### SetTeamsAppExplicitNull

`func (o *MicrosoftGraphTeamsAppInstallation) SetTeamsAppExplicitNull(b bool)`

SetTeamsAppExplicitNull (un)sets TeamsApp to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TeamsApp value is set to nil even if false is passed
### GetTeamsAppDefinition

`func (o *MicrosoftGraphTeamsAppInstallation) GetTeamsAppDefinition() AnyOfmicrosoftGraphTeamsAppDefinition`

GetTeamsAppDefinition returns the TeamsAppDefinition field if non-nil, zero value otherwise.

### GetTeamsAppDefinitionOk

`func (o *MicrosoftGraphTeamsAppInstallation) GetTeamsAppDefinitionOk() (AnyOfmicrosoftGraphTeamsAppDefinition, bool)`

GetTeamsAppDefinitionOk returns a tuple with the TeamsAppDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeamsAppDefinition

`func (o *MicrosoftGraphTeamsAppInstallation) HasTeamsAppDefinition() bool`

HasTeamsAppDefinition returns a boolean if a field has been set.

### SetTeamsAppDefinition

`func (o *MicrosoftGraphTeamsAppInstallation) SetTeamsAppDefinition(v AnyOfmicrosoftGraphTeamsAppDefinition)`

SetTeamsAppDefinition gets a reference to the given AnyOfmicrosoftGraphTeamsAppDefinition and assigns it to the TeamsAppDefinition field.

### SetTeamsAppDefinitionExplicitNull

`func (o *MicrosoftGraphTeamsAppInstallation) SetTeamsAppDefinitionExplicitNull(b bool)`

SetTeamsAppDefinitionExplicitNull (un)sets TeamsAppDefinition to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TeamsAppDefinition value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


