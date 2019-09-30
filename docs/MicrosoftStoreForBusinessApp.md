# MicrosoftStoreForBusinessApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedLicenseCount** | Pointer to **int32** | The number of Microsoft Store for Business licenses in use. | [optional] 
**TotalLicenseCount** | Pointer to **int32** | The total number of Microsoft Store for Business licenses. | [optional] 
**ProductKey** | Pointer to **string** | The app product key | [optional] 
**LicenseType** | Pointer to [**AnyOfmicrosoftGraphMicrosoftStoreForBusinessLicenseType**](anyOf&lt;microsoft.graph.microsoftStoreForBusinessLicenseType&gt;.md) | The app license type | [optional] 
**PackageIdentityName** | Pointer to **string** | The app package identifier | [optional] 

## Methods

### GetUsedLicenseCount

`func (o *MicrosoftStoreForBusinessApp) GetUsedLicenseCount() int32`

GetUsedLicenseCount returns the UsedLicenseCount field if non-nil, zero value otherwise.

### GetUsedLicenseCountOk

`func (o *MicrosoftStoreForBusinessApp) GetUsedLicenseCountOk() (int32, bool)`

GetUsedLicenseCountOk returns a tuple with the UsedLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsedLicenseCount

`func (o *MicrosoftStoreForBusinessApp) HasUsedLicenseCount() bool`

HasUsedLicenseCount returns a boolean if a field has been set.

### SetUsedLicenseCount

`func (o *MicrosoftStoreForBusinessApp) SetUsedLicenseCount(v int32)`

SetUsedLicenseCount gets a reference to the given int32 and assigns it to the UsedLicenseCount field.

### GetTotalLicenseCount

`func (o *MicrosoftStoreForBusinessApp) GetTotalLicenseCount() int32`

GetTotalLicenseCount returns the TotalLicenseCount field if non-nil, zero value otherwise.

### GetTotalLicenseCountOk

`func (o *MicrosoftStoreForBusinessApp) GetTotalLicenseCountOk() (int32, bool)`

GetTotalLicenseCountOk returns a tuple with the TotalLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalLicenseCount

`func (o *MicrosoftStoreForBusinessApp) HasTotalLicenseCount() bool`

HasTotalLicenseCount returns a boolean if a field has been set.

### SetTotalLicenseCount

`func (o *MicrosoftStoreForBusinessApp) SetTotalLicenseCount(v int32)`

SetTotalLicenseCount gets a reference to the given int32 and assigns it to the TotalLicenseCount field.

### GetProductKey

`func (o *MicrosoftStoreForBusinessApp) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *MicrosoftStoreForBusinessApp) GetProductKeyOk() (string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductKey

`func (o *MicrosoftStoreForBusinessApp) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### SetProductKey

`func (o *MicrosoftStoreForBusinessApp) SetProductKey(v string)`

SetProductKey gets a reference to the given string and assigns it to the ProductKey field.

### SetProductKeyExplicitNull

`func (o *MicrosoftStoreForBusinessApp) SetProductKeyExplicitNull(b bool)`

SetProductKeyExplicitNull (un)sets ProductKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProductKey value is set to nil even if false is passed
### GetLicenseType

`func (o *MicrosoftStoreForBusinessApp) GetLicenseType() AnyOfmicrosoftGraphMicrosoftStoreForBusinessLicenseType`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *MicrosoftStoreForBusinessApp) GetLicenseTypeOk() (AnyOfmicrosoftGraphMicrosoftStoreForBusinessLicenseType, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicenseType

`func (o *MicrosoftStoreForBusinessApp) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### SetLicenseType

`func (o *MicrosoftStoreForBusinessApp) SetLicenseType(v AnyOfmicrosoftGraphMicrosoftStoreForBusinessLicenseType)`

SetLicenseType gets a reference to the given AnyOfmicrosoftGraphMicrosoftStoreForBusinessLicenseType and assigns it to the LicenseType field.

### GetPackageIdentityName

`func (o *MicrosoftStoreForBusinessApp) GetPackageIdentityName() string`

GetPackageIdentityName returns the PackageIdentityName field if non-nil, zero value otherwise.

### GetPackageIdentityNameOk

`func (o *MicrosoftStoreForBusinessApp) GetPackageIdentityNameOk() (string, bool)`

GetPackageIdentityNameOk returns a tuple with the PackageIdentityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPackageIdentityName

`func (o *MicrosoftStoreForBusinessApp) HasPackageIdentityName() bool`

HasPackageIdentityName returns a boolean if a field has been set.

### SetPackageIdentityName

`func (o *MicrosoftStoreForBusinessApp) SetPackageIdentityName(v string)`

SetPackageIdentityName gets a reference to the given string and assigns it to the PackageIdentityName field.

### SetPackageIdentityNameExplicitNull

`func (o *MicrosoftStoreForBusinessApp) SetPackageIdentityNameExplicitNull(b bool)`

SetPackageIdentityNameExplicitNull (un)sets PackageIdentityName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PackageIdentityName value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


