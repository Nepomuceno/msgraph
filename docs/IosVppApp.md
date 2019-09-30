# IosVppApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedLicenseCount** | Pointer to **int32** | The number of VPP licenses in use. | [optional] 
**TotalLicenseCount** | Pointer to **int32** | The total number of VPP licenses. | [optional] 
**ReleaseDateTime** | Pointer to [**time.Time**](time.Time.md) | The VPP application release date and time. | [optional] 
**AppStoreUrl** | Pointer to **string** | The store URL. | [optional] 
**LicensingType** | Pointer to [**AnyOfmicrosoftGraphVppLicensingType**](anyOf&lt;microsoft.graph.vppLicensingType&gt;.md) | The supported License Type. | [optional] 
**ApplicableDeviceType** | Pointer to [**AnyOfmicrosoftGraphIosDeviceType**](anyOf&lt;microsoft.graph.iosDeviceType&gt;.md) | The applicable iOS Device Type. | [optional] 
**VppTokenOrganizationName** | Pointer to **string** | The organization associated with the Apple Volume Purchase Program Token | [optional] 
**VppTokenAccountType** | Pointer to [**AnyOfmicrosoftGraphVppTokenAccountType**](anyOf&lt;microsoft.graph.vppTokenAccountType&gt;.md) | The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: &#x60;business&#x60;, &#x60;education&#x60;. | [optional] 
**VppTokenAppleId** | Pointer to **string** | The Apple Id associated with the given Apple Volume Purchase Program Token. | [optional] 
**BundleId** | Pointer to **string** | The Identity Name. | [optional] 

## Methods

### GetUsedLicenseCount

`func (o *IosVppApp) GetUsedLicenseCount() int32`

GetUsedLicenseCount returns the UsedLicenseCount field if non-nil, zero value otherwise.

### GetUsedLicenseCountOk

`func (o *IosVppApp) GetUsedLicenseCountOk() (int32, bool)`

GetUsedLicenseCountOk returns a tuple with the UsedLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsedLicenseCount

`func (o *IosVppApp) HasUsedLicenseCount() bool`

HasUsedLicenseCount returns a boolean if a field has been set.

### SetUsedLicenseCount

`func (o *IosVppApp) SetUsedLicenseCount(v int32)`

SetUsedLicenseCount gets a reference to the given int32 and assigns it to the UsedLicenseCount field.

### GetTotalLicenseCount

`func (o *IosVppApp) GetTotalLicenseCount() int32`

GetTotalLicenseCount returns the TotalLicenseCount field if non-nil, zero value otherwise.

### GetTotalLicenseCountOk

`func (o *IosVppApp) GetTotalLicenseCountOk() (int32, bool)`

GetTotalLicenseCountOk returns a tuple with the TotalLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalLicenseCount

`func (o *IosVppApp) HasTotalLicenseCount() bool`

HasTotalLicenseCount returns a boolean if a field has been set.

### SetTotalLicenseCount

`func (o *IosVppApp) SetTotalLicenseCount(v int32)`

SetTotalLicenseCount gets a reference to the given int32 and assigns it to the TotalLicenseCount field.

### GetReleaseDateTime

`func (o *IosVppApp) GetReleaseDateTime() time.Time`

GetReleaseDateTime returns the ReleaseDateTime field if non-nil, zero value otherwise.

### GetReleaseDateTimeOk

`func (o *IosVppApp) GetReleaseDateTimeOk() (time.Time, bool)`

GetReleaseDateTimeOk returns a tuple with the ReleaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReleaseDateTime

`func (o *IosVppApp) HasReleaseDateTime() bool`

HasReleaseDateTime returns a boolean if a field has been set.

### SetReleaseDateTime

`func (o *IosVppApp) SetReleaseDateTime(v time.Time)`

SetReleaseDateTime gets a reference to the given time.Time and assigns it to the ReleaseDateTime field.

### SetReleaseDateTimeExplicitNull

`func (o *IosVppApp) SetReleaseDateTimeExplicitNull(b bool)`

SetReleaseDateTimeExplicitNull (un)sets ReleaseDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReleaseDateTime value is set to nil even if false is passed
### GetAppStoreUrl

`func (o *IosVppApp) GetAppStoreUrl() string`

GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.

### GetAppStoreUrlOk

`func (o *IosVppApp) GetAppStoreUrlOk() (string, bool)`

GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppStoreUrl

`func (o *IosVppApp) HasAppStoreUrl() bool`

HasAppStoreUrl returns a boolean if a field has been set.

### SetAppStoreUrl

`func (o *IosVppApp) SetAppStoreUrl(v string)`

SetAppStoreUrl gets a reference to the given string and assigns it to the AppStoreUrl field.

### SetAppStoreUrlExplicitNull

`func (o *IosVppApp) SetAppStoreUrlExplicitNull(b bool)`

SetAppStoreUrlExplicitNull (un)sets AppStoreUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppStoreUrl value is set to nil even if false is passed
### GetLicensingType

`func (o *IosVppApp) GetLicensingType() AnyOfmicrosoftGraphVppLicensingType`

GetLicensingType returns the LicensingType field if non-nil, zero value otherwise.

### GetLicensingTypeOk

`func (o *IosVppApp) GetLicensingTypeOk() (AnyOfmicrosoftGraphVppLicensingType, bool)`

GetLicensingTypeOk returns a tuple with the LicensingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicensingType

`func (o *IosVppApp) HasLicensingType() bool`

