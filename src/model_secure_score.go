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
// SecureScore struct for SecureScore
type SecureScore struct {
	ActiveUserCount *int32 `json:"activeUserCount,omitempty"`
	isExplicitNullActiveUserCount bool `json:"-"`
	AverageComparativeScores *[]AnyOfmicrosoftGraphAverageComparativeScore `json:"averageComparativeScores,omitempty"`

	AzureTenantId *string `json:"azureTenantId,omitempty"`

	ControlScores *[]AnyOfmicrosoftGraphControlScore `json:"controlScores,omitempty"`

	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	isExplicitNullCreatedDateTime bool `json:"-"`
	CurrentScore *AnyOfnumberstringstring `json:"currentScore,omitempty"`
	isExplicitNullCurrentScore bool `json:"-"`
	EnabledServices *[]string `json:"enabledServices,omitempty"`

	LicensedUserCount *int32 `json:"licensedUserCount,omitempty"`
	isExplicitNullLicensedUserCount bool `json:"-"`
	MaxScore *AnyOfnumberstringstring `json:"maxScore,omitempty"`
	isExplicitNullMaxScore bool `json:"-"`
	VendorInformation *AnyOfmicrosoftGraphSecurityVendorInformation `json:"vendorInformation,omitempty"`
	isExplicitNullVendorInformation bool `json:"-"`
}

// GetActiveUserCount returns the ActiveUserCount field if non-nil, zero value otherwise.
func (o *SecureScore) GetActiveUserCount() int32 {
	if o == nil || o.ActiveUserCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveUserCount
}

// GetActiveUserCountOk returns a tuple with the ActiveUserCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetActiveUserCountOk() (int32, bool) {
	if o == nil || o.ActiveUserCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ActiveUserCount, true
}

// HasActiveUserCount returns a boolean if a field has been set.
func (o *SecureScore) HasActiveUserCount() bool {
	if o != nil && o.ActiveUserCount != nil {
		return true
	}

	return false
}

// SetActiveUserCount gets a reference to the given int32 and assigns it to the ActiveUserCount field.
func (o *SecureScore) SetActiveUserCount(v int32) {
	o.ActiveUserCount = &v
}

// SetActiveUserCountExplicitNull (un)sets ActiveUserCount to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ActiveUserCount value is set to nil even if false is passed
func (o *SecureScore) SetActiveUserCountExplicitNull(b bool) {
	o.ActiveUserCount = nil
	o.isExplicitNullActiveUserCount = b
}
// GetAverageComparativeScores returns the AverageComparativeScores field if non-nil, zero value otherwise.
func (o *SecureScore) GetAverageComparativeScores() []AnyOfmicrosoftGraphAverageComparativeScore {
	if o == nil || o.AverageComparativeScores == nil {
		var ret []AnyOfmicrosoftGraphAverageComparativeScore
		return ret
	}
	return *o.AverageComparativeScores
}

// GetAverageComparativeScoresOk returns a tuple with the AverageComparativeScores field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetAverageComparativeScoresOk() ([]AnyOfmicrosoftGraphAverageComparativeScore, bool) {
	if o == nil || o.AverageComparativeScores == nil {
		var ret []AnyOfmicrosoftGraphAverageComparativeScore
		return ret, false
	}
	return *o.AverageComparativeScores, true
}

// HasAverageComparativeScores returns a boolean if a field has been set.
func (o *SecureScore) HasAverageComparativeScores() bool {
	if o != nil && o.AverageComparativeScores != nil {
		return true
	}

	return false
}

// SetAverageComparativeScores gets a reference to the given []AnyOfmicrosoftGraphAverageComparativeScore and assigns it to the AverageComparativeScores field.
func (o *SecureScore) SetAverageComparativeScores(v []AnyOfmicrosoftGraphAverageComparativeScore) {
	o.AverageComparativeScores = &v
}

// GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.
func (o *SecureScore) GetAzureTenantId() string {
	if o == nil || o.AzureTenantId == nil {
		var ret string
		return ret
	}
	return *o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetAzureTenantIdOk() (string, bool) {
	if o == nil || o.AzureTenantId == nil {
		var ret string
		return ret, false
	}
	return *o.AzureTenantId, true
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *SecureScore) HasAzureTenantId() bool {
	if o != nil && o.AzureTenantId != nil {
		return true
	}

	return false
}

// SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.
func (o *SecureScore) SetAzureTenantId(v string) {
	o.AzureTenantId = &v
}

// GetControlScores returns the ControlScores field if non-nil, zero value otherwise.
func (o *SecureScore) GetControlScores() []AnyOfmicrosoftGraphControlScore {
	if o == nil || o.ControlScores == nil {
		var ret []AnyOfmicrosoftGraphControlScore
		return ret
	}
	return *o.ControlScores
}

// GetControlScoresOk returns a tuple with the ControlScores field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetControlScoresOk() ([]AnyOfmicrosoftGraphControlScore, bool) {
	if o == nil || o.ControlScores == nil {
		var ret []AnyOfmicrosoftGraphControlScore
		return ret, false
	}
	return *o.ControlScores, true
}

