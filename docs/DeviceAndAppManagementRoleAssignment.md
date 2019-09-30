# DeviceAndAppManagementRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to **[]string** | The list of ids of role member security groups. These are IDs from Azure Active Directory. | [optional] 

## Methods

### GetMembers

`func (o *DeviceAndAppManagementRoleAssignment) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DeviceAndAppManagementRoleAssignment) GetMembersOk() ([]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *DeviceAndAppManagementRoleAssignment) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *DeviceAndAppManagementRoleAssignment) SetMembers(v []string)`

SetMembers gets a reference to the given []string and assigns it to the Members field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


