# WorkbookRangeFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnWidth** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**HorizontalAlignment** | Pointer to **string** |  | [optional] 
**RowHeight** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**VerticalAlignment** | Pointer to **string** |  | [optional] 
**WrapText** | Pointer to **bool** |  | [optional] 
**Borders** | Pointer to [**[]MicrosoftGraphWorkbookRangeBorder**](microsoft.graph.workbookRangeBorder.md) |  | [optional] 
**Fill** | Pointer to [**AnyOfmicrosoftGraphWorkbookRangeFill**](anyOf&lt;microsoft.graph.workbookRangeFill&gt;.md) |  | [optional] 
**Font** | Pointer to [**AnyOfmicrosoftGraphWorkbookRangeFont**](anyOf&lt;microsoft.graph.workbookRangeFont&gt;.md) |  | [optional] 
**Protection** | Pointer to [**AnyOfmicrosoftGraphWorkbookFormatProtection**](anyOf&lt;microsoft.graph.workbookFormatProtection&gt;.md) |  | [optional] 

## Methods

### GetColumnWidth

`func (o *WorkbookRangeFormat) GetColumnWidth() AnyOfnumberstringstring`

GetColumnWidth returns the ColumnWidth field if non-nil, zero value otherwise.

### GetColumnWidthOk

`func (o *WorkbookRangeFormat) GetColumnWidthOk() (AnyOfnumberstringstring, bool)`

GetColumnWidthOk returns a tuple with the ColumnWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnWidth

`func (o *WorkbookRangeFormat) HasColumnWidth() bool`

HasColumnWidth returns a boolean if a field has been set.

### SetColumnWidth

`func (o *WorkbookRangeFormat) SetColumnWidth(v AnyOfnumberstringstring)`

SetColumnWidth gets a reference to the given AnyOfnumberstringstring and assigns it to the ColumnWidth field.

### SetColumnWidthExplicitNull

`func (o *WorkbookRangeFormat) SetColumnWidthExplicitNull(b bool)`

SetColumnWidthExplicitNull (un)sets ColumnWidth to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ColumnWidth value is set to nil even if false is passed
### GetHorizontalAlignment

`func (o *WorkbookRangeFormat) GetHorizontalAlignment() string`

GetHorizontalAlignment returns the HorizontalAlignment field if non-nil, zero value otherwise.

### GetHorizontalAlignmentOk

`func (o *WorkbookRangeFormat) GetHorizontalAlignmentOk() (string, bool)`

GetHorizontalAlignmentOk returns a tuple with the HorizontalAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHorizontalAlignment

`func (o *WorkbookRangeFormat) HasHorizontalAlignment() bool`

HasHorizontalAlignment returns a boolean if a field has been set.

### SetHorizontalAlignment

`func (o *WorkbookRangeFormat) SetHorizontalAlignment(v string)`

SetHorizontalAlignment gets a reference to the given string and assigns it to the HorizontalAlignment field.

### SetHorizontalAlignmentExplicitNull

`func (o *WorkbookRangeFormat) SetHorizontalAlignmentExplicitNull(b bool)`

SetHorizontalAlignmentExplicitNull (un)sets HorizontalAlignment to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HorizontalAlignment value is set to nil even if false is passed
### GetRowHeight

`func (o *WorkbookRangeFormat) GetRowHeight() AnyOfnumberstringstring`

GetRowHeight returns the RowHeight field if non-nil, zero value otherwise.

### GetRowHeightOk

`func (o *WorkbookRangeFormat) GetRowHeightOk() (AnyOfnumberstringstring, bool)`

GetRowHeightOk returns a tuple with the RowHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRowHeight

`func (o *WorkbookRangeFormat) HasRowHeight() bool`

HasRowHeight returns a boolean if a field has been set.

### SetRowHeight

`func (o *WorkbookRangeFormat) SetRowHeight(v AnyOfnumberstringstring)`

SetRowHeight gets a reference to the given AnyOfnumberstringstring and assigns it to the RowHeight field.

### SetRowHeightExplicitNull

`func (o *WorkbookRangeFormat) SetRowHeightExplicitNull(b bool)`

SetRowHeightExplicitNull (un)sets RowHeight to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RowHeight value is set to nil even if false is passed
### GetVerticalAlignment

`func (o *WorkbookRangeFormat) GetVerticalAlignment() string`

GetVerticalAlignment returns the VerticalAlignment field if non-nil, zero value otherwise.

### GetVerticalAlignmentOk

`func (o *WorkbookRangeFormat) GetVerticalAlignmentOk() (string, bool)`

GetVerticalAlignmentOk returns a tuple with the VerticalAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVerticalAlignment

