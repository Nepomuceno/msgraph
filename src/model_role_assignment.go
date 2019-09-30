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
// RoleAssignment The Role Assignment resource. Role assignments tie together a role definition with members and scopes. There can be one or more role assignments per role. This applies to custom and built-in roles.
type RoleAssignment struct {
	// The display or friendly name of the role Assignment.
	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	// Description of the Role Assignment.
	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
	// List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
	ResourceScopes *[]string `json:"resourceScopes,omitempty"`

	RoleDefinition *AnyOfmicrosoftGraphRoleDefinition `json:"roleDefinition,omitempty"`
	isExplicitNullRoleDefinition bool `json:"-"`
}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *RoleAssignment) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *RoleAssignment) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *RoleAssignment) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *RoleAssignment) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *RoleAssignment) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleAssignment) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleAssignment) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *RoleAssignment) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetResourceScopes returns the ResourceScopes field if non-nil, zero value otherwise.
func (o *RoleAssignment) GetResourceScopes() []string {
	if o == nil || o.ResourceScopes == nil {
		var ret []string
		return ret
	}
	return *o.ResourceScopes
}

// GetResourceScopesOk returns a tuple with the ResourceScopes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetResourceScopesOk() ([]string, bool) {
	if o == nil || o.ResourceScopes == nil {
		var ret []string
		return ret, false
	}
	return *o.ResourceScopes, true
}

// HasResourceScopes returns a boolean if a field has been set.
func (o *RoleAssignment) HasResourceScopes() bool {
	if o != nil && o.ResourceScopes != nil {
		return true
	}

	return false
}

// SetResourceScopes gets a reference to the given []string and assigns it to the ResourceScopes field.
func (o *RoleAssignment) SetResourceScopes(v []string) {
	o.ResourceScopes = &v
}

// GetRoleDefinition returns the RoleDefinition field if non-nil, zero value otherwise.
func (o *RoleAssignment) GetRoleDefinition() AnyOfmicrosoftGraphRoleDefinition {
	if o == nil || o.RoleDefinition == nil {
		var ret AnyOfmicrosoftGraphRoleDefinition
		return ret
	}
	return *o.RoleDefinition
}

// GetRoleDefinitionOk returns a tuple with the RoleDefinition field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignment) GetRoleDefinitionOk() (AnyOfmicrosoftGraphRoleDefinition, bool) {
	if o == nil || o.RoleDefinition == nil {
		var ret AnyOfmicrosoftGraphRoleDefinition
		return ret, false
	}
	return *o.RoleDefinition, true
}

// HasRoleDefinition returns a boolean if a field has been set.
func (o *RoleAssignment) HasRoleDefinition() bool {
	if o != nil && o.RoleDefinition != nil {
		return true
	}

	return false
}

// SetRoleDefinition gets a reference to the given AnyOfmicrosoftGraphRoleDefinition and assigns it to the RoleDefinition field.
func (o *RoleAssignment) SetRoleDefinition(v AnyOfmicrosoftGraphRoleDefinition) {
	o.RoleDefinition = &v
}

// SetRoleDefinitionExplicitNull (un)sets RoleDefinition to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The RoleDefinition value is set to nil even if false is passed
func (o *RoleAssignment) SetRoleDefinitionExplicitNull(b bool) {
	o.RoleDefinition = nil
	o.isExplicitNullRoleDefinition = b
}

// MarshalJSON returns the JSON representation of the model.
func (o RoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description == nil {
		if o.isExplicitNullDescription {
			toSerialize["description"] = o.Description
		}
	} else {
		toSerialize["description"] = o.Description
	}
	if o.ResourceScopes != nil {
		toSerialize["resourceScopes"] = o.ResourceScopes
	}
	if o.RoleDefinition == nil {
		if o.isExplicitNullRoleDefinition {
			toSerialize["roleDefinition"] = o.RoleDefinition
		}
	} else {
		toSerialize["roleDefinition"] = o.RoleDefinition
	}
	return json.Marshal(toSerialize)
}

