# MicrosoftGraphMailSearchFolder

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
**IsSupported** | Pointer to **bool** |  | [optional] 
**IncludeNestedFolders** | Pointer to **bool** |  | [optional] 
**SourceFolderIds** | Pointer to **[]string** |  | [optional] 
**FilterQuery** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphMailSearchFolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMailSearchFolder) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphMailSearchFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphMailSearchFolder) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphMailSearchFolder) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphMailSearchFolder) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphMailSearchFolder) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphMailSearchFolder) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphMailSearchFolder) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetParentFolderId

`func (o *MicrosoftGraphMailSearchFolder) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MicrosoftGraphMailSearchFolder) GetParentFolderIdOk() (string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentFolderId

`func (o *MicrosoftGraphMailSearchFolder) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderId

`func (o *MicrosoftGraphMailSearchFolder) SetParentFolderId(v string)`

SetParentFolderId gets a reference to the given string and assigns it to the ParentFolderId field.

### SetParentFolderIdExplicitNull

`func (o *MicrosoftGraphMailSearchFolder) SetParentFolderIdExplicitNull(b bool)`

SetParentFolderIdExplicitNull (un)sets ParentFolderId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentFolderId value is set to nil even if false is passed
### GetChildFolderCount

`func (o *MicrosoftGraphMailSearchFolder) GetChildFolderCount() int32`

GetChildFolderCount returns the ChildFolderCount field if non-nil, zero value otherwise.

### GetChildFolderCountOk

`func (o *MicrosoftGraphMailSearchFolder) GetChildFolderCountOk() (int32, bool)`

GetChildFolderCountOk returns a tuple with the ChildFolderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildFolderCount

`func (o *MicrosoftGraphMailSearchFolder) HasChildFolderCount() bool`

HasChildFolderCount returns a boolean if a field has been set.

### SetChildFolderCount

`func (o *MicrosoftGraphMailSearchFolder) SetChildFolderCount(v int32)`

SetChildFolderCount gets a reference to the given int32 and assigns it to the ChildFolderCount field.

### SetChildFolderCountExplicitNull

`func (o *MicrosoftGraphMailSearchFolder) SetChildFolderCountExplicitNull(b bool)`

SetChildFolderCountExplicitNull (un)sets ChildFolderCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChildFolderCount value is set to nil even if false is passed
### GetUnreadItemCount

`func (o *MicrosoftGraphMailSearchFolder) GetUnreadItemCount() int32`

GetUnreadItemCount returns the UnreadItemCount field if non-nil, zero value otherwise.

### GetUnreadItemCountOk

`func (o *MicrosoftGraphMailSearchFolder) GetUnreadItemCountOk() (int32, bool)`

GetUnreadItemCountOk returns a tuple with the UnreadItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnreadItemCount

`func (o *MicrosoftGraphMailSearchFolder) HasUnreadItemCount() bool`

HasUnreadItemCount returns a boolean if a field has been set.

### SetUnreadItemCount

`func (o *MicrosoftGraphMailSearchFolder) SetUnreadItemCount(v int32)`

SetUnreadItemCount gets a reference to the given int32 and assigns it to the UnreadItemCount field.

### SetUnreadItemCountExplicitNull

`func (o *MicrosoftGraphMailSearchFolder) SetUnreadItemCountExplicitNull(b bool)`

SetUnreadItemCountExplicitNull (un)sets UnreadItemCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UnreadItemCount value is set to nil even if false is passed
### GetTotalItemCount

`func (o *MicrosoftGraphMailSearchFolder) GetTotalItemCount() int32`

GetTotalItemCount returns the TotalItemCount field if non-nil, zero value otherwise.

### GetTotalItemCountOk

`func (o *MicrosoftGraphMailSearchFolder) GetTotalItemCountOk() (int32, bool)`

GetTotalItemCountOk returns a tuple with the TotalItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalItemCount

`func (o *MicrosoftGraphMailSearchFolder) HasTotalItemCount() bool`

HasTotalItemCount returns a boolean if a field has been set.

### SetTotalItemCount

`func (o *MicrosoftGraphMailSearchFolder) SetTotalItemCount(v int32)`

SetTotalItemCount gets a reference to the given int32 and assigns it to the TotalItemCount field.

### SetTotalItemCountExplicitNull

`func (o *MicrosoftGraphMailSearchFolder) SetTotalItemCountExplicitNull(b bool)`

SetTotalItemCountExplicitNull (un)sets TotalItemCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TotalItemCount value is set to nil even if false is passed
### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphMailSearchFolder) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphMailSearchFolder) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphMailSearchFolder) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphMailSearchFolder) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphMailSearchFolder) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphMailSearchFolder) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphMailSearchFolder) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphMailSearchFolder) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetMessages

`func (o *MicrosoftGraphMailSearchFolder) GetMessages() []MicrosoftGraphMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MicrosoftGraphMailSearchFolder) GetMessagesOk() ([]MicrosoftGraphMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessages

`func (o *MicrosoftGraphMailSearchFolder) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessages

`func (o *MicrosoftGraphMailSearchFolder) SetMessages(v []MicrosoftGraphMessage)`

SetMessages gets a reference to the given []MicrosoftGraphMessage and assigns it to the Messages field.

### GetMessageRules

`func (o *MicrosoftGraphMailSearchFolder) GetMessageRules() []MicrosoftGraphMessageRule`

GetMessageRules returns the MessageRules field if non-nil, zero value otherwise.

### GetMessageRulesOk

