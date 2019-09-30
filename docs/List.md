# List

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**List** | Pointer to [**AnyOfmicrosoftGraphListInfo**](anyOf&lt;microsoft.graph.listInfo&gt;.md) |  | [optional] 
**SharepointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) |  | [optional] 
**System** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Columns** | Pointer to [**[]MicrosoftGraphColumnDefinition**](microsoft.graph.columnDefinition.md) |  | [optional] 
**ContentTypes** | Pointer to [**[]MicrosoftGraphContentType**](microsoft.graph.contentType.md) |  | [optional] 
**Drive** | Pointer to [**AnyOfmicrosoftGraphDrive**](anyOf&lt;microsoft.graph.drive&gt;.md) |  | [optional] 
**Items** | Pointer to [**[]MicrosoftGraphListItem**](microsoft.graph.listItem.md) |  | [optional] 

## Methods

### GetDisplayName

`func (o *List) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *List) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *List) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *List) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *List) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetList

`func (o *List) GetList() AnyOfmicrosoftGraphListInfo`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *List) GetListOk() (AnyOfmicrosoftGraphListInfo, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasList

`func (o *List) HasList() bool`

HasList returns a boolean if a field has been set.

### SetList

`func (o *List) SetList(v AnyOfmicrosoftGraphListInfo)`

SetList gets a reference to the given AnyOfmicrosoftGraphListInfo and assigns it to the List field.

### SetListExplicitNull

`func (o *List) SetListExplicitNull(b bool)`

SetListExplicitNull (un)sets List to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The List value is set to nil even if false is passed
### GetSharepointIds

`func (o *List) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *List) GetSharepointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharepointIds

`func (o *List) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIds

`func (o *List) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.

### SetSharepointIdsExplicitNull

`func (o *List) SetSharepointIdsExplicitNull(b bool)`

SetSharepointIdsExplicitNull (un)sets SharepointIds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharepointIds value is set to nil even if false is passed
### GetSystem

`func (o *List) GetSystem() AnyOfobject`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *List) GetSystemOk() (AnyOfobject, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSystem

`func (o *List) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystem

`func (o *List) SetSystem(v AnyOfobject)`

SetSystem gets a reference to the given AnyOfobject and assigns it to the System field.

### SetSystemExplicitNull

`func (o *List) SetSystemExplicitNull(b bool)`

SetSystemExplicitNull (un)sets System to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The System value is set to nil even if false is passed
### GetColumns

`func (o *List) GetColumns() []MicrosoftGraphColumnDefinition`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *List) GetColumnsOk() ([]MicrosoftGraphColumnDefinition, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumns

`func (o *List) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumns

`func (o *List) SetColumns(v []MicrosoftGraphColumnDefinition)`

SetColumns gets a reference to the given []MicrosoftGraphColumnDefinition and assigns it to the Columns field.

### GetContentTypes

`func (o *List) GetContentTypes() []MicrosoftGraphContentType`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *List) GetContentTypesOk() ([]MicrosoftGraphContentType, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentTypes

`func (o *List) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### SetContentTypes

`func (o *List) SetContentTypes(v []MicrosoftGraphContentType)`

SetContentTypes gets a reference to the given []MicrosoftGraphContentType and assigns it to the ContentTypes field.

### GetDrive

`func (o *List) GetDrive() AnyOfmicrosoftGraphDrive`

GetDrive returns the Drive field if non-nil, zero value otherwise.

### GetDriveOk

`func (o *List) GetDriveOk() (AnyOfmicrosoftGraphDrive, bool)`

GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDrive

`func (o *List) HasDrive() bool`

HasDrive returns a boolean if a field has been set.

### SetDrive

`func (o *List) SetDrive(v AnyOfmicrosoftGraphDrive)`

SetDrive gets a reference to the given AnyOfmicrosoftGraphDrive and assigns it to the Drive field.

### SetDriveExplicitNull

`func (o *List) SetDriveExplicitNull(b bool)`

SetDriveExplicitNull (un)sets Drive to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Drive value is set to nil even if false is passed
### GetItems

`func (o *List) GetItems() []MicrosoftGraphListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *List) GetItemsOk() ([]MicrosoftGraphListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItems

`func (o *List) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItems

`func (o *List) SetItems(v []MicrosoftGraphListItem)`

SetItems gets a reference to the given []MicrosoftGraphListItem and assigns it to the Items field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


