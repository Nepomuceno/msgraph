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
// MicrosoftGraphOmaSettingStringXml struct for MicrosoftGraphOmaSettingStringXml
type MicrosoftGraphOmaSettingStringXml struct {
	// Display Name.
	DisplayName *string `json:"displayName,omitempty"`

	// Description.
	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
	// OMA.
	OmaUri *string `json:"omaUri,omitempty"`

	// File name associated with the Value property (*.xml).
	FileName *string `json:"fileName,omitempty"`
	isExplicitNullFileName bool `json:"-"`
	// Value. (UTF8 encoded byte array)
	Value *string `json:"value,omitempty"`

}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOmaSettingStringXml) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOmaSettingStringXml) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphOmaSettingStringXml) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphOmaSettingStringXml) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOmaSettingStringXml) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOmaSettingStringXml) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphOmaSettingStringXml) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphOmaSettingStringXml) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphOmaSettingStringXml) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetOmaUri returns the OmaUri field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOmaSettingStringXml) GetOmaUri() string {
	if o == nil || o.OmaUri == nil {
		var ret string
		return ret
	}
	return *o.OmaUri
}

// GetOmaUriOk returns a tuple with the OmaUri field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOmaSettingStringXml) GetOmaUriOk() (string, bool) {
	if o == nil || o.OmaUri == nil {
		var ret string
		return ret, false
	}
	return *o.OmaUri, true
}

// HasOmaUri returns a boolean if a field has been set.
func (o *MicrosoftGraphOmaSettingStringXml) HasOmaUri() bool {
	if o != nil && o.OmaUri != nil {
		return true
	}

	return false
}

// SetOmaUri gets a reference to the given string and assigns it to the OmaUri field.
func (o *MicrosoftGraphOmaSettingStringXml) SetOmaUri(v string) {
	o.OmaUri = &v
}

// GetFileName returns the FileName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOmaSettingStringXml) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOmaSettingStringXml) GetFileNameOk() (string, bool) {
	if o == nil || o.FileName == nil {
		var ret string
		return ret, false
	}
	return *o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *MicrosoftGraphOmaSettingStringXml) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *MicrosoftGraphOmaSettingStringXml) SetFileName(v string) {
	o.FileName = &v
}

// SetFileNameExplicitNull (un)sets FileName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The FileName value is set to nil even if false is passed
func (o *MicrosoftGraphOmaSettingStringXml) SetFileNameExplicitNull(b bool) {
	o.FileName = nil
	o.isExplicitNullFileName = b
}
// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOmaSettingStringXml) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOmaSettingStringXml) GetValueOk() (string, bool) {
	if o == nil || o.Value == nil {
		var ret string
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MicrosoftGraphOmaSettingStringXml) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *MicrosoftGraphOmaSettingStringXml) SetValue(v string) {
	o.Value = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphOmaSettingStringXml) MarshalJSON() ([]byte, error) {
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
	if o.FileName == nil {
		if o.isExplicitNullFileName {
			toSerialize["fileName"] = o.FileName
		}
	} else {
		toSerialize["fileName"] = o.FileName
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}


