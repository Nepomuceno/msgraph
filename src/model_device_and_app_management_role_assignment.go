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
// DeviceAndAppManagementRoleAssignment The Role Assignment resource. Role assignments tie together a role definition with members and scopes. There can be one or more role assignments per role. This applies to custom and built-in roles.
type DeviceAndAppManagementRoleAssignment struct {
	// The list of ids of role member security groups. These are IDs from Azure Active Directory.
	Members *[]string `json:"members,omitempty"`

}

// GetMembers returns the Members field if non-nil, zero value otherwise.
func (o *DeviceAndAppManagementRoleAssignment) GetMembers() []string {
	if o == nil || o.Members == nil {
		var ret []string
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAndAppManagementRoleAssignment) GetMembersOk() ([]string, bool) {
	if o == nil || o.Members == nil {
		var ret []string
		return ret, false
	}
	return *o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *DeviceAndAppManagementRoleAssignment) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *DeviceAndAppManagementRoleAssignment) SetMembers(v []string) {
	o.Members = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o DeviceAndAppManagementRoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	return json.Marshal(toSerialize)
}


