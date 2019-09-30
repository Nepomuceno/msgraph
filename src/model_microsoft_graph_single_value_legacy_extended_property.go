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
// MicrosoftGraphSingleValueLegacyExtendedProperty struct for MicrosoftGraphSingleValueLegacyExtendedProperty
type MicrosoftGraphSingleValueLegacyExtendedProperty struct {
	Id *string `json:"id,omitempty"`

	Value *string `json:"value,omitempty"`
	isExplicitNullValue bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) SetId(v string) {
	o.Id = &v
}

// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) GetValueOk() (string, bool) {
	if o == nil || o.Value == nil {
		var ret string
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) SetValue(v string) {
	o.Value = &v
}

// SetValueExplicitNull (un)sets Value to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Value value is set to nil even if false is passed
func (o *MicrosoftGraphSingleValueLegacyExtendedProperty) SetValueExplicitNull(b bool) {
	o.Value = nil
	o.isExplicitNullValue = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphSingleValueLegacyExtendedProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Value == nil {
		if o.isExplicitNullValue {
			toSerialize["value"] = o.Value
		}
	} else {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

