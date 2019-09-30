# MicrosoftGraphTeamsApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DistributionMethod** | Pointer to [**AnyOfmicrosoftGraphTeamsAppDistributionMethod**](anyOf&lt;microsoft.graph.teamsAppDistributionMethod&gt;.md) |  | [optional] 
**AppDefinitions** | Pointer to [**[]MicrosoftGraphTeamsAppDefinition**](microsoft.graph.teamsAppDefinition.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphTeamsApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTeamsApp) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphTeamsApp) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphTeamsApp) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetExternalId

`func (o *MicrosoftGraphTeamsApp) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *MicrosoftGraphTeamsApp) GetExternalIdOk() (string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalId

`func (o *MicrosoftGraphTeamsApp) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalId

`func (o *MicrosoftGraphTeamsApp) SetExternalId(v string)`

SetExternalId gets a reference to the given string and assigns it to the ExternalId field.

### SetExternalIdExplicitNull

`func (o *MicrosoftGraphTeamsApp) SetExternalIdExplicitNull(b bool)`

SetExternalIdExplicitNull (un)sets ExternalId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalId value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphTeamsApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTeamsApp) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphTeamsApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphTeamsApp) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphTeamsApp) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDistributionMethod

`func (o *MicrosoftGraphTeamsApp) GetDistributionMethod() AnyOfmicrosoftGraphTeamsAppDistributionMethod`

GetDistributionMethod returns the DistributionMethod field if non-nil, zero value otherwise.

### GetDistributionMethodOk

`func (o *MicrosoftGraphTeamsApp) GetDistributionMethodOk() (AnyOfmicrosoftGraphTeamsAppDistributionMethod, bool)`

GetDistributionMethodOk returns a tuple with the DistributionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDistributionMethod

`func (o *MicrosoftGraphTeamsApp) HasDistributionMethod() bool`

HasDistributionMethod returns a boolean if a field has been set.

### SetDistributionMethod

`func (o *MicrosoftGraphTeamsApp) SetDistributionMethod(v AnyOfmicrosoftGraphTeamsAppDistributionMethod)`

SetDistributionMethod gets a reference to the given AnyOfmicrosoftGraphTeamsAppDistributionMethod and assigns it to the DistributionMethod field.

### GetAppDefinitions

`func (o *MicrosoftGraphTeamsApp) GetAppDefinitions() []MicrosoftGraphTeamsAppDefinition`

GetAppDefinitions returns the AppDefinitions field if non-nil, zero value otherwise.

### GetAppDefinitionsOk

`func (o *MicrosoftGraphTeamsApp) GetAppDefinitionsOk() ([]MicrosoftGraphTeamsAppDefinition, bool)`

GetAppDefinitionsOk returns a tuple with the AppDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppDefinitions

`func (o *MicrosoftGraphTeamsApp) HasAppDefinitions() bool`

HasAppDefinitions returns a boolean if a field has been set.

### SetAppDefinitions

`func (o *MicrosoftGraphTeamsApp) SetAppDefinitions(v []MicrosoftGraphTeamsAppDefinition)`

SetAppDefinitions gets a reference to the given []MicrosoftGraphTeamsAppDefinition and assigns it to the AppDefinitions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