`func (o *WorkbookRangeFormat) HasVerticalAlignment() bool`

HasVerticalAlignment returns a boolean if a field has been set.

### SetVerticalAlignment

`func (o *WorkbookRangeFormat) SetVerticalAlignment(v string)`

SetVerticalAlignment gets a reference to the given string and assigns it to the VerticalAlignment field.

### SetVerticalAlignmentExplicitNull

`func (o *WorkbookRangeFormat) SetVerticalAlignmentExplicitNull(b bool)`

SetVerticalAlignmentExplicitNull (un)sets VerticalAlignment to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VerticalAlignment value is set to nil even if false is passed
### GetWrapText

`func (o *WorkbookRangeFormat) GetWrapText() bool`

GetWrapText returns the WrapText field if non-nil, zero value otherwise.

### GetWrapTextOk

`func (o *WorkbookRangeFormat) GetWrapTextOk() (bool, bool)`

GetWrapTextOk returns a tuple with the WrapText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWrapText

`func (o *WorkbookRangeFormat) HasWrapText() bool`

HasWrapText returns a boolean if a field has been set.

### SetWrapText

`func (o *WorkbookRangeFormat) SetWrapText(v bool)`

SetWrapText gets a reference to the given bool and assigns it to the WrapText field.

### SetWrapTextExplicitNull

`func (o *WorkbookRangeFormat) SetWrapTextExplicitNull(b bool)`

SetWrapTextExplicitNull (un)sets WrapText to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WrapText value is set to nil even if false is passed
### GetBorders

`func (o *WorkbookRangeFormat) GetBorders() []MicrosoftGraphWorkbookRangeBorder`

GetBorders returns the Borders field if non-nil, zero value otherwise.

### GetBordersOk

`func (o *WorkbookRangeFormat) GetBordersOk() ([]MicrosoftGraphWorkbookRangeBorder, bool)`

GetBordersOk returns a tuple with the Borders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBorders

`func (o *WorkbookRangeFormat) HasBorders() bool`

HasBorders returns a boolean if a field has been set.

### SetBorders

`func (o *WorkbookRangeFormat) SetBorders(v []MicrosoftGraphWorkbookRangeBorder)`

SetBorders gets a reference to the given []MicrosoftGraphWorkbookRangeBorder and assigns it to the Borders field.

### GetFill

`func (o *WorkbookRangeFormat) GetFill() AnyOfmicrosoftGraphWorkbookRangeFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *WorkbookRangeFormat) GetFillOk() (AnyOfmicrosoftGraphWorkbookRangeFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFill

`func (o *WorkbookRangeFormat) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFill

`func (o *WorkbookRangeFormat) SetFill(v AnyOfmicrosoftGraphWorkbookRangeFill)`

SetFill gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeFill and assigns it to the Fill field.

### SetFillExplicitNull

`func (o *WorkbookRangeFormat) SetFillExplicitNull(b bool)`

SetFillExplicitNull (un)sets Fill to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Fill value is set to nil even if false is passed
### GetFont

`func (o *WorkbookRangeFormat) GetFont() AnyOfmicrosoftGraphWorkbookRangeFont`

GetFont returns the Font field if non-nil, zero value otherwise.

### GetFontOk

`func (o *WorkbookRangeFormat) GetFontOk() (AnyOfmicrosoftGraphWorkbookRangeFont, bool)`

GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFont

`func (o *WorkbookRangeFormat) HasFont() bool`

HasFont returns a boolean if a field has been set.

### SetFont

`func (o *WorkbookRangeFormat) SetFont(v AnyOfmicrosoftGraphWorkbookRangeFont)`

SetFont gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeFont and assigns it to the Font field.

### SetFontExplicitNull

`func (o *WorkbookRangeFormat) SetFontExplicitNull(b bool)`

SetFontExplicitNull (un)sets Font to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Font value is set to nil even if false is passed
### GetProtection

`func (o *WorkbookRangeFormat) GetProtection() AnyOfmicrosoftGraphWorkbookFormatProtection`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *WorkbookRangeFormat) GetProtectionOk() (AnyOfmicrosoftGraphWorkbookFormatProtection, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtection

`func (o *WorkbookRangeFormat) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### SetProtection

`func (o *WorkbookRangeFormat) SetProtection(v AnyOfmicrosoftGraphWorkbookFormatProtection)`

SetProtection gets a reference to the given AnyOfmicrosoftGraphWorkbookFormatProtection and assigns it to the Protection field.

### SetProtectionExplicitNull

`func (o *WorkbookRangeFormat) SetProtectionExplicitNull(b bool)`

SetProtectionExplicitNull (un)sets Protection to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Protection value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


