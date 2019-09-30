# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedPlans** | Pointer to [**[]MicrosoftGraphAssignedPlan**](microsoft.graph.assignedPlan.md) |  | [optional] 
**BusinessPhones** | Pointer to **[]string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryLetterCode** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**MarketingNotificationEmails** | Pointer to **[]string** |  | [optional] 
**OnPremisesLastSyncDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**OnPremisesSyncEnabled** | Pointer to **bool** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**PreferredLanguage** | Pointer to **string** |  | [optional] 
**PrivacyProfile** | Pointer to [**AnyOfmicrosoftGraphPrivacyProfile**](anyOf&lt;microsoft.graph.privacyProfile&gt;.md) |  | [optional] 
**ProvisionedPlans** | Pointer to [**[]MicrosoftGraphProvisionedPlan**](microsoft.graph.provisionedPlan.md) |  | [optional] 
**SecurityComplianceNotificationMails** | Pointer to **[]string** |  | [optional] 
**SecurityComplianceNotificationPhones** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Street** | Pointer to **string** |  | [optional] 
**TechnicalNotificationMails** | Pointer to **[]string** |  | [optional] 
**VerifiedDomains** | Pointer to [**[]MicrosoftGraphVerifiedDomain**](microsoft.graph.verifiedDomain.md) |  | [optional] 
**MobileDeviceManagementAuthority** | Pointer to [**AnyOfmicrosoftGraphMdmAuthority**](anyOf&lt;microsoft.graph.mdmAuthority&gt;.md) | Mobile device management authority. | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](microsoft.graph.extension.md) |  | [optional] 

## Methods

### GetAssignedPlans

`func (o *Organization) GetAssignedPlans() []MicrosoftGraphAssignedPlan`

GetAssignedPlans returns the AssignedPlans field if non-nil, zero value otherwise.

### GetAssignedPlansOk

`func (o *Organization) GetAssignedPlansOk() ([]MicrosoftGraphAssignedPlan, bool)`

GetAssignedPlansOk returns a tuple with the AssignedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedPlans

`func (o *Organization) HasAssignedPlans() bool`

HasAssignedPlans returns a boolean if a field has been set.

### SetAssignedPlans

`func (o *Organization) SetAssignedPlans(v []MicrosoftGraphAssignedPlan)`

SetAssignedPlans gets a reference to the given []MicrosoftGraphAssignedPlan and assigns it to the AssignedPlans field.

### GetBusinessPhones

`func (o *Organization) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *Organization) GetBusinessPhonesOk() ([]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessPhones

`func (o *Organization) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### SetBusinessPhones

`func (o *Organization) SetBusinessPhones(v []string)`

SetBusinessPhones gets a reference to the given []string and assigns it to the BusinessPhones field.

### GetCity

`func (o *Organization) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Organization) GetCityOk() (string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCity

`func (o *Organization) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCity

`func (o *Organization) SetCity(v string)`

SetCity gets a reference to the given string and assigns it to the City field.

### SetCityExplicitNull

`func (o *Organization) SetCityExplicitNull(b bool)`

SetCityExplicitNull (un)sets City to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The City value is set to nil even if false is passed
### GetCountry

`func (o *Organization) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Organization) GetCountryOk() (string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCountry

`func (o *Organization) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountry

`func (o *Organization) SetCountry(v string)`

SetCountry gets a reference to the given string and assigns it to the Country field.

### SetCountryExplicitNull

`func (o *Organization) SetCountryExplicitNull(b bool)`

SetCountryExplicitNull (un)sets Country to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Country value is set to nil even if false is passed
### GetCountryLetterCode

`func (o *Organization) GetCountryLetterCode() string`

GetCountryLetterCode returns the CountryLetterCode field if non-nil, zero value otherwise.

### GetCountryLetterCodeOk

`func (o *Organization) GetCountryLetterCodeOk() (string, bool)`

GetCountryLetterCodeOk returns a tuple with the CountryLetterCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCountryLetterCode

`func (o *Organization) HasCountryLetterCode() bool`

HasCountryLetterCode returns a boolean if a field has been set.

### SetCountryLetterCode

`func (o *Organization) SetCountryLetterCode(v string)`

SetCountryLetterCode gets a reference to the given string and assigns it to the CountryLetterCode field.

### SetCountryLetterCodeExplicitNull

`func (o *Organization) SetCountryLetterCodeExplicitNull(b bool)`

SetCountryLetterCodeExplicitNull (un)sets CountryLetterCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CountryLetterCode value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *Organization) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Organization) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *Organization) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *Organization) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *Organization) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetDisplayName

