# DeviceManagementPartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastHeartbeatDateTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp of last heartbeat after admin enabled option Connect to Device management Partner | [optional] 
**PartnerState** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementPartnerTenantState**](anyOf&lt;microsoft.graph.deviceManagementPartnerTenantState&gt;.md) | Partner state of this tenant | [optional] 
**PartnerAppType** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementPartnerAppType**](anyOf&lt;microsoft.graph.deviceManagementPartnerAppType&gt;.md) | Partner App type | [optional] 
**SingleTenantAppId** | Pointer to **string** | Partner Single tenant App id | [optional] 
**DisplayName** | Pointer to **string** | Partner display name | [optional] 
**IsConfigured** | Pointer to **bool** | Whether device management partner is configured or not | [optional] 
**WhenPartnerDevicesWillBeRemovedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime in UTC when PartnerDevices will be removed | [optional] 
**WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime in UTC when PartnerDevices will be marked as NonCompliant | [optional] 

## Methods

### GetLastHeartbeatDateTime

`func (o *DeviceManagementPartner) GetLastHeartbeatDateTime() time.Time`

GetLastHeartbeatDateTime returns the LastHeartbeatDateTime field if non-nil, zero value otherwise.

### GetLastHeartbeatDateTimeOk

`func (o *DeviceManagementPartner) GetLastHeartbeatDateTimeOk() (time.Time, bool)`

GetLastHeartbeatDateTimeOk returns a tuple with the LastHeartbeatDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastHeartbeatDateTime

`func (o *DeviceManagementPartner) HasLastHeartbeatDateTime() bool`

HasLastHeartbeatDateTime returns a boolean if a field has been set.

### SetLastHeartbeatDateTime

`func (o *DeviceManagementPartner) SetLastHeartbeatDateTime(v time.Time)`

SetLastHeartbeatDateTime gets a reference to the given time.Time and assigns it to the LastHeartbeatDateTime field.

### GetPartnerState

`func (o *DeviceManagementPartner) GetPartnerState() AnyOfmicrosoftGraphDeviceManagementPartnerTenantState`

GetPartnerState returns the PartnerState field if non-nil, zero value otherwise.

### GetPartnerStateOk

`func (o *DeviceManagementPartner) GetPartnerStateOk() (AnyOfmicrosoftGraphDeviceManagementPartnerTenantState, bool)`

GetPartnerStateOk returns a tuple with the PartnerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPartnerState

`func (o *DeviceManagementPartner) HasPartnerState() bool`

HasPartnerState returns a boolean if a field has been set.

### SetPartnerState

`func (o *DeviceManagementPartner) SetPartnerState(v AnyOfmicrosoftGraphDeviceManagementPartnerTenantState)`

SetPartnerState gets a reference to the given AnyOfmicrosoftGraphDeviceManagementPartnerTenantState and assigns it to the PartnerState field.

### GetPartnerAppType

`func (o *DeviceManagementPartner) GetPartnerAppType() AnyOfmicrosoftGraphDeviceManagementPartnerAppType`

GetPartnerAppType returns the PartnerAppType field if non-nil, zero value otherwise.

### GetPartnerAppTypeOk

`func (o *DeviceManagementPartner) GetPartnerAppTypeOk() (AnyOfmicrosoftGraphDeviceManagementPartnerAppType, bool)`

GetPartnerAppTypeOk returns a tuple with the PartnerAppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPartnerAppType

`func (o *DeviceManagementPartner) HasPartnerAppType() bool`

HasPartnerAppType returns a boolean if a field has been set.

### SetPartnerAppType

`func (o *DeviceManagementPartner) SetPartnerAppType(v AnyOfmicrosoftGraphDeviceManagementPartnerAppType)`

SetPartnerAppType gets a reference to the given AnyOfmicrosoftGraphDeviceManagementPartnerAppType and assigns it to the PartnerAppType field.

### GetSingleTenantAppId

`func (o *DeviceManagementPartner) GetSingleTenantAppId() string`

GetSingleTenantAppId returns the SingleTenantAppId field if non-nil, zero value otherwise.

### GetSingleTenantAppIdOk

`func (o *DeviceManagementPartner) GetSingleTenantAppIdOk() (string, bool)`

