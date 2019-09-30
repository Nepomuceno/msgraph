# OmaSettingStringXml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | File name associated with the Value property (*.xml). | [optional] 
**Value** | Pointer to **string** | Value. (UTF8 encoded byte array) | [optional] 

## Methods

### GetFileName

`func (o *OmaSettingStringXml) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *OmaSettingStringXml) GetFileNameOk() (string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileName

`func (o *OmaSettingStringXml) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileName

`func (o *OmaSettingStringXml) SetFileName(v string)`

SetFileName gets a reference to the given string and assigns it to the FileName field.

### SetFileNameExplicitNull

`func (o *OmaSettingStringXml) SetFileNameExplicitNull(b bool)`

SetFileNameExplicitNull (un)sets FileName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileName value is set to nil even if false is passed
### GetValue

`func (o *OmaSettingStringXml) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OmaSettingStringXml) GetValueOk() (string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *OmaSettingStringXml) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *OmaSettingStringXml) SetValue(v string)`

SetValue gets a reference to the given string and assigns it to the Value field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


