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
// InlineObject1065 struct for InlineObject1065
type InlineObject1065 struct {
	Probability *AnyOfobject `json:"probability,omitempty"`
	isExplicitNullProbability bool `json:"-"`
	DegFreedom *AnyOfobject `json:"degFreedom,omitempty"`
	isExplicitNullDegFreedom bool `json:"-"`
}

// GetProbability returns the Probability field if non-nil, zero value otherwise.
func (o *InlineObject1065) GetProbability() AnyOfobject {
	if o == nil || o.Probability == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Probability
}

// GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1065) GetProbabilityOk() (AnyOfobject, bool) {
	if o == nil || o.Probability == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Probability, true
}

// HasProbability returns a boolean if a field has been set.
func (o *InlineObject1065) HasProbability() bool {
	if o != nil && o.Probability != nil {
		return true
	}

	return false
}

// SetProbability gets a reference to the given AnyOfobject and assigns it to the Probability field.
func (o *InlineObject1065) SetProbability(v AnyOfobject) {
	o.Probability = &v
}

// SetProbabilityExplicitNull (un)sets Probability to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Probability value is set to nil even if false is passed
func (o *InlineObject1065) SetProbabilityExplicitNull(b bool) {
	o.Probability = nil
	o.isExplicitNullProbability = b
}
// GetDegFreedom returns the DegFreedom field if non-nil, zero value otherwise.
func (o *InlineObject1065) GetDegFreedom() AnyOfobject {
	if o == nil || o.DegFreedom == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.DegFreedom
}

// GetDegFreedomOk returns a tuple with the DegFreedom field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1065) GetDegFreedomOk() (AnyOfobject, bool) {
	if o == nil || o.DegFreedom == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.DegFreedom, true
}

// HasDegFreedom returns a boolean if a field has been set.
func (o *InlineObject1065) HasDegFreedom() bool {
	if o != nil && o.DegFreedom != nil {
		return true
	}

	return false
}

// SetDegFreedom gets a reference to the given AnyOfobject and assigns it to the DegFreedom field.
func (o *InlineObject1065) SetDegFreedom(v AnyOfobject) {
	o.DegFreedom = &v
}

// SetDegFreedomExplicitNull (un)sets DegFreedom to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DegFreedom value is set to nil even if false is passed
func (o *InlineObject1065) SetDegFreedomExplicitNull(b bool) {
	o.DegFreedom = nil
	o.isExplicitNullDegFreedom = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1065) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Probability == nil {
		if o.isExplicitNullProbability {
			toSerialize["probability"] = o.Probability
		}
	} else {
		toSerialize["probability"] = o.Probability
	}
	if o.DegFreedom == nil {
		if o.isExplicitNullDegFreedom {
			toSerialize["degFreedom"] = o.DegFreedom
		}
	} else {
		toSerialize["degFreedom"] = o.DegFreedom
	}
	return json.Marshal(toSerialize)
}


