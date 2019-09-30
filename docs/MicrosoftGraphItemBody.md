# MicrosoftGraphItemBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to [**AnyOfmicrosoftGraphBodyType**](anyOf&lt;microsoft.graph.bodyType&gt;.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### GetContentType

`func (o *MicrosoftGraphItemBody) GetContentType() AnyOfmicrosoftGraphBodyType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MicrosoftGraphItemBody) GetContentTypeOk() (AnyOfmicrosoftGraphBodyType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentType

`func (o *MicrosoftGraphItemBody) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentType

`func (o *MicrosoftGraphItemBody) SetContentType(v AnyOfmicrosoftGraphBodyType)`

SetContentType gets a reference to the given AnyOfmicrosoftGraphBodyType and assigns it to the ContentType field.

### SetContentTypeExplicitNull

`func (o *MicrosoftGraphItemBody) SetContentTypeExplicitNull(b bool)`

SetContentTypeExplicitNull (un)sets ContentType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentType value is set to nil even if false is passed
### GetContent

`func (o *MicrosoftGraphItemBody) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphItemBody) GetContentOk() (string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContent

`func (o *MicrosoftGraphItemBody) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContent

`func (o *MicrosoftGraphItemBody) SetContent(v string)`

SetContent gets a reference to the given string and assigns it to the Content field.

### SetContentExplicitNull

`func (o *MicrosoftGraphItemBody) SetContentExplicitNull(b bool)`

SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Content value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


