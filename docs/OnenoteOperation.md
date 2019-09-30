# OnenoteOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceLocation** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**AnyOfmicrosoftGraphOnenoteOperationError**](anyOf&lt;microsoft.graph.onenoteOperationError&gt;.md) |  | [optional] 
**PercentComplete** | Pointer to **string** |  | [optional] 

## Methods

### GetResourceLocation

`func (o *OnenoteOperation) GetResourceLocation() string`

GetResourceLocation returns the ResourceLocation field if non-nil, zero value otherwise.

### GetResourceLocationOk

`func (o *OnenoteOperation) GetResourceLocationOk() (string, bool)`

GetResourceLocationOk returns a tuple with the ResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceLocation

`func (o *OnenoteOperation) HasResourceLocation() bool`

HasResourceLocation returns a boolean if a field has been set.

### SetResourceLocation

`func (o *OnenoteOperation) SetResourceLocation(v string)`

SetResourceLocation gets a reference to the given string and assigns it to the ResourceLocation field.

### SetResourceLocationExplicitNull

`func (o *OnenoteOperation) SetResourceLocationExplicitNull(b bool)`

SetResourceLocationExplicitNull (un)sets ResourceLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceLocation value is set to nil even if false is passed
### GetResourceId

`func (o *OnenoteOperation) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *OnenoteOperation) GetResourceIdOk() (string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceId

`func (o *OnenoteOperation) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceId

`func (o *OnenoteOperation) SetResourceId(v string)`

SetResourceId gets a reference to the given string and assigns it to the ResourceId field.

### SetResourceIdExplicitNull

`func (o *OnenoteOperation) SetResourceIdExplicitNull(b bool)`

SetResourceIdExplicitNull (un)sets ResourceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceId value is set to nil even if false is passed
### GetError

`func (o *OnenoteOperation) GetError() AnyOfmicrosoftGraphOnenoteOperationError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OnenoteOperation) GetErrorOk() (AnyOfmicrosoftGraphOnenoteOperationError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *OnenoteOperation) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *OnenoteOperation) SetError(v AnyOfmicrosoftGraphOnenoteOperationError)`

SetError gets a reference to the given AnyOfmicrosoftGraphOnenoteOperationError and assigns it to the Error field.

### SetErrorExplicitNull

`func (o *OnenoteOperation) SetErrorExplicitNull(b bool)`

SetErrorExplicitNull (un)sets Error to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Error value is set to nil even if false is passed
### GetPercentComplete

`func (o *OnenoteOperation) GetPercentComplete() string`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *OnenoteOperation) GetPercentCompleteOk() (string, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPercentComplete

`func (o *OnenoteOperation) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentComplete

`func (o *OnenoteOperation) SetPercentComplete(v string)`

SetPercentComplete gets a reference to the given string and assigns it to the PercentComplete field.

### SetPercentCompleteExplicitNull

`func (o *OnenoteOperation) SetPercentCompleteExplicitNull(b bool)`

SetPercentCompleteExplicitNull (un)sets PercentComplete to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PercentComplete value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


