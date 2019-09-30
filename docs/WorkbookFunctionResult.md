# WorkbookFunctionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 

## Methods

### GetError

`func (o *WorkbookFunctionResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WorkbookFunctionResult) GetErrorOk() (string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *WorkbookFunctionResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *WorkbookFunctionResult) SetError(v string)`

SetError gets a reference to the given string and assigns it to the Error field.

### SetErrorExplicitNull

`func (o *WorkbookFunctionResult) SetErrorExplicitNull(b bool)`

SetErrorExplicitNull (un)sets Error to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Error value is set to nil even if false is passed
### GetValue

`func (o *WorkbookFunctionResult) GetValue() AnyOfobject`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkbookFunctionResult) GetValueOk() (AnyOfobject, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *WorkbookFunctionResult) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *WorkbookFunctionResult) SetValue(v AnyOfobject)`

SetValue gets a reference to the given AnyOfobject and assigns it to the Value field.

### SetValueExplicitNull

`func (o *WorkbookFunctionResult) SetValueExplicitNull(b bool)`

SetValueExplicitNull (un)sets Value to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Value value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


