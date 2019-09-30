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
// GroupAssignmentTarget struct for GroupAssignmentTarget
type GroupAssignmentTarget struct {
	// The group Id that is the target of the assignment.
	GroupId *string `json:"groupId,omitempty"`
	isExplicitNullGroupId bool `json:"-"`
}

// GetGroupId returns the GroupId field if non-nil, zero value otherwise.
func (o *GroupAssignmentTarget) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *GroupAssignmentTarget) GetGroupIdOk() (string, bool) {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret, false
	}
	return *o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupAssignmentTarget) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupAssignmentTarget) SetGroupId(v string) {
	o.GroupId = &v
}

// SetGroupIdExplicitNull (un)sets GroupId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The GroupId value is set to nil even if false is passed
func (o *GroupAssignmentTarget) SetGroupIdExplicitNull(b bool) {
	o.GroupId = nil
	o.isExplicitNullGroupId = b
}

// MarshalJSON returns the JSON representation of the model.
func (o GroupAssignmentTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupId == nil {
		if o.isExplicitNullGroupId {
			toSerialize["groupId"] = o.GroupId
		}
	} else {
		toSerialize["groupId"] = o.GroupId
	}
	return json.Marshal(toSerialize)
}

