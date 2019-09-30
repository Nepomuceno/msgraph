# MicrosoftGraphIosHomeScreenFolderPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Name of the folder page | [optional] 
**Apps** | Pointer to [**[]MicrosoftGraphIosHomeScreenApp**](microsoft.graph.iosHomeScreenApp.md) | A list of apps to appear on a page within a folder. This collection can contain a maximum of 500 elements. | [optional] 

## Methods

### GetDisplayName

`func (o *MicrosoftGraphIosHomeScreenFolderPage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIosHomeScreenFolderPage) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphIosHomeScreenFolderPage) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphIosHomeScreenFolderPage) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphIosHomeScreenFolderPage) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetApps

`func (o *MicrosoftGraphIosHomeScreenFolderPage) GetApps() []MicrosoftGraphIosHomeScreenApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *MicrosoftGraphIosHomeScreenFolderPage) GetAppsOk() ([]MicrosoftGraphIosHomeScreenApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApps

`func (o *MicrosoftGraphIosHomeScreenFolderPage) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetApps

`func (o *MicrosoftGraphIosHomeScreenFolderPage) SetApps(v []MicrosoftGraphIosHomeScreenApp)`

SetApps gets a reference to the given []MicrosoftGraphIosHomeScreenApp and assigns it to the Apps field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


