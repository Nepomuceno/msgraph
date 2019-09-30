# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetAuthenticationType

`func (o *Domain) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *Domain) GetAuthenticationTypeOk() (string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAuthenticationType

`func (o *Domain) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationType

`func (o *Domain) SetAuthenticationType(v string)`

SetAuthenticationType gets a reference to the given string and assigns it to the AuthenticationType field.

### GetAvailabilityStatus

`func (o *Domain) GetAvailabilityStatus() string`

GetAvailabilityStatus returns the AvailabilityStatus field if non-nil, zero value otherwise.

### GetAvailabilityStatusOk

`func (o *Domain) GetAvailabilityStatusOk() (string, bool)`

GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAvailabilityStatus

`func (o *Domain) HasAvailabilityStatus() bool`

HasAvailabilityStatus returns a boolean if a field has been set.

### SetAvailabilityStatus

`func (o *Domain) SetAvailabilityStatus(v string)`

SetAvailabilityStatus gets a reference to the given string and assigns it to the AvailabilityStatus field.

### SetAvailabilityStatusExplicitNull

`func (o *Domain) SetAvailabilityStatusExplicitNull(b bool)`

SetAvailabilityStatusExplicitNull (un)sets AvailabilityStatus to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AvailabilityStatus value is set to nil even if false is passed
### GetIsAdminManaged

`func (o *Domain) GetIsAdminManaged() bool`

GetIsAdminManaged returns the IsAdminManaged field if non-nil, zero value otherwise.

### GetIsAdminManagedOk

`func (o *Domain) GetIsAdminManagedOk() (bool, bool)`

GetIsAdminManagedOk returns a tuple with the IsAdminManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAdminManaged

`func (o *Domain) HasIsAdminManaged() bool`

HasIsAdminManaged returns a boolean if a field has been set.

### SetIsAdminManaged

`func (o *Domain) SetIsAdminManaged(v bool)`

SetIsAdminManaged gets a reference to the given bool and assigns it to the IsAdminManaged field.

### GetIsDefault

`func (o *Domain) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Domain) GetIsDefaultOk() (bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDefault

`func (o *Domain) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefault

`func (o *Domain) SetIsDefault(v bool)`

SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.

### GetIsInitial

`func (o *Domain) GetIsInitial() bool`

GetIsInitial returns the IsInitial field if non-nil, zero value otherwise.

### GetIsInitialOk

`func (o *Domain) GetIsInitialOk() (bool, bool)`

GetIsInitialOk returns a tuple with the IsInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsInitial

`func (o *Domain) HasIsInitial() bool`

HasIsInitial returns a boolean if a field has been set.

### SetIsInitial

`func (o *Domain) SetIsInitial(v bool)`

SetIsInitial gets a reference to the given bool and assigns it to the IsInitial field.

### GetIsRoot

`func (o *Domain) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *Domain) GetIsRootOk() (bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsRoot

`func (o *Domain) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### SetIsRoot

`func (o *Domain) SetIsRoot(v bool)`

SetIsRoot gets a reference to the given bool and assigns it to the IsRoot field.

### GetIsVerified

`func (o *Domain) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *Domain) GetIsVerifiedOk() (bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsVerified

`func (o *Domain) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### SetIsVerified

`func (o *Domain) SetIsVerified(v bool)`

SetIsVerified gets a reference to the given bool and assigns it to the IsVerified field.

### GetPasswordNotificationWindowInDays

`func (o *Domain) GetPasswordNotificationWindowInDays() int32`

GetPasswordNotificationWindowInDays returns the PasswordNotificationWindowInDays field if non-nil, zero value otherwise.

### GetPasswordNotificationWindowInDaysOk

`func (o *Domain) GetPasswordNotificationWindowInDaysOk() (int32, bool)`

GetPasswordNotificationWindowInDaysOk returns a tuple with the PasswordNotificationWindowInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordNotificationWindowInDays

`func (o *Domain) HasPasswordNotificationWindowInDays() bool`

HasPasswordNotificationWindowInDays returns a boolean if a field has been set.

### SetPasswordNotificationWindowInDays

`func (o *Domain) SetPasswordNotificationWindowInDays(v int32)`

SetPasswordNotificationWindowInDays gets a reference to the given int32 and assigns it to the PasswordNotificationWindowInDays field.

### SetPasswordNotificationWindowInDaysExplicitNull

`func (o *Domain) SetPasswordNotificationWindowInDaysExplicitNull(b bool)`

SetPasswordNotificationWindowInDaysExplicitNull (un)sets PasswordNotificationWindowInDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordNotificationWindowInDays value is set to nil even if false is passed
### GetPasswordValidityPeriodInDays

`func (o *Domain) GetPasswordValidityPeriodInDays() int32`

GetPasswordValidityPeriodInDays returns the PasswordValidityPeriodInDays field if non-nil, zero value otherwise.

