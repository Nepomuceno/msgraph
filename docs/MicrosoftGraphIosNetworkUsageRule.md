# MicrosoftGraphIosNetworkUsageRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagedApps** | Pointer to [**[]AnyOfmicrosoftGraphAppListItem**](anyOf&lt;microsoft.graph.appListItem&gt;.md) | Information about the managed apps that this rule is going to apply to. This collection can contain a maximum of 500 elements. | [optional] 
**CellularDataBlockWhenRoaming** | Pointer to **bool** | If set to true, corresponding managed apps will not be allowed to use cellular data when roaming. | [optional] 
**CellularDataBlocked** | Pointer to **bool** | If set to true, corresponding managed apps will not be allowed to use cellular data at any time. | [optional] 

## Methods

### GetManagedApps

`func (o *MicrosoftGraphIosNetworkUsageRule) GetManagedApps() []AnyOfmicrosoftGraphAppListItem`

GetManagedApps returns the ManagedApps field if non-nil, zero value otherwise.

### GetManagedAppsOk

`func (o *MicrosoftGraphIosNetworkUsageRule) GetManagedAppsOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetManagedAppsOk returns a tuple with the ManagedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedApps

`func (o *MicrosoftGraphIosNetworkUsageRule) HasManagedApps() bool`

HasManagedApps returns a boolean if a field has been set.

### SetManagedApps

`func (o *MicrosoftGraphIosNetworkUsageRule) SetManagedApps(v []AnyOfmicrosoftGraphAppListItem)`

SetManagedApps gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the ManagedApps field.

### GetCellularDataBlockWhenRoaming

`func (o *MicrosoftGraphIosNetworkUsageRule) GetCellularDataBlockWhenRoaming() bool`

GetCellularDataBlockWhenRoaming returns the CellularDataBlockWhenRoaming field if non-nil, zero value otherwise.

### GetCellularDataBlockWhenRoamingOk

`func (o *MicrosoftGraphIosNetworkUsageRule) GetCellularDataBlockWhenRoamingOk() (bool, bool)`

GetCellularDataBlockWhenRoamingOk returns a tuple with the CellularDataBlockWhenRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularDataBlockWhenRoaming

`func (o *MicrosoftGraphIosNetworkUsageRule) HasCellularDataBlockWhenRoaming() bool`

HasCellularDataBlockWhenRoaming returns a boolean if a field has been set.

### SetCellularDataBlockWhenRoaming

`func (o *MicrosoftGraphIosNetworkUsageRule) SetCellularDataBlockWhenRoaming(v bool)`

SetCellularDataBlockWhenRoaming gets a reference to the given bool and assigns it to the CellularDataBlockWhenRoaming field.

### GetCellularDataBlocked

`func (o *MicrosoftGraphIosNetworkUsageRule) GetCellularDataBlocked() bool`

GetCellularDataBlocked returns the CellularDataBlocked field if non-nil, zero value otherwise.

### GetCellularDataBlockedOk

`func (o *MicrosoftGraphIosNetworkUsageRule) GetCellularDataBlockedOk() (bool, bool)`

GetCellularDataBlockedOk returns a tuple with the CellularDataBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularDataBlocked

`func (o *MicrosoftGraphIosNetworkUsageRule) HasCellularDataBlocked() bool`

HasCellularDataBlocked returns a boolean if a field has been set.

### SetCellularDataBlocked

`func (o *MicrosoftGraphIosNetworkUsageRule) SetCellularDataBlocked(v bool)`

SetCellularDataBlocked gets a reference to the given bool and assigns it to the CellularDataBlocked field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


