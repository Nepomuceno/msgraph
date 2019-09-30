# ManagedIosLobApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | Pointer to **string** | The Identity Name. | [optional] 
**ApplicableDeviceType** | Pointer to [**MicrosoftGraphIosDeviceType**](microsoft.graph.iosDeviceType.md) |  | [optional] 
**MinimumSupportedOperatingSystem** | Pointer to [**AnyOfmicrosoftGraphIosMinimumOperatingSystem**](anyOf&lt;microsoft.graph.iosMinimumOperatingSystem&gt;.md) | The value for the minimum applicable operating system. | [optional] 
**ExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) | The expiration time. | [optional] 
**VersionNumber** | Pointer to **string** | The version number of managed iOS Line of Business (LoB) app. | [optional] 
**BuildNumber** | Pointer to **string** | The build number of managed iOS Line of Business (LoB) app. | [optional] 

## Methods

### GetBundleId

`func (o *ManagedIosLobApp) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *ManagedIosLobApp) GetBundleIdOk() (string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleId

`func (o *ManagedIosLobApp) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### SetBundleId

`func (o *ManagedIosLobApp) SetBundleId(v string)`

SetBundleId gets a reference to the given string and assigns it to the BundleId field.

### SetBundleIdExplicitNull

`func (o *ManagedIosLobApp) SetBundleIdExplicitNull(b bool)`

SetBundleIdExplicitNull (un)sets BundleId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BundleId value is set to nil even if false is passed
### GetApplicableDeviceType

`func (o *ManagedIosLobApp) GetApplicableDeviceType() MicrosoftGraphIosDeviceType`

GetApplicableDeviceType returns the ApplicableDeviceType field if non-nil, zero value otherwise.

### GetApplicableDeviceTypeOk

`func (o *ManagedIosLobApp) GetApplicableDeviceTypeOk() (MicrosoftGraphIosDeviceType, bool)`

GetApplicableDeviceTypeOk returns a tuple with the ApplicableDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicableDeviceType

`func (o *ManagedIosLobApp) HasApplicableDeviceType() bool`

HasApplicableDeviceType returns a boolean if a field has been set.

### SetApplicableDeviceType

`func (o *ManagedIosLobApp) SetApplicableDeviceType(v MicrosoftGraphIosDeviceType)`

SetApplicableDeviceType gets a reference to the given MicrosoftGraphIosDeviceType and assigns it to the ApplicableDeviceType field.

### GetMinimumSupportedOperatingSystem

`func (o *ManagedIosLobApp) GetMinimumSupportedOperatingSystem() AnyOfmicrosoftGraphIosMinimumOperatingSystem`

GetMinimumSupportedOperatingSystem returns the MinimumSupportedOperatingSystem field if non-nil, zero value otherwise.

### GetMinimumSupportedOperatingSystemOk

`func (o *ManagedIosLobApp) GetMinimumSupportedOperatingSystemOk() (AnyOfmicrosoftGraphIosMinimumOperatingSystem, bool)`

GetMinimumSupportedOperatingSystemOk returns a tuple with the MinimumSupportedOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumSupportedOperatingSystem

`func (o *ManagedIosLobApp) HasMinimumSupportedOperatingSystem() bool`

HasMinimumSupportedOperatingSystem returns a boolean if a field has been set.

### SetMinimumSupportedOperatingSystem

`func (o *ManagedIosLobApp) SetMinimumSupportedOperatingSystem(v AnyOfmicrosoftGraphIosMinimumOperatingSystem)`

SetMinimumSupportedOperatingSystem gets a reference to the given AnyOfmicrosoftGraphIosMinimumOperatingSystem and assigns it to the MinimumSupportedOperatingSystem field.

### SetMinimumSupportedOperatingSystemExplicitNull

`func (o *ManagedIosLobApp) SetMinimumSupportedOperatingSystemExplicitNull(b bool)`

SetMinimumSupportedOperatingSystemExplicitNull (un)sets MinimumSupportedOperatingSystem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumSupportedOperatingSystem value is set to nil even if false is passed
### GetExpirationDateTime

`func (o *ManagedIosLobApp) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *ManagedIosLobApp) GetExpirationDateTimeOk() (time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpirationDateTime

`func (o *ManagedIosLobApp) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTime

`func (o *ManagedIosLobApp) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.

### SetExpirationDateTimeExplicitNull

`func (o *ManagedIosLobApp) SetExpirationDateTimeExplicitNull(b bool)`

SetExpirationDateTimeExplicitNull (un)sets ExpirationDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExpirationDateTime value is set to nil even if false is passed
### GetVersionNumber

`func (o *ManagedIosLobApp) GetVersionNumber() string`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *ManagedIosLobApp) GetVersionNumberOk() (string, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersionNumber

`func (o *ManagedIosLobApp) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### SetVersionNumber

`func (o *ManagedIosLobApp) SetVersionNumber(v string)`

SetVersionNumber gets a reference to the given string and assigns it to the VersionNumber field.

### SetVersionNumberExplicitNull

`func (o *ManagedIosLobApp) SetVersionNumberExplicitNull(b bool)`

SetVersionNumberExplicitNull (un)sets VersionNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VersionNumber value is set to nil even if false is passed
### GetBuildNumber

`func (o *ManagedIosLobApp) GetBuildNumber() string`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *ManagedIosLobApp) GetBuildNumberOk() (string, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBuildNumber

`func (o *ManagedIosLobApp) HasBuildNumber() bool`

HasBuildNumber returns a boolean if a field has been set.

### SetBuildNumber

`func (o *ManagedIosLobApp) SetBuildNumber(v string)`

SetBuildNumber gets a reference to the given string and assigns it to the BuildNumber field.

### SetBuildNumberExplicitNull

`func (o *ManagedIosLobApp) SetBuildNumberExplicitNull(b bool)`

SetBuildNumberExplicitNull (un)sets BuildNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BuildNumber value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


