# DomainDnsMxRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailExchange** | Pointer to **string** |  | [optional] 
**Preference** | Pointer to **int32** |  | [optional] 

## Methods

### GetMailExchange

`func (o *DomainDnsMxRecord) GetMailExchange() string`

GetMailExchange returns the MailExchange field if non-nil, zero value otherwise.

### GetMailExchangeOk

`func (o *DomainDnsMxRecord) GetMailExchangeOk() (string, bool)`

GetMailExchangeOk returns a tuple with the MailExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailExchange

`func (o *DomainDnsMxRecord) HasMailExchange() bool`

HasMailExchange returns a boolean if a field has been set.

### SetMailExchange

`func (o *DomainDnsMxRecord) SetMailExchange(v string)`

SetMailExchange gets a reference to the given string and assigns it to the MailExchange field.

### GetPreference

`func (o *DomainDnsMxRecord) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *DomainDnsMxRecord) GetPreferenceOk() (int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreference

`func (o *DomainDnsMxRecord) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### SetPreference

`func (o *DomainDnsMxRecord) SetPreference(v int32)`

SetPreference gets a reference to the given int32 and assigns it to the Preference field.

### SetPreferenceExplicitNull

`func (o *DomainDnsMxRecord) SetPreferenceExplicitNull(b bool)`

SetPreferenceExplicitNull (un)sets Preference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Preference value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


