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
// MicrosoftGraphVerifiedDomain struct for MicrosoftGraphVerifiedDomain
type MicrosoftGraphVerifiedDomain struct {
	Capabilities *string `json:"capabilities,omitempty"`
	isExplicitNullCapabilities bool `json:"-"`
	IsDefault *bool `json:"isDefault,omitempty"`
	isExplicitNullIsDefault bool `json:"-"`
	IsInitial *bool `json:"isInitial,omitempty"`
	isExplicitNullIsInitial bool `json:"-"`
	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
	Type *string `json:"type,omitempty"`
	isExplicitNullType bool `json:"-"`
}

// GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVerifiedDomain) GetCapabilities() string {
	if o == nil || o.Capabilities == nil {
		var ret string
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVerifiedDomain) GetCapabilitiesOk() (string, bool) {
	if o == nil || o.Capabilities == nil {
		var ret string
		return ret, false
	}
	return *o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *MicrosoftGraphVerifiedDomain) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given string and assigns it to the Capabilities field.
func (o *MicrosoftGraphVerifiedDomain) SetCapabilities(v string) {
	o.Capabilities = &v
}

// SetCapabilitiesExplicitNull (un)sets Capabilities to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Capabilities value is set to nil even if false is passed
func (o *MicrosoftGraphVerifiedDomain) SetCapabilitiesExplicitNull(b bool) {
	o.Capabilities = nil
	o.isExplicitNullCapabilities = b
}
// GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVerifiedDomain) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVerifiedDomain) GetIsDefaultOk() (bool, bool) {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret, false
	}
	return *o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *MicrosoftGraphVerifiedDomain) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *MicrosoftGraphVerifiedDomain) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// SetIsDefaultExplicitNull (un)sets IsDefault to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The IsDefault value is set to nil even if false is passed
func (o *MicrosoftGraphVerifiedDomain) SetIsDefaultExplicitNull(b bool) {
	o.IsDefault = nil
	o.isExplicitNullIsDefault = b
}
// GetIsInitial returns the IsInitial field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVerifiedDomain) GetIsInitial() bool {
	if o == nil || o.IsInitial == nil {
		var ret bool
		return ret
	}
	return *o.IsInitial
}

// GetIsInitialOk returns a tuple with the IsInitial field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVerifiedDomain) GetIsInitialOk() (bool, bool) {
	if o == nil || o.IsInitial == nil {
		var ret bool
		return ret, false
	}
	return *o.IsInitial, true
}

// HasIsInitial returns a boolean if a field has been set.
func (o *MicrosoftGraphVerifiedDomain) HasIsInitial() bool {
	if o != nil && o.IsInitial != nil {
		return true
	}

	return false
}

// SetIsInitial gets a reference to the given bool and assigns it to the IsInitial field.
func (o *MicrosoftGraphVerifiedDomain) SetIsInitial(v bool) {
	o.IsInitial = &v
}

// SetIsInitialExplicitNull (un)sets IsInitial to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The IsInitial value is set to nil even if false is passed
func (o *MicrosoftGraphVerifiedDomain) SetIsInitialExplicitNull(b bool) {
	o.IsInitial = nil
	o.isExplicitNullIsInitial = b
}
// GetName returns the Name field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVerifiedDomain) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVerifiedDomain) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphVerifiedDomain) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftGraphVerifiedDomain) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *MicrosoftGraphVerifiedDomain) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetType returns the Type field if non-nil, zero value otherwise.
func (o *MicrosoftGraphVerifiedDomain) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVerifiedDomain) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphVerifiedDomain) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MicrosoftGraphVerifiedDomain) SetType(v string) {
	o.Type = &v
}

// SetTypeExplicitNull (un)sets Type to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Type value is set to nil even if false is passed
func (o *MicrosoftGraphVerifiedDomain) SetTypeExplicitNull(b bool) {
	o.Type = nil
	o.isExplicitNullType = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphVerifiedDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Capabilities == nil {
		if o.isExplicitNullCapabilities {
			toSerialize["capabilities"] = o.Capabilities
		}
	} else {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.IsDefault == nil {
		if o.isExplicitNullIsDefault {
			toSerialize["isDefault"] = o.IsDefault
		}
	} else {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.IsInitial == nil {
		if o.isExplicitNullIsInitial {
			toSerialize["isInitial"] = o.IsInitial
		}
	} else {
		toSerialize["isInitial"] = o.IsInitial
	}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
	}
	if o.Type == nil {
		if o.isExplicitNullType {
			toSerialize["type"] = o.Type
		}
	} else {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}


