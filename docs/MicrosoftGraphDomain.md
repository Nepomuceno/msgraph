# MicrosoftGraphDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AuthenticationType** | Pointer to **string** |  | [optional] 
**AvailabilityStatus** | Pointer to **string** |  | [optional] 
**IsAdminManaged** | Pointer to **bool** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsInitial** | Pointer to **bool** |  | [optional] 
**IsRoot** | Pointer to **bool** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] 
**PasswordNotificationWindowInDays** | Pointer to **int32** |  | [optional] 
**PasswordValidityPeriodInDays** | Pointer to **int32** |  | [optional] 
**SupportedServices** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphDomainState**](anyOf&lt;microsoft.graph.domainState&gt;.md) |  | [optional] 
**ServiceConfigurationRecords** | Pointer to [**[]MicrosoftGraphDomainDnsRecord**](microsoft.graph.domainDnsRecord.md) |  | [optional] 
**VerificationDnsRecords** | Pointer to [**[]MicrosoftGraphDomainDnsRecord**](microsoft.graph.domainDnsRecord.md) |  | [optional] 
**DomainNameReferences** | Pointer to [**[]MicrosoftGraphDirectoryObject**](microsoft.graph.directoryObject.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDomain) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDomain) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetAuthenticationType

`func (o *MicrosoftGraphDomain) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *MicrosoftGraphDomain) GetAuthenticationTypeOk() (string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAuthenticationType

`func (o *MicrosoftGraphDomain) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationType

`func (o *MicrosoftGraphDomain) SetAuthenticationType(v string)`

SetAuthenticationType gets a reference to the given string and assigns it to the AuthenticationType field.

### GetAvailabilityStatus

`func (o *MicrosoftGraphDomain) GetAvailabilityStatus() string`

GetAvailabilityStatus returns the AvailabilityStatus field if non-nil, zero value otherwise.

### GetAvailabilityStatusOk

`func (o *MicrosoftGraphDomain) GetAvailabilityStatusOk() (string, bool)`

GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAvailabilityStatus

`func (o *MicrosoftGraphDomain) HasAvailabilityStatus() bool`

HasAvailabilityStatus returns a boolean if a field has been set.

### SetAvailabilityStatus

`func (o *MicrosoftGraphDomain) SetAvailabilityStatus(v string)`

SetAvailabilityStatus gets a reference to the given string and assigns it to the AvailabilityStatus field.

### SetAvailabilityStatusExplicitNull

`func (o *MicrosoftGraphDomain) SetAvailabilityStatusExplicitNull(b bool)`

SetAvailabilityStatusExplicitNull (un)sets AvailabilityStatus to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AvailabilityStatus value is set to nil even if false is passed
### GetIsAdminManaged

`func (o *MicrosoftGraphDomain) GetIsAdminManaged() bool`

GetIsAdminManaged returns the IsAdminManaged field if non-nil, zero value otherwise.

### GetIsAdminManagedOk

`func (o *MicrosoftGraphDomain) GetIsAdminManagedOk() (bool, bool)`

GetIsAdminManagedOk returns a tuple with the IsAdminManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAdminManaged

`func (o *MicrosoftGraphDomain) HasIsAdminManaged() bool`

HasIsAdminManaged returns a boolean if a field has been set.

### SetIsAdminManaged

`func (o *MicrosoftGraphDomain) SetIsAdminManaged(v bool)`

SetIsAdminManaged gets a reference to the given bool and assigns it to the IsAdminManaged field.

### GetIsDefault

`func (o *MicrosoftGraphDomain) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphDomain) GetIsDefaultOk() (bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDefault

`func (o *MicrosoftGraphDomain) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefault

`func (o *MicrosoftGraphDomain) SetIsDefault(v bool)`

SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.

### GetIsInitial

`func (o *MicrosoftGraphDomain) GetIsInitial() bool`

GetIsInitial returns the IsInitial field if non-nil, zero value otherwise.

### GetIsInitialOk

`func (o *MicrosoftGraphDomain) GetIsInitialOk() (bool, bool)`

GetIsInitialOk returns a tuple with the IsInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsInitial

