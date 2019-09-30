# TeamsAsyncOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | Pointer to [**AnyOfmicrosoftGraphTeamsAsyncOperationType**](anyOf&lt;microsoft.graph.teamsAsyncOperationType&gt;.md) |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphTeamsAsyncOperationStatus**](anyOf&lt;microsoft.graph.teamsAsyncOperationStatus&gt;.md) |  | [optional] 
**LastActionDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**AttemptsCount** | Pointer to **int32** |  | [optional] 
**TargetResourceId** | Pointer to **string** |  | [optional] 
**TargetResourceLocation** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**AnyOfmicrosoftGraphOperationError**](anyOf&lt;microsoft.graph.operationError&gt;.md) |  | [optional] 

## Methods

### GetOperationType

`func (o *TeamsAsyncOperation) GetOperationType() AnyOfmicrosoftGraphTeamsAsyncOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *TeamsAsyncOperation) GetOperationTypeOk() (AnyOfmicrosoftGraphTeamsAsyncOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperationType

`func (o *TeamsAsyncOperation) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationType

`func (o *TeamsAsyncOperation) SetOperationType(v AnyOfmicrosoftGraphTeamsAsyncOperationType)`

SetOperationType gets a reference to the given AnyOfmicrosoftGraphTeamsAsyncOperationType and assigns it to the OperationType field.

### GetCreatedDateTime

`func (o *TeamsAsyncOperation) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *TeamsAsyncOperation) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *TeamsAsyncOperation) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *TeamsAsyncOperation) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetStatus

`func (o *TeamsAsyncOperation) GetStatus() AnyOfmicrosoftGraphTeamsAsyncOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TeamsAsyncOperation) GetStatusOk() (AnyOfmicrosoftGraphTeamsAsyncOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *TeamsAsyncOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *TeamsAsyncOperation) SetStatus(v AnyOfmicrosoftGraphTeamsAsyncOperationStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphTeamsAsyncOperationStatus and assigns it to the Status field.

### GetLastActionDateTime

`func (o *TeamsAsyncOperation) GetLastActionDateTime() time.Time`

GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.

### GetLastActionDateTimeOk

`func (o *TeamsAsyncOperation) GetLastActionDateTimeOk() (time.Time, bool)`

GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActionDateTime

`func (o *TeamsAsyncOperation) HasLastActionDateTime() bool`

HasLastActionDateTime returns a boolean if a field has been set.

### SetLastActionDateTime

`func (o *TeamsAsyncOperation) SetLastActionDateTime(v time.Time)`

SetLastActionDateTime gets a reference to the given time.Time and assigns it to the LastActionDateTime field.

### GetAttemptsCount

`func (o *TeamsAsyncOperation) GetAttemptsCount() int32`

GetAttemptsCount returns the AttemptsCount field if non-nil, zero value otherwise.

### GetAttemptsCountOk

`func (o *TeamsAsyncOperation) GetAttemptsCountOk() (int32, bool)`

GetAttemptsCountOk returns a tuple with the AttemptsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttemptsCount

`func (o *TeamsAsyncOperation) HasAttemptsCount() bool`

HasAttemptsCount returns a boolean if a field has been set.

### SetAttemptsCount

`func (o *TeamsAsyncOperation) SetAttemptsCount(v int32)`

SetAttemptsCount gets a reference to the given int32 and assigns it to the AttemptsCount field.

### GetTargetResourceId

`func (o *TeamsAsyncOperation) GetTargetResourceId() string`

GetTargetResourceId returns the TargetResourceId field if non-nil, zero value otherwise.

### GetTargetResourceIdOk

`func (o *TeamsAsyncOperation) GetTargetResourceIdOk() (string, bool)`

GetTargetResourceIdOk returns a tuple with the TargetResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetResourceId

`func (o *TeamsAsyncOperation) HasTargetResourceId() bool`

HasTargetResourceId returns a boolean if a field has been set.

### SetTargetResourceId

`func (o *TeamsAsyncOperation) SetTargetResourceId(v string)`

SetTargetResourceId gets a reference to the given string and assigns it to the TargetResourceId field.

### SetTargetResourceIdExplicitNull

`func (o *TeamsAsyncOperation) SetTargetResourceIdExplicitNull(b bool)`

SetTargetResourceIdExplicitNull (un)sets TargetResourceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TargetResourceId value is set to nil even if false is passed
### GetTargetResourceLocation

`func (o *TeamsAsyncOperation) GetTargetResourceLocation() string`

GetTargetResourceLocation returns the TargetResourceLocation field if non-nil, zero value otherwise.

### GetTargetResourceLocationOk

`func (o *TeamsAsyncOperation) GetTargetResourceLocationOk() (string, bool)`

GetTargetResourceLocationOk returns a tuple with the TargetResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetResourceLocation

`func (o *TeamsAsyncOperation) HasTargetResourceLocation() bool`

HasTargetResourceLocation returns a boolean if a field has been set.

### SetTargetResourceLocation

`func (o *TeamsAsyncOperation) SetTargetResourceLocation(v string)`

SetTargetResourceLocation gets a reference to the given string and assigns it to the TargetResourceLocation field.

### SetTargetResourceLocationExplicitNull

`func (o *TeamsAsyncOperation) SetTargetResourceLocationExplicitNull(b bool)`

SetTargetResourceLocationExplicitNull (un)sets TargetResourceLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TargetResourceLocation value is set to nil even if false is passed
### GetError

`func (o *TeamsAsyncOperation) GetError() AnyOfmicrosoftGraphOperationError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TeamsAsyncOperation) GetErrorOk() (AnyOfmicrosoftGraphOperationError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *TeamsAsyncOperation) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *TeamsAsyncOperation) SetError(v AnyOfmicrosoftGraphOperationError)`

SetError gets a reference to the given AnyOfmicrosoftGraphOperationError and assigns it to the Error field.

### SetErrorExplicitNull

`func (o *TeamsAsyncOperation) SetErrorExplicitNull(b bool)`

SetErrorExplicitNull (un)sets Error to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Error value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


