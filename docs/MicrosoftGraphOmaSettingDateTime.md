# MicrosoftGraphOmaSettingDateTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display Name. | [optional] 
**Description** | Pointer to **string** | Description. | [optional] 
**OmaUri** | Pointer to **string** | OMA. | [optional] 
**Value** | Pointer to [**time.Time**](time.Time.md) | Value. | [optional] 

## Methods

### GetDisplayName

`func (o *MicrosoftGraphOmaSettingDateTime) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphOmaSettingDateTime) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphOmaSettingDateTime) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphOmaSettingDateTime) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphOmaSettingDateTime) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphOmaSettingDateTime) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphOmaSettingDateTime) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphOmaSettingDateTime) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphOmaSettingDateTime) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetOmaUri

`func (o *MicrosoftGraphOmaSettingDateTime) GetOmaUri() string`

GetOmaUri returns the OmaUri field if non-nil, zero value otherwise.

### GetOmaUriOk

`func (o *MicrosoftGraphOmaSettingDateTime) GetOmaUriOk() (string, bool)`

GetOmaUriOk returns a tuple with the OmaUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOmaUri

`func (o *MicrosoftGraphOmaSettingDateTime) HasOmaUri() bool`

HasOmaUri returns a boolean if a field has been set.

### SetOmaUri

`func (o *MicrosoftGraphOmaSettingDateTime) SetOmaUri(v string)`

SetOmaUri gets a reference to the given string and assigns it to the OmaUri field.

### GetValue

`func (o *MicrosoftGraphOmaSettingDateTime) GetValue() time.Time`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphOmaSettingDateTime) GetValueOk() (time.Time, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *MicrosoftGraphOmaSettingDateTime) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *MicrosoftGraphOmaSettingDateTime) SetValue(v time.Time)`

SetValue gets a reference to the given time.Time and assigns it to the Value field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


