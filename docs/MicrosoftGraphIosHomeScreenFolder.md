# MicrosoftGraphIosHomeScreenFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Name of the app | [optional] 
**Pages** | Pointer to [**[]MicrosoftGraphIosHomeScreenFolderPage**](microsoft.graph.iosHomeScreenFolderPage.md) | Pages of Home Screen Layout Icons which must be Application Type. This collection can contain a maximum of 500 elements. | [optional] 

## Methods

### GetDisplayName

`func (o *MicrosoftGraphIosHomeScreenFolder) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIosHomeScreenFolder) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphIosHomeScreenFolder) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphIosHomeScreenFolder) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphIosHomeScreenFolder) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetPages

`func (o *MicrosoftGraphIosHomeScreenFolder) GetPages() []MicrosoftGraphIosHomeScreenFolderPage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *MicrosoftGraphIosHomeScreenFolder) GetPagesOk() ([]MicrosoftGraphIosHomeScreenFolderPage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPages

`func (o *MicrosoftGraphIosHomeScreenFolder) HasPages() bool`

HasPages returns a boolean if a field has been set.

### SetPages

`func (o *MicrosoftGraphIosHomeScreenFolder) SetPages(v []MicrosoftGraphIosHomeScreenFolderPage)`

SetPages gets a reference to the given []MicrosoftGraphIosHomeScreenFolderPage and assigns it to the Pages field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


