# MicrosoftGraphWorkbookChartAxis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**MajorUnit** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Maximum** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Minimum** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**MinorUnit** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartAxisFormat**](anyOf&lt;microsoft.graph.workbookChartAxisFormat&gt;.md) |  | [optional] 
**MajorGridlines** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartGridlines**](anyOf&lt;microsoft.graph.workbookChartGridlines&gt;.md) |  | [optional] 
**MinorGridlines** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartGridlines**](anyOf&lt;microsoft.graph.workbookChartGridlines&gt;.md) |  | [optional] 
**Title** | Pointer to [**AnyOfmicrosoftGraphWorkbookChartAxisTitle**](anyOf&lt;microsoft.graph.workbookChartAxisTitle&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWorkbookChartAxis) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookChartAxis) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookChartAxis) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetMajorUnit

`func (o *MicrosoftGraphWorkbookChartAxis) GetMajorUnit() AnyOfobject`

GetMajorUnit returns the MajorUnit field if non-nil, zero value otherwise.

### GetMajorUnitOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetMajorUnitOk() (AnyOfobject, bool)`

GetMajorUnitOk returns a tuple with the MajorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMajorUnit

`func (o *MicrosoftGraphWorkbookChartAxis) HasMajorUnit() bool`

HasMajorUnit returns a boolean if a field has been set.

### SetMajorUnit

`func (o *MicrosoftGraphWorkbookChartAxis) SetMajorUnit(v AnyOfobject)`

SetMajorUnit gets a reference to the given AnyOfobject and assigns it to the MajorUnit field.

### SetMajorUnitExplicitNull

`func (o *MicrosoftGraphWorkbookChartAxis) SetMajorUnitExplicitNull(b bool)`

SetMajorUnitExplicitNull (un)sets MajorUnit to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MajorUnit value is set to nil even if false is passed
### GetMaximum

