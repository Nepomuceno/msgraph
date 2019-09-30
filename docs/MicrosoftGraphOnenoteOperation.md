# MicrosoftGraphOnenoteOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphOperationStatus**](anyOf&lt;microsoft.graph.operationStatus&gt;.md) |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastActionDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ResourceLocation** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**AnyOfmicrosoftGraphOnenoteOperationError**](anyOf&lt;microsoft.graph.onenoteOperationError&gt;.md) |  | [optional] 
**PercentComplete** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphOnenoteOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOnenoteOperation) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphOnenoteOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphOnenoteOperation) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetStatus

`func (o *MicrosoftGraphOnenoteOperation) GetStatus() AnyOfmicrosoftGraphOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphOnenoteOperation) GetStatusOk() (AnyOfmicrosoftGraphOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphOnenoteOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphOnenoteOperation) SetStatus(v AnyOfmicrosoftGraphOperationStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphOperationStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphOnenoteOperation) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphOnenoteOperation) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOnenoteOperation) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphOnenoteOperation) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOnenoteOperation) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphOnenoteOperation) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetLastActionDateTime

`func (o *MicrosoftGraphOnenoteOperation) GetLastActionDateTime() time.Time`

GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.

### GetLastActionDateTimeOk

`func (o *MicrosoftGraphOnenoteOperation) GetLastActionDateTimeOk() (time.Time, bool)`

GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastActionDateTime

`func (o *MicrosoftGraphOnenoteOperation) HasLastActionDateTime() bool`

HasLastActionDateTime returns a boolean if a field has been set.

### SetLastActionDateTime

`func (o *MicrosoftGraphOnenoteOperation) SetLastActionDateTime(v time.Time)`

SetLastActionDateTime gets a reference to the given time.Time and assigns it to the LastActionDateTime field.

### SetLastActionDateTimeExplicitNull

`func (o *MicrosoftGraphOnenoteOperation) SetLastActionDateTimeExplicitNull(b bool)`

SetLastActionDateTimeExplicitNull (un)sets LastActionDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastActionDateTime value is set to nil even if false is passed
### GetResourceLocation

`func (o *MicrosoftGraphOnenoteOperation) GetResourceLocation() string`

GetResourceLocation returns the ResourceLocation field if non-nil, zero value otherwise.

### GetResourceLocationOk

`func (o *MicrosoftGraphOnenoteOperation) GetResourceLocationOk() (string, bool)`

GetResourceLocationOk returns a tuple with the ResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceLocation

`func (o *MicrosoftGraphOnenoteOperation) HasResourceLocation() bool`

HasResourceLocation returns a boolean if a field has been set.

### SetResourceLocation

`func (o *MicrosoftGraphOnenoteOperation) SetResourceLocation(v string)`

SetResourceLocation gets a reference to the given string and assigns it to the ResourceLocation field.

### SetResourceLocationExplicitNull

`func (o *MicrosoftGraphOnenoteOperation) SetResourceLocationExplicitNull(b bool)`

SetResourceLocationExplicitNull (un)sets ResourceLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceLocation value is set to nil even if false is passed
### GetResourceId

`func (o *MicrosoftGraphOnenoteOperation) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *MicrosoftGraphOnenoteOperation) GetResourceIdOk() (string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceId

`func (o *MicrosoftGraphOnenoteOperation) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceId

`func (o *MicrosoftGraphOnenoteOperation) SetResourceId(v string)`

SetResourceId gets a reference to the given string and assigns it to the ResourceId field.

### SetResourceIdExplicitNull

`func (o *MicrosoftGraphOnenoteOperation) SetResourceIdExplicitNull(b bool)`

SetResourceIdExplicitNull (un)sets ResourceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceId value is set to nil even if false is passed
### GetError

`func (o *MicrosoftGraphOnenoteOperation) GetError() AnyOfmicrosoftGraphOnenoteOperationError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphOnenoteOperation) GetErrorOk() (AnyOfmicrosoftGraphOnenoteOperationError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *MicrosoftGraphOnenoteOperation) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *MicrosoftGraphOnenoteOperation) SetError(v AnyOfmicrosoftGraphOnenoteOperationError)`

SetError gets a reference to the given AnyOfmicrosoftGraphOnenoteOperationError and assigns it to the Error field.

### SetErrorExplicitNull

`func (o *MicrosoftGraphOnenoteOperation) SetErrorExplicitNull(b bool)`

SetErrorExplicitNull (un)sets Error to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Error value is set to nil even if false is passed
### GetPercentComplete

`func (o *MicrosoftGraphOnenoteOperation) GetPercentComplete() string`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *MicrosoftGraphOnenoteOperation) GetPercentCompleteOk() (string, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPercentComplete

`func (o *MicrosoftGraphOnenoteOperation) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentComplete

`func (o *MicrosoftGraphOnenoteOperation) SetPercentComplete(v string)`

SetPercentComplete gets a reference to the given string and assigns it to the PercentComplete field.

### SetPercentCompleteExplicitNull

`func (o *MicrosoftGraphOnenoteOperation) SetPercentCompleteExplicitNull(b bool)`

SetPercentCompleteExplicitNull (un)sets PercentComplete to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PercentComplete value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


