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
// MicrosoftGraphIosNetworkUsageRule struct for MicrosoftGraphIosNetworkUsageRule
type MicrosoftGraphIosNetworkUsageRule struct {
	// Information about the managed apps that this rule is going to apply to. This collection can contain a maximum of 500 elements.
	ManagedApps *[]AnyOfmicrosoftGraphAppListItem `json:"managedApps,omitempty"`

	// If set to true, corresponding managed apps will not be allowed to use cellular data when roaming.
	CellularDataBlockWhenRoaming *bool `json:"cellularDataBlockWhenRoaming,omitempty"`

	// If set to true, corresponding managed apps will not be allowed to use cellular data at any time.
	CellularDataBlocked *bool `json:"cellularDataBlocked,omitempty"`

}

// GetManagedApps returns the ManagedApps field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosNetworkUsageRule) GetManagedApps() []AnyOfmicrosoftGraphAppListItem {
	if o == nil || o.ManagedApps == nil {
		var ret []AnyOfmicrosoftGraphAppListItem
		return ret
	}
	return *o.ManagedApps
}

// GetManagedAppsOk returns a tuple with the ManagedApps field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosNetworkUsageRule) GetManagedAppsOk() ([]AnyOfmicrosoftGraphAppListItem, bool) {
	if o == nil || o.ManagedApps == nil {
		var ret []AnyOfmicrosoftGraphAppListItem
		return ret, false
	}
	return *o.ManagedApps, true
}

// HasManagedApps returns a boolean if a field has been set.
func (o *MicrosoftGraphIosNetworkUsageRule) HasManagedApps() bool {
	if o != nil && o.ManagedApps != nil {
		return true
	}

	return false
}

// SetManagedApps gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the ManagedApps field.
func (o *MicrosoftGraphIosNetworkUsageRule) SetManagedApps(v []AnyOfmicrosoftGraphAppListItem) {
	o.ManagedApps = &v
}

// GetCellularDataBlockWhenRoaming returns the CellularDataBlockWhenRoaming field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosNetworkUsageRule) GetCellularDataBlockWhenRoaming() bool {
	if o == nil || o.CellularDataBlockWhenRoaming == nil {
		var ret bool
		return ret
	}
	return *o.CellularDataBlockWhenRoaming
}

// GetCellularDataBlockWhenRoamingOk returns a tuple with the CellularDataBlockWhenRoaming field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosNetworkUsageRule) GetCellularDataBlockWhenRoamingOk() (bool, bool) {
	if o == nil || o.CellularDataBlockWhenRoaming == nil {
		var ret bool
		return ret, false
	}
	return *o.CellularDataBlockWhenRoaming, true
}

// HasCellularDataBlockWhenRoaming returns a boolean if a field has been set.
func (o *MicrosoftGraphIosNetworkUsageRule) HasCellularDataBlockWhenRoaming() bool {
	if o != nil && o.CellularDataBlockWhenRoaming != nil {
		return true
	}

	return false
}

// SetCellularDataBlockWhenRoaming gets a reference to the given bool and assigns it to the CellularDataBlockWhenRoaming field.
func (o *MicrosoftGraphIosNetworkUsageRule) SetCellularDataBlockWhenRoaming(v bool) {
	o.CellularDataBlockWhenRoaming = &v
}

// GetCellularDataBlocked returns the CellularDataBlocked field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosNetworkUsageRule) GetCellularDataBlocked() bool {
	if o == nil || o.CellularDataBlocked == nil {
		var ret bool
		return ret
	}
	return *o.CellularDataBlocked
}

// GetCellularDataBlockedOk returns a tuple with the CellularDataBlocked field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosNetworkUsageRule) GetCellularDataBlockedOk() (bool, bool) {
	if o == nil || o.CellularDataBlocked == nil {
		var ret bool
		return ret, false
	}
	return *o.CellularDataBlocked, true
}

// HasCellularDataBlocked returns a boolean if a field has been set.
func (o *MicrosoftGraphIosNetworkUsageRule) HasCellularDataBlocked() bool {
	if o != nil && o.CellularDataBlocked != nil {
		return true
	}

	return false
}

// SetCellularDataBlocked gets a reference to the given bool and assigns it to the CellularDataBlocked field.
func (o *MicrosoftGraphIosNetworkUsageRule) SetCellularDataBlocked(v bool) {
	o.CellularDataBlocked = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphIosNetworkUsageRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ManagedApps != nil {
		toSerialize["managedApps"] = o.ManagedApps
	}
	if o.CellularDataBlockWhenRoaming != nil {
		toSerialize["cellularDataBlockWhenRoaming"] = o.CellularDataBlockWhenRoaming
	}
	if o.CellularDataBlocked != nil {
		toSerialize["cellularDataBlocked"] = o.CellularDataBlocked
	}
	return json.Marshal(toSerialize)
}


