# RoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display or friendly name of the role Assignment. | [optional] 
**Description** | Pointer to **string** | Description of the Role Assignment. | [optional] 
**ResourceScopes** | Pointer to **[]string** | List of ids of role scope member security groups.  These are IDs from Azure Active Directory. | [optional] 
**RoleDefinition** | Pointer to [**AnyOfmicrosoftGraphRoleDefinition**](anyOf&lt;microsoft.graph.roleDefinition&gt;.md) |  | [optional] 

## Methods

### GetDisplayName

`func (o *RoleAssignment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RoleAssignment) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *RoleAssignment) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *RoleAssignment) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *RoleAssignment) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *RoleAssignment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleAssignment) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *RoleAssignment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *RoleAssignment) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *RoleAssignment) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetResourceScopes

`func (o *RoleAssignment) GetResourceScopes() []string`

GetResourceScopes returns the ResourceScopes field if non-nil, zero value otherwise.

### GetResourceScopesOk

`func (o *RoleAssignment) GetResourceScopesOk() ([]string, bool)`

GetResourceScopesOk returns a tuple with the ResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceScopes

`func (o *RoleAssignment) HasResourceScopes() bool`

HasResourceScopes returns a boolean if a field has been set.

### SetResourceScopes

`func (o *RoleAssignment) SetResourceScopes(v []string)`

SetResourceScopes gets a reference to the given []string and assigns it to the ResourceScopes field.

### GetRoleDefinition

`func (o *RoleAssignment) GetRoleDefinition() AnyOfmicrosoftGraphRoleDefinition`

GetRoleDefinition returns the RoleDefinition field if non-nil, zero value otherwise.

### GetRoleDefinitionOk

`func (o *RoleAssignment) GetRoleDefinitionOk() (AnyOfmicrosoftGraphRoleDefinition, bool)`

GetRoleDefinitionOk returns a tuple with the RoleDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoleDefinition

`func (o *RoleAssignment) HasRoleDefinition() bool`

HasRoleDefinition returns a boolean if a field has been set.

### SetRoleDefinition

`func (o *RoleAssignment) SetRoleDefinition(v AnyOfmicrosoftGraphRoleDefinition)`

SetRoleDefinition gets a reference to the given AnyOfmicrosoftGraphRoleDefinition and assigns it to the RoleDefinition field.

### SetRoleDefinitionExplicitNull

`func (o *RoleAssignment) SetRoleDefinitionExplicitNull(b bool)`

SetRoleDefinitionExplicitNull (un)sets RoleDefinition to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RoleDefinition value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


