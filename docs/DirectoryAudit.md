# DirectoryAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetCategory

`func (o *DirectoryAudit) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DirectoryAudit) GetCategoryOk() (string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategory

`func (o *DirectoryAudit) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategory

`func (o *DirectoryAudit) SetCategory(v string)`

SetCategory gets a reference to the given string and assigns it to the Category field.

### GetCorrelationId

`func (o *DirectoryAudit) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *DirectoryAudit) GetCorrelationIdOk() (string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCorrelationId

`func (o *DirectoryAudit) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationId

`func (o *DirectoryAudit) SetCorrelationId(v string)`

SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.

### SetCorrelationIdExplicitNull

`func (o *DirectoryAudit) SetCorrelationIdExplicitNull(b bool)`

SetCorrelationIdExplicitNull (un)sets CorrelationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CorrelationId value is set to nil even if false is passed
### GetResult

`func (o *DirectoryAudit) GetResult() AnyOfmicrosoftGraphOperationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DirectoryAudit) GetResultOk() (AnyOfmicrosoftGraphOperationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResult

`func (o *DirectoryAudit) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResult

`func (o *DirectoryAudit) SetResult(v AnyOfmicrosoftGraphOperationResult)`

SetResult gets a reference to the given AnyOfmicrosoftGraphOperationResult and assigns it to the Result field.

### SetResultExplicitNull

`func (o *DirectoryAudit) SetResultExplicitNull(b bool)`

SetResultExplicitNull (un)sets Result to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Result value is set to nil even if false is passed
### GetResultReason

`func (o *DirectoryAudit) GetResultReason() string`

GetResultReason returns the ResultReason field if non-nil, zero value otherwise.

### GetResultReasonOk

`func (o *DirectoryAudit) GetResultReasonOk() (string, bool)`

GetResultReasonOk returns a tuple with the ResultReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResultReason

`func (o *DirectoryAudit) HasResultReason() bool`

HasResultReason returns a boolean if a field has been set.

### SetResultReason

`func (o *DirectoryAudit) SetResultReason(v string)`

SetResultReason gets a reference to the given string and assigns it to the ResultReason field.

### SetResultReasonExplicitNull

`func (o *DirectoryAudit) SetResultReasonExplicitNull(b bool)`

SetResultReasonExplicitNull (un)sets ResultReason to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResultReason value is set to nil even if false is passed
### GetActivityDisplayName

`func (o *DirectoryAudit) GetActivityDisplayName() string`

GetActivityDisplayName returns the ActivityDisplayName field if non-nil, zero value otherwise.

### GetActivityDisplayNameOk

`func (o *DirectoryAudit) GetActivityDisplayNameOk() (string, bool)`

GetActivityDisplayNameOk returns a tuple with the ActivityDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivityDisplayName

`func (o *DirectoryAudit) HasActivityDisplayName() bool`

HasActivityDisplayName returns a boolean if a field has been set.

### SetActivityDisplayName

`func (o *DirectoryAudit) SetActivityDisplayName(v string)`

SetActivityDisplayName gets a reference to the given string and assigns it to the ActivityDisplayName field.

### GetActivityDateTime

`func (o *DirectoryAudit) GetActivityDateTime() time.Time`

GetActivityDateTime returns the ActivityDateTime field if non-nil, zero value otherwise.

### GetActivityDateTimeOk

`func (o *DirectoryAudit) GetActivityDateTimeOk() (time.Time, bool)`

GetActivityDateTimeOk returns a tuple with the ActivityDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivityDateTime

`func (o *DirectoryAudit) HasActivityDateTime() bool`

HasActivityDateTime returns a boolean if a field has been set.

### SetActivityDateTime

`func (o *DirectoryAudit) SetActivityDateTime(v time.Time)`

SetActivityDateTime gets a reference to the given time.Time and assigns it to the ActivityDateTime field.

### GetLoggedByService

`func (o *DirectoryAudit) GetLoggedByService() string`

GetLoggedByService returns the LoggedByService field if non-nil, zero value otherwise.

### GetLoggedByServiceOk

`func (o *DirectoryAudit) GetLoggedByServiceOk() (string, bool)`

GetLoggedByServiceOk returns a tuple with the LoggedByService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoggedByService

`func (o *DirectoryAudit) HasLoggedByService() bool`

HasLoggedByService returns a boolean if a field has been set.

### SetLoggedByService

`func (o *DirectoryAudit) SetLoggedByService(v string)`

SetLoggedByService gets a reference to the given string and assigns it to the LoggedByService field.

### SetLoggedByServiceExplicitNull

`func (o *DirectoryAudit) SetLoggedByServiceExplicitNull(b bool)`

SetLoggedByServiceExplicitNull (un)sets LoggedByService to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LoggedByService value is set to nil even if false is passed
### GetOperationType

`func (o *DirectoryAudit) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *DirectoryAudit) GetOperationTypeOk() (string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperationType

`func (o *DirectoryAudit) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationType

`func (o *DirectoryAudit) SetOperationType(v string)`

SetOperationType gets a reference to the given string and assigns it to the OperationType field.

### SetOperationTypeExplicitNull

`func (o *DirectoryAudit) SetOperationTypeExplicitNull(b bool)`

SetOperationTypeExplicitNull (un)sets OperationType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperationType value is set to nil even if false is passed
### GetInitiatedBy

`func (o *DirectoryAudit) GetInitiatedBy() MicrosoftGraphAuditActivityInitiator`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *DirectoryAudit) GetInitiatedByOk() (MicrosoftGraphAuditActivityInitiator, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInitiatedBy

`func (o *DirectoryAudit) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### SetInitiatedBy

`func (o *DirectoryAudit) SetInitiatedBy(v MicrosoftGraphAuditActivityInitiator)`

SetInitiatedBy gets a reference to the given MicrosoftGraphAuditActivityInitiator and assigns it to the InitiatedBy field.

### GetTargetResources

`func (o *DirectoryAudit) GetTargetResources() []AnyOfmicrosoftGraphTargetResource`

GetTargetResources returns the TargetResources field if non-nil, zero value otherwise.

### GetTargetResourcesOk

`func (o *DirectoryAudit) GetTargetResourcesOk() ([]AnyOfmicrosoftGraphTargetResource, bool)`

GetTargetResourcesOk returns a tuple with the TargetResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetResources

`func (o *DirectoryAudit) HasTargetResources() bool`

HasTargetResources returns a boolean if a field has been set.

### SetTargetResources

`func (o *DirectoryAudit) SetTargetResources(v []AnyOfmicrosoftGraphTargetResource)`

SetTargetResources gets a reference to the given []AnyOfmicrosoftGraphTargetResource and assigns it to the TargetResources field.

### GetAdditionalDetails

`func (o *DirectoryAudit) GetAdditionalDetails() []AnyOfmicrosoftGraphKeyValue`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *DirectoryAudit) GetAdditionalDetailsOk() ([]AnyOfmicrosoftGraphKeyValue, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdditionalDetails

`func (o *DirectoryAudit) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### SetAdditionalDetails

`func (o *DirectoryAudit) SetAdditionalDetails(v []AnyOfmicrosoftGraphKeyValue)`

SetAdditionalDetails gets a reference to the given []AnyOfmicrosoftGraphKeyValue and assigns it to the AdditionalDetails field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


