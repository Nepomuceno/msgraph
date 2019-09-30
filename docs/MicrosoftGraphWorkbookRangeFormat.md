# MicrosoftGraphWorkbookRangeFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphWorkbookRangeFormat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookRangeFormat) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookRangeFormat) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetColumnWidth

`func (o *MicrosoftGraphWorkbookRangeFormat) GetColumnWidth() AnyOfnumberstringstring`

GetColumnWidth returns the ColumnWidth field if non-nil, zero value otherwise.

### GetColumnWidthOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetColumnWidthOk() (AnyOfnumberstringstring, bool)`

GetColumnWidthOk returns a tuple with the ColumnWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnWidth

`func (o *MicrosoftGraphWorkbookRangeFormat) HasColumnWidth() bool`

HasColumnWidth returns a boolean if a field has been set.

### SetColumnWidth

`func (o *MicrosoftGraphWorkbookRangeFormat) SetColumnWidth(v AnyOfnumberstringstring)`

SetColumnWidth gets a reference to the given AnyOfnumberstringstring and assigns it to the ColumnWidth field.

### SetColumnWidthExplicitNull

`func (o *MicrosoftGraphWorkbookRangeFormat) SetColumnWidthExplicitNull(b bool)`

SetColumnWidthExplicitNull (un)sets ColumnWidth to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ColumnWidth value is set to nil even if false is passed
### GetHorizontalAlignment

`func (o *MicrosoftGraphWorkbookRangeFormat) GetHorizontalAlignment() string`

GetHorizontalAlignment returns the HorizontalAlignment field if non-nil, zero value otherwise.

### GetHorizontalAlignmentOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetHorizontalAlignmentOk() (string, bool)`

GetHorizontalAlignmentOk returns a tuple with the HorizontalAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHorizontalAlignment

`func (o *MicrosoftGraphWorkbookRangeFormat) HasHorizontalAlignment() bool`

HasHorizontalAlignment returns a boolean if a field has been set.

### SetHorizontalAlignment

`func (o *MicrosoftGraphWorkbookRangeFormat) SetHorizontalAlignment(v string)`

SetHorizontalAlignment gets a reference to the given string and assigns it to the HorizontalAlignment field.

### SetHorizontalAlignmentExplicitNull

`func (o *MicrosoftGraphWorkbookRangeFormat) SetHorizontalAlignmentExplicitNull(b bool)`

SetHorizontalAlignmentExplicitNull (un)sets HorizontalAlignment to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HorizontalAlignment value is set to nil even if false is passed
### GetRowHeight

`func (o *MicrosoftGraphWorkbookRangeFormat) GetRowHeight() AnyOfnumberstringstring`

GetRowHeight returns the RowHeight field if non-nil, zero value otherwise.

### GetRowHeightOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetRowHeightOk() (AnyOfnumberstringstring, bool)`

GetRowHeightOk returns a tuple with the RowHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRowHeight

`func (o *MicrosoftGraphWorkbookRangeFormat) HasRowHeight() bool`

HasRowHeight returns a boolean if a field has been set.

### SetRowHeight

`func (o *MicrosoftGraphWorkbookRangeFormat) SetRowHeight(v AnyOfnumberstringstring)`

SetRowHeight gets a reference to the given AnyOfnumberstringstring and assigns it to the RowHeight field.

### SetRowHeightExplicitNull

`func (o *MicrosoftGraphWorkbookRangeFormat) SetRowHeightExplicitNull(b bool)`

SetRowHeightExplicitNull (un)sets RowHeight to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RowHeight value is set to nil even if false is passed
### GetVerticalAlignment

`func (o *MicrosoftGraphWorkbookRangeFormat) GetVerticalAlignment() string`

GetVerticalAlignment returns the VerticalAlignment field if non-nil, zero value otherwise.

### GetVerticalAlignmentOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetVerticalAlignmentOk() (string, bool)`

GetVerticalAlignmentOk returns a tuple with the VerticalAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVerticalAlignment

`func (o *MicrosoftGraphWorkbookRangeFormat) HasVerticalAlignment() bool`

HasVerticalAlignment returns a boolean if a field has been set.

### SetVerticalAlignment

`func (o *MicrosoftGraphWorkbookRangeFormat) SetVerticalAlignment(v string)`

SetVerticalAlignment gets a reference to the given string and assigns it to the VerticalAlignment field.

### SetVerticalAlignmentExplicitNull

`func (o *MicrosoftGraphWorkbookRangeFormat) SetVerticalAlignmentExplicitNull(b bool)`

