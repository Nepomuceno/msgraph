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
	"encoding/json"
)
// DomainDnsMxRecord struct for DomainDnsMxRecord
type DomainDnsMxRecord struct {
	MailExchange *string `json:"mailExchange,omitempty"`

	Preference *int32 `json:"preference,omitempty"`
	isExplicitNullPreference bool `json:"-"`
}

// GetMailExchange returns the MailExchange field if non-nil, zero value otherwise.
func (o *DomainDnsMxRecord) GetMailExchange() string {
	if o == nil || o.MailExchange == nil {
		var ret string
		return ret
	}
	return *o.MailExchange
}

// GetMailExchangeOk returns a tuple with the MailExchange field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DomainDnsMxRecord) GetMailExchangeOk() (string, bool) {
	if o == nil || o.MailExchange == nil {
		var ret string
		return ret, false
	}
	return *o.MailExchange, true
}

// HasMailExchange returns a boolean if a field has been set.
func (o *DomainDnsMxRecord) HasMailExchange() bool {
	if o != nil && o.MailExchange != nil {
		return true
	}

	return false
}

// SetMailExchange gets a reference to the given string and assigns it to the MailExchange field.
func (o *DomainDnsMxRecord) SetMailExchange(v string) {
	o.MailExchange = &v
}

// GetPreference returns the Preference field if non-nil, zero value otherwise.
func (o *DomainDnsMxRecord) GetPreference() int32 {
	if o == nil || o.Preference == nil {
		var ret int32
		return ret
	}
	return *o.Preference
}

// GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DomainDnsMxRecord) GetPreferenceOk() (int32, bool) {
	if o == nil || o.Preference == nil {
		var ret int32
		return ret, false
	}
	return *o.Preference, true
}

// HasPreference returns a boolean if a field has been set.
func (o *DomainDnsMxRecord) HasPreference() bool {
	if o != nil && o.Preference != nil {
		return true
	}

	return false
}

// SetPreference gets a reference to the given int32 and assigns it to the Preference field.
func (o *DomainDnsMxRecord) SetPreference(v int32) {
	o.Preference = &v
}

// SetPreferenceExplicitNull (un)sets Preference to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Preference value is set to nil even if false is passed
func (o *DomainDnsMxRecord) SetPreferenceExplicitNull(b bool) {
	o.Preference = nil
	o.isExplicitNullPreference = b
}

// MarshalJSON returns the JSON representation of the model.
func (o DomainDnsMxRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MailExchange != nil {
		toSerialize["mailExchange"] = o.MailExchange
	}
	if o.Preference == nil {
		if o.isExplicitNullPreference {
			toSerialize["preference"] = o.Preference
		}
	} else {
		toSerialize["preference"] = o.Preference
	}
	return json.Marshal(toSerialize)
}


