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
// MicrosoftGraphWorkbookChartFont struct for MicrosoftGraphWorkbookChartFont
type MicrosoftGraphWorkbookChartFont struct {
	Id *string `json:"id,omitempty"`

	Bold *bool `json:"bold,omitempty"`
	isExplicitNullBold bool `json:"-"`
	Color *string `json:"color,omitempty"`
	isExplicitNullColor bool `json:"-"`
	Italic *bool `json:"italic,omitempty"`
	isExplicitNullItalic bool `json:"-"`
	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
	Size *AnyOfnumberstringstring `json:"size,omitempty"`
	isExplicitNullSize bool `json:"-"`
	Underline *string `json:"underline,omitempty"`
	isExplicitNullUnderline bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartFont) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartFont) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartFont) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookChartFont) SetId(v string) {
	o.Id = &v
}

// GetBold returns the Bold field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartFont) GetBold() bool {
	if o == nil || o.Bold == nil {
		var ret bool
		return ret
	}
	return *o.Bold
}

// GetBoldOk returns a tuple with the Bold field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartFont) GetBoldOk() (bool, bool) {
	if o == nil || o.Bold == nil {
		var ret bool
		return ret, false
	}
	return *o.Bold, true
}

// HasBold returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartFont) HasBold() bool {
	if o != nil && o.Bold != nil {
		return true
	}

	return false
}

// SetBold gets a reference to the given bool and assigns it to the Bold field.
func (o *MicrosoftGraphWorkbookChartFont) SetBold(v bool) {
	o.Bold = &v
}

// SetBoldExplicitNull (un)sets Bold to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Bold value is set to nil even if false is passed
func (o *MicrosoftGraphWorkbookChartFont) SetBoldExplicitNull(b bool) {
	o.Bold = nil
	o.isExplicitNullBold = b
}
// GetColor returns the Color field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartFont) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartFont) GetColorOk() (string, bool) {
	if o == nil || o.Color == nil {
		var ret string
		return ret, false
	}
	return *o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartFont) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *MicrosoftGraphWorkbookChartFont) SetColor(v string) {
	o.Color = &v
}

// SetColorExplicitNull (un)sets Color to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Color value is set to nil even if false is passed
func (o *MicrosoftGraphWorkbookChartFont) SetColorExplicitNull(b bool) {
	o.Color = nil
	o.isExplicitNullColor = b
}
// GetItalic returns the Italic field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartFont) GetItalic() bool {
	if o == nil || o.Italic == nil {
		var ret bool
		return ret
	}
	return *o.Italic
}

// GetItalicOk returns a tuple with the Italic field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartFont) GetItalicOk() (bool, bool) {
	if o == nil || o.Italic == nil {
		var ret bool
		return ret, false
	}
	return *o.Italic, true
}

// HasItalic returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartFont) HasItalic() bool {
	if o != nil && o.Italic != nil {
		return true
	}

	return false
}

// SetItalic gets a reference to the given bool and assigns it to the Italic field.
func (o *MicrosoftGraphWorkbookChartFont) SetItalic(v bool) {
	o.Italic = &v
}

// SetItalicExplicitNull (un)sets Italic to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Italic value is set to nil even if false is passed
func (o *MicrosoftGraphWorkbookChartFont) SetItalicExplicitNull(b bool) {
	o.Italic = nil
	o.isExplicitNullItalic = b
}
// GetName returns the Name field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartFont) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartFont) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartFont) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftGraphWorkbookChartFont) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *MicrosoftGraphWorkbookChartFont) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetSize returns the Size field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartFont) GetSize() AnyOfnumberstringstring {
	if o == nil || o.Size == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartFont) GetSizeOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Size == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartFont) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given AnyOfnumberstringstring and assigns it to the Size field.
func (o *MicrosoftGraphWorkbookChartFont) SetSize(v AnyOfnumberstringstring) {
	o.Size = &v
}

// SetSizeExplicitNull (un)sets Size to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Size value is set to nil even if false is passed
func (o *MicrosoftGraphWorkbookChartFont) SetSizeExplicitNull(b bool) {
	o.Size = nil
	o.isExplicitNullSize = b
}
// GetUnderline returns the Underline field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartFont) GetUnderline() string {
	if o == nil || o.Underline == nil {
		var ret string
		return ret
	}
	return *o.Underline
}

// GetUnderlineOk returns a tuple with the Underline field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartFont) GetUnderlineOk() (string, bool) {
	if o == nil || o.Underline == nil {
		var ret string
		return ret, false
	}
	return *o.Underline, true
}

// HasUnderline returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartFont) HasUnderline() bool {
	if o != nil && o.Underline != nil {
		return true
	}

	return false
}

// SetUnderline gets a reference to the given string and assigns it to the Underline field.
func (o *MicrosoftGraphWorkbookChartFont) SetUnderline(v string) {
	o.Underline = &v
}

// SetUnderlineExplicitNull (un)sets Underline to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Underline value is set to nil even if false is passed
func (o *MicrosoftGraphWorkbookChartFont) SetUnderlineExplicitNull(b bool) {
	o.Underline = nil
	o.isExplicitNullUnderline = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphWorkbookChartFont) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Bold == nil {
		if o.isExplicitNullBold {
			toSerialize["bold"] = o.Bold
		}
	} else {
		toSerialize["bold"] = o.Bold
	}
	if o.Color == nil {
		if o.isExplicitNullColor {
			toSerialize["color"] = o.Color
		}
	} else {
		toSerialize["color"] = o.Color
	}
	if o.Italic == nil {
		if o.isExplicitNullItalic {
			toSerialize["italic"] = o.Italic
		}
	} else {
		toSerialize["italic"] = o.Italic
	}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
	}
	if o.Size == nil {
		if o.isExplicitNullSize {
			toSerialize["size"] = o.Size
		}
	} else {
		toSerialize["size"] = o.Size
	}
	if o.Underline == nil {
		if o.isExplicitNullUnderline {
			toSerialize["underline"] = o.Underline
		}
	} else {
		toSerialize["underline"] = o.Underline
	}
	return json.Marshal(toSerialize)
}


