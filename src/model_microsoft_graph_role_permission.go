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
// MicrosoftGraphRolePermission struct for MicrosoftGraphRolePermission
type MicrosoftGraphRolePermission struct {
	// Actions
	ResourceActions *[]AnyOfmicrosoftGraphResourceAction `json:"resourceActions,omitempty"`

}

// GetResourceActions returns the ResourceActions field if non-nil, zero value otherwise.
func (o *MicrosoftGraphRolePermission) GetResourceActions() []AnyOfmicrosoftGraphResourceAction {
	if o == nil || o.ResourceActions == nil {
		var ret []AnyOfmicrosoftGraphResourceAction
		return ret
	}
	return *o.ResourceActions
}

// GetResourceActionsOk returns a tuple with the ResourceActions field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRolePermission) GetResourceActionsOk() ([]AnyOfmicrosoftGraphResourceAction, bool) {
	if o == nil || o.ResourceActions == nil {
		var ret []AnyOfmicrosoftGraphResourceAction
		return ret, false
	}
	return *o.ResourceActions, true
}

// HasResourceActions returns a boolean if a field has been set.
func (o *MicrosoftGraphRolePermission) HasResourceActions() bool {
	if o != nil && o.ResourceActions != nil {
		return true
	}

	return false
}

// SetResourceActions gets a reference to the given []AnyOfmicrosoftGraphResourceAction and assigns it to the ResourceActions field.
func (o *MicrosoftGraphRolePermission) SetResourceActions(v []AnyOfmicrosoftGraphResourceAction) {
	o.ResourceActions = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphRolePermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceActions != nil {
		toSerialize["resourceActions"] = o.ResourceActions
	}
	return json.Marshal(toSerialize)
}

