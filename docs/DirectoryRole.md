# DirectoryRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**RoleTemplateId** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 

## Methods

### GetDescription

`func (o *DirectoryRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DirectoryRole) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *DirectoryRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *DirectoryRole) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *DirectoryRole) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *DirectoryRole) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DirectoryRole) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *DirectoryRole) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *DirectoryRole) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *DirectoryRole) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetRoleTemplateId

`func (o *DirectoryRole) GetRoleTemplateId() string`

GetRoleTemplateId returns the RoleTemplateId field if non-nil, zero value otherwise.

### GetRoleTemplateIdOk

`func (o *DirectoryRole) GetRoleTemplateIdOk() (string, bool)`

GetRoleTemplateIdOk returns a tuple with the RoleTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoleTemplateId

`func (o *DirectoryRole) HasRoleTemplateId() bool`

HasRoleTemplateId returns a boolean if a field has been set.

### SetRoleTemplateId

`func (o *DirectoryRole) SetRoleTemplateId(v string)`

SetRoleTemplateId gets a reference to the given string and assigns it to the RoleTemplateId field.

### SetRoleTemplateIdExplicitNull

`func (o *DirectoryRole) SetRoleTemplateIdExplicitNull(b bool)`

SetRoleTemplateIdExplicitNull (un)sets RoleTemplateId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RoleTemplateId value is set to nil even if false is passed
### GetMembers

`func (o *DirectoryRole) GetMembers() []MicrosoftGraphDirectoryObject`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DirectoryRole) GetMembersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *DirectoryRole) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *DirectoryRole) SetMembers(v []MicrosoftGraphDirectoryObject)`

SetMembers gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the Members field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


