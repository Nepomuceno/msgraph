# MicrosoftGraphTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**WebUrl** | Pointer to **string** |  | [optional] 
**MemberSettings** | Pointer to [**AnyOfmicrosoftGraphTeamMemberSettings**](anyOf&lt;microsoft.graph.teamMemberSettings&gt;.md) |  | [optional] 
**GuestSettings** | Pointer to [**AnyOfmicrosoftGraphTeamGuestSettings**](anyOf&lt;microsoft.graph.teamGuestSettings&gt;.md) |  | [optional] 
**MessagingSettings** | Pointer to [**AnyOfmicrosoftGraphTeamMessagingSettings**](anyOf&lt;microsoft.graph.teamMessagingSettings&gt;.md) |  | [optional] 
**FunSettings** | Pointer to [**AnyOfmicrosoftGraphTeamFunSettings**](anyOf&lt;microsoft.graph.teamFunSettings&gt;.md) |  | [optional] 
**IsArchived** | Pointer to **bool** |  | [optional] 
**Channels** | Pointer to [**[]MicrosoftGraphChannel**](microsoft.graph.channel.md) |  | [optional] 
**InstalledApps** | Pointer to [**[]MicrosoftGraphTeamsAppInstallation**](microsoft.graph.teamsAppInstallation.md) |  | [optional] 
**Operations** | Pointer to [**[]MicrosoftGraphTeamsAsyncOperation**](microsoft.graph.teamsAsyncOperation.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphTeam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTeam) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphTeam) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetWebUrl

`func (o *MicrosoftGraphTeam) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphTeam) GetWebUrlOk() (string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebUrl

`func (o *MicrosoftGraphTeam) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrl

`func (o *MicrosoftGraphTeam) SetWebUrl(v string)`

SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.

### SetWebUrlExplicitNull

`func (o *MicrosoftGraphTeam) SetWebUrlExplicitNull(b bool)`

SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebUrl value is set to nil even if false is passed
### GetMemberSettings

`func (o *MicrosoftGraphTeam) GetMemberSettings() AnyOfmicrosoftGraphTeamMemberSettings`

GetMemberSettings returns the MemberSettings field if non-nil, zero value otherwise.

### GetMemberSettingsOk

`func (o *MicrosoftGraphTeam) GetMemberSettingsOk() (AnyOfmicrosoftGraphTeamMemberSettings, bool)`

GetMemberSettingsOk returns a tuple with the MemberSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMemberSettings

`func (o *MicrosoftGraphTeam) HasMemberSettings() bool`

HasMemberSettings returns a boolean if a field has been set.

### SetMemberSettings

`func (o *MicrosoftGraphTeam) SetMemberSettings(v AnyOfmicrosoftGraphTeamMemberSettings)`

SetMemberSettings gets a reference to the given AnyOfmicrosoftGraphTeamMemberSettings and assigns it to the MemberSettings field.

### SetMemberSettingsExplicitNull

`func (o *MicrosoftGraphTeam) SetMemberSettingsExplicitNull(b bool)`

SetMemberSettingsExplicitNull (un)sets MemberSettings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MemberSettings value is set to nil even if false is passed
### GetGuestSettings

`func (o *MicrosoftGraphTeam) GetGuestSettings() AnyOfmicrosoftGraphTeamGuestSettings`

GetGuestSettings returns the GuestSettings field if non-nil, zero value otherwise.

### GetGuestSettingsOk

`func (o *MicrosoftGraphTeam) GetGuestSettingsOk() (AnyOfmicrosoftGraphTeamGuestSettings, bool)`

GetGuestSettingsOk returns a tuple with the GuestSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGuestSettings

`func (o *MicrosoftGraphTeam) HasGuestSettings() bool`

HasGuestSettings returns a boolean if a field has been set.

### SetGuestSettings

`func (o *MicrosoftGraphTeam) SetGuestSettings(v AnyOfmicrosoftGraphTeamGuestSettings)`

SetGuestSettings gets a reference to the given AnyOfmicrosoftGraphTeamGuestSettings and assigns it to the GuestSettings field.

### SetGuestSettingsExplicitNull

`func (o *MicrosoftGraphTeam) SetGuestSettingsExplicitNull(b bool)`

SetGuestSettingsExplicitNull (un)sets GuestSettings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GuestSettings value is set to nil even if false is passed
### GetMessagingSettings

`func (o *MicrosoftGraphTeam) GetMessagingSettings() AnyOfmicrosoftGraphTeamMessagingSettings`

GetMessagingSettings returns the MessagingSettings field if non-nil, zero value otherwise.

### GetMessagingSettingsOk

