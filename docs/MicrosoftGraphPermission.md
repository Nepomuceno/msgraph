# MicrosoftGraphPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GrantedTo** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**InheritedFrom** | Pointer to [**AnyOfmicrosoftGraphItemReference**](anyOf&lt;microsoft.graph.itemReference&gt;.md) |  | [optional] 
**Invitation** | Pointer to [**AnyOfmicrosoftGraphSharingInvitation**](anyOf&lt;microsoft.graph.sharingInvitation&gt;.md) |  | [optional] 
**Link** | Pointer to [**AnyOfmicrosoftGraphSharingLink**](anyOf&lt;microsoft.graph.sharingLink&gt;.md) |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**ShareId** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphPermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPermission) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphPermission) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetGrantedTo

`func (o *MicrosoftGraphPermission) GetGrantedTo() AnyOfmicrosoftGraphIdentitySet`

GetGrantedTo returns the GrantedTo field if non-nil, zero value otherwise.

### GetGrantedToOk

`func (o *MicrosoftGraphPermission) GetGrantedToOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetGrantedToOk returns a tuple with the GrantedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGrantedTo

`func (o *MicrosoftGraphPermission) HasGrantedTo() bool`

HasGrantedTo returns a boolean if a field has been set.

### SetGrantedTo

`func (o *MicrosoftGraphPermission) SetGrantedTo(v AnyOfmicrosoftGraphIdentitySet)`

SetGrantedTo gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the GrantedTo field.

### SetGrantedToExplicitNull

`func (o *MicrosoftGraphPermission) SetGrantedToExplicitNull(b bool)`

SetGrantedToExplicitNull (un)sets GrantedTo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GrantedTo value is set to nil even if false is passed
### GetInheritedFrom

`func (o *MicrosoftGraphPermission) GetInheritedFrom() AnyOfmicrosoftGraphItemReference`

GetInheritedFrom returns the InheritedFrom field if non-nil, zero value otherwise.

### GetInheritedFromOk

`func (o *MicrosoftGraphPermission) GetInheritedFromOk() (AnyOfmicrosoftGraphItemReference, bool)`

GetInheritedFromOk returns a tuple with the InheritedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInheritedFrom

`func (o *MicrosoftGraphPermission) HasInheritedFrom() bool`

HasInheritedFrom returns a boolean if a field has been set.

### SetInheritedFrom

`func (o *MicrosoftGraphPermission) SetInheritedFrom(v AnyOfmicrosoftGraphItemReference)`

SetInheritedFrom gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the InheritedFrom field.

### SetInheritedFromExplicitNull

`func (o *MicrosoftGraphPermission) SetInheritedFromExplicitNull(b bool)`

SetInheritedFromExplicitNull (un)sets InheritedFrom to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InheritedFrom value is set to nil even if false is passed
### GetInvitation

`func (o *MicrosoftGraphPermission) GetInvitation() AnyOfmicrosoftGraphSharingInvitation`

GetInvitation returns the Invitation field if non-nil, zero value otherwise.

### GetInvitationOk

`func (o *MicrosoftGraphPermission) GetInvitationOk() (AnyOfmicrosoftGraphSharingInvitation, bool)`

GetInvitationOk returns a tuple with the Invitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInvitation

`func (o *MicrosoftGraphPermission) HasInvitation() bool`

HasInvitation returns a boolean if a field has been set.

### SetInvitation

`func (o *MicrosoftGraphPermission) SetInvitation(v AnyOfmicrosoftGraphSharingInvitation)`

SetInvitation gets a reference to the given AnyOfmicrosoftGraphSharingInvitation and assigns it to the Invitation field.

### SetInvitationExplicitNull

`func (o *MicrosoftGraphPermission) SetInvitationExplicitNull(b bool)`

SetInvitationExplicitNull (un)sets Invitation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Invitation value is set to nil even if false is passed
### GetLink

`func (o *MicrosoftGraphPermission) GetLink() AnyOfmicrosoftGraphSharingLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *MicrosoftGraphPermission) GetLinkOk() (AnyOfmicrosoftGraphSharingLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLink

`func (o *MicrosoftGraphPermission) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLink

`func (o *MicrosoftGraphPermission) SetLink(v AnyOfmicrosoftGraphSharingLink)`

SetLink gets a reference to the given AnyOfmicrosoftGraphSharingLink and assigns it to the Link field.

### SetLinkExplicitNull

`func (o *MicrosoftGraphPermission) SetLinkExplicitNull(b bool)`

SetLinkExplicitNull (un)sets Link to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Link value is set to nil even if false is passed
### GetRoles

`func (o *MicrosoftGraphPermission) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MicrosoftGraphPermission) GetRolesOk() ([]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoles

`func (o *MicrosoftGraphPermission) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRoles

`func (o *MicrosoftGraphPermission) SetRoles(v []string)`

SetRoles gets a reference to the given []string and assigns it to the Roles field.

### GetShareId

`func (o *MicrosoftGraphPermission) GetShareId() string`

GetShareId returns the ShareId field if non-nil, zero value otherwise.

### GetShareIdOk

`func (o *MicrosoftGraphPermission) GetShareIdOk() (string, bool)`

GetShareIdOk returns a tuple with the ShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShareId

`func (o *MicrosoftGraphPermission) HasShareId() bool`

HasShareId returns a boolean if a field has been set.

### SetShareId

`func (o *MicrosoftGraphPermission) SetShareId(v string)`

SetShareId gets a reference to the given string and assigns it to the ShareId field.

### SetShareIdExplicitNull

`func (o *MicrosoftGraphPermission) SetShareIdExplicitNull(b bool)`

SetShareIdExplicitNull (un)sets ShareId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ShareId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


