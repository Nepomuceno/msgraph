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
// MicrosoftGraphVppToken struct for MicrosoftGraphVppToken
type MicrosoftGraphVppToken struct {
	Id *string `json:"id,omitempty"`

	// The organization associated with the Apple Volume Purchase Program Token
	OrganizationName *string `json:"organizationName,omitempty"`
	isExplicitNullOrganizationName bool `json:"-"`
	// The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: `business`, `education`.
	VppTokenAccountType *AnyOfmicrosoftGraphVppTokenAccountType `json:"vppTokenAccountType,omitempty"`

	// The apple Id associated with the given Apple Volume Purchase Program Token.
	AppleId *string `json:"appleId,omitempty"`
	isExplicitNullAppleId bool `json:"-"`
	// The expiration date time of the Apple Volume Purchase Program Token.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`

	// The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`

	// The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
	Token *string `json:"token,omitempty"`
	isExplicitNullToken bool `json:"-"`
	// Last modification date time associated with the Apple Volume Purchase Program Token.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`

	// Current state of the Apple Volume Purchase Program Token. Possible values are: `unknown`, `valid`, `expired`, `invalid`, `assignedToExternalMDM`.
	State *AnyOfmicrosoftGraphVppTokenState `json:"state,omitempty"`

	// Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: `none`, `inProgress`, `completed`, `failed`.
	LastSyncStatus *AnyOfmicrosoftGraphVppTokenSyncStatus `json:"lastSyncStatus,omitempty"`

	// Whether or not apps for the VPP token will be automatically updated.
	AutomaticallyUpdateApps *bool `json:"automaticallyUpdateApps,omitempty"`

	// Whether or not apps for the VPP token will be automatically updated.
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	isExplicitNullCountryOrRegion bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphVppToken) SetId(v string) {
	o.Id = &v
}

// GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetOrganizationName() string {
	if o == nil || o.OrganizationName == nil {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetOrganizationNameOk() (string, bool) {
	if o == nil || o.OrganizationName == nil {
		var ret string
		return ret, false
	}
	return *o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasOrganizationName() bool {
	if o != nil && o.OrganizationName != nil {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *MicrosoftGraphVppToken) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// SetOrganizationNameExplicitNull (un)sets OrganizationName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OrganizationName value is set to nil even if false is passed
func (o *MicrosoftGraphVppToken) SetOrganizationNameExplicitNull(b bool) {
	o.OrganizationName = nil
	o.isExplicitNullOrganizationName = b
}
// GetVppTokenAccountType returns the VppTokenAccountType field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetVppTokenAccountType() AnyOfmicrosoftGraphVppTokenAccountType {
	if o == nil || o.VppTokenAccountType == nil {
		var ret AnyOfmicrosoftGraphVppTokenAccountType
		return ret
	}
	return *o.VppTokenAccountType
}

// GetVppTokenAccountTypeOk returns a tuple with the VppTokenAccountType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetVppTokenAccountTypeOk() (AnyOfmicrosoftGraphVppTokenAccountType, bool) {
	if o == nil || o.VppTokenAccountType == nil {
		var ret AnyOfmicrosoftGraphVppTokenAccountType
		return ret, false
	}
	return *o.VppTokenAccountType, true
}

// HasVppTokenAccountType returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasVppTokenAccountType() bool {
	if o != nil && o.VppTokenAccountType != nil {
		return true
	}

	return false
}

// SetVppTokenAccountType gets a reference to the given AnyOfmicrosoftGraphVppTokenAccountType and assigns it to the VppTokenAccountType field.
func (o *MicrosoftGraphVppToken) SetVppTokenAccountType(v AnyOfmicrosoftGraphVppTokenAccountType) {
	o.VppTokenAccountType = &v
}

// GetAppleId returns the AppleId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetAppleId() string {
	if o == nil || o.AppleId == nil {
		var ret string
		return ret
	}
	return *o.AppleId
}

