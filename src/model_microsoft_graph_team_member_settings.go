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
// MicrosoftGraphTeamMemberSettings struct for MicrosoftGraphTeamMemberSettings
type MicrosoftGraphTeamMemberSettings struct {
	AllowCreateUpdateChannels *bool `json:"allowCreateUpdateChannels,omitempty"`
	isExplicitNullAllowCreateUpdateChannels bool `json:"-"`
	AllowDeleteChannels *bool `json:"allowDeleteChannels,omitempty"`
	isExplicitNullAllowDeleteChannels bool `json:"-"`
	AllowAddRemoveApps *bool `json:"allowAddRemoveApps,omitempty"`
	isExplicitNullAllowAddRemoveApps bool `json:"-"`
	AllowCreateUpdateRemoveTabs *bool `json:"allowCreateUpdateRemoveTabs,omitempty"`
	isExplicitNullAllowCreateUpdateRemoveTabs bool `json:"-"`
	AllowCreateUpdateRemoveConnectors *bool `json:"allowCreateUpdateRemoveConnectors,omitempty"`
	isExplicitNullAllowCreateUpdateRemoveConnectors bool `json:"-"`
}

// GetAllowCreateUpdateChannels returns the AllowCreateUpdateChannels field if non-nil, zero value otherwise.
func (o *MicrosoftGraphTeamMemberSettings) GetAllowCreateUpdateChannels() bool {
	if o == nil || o.AllowCreateUpdateChannels == nil {
		var ret bool
		return ret
	}
	return *o.AllowCreateUpdateChannels
}

// GetAllowCreateUpdateChannelsOk returns a tuple with the AllowCreateUpdateChannels field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamMemberSettings) GetAllowCreateUpdateChannelsOk() (bool, bool) {
	if o == nil || o.AllowCreateUpdateChannels == nil {
		var ret bool
		return ret, false
	}
	return *o.AllowCreateUpdateChannels, true
}

// HasAllowCreateUpdateChannels returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamMemberSettings) HasAllowCreateUpdateChannels() bool {
	if o != nil && o.AllowCreateUpdateChannels != nil {
		return true
	}

	return false
}

// SetAllowCreateUpdateChannels gets a reference to the given bool and assigns it to the AllowCreateUpdateChannels field.
func (o *MicrosoftGraphTeamMemberSettings) SetAllowCreateUpdateChannels(v bool) {
	o.AllowCreateUpdateChannels = &v
}

// SetAllowCreateUpdateChannelsExplicitNull (un)sets AllowCreateUpdateChannels to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AllowCreateUpdateChannels value is set to nil even if false is passed
func (o *MicrosoftGraphTeamMemberSettings) SetAllowCreateUpdateChannelsExplicitNull(b bool) {
	o.AllowCreateUpdateChannels = nil
	o.isExplicitNullAllowCreateUpdateChannels = b
}
// GetAllowDeleteChannels returns the AllowDeleteChannels field if non-nil, zero value otherwise.
func (o *MicrosoftGraphTeamMemberSettings) GetAllowDeleteChannels() bool {
	if o == nil || o.AllowDeleteChannels == nil {
		var ret bool
		return ret
	}
	return *o.AllowDeleteChannels
}

// GetAllowDeleteChannelsOk returns a tuple with the AllowDeleteChannels field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamMemberSettings) GetAllowDeleteChannelsOk() (bool, bool) {
	if o == nil || o.AllowDeleteChannels == nil {
		var ret bool
		return ret, false
	}
	return *o.AllowDeleteChannels, true
}

// HasAllowDeleteChannels returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamMemberSettings) HasAllowDeleteChannels() bool {
	if o != nil && o.AllowDeleteChannels != nil {
		return true
	}

	return false
}

// SetAllowDeleteChannels gets a reference to the given bool and assigns it to the AllowDeleteChannels field.
func (o *MicrosoftGraphTeamMemberSettings) SetAllowDeleteChannels(v bool) {
	o.AllowDeleteChannels = &v
}

