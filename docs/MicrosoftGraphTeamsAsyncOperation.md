# MicrosoftGraphTeamsAsyncOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OperationType** | Pointer to [**AnyOfmicrosoftGraphTeamsAsyncOperationType**](anyOf&lt;microsoft.graph.teamsAsyncOperationType&gt;.md) |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphTeamsAsyncOperationStatus**](anyOf&lt;microsoft.graph.teamsAsyncOperationStatus&gt;.md) |  | [optional] 
**LastActionDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**AttemptsCount** | Pointer to **int32** |  | [optional] 
**TargetResourceId** | Pointer to **string** |  | [optional] 
**TargetResourceLocation** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**AnyOfmicrosoftGraphOperationError**](anyOf&lt;microsoft.graph.operationError&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphTeamsAsyncOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTeamsAsyncOperation) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphTeamsAsyncOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphTeamsAsyncOperation) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetOperationType

`func (o *MicrosoftGraphTeamsAsyncOperation) GetOperationType() AnyOfmicrosoftGraphTeamsAsyncOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *MicrosoftGraphTeamsAsyncOperation) GetOperationTypeOk() (AnyOfmicrosoftGraphTeamsAsyncOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperationType

`func (o *MicrosoftGraphTeamsAsyncOperation) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationType

`func (o *MicrosoftGraphTeamsAsyncOperation) SetOperationType(v AnyOfmicrosoftGraphTeamsAsyncOperationType)`

SetOperationType gets a reference to the given AnyOfmicrosoftGraphTeamsAsyncOperationType and assigns it to the OperationType field.

### GetCreatedDateTime

`func (o *MicrosoftGraphTeamsAsyncOperation) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphTeamsAsyncOperation) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphTeamsAsyncOperation) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphTeamsAsyncOperation) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetStatus

`func (o *MicrosoftGraphTeamsAsyncOperation) GetStatus() AnyOfmicrosoftGraphTeamsAsyncOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphTeamsAsyncOperation) GetStatusOk() (AnyOfmicrosoftGraphTeamsAsyncOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphTeamsAsyncOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphTeamsAsyncOperation) SetStatus(v AnyOfmicrosoftGraphTeamsAsyncOperationStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphTeamsAsyncOperationStatus and assigns it to the Status field.

### GetLastActionDateTime

`func (o *MicrosoftGraphTeamsAsyncOperation) GetLastActionDateTime() time.Time`

GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.

### GetLastActionDateTimeOk

`func (o *MicrosoftGraphTeamsAsyncOperation) GetLastActionDateTimeOk() (time.Time, bool)`

GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActionDateTime

`func (o *MicrosoftGraphTeamsAsyncOperation) HasLastActionDateTime() bool`

HasLastActionDateTime returns a boolean if a field has been set.

### SetLastActionDateTime

`func (o *MicrosoftGraphTeamsAsyncOperation) SetLastActionDateTime(v time.Time)`

SetLastActionDateTime gets a reference to the given time.Time and assigns it to the LastActionDateTime field.

### GetAttemptsCount

`func (o *MicrosoftGraphTeamsAsyncOperation) GetAttemptsCount() int32`

GetAttemptsCount returns the AttemptsCount field if non-nil, zero value otherwise.

### GetAttemptsCountOk

`func (o *MicrosoftGraphTeamsAsyncOperation) GetAttemptsCountOk() (int32, bool)`

GetAttemptsCountOk returns a tuple with the AttemptsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttemptsCount

`func (o *MicrosoftGraphTeamsAsyncOperation) HasAttemptsCount() bool`

HasAttemptsCount returns a boolean if a field has been set.

### SetAttemptsCount

`func (o *MicrosoftGraphTeamsAsyncOperation) SetAttemptsCount(v int32)`

SetAttemptsCount gets a reference to the given int32 and assigns it to the AttemptsCount field.

### GetTargetResourceId

`func (o *MicrosoftGraphTeamsAsyncOperation) GetTargetResourceId() string`

GetTargetResourceId returns the TargetResourceId field if non-nil, zero value otherwise.

### GetTargetResourceIdOk

`func (o *MicrosoftGraphTeamsAsyncOperation) GetTargetResourceIdOk() (string, bool)`

GetTargetResourceIdOk returns a tuple with the TargetResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetResourceId

`func (o *MicrosoftGraphTeamsAsyncOperation) HasTargetResourceId() bool`

HasTargetResourceId returns a boolean if a field has been set.

### SetTargetResourceId

`func (o *MicrosoftGraphTeamsAsyncOperation) SetTargetResourceId(v string)`

SetTargetResourceId gets a reference to the given string and assigns it to the TargetResourceId field.

### SetTargetResourceIdExplicitNull

`func (o *MicrosoftGraphTeamsAsyncOperation) SetTargetResourceIdExplicitNull(b bool)`

SetTargetResourceIdExplicitNull (un)sets TargetResourceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TargetResourceId value is set to nil even if false is passed
### GetTargetResourceLocation

`func (o *MicrosoftGraphTeamsAsyncOperation) GetTargetResourceLocation() string`

GetTargetResourceLocation returns the TargetResourceLocation field if non-nil, zero value otherwise.

### GetTargetResourceLocationOk

`func (o *MicrosoftGraphTeamsAsyncOperation) GetTargetResourceLocationOk() (string, bool)`

GetTargetResourceLocationOk returns a tuple with the TargetResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetResourceLocation

`func (o *MicrosoftGraphTeamsAsyncOperation) HasTargetResourceLocation() bool`

HasTargetResourceLocation returns a boolean if a field has been set.

### SetTargetResourceLocation

`func (o *MicrosoftGraphTeamsAsyncOperation) SetTargetResourceLocation(v string)`

SetTargetResourceLocation gets a reference to the given string and assigns it to the TargetResourceLocation field.

### SetTargetResourceLocationExplicitNull

`func (o *MicrosoftGraphTeamsAsyncOperation) SetTargetResourceLocationExplicitNull(b bool)`

SetTargetResourceLocationExplicitNull (un)sets TargetResourceLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TargetResourceLocation value is set to nil even if false is passed
### GetError

`func (o *MicrosoftGraphTeamsAsyncOperation) GetError() AnyOfmicrosoftGraphOperationError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphTeamsAsyncOperation) GetErrorOk() (AnyOfmicrosoftGraphOperationError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *MicrosoftGraphTeamsAsyncOperation) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *MicrosoftGraphTeamsAsyncOperation) SetError(v AnyOfmicrosoftGraphOperationError)`

SetError gets a reference to the given AnyOfmicrosoftGraphOperationError and assigns it to the Error field.

### SetErrorExplicitNull

`func (o *MicrosoftGraphTeamsAsyncOperation) SetErrorExplicitNull(b bool)`

SetErrorExplicitNull (un)sets Error to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Error value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


