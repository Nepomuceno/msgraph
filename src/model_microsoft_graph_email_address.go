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
// MicrosoftGraphEmailAddress struct for MicrosoftGraphEmailAddress
type MicrosoftGraphEmailAddress struct {
	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
	Address *string `json:"address,omitempty"`
	isExplicitNullAddress bool `json:"-"`
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEmailAddress) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEmailAddress) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphEmailAddress) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftGraphEmailAddress) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *MicrosoftGraphEmailAddress) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetAddress returns the Address field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEmailAddress) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEmailAddress) GetAddressOk() (string, bool) {
	if o == nil || o.Address == nil {
		var ret string
		return ret, false
	}
	return *o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphEmailAddress) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *MicrosoftGraphEmailAddress) SetAddress(v string) {
	o.Address = &v
}

// SetAddressExplicitNull (un)sets Address to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Address value is set to nil even if false is passed
func (o *MicrosoftGraphEmailAddress) SetAddressExplicitNull(b bool) {
	o.Address = nil
	o.isExplicitNullAddress = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphEmailAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
	}
	if o.Address == nil {
		if o.isExplicitNullAddress {
			toSerialize["address"] = o.Address
		}
	} else {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

