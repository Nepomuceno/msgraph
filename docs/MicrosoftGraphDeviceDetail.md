# MicrosoftGraphDeviceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**Browser** | Pointer to **string** |  | [optional] 
**IsCompliant** | Pointer to **bool** |  | [optional] 
**IsManaged** | Pointer to **bool** |  | [optional] 
**TrustType** | Pointer to **string** |  | [optional] 

## Methods

### GetDeviceId

`func (o *MicrosoftGraphDeviceDetail) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MicrosoftGraphDeviceDetail) GetDeviceIdOk() (string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceId

`func (o *MicrosoftGraphDeviceDetail) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceId

`func (o *MicrosoftGraphDeviceDetail) SetDeviceId(v string)`

SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.

### SetDeviceIdExplicitNull

`func (o *MicrosoftGraphDeviceDetail) SetDeviceIdExplicitNull(b bool)`

SetDeviceIdExplicitNull (un)sets DeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceId value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphDeviceDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDeviceDetail) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphDeviceDetail) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphDeviceDetail) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphDeviceDetail) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetOperatingSystem

`func (o *MicrosoftGraphDeviceDetail) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *MicrosoftGraphDeviceDetail) GetOperatingSystemOk() (string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperatingSystem

`func (o *MicrosoftGraphDeviceDetail) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystem

`func (o *MicrosoftGraphDeviceDetail) SetOperatingSystem(v string)`

SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.

### SetOperatingSystemExplicitNull

`func (o *MicrosoftGraphDeviceDetail) SetOperatingSystemExplicitNull(b bool)`

SetOperatingSystemExplicitNull (un)sets OperatingSystem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OperatingSystem value is set to nil even if false is passed
### GetBrowser

`func (o *MicrosoftGraphDeviceDetail) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *MicrosoftGraphDeviceDetail) GetBrowserOk() (string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrowser

`func (o *MicrosoftGraphDeviceDetail) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### SetBrowser

`func (o *MicrosoftGraphDeviceDetail) SetBrowser(v string)`

SetBrowser gets a reference to the given string and assigns it to the Browser field.

### SetBrowserExplicitNull

`func (o *MicrosoftGraphDeviceDetail) SetBrowserExplicitNull(b bool)`

SetBrowserExplicitNull (un)sets Browser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Browser value is set to nil even if false is passed
### GetIsCompliant

`func (o *MicrosoftGraphDeviceDetail) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *MicrosoftGraphDeviceDetail) GetIsCompliantOk() (bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsCompliant

`func (o *MicrosoftGraphDeviceDetail) HasIsCompliant() bool`

HasIsCompliant returns a boolean if a field has been set.

### SetIsCompliant

`func (o *MicrosoftGraphDeviceDetail) SetIsCompliant(v bool)`

SetIsCompliant gets a reference to the given bool and assigns it to the IsCompliant field.

### SetIsCompliantExplicitNull

`func (o *MicrosoftGraphDeviceDetail) SetIsCompliantExplicitNull(b bool)`

SetIsCompliantExplicitNull (un)sets IsCompliant to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsCompliant value is set to nil even if false is passed
### GetIsManaged

`func (o *MicrosoftGraphDeviceDetail) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *MicrosoftGraphDeviceDetail) GetIsManagedOk() (bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsManaged

`func (o *MicrosoftGraphDeviceDetail) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### SetIsManaged

`func (o *MicrosoftGraphDeviceDetail) SetIsManaged(v bool)`

SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.

### SetIsManagedExplicitNull

`func (o *MicrosoftGraphDeviceDetail) SetIsManagedExplicitNull(b bool)`

SetIsManagedExplicitNull (un)sets IsManaged to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsManaged value is set to nil even if false is passed
### GetTrustType

`func (o *MicrosoftGraphDeviceDetail) GetTrustType() string`

GetTrustType returns the TrustType field if non-nil, zero value otherwise.

### GetTrustTypeOk

`func (o *MicrosoftGraphDeviceDetail) GetTrustTypeOk() (string, bool)`

GetTrustTypeOk returns a tuple with the TrustType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTrustType

`func (o *MicrosoftGraphDeviceDetail) HasTrustType() bool`

HasTrustType returns a boolean if a field has been set.

### SetTrustType

`func (o *MicrosoftGraphDeviceDetail) SetTrustType(v string)`

SetTrustType gets a reference to the given string and assigns it to the TrustType field.

### SetTrustTypeExplicitNull

`func (o *MicrosoftGraphDeviceDetail) SetTrustTypeExplicitNull(b bool)`

SetTrustTypeExplicitNull (un)sets TrustType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TrustType value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