// GetAppleIdOk returns a tuple with the AppleId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetAppleIdOk() (string, bool) {
	if o == nil || o.AppleId == nil {
		var ret string
		return ret, false
	}
	return *o.AppleId, true
}

// HasAppleId returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasAppleId() bool {
	if o != nil && o.AppleId != nil {
		return true
	}

	return false
}

// SetAppleId gets a reference to the given string and assigns it to the AppleId field.
func (o *MicrosoftGraphVppToken) SetAppleId(v string) {
	o.AppleId = &v
}

// SetAppleIdExplicitNull (un)sets AppleId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AppleId value is set to nil even if false is passed
func (o *MicrosoftGraphVppToken) SetAppleIdExplicitNull(b bool) {
	o.AppleId = nil
	o.isExplicitNullAppleId = b
}
// GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetExpirationDateTime() time.Time {
	if o == nil || o.ExpirationDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetExpirationDateTimeOk() (time.Time, bool) {
	if o == nil || o.ExpirationDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpirationDateTime, true
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasExpirationDateTime() bool {
	if o != nil && o.ExpirationDateTime != nil {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.
func (o *MicrosoftGraphVppToken) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime = &v
}

// GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetLastSyncDateTime() time.Time {
	if o == nil || o.LastSyncDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSyncDateTime
}

// GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetLastSyncDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastSyncDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastSyncDateTime, true
}

// HasLastSyncDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasLastSyncDateTime() bool {
	if o != nil && o.LastSyncDateTime != nil {
		return true
	}

	return false
}

// SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.
func (o *MicrosoftGraphVppToken) SetLastSyncDateTime(v time.Time) {
	o.LastSyncDateTime = &v
}

// GetToken returns the Token field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetTokenOk() (string, bool) {
	if o == nil || o.Token == nil {
		var ret string
		return ret, false
	}
	return *o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *MicrosoftGraphVppToken) SetToken(v string) {
	o.Token = &v
}

// SetTokenExplicitNull (un)sets Token to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Token value is set to nil even if false is passed
func (o *MicrosoftGraphVppToken) SetTokenExplicitNull(b bool) {
	o.Token = nil
	o.isExplicitNullToken = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphVppToken) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetState returns the State field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetState() AnyOfmicrosoftGraphVppTokenState {
	if o == nil || o.State == nil {
		var ret AnyOfmicrosoftGraphVppTokenState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetStateOk() (AnyOfmicrosoftGraphVppTokenState, bool) {
	if o == nil || o.State == nil {
		var ret AnyOfmicrosoftGraphVppTokenState
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphVppTokenState and assigns it to the State field.
func (o *MicrosoftGraphVppToken) SetState(v AnyOfmicrosoftGraphVppTokenState) {
	o.State = &v
}

// GetLastSyncStatus returns the LastSyncStatus field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetLastSyncStatus() AnyOfmicrosoftGraphVppTokenSyncStatus {
	if o == nil || o.LastSyncStatus == nil {
		var ret AnyOfmicrosoftGraphVppTokenSyncStatus
		return ret
	}
	return *o.LastSyncStatus
}

// GetLastSyncStatusOk returns a tuple with the LastSyncStatus field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetLastSyncStatusOk() (AnyOfmicrosoftGraphVppTokenSyncStatus, bool) {
	if o == nil || o.LastSyncStatus == nil {
		var ret AnyOfmicrosoftGraphVppTokenSyncStatus
		return ret, false
	}
	return *o.LastSyncStatus, true
}

// HasLastSyncStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasLastSyncStatus() bool {
	if o != nil && o.LastSyncStatus != nil {
		return true
	}

	return false
}

// SetLastSyncStatus gets a reference to the given AnyOfmicrosoftGraphVppTokenSyncStatus and assigns it to the LastSyncStatus field.
func (o *MicrosoftGraphVppToken) SetLastSyncStatus(v AnyOfmicrosoftGraphVppTokenSyncStatus) {
	o.LastSyncStatus = &v
}

// GetAutomaticallyUpdateApps returns the AutomaticallyUpdateApps field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetAutomaticallyUpdateApps() bool {
	if o == nil || o.AutomaticallyUpdateApps == nil {
		var ret bool
		return ret
	}
	return *o.AutomaticallyUpdateApps
}

