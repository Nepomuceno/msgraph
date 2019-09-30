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
// MicrosoftGraphCurrencyColumn struct for MicrosoftGraphCurrencyColumn
type MicrosoftGraphCurrencyColumn struct {
	Locale *string `json:"locale,omitempty"`
	isExplicitNullLocale bool `json:"-"`
}

// GetLocale returns the Locale field if non-nil, zero value otherwise.
func (o *MicrosoftGraphCurrencyColumn) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphCurrencyColumn) GetLocaleOk() (string, bool) {
	if o == nil || o.Locale == nil {
		var ret string
		return ret, false
	}
	return *o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *MicrosoftGraphCurrencyColumn) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *MicrosoftGraphCurrencyColumn) SetLocale(v string) {
	o.Locale = &v
}

// SetLocaleExplicitNull (un)sets Locale to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Locale value is set to nil even if false is passed
func (o *MicrosoftGraphCurrencyColumn) SetLocaleExplicitNull(b bool) {
	o.Locale = nil
	o.isExplicitNullLocale = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphCurrencyColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Locale == nil {
		if o.isExplicitNullLocale {
			toSerialize["locale"] = o.Locale
		}
	} else {
		toSerialize["locale"] = o.Locale
	}
	return json.Marshal(toSerialize)
}


