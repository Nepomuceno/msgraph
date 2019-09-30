# MicrosoftGraphOutlookCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Color** | Pointer to [**AnyOfmicrosoftGraphCategoryColor**](anyOf&lt;microsoft.graph.categoryColor&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphOutlookCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOutlookCategory) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphOutlookCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphOutlookCategory) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphOutlookCategory) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphOutlookCategory) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphOutlookCategory) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphOutlookCategory) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphOutlookCategory) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetColor

`func (o *MicrosoftGraphOutlookCategory) GetColor() AnyOfmicrosoftGraphCategoryColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *MicrosoftGraphOutlookCategory) GetColorOk() (AnyOfmicrosoftGraphCategoryColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColor

`func (o *MicrosoftGraphOutlookCategory) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColor

`func (o *MicrosoftGraphOutlookCategory) SetColor(v AnyOfmicrosoftGraphCategoryColor)`

SetColor gets a reference to the given AnyOfmicrosoftGraphCategoryColor and assigns it to the Color field.

### SetColorExplicitNull

`func (o *MicrosoftGraphOutlookCategory) SetColorExplicitNull(b bool)`

SetColorExplicitNull (un)sets Color to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Color value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


