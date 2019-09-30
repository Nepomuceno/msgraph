# LocateDeviceActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceLocation** | Pointer to [**AnyOfmicrosoftGraphDeviceGeoLocation**](anyOf&lt;microsoft.graph.deviceGeoLocation&gt;.md) | device location | [optional] 

## Methods

### GetDeviceLocation

`func (o *LocateDeviceActionResult) GetDeviceLocation() AnyOfmicrosoftGraphDeviceGeoLocation`

GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.

### GetDeviceLocationOk

`func (o *LocateDeviceActionResult) GetDeviceLocationOk() (AnyOfmicrosoftGraphDeviceGeoLocation, bool)`

GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceLocation

`func (o *LocateDeviceActionResult) HasDeviceLocation() bool`

HasDeviceLocation returns a boolean if a field has been set.

### SetDeviceLocation

`func (o *LocateDeviceActionResult) SetDeviceLocation(v AnyOfmicrosoftGraphDeviceGeoLocation)`

SetDeviceLocation gets a reference to the given AnyOfmicrosoftGraphDeviceGeoLocation and assigns it to the DeviceLocation field.

### SetDeviceLocationExplicitNull

`func (o *LocateDeviceActionResult) SetDeviceLocationExplicitNull(b bool)`

SetDeviceLocationExplicitNull (un)sets DeviceLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceLocation value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


