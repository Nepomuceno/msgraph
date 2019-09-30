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
// InlineObject4 struct for InlineObject4
type InlineObject4 struct {
	Assignments *[]AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationAssignment `json:"assignments,omitempty"`

}

// GetAssignments returns the Assignments field if non-nil, zero value otherwise.
func (o *InlineObject4) GetAssignments() []AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationAssignment {
	if o == nil || o.Assignments == nil {
		var ret []AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetAssignmentsOk() ([]AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationAssignment, bool) {
	if o == nil || o.Assignments == nil {
		var ret []AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationAssignment
		return ret, false
	}
	return *o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *InlineObject4) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationAssignment and assigns it to the Assignments field.
func (o *InlineObject4) SetAssignments(v []AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationAssignment) {
	o.Assignments = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	return json.Marshal(toSerialize)
}


