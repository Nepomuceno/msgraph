# MicrosoftGraphDomainDnsSrvRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IsOptional** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**SupportedService** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**NameTarget** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDomainDnsSrvRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetIsOptional

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetIsOptionalOk() (bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsOptional

`func (o *MicrosoftGraphDomainDnsSrvRecord) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.

### SetIsOptional

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetIsOptional(v bool)`

SetIsOptional gets a reference to the given bool and assigns it to the IsOptional field.

### GetLabel

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetLabelOk() (string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLabel

`func (o *MicrosoftGraphDomainDnsSrvRecord) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabel

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetLabel(v string)`

SetLabel gets a reference to the given string and assigns it to the Label field.

### GetRecordType

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetRecordTypeOk() (string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecordType

`func (o *MicrosoftGraphDomainDnsSrvRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### SetRecordType

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetRecordType(v string)`

SetRecordType gets a reference to the given string and assigns it to the RecordType field.

### SetRecordTypeExplicitNull

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetRecordTypeExplicitNull(b bool)`

SetRecordTypeExplicitNull (un)sets RecordType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RecordType value is set to nil even if false is passed
### GetSupportedService

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetSupportedService() string`

GetSupportedService returns the SupportedService field if non-nil, zero value otherwise.

### GetSupportedServiceOk

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetSupportedServiceOk() (string, bool)`

GetSupportedServiceOk returns a tuple with the SupportedService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSupportedService

`func (o *MicrosoftGraphDomainDnsSrvRecord) HasSupportedService() bool`

HasSupportedService returns a boolean if a field has been set.

### SetSupportedService

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetSupportedService(v string)`

SetSupportedService gets a reference to the given string and assigns it to the SupportedService field.

### GetTtl

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetTtlOk() (int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTtl

`func (o *MicrosoftGraphDomainDnsSrvRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtl

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetTtl(v int32)`

SetTtl gets a reference to the given int32 and assigns it to the Ttl field.

### GetNameTarget

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetNameTarget() string`

GetNameTarget returns the NameTarget field if non-nil, zero value otherwise.

### GetNameTargetOk

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetNameTargetOk() (string, bool)`

GetNameTargetOk returns a tuple with the NameTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNameTarget

`func (o *MicrosoftGraphDomainDnsSrvRecord) HasNameTarget() bool`

HasNameTarget returns a boolean if a field has been set.

### SetNameTarget

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetNameTarget(v string)`

SetNameTarget gets a reference to the given string and assigns it to the NameTarget field.

### SetNameTargetExplicitNull

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetNameTargetExplicitNull(b bool)`

SetNameTargetExplicitNull (un)sets NameTarget to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NameTarget value is set to nil even if false is passed
### GetPort

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetPortOk() (int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPort

`func (o *MicrosoftGraphDomainDnsSrvRecord) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPort

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetPort(v int32)`

SetPort gets a reference to the given int32 and assigns it to the Port field.

### SetPortExplicitNull

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetPortExplicitNull(b bool)`

SetPortExplicitNull (un)sets Port to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Port value is set to nil even if false is passed
### GetPriority

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetPriorityOk() (int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPriority

`func (o *MicrosoftGraphDomainDnsSrvRecord) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriority

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetPriority(v int32)`

SetPriority gets a reference to the given int32 and assigns it to the Priority field.

### SetPriorityExplicitNull

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetPriorityExplicitNull(b bool)`

SetPriorityExplicitNull (un)sets Priority to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Priority value is set to nil even if false is passed
### GetProtocol

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetProtocolOk() (string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtocol

`func (o *MicrosoftGraphDomainDnsSrvRecord) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocol

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetProtocol(v string)`

SetProtocol gets a reference to the given string and assigns it to the Protocol field.

### SetProtocolExplicitNull

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetProtocolExplicitNull(b bool)`

SetProtocolExplicitNull (un)sets Protocol to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Protocol value is set to nil even if false is passed
### GetService

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetServiceOk() (string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasService

`func (o *MicrosoftGraphDomainDnsSrvRecord) HasService() bool`

HasService returns a boolean if a field has been set.

### SetService

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetService(v string)`

SetService gets a reference to the given string and assigns it to the Service field.

### SetServiceExplicitNull

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetServiceExplicitNull(b bool)`

SetServiceExplicitNull (un)sets Service to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Service value is set to nil even if false is passed
### GetWeight

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *MicrosoftGraphDomainDnsSrvRecord) GetWeightOk() (int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWeight

`func (o *MicrosoftGraphDomainDnsSrvRecord) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeight

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetWeight(v int32)`

SetWeight gets a reference to the given int32 and assigns it to the Weight field.

### SetWeightExplicitNull

`func (o *MicrosoftGraphDomainDnsSrvRecord) SetWeightExplicitNull(b bool)`

SetWeightExplicitNull (un)sets Weight to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Weight value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


