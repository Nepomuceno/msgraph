# OnenoteEntitySchemaObjectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### GetCreatedDateTime

`func (o *OnenoteEntitySchemaObjectModel) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *OnenoteEntitySchemaObjectModel) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *OnenoteEntitySchemaObjectModel) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *OnenoteEntitySchemaObjectModel) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *OnenoteEntitySchemaObjectModel) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


