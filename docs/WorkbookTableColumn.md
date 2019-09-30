# WorkbookTableColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Values** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Filter** | Pointer to [**AnyOfmicrosoftGraphWorkbookFilter**](anyOf&lt;microsoft.graph.workbookFilter&gt;.md) |  | [optional] 

## Methods

### GetIndex

`func (o *WorkbookTableColumn) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *WorkbookTableColumn) GetIndexOk() (int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndex

`func (o *WorkbookTableColumn) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndex

`func (o *WorkbookTableColumn) SetIndex(v int32)`

SetIndex gets a reference to the given int32 and assigns it to the Index field.

### GetName

`func (o *WorkbookTableColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkbookTableColumn) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *WorkbookTableColumn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *WorkbookTableColumn) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *WorkbookTableColumn) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetValues

`func (o *WorkbookTableColumn) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *WorkbookTableColumn) GetValuesOk() (AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *WorkbookTableColumn) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *WorkbookTableColumn) SetValues(v AnyOfobject)`

SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.

### SetValuesExplicitNull

`func (o *WorkbookTableColumn) SetValuesExplicitNull(b bool)`

SetValuesExplicitNull (un)sets Values to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Values value is set to nil even if false is passed
### GetFilter

`func (o *WorkbookTableColumn) GetFilter() AnyOfmicrosoftGraphWorkbookFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *WorkbookTableColumn) GetFilterOk() (AnyOfmicrosoftGraphWorkbookFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilter

`func (o *WorkbookTableColumn) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilter

`func (o *WorkbookTableColumn) SetFilter(v AnyOfmicrosoftGraphWorkbookFilter)`

SetFilter gets a reference to the given AnyOfmicrosoftGraphWorkbookFilter and assigns it to the Filter field.

### SetFilterExplicitNull

`func (o *WorkbookTableColumn) SetFilterExplicitNull(b bool)`

SetFilterExplicitNull (un)sets Filter to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Filter value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


