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
// InlineObject744 struct for InlineObject744
type InlineObject744 struct {
	CalculationType *string `json:"calculationType,omitempty"`

}

// GetCalculationType returns the CalculationType field if non-nil, zero value otherwise.
func (o *InlineObject744) GetCalculationType() string {
	if o == nil || o.CalculationType == nil {
		var ret string
		return ret
	}
	return *o.CalculationType
}

// GetCalculationTypeOk returns a tuple with the CalculationType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject744) GetCalculationTypeOk() (string, bool) {
	if o == nil || o.CalculationType == nil {
		var ret string
		return ret, false
	}
	return *o.CalculationType, true
}

// HasCalculationType returns a boolean if a field has been set.
func (o *InlineObject744) HasCalculationType() bool {
	if o != nil && o.CalculationType != nil {
		return true
	}

	return false
}

// SetCalculationType gets a reference to the given string and assigns it to the CalculationType field.
func (o *InlineObject744) SetCalculationType(v string) {
	o.CalculationType = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject744) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CalculationType != nil {
		toSerialize["calculationType"] = o.CalculationType
	}
	return json.Marshal(toSerialize)
}


