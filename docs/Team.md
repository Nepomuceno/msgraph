# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetWebUrl

`func (o *Team) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *Team) GetWebUrlOk() (string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebUrl

`func (o *Team) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrl

`func (o *Team) SetWebUrl(v string)`

SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.

### SetWebUrlExplicitNull

`func (o *Team) SetWebUrlExplicitNull(b bool)`

SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebUrl value is set to nil even if false is passed
### GetMemberSettings

`func (o *Team) GetMemberSettings() AnyOfmicrosoftGraphTeamMemberSettings`

GetMemberSettings returns the MemberSettings field if non-nil, zero value otherwise.

### GetMemberSettingsOk

`func (o *Team) GetMemberSettingsOk() (AnyOfmicrosoftGraphTeamMemberSettings, bool)`

GetMemberSettingsOk returns a tuple with the MemberSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMemberSettings

`func (o *Team) HasMemberSettings() bool`

HasMemberSettings returns a boolean if a field has been set.

### SetMemberSettings

`func (o *Team) SetMemberSettings(v AnyOfmicrosoftGraphTeamMemberSettings)`

SetMemberSettings gets a reference to the given AnyOfmicrosoftGraphTeamMemberSettings and assigns it to the MemberSettings field.

### SetMemberSettingsExplicitNull

`func (o *Team) SetMemberSettingsExplicitNull(b bool)`

SetMemberSettingsExplicitNull (un)sets MemberSettings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MemberSettings value is set to nil even if false is passed
### GetGuestSettings

`func (o *Team) GetGuestSettings() AnyOfmicrosoftGraphTeamGuestSettings`

GetGuestSettings returns the GuestSettings field if non-nil, zero value otherwise.

### GetGuestSettingsOk

`func (o *Team) GetGuestSettingsOk() (AnyOfmicrosoftGraphTeamGuestSettings, bool)`

GetGuestSettingsOk returns a tuple with the GuestSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGuestSettings

`func (o *Team) HasGuestSettings() bool`

HasGuestSettings returns a boolean if a field has been set.

### SetGuestSettings

`func (o *Team) SetGuestSettings(v AnyOfmicrosoftGraphTeamGuestSettings)`

SetGuestSettings gets a reference to the given AnyOfmicrosoftGraphTeamGuestSettings and assigns it to the GuestSettings field.

### SetGuestSettingsExplicitNull

`func (o *Team) SetGuestSettingsExplicitNull(b bool)`

SetGuestSettingsExplicitNull (un)sets GuestSettings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GuestSettings value is set to nil even if false is passed
### GetMessagingSettings

`func (o *Team) GetMessagingSettings() AnyOfmicrosoftGraphTeamMessagingSettings`

GetMessagingSettings returns the MessagingSettings field if non-nil, zero value otherwise.

### GetMessagingSettingsOk

`func (o *Team) GetMessagingSettingsOk() (AnyOfmicrosoftGraphTeamMessagingSettings, bool)`

GetMessagingSettingsOk returns a tuple with the MessagingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessagingSettings

`func (o *Team) HasMessagingSettings() bool`

HasMessagingSettings returns a boolean if a field has been set.

### SetMessagingSettings

`func (o *Team) SetMessagingSettings(v AnyOfmicrosoftGraphTeamMessagingSettings)`

SetMessagingSettings gets a reference to the given AnyOfmicrosoftGraphTeamMessagingSettings and assigns it to the MessagingSettings field.

### SetMessagingSettingsExplicitNull

`func (o *Team) SetMessagingSettingsExplicitNull(b bool)`

SetMessagingSettingsExplicitNull (un)sets MessagingSettings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MessagingSettings value is set to nil even if false is passed
### GetFunSettings

`func (o *Team) GetFunSettings() AnyOfmicrosoftGraphTeamFunSettings`

GetFunSettings returns the FunSettings field if non-nil, zero value otherwise.

### GetFunSettingsOk

`func (o *Team) GetFunSettingsOk() (AnyOfmicrosoftGraphTeamFunSettings, bool)`

GetFunSettingsOk returns a tuple with the FunSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFunSettings

`func (o *Team) HasFunSettings() bool`

HasFunSettings returns a boolean if a field has been set.

### SetFunSettings

`func (o *Team) SetFunSettings(v AnyOfmicrosoftGraphTeamFunSettings)`

SetFunSettings gets a reference to the given AnyOfmicrosoftGraphTeamFunSettings and assigns it to the FunSettings field.

### SetFunSettingsExplicitNull

`func (o *Team) SetFunSettingsExplicitNull(b bool)`

SetFunSettingsExplicitNull (un)sets FunSettings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FunSettings value is set to nil even if false is passed
### GetIsArchived

`func (o *Team) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Team) GetIsArchivedOk() (bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsArchived

`func (o *Team) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### SetIsArchived

`func (o *Team) SetIsArchived(v bool)`

SetIsArchived gets a reference to the given bool and assigns it to the IsArchived field.

### SetIsArchivedExplicitNull

`func (o *Team) SetIsArchivedExplicitNull(b bool)`

SetIsArchivedExplicitNull (un)sets IsArchived to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsArchived value is set to nil even if false is passed
### GetChannels

`func (o *Team) GetChannels() []MicrosoftGraphChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *Team) GetChannelsOk() ([]MicrosoftGraphChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChannels

`func (o *Team) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannels

`func (o *Team) SetChannels(v []MicrosoftGraphChannel)`

SetChannels gets a reference to the given []MicrosoftGraphChannel and assigns it to the Channels field.

### GetInstalledApps

`func (o *Team) GetInstalledApps() []MicrosoftGraphTeamsAppInstallation`

GetInstalledApps returns the InstalledApps field if non-nil, zero value otherwise.

### GetInstalledAppsOk

`func (o *Team) GetInstalledAppsOk() ([]MicrosoftGraphTeamsAppInstallation, bool)`

GetInstalledAppsOk returns a tuple with the InstalledApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstalledApps

`func (o *Team) HasInstalledApps() bool`

HasInstalledApps returns a boolean if a field has been set.

### SetInstalledApps

`func (o *Team) SetInstalledApps(v []MicrosoftGraphTeamsAppInstallation)`

SetInstalledApps gets a reference to the given []MicrosoftGraphTeamsAppInstallation and assigns it to the InstalledApps field.

### GetOperations

`func (o *Team) GetOperations() []MicrosoftGraphTeamsAsyncOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Team) GetOperationsOk() ([]MicrosoftGraphTeamsAsyncOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperations

`func (o *Team) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### SetOperations

`func (o *Team) SetOperations(v []MicrosoftGraphTeamsAsyncOperation)`

SetOperations gets a reference to the given []MicrosoftGraphTeamsAsyncOperation and assigns it to the Operations field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


