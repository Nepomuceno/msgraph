# OmaSettingBase64

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | File name associated with the Value property (*.cer | *.crt | *.p7b | *.bin). | [optional] 
**Value** | Pointer to **string** | Value. (Base64 encoded string) | [optional] 

## Methods

### GetFileName

`func (o *OmaSettingBase64) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *OmaSettingBase64) GetFileNameOk() (string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileName

`func (o *OmaSettingBase64) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileName

`func (o *OmaSettingBase64) SetFileName(v string)`

SetFileName gets a reference to the given string and assigns it to the FileName field.

### SetFileNameExplicitNull

`func (o *OmaSettingBase64) SetFileNameExplicitNull(b bool)`

SetFileNameExplicitNull (un)sets FileName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileName value is set to nil even if false is passed
### GetValue

`func (o *OmaSettingBase64) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OmaSettingBase64) GetValueOk() (string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *OmaSettingBase64) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *OmaSettingBase64) SetValue(v string)`

SetValue gets a reference to the given string and assigns it to the Value field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


