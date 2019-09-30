# MicrosoftGraphAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ActivityGroupName** | Pointer to **string** |  | [optional] 
**AssignedTo** | Pointer to **string** |  | [optional] 
**AzureSubscriptionId** | Pointer to **string** |  | [optional] 
**AzureTenantId** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ClosedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**CloudAppStates** | Pointer to [**[]AnyOfmicrosoftGraphCloudAppSecurityState**](anyOf&lt;microsoft.graph.cloudAppSecurityState&gt;.md) |  | [optional] 
**Comments** | Pointer to **[]string** |  | [optional] 
**Confidence** | Pointer to **int32** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DetectionIds** | Pointer to **[]string** |  | [optional] 
**EventDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Feedback** | Pointer to [**AnyOfmicrosoftGraphAlertFeedback**](anyOf&lt;microsoft.graph.alertFeedback&gt;.md) |  | [optional] 
**FileStates** | Pointer to [**[]AnyOfmicrosoftGraphFileSecurityState**](anyOf&lt;microsoft.graph.fileSecurityState&gt;.md) |  | [optional] 
**HistoryStates** | Pointer to [**[]AnyOfmicrosoftGraphAlertHistoryState**](anyOf&lt;microsoft.graph.alertHistoryState&gt;.md) |  | [optional] 
**HostStates** | Pointer to [**[]AnyOfmicrosoftGraphHostSecurityState**](anyOf&lt;microsoft.graph.hostSecurityState&gt;.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**MalwareStates** | Pointer to [**[]AnyOfmicrosoftGraphMalwareState**](anyOf&lt;microsoft.graph.malwareState&gt;.md) |  | [optional] 
**NetworkConnections** | Pointer to [**[]AnyOfmicrosoftGraphNetworkConnection**](anyOf&lt;microsoft.graph.networkConnection&gt;.md) |  | [optional] 
**Processes** | Pointer to [**[]AnyOfmicrosoftGraphProcess**](anyOf&lt;microsoft.graph.process&gt;.md) |  | [optional] 
**RecommendedActions** | Pointer to **[]string** |  | [optional] 
**RegistryKeyStates** | Pointer to [**[]AnyOfmicrosoftGraphRegistryKeyState**](anyOf&lt;microsoft.graph.registryKeyState&gt;.md) |  | [optional] 
**Severity** | Pointer to [**AnyOfmicrosoftGraphAlertSeverity**](anyOf&lt;microsoft.graph.alertSeverity&gt;.md) |  | [optional] 
**SourceMaterials** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphAlertStatus**](anyOf&lt;microsoft.graph.alertStatus&gt;.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]AnyOfmicrosoftGraphAlertTrigger**](anyOf&lt;microsoft.graph.alertTrigger&gt;.md) |  | [optional] 
**UserStates** | Pointer to [**[]AnyOfmicrosoftGraphUserSecurityState**](anyOf&lt;microsoft.graph.userSecurityState&gt;.md) |  | [optional] 
**VendorInformation** | Pointer to [**AnyOfmicrosoftGraphSecurityVendorInformation**](anyOf&lt;microsoft.graph.securityVendorInformation&gt;.md) |  | [optional] 
**VulnerabilityStates** | Pointer to [**[]AnyOfmicrosoftGraphVulnerabilityState**](anyOf&lt;microsoft.graph.vulnerabilityState&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphAlert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAlert) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphAlert) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphAlert) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetActivityGroupName

`func (o *MicrosoftGraphAlert) GetActivityGroupName() string`

GetActivityGroupName returns the ActivityGroupName field if non-nil, zero value otherwise.

### GetActivityGroupNameOk

`func (o *MicrosoftGraphAlert) GetActivityGroupNameOk() (string, bool)`

GetActivityGroupNameOk returns a tuple with the ActivityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivityGroupName

`func (o *MicrosoftGraphAlert) HasActivityGroupName() bool`

HasActivityGroupName returns a boolean if a field has been set.

### SetActivityGroupName

`func (o *MicrosoftGraphAlert) SetActivityGroupName(v string)`

SetActivityGroupName gets a reference to the given string and assigns it to the ActivityGroupName field.

