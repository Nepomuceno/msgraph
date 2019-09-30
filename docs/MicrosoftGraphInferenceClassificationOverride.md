# MicrosoftGraphInferenceClassificationOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ClassifyAs** | Pointer to [**AnyOfmicrosoftGraphInferenceClassificationType**](anyOf&lt;microsoft.graph.inferenceClassificationType&gt;.md) |  | [optional] 
**SenderEmailAddress** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphInferenceClassificationOverride) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphInferenceClassificationOverride) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphInferenceClassificationOverride) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphInferenceClassificationOverride) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetClassifyAs

`func (o *MicrosoftGraphInferenceClassificationOverride) GetClassifyAs() AnyOfmicrosoftGraphInferenceClassificationType`

GetClassifyAs returns the ClassifyAs field if non-nil, zero value otherwise.

### GetClassifyAsOk

`func (o *MicrosoftGraphInferenceClassificationOverride) GetClassifyAsOk() (AnyOfmicrosoftGraphInferenceClassificationType, bool)`

GetClassifyAsOk returns a tuple with the ClassifyAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClassifyAs

`func (o *MicrosoftGraphInferenceClassificationOverride) HasClassifyAs() bool`

HasClassifyAs returns a boolean if a field has been set.

### SetClassifyAs

`func (o *MicrosoftGraphInferenceClassificationOverride) SetClassifyAs(v AnyOfmicrosoftGraphInferenceClassificationType)`

SetClassifyAs gets a reference to the given AnyOfmicrosoftGraphInferenceClassificationType and assigns it to the ClassifyAs field.

### SetClassifyAsExplicitNull

`func (o *MicrosoftGraphInferenceClassificationOverride) SetClassifyAsExplicitNull(b bool)`

SetClassifyAsExplicitNull (un)sets ClassifyAs to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ClassifyAs value is set to nil even if false is passed
### GetSenderEmailAddress

`func (o *MicrosoftGraphInferenceClassificationOverride) GetSenderEmailAddress() AnyOfmicrosoftGraphEmailAddress`

GetSenderEmailAddress returns the SenderEmailAddress field if non-nil, zero value otherwise.

### GetSenderEmailAddressOk

`func (o *MicrosoftGraphInferenceClassificationOverride) GetSenderEmailAddressOk() (AnyOfmicrosoftGraphEmailAddress, bool)`

GetSenderEmailAddressOk returns a tuple with the SenderEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSenderEmailAddress

`func (o *MicrosoftGraphInferenceClassificationOverride) HasSenderEmailAddress() bool`

HasSenderEmailAddress returns a boolean if a field has been set.

### SetSenderEmailAddress

`func (o *MicrosoftGraphInferenceClassificationOverride) SetSenderEmailAddress(v AnyOfmicrosoftGraphEmailAddress)`

SetSenderEmailAddress gets a reference to the given AnyOfmicrosoftGraphEmailAddress and assigns it to the SenderEmailAddress field.

### SetSenderEmailAddressExplicitNull

`func (o *MicrosoftGraphInferenceClassificationOverride) SetSenderEmailAddressExplicitNull(b bool)`

SetSenderEmailAddressExplicitNull (un)sets SenderEmailAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SenderEmailAddress value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


