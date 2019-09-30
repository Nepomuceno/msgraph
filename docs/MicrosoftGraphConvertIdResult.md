# MicrosoftGraphConvertIdResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**ErrorDetails** | Pointer to [**AnyOfmicrosoftGraphGenericError**](anyOf&lt;microsoft.graph.genericError&gt;.md) |  | [optional] 

## Methods

### GetSourceId

`func (o *MicrosoftGraphConvertIdResult) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *MicrosoftGraphConvertIdResult) GetSourceIdOk() (string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceId

`func (o *MicrosoftGraphConvertIdResult) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceId

`func (o *MicrosoftGraphConvertIdResult) SetSourceId(v string)`

SetSourceId gets a reference to the given string and assigns it to the SourceId field.

### SetSourceIdExplicitNull

`func (o *MicrosoftGraphConvertIdResult) SetSourceIdExplicitNull(b bool)`

SetSourceIdExplicitNull (un)sets SourceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SourceId value is set to nil even if false is passed
### GetTargetId

`func (o *MicrosoftGraphConvertIdResult) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *MicrosoftGraphConvertIdResult) GetTargetIdOk() (string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetId

`func (o *MicrosoftGraphConvertIdResult) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetId

`func (o *MicrosoftGraphConvertIdResult) SetTargetId(v string)`

SetTargetId gets a reference to the given string and assigns it to the TargetId field.

### SetTargetIdExplicitNull

`func (o *MicrosoftGraphConvertIdResult) SetTargetIdExplicitNull(b bool)`

SetTargetIdExplicitNull (un)sets TargetId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TargetId value is set to nil even if false is passed
### GetErrorDetails

`func (o *MicrosoftGraphConvertIdResult) GetErrorDetails() AnyOfmicrosoftGraphGenericError`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *MicrosoftGraphConvertIdResult) GetErrorDetailsOk() (AnyOfmicrosoftGraphGenericError, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorDetails

`func (o *MicrosoftGraphConvertIdResult) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### SetErrorDetails

`func (o *MicrosoftGraphConvertIdResult) SetErrorDetails(v AnyOfmicrosoftGraphGenericError)`

SetErrorDetails gets a reference to the given AnyOfmicrosoftGraphGenericError and assigns it to the ErrorDetails field.

### SetErrorDetailsExplicitNull

`func (o *MicrosoftGraphConvertIdResult) SetErrorDetailsExplicitNull(b bool)`

SetErrorDetailsExplicitNull (un)sets ErrorDetails to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ErrorDetails value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


