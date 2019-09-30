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
// MicrosoftGraphRgbColor struct for MicrosoftGraphRgbColor
type MicrosoftGraphRgbColor struct {
	// Red value
	R *int32 `json:"r,omitempty"`

	// Green value
	G *int32 `json:"g,omitempty"`

	// Blue value
	B *int32 `json:"b,omitempty"`

}

// GetR returns the R field if non-nil, zero value otherwise.
func (o *MicrosoftGraphRgbColor) GetR() int32 {
	if o == nil || o.R == nil {
		var ret int32
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRgbColor) GetROk() (int32, bool) {
	if o == nil || o.R == nil {
		var ret int32
		return ret, false
	}
	return *o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *MicrosoftGraphRgbColor) HasR() bool {
	if o != nil && o.R != nil {
		return true
	}

	return false
}

// SetR gets a reference to the given int32 and assigns it to the R field.
func (o *MicrosoftGraphRgbColor) SetR(v int32) {
	o.R = &v
}

// GetG returns the G field if non-nil, zero value otherwise.
func (o *MicrosoftGraphRgbColor) GetG() int32 {
	if o == nil || o.G == nil {
		var ret int32
		return ret
	}
	return *o.G
}

// GetGOk returns a tuple with the G field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRgbColor) GetGOk() (int32, bool) {
	if o == nil || o.G == nil {
		var ret int32
		return ret, false
	}
	return *o.G, true
}

// HasG returns a boolean if a field has been set.
func (o *MicrosoftGraphRgbColor) HasG() bool {
	if o != nil && o.G != nil {
		return true
	}

	return false
}

// SetG gets a reference to the given int32 and assigns it to the G field.
func (o *MicrosoftGraphRgbColor) SetG(v int32) {
	o.G = &v
}

// GetB returns the B field if non-nil, zero value otherwise.
func (o *MicrosoftGraphRgbColor) GetB() int32 {
	if o == nil || o.B == nil {
		var ret int32
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRgbColor) GetBOk() (int32, bool) {
	if o == nil || o.B == nil {
		var ret int32
		return ret, false
	}
	return *o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *MicrosoftGraphRgbColor) HasB() bool {
	if o != nil && o.B != nil {
		return true
	}

	return false
}

// SetB gets a reference to the given int32 and assigns it to the B field.
func (o *MicrosoftGraphRgbColor) SetB(v int32) {
	o.B = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphRgbColor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.R != nil {
		toSerialize["r"] = o.R
	}
	if o.G != nil {
		toSerialize["g"] = o.G
	}
	if o.B != nil {
		toSerialize["b"] = o.B
	}
	return json.Marshal(toSerialize)
}