`func (o *MicrosoftGraphMailSearchFolder) GetMessageRulesOk() ([]MicrosoftGraphMessageRule, bool)`

GetMessageRulesOk returns a tuple with the MessageRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessageRules

`func (o *MicrosoftGraphMailSearchFolder) HasMessageRules() bool`

HasMessageRules returns a boolean if a field has been set.

### SetMessageRules

`func (o *MicrosoftGraphMailSearchFolder) SetMessageRules(v []MicrosoftGraphMessageRule)`

SetMessageRules gets a reference to the given []MicrosoftGraphMessageRule and assigns it to the MessageRules field.

### GetChildFolders

`func (o *MicrosoftGraphMailSearchFolder) GetChildFolders() []MicrosoftGraphMailFolder`

GetChildFolders returns the ChildFolders field if non-nil, zero value otherwise.

### GetChildFoldersOk

`func (o *MicrosoftGraphMailSearchFolder) GetChildFoldersOk() ([]MicrosoftGraphMailFolder, bool)`

GetChildFoldersOk returns a tuple with the ChildFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildFolders

`func (o *MicrosoftGraphMailSearchFolder) HasChildFolders() bool`

HasChildFolders returns a boolean if a field has been set.

### SetChildFolders

`func (o *MicrosoftGraphMailSearchFolder) SetChildFolders(v []MicrosoftGraphMailFolder)`

SetChildFolders gets a reference to the given []MicrosoftGraphMailFolder and assigns it to the ChildFolders field.

### GetIsSupported

`func (o *MicrosoftGraphMailSearchFolder) GetIsSupported() bool`

GetIsSupported returns the IsSupported field if non-nil, zero value otherwise.

### GetIsSupportedOk

`func (o *MicrosoftGraphMailSearchFolder) GetIsSupportedOk() (bool, bool)`

GetIsSupportedOk returns a tuple with the IsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsSupported

`func (o *MicrosoftGraphMailSearchFolder) HasIsSupported() bool`

HasIsSupported returns a boolean if a field has been set.

### SetIsSupported

`func (o *MicrosoftGraphMailSearchFolder) SetIsSupported(v bool)`

SetIsSupported gets a reference to the given bool and assigns it to the IsSupported field.

### SetIsSupportedExplicitNull

`func (o *MicrosoftGraphMailSearchFolder) SetIsSupportedExplicitNull(b bool)`

SetIsSupportedExplicitNull (un)sets IsSupported to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsSupported value is set to nil even if false is passed
### GetIncludeNestedFolders

`func (o *MicrosoftGraphMailSearchFolder) GetIncludeNestedFolders() bool`

GetIncludeNestedFolders returns the IncludeNestedFolders field if non-nil, zero value otherwise.

### GetIncludeNestedFoldersOk

`func (o *MicrosoftGraphMailSearchFolder) GetIncludeNestedFoldersOk() (bool, bool)`

GetIncludeNestedFoldersOk returns a tuple with the IncludeNestedFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIncludeNestedFolders

`func (o *MicrosoftGraphMailSearchFolder) HasIncludeNestedFolders() bool`

HasIncludeNestedFolders returns a boolean if a field has been set.

### SetIncludeNestedFolders

`func (o *MicrosoftGraphMailSearchFolder) SetIncludeNestedFolders(v bool)`

SetIncludeNestedFolders gets a reference to the given bool and assigns it to the IncludeNestedFolders field.

### SetIncludeNestedFoldersExplicitNull

`func (o *MicrosoftGraphMailSearchFolder) SetIncludeNestedFoldersExplicitNull(b bool)`

SetIncludeNestedFoldersExplicitNull (un)sets IncludeNestedFolders to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IncludeNestedFolders value is set to nil even if false is passed
### GetSourceFolderIds

`func (o *MicrosoftGraphMailSearchFolder) GetSourceFolderIds() []string`

GetSourceFolderIds returns the SourceFolderIds field if non-nil, zero value otherwise.

### GetSourceFolderIdsOk

`func (o *MicrosoftGraphMailSearchFolder) GetSourceFolderIdsOk() ([]string, bool)`

GetSourceFolderIdsOk returns a tuple with the SourceFolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceFolderIds

`func (o *MicrosoftGraphMailSearchFolder) HasSourceFolderIds() bool`

HasSourceFolderIds returns a boolean if a field has been set.

### SetSourceFolderIds

`func (o *MicrosoftGraphMailSearchFolder) SetSourceFolderIds(v []string)`

SetSourceFolderIds gets a reference to the given []string and assigns it to the SourceFolderIds field.

### GetFilterQuery

`func (o *MicrosoftGraphMailSearchFolder) GetFilterQuery() string`

GetFilterQuery returns the FilterQuery field if non-nil, zero value otherwise.

### GetFilterQueryOk

`func (o *MicrosoftGraphMailSearchFolder) GetFilterQueryOk() (string, bool)`

GetFilterQueryOk returns a tuple with the FilterQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilterQuery

`func (o *MicrosoftGraphMailSearchFolder) HasFilterQuery() bool`

HasFilterQuery returns a boolean if a field has been set.

### SetFilterQuery

`func (o *MicrosoftGraphMailSearchFolder) SetFilterQuery(v string)`

SetFilterQuery gets a reference to the given string and assigns it to the FilterQuery field.

### SetFilterQueryExplicitNull

`func (o *MicrosoftGraphMailSearchFolder) SetFilterQueryExplicitNull(b bool)`

SetFilterQueryExplicitNull (un)sets FilterQuery to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FilterQuery value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


