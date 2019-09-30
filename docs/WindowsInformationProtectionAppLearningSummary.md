# WindowsInformationProtectionAppLearningSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **string** | Application Name | [optional] 
**ApplicationType** | Pointer to [**AnyOfmicrosoftGraphApplicationType**](anyOf&lt;microsoft.graph.applicationType&gt;.md) | Application Type | [optional] 
**DeviceCount** | Pointer to **int32** | Device Count | [optional] 

## Methods

### GetApplicationName

`func (o *WindowsInformationProtectionAppLearningSummary) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *WindowsInformationProtectionAppLearningSummary) GetApplicationNameOk() (string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationName

`func (o *WindowsInformationProtectionAppLearningSummary) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### SetApplicationName

`func (o *WindowsInformationProtectionAppLearningSummary) SetApplicationName(v string)`

SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.

### SetApplicationNameExplicitNull

`func (o *WindowsInformationProtectionAppLearningSummary) SetApplicationNameExplicitNull(b bool)`

SetApplicationNameExplicitNull (un)sets ApplicationName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ApplicationName value is set to nil even if false is passed
### GetApplicationType

`func (o *WindowsInformationProtectionAppLearningSummary) GetApplicationType() AnyOfmicrosoftGraphApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *WindowsInformationProtectionAppLearningSummary) GetApplicationTypeOk() (AnyOfmicrosoftGraphApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationType

`func (o *WindowsInformationProtectionAppLearningSummary) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### SetApplicationType

`func (o *WindowsInformationProtectionAppLearningSummary) SetApplicationType(v AnyOfmicrosoftGraphApplicationType)`

SetApplicationType gets a reference to the given AnyOfmicrosoftGraphApplicationType and assigns it to the ApplicationType field.

### GetDeviceCount

`func (o *WindowsInformationProtectionAppLearningSummary) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *WindowsInformationProtectionAppLearningSummary) GetDeviceCountOk() (int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCount

`func (o *WindowsInformationProtectionAppLearningSummary) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### SetDeviceCount

`func (o *WindowsInformationProtectionAppLearningSummary) SetDeviceCount(v int32)`

SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


