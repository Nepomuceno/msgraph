# AndroidStoreApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageId** | Pointer to **string** | The package identifier. | [optional] 
**AppStoreUrl** | Pointer to **string** | The Android app store URL. | [optional] 
**MinimumSupportedOperatingSystem** | Pointer to [**AnyOfmicrosoftGraphAndroidMinimumOperatingSystem**](anyOf&lt;microsoft.graph.androidMinimumOperatingSystem&gt;.md) | The value for the minimum applicable operating system. | [optional] 

## Methods

### GetPackageId

`func (o *AndroidStoreApp) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *AndroidStoreApp) GetPackageIdOk() (string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPackageId

`func (o *AndroidStoreApp) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### SetPackageId

`func (o *AndroidStoreApp) SetPackageId(v string)`

SetPackageId gets a reference to the given string and assigns it to the PackageId field.

### SetPackageIdExplicitNull

`func (o *AndroidStoreApp) SetPackageIdExplicitNull(b bool)`

SetPackageIdExplicitNull (un)sets PackageId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PackageId value is set to nil even if false is passed
### GetAppStoreUrl

`func (o *AndroidStoreApp) GetAppStoreUrl() string`

GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.

### GetAppStoreUrlOk

`func (o *AndroidStoreApp) GetAppStoreUrlOk() (string, bool)`

GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreUrl

`func (o *AndroidStoreApp) HasAppStoreUrl() bool`

HasAppStoreUrl returns a boolean if a field has been set.

### SetAppStoreUrl

`func (o *AndroidStoreApp) SetAppStoreUrl(v string)`

SetAppStoreUrl gets a reference to the given string and assigns it to the AppStoreUrl field.

### SetAppStoreUrlExplicitNull

`func (o *AndroidStoreApp) SetAppStoreUrlExplicitNull(b bool)`

SetAppStoreUrlExplicitNull (un)sets AppStoreUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppStoreUrl value is set to nil even if false is passed
### GetMinimumSupportedOperatingSystem

`func (o *AndroidStoreApp) GetMinimumSupportedOperatingSystem() AnyOfmicrosoftGraphAndroidMinimumOperatingSystem`

GetMinimumSupportedOperatingSystem returns the MinimumSupportedOperatingSystem field if non-nil, zero value otherwise.

### GetMinimumSupportedOperatingSystemOk

`func (o *AndroidStoreApp) GetMinimumSupportedOperatingSystemOk() (AnyOfmicrosoftGraphAndroidMinimumOperatingSystem, bool)`

GetMinimumSupportedOperatingSystemOk returns a tuple with the MinimumSupportedOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumSupportedOperatingSystem

`func (o *AndroidStoreApp) HasMinimumSupportedOperatingSystem() bool`

HasMinimumSupportedOperatingSystem returns a boolean if a field has been set.

### SetMinimumSupportedOperatingSystem

`func (o *AndroidStoreApp) SetMinimumSupportedOperatingSystem(v AnyOfmicrosoftGraphAndroidMinimumOperatingSystem)`

SetMinimumSupportedOperatingSystem gets a reference to the given AnyOfmicrosoftGraphAndroidMinimumOperatingSystem and assigns it to the MinimumSupportedOperatingSystem field.

### SetMinimumSupportedOperatingSystemExplicitNull

`func (o *AndroidStoreApp) SetMinimumSupportedOperatingSystemExplicitNull(b bool)`

SetMinimumSupportedOperatingSystemExplicitNull (un)sets MinimumSupportedOperatingSystem to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumSupportedOperatingSystem value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