HasLicensingType returns a boolean if a field has been set.

### SetLicensingType

`func (o *IosVppApp) SetLicensingType(v AnyOfmicrosoftGraphVppLicensingType)`

SetLicensingType gets a reference to the given AnyOfmicrosoftGraphVppLicensingType and assigns it to the LicensingType field.

### SetLicensingTypeExplicitNull

`func (o *IosVppApp) SetLicensingTypeExplicitNull(b bool)`

SetLicensingTypeExplicitNull (un)sets LicensingType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LicensingType value is set to nil even if false is passed
### GetApplicableDeviceType

`func (o *IosVppApp) GetApplicableDeviceType() AnyOfmicrosoftGraphIosDeviceType`

GetApplicableDeviceType returns the ApplicableDeviceType field if non-nil, zero value otherwise.

### GetApplicableDeviceTypeOk

`func (o *IosVppApp) GetApplicableDeviceTypeOk() (AnyOfmicrosoftGraphIosDeviceType, bool)`

GetApplicableDeviceTypeOk returns a tuple with the ApplicableDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicableDeviceType

`func (o *IosVppApp) HasApplicableDeviceType() bool`

HasApplicableDeviceType returns a boolean if a field has been set.

### SetApplicableDeviceType

`func (o *IosVppApp) SetApplicableDeviceType(v AnyOfmicrosoftGraphIosDeviceType)`

SetApplicableDeviceType gets a reference to the given AnyOfmicrosoftGraphIosDeviceType and assigns it to the ApplicableDeviceType field.

### SetApplicableDeviceTypeExplicitNull

`func (o *IosVppApp) SetApplicableDeviceTypeExplicitNull(b bool)`

SetApplicableDeviceTypeExplicitNull (un)sets ApplicableDeviceType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ApplicableDeviceType value is set to nil even if false is passed
### GetVppTokenOrganizationName

`func (o *IosVppApp) GetVppTokenOrganizationName() string`

GetVppTokenOrganizationName returns the VppTokenOrganizationName field if non-nil, zero value otherwise.

### GetVppTokenOrganizationNameOk

`func (o *IosVppApp) GetVppTokenOrganizationNameOk() (string, bool)`

GetVppTokenOrganizationNameOk returns a tuple with the VppTokenOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVppTokenOrganizationName

`func (o *IosVppApp) HasVppTokenOrganizationName() bool`

HasVppTokenOrganizationName returns a boolean if a field has been set.

### SetVppTokenOrganizationName

`func (o *IosVppApp) SetVppTokenOrganizationName(v string)`

SetVppTokenOrganizationName gets a reference to the given string and assigns it to the VppTokenOrganizationName field.

### SetVppTokenOrganizationNameExplicitNull

`func (o *IosVppApp) SetVppTokenOrganizationNameExplicitNull(b bool)`

SetVppTokenOrganizationNameExplicitNull (un)sets VppTokenOrganizationName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VppTokenOrganizationName value is set to nil even if false is passed
### GetVppTokenAccountType

`func (o *IosVppApp) GetVppTokenAccountType() AnyOfmicrosoftGraphVppTokenAccountType`

GetVppTokenAccountType returns the VppTokenAccountType field if non-nil, zero value otherwise.

### GetVppTokenAccountTypeOk

`func (o *IosVppApp) GetVppTokenAccountTypeOk() (AnyOfmicrosoftGraphVppTokenAccountType, bool)`

GetVppTokenAccountTypeOk returns a tuple with the VppTokenAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVppTokenAccountType

`func (o *IosVppApp) HasVppTokenAccountType() bool`

HasVppTokenAccountType returns a boolean if a field has been set.

### SetVppTokenAccountType

`func (o *IosVppApp) SetVppTokenAccountType(v AnyOfmicrosoftGraphVppTokenAccountType)`

SetVppTokenAccountType gets a reference to the given AnyOfmicrosoftGraphVppTokenAccountType and assigns it to the VppTokenAccountType field.

### GetVppTokenAppleId

`func (o *IosVppApp) GetVppTokenAppleId() string`

GetVppTokenAppleId returns the VppTokenAppleId field if non-nil, zero value otherwise.

### GetVppTokenAppleIdOk

`func (o *IosVppApp) GetVppTokenAppleIdOk() (string, bool)`

GetVppTokenAppleIdOk returns a tuple with the VppTokenAppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVppTokenAppleId

`func (o *IosVppApp) HasVppTokenAppleId() bool`

HasVppTokenAppleId returns a boolean if a field has been set.

### SetVppTokenAppleId

`func (o *IosVppApp) SetVppTokenAppleId(v string)`

SetVppTokenAppleId gets a reference to the given string and assigns it to the VppTokenAppleId field.

### SetVppTokenAppleIdExplicitNull

`func (o *IosVppApp) SetVppTokenAppleIdExplicitNull(b bool)`

SetVppTokenAppleIdExplicitNull (un)sets VppTokenAppleId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VppTokenAppleId value is set to nil even if false is passed
### GetBundleId

`func (o *IosVppApp) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *IosVppApp) GetBundleIdOk() (string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleId

`func (o *IosVppApp) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### SetBundleId

`func (o *IosVppApp) SetBundleId(v string)`

SetBundleId gets a reference to the given string and assigns it to the BundleId field.

### SetBundleIdExplicitNull

`func (o *IosVppApp) SetBundleIdExplicitNull(b bool)`

SetBundleIdExplicitNull (un)sets BundleId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BundleId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


