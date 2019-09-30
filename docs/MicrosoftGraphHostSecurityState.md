# MicrosoftGraphHostSecurityState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** |  | [optional] 
**IsAzureAdJoined** | Pointer to **bool** |  | [optional] 
**IsAzureAdRegistered** | Pointer to **bool** |  | [optional] 
**IsHybridAzureDomainJoined** | Pointer to **bool** |  | [optional] 
**NetBiosName** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**PrivateIpAddress** | Pointer to **string** |  | [optional] 
**PublicIpAddress** | Pointer to **string** |  | [optional] 
**RiskScore** | Pointer to **string** |  | [optional] 

## Methods

### GetFqdn

`func (o *MicrosoftGraphHostSecurityState) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *MicrosoftGraphHostSecurityState) GetFqdnOk() (string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFqdn

`func (o *MicrosoftGraphHostSecurityState) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdn

`func (o *MicrosoftGraphHostSecurityState) SetFqdn(v string)`

SetFqdn gets a reference to the given string and assigns it to the Fqdn field.

### SetFqdnExplicitNull

`func (o *MicrosoftGraphHostSecurityState) SetFqdnExplicitNull(b bool)`

SetFqdnExplicitNull (un)sets Fqdn to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Fqdn value is set to nil even if false is passed
### GetIsAzureAdJoined

`func (o *MicrosoftGraphHostSecurityState) GetIsAzureAdJoined() bool`

GetIsAzureAdJoined returns the IsAzureAdJoined field if non-nil, zero value otherwise.

### GetIsAzureAdJoinedOk

`func (o *MicrosoftGraphHostSecurityState) GetIsAzureAdJoinedOk() (bool, bool)`

GetIsAzureAdJoinedOk returns a tuple with the IsAzureAdJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAzureAdJoined

`func (o *MicrosoftGraphHostSecurityState) HasIsAzureAdJoined() bool`

HasIsAzureAdJoined returns a boolean if a field has been set.

### SetIsAzureAdJoined

`func (o *MicrosoftGraphHostSecurityState) SetIsAzureAdJoined(v bool)`

SetIsAzureAdJoined gets a reference to the given bool and assigns it to the IsAzureAdJoined field.

### SetIsAzureAdJoinedExplicitNull

`func (o *MicrosoftGraphHostSecurityState) SetIsAzureAdJoinedExplicitNull(b bool)`

SetIsAzureAdJoinedExplicitNull (un)sets IsAzureAdJoined to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsAzureAdJoined value is set to nil even if false is passed
### GetIsAzureAdRegistered

`func (o *MicrosoftGraphHostSecurityState) GetIsAzureAdRegistered() bool`

GetIsAzureAdRegistered returns the IsAzureAdRegistered field if non-nil, zero value otherwise.

### GetIsAzureAdRegisteredOk

`func (o *MicrosoftGraphHostSecurityState) GetIsAzureAdRegisteredOk() (bool, bool)`

GetIsAzureAdRegisteredOk returns a tuple with the IsAzureAdRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAzureAdRegistered

`func (o *MicrosoftGraphHostSecurityState) HasIsAzureAdRegistered() bool`

HasIsAzureAdRegistered returns a boolean if a field has been set.

### SetIsAzureAdRegistered

`func (o *MicrosoftGraphHostSecurityState) SetIsAzureAdRegistered(v bool)`

SetIsAzureAdRegistered gets a reference to the given bool and assigns it to the IsAzureAdRegistered field.

### SetIsAzureAdRegisteredExplicitNull

`func (o *MicrosoftGraphHostSecurityState) SetIsAzureAdRegisteredExplicitNull(b bool)`

SetIsAzureAdRegisteredExplicitNull (un)sets IsAzureAdRegistered to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsAzureAdRegistered value is set to nil even if false is passed
### GetIsHybridAzureDomainJoined

`func (o *MicrosoftGraphHostSecurityState) GetIsHybridAzureDomainJoined() bool`

GetIsHybridAzureDomainJoined returns the IsHybridAzureDomainJoined field if non-nil, zero value otherwise.

### GetIsHybridAzureDomainJoinedOk

`func (o *MicrosoftGraphHostSecurityState) GetIsHybridAzureDomainJoinedOk() (bool, bool)`

GetIsHybridAzureDomainJoinedOk returns a tuple with the IsHybridAzureDomainJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsHybridAzureDomainJoined

`func (o *MicrosoftGraphHostSecurityState) HasIsHybridAzureDomainJoined() bool`

HasIsHybridAzureDomainJoined returns a boolean if a field has been set.

### SetIsHybridAzureDomainJoined

`func (o *MicrosoftGraphHostSecurityState) SetIsHybridAzureDomainJoined(v bool)`

