# MicrosoftGraphRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | The display or friendly name of the role Assignment. | [optional] 
**Description** | Pointer to **string** | Description of the Role Assignment. | [optional] 
**ResourceScopes** | Pointer to **[]string** | List of ids of role scope member security groups.  These are IDs from Azure Active Directory. | [optional] 
**RoleDefinition** | Pointer to [**AnyOfmicrosoftGraphRoleDefinition**](anyOf&lt;microsoft.graph.roleDefinition&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphRoleAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphRoleAssignment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphRoleAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphRoleAssignment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphRoleAssignment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphRoleAssignment) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphRoleAssignment) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphRoleAssignment) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphRoleAssignment) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphRoleAssignment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphRoleAssignment) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphRoleAssignment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphRoleAssignment) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphRoleAssignment) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetResourceScopes

`func (o *MicrosoftGraphRoleAssignment) GetResourceScopes() []string`

GetResourceScopes returns the ResourceScopes field if non-nil, zero value otherwise.

### GetResourceScopesOk

`func (o *MicrosoftGraphRoleAssignment) GetResourceScopesOk() ([]string, bool)`

GetResourceScopesOk returns a tuple with the ResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceScopes

`func (o *MicrosoftGraphRoleAssignment) HasResourceScopes() bool`

HasResourceScopes returns a boolean if a field has been set.

### SetResourceScopes

`func (o *MicrosoftGraphRoleAssignment) SetResourceScopes(v []string)`

SetResourceScopes gets a reference to the given []string and assigns it to the ResourceScopes field.

### GetRoleDefinition

`func (o *MicrosoftGraphRoleAssignment) GetRoleDefinition() AnyOfmicrosoftGraphRoleDefinition`

GetRoleDefinition returns the RoleDefinition field if non-nil, zero value otherwise.

### GetRoleDefinitionOk

`func (o *MicrosoftGraphRoleAssignment) GetRoleDefinitionOk() (AnyOfmicrosoftGraphRoleDefinition, bool)`

GetRoleDefinitionOk returns a tuple with the RoleDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoleDefinition

`func (o *MicrosoftGraphRoleAssignment) HasRoleDefinition() bool`

HasRoleDefinition returns a boolean if a field has been set.

### SetRoleDefinition

`func (o *MicrosoftGraphRoleAssignment) SetRoleDefinition(v AnyOfmicrosoftGraphRoleDefinition)`

SetRoleDefinition gets a reference to the given AnyOfmicrosoftGraphRoleDefinition and assigns it to the RoleDefinition field.

### SetRoleDefinitionExplicitNull

`func (o *MicrosoftGraphRoleAssignment) SetRoleDefinitionExplicitNull(b bool)`

SetRoleDefinitionExplicitNull (un)sets RoleDefinition to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RoleDefinition value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


