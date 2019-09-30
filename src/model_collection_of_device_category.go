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
// CollectionOfDeviceCategory struct for CollectionOfDeviceCategory
type CollectionOfDeviceCategory struct {
	Value *[]MicrosoftGraphDeviceCategory `json:"value,omitempty"`

}

// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *CollectionOfDeviceCategory) GetValue() []MicrosoftGraphDeviceCategory {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphDeviceCategory
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceCategory) GetValueOk() ([]MicrosoftGraphDeviceCategory, bool) {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphDeviceCategory
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDeviceCategory) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphDeviceCategory and assigns it to the Value field.
func (o *CollectionOfDeviceCategory) SetValue(v []MicrosoftGraphDeviceCategory) {
	o.Value = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o CollectionOfDeviceCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}


