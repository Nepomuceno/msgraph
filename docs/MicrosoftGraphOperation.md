# MicrosoftGraphOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphOperationStatus**](anyOf&lt;microsoft.graph.operationStatus&gt;.md) |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastActionDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOperation) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphOperation) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetStatus

`func (o *MicrosoftGraphOperation) GetStatus() AnyOfmicrosoftGraphOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphOperation) GetStatusOk() (AnyOfmicrosoftGraphOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphOperation) SetStatus(v AnyOfmicrosoftGraphOperationStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphOperationStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphOperation) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphOperation) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOperation) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphOperation) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOperation) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphOperation) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetLastActionDateTime

`func (o *MicrosoftGraphOperation) GetLastActionDateTime() time.Time`

GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.

### GetLastActionDateTimeOk

`func (o *MicrosoftGraphOperation) GetLastActionDateTimeOk() (time.Time, bool)`

GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActionDateTime

`func (o *MicrosoftGraphOperation) HasLastActionDateTime() bool`

HasLastActionDateTime returns a boolean if a field has been set.

### SetLastActionDateTime

`func (o *MicrosoftGraphOperation) SetLastActionDateTime(v time.Time)`

SetLastActionDateTime gets a reference to the given time.Time and assigns it to the LastActionDateTime field.

### SetLastActionDateTimeExplicitNull

`func (o *MicrosoftGraphOperation) SetLastActionDateTimeExplicitNull(b bool)`

SetLastActionDateTimeExplicitNull (un)sets LastActionDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastActionDateTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


