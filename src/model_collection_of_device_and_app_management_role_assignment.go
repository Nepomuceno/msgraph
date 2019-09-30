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
// CollectionOfDeviceAndAppManagementRoleAssignment struct for CollectionOfDeviceAndAppManagementRoleAssignment
type CollectionOfDeviceAndAppManagementRoleAssignment struct {
	Value *[]MicrosoftGraphDeviceAndAppManagementRoleAssignment `json:"value,omitempty"`

}

// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *CollectionOfDeviceAndAppManagementRoleAssignment) GetValue() []MicrosoftGraphDeviceAndAppManagementRoleAssignment {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphDeviceAndAppManagementRoleAssignment
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceAndAppManagementRoleAssignment) GetValueOk() ([]MicrosoftGraphDeviceAndAppManagementRoleAssignment, bool) {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphDeviceAndAppManagementRoleAssignment
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDeviceAndAppManagementRoleAssignment) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphDeviceAndAppManagementRoleAssignment and assigns it to the Value field.
func (o *CollectionOfDeviceAndAppManagementRoleAssignment) SetValue(v []MicrosoftGraphDeviceAndAppManagementRoleAssignment) {
	o.Value = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o CollectionOfDeviceAndAppManagementRoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}


