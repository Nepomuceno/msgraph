# MicrosoftGraphInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphInvitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphInvitation) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphInvitation) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphInvitation) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetInvitedUserDisplayName

`func (o *MicrosoftGraphInvitation) GetInvitedUserDisplayName() string`

GetInvitedUserDisplayName returns the InvitedUserDisplayName field if non-nil, zero value otherwise.

### GetInvitedUserDisplayNameOk

`func (o *MicrosoftGraphInvitation) GetInvitedUserDisplayNameOk() (string, bool)`

GetInvitedUserDisplayNameOk returns a tuple with the InvitedUserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitedUserDisplayName

`func (o *MicrosoftGraphInvitation) HasInvitedUserDisplayName() bool`

HasInvitedUserDisplayName returns a boolean if a field has been set.

### SetInvitedUserDisplayName

`func (o *MicrosoftGraphInvitation) SetInvitedUserDisplayName(v string)`

SetInvitedUserDisplayName gets a reference to the given string and assigns it to the InvitedUserDisplayName field.

### SetInvitedUserDisplayNameExplicitNull

`func (o *MicrosoftGraphInvitation) SetInvitedUserDisplayNameExplicitNull(b bool)`

SetInvitedUserDisplayNameExplicitNull (un)sets InvitedUserDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InvitedUserDisplayName value is set to nil even if false is passed
### GetInvitedUserType

`func (o *MicrosoftGraphInvitation) GetInvitedUserType() string`

GetInvitedUserType returns the InvitedUserType field if non-nil, zero value otherwise.

### GetInvitedUserTypeOk

`func (o *MicrosoftGraphInvitation) GetInvitedUserTypeOk() (string, bool)`

GetInvitedUserTypeOk returns a tuple with the InvitedUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitedUserType

`func (o *MicrosoftGraphInvitation) HasInvitedUserType() bool`

HasInvitedUserType returns a boolean if a field has been set.

### SetInvitedUserType

`func (o *MicrosoftGraphInvitation) SetInvitedUserType(v string)`

SetInvitedUserType gets a reference to the given string and assigns it to the InvitedUserType field.

### SetInvitedUserTypeExplicitNull

`func (o *MicrosoftGraphInvitation) SetInvitedUserTypeExplicitNull(b bool)`

SetInvitedUserTypeExplicitNull (un)sets InvitedUserType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InvitedUserType value is set to nil even if false is passed
### GetInvitedUserEmailAddress

`func (o *MicrosoftGraphInvitation) GetInvitedUserEmailAddress() string`

GetInvitedUserEmailAddress returns the InvitedUserEmailAddress field if non-nil, zero value otherwise.

### GetInvitedUserEmailAddressOk

`func (o *MicrosoftGraphInvitation) GetInvitedUserEmailAddressOk() (string, bool)`

GetInvitedUserEmailAddressOk returns a tuple with the InvitedUserEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitedUserEmailAddress

`func (o *MicrosoftGraphInvitation) HasInvitedUserEmailAddress() bool`

HasInvitedUserEmailAddress returns a boolean if a field has been set.

### SetInvitedUserEmailAddress

`func (o *MicrosoftGraphInvitation) SetInvitedUserEmailAddress(v string)`

SetInvitedUserEmailAddress gets a reference to the given string and assigns it to the InvitedUserEmailAddress field.

### GetInvitedUserMessageInfo

`func (o *MicrosoftGraphInvitation) GetInvitedUserMessageInfo() AnyOfmicrosoftGraphInvitedUserMessageInfo`

GetInvitedUserMessageInfo returns the InvitedUserMessageInfo field if non-nil, zero value otherwise.

### GetInvitedUserMessageInfoOk

`func (o *MicrosoftGraphInvitation) GetInvitedUserMessageInfoOk() (AnyOfmicrosoftGraphInvitedUserMessageInfo, bool)`

GetInvitedUserMessageInfoOk returns a tuple with the InvitedUserMessageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitedUserMessageInfo

`func (o *MicrosoftGraphInvitation) HasInvitedUserMessageInfo() bool`

HasInvitedUserMessageInfo returns a boolean if a field has been set.

### SetInvitedUserMessageInfo

`func (o *MicrosoftGraphInvitation) SetInvitedUserMessageInfo(v AnyOfmicrosoftGraphInvitedUserMessageInfo)`

SetInvitedUserMessageInfo gets a reference to the given AnyOfmicrosoftGraphInvitedUserMessageInfo and assigns it to the InvitedUserMessageInfo field.

