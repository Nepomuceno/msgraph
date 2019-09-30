# MicrosoftGraphWorkbookChart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Height** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**Left** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Top** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**Width** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**Axes** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartAxes**](anyOf&lt;microsoft.graph.workbookChartAxes&gt;.md) |  | [optional] 
**DataLabels** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartDataLabels**](anyOf&lt;microsoft.graph.workbookChartDataLabels&gt;.md) |  | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartAreaFormat**](anyOf&lt;microsoft.graph.workbookChartAreaFormat&gt;.md) |  | [optional] 
**Legend** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartLegend**](anyOf&lt;microsoft.graph.workbookChartLegend&gt;.md) |  | [optional] 
**Series** | Pointer to [**[]MicrosoftGraphWorkbookChartSeries**](microsoft.graph.workbookChartSeries.md) |  | [optional] 
**Title** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartTitle**](anyOf&lt;microsoft.graph.workbookChartTitle&gt;.md) |  | [optional] 
**Worksheet** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheet**](anyOf&lt;microsoft.graph.workbookWorksheet&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWorkbookChart) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChart) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookChart) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChart) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetHeight

`func (o *MicrosoftGraphWorkbookChart) GetHeight() AnyOfnumberstringstring`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *MicrosoftGraphWorkbookChart) GetHeightOk() (AnyOfnumberstringstring, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeight

`func (o *MicrosoftGraphWorkbookChart) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeight

`func (o *MicrosoftGraphWorkbookChart) SetHeight(v AnyOfnumberstringstring)`

SetHeight gets a reference to the given AnyOfnumberstringstring and assigns it to the Height field.

### GetLeft

`func (o *MicrosoftGraphWorkbookChart) GetLeft() AnyOfnumberstringstring`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *MicrosoftGraphWorkbookChart) GetLeftOk() (AnyOfnumberstringstring, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLeft

`func (o *MicrosoftGraphWorkbookChart) HasLeft() bool`

HasLeft returns a boolean if a field has been set.

### SetLeft

`func (o *MicrosoftGraphWorkbookChart) SetLeft(v AnyOfnumberstringstring)`

SetLeft gets a reference to the given AnyOfnumberstringstring and assigns it to the Left field.

### GetName

`func (o *MicrosoftGraphWorkbookChart) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphWorkbookChart) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphWorkbookChart) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphWorkbookChart) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphWorkbookChart) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetTop

