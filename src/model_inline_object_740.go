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
// InlineObject740 struct for InlineObject740
type InlineObject740 struct {
	Type *string `json:"type,omitempty"`

	Scope *string `json:"scope,omitempty"`
	isExplicitNullScope bool `json:"-"`
}

// GetType returns the Type field if non-nil, zero value otherwise.
func (o *InlineObject740) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject740) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject740) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject740) SetType(v string) {
	o.Type = &v
}

// GetScope returns the Scope field if non-nil, zero value otherwise.
func (o *InlineObject740) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject740) GetScopeOk() (string, bool) {
	if o == nil || o.Scope == nil {
		var ret string
		return ret, false
	}
	return *o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineObject740) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InlineObject740) SetScope(v string) {
	o.Scope = &v
}

// SetScopeExplicitNull (un)sets Scope to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Scope value is set to nil even if false is passed
func (o *InlineObject740) SetScopeExplicitNull(b bool) {
	o.Scope = nil
	o.isExplicitNullScope = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject740) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Scope == nil {
		if o.isExplicitNullScope {
			toSerialize["scope"] = o.Scope
		}
	} else {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

