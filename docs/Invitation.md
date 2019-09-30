# Invitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvitedUserDisplayName** | Pointer to **string** |  | [optional] 
**InvitedUserType** | Pointer to **string** |  | [optional] 
**InvitedUserEmailAddress** | Pointer to **string** |  | [optional] 
**InvitedUserMessageInfo** | Pointer to [**AnyOfmicrosoftGraphInvitedUserMessageInfo**](anyOf&lt;microsoft.graph.invitedUserMessageInfo&gt;.md) |  | [optional] 
**SendInvitationMessage** | Pointer to **bool** |  | [optional] 
**InviteRedirectUrl** | Pointer to **string** |  | [optional] 
**InviteRedeemUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**InvitedUser** | Pointer to [**AnyOfmicrosoftGraphUser**](anyOf&lt;microsoft.graph.user&gt;.md) |  | [optional] 

## Methods

### GetInvitedUserDisplayName

`func (o *Invitation) GetInvitedUserDisplayName() string`

GetInvitedUserDisplayName returns the InvitedUserDisplayName field if non-nil, zero value otherwise.

### GetInvitedUserDisplayNameOk

`func (o *Invitation) GetInvitedUserDisplayNameOk() (string, bool)`

GetInvitedUserDisplayNameOk returns a tuple with the InvitedUserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitedUserDisplayName

`func (o *Invitation) HasInvitedUserDisplayName() bool`

HasInvitedUserDisplayName returns a boolean if a field has been set.

### SetInvitedUserDisplayName

`func (o *Invitation) SetInvitedUserDisplayName(v string)`

SetInvitedUserDisplayName gets a reference to the given string and assigns it to the InvitedUserDisplayName field.

### SetInvitedUserDisplayNameExplicitNull

`func (o *Invitation) SetInvitedUserDisplayNameExplicitNull(b bool)`

SetInvitedUserDisplayNameExplicitNull (un)sets InvitedUserDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InvitedUserDisplayName value is set to nil even if false is passed
### GetInvitedUserType

`func (o *Invitation) GetInvitedUserType() string`

GetInvitedUserType returns the InvitedUserType field if non-nil, zero value otherwise.

### GetInvitedUserTypeOk

`func (o *Invitation) GetInvitedUserTypeOk() (string, bool)`

GetInvitedUserTypeOk returns a tuple with the InvitedUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitedUserType

`func (o *Invitation) HasInvitedUserType() bool`

HasInvitedUserType returns a boolean if a field has been set.

### SetInvitedUserType

`func (o *Invitation) SetInvitedUserType(v string)`

SetInvitedUserType gets a reference to the given string and assigns it to the InvitedUserType field.

### SetInvitedUserTypeExplicitNull

`func (o *Invitation) SetInvitedUserTypeExplicitNull(b bool)`

SetInvitedUserTypeExplicitNull (un)sets InvitedUserType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InvitedUserType value is set to nil even if false is passed
### GetInvitedUserEmailAddress

`func (o *Invitation) GetInvitedUserEmailAddress() string`

GetInvitedUserEmailAddress returns the InvitedUserEmailAddress field if non-nil, zero value otherwise.

### GetInvitedUserEmailAddressOk

`func (o *Invitation) GetInvitedUserEmailAddressOk() (string, bool)`

GetInvitedUserEmailAddressOk returns a tuple with the InvitedUserEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitedUserEmailAddress

`func (o *Invitation) HasInvitedUserEmailAddress() bool`

HasInvitedUserEmailAddress returns a boolean if a field has been set.

### SetInvitedUserEmailAddress

`func (o *Invitation) SetInvitedUserEmailAddress(v string)`

SetInvitedUserEmailAddress gets a reference to the given string and assigns it to the InvitedUserEmailAddress field.

### GetInvitedUserMessageInfo

`func (o *Invitation) GetInvitedUserMessageInfo() AnyOfmicrosoftGraphInvitedUserMessageInfo`

GetInvitedUserMessageInfo returns the InvitedUserMessageInfo field if non-nil, zero value otherwise.

### GetInvitedUserMessageInfoOk

`func (o *Invitation) GetInvitedUserMessageInfoOk() (AnyOfmicrosoftGraphInvitedUserMessageInfo, bool)`

GetInvitedUserMessageInfoOk returns a tuple with the InvitedUserMessageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitedUserMessageInfo

`func (o *Invitation) HasInvitedUserMessageInfo() bool`

HasInvitedUserMessageInfo returns a boolean if a field has been set.

