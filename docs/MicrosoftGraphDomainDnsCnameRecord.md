# MicrosoftGraphDomainDnsCnameRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IsOptional** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**SupportedService** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**CanonicalName** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDomainDnsCnameRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDomainDnsCnameRecord) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetIsOptional

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetIsOptionalOk() (bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsOptional

`func (o *MicrosoftGraphDomainDnsCnameRecord) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.

### SetIsOptional

`func (o *MicrosoftGraphDomainDnsCnameRecord) SetIsOptional(v bool)`

SetIsOptional gets a reference to the given bool and assigns it to the IsOptional field.

### GetLabel

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetLabelOk() (string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLabel

`func (o *MicrosoftGraphDomainDnsCnameRecord) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabel

`func (o *MicrosoftGraphDomainDnsCnameRecord) SetLabel(v string)`

SetLabel gets a reference to the given string and assigns it to the Label field.

### GetRecordType

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetRecordTypeOk() (string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecordType

`func (o *MicrosoftGraphDomainDnsCnameRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### SetRecordType

`func (o *MicrosoftGraphDomainDnsCnameRecord) SetRecordType(v string)`

SetRecordType gets a reference to the given string and assigns it to the RecordType field.

### SetRecordTypeExplicitNull

`func (o *MicrosoftGraphDomainDnsCnameRecord) SetRecordTypeExplicitNull(b bool)`

SetRecordTypeExplicitNull (un)sets RecordType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RecordType value is set to nil even if false is passed
### GetSupportedService

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetSupportedService() string`

GetSupportedService returns the SupportedService field if non-nil, zero value otherwise.

### GetSupportedServiceOk

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetSupportedServiceOk() (string, bool)`

GetSupportedServiceOk returns a tuple with the SupportedService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSupportedService

`func (o *MicrosoftGraphDomainDnsCnameRecord) HasSupportedService() bool`

HasSupportedService returns a boolean if a field has been set.

### SetSupportedService

`func (o *MicrosoftGraphDomainDnsCnameRecord) SetSupportedService(v string)`

SetSupportedService gets a reference to the given string and assigns it to the SupportedService field.

### GetTtl

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetTtlOk() (int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTtl

`func (o *MicrosoftGraphDomainDnsCnameRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtl

`func (o *MicrosoftGraphDomainDnsCnameRecord) SetTtl(v int32)`

SetTtl gets a reference to the given int32 and assigns it to the Ttl field.

### GetCanonicalName

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetCanonicalName() string`

GetCanonicalName returns the CanonicalName field if non-nil, zero value otherwise.

### GetCanonicalNameOk

`func (o *MicrosoftGraphDomainDnsCnameRecord) GetCanonicalNameOk() (string, bool)`

GetCanonicalNameOk returns a tuple with the CanonicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanonicalName

`func (o *MicrosoftGraphDomainDnsCnameRecord) HasCanonicalName() bool`

HasCanonicalName returns a boolean if a field has been set.

### SetCanonicalName

`func (o *MicrosoftGraphDomainDnsCnameRecord) SetCanonicalName(v string)`

SetCanonicalName gets a reference to the given string and assigns it to the CanonicalName field.

### SetCanonicalNameExplicitNull

`func (o *MicrosoftGraphDomainDnsCnameRecord) SetCanonicalNameExplicitNull(b bool)`

SetCanonicalNameExplicitNull (un)sets CanonicalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CanonicalName value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


