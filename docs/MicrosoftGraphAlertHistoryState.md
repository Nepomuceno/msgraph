# MicrosoftGraphAlertHistoryState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** |  | [optional] 
**AssignedTo** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **[]string** |  | [optional] 
**Feedback** | Pointer to [**AnyOfmicrosoftGraphAlertFeedback**](anyOf&lt;microsoft.graph.alertFeedback&gt;.md) |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphAlertStatus**](anyOf&lt;microsoft.graph.alertStatus&gt;.md) |  | [optional] 
**UpdatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### GetAppId

`func (o *MicrosoftGraphAlertHistoryState) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MicrosoftGraphAlertHistoryState) GetAppIdOk() (string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppId

`func (o *MicrosoftGraphAlertHistoryState) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppId

`func (o *MicrosoftGraphAlertHistoryState) SetAppId(v string)`

SetAppId gets a reference to the given string and assigns it to the AppId field.

### SetAppIdExplicitNull

`func (o *MicrosoftGraphAlertHistoryState) SetAppIdExplicitNull(b bool)`

SetAppIdExplicitNull (un)sets AppId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppId value is set to nil even if false is passed
### GetAssignedTo

`func (o *MicrosoftGraphAlertHistoryState) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *MicrosoftGraphAlertHistoryState) GetAssignedToOk() (string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedTo

`func (o *MicrosoftGraphAlertHistoryState) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### SetAssignedTo

`func (o *MicrosoftGraphAlertHistoryState) SetAssignedTo(v string)`

SetAssignedTo gets a reference to the given string and assigns it to the AssignedTo field.

### SetAssignedToExplicitNull

`func (o *MicrosoftGraphAlertHistoryState) SetAssignedToExplicitNull(b bool)`

SetAssignedToExplicitNull (un)sets AssignedTo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AssignedTo value is set to nil even if false is passed
### GetComments

`func (o *MicrosoftGraphAlertHistoryState) GetComments() []string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *MicrosoftGraphAlertHistoryState) GetCommentsOk() ([]string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComments

`func (o *MicrosoftGraphAlertHistoryState) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetComments

`func (o *MicrosoftGraphAlertHistoryState) SetComments(v []string)`

SetComments gets a reference to the given []string and assigns it to the Comments field.

### GetFeedback

`func (o *MicrosoftGraphAlertHistoryState) GetFeedback() AnyOfmicrosoftGraphAlertFeedback`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *MicrosoftGraphAlertHistoryState) GetFeedbackOk() (AnyOfmicrosoftGraphAlertFeedback, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeedback

`func (o *MicrosoftGraphAlertHistoryState) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### SetFeedback

`func (o *MicrosoftGraphAlertHistoryState) SetFeedback(v AnyOfmicrosoftGraphAlertFeedback)`

SetFeedback gets a reference to the given AnyOfmicrosoftGraphAlertFeedback and assigns it to the Feedback field.

### SetFeedbackExplicitNull

`func (o *MicrosoftGraphAlertHistoryState) SetFeedbackExplicitNull(b bool)`

SetFeedbackExplicitNull (un)sets Feedback to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Feedback value is set to nil even if false is passed
### GetStatus

`func (o *MicrosoftGraphAlertHistoryState) GetStatus() AnyOfmicrosoftGraphAlertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphAlertHistoryState) GetStatusOk() (AnyOfmicrosoftGraphAlertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphAlertHistoryState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphAlertHistoryState) SetStatus(v AnyOfmicrosoftGraphAlertStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphAlertStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphAlertHistoryState) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetUpdatedDateTime

`func (o *MicrosoftGraphAlertHistoryState) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *MicrosoftGraphAlertHistoryState) GetUpdatedDateTimeOk() (time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdatedDateTime

`func (o *MicrosoftGraphAlertHistoryState) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### SetUpdatedDateTime

`func (o *MicrosoftGraphAlertHistoryState) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime gets a reference to the given time.Time and assigns it to the UpdatedDateTime field.

### SetUpdatedDateTimeExplicitNull

`func (o *MicrosoftGraphAlertHistoryState) SetUpdatedDateTimeExplicitNull(b bool)`

SetUpdatedDateTimeExplicitNull (un)sets UpdatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UpdatedDateTime value is set to nil even if false is passed
### GetUser

`func (o *MicrosoftGraphAlertHistoryState) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MicrosoftGraphAlertHistoryState) GetUserOk() (string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUser

`func (o *MicrosoftGraphAlertHistoryState) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUser

`func (o *MicrosoftGraphAlertHistoryState) SetUser(v string)`

SetUser gets a reference to the given string and assigns it to the User field.

### SetUserExplicitNull

`func (o *MicrosoftGraphAlertHistoryState) SetUserExplicitNull(b bool)`

SetUserExplicitNull (un)sets User to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The User value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


