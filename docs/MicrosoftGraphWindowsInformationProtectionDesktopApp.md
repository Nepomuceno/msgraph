# MicrosoftGraphWindowsInformationProtectionDesktopApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | App display name. | [optional] 
**Description** | Pointer to **string** | The app&#39;s description. | [optional] 
**PublisherName** | Pointer to **string** | The publisher name | [optional] 
**ProductName** | Pointer to **string** | The product name. | [optional] 
**Denied** | Pointer to **bool** | If true, app is denied protection or exemption. | [optional] 
**BinaryName** | Pointer to **string** | The binary name. | [optional] 
**BinaryVersionLow** | Pointer to **string** | The lower binary version. | [optional] 
**BinaryVersionHigh** | Pointer to **string** | The high binary version. | [optional] 

## Methods

### GetDisplayName

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetPublisherName

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetPublisherName() string`

GetPublisherName returns the PublisherName field if non-nil, zero value otherwise.

### GetPublisherNameOk

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetPublisherNameOk() (string, bool)`

GetPublisherNameOk returns a tuple with the PublisherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublisherName

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasPublisherName() bool`

HasPublisherName returns a boolean if a field has been set.

### SetPublisherName

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetPublisherName(v string)`

SetPublisherName gets a reference to the given string and assigns it to the PublisherName field.

### SetPublisherNameExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetPublisherNameExplicitNull(b bool)`

SetPublisherNameExplicitNull (un)sets PublisherName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PublisherName value is set to nil even if false is passed
### GetProductName

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetProductNameOk() (string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductName

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductName

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetProductName(v string)`

SetProductName gets a reference to the given string and assigns it to the ProductName field.

### SetProductNameExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetProductNameExplicitNull(b bool)`

SetProductNameExplicitNull (un)sets ProductName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProductName value is set to nil even if false is passed
### GetDenied

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetDenied() bool`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetDeniedOk() (bool, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDenied

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasDenied() bool`

HasDenied returns a boolean if a field has been set.

### SetDenied

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetDenied(v bool)`

SetDenied gets a reference to the given bool and assigns it to the Denied field.

### GetBinaryName

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetBinaryNameOk() (string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBinaryName

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasBinaryName() bool`

HasBinaryName returns a boolean if a field has been set.

### SetBinaryName

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetBinaryName(v string)`

SetBinaryName gets a reference to the given string and assigns it to the BinaryName field.

### GetBinaryVersionLow

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetBinaryVersionLow() string`

GetBinaryVersionLow returns the BinaryVersionLow field if non-nil, zero value otherwise.

### GetBinaryVersionLowOk

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetBinaryVersionLowOk() (string, bool)`

GetBinaryVersionLowOk returns a tuple with the BinaryVersionLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBinaryVersionLow

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasBinaryVersionLow() bool`

HasBinaryVersionLow returns a boolean if a field has been set.

### SetBinaryVersionLow

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetBinaryVersionLow(v string)`

SetBinaryVersionLow gets a reference to the given string and assigns it to the BinaryVersionLow field.

### SetBinaryVersionLowExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetBinaryVersionLowExplicitNull(b bool)`

SetBinaryVersionLowExplicitNull (un)sets BinaryVersionLow to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BinaryVersionLow value is set to nil even if false is passed
### GetBinaryVersionHigh

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetBinaryVersionHigh() string`

GetBinaryVersionHigh returns the BinaryVersionHigh field if non-nil, zero value otherwise.

### GetBinaryVersionHighOk

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetBinaryVersionHighOk() (string, bool)`

GetBinaryVersionHighOk returns a tuple with the BinaryVersionHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBinaryVersionHigh

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasBinaryVersionHigh() bool`

HasBinaryVersionHigh returns a boolean if a field has been set.

### SetBinaryVersionHigh

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetBinaryVersionHigh(v string)`

SetBinaryVersionHigh gets a reference to the given string and assigns it to the BinaryVersionHigh field.

### SetBinaryVersionHighExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetBinaryVersionHighExplicitNull(b bool)`

SetBinaryVersionHighExplicitNull (un)sets BinaryVersionHigh to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BinaryVersionHigh value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