### SetInvitedUserMessageInfoExplicitNull

`func (o *MicrosoftGraphInvitation) SetInvitedUserMessageInfoExplicitNull(b bool)`

SetInvitedUserMessageInfoExplicitNull (un)sets InvitedUserMessageInfo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InvitedUserMessageInfo value is set to nil even if false is passed
### GetSendInvitationMessage

`func (o *MicrosoftGraphInvitation) GetSendInvitationMessage() bool`

GetSendInvitationMessage returns the SendInvitationMessage field if non-nil, zero value otherwise.

### GetSendInvitationMessageOk

`func (o *MicrosoftGraphInvitation) GetSendInvitationMessageOk() (bool, bool)`

GetSendInvitationMessageOk returns a tuple with the SendInvitationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSendInvitationMessage

`func (o *MicrosoftGraphInvitation) HasSendInvitationMessage() bool`

HasSendInvitationMessage returns a boolean if a field has been set.

### SetSendInvitationMessage

`func (o *MicrosoftGraphInvitation) SetSendInvitationMessage(v bool)`

SetSendInvitationMessage gets a reference to the given bool and assigns it to the SendInvitationMessage field.

### SetSendInvitationMessageExplicitNull

`func (o *MicrosoftGraphInvitation) SetSendInvitationMessageExplicitNull(b bool)`

SetSendInvitationMessageExplicitNull (un)sets SendInvitationMessage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SendInvitationMessage value is set to nil even if false is passed
### GetInviteRedirectUrl

`func (o *MicrosoftGraphInvitation) GetInviteRedirectUrl() string`

GetInviteRedirectUrl returns the InviteRedirectUrl field if non-nil, zero value otherwise.

### GetInviteRedirectUrlOk

`func (o *MicrosoftGraphInvitation) GetInviteRedirectUrlOk() (string, bool)`

GetInviteRedirectUrlOk returns a tuple with the InviteRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInviteRedirectUrl

`func (o *MicrosoftGraphInvitation) HasInviteRedirectUrl() bool`

HasInviteRedirectUrl returns a boolean if a field has been set.

### SetInviteRedirectUrl

`func (o *MicrosoftGraphInvitation) SetInviteRedirectUrl(v string)`

SetInviteRedirectUrl gets a reference to the given string and assigns it to the InviteRedirectUrl field.

### GetInviteRedeemUrl

`func (o *MicrosoftGraphInvitation) GetInviteRedeemUrl() string`

GetInviteRedeemUrl returns the InviteRedeemUrl field if non-nil, zero value otherwise.

### GetInviteRedeemUrlOk

`func (o *MicrosoftGraphInvitation) GetInviteRedeemUrlOk() (string, bool)`

GetInviteRedeemUrlOk returns a tuple with the InviteRedeemUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInviteRedeemUrl

`func (o *MicrosoftGraphInvitation) HasInviteRedeemUrl() bool`

HasInviteRedeemUrl returns a boolean if a field has been set.

### SetInviteRedeemUrl

`func (o *MicrosoftGraphInvitation) SetInviteRedeemUrl(v string)`

SetInviteRedeemUrl gets a reference to the given string and assigns it to the InviteRedeemUrl field.

### SetInviteRedeemUrlExplicitNull

`func (o *MicrosoftGraphInvitation) SetInviteRedeemUrlExplicitNull(b bool)`

SetInviteRedeemUrlExplicitNull (un)sets InviteRedeemUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InviteRedeemUrl value is set to nil even if false is passed
### GetStatus

`func (o *MicrosoftGraphInvitation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphInvitation) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphInvitation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphInvitation) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphInvitation) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetInvitedUser

`func (o *MicrosoftGraphInvitation) GetInvitedUser() AnyOfmicrosoftGraphUser`

GetInvitedUser returns the InvitedUser field if non-nil, zero value otherwise.

### GetInvitedUserOk

`func (o *MicrosoftGraphInvitation) GetInvitedUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetInvitedUserOk returns a tuple with the InvitedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitedUser

`func (o *MicrosoftGraphInvitation) HasInvitedUser() bool`

HasInvitedUser returns a boolean if a field has been set.

### SetInvitedUser

`func (o *MicrosoftGraphInvitation) SetInvitedUser(v AnyOfmicrosoftGraphUser)`

SetInvitedUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the InvitedUser field.

### SetInvitedUserExplicitNull

`func (o *MicrosoftGraphInvitation) SetInvitedUserExplicitNull(b bool)`

SetInvitedUserExplicitNull (un)sets InvitedUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InvitedUser value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