SetIsHybridAzureDomainJoined gets a reference to the given bool and assigns it to the IsHybridAzureDomainJoined field.

### SetIsHybridAzureDomainJoinedExplicitNull

`func (o *MicrosoftGraphHostSecurityState) SetIsHybridAzureDomainJoinedExplicitNull(b bool)`

SetIsHybridAzureDomainJoinedExplicitNull (un)sets IsHybridAzureDomainJoined to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsHybridAzureDomainJoined value is set to nil even if false is passed
### GetNetBiosName

`func (o *MicrosoftGraphHostSecurityState) GetNetBiosName() string`

GetNetBiosName returns the NetBiosName field if non-nil, zero value otherwise.

### GetNetBiosNameOk

`func (o *MicrosoftGraphHostSecurityState) GetNetBiosNameOk() (string, bool)`

GetNetBiosNameOk returns a tuple with the NetBiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetBiosName

`func (o *MicrosoftGraphHostSecurityState) HasNetBiosName() bool`

HasNetBiosName returns a boolean if a field has been set.

### SetNetBiosName

`func (o *MicrosoftGraphHostSecurityState) SetNetBiosName(v string)`

SetNetBiosName gets a reference to the given string and assigns it to the NetBiosName field.

### SetNetBiosNameExplicitNull

`func (o *MicrosoftGraphHostSecurityState) SetNetBiosNameExplicitNull(b bool)`

SetNetBiosNameExplicitNull (un)sets NetBiosName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NetBiosName value is set to nil even if false is passed
### GetOs

`func (o *MicrosoftGraphHostSecurityState) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *MicrosoftGraphHostSecurityState) GetOsOk() (string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOs

`func (o *MicrosoftGraphHostSecurityState) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOs

`func (o *MicrosoftGraphHostSecurityState) SetOs(v string)`

SetOs gets a reference to the given string and assigns it to the Os field.

### SetOsExplicitNull

`func (o *MicrosoftGraphHostSecurityState) SetOsExplicitNull(b bool)`

SetOsExplicitNull (un)sets Os to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Os value is set to nil even if false is passed
### GetPrivateIpAddress

`func (o *MicrosoftGraphHostSecurityState) GetPrivateIpAddress() string`

GetPrivateIpAddress returns the PrivateIpAddress field if non-nil, zero value otherwise.

### GetPrivateIpAddressOk

`func (o *MicrosoftGraphHostSecurityState) GetPrivateIpAddressOk() (string, bool)`

GetPrivateIpAddressOk returns a tuple with the PrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateIpAddress

`func (o *MicrosoftGraphHostSecurityState) HasPrivateIpAddress() bool`

HasPrivateIpAddress returns a boolean if a field has been set.

### SetPrivateIpAddress

`func (o *MicrosoftGraphHostSecurityState) SetPrivateIpAddress(v string)`

SetPrivateIpAddress gets a reference to the given string and assigns it to the PrivateIpAddress field.

### SetPrivateIpAddressExplicitNull

`func (o *MicrosoftGraphHostSecurityState) SetPrivateIpAddressExplicitNull(b bool)`

SetPrivateIpAddressExplicitNull (un)sets PrivateIpAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivateIpAddress value is set to nil even if false is passed
### GetPublicIpAddress

`func (o *MicrosoftGraphHostSecurityState) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *MicrosoftGraphHostSecurityState) GetPublicIpAddressOk() (string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIpAddress

`func (o *MicrosoftGraphHostSecurityState) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### SetPublicIpAddress

`func (o *MicrosoftGraphHostSecurityState) SetPublicIpAddress(v string)`

SetPublicIpAddress gets a reference to the given string and assigns it to the PublicIpAddress field.

### SetPublicIpAddressExplicitNull

`func (o *MicrosoftGraphHostSecurityState) SetPublicIpAddressExplicitNull(b bool)`

SetPublicIpAddressExplicitNull (un)sets PublicIpAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PublicIpAddress value is set to nil even if false is passed
### GetRiskScore

`func (o *MicrosoftGraphHostSecurityState) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *MicrosoftGraphHostSecurityState) GetRiskScoreOk() (string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskScore

`func (o *MicrosoftGraphHostSecurityState) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### SetRiskScore

`func (o *MicrosoftGraphHostSecurityState) SetRiskScore(v string)`

SetRiskScore gets a reference to the given string and assigns it to the RiskScore field.

### SetRiskScoreExplicitNull

`func (o *MicrosoftGraphHostSecurityState) SetRiskScoreExplicitNull(b bool)`

SetRiskScoreExplicitNull (un)sets RiskScore to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskScore value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


