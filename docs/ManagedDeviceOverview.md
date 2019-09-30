# ManagedDeviceOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnrolledDeviceCount** | Pointer to **int32** | Total enrolled device count. Does not include PC devices managed via Intune PC Agent | [optional] 
**MdmEnrolledCount** | Pointer to **int32** | The number of devices enrolled in MDM | [optional] 
**DualEnrolledDeviceCount** | Pointer to **int32** | The number of devices enrolled in both MDM and EAS | [optional] 
**DeviceOperatingSystemSummary** | Pointer to [**AnyOfmicrosoftGraphDeviceOperatingSystemSummary**](anyOf&lt;microsoft.graph.deviceOperatingSystemSummary&gt;.md) | Device operating system summary. | [optional] 
**DeviceExchangeAccessStateSummary** | Pointer to [**AnyOfmicrosoftGraphDeviceExchangeAccessStateSummary**](anyOf&lt;microsoft.graph.deviceExchangeAccessStateSummary&gt;.md) | Distribution of Exchange Access State in Intune | [optional] 

## Methods

### GetEnrolledDeviceCount

`func (o *ManagedDeviceOverview) GetEnrolledDeviceCount() int32`

GetEnrolledDeviceCount returns the EnrolledDeviceCount field if non-nil, zero value otherwise.

### GetEnrolledDeviceCountOk

`func (o *ManagedDeviceOverview) GetEnrolledDeviceCountOk() (int32, bool)`

GetEnrolledDeviceCountOk returns a tuple with the EnrolledDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnrolledDeviceCount

`func (o *ManagedDeviceOverview) HasEnrolledDeviceCount() bool`

HasEnrolledDeviceCount returns a boolean if a field has been set.

### SetEnrolledDeviceCount

`func (o *ManagedDeviceOverview) SetEnrolledDeviceCount(v int32)`

SetEnrolledDeviceCount gets a reference to the given int32 and assigns it to the EnrolledDeviceCount field.

### GetMdmEnrolledCount

`func (o *ManagedDeviceOverview) GetMdmEnrolledCount() int32`

GetMdmEnrolledCount returns the MdmEnrolledCount field if non-nil, zero value otherwise.

### GetMdmEnrolledCountOk

`func (o *ManagedDeviceOverview) GetMdmEnrolledCountOk() (int32, bool)`

GetMdmEnrolledCountOk returns a tuple with the MdmEnrolledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMdmEnrolledCount

`func (o *ManagedDeviceOverview) HasMdmEnrolledCount() bool`

HasMdmEnrolledCount returns a boolean if a field has been set.

### SetMdmEnrolledCount

`func (o *ManagedDeviceOverview) SetMdmEnrolledCount(v int32)`

SetMdmEnrolledCount gets a reference to the given int32 and assigns it to the MdmEnrolledCount field.

### GetDualEnrolledDeviceCount

`func (o *ManagedDeviceOverview) GetDualEnrolledDeviceCount() int32`

GetDualEnrolledDeviceCount returns the DualEnrolledDeviceCount field if non-nil, zero value otherwise.

### GetDualEnrolledDeviceCountOk

`func (o *ManagedDeviceOverview) GetDualEnrolledDeviceCountOk() (int32, bool)`

GetDualEnrolledDeviceCountOk returns a tuple with the DualEnrolledDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDualEnrolledDeviceCount

`func (o *ManagedDeviceOverview) HasDualEnrolledDeviceCount() bool`

HasDualEnrolledDeviceCount returns a boolean if a field has been set.

### SetDualEnrolledDeviceCount

`func (o *ManagedDeviceOverview) SetDualEnrolledDeviceCount(v int32)`

SetDualEnrolledDeviceCount gets a reference to the given int32 and assigns it to the DualEnrolledDeviceCount field.

### GetDeviceOperatingSystemSummary

`func (o *ManagedDeviceOverview) GetDeviceOperatingSystemSummary() AnyOfmicrosoftGraphDeviceOperatingSystemSummary`

GetDeviceOperatingSystemSummary returns the DeviceOperatingSystemSummary field if non-nil, zero value otherwise.

### GetDeviceOperatingSystemSummaryOk

`func (o *ManagedDeviceOverview) GetDeviceOperatingSystemSummaryOk() (AnyOfmicrosoftGraphDeviceOperatingSystemSummary, bool)`

GetDeviceOperatingSystemSummaryOk returns a tuple with the DeviceOperatingSystemSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceOperatingSystemSummary

`func (o *ManagedDeviceOverview) HasDeviceOperatingSystemSummary() bool`

HasDeviceOperatingSystemSummary returns a boolean if a field has been set.

### SetDeviceOperatingSystemSummary

`func (o *ManagedDeviceOverview) SetDeviceOperatingSystemSummary(v AnyOfmicrosoftGraphDeviceOperatingSystemSummary)`

SetDeviceOperatingSystemSummary gets a reference to the given AnyOfmicrosoftGraphDeviceOperatingSystemSummary and assigns it to the DeviceOperatingSystemSummary field.

### SetDeviceOperatingSystemSummaryExplicitNull

`func (o *ManagedDeviceOverview) SetDeviceOperatingSystemSummaryExplicitNull(b bool)`

SetDeviceOperatingSystemSummaryExplicitNull (un)sets DeviceOperatingSystemSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceOperatingSystemSummary value is set to nil even if false is passed
### GetDeviceExchangeAccessStateSummary

`func (o *ManagedDeviceOverview) GetDeviceExchangeAccessStateSummary() AnyOfmicrosoftGraphDeviceExchangeAccessStateSummary`

GetDeviceExchangeAccessStateSummary returns the DeviceExchangeAccessStateSummary field if non-nil, zero value otherwise.

### GetDeviceExchangeAccessStateSummaryOk

`func (o *ManagedDeviceOverview) GetDeviceExchangeAccessStateSummaryOk() (AnyOfmicrosoftGraphDeviceExchangeAccessStateSummary, bool)`

GetDeviceExchangeAccessStateSummaryOk returns a tuple with the DeviceExchangeAccessStateSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceExchangeAccessStateSummary

`func (o *ManagedDeviceOverview) HasDeviceExchangeAccessStateSummary() bool`

HasDeviceExchangeAccessStateSummary returns a boolean if a field has been set.

### SetDeviceExchangeAccessStateSummary

`func (o *ManagedDeviceOverview) SetDeviceExchangeAccessStateSummary(v AnyOfmicrosoftGraphDeviceExchangeAccessStateSummary)`

SetDeviceExchangeAccessStateSummary gets a reference to the given AnyOfmicrosoftGraphDeviceExchangeAccessStateSummary and assigns it to the DeviceExchangeAccessStateSummary field.

### SetDeviceExchangeAccessStateSummaryExplicitNull

`func (o *ManagedDeviceOverview) SetDeviceExchangeAccessStateSummaryExplicitNull(b bool)`

SetDeviceExchangeAccessStateSummaryExplicitNull (un)sets DeviceExchangeAccessStateSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceExchangeAccessStateSummary value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


