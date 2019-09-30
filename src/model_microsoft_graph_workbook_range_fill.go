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
// MicrosoftGraphWorkbookRangeFill struct for MicrosoftGraphWorkbookRangeFill
type MicrosoftGraphWorkbookRangeFill struct {
	Id *string `json:"id,omitempty"`

	Color *string `json:"color,omitempty"`
	isExplicitNullColor bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookRangeFill) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookRangeFill) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRangeFill) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookRangeFill) SetId(v string) {
	o.Id = &v
}

// GetColor returns the Color field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookRangeFill) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookRangeFill) GetColorOk() (string, bool) {
	if o == nil || o.Color == nil {
		var ret string
		return ret, false
	}
	return *o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookRangeFill) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *MicrosoftGraphWorkbookRangeFill) SetColor(v string) {
	o.Color = &v
}

// SetColorExplicitNull (un)sets Color to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Color value is set to nil even if false is passed
func (o *MicrosoftGraphWorkbookRangeFill) SetColorExplicitNull(b bool) {
	o.Color = nil
	o.isExplicitNullColor = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphWorkbookRangeFill) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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

