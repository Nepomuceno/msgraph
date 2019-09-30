# MicrosoftGraphProxiedDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddressOrFQDN** | Pointer to **string** | The IP address or FQDN | [optional] 
**Proxy** | Pointer to **string** | Proxy IP or FQDN | [optional] 

## Methods

### GetIpAddressOrFQDN

`func (o *MicrosoftGraphProxiedDomain) GetIpAddressOrFQDN() string`

GetIpAddressOrFQDN returns the IpAddressOrFQDN field if non-nil, zero value otherwise.

### GetIpAddressOrFQDNOk

`func (o *MicrosoftGraphProxiedDomain) GetIpAddressOrFQDNOk() (string, bool)`

GetIpAddressOrFQDNOk returns a tuple with the IpAddressOrFQDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpAddressOrFQDN

`func (o *MicrosoftGraphProxiedDomain) HasIpAddressOrFQDN() bool`

HasIpAddressOrFQDN returns a boolean if a field has been set.

### SetIpAddressOrFQDN

`func (o *MicrosoftGraphProxiedDomain) SetIpAddressOrFQDN(v string)`

SetIpAddressOrFQDN gets a reference to the given string and assigns it to the IpAddressOrFQDN field.

### GetProxy

`func (o *MicrosoftGraphProxiedDomain) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *MicrosoftGraphProxiedDomain) GetProxyOk() (string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProxy

`func (o *MicrosoftGraphProxiedDomain) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxy

`func (o *MicrosoftGraphProxiedDomain) SetProxy(v string)`

SetProxy gets a reference to the given string and assigns it to the Proxy field.

### SetProxyExplicitNull

`func (o *MicrosoftGraphProxiedDomain) SetProxyExplicitNull(b bool)`

SetProxyExplicitNull (un)sets Proxy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Proxy value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


