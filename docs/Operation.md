# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AnyOfmicrosoftGraphOperationStatus**](anyOf&lt;microsoft.graph.operationStatus&gt;.md) |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastActionDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### GetStatus

`func (o *Operation) GetStatus() AnyOfmicrosoftGraphOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Operation) GetStatusOk() (AnyOfmicrosoftGraphOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *Operation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *Operation) SetStatus(v AnyOfmicrosoftGraphOperationStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphOperationStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *Operation) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *Operation) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Operation) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *Operation) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *Operation) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *Operation) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetLastActionDateTime

`func (o *Operation) GetLastActionDateTime() time.Time`

GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.

### GetLastActionDateTimeOk

`func (o *Operation) GetLastActionDateTimeOk() (time.Time, bool)`

GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActionDateTime

`func (o *Operation) HasLastActionDateTime() bool`

HasLastActionDateTime returns a boolean if a field has been set.

### SetLastActionDateTime

`func (o *Operation) SetLastActionDateTime(v time.Time)`

SetLastActionDateTime gets a reference to the given time.Time and assigns it to the LastActionDateTime field.

### SetLastActionDateTimeExplicitNull

`func (o *Operation) SetLastActionDateTimeExplicitNull(b bool)`

SetLastActionDateTimeExplicitNull (un)sets LastActionDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastActionDateTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


