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
// InlineObject934 struct for InlineObject934
type InlineObject934 struct {
	Reference *AnyOfobject `json:"reference,omitempty"`
	isExplicitNullReference bool `json:"-"`
}

// GetReference returns the Reference field if non-nil, zero value otherwise.
func (o *InlineObject934) GetReference() AnyOfobject {
	if o == nil || o.Reference == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject934) GetReferenceOk() (AnyOfobject, bool) {
	if o == nil || o.Reference == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *InlineObject934) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given AnyOfobject and assigns it to the Reference field.
func (o *InlineObject934) SetReference(v AnyOfobject) {
	o.Reference = &v
}

// SetReferenceExplicitNull (un)sets Reference to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Reference value is set to nil even if false is passed
func (o *InlineObject934) SetReferenceExplicitNull(b bool) {
	o.Reference = nil
	o.isExplicitNullReference = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject934) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reference == nil {
		if o.isExplicitNullReference {
			toSerialize["reference"] = o.Reference
		}
	} else {
		toSerialize["reference"] = o.Reference
	}
	return json.Marshal(toSerialize)
}


