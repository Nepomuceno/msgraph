# InferenceClassificationOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassifyAs** | Pointer to [**AnyOfmicrosoftGraphInferenceClassificationType**](anyOf&lt;microsoft.graph.inferenceClassificationType&gt;.md) |  | [optional] 
**SenderEmailAddress** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) |  | [optional] 

## Methods

### GetClassifyAs

`func (o *InferenceClassificationOverride) GetClassifyAs() AnyOfmicrosoftGraphInferenceClassificationType`

GetClassifyAs returns the ClassifyAs field if non-nil, zero value otherwise.

### GetClassifyAsOk

`func (o *InferenceClassificationOverride) GetClassifyAsOk() (AnyOfmicrosoftGraphInferenceClassificationType, bool)`

GetClassifyAsOk returns a tuple with the ClassifyAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClassifyAs

`func (o *InferenceClassificationOverride) HasClassifyAs() bool`

HasClassifyAs returns a boolean if a field has been set.

### SetClassifyAs

`func (o *InferenceClassificationOverride) SetClassifyAs(v AnyOfmicrosoftGraphInferenceClassificationType)`

SetClassifyAs gets a reference to the given AnyOfmicrosoftGraphInferenceClassificationType and assigns it to the ClassifyAs field.

### SetClassifyAsExplicitNull

`func (o *InferenceClassificationOverride) SetClassifyAsExplicitNull(b bool)`

SetClassifyAsExplicitNull (un)sets ClassifyAs to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ClassifyAs value is set to nil even if false is passed
### GetSenderEmailAddress

`func (o *InferenceClassificationOverride) GetSenderEmailAddress() AnyOfmicrosoftGraphEmailAddress`

GetSenderEmailAddress returns the SenderEmailAddress field if non-nil, zero value otherwise.

### GetSenderEmailAddressOk

`func (o *InferenceClassificationOverride) GetSenderEmailAddressOk() (AnyOfmicrosoftGraphEmailAddress, bool)`

GetSenderEmailAddressOk returns a tuple with the SenderEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSenderEmailAddress

`func (o *InferenceClassificationOverride) HasSenderEmailAddress() bool`

HasSenderEmailAddress returns a boolean if a field has been set.

### SetSenderEmailAddress

`func (o *InferenceClassificationOverride) SetSenderEmailAddress(v AnyOfmicrosoftGraphEmailAddress)`

SetSenderEmailAddress gets a reference to the given AnyOfmicrosoftGraphEmailAddress and assigns it to the SenderEmailAddress field.

### SetSenderEmailAddressExplicitNull

`func (o *InferenceClassificationOverride) SetSenderEmailAddressExplicitNull(b bool)`

SetSenderEmailAddressExplicitNull (un)sets SenderEmailAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SenderEmailAddress value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


