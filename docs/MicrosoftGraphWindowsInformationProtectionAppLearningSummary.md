# MicrosoftGraphWindowsInformationProtectionAppLearningSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ApplicationName** | Pointer to **string** | Application Name | [optional] 
**ApplicationType** | Pointer to [**AnyOfmicrosoftGraphApplicationType**](anyOf&lt;microsoft.graph.applicationType&gt;.md) | Application Type | [optional] 
**DeviceCount** | Pointer to **int32** | Device Count | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetApplicationName

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetApplicationNameOk() (string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationName

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### SetApplicationName

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) SetApplicationName(v string)`

SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.

### SetApplicationNameExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) SetApplicationNameExplicitNull(b bool)`

SetApplicationNameExplicitNull (un)sets ApplicationName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ApplicationName value is set to nil even if false is passed
### GetApplicationType

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetApplicationType() AnyOfmicrosoftGraphApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetApplicationTypeOk() (AnyOfmicrosoftGraphApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationType

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### SetApplicationType

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) SetApplicationType(v AnyOfmicrosoftGraphApplicationType)`

SetApplicationType gets a reference to the given AnyOfmicrosoftGraphApplicationType and assigns it to the ApplicationType field.

### GetDeviceCount

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) GetDeviceCountOk() (int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCount

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### SetDeviceCount

`func (o *MicrosoftGraphWindowsInformationProtectionAppLearningSummary) SetDeviceCount(v int32)`

SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


