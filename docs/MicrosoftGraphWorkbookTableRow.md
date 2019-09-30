# MicrosoftGraphWorkbookTableRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Values** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWorkbookTableRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookTableRow) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookTableRow) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookTableRow) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetIndex

`func (o *MicrosoftGraphWorkbookTableRow) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MicrosoftGraphWorkbookTableRow) GetIndexOk() (int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndex

`func (o *MicrosoftGraphWorkbookTableRow) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndex

`func (o *MicrosoftGraphWorkbookTableRow) SetIndex(v int32)`

SetIndex gets a reference to the given int32 and assigns it to the Index field.

### GetValues

`func (o *MicrosoftGraphWorkbookTableRow) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MicrosoftGraphWorkbookTableRow) GetValuesOk() (AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *MicrosoftGraphWorkbookTableRow) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *MicrosoftGraphWorkbookTableRow) SetValues(v AnyOfobject)`

SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.

### SetValuesExplicitNull

`func (o *MicrosoftGraphWorkbookTableRow) SetValuesExplicitNull(b bool)`

SetValuesExplicitNull (un)sets Values to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Values value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


