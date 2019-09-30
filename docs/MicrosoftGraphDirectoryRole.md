# MicrosoftGraphDirectoryRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DeletedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**RoleTemplateId** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDirectoryRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDirectoryRole) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDirectoryRole) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDirectoryRole) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDeletedDateTime

`func (o *MicrosoftGraphDirectoryRole) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphDirectoryRole) GetDeletedDateTimeOk() (time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeletedDateTime

`func (o *MicrosoftGraphDirectoryRole) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphDirectoryRole) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime gets a reference to the given time.Time and assigns it to the DeletedDateTime field.

### SetDeletedDateTimeExplicitNull

`func (o *MicrosoftGraphDirectoryRole) SetDeletedDateTimeExplicitNull(b bool)`

SetDeletedDateTimeExplicitNull (un)sets DeletedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeletedDateTime value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphDirectoryRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDirectoryRole) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphDirectoryRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphDirectoryRole) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphDirectoryRole) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphDirectoryRole) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDirectoryRole) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphDirectoryRole) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphDirectoryRole) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphDirectoryRole) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetRoleTemplateId

`func (o *MicrosoftGraphDirectoryRole) GetRoleTemplateId() string`

GetRoleTemplateId returns the RoleTemplateId field if non-nil, zero value otherwise.

### GetRoleTemplateIdOk

`func (o *MicrosoftGraphDirectoryRole) GetRoleTemplateIdOk() (string, bool)`

GetRoleTemplateIdOk returns a tuple with the RoleTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoleTemplateId

`func (o *MicrosoftGraphDirectoryRole) HasRoleTemplateId() bool`

HasRoleTemplateId returns a boolean if a field has been set.

### SetRoleTemplateId

`func (o *MicrosoftGraphDirectoryRole) SetRoleTemplateId(v string)`

SetRoleTemplateId gets a reference to the given string and assigns it to the RoleTemplateId field.

### SetRoleTemplateIdExplicitNull

`func (o *MicrosoftGraphDirectoryRole) SetRoleTemplateIdExplicitNull(b bool)`

SetRoleTemplateIdExplicitNull (un)sets RoleTemplateId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RoleTemplateId value is set to nil even if false is passed
### GetMembers

`func (o *MicrosoftGraphDirectoryRole) GetMembers() []MicrosoftGraphDirectoryObject`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MicrosoftGraphDirectoryRole) GetMembersOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *MicrosoftGraphDirectoryRole) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *MicrosoftGraphDirectoryRole) SetMembers(v []MicrosoftGraphDirectoryObject)`

SetMembers gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the Members field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


