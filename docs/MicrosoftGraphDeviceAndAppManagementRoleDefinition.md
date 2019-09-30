# MicrosoftGraphDeviceAndAppManagementRoleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | Display Name of the Role definition. | [optional] 
**Description** | Pointer to **string** | Description of the Role definition. | [optional] 
**RolePermissions** | Pointer to [**[]AnyOfmicrosoftGraphRolePermission**](anyOf&lt;microsoft.graph.rolePermission&gt;.md) | List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission. | [optional] 
**IsBuiltIn** | Pointer to **bool** | Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition. | [optional] 
**RoleAssignments** | Pointer to [**[]MicrosoftGraphRoleAssignment**](microsoft.graph.roleAssignment.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetRolePermissions

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) GetRolePermissions() []AnyOfmicrosoftGraphRolePermission`

GetRolePermissions returns the RolePermissions field if non-nil, zero value otherwise.

### GetRolePermissionsOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) GetRolePermissionsOk() ([]AnyOfmicrosoftGraphRolePermission, bool)`

GetRolePermissionsOk returns a tuple with the RolePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRolePermissions

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) HasRolePermissions() bool`

HasRolePermissions returns a boolean if a field has been set.

### SetRolePermissions

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) SetRolePermissions(v []AnyOfmicrosoftGraphRolePermission)`

SetRolePermissions gets a reference to the given []AnyOfmicrosoftGraphRolePermission and assigns it to the RolePermissions field.

### GetIsBuiltIn

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) GetIsBuiltIn() bool`

GetIsBuiltIn returns the IsBuiltIn field if non-nil, zero value otherwise.

### GetIsBuiltInOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) GetIsBuiltInOk() (bool, bool)`

GetIsBuiltInOk returns a tuple with the IsBuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsBuiltIn

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) HasIsBuiltIn() bool`

HasIsBuiltIn returns a boolean if a field has been set.

### SetIsBuiltIn

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) SetIsBuiltIn(v bool)`

SetIsBuiltIn gets a reference to the given bool and assigns it to the IsBuiltIn field.

### GetRoleAssignments

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) GetRoleAssignments() []MicrosoftGraphRoleAssignment`

GetRoleAssignments returns the RoleAssignments field if non-nil, zero value otherwise.

### GetRoleAssignmentsOk

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) GetRoleAssignmentsOk() ([]MicrosoftGraphRoleAssignment, bool)`

GetRoleAssignmentsOk returns a tuple with the RoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoleAssignments

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) HasRoleAssignments() bool`

HasRoleAssignments returns a boolean if a field has been set.

### SetRoleAssignments

`func (o *MicrosoftGraphDeviceAndAppManagementRoleDefinition) SetRoleAssignments(v []MicrosoftGraphRoleAssignment)`

SetRoleAssignments gets a reference to the given []MicrosoftGraphRoleAssignment and assigns it to the RoleAssignments field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


