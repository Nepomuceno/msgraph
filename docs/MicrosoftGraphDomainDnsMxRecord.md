# MicrosoftGraphDomainDnsMxRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IsOptional** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**SupportedService** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**MailExchange** | Pointer to **string** |  | [optional] 
**Preference** | Pointer to **int32** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDomainDnsMxRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDomainDnsMxRecord) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDomainDnsMxRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDomainDnsMxRecord) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetIsOptional

`func (o *MicrosoftGraphDomainDnsMxRecord) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *MicrosoftGraphDomainDnsMxRecord) GetIsOptionalOk() (bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsOptional

`func (o *MicrosoftGraphDomainDnsMxRecord) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.

### SetIsOptional

`func (o *MicrosoftGraphDomainDnsMxRecord) SetIsOptional(v bool)`

SetIsOptional gets a reference to the given bool and assigns it to the IsOptional field.

### GetLabel

`func (o *MicrosoftGraphDomainDnsMxRecord) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MicrosoftGraphDomainDnsMxRecord) GetLabelOk() (string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLabel

`func (o *MicrosoftGraphDomainDnsMxRecord) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabel

`func (o *MicrosoftGraphDomainDnsMxRecord) SetLabel(v string)`

SetLabel gets a reference to the given string and assigns it to the Label field.

### GetRecordType

`func (o *MicrosoftGraphDomainDnsMxRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *MicrosoftGraphDomainDnsMxRecord) GetRecordTypeOk() (string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecordType

`func (o *MicrosoftGraphDomainDnsMxRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### SetRecordType

`func (o *MicrosoftGraphDomainDnsMxRecord) SetRecordType(v string)`

SetRecordType gets a reference to the given string and assigns it to the RecordType field.

### SetRecordTypeExplicitNull

`func (o *MicrosoftGraphDomainDnsMxRecord) SetRecordTypeExplicitNull(b bool)`

SetRecordTypeExplicitNull (un)sets RecordType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RecordType value is set to nil even if false is passed
### GetSupportedService

`func (o *MicrosoftGraphDomainDnsMxRecord) GetSupportedService() string`

GetSupportedService returns the SupportedService field if non-nil, zero value otherwise.

### GetSupportedServiceOk

`func (o *MicrosoftGraphDomainDnsMxRecord) GetSupportedServiceOk() (string, bool)`

GetSupportedServiceOk returns a tuple with the SupportedService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSupportedService

`func (o *MicrosoftGraphDomainDnsMxRecord) HasSupportedService() bool`

HasSupportedService returns a boolean if a field has been set.

### SetSupportedService

`func (o *MicrosoftGraphDomainDnsMxRecord) SetSupportedService(v string)`

SetSupportedService gets a reference to the given string and assigns it to the SupportedService field.

### GetTtl

`func (o *MicrosoftGraphDomainDnsMxRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *MicrosoftGraphDomainDnsMxRecord) GetTtlOk() (int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTtl

`func (o *MicrosoftGraphDomainDnsMxRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtl

`func (o *MicrosoftGraphDomainDnsMxRecord) SetTtl(v int32)`

SetTtl gets a reference to the given int32 and assigns it to the Ttl field.

### GetMailExchange

`func (o *MicrosoftGraphDomainDnsMxRecord) GetMailExchange() string`

GetMailExchange returns the MailExchange field if non-nil, zero value otherwise.

### GetMailExchangeOk

`func (o *MicrosoftGraphDomainDnsMxRecord) GetMailExchangeOk() (string, bool)`

GetMailExchangeOk returns a tuple with the MailExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailExchange

`func (o *MicrosoftGraphDomainDnsMxRecord) HasMailExchange() bool`

HasMailExchange returns a boolean if a field has been set.

### SetMailExchange

`func (o *MicrosoftGraphDomainDnsMxRecord) SetMailExchange(v string)`

SetMailExchange gets a reference to the given string and assigns it to the MailExchange field.

### GetPreference

`func (o *MicrosoftGraphDomainDnsMxRecord) GetPreference() int32`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *MicrosoftGraphDomainDnsMxRecord) GetPreferenceOk() (int32, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreference

`func (o *MicrosoftGraphDomainDnsMxRecord) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### SetPreference

`func (o *MicrosoftGraphDomainDnsMxRecord) SetPreference(v int32)`

SetPreference gets a reference to the given int32 and assigns it to the Preference field.

### SetPreferenceExplicitNull

`func (o *MicrosoftGraphDomainDnsMxRecord) SetPreferenceExplicitNull(b bool)`

SetPreferenceExplicitNull (un)sets Preference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Preference value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


