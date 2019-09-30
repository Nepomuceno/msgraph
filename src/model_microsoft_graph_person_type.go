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
// MicrosoftGraphPersonType struct for MicrosoftGraphPersonType
type MicrosoftGraphPersonType struct {
	Class *string `json:"class,omitempty"`
	isExplicitNullClass bool `json:"-"`
	Subclass *string `json:"subclass,omitempty"`
	isExplicitNullSubclass bool `json:"-"`
}

// GetClass returns the Class field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPersonType) GetClass() string {
	if o == nil || o.Class == nil {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPersonType) GetClassOk() (string, bool) {
	if o == nil || o.Class == nil {
		var ret string
		return ret, false
	}
	return *o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *MicrosoftGraphPersonType) HasClass() bool {
	if o != nil && o.Class != nil {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *MicrosoftGraphPersonType) SetClass(v string) {
	o.Class = &v
}

// SetClassExplicitNull (un)sets Class to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Class value is set to nil even if false is passed
func (o *MicrosoftGraphPersonType) SetClassExplicitNull(b bool) {
	o.Class = nil
	o.isExplicitNullClass = b
}
// GetSubclass returns the Subclass field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPersonType) GetSubclass() string {
	if o == nil || o.Subclass == nil {
		var ret string
		return ret
	}
	return *o.Subclass
}

// GetSubclassOk returns a tuple with the Subclass field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPersonType) GetSubclassOk() (string, bool) {
	if o == nil || o.Subclass == nil {
		var ret string
		return ret, false
	}
	return *o.Subclass, true
}

// HasSubclass returns a boolean if a field has been set.
func (o *MicrosoftGraphPersonType) HasSubclass() bool {
	if o != nil && o.Subclass != nil {
		return true
	}

	return false
}

// SetSubclass gets a reference to the given string and assigns it to the Subclass field.
func (o *MicrosoftGraphPersonType) SetSubclass(v string) {
	o.Subclass = &v
}

// SetSubclassExplicitNull (un)sets Subclass to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Subclass value is set to nil even if false is passed
func (o *MicrosoftGraphPersonType) SetSubclassExplicitNull(b bool) {
	o.Subclass = nil
	o.isExplicitNullSubclass = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphPersonType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Class == nil {
		if o.isExplicitNullClass {
			toSerialize["class"] = o.Class
		}
	} else {
		toSerialize["class"] = o.Class
	}
	if o.Subclass == nil {
		if o.isExplicitNullSubclass {
			toSerialize["subclass"] = o.Subclass
		}
	} else {
		toSerialize["subclass"] = o.Subclass
	}
	return json.Marshal(toSerialize)
}

