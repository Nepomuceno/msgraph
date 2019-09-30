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
// AndroidWorkProfileCustomConfiguration Android Work Profile custom configuration
type AndroidWorkProfileCustomConfiguration struct {
	// OMA settings. This collection can contain a maximum of 500 elements.
	OmaSettings *[]AnyOfmicrosoftGraphOmaSetting `json:"omaSettings,omitempty"`

}

// GetOmaSettings returns the OmaSettings field if non-nil, zero value otherwise.
func (o *AndroidWorkProfileCustomConfiguration) GetOmaSettings() []AnyOfmicrosoftGraphOmaSetting {
	if o == nil || o.OmaSettings == nil {
		var ret []AnyOfmicrosoftGraphOmaSetting
		return ret
	}
	return *o.OmaSettings
}

// GetOmaSettingsOk returns a tuple with the OmaSettings field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *AndroidWorkProfileCustomConfiguration) GetOmaSettingsOk() ([]AnyOfmicrosoftGraphOmaSetting, bool) {
	if o == nil || o.OmaSettings == nil {
		var ret []AnyOfmicrosoftGraphOmaSetting
		return ret, false
	}
	return *o.OmaSettings, true
}

// HasOmaSettings returns a boolean if a field has been set.
func (o *AndroidWorkProfileCustomConfiguration) HasOmaSettings() bool {
	if o != nil && o.OmaSettings != nil {
		return true
	}

	return false
}

// SetOmaSettings gets a reference to the given []AnyOfmicrosoftGraphOmaSetting and assigns it to the OmaSettings field.
func (o *AndroidWorkProfileCustomConfiguration) SetOmaSettings(v []AnyOfmicrosoftGraphOmaSetting) {
	o.OmaSettings = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o AndroidWorkProfileCustomConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OmaSettings != nil {
		toSerialize["omaSettings"] = o.OmaSettings
	}
	return json.Marshal(toSerialize)
}

