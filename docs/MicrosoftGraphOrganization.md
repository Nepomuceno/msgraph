# MicrosoftGraphOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DeletedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOrganization) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphOrganization) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDeletedDateTime

`func (o *MicrosoftGraphOrganization) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphOrganization) GetDeletedDateTimeOk() (time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeletedDateTime

`func (o *MicrosoftGraphOrganization) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphOrganization) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime gets a reference to the given time.Time and assigns it to the DeletedDateTime field.

### SetDeletedDateTimeExplicitNull

`func (o *MicrosoftGraphOrganization) SetDeletedDateTimeExplicitNull(b bool)`

SetDeletedDateTimeExplicitNull (un)sets DeletedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeletedDateTime value is set to nil even if false is passed
### GetAssignedPlans

`func (o *MicrosoftGraphOrganization) GetAssignedPlans() []MicrosoftGraphAssignedPlan`

GetAssignedPlans returns the AssignedPlans field if non-nil, zero value otherwise.

### GetAssignedPlansOk

`func (o *MicrosoftGraphOrganization) GetAssignedPlansOk() ([]MicrosoftGraphAssignedPlan, bool)`

GetAssignedPlansOk returns a tuple with the AssignedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedPlans

`func (o *MicrosoftGraphOrganization) HasAssignedPlans() bool`

HasAssignedPlans returns a boolean if a field has been set.

### SetAssignedPlans

`func (o *MicrosoftGraphOrganization) SetAssignedPlans(v []MicrosoftGraphAssignedPlan)`

SetAssignedPlans gets a reference to the given []MicrosoftGraphAssignedPlan and assigns it to the AssignedPlans field.

### GetBusinessPhones

`func (o *MicrosoftGraphOrganization) GetBusinessPhones() []string`

GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.

### GetBusinessPhonesOk

`func (o *MicrosoftGraphOrganization) GetBusinessPhonesOk() ([]string, bool)`

GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessPhones

`func (o *MicrosoftGraphOrganization) HasBusinessPhones() bool`

HasBusinessPhones returns a boolean if a field has been set.

### SetBusinessPhones

`func (o *MicrosoftGraphOrganization) SetBusinessPhones(v []string)`

SetBusinessPhones gets a reference to the given []string and assigns it to the BusinessPhones field.

### GetCity

`func (o *MicrosoftGraphOrganization) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MicrosoftGraphOrganization) GetCityOk() (string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCity

`func (o *MicrosoftGraphOrganization) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCity

`func (o *MicrosoftGraphOrganization) SetCity(v string)`

SetCity gets a reference to the given string and assigns it to the City field.

### SetCityExplicitNull

`func (o *MicrosoftGraphOrganization) SetCityExplicitNull(b bool)`

SetCityExplicitNull (un)sets City to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The City value is set to nil even if false is passed
### GetCountry

`func (o *MicrosoftGraphOrganization) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MicrosoftGraphOrganization) GetCountryOk() (string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCountry

`func (o *MicrosoftGraphOrganization) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountry

`func (o *MicrosoftGraphOrganization) SetCountry(v string)`

SetCountry gets a reference to the given string and assigns it to the Country field.

### SetCountryExplicitNull

`func (o *MicrosoftGraphOrganization) SetCountryExplicitNull(b bool)`

SetCountryExplicitNull (un)sets Country to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Country value is set to nil even if false is passed
### GetCountryLetterCode

`func (o *MicrosoftGraphOrganization) GetCountryLetterCode() string`

GetCountryLetterCode returns the CountryLetterCode field if non-nil, zero value otherwise.

### GetCountryLetterCodeOk

`func (o *MicrosoftGraphOrganization) GetCountryLetterCodeOk() (string, bool)`

GetCountryLetterCodeOk returns a tuple with the CountryLetterCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCountryLetterCode

`func (o *MicrosoftGraphOrganization) HasCountryLetterCode() bool`

HasCountryLetterCode returns a boolean if a field has been set.

### SetCountryLetterCode

`func (o *MicrosoftGraphOrganization) SetCountryLetterCode(v string)`

SetCountryLetterCode gets a reference to the given string and assigns it to the CountryLetterCode field.

