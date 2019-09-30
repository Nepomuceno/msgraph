# InlineObject742

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireSignIn** | Pointer to **bool** |  | [optional] [default to false]
**Roles** | Pointer to **[]string** |  | [optional] 
**SendInvitation** | Pointer to **bool** |  | [optional] [default to false]
**Message** | Pointer to **string** |  | [optional] 
**Recipients** | Pointer to [**[]MicrosoftGraphDriveRecipient**](microsoft.graph.driveRecipient.md) |  | [optional] 

## Methods

### GetRequireSignIn

`func (o *InlineObject742) GetRequireSignIn() bool`

GetRequireSignIn returns the RequireSignIn field if non-nil, zero value otherwise.

### GetRequireSignInOk

`func (o *InlineObject742) GetRequireSignInOk() (bool, bool)`

GetRequireSignInOk returns a tuple with the RequireSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequireSignIn

`func (o *InlineObject742) HasRequireSignIn() bool`

HasRequireSignIn returns a boolean if a field has been set.

### SetRequireSignIn

`func (o *InlineObject742) SetRequireSignIn(v bool)`

SetRequireSignIn gets a reference to the given bool and assigns it to the RequireSignIn field.

### SetRequireSignInExplicitNull

`func (o *InlineObject742) SetRequireSignInExplicitNull(b bool)`

SetRequireSignInExplicitNull (un)sets RequireSignIn to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RequireSignIn value is set to nil even if false is passed
### GetRoles

`func (o *InlineObject742) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *InlineObject742) GetRolesOk() ([]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoles

`func (o *InlineObject742) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRoles

`func (o *InlineObject742) SetRoles(v []string)`

SetRoles gets a reference to the given []string and assigns it to the Roles field.

### GetSendInvitation

`func (o *InlineObject742) GetSendInvitation() bool`

GetSendInvitation returns the SendInvitation field if non-nil, zero value otherwise.

### GetSendInvitationOk

`func (o *InlineObject742) GetSendInvitationOk() (bool, bool)`

GetSendInvitationOk returns a tuple with the SendInvitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSendInvitation

`func (o *InlineObject742) HasSendInvitation() bool`

HasSendInvitation returns a boolean if a field has been set.

### SetSendInvitation

`func (o *InlineObject742) SetSendInvitation(v bool)`

SetSendInvitation gets a reference to the given bool and assigns it to the SendInvitation field.

### SetSendInvitationExplicitNull

`func (o *InlineObject742) SetSendInvitationExplicitNull(b bool)`

SetSendInvitationExplicitNull (un)sets SendInvitation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SendInvitation value is set to nil even if false is passed
### GetMessage

`func (o *InlineObject742) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineObject742) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *InlineObject742) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *InlineObject742) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### SetMessageExplicitNull

`func (o *InlineObject742) SetMessageExplicitNull(b bool)`

SetMessageExplicitNull (un)sets Message to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Message value is set to nil even if false is passed
### GetRecipients

`func (o *InlineObject742) GetRecipients() []MicrosoftGraphDriveRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineObject742) GetRecipientsOk() ([]MicrosoftGraphDriveRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipients

`func (o *InlineObject742) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipients

`func (o *InlineObject742) SetRecipients(v []MicrosoftGraphDriveRecipient)`

SetRecipients gets a reference to the given []MicrosoftGraphDriveRecipient and assigns it to the Recipients field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


