# ResourceOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceName** | Pointer to **string** | Name of the Resource this operation is performed on. | [optional] 
**ActionName** | Pointer to **string** | Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible. | [optional] 
**Description** | Pointer to **string** | Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal. | [optional] 

## Methods

### GetResourceName

`func (o *ResourceOperation) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceOperation) GetResourceNameOk() (string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceName

`func (o *ResourceOperation) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceName

`func (o *ResourceOperation) SetResourceName(v string)`

SetResourceName gets a reference to the given string and assigns it to the ResourceName field.

### SetResourceNameExplicitNull

`func (o *ResourceOperation) SetResourceNameExplicitNull(b bool)`

SetResourceNameExplicitNull (un)sets ResourceName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceName value is set to nil even if false is passed
### GetActionName

`func (o *ResourceOperation) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *ResourceOperation) GetActionNameOk() (string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionName

`func (o *ResourceOperation) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### SetActionName

`func (o *ResourceOperation) SetActionName(v string)`

SetActionName gets a reference to the given string and assigns it to the ActionName field.

### SetActionNameExplicitNull

`func (o *ResourceOperation) SetActionNameExplicitNull(b bool)`

SetActionNameExplicitNull (un)sets ActionName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActionName value is set to nil even if false is passed
### GetDescription

`func (o *ResourceOperation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceOperation) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *ResourceOperation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *ResourceOperation) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *ResourceOperation) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


