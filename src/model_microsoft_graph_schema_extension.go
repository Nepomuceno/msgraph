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
// MicrosoftGraphSchemaExtension struct for MicrosoftGraphSchemaExtension
type MicrosoftGraphSchemaExtension struct {
	Id *string `json:"id,omitempty"`

	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
	TargetTypes *[]string `json:"targetTypes,omitempty"`

	Properties *[]MicrosoftGraphExtensionSchemaProperty `json:"properties,omitempty"`

	Status *string `json:"status,omitempty"`

	Owner *string `json:"owner,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSchemaExtension) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSchemaExtension) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphSchemaExtension) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphSchemaExtension) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSchemaExtension) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSchemaExtension) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphSchemaExtension) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphSchemaExtension) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphSchemaExtension) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetTargetTypes returns the TargetTypes field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSchemaExtension) GetTargetTypes() []string {
	if o == nil || o.TargetTypes == nil {
		var ret []string
		return ret
	}
	return *o.TargetTypes
}

// GetTargetTypesOk returns a tuple with the TargetTypes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSchemaExtension) GetTargetTypesOk() ([]string, bool) {
	if o == nil || o.TargetTypes == nil {
		var ret []string
		return ret, false
	}
	return *o.TargetTypes, true
}

// HasTargetTypes returns a boolean if a field has been set.
func (o *MicrosoftGraphSchemaExtension) HasTargetTypes() bool {
	if o != nil && o.TargetTypes != nil {
		return true
	}

	return false
}

// SetTargetTypes gets a reference to the given []string and assigns it to the TargetTypes field.
func (o *MicrosoftGraphSchemaExtension) SetTargetTypes(v []string) {
	o.TargetTypes = &v
}

// GetProperties returns the Properties field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSchemaExtension) GetProperties() []MicrosoftGraphExtensionSchemaProperty {
	if o == nil || o.Properties == nil {
		var ret []MicrosoftGraphExtensionSchemaProperty
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSchemaExtension) GetPropertiesOk() ([]MicrosoftGraphExtensionSchemaProperty, bool) {
	if o == nil || o.Properties == nil {
		var ret []MicrosoftGraphExtensionSchemaProperty
		return ret, false
	}
	return *o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *MicrosoftGraphSchemaExtension) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []MicrosoftGraphExtensionSchemaProperty and assigns it to the Properties field.
func (o *MicrosoftGraphSchemaExtension) SetProperties(v []MicrosoftGraphExtensionSchemaProperty) {
	o.Properties = &v
}

// GetStatus returns the Status field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSchemaExtension) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSchemaExtension) GetStatusOk() (string, bool) {
	if o == nil || o.Status == nil {
		var ret string
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphSchemaExtension) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MicrosoftGraphSchemaExtension) SetStatus(v string) {
	o.Status = &v
}

// GetOwner returns the Owner field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSchemaExtension) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSchemaExtension) GetOwnerOk() (string, bool) {
	if o == nil || o.Owner == nil {
		var ret string
		return ret, false
	}
	return *o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *MicrosoftGraphSchemaExtension) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *MicrosoftGraphSchemaExtension) SetOwner(v string) {
	o.Owner = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphSchemaExtension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Description == nil {
		if o.isExplicitNullDescription {
			toSerialize["description"] = o.Description
		}
	} else {
		toSerialize["description"] = o.Description
	}
	if o.TargetTypes != nil {
		toSerialize["targetTypes"] = o.TargetTypes
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}