`func (o *MicrosoftGraphDomain) HasIsInitial() bool`

HasIsInitial returns a boolean if a field has been set.

### SetIsInitial

`func (o *MicrosoftGraphDomain) SetIsInitial(v bool)`

SetIsInitial gets a reference to the given bool and assigns it to the IsInitial field.

### GetIsRoot

`func (o *MicrosoftGraphDomain) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *MicrosoftGraphDomain) GetIsRootOk() (bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsRoot

`func (o *MicrosoftGraphDomain) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### SetIsRoot

`func (o *MicrosoftGraphDomain) SetIsRoot(v bool)`

SetIsRoot gets a reference to the given bool and assigns it to the IsRoot field.

### GetIsVerified

`func (o *MicrosoftGraphDomain) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *MicrosoftGraphDomain) GetIsVerifiedOk() (bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsVerified

`func (o *MicrosoftGraphDomain) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### SetIsVerified

`func (o *MicrosoftGraphDomain) SetIsVerified(v bool)`

SetIsVerified gets a reference to the given bool and assigns it to the IsVerified field.

### GetPasswordNotificationWindowInDays

`func (o *MicrosoftGraphDomain) GetPasswordNotificationWindowInDays() int32`

GetPasswordNotificationWindowInDays returns the PasswordNotificationWindowInDays field if non-nil, zero value otherwise.

### GetPasswordNotificationWindowInDaysOk

`func (o *MicrosoftGraphDomain) GetPasswordNotificationWindowInDaysOk() (int32, bool)`

GetPasswordNotificationWindowInDaysOk returns a tuple with the PasswordNotificationWindowInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordNotificationWindowInDays

`func (o *MicrosoftGraphDomain) HasPasswordNotificationWindowInDays() bool`

HasPasswordNotificationWindowInDays returns a boolean if a field has been set.

### SetPasswordNotificationWindowInDays

`func (o *MicrosoftGraphDomain) SetPasswordNotificationWindowInDays(v int32)`

SetPasswordNotificationWindowInDays gets a reference to the given int32 and assigns it to the PasswordNotificationWindowInDays field.

### SetPasswordNotificationWindowInDaysExplicitNull

`func (o *MicrosoftGraphDomain) SetPasswordNotificationWindowInDaysExplicitNull(b bool)`

SetPasswordNotificationWindowInDaysExplicitNull (un)sets PasswordNotificationWindowInDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordNotificationWindowInDays value is set to nil even if false is passed
### GetPasswordValidityPeriodInDays

`func (o *MicrosoftGraphDomain) GetPasswordValidityPeriodInDays() int32`

GetPasswordValidityPeriodInDays returns the PasswordValidityPeriodInDays field if non-nil, zero value otherwise.

### GetPasswordValidityPeriodInDaysOk

`func (o *MicrosoftGraphDomain) GetPasswordValidityPeriodInDaysOk() (int32, bool)`

GetPasswordValidityPeriodInDaysOk returns a tuple with the PasswordValidityPeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordValidityPeriodInDays

`func (o *MicrosoftGraphDomain) HasPasswordValidityPeriodInDays() bool`

HasPasswordValidityPeriodInDays returns a boolean if a field has been set.

### SetPasswordValidityPeriodInDays

`func (o *MicrosoftGraphDomain) SetPasswordValidityPeriodInDays(v int32)`

SetPasswordValidityPeriodInDays gets a reference to the given int32 and assigns it to the PasswordValidityPeriodInDays field.

### SetPasswordValidityPeriodInDaysExplicitNull

`func (o *MicrosoftGraphDomain) SetPasswordValidityPeriodInDaysExplicitNull(b bool)`

SetPasswordValidityPeriodInDaysExplicitNull (un)sets PasswordValidityPeriodInDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordValidityPeriodInDays value is set to nil even if false is passed
### GetSupportedServices

`func (o *MicrosoftGraphDomain) GetSupportedServices() []string`

GetSupportedServices returns the SupportedServices field if non-nil, zero value otherwise.

### GetSupportedServicesOk

