# MicrosoftGraphWindows10NetworkProxyServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address to the proxy server. Specify an address in the format &lt;server&gt;[“:”&lt;port&gt;] | [optional] 
**Exceptions** | Pointer to **[]string** | Addresses that should not use the proxy server. The system will not use the proxy server for addresses beginning with what is specified in this node. | [optional] 
**UseForLocalAddresses** | Pointer to **bool** | Specifies whether the proxy server should be used for local (intranet) addresses. | [optional] 

## Methods

### GetAddress

`func (o *MicrosoftGraphWindows10NetworkProxyServer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphWindows10NetworkProxyServer) GetAddressOk() (string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddress

`func (o *MicrosoftGraphWindows10NetworkProxyServer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddress

`func (o *MicrosoftGraphWindows10NetworkProxyServer) SetAddress(v string)`

SetAddress gets a reference to the given string and assigns it to the Address field.

### GetExceptions

`func (o *MicrosoftGraphWindows10NetworkProxyServer) GetExceptions() []string`

GetExceptions returns the Exceptions field if non-nil, zero value otherwise.

### GetExceptionsOk

`func (o *MicrosoftGraphWindows10NetworkProxyServer) GetExceptionsOk() ([]string, bool)`

GetExceptionsOk returns a tuple with the Exceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExceptions

`func (o *MicrosoftGraphWindows10NetworkProxyServer) HasExceptions() bool`

HasExceptions returns a boolean if a field has been set.

### SetExceptions

`func (o *MicrosoftGraphWindows10NetworkProxyServer) SetExceptions(v []string)`

SetExceptions gets a reference to the given []string and assigns it to the Exceptions field.

### GetUseForLocalAddresses

`func (o *MicrosoftGraphWindows10NetworkProxyServer) GetUseForLocalAddresses() bool`

GetUseForLocalAddresses returns the UseForLocalAddresses field if non-nil, zero value otherwise.

### GetUseForLocalAddressesOk

`func (o *MicrosoftGraphWindows10NetworkProxyServer) GetUseForLocalAddressesOk() (bool, bool)`

GetUseForLocalAddressesOk returns a tuple with the UseForLocalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUseForLocalAddresses

`func (o *MicrosoftGraphWindows10NetworkProxyServer) HasUseForLocalAddresses() bool`

HasUseForLocalAddresses returns a boolean if a field has been set.

### SetUseForLocalAddresses

`func (o *MicrosoftGraphWindows10NetworkProxyServer) SetUseForLocalAddresses(v bool)`

SetUseForLocalAddresses gets a reference to the given bool and assigns it to the UseForLocalAddresses field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


