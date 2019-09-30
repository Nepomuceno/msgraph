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
// MicrosoftGraphResourceOperation struct for MicrosoftGraphResourceOperation
type MicrosoftGraphResourceOperation struct {
	Id *string `json:"id,omitempty"`

	// Name of the Resource this operation is performed on.
	ResourceName *string `json:"resourceName,omitempty"`
	isExplicitNullResourceName bool `json:"-"`
	// Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
	ActionName *string `json:"actionName,omitempty"`
	isExplicitNullActionName bool `json:"-"`
	// Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphResourceOperation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResourceOperation) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceOperation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphResourceOperation) SetId(v string) {
	o.Id = &v
}

// GetResourceName returns the ResourceName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphResourceOperation) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResourceOperation) GetResourceNameOk() (string, bool) {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret, false
	}
	return *o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceOperation) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *MicrosoftGraphResourceOperation) SetResourceName(v string) {
	o.ResourceName = &v
}

// SetResourceNameExplicitNull (un)sets ResourceName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ResourceName value is set to nil even if false is passed
func (o *MicrosoftGraphResourceOperation) SetResourceNameExplicitNull(b bool) {
	o.ResourceName = nil
	o.isExplicitNullResourceName = b
}
// GetActionName returns the ActionName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphResourceOperation) GetActionName() string {
	if o == nil || o.ActionName == nil {
		var ret string
		return ret
	}
	return *o.ActionName
}

// GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResourceOperation) GetActionNameOk() (string, bool) {
	if o == nil || o.ActionName == nil {
		var ret string
		return ret, false
	}
	return *o.ActionName, true
}

// HasActionName returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceOperation) HasActionName() bool {
	if o != nil && o.ActionName != nil {
		return true
	}

	return false
}

// SetActionName gets a reference to the given string and assigns it to the ActionName field.
func (o *MicrosoftGraphResourceOperation) SetActionName(v string) {
	o.ActionName = &v
}

// SetActionNameExplicitNull (un)sets ActionName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ActionName value is set to nil even if false is passed
func (o *MicrosoftGraphResourceOperation) SetActionNameExplicitNull(b bool) {
	o.ActionName = nil
	o.isExplicitNullActionName = b
}
// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphResourceOperation) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResourceOperation) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceOperation) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphResourceOperation) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphResourceOperation) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphResourceOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ResourceName == nil {
		if o.isExplicitNullResourceName {
			toSerialize["resourceName"] = o.ResourceName
		}
	} else {
		toSerialize["resourceName"] = o.ResourceName
	}
	if o.ActionName == nil {
		if o.isExplicitNullActionName {
			toSerialize["actionName"] = o.ActionName
		}
	} else {
		toSerialize["actionName"] = o.ActionName
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

