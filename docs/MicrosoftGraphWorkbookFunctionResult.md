# MicrosoftGraphWorkbookFunctionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWorkbookFunctionResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookFunctionResult) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookFunctionResult) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookFunctionResult) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetError

`func (o *MicrosoftGraphWorkbookFunctionResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphWorkbookFunctionResult) GetErrorOk() (string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *MicrosoftGraphWorkbookFunctionResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *MicrosoftGraphWorkbookFunctionResult) SetError(v string)`

SetError gets a reference to the given string and assigns it to the Error field.

### SetErrorExplicitNull

`func (o *MicrosoftGraphWorkbookFunctionResult) SetErrorExplicitNull(b bool)`

SetErrorExplicitNull (un)sets Error to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Error value is set to nil even if false is passed
### GetValue

`func (o *MicrosoftGraphWorkbookFunctionResult) GetValue() AnyOfobject`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphWorkbookFunctionResult) GetValueOk() (AnyOfobject, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *MicrosoftGraphWorkbookFunctionResult) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *MicrosoftGraphWorkbookFunctionResult) SetValue(v AnyOfobject)`

SetValue gets a reference to the given AnyOfobject and assigns it to the Value field.

### SetValueExplicitNull

`func (o *MicrosoftGraphWorkbookFunctionResult) SetValueExplicitNull(b bool)`

SetValueExplicitNull (un)sets Value to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Value value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


