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
// InlineObject17 struct for InlineObject17
type InlineObject17 struct {
	Priority *int32 `json:"priority,omitempty"`

}

// GetPriority returns the Priority field if non-nil, zero value otherwise.
func (o *InlineObject17) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject17) GetPriorityOk() (int32, bool) {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret, false
	}
	return *o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *InlineObject17) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *InlineObject17) SetPriority(v int32) {
	o.Priority = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject17) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	return json.Marshal(toSerialize)
}


