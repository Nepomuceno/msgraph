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
// TelecomExpenseManagementPartner telecomExpenseManagementPartner resources represent the metadata and status of a given TEM service. Once your organization has onboarded with a partner, the partner can be enabled or disabled to switch TEM functionality on or off.
type TelecomExpenseManagementPartner struct {
	// Display name of the TEM partner.
	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	// URL of the TEM partner's administrative control panel, where an administrator can configure their TEM service.
	Url *string `json:"url,omitempty"`
	isExplicitNullUrl bool `json:"-"`
	// Whether the partner's AAD app has been authorized to access Intune.
	AppAuthorized *bool `json:"appAuthorized,omitempty"`

	// Whether Intune's connection to the TEM service is currently enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Timestamp of the last request sent to Intune by the TEM partner.
	LastConnectionDateTime *time.Time `json:"lastConnectionDateTime,omitempty"`

}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *TelecomExpenseManagementPartner) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TelecomExpenseManagementPartner) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TelecomExpenseManagementPartner) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TelecomExpenseManagementPartner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *TelecomExpenseManagementPartner) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetUrl returns the Url field if non-nil, zero value otherwise.
func (o *TelecomExpenseManagementPartner) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TelecomExpenseManagementPartner) GetUrlOk() (string, bool) {
	if o == nil || o.Url == nil {
		var ret string
		return ret, false
	}
	return *o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TelecomExpenseManagementPartner) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *TelecomExpenseManagementPartner) SetUrl(v string) {
	o.Url = &v
}

// SetUrlExplicitNull (un)sets Url to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Url value is set to nil even if false is passed
func (o *TelecomExpenseManagementPartner) SetUrlExplicitNull(b bool) {
	o.Url = nil
	o.isExplicitNullUrl = b
}
// GetAppAuthorized returns the AppAuthorized field if non-nil, zero value otherwise.
func (o *TelecomExpenseManagementPartner) GetAppAuthorized() bool {
	if o == nil || o.AppAuthorized == nil {
		var ret bool
		return ret
	}
	return *o.AppAuthorized
}

// GetAppAuthorizedOk returns a tuple with the AppAuthorized field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TelecomExpenseManagementPartner) GetAppAuthorizedOk() (bool, bool) {
	if o == nil || o.AppAuthorized == nil {
		var ret bool
		return ret, false
	}
	return *o.AppAuthorized, true
}

// HasAppAuthorized returns a boolean if a field has been set.
func (o *TelecomExpenseManagementPartner) HasAppAuthorized() bool {
	if o != nil && o.AppAuthorized != nil {
		return true
	}

	return false
}

// SetAppAuthorized gets a reference to the given bool and assigns it to the AppAuthorized field.
func (o *TelecomExpenseManagementPartner) SetAppAuthorized(v bool) {
	o.AppAuthorized = &v
}

// GetEnabled returns the Enabled field if non-nil, zero value otherwise.
func (o *TelecomExpenseManagementPartner) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TelecomExpenseManagementPartner) GetEnabledOk() (bool, bool) {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret, false
	}
	return *o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TelecomExpenseManagementPartner) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TelecomExpenseManagementPartner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLastConnectionDateTime returns the LastConnectionDateTime field if non-nil, zero value otherwise.
func (o *TelecomExpenseManagementPartner) GetLastConnectionDateTime() time.Time {
	if o == nil || o.LastConnectionDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastConnectionDateTime
}

// GetLastConnectionDateTimeOk returns a tuple with the LastConnectionDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TelecomExpenseManagementPartner) GetLastConnectionDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastConnectionDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastConnectionDateTime, true
}

// HasLastConnectionDateTime returns a boolean if a field has been set.
func (o *TelecomExpenseManagementPartner) HasLastConnectionDateTime() bool {
	if o != nil && o.LastConnectionDateTime != nil {
		return true
	}

	return false
}

// SetLastConnectionDateTime gets a reference to the given time.Time and assigns it to the LastConnectionDateTime field.
func (o *TelecomExpenseManagementPartner) SetLastConnectionDateTime(v time.Time) {
	o.LastConnectionDateTime = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o TelecomExpenseManagementPartner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Url == nil {
		if o.isExplicitNullUrl {
			toSerialize["url"] = o.Url
		}
	} else {
		toSerialize["url"] = o.Url
	}
	if o.AppAuthorized != nil {
		toSerialize["appAuthorized"] = o.AppAuthorized
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.LastConnectionDateTime != nil {
		toSerialize["lastConnectionDateTime"] = o.LastConnectionDateTime
	}
	return json.Marshal(toSerialize)
}

