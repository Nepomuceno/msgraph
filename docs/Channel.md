# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**WebUrl** | Pointer to **string** |  | [optional] 
**Tabs** | Pointer to [**[]MicrosoftGraphTeamsTab**](microsoft.graph.teamsTab.md) |  | [optional] 

## Methods

### GetDisplayName

`func (o *Channel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Channel) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *Channel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *Channel) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *Channel) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *Channel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Channel) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Channel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Channel) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *Channel) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetEmail

`func (o *Channel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Channel) GetEmailOk() (string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmail

`func (o *Channel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmail

`func (o *Channel) SetEmail(v string)`

SetEmail gets a reference to the given string and assigns it to the Email field.

### SetEmailExplicitNull

`func (o *Channel) SetEmailExplicitNull(b bool)`

SetEmailExplicitNull (un)sets Email to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Email value is set to nil even if false is passed
### GetWebUrl

`func (o *Channel) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *Channel) GetWebUrlOk() (string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebUrl

`func (o *Channel) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrl

`func (o *Channel) SetWebUrl(v string)`

SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.

### SetWebUrlExplicitNull

`func (o *Channel) SetWebUrlExplicitNull(b bool)`

SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebUrl value is set to nil even if false is passed
### GetTabs

`func (o *Channel) GetTabs() []MicrosoftGraphTeamsTab`

GetTabs returns the Tabs field if non-nil, zero value otherwise.

### GetTabsOk

`func (o *Channel) GetTabsOk() ([]MicrosoftGraphTeamsTab, bool)`

GetTabsOk returns a tuple with the Tabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTabs

`func (o *Channel) HasTabs() bool`

HasTabs returns a boolean if a field has been set.

### SetTabs

`func (o *Channel) SetTabs(v []MicrosoftGraphTeamsTab)`

SetTabs gets a reference to the given []MicrosoftGraphTeamsTab and assigns it to the Tabs field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