// SetAllowDeleteChannelsExplicitNull (un)sets AllowDeleteChannels to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AllowDeleteChannels value is set to nil even if false is passed
func (o *MicrosoftGraphTeamMemberSettings) SetAllowDeleteChannelsExplicitNull(b bool) {
	o.AllowDeleteChannels = nil
	o.isExplicitNullAllowDeleteChannels = b
}
// GetAllowAddRemoveApps returns the AllowAddRemoveApps field if non-nil, zero value otherwise.
func (o *MicrosoftGraphTeamMemberSettings) GetAllowAddRemoveApps() bool {
	if o == nil || o.AllowAddRemoveApps == nil {
		var ret bool
		return ret
	}
	return *o.AllowAddRemoveApps
}

// GetAllowAddRemoveAppsOk returns a tuple with the AllowAddRemoveApps field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamMemberSettings) GetAllowAddRemoveAppsOk() (bool, bool) {
	if o == nil || o.AllowAddRemoveApps == nil {
		var ret bool
		return ret, false
	}
	return *o.AllowAddRemoveApps, true
}

// HasAllowAddRemoveApps returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamMemberSettings) HasAllowAddRemoveApps() bool {
	if o != nil && o.AllowAddRemoveApps != nil {
		return true
	}

	return false
}

// SetAllowAddRemoveApps gets a reference to the given bool and assigns it to the AllowAddRemoveApps field.
func (o *MicrosoftGraphTeamMemberSettings) SetAllowAddRemoveApps(v bool) {
	o.AllowAddRemoveApps = &v
}

// SetAllowAddRemoveAppsExplicitNull (un)sets AllowAddRemoveApps to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AllowAddRemoveApps value is set to nil even if false is passed
func (o *MicrosoftGraphTeamMemberSettings) SetAllowAddRemoveAppsExplicitNull(b bool) {
	o.AllowAddRemoveApps = nil
	o.isExplicitNullAllowAddRemoveApps = b
}
// GetAllowCreateUpdateRemoveTabs returns the AllowCreateUpdateRemoveTabs field if non-nil, zero value otherwise.
func (o *MicrosoftGraphTeamMemberSettings) GetAllowCreateUpdateRemoveTabs() bool {
	if o == nil || o.AllowCreateUpdateRemoveTabs == nil {
		var ret bool
		return ret
	}
	return *o.AllowCreateUpdateRemoveTabs
}

// GetAllowCreateUpdateRemoveTabsOk returns a tuple with the AllowCreateUpdateRemoveTabs field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamMemberSettings) GetAllowCreateUpdateRemoveTabsOk() (bool, bool) {
	if o == nil || o.AllowCreateUpdateRemoveTabs == nil {
		var ret bool
		return ret, false
	}
	return *o.AllowCreateUpdateRemoveTabs, true
}

// HasAllowCreateUpdateRemoveTabs returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamMemberSettings) HasAllowCreateUpdateRemoveTabs() bool {
	if o != nil && o.AllowCreateUpdateRemoveTabs != nil {
		return true
	}

	return false
}

// SetAllowCreateUpdateRemoveTabs gets a reference to the given bool and assigns it to the AllowCreateUpdateRemoveTabs field.
func (o *MicrosoftGraphTeamMemberSettings) SetAllowCreateUpdateRemoveTabs(v bool) {
	o.AllowCreateUpdateRemoveTabs = &v
}