`func (o *Organization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Organization) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *Organization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *Organization) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *Organization) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetMarketingNotificationEmails

`func (o *Organization) GetMarketingNotificationEmails() []string`

GetMarketingNotificationEmails returns the MarketingNotificationEmails field if non-nil, zero value otherwise.

### GetMarketingNotificationEmailsOk

`func (o *Organization) GetMarketingNotificationEmailsOk() ([]string, bool)`

GetMarketingNotificationEmailsOk returns a tuple with the MarketingNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMarketingNotificationEmails

`func (o *Organization) HasMarketingNotificationEmails() bool`

HasMarketingNotificationEmails returns a boolean if a field has been set.

### SetMarketingNotificationEmails

`func (o *Organization) SetMarketingNotificationEmails(v []string)`

SetMarketingNotificationEmails gets a reference to the given []string and assigns it to the MarketingNotificationEmails field.

### GetOnPremisesLastSyncDateTime

`func (o *Organization) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *Organization) GetOnPremisesLastSyncDateTimeOk() (time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesLastSyncDateTime

`func (o *Organization) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTime

`func (o *Organization) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime gets a reference to the given time.Time and assigns it to the OnPremisesLastSyncDateTime field.

### SetOnPremisesLastSyncDateTimeExplicitNull

`func (o *Organization) SetOnPremisesLastSyncDateTimeExplicitNull(b bool)`

SetOnPremisesLastSyncDateTimeExplicitNull (un)sets OnPremisesLastSyncDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesLastSyncDateTime value is set to nil even if false is passed
### GetOnPremisesSyncEnabled

`func (o *Organization) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *Organization) GetOnPremisesSyncEnabledOk() (bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSyncEnabled

`func (o *Organization) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabled

`func (o *Organization) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled gets a reference to the given bool and assigns it to the OnPremisesSyncEnabled field.

### SetOnPremisesSyncEnabledExplicitNull

`func (o *Organization) SetOnPremisesSyncEnabledExplicitNull(b bool)`

SetOnPremisesSyncEnabledExplicitNull (un)sets OnPremisesSyncEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSyncEnabled value is set to nil even if false is passed
### GetPostalCode

`func (o *Organization) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Organization) GetPostalCodeOk() (string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPostalCode

`func (o *Organization) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCode

`func (o *Organization) SetPostalCode(v string)`

SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.

### SetPostalCodeExplicitNull

`func (o *Organization) SetPostalCodeExplicitNull(b bool)`

SetPostalCodeExplicitNull (un)sets PostalCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PostalCode value is set to nil even if false is passed
### GetPreferredLanguage

`func (o *Organization) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *Organization) GetPreferredLanguageOk() (string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreferredLanguage

`func (o *Organization) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguage

`func (o *Organization) SetPreferredLanguage(v string)`

SetPreferredLanguage gets a reference to the given string and assigns it to the PreferredLanguage field.

### SetPreferredLanguageExplicitNull

`func (o *Organization) SetPreferredLanguageExplicitNull(b bool)`

SetPreferredLanguageExplicitNull (un)sets PreferredLanguage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreferredLanguage value is set to nil even if false is passed
### GetPrivacyProfile

`func (o *Organization) GetPrivacyProfile() AnyOfmicrosoftGraphPrivacyProfile`

GetPrivacyProfile returns the PrivacyProfile field if non-nil, zero value otherwise.

### GetPrivacyProfileOk

`func (o *Organization) GetPrivacyProfileOk() (AnyOfmicrosoftGraphPrivacyProfile, bool)`

GetPrivacyProfileOk returns a tuple with the PrivacyProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyProfile

`func (o *Organization) HasPrivacyProfile() bool`

HasPrivacyProfile returns a boolean if a field has been set.

### SetPrivacyProfile

`func (o *Organization) SetPrivacyProfile(v AnyOfmicrosoftGraphPrivacyProfile)`

SetPrivacyProfile gets a reference to the given AnyOfmicrosoftGraphPrivacyProfile and assigns it to the PrivacyProfile field.

### SetPrivacyProfileExplicitNull

`func (o *Organization) SetPrivacyProfileExplicitNull(b bool)`

SetPrivacyProfileExplicitNull (un)sets PrivacyProfile to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyProfile value is set to nil even if false is passed
### GetProvisionedPlans

`func (o *Organization) GetProvisionedPlans() []MicrosoftGraphProvisionedPlan`

GetProvisionedPlans returns the ProvisionedPlans field if non-nil, zero value otherwise.

### GetProvisionedPlansOk

`func (o *Organization) GetProvisionedPlansOk() ([]MicrosoftGraphProvisionedPlan, bool)`

GetProvisionedPlansOk returns a tuple with the ProvisionedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProvisionedPlans

`func (o *Organization) HasProvisionedPlans() bool`

HasProvisionedPlans returns a boolean if a field has been set.

### SetProvisionedPlans

`func (o *Organization) SetProvisionedPlans(v []MicrosoftGraphProvisionedPlan)`

SetProvisionedPlans gets a reference to the given []MicrosoftGraphProvisionedPlan and assigns it to the ProvisionedPlans field.

### GetSecurityComplianceNotificationMails

`func (o *Organization) GetSecurityComplianceNotificationMails() []string`

GetSecurityComplianceNotificationMails returns the SecurityComplianceNotificationMails field if non-nil, zero value otherwise.

### GetSecurityComplianceNotificationMailsOk

`func (o *Organization) GetSecurityComplianceNotificationMailsOk() ([]string, bool)`

GetSecurityComplianceNotificationMailsOk returns a tuple with the SecurityComplianceNotificationMails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityComplianceNotificationMails

`func (o *Organization) HasSecurityComplianceNotificationMails() bool`

HasSecurityComplianceNotificationMails returns a boolean if a field has been set.

### SetSecurityComplianceNotificationMails

`func (o *Organization) SetSecurityComplianceNotificationMails(v []string)`

SetSecurityComplianceNotificationMails gets a reference to the given []string and assigns it to the SecurityComplianceNotificationMails field.

### GetSecurityComplianceNotificationPhones

`func (o *Organization) GetSecurityComplianceNotificationPhones() []string`

GetSecurityComplianceNotificationPhones returns the SecurityComplianceNotificationPhones field if non-nil, zero value otherwise.

### GetSecurityComplianceNotificationPhonesOk

`func (o *Organization) GetSecurityComplianceNotificationPhonesOk() ([]string, bool)`

GetSecurityComplianceNotificationPhonesOk returns a tuple with the SecurityComplianceNotificationPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityComplianceNotificationPhones

`func (o *Organization) HasSecurityComplianceNotificationPhones() bool`

HasSecurityComplianceNotificationPhones returns a boolean if a field has been set.

### SetSecurityComplianceNotificationPhones

`func (o *Organization) SetSecurityComplianceNotificationPhones(v []string)`

SetSecurityComplianceNotificationPhones gets a reference to the given []string and assigns it to the SecurityComplianceNotificationPhones field.

### GetState

`func (o *Organization) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Organization) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *Organization) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *Organization) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### SetStateExplicitNull

