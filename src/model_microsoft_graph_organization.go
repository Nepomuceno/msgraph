/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
import (
	"time"
	"encoding/json"
)
// MicrosoftGraphOrganization struct for MicrosoftGraphOrganization
type MicrosoftGraphOrganization struct {
	Id *string `json:"id,omitempty"`

	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
	isExplicitNullDeletedDateTime bool `json:"-"`
	AssignedPlans *[]MicrosoftGraphAssignedPlan `json:"assignedPlans,omitempty"`

	BusinessPhones *[]string `json:"businessPhones,omitempty"`

	City *string `json:"city,omitempty"`
	isExplicitNullCity bool `json:"-"`
	Country *string `json:"country,omitempty"`
	isExplicitNullCountry bool `json:"-"`
	CountryLetterCode *string `json:"countryLetterCode,omitempty"`
	isExplicitNullCountryLetterCode bool `json:"-"`
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	isExplicitNullCreatedDateTime bool `json:"-"`
	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	MarketingNotificationEmails *[]string `json:"marketingNotificationEmails,omitempty"`

	OnPremisesLastSyncDateTime *time.Time `json:"onPremisesLastSyncDateTime,omitempty"`
	isExplicitNullOnPremisesLastSyncDateTime bool `json:"-"`
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`
	isExplicitNullOnPremisesSyncEnabled bool `json:"-"`
	PostalCode *string `json:"postalCode,omitempty"`
	isExplicitNullPostalCode bool `json:"-"`
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	isExplicitNullPreferredLanguage bool `json:"-"`
	PrivacyProfile *AnyOfmicrosoftGraphPrivacyProfile `json:"privacyProfile,omitempty"`
	isExplicitNullPrivacyProfile bool `json:"-"`
	ProvisionedPlans *[]MicrosoftGraphProvisionedPlan `json:"provisionedPlans,omitempty"`

	SecurityComplianceNotificationMails *[]string `json:"securityComplianceNotificationMails,omitempty"`

	SecurityComplianceNotificationPhones *[]string `json:"securityComplianceNotificationPhones,omitempty"`

	State *string `json:"state,omitempty"`
	isExplicitNullState bool `json:"-"`
	Street *string `json:"street,omitempty"`
	isExplicitNullStreet bool `json:"-"`
	TechnicalNotificationMails *[]string `json:"technicalNotificationMails,omitempty"`

	VerifiedDomains *[]MicrosoftGraphVerifiedDomain `json:"verifiedDomains,omitempty"`

	// Mobile device management authority.
	MobileDeviceManagementAuthority *AnyOfmicrosoftGraphMdmAuthority `json:"mobileDeviceManagementAuthority,omitempty"`

	Extensions *[]MicrosoftGraphExtension `json:"extensions,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOrganization) SetId(v string) {
	o.Id = &v
}

// GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetDeletedDateTime() time.Time {
	if o == nil || o.DeletedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedDateTime
}

// GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetDeletedDateTimeOk() (time.Time, bool) {
	if o == nil || o.DeletedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.DeletedDateTime, true
}

// HasDeletedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasDeletedDateTime() bool {
	if o != nil && o.DeletedDateTime != nil {
		return true
	}

	return false
}

// SetDeletedDateTime gets a reference to the given time.Time and assigns it to the DeletedDateTime field.
func (o *MicrosoftGraphOrganization) SetDeletedDateTime(v time.Time) {
	o.DeletedDateTime = &v
}

// SetDeletedDateTimeExplicitNull (un)sets DeletedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DeletedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetDeletedDateTimeExplicitNull(b bool) {
	o.DeletedDateTime = nil
	o.isExplicitNullDeletedDateTime = b
}
// GetAssignedPlans returns the AssignedPlans field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetAssignedPlans() []MicrosoftGraphAssignedPlan {
	if o == nil || o.AssignedPlans == nil {
		var ret []MicrosoftGraphAssignedPlan
		return ret
	}
	return *o.AssignedPlans
}

// GetAssignedPlansOk returns a tuple with the AssignedPlans field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetAssignedPlansOk() ([]MicrosoftGraphAssignedPlan, bool) {
	if o == nil || o.AssignedPlans == nil {
		var ret []MicrosoftGraphAssignedPlan
		return ret, false
	}
	return *o.AssignedPlans, true
}

// HasAssignedPlans returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasAssignedPlans() bool {
	if o != nil && o.AssignedPlans != nil {
		return true
	}

	return false
}

