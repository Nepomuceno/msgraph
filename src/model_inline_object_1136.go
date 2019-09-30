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
// InlineObject1136 struct for InlineObject1136
type InlineObject1136 struct {
	Percent *int32 `json:"percent,omitempty"`

}

// GetPercent returns the Percent field if non-nil, zero value otherwise.
func (o *InlineObject1136) GetPercent() int32 {
	if o == nil || o.Percent == nil {
		var ret int32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1136) GetPercentOk() (int32, bool) {
	if o == nil || o.Percent == nil {
		var ret int32
		return ret, false
	}
	return *o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *InlineObject1136) HasPercent() bool {
	if o != nil && o.Percent != nil {
		return true
	}

	return false
}

// SetPercent gets a reference to the given int32 and assigns it to the Percent field.
func (o *InlineObject1136) SetPercent(v int32) {
	o.Percent = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1136) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Percent != nil {
		toSerialize["percent"] = o.Percent
	}
	return json.Marshal(toSerialize)
}