### SetInvitedUserMessageInfo

`func (o *Invitation) SetInvitedUserMessageInfo(v AnyOfmicrosoftGraphInvitedUserMessageInfo)`

SetInvitedUserMessageInfo gets a reference to the given AnyOfmicrosoftGraphInvitedUserMessageInfo and assigns it to the InvitedUserMessageInfo field.

### SetInvitedUserMessageInfoExplicitNull

`func (o *Invitation) SetInvitedUserMessageInfoExplicitNull(b bool)`

SetInvitedUserMessageInfoExplicitNull (un)sets InvitedUserMessageInfo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InvitedUserMessageInfo value is set to nil even if false is passed
### GetSendInvitationMessage

`func (o *Invitation) GetSendInvitationMessage() bool`

GetSendInvitationMessage returns the SendInvitationMessage field if non-nil, zero value otherwise.

### GetSendInvitationMessageOk

`func (o *Invitation) GetSendInvitationMessageOk() (bool, bool)`

GetSendInvitationMessageOk returns a tuple with the SendInvitationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSendInvitationMessage

`func (o *Invitation) HasSendInvitationMessage() bool`

HasSendInvitationMessage returns a boolean if a field has been set.

### SetSendInvitationMessage

`func (o *Invitation) SetSendInvitationMessage(v bool)`

SetSendInvitationMessage gets a reference to the given bool and assigns it to the SendInvitationMessage field.

### SetSendInvitationMessageExplicitNull

`func (o *Invitation) SetSendInvitationMessageExplicitNull(b bool)`

SetSendInvitationMessageExplicitNull (un)sets SendInvitationMessage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SendInvitationMessage value is set to nil even if false is passed
### GetInviteRedirectUrl

`func (o *Invitation) GetInviteRedirectUrl() string`

GetInviteRedirectUrl returns the InviteRedirectUrl field if non-nil, zero value otherwise.

### GetInviteRedirectUrlOk

`func (o *Invitation) GetInviteRedirectUrlOk() (string, bool)`

GetInviteRedirectUrlOk returns a tuple with the InviteRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInviteRedirectUrl

`func (o *Invitation) HasInviteRedirectUrl() bool`

HasInviteRedirectUrl returns a boolean if a field has been set.

### SetInviteRedirectUrl

`func (o *Invitation) SetInviteRedirectUrl(v string)`

SetInviteRedirectUrl gets a reference to the given string and assigns it to the InviteRedirectUrl field.

### GetInviteRedeemUrl

`func (o *Invitation) GetInviteRedeemUrl() string`

GetInviteRedeemUrl returns the InviteRedeemUrl field if non-nil, zero value otherwise.

### GetInviteRedeemUrlOk

`func (o *Invitation) GetInviteRedeemUrlOk() (string, bool)`

GetInviteRedeemUrlOk returns a tuple with the InviteRedeemUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInviteRedeemUrl

`func (o *Invitation) HasInviteRedeemUrl() bool`

HasInviteRedeemUrl returns a boolean if a field has been set.

### SetInviteRedeemUrl

`func (o *Invitation) SetInviteRedeemUrl(v string)`

SetInviteRedeemUrl gets a reference to the given string and assigns it to the InviteRedeemUrl field.

### SetInviteRedeemUrlExplicitNull

`func (o *Invitation) SetInviteRedeemUrlExplicitNull(b bool)`

SetInviteRedeemUrlExplicitNull (un)sets InviteRedeemUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InviteRedeemUrl value is set to nil even if false is passed
### GetStatus

`func (o *Invitation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invitation) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *Invitation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *Invitation) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *Invitation) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetInvitedUser

`func (o *Invitation) GetInvitedUser() AnyOfmicrosoftGraphUser`

GetInvitedUser returns the InvitedUser field if non-nil, zero value otherwise.

### GetInvitedUserOk

`func (o *Invitation) GetInvitedUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetInvitedUserOk returns a tuple with the InvitedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitedUser

`func (o *Invitation) HasInvitedUser() bool`

HasInvitedUser returns a boolean if a field has been set.

### SetInvitedUser

`func (o *Invitation) SetInvitedUser(v AnyOfmicrosoftGraphUser)`

SetInvitedUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the InvitedUser field.

### SetInvitedUserExplicitNull

`func (o *Invitation) SetInvitedUserExplicitNull(b bool)`

SetInvitedUserExplicitNull (un)sets InvitedUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InvitedUser value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


