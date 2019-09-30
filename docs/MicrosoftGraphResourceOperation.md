# MicrosoftGraphResourceOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** | Name of the Resource this operation is performed on. | [optional] 
**ActionName** | Pointer to **string** | Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible. | [optional] 
**Description** | Pointer to **string** | Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphResourceOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphResourceOperation) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphResourceOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphResourceOperation) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetResourceName

`func (o *MicrosoftGraphResourceOperation) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *MicrosoftGraphResourceOperation) GetResourceNameOk() (string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceName

`func (o *MicrosoftGraphResourceOperation) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceName

`func (o *MicrosoftGraphResourceOperation) SetResourceName(v string)`

SetResourceName gets a reference to the given string and assigns it to the ResourceName field.

### SetResourceNameExplicitNull

`func (o *MicrosoftGraphResourceOperation) SetResourceNameExplicitNull(b bool)`

SetResourceNameExplicitNull (un)sets ResourceName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceName value is set to nil even if false is passed
### GetActionName

`func (o *MicrosoftGraphResourceOperation) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *MicrosoftGraphResourceOperation) GetActionNameOk() (string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionName

`func (o *MicrosoftGraphResourceOperation) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### SetActionName

`func (o *MicrosoftGraphResourceOperation) SetActionName(v string)`

SetActionName gets a reference to the given string and assigns it to the ActionName field.

### SetActionNameExplicitNull

`func (o *MicrosoftGraphResourceOperation) SetActionNameExplicitNull(b bool)`

SetActionNameExplicitNull (un)sets ActionName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActionName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphResourceOperation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphResourceOperation) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphResourceOperation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphResourceOperation) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphResourceOperation) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


