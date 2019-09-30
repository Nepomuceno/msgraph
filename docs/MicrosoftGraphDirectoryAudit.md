# MicrosoftGraphDirectoryAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CorrelationId** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**AnyOfmicrosoftGraphOperationResult**](anyOf&lt;microsoft.graph.operationResult&gt;.md) |  | [optional] 
**ResultReason** | Pointer to **string** |  | [optional] 
**ActivityDisplayName** | Pointer to **string** |  | [optional] 
**ActivityDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LoggedByService** | Pointer to **string** |  | [optional] 
**OperationType** | Pointer to **string** |  | [optional] 
**InitiatedBy** | Pointer to [**MicrosoftGraphAuditActivityInitiator**](microsoft.graph.auditActivityInitiator.md) |  | [optional] 
**TargetResources** | Pointer to [**[]AnyOfmicrosoftGraphTargetResource**](anyOf&lt;microsoft.graph.targetResource&gt;.md) |  | [optional] 
**AdditionalDetails** | Pointer to [**[]AnyOfmicrosoftGraphKeyValue**](anyOf&lt;microsoft.graph.keyValue&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDirectoryAudit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDirectoryAudit) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDirectoryAudit) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDirectoryAudit) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCategory

`func (o *MicrosoftGraphDirectoryAudit) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MicrosoftGraphDirectoryAudit) GetCategoryOk() (string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategory

`func (o *MicrosoftGraphDirectoryAudit) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategory

`func (o *MicrosoftGraphDirectoryAudit) SetCategory(v string)`

SetCategory gets a reference to the given string and assigns it to the Category field.

### GetCorrelationId

`func (o *MicrosoftGraphDirectoryAudit) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MicrosoftGraphDirectoryAudit) GetCorrelationIdOk() (string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCorrelationId

`func (o *MicrosoftGraphDirectoryAudit) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationId

`func (o *MicrosoftGraphDirectoryAudit) SetCorrelationId(v string)`

SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.

### SetCorrelationIdExplicitNull

`func (o *MicrosoftGraphDirectoryAudit) SetCorrelationIdExplicitNull(b bool)`

SetCorrelationIdExplicitNull (un)sets CorrelationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CorrelationId value is set to nil even if false is passed
### GetResult

`func (o *MicrosoftGraphDirectoryAudit) GetResult() AnyOfmicrosoftGraphOperationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MicrosoftGraphDirectoryAudit) GetResultOk() (AnyOfmicrosoftGraphOperationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResult

`func (o *MicrosoftGraphDirectoryAudit) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResult

`func (o *MicrosoftGraphDirectoryAudit) SetResult(v AnyOfmicrosoftGraphOperationResult)`

SetResult gets a reference to the given AnyOfmicrosoftGraphOperationResult and assigns it to the Result field.

### SetResultExplicitNull

`func (o *MicrosoftGraphDirectoryAudit) SetResultExplicitNull(b bool)`

SetResultExplicitNull (un)sets Result to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Result value is set to nil even if false is passed
### GetResultReason

`func (o *MicrosoftGraphDirectoryAudit) GetResultReason() string`

GetResultReason returns the ResultReason field if non-nil, zero value otherwise.

### GetResultReasonOk

`func (o *MicrosoftGraphDirectoryAudit) GetResultReasonOk() (string, bool)`

GetResultReasonOk returns a tuple with the ResultReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResultReason

`func (o *MicrosoftGraphDirectoryAudit) HasResultReason() bool`

HasResultReason returns a boolean if a field has been set.

### SetResultReason

`func (o *MicrosoftGraphDirectoryAudit) SetResultReason(v string)`

SetResultReason gets a reference to the given string and assigns it to the ResultReason field.

### SetResultReasonExplicitNull

`func (o *MicrosoftGraphDirectoryAudit) SetResultReasonExplicitNull(b bool)`

SetResultReasonExplicitNull (un)sets ResultReason to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResultReason value is set to nil even if false is passed
### GetActivityDisplayName

`func (o *MicrosoftGraphDirectoryAudit) GetActivityDisplayName() string`

GetActivityDisplayName returns the ActivityDisplayName field if non-nil, zero value otherwise.

### GetActivityDisplayNameOk

`func (o *MicrosoftGraphDirectoryAudit) GetActivityDisplayNameOk() (string, bool)`

GetActivityDisplayNameOk returns a tuple with the ActivityDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivityDisplayName

`func (o *MicrosoftGraphDirectoryAudit) HasActivityDisplayName() bool`

HasActivityDisplayName returns a boolean if a field has been set.

### SetActivityDisplayName

`func (o *MicrosoftGraphDirectoryAudit) SetActivityDisplayName(v string)`

SetActivityDisplayName gets a reference to the given string and assigns it to the ActivityDisplayName field.

### GetActivityDateTime

