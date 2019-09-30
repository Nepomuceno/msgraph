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
// MicrosoftGraphIosDeviceType struct for MicrosoftGraphIosDeviceType
type MicrosoftGraphIosDeviceType struct {
	// Whether the app should run on iPads.
	IPad *bool `json:"iPad,omitempty"`

	// Whether the app should run on iPhones and iPods.
	IPhoneAndIPod *bool `json:"iPhoneAndIPod,omitempty"`

}

// GetIPad returns the IPad field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosDeviceType) GetIPad() bool {
	if o == nil || o.IPad == nil {
		var ret bool
		return ret
	}
	return *o.IPad
}

// GetIPadOk returns a tuple with the IPad field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosDeviceType) GetIPadOk() (bool, bool) {
	if o == nil || o.IPad == nil {
		var ret bool
		return ret, false
	}
	return *o.IPad, true
}

// HasIPad returns a boolean if a field has been set.
func (o *MicrosoftGraphIosDeviceType) HasIPad() bool {
	if o != nil && o.IPad != nil {
		return true
	}

	return false
}

// SetIPad gets a reference to the given bool and assigns it to the IPad field.
func (o *MicrosoftGraphIosDeviceType) SetIPad(v bool) {
	o.IPad = &v
}

// GetIPhoneAndIPod returns the IPhoneAndIPod field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosDeviceType) GetIPhoneAndIPod() bool {
	if o == nil || o.IPhoneAndIPod == nil {
		var ret bool
		return ret
	}
	return *o.IPhoneAndIPod
}

// GetIPhoneAndIPodOk returns a tuple with the IPhoneAndIPod field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosDeviceType) GetIPhoneAndIPodOk() (bool, bool) {
	if o == nil || o.IPhoneAndIPod == nil {
		var ret bool
		return ret, false
	}
	return *o.IPhoneAndIPod, true
}

// HasIPhoneAndIPod returns a boolean if a field has been set.
func (o *MicrosoftGraphIosDeviceType) HasIPhoneAndIPod() bool {
	if o != nil && o.IPhoneAndIPod != nil {
		return true
	}

	return false
}

// SetIPhoneAndIPod gets a reference to the given bool and assigns it to the IPhoneAndIPod field.
func (o *MicrosoftGraphIosDeviceType) SetIPhoneAndIPod(v bool) {
	o.IPhoneAndIPod = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphIosDeviceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IPad != nil {
		toSerialize["iPad"] = o.IPad
	}
	if o.IPhoneAndIPod != nil {
		toSerialize["iPhoneAndIPod"] = o.IPhoneAndIPod
	}
	return json.Marshal(toSerialize)
}


