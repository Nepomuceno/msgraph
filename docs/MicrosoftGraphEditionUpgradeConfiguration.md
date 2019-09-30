# MicrosoftGraphEditionUpgradeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was last modified. | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was created. | [optional] 
**Description** | Pointer to **string** | Admin provided description of the Device Configuration. | [optional] 
**DisplayName** | Pointer to **string** | Admin provided name of the device configuration. | [optional] 
**Version** | Pointer to **int32** | Version of the device configuration. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphDeviceConfigurationAssignment**](microsoft.graph.deviceConfigurationAssignment.md) |  | [optional] 
**DeviceStatuses** | Pointer to [**[]MicrosoftGraphDeviceConfigurationDeviceStatus**](microsoft.graph.deviceConfigurationDeviceStatus.md) |  | [optional] 
**UserStatuses** | Pointer to [**[]MicrosoftGraphDeviceConfigurationUserStatus**](microsoft.graph.deviceConfigurationUserStatus.md) |  | [optional] 
**DeviceStatusOverview** | Pointer to [**AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview**](anyOf&lt;microsoft.graph.deviceConfigurationDeviceOverview&gt;.md) |  | [optional] 
**UserStatusOverview** | Pointer to [**AnyOfmicrosoftGraphDeviceConfigurationUserOverview**](anyOf&lt;microsoft.graph.deviceConfigurationUserOverview&gt;.md) |  | [optional] 
**DeviceSettingStateSummaries** | Pointer to [**[]MicrosoftGraphSettingStateDeviceSummary**](microsoft.graph.settingStateDeviceSummary.md) |  | [optional] 
**LicenseType** | Pointer to [**AnyOfmicrosoftGraphEditionUpgradeLicenseType**](anyOf&lt;microsoft.graph.editionUpgradeLicenseType&gt;.md) | Edition Upgrade License Type. | [optional] 
**TargetEdition** | Pointer to [**AnyOfmicrosoftGraphWindows10EditionType**](anyOf&lt;microsoft.graph.windows10EditionType&gt;.md) | Edition Upgrade Target Edition. | [optional] 
**License** | Pointer to **string** | Edition Upgrade License File Content. | [optional] 
**ProductKey** | Pointer to **string** | Edition Upgrade Product Key. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetLicenseType

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetLicenseType() AnyOfmicrosoftGraphEditionUpgradeLicenseType`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetLicenseTypeOk() (AnyOfmicrosoftGraphEditionUpgradeLicenseType, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicenseType

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### SetLicenseType

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetLicenseType(v AnyOfmicrosoftGraphEditionUpgradeLicenseType)`

SetLicenseType gets a reference to the given AnyOfmicrosoftGraphEditionUpgradeLicenseType and assigns it to the LicenseType field.

### GetTargetEdition

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetTargetEdition() AnyOfmicrosoftGraphWindows10EditionType`

GetTargetEdition returns the TargetEdition field if non-nil, zero value otherwise.

### GetTargetEditionOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetTargetEditionOk() (AnyOfmicrosoftGraphWindows10EditionType, bool)`

GetTargetEditionOk returns a tuple with the TargetEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetEdition

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasTargetEdition() bool`

HasTargetEdition returns a boolean if a field has been set.

### SetTargetEdition

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetTargetEdition(v AnyOfmicrosoftGraphWindows10EditionType)`

SetTargetEdition gets a reference to the given AnyOfmicrosoftGraphWindows10EditionType and assigns it to the TargetEdition field.

### GetLicense

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetLicenseOk() (string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicense

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### SetLicense

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetLicense(v string)`

SetLicense gets a reference to the given string and assigns it to the License field.

### SetLicenseExplicitNull

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetLicenseExplicitNull(b bool)`

SetLicenseExplicitNull (un)sets License to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The License value is set to nil even if false is passed
### GetProductKey

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *MicrosoftGraphEditionUpgradeConfiguration) GetProductKeyOk() (string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductKey

`func (o *MicrosoftGraphEditionUpgradeConfiguration) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### SetProductKey

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetProductKey(v string)`

SetProductKey gets a reference to the given string and assigns it to the ProductKey field.

### SetProductKeyExplicitNull

`func (o *MicrosoftGraphEditionUpgradeConfiguration) SetProductKeyExplicitNull(b bool)`

SetProductKeyExplicitNull (un)sets ProductKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProductKey value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