`func (o *MicrosoftGraphWorkbookChart) GetTop() AnyOfnumberstringstring`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *MicrosoftGraphWorkbookChart) GetTopOk() (AnyOfnumberstringstring, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTop

`func (o *MicrosoftGraphWorkbookChart) HasTop() bool`

HasTop returns a boolean if a field has been set.

### SetTop

`func (o *MicrosoftGraphWorkbookChart) SetTop(v AnyOfnumberstringstring)`

SetTop gets a reference to the given AnyOfnumberstringstring and assigns it to the Top field.

### GetWidth

`func (o *MicrosoftGraphWorkbookChart) GetWidth() AnyOfnumberstringstring`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MicrosoftGraphWorkbookChart) GetWidthOk() (AnyOfnumberstringstring, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWidth

`func (o *MicrosoftGraphWorkbookChart) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidth

`func (o *MicrosoftGraphWorkbookChart) SetWidth(v AnyOfnumberstringstring)`

SetWidth gets a reference to the given AnyOfnumberstringstring and assigns it to the Width field.

### GetAxes

`func (o *MicrosoftGraphWorkbookChart) GetAxes() AnyOfmicrosoftGraphWorkbookChartAxes`

GetAxes returns the Axes field if non-nil, zero value otherwise.

### GetAxesOk

`func (o *MicrosoftGraphWorkbookChart) GetAxesOk() (AnyOfmicrosoftGraphWorkbookChartAxes, bool)`

GetAxesOk returns a tuple with the Axes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAxes

`func (o *MicrosoftGraphWorkbookChart) HasAxes() bool`

HasAxes returns a boolean if a field has been set.

### SetAxes

`func (o *MicrosoftGraphWorkbookChart) SetAxes(v AnyOfmicrosoftGraphWorkbookChartAxes)`

SetAxes gets a reference to the given AnyOfmicrosoftGraphWorkbookChartAxes and assigns it to the Axes field.

### SetAxesExplicitNull

`func (o *MicrosoftGraphWorkbookChart) SetAxesExplicitNull(b bool)`

SetAxesExplicitNull (un)sets Axes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Axes value is set to nil even if false is passed
### GetDataLabels

`func (o *MicrosoftGraphWorkbookChart) GetDataLabels() AnyOfmicrosoftGraphWorkbookChartDataLabels`

GetDataLabels returns the DataLabels field if non-nil, zero value otherwise.

### GetDataLabelsOk

`func (o *MicrosoftGraphWorkbookChart) GetDataLabelsOk() (AnyOfmicrosoftGraphWorkbookChartDataLabels, bool)`

GetDataLabelsOk returns a tuple with the DataLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDataLabels

`func (o *MicrosoftGraphWorkbookChart) HasDataLabels() bool`

HasDataLabels returns a boolean if a field has been set.

### SetDataLabels

`func (o *MicrosoftGraphWorkbookChart) SetDataLabels(v AnyOfmicrosoftGraphWorkbookChartDataLabels)`

SetDataLabels gets a reference to the given AnyOfmicrosoftGraphWorkbookChartDataLabels and assigns it to the DataLabels field.

### SetDataLabelsExplicitNull

`func (o *MicrosoftGraphWorkbookChart) SetDataLabelsExplicitNull(b bool)`

SetDataLabelsExplicitNull (un)sets DataLabels to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DataLabels value is set to nil even if false is passed
### GetFormat

`func (o *MicrosoftGraphWorkbookChart) GetFormat() AnyOfmicrosoftGraphWorkbookChartAreaFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphWorkbookChart) GetFormatOk() (AnyOfmicrosoftGraphWorkbookChartAreaFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormat

`func (o *MicrosoftGraphWorkbookChart) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormat

`func (o *MicrosoftGraphWorkbookChart) SetFormat(v AnyOfmicrosoftGraphWorkbookChartAreaFormat)`

SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartAreaFormat and assigns it to the Format field.

### SetFormatExplicitNull

`func (o *MicrosoftGraphWorkbookChart) SetFormatExplicitNull(b bool)`

SetFormatExplicitNull (un)sets Format to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Format value is set to nil even if false is passed
### GetLegend

`func (o *MicrosoftGraphWorkbookChart) GetLegend() AnyOfmicrosoftGraphWorkbookChartLegend`

GetLegend returns the Legend field if non-nil, zero value otherwise.

### GetLegendOk

`func (o *MicrosoftGraphWorkbookChart) GetLegendOk() (AnyOfmicrosoftGraphWorkbookChartLegend, bool)`

GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLegend

`func (o *MicrosoftGraphWorkbookChart) HasLegend() bool`

HasLegend returns a boolean if a field has been set.

### SetLegend

`func (o *MicrosoftGraphWorkbookChart) SetLegend(v AnyOfmicrosoftGraphWorkbookChartLegend)`

SetLegend gets a reference to the given AnyOfmicrosoftGraphWorkbookChartLegend and assigns it to the Legend field.

### SetLegendExplicitNull

`func (o *MicrosoftGraphWorkbookChart) SetLegendExplicitNull(b bool)`

SetLegendExplicitNull (un)sets Legend to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Legend value is set to nil even if false is passed
### GetSeries

`func (o *MicrosoftGraphWorkbookChart) GetSeries() []MicrosoftGraphWorkbookChartSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *MicrosoftGraphWorkbookChart) GetSeriesOk() ([]MicrosoftGraphWorkbookChartSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSeries

`func (o *MicrosoftGraphWorkbookChart) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeries

`func (o *MicrosoftGraphWorkbookChart) SetSeries(v []MicrosoftGraphWorkbookChartSeries)`

SetSeries gets a reference to the given []MicrosoftGraphWorkbookChartSeries and assigns it to the Series field.

### GetTitle

`func (o *MicrosoftGraphWorkbookChart) GetTitle() AnyOfmicrosoftGraphWorkbookChartTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphWorkbookChart) GetTitleOk() (AnyOfmicrosoftGraphWorkbookChartTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *MicrosoftGraphWorkbookChart) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *MicrosoftGraphWorkbookChart) SetTitle(v AnyOfmicrosoftGraphWorkbookChartTitle)`

SetTitle gets a reference to the given AnyOfmicrosoftGraphWorkbookChartTitle and assigns it to the Title field.

### SetTitleExplicitNull

`func (o *MicrosoftGraphWorkbookChart) SetTitleExplicitNull(b bool)`

SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Title value is set to nil even if false is passed
### GetWorksheet

`func (o *MicrosoftGraphWorkbookChart) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *MicrosoftGraphWorkbookChart) GetWorksheetOk() (AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorksheet

`func (o *MicrosoftGraphWorkbookChart) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheet

`func (o *MicrosoftGraphWorkbookChart) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.

### SetWorksheetExplicitNull

`func (o *MicrosoftGraphWorkbookChart) SetWorksheetExplicitNull(b bool)`

SetWorksheetExplicitNull (un)sets Worksheet to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Worksheet value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