// SetAssignedPlans gets a reference to the given []MicrosoftGraphAssignedPlan and assigns it to the AssignedPlans field.
func (o *MicrosoftGraphOrganization) SetAssignedPlans(v []MicrosoftGraphAssignedPlan) {
	o.AssignedPlans = &v
}

// GetBusinessPhones returns the BusinessPhones field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetBusinessPhones() []string {
	if o == nil || o.BusinessPhones == nil {
		var ret []string
		return ret
	}
	return *o.BusinessPhones
}

// GetBusinessPhonesOk returns a tuple with the BusinessPhones field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetBusinessPhonesOk() ([]string, bool) {
	if o == nil || o.BusinessPhones == nil {
		var ret []string
		return ret, false
	}
	return *o.BusinessPhones, true
}

// HasBusinessPhones returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasBusinessPhones() bool {
	if o != nil && o.BusinessPhones != nil {
		return true
	}

	return false
}

// SetBusinessPhones gets a reference to the given []string and assigns it to the BusinessPhones field.
func (o *MicrosoftGraphOrganization) SetBusinessPhones(v []string) {
	o.BusinessPhones = &v
}

// GetCity returns the City field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetCityOk() (string, bool) {
	if o == nil || o.City == nil {
		var ret string
		return ret, false
	}
	return *o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *MicrosoftGraphOrganization) SetCity(v string) {
	o.City = &v
}

// SetCityExplicitNull (un)sets City to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The City value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetCityExplicitNull(b bool) {
	o.City = nil
	o.isExplicitNullCity = b
}
// GetCountry returns the Country field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetCountryOk() (string, bool) {
	if o == nil || o.Country == nil {
		var ret string
		return ret, false
	}
	return *o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *MicrosoftGraphOrganization) SetCountry(v string) {
	o.Country = &v
}

// SetCountryExplicitNull (un)sets Country to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Country value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetCountryExplicitNull(b bool) {
	o.Country = nil
	o.isExplicitNullCountry = b
}
// GetCountryLetterCode returns the CountryLetterCode field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetCountryLetterCode() string {
	if o == nil || o.CountryLetterCode == nil {
		var ret string
		return ret
	}
	return *o.CountryLetterCode
}

// GetCountryLetterCodeOk returns a tuple with the CountryLetterCode field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetCountryLetterCodeOk() (string, bool) {
	if o == nil || o.CountryLetterCode == nil {
		var ret string
		return ret, false
	}
	return *o.CountryLetterCode, true
}

// HasCountryLetterCode returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasCountryLetterCode() bool {
	if o != nil && o.CountryLetterCode != nil {
		return true
	}

	return false
}

// SetCountryLetterCode gets a reference to the given string and assigns it to the CountryLetterCode field.
func (o *MicrosoftGraphOrganization) SetCountryLetterCode(v string) {
	o.CountryLetterCode = &v
}

// SetCountryLetterCodeExplicitNull (un)sets CountryLetterCode to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CountryLetterCode value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetCountryLetterCodeExplicitNull(b bool) {
	o.CountryLetterCode = nil
	o.isExplicitNullCountryLetterCode = b
}
// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphOrganization) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetCreatedDateTimeExplicitNull(b bool) {
	o.CreatedDateTime = nil
	o.isExplicitNullCreatedDateTime = b
}
// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphOrganization) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetMarketingNotificationEmails returns the MarketingNotificationEmails field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetMarketingNotificationEmails() []string {
	if o == nil || o.MarketingNotificationEmails == nil {
		var ret []string
		return ret
	}
	return *o.MarketingNotificationEmails
}

// GetMarketingNotificationEmailsOk returns a tuple with the MarketingNotificationEmails field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetMarketingNotificationEmailsOk() ([]string, bool) {
	if o == nil || o.MarketingNotificationEmails == nil {
		var ret []string
		return ret, false
	}
	return *o.MarketingNotificationEmails, true
}

// HasMarketingNotificationEmails returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasMarketingNotificationEmails() bool {
	if o != nil && o.MarketingNotificationEmails != nil {
		return true
	}

	return false
}

// SetMarketingNotificationEmails gets a reference to the given []string and assigns it to the MarketingNotificationEmails field.
func (o *MicrosoftGraphOrganization) SetMarketingNotificationEmails(v []string) {
	o.MarketingNotificationEmails = &v
}

// GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetOnPremisesLastSyncDateTime() time.Time {
	if o == nil || o.OnPremisesLastSyncDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.OnPremisesLastSyncDateTime
}

// GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetOnPremisesLastSyncDateTimeOk() (time.Time, bool) {
	if o == nil || o.OnPremisesLastSyncDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.OnPremisesLastSyncDateTime, true
}

// HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasOnPremisesLastSyncDateTime() bool {
	if o != nil && o.OnPremisesLastSyncDateTime != nil {
		return true
	}

	return false
}

// SetOnPremisesLastSyncDateTime gets a reference to the given time.Time and assigns it to the OnPremisesLastSyncDateTime field.
func (o *MicrosoftGraphOrganization) SetOnPremisesLastSyncDateTime(v time.Time) {
	o.OnPremisesLastSyncDateTime = &v
}

// SetOnPremisesLastSyncDateTimeExplicitNull (un)sets OnPremisesLastSyncDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OnPremisesLastSyncDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetOnPremisesLastSyncDateTimeExplicitNull(b bool) {
	o.OnPremisesLastSyncDateTime = nil
	o.isExplicitNullOnPremisesLastSyncDateTime = b
}
// GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetOnPremisesSyncEnabled() bool {
	if o == nil || o.OnPremisesSyncEnabled == nil {
		var ret bool
		return ret
	}
	return *o.OnPremisesSyncEnabled
}

// GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetOnPremisesSyncEnabledOk() (bool, bool) {
	if o == nil || o.OnPremisesSyncEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.OnPremisesSyncEnabled, true
}

// HasOnPremisesSyncEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasOnPremisesSyncEnabled() bool {
	if o != nil && o.OnPremisesSyncEnabled != nil {
		return true
	}

	return false
}

// SetOnPremisesSyncEnabled gets a reference to the given bool and assigns it to the OnPremisesSyncEnabled field.
func (o *MicrosoftGraphOrganization) SetOnPremisesSyncEnabled(v bool) {
	o.OnPremisesSyncEnabled = &v
}

// SetOnPremisesSyncEnabledExplicitNull (un)sets OnPremisesSyncEnabled to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OnPremisesSyncEnabled value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetOnPremisesSyncEnabledExplicitNull(b bool) {
	o.OnPremisesSyncEnabled = nil
	o.isExplicitNullOnPremisesSyncEnabled = b
}
// GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetPostalCodeOk() (string, bool) {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret, false
	}
	return *o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *MicrosoftGraphOrganization) SetPostalCode(v string) {
	o.PostalCode = &v
}

// SetPostalCodeExplicitNull (un)sets PostalCode to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PostalCode value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetPostalCodeExplicitNull(b bool) {
	o.PostalCode = nil
	o.isExplicitNullPostalCode = b
}
// GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetPreferredLanguage() string {
	if o == nil || o.PreferredLanguage == nil {
		var ret string
		return ret
	}
	return *o.PreferredLanguage
}

// GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetPreferredLanguageOk() (string, bool) {
	if o == nil || o.PreferredLanguage == nil {
		var ret string
		return ret, false
	}
	return *o.PreferredLanguage, true
}

// HasPreferredLanguage returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasPreferredLanguage() bool {
	if o != nil && o.PreferredLanguage != nil {
		return true
	}

	return false
}

// SetPreferredLanguage gets a reference to the given string and assigns it to the PreferredLanguage field.
func (o *MicrosoftGraphOrganization) SetPreferredLanguage(v string) {
	o.PreferredLanguage = &v
}

// SetPreferredLanguageExplicitNull (un)sets PreferredLanguage to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PreferredLanguage value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetPreferredLanguageExplicitNull(b bool) {
	o.PreferredLanguage = nil
	o.isExplicitNullPreferredLanguage = b
}
// GetPrivacyProfile returns the PrivacyProfile field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetPrivacyProfile() AnyOfmicrosoftGraphPrivacyProfile {
	if o == nil || o.PrivacyProfile == nil {
		var ret AnyOfmicrosoftGraphPrivacyProfile
		return ret
	}
	return *o.PrivacyProfile
}

// GetPrivacyProfileOk returns a tuple with the PrivacyProfile field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetPrivacyProfileOk() (AnyOfmicrosoftGraphPrivacyProfile, bool) {
	if o == nil || o.PrivacyProfile == nil {
		var ret AnyOfmicrosoftGraphPrivacyProfile
		return ret, false
	}
	return *o.PrivacyProfile, true
}

// HasPrivacyProfile returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasPrivacyProfile() bool {
	if o != nil && o.PrivacyProfile != nil {
		return true
	}

	return false
}

