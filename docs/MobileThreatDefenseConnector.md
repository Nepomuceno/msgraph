# MobileThreatDefenseConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastHeartbeatDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime of last Heartbeat recieved from the Data Sync Partner | [optional] 
**PartnerState** | Pointer to [**AnyOfmicrosoftGraphMobileThreatPartnerTenantState**](anyOf&lt;microsoft.graph.mobileThreatPartnerTenantState&gt;.md) | Data Sync Partner state for this account | [optional] 
**AndroidEnabled** | Pointer to **bool** | For Android, set whether data from the data sync partner should be used during compliance evaluations | [optional] 
**IosEnabled** | Pointer to **bool** | For IOS, get or set whether data from the data sync partner should be used during compliance evaluations | [optional] 
**AndroidDeviceBlockedOnMissingPartnerData** | Pointer to **bool** | For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant | [optional] 
**IosDeviceBlockedOnMissingPartnerData** | Pointer to **bool** | For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant | [optional] 
**PartnerUnsupportedOsVersionBlocked** | Pointer to **bool** | Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner | [optional] 
**PartnerUnresponsivenessThresholdInDays** | Pointer to **int32** | Get or Set days the per tenant tolerance to unresponsiveness for this partner integration | [optional] 

## Methods

### GetLastHeartbeatDateTime

`func (o *MobileThreatDefenseConnector) GetLastHeartbeatDateTime() time.Time`

GetLastHeartbeatDateTime returns the LastHeartbeatDateTime field if non-nil, zero value otherwise.

### GetLastHeartbeatDateTimeOk

`func (o *MobileThreatDefenseConnector) GetLastHeartbeatDateTimeOk() (time.Time, bool)`

GetLastHeartbeatDateTimeOk returns a tuple with the LastHeartbeatDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastHeartbeatDateTime

`func (o *MobileThreatDefenseConnector) HasLastHeartbeatDateTime() bool`

HasLastHeartbeatDateTime returns a boolean if a field has been set.

### SetLastHeartbeatDateTime

`func (o *MobileThreatDefenseConnector) SetLastHeartbeatDateTime(v time.Time)`

SetLastHeartbeatDateTime gets a reference to the given time.Time and assigns it to the LastHeartbeatDateTime field.

### GetPartnerState

`func (o *MobileThreatDefenseConnector) GetPartnerState() AnyOfmicrosoftGraphMobileThreatPartnerTenantState`

GetPartnerState returns the PartnerState field if non-nil, zero value otherwise.

### GetPartnerStateOk

`func (o *MobileThreatDefenseConnector) GetPartnerStateOk() (AnyOfmicrosoftGraphMobileThreatPartnerTenantState, bool)`

GetPartnerStateOk returns a tuple with the PartnerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPartnerState

`func (o *MobileThreatDefenseConnector) HasPartnerState() bool`

HasPartnerState returns a boolean if a field has been set.

### SetPartnerState

`func (o *MobileThreatDefenseConnector) SetPartnerState(v AnyOfmicrosoftGraphMobileThreatPartnerTenantState)`

SetPartnerState gets a reference to the given AnyOfmicrosoftGraphMobileThreatPartnerTenantState and assigns it to the PartnerState field.

### GetAndroidEnabled

`func (o *MobileThreatDefenseConnector) GetAndroidEnabled() bool`

GetAndroidEnabled returns the AndroidEnabled field if non-nil, zero value otherwise.

### GetAndroidEnabledOk

`func (o *MobileThreatDefenseConnector) GetAndroidEnabledOk() (bool, bool)`

GetAndroidEnabledOk returns a tuple with the AndroidEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAndroidEnabled

`func (o *MobileThreatDefenseConnector) HasAndroidEnabled() bool`

HasAndroidEnabled returns a boolean if a field has been set.

### SetAndroidEnabled

`func (o *MobileThreatDefenseConnector) SetAndroidEnabled(v bool)`

SetAndroidEnabled gets a reference to the given bool and assigns it to the AndroidEnabled field.

### GetIosEnabled

`func (o *MobileThreatDefenseConnector) GetIosEnabled() bool`

GetIosEnabled returns the IosEnabled field if non-nil, zero value otherwise.

### GetIosEnabledOk

`func (o *MobileThreatDefenseConnector) GetIosEnabledOk() (bool, bool)`

GetIosEnabledOk returns a tuple with the IosEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIosEnabled

`func (o *MobileThreatDefenseConnector) HasIosEnabled() bool`

HasIosEnabled returns a boolean if a field has been set.

### SetIosEnabled

`func (o *MobileThreatDefenseConnector) SetIosEnabled(v bool)`

SetIosEnabled gets a reference to the given bool and assigns it to the IosEnabled field.

### GetAndroidDeviceBlockedOnMissingPartnerData

`func (o *MobileThreatDefenseConnector) GetAndroidDeviceBlockedOnMissingPartnerData() bool`

