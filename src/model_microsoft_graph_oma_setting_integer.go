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
// MicrosoftGraphOmaSettingInteger struct for MicrosoftGraphOmaSettingInteger
type MicrosoftGraphOmaSettingInteger struct {
	// Display Name.
	DisplayName *string `json:"displayName,omitempty"`

	// Description.
	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
	// OMA.
	OmaUri *string `json:"omaUri,omitempty"`

	// Value.
	Value *int32 `json:"value,omitempty"`

}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOmaSettingInteger) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOmaSettingInteger) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphOmaSettingInteger) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphOmaSettingInteger) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOmaSettingInteger) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOmaSettingInteger) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphOmaSettingInteger) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphOmaSettingInteger) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphOmaSettingInteger) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetOmaUri returns the OmaUri field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOmaSettingInteger) GetOmaUri() string {
	if o == nil || o.OmaUri == nil {
		var ret string
		return ret
	}
	return *o.OmaUri
}

// GetOmaUriOk returns a tuple with the OmaUri field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOmaSettingInteger) GetOmaUriOk() (string, bool) {
	if o == nil || o.OmaUri == nil {
		var ret string
		return ret, false
	}
	return *o.OmaUri, true
}

// HasOmaUri returns a boolean if a field has been set.
func (o *MicrosoftGraphOmaSettingInteger) HasOmaUri() bool {
	if o != nil && o.OmaUri != nil {
		return true
	}

	return false
}

// SetOmaUri gets a reference to the given string and assigns it to the OmaUri field.
func (o *MicrosoftGraphOmaSettingInteger) SetOmaUri(v string) {
	o.OmaUri = &v
}

// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOmaSettingInteger) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOmaSettingInteger) GetValueOk() (int32, bool) {
	if o == nil || o.Value == nil {
		var ret int32
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MicrosoftGraphOmaSettingInteger) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *MicrosoftGraphOmaSettingInteger) SetValue(v int32) {
	o.Value = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphOmaSettingInteger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description == nil {
		if o.isExplicitNullDescription {
			toSerialize["description"] = o.Description
		}
	} else {
		toSerialize["description"] = o.Description
	}
	if o.OmaUri != nil {
		toSerialize["omaUri"] = o.OmaUri
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

