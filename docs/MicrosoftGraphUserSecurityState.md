# MicrosoftGraphUserSecurityState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AadUserId** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**EmailRole** | Pointer to [**AnyOfmicrosoftGraphEmailRole**](anyOf&lt;microsoft.graph.emailRole&gt;.md) |  | [optional] 
**IsVpn** | Pointer to **bool** |  | [optional] 
**LogonDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LogonId** | Pointer to **string** |  | [optional] 
**LogonIp** | Pointer to **string** |  | [optional] 
**LogonLocation** | Pointer to **string** |  | [optional] 
**LogonType** | Pointer to [**AnyOfmicrosoftGraphLogonType**](anyOf&lt;microsoft.graph.logonType&gt;.md) |  | [optional] 
**OnPremisesSecurityIdentifier** | Pointer to **string** |  | [optional] 
**RiskScore** | Pointer to **string** |  | [optional] 
**UserAccountType** | Pointer to [**AnyOfmicrosoftGraphUserAccountSecurityType**](anyOf&lt;microsoft.graph.userAccountSecurityType&gt;.md) |  | [optional] 
**UserPrincipalName** | Pointer to **string** |  | [optional] 

## Methods

### GetAadUserId

`func (o *MicrosoftGraphUserSecurityState) GetAadUserId() string`

GetAadUserId returns the AadUserId field if non-nil, zero value otherwise.

### GetAadUserIdOk

`func (o *MicrosoftGraphUserSecurityState) GetAadUserIdOk() (string, bool)`

GetAadUserIdOk returns a tuple with the AadUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAadUserId

`func (o *MicrosoftGraphUserSecurityState) HasAadUserId() bool`

HasAadUserId returns a boolean if a field has been set.

### SetAadUserId

`func (o *MicrosoftGraphUserSecurityState) SetAadUserId(v string)`

SetAadUserId gets a reference to the given string and assigns it to the AadUserId field.

### SetAadUserIdExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetAadUserIdExplicitNull(b bool)`

SetAadUserIdExplicitNull (un)sets AadUserId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AadUserId value is set to nil even if false is passed
### GetAccountName

`func (o *MicrosoftGraphUserSecurityState) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *MicrosoftGraphUserSecurityState) GetAccountNameOk() (string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountName

`func (o *MicrosoftGraphUserSecurityState) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountName

`func (o *MicrosoftGraphUserSecurityState) SetAccountName(v string)`

SetAccountName gets a reference to the given string and assigns it to the AccountName field.

### SetAccountNameExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetAccountNameExplicitNull(b bool)`

SetAccountNameExplicitNull (un)sets AccountName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AccountName value is set to nil even if false is passed
### GetDomainName

`func (o *MicrosoftGraphUserSecurityState) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *MicrosoftGraphUserSecurityState) GetDomainNameOk() (string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomainName

`func (o *MicrosoftGraphUserSecurityState) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainName

`func (o *MicrosoftGraphUserSecurityState) SetDomainName(v string)`

SetDomainName gets a reference to the given string and assigns it to the DomainName field.

### SetDomainNameExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetDomainNameExplicitNull(b bool)`

SetDomainNameExplicitNull (un)sets DomainName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DomainName value is set to nil even if false is passed
### GetEmailRole

`func (o *MicrosoftGraphUserSecurityState) GetEmailRole() AnyOfmicrosoftGraphEmailRole`

GetEmailRole returns the EmailRole field if non-nil, zero value otherwise.

### GetEmailRoleOk

`func (o *MicrosoftGraphUserSecurityState) GetEmailRoleOk() (AnyOfmicrosoftGraphEmailRole, bool)`

GetEmailRoleOk returns a tuple with the EmailRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailRole

`func (o *MicrosoftGraphUserSecurityState) HasEmailRole() bool`

HasEmailRole returns a boolean if a field has been set.

### SetEmailRole

`func (o *MicrosoftGraphUserSecurityState) SetEmailRole(v AnyOfmicrosoftGraphEmailRole)`

SetEmailRole gets a reference to the given AnyOfmicrosoftGraphEmailRole and assigns it to the EmailRole field.

### SetEmailRoleExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetEmailRoleExplicitNull(b bool)`

SetEmailRoleExplicitNull (un)sets EmailRole to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EmailRole value is set to nil even if false is passed
### GetIsVpn

