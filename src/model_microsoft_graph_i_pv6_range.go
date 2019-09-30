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
// MicrosoftGraphIPv6Range struct for MicrosoftGraphIPv6Range
type MicrosoftGraphIPv6Range struct {
	// Lower address
	LowerAddress *string `json:"lowerAddress,omitempty"`

	// Upper address
	UpperAddress *string `json:"upperAddress,omitempty"`

}

// GetLowerAddress returns the LowerAddress field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIPv6Range) GetLowerAddress() string {
	if o == nil || o.LowerAddress == nil {
		var ret string
		return ret
	}
	return *o.LowerAddress
}

// GetLowerAddressOk returns a tuple with the LowerAddress field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIPv6Range) GetLowerAddressOk() (string, bool) {
	if o == nil || o.LowerAddress == nil {
		var ret string
		return ret, false
	}
	return *o.LowerAddress, true
}

// HasLowerAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphIPv6Range) HasLowerAddress() bool {
	if o != nil && o.LowerAddress != nil {
		return true
	}

	return false
}

// SetLowerAddress gets a reference to the given string and assigns it to the LowerAddress field.
func (o *MicrosoftGraphIPv6Range) SetLowerAddress(v string) {
	o.LowerAddress = &v
}

// GetUpperAddress returns the UpperAddress field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIPv6Range) GetUpperAddress() string {
	if o == nil || o.UpperAddress == nil {
		var ret string
		return ret
	}
	return *o.UpperAddress
}

// GetUpperAddressOk returns a tuple with the UpperAddress field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIPv6Range) GetUpperAddressOk() (string, bool) {
	if o == nil || o.UpperAddress == nil {
		var ret string
		return ret, false
	}
	return *o.UpperAddress, true
}

// HasUpperAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphIPv6Range) HasUpperAddress() bool {
	if o != nil && o.UpperAddress != nil {
		return true
	}

	return false
}

// SetUpperAddress gets a reference to the given string and assigns it to the UpperAddress field.
func (o *MicrosoftGraphIPv6Range) SetUpperAddress(v string) {
	o.UpperAddress = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphIPv6Range) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LowerAddress != nil {
		toSerialize["lowerAddress"] = o.LowerAddress
	}
	if o.UpperAddress != nil {
		toSerialize["upperAddress"] = o.UpperAddress
	}
	return json.Marshal(toSerialize)
}

