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
// MicrosoftGraphAndroidMinimumOperatingSystem struct for MicrosoftGraphAndroidMinimumOperatingSystem
type MicrosoftGraphAndroidMinimumOperatingSystem struct {
	// Version 4.0 or later.
	V40 *bool `json:"v4_0,omitempty"`

	// Version 4.0.3 or later.
	V403 *bool `json:"v4_0_3,omitempty"`

	// Version 4.1 or later.
	V41 *bool `json:"v4_1,omitempty"`

	// Version 4.2 or later.
	V42 *bool `json:"v4_2,omitempty"`

	// Version 4.3 or later.
	V43 *bool `json:"v4_3,omitempty"`

	// Version 4.4 or later.
	V44 *bool `json:"v4_4,omitempty"`

	// Version 5.0 or later.
	V50 *bool `json:"v5_0,omitempty"`

	// Version 5.1 or later.
	V51 *bool `json:"v5_1,omitempty"`

}

// GetV40 returns the V40 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV40() bool {
	if o == nil || o.V40 == nil {
		var ret bool
		return ret
	}
	return *o.V40
}

// GetV40Ok returns a tuple with the V40 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV40Ok() (bool, bool) {
	if o == nil || o.V40 == nil {
		var ret bool
		return ret, false
	}
	return *o.V40, true
}

// HasV40 returns a boolean if a field has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) HasV40() bool {
	if o != nil && o.V40 != nil {
		return true
	}

	return false
}

// SetV40 gets a reference to the given bool and assigns it to the V40 field.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) SetV40(v bool) {
	o.V40 = &v
}

// GetV403 returns the V403 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV403() bool {
	if o == nil || o.V403 == nil {
		var ret bool
		return ret
	}
	return *o.V403
}

// GetV403Ok returns a tuple with the V403 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV403Ok() (bool, bool) {
	if o == nil || o.V403 == nil {
		var ret bool
		return ret, false
	}
	return *o.V403, true
}

// HasV403 returns a boolean if a field has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) HasV403() bool {
	if o != nil && o.V403 != nil {
		return true
	}

	return false
}

// SetV403 gets a reference to the given bool and assigns it to the V403 field.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) SetV403(v bool) {
	o.V403 = &v
}

// GetV41 returns the V41 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV41() bool {
	if o == nil || o.V41 == nil {
		var ret bool
		return ret
	}
	return *o.V41
}

// GetV41Ok returns a tuple with the V41 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV41Ok() (bool, bool) {
	if o == nil || o.V41 == nil {
		var ret bool
		return ret, false
	}
	return *o.V41, true
}

// HasV41 returns a boolean if a field has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) HasV41() bool {
	if o != nil && o.V41 != nil {
		return true
	}

	return false
}

// SetV41 gets a reference to the given bool and assigns it to the V41 field.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) SetV41(v bool) {
	o.V41 = &v
}

// GetV42 returns the V42 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV42() bool {
	if o == nil || o.V42 == nil {
		var ret bool
		return ret
	}
	return *o.V42
}

// GetV42Ok returns a tuple with the V42 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV42Ok() (bool, bool) {
	if o == nil || o.V42 == nil {
		var ret bool
		return ret, false
	}
	return *o.V42, true
}

// HasV42 returns a boolean if a field has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) HasV42() bool {
	if o != nil && o.V42 != nil {
		return true
	}

	return false
}

// SetV42 gets a reference to the given bool and assigns it to the V42 field.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) SetV42(v bool) {
	o.V42 = &v
}

// GetV43 returns the V43 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV43() bool {
	if o == nil || o.V43 == nil {
		var ret bool
		return ret
	}
	return *o.V43
}

// GetV43Ok returns a tuple with the V43 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV43Ok() (bool, bool) {
	if o == nil || o.V43 == nil {
		var ret bool
		return ret, false
	}
	return *o.V43, true
}

// HasV43 returns a boolean if a field has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) HasV43() bool {
	if o != nil && o.V43 != nil {
		return true
	}

	return false
}

// SetV43 gets a reference to the given bool and assigns it to the V43 field.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) SetV43(v bool) {
	o.V43 = &v
}

// GetV44 returns the V44 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV44() bool {
	if o == nil || o.V44 == nil {
		var ret bool
		return ret
	}
	return *o.V44
}

// GetV44Ok returns a tuple with the V44 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV44Ok() (bool, bool) {
	if o == nil || o.V44 == nil {
		var ret bool
		return ret, false
	}
	return *o.V44, true
}

// HasV44 returns a boolean if a field has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) HasV44() bool {
	if o != nil && o.V44 != nil {
		return true
	}

	return false
}

// SetV44 gets a reference to the given bool and assigns it to the V44 field.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) SetV44(v bool) {
	o.V44 = &v
}

// GetV50 returns the V50 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV50() bool {
	if o == nil || o.V50 == nil {
		var ret bool
		return ret
	}
	return *o.V50
}

// GetV50Ok returns a tuple with the V50 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV50Ok() (bool, bool) {
	if o == nil || o.V50 == nil {
		var ret bool
		return ret, false
	}
	return *o.V50, true
}

// HasV50 returns a boolean if a field has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) HasV50() bool {
	if o != nil && o.V50 != nil {
		return true
	}

	return false
}

// SetV50 gets a reference to the given bool and assigns it to the V50 field.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) SetV50(v bool) {
	o.V50 = &v
}

// GetV51 returns the V51 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV51() bool {
	if o == nil || o.V51 == nil {
		var ret bool
		return ret
	}
	return *o.V51
}

// GetV51Ok returns a tuple with the V51 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) GetV51Ok() (bool, bool) {
	if o == nil || o.V51 == nil {
		var ret bool
		return ret, false
	}
	return *o.V51, true
}

// HasV51 returns a boolean if a field has been set.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) HasV51() bool {
	if o != nil && o.V51 != nil {
		return true
	}

	return false
}

// SetV51 gets a reference to the given bool and assigns it to the V51 field.
func (o *MicrosoftGraphAndroidMinimumOperatingSystem) SetV51(v bool) {
	o.V51 = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphAndroidMinimumOperatingSystem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.V40 != nil {
		toSerialize["v4_0"] = o.V40
	}
	if o.V403 != nil {
		toSerialize["v4_0_3"] = o.V403
	}
	if o.V41 != nil {
		toSerialize["v4_1"] = o.V41
	}
	if o.V42 != nil {
		toSerialize["v4_2"] = o.V42
	}
	if o.V43 != nil {
		toSerialize["v4_3"] = o.V43
	}
	if o.V44 != nil {
		toSerialize["v4_4"] = o.V44
	}
	if o.V50 != nil {
		toSerialize["v5_0"] = o.V50
	}
	if o.V51 != nil {
		toSerialize["v5_1"] = o.V51
	}
	return json.Marshal(toSerialize)
}