// SetPrivacyProfile gets a reference to the given AnyOfmicrosoftGraphPrivacyProfile and assigns it to the PrivacyProfile field.
func (o *MicrosoftGraphOrganization) SetPrivacyProfile(v AnyOfmicrosoftGraphPrivacyProfile) {
	o.PrivacyProfile = &v
}

// SetPrivacyProfileExplicitNull (un)sets PrivacyProfile to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PrivacyProfile value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetPrivacyProfileExplicitNull(b bool) {
	o.PrivacyProfile = nil
	o.isExplicitNullPrivacyProfile = b
}
// GetProvisionedPlans returns the ProvisionedPlans field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetProvisionedPlans() []MicrosoftGraphProvisionedPlan {
	if o == nil || o.ProvisionedPlans == nil {
		var ret []MicrosoftGraphProvisionedPlan
		return ret
	}
	return *o.ProvisionedPlans
}

// GetProvisionedPlansOk returns a tuple with the ProvisionedPlans field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetProvisionedPlansOk() ([]MicrosoftGraphProvisionedPlan, bool) {
	if o == nil || o.ProvisionedPlans == nil {
		var ret []MicrosoftGraphProvisionedPlan
		return ret, false
	}
	return *o.ProvisionedPlans, true
}

// HasProvisionedPlans returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasProvisionedPlans() bool {
	if o != nil && o.ProvisionedPlans != nil {
		return true
	}

	return false
}

// SetProvisionedPlans gets a reference to the given []MicrosoftGraphProvisionedPlan and assigns it to the ProvisionedPlans field.
func (o *MicrosoftGraphOrganization) SetProvisionedPlans(v []MicrosoftGraphProvisionedPlan) {
	o.ProvisionedPlans = &v
}

// GetSecurityComplianceNotificationMails returns the SecurityComplianceNotificationMails field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetSecurityComplianceNotificationMails() []string {
	if o == nil || o.SecurityComplianceNotificationMails == nil {
		var ret []string
		return ret
	}
	return *o.SecurityComplianceNotificationMails
}

// GetSecurityComplianceNotificationMailsOk returns a tuple with the SecurityComplianceNotificationMails field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetSecurityComplianceNotificationMailsOk() ([]string, bool) {
	if o == nil || o.SecurityComplianceNotificationMails == nil {
		var ret []string
		return ret, false
	}
	return *o.SecurityComplianceNotificationMails, true
}

// HasSecurityComplianceNotificationMails returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasSecurityComplianceNotificationMails() bool {
	if o != nil && o.SecurityComplianceNotificationMails != nil {
		return true
	}

	return false
}

// SetSecurityComplianceNotificationMails gets a reference to the given []string and assigns it to the SecurityComplianceNotificationMails field.
func (o *MicrosoftGraphOrganization) SetSecurityComplianceNotificationMails(v []string) {
	o.SecurityComplianceNotificationMails = &v
}

// GetSecurityComplianceNotificationPhones returns the SecurityComplianceNotificationPhones field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetSecurityComplianceNotificationPhones() []string {
	if o == nil || o.SecurityComplianceNotificationPhones == nil {
		var ret []string
		return ret
	}
	return *o.SecurityComplianceNotificationPhones
}

// GetSecurityComplianceNotificationPhonesOk returns a tuple with the SecurityComplianceNotificationPhones field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetSecurityComplianceNotificationPhonesOk() ([]string, bool) {
	if o == nil || o.SecurityComplianceNotificationPhones == nil {
		var ret []string
		return ret, false
	}
	return *o.SecurityComplianceNotificationPhones, true
}

// HasSecurityComplianceNotificationPhones returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasSecurityComplianceNotificationPhones() bool {
	if o != nil && o.SecurityComplianceNotificationPhones != nil {
		return true
	}

	return false
}

// SetSecurityComplianceNotificationPhones gets a reference to the given []string and assigns it to the SecurityComplianceNotificationPhones field.
func (o *MicrosoftGraphOrganization) SetSecurityComplianceNotificationPhones(v []string) {
	o.SecurityComplianceNotificationPhones = &v
}

// GetState returns the State field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MicrosoftGraphOrganization) SetState(v string) {
	o.State = &v
}

