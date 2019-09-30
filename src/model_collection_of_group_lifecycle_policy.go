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
// CollectionOfGroupLifecyclePolicy struct for CollectionOfGroupLifecyclePolicy
type CollectionOfGroupLifecyclePolicy struct {
	Value *[]MicrosoftGraphGroupLifecyclePolicy `json:"value,omitempty"`

}

// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *CollectionOfGroupLifecyclePolicy) GetValue() []MicrosoftGraphGroupLifecyclePolicy {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphGroupLifecyclePolicy
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfGroupLifecyclePolicy) GetValueOk() ([]MicrosoftGraphGroupLifecyclePolicy, bool) {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphGroupLifecyclePolicy
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfGroupLifecyclePolicy) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphGroupLifecyclePolicy and assigns it to the Value field.
func (o *CollectionOfGroupLifecyclePolicy) SetValue(v []MicrosoftGraphGroupLifecyclePolicy) {
	o.Value = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o CollectionOfGroupLifecyclePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