// SetAllowCreateUpdateRemoveTabsExplicitNull (un)sets AllowCreateUpdateRemoveTabs to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AllowCreateUpdateRemoveTabs value is set to nil even if false is passed
func (o *MicrosoftGraphTeamMemberSettings) SetAllowCreateUpdateRemoveTabsExplicitNull(b bool) {
	o.AllowCreateUpdateRemoveTabs = nil
	o.isExplicitNullAllowCreateUpdateRemoveTabs = b
}
// GetAllowCreateUpdateRemoveConnectors returns the AllowCreateUpdateRemoveConnectors field if non-nil, zero value otherwise.
func (o *MicrosoftGraphTeamMemberSettings) GetAllowCreateUpdateRemoveConnectors() bool {
	if o == nil || o.AllowCreateUpdateRemoveConnectors == nil {
		var ret bool
		return ret
	}
	return *o.AllowCreateUpdateRemoveConnectors
}

// GetAllowCreateUpdateRemoveConnectorsOk returns a tuple with the AllowCreateUpdateRemoveConnectors field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamMemberSettings) GetAllowCreateUpdateRemoveConnectorsOk() (bool, bool) {
	if o == nil || o.AllowCreateUpdateRemoveConnectors == nil {
		var ret bool
		return ret, false
	}
	return *o.AllowCreateUpdateRemoveConnectors, true
}

// HasAllowCreateUpdateRemoveConnectors returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamMemberSettings) HasAllowCreateUpdateRemoveConnectors() bool {
	if o != nil && o.AllowCreateUpdateRemoveConnectors != nil {
		return true
	}

	return false
}

// SetAllowCreateUpdateRemoveConnectors gets a reference to the given bool and assigns it to the AllowCreateUpdateRemoveConnectors field.
func (o *MicrosoftGraphTeamMemberSettings) SetAllowCreateUpdateRemoveConnectors(v bool) {
	o.AllowCreateUpdateRemoveConnectors = &v
}

// SetAllowCreateUpdateRemoveConnectorsExplicitNull (un)sets AllowCreateUpdateRemoveConnectors to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AllowCreateUpdateRemoveConnectors value is set to nil even if false is passed
func (o *MicrosoftGraphTeamMemberSettings) SetAllowCreateUpdateRemoveConnectorsExplicitNull(b bool) {
	o.AllowCreateUpdateRemoveConnectors = nil
	o.isExplicitNullAllowCreateUpdateRemoveConnectors = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphTeamMemberSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowCreateUpdateChannels == nil {
		if o.isExplicitNullAllowCreateUpdateChannels {
			toSerialize["allowCreateUpdateChannels"] = o.AllowCreateUpdateChannels
		}
	} else {
		toSerialize["allowCreateUpdateChannels"] = o.AllowCreateUpdateChannels
	}
	if o.AllowDeleteChannels == nil {
		if o.isExplicitNullAllowDeleteChannels {
			toSerialize["allowDeleteChannels"] = o.AllowDeleteChannels
		}
	} else {
		toSerialize["allowDeleteChannels"] = o.AllowDeleteChannels
	}
	if o.AllowAddRemoveApps == nil {
		if o.isExplicitNullAllowAddRemoveApps {
			toSerialize["allowAddRemoveApps"] = o.AllowAddRemoveApps
		}
	} else {
		toSerialize["allowAddRemoveApps"] = o.AllowAddRemoveApps
	}
	if o.AllowCreateUpdateRemoveTabs == nil {
		if o.isExplicitNullAllowCreateUpdateRemoveTabs {
			toSerialize["allowCreateUpdateRemoveTabs"] = o.AllowCreateUpdateRemoveTabs
		}
	} else {
		toSerialize["allowCreateUpdateRemoveTabs"] = o.AllowCreateUpdateRemoveTabs
	}
	if o.AllowCreateUpdateRemoveConnectors == nil {
		if o.isExplicitNullAllowCreateUpdateRemoveConnectors {
			toSerialize["allowCreateUpdateRemoveConnectors"] = o.AllowCreateUpdateRemoveConnectors
		}
	} else {
		toSerialize["allowCreateUpdateRemoveConnectors"] = o.AllowCreateUpdateRemoveConnectors
	}
	return json.Marshal(toSerialize)
}

