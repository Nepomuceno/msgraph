# MailFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetDisplayName

`func (o *MailFolder) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MailFolder) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MailFolder) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MailFolder) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MailFolder) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetParentFolderId

`func (o *MailFolder) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MailFolder) GetParentFolderIdOk() (string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentFolderId

`func (o *MailFolder) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderId

`func (o *MailFolder) SetParentFolderId(v string)`

SetParentFolderId gets a reference to the given string and assigns it to the ParentFolderId field.

### SetParentFolderIdExplicitNull

`func (o *MailFolder) SetParentFolderIdExplicitNull(b bool)`

SetParentFolderIdExplicitNull (un)sets ParentFolderId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentFolderId value is set to nil even if false is passed
### GetChildFolderCount

`func (o *MailFolder) GetChildFolderCount() int32`

GetChildFolderCount returns the ChildFolderCount field if non-nil, zero value otherwise.

### GetChildFolderCountOk

`func (o *MailFolder) GetChildFolderCountOk() (int32, bool)`

GetChildFolderCountOk returns a tuple with the ChildFolderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildFolderCount

`func (o *MailFolder) HasChildFolderCount() bool`

HasChildFolderCount returns a boolean if a field has been set.

### SetChildFolderCount

`func (o *MailFolder) SetChildFolderCount(v int32)`

SetChildFolderCount gets a reference to the given int32 and assigns it to the ChildFolderCount field.

### SetChildFolderCountExplicitNull

`func (o *MailFolder) SetChildFolderCountExplicitNull(b bool)`

SetChildFolderCountExplicitNull (un)sets ChildFolderCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChildFolderCount value is set to nil even if false is passed
### GetUnreadItemCount

`func (o *MailFolder) GetUnreadItemCount() int32`

GetUnreadItemCount returns the UnreadItemCount field if non-nil, zero value otherwise.

### GetUnreadItemCountOk

`func (o *MailFolder) GetUnreadItemCountOk() (int32, bool)`

GetUnreadItemCountOk returns a tuple with the UnreadItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnreadItemCount

`func (o *MailFolder) HasUnreadItemCount() bool`

HasUnreadItemCount returns a boolean if a field has been set.

### SetUnreadItemCount

`func (o *MailFolder) SetUnreadItemCount(v int32)`

SetUnreadItemCount gets a reference to the given int32 and assigns it to the UnreadItemCount field.

### SetUnreadItemCountExplicitNull

`func (o *MailFolder) SetUnreadItemCountExplicitNull(b bool)`

SetUnreadItemCountExplicitNull (un)sets UnreadItemCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UnreadItemCount value is set to nil even if false is passed
### GetTotalItemCount

`func (o *MailFolder) GetTotalItemCount() int32`

GetTotalItemCount returns the TotalItemCount field if non-nil, zero value otherwise.

### GetTotalItemCountOk

`func (o *MailFolder) GetTotalItemCountOk() (int32, bool)`

GetTotalItemCountOk returns a tuple with the TotalItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalItemCount

`func (o *MailFolder) HasTotalItemCount() bool`

HasTotalItemCount returns a boolean if a field has been set.

### SetTotalItemCount

`func (o *MailFolder) SetTotalItemCount(v int32)`

SetTotalItemCount gets a reference to the given int32 and assigns it to the TotalItemCount field.

### SetTotalItemCountExplicitNull

`func (o *MailFolder) SetTotalItemCountExplicitNull(b bool)`

SetTotalItemCountExplicitNull (un)sets TotalItemCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TotalItemCount value is set to nil even if false is passed
### GetSingleValueExtendedProperties

`func (o *MailFolder) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MailFolder) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *MailFolder) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *MailFolder) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *MailFolder) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MailFolder) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *MailFolder) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *MailFolder) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetMessages

`func (o *MailFolder) GetMessages() []MicrosoftGraphMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MailFolder) GetMessagesOk() ([]MicrosoftGraphMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessages

`func (o *MailFolder) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessages

`func (o *MailFolder) SetMessages(v []MicrosoftGraphMessage)`

SetMessages gets a reference to the given []MicrosoftGraphMessage and assigns it to the Messages field.

### GetMessageRules

`func (o *MailFolder) GetMessageRules() []MicrosoftGraphMessageRule`

GetMessageRules returns the MessageRules field if non-nil, zero value otherwise.

### GetMessageRulesOk

`func (o *MailFolder) GetMessageRulesOk() ([]MicrosoftGraphMessageRule, bool)`

GetMessageRulesOk returns a tuple with the MessageRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessageRules

`func (o *MailFolder) HasMessageRules() bool`

HasMessageRules returns a boolean if a field has been set.

### SetMessageRules

`func (o *MailFolder) SetMessageRules(v []MicrosoftGraphMessageRule)`

SetMessageRules gets a reference to the given []MicrosoftGraphMessageRule and assigns it to the MessageRules field.

### GetChildFolders

`func (o *MailFolder) GetChildFolders() []MicrosoftGraphMailFolder`

GetChildFolders returns the ChildFolders field if non-nil, zero value otherwise.

### GetChildFoldersOk

`func (o *MailFolder) GetChildFoldersOk() ([]MicrosoftGraphMailFolder, bool)`

GetChildFoldersOk returns a tuple with the ChildFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildFolders

`func (o *MailFolder) HasChildFolders() bool`

HasChildFolders returns a boolean if a field has been set.

### SetChildFolders

`func (o *MailFolder) SetChildFolders(v []MicrosoftGraphMailFolder)`

SetChildFolders gets a reference to the given []MicrosoftGraphMailFolder and assigns it to the ChildFolders field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