`func (o *MicrosoftGraphDirectoryAudit) GetActivityDateTime() time.Time`

GetActivityDateTime returns the ActivityDateTime field if non-nil, zero value otherwise.

### GetActivityDateTimeOk

`func (o *MicrosoftGraphDirectoryAudit) GetActivityDateTimeOk() (time.Time, bool)`

GetActivityDateTimeOk returns a tuple with the ActivityDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivityDateTime

`func (o *MicrosoftGraphDirectoryAudit) HasActivityDateTime() bool`

HasActivityDateTime returns a boolean if a field has been set.

### SetActivityDateTime

`func (o *MicrosoftGraphDirectoryAudit) SetActivityDateTime(v time.Time)`

SetActivityDateTime gets a reference to the given time.Time and assigns it to the ActivityDateTime field.

### GetLoggedByService

`func (o *MicrosoftGraphDirectoryAudit) GetLoggedByService() string`

GetLoggedByService returns the LoggedByService field if non-nil, zero value otherwise.

### GetLoggedByServiceOk

`func (o *MicrosoftGraphDirectoryAudit) GetLoggedByServiceOk() (string, bool)`

GetLoggedByServiceOk returns a tuple with the LoggedByService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoggedByService

`func (o *MicrosoftGraphDirectoryAudit) HasLoggedByService() bool`

HasLoggedByService returns a boolean if a field has been set.

### SetLoggedByService

`func (o *MicrosoftGraphDirectoryAudit) SetLoggedByService(v string)`

SetLoggedByService gets a reference to the given string and assigns it to the LoggedByService field.

### SetLoggedByServiceExplicitNull

`func (o *MicrosoftGraphDirectoryAudit) SetLoggedByServiceExplicitNull(b bool)`

SetLoggedByServiceExplicitNull (un)sets LoggedByService to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LoggedByService value is set to nil even if false is passed
### GetOperationType

`func (o *MicrosoftGraphDirectoryAudit) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *MicrosoftGraphDirectoryAudit) GetOperationTypeOk() (string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperationType

`func (o *MicrosoftGraphDirectoryAudit) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationType

`func (o *MicrosoftGraphDirectoryAudit) SetOperationType(v string)`

SetOperationType gets a reference to the given string and assigns it to the OperationType field.

### SetOperationTypeExplicitNull

`func (o *MicrosoftGraphDirectoryAudit) SetOperationTypeExplicitNull(b bool)`

SetOperationTypeExplicitNull (un)sets OperationType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperationType value is set to nil even if false is passed
### GetInitiatedBy

`func (o *MicrosoftGraphDirectoryAudit) GetInitiatedBy() MicrosoftGraphAuditActivityInitiator`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *MicrosoftGraphDirectoryAudit) GetInitiatedByOk() (MicrosoftGraphAuditActivityInitiator, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInitiatedBy

`func (o *MicrosoftGraphDirectoryAudit) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### SetInitiatedBy

`func (o *MicrosoftGraphDirectoryAudit) SetInitiatedBy(v MicrosoftGraphAuditActivityInitiator)`

SetInitiatedBy gets a reference to the given MicrosoftGraphAuditActivityInitiator and assigns it to the InitiatedBy field.

### GetTargetResources

`func (o *MicrosoftGraphDirectoryAudit) GetTargetResources() []AnyOfmicrosoftGraphTargetResource`

GetTargetResources returns the TargetResources field if non-nil, zero value otherwise.

### GetTargetResourcesOk

`func (o *MicrosoftGraphDirectoryAudit) GetTargetResourcesOk() ([]AnyOfmicrosoftGraphTargetResource, bool)`

GetTargetResourcesOk returns a tuple with the TargetResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetResources

`func (o *MicrosoftGraphDirectoryAudit) HasTargetResources() bool`

HasTargetResources returns a boolean if a field has been set.

### SetTargetResources

`func (o *MicrosoftGraphDirectoryAudit) SetTargetResources(v []AnyOfmicrosoftGraphTargetResource)`

SetTargetResources gets a reference to the given []AnyOfmicrosoftGraphTargetResource and assigns it to the TargetResources field.

### GetAdditionalDetails

`func (o *MicrosoftGraphDirectoryAudit) GetAdditionalDetails() []AnyOfmicrosoftGraphKeyValue`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *MicrosoftGraphDirectoryAudit) GetAdditionalDetailsOk() ([]AnyOfmicrosoftGraphKeyValue, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalDetails

`func (o *MicrosoftGraphDirectoryAudit) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### SetAdditionalDetails

`func (o *MicrosoftGraphDirectoryAudit) SetAdditionalDetails(v []AnyOfmicrosoftGraphKeyValue)`

SetAdditionalDetails gets a reference to the given []AnyOfmicrosoftGraphKeyValue and assigns it to the AdditionalDetails field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