// HasControlScores returns a boolean if a field has been set.
func (o *SecureScore) HasControlScores() bool {
	if o != nil && o.ControlScores != nil {
		return true
	}

	return false
}

// SetControlScores gets a reference to the given []AnyOfmicrosoftGraphControlScore and assigns it to the ControlScores field.
func (o *SecureScore) SetControlScores(v []AnyOfmicrosoftGraphControlScore) {
	o.ControlScores = &v
}

// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *SecureScore) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *SecureScore) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *SecureScore) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedDateTime value is set to nil even if false is passed
func (o *SecureScore) SetCreatedDateTimeExplicitNull(b bool) {
	o.CreatedDateTime = nil
	o.isExplicitNullCreatedDateTime = b
}
// GetCurrentScore returns the CurrentScore field if non-nil, zero value otherwise.
func (o *SecureScore) GetCurrentScore() AnyOfnumberstringstring {
	if o == nil || o.CurrentScore == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.CurrentScore
}

// GetCurrentScoreOk returns a tuple with the CurrentScore field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetCurrentScoreOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.CurrentScore == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.CurrentScore, true
}

// HasCurrentScore returns a boolean if a field has been set.
func (o *SecureScore) HasCurrentScore() bool {
	if o != nil && o.CurrentScore != nil {
		return true
	}

	return false
}

// SetCurrentScore gets a reference to the given AnyOfnumberstringstring and assigns it to the CurrentScore field.
func (o *SecureScore) SetCurrentScore(v AnyOfnumberstringstring) {
	o.CurrentScore = &v
}

// SetCurrentScoreExplicitNull (un)sets CurrentScore to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CurrentScore value is set to nil even if false is passed
func (o *SecureScore) SetCurrentScoreExplicitNull(b bool) {
	o.CurrentScore = nil
	o.isExplicitNullCurrentScore = b
}
// GetEnabledServices returns the EnabledServices field if non-nil, zero value otherwise.
func (o *SecureScore) GetEnabledServices() []string {
	if o == nil || o.EnabledServices == nil {
		var ret []string
		return ret
	}
	return *o.EnabledServices
}

// GetEnabledServicesOk returns a tuple with the EnabledServices field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetEnabledServicesOk() ([]string, bool) {
	if o == nil || o.EnabledServices == nil {
		var ret []string
		return ret, false
	}
	return *o.EnabledServices, true
}

// HasEnabledServices returns a boolean if a field has been set.
func (o *SecureScore) HasEnabledServices() bool {
	if o != nil && o.EnabledServices != nil {
		return true
	}

	return false
}

// SetEnabledServices gets a reference to the given []string and assigns it to the EnabledServices field.
func (o *SecureScore) SetEnabledServices(v []string) {
	o.EnabledServices = &v
}

// GetLicensedUserCount returns the LicensedUserCount field if non-nil, zero value otherwise.
func (o *SecureScore) GetLicensedUserCount() int32 {
	if o == nil || o.LicensedUserCount == nil {
		var ret int32
		return ret
	}
	return *o.LicensedUserCount
}

// GetLicensedUserCountOk returns a tuple with the LicensedUserCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetLicensedUserCountOk() (int32, bool) {
	if o == nil || o.LicensedUserCount == nil {
		var ret int32
		return ret, false
	}
	return *o.LicensedUserCount, true
}

// HasLicensedUserCount returns a boolean if a field has been set.
func (o *SecureScore) HasLicensedUserCount() bool {
	if o != nil && o.LicensedUserCount != nil {
		return true
	}

	return false
}

// SetLicensedUserCount gets a reference to the given int32 and assigns it to the LicensedUserCount field.
func (o *SecureScore) SetLicensedUserCount(v int32) {
	o.LicensedUserCount = &v
}

