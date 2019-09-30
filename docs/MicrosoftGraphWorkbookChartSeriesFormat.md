# MicrosoftGraphWorkbookChartSeriesFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Fill** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartFill**](anyOf&lt;microsoft.graph.workbookChartFill&gt;.md) |  | [optional] 
**Line** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartLineFormat**](anyOf&lt;microsoft.graph.workbookChartLineFormat&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetFill

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) GetFill() AnyOfmicrosoftGraphWorkbookChartFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) GetFillOk() (AnyOfmicrosoftGraphWorkbookChartFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFill

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFill

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) SetFill(v AnyOfmicrosoftGraphWorkbookChartFill)`

SetFill gets a reference to the given AnyOfmicrosoftGraphWorkbookChartFill and assigns it to the Fill field.

### SetFillExplicitNull

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) SetFillExplicitNull(b bool)`

SetFillExplicitNull (un)sets Fill to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Fill value is set to nil even if false is passed
### GetLine

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) GetLine() AnyOfmicrosoftGraphWorkbookChartLineFormat`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) GetLineOk() (AnyOfmicrosoftGraphWorkbookChartLineFormat, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLine

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) HasLine() bool`

HasLine returns a boolean if a field has been set.

### SetLine

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) SetLine(v AnyOfmicrosoftGraphWorkbookChartLineFormat)`

SetLine gets a reference to the given AnyOfmicrosoftGraphWorkbookChartLineFormat and assigns it to the Line field.

### SetLineExplicitNull

`func (o *MicrosoftGraphWorkbookChartSeriesFormat) SetLineExplicitNull(b bool)`

SetLineExplicitNull (un)sets Line to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Line value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