`func (o *MicrosoftGraphDomain) GetSupportedServicesOk() ([]string, bool)`

GetSupportedServicesOk returns a tuple with the SupportedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSupportedServices

`func (o *MicrosoftGraphDomain) HasSupportedServices() bool`

HasSupportedServices returns a boolean if a field has been set.

### SetSupportedServices

`func (o *MicrosoftGraphDomain) SetSupportedServices(v []string)`

SetSupportedServices gets a reference to the given []string and assigns it to the SupportedServices field.

### GetState

`func (o *MicrosoftGraphDomain) GetState() AnyOfmicrosoftGraphDomainState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphDomain) GetStateOk() (AnyOfmicrosoftGraphDomainState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *MicrosoftGraphDomain) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *MicrosoftGraphDomain) SetState(v AnyOfmicrosoftGraphDomainState)`

SetState gets a reference to the given AnyOfmicrosoftGraphDomainState and assigns it to the State field.

### SetStateExplicitNull

`func (o *MicrosoftGraphDomain) SetStateExplicitNull(b bool)`

SetStateExplicitNull (un)sets State to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The State value is set to nil even if false is passed
### GetServiceConfigurationRecords

`func (o *MicrosoftGraphDomain) GetServiceConfigurationRecords() []MicrosoftGraphDomainDnsRecord`

GetServiceConfigurationRecords returns the ServiceConfigurationRecords field if non-nil, zero value otherwise.

### GetServiceConfigurationRecordsOk

`func (o *MicrosoftGraphDomain) GetServiceConfigurationRecordsOk() ([]MicrosoftGraphDomainDnsRecord, bool)`

GetServiceConfigurationRecordsOk returns a tuple with the ServiceConfigurationRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceConfigurationRecords

`func (o *MicrosoftGraphDomain) HasServiceConfigurationRecords() bool`

HasServiceConfigurationRecords returns a boolean if a field has been set.

### SetServiceConfigurationRecords

`func (o *MicrosoftGraphDomain) SetServiceConfigurationRecords(v []MicrosoftGraphDomainDnsRecord)`

SetServiceConfigurationRecords gets a reference to the given []MicrosoftGraphDomainDnsRecord and assigns it to the ServiceConfigurationRecords field.

### GetVerificationDnsRecords

`func (o *MicrosoftGraphDomain) GetVerificationDnsRecords() []MicrosoftGraphDomainDnsRecord`

GetVerificationDnsRecords returns the VerificationDnsRecords field if non-nil, zero value otherwise.

### GetVerificationDnsRecordsOk

`func (o *MicrosoftGraphDomain) GetVerificationDnsRecordsOk() ([]MicrosoftGraphDomainDnsRecord, bool)`

GetVerificationDnsRecordsOk returns a tuple with the VerificationDnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVerificationDnsRecords

`func (o *MicrosoftGraphDomain) HasVerificationDnsRecords() bool`

HasVerificationDnsRecords returns a boolean if a field has been set.

### SetVerificationDnsRecords

`func (o *MicrosoftGraphDomain) SetVerificationDnsRecords(v []MicrosoftGraphDomainDnsRecord)`

SetVerificationDnsRecords gets a reference to the given []MicrosoftGraphDomainDnsRecord and assigns it to the VerificationDnsRecords field.

### GetDomainNameReferences

`func (o *MicrosoftGraphDomain) GetDomainNameReferences() []MicrosoftGraphDirectoryObject`

GetDomainNameReferences returns the DomainNameReferences field if non-nil, zero value otherwise.

### GetDomainNameReferencesOk

`func (o *MicrosoftGraphDomain) GetDomainNameReferencesOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetDomainNameReferencesOk returns a tuple with the DomainNameReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomainNameReferences

`func (o *MicrosoftGraphDomain) HasDomainNameReferences() bool`

HasDomainNameReferences returns a boolean if a field has been set.

### SetDomainNameReferences

`func (o *MicrosoftGraphDomain) SetDomainNameReferences(v []MicrosoftGraphDirectoryObject)`

SetDomainNameReferences gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the DomainNameReferences field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


