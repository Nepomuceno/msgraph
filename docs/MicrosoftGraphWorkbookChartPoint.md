# MicrosoftGraphWorkbookChartPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartPointFormat**](anyOf&lt;microsoft.graph.workbookChartPointFormat&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWorkbookChartPoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartPoint) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookChartPoint) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartPoint) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetValue

`func (o *MicrosoftGraphWorkbookChartPoint) GetValue() AnyOfobject`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphWorkbookChartPoint) GetValueOk() (AnyOfobject, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *MicrosoftGraphWorkbookChartPoint) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *MicrosoftGraphWorkbookChartPoint) SetValue(v AnyOfobject)`

SetValue gets a reference to the given AnyOfobject and assigns it to the Value field.

### SetValueExplicitNull

`func (o *MicrosoftGraphWorkbookChartPoint) SetValueExplicitNull(b bool)`

SetValueExplicitNull (un)sets Value to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Value value is set to nil even if false is passed
### GetFormat

`func (o *MicrosoftGraphWorkbookChartPoint) GetFormat() AnyOfmicrosoftGraphWorkbookChartPointFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphWorkbookChartPoint) GetFormatOk() (AnyOfmicrosoftGraphWorkbookChartPointFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormat

`func (o *MicrosoftGraphWorkbookChartPoint) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormat

`func (o *MicrosoftGraphWorkbookChartPoint) SetFormat(v AnyOfmicrosoftGraphWorkbookChartPointFormat)`

SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartPointFormat and assigns it to the Format field.

### SetFormatExplicitNull

`func (o *MicrosoftGraphWorkbookChartPoint) SetFormatExplicitNull(b bool)`

SetFormatExplicitNull (un)sets Format to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Format value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


