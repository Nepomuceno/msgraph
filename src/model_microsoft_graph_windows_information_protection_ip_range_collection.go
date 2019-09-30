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
// MicrosoftGraphWindowsInformationProtectionIpRangeCollection struct for MicrosoftGraphWindowsInformationProtectionIpRangeCollection
type MicrosoftGraphWindowsInformationProtectionIpRangeCollection struct {
	// Display name
	DisplayName *string `json:"displayName,omitempty"`

	// Collection of ip ranges
	Ranges *[]map[string]interface{} `json:"ranges,omitempty"`

}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionIpRangeCollection) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionIpRangeCollection) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionIpRangeCollection) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphWindowsInformationProtectionIpRangeCollection) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetRanges returns the Ranges field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionIpRangeCollection) GetRanges() []map[string]interface{} {
	if o == nil || o.Ranges == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionIpRangeCollection) GetRangesOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Ranges == nil {
		var ret []map[string]interface{}
		return ret, false
	}
	return *o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionIpRangeCollection) HasRanges() bool {
	if o != nil && o.Ranges != nil {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []map[string]interface{} and assigns it to the Ranges field.
func (o *MicrosoftGraphWindowsInformationProtectionIpRangeCollection) SetRanges(v []map[string]interface{}) {
	o.Ranges = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphWindowsInformationProtectionIpRangeCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Ranges != nil {
		toSerialize["ranges"] = o.Ranges
	}
	return json.Marshal(toSerialize)
}