GetSingleTenantAppIdOk returns a tuple with the SingleTenantAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleTenantAppId

`func (o *DeviceManagementPartner) HasSingleTenantAppId() bool`

HasSingleTenantAppId returns a boolean if a field has been set.

### SetSingleTenantAppId

`func (o *DeviceManagementPartner) SetSingleTenantAppId(v string)`

SetSingleTenantAppId gets a reference to the given string and assigns it to the SingleTenantAppId field.

### SetSingleTenantAppIdExplicitNull

`func (o *DeviceManagementPartner) SetSingleTenantAppIdExplicitNull(b bool)`

SetSingleTenantAppIdExplicitNull (un)sets SingleTenantAppId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SingleTenantAppId value is set to nil even if false is passed
### GetDisplayName

`func (o *DeviceManagementPartner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceManagementPartner) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *DeviceManagementPartner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *DeviceManagementPartner) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *DeviceManagementPartner) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetIsConfigured

`func (o *DeviceManagementPartner) GetIsConfigured() bool`

GetIsConfigured returns the IsConfigured field if non-nil, zero value otherwise.

### GetIsConfiguredOk

`func (o *DeviceManagementPartner) GetIsConfiguredOk() (bool, bool)`

GetIsConfiguredOk returns a tuple with the IsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsConfigured

`func (o *DeviceManagementPartner) HasIsConfigured() bool`

HasIsConfigured returns a boolean if a field has been set.

### SetIsConfigured

`func (o *DeviceManagementPartner) SetIsConfigured(v bool)`

SetIsConfigured gets a reference to the given bool and assigns it to the IsConfigured field.

### GetWhenPartnerDevicesWillBeRemovedDateTime

`func (o *DeviceManagementPartner) GetWhenPartnerDevicesWillBeRemovedDateTime() time.Time`

GetWhenPartnerDevicesWillBeRemovedDateTime returns the WhenPartnerDevicesWillBeRemovedDateTime field if non-nil, zero value otherwise.

### GetWhenPartnerDevicesWillBeRemovedDateTimeOk

`func (o *DeviceManagementPartner) GetWhenPartnerDevicesWillBeRemovedDateTimeOk() (time.Time, bool)`

GetWhenPartnerDevicesWillBeRemovedDateTimeOk returns a tuple with the WhenPartnerDevicesWillBeRemovedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWhenPartnerDevicesWillBeRemovedDateTime

`func (o *DeviceManagementPartner) HasWhenPartnerDevicesWillBeRemovedDateTime() bool`

HasWhenPartnerDevicesWillBeRemovedDateTime returns a boolean if a field has been set.

### SetWhenPartnerDevicesWillBeRemovedDateTime

`func (o *DeviceManagementPartner) SetWhenPartnerDevicesWillBeRemovedDateTime(v time.Time)`

SetWhenPartnerDevicesWillBeRemovedDateTime gets a reference to the given time.Time and assigns it to the WhenPartnerDevicesWillBeRemovedDateTime field.

### SetWhenPartnerDevicesWillBeRemovedDateTimeExplicitNull

`func (o *DeviceManagementPartner) SetWhenPartnerDevicesWillBeRemovedDateTimeExplicitNull(b bool)`

SetWhenPartnerDevicesWillBeRemovedDateTimeExplicitNull (un)sets WhenPartnerDevicesWillBeRemovedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WhenPartnerDevicesWillBeRemovedDateTime value is set to nil even if false is passed
### GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime

`func (o *DeviceManagementPartner) GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime() time.Time`

GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime returns the WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime field if non-nil, zero value otherwise.

### GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTimeOk

`func (o *DeviceManagementPartner) GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTimeOk() (time.Time, bool)`

GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTimeOk returns a tuple with the WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime

`func (o *DeviceManagementPartner) HasWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime() bool`

HasWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime returns a boolean if a field has been set.

### SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime

`func (o *DeviceManagementPartner) SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime(v time.Time)`

SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime gets a reference to the given time.Time and assigns it to the WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime field.

### SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTimeExplicitNull

`func (o *DeviceManagementPartner) SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTimeExplicitNull(b bool)`

SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTimeExplicitNull (un)sets WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


