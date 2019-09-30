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
// InlineObject629 struct for InlineObject629
type InlineObject629 struct {
	DestinationId *string `json:"DestinationId,omitempty"`

}

// GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.
func (o *InlineObject629) GetDestinationId() string {
	if o == nil || o.DestinationId == nil {
		var ret string
		return ret
	}
	return *o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject629) GetDestinationIdOk() (string, bool) {
	if o == nil || o.DestinationId == nil {
		var ret string
		return ret, false
	}
	return *o.DestinationId, true
}

// HasDestinationId returns a boolean if a field has been set.
func (o *InlineObject629) HasDestinationId() bool {
	if o != nil && o.DestinationId != nil {
		return true
	}

	return false
}

// SetDestinationId gets a reference to the given string and assigns it to the DestinationId field.
func (o *InlineObject629) SetDestinationId(v string) {
	o.DestinationId = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject629) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DestinationId != nil {
		toSerialize["DestinationId"] = o.DestinationId
	}
	return json.Marshal(toSerialize)
}


