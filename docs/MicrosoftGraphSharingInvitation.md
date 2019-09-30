# MicrosoftGraphSharingInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**InvitedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**RedeemedBy** | Pointer to **string** |  | [optional] 
**SignInRequired** | Pointer to **bool** |  | [optional] 

## Methods

### GetEmail

`func (o *MicrosoftGraphSharingInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MicrosoftGraphSharingInvitation) GetEmailOk() (string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmail

`func (o *MicrosoftGraphSharingInvitation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmail

`func (o *MicrosoftGraphSharingInvitation) SetEmail(v string)`

SetEmail gets a reference to the given string and assigns it to the Email field.

### SetEmailExplicitNull

`func (o *MicrosoftGraphSharingInvitation) SetEmailExplicitNull(b bool)`

SetEmailExplicitNull (un)sets Email to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Email value is set to nil even if false is passed
### GetInvitedBy

`func (o *MicrosoftGraphSharingInvitation) GetInvitedBy() AnyOfmicrosoftGraphIdentitySet`

GetInvitedBy returns the InvitedBy field if non-nil, zero value otherwise.

### GetInvitedByOk

`func (o *MicrosoftGraphSharingInvitation) GetInvitedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetInvitedByOk returns a tuple with the InvitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitedBy

`func (o *MicrosoftGraphSharingInvitation) HasInvitedBy() bool`

HasInvitedBy returns a boolean if a field has been set.

### SetInvitedBy

`func (o *MicrosoftGraphSharingInvitation) SetInvitedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetInvitedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the InvitedBy field.

### SetInvitedByExplicitNull

`func (o *MicrosoftGraphSharingInvitation) SetInvitedByExplicitNull(b bool)`

SetInvitedByExplicitNull (un)sets InvitedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InvitedBy value is set to nil even if false is passed
### GetRedeemedBy

`func (o *MicrosoftGraphSharingInvitation) GetRedeemedBy() string`

GetRedeemedBy returns the RedeemedBy field if non-nil, zero value otherwise.

### GetRedeemedByOk

`func (o *MicrosoftGraphSharingInvitation) GetRedeemedByOk() (string, bool)`

GetRedeemedByOk returns a tuple with the RedeemedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRedeemedBy

`func (o *MicrosoftGraphSharingInvitation) HasRedeemedBy() bool`

HasRedeemedBy returns a boolean if a field has been set.

### SetRedeemedBy

`func (o *MicrosoftGraphSharingInvitation) SetRedeemedBy(v string)`

SetRedeemedBy gets a reference to the given string and assigns it to the RedeemedBy field.

### SetRedeemedByExplicitNull

`func (o *MicrosoftGraphSharingInvitation) SetRedeemedByExplicitNull(b bool)`

SetRedeemedByExplicitNull (un)sets RedeemedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RedeemedBy value is set to nil even if false is passed
### GetSignInRequired

`func (o *MicrosoftGraphSharingInvitation) GetSignInRequired() bool`

GetSignInRequired returns the SignInRequired field if non-nil, zero value otherwise.

### GetSignInRequiredOk

`func (o *MicrosoftGraphSharingInvitation) GetSignInRequiredOk() (bool, bool)`

GetSignInRequiredOk returns a tuple with the SignInRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSignInRequired

`func (o *MicrosoftGraphSharingInvitation) HasSignInRequired() bool`

HasSignInRequired returns a boolean if a field has been set.

### SetSignInRequired

`func (o *MicrosoftGraphSharingInvitation) SetSignInRequired(v bool)`

SetSignInRequired gets a reference to the given bool and assigns it to the SignInRequired field.

### SetSignInRequiredExplicitNull

`func (o *MicrosoftGraphSharingInvitation) SetSignInRequiredExplicitNull(b bool)`

SetSignInRequiredExplicitNull (un)sets SignInRequired to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SignInRequired value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


