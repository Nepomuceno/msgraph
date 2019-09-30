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
// MicrosoftGraphNumberColumn struct for MicrosoftGraphNumberColumn
type MicrosoftGraphNumberColumn struct {
	DecimalPlaces *string `json:"decimalPlaces,omitempty"`
	isExplicitNullDecimalPlaces bool `json:"-"`
	DisplayAs *string `json:"displayAs,omitempty"`
	isExplicitNullDisplayAs bool `json:"-"`
	Maximum *AnyOfnumberstringstring `json:"maximum,omitempty"`
	isExplicitNullMaximum bool `json:"-"`
	Minimum *AnyOfnumberstringstring `json:"minimum,omitempty"`
	isExplicitNullMinimum bool `json:"-"`
}

// GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.
func (o *MicrosoftGraphNumberColumn) GetDecimalPlaces() string {
	if o == nil || o.DecimalPlaces == nil {
		var ret string
		return ret
	}
	return *o.DecimalPlaces
}

// GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphNumberColumn) GetDecimalPlacesOk() (string, bool) {
	if o == nil || o.DecimalPlaces == nil {
		var ret string
		return ret, false
	}
	return *o.DecimalPlaces, true
}

// HasDecimalPlaces returns a boolean if a field has been set.
func (o *MicrosoftGraphNumberColumn) HasDecimalPlaces() bool {
	if o != nil && o.DecimalPlaces != nil {
		return true
	}

	return false
}

// SetDecimalPlaces gets a reference to the given string and assigns it to the DecimalPlaces field.
func (o *MicrosoftGraphNumberColumn) SetDecimalPlaces(v string) {
	o.DecimalPlaces = &v
}

// SetDecimalPlacesExplicitNull (un)sets DecimalPlaces to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DecimalPlaces value is set to nil even if false is passed
func (o *MicrosoftGraphNumberColumn) SetDecimalPlacesExplicitNull(b bool) {
	o.DecimalPlaces = nil
	o.isExplicitNullDecimalPlaces = b
}
// GetDisplayAs returns the DisplayAs field if non-nil, zero value otherwise.
func (o *MicrosoftGraphNumberColumn) GetDisplayAs() string {
	if o == nil || o.DisplayAs == nil {
		var ret string
		return ret
	}
	return *o.DisplayAs
}

// GetDisplayAsOk returns a tuple with the DisplayAs field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphNumberColumn) GetDisplayAsOk() (string, bool) {
	if o == nil || o.DisplayAs == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayAs, true
}

// HasDisplayAs returns a boolean if a field has been set.
func (o *MicrosoftGraphNumberColumn) HasDisplayAs() bool {
	if o != nil && o.DisplayAs != nil {
		return true
	}

	return false
}

// SetDisplayAs gets a reference to the given string and assigns it to the DisplayAs field.
func (o *MicrosoftGraphNumberColumn) SetDisplayAs(v string) {
	o.DisplayAs = &v
}

// SetDisplayAsExplicitNull (un)sets DisplayAs to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayAs value is set to nil even if false is passed
func (o *MicrosoftGraphNumberColumn) SetDisplayAsExplicitNull(b bool) {
	o.DisplayAs = nil
	o.isExplicitNullDisplayAs = b
}
// GetMaximum returns the Maximum field if non-nil, zero value otherwise.
func (o *MicrosoftGraphNumberColumn) GetMaximum() AnyOfnumberstringstring {
	if o == nil || o.Maximum == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphNumberColumn) GetMaximumOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Maximum == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *MicrosoftGraphNumberColumn) HasMaximum() bool {
	if o != nil && o.Maximum != nil {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given AnyOfnumberstringstring and assigns it to the Maximum field.
func (o *MicrosoftGraphNumberColumn) SetMaximum(v AnyOfnumberstringstring) {
	o.Maximum = &v
}

// SetMaximumExplicitNull (un)sets Maximum to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Maximum value is set to nil even if false is passed
func (o *MicrosoftGraphNumberColumn) SetMaximumExplicitNull(b bool) {
	o.Maximum = nil
	o.isExplicitNullMaximum = b
}
// GetMinimum returns the Minimum field if non-nil, zero value otherwise.
func (o *MicrosoftGraphNumberColumn) GetMinimum() AnyOfnumberstringstring {
	if o == nil || o.Minimum == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphNumberColumn) GetMinimumOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Minimum == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *MicrosoftGraphNumberColumn) HasMinimum() bool {
	if o != nil && o.Minimum != nil {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given AnyOfnumberstringstring and assigns it to the Minimum field.
func (o *MicrosoftGraphNumberColumn) SetMinimum(v AnyOfnumberstringstring) {
	o.Minimum = &v
}

// SetMinimumExplicitNull (un)sets Minimum to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Minimum value is set to nil even if false is passed
func (o *MicrosoftGraphNumberColumn) SetMinimumExplicitNull(b bool) {
	o.Minimum = nil
	o.isExplicitNullMinimum = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphNumberColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DecimalPlaces == nil {
		if o.isExplicitNullDecimalPlaces {
			toSerialize["decimalPlaces"] = o.DecimalPlaces
		}
	} else {
		toSerialize["decimalPlaces"] = o.DecimalPlaces
	}
	if o.DisplayAs == nil {
		if o.isExplicitNullDisplayAs {
			toSerialize["displayAs"] = o.DisplayAs
		}
	} else {
		toSerialize["displayAs"] = o.DisplayAs
	}
	if o.Maximum == nil {
		if o.isExplicitNullMaximum {
			toSerialize["maximum"] = o.Maximum
		}
	} else {
		toSerialize["maximum"] = o.Maximum
	}
	if o.Minimum == nil {
		if o.isExplicitNullMinimum {
			toSerialize["minimum"] = o.Minimum
		}
	} else {
		toSerialize["minimum"] = o.Minimum
	}
	return json.Marshal(toSerialize)
}

