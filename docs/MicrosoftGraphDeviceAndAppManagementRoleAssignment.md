# MicrosoftGraphDeviceAndAppManagementRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | The display or friendly name of the role Assignment. | [optional] 
**Description** | Pointer to **string** | Description of the Role Assignment. | [optional] 
**ResourceScopes** | Pointer to **[]string** | List of ids of role scope member security groups.  These are IDs from Azure Active Directory. | [optional] 
**RoleDefinition** | Pointer to [**AnyOfmicrosoftGraphRoleDefinition**](anyOf&lt;microsoft.graph.roleDefinition&gt;.md) |  | [optional] 
**Members** | Pointer to **[]string** | The list of ids of role member security groups. These are IDs from Azure Active Directory. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetResourceScopes

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetResourceScopes() []string`

GetResourceScopes returns the ResourceScopes field if non-nil, zero value otherwise.

### GetResourceScopesOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetResourceScopesOk() ([]string, bool)`

GetResourceScopesOk returns a tuple with the ResourceScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceScopes

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) HasResourceScopes() bool`

HasResourceScopes returns a boolean if a field has been set.

### SetResourceScopes

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetResourceScopes(v []string)`

SetResourceScopes gets a reference to the given []string and assigns it to the ResourceScopes field.

### GetRoleDefinition

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetRoleDefinition() AnyOfmicrosoftGraphRoleDefinition`

GetRoleDefinition returns the RoleDefinition field if non-nil, zero value otherwise.

### GetRoleDefinitionOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetRoleDefinitionOk() (AnyOfmicrosoftGraphRoleDefinition, bool)`

GetRoleDefinitionOk returns a tuple with the RoleDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoleDefinition

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) HasRoleDefinition() bool`

HasRoleDefinition returns a boolean if a field has been set.

### SetRoleDefinition

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetRoleDefinition(v AnyOfmicrosoftGraphRoleDefinition)`

SetRoleDefinition gets a reference to the given AnyOfmicrosoftGraphRoleDefinition and assigns it to the RoleDefinition field.

### SetRoleDefinitionExplicitNull

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetRoleDefinitionExplicitNull(b bool)`

SetRoleDefinitionExplicitNull (un)sets RoleDefinition to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RoleDefinition value is set to nil even if false is passed
### GetMembers

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) GetMembersOk() ([]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *MicrosoftGraphDeviceAndAppManagementRoleAssignment) SetMembers(v []string)`

SetMembers gets a reference to the given []string and assigns it to the Members field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


