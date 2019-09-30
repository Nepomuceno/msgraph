# WindowsUniversalAppX

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicableArchitectures** | Pointer to [**AnyOfmicrosoftGraphWindowsArchitecture**](anyOf&lt;microsoft.graph.windowsArchitecture&gt;.md) | The Windows architecture(s) for which this app can run on. | [optional] 
**ApplicableDeviceTypes** | Pointer to [**AnyOfmicrosoftGraphWindowsDeviceType**](anyOf&lt;microsoft.graph.windowsDeviceType&gt;.md) | The Windows device type(s) for which this app can run on. | [optional] 
**IdentityName** | Pointer to **string** | The Identity Name. | [optional] 
**IdentityPublisherHash** | Pointer to **string** | The Identity Publisher Hash. | [optional] 
**IdentityResourceIdentifier** | Pointer to **string** | The Identity Resource Identifier. | [optional] 
**IsBundle** | Pointer to **bool** | Whether or not the app is a bundle. | [optional] 
**MinimumSupportedOperatingSystem** | Pointer to [**MicrosoftGraphWindowsMinimumOperatingSystem**](microsoft.graph.windowsMinimumOperatingSystem.md) |  | [optional] 
**IdentityVersion** | Pointer to **string** | The identity version. | [optional] 

## Methods

### GetApplicableArchitectures

`func (o *WindowsUniversalAppX) GetApplicableArchitectures() AnyOfmicrosoftGraphWindowsArchitecture`

GetApplicableArchitectures returns the ApplicableArchitectures field if non-nil, zero value otherwise.

### GetApplicableArchitecturesOk

`func (o *WindowsUniversalAppX) GetApplicableArchitecturesOk() (AnyOfmicrosoftGraphWindowsArchitecture, bool)`

GetApplicableArchitecturesOk returns a tuple with the ApplicableArchitectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicableArchitectures

`func (o *WindowsUniversalAppX) HasApplicableArchitectures() bool`

HasApplicableArchitectures returns a boolean if a field has been set.

### SetApplicableArchitectures

`func (o *WindowsUniversalAppX) SetApplicableArchitectures(v AnyOfmicrosoftGraphWindowsArchitecture)`

SetApplicableArchitectures gets a reference to the given AnyOfmicrosoftGraphWindowsArchitecture and assigns it to the ApplicableArchitectures field.

### GetApplicableDeviceTypes

`func (o *WindowsUniversalAppX) GetApplicableDeviceTypes() AnyOfmicrosoftGraphWindowsDeviceType`

GetApplicableDeviceTypes returns the ApplicableDeviceTypes field if non-nil, zero value otherwise.

### GetApplicableDeviceTypesOk

`func (o *WindowsUniversalAppX) GetApplicableDeviceTypesOk() (AnyOfmicrosoftGraphWindowsDeviceType, bool)`

GetApplicableDeviceTypesOk returns a tuple with the ApplicableDeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicableDeviceTypes

`func (o *WindowsUniversalAppX) HasApplicableDeviceTypes() bool`

HasApplicableDeviceTypes returns a boolean if a field has been set.

### SetApplicableDeviceTypes

`func (o *WindowsUniversalAppX) SetApplicableDeviceTypes(v AnyOfmicrosoftGraphWindowsDeviceType)`

SetApplicableDeviceTypes gets a reference to the given AnyOfmicrosoftGraphWindowsDeviceType and assigns it to the ApplicableDeviceTypes field.

### GetIdentityName

`func (o *WindowsUniversalAppX) GetIdentityName() string`

GetIdentityName returns the IdentityName field if non-nil, zero value otherwise.

### GetIdentityNameOk

`func (o *WindowsUniversalAppX) GetIdentityNameOk() (string, bool)`

GetIdentityNameOk returns a tuple with the IdentityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentityName

`func (o *WindowsUniversalAppX) HasIdentityName() bool`

HasIdentityName returns a boolean if a field has been set.

### SetIdentityName

`func (o *WindowsUniversalAppX) SetIdentityName(v string)`

SetIdentityName gets a reference to the given string and assigns it to the IdentityName field.

### SetIdentityNameExplicitNull

`func (o *WindowsUniversalAppX) SetIdentityNameExplicitNull(b bool)`

SetIdentityNameExplicitNull (un)sets IdentityName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IdentityName value is set to nil even if false is passed
### GetIdentityPublisherHash

`func (o *WindowsUniversalAppX) GetIdentityPublisherHash() string`

GetIdentityPublisherHash returns the IdentityPublisherHash field if non-nil, zero value otherwise.

### GetIdentityPublisherHashOk

