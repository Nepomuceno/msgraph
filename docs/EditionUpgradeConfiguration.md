# EditionUpgradeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseType** | Pointer to [**AnyOfmicrosoftGraphEditionUpgradeLicenseType**](anyOf&lt;microsoft.graph.editionUpgradeLicenseType&gt;.md) | Edition Upgrade License Type. | [optional] 
**TargetEdition** | Pointer to [**AnyOfmicrosoftGraphWindows10EditionType**](anyOf&lt;microsoft.graph.windows10EditionType&gt;.md) | Edition Upgrade Target Edition. | [optional] 
**License** | Pointer to **string** | Edition Upgrade License File Content. | [optional] 
**ProductKey** | Pointer to **string** | Edition Upgrade Product Key. | [optional] 

## Methods

### GetLicenseType

`func (o *EditionUpgradeConfiguration) GetLicenseType() AnyOfmicrosoftGraphEditionUpgradeLicenseType`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *EditionUpgradeConfiguration) GetLicenseTypeOk() (AnyOfmicrosoftGraphEditionUpgradeLicenseType, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicenseType

`func (o *EditionUpgradeConfiguration) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### SetLicenseType

`func (o *EditionUpgradeConfiguration) SetLicenseType(v AnyOfmicrosoftGraphEditionUpgradeLicenseType)`

SetLicenseType gets a reference to the given AnyOfmicrosoftGraphEditionUpgradeLicenseType and assigns it to the LicenseType field.

### GetTargetEdition

`func (o *EditionUpgradeConfiguration) GetTargetEdition() AnyOfmicrosoftGraphWindows10EditionType`

GetTargetEdition returns the TargetEdition field if non-nil, zero value otherwise.

### GetTargetEditionOk

`func (o *EditionUpgradeConfiguration) GetTargetEditionOk() (AnyOfmicrosoftGraphWindows10EditionType, bool)`

GetTargetEditionOk returns a tuple with the TargetEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetEdition

`func (o *EditionUpgradeConfiguration) HasTargetEdition() bool`

HasTargetEdition returns a boolean if a field has been set.

### SetTargetEdition

`func (o *EditionUpgradeConfiguration) SetTargetEdition(v AnyOfmicrosoftGraphWindows10EditionType)`

SetTargetEdition gets a reference to the given AnyOfmicrosoftGraphWindows10EditionType and assigns it to the TargetEdition field.

### GetLicense

`func (o *EditionUpgradeConfiguration) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *EditionUpgradeConfiguration) GetLicenseOk() (string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicense

`func (o *EditionUpgradeConfiguration) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### SetLicense

`func (o *EditionUpgradeConfiguration) SetLicense(v string)`

SetLicense gets a reference to the given string and assigns it to the License field.

### SetLicenseExplicitNull

`func (o *EditionUpgradeConfiguration) SetLicenseExplicitNull(b bool)`

SetLicenseExplicitNull (un)sets License to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The License value is set to nil even if false is passed
### GetProductKey

`func (o *EditionUpgradeConfiguration) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *EditionUpgradeConfiguration) GetProductKeyOk() (string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductKey

`func (o *EditionUpgradeConfiguration) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### SetProductKey

`func (o *EditionUpgradeConfiguration) SetProductKey(v string)`

SetProductKey gets a reference to the given string and assigns it to the ProductKey field.

### SetProductKeyExplicitNull

`func (o *EditionUpgradeConfiguration) SetProductKeyExplicitNull(b bool)`

SetProductKeyExplicitNull (un)sets ProductKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProductKey value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