// SetStateExplicitNull (un)sets State to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The State value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetStateExplicitNull(b bool) {
	o.State = nil
	o.isExplicitNullState = b
}
// GetStreet returns the Street field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetStreet() string {
	if o == nil || o.Street == nil {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetStreetOk() (string, bool) {
	if o == nil || o.Street == nil {
		var ret string
		return ret, false
	}
	return *o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *MicrosoftGraphOrganization) SetStreet(v string) {
	o.Street = &v
}

// SetStreetExplicitNull (un)sets Street to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Street value is set to nil even if false is passed
func (o *MicrosoftGraphOrganization) SetStreetExplicitNull(b bool) {
	o.Street = nil
	o.isExplicitNullStreet = b
}
// GetTechnicalNotificationMails returns the TechnicalNotificationMails field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetTechnicalNotificationMails() []string {
	if o == nil || o.TechnicalNotificationMails == nil {
		var ret []string
		return ret
	}
	return *o.TechnicalNotificationMails
}

// GetTechnicalNotificationMailsOk returns a tuple with the TechnicalNotificationMails field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetTechnicalNotificationMailsOk() ([]string, bool) {
	if o == nil || o.TechnicalNotificationMails == nil {
		var ret []string
		return ret, false
	}
	return *o.TechnicalNotificationMails, true
}

// HasTechnicalNotificationMails returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasTechnicalNotificationMails() bool {
	if o != nil && o.TechnicalNotificationMails != nil {
		return true
	}

	return false
}

// SetTechnicalNotificationMails gets a reference to the given []string and assigns it to the TechnicalNotificationMails field.
func (o *MicrosoftGraphOrganization) SetTechnicalNotificationMails(v []string) {
	o.TechnicalNotificationMails = &v
}

// GetVerifiedDomains returns the VerifiedDomains field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetVerifiedDomains() []MicrosoftGraphVerifiedDomain {
	if o == nil || o.VerifiedDomains == nil {
		var ret []MicrosoftGraphVerifiedDomain
		return ret
	}
	return *o.VerifiedDomains
}

// GetVerifiedDomainsOk returns a tuple with the VerifiedDomains field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetVerifiedDomainsOk() ([]MicrosoftGraphVerifiedDomain, bool) {
	if o == nil || o.VerifiedDomains == nil {
		var ret []MicrosoftGraphVerifiedDomain
		return ret, false
	}
	return *o.VerifiedDomains, true
}

// HasVerifiedDomains returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasVerifiedDomains() bool {
	if o != nil && o.VerifiedDomains != nil {
		return true
	}

	return false
}

// SetVerifiedDomains gets a reference to the given []MicrosoftGraphVerifiedDomain and assigns it to the VerifiedDomains field.
func (o *MicrosoftGraphOrganization) SetVerifiedDomains(v []MicrosoftGraphVerifiedDomain) {
	o.VerifiedDomains = &v
}

// GetMobileDeviceManagementAuthority returns the MobileDeviceManagementAuthority field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetMobileDeviceManagementAuthority() AnyOfmicrosoftGraphMdmAuthority {
	if o == nil || o.MobileDeviceManagementAuthority == nil {
		var ret AnyOfmicrosoftGraphMdmAuthority
		return ret
	}
	return *o.MobileDeviceManagementAuthority
}

// GetMobileDeviceManagementAuthorityOk returns a tuple with the MobileDeviceManagementAuthority field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetMobileDeviceManagementAuthorityOk() (AnyOfmicrosoftGraphMdmAuthority, bool) {
	if o == nil || o.MobileDeviceManagementAuthority == nil {
		var ret AnyOfmicrosoftGraphMdmAuthority
		return ret, false
	}
	return *o.MobileDeviceManagementAuthority, true
}

// HasMobileDeviceManagementAuthority returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasMobileDeviceManagementAuthority() bool {
	if o != nil && o.MobileDeviceManagementAuthority != nil {
		return true
	}

	return false
}

// SetMobileDeviceManagementAuthority gets a reference to the given AnyOfmicrosoftGraphMdmAuthority and assigns it to the MobileDeviceManagementAuthority field.
func (o *MicrosoftGraphOrganization) SetMobileDeviceManagementAuthority(v AnyOfmicrosoftGraphMdmAuthority) {
	o.MobileDeviceManagementAuthority = &v
}

// GetExtensions returns the Extensions field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOrganization) GetExtensions() []MicrosoftGraphExtension {
	if o == nil || o.Extensions == nil {
		var ret []MicrosoftGraphExtension
		return ret
	}
	return *o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOrganization) GetExtensionsOk() ([]MicrosoftGraphExtension, bool) {
	if o == nil || o.Extensions == nil {
		var ret []MicrosoftGraphExtension
		return ret, false
	}
	return *o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *MicrosoftGraphOrganization) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.
