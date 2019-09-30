# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantedTo** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**InheritedFrom** | Pointer to [**AnyOfmicrosoftGraphItemReference**](anyOf&lt;microsoft.graph.itemReference&gt;.md) |  | [optional] 
**Invitation** | Pointer to [**AnyOfmicrosoftGraphSharingInvitation**](anyOf&lt;microsoft.graph.sharingInvitation&gt;.md) |  | [optional] 
**Link** | Pointer to [**AnyOfmicrosoftGraphSharingLink**](anyOf&lt;microsoft.graph.sharingLink&gt;.md) |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**ShareId** | Pointer to **string** |  | [optional] 

## Methods

### GetGrantedTo

`func (o *Permission) GetGrantedTo() AnyOfmicrosoftGraphIdentitySet`

GetGrantedTo returns the GrantedTo field if non-nil, zero value otherwise.

### GetGrantedToOk

`func (o *Permission) GetGrantedToOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetGrantedToOk returns a tuple with the GrantedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGrantedTo

`func (o *Permission) HasGrantedTo() bool`

HasGrantedTo returns a boolean if a field has been set.

### SetGrantedTo

`func (o *Permission) SetGrantedTo(v AnyOfmicrosoftGraphIdentitySet)`

SetGrantedTo gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the GrantedTo field.

### SetGrantedToExplicitNull

`func (o *Permission) SetGrantedToExplicitNull(b bool)`

SetGrantedToExplicitNull (un)sets GrantedTo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GrantedTo value is set to nil even if false is passed
### GetInheritedFrom

`func (o *Permission) GetInheritedFrom() AnyOfmicrosoftGraphItemReference`

GetInheritedFrom returns the InheritedFrom field if non-nil, zero value otherwise.

### GetInheritedFromOk

`func (o *Permission) GetInheritedFromOk() (AnyOfmicrosoftGraphItemReference, bool)`

GetInheritedFromOk returns a tuple with the InheritedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInheritedFrom

`func (o *Permission) HasInheritedFrom() bool`

HasInheritedFrom returns a boolean if a field has been set.

### SetInheritedFrom

`func (o *Permission) SetInheritedFrom(v AnyOfmicrosoftGraphItemReference)`

SetInheritedFrom gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the InheritedFrom field.

### SetInheritedFromExplicitNull

`func (o *Permission) SetInheritedFromExplicitNull(b bool)`

SetInheritedFromExplicitNull (un)sets InheritedFrom to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InheritedFrom value is set to nil even if false is passed
### GetInvitation

`func (o *Permission) GetInvitation() AnyOfmicrosoftGraphSharingInvitation`

GetInvitation returns the Invitation field if non-nil, zero value otherwise.

### GetInvitationOk

`func (o *Permission) GetInvitationOk() (AnyOfmicrosoftGraphSharingInvitation, bool)`

GetInvitationOk returns a tuple with the Invitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitation

`func (o *Permission) HasInvitation() bool`

HasInvitation returns a boolean if a field has been set.

### SetInvitation

`func (o *Permission) SetInvitation(v AnyOfmicrosoftGraphSharingInvitation)`

SetInvitation gets a reference to the given AnyOfmicrosoftGraphSharingInvitation and assigns it to the Invitation field.

### SetInvitationExplicitNull

`func (o *Permission) SetInvitationExplicitNull(b bool)`

SetInvitationExplicitNull (un)sets Invitation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Invitation value is set to nil even if false is passed
### GetLink

`func (o *Permission) GetLink() AnyOfmicrosoftGraphSharingLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Permission) GetLinkOk() (AnyOfmicrosoftGraphSharingLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLink

`func (o *Permission) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLink

`func (o *Permission) SetLink(v AnyOfmicrosoftGraphSharingLink)`

SetLink gets a reference to the given AnyOfmicrosoftGraphSharingLink and assigns it to the Link field.

### SetLinkExplicitNull

`func (o *Permission) SetLinkExplicitNull(b bool)`

SetLinkExplicitNull (un)sets Link to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Link value is set to nil even if false is passed
### GetRoles

`func (o *Permission) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Permission) GetRolesOk() ([]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoles

`func (o *Permission) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRoles

`func (o *Permission) SetRoles(v []string)`

SetRoles gets a reference to the given []string and assigns it to the Roles field.

### GetShareId

`func (o *Permission) GetShareId() string`

GetShareId returns the ShareId field if non-nil, zero value otherwise.

### GetShareIdOk

`func (o *Permission) GetShareIdOk() (string, bool)`

GetShareIdOk returns a tuple with the ShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShareId

`func (o *Permission) HasShareId() bool`

HasShareId returns a boolean if a field has been set.

### SetShareId

`func (o *Permission) SetShareId(v string)`

SetShareId gets a reference to the given string and assigns it to the ShareId field.

### SetShareIdExplicitNull

`func (o *Permission) SetShareIdExplicitNull(b bool)`

SetShareIdExplicitNull (un)sets ShareId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ShareId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


