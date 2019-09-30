# MicrosoftGraphWorkbookTableColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Values** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Filter** | Pointer to [**AnyOfmicrosoftGraphWorkbookFilter**](anyOf&lt;microsoft.graph.workbookFilter&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWorkbookTableColumn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookTableColumn) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookTableColumn) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookTableColumn) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetIndex

`func (o *MicrosoftGraphWorkbookTableColumn) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MicrosoftGraphWorkbookTableColumn) GetIndexOk() (int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndex

`func (o *MicrosoftGraphWorkbookTableColumn) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndex

`func (o *MicrosoftGraphWorkbookTableColumn) SetIndex(v int32)`

SetIndex gets a reference to the given int32 and assigns it to the Index field.

### GetName

`func (o *MicrosoftGraphWorkbookTableColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphWorkbookTableColumn) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphWorkbookTableColumn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphWorkbookTableColumn) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphWorkbookTableColumn) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetValues

`func (o *MicrosoftGraphWorkbookTableColumn) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MicrosoftGraphWorkbookTableColumn) GetValuesOk() (AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *MicrosoftGraphWorkbookTableColumn) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *MicrosoftGraphWorkbookTableColumn) SetValues(v AnyOfobject)`

SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.

### SetValuesExplicitNull

`func (o *MicrosoftGraphWorkbookTableColumn) SetValuesExplicitNull(b bool)`

SetValuesExplicitNull (un)sets Values to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Values value is set to nil even if false is passed
### GetFilter

`func (o *MicrosoftGraphWorkbookTableColumn) GetFilter() AnyOfmicrosoftGraphWorkbookFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MicrosoftGraphWorkbookTableColumn) GetFilterOk() (AnyOfmicrosoftGraphWorkbookFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilter

`func (o *MicrosoftGraphWorkbookTableColumn) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilter

`func (o *MicrosoftGraphWorkbookTableColumn) SetFilter(v AnyOfmicrosoftGraphWorkbookFilter)`

SetFilter gets a reference to the given AnyOfmicrosoftGraphWorkbookFilter and assigns it to the Filter field.

### SetFilterExplicitNull

`func (o *MicrosoftGraphWorkbookTableColumn) SetFilterExplicitNull(b bool)`

SetFilterExplicitNull (un)sets Filter to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Filter value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


