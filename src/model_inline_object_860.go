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
// InlineObject860 struct for InlineObject860
type InlineObject860 struct {
	LowerLimit *AnyOfobject `json:"lowerLimit,omitempty"`
	isExplicitNullLowerLimit bool `json:"-"`
	UpperLimit *AnyOfobject `json:"upperLimit,omitempty"`
	isExplicitNullUpperLimit bool `json:"-"`
}

// GetLowerLimit returns the LowerLimit field if non-nil, zero value otherwise.
func (o *InlineObject860) GetLowerLimit() AnyOfobject {
	if o == nil || o.LowerLimit == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.LowerLimit
}

// GetLowerLimitOk returns a tuple with the LowerLimit field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject860) GetLowerLimitOk() (AnyOfobject, bool) {
	if o == nil || o.LowerLimit == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.LowerLimit, true
}

// HasLowerLimit returns a boolean if a field has been set.
func (o *InlineObject860) HasLowerLimit() bool {
	if o != nil && o.LowerLimit != nil {
		return true
	}

	return false
}

// SetLowerLimit gets a reference to the given AnyOfobject and assigns it to the LowerLimit field.
func (o *InlineObject860) SetLowerLimit(v AnyOfobject) {
	o.LowerLimit = &v
}

// SetLowerLimitExplicitNull (un)sets LowerLimit to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LowerLimit value is set to nil even if false is passed
func (o *InlineObject860) SetLowerLimitExplicitNull(b bool) {
	o.LowerLimit = nil
	o.isExplicitNullLowerLimit = b
}
// GetUpperLimit returns the UpperLimit field if non-nil, zero value otherwise.
func (o *InlineObject860) GetUpperLimit() AnyOfobject {
	if o == nil || o.UpperLimit == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.UpperLimit
}

// GetUpperLimitOk returns a tuple with the UpperLimit field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject860) GetUpperLimitOk() (AnyOfobject, bool) {
	if o == nil || o.UpperLimit == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.UpperLimit, true
}

// HasUpperLimit returns a boolean if a field has been set.
func (o *InlineObject860) HasUpperLimit() bool {
	if o != nil && o.UpperLimit != nil {
		return true
	}

	return false
}

// SetUpperLimit gets a reference to the given AnyOfobject and assigns it to the UpperLimit field.
func (o *InlineObject860) SetUpperLimit(v AnyOfobject) {
	o.UpperLimit = &v
}

// SetUpperLimitExplicitNull (un)sets UpperLimit to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UpperLimit value is set to nil even if false is passed
func (o *InlineObject860) SetUpperLimitExplicitNull(b bool) {
	o.UpperLimit = nil
	o.isExplicitNullUpperLimit = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject860) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LowerLimit == nil {
		if o.isExplicitNullLowerLimit {
			toSerialize["lowerLimit"] = o.LowerLimit
		}
	} else {
		toSerialize["lowerLimit"] = o.LowerLimit
	}
	if o.UpperLimit == nil {
		if o.isExplicitNullUpperLimit {
			toSerialize["upperLimit"] = o.UpperLimit
		}
	} else {
		toSerialize["upperLimit"] = o.UpperLimit
	}
	return json.Marshal(toSerialize)
}


