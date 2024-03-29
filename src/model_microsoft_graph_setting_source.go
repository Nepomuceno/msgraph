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
// MicrosoftGraphSettingSource struct for MicrosoftGraphSettingSource
type MicrosoftGraphSettingSource struct {
	Id *string `json:"id,omitempty"`
	isExplicitNullId bool `json:"-"`
	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSettingSource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSettingSource) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphSettingSource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphSettingSource) SetId(v string) {
	o.Id = &v
}

// SetIdExplicitNull (un)sets Id to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Id value is set to nil even if false is passed
func (o *MicrosoftGraphSettingSource) SetIdExplicitNull(b bool) {
	o.Id = nil
	o.isExplicitNullId = b
}
// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSettingSource) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSettingSource) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphSettingSource) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphSettingSource) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphSettingSource) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphSettingSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id == nil {
		if o.isExplicitNullId {
			toSerialize["id"] = o.Id
		}
	} else {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}


