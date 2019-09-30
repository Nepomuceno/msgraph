# MicrosoftGraphContactFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ParentFolderId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md) |  | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md) |  | [optional] 
**Contacts** | Pointer to [**[]MicrosoftGraphContact**](microsoft.graph.contact.md) |  | [optional] 
**ChildFolders** | Pointer to [**[]MicrosoftGraphContactFolder**](microsoft.graph.contactFolder.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphContactFolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphContactFolder) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphContactFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphContactFolder) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetParentFolderId

`func (o *MicrosoftGraphContactFolder) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MicrosoftGraphContactFolder) GetParentFolderIdOk() (string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentFolderId

`func (o *MicrosoftGraphContactFolder) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderId

`func (o *MicrosoftGraphContactFolder) SetParentFolderId(v string)`

SetParentFolderId gets a reference to the given string and assigns it to the ParentFolderId field.

### SetParentFolderIdExplicitNull

`func (o *MicrosoftGraphContactFolder) SetParentFolderIdExplicitNull(b bool)`

SetParentFolderIdExplicitNull (un)sets ParentFolderId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentFolderId value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphContactFolder) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphContactFolder) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphContactFolder) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphContactFolder) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphContactFolder) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphContactFolder) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphContactFolder) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphContactFolder) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphContactFolder) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphContactFolder) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphContactFolder) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphContactFolder) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphContactFolder) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetContacts

`func (o *MicrosoftGraphContactFolder) GetContacts() []MicrosoftGraphContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *MicrosoftGraphContactFolder) GetContactsOk() ([]MicrosoftGraphContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContacts

`func (o *MicrosoftGraphContactFolder) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContacts

`func (o *MicrosoftGraphContactFolder) SetContacts(v []MicrosoftGraphContact)`

SetContacts gets a reference to the given []MicrosoftGraphContact and assigns it to the Contacts field.

### GetChildFolders

`func (o *MicrosoftGraphContactFolder) GetChildFolders() []MicrosoftGraphContactFolder`

GetChildFolders returns the ChildFolders field if non-nil, zero value otherwise.

### GetChildFoldersOk

`func (o *MicrosoftGraphContactFolder) GetChildFoldersOk() ([]MicrosoftGraphContactFolder, bool)`

GetChildFoldersOk returns a tuple with the ChildFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildFolders

`func (o *MicrosoftGraphContactFolder) HasChildFolders() bool`

HasChildFolders returns a boolean if a field has been set.

### SetChildFolders

`func (o *MicrosoftGraphContactFolder) SetChildFolders(v []MicrosoftGraphContactFolder)`

SetChildFolders gets a reference to the given []MicrosoftGraphContactFolder and assigns it to the ChildFolders field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