`func (o *MicrosoftGraphWorkbookChartAxis) GetMaximum() AnyOfobject`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetMaximumOk() (AnyOfobject, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaximum

`func (o *MicrosoftGraphWorkbookChartAxis) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### SetMaximum

`func (o *MicrosoftGraphWorkbookChartAxis) SetMaximum(v AnyOfobject)`

SetMaximum gets a reference to the given AnyOfobject and assigns it to the Maximum field.

### SetMaximumExplicitNull

`func (o *MicrosoftGraphWorkbookChartAxis) SetMaximumExplicitNull(b bool)`

SetMaximumExplicitNull (un)sets Maximum to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Maximum value is set to nil even if false is passed
### GetMinimum

`func (o *MicrosoftGraphWorkbookChartAxis) GetMinimum() AnyOfobject`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetMinimumOk() (AnyOfobject, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimum

`func (o *MicrosoftGraphWorkbookChartAxis) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### SetMinimum

`func (o *MicrosoftGraphWorkbookChartAxis) SetMinimum(v AnyOfobject)`

SetMinimum gets a reference to the given AnyOfobject and assigns it to the Minimum field.

### SetMinimumExplicitNull

`func (o *MicrosoftGraphWorkbookChartAxis) SetMinimumExplicitNull(b bool)`

SetMinimumExplicitNull (un)sets Minimum to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Minimum value is set to nil even if false is passed
### GetMinorUnit

`func (o *MicrosoftGraphWorkbookChartAxis) GetMinorUnit() AnyOfobject`

GetMinorUnit returns the MinorUnit field if non-nil, zero value otherwise.

### GetMinorUnitOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetMinorUnitOk() (AnyOfobject, bool)`

GetMinorUnitOk returns a tuple with the MinorUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinorUnit

`func (o *MicrosoftGraphWorkbookChartAxis) HasMinorUnit() bool`

HasMinorUnit returns a boolean if a field has been set.

### SetMinorUnit

`func (o *MicrosoftGraphWorkbookChartAxis) SetMinorUnit(v AnyOfobject)`

SetMinorUnit gets a reference to the given AnyOfobject and assigns it to the MinorUnit field.

### SetMinorUnitExplicitNull

`func (o *MicrosoftGraphWorkbookChartAxis) SetMinorUnitExplicitNull(b bool)`

SetMinorUnitExplicitNull (un)sets MinorUnit to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinorUnit value is set to nil even if false is passed
### GetFormat

`func (o *MicrosoftGraphWorkbookChartAxis) GetFormat() AnyOfmicrosoftGraphWorkbookChartAxisFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetFormatOk() (AnyOfmicrosoftGraphWorkbookChartAxisFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormat

`func (o *MicrosoftGraphWorkbookChartAxis) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormat

`func (o *MicrosoftGraphWorkbookChartAxis) SetFormat(v AnyOfmicrosoftGraphWorkbookChartAxisFormat)`

SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartAxisFormat and assigns it to the Format field.

### SetFormatExplicitNull

`func (o *MicrosoftGraphWorkbookChartAxis) SetFormatExplicitNull(b bool)`

SetFormatExplicitNull (un)sets Format to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Format value is set to nil even if false is passed
### GetMajorGridlines

`func (o *MicrosoftGraphWorkbookChartAxis) GetMajorGridlines() AnyOfmicrosoftGraphWorkbookChartGridlines`

GetMajorGridlines returns the MajorGridlines field if non-nil, zero value otherwise.

### GetMajorGridlinesOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetMajorGridlinesOk() (AnyOfmicrosoftGraphWorkbookChartGridlines, bool)`

GetMajorGridlinesOk returns a tuple with the MajorGridlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMajorGridlines

`func (o *MicrosoftGraphWorkbookChartAxis) HasMajorGridlines() bool`

HasMajorGridlines returns a boolean if a field has been set.

### SetMajorGridlines

`func (o *MicrosoftGraphWorkbookChartAxis) SetMajorGridlines(v AnyOfmicrosoftGraphWorkbookChartGridlines)`

SetMajorGridlines gets a reference to the given AnyOfmicrosoftGraphWorkbookChartGridlines and assigns it to the MajorGridlines field.

### SetMajorGridlinesExplicitNull

`func (o *MicrosoftGraphWorkbookChartAxis) SetMajorGridlinesExplicitNull(b bool)`

SetMajorGridlinesExplicitNull (un)sets MajorGridlines to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MajorGridlines value is set to nil even if false is passed
### GetMinorGridlines

`func (o *MicrosoftGraphWorkbookChartAxis) GetMinorGridlines() AnyOfmicrosoftGraphWorkbookChartGridlines`

GetMinorGridlines returns the MinorGridlines field if non-nil, zero value otherwise.

### GetMinorGridlinesOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetMinorGridlinesOk() (AnyOfmicrosoftGraphWorkbookChartGridlines, bool)`

GetMinorGridlinesOk returns a tuple with the MinorGridlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinorGridlines

`func (o *MicrosoftGraphWorkbookChartAxis) HasMinorGridlines() bool`

HasMinorGridlines returns a boolean if a field has been set.

### SetMinorGridlines

`func (o *MicrosoftGraphWorkbookChartAxis) SetMinorGridlines(v AnyOfmicrosoftGraphWorkbookChartGridlines)`

SetMinorGridlines gets a reference to the given AnyOfmicrosoftGraphWorkbookChartGridlines and assigns it to the MinorGridlines field.

### SetMinorGridlinesExplicitNull

`func (o *MicrosoftGraphWorkbookChartAxis) SetMinorGridlinesExplicitNull(b bool)`

SetMinorGridlinesExplicitNull (un)sets MinorGridlines to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinorGridlines value is set to nil even if false is passed
### GetTitle

`func (o *MicrosoftGraphWorkbookChartAxis) GetTitle() AnyOfmicrosoftGraphWorkbookChartAxisTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphWorkbookChartAxis) GetTitleOk() (AnyOfmicrosoftGraphWorkbookChartAxisTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *MicrosoftGraphWorkbookChartAxis) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *MicrosoftGraphWorkbookChartAxis) SetTitle(v AnyOfmicrosoftGraphWorkbookChartAxisTitle)`

SetTitle gets a reference to the given AnyOfmicrosoftGraphWorkbookChartAxisTitle and assigns it to the Title field.

### SetTitleExplicitNull

`func (o *MicrosoftGraphWorkbookChartAxis) SetTitleExplicitNull(b bool)`

SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Title value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


