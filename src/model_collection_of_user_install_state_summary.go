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
// CollectionOfUserInstallStateSummary struct for CollectionOfUserInstallStateSummary
type CollectionOfUserInstallStateSummary struct {
	Value *[]MicrosoftGraphUserInstallStateSummary `json:"value,omitempty"`

}

// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *CollectionOfUserInstallStateSummary) GetValue() []MicrosoftGraphUserInstallStateSummary {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphUserInstallStateSummary
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfUserInstallStateSummary) GetValueOk() ([]MicrosoftGraphUserInstallStateSummary, bool) {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphUserInstallStateSummary
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfUserInstallStateSummary) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphUserInstallStateSummary and assigns it to the Value field.
func (o *CollectionOfUserInstallStateSummary) SetValue(v []MicrosoftGraphUserInstallStateSummary) {
	o.Value = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o CollectionOfUserInstallStateSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

