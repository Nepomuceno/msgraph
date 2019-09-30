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
// MicrosoftGraphOutlookCategory struct for MicrosoftGraphOutlookCategory
type MicrosoftGraphOutlookCategory struct {
	Id *string `json:"id,omitempty"`

	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	Color *AnyOfmicrosoftGraphCategoryColor `json:"color,omitempty"`
	isExplicitNullColor bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookCategory) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookCategory) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookCategory) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOutlookCategory) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookCategory) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookCategory) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookCategory) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphOutlookCategory) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphOutlookCategory) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetColor returns the Color field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookCategory) GetColor() AnyOfmicrosoftGraphCategoryColor {
	if o == nil || o.Color == nil {
		var ret AnyOfmicrosoftGraphCategoryColor
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookCategory) GetColorOk() (AnyOfmicrosoftGraphCategoryColor, bool) {
	if o == nil || o.Color == nil {
		var ret AnyOfmicrosoftGraphCategoryColor
		return ret, false
	}
	return *o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookCategory) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given AnyOfmicrosoftGraphCategoryColor and assigns it to the Color field.
func (o *MicrosoftGraphOutlookCategory) SetColor(v AnyOfmicrosoftGraphCategoryColor) {
	o.Color = &v
}

// SetColorExplicitNull (un)sets Color to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Color value is set to nil even if false is passed
func (o *MicrosoftGraphOutlookCategory) SetColorExplicitNull(b bool) {
	o.Color = nil
	o.isExplicitNullColor = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphOutlookCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Color == nil {
		if o.isExplicitNullColor {
			toSerialize["color"] = o.Color
		}
	} else {
		toSerialize["color"] = o.Color
	}
	return json.Marshal(toSerialize)
}

