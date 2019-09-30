# MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**InGracePeriodCount** | Pointer to **int32** | Number of devices that are in grace period | [optional] 
**ConfigManagerCount** | Pointer to **int32** | Number of devices that have compliance managed by System Center Configuration Manager | [optional] 
**UnknownDeviceCount** | Pointer to **int32** | Number of unknown devices | [optional] 
**NotApplicableDeviceCount** | Pointer to **int32** | Number of not applicable devices | [optional] 
**CompliantDeviceCount** | Pointer to **int32** | Number of compliant devices | [optional] 
**RemediatedDeviceCount** | Pointer to **int32** | Number of remediated devices | [optional] 
**NonCompliantDeviceCount** | Pointer to **int32** | Number of NonCompliant devices | [optional] 
**ErrorDeviceCount** | Pointer to **int32** | Number of error devices | [optional] 
**ConflictDeviceCount** | Pointer to **int32** | Number of conflict devices | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetInGracePeriodCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetInGracePeriodCount() int32`

GetInGracePeriodCount returns the InGracePeriodCount field if non-nil, zero value otherwise.

### GetInGracePeriodCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetInGracePeriodCountOk() (int32, bool)`

GetInGracePeriodCountOk returns a tuple with the InGracePeriodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInGracePeriodCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasInGracePeriodCount() bool`

HasInGracePeriodCount returns a boolean if a field has been set.

### SetInGracePeriodCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetInGracePeriodCount(v int32)`

SetInGracePeriodCount gets a reference to the given int32 and assigns it to the InGracePeriodCount field.

### GetConfigManagerCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetConfigManagerCount() int32`

GetConfigManagerCount returns the ConfigManagerCount field if non-nil, zero value otherwise.

### GetConfigManagerCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetConfigManagerCountOk() (int32, bool)`

GetConfigManagerCountOk returns a tuple with the ConfigManagerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfigManagerCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasConfigManagerCount() bool`

HasConfigManagerCount returns a boolean if a field has been set.

### SetConfigManagerCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetConfigManagerCount(v int32)`

SetConfigManagerCount gets a reference to the given int32 and assigns it to the ConfigManagerCount field.

### GetUnknownDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetUnknownDeviceCount() int32`

GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.

### GetUnknownDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetUnknownDeviceCountOk() (int32, bool)`

GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnknownDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasUnknownDeviceCount() bool`

HasUnknownDeviceCount returns a boolean if a field has been set.

### SetUnknownDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetUnknownDeviceCount(v int32)`

SetUnknownDeviceCount gets a reference to the given int32 and assigns it to the UnknownDeviceCount field.

### GetNotApplicableDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetNotApplicableDeviceCount() int32`

GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.

### GetNotApplicableDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetNotApplicableDeviceCountOk() (int32, bool)`

GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotApplicableDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasNotApplicableDeviceCount() bool`

HasNotApplicableDeviceCount returns a boolean if a field has been set.

### SetNotApplicableDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetNotApplicableDeviceCount(v int32)`

SetNotApplicableDeviceCount gets a reference to the given int32 and assigns it to the NotApplicableDeviceCount field.

### GetCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetCompliantDeviceCount() int32`

GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.

### GetCompliantDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetCompliantDeviceCountOk() (int32, bool)`

GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasCompliantDeviceCount() bool`

HasCompliantDeviceCount returns a boolean if a field has been set.

### SetCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetCompliantDeviceCount(v int32)`

SetCompliantDeviceCount gets a reference to the given int32 and assigns it to the CompliantDeviceCount field.

### GetRemediatedDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetRemediatedDeviceCount() int32`

GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.

### GetRemediatedDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetRemediatedDeviceCountOk() (int32, bool)`

GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemediatedDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasRemediatedDeviceCount() bool`

HasRemediatedDeviceCount returns a boolean if a field has been set.

### SetRemediatedDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetRemediatedDeviceCount(v int32)`

SetRemediatedDeviceCount gets a reference to the given int32 and assigns it to the RemediatedDeviceCount field.

### GetNonCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetNonCompliantDeviceCount() int32`

GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.

### GetNonCompliantDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetNonCompliantDeviceCountOk() (int32, bool)`

GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNonCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasNonCompliantDeviceCount() bool`

HasNonCompliantDeviceCount returns a boolean if a field has been set.

### SetNonCompliantDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetNonCompliantDeviceCount(v int32)`

SetNonCompliantDeviceCount gets a reference to the given int32 and assigns it to the NonCompliantDeviceCount field.

### GetErrorDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetErrorDeviceCount() int32`

GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.

### GetErrorDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetErrorDeviceCountOk() (int32, bool)`

GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasErrorDeviceCount() bool`

HasErrorDeviceCount returns a boolean if a field has been set.

### SetErrorDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetErrorDeviceCount(v int32)`

SetErrorDeviceCount gets a reference to the given int32 and assigns it to the ErrorDeviceCount field.

### GetConflictDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetConflictDeviceCount() int32`

GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.

### GetConflictDeviceCountOk

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) GetConflictDeviceCountOk() (int32, bool)`

GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConflictDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) HasConflictDeviceCount() bool`

HasConflictDeviceCount returns a boolean if a field has been set.

### SetConflictDeviceCount

`func (o *MicrosoftGraphDeviceCompliancePolicyDeviceStateSummary) SetConflictDeviceCount(v int32)`

SetConflictDeviceCount gets a reference to the given int32 and assigns it to the ConflictDeviceCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