### SetActivityGroupNameExplicitNull

`func (o *MicrosoftGraphAlert) SetActivityGroupNameExplicitNull(b bool)`

SetActivityGroupNameExplicitNull (un)sets ActivityGroupName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActivityGroupName value is set to nil even if false is passed
### GetAssignedTo

`func (o *MicrosoftGraphAlert) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *MicrosoftGraphAlert) GetAssignedToOk() (string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedTo

`func (o *MicrosoftGraphAlert) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### SetAssignedTo

`func (o *MicrosoftGraphAlert) SetAssignedTo(v string)`

SetAssignedTo gets a reference to the given string and assigns it to the AssignedTo field.

### SetAssignedToExplicitNull

`func (o *MicrosoftGraphAlert) SetAssignedToExplicitNull(b bool)`

SetAssignedToExplicitNull (un)sets AssignedTo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AssignedTo value is set to nil even if false is passed
### GetAzureSubscriptionId

`func (o *MicrosoftGraphAlert) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *MicrosoftGraphAlert) GetAzureSubscriptionIdOk() (string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureSubscriptionId

`func (o *MicrosoftGraphAlert) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.

### SetAzureSubscriptionId

`func (o *MicrosoftGraphAlert) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId gets a reference to the given string and assigns it to the AzureSubscriptionId field.

### SetAzureSubscriptionIdExplicitNull

`func (o *MicrosoftGraphAlert) SetAzureSubscriptionIdExplicitNull(b bool)`

SetAzureSubscriptionIdExplicitNull (un)sets AzureSubscriptionId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AzureSubscriptionId value is set to nil even if false is passed
### GetAzureTenantId

`func (o *MicrosoftGraphAlert) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *MicrosoftGraphAlert) GetAzureTenantIdOk() (string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureTenantId

`func (o *MicrosoftGraphAlert) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### SetAzureTenantId

`func (o *MicrosoftGraphAlert) SetAzureTenantId(v string)`

SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.

### GetCategory

`func (o *MicrosoftGraphAlert) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MicrosoftGraphAlert) GetCategoryOk() (string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategory

`func (o *MicrosoftGraphAlert) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategory

`func (o *MicrosoftGraphAlert) SetCategory(v string)`

SetCategory gets a reference to the given string and assigns it to the Category field.

### SetCategoryExplicitNull

`func (o *MicrosoftGraphAlert) SetCategoryExplicitNull(b bool)`

SetCategoryExplicitNull (un)sets Category to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Category value is set to nil even if false is passed
### GetClosedDateTime

`func (o *MicrosoftGraphAlert) GetClosedDateTime() time.Time`

GetClosedDateTime returns the ClosedDateTime field if non-nil, zero value otherwise.

### GetClosedDateTimeOk

`func (o *MicrosoftGraphAlert) GetClosedDateTimeOk() (time.Time, bool)`

GetClosedDateTimeOk returns a tuple with the ClosedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClosedDateTime

`func (o *MicrosoftGraphAlert) HasClosedDateTime() bool`

HasClosedDateTime returns a boolean if a field has been set.

### SetClosedDateTime

`func (o *MicrosoftGraphAlert) SetClosedDateTime(v time.Time)`

SetClosedDateTime gets a reference to the given time.Time and assigns it to the ClosedDateTime field.

### SetClosedDateTimeExplicitNull

`func (o *MicrosoftGraphAlert) SetClosedDateTimeExplicitNull(b bool)`

SetClosedDateTimeExplicitNull (un)sets ClosedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ClosedDateTime value is set to nil even if false is passed
### GetCloudAppStates

`func (o *MicrosoftGraphAlert) GetCloudAppStates() []AnyOfmicrosoftGraphCloudAppSecurityState`

GetCloudAppStates returns the CloudAppStates field if non-nil, zero value otherwise.

### GetCloudAppStatesOk

`func (o *MicrosoftGraphAlert) GetCloudAppStatesOk() ([]AnyOfmicrosoftGraphCloudAppSecurityState, bool)`

GetCloudAppStatesOk returns a tuple with the CloudAppStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCloudAppStates

`func (o *MicrosoftGraphAlert) HasCloudAppStates() bool`

HasCloudAppStates returns a boolean if a field has been set.

### SetCloudAppStates

`func (o *MicrosoftGraphAlert) SetCloudAppStates(v []AnyOfmicrosoftGraphCloudAppSecurityState)`

SetCloudAppStates gets a reference to the given []AnyOfmicrosoftGraphCloudAppSecurityState and assigns it to the CloudAppStates field.

### GetComments

`func (o *MicrosoftGraphAlert) GetComments() []string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *MicrosoftGraphAlert) GetCommentsOk() ([]string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComments

`func (o *MicrosoftGraphAlert) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetComments

`func (o *MicrosoftGraphAlert) SetComments(v []string)`

SetComments gets a reference to the given []string and assigns it to the Comments field.

### GetConfidence

`func (o *MicrosoftGraphAlert) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *MicrosoftGraphAlert) GetConfidenceOk() (int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfidence

`func (o *MicrosoftGraphAlert) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidence

`func (o *MicrosoftGraphAlert) SetConfidence(v int32)`

SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.

### SetConfidenceExplicitNull

`func (o *MicrosoftGraphAlert) SetConfidenceExplicitNull(b bool)`

SetConfidenceExplicitNull (un)sets Confidence to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Confidence value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphAlert) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAlert) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphAlert) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAlert) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphAlert) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphAlert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAlert) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphAlert) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphAlert) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphAlert) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDetectionIds

`func (o *MicrosoftGraphAlert) GetDetectionIds() []string`

GetDetectionIds returns the DetectionIds field if non-nil, zero value otherwise.

### GetDetectionIdsOk

`func (o *MicrosoftGraphAlert) GetDetectionIdsOk() ([]string, bool)`

GetDetectionIdsOk returns a tuple with the DetectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetectionIds

`func (o *MicrosoftGraphAlert) HasDetectionIds() bool`

HasDetectionIds returns a boolean if a field has been set.

### SetDetectionIds

`func (o *MicrosoftGraphAlert) SetDetectionIds(v []string)`

SetDetectionIds gets a reference to the given []string and assigns it to the DetectionIds field.

### GetEventDateTime

`func (o *MicrosoftGraphAlert) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *MicrosoftGraphAlert) GetEventDateTimeOk() (time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventDateTime

`func (o *MicrosoftGraphAlert) HasEventDateTime() bool`

HasEventDateTime returns a boolean if a field has been set.

### SetEventDateTime

`func (o *MicrosoftGraphAlert) SetEventDateTime(v time.Time)`

SetEventDateTime gets a reference to the given time.Time and assigns it to the EventDateTime field.

### SetEventDateTimeExplicitNull

`func (o *MicrosoftGraphAlert) SetEventDateTimeExplicitNull(b bool)`

SetEventDateTimeExplicitNull (un)sets EventDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EventDateTime value is set to nil even if false is passed
### GetFeedback

`func (o *MicrosoftGraphAlert) GetFeedback() AnyOfmicrosoftGraphAlertFeedback`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *MicrosoftGraphAlert) GetFeedbackOk() (AnyOfmicrosoftGraphAlertFeedback, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeedback

`func (o *MicrosoftGraphAlert) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### SetFeedback

`func (o *MicrosoftGraphAlert) SetFeedback(v AnyOfmicrosoftGraphAlertFeedback)`

SetFeedback gets a reference to the given AnyOfmicrosoftGraphAlertFeedback and assigns it to the Feedback field.

### SetFeedbackExplicitNull

`func (o *MicrosoftGraphAlert) SetFeedbackExplicitNull(b bool)`

SetFeedbackExplicitNull (un)sets Feedback to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Feedback value is set to nil even if false is passed
### GetFileStates

`func (o *MicrosoftGraphAlert) GetFileStates() []AnyOfmicrosoftGraphFileSecurityState`

GetFileStates returns the FileStates field if non-nil, zero value otherwise.

### GetFileStatesOk

`func (o *MicrosoftGraphAlert) GetFileStatesOk() ([]AnyOfmicrosoftGraphFileSecurityState, bool)`

GetFileStatesOk returns a tuple with the FileStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileStates

`func (o *MicrosoftGraphAlert) HasFileStates() bool`

HasFileStates returns a boolean if a field has been set.

### SetFileStates

`func (o *MicrosoftGraphAlert) SetFileStates(v []AnyOfmicrosoftGraphFileSecurityState)`

SetFileStates gets a reference to the given []AnyOfmicrosoftGraphFileSecurityState and assigns it to the FileStates field.

### GetHistoryStates

`func (o *MicrosoftGraphAlert) GetHistoryStates() []AnyOfmicrosoftGraphAlertHistoryState`

GetHistoryStates returns the HistoryStates field if non-nil, zero value otherwise.

### GetHistoryStatesOk

`func (o *MicrosoftGraphAlert) GetHistoryStatesOk() ([]AnyOfmicrosoftGraphAlertHistoryState, bool)`

GetHistoryStatesOk returns a tuple with the HistoryStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHistoryStates

`func (o *MicrosoftGraphAlert) HasHistoryStates() bool`

HasHistoryStates returns a boolean if a field has been set.

### SetHistoryStates

`func (o *MicrosoftGraphAlert) SetHistoryStates(v []AnyOfmicrosoftGraphAlertHistoryState)`

SetHistoryStates gets a reference to the given []AnyOfmicrosoftGraphAlertHistoryState and assigns it to the HistoryStates field.

### GetHostStates

`func (o *MicrosoftGraphAlert) GetHostStates() []AnyOfmicrosoftGraphHostSecurityState`

GetHostStates returns the HostStates field if non-nil, zero value otherwise.

### GetHostStatesOk

`func (o *MicrosoftGraphAlert) GetHostStatesOk() ([]AnyOfmicrosoftGraphHostSecurityState, bool)`

GetHostStatesOk returns a tuple with the HostStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostStates

`func (o *MicrosoftGraphAlert) HasHostStates() bool`

HasHostStates returns a boolean if a field has been set.

### SetHostStates

`func (o *MicrosoftGraphAlert) SetHostStates(v []AnyOfmicrosoftGraphHostSecurityState)`

SetHostStates gets a reference to the given []AnyOfmicrosoftGraphHostSecurityState and assigns it to the HostStates field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphAlert) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphAlert) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphAlert) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphAlert) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphAlert) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetMalwareStates

`func (o *MicrosoftGraphAlert) GetMalwareStates() []AnyOfmicrosoftGraphMalwareState`

GetMalwareStates returns the MalwareStates field if non-nil, zero value otherwise.

### GetMalwareStatesOk

`func (o *MicrosoftGraphAlert) GetMalwareStatesOk() ([]AnyOfmicrosoftGraphMalwareState, bool)`

GetMalwareStatesOk returns a tuple with the MalwareStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMalwareStates

`func (o *MicrosoftGraphAlert) HasMalwareStates() bool`

HasMalwareStates returns a boolean if a field has been set.

### SetMalwareStates

`func (o *MicrosoftGraphAlert) SetMalwareStates(v []AnyOfmicrosoftGraphMalwareState)`

SetMalwareStates gets a reference to the given []AnyOfmicrosoftGraphMalwareState and assigns it to the MalwareStates field.

### GetNetworkConnections

`func (o *MicrosoftGraphAlert) GetNetworkConnections() []AnyOfmicrosoftGraphNetworkConnection`

GetNetworkConnections returns the NetworkConnections field if non-nil, zero value otherwise.

### GetNetworkConnectionsOk

`func (o *MicrosoftGraphAlert) GetNetworkConnectionsOk() ([]AnyOfmicrosoftGraphNetworkConnection, bool)`

GetNetworkConnectionsOk returns a tuple with the NetworkConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetworkConnections

`func (o *MicrosoftGraphAlert) HasNetworkConnections() bool`

HasNetworkConnections returns a boolean if a field has been set.

### SetNetworkConnections

`func (o *MicrosoftGraphAlert) SetNetworkConnections(v []AnyOfmicrosoftGraphNetworkConnection)`

SetNetworkConnections gets a reference to the given []AnyOfmicrosoftGraphNetworkConnection and assigns it to the NetworkConnections field.

### GetProcesses

`func (o *MicrosoftGraphAlert) GetProcesses() []AnyOfmicrosoftGraphProcess`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *MicrosoftGraphAlert) GetProcessesOk() ([]AnyOfmicrosoftGraphProcess, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProcesses

`func (o *MicrosoftGraphAlert) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### SetProcesses

`func (o *MicrosoftGraphAlert) SetProcesses(v []AnyOfmicrosoftGraphProcess)`

SetProcesses gets a reference to the given []AnyOfmicrosoftGraphProcess and assigns it to the Processes field.

### GetRecommendedActions

`func (o *MicrosoftGraphAlert) GetRecommendedActions() []string`

GetRecommendedActions returns the RecommendedActions field if non-nil, zero value otherwise.

### GetRecommendedActionsOk

`func (o *MicrosoftGraphAlert) GetRecommendedActionsOk() ([]string, bool)`

GetRecommendedActionsOk returns a tuple with the RecommendedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecommendedActions

`func (o *MicrosoftGraphAlert) HasRecommendedActions() bool`

HasRecommendedActions returns a boolean if a field has been set.

### SetRecommendedActions

`func (o *MicrosoftGraphAlert) SetRecommendedActions(v []string)`

SetRecommendedActions gets a reference to the given []string and assigns it to the RecommendedActions field.

### GetRegistryKeyStates

`func (o *MicrosoftGraphAlert) GetRegistryKeyStates() []AnyOfmicrosoftGraphRegistryKeyState`

GetRegistryKeyStates returns the RegistryKeyStates field if non-nil, zero value otherwise.

### GetRegistryKeyStatesOk

`func (o *MicrosoftGraphAlert) GetRegistryKeyStatesOk() ([]AnyOfmicrosoftGraphRegistryKeyState, bool)`

GetRegistryKeyStatesOk returns a tuple with the RegistryKeyStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRegistryKeyStates

`func (o *MicrosoftGraphAlert) HasRegistryKeyStates() bool`

HasRegistryKeyStates returns a boolean if a field has been set.

### SetRegistryKeyStates

`func (o *MicrosoftGraphAlert) SetRegistryKeyStates(v []AnyOfmicrosoftGraphRegistryKeyState)`

SetRegistryKeyStates gets a reference to the given []AnyOfmicrosoftGraphRegistryKeyState and assigns it to the RegistryKeyStates field.

### GetSeverity

`func (o *MicrosoftGraphAlert) GetSeverity() AnyOfmicrosoftGraphAlertSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *MicrosoftGraphAlert) GetSeverityOk() (AnyOfmicrosoftGraphAlertSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSeverity

`func (o *MicrosoftGraphAlert) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverity

`func (o *MicrosoftGraphAlert) SetSeverity(v AnyOfmicrosoftGraphAlertSeverity)`

SetSeverity gets a reference to the given AnyOfmicrosoftGraphAlertSeverity and assigns it to the Severity field.

### GetSourceMaterials

`func (o *MicrosoftGraphAlert) GetSourceMaterials() []string`

GetSourceMaterials returns the SourceMaterials field if non-nil, zero value otherwise.

### GetSourceMaterialsOk

`func (o *MicrosoftGraphAlert) GetSourceMaterialsOk() ([]string, bool)`

GetSourceMaterialsOk returns a tuple with the SourceMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceMaterials

`func (o *MicrosoftGraphAlert) HasSourceMaterials() bool`

HasSourceMaterials returns a boolean if a field has been set.

### SetSourceMaterials

`func (o *MicrosoftGraphAlert) SetSourceMaterials(v []string)`

SetSourceMaterials gets a reference to the given []string and assigns it to the SourceMaterials field.

### GetStatus

`func (o *MicrosoftGraphAlert) GetStatus() AnyOfmicrosoftGraphAlertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphAlert) GetStatusOk() (AnyOfmicrosoftGraphAlertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphAlert) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphAlert) SetStatus(v AnyOfmicrosoftGraphAlertStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphAlertStatus and assigns it to the Status field.

### GetTags

`func (o *MicrosoftGraphAlert) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MicrosoftGraphAlert) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *MicrosoftGraphAlert) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *MicrosoftGraphAlert) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetTitle

