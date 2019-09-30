# MicrosoftGraphOnenotePagePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviewText** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphOnenotePagePreviewLinks**](anyOf&lt;microsoft.graph.onenotePagePreviewLinks&gt;.md) |  | [optional] 

## Methods

### GetPreviewText

`func (o *MicrosoftGraphOnenotePagePreview) GetPreviewText() string`

GetPreviewText returns the PreviewText field if non-nil, zero value otherwise.

### GetPreviewTextOk

`func (o *MicrosoftGraphOnenotePagePreview) GetPreviewTextOk() (string, bool)`

GetPreviewTextOk returns a tuple with the PreviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreviewText

`func (o *MicrosoftGraphOnenotePagePreview) HasPreviewText() bool`

HasPreviewText returns a boolean if a field has been set.

### SetPreviewText

`func (o *MicrosoftGraphOnenotePagePreview) SetPreviewText(v string)`

SetPreviewText gets a reference to the given string and assigns it to the PreviewText field.

### SetPreviewTextExplicitNull

`func (o *MicrosoftGraphOnenotePagePreview) SetPreviewTextExplicitNull(b bool)`

SetPreviewTextExplicitNull (un)sets PreviewText to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreviewText value is set to nil even if false is passed
### GetLinks

`func (o *MicrosoftGraphOnenotePagePreview) GetLinks() AnyOfmicrosoftGraphOnenotePagePreviewLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MicrosoftGraphOnenotePagePreview) GetLinksOk() (AnyOfmicrosoftGraphOnenotePagePreviewLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinks

`func (o *MicrosoftGraphOnenotePagePreview) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinks

`func (o *MicrosoftGraphOnenotePagePreview) SetLinks(v AnyOfmicrosoftGraphOnenotePagePreviewLinks)`

SetLinks gets a reference to the given AnyOfmicrosoftGraphOnenotePagePreviewLinks and assigns it to the Links field.

### SetLinksExplicitNull

`func (o *MicrosoftGraphOnenotePagePreview) SetLinksExplicitNull(b bool)`

SetLinksExplicitNull (un)sets Links to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Links value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


