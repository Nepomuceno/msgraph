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
// MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary struct for MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary
type MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary struct {
	Id *string `json:"id,omitempty"`

	// Website url
	Url *string `json:"url,omitempty"`
	isExplicitNullUrl bool `json:"-"`
	// Device Count
	DeviceCount *int32 `json:"deviceCount,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) GetUrlOk() (string, bool) {
	if o == nil || o.Url == nil {
		var ret string
		return ret, false
	}
	return *o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) SetUrl(v string) {
	o.Url = &v
}

// SetUrlExplicitNull (un)sets Url to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Url value is set to nil even if false is passed
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) SetUrlExplicitNull(b bool) {
	o.Url = nil
	o.isExplicitNullUrl = b
}
// GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) GetDeviceCount() int32 {
	if o == nil || o.DeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) GetDeviceCountOk() (int32, bool) {
	if o == nil || o.DeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) HasDeviceCount() bool {
	if o != nil && o.DeviceCount != nil {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url == nil {
		if o.isExplicitNullUrl {
			toSerialize["url"] = o.Url
		}
	} else {
		toSerialize["url"] = o.Url
	}
	if o.DeviceCount != nil {
		toSerialize["deviceCount"] = o.DeviceCount
	}
	return json.Marshal(toSerialize)
}

