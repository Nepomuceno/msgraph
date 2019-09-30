# DetectedApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Name of the discovered application. Read-only | [optional] 
**Version** | Pointer to **string** | Version of the discovered application. Read-only | [optional] 
**SizeInByte** | Pointer to **int64** | Discovered application size in bytes. Read-only | [optional] 
**DeviceCount** | Pointer to **int32** | The number of devices that have installed this application | [optional] 
**ManagedDevices** | Pointer to [**[]MicrosoftGraphManagedDevice**](microsoft.graph.managedDevice.md) |  | [optional] 

## Methods

### GetDisplayName

`func (o *DetectedApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DetectedApp) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *DetectedApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *DetectedApp) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *DetectedApp) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetVersion

`func (o *DetectedApp) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DetectedApp) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *DetectedApp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *DetectedApp) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *DetectedApp) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetSizeInByte

`func (o *DetectedApp) GetSizeInByte() int64`

GetSizeInByte returns the SizeInByte field if non-nil, zero value otherwise.

### GetSizeInByteOk

`func (o *DetectedApp) GetSizeInByteOk() (int64, bool)`

GetSizeInByteOk returns a tuple with the SizeInByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSizeInByte

`func (o *DetectedApp) HasSizeInByte() bool`

HasSizeInByte returns a boolean if a field has been set.

### SetSizeInByte

`func (o *DetectedApp) SetSizeInByte(v int64)`

SetSizeInByte gets a reference to the given int64 and assigns it to the SizeInByte field.

### GetDeviceCount

`func (o *DetectedApp) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *DetectedApp) GetDeviceCountOk() (int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCount

`func (o *DetectedApp) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### SetDeviceCount

`func (o *DetectedApp) SetDeviceCount(v int32)`

SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.

### GetManagedDevices

`func (o *DetectedApp) GetManagedDevices() []MicrosoftGraphManagedDevice`

GetManagedDevices returns the ManagedDevices field if non-nil, zero value otherwise.

### GetManagedDevicesOk

`func (o *DetectedApp) GetManagedDevicesOk() ([]MicrosoftGraphManagedDevice, bool)`

GetManagedDevicesOk returns a tuple with the ManagedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedDevices

`func (o *DetectedApp) HasManagedDevices() bool`

HasManagedDevices returns a boolean if a field has been set.

### SetManagedDevices

`func (o *DetectedApp) SetManagedDevices(v []MicrosoftGraphManagedDevice)`

SetManagedDevices gets a reference to the given []MicrosoftGraphManagedDevice and assigns it to the ManagedDevices field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


