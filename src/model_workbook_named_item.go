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
// WorkbookNamedItem struct for WorkbookNamedItem
type WorkbookNamedItem struct {
	Comment *string `json:"comment,omitempty"`
	isExplicitNullComment bool `json:"-"`
	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
	Scope *string `json:"scope,omitempty"`

	Type *string `json:"type,omitempty"`
	isExplicitNullType bool `json:"-"`
	Value *AnyOfobject `json:"value,omitempty"`
	isExplicitNullValue bool `json:"-"`
	Visible *bool `json:"visible,omitempty"`

	Worksheet *AnyOfmicrosoftGraphWorkbookWorksheet `json:"worksheet,omitempty"`
	isExplicitNullWorksheet bool `json:"-"`
}

// GetComment returns the Comment field if non-nil, zero value otherwise.
func (o *WorkbookNamedItem) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookNamedItem) GetCommentOk() (string, bool) {
	if o == nil || o.Comment == nil {
		var ret string
		return ret, false
	}
	return *o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *WorkbookNamedItem) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *WorkbookNamedItem) SetComment(v string) {
	o.Comment = &v
}

// SetCommentExplicitNull (un)sets Comment to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Comment value is set to nil even if false is passed
func (o *WorkbookNamedItem) SetCommentExplicitNull(b bool) {
	o.Comment = nil
	o.isExplicitNullComment = b
}
// GetName returns the Name field if non-nil, zero value otherwise.
func (o *WorkbookNamedItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookNamedItem) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkbookNamedItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkbookNamedItem) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *WorkbookNamedItem) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetScope returns the Scope field if non-nil, zero value otherwise.
func (o *WorkbookNamedItem) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookNamedItem) GetScopeOk() (string, bool) {
	if o == nil || o.Scope == nil {
		var ret string
		return ret, false
	}
	return *o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *WorkbookNamedItem) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *WorkbookNamedItem) SetScope(v string) {
	o.Scope = &v
}

// GetType returns the Type field if non-nil, zero value otherwise.
func (o *WorkbookNamedItem) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookNamedItem) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkbookNamedItem) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkbookNamedItem) SetType(v string) {
	o.Type = &v
}

// SetTypeExplicitNull (un)sets Type to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Type value is set to nil even if false is passed
func (o *WorkbookNamedItem) SetTypeExplicitNull(b bool) {
	o.Type = nil
	o.isExplicitNullType = b
}
// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *WorkbookNamedItem) GetValue() AnyOfobject {
	if o == nil || o.Value == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookNamedItem) GetValueOk() (AnyOfobject, bool) {
	if o == nil || o.Value == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkbookNamedItem) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given AnyOfobject and assigns it to the Value field.
func (o *WorkbookNamedItem) SetValue(v AnyOfobject) {
	o.Value = &v
}

// SetValueExplicitNull (un)sets Value to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Value value is set to nil even if false is passed
func (o *WorkbookNamedItem) SetValueExplicitNull(b bool) {
	o.Value = nil
	o.isExplicitNullValue = b
}
// GetVisible returns the Visible field if non-nil, zero value otherwise.
func (o *WorkbookNamedItem) GetVisible() bool {
	if o == nil || o.Visible == nil {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookNamedItem) GetVisibleOk() (bool, bool) {
	if o == nil || o.Visible == nil {
		var ret bool
		return ret, false
	}
	return *o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *WorkbookNamedItem) HasVisible() bool {
	if o != nil && o.Visible != nil {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *WorkbookNamedItem) SetVisible(v bool) {
	o.Visible = &v
}

// GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.
func (o *WorkbookNamedItem) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet {
	if o == nil || o.Worksheet == nil {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret
	}
	return *o.Worksheet
}

// GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookNamedItem) GetWorksheetOk() (AnyOfmicrosoftGraphWorkbookWorksheet, bool) {
	if o == nil || o.Worksheet == nil {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret, false
	}
	return *o.Worksheet, true
}

// HasWorksheet returns a boolean if a field has been set.
func (o *WorkbookNamedItem) HasWorksheet() bool {
	if o != nil && o.Worksheet != nil {
		return true
	}

	return false
}

// SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.
func (o *WorkbookNamedItem) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet) {
	o.Worksheet = &v
}

// SetWorksheetExplicitNull (un)sets Worksheet to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Worksheet value is set to nil even if false is passed
func (o *WorkbookNamedItem) SetWorksheetExplicitNull(b bool) {
	o.Worksheet = nil
	o.isExplicitNullWorksheet = b
}

// MarshalJSON returns the JSON representation of the model.
func (o WorkbookNamedItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment == nil {
		if o.isExplicitNullComment {
			toSerialize["comment"] = o.Comment
		}
	} else {
		toSerialize["comment"] = o.Comment
	}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Type == nil {
		if o.isExplicitNullType {
			toSerialize["type"] = o.Type
		}
	} else {
		toSerialize["type"] = o.Type
	}
	if o.Value == nil {
		if o.isExplicitNullValue {
			toSerialize["value"] = o.Value
		}
	} else {
		toSerialize["value"] = o.Value
	}
	if o.Visible != nil {
		toSerialize["visible"] = o.Visible
	}
	if o.Worksheet == nil {
		if o.isExplicitNullWorksheet {
			toSerialize["worksheet"] = o.Worksheet
		}
	} else {
		toSerialize["worksheet"] = o.Worksheet
	}
	return json.Marshal(toSerialize)
}