// GetAutomaticallyUpdateAppsOk returns a tuple with the AutomaticallyUpdateApps field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetAutomaticallyUpdateAppsOk() (bool, bool) {
	if o == nil || o.AutomaticallyUpdateApps == nil {
		var ret bool
		return ret, false
	}
	return *o.AutomaticallyUpdateApps, true
}

// HasAutomaticallyUpdateApps returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasAutomaticallyUpdateApps() bool {
	if o != nil && o.AutomaticallyUpdateApps != nil {
		return true
	}

	return false
}

// SetAutomaticallyUpdateApps gets a reference to the given bool and assigns it to the AutomaticallyUpdateApps field.
func (o *MicrosoftGraphVppToken) SetAutomaticallyUpdateApps(v bool) {
	o.AutomaticallyUpdateApps = &v
}

// GetCountryOrRegion returns the CountryOrRegion field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetCountryOrRegion() string {
	if o == nil || o.CountryOrRegion == nil {
		var ret string
		return ret
	}
	return *o.CountryOrRegion
}

// GetCountryOrRegionOk returns a tuple with the CountryOrRegion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetCountryOrRegionOk() (string, bool) {
	if o == nil || o.CountryOrRegion == nil {
		var ret string
		return ret, false
	}
	return *o.CountryOrRegion, true
}

// HasCountryOrRegion returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasCountryOrRegion() bool {
	if o != nil && o.CountryOrRegion != nil {
		return true
	}

	return false
}

// SetCountryOrRegion gets a reference to the given string and assigns it to the CountryOrRegion field.
func (o *MicrosoftGraphVppToken) SetCountryOrRegion(v string) {
	o.CountryOrRegion = &v
}

// SetCountryOrRegionExplicitNull (un)sets CountryOrRegion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CountryOrRegion value is set to nil even if false is passed
func (o *MicrosoftGraphVppToken) SetCountryOrRegionExplicitNull(b bool) {
	o.CountryOrRegion = nil
	o.isExplicitNullCountryOrRegion = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphVppToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrganizationName == nil {
		if o.isExplicitNullOrganizationName {
			toSerialize["organizationName"] = o.OrganizationName
		}
	} else {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if o.VppTokenAccountType != nil {
		toSerialize["vppTokenAccountType"] = o.VppTokenAccountType
	}
	if o.AppleId == nil {
		if o.isExplicitNullAppleId {
			toSerialize["appleId"] = o.AppleId
		}
	} else {
		toSerialize["appleId"] = o.AppleId
	}
	if o.ExpirationDateTime != nil {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime
	}
	if o.LastSyncDateTime != nil {
		toSerialize["lastSyncDateTime"] = o.LastSyncDateTime
	}
	if o.Token == nil {
		if o.isExplicitNullToken {
			toSerialize["token"] = o.Token
		}
	} else {
		toSerialize["token"] = o.Token
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.LastSyncStatus != nil {
		toSerialize["lastSyncStatus"] = o.LastSyncStatus
	}
	if o.AutomaticallyUpdateApps != nil {
		toSerialize["automaticallyUpdateApps"] = o.AutomaticallyUpdateApps
	}
	if o.CountryOrRegion == nil {
		if o.isExplicitNullCountryOrRegion {
			toSerialize["countryOrRegion"] = o.CountryOrRegion
		}
	} else {
		toSerialize["countryOrRegion"] = o.CountryOrRegion
	}
	return json.Marshal(toSerialize)
}

