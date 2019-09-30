# MicrosoftGraphIosHomeScreenPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Name of the page | [optional] 
**Icons** | Pointer to [**[]MicrosoftGraphIosHomeScreenItem**](microsoft.graph.iosHomeScreenItem.md) | A list of apps and folders to appear on a page. This collection can contain a maximum of 500 elements. | [optional] 

## Methods

### GetDisplayName

`func (o *MicrosoftGraphIosHomeScreenPage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIosHomeScreenPage) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphIosHomeScreenPage) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphIosHomeScreenPage) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphIosHomeScreenPage) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetIcons

`func (o *MicrosoftGraphIosHomeScreenPage) GetIcons() []MicrosoftGraphIosHomeScreenItem`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *MicrosoftGraphIosHomeScreenPage) GetIconsOk() ([]MicrosoftGraphIosHomeScreenItem, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIcons

`func (o *MicrosoftGraphIosHomeScreenPage) HasIcons() bool`

HasIcons returns a boolean if a field has been set.

### SetIcons

`func (o *MicrosoftGraphIosHomeScreenPage) SetIcons(v []MicrosoftGraphIosHomeScreenItem)`

SetIcons gets a reference to the given []MicrosoftGraphIosHomeScreenItem and assigns it to the Icons field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