`func (o *MicrosoftGraphAlert) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphAlert) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *MicrosoftGraphAlert) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *MicrosoftGraphAlert) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### SetTitleExplicitNull

`func (o *MicrosoftGraphAlert) SetTitleExplicitNull(b bool)`

SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Title value is set to nil even if false is passed
### GetTriggers

`func (o *MicrosoftGraphAlert) GetTriggers() []AnyOfmicrosoftGraphAlertTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *MicrosoftGraphAlert) GetTriggersOk() ([]AnyOfmicrosoftGraphAlertTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTriggers

`func (o *MicrosoftGraphAlert) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggers

`func (o *MicrosoftGraphAlert) SetTriggers(v []AnyOfmicrosoftGraphAlertTrigger)`

SetTriggers gets a reference to the given []AnyOfmicrosoftGraphAlertTrigger and assigns it to the Triggers field.

### GetUserStates

`func (o *MicrosoftGraphAlert) GetUserStates() []AnyOfmicrosoftGraphUserSecurityState`

GetUserStates returns the UserStates field if non-nil, zero value otherwise.

### GetUserStatesOk

`func (o *MicrosoftGraphAlert) GetUserStatesOk() ([]AnyOfmicrosoftGraphUserSecurityState, bool)`

GetUserStatesOk returns a tuple with the UserStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStates

`func (o *MicrosoftGraphAlert) HasUserStates() bool`

HasUserStates returns a boolean if a field has been set.

### SetUserStates

`func (o *MicrosoftGraphAlert) SetUserStates(v []AnyOfmicrosoftGraphUserSecurityState)`

SetUserStates gets a reference to the given []AnyOfmicrosoftGraphUserSecurityState and assigns it to the UserStates field.

### GetVendorInformation

`func (o *MicrosoftGraphAlert) GetVendorInformation() AnyOfmicrosoftGraphSecurityVendorInformation`

GetVendorInformation returns the VendorInformation field if non-nil, zero value otherwise.

### GetVendorInformationOk

`func (o *MicrosoftGraphAlert) GetVendorInformationOk() (AnyOfmicrosoftGraphSecurityVendorInformation, bool)`

GetVendorInformationOk returns a tuple with the VendorInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVendorInformation

`func (o *MicrosoftGraphAlert) HasVendorInformation() bool`

HasVendorInformation returns a boolean if a field has been set.

### SetVendorInformation

`func (o *MicrosoftGraphAlert) SetVendorInformation(v AnyOfmicrosoftGraphSecurityVendorInformation)`

SetVendorInformation gets a reference to the given AnyOfmicrosoftGraphSecurityVendorInformation and assigns it to the VendorInformation field.

### SetVendorInformationExplicitNull

`func (o *MicrosoftGraphAlert) SetVendorInformationExplicitNull(b bool)`

SetVendorInformationExplicitNull (un)sets VendorInformation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VendorInformation value is set to nil even if false is passed
### GetVulnerabilityStates

`func (o *MicrosoftGraphAlert) GetVulnerabilityStates() []AnyOfmicrosoftGraphVulnerabilityState`

GetVulnerabilityStates returns the VulnerabilityStates field if non-nil, zero value otherwise.

### GetVulnerabilityStatesOk

`func (o *MicrosoftGraphAlert) GetVulnerabilityStatesOk() ([]AnyOfmicrosoftGraphVulnerabilityState, bool)`

GetVulnerabilityStatesOk returns a tuple with the VulnerabilityStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVulnerabilityStates

`func (o *MicrosoftGraphAlert) HasVulnerabilityStates() bool`

HasVulnerabilityStates returns a boolean if a field has been set.

### SetVulnerabilityStates

`func (o *MicrosoftGraphAlert) SetVulnerabilityStates(v []AnyOfmicrosoftGraphVulnerabilityState)`

SetVulnerabilityStates gets a reference to the given []AnyOfmicrosoftGraphVulnerabilityState and assigns it to the VulnerabilityStates field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


