# WindowsMobileMsi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandLine** | Pointer to **string** | The command line. | [optional] 
**ProductCode** | Pointer to **string** | The product code. | [optional] 
**ProductVersion** | Pointer to **string** | The product version of Windows Mobile MSI Line of Business (LoB) app. | [optional] 
**IgnoreVersionDetection** | Pointer to **bool** | A boolean to control whether the app&#39;s version will be used to detect the app after it is installed on a device. Set this to true for Windows Mobile MSI Line of Business (LoB) apps that use a self update feature. | [optional] 

## Methods

### GetCommandLine

`func (o *WindowsMobileMsi) GetCommandLine() string`

GetCommandLine returns the CommandLine field if non-nil, zero value otherwise.

### GetCommandLineOk

`func (o *WindowsMobileMsi) GetCommandLineOk() (string, bool)`

GetCommandLineOk returns a tuple with the CommandLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCommandLine

`func (o *WindowsMobileMsi) HasCommandLine() bool`

HasCommandLine returns a boolean if a field has been set.

### SetCommandLine

`func (o *WindowsMobileMsi) SetCommandLine(v string)`

SetCommandLine gets a reference to the given string and assigns it to the CommandLine field.

### SetCommandLineExplicitNull

`func (o *WindowsMobileMsi) SetCommandLineExplicitNull(b bool)`

SetCommandLineExplicitNull (un)sets CommandLine to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CommandLine value is set to nil even if false is passed
### GetProductCode

`func (o *WindowsMobileMsi) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *WindowsMobileMsi) GetProductCodeOk() (string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductCode

`func (o *WindowsMobileMsi) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### SetProductCode

`func (o *WindowsMobileMsi) SetProductCode(v string)`

SetProductCode gets a reference to the given string and assigns it to the ProductCode field.

### SetProductCodeExplicitNull

`func (o *WindowsMobileMsi) SetProductCodeExplicitNull(b bool)`

SetProductCodeExplicitNull (un)sets ProductCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProductCode value is set to nil even if false is passed
### GetProductVersion

`func (o *WindowsMobileMsi) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *WindowsMobileMsi) GetProductVersionOk() (string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductVersion

`func (o *WindowsMobileMsi) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### SetProductVersion

`func (o *WindowsMobileMsi) SetProductVersion(v string)`

SetProductVersion gets a reference to the given string and assigns it to the ProductVersion field.

### SetProductVersionExplicitNull

`func (o *WindowsMobileMsi) SetProductVersionExplicitNull(b bool)`

SetProductVersionExplicitNull (un)sets ProductVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProductVersion value is set to nil even if false is passed
### GetIgnoreVersionDetection

`func (o *WindowsMobileMsi) GetIgnoreVersionDetection() bool`

GetIgnoreVersionDetection returns the IgnoreVersionDetection field if non-nil, zero value otherwise.

### GetIgnoreVersionDetectionOk

`func (o *WindowsMobileMsi) GetIgnoreVersionDetectionOk() (bool, bool)`

GetIgnoreVersionDetectionOk returns a tuple with the IgnoreVersionDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIgnoreVersionDetection

`func (o *WindowsMobileMsi) HasIgnoreVersionDetection() bool`

HasIgnoreVersionDetection returns a boolean if a field has been set.

### SetIgnoreVersionDetection

`func (o *WindowsMobileMsi) SetIgnoreVersionDetection(v bool)`

SetIgnoreVersionDetection gets a reference to the given bool and assigns it to the IgnoreVersionDetection field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


