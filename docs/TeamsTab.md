# TeamsTab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**WebUrl** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to [**AnyOfmicrosoftGraphTeamsTabConfiguration**](anyOf&lt;microsoft.graph.teamsTabConfiguration&gt;.md) |  | [optional] 
**TeamsApp** | Pointer to [**AnyOfmicrosoftGraphTeamsApp**](anyOf&lt;microsoft.graph.teamsApp&gt;.md) |  | [optional] 

## Methods

### GetDisplayName

`func (o *TeamsTab) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TeamsTab) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *TeamsTab) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *TeamsTab) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *TeamsTab) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetWebUrl

`func (o *TeamsTab) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *TeamsTab) GetWebUrlOk() (string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebUrl

`func (o *TeamsTab) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrl

`func (o *TeamsTab) SetWebUrl(v string)`

SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.

### SetWebUrlExplicitNull

`func (o *TeamsTab) SetWebUrlExplicitNull(b bool)`

SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebUrl value is set to nil even if false is passed
### GetConfiguration

`func (o *TeamsTab) GetConfiguration() AnyOfmicrosoftGraphTeamsTabConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *TeamsTab) GetConfigurationOk() (AnyOfmicrosoftGraphTeamsTabConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfiguration

`func (o *TeamsTab) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfiguration

`func (o *TeamsTab) SetConfiguration(v AnyOfmicrosoftGraphTeamsTabConfiguration)`

SetConfiguration gets a reference to the given AnyOfmicrosoftGraphTeamsTabConfiguration and assigns it to the Configuration field.

### SetConfigurationExplicitNull

`func (o *TeamsTab) SetConfigurationExplicitNull(b bool)`

SetConfigurationExplicitNull (un)sets Configuration to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Configuration value is set to nil even if false is passed
### GetTeamsApp

`func (o *TeamsTab) GetTeamsApp() AnyOfmicrosoftGraphTeamsApp`

GetTeamsApp returns the TeamsApp field if non-nil, zero value otherwise.

### GetTeamsAppOk

`func (o *TeamsTab) GetTeamsAppOk() (AnyOfmicrosoftGraphTeamsApp, bool)`

GetTeamsAppOk returns a tuple with the TeamsApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTeamsApp

`func (o *TeamsTab) HasTeamsApp() bool`

HasTeamsApp returns a boolean if a field has been set.

### SetTeamsApp

`func (o *TeamsTab) SetTeamsApp(v AnyOfmicrosoftGraphTeamsApp)`

SetTeamsApp gets a reference to the given AnyOfmicrosoftGraphTeamsApp and assigns it to the TeamsApp field.

### SetTeamsAppExplicitNull

`func (o *TeamsTab) SetTeamsAppExplicitNull(b bool)`

SetTeamsAppExplicitNull (un)sets TeamsApp to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TeamsApp value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


