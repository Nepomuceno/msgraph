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
// InlineObject923 struct for InlineObject923
type InlineObject923 struct {
	Inumber *AnyOfobject `json:"inumber,omitempty"`
	isExplicitNullInumber bool `json:"-"`
}

// GetInumber returns the Inumber field if non-nil, zero value otherwise.
func (o *InlineObject923) GetInumber() AnyOfobject {
	if o == nil || o.Inumber == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Inumber
}

// GetInumberOk returns a tuple with the Inumber field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject923) GetInumberOk() (AnyOfobject, bool) {
	if o == nil || o.Inumber == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Inumber, true
}

// HasInumber returns a boolean if a field has been set.
func (o *InlineObject923) HasInumber() bool {
	if o != nil && o.Inumber != nil {
		return true
	}

	return false
}

// SetInumber gets a reference to the given AnyOfobject and assigns it to the Inumber field.
func (o *InlineObject923) SetInumber(v AnyOfobject) {
	o.Inumber = &v
}

// SetInumberExplicitNull (un)sets Inumber to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Inumber value is set to nil even if false is passed
func (o *InlineObject923) SetInumberExplicitNull(b bool) {
	o.Inumber = nil
	o.isExplicitNullInumber = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject923) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Inumber == nil {
		if o.isExplicitNullInumber {
			toSerialize["inumber"] = o.Inumber
		}
	} else {
		toSerialize["inumber"] = o.Inumber
	}
	return json.Marshal(toSerialize)
}