func (o *MicrosoftGraphOrganization) SetExtensions(v []MicrosoftGraphExtension) {
	o.Extensions = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphOrganization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DeletedDateTime == nil {
		if o.isExplicitNullDeletedDateTime {
			toSerialize["deletedDateTime"] = o.DeletedDateTime
		}
	} else {
		toSerialize["deletedDateTime"] = o.DeletedDateTime
	}
	if o.AssignedPlans != nil {
		toSerialize["assignedPlans"] = o.AssignedPlans
	}
	if o.BusinessPhones != nil {
		toSerialize["businessPhones"] = o.BusinessPhones
	}
	if o.City == nil {
		if o.isExplicitNullCity {
			toSerialize["city"] = o.City
		}
	} else {
		toSerialize["city"] = o.City
	}
	if o.Country == nil {
		if o.isExplicitNullCountry {
			toSerialize["country"] = o.Country
		}
	} else {
		toSerialize["country"] = o.Country
	}
	if o.CountryLetterCode == nil {
		if o.isExplicitNullCountryLetterCode {
			toSerialize["countryLetterCode"] = o.CountryLetterCode
		}
	} else {
		toSerialize["countryLetterCode"] = o.CountryLetterCode
	}
	if o.CreatedDateTime == nil {
		if o.isExplicitNullCreatedDateTime {
			toSerialize["createdDateTime"] = o.CreatedDateTime
		}
	} else {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.MarketingNotificationEmails != nil {
		toSerialize["marketingNotificationEmails"] = o.MarketingNotificationEmails
	}
	if o.OnPremisesLastSyncDateTime == nil {
		if o.isExplicitNullOnPremisesLastSyncDateTime {
			toSerialize["onPremisesLastSyncDateTime"] = o.OnPremisesLastSyncDateTime
		}
	} else {
		toSerialize["onPremisesLastSyncDateTime"] = o.OnPremisesLastSyncDateTime
	}
	if o.OnPremisesSyncEnabled == nil {
		if o.isExplicitNullOnPremisesSyncEnabled {
			toSerialize["onPremisesSyncEnabled"] = o.OnPremisesSyncEnabled
		}
	} else {
		toSerialize["onPremisesSyncEnabled"] = o.OnPremisesSyncEnabled
	}
	if o.PostalCode == nil {
		if o.isExplicitNullPostalCode {
			toSerialize["postalCode"] = o.PostalCode
		}
	} else {
		toSerialize["postalCode"] = o.PostalCode
	}
	if o.PreferredLanguage == nil {
		if o.isExplicitNullPreferredLanguage {
			toSerialize["preferredLanguage"] = o.PreferredLanguage
		}
	} else {
		toSerialize["preferredLanguage"] = o.PreferredLanguage
	}
	if o.PrivacyProfile == nil {
		if o.isExplicitNullPrivacyProfile {
			toSerialize["privacyProfile"] = o.PrivacyProfile
		}
	} else {
		toSerialize["privacyProfile"] = o.PrivacyProfile
	}
	if o.ProvisionedPlans != nil {
		toSerialize["provisionedPlans"] = o.ProvisionedPlans
	}
	if o.SecurityComplianceNotificationMails != nil {
		toSerialize["securityComplianceNotificationMails"] = o.SecurityComplianceNotificationMails
	}
	if o.SecurityComplianceNotificationPhones != nil {
		toSerialize["securityComplianceNotificationPhones"] = o.SecurityComplianceNotificationPhones
	}
	if o.State == nil {
		if o.isExplicitNullState {
			toSerialize["state"] = o.State
		}
	} else {
		toSerialize["state"] = o.State
	}
	if o.Street == nil {
		if o.isExplicitNullStreet {
			toSerialize["street"] = o.Street
		}
	} else {
		toSerialize["street"] = o.Street
	}
	if o.TechnicalNotificationMails != nil {
		toSerialize["technicalNotificationMails"] = o.TechnicalNotificationMails
	}
	if o.VerifiedDomains != nil {
		toSerialize["verifiedDomains"] = o.VerifiedDomains
	}
	if o.MobileDeviceManagementAuthority != nil {
		toSerialize["mobileDeviceManagementAuthority"] = o.MobileDeviceManagementAuthority
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	return json.Marshal(toSerialize)
}

