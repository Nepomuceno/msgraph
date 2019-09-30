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
// InlineObject1131 struct for InlineObject1131
type InlineObject1131 struct {
	Address *string `json:"address,omitempty"`
	isExplicitNullAddress bool `json:"-"`
	HasHeaders *bool `json:"hasHeaders,omitempty"`

}

// GetAddress returns the Address field if non-nil, zero value otherwise.
func (o *InlineObject1131) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1131) GetAddressOk() (string, bool) {
	if o == nil || o.Address == nil {
		var ret string
		return ret, false
	}
	return *o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineObject1131) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineObject1131) SetAddress(v string) {
	o.Address = &v
}

// SetAddressExplicitNull (un)sets Address to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Address value is set to nil even if false is passed
func (o *InlineObject1131) SetAddressExplicitNull(b bool) {
	o.Address = nil
	o.isExplicitNullAddress = b
}
// GetHasHeaders returns the HasHeaders field if non-nil, zero value otherwise.
func (o *InlineObject1131) GetHasHeaders() bool {
	if o == nil || o.HasHeaders == nil {
		var ret bool
		return ret
	}
	return *o.HasHeaders
}

// GetHasHeadersOk returns a tuple with the HasHeaders field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1131) GetHasHeadersOk() (bool, bool) {
	if o == nil || o.HasHeaders == nil {
		var ret bool
		return ret, false
	}
	return *o.HasHeaders, true
}

// HasHasHeaders returns a boolean if a field has been set.
func (o *InlineObject1131) HasHasHeaders() bool {
	if o != nil && o.HasHeaders != nil {
		return true
	}

	return false
}

// SetHasHeaders gets a reference to the given bool and assigns it to the HasHeaders field.
func (o *InlineObject1131) SetHasHeaders(v bool) {
	o.HasHeaders = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1131) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address == nil {
		if o.isExplicitNullAddress {
			toSerialize["address"] = o.Address
		}
	} else {
		toSerialize["address"] = o.Address
	}
	if o.HasHeaders != nil {
		toSerialize["hasHeaders"] = o.HasHeaders
	}
	return json.Marshal(toSerialize)
}


