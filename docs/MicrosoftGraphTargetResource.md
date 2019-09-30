# MicrosoftGraphTargetResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserPrincipalName** | Pointer to **string** |  | [optional] 
**GroupType** | Pointer to [**AnyOfmicrosoftGraphGroupType**](anyOf&lt;microsoft.graph.groupType&gt;.md) |  | [optional] 
**ModifiedProperties** | Pointer to [**[]AnyOfmicrosoftGraphModifiedProperty**](anyOf&lt;microsoft.graph.modifiedProperty&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphTargetResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTargetResource) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphTargetResource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphTargetResource) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### SetIdExplicitNull

`func (o *MicrosoftGraphTargetResource) SetIdExplicitNull(b bool)`

SetIdExplicitNull (un)sets Id to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Id value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphTargetResource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTargetResource) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphTargetResource) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphTargetResource) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphTargetResource) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetType

`func (o *MicrosoftGraphTargetResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphTargetResource) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *MicrosoftGraphTargetResource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *MicrosoftGraphTargetResource) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### SetTypeExplicitNull

`func (o *MicrosoftGraphTargetResource) SetTypeExplicitNull(b bool)`

SetTypeExplicitNull (un)sets Type to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Type value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *MicrosoftGraphTargetResource) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphTargetResource) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *MicrosoftGraphTargetResource) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphTargetResource) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphTargetResource) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetGroupType

`func (o *MicrosoftGraphTargetResource) GetGroupType() AnyOfmicrosoftGraphGroupType`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *MicrosoftGraphTargetResource) GetGroupTypeOk() (AnyOfmicrosoftGraphGroupType, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroupType

`func (o *MicrosoftGraphTargetResource) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### SetGroupType

`func (o *MicrosoftGraphTargetResource) SetGroupType(v AnyOfmicrosoftGraphGroupType)`

SetGroupType gets a reference to the given AnyOfmicrosoftGraphGroupType and assigns it to the GroupType field.

### SetGroupTypeExplicitNull

`func (o *MicrosoftGraphTargetResource) SetGroupTypeExplicitNull(b bool)`

SetGroupTypeExplicitNull (un)sets GroupType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GroupType value is set to nil even if false is passed
### GetModifiedProperties

`func (o *MicrosoftGraphTargetResource) GetModifiedProperties() []AnyOfmicrosoftGraphModifiedProperty`

GetModifiedProperties returns the ModifiedProperties field if non-nil, zero value otherwise.

### GetModifiedPropertiesOk

`func (o *MicrosoftGraphTargetResource) GetModifiedPropertiesOk() ([]AnyOfmicrosoftGraphModifiedProperty, bool)`

GetModifiedPropertiesOk returns a tuple with the ModifiedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModifiedProperties

`func (o *MicrosoftGraphTargetResource) HasModifiedProperties() bool`

HasModifiedProperties returns a boolean if a field has been set.

### SetModifiedProperties

`func (o *MicrosoftGraphTargetResource) SetModifiedProperties(v []AnyOfmicrosoftGraphModifiedProperty)`

SetModifiedProperties gets a reference to the given []AnyOfmicrosoftGraphModifiedProperty and assigns it to the ModifiedProperties field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