GetAndroidDeviceBlockedOnMissingPartnerData returns the AndroidDeviceBlockedOnMissingPartnerData field if non-nil, zero value otherwise.

### GetAndroidDeviceBlockedOnMissingPartnerDataOk

`func (o *MobileThreatDefenseConnector) GetAndroidDeviceBlockedOnMissingPartnerDataOk() (bool, bool)`

GetAndroidDeviceBlockedOnMissingPartnerDataOk returns a tuple with the AndroidDeviceBlockedOnMissingPartnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAndroidDeviceBlockedOnMissingPartnerData

`func (o *MobileThreatDefenseConnector) HasAndroidDeviceBlockedOnMissingPartnerData() bool`

HasAndroidDeviceBlockedOnMissingPartnerData returns a boolean if a field has been set.

### SetAndroidDeviceBlockedOnMissingPartnerData

`func (o *MobileThreatDefenseConnector) SetAndroidDeviceBlockedOnMissingPartnerData(v bool)`

SetAndroidDeviceBlockedOnMissingPartnerData gets a reference to the given bool and assigns it to the AndroidDeviceBlockedOnMissingPartnerData field.

### GetIosDeviceBlockedOnMissingPartnerData

`func (o *MobileThreatDefenseConnector) GetIosDeviceBlockedOnMissingPartnerData() bool`

GetIosDeviceBlockedOnMissingPartnerData returns the IosDeviceBlockedOnMissingPartnerData field if non-nil, zero value otherwise.

### GetIosDeviceBlockedOnMissingPartnerDataOk

`func (o *MobileThreatDefenseConnector) GetIosDeviceBlockedOnMissingPartnerDataOk() (bool, bool)`

GetIosDeviceBlockedOnMissingPartnerDataOk returns a tuple with the IosDeviceBlockedOnMissingPartnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIosDeviceBlockedOnMissingPartnerData

`func (o *MobileThreatDefenseConnector) HasIosDeviceBlockedOnMissingPartnerData() bool`

HasIosDeviceBlockedOnMissingPartnerData returns a boolean if a field has been set.

### SetIosDeviceBlockedOnMissingPartnerData

`func (o *MobileThreatDefenseConnector) SetIosDeviceBlockedOnMissingPartnerData(v bool)`

SetIosDeviceBlockedOnMissingPartnerData gets a reference to the given bool and assigns it to the IosDeviceBlockedOnMissingPartnerData field.

### GetPartnerUnsupportedOsVersionBlocked

`func (o *MobileThreatDefenseConnector) GetPartnerUnsupportedOsVersionBlocked() bool`

GetPartnerUnsupportedOsVersionBlocked returns the PartnerUnsupportedOsVersionBlocked field if non-nil, zero value otherwise.

### GetPartnerUnsupportedOsVersionBlockedOk

`func (o *MobileThreatDefenseConnector) GetPartnerUnsupportedOsVersionBlockedOk() (bool, bool)`

GetPartnerUnsupportedOsVersionBlockedOk returns a tuple with the PartnerUnsupportedOsVersionBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPartnerUnsupportedOsVersionBlocked

`func (o *MobileThreatDefenseConnector) HasPartnerUnsupportedOsVersionBlocked() bool`

HasPartnerUnsupportedOsVersionBlocked returns a boolean if a field has been set.

### SetPartnerUnsupportedOsVersionBlocked

`func (o *MobileThreatDefenseConnector) SetPartnerUnsupportedOsVersionBlocked(v bool)`

SetPartnerUnsupportedOsVersionBlocked gets a reference to the given bool and assigns it to the PartnerUnsupportedOsVersionBlocked field.

### GetPartnerUnresponsivenessThresholdInDays

`func (o *MobileThreatDefenseConnector) GetPartnerUnresponsivenessThresholdInDays() int32`

GetPartnerUnresponsivenessThresholdInDays returns the PartnerUnresponsivenessThresholdInDays field if non-nil, zero value otherwise.

### GetPartnerUnresponsivenessThresholdInDaysOk

`func (o *MobileThreatDefenseConnector) GetPartnerUnresponsivenessThresholdInDaysOk() (int32, bool)`

GetPartnerUnresponsivenessThresholdInDaysOk returns a tuple with the PartnerUnresponsivenessThresholdInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPartnerUnresponsivenessThresholdInDays

`func (o *MobileThreatDefenseConnector) HasPartnerUnresponsivenessThresholdInDays() bool`

HasPartnerUnresponsivenessThresholdInDays returns a boolean if a field has been set.

### SetPartnerUnresponsivenessThresholdInDays

`func (o *MobileThreatDefenseConnector) SetPartnerUnresponsivenessThresholdInDays(v int32)`

SetPartnerUnresponsivenessThresholdInDays gets a reference to the given int32 and assigns it to the PartnerUnresponsivenessThresholdInDays field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