### SetCountryLetterCodeExplicitNull

`func (o *MicrosoftGraphOrganization) SetCountryLetterCodeExplicitNull(b bool)`

SetCountryLetterCodeExplicitNull (un)sets CountryLetterCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CountryLetterCode value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphOrganization) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOrganization) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphOrganization) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOrganization) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphOrganization) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphOrganization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphOrganization) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphOrganization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphOrganization) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphOrganization) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetMarketingNotificationEmails

`func (o *MicrosoftGraphOrganization) GetMarketingNotificationEmails() []string`

GetMarketingNotificationEmails returns the MarketingNotificationEmails field if non-nil, zero value otherwise.

### GetMarketingNotificationEmailsOk

`func (o *MicrosoftGraphOrganization) GetMarketingNotificationEmailsOk() ([]string, bool)`

GetMarketingNotificationEmailsOk returns a tuple with the MarketingNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMarketingNotificationEmails

`func (o *MicrosoftGraphOrganization) HasMarketingNotificationEmails() bool`

HasMarketingNotificationEmails returns a boolean if a field has been set.

### SetMarketingNotificationEmails

`func (o *MicrosoftGraphOrganization) SetMarketingNotificationEmails(v []string)`

SetMarketingNotificationEmails gets a reference to the given []string and assigns it to the MarketingNotificationEmails field.

### GetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphOrganization) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *MicrosoftGraphOrganization) GetOnPremisesLastSyncDateTimeOk() (time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphOrganization) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphOrganization) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime gets a reference to the given time.Time and assigns it to the OnPremisesLastSyncDateTime field.

### SetOnPremisesLastSyncDateTimeExplicitNull

`func (o *MicrosoftGraphOrganization) SetOnPremisesLastSyncDateTimeExplicitNull(b bool)`

SetOnPremisesLastSyncDateTimeExplicitNull (un)sets OnPremisesLastSyncDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesLastSyncDateTime value is set to nil even if false is passed
### GetOnPremisesSyncEnabled

`func (o *MicrosoftGraphOrganization) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *MicrosoftGraphOrganization) GetOnPremisesSyncEnabledOk() (bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnPremisesSyncEnabled

`func (o *MicrosoftGraphOrganization) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabled

