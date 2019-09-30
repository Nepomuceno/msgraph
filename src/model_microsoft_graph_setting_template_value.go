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
// MicrosoftGraphSettingTemplateValue struct for MicrosoftGraphSettingTemplateValue
type MicrosoftGraphSettingTemplateValue struct {
	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
	Type *string `json:"type,omitempty"`
	isExplicitNullType bool `json:"-"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	isExplicitNullDefaultValue bool `json:"-"`
	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSettingTemplateValue) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSettingTemplateValue) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphSettingTemplateValue) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftGraphSettingTemplateValue) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *MicrosoftGraphSettingTemplateValue) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetType returns the Type field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSettingTemplateValue) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSettingTemplateValue) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphSettingTemplateValue) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MicrosoftGraphSettingTemplateValue) SetType(v string) {
	o.Type = &v
}

// SetTypeExplicitNull (un)sets Type to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Type value is set to nil even if false is passed
func (o *MicrosoftGraphSettingTemplateValue) SetTypeExplicitNull(b bool) {
	o.Type = nil
	o.isExplicitNullType = b
}
// GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSettingTemplateValue) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSettingTemplateValue) GetDefaultValueOk() (string, bool) {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret, false
	}
	return *o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *MicrosoftGraphSettingTemplateValue) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *MicrosoftGraphSettingTemplateValue) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// SetDefaultValueExplicitNull (un)sets DefaultValue to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DefaultValue value is set to nil even if false is passed
func (o *MicrosoftGraphSettingTemplateValue) SetDefaultValueExplicitNull(b bool) {
	o.DefaultValue = nil
	o.isExplicitNullDefaultValue = b
}
// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSettingTemplateValue) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSettingTemplateValue) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphSettingTemplateValue) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphSettingTemplateValue) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphSettingTemplateValue) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphSettingTemplateValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
	}
	if o.Type == nil {
		if o.isExplicitNullType {
			toSerialize["type"] = o.Type
		}
	} else {
		toSerialize["type"] = o.Type
	}
	if o.DefaultValue == nil {
		if o.isExplicitNullDefaultValue {
			toSerialize["defaultValue"] = o.DefaultValue
		}
	} else {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.Description == nil {
		if o.isExplicitNullDescription {
			toSerialize["description"] = o.Description
		}
	} else {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