### GetPasswordValidityPeriodInDaysOk

`func (o *Domain) GetPasswordValidityPeriodInDaysOk() (int32, bool)`

GetPasswordValidityPeriodInDaysOk returns a tuple with the PasswordValidityPeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordValidityPeriodInDays

`func (o *Domain) HasPasswordValidityPeriodInDays() bool`

HasPasswordValidityPeriodInDays returns a boolean if a field has been set.

### SetPasswordValidityPeriodInDays

`func (o *Domain) SetPasswordValidityPeriodInDays(v int32)`

SetPasswordValidityPeriodInDays gets a reference to the given int32 and assigns it to the PasswordValidityPeriodInDays field.

### SetPasswordValidityPeriodInDaysExplicitNull

`func (o *Domain) SetPasswordValidityPeriodInDaysExplicitNull(b bool)`

SetPasswordValidityPeriodInDaysExplicitNull (un)sets PasswordValidityPeriodInDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordValidityPeriodInDays value is set to nil even if false is passed
### GetSupportedServices

`func (o *Domain) GetSupportedServices() []string`

GetSupportedServices returns the SupportedServices field if non-nil, zero value otherwise.

### GetSupportedServicesOk

`func (o *Domain) GetSupportedServicesOk() ([]string, bool)`

GetSupportedServicesOk returns a tuple with the SupportedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSupportedServices

`func (o *Domain) HasSupportedServices() bool`

HasSupportedServices returns a boolean if a field has been set.

### SetSupportedServices

`func (o *Domain) SetSupportedServices(v []string)`

SetSupportedServices gets a reference to the given []string and assigns it to the SupportedServices field.

### GetState

`func (o *Domain) GetState() AnyOfmicrosoftGraphDomainState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Domain) GetStateOk() (AnyOfmicrosoftGraphDomainState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *Domain) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *Domain) SetState(v AnyOfmicrosoftGraphDomainState)`

SetState gets a reference to the given AnyOfmicrosoftGraphDomainState and assigns it to the State field.

### SetStateExplicitNull

`func (o *Domain) SetStateExplicitNull(b bool)`

SetStateExplicitNull (un)sets State to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The State value is set to nil even if false is passed
### GetServiceConfigurationRecords

`func (o *Domain) GetServiceConfigurationRecords() []MicrosoftGraphDomainDnsRecord`

GetServiceConfigurationRecords returns the ServiceConfigurationRecords field if non-nil, zero value otherwise.

### GetServiceConfigurationRecordsOk

`func (o *Domain) GetServiceConfigurationRecordsOk() ([]MicrosoftGraphDomainDnsRecord, bool)`

GetServiceConfigurationRecordsOk returns a tuple with the ServiceConfigurationRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceConfigurationRecords

`func (o *Domain) HasServiceConfigurationRecords() bool`

HasServiceConfigurationRecords returns a boolean if a field has been set.

### SetServiceConfigurationRecords

`func (o *Domain) SetServiceConfigurationRecords(v []MicrosoftGraphDomainDnsRecord)`

SetServiceConfigurationRecords gets a reference to the given []MicrosoftGraphDomainDnsRecord and assigns it to the ServiceConfigurationRecords field.

### GetVerificationDnsRecords

`func (o *Domain) GetVerificationDnsRecords() []MicrosoftGraphDomainDnsRecord`

GetVerificationDnsRecords returns the VerificationDnsRecords field if non-nil, zero value otherwise.

### GetVerificationDnsRecordsOk

`func (o *Domain) GetVerificationDnsRecordsOk() ([]MicrosoftGraphDomainDnsRecord, bool)`

GetVerificationDnsRecordsOk returns a tuple with the VerificationDnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVerificationDnsRecords

`func (o *Domain) HasVerificationDnsRecords() bool`

HasVerificationDnsRecords returns a boolean if a field has been set.

### SetVerificationDnsRecords

`func (o *Domain) SetVerificationDnsRecords(v []MicrosoftGraphDomainDnsRecord)`

SetVerificationDnsRecords gets a reference to the given []MicrosoftGraphDomainDnsRecord and assigns it to the VerificationDnsRecords field.

### GetDomainNameReferences

`func (o *Domain) GetDomainNameReferences() []MicrosoftGraphDirectoryObject`

GetDomainNameReferences returns the DomainNameReferences field if non-nil, zero value otherwise.

### GetDomainNameReferencesOk

`func (o *Domain) GetDomainNameReferencesOk() ([]MicrosoftGraphDirectoryObject, bool)`

GetDomainNameReferencesOk returns a tuple with the DomainNameReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomainNameReferences

`func (o *Domain) HasDomainNameReferences() bool`

HasDomainNameReferences returns a boolean if a field has been set.

### SetDomainNameReferences

`func (o *Domain) SetDomainNameReferences(v []MicrosoftGraphDirectoryObject)`

SetDomainNameReferences gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the DomainNameReferences field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


