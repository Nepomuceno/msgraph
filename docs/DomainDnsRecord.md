# DomainDnsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOptional** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**SupportedService** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 

## Methods

### GetIsOptional

`func (o *DomainDnsRecord) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *DomainDnsRecord) GetIsOptionalOk() (bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsOptional

`func (o *DomainDnsRecord) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.

### SetIsOptional

`func (o *DomainDnsRecord) SetIsOptional(v bool)`

SetIsOptional gets a reference to the given bool and assigns it to the IsOptional field.

### GetLabel

`func (o *DomainDnsRecord) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DomainDnsRecord) GetLabelOk() (string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLabel

`func (o *DomainDnsRecord) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabel

`func (o *DomainDnsRecord) SetLabel(v string)`

SetLabel gets a reference to the given string and assigns it to the Label field.

### GetRecordType

`func (o *DomainDnsRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *DomainDnsRecord) GetRecordTypeOk() (string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecordType

`func (o *DomainDnsRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### SetRecordType

`func (o *DomainDnsRecord) SetRecordType(v string)`

SetRecordType gets a reference to the given string and assigns it to the RecordType field.

### SetRecordTypeExplicitNull

`func (o *DomainDnsRecord) SetRecordTypeExplicitNull(b bool)`

SetRecordTypeExplicitNull (un)sets RecordType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RecordType value is set to nil even if false is passed
### GetSupportedService

`func (o *DomainDnsRecord) GetSupportedService() string`

GetSupportedService returns the SupportedService field if non-nil, zero value otherwise.

### GetSupportedServiceOk

`func (o *DomainDnsRecord) GetSupportedServiceOk() (string, bool)`

GetSupportedServiceOk returns a tuple with the SupportedService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSupportedService

`func (o *DomainDnsRecord) HasSupportedService() bool`

HasSupportedService returns a boolean if a field has been set.

### SetSupportedService

`func (o *DomainDnsRecord) SetSupportedService(v string)`

SetSupportedService gets a reference to the given string and assigns it to the SupportedService field.

### GetTtl

`func (o *DomainDnsRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DomainDnsRecord) GetTtlOk() (int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTtl

`func (o *DomainDnsRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtl

`func (o *DomainDnsRecord) SetTtl(v int32)`

SetTtl gets a reference to the given int32 and assigns it to the Ttl field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