SetVerticalAlignmentExplicitNull (un)sets VerticalAlignment to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VerticalAlignment value is set to nil even if false is passed
### GetWrapText

`func (o *MicrosoftGraphWorkbookRangeFormat) GetWrapText() bool`

GetWrapText returns the WrapText field if non-nil, zero value otherwise.

### GetWrapTextOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetWrapTextOk() (bool, bool)`

GetWrapTextOk returns a tuple with the WrapText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWrapText

`func (o *MicrosoftGraphWorkbookRangeFormat) HasWrapText() bool`

HasWrapText returns a boolean if a field has been set.

### SetWrapText

`func (o *MicrosoftGraphWorkbookRangeFormat) SetWrapText(v bool)`

SetWrapText gets a reference to the given bool and assigns it to the WrapText field.

### SetWrapTextExplicitNull

`func (o *MicrosoftGraphWorkbookRangeFormat) SetWrapTextExplicitNull(b bool)`

SetWrapTextExplicitNull (un)sets WrapText to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WrapText value is set to nil even if false is passed
### GetBorders

`func (o *MicrosoftGraphWorkbookRangeFormat) GetBorders() []MicrosoftGraphWorkbookRangeBorder`

GetBorders returns the Borders field if non-nil, zero value otherwise.

### GetBordersOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetBordersOk() ([]MicrosoftGraphWorkbookRangeBorder, bool)`

GetBordersOk returns a tuple with the Borders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBorders

`func (o *MicrosoftGraphWorkbookRangeFormat) HasBorders() bool`

HasBorders returns a boolean if a field has been set.

### SetBorders

`func (o *MicrosoftGraphWorkbookRangeFormat) SetBorders(v []MicrosoftGraphWorkbookRangeBorder)`

SetBorders gets a reference to the given []MicrosoftGraphWorkbookRangeBorder and assigns it to the Borders field.

### GetFill

`func (o *MicrosoftGraphWorkbookRangeFormat) GetFill() AnyOfmicrosoftGraphWorkbookRangeFill`

GetFill returns the Fill field if non-nil, zero value otherwise.

### GetFillOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetFillOk() (AnyOfmicrosoftGraphWorkbookRangeFill, bool)`

GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFill

`func (o *MicrosoftGraphWorkbookRangeFormat) HasFill() bool`

HasFill returns a boolean if a field has been set.

### SetFill

`func (o *MicrosoftGraphWorkbookRangeFormat) SetFill(v AnyOfmicrosoftGraphWorkbookRangeFill)`

SetFill gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeFill and assigns it to the Fill field.

### SetFillExplicitNull

`func (o *MicrosoftGraphWorkbookRangeFormat) SetFillExplicitNull(b bool)`

SetFillExplicitNull (un)sets Fill to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Fill value is set to nil even if false is passed
### GetFont

`func (o *MicrosoftGraphWorkbookRangeFormat) GetFont() AnyOfmicrosoftGraphWorkbookRangeFont`

GetFont returns the Font field if non-nil, zero value otherwise.

### GetFontOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetFontOk() (AnyOfmicrosoftGraphWorkbookRangeFont, bool)`

GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFont

`func (o *MicrosoftGraphWorkbookRangeFormat) HasFont() bool`

HasFont returns a boolean if a field has been set.

### SetFont

`func (o *MicrosoftGraphWorkbookRangeFormat) SetFont(v AnyOfmicrosoftGraphWorkbookRangeFont)`

SetFont gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeFont and assigns it to the Font field.

### SetFontExplicitNull

`func (o *MicrosoftGraphWorkbookRangeFormat) SetFontExplicitNull(b bool)`

SetFontExplicitNull (un)sets Font to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Font value is set to nil even if false is passed
### GetProtection

`func (o *MicrosoftGraphWorkbookRangeFormat) GetProtection() AnyOfmicrosoftGraphWorkbookFormatProtection`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *MicrosoftGraphWorkbookRangeFormat) GetProtectionOk() (AnyOfmicrosoftGraphWorkbookFormatProtection, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtection

`func (o *MicrosoftGraphWorkbookRangeFormat) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### SetProtection

`func (o *MicrosoftGraphWorkbookRangeFormat) SetProtection(v AnyOfmicrosoftGraphWorkbookFormatProtection)`

SetProtection gets a reference to the given AnyOfmicrosoftGraphWorkbookFormatProtection and assigns it to the Protection field.

### SetProtectionExplicitNull

`func (o *MicrosoftGraphWorkbookRangeFormat) SetProtectionExplicitNull(b bool)`

SetProtectionExplicitNull (un)sets Protection to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Protection value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


