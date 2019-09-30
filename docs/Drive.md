# Drive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveType** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**Quota** | Pointer to [**AnyOfmicrosoftGraphQuota**](anyOf&lt;microsoft.graph.quota&gt;.md) |  | [optional] 
**SharePointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) |  | [optional] 
**System** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Items** | Pointer to [**[]MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md) |  | [optional] 
**List** | Pointer to [**AnyOfmicrosoftGraphList**](anyOf&lt;microsoft.graph.list&gt;.md) |  | [optional] 
**Root** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) |  | [optional] 
**Special** | Pointer to [**[]MicrosoftGraphDriveItem**](microsoft.graph.driveItem.md) |  | [optional] 

## Methods

### GetDriveType

`func (o *Drive) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *Drive) GetDriveTypeOk() (string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDriveType

`func (o *Drive) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### SetDriveType

`func (o *Drive) SetDriveType(v string)`

SetDriveType gets a reference to the given string and assigns it to the DriveType field.

### SetDriveTypeExplicitNull

`func (o *Drive) SetDriveTypeExplicitNull(b bool)`

SetDriveTypeExplicitNull (un)sets DriveType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DriveType value is set to nil even if false is passed
### GetOwner

`func (o *Drive) GetOwner() AnyOfmicrosoftGraphIdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Drive) GetOwnerOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *Drive) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *Drive) SetOwner(v AnyOfmicrosoftGraphIdentitySet)`

SetOwner gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *Drive) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetQuota

`func (o *Drive) GetQuota() AnyOfmicrosoftGraphQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Drive) GetQuotaOk() (AnyOfmicrosoftGraphQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuota

`func (o *Drive) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuota

`func (o *Drive) SetQuota(v AnyOfmicrosoftGraphQuota)`

SetQuota gets a reference to the given AnyOfmicrosoftGraphQuota and assigns it to the Quota field.

### SetQuotaExplicitNull

`func (o *Drive) SetQuotaExplicitNull(b bool)`

SetQuotaExplicitNull (un)sets Quota to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Quota value is set to nil even if false is passed
### GetSharePointIds

`func (o *Drive) GetSharePointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharePointIds returns the SharePointIds field if non-nil, zero value otherwise.

### GetSharePointIdsOk

`func (o *Drive) GetSharePointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharePointIdsOk returns a tuple with the SharePointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharePointIds

`func (o *Drive) HasSharePointIds() bool`

HasSharePointIds returns a boolean if a field has been set.

### SetSharePointIds

`func (o *Drive) SetSharePointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharePointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharePointIds field.

### SetSharePointIdsExplicitNull

`func (o *Drive) SetSharePointIdsExplicitNull(b bool)`

SetSharePointIdsExplicitNull (un)sets SharePointIds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharePointIds value is set to nil even if false is passed
### GetSystem

`func (o *Drive) GetSystem() AnyOfobject`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Drive) GetSystemOk() (AnyOfobject, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSystem

`func (o *Drive) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystem

`func (o *Drive) SetSystem(v AnyOfobject)`

SetSystem gets a reference to the given AnyOfobject and assigns it to the System field.

### SetSystemExplicitNull

`func (o *Drive) SetSystemExplicitNull(b bool)`

SetSystemExplicitNull (un)sets System to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The System value is set to nil even if false is passed
### GetItems

`func (o *Drive) GetItems() []MicrosoftGraphDriveItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Drive) GetItemsOk() ([]MicrosoftGraphDriveItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItems

`func (o *Drive) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItems

`func (o *Drive) SetItems(v []MicrosoftGraphDriveItem)`

SetItems gets a reference to the given []MicrosoftGraphDriveItem and assigns it to the Items field.

### GetList

`func (o *Drive) GetList() AnyOfmicrosoftGraphList`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *Drive) GetListOk() (AnyOfmicrosoftGraphList, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasList

`func (o *Drive) HasList() bool`

HasList returns a boolean if a field has been set.

### SetList

`func (o *Drive) SetList(v AnyOfmicrosoftGraphList)`

SetList gets a reference to the given AnyOfmicrosoftGraphList and assigns it to the List field.

### SetListExplicitNull

`func (o *Drive) SetListExplicitNull(b bool)`

SetListExplicitNull (un)sets List to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The List value is set to nil even if false is passed
### GetRoot

`func (o *Drive) GetRoot() AnyOfmicrosoftGraphDriveItem`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *Drive) GetRootOk() (AnyOfmicrosoftGraphDriveItem, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoot

`func (o *Drive) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRoot

`func (o *Drive) SetRoot(v AnyOfmicrosoftGraphDriveItem)`

SetRoot gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the Root field.

### SetRootExplicitNull

`func (o *Drive) SetRootExplicitNull(b bool)`

SetRootExplicitNull (un)sets Root to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Root value is set to nil even if false is passed
### GetSpecial

`func (o *Drive) GetSpecial() []MicrosoftGraphDriveItem`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *Drive) GetSpecialOk() ([]MicrosoftGraphDriveItem, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpecial

`func (o *Drive) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.

### SetSpecial

`func (o *Drive) SetSpecial(v []MicrosoftGraphDriveItem)`

SetSpecial gets a reference to the given []MicrosoftGraphDriveItem and assigns it to the Special field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


