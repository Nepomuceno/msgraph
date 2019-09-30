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
// MicrosoftGraphWorkbookChartAxisTitleFormat struct for MicrosoftGraphWorkbookChartAxisTitleFormat
type MicrosoftGraphWorkbookChartAxisTitleFormat struct {
	Id *string `json:"id,omitempty"`

	Font *AnyOfmicrosoftGraphWorkbookChartFont `json:"font,omitempty"`
	isExplicitNullFont bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) SetId(v string) {
	o.Id = &v
}

// GetFont returns the Font field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) GetFont() AnyOfmicrosoftGraphWorkbookChartFont {
	if o == nil || o.Font == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartFont
		return ret
	}
	return *o.Font
}

// GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) GetFontOk() (AnyOfmicrosoftGraphWorkbookChartFont, bool) {
	if o == nil || o.Font == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartFont
		return ret, false
	}
	return *o.Font, true
}

// HasFont returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) HasFont() bool {
	if o != nil && o.Font != nil {
		return true
	}

	return false
}

// SetFont gets a reference to the given AnyOfmicrosoftGraphWorkbookChartFont and assigns it to the Font field.
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) SetFont(v AnyOfmicrosoftGraphWorkbookChartFont) {
	o.Font = &v
}

// SetFontExplicitNull (un)sets Font to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Font value is set to nil even if false is passed
func (o *MicrosoftGraphWorkbookChartAxisTitleFormat) SetFontExplicitNull(b bool) {
	o.Font = nil
	o.isExplicitNullFont = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphWorkbookChartAxisTitleFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Font == nil {
		if o.isExplicitNullFont {
			toSerialize["font"] = o.Font
		}
	} else {
		toSerialize["font"] = o.Font
	}
	return json.Marshal(toSerialize)
}


