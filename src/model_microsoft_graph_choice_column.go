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
// MicrosoftGraphChoiceColumn struct for MicrosoftGraphChoiceColumn
type MicrosoftGraphChoiceColumn struct {
	AllowTextEntry *bool `json:"allowTextEntry,omitempty"`
	isExplicitNullAllowTextEntry bool `json:"-"`
	Choices *[]string `json:"choices,omitempty"`

	DisplayAs *string `json:"displayAs,omitempty"`
	isExplicitNullDisplayAs bool `json:"-"`
}

// GetAllowTextEntry returns the AllowTextEntry field if non-nil, zero value otherwise.
func (o *MicrosoftGraphChoiceColumn) GetAllowTextEntry() bool {
	if o == nil || o.AllowTextEntry == nil {
		var ret bool
		return ret
	}
	return *o.AllowTextEntry
}

// GetAllowTextEntryOk returns a tuple with the AllowTextEntry field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphChoiceColumn) GetAllowTextEntryOk() (bool, bool) {
	if o == nil || o.AllowTextEntry == nil {
		var ret bool
		return ret, false
	}
	return *o.AllowTextEntry, true
}

// HasAllowTextEntry returns a boolean if a field has been set.
func (o *MicrosoftGraphChoiceColumn) HasAllowTextEntry() bool {
	if o != nil && o.AllowTextEntry != nil {
		return true
	}

	return false
}

// SetAllowTextEntry gets a reference to the given bool and assigns it to the AllowTextEntry field.
func (o *MicrosoftGraphChoiceColumn) SetAllowTextEntry(v bool) {
	o.AllowTextEntry = &v
}

// SetAllowTextEntryExplicitNull (un)sets AllowTextEntry to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AllowTextEntry value is set to nil even if false is passed
func (o *MicrosoftGraphChoiceColumn) SetAllowTextEntryExplicitNull(b bool) {
	o.AllowTextEntry = nil
	o.isExplicitNullAllowTextEntry = b
}
// GetChoices returns the Choices field if non-nil, zero value otherwise.
func (o *MicrosoftGraphChoiceColumn) GetChoices() []string {
	if o == nil || o.Choices == nil {
		var ret []string
		return ret
	}
	return *o.Choices
}

// GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphChoiceColumn) GetChoicesOk() ([]string, bool) {
	if o == nil || o.Choices == nil {
		var ret []string
		return ret, false
	}
	return *o.Choices, true
}

// HasChoices returns a boolean if a field has been set.
func (o *MicrosoftGraphChoiceColumn) HasChoices() bool {
	if o != nil && o.Choices != nil {
		return true
	}

	return false
}

// SetChoices gets a reference to the given []string and assigns it to the Choices field.
func (o *MicrosoftGraphChoiceColumn) SetChoices(v []string) {
	o.Choices = &v
}

// GetDisplayAs returns the DisplayAs field if non-nil, zero value otherwise.
func (o *MicrosoftGraphChoiceColumn) GetDisplayAs() string {
	if o == nil || o.DisplayAs == nil {
		var ret string
		return ret
	}
	return *o.DisplayAs
}

// GetDisplayAsOk returns a tuple with the DisplayAs field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphChoiceColumn) GetDisplayAsOk() (string, bool) {
	if o == nil || o.DisplayAs == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayAs, true
}

// HasDisplayAs returns a boolean if a field has been set.
func (o *MicrosoftGraphChoiceColumn) HasDisplayAs() bool {
	if o != nil && o.DisplayAs != nil {
		return true
	}

	return false
}

// SetDisplayAs gets a reference to the given string and assigns it to the DisplayAs field.
func (o *MicrosoftGraphChoiceColumn) SetDisplayAs(v string) {
	o.DisplayAs = &v
}

// SetDisplayAsExplicitNull (un)sets DisplayAs to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayAs value is set to nil even if false is passed
func (o *MicrosoftGraphChoiceColumn) SetDisplayAsExplicitNull(b bool) {
	o.DisplayAs = nil
	o.isExplicitNullDisplayAs = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphChoiceColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowTextEntry == nil {
		if o.isExplicitNullAllowTextEntry {
			toSerialize["allowTextEntry"] = o.AllowTextEntry
		}
	} else {
		toSerialize["allowTextEntry"] = o.AllowTextEntry
	}
	if o.Choices != nil {
		toSerialize["choices"] = o.Choices
	}
	if o.DisplayAs == nil {
		if o.isExplicitNullDisplayAs {
			toSerialize["displayAs"] = o.DisplayAs
		}
	} else {
		toSerialize["displayAs"] = o.DisplayAs
	}
	return json.Marshal(toSerialize)
}


