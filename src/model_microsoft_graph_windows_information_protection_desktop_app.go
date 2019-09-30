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
// MicrosoftGraphWindowsInformationProtectionDesktopApp struct for MicrosoftGraphWindowsInformationProtectionDesktopApp
type MicrosoftGraphWindowsInformationProtectionDesktopApp struct {
	// App display name.
	DisplayName *string `json:"displayName,omitempty"`

	// The app's description.
	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
	// The publisher name
	PublisherName *string `json:"publisherName,omitempty"`
	isExplicitNullPublisherName bool `json:"-"`
	// The product name.
	ProductName *string `json:"productName,omitempty"`
	isExplicitNullProductName bool `json:"-"`
	// If true, app is denied protection or exemption.
	Denied *bool `json:"denied,omitempty"`

	// The binary name.
	BinaryName *string `json:"binaryName,omitempty"`

	// The lower binary version.
	BinaryVersionLow *string `json:"binaryVersionLow,omitempty"`
	isExplicitNullBinaryVersionLow bool `json:"-"`
	// The high binary version.
	BinaryVersionHigh *string `json:"binaryVersionHigh,omitempty"`
	isExplicitNullBinaryVersionHigh bool `json:"-"`
}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetPublisherName returns the PublisherName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetPublisherName() string {
	if o == nil || o.PublisherName == nil {
		var ret string
		return ret
	}
	return *o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetPublisherNameOk() (string, bool) {
	if o == nil || o.PublisherName == nil {
		var ret string
		return ret, false
	}
	return *o.PublisherName, true
}

// HasPublisherName returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasPublisherName() bool {
	if o != nil && o.PublisherName != nil {
		return true
	}

	return false
}

// SetPublisherName gets a reference to the given string and assigns it to the PublisherName field.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetPublisherName(v string) {
	o.PublisherName = &v
}

// SetPublisherNameExplicitNull (un)sets PublisherName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PublisherName value is set to nil even if false is passed
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetPublisherNameExplicitNull(b bool) {
	o.PublisherName = nil
	o.isExplicitNullPublisherName = b
}
// GetProductName returns the ProductName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetProductNameOk() (string, bool) {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret, false
	}
	return *o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetProductName(v string) {
	o.ProductName = &v
}

// SetProductNameExplicitNull (un)sets ProductName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ProductName value is set to nil even if false is passed
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetProductNameExplicitNull(b bool) {
	o.ProductName = nil
	o.isExplicitNullProductName = b
}
// GetDenied returns the Denied field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetDenied() bool {
	if o == nil || o.Denied == nil {
		var ret bool
		return ret
	}
	return *o.Denied
}

// GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetDeniedOk() (bool, bool) {
	if o == nil || o.Denied == nil {
		var ret bool
		return ret, false
	}
	return *o.Denied, true
}

// HasDenied returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasDenied() bool {
	if o != nil && o.Denied != nil {
		return true
	}

	return false
}

// SetDenied gets a reference to the given bool and assigns it to the Denied field.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetDenied(v bool) {
	o.Denied = &v
}

// GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetBinaryName() string {
	if o == nil || o.BinaryName == nil {
		var ret string
		return ret
	}
	return *o.BinaryName
}

// GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetBinaryNameOk() (string, bool) {
	if o == nil || o.BinaryName == nil {
		var ret string
		return ret, false
	}
	return *o.BinaryName, true
}

// HasBinaryName returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasBinaryName() bool {
	if o != nil && o.BinaryName != nil {
		return true
	}

	return false
}

// SetBinaryName gets a reference to the given string and assigns it to the BinaryName field.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetBinaryName(v string) {
	o.BinaryName = &v
}

// GetBinaryVersionLow returns the BinaryVersionLow field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetBinaryVersionLow() string {
	if o == nil || o.BinaryVersionLow == nil {
		var ret string
		return ret
	}
	return *o.BinaryVersionLow
}

// GetBinaryVersionLowOk returns a tuple with the BinaryVersionLow field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetBinaryVersionLowOk() (string, bool) {
	if o == nil || o.BinaryVersionLow == nil {
		var ret string
		return ret, false
	}
	return *o.BinaryVersionLow, true
}

// HasBinaryVersionLow returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasBinaryVersionLow() bool {
	if o != nil && o.BinaryVersionLow != nil {
		return true
	}

	return false
}

// SetBinaryVersionLow gets a reference to the given string and assigns it to the BinaryVersionLow field.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetBinaryVersionLow(v string) {
	o.BinaryVersionLow = &v
}

// SetBinaryVersionLowExplicitNull (un)sets BinaryVersionLow to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The BinaryVersionLow value is set to nil even if false is passed
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetBinaryVersionLowExplicitNull(b bool) {
	o.BinaryVersionLow = nil
	o.isExplicitNullBinaryVersionLow = b
}
// GetBinaryVersionHigh returns the BinaryVersionHigh field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetBinaryVersionHigh() string {
	if o == nil || o.BinaryVersionHigh == nil {
		var ret string
		return ret
	}
	return *o.BinaryVersionHigh
}

// GetBinaryVersionHighOk returns a tuple with the BinaryVersionHigh field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) GetBinaryVersionHighOk() (string, bool) {
	if o == nil || o.BinaryVersionHigh == nil {
		var ret string
		return ret, false
	}
	return *o.BinaryVersionHigh, true
}

// HasBinaryVersionHigh returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) HasBinaryVersionHigh() bool {
	if o != nil && o.BinaryVersionHigh != nil {
		return true
	}

	return false
}

// SetBinaryVersionHigh gets a reference to the given string and assigns it to the BinaryVersionHigh field.
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetBinaryVersionHigh(v string) {
	o.BinaryVersionHigh = &v
}

// SetBinaryVersionHighExplicitNull (un)sets BinaryVersionHigh to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The BinaryVersionHigh value is set to nil even if false is passed
func (o *MicrosoftGraphWindowsInformationProtectionDesktopApp) SetBinaryVersionHighExplicitNull(b bool) {
	o.BinaryVersionHigh = nil
	o.isExplicitNullBinaryVersionHigh = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphWindowsInformationProtectionDesktopApp) MarshalJSON() ([]byte, error) {
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
	if o.PublisherName == nil {
		if o.isExplicitNullPublisherName {
			toSerialize["publisherName"] = o.PublisherName
		}
	} else {
		toSerialize["publisherName"] = o.PublisherName
	}
	if o.ProductName == nil {
		if o.isExplicitNullProductName {
			toSerialize["productName"] = o.ProductName
		}
	} else {
		toSerialize["productName"] = o.ProductName
	}
	if o.Denied != nil {
		toSerialize["denied"] = o.Denied
	}
	if o.BinaryName != nil {
		toSerialize["binaryName"] = o.BinaryName
	}
	if o.BinaryVersionLow == nil {
		if o.isExplicitNullBinaryVersionLow {
			toSerialize["binaryVersionLow"] = o.BinaryVersionLow
		}
	} else {
		toSerialize["binaryVersionLow"] = o.BinaryVersionLow
	}
	if o.BinaryVersionHigh == nil {
		if o.isExplicitNullBinaryVersionHigh {
			toSerialize["binaryVersionHigh"] = o.BinaryVersionHigh
		}
	} else {
		toSerialize["binaryVersionHigh"] = o.BinaryVersionHigh
	}
	return json.Marshal(toSerialize)
}

