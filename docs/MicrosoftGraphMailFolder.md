# MicrosoftGraphMailFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ParentFolderId** | Pointer to **string** |  | [optional] 
**ChildFolderCount** | Pointer to **int32** |  | [optional] 
**UnreadItemCount** | Pointer to **int32** |  | [optional] 
**TotalItemCount** | Pointer to **int32** |  | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md) |  | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md) |  | [optional] 
**Messages** | Pointer to [**[]MicrosoftGraphMessage**](microsoft.graph.message.md) |  | [optional] 
**MessageRules** | Pointer to [**[]MicrosoftGraphMessageRule**](microsoft.graph.messageRule.md) |  | [optional] 
**ChildFolders** | Pointer to [**[]MicrosoftGraphMailFolder**](microsoft.graph.mailFolder.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphMailFolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMailFolder) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphMailFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphMailFolder) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphMailFolder) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphMailFolder) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphMailFolder) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphMailFolder) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphMailFolder) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetParentFolderId

`func (o *MicrosoftGraphMailFolder) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MicrosoftGraphMailFolder) GetParentFolderIdOk() (string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentFolderId

`func (o *MicrosoftGraphMailFolder) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderId

`func (o *MicrosoftGraphMailFolder) SetParentFolderId(v string)`

SetParentFolderId gets a reference to the given string and assigns it to the ParentFolderId field.

### SetParentFolderIdExplicitNull

`func (o *MicrosoftGraphMailFolder) SetParentFolderIdExplicitNull(b bool)`

SetParentFolderIdExplicitNull (un)sets ParentFolderId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentFolderId value is set to nil even if false is passed
### GetChildFolderCount

`func (o *MicrosoftGraphMailFolder) GetChildFolderCount() int32`

GetChildFolderCount returns the ChildFolderCount field if non-nil, zero value otherwise.

### GetChildFolderCountOk

`func (o *MicrosoftGraphMailFolder) GetChildFolderCountOk() (int32, bool)`

GetChildFolderCountOk returns a tuple with the ChildFolderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildFolderCount

`func (o *MicrosoftGraphMailFolder) HasChildFolderCount() bool`

HasChildFolderCount returns a boolean if a field has been set.

### SetChildFolderCount

`func (o *MicrosoftGraphMailFolder) SetChildFolderCount(v int32)`

SetChildFolderCount gets a reference to the given int32 and assigns it to the ChildFolderCount field.

### SetChildFolderCountExplicitNull

`func (o *MicrosoftGraphMailFolder) SetChildFolderCountExplicitNull(b bool)`

SetChildFolderCountExplicitNull (un)sets ChildFolderCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChildFolderCount value is set to nil even if false is passed
### GetUnreadItemCount

`func (o *MicrosoftGraphMailFolder) GetUnreadItemCount() int32`

GetUnreadItemCount returns the UnreadItemCount field if non-nil, zero value otherwise.

### GetUnreadItemCountOk

`func (o *MicrosoftGraphMailFolder) GetUnreadItemCountOk() (int32, bool)`

GetUnreadItemCountOk returns a tuple with the UnreadItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnreadItemCount

`func (o *MicrosoftGraphMailFolder) HasUnreadItemCount() bool`

HasUnreadItemCount returns a boolean if a field has been set.

### SetUnreadItemCount

`func (o *MicrosoftGraphMailFolder) SetUnreadItemCount(v int32)`

SetUnreadItemCount gets a reference to the given int32 and assigns it to the UnreadItemCount field.

### SetUnreadItemCountExplicitNull

`func (o *MicrosoftGraphMailFolder) SetUnreadItemCountExplicitNull(b bool)`

SetUnreadItemCountExplicitNull (un)sets UnreadItemCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UnreadItemCount value is set to nil even if false is passed
### GetTotalItemCount

`func (o *MicrosoftGraphMailFolder) GetTotalItemCount() int32`

GetTotalItemCount returns the TotalItemCount field if non-nil, zero value otherwise.

### GetTotalItemCountOk

`func (o *MicrosoftGraphMailFolder) GetTotalItemCountOk() (int32, bool)`

GetTotalItemCountOk returns a tuple with the TotalItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalItemCount

`func (o *MicrosoftGraphMailFolder) HasTotalItemCount() bool`

HasTotalItemCount returns a boolean if a field has been set.

### SetTotalItemCount

`func (o *MicrosoftGraphMailFolder) SetTotalItemCount(v int32)`

SetTotalItemCount gets a reference to the given int32 and assigns it to the TotalItemCount field.

### SetTotalItemCountExplicitNull

`func (o *MicrosoftGraphMailFolder) SetTotalItemCountExplicitNull(b bool)`

SetTotalItemCountExplicitNull (un)sets TotalItemCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TotalItemCount value is set to nil even if false is passed
### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphMailFolder) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphMailFolder) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphMailFolder) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphMailFolder) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphMailFolder) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphMailFolder) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphMailFolder) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphMailFolder) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetMessages

`func (o *MicrosoftGraphMailFolder) GetMessages() []MicrosoftGraphMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MicrosoftGraphMailFolder) GetMessagesOk() ([]MicrosoftGraphMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessages

`func (o *MicrosoftGraphMailFolder) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessages

`func (o *MicrosoftGraphMailFolder) SetMessages(v []MicrosoftGraphMessage)`

SetMessages gets a reference to the given []MicrosoftGraphMessage and assigns it to the Messages field.

### GetMessageRules

`func (o *MicrosoftGraphMailFolder) GetMessageRules() []MicrosoftGraphMessageRule`

GetMessageRules returns the MessageRules field if non-nil, zero value otherwise.

### GetMessageRulesOk

`func (o *MicrosoftGraphMailFolder) GetMessageRulesOk() ([]MicrosoftGraphMessageRule, bool)`

GetMessageRulesOk returns a tuple with the MessageRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessageRules

`func (o *MicrosoftGraphMailFolder) HasMessageRules() bool`

HasMessageRules returns a boolean if a field has been set.

### SetMessageRules

`func (o *MicrosoftGraphMailFolder) SetMessageRules(v []MicrosoftGraphMessageRule)`

SetMessageRules gets a reference to the given []MicrosoftGraphMessageRule and assigns it to the MessageRules field.

### GetChildFolders

`func (o *MicrosoftGraphMailFolder) GetChildFolders() []MicrosoftGraphMailFolder`

GetChildFolders returns the ChildFolders field if non-nil, zero value otherwise.

### GetChildFoldersOk

`func (o *MicrosoftGraphMailFolder) GetChildFoldersOk() ([]MicrosoftGraphMailFolder, bool)`

GetChildFoldersOk returns a tuple with the ChildFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildFolders

`func (o *MicrosoftGraphMailFolder) HasChildFolders() bool`

HasChildFolders returns a boolean if a field has been set.

### SetChildFolders

`func (o *MicrosoftGraphMailFolder) SetChildFolders(v []MicrosoftGraphMailFolder)`

SetChildFolders gets a reference to the given []MicrosoftGraphMailFolder and assigns it to the ChildFolders field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


