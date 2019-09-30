# MicrosoftGraphNetworkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **string** |  | [optional] 
**DestinationAddress** | Pointer to **string** |  | [optional] 
**DestinationDomain** | Pointer to **string** |  | [optional] 
**DestinationPort** | Pointer to **string** |  | [optional] 
**DestinationUrl** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to [**AnyOfmicrosoftGraphConnectionDirection**](anyOf&lt;microsoft.graph.connectionDirection&gt;.md) |  | [optional] 
**DomainRegisteredDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LocalDnsName** | Pointer to **string** |  | [optional] 
**NatDestinationAddress** | Pointer to **string** |  | [optional] 
**NatDestinationPort** | Pointer to **string** |  | [optional] 
**NatSourceAddress** | Pointer to **string** |  | [optional] 
**NatSourcePort** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to [**AnyOfmicrosoftGraphSecurityNetworkProtocol**](anyOf&lt;microsoft.graph.securityNetworkProtocol&gt;.md) |  | [optional] 
**RiskScore** | Pointer to **string** |  | [optional] 
**SourceAddress** | Pointer to **string** |  | [optional] 
**SourcePort** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphConnectionStatus**](anyOf&lt;microsoft.graph.connectionStatus&gt;.md) |  | [optional] 
**UrlParameters** | Pointer to **string** |  | [optional] 

## Methods

### GetApplicationName

`func (o *MicrosoftGraphNetworkConnection) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *MicrosoftGraphNetworkConnection) GetApplicationNameOk() (string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationName

`func (o *MicrosoftGraphNetworkConnection) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### SetApplicationName

`func (o *MicrosoftGraphNetworkConnection) SetApplicationName(v string)`

SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.

### SetApplicationNameExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetApplicationNameExplicitNull(b bool)`

SetApplicationNameExplicitNull (un)sets ApplicationName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ApplicationName value is set to nil even if false is passed
### GetDestinationAddress

`func (o *MicrosoftGraphNetworkConnection) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *MicrosoftGraphNetworkConnection) GetDestinationAddressOk() (string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDestinationAddress

`func (o *MicrosoftGraphNetworkConnection) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### SetDestinationAddress

`func (o *MicrosoftGraphNetworkConnection) SetDestinationAddress(v string)`

SetDestinationAddress gets a reference to the given string and assigns it to the DestinationAddress field.

### SetDestinationAddressExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetDestinationAddressExplicitNull(b bool)`

SetDestinationAddressExplicitNull (un)sets DestinationAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DestinationAddress value is set to nil even if false is passed
### GetDestinationDomain

`func (o *MicrosoftGraphNetworkConnection) GetDestinationDomain() string`

GetDestinationDomain returns the DestinationDomain field if non-nil, zero value otherwise.

### GetDestinationDomainOk

`func (o *MicrosoftGraphNetworkConnection) GetDestinationDomainOk() (string, bool)`

GetDestinationDomainOk returns a tuple with the DestinationDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDestinationDomain

`func (o *MicrosoftGraphNetworkConnection) HasDestinationDomain() bool`

HasDestinationDomain returns a boolean if a field has been set.

### SetDestinationDomain

`func (o *MicrosoftGraphNetworkConnection) SetDestinationDomain(v string)`

SetDestinationDomain gets a reference to the given string and assigns it to the DestinationDomain field.

### SetDestinationDomainExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetDestinationDomainExplicitNull(b bool)`

SetDestinationDomainExplicitNull (un)sets DestinationDomain to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DestinationDomain value is set to nil even if false is passed
### GetDestinationPort