`func (o *Organization) SetStateExplicitNull(b bool)`

SetStateExplicitNull (un)sets State to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The State value is set to nil even if false is passed
### GetStreet

`func (o *Organization) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *Organization) GetStreetOk() (string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStreet

`func (o *Organization) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreet

`func (o *Organization) SetStreet(v string)`

SetStreet gets a reference to the given string and assigns it to the Street field.

### SetStreetExplicitNull

`func (o *Organization) SetStreetExplicitNull(b bool)`

SetStreetExplicitNull (un)sets Street to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Street value is set to nil even if false is passed
### GetTechnicalNotificationMails

`func (o *Organization) GetTechnicalNotificationMails() []string`

GetTechnicalNotificationMails returns the TechnicalNotificationMails field if non-nil, zero value otherwise.

### GetTechnicalNotificationMailsOk

`func (o *Organization) GetTechnicalNotificationMailsOk() ([]string, bool)`

GetTechnicalNotificationMailsOk returns a tuple with the TechnicalNotificationMails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTechnicalNotificationMails

`func (o *Organization) HasTechnicalNotificationMails() bool`

HasTechnicalNotificationMails returns a boolean if a field has been set.

### SetTechnicalNotificationMails

`func (o *Organization) SetTechnicalNotificationMails(v []string)`

SetTechnicalNotificationMails gets a reference to the given []string and assigns it to the TechnicalNotificationMails field.

### GetVerifiedDomains

`func (o *Organization) GetVerifiedDomains() []MicrosoftGraphVerifiedDomain`

GetVerifiedDomains returns the VerifiedDomains field if non-nil, zero value otherwise.

### GetVerifiedDomainsOk

`func (o *Organization) GetVerifiedDomainsOk() ([]MicrosoftGraphVerifiedDomain, bool)`

GetVerifiedDomainsOk returns a tuple with the VerifiedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVerifiedDomains

`func (o *Organization) HasVerifiedDomains() bool`

HasVerifiedDomains returns a boolean if a field has been set.

### SetVerifiedDomains

`func (o *Organization) SetVerifiedDomains(v []MicrosoftGraphVerifiedDomain)`

SetVerifiedDomains gets a reference to the given []MicrosoftGraphVerifiedDomain and assigns it to the VerifiedDomains field.

### GetMobileDeviceManagementAuthority

`func (o *Organization) GetMobileDeviceManagementAuthority() AnyOfmicrosoftGraphMdmAuthority`

GetMobileDeviceManagementAuthority returns the MobileDeviceManagementAuthority field if non-nil, zero value otherwise.

### GetMobileDeviceManagementAuthorityOk

`func (o *Organization) GetMobileDeviceManagementAuthorityOk() (AnyOfmicrosoftGraphMdmAuthority, bool)`

GetMobileDeviceManagementAuthorityOk returns a tuple with the MobileDeviceManagementAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobileDeviceManagementAuthority

`func (o *Organization) HasMobileDeviceManagementAuthority() bool`

HasMobileDeviceManagementAuthority returns a boolean if a field has been set.

### SetMobileDeviceManagementAuthority

`func (o *Organization) SetMobileDeviceManagementAuthority(v AnyOfmicrosoftGraphMdmAuthority)`

SetMobileDeviceManagementAuthority gets a reference to the given AnyOfmicrosoftGraphMdmAuthority and assigns it to the MobileDeviceManagementAuthority field.

### GetExtensions

`func (o *Organization) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Organization) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *Organization) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *Organization) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


