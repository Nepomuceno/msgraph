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
// MicrosoftGraphKeyValue struct for MicrosoftGraphKeyValue
type MicrosoftGraphKeyValue struct {
	Key *string `json:"key,omitempty"`
	isExplicitNullKey bool `json:"-"`
	Value *string `json:"value,omitempty"`
	isExplicitNullValue bool `json:"-"`
}

// GetKey returns the Key field if non-nil, zero value otherwise.
func (o *MicrosoftGraphKeyValue) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphKeyValue) GetKeyOk() (string, bool) {
	if o == nil || o.Key == nil {
		var ret string
		return ret, false
	}
	return *o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *MicrosoftGraphKeyValue) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *MicrosoftGraphKeyValue) SetKey(v string) {
	o.Key = &v
}

// SetKeyExplicitNull (un)sets Key to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Key value is set to nil even if false is passed
func (o *MicrosoftGraphKeyValue) SetKeyExplicitNull(b bool) {
	o.Key = nil
	o.isExplicitNullKey = b
}
// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *MicrosoftGraphKeyValue) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphKeyValue) GetValueOk() (string, bool) {
	if o == nil || o.Value == nil {
		var ret string
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MicrosoftGraphKeyValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *MicrosoftGraphKeyValue) SetValue(v string) {
	o.Value = &v
}

// SetValueExplicitNull (un)sets Value to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Value value is set to nil even if false is passed
func (o *MicrosoftGraphKeyValue) SetValueExplicitNull(b bool) {
	o.Value = nil
	o.isExplicitNullValue = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphKeyValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key == nil {
		if o.isExplicitNullKey {
			toSerialize["key"] = o.Key
		}
	} else {
		toSerialize["key"] = o.Key
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