`func (o *MicrosoftGraphNetworkConnection) GetDestinationPort() string`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *MicrosoftGraphNetworkConnection) GetDestinationPortOk() (string, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDestinationPort

`func (o *MicrosoftGraphNetworkConnection) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### SetDestinationPort

`func (o *MicrosoftGraphNetworkConnection) SetDestinationPort(v string)`

SetDestinationPort gets a reference to the given string and assigns it to the DestinationPort field.

### SetDestinationPortExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetDestinationPortExplicitNull(b bool)`

SetDestinationPortExplicitNull (un)sets DestinationPort to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DestinationPort value is set to nil even if false is passed
### GetDestinationUrl

`func (o *MicrosoftGraphNetworkConnection) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *MicrosoftGraphNetworkConnection) GetDestinationUrlOk() (string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDestinationUrl

`func (o *MicrosoftGraphNetworkConnection) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.

### SetDestinationUrl

`func (o *MicrosoftGraphNetworkConnection) SetDestinationUrl(v string)`

SetDestinationUrl gets a reference to the given string and assigns it to the DestinationUrl field.

### SetDestinationUrlExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetDestinationUrlExplicitNull(b bool)`

SetDestinationUrlExplicitNull (un)sets DestinationUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DestinationUrl value is set to nil even if false is passed
### GetDirection

`func (o *MicrosoftGraphNetworkConnection) GetDirection() AnyOfmicrosoftGraphConnectionDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MicrosoftGraphNetworkConnection) GetDirectionOk() (AnyOfmicrosoftGraphConnectionDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirection

`func (o *MicrosoftGraphNetworkConnection) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirection

`func (o *MicrosoftGraphNetworkConnection) SetDirection(v AnyOfmicrosoftGraphConnectionDirection)`

SetDirection gets a reference to the given AnyOfmicrosoftGraphConnectionDirection and assigns it to the Direction field.

### SetDirectionExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetDirectionExplicitNull(b bool)`

SetDirectionExplicitNull (un)sets Direction to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Direction value is set to nil even if false is passed
### GetDomainRegisteredDateTime

`func (o *MicrosoftGraphNetworkConnection) GetDomainRegisteredDateTime() time.Time`

GetDomainRegisteredDateTime returns the DomainRegisteredDateTime field if non-nil, zero value otherwise.

### GetDomainRegisteredDateTimeOk

`func (o *MicrosoftGraphNetworkConnection) GetDomainRegisteredDateTimeOk() (time.Time, bool)`

GetDomainRegisteredDateTimeOk returns a tuple with the DomainRegisteredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomainRegisteredDateTime

`func (o *MicrosoftGraphNetworkConnection) HasDomainRegisteredDateTime() bool`

HasDomainRegisteredDateTime returns a boolean if a field has been set.

### SetDomainRegisteredDateTime

`func (o *MicrosoftGraphNetworkConnection) SetDomainRegisteredDateTime(v time.Time)`

SetDomainRegisteredDateTime gets a reference to the given time.Time and assigns it to the DomainRegisteredDateTime field.

### SetDomainRegisteredDateTimeExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetDomainRegisteredDateTimeExplicitNull(b bool)`

SetDomainRegisteredDateTimeExplicitNull (un)sets DomainRegisteredDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DomainRegisteredDateTime value is set to nil even if false is passed
### GetLocalDnsName

`func (o *MicrosoftGraphNetworkConnection) GetLocalDnsName() string`

GetLocalDnsName returns the LocalDnsName field if non-nil, zero value otherwise.

### GetLocalDnsNameOk

`func (o *MicrosoftGraphNetworkConnection) GetLocalDnsNameOk() (string, bool)`

GetLocalDnsNameOk returns a tuple with the LocalDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocalDnsName

`func (o *MicrosoftGraphNetworkConnection) HasLocalDnsName() bool`

HasLocalDnsName returns a boolean if a field has been set.

### SetLocalDnsName

`func (o *MicrosoftGraphNetworkConnection) SetLocalDnsName(v string)`

SetLocalDnsName gets a reference to the given string and assigns it to the LocalDnsName field.

### SetLocalDnsNameExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetLocalDnsNameExplicitNull(b bool)`

SetLocalDnsNameExplicitNull (un)sets LocalDnsName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LocalDnsName value is set to nil even if false is passed
### GetNatDestinationAddress

`func (o *MicrosoftGraphNetworkConnection) GetNatDestinationAddress() string`

GetNatDestinationAddress returns the NatDestinationAddress field if non-nil, zero value otherwise.

### GetNatDestinationAddressOk

`func (o *MicrosoftGraphNetworkConnection) GetNatDestinationAddressOk() (string, bool)`

GetNatDestinationAddressOk returns a tuple with the NatDestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNatDestinationAddress

`func (o *MicrosoftGraphNetworkConnection) HasNatDestinationAddress() bool`

HasNatDestinationAddress returns a boolean if a field has been set.

### SetNatDestinationAddress

`func (o *MicrosoftGraphNetworkConnection) SetNatDestinationAddress(v string)`

SetNatDestinationAddress gets a reference to the given string and assigns it to the NatDestinationAddress field.

### SetNatDestinationAddressExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetNatDestinationAddressExplicitNull(b bool)`

SetNatDestinationAddressExplicitNull (un)sets NatDestinationAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NatDestinationAddress value is set to nil even if false is passed
### GetNatDestinationPort

`func (o *MicrosoftGraphNetworkConnection) GetNatDestinationPort() string`

GetNatDestinationPort returns the NatDestinationPort field if non-nil, zero value otherwise.

### GetNatDestinationPortOk

`func (o *MicrosoftGraphNetworkConnection) GetNatDestinationPortOk() (string, bool)`

GetNatDestinationPortOk returns a tuple with the NatDestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNatDestinationPort

`func (o *MicrosoftGraphNetworkConnection) HasNatDestinationPort() bool`

HasNatDestinationPort returns a boolean if a field has been set.

### SetNatDestinationPort

`func (o *MicrosoftGraphNetworkConnection) SetNatDestinationPort(v string)`

SetNatDestinationPort gets a reference to the given string and assigns it to the NatDestinationPort field.

### SetNatDestinationPortExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetNatDestinationPortExplicitNull(b bool)`

SetNatDestinationPortExplicitNull (un)sets NatDestinationPort to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NatDestinationPort value is set to nil even if false is passed
### GetNatSourceAddress

`func (o *MicrosoftGraphNetworkConnection) GetNatSourceAddress() string`

GetNatSourceAddress returns the NatSourceAddress field if non-nil, zero value otherwise.

### GetNatSourceAddressOk

`func (o *MicrosoftGraphNetworkConnection) GetNatSourceAddressOk() (string, bool)`

GetNatSourceAddressOk returns a tuple with the NatSourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNatSourceAddress

`func (o *MicrosoftGraphNetworkConnection) HasNatSourceAddress() bool`

HasNatSourceAddress returns a boolean if a field has been set.

### SetNatSourceAddress

`func (o *MicrosoftGraphNetworkConnection) SetNatSourceAddress(v string)`

SetNatSourceAddress gets a reference to the given string and assigns it to the NatSourceAddress field.

### SetNatSourceAddressExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetNatSourceAddressExplicitNull(b bool)`

SetNatSourceAddressExplicitNull (un)sets NatSourceAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NatSourceAddress value is set to nil even if false is passed
### GetNatSourcePort

`func (o *MicrosoftGraphNetworkConnection) GetNatSourcePort() string`

GetNatSourcePort returns the NatSourcePort field if non-nil, zero value otherwise.

### GetNatSourcePortOk

`func (o *MicrosoftGraphNetworkConnection) GetNatSourcePortOk() (string, bool)`

GetNatSourcePortOk returns a tuple with the NatSourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNatSourcePort

`func (o *MicrosoftGraphNetworkConnection) HasNatSourcePort() bool`

HasNatSourcePort returns a boolean if a field has been set.

### SetNatSourcePort

`func (o *MicrosoftGraphNetworkConnection) SetNatSourcePort(v string)`

SetNatSourcePort gets a reference to the given string and assigns it to the NatSourcePort field.

### SetNatSourcePortExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetNatSourcePortExplicitNull(b bool)`

SetNatSourcePortExplicitNull (un)sets NatSourcePort to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NatSourcePort value is set to nil even if false is passed
### GetProtocol

`func (o *MicrosoftGraphNetworkConnection) GetProtocol() AnyOfmicrosoftGraphSecurityNetworkProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *MicrosoftGraphNetworkConnection) GetProtocolOk() (AnyOfmicrosoftGraphSecurityNetworkProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtocol

`func (o *MicrosoftGraphNetworkConnection) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocol

`func (o *MicrosoftGraphNetworkConnection) SetProtocol(v AnyOfmicrosoftGraphSecurityNetworkProtocol)`

SetProtocol gets a reference to the given AnyOfmicrosoftGraphSecurityNetworkProtocol and assigns it to the Protocol field.

### SetProtocolExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetProtocolExplicitNull(b bool)`

SetProtocolExplicitNull (un)sets Protocol to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Protocol value is set to nil even if false is passed
### GetRiskScore

`func (o *MicrosoftGraphNetworkConnection) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *MicrosoftGraphNetworkConnection) GetRiskScoreOk() (string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskScore

`func (o *MicrosoftGraphNetworkConnection) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### SetRiskScore

`func (o *MicrosoftGraphNetworkConnection) SetRiskScore(v string)`

SetRiskScore gets a reference to the given string and assigns it to the RiskScore field.

### SetRiskScoreExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetRiskScoreExplicitNull(b bool)`

SetRiskScoreExplicitNull (un)sets RiskScore to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskScore value is set to nil even if false is passed
### GetSourceAddress

`func (o *MicrosoftGraphNetworkConnection) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *MicrosoftGraphNetworkConnection) GetSourceAddressOk() (string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceAddress

`func (o *MicrosoftGraphNetworkConnection) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### SetSourceAddress

`func (o *MicrosoftGraphNetworkConnection) SetSourceAddress(v string)`

SetSourceAddress gets a reference to the given string and assigns it to the SourceAddress field.

### SetSourceAddressExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetSourceAddressExplicitNull(b bool)`

SetSourceAddressExplicitNull (un)sets SourceAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SourceAddress value is set to nil even if false is passed
### GetSourcePort

`func (o *MicrosoftGraphNetworkConnection) GetSourcePort() string`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *MicrosoftGraphNetworkConnection) GetSourcePortOk() (string, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourcePort

`func (o *MicrosoftGraphNetworkConnection) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### SetSourcePort

`func (o *MicrosoftGraphNetworkConnection) SetSourcePort(v string)`

SetSourcePort gets a reference to the given string and assigns it to the SourcePort field.

### SetSourcePortExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetSourcePortExplicitNull(b bool)`

SetSourcePortExplicitNull (un)sets SourcePort to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SourcePort value is set to nil even if false is passed
### GetStatus

`func (o *MicrosoftGraphNetworkConnection) GetStatus() AnyOfmicrosoftGraphConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphNetworkConnection) GetStatusOk() (AnyOfmicrosoftGraphConnectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphNetworkConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphNetworkConnection) SetStatus(v AnyOfmicrosoftGraphConnectionStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphConnectionStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetUrlParameters

`func (o *MicrosoftGraphNetworkConnection) GetUrlParameters() string`

GetUrlParameters returns the UrlParameters field if non-nil, zero value otherwise.

### GetUrlParametersOk

`func (o *MicrosoftGraphNetworkConnection) GetUrlParametersOk() (string, bool)`

GetUrlParametersOk returns a tuple with the UrlParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrlParameters

`func (o *MicrosoftGraphNetworkConnection) HasUrlParameters() bool`

HasUrlParameters returns a boolean if a field has been set.

### SetUrlParameters

`func (o *MicrosoftGraphNetworkConnection) SetUrlParameters(v string)`

SetUrlParameters gets a reference to the given string and assigns it to the UrlParameters field.

### SetUrlParametersExplicitNull

`func (o *MicrosoftGraphNetworkConnection) SetUrlParametersExplicitNull(b bool)`

SetUrlParametersExplicitNull (un)sets UrlParameters to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UrlParameters value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