`func (o *MicrosoftGraphOrganization) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled gets a reference to the given bool and assigns it to the OnPremisesSyncEnabled field.

### SetOnPremisesSyncEnabledExplicitNull

`func (o *MicrosoftGraphOrganization) SetOnPremisesSyncEnabledExplicitNull(b bool)`

SetOnPremisesSyncEnabledExplicitNull (un)sets OnPremisesSyncEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnPremisesSyncEnabled value is set to nil even if false is passed
### GetPostalCode

`func (o *MicrosoftGraphOrganization) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *MicrosoftGraphOrganization) GetPostalCodeOk() (string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPostalCode

`func (o *MicrosoftGraphOrganization) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCode

`func (o *MicrosoftGraphOrganization) SetPostalCode(v string)`

SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.

### SetPostalCodeExplicitNull

`func (o *MicrosoftGraphOrganization) SetPostalCodeExplicitNull(b bool)`

SetPostalCodeExplicitNull (un)sets PostalCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PostalCode value is set to nil even if false is passed
### GetPreferredLanguage

`func (o *MicrosoftGraphOrganization) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *MicrosoftGraphOrganization) GetPreferredLanguageOk() (string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreferredLanguage

`func (o *MicrosoftGraphOrganization) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguage

`func (o *MicrosoftGraphOrganization) SetPreferredLanguage(v string)`

SetPreferredLanguage gets a reference to the given string and assigns it to the PreferredLanguage field.

### SetPreferredLanguageExplicitNull

`func (o *MicrosoftGraphOrganization) SetPreferredLanguageExplicitNull(b bool)`

SetPreferredLanguageExplicitNull (un)sets PreferredLanguage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreferredLanguage value is set to nil even if false is passed
### GetPrivacyProfile

`func (o *MicrosoftGraphOrganization) GetPrivacyProfile() AnyOfmicrosoftGraphPrivacyProfile`

GetPrivacyProfile returns the PrivacyProfile field if non-nil, zero value otherwise.

### GetPrivacyProfileOk

`func (o *MicrosoftGraphOrganization) GetPrivacyProfileOk() (AnyOfmicrosoftGraphPrivacyProfile, bool)`

GetPrivacyProfileOk returns a tuple with the PrivacyProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyProfile

`func (o *MicrosoftGraphOrganization) HasPrivacyProfile() bool`

HasPrivacyProfile returns a boolean if a field has been set.

### SetPrivacyProfile

`func (o *MicrosoftGraphOrganization) SetPrivacyProfile(v AnyOfmicrosoftGraphPrivacyProfile)`

SetPrivacyProfile gets a reference to the given AnyOfmicrosoftGraphPrivacyProfile and assigns it to the PrivacyProfile field.

### SetPrivacyProfileExplicitNull

`func (o *MicrosoftGraphOrganization) SetPrivacyProfileExplicitNull(b bool)`

SetPrivacyProfileExplicitNull (un)sets PrivacyProfile to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyProfile value is set to nil even if false is passed
### GetProvisionedPlans

`func (o *MicrosoftGraphOrganization) GetProvisionedPlans() []MicrosoftGraphProvisionedPlan`

GetProvisionedPlans returns the ProvisionedPlans field if non-nil, zero value otherwise.

### GetProvisionedPlansOk

`func (o *MicrosoftGraphOrganization) GetProvisionedPlansOk() ([]MicrosoftGraphProvisionedPlan, bool)`

GetProvisionedPlansOk returns a tuple with the ProvisionedPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProvisionedPlans

`func (o *MicrosoftGraphOrganization) HasProvisionedPlans() bool`

HasProvisionedPlans returns a boolean if a field has been set.

### SetProvisionedPlans

`func (o *MicrosoftGraphOrganization) SetProvisionedPlans(v []MicrosoftGraphProvisionedPlan)`

SetProvisionedPlans gets a reference to the given []MicrosoftGraphProvisionedPlan and assigns it to the ProvisionedPlans field.

### GetSecurityComplianceNotificationMails

`func (o *MicrosoftGraphOrganization) GetSecurityComplianceNotificationMails() []string`

GetSecurityComplianceNotificationMails returns the SecurityComplianceNotificationMails field if non-nil, zero value otherwise.

### GetSecurityComplianceNotificationMailsOk

`func (o *MicrosoftGraphOrganization) GetSecurityComplianceNotificationMailsOk() ([]string, bool)`

GetSecurityComplianceNotificationMailsOk returns a tuple with the SecurityComplianceNotificationMails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityComplianceNotificationMails

`func (o *MicrosoftGraphOrganization) HasSecurityComplianceNotificationMails() bool`

HasSecurityComplianceNotificationMails returns a boolean if a field has been set.

### SetSecurityComplianceNotificationMails

`func (o *MicrosoftGraphOrganization) SetSecurityComplianceNotificationMails(v []string)`

SetSecurityComplianceNotificationMails gets a reference to the given []string and assigns it to the SecurityComplianceNotificationMails field.

### GetSecurityComplianceNotificationPhones

`func (o *MicrosoftGraphOrganization) GetSecurityComplianceNotificationPhones() []string`

GetSecurityComplianceNotificationPhones returns the SecurityComplianceNotificationPhones field if non-nil, zero value otherwise.

### GetSecurityComplianceNotificationPhonesOk

`func (o *MicrosoftGraphOrganization) GetSecurityComplianceNotificationPhonesOk() ([]string, bool)`

GetSecurityComplianceNotificationPhonesOk returns a tuple with the SecurityComplianceNotificationPhones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityComplianceNotificationPhones

`func (o *MicrosoftGraphOrganization) HasSecurityComplianceNotificationPhones() bool`

HasSecurityComplianceNotificationPhones returns a boolean if a field has been set.

### SetSecurityComplianceNotificationPhones

`func (o *MicrosoftGraphOrganization) SetSecurityComplianceNotificationPhones(v []string)`

SetSecurityComplianceNotificationPhones gets a reference to the given []string and assigns it to the SecurityComplianceNotificationPhones field.

### GetState

`func (o *MicrosoftGraphOrganization) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphOrganization) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *MicrosoftGraphOrganization) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *MicrosoftGraphOrganization) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### SetStateExplicitNull

