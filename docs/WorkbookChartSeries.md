# WorkbookChartSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartSeriesFormat**](anyOf&lt;microsoft.graph.workbookChartSeriesFormat&gt;.md) |  | [optional] 
**Points** | Pointer to [**[]MicrosoftGraphWorkbookChartPoint**](microsoft.graph.workbookChartPoint.md) |  | [optional] 

## Methods

### GetName

`func (o *WorkbookChartSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkbookChartSeries) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *WorkbookChartSeries) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *WorkbookChartSeries) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *WorkbookChartSeries) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetFormat

`func (o *WorkbookChartSeries) GetFormat() AnyOfmicrosoftGraphWorkbookChartSeriesFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WorkbookChartSeries) GetFormatOk() (AnyOfmicrosoftGraphWorkbookChartSeriesFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormat

`func (o *WorkbookChartSeries) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormat

`func (o *WorkbookChartSeries) SetFormat(v AnyOfmicrosoftGraphWorkbookChartSeriesFormat)`

SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartSeriesFormat and assigns it to the Format field.

### SetFormatExplicitNull

`func (o *WorkbookChartSeries) SetFormatExplicitNull(b bool)`

SetFormatExplicitNull (un)sets Format to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Format value is set to nil even if false is passed
### GetPoints

`func (o *WorkbookChartSeries) GetPoints() []MicrosoftGraphWorkbookChartPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *WorkbookChartSeries) GetPointsOk() ([]MicrosoftGraphWorkbookChartPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPoints

`func (o *WorkbookChartSeries) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### SetPoints

`func (o *WorkbookChartSeries) SetPoints(v []MicrosoftGraphWorkbookChartPoint)`

SetPoints gets a reference to the given []MicrosoftGraphWorkbookChartPoint and assigns it to the Points field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