`func (o *WindowsUniversalAppX) GetIdentityPublisherHashOk() (string, bool)`

GetIdentityPublisherHashOk returns a tuple with the IdentityPublisherHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentityPublisherHash

`func (o *WindowsUniversalAppX) HasIdentityPublisherHash() bool`

HasIdentityPublisherHash returns a boolean if a field has been set.

### SetIdentityPublisherHash

`func (o *WindowsUniversalAppX) SetIdentityPublisherHash(v string)`

SetIdentityPublisherHash gets a reference to the given string and assigns it to the IdentityPublisherHash field.

### GetIdentityResourceIdentifier

`func (o *WindowsUniversalAppX) GetIdentityResourceIdentifier() string`

GetIdentityResourceIdentifier returns the IdentityResourceIdentifier field if non-nil, zero value otherwise.

### GetIdentityResourceIdentifierOk

`func (o *WindowsUniversalAppX) GetIdentityResourceIdentifierOk() (string, bool)`

GetIdentityResourceIdentifierOk returns a tuple with the IdentityResourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentityResourceIdentifier

`func (o *WindowsUniversalAppX) HasIdentityResourceIdentifier() bool`

HasIdentityResourceIdentifier returns a boolean if a field has been set.

### SetIdentityResourceIdentifier

`func (o *WindowsUniversalAppX) SetIdentityResourceIdentifier(v string)`

SetIdentityResourceIdentifier gets a reference to the given string and assigns it to the IdentityResourceIdentifier field.

### SetIdentityResourceIdentifierExplicitNull

`func (o *WindowsUniversalAppX) SetIdentityResourceIdentifierExplicitNull(b bool)`

SetIdentityResourceIdentifierExplicitNull (un)sets IdentityResourceIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IdentityResourceIdentifier value is set to nil even if false is passed
### GetIsBundle

`func (o *WindowsUniversalAppX) GetIsBundle() bool`

GetIsBundle returns the IsBundle field if non-nil, zero value otherwise.

### GetIsBundleOk

`func (o *WindowsUniversalAppX) GetIsBundleOk() (bool, bool)`

GetIsBundleOk returns a tuple with the IsBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsBundle

`func (o *WindowsUniversalAppX) HasIsBundle() bool`

HasIsBundle returns a boolean if a field has been set.

### SetIsBundle

`func (o *WindowsUniversalAppX) SetIsBundle(v bool)`

SetIsBundle gets a reference to the given bool and assigns it to the IsBundle field.

### GetMinimumSupportedOperatingSystem

`func (o *WindowsUniversalAppX) GetMinimumSupportedOperatingSystem() MicrosoftGraphWindowsMinimumOperatingSystem`

GetMinimumSupportedOperatingSystem returns the MinimumSupportedOperatingSystem field if non-nil, zero value otherwise.

### GetMinimumSupportedOperatingSystemOk

`func (o *WindowsUniversalAppX) GetMinimumSupportedOperatingSystemOk() (MicrosoftGraphWindowsMinimumOperatingSystem, bool)`

GetMinimumSupportedOperatingSystemOk returns a tuple with the MinimumSupportedOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumSupportedOperatingSystem

`func (o *WindowsUniversalAppX) HasMinimumSupportedOperatingSystem() bool`

HasMinimumSupportedOperatingSystem returns a boolean if a field has been set.

### SetMinimumSupportedOperatingSystem

`func (o *WindowsUniversalAppX) SetMinimumSupportedOperatingSystem(v MicrosoftGraphWindowsMinimumOperatingSystem)`

SetMinimumSupportedOperatingSystem gets a reference to the given MicrosoftGraphWindowsMinimumOperatingSystem and assigns it to the MinimumSupportedOperatingSystem field.

### GetIdentityVersion

`func (o *WindowsUniversalAppX) GetIdentityVersion() string`

GetIdentityVersion returns the IdentityVersion field if non-nil, zero value otherwise.

### GetIdentityVersionOk

`func (o *WindowsUniversalAppX) GetIdentityVersionOk() (string, bool)`

GetIdentityVersionOk returns a tuple with the IdentityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdentityVersion

`func (o *WindowsUniversalAppX) HasIdentityVersion() bool`

HasIdentityVersion returns a boolean if a field has been set.

### SetIdentityVersion

`func (o *WindowsUniversalAppX) SetIdentityVersion(v string)`

SetIdentityVersion gets a reference to the given string and assigns it to the IdentityVersion field.

### SetIdentityVersionExplicitNull

`func (o *WindowsUniversalAppX) SetIdentityVersionExplicitNull(b bool)`

SetIdentityVersionExplicitNull (un)sets IdentityVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IdentityVersion value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