`func (o *MicrosoftGraphUserSecurityState) GetIsVpn() bool`

GetIsVpn returns the IsVpn field if non-nil, zero value otherwise.

### GetIsVpnOk

`func (o *MicrosoftGraphUserSecurityState) GetIsVpnOk() (bool, bool)`

GetIsVpnOk returns a tuple with the IsVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsVpn

`func (o *MicrosoftGraphUserSecurityState) HasIsVpn() bool`

HasIsVpn returns a boolean if a field has been set.

### SetIsVpn

`func (o *MicrosoftGraphUserSecurityState) SetIsVpn(v bool)`

SetIsVpn gets a reference to the given bool and assigns it to the IsVpn field.

### SetIsVpnExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetIsVpnExplicitNull(b bool)`

SetIsVpnExplicitNull (un)sets IsVpn to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsVpn value is set to nil even if false is passed
### GetLogonDateTime

`func (o *MicrosoftGraphUserSecurityState) GetLogonDateTime() time.Time`

GetLogonDateTime returns the LogonDateTime field if non-nil, zero value otherwise.

### GetLogonDateTimeOk

`func (o *MicrosoftGraphUserSecurityState) GetLogonDateTimeOk() (time.Time, bool)`

GetLogonDateTimeOk returns a tuple with the LogonDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogonDateTime

`func (o *MicrosoftGraphUserSecurityState) HasLogonDateTime() bool`

HasLogonDateTime returns a boolean if a field has been set.

### SetLogonDateTime

`func (o *MicrosoftGraphUserSecurityState) SetLogonDateTime(v time.Time)`

SetLogonDateTime gets a reference to the given time.Time and assigns it to the LogonDateTime field.

### SetLogonDateTimeExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetLogonDateTimeExplicitNull(b bool)`

SetLogonDateTimeExplicitNull (un)sets LogonDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LogonDateTime value is set to nil even if false is passed
### GetLogonId

`func (o *MicrosoftGraphUserSecurityState) GetLogonId() string`

GetLogonId returns the LogonId field if non-nil, zero value otherwise.

### GetLogonIdOk

`func (o *MicrosoftGraphUserSecurityState) GetLogonIdOk() (string, bool)`

GetLogonIdOk returns a tuple with the LogonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogonId

`func (o *MicrosoftGraphUserSecurityState) HasLogonId() bool`

HasLogonId returns a boolean if a field has been set.

### SetLogonId

`func (o *MicrosoftGraphUserSecurityState) SetLogonId(v string)`

SetLogonId gets a reference to the given string and assigns it to the LogonId field.

### SetLogonIdExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetLogonIdExplicitNull(b bool)`

SetLogonIdExplicitNull (un)sets LogonId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LogonId value is set to nil even if false is passed
### GetLogonIp

`func (o *MicrosoftGraphUserSecurityState) GetLogonIp() string`

GetLogonIp returns the LogonIp field if non-nil, zero value otherwise.

### GetLogonIpOk

`func (o *MicrosoftGraphUserSecurityState) GetLogonIpOk() (string, bool)`

GetLogonIpOk returns a tuple with the LogonIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogonIp

`func (o *MicrosoftGraphUserSecurityState) HasLogonIp() bool`

HasLogonIp returns a boolean if a field has been set.

### SetLogonIp

`func (o *MicrosoftGraphUserSecurityState) SetLogonIp(v string)`

SetLogonIp gets a reference to the given string and assigns it to the LogonIp field.

### SetLogonIpExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetLogonIpExplicitNull(b bool)`

SetLogonIpExplicitNull (un)sets LogonIp to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LogonIp value is set to nil even if false is passed
### GetLogonLocation

`func (o *MicrosoftGraphUserSecurityState) GetLogonLocation() string`

GetLogonLocation returns the LogonLocation field if non-nil, zero value otherwise.

### GetLogonLocationOk

`func (o *MicrosoftGraphUserSecurityState) GetLogonLocationOk() (string, bool)`

GetLogonLocationOk returns a tuple with the LogonLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogonLocation

`func (o *MicrosoftGraphUserSecurityState) HasLogonLocation() bool`

HasLogonLocation returns a boolean if a field has been set.

### SetLogonLocation

`func (o *MicrosoftGraphUserSecurityState) SetLogonLocation(v string)`

SetLogonLocation gets a reference to the given string and assigns it to the LogonLocation field.

### SetLogonLocationExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetLogonLocationExplicitNull(b bool)`

SetLogonLocationExplicitNull (un)sets LogonLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LogonLocation value is set to nil even if false is passed
### GetLogonType

`func (o *MicrosoftGraphUserSecurityState) GetLogonType() AnyOfmicrosoftGraphLogonType`

GetLogonType returns the LogonType field if non-nil, zero value otherwise.

### GetLogonTypeOk

`func (o *MicrosoftGraphUserSecurityState) GetLogonTypeOk() (AnyOfmicrosoftGraphLogonType, bool)`

GetLogonTypeOk returns a tuple with the LogonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogonType

`func (o *MicrosoftGraphUserSecurityState) HasLogonType() bool`

HasLogonType returns a boolean if a field has been set.

### SetLogonType

`func (o *MicrosoftGraphUserSecurityState) SetLogonType(v AnyOfmicrosoftGraphLogonType)`

SetLogonType gets a reference to the given AnyOfmicrosoftGraphLogonType and assigns it to the LogonType field.

### SetLogonTypeExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetLogonTypeExplicitNull(b bool)`

SetLogonTypeExplicitNull (un)sets LogonType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LogonType value is set to nil even if false is passed
### GetOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphUserSecurityState) GetOnPremisesSecurityIdentifier() string`

GetOnPremisesSecurityIdentifier returns the OnPremisesSecurityIdentifier field if non-nil, zero value otherwise.

### GetOnPremisesSecurityIdentifierOk

`func (o *MicrosoftGraphUserSecurityState) GetOnPremisesSecurityIdentifierOk() (string, bool)`

GetOnPremisesSecurityIdentifierOk returns a tuple with the OnPremisesSecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphUserSecurityState) HasOnPremisesSecurityIdentifier() bool`

HasOnPremisesSecurityIdentifier returns a boolean if a field has been set.

### SetOnPremisesSecurityIdentifier

`func (o *MicrosoftGraphUserSecurityState) SetOnPremisesSecurityIdentifier(v string)`

SetOnPremisesSecurityIdentifier gets a reference to the given string and assigns it to the OnPremisesSecurityIdentifier field.

### SetOnPremisesSecurityIdentifierExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetOnPremisesSecurityIdentifierExplicitNull(b bool)`

SetOnPremisesSecurityIdentifierExplicitNull (un)sets OnPremisesSecurityIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSecurityIdentifier value is set to nil even if false is passed
### GetRiskScore

`func (o *MicrosoftGraphUserSecurityState) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *MicrosoftGraphUserSecurityState) GetRiskScoreOk() (string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskScore

`func (o *MicrosoftGraphUserSecurityState) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### SetRiskScore

`func (o *MicrosoftGraphUserSecurityState) SetRiskScore(v string)`

SetRiskScore gets a reference to the given string and assigns it to the RiskScore field.

### SetRiskScoreExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetRiskScoreExplicitNull(b bool)`

SetRiskScoreExplicitNull (un)sets RiskScore to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskScore value is set to nil even if false is passed
### GetUserAccountType

`func (o *MicrosoftGraphUserSecurityState) GetUserAccountType() AnyOfmicrosoftGraphUserAccountSecurityType`

GetUserAccountType returns the UserAccountType field if non-nil, zero value otherwise.

### GetUserAccountTypeOk

`func (o *MicrosoftGraphUserSecurityState) GetUserAccountTypeOk() (AnyOfmicrosoftGraphUserAccountSecurityType, bool)`

GetUserAccountTypeOk returns a tuple with the UserAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserAccountType

`func (o *MicrosoftGraphUserSecurityState) HasUserAccountType() bool`

HasUserAccountType returns a boolean if a field has been set.

### SetUserAccountType

`func (o *MicrosoftGraphUserSecurityState) SetUserAccountType(v AnyOfmicrosoftGraphUserAccountSecurityType)`

SetUserAccountType gets a reference to the given AnyOfmicrosoftGraphUserAccountSecurityType and assigns it to the UserAccountType field.

### SetUserAccountTypeExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetUserAccountTypeExplicitNull(b bool)`

SetUserAccountTypeExplicitNull (un)sets UserAccountType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserAccountType value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *MicrosoftGraphUserSecurityState) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphUserSecurityState) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *MicrosoftGraphUserSecurityState) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphUserSecurityState) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphUserSecurityState) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


