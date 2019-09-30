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
// InlineObject634 struct for InlineObject634
type InlineObject634 struct {
	StorageLocation *string `json:"storageLocation,omitempty"`
	isExplicitNullStorageLocation bool `json:"-"`
}

// GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.
func (o *InlineObject634) GetStorageLocation() string {
	if o == nil || o.StorageLocation == nil {
		var ret string
		return ret
	}
	return *o.StorageLocation
}

// GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject634) GetStorageLocationOk() (string, bool) {
	if o == nil || o.StorageLocation == nil {
		var ret string
		return ret, false
	}
	return *o.StorageLocation, true
}

// HasStorageLocation returns a boolean if a field has been set.
func (o *InlineObject634) HasStorageLocation() bool {
	if o != nil && o.StorageLocation != nil {
		return true
	}

	return false
}

// SetStorageLocation gets a reference to the given string and assigns it to the StorageLocation field.
func (o *InlineObject634) SetStorageLocation(v string) {
	o.StorageLocation = &v
}

// SetStorageLocationExplicitNull (un)sets StorageLocation to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The StorageLocation value is set to nil even if false is passed
func (o *InlineObject634) SetStorageLocationExplicitNull(b bool) {
	o.StorageLocation = nil
	o.isExplicitNullStorageLocation = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject634) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StorageLocation == nil {
		if o.isExplicitNullStorageLocation {
			toSerialize["storageLocation"] = o.StorageLocation
		}
	} else {
		toSerialize["storageLocation"] = o.StorageLocation
	}
	return json.Marshal(toSerialize)
}


