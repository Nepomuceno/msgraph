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
// InlineObject1124 struct for InlineObject1124
type InlineObject1124 struct {
	Icon *AnyOfmicrosoftGraphWorkbookIcon `json:"icon,omitempty"`
	isExplicitNullIcon bool `json:"-"`
}

// GetIcon returns the Icon field if non-nil, zero value otherwise.
func (o *InlineObject1124) GetIcon() AnyOfmicrosoftGraphWorkbookIcon {
	if o == nil || o.Icon == nil {
		var ret AnyOfmicrosoftGraphWorkbookIcon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1124) GetIconOk() (AnyOfmicrosoftGraphWorkbookIcon, bool) {
	if o == nil || o.Icon == nil {
		var ret AnyOfmicrosoftGraphWorkbookIcon
		return ret, false
	}
	return *o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *InlineObject1124) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given AnyOfmicrosoftGraphWorkbookIcon and assigns it to the Icon field.
func (o *InlineObject1124) SetIcon(v AnyOfmicrosoftGraphWorkbookIcon) {
	o.Icon = &v
}

// SetIconExplicitNull (un)sets Icon to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Icon value is set to nil even if false is passed
func (o *InlineObject1124) SetIconExplicitNull(b bool) {
	o.Icon = nil
	o.isExplicitNullIcon = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1124) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Icon == nil {
		if o.isExplicitNullIcon {
			toSerialize["icon"] = o.Icon
		}
	} else {
		toSerialize["icon"] = o.Icon
	}
	return json.Marshal(toSerialize)
}

