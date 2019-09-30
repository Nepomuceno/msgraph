# DataPolicyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphDataPolicyOperationStatus**](anyOf&lt;microsoft.graph.dataPolicyOperationStatus&gt;.md) |  | [optional] 
**StorageLocation** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**SubmittedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Progress** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 

## Methods

### GetCompletedDateTime

`func (o *DataPolicyOperation) GetCompletedDateTime() time.Time`

GetCompletedDateTime returns the CompletedDateTime field if non-nil, zero value otherwise.

### GetCompletedDateTimeOk

`func (o *DataPolicyOperation) GetCompletedDateTimeOk() (time.Time, bool)`

GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompletedDateTime

`func (o *DataPolicyOperation) HasCompletedDateTime() bool`

HasCompletedDateTime returns a boolean if a field has been set.

### SetCompletedDateTime

`func (o *DataPolicyOperation) SetCompletedDateTime(v time.Time)`

SetCompletedDateTime gets a reference to the given time.Time and assigns it to the CompletedDateTime field.

### SetCompletedDateTimeExplicitNull

`func (o *DataPolicyOperation) SetCompletedDateTimeExplicitNull(b bool)`

SetCompletedDateTimeExplicitNull (un)sets CompletedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CompletedDateTime value is set to nil even if false is passed
### GetStatus

`func (o *DataPolicyOperation) GetStatus() AnyOfmicrosoftGraphDataPolicyOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataPolicyOperation) GetStatusOk() (AnyOfmicrosoftGraphDataPolicyOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *DataPolicyOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *DataPolicyOperation) SetStatus(v AnyOfmicrosoftGraphDataPolicyOperationStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphDataPolicyOperationStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *DataPolicyOperation) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetStorageLocation

`func (o *DataPolicyOperation) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *DataPolicyOperation) GetStorageLocationOk() (string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageLocation

`func (o *DataPolicyOperation) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.

### SetStorageLocation

`func (o *DataPolicyOperation) SetStorageLocation(v string)`

SetStorageLocation gets a reference to the given string and assigns it to the StorageLocation field.

### SetStorageLocationExplicitNull

`func (o *DataPolicyOperation) SetStorageLocationExplicitNull(b bool)`

SetStorageLocationExplicitNull (un)sets StorageLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StorageLocation value is set to nil even if false is passed
### GetUserId

`func (o *DataPolicyOperation) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DataPolicyOperation) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *DataPolicyOperation) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *DataPolicyOperation) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### GetSubmittedDateTime

`func (o *DataPolicyOperation) GetSubmittedDateTime() time.Time`

GetSubmittedDateTime returns the SubmittedDateTime field if non-nil, zero value otherwise.

### GetSubmittedDateTimeOk

`func (o *DataPolicyOperation) GetSubmittedDateTimeOk() (time.Time, bool)`

GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubmittedDateTime

`func (o *DataPolicyOperation) HasSubmittedDateTime() bool`

HasSubmittedDateTime returns a boolean if a field has been set.

### SetSubmittedDateTime

`func (o *DataPolicyOperation) SetSubmittedDateTime(v time.Time)`

SetSubmittedDateTime gets a reference to the given time.Time and assigns it to the SubmittedDateTime field.

### GetProgress

`func (o *DataPolicyOperation) GetProgress() AnyOfnumberstringstring`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DataPolicyOperation) GetProgressOk() (AnyOfnumberstringstring, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgress

`func (o *DataPolicyOperation) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgress

`func (o *DataPolicyOperation) SetProgress(v AnyOfnumberstringstring)`

SetProgress gets a reference to the given AnyOfnumberstringstring and assigns it to the Progress field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