// SetLicensedUserCountExplicitNull (un)sets LicensedUserCount to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LicensedUserCount value is set to nil even if false is passed
func (o *SecureScore) SetLicensedUserCountExplicitNull(b bool) {
	o.LicensedUserCount = nil
	o.isExplicitNullLicensedUserCount = b
}
// GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.
func (o *SecureScore) GetMaxScore() AnyOfnumberstringstring {
	if o == nil || o.MaxScore == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.MaxScore
}

// GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetMaxScoreOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.MaxScore == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.MaxScore, true
}

// HasMaxScore returns a boolean if a field has been set.
func (o *SecureScore) HasMaxScore() bool {
	if o != nil && o.MaxScore != nil {
		return true
	}

	return false
}

// SetMaxScore gets a reference to the given AnyOfnumberstringstring and assigns it to the MaxScore field.
func (o *SecureScore) SetMaxScore(v AnyOfnumberstringstring) {
	o.MaxScore = &v
}

// SetMaxScoreExplicitNull (un)sets MaxScore to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The MaxScore value is set to nil even if false is passed
func (o *SecureScore) SetMaxScoreExplicitNull(b bool) {
	o.MaxScore = nil
	o.isExplicitNullMaxScore = b
}
// GetVendorInformation returns the VendorInformation field if non-nil, zero value otherwise.
func (o *SecureScore) GetVendorInformation() AnyOfmicrosoftGraphSecurityVendorInformation {
	if o == nil || o.VendorInformation == nil {
		var ret AnyOfmicrosoftGraphSecurityVendorInformation
		return ret
	}
	return *o.VendorInformation
}

// GetVendorInformationOk returns a tuple with the VendorInformation field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecureScore) GetVendorInformationOk() (AnyOfmicrosoftGraphSecurityVendorInformation, bool) {
	if o == nil || o.VendorInformation == nil {
		var ret AnyOfmicrosoftGraphSecurityVendorInformation
		return ret, false
	}
	return *o.VendorInformation, true
}

// HasVendorInformation returns a boolean if a field has been set.
func (o *SecureScore) HasVendorInformation() bool {
	if o != nil && o.VendorInformation != nil {
		return true
	}

	return false
}

// SetVendorInformation gets a reference to the given AnyOfmicrosoftGraphSecurityVendorInformation and assigns it to the VendorInformation field.
func (o *SecureScore) SetVendorInformation(v AnyOfmicrosoftGraphSecurityVendorInformation) {
	o.VendorInformation = &v
}

// SetVendorInformationExplicitNull (un)sets VendorInformation to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The VendorInformation value is set to nil even if false is passed
func (o *SecureScore) SetVendorInformationExplicitNull(b bool) {
	o.VendorInformation = nil
	o.isExplicitNullVendorInformation = b
}

// MarshalJSON returns the JSON representation of the model.
func (o SecureScore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveUserCount == nil {
		if o.isExplicitNullActiveUserCount {
			toSerialize["activeUserCount"] = o.ActiveUserCount
		}
	} else {
		toSerialize["activeUserCount"] = o.ActiveUserCount
	}
	if o.AverageComparativeScores != nil {
		toSerialize["averageComparativeScores"] = o.AverageComparativeScores
	}
	if o.AzureTenantId != nil {
		toSerialize["azureTenantId"] = o.AzureTenantId
	}
	if o.ControlScores != nil {
		toSerialize["controlScores"] = o.ControlScores
	}
	if o.CreatedDateTime == nil {
		if o.isExplicitNullCreatedDateTime {
			toSerialize["createdDateTime"] = o.CreatedDateTime
		}
	} else {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.CurrentScore == nil {
		if o.isExplicitNullCurrentScore {
			toSerialize["currentScore"] = o.CurrentScore
		}
	} else {
		toSerialize["currentScore"] = o.CurrentScore
	}
	if o.EnabledServices != nil {
		toSerialize["enabledServices"] = o.EnabledServices
	}
	if o.LicensedUserCount == nil {
		if o.isExplicitNullLicensedUserCount {
			toSerialize["licensedUserCount"] = o.LicensedUserCount
		}
	} else {
		toSerialize["licensedUserCount"] = o.LicensedUserCount
	}
	if o.MaxScore == nil {
		if o.isExplicitNullMaxScore {
			toSerialize["maxScore"] = o.MaxScore
		}
	} else {
		toSerialize["maxScore"] = o.MaxScore
	}
	if o.VendorInformation == nil {
		if o.isExplicitNullVendorInformation {
			toSerialize["vendorInformation"] = o.VendorInformation
		}
	} else {
		toSerialize["vendorInformation"] = o.VendorInformation
	}
	return json.Marshal(toSerialize)
}