`func (o *MicrosoftGraphTeam) GetMessagingSettingsOk() (AnyOfmicrosoftGraphTeamMessagingSettings, bool)`

GetMessagingSettingsOk returns a tuple with the MessagingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessagingSettings

`func (o *MicrosoftGraphTeam) HasMessagingSettings() bool`

HasMessagingSettings returns a boolean if a field has been set.

### SetMessagingSettings

`func (o *MicrosoftGraphTeam) SetMessagingSettings(v AnyOfmicrosoftGraphTeamMessagingSettings)`

SetMessagingSettings gets a reference to the given AnyOfmicrosoftGraphTeamMessagingSettings and assigns it to the MessagingSettings field.

### SetMessagingSettingsExplicitNull

`func (o *MicrosoftGraphTeam) SetMessagingSettingsExplicitNull(b bool)`

SetMessagingSettingsExplicitNull (un)sets MessagingSettings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MessagingSettings value is set to nil even if false is passed
### GetFunSettings

`func (o *MicrosoftGraphTeam) GetFunSettings() AnyOfmicrosoftGraphTeamFunSettings`

GetFunSettings returns the FunSettings field if non-nil, zero value otherwise.

### GetFunSettingsOk

`func (o *MicrosoftGraphTeam) GetFunSettingsOk() (AnyOfmicrosoftGraphTeamFunSettings, bool)`

GetFunSettingsOk returns a tuple with the FunSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFunSettings

`func (o *MicrosoftGraphTeam) HasFunSettings() bool`

HasFunSettings returns a boolean if a field has been set.

### SetFunSettings

`func (o *MicrosoftGraphTeam) SetFunSettings(v AnyOfmicrosoftGraphTeamFunSettings)`

SetFunSettings gets a reference to the given AnyOfmicrosoftGraphTeamFunSettings and assigns it to the FunSettings field.

### SetFunSettingsExplicitNull

`func (o *MicrosoftGraphTeam) SetFunSettingsExplicitNull(b bool)`

SetFunSettingsExplicitNull (un)sets FunSettings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FunSettings value is set to nil even if false is passed
### GetIsArchived

`func (o *MicrosoftGraphTeam) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *MicrosoftGraphTeam) GetIsArchivedOk() (bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsArchived

`func (o *MicrosoftGraphTeam) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### SetIsArchived

`func (o *MicrosoftGraphTeam) SetIsArchived(v bool)`

SetIsArchived gets a reference to the given bool and assigns it to the IsArchived field.

### SetIsArchivedExplicitNull

`func (o *MicrosoftGraphTeam) SetIsArchivedExplicitNull(b bool)`

SetIsArchivedExplicitNull (un)sets IsArchived to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsArchived value is set to nil even if false is passed
### GetChannels

`func (o *MicrosoftGraphTeam) GetChannels() []MicrosoftGraphChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *MicrosoftGraphTeam) GetChannelsOk() ([]MicrosoftGraphChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChannels

`func (o *MicrosoftGraphTeam) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannels

`func (o *MicrosoftGraphTeam) SetChannels(v []MicrosoftGraphChannel)`

SetChannels gets a reference to the given []MicrosoftGraphChannel and assigns it to the Channels field.

### GetInstalledApps

`func (o *MicrosoftGraphTeam) GetInstalledApps() []MicrosoftGraphTeamsAppInstallation`

GetInstalledApps returns the InstalledApps field if non-nil, zero value otherwise.

### GetInstalledAppsOk

`func (o *MicrosoftGraphTeam) GetInstalledAppsOk() ([]MicrosoftGraphTeamsAppInstallation, bool)`

GetInstalledAppsOk returns a tuple with the InstalledApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstalledApps

`func (o *MicrosoftGraphTeam) HasInstalledApps() bool`

HasInstalledApps returns a boolean if a field has been set.

### SetInstalledApps

`func (o *MicrosoftGraphTeam) SetInstalledApps(v []MicrosoftGraphTeamsAppInstallation)`

SetInstalledApps gets a reference to the given []MicrosoftGraphTeamsAppInstallation and assigns it to the InstalledApps field.

### GetOperations

`func (o *MicrosoftGraphTeam) GetOperations() []MicrosoftGraphTeamsAsyncOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *MicrosoftGraphTeam) GetOperationsOk() ([]MicrosoftGraphTeamsAsyncOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperations

`func (o *MicrosoftGraphTeam) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### SetOperations

`func (o *MicrosoftGraphTeam) SetOperations(v []MicrosoftGraphTeamsAsyncOperation)`

SetOperations gets a reference to the given []MicrosoftGraphTeamsAsyncOperation and assigns it to the Operations field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


