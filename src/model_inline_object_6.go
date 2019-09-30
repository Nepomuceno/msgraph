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
// InlineObject6 struct for InlineObject6
type InlineObject6 struct {
	Assignments *[]AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`

}

// GetAssignments returns the Assignments field if non-nil, zero value otherwise.
func (o *InlineObject6) GetAssignments() []AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment {
	if o == nil || o.Assignments == nil {
		var ret []AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetAssignmentsOk() ([]AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment, bool) {
	if o == nil || o.Assignments == nil {
		var ret []AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment
		return ret, false
	}
	return *o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *InlineObject6) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.
func (o *InlineObject6) SetAssignments(v []AnyOfmicrosoftGraphTargetedManagedAppPolicyAssignment) {
	o.Assignments = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	return json.Marshal(toSerialize)
}