`func (o *MicrosoftGraphOrganization) SetStateExplicitNull(b bool)`

SetStateExplicitNull (un)sets State to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The State value is set to nil even if false is passed
### GetStreet

`func (o *MicrosoftGraphOrganization) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *MicrosoftGraphOrganization) GetStreetOk() (string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStreet

`func (o *MicrosoftGraphOrganization) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreet

`func (o *MicrosoftGraphOrganization) SetStreet(v string)`

SetStreet gets a reference to the given string and assigns it to the Street field.

### SetStreetExplicitNull

`func (o *MicrosoftGraphOrganization) SetStreetExplicitNull(b bool)`

SetStreetExplicitNull (un)sets Street to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Street value is set to nil even if false is passed
### GetTechnicalNotificationMails

`func (o *MicrosoftGraphOrganization) GetTechnicalNotificationMails() []string`

GetTechnicalNotificationMails returns the TechnicalNotificationMails field if non-nil, zero value otherwise.

### GetTechnicalNotificationMailsOk

`func (o *MicrosoftGraphOrganization) GetTechnicalNotificationMailsOk() ([]string, bool)`

GetTechnicalNotificationMailsOk returns a tuple with the TechnicalNotificationMails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTechnicalNotificationMails

`func (o *MicrosoftGraphOrganization) HasTechnicalNotificationMails() bool`

HasTechnicalNotificationMails returns a boolean if a field has been set.

### SetTechnicalNotificationMails

`func (o *MicrosoftGraphOrganization) SetTechnicalNotificationMails(v []string)`

SetTechnicalNotificationMails gets a reference to the given []string and assigns it to the TechnicalNotificationMails field.

### GetVerifiedDomains

`func (o *MicrosoftGraphOrganization) GetVerifiedDomains() []MicrosoftGraphVerifiedDomain`

GetVerifiedDomains returns the VerifiedDomains field if non-nil, zero value otherwise.

### GetVerifiedDomainsOk

`func (o *MicrosoftGraphOrganization) GetVerifiedDomainsOk() ([]MicrosoftGraphVerifiedDomain, bool)`

GetVerifiedDomainsOk returns a tuple with the VerifiedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVerifiedDomains

`func (o *MicrosoftGraphOrganization) HasVerifiedDomains() bool`

HasVerifiedDomains returns a boolean if a field has been set.

### SetVerifiedDomains

`func (o *MicrosoftGraphOrganization) SetVerifiedDomains(v []MicrosoftGraphVerifiedDomain)`

SetVerifiedDomains gets a reference to the given []MicrosoftGraphVerifiedDomain and assigns it to the VerifiedDomains field.

### GetMobileDeviceManagementAuthority

`func (o *MicrosoftGraphOrganization) GetMobileDeviceManagementAuthority() AnyOfmicrosoftGraphMdmAuthority`

GetMobileDeviceManagementAuthority returns the MobileDeviceManagementAuthority field if non-nil, zero value otherwise.

### GetMobileDeviceManagementAuthorityOk

`func (o *MicrosoftGraphOrganization) GetMobileDeviceManagementAuthorityOk() (AnyOfmicrosoftGraphMdmAuthority, bool)`

GetMobileDeviceManagementAuthorityOk returns a tuple with the MobileDeviceManagementAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobileDeviceManagementAuthority

`func (o *MicrosoftGraphOrganization) HasMobileDeviceManagementAuthority() bool`

HasMobileDeviceManagementAuthority returns a boolean if a field has been set.

### SetMobileDeviceManagementAuthority

`func (o *MicrosoftGraphOrganization) SetMobileDeviceManagementAuthority(v AnyOfmicrosoftGraphMdmAuthority)`

SetMobileDeviceManagementAuthority gets a reference to the given AnyOfmicrosoftGraphMdmAuthority and assigns it to the MobileDeviceManagementAuthority field.

### GetExtensions

`func (o *MicrosoftGraphOrganization) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphOrganization) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *MicrosoftGraphOrganization) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *MicrosoftGraphOrganization) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


