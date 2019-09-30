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
// MicrosoftGraphConfigurationManagerClientEnabledFeatures struct for MicrosoftGraphConfigurationManagerClientEnabledFeatures
type MicrosoftGraphConfigurationManagerClientEnabledFeatures struct {
	// Whether inventory is managed by Intune
	Inventory *bool `json:"inventory,omitempty"`

	// Whether modern application is managed by Intune
	ModernApps *bool `json:"modernApps,omitempty"`

	// Whether resource access is managed by Intune
	ResourceAccess *bool `json:"resourceAccess,omitempty"`

	// Whether device configuration is managed by Intune
	DeviceConfiguration *bool `json:"deviceConfiguration,omitempty"`

	// Whether compliance policy is managed by Intune
	CompliancePolicy *bool `json:"compliancePolicy,omitempty"`

	// Whether Windows Update for Business is managed by Intune
	WindowsUpdateForBusiness *bool `json:"windowsUpdateForBusiness,omitempty"`

}

// GetInventory returns the Inventory field if non-nil, zero value otherwise.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) GetInventory() bool {
	if o == nil || o.Inventory == nil {
		var ret bool
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) GetInventoryOk() (bool, bool) {
	if o == nil || o.Inventory == nil {
		var ret bool
		return ret, false
	}
	return *o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given bool and assigns it to the Inventory field.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) SetInventory(v bool) {
	o.Inventory = &v
}

// GetModernApps returns the ModernApps field if non-nil, zero value otherwise.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) GetModernApps() bool {
	if o == nil || o.ModernApps == nil {
		var ret bool
		return ret
	}
	return *o.ModernApps
}

// GetModernAppsOk returns a tuple with the ModernApps field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) GetModernAppsOk() (bool, bool) {
	if o == nil || o.ModernApps == nil {
		var ret bool
		return ret, false
	}
	return *o.ModernApps, true
}

// HasModernApps returns a boolean if a field has been set.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) HasModernApps() bool {
	if o != nil && o.ModernApps != nil {
		return true
	}

	return false
}

// SetModernApps gets a reference to the given bool and assigns it to the ModernApps field.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) SetModernApps(v bool) {
	o.ModernApps = &v
}

// GetResourceAccess returns the ResourceAccess field if non-nil, zero value otherwise.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) GetResourceAccess() bool {
	if o == nil || o.ResourceAccess == nil {
		var ret bool
		return ret
	}
	return *o.ResourceAccess
}

// GetResourceAccessOk returns a tuple with the ResourceAccess field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) GetResourceAccessOk() (bool, bool) {
	if o == nil || o.ResourceAccess == nil {
		var ret bool
		return ret, false
	}
	return *o.ResourceAccess, true
}

// HasResourceAccess returns a boolean if a field has been set.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) HasResourceAccess() bool {
	if o != nil && o.ResourceAccess != nil {
		return true
	}

	return false
}

// SetResourceAccess gets a reference to the given bool and assigns it to the ResourceAccess field.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) SetResourceAccess(v bool) {
	o.ResourceAccess = &v
}

// GetDeviceConfiguration returns the DeviceConfiguration field if non-nil, zero value otherwise.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) GetDeviceConfiguration() bool {
	if o == nil || o.DeviceConfiguration == nil {
		var ret bool
		return ret
	}
	return *o.DeviceConfiguration
}

// GetDeviceConfigurationOk returns a tuple with the DeviceConfiguration field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) GetDeviceConfigurationOk() (bool, bool) {
	if o == nil || o.DeviceConfiguration == nil {
		var ret bool
		return ret, false
	}
	return *o.DeviceConfiguration, true
}

// HasDeviceConfiguration returns a boolean if a field has been set.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) HasDeviceConfiguration() bool {
	if o != nil && o.DeviceConfiguration != nil {
		return true
	}

	return false
}

// SetDeviceConfiguration gets a reference to the given bool and assigns it to the DeviceConfiguration field.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) SetDeviceConfiguration(v bool) {
	o.DeviceConfiguration = &v
}

// GetCompliancePolicy returns the CompliancePolicy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) GetCompliancePolicy() bool {
	if o == nil || o.CompliancePolicy == nil {
		var ret bool
		return ret
	}
	return *o.CompliancePolicy
}

// GetCompliancePolicyOk returns a tuple with the CompliancePolicy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) GetCompliancePolicyOk() (bool, bool) {
	if o == nil || o.CompliancePolicy == nil {
		var ret bool
		return ret, false
	}
	return *o.CompliancePolicy, true
}

// HasCompliancePolicy returns a boolean if a field has been set.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) HasCompliancePolicy() bool {
	if o != nil && o.CompliancePolicy != nil {
		return true
	}

	return false
}

// SetCompliancePolicy gets a reference to the given bool and assigns it to the CompliancePolicy field.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) SetCompliancePolicy(v bool) {
	o.CompliancePolicy = &v
}

// GetWindowsUpdateForBusiness returns the WindowsUpdateForBusiness field if non-nil, zero value otherwise.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) GetWindowsUpdateForBusiness() bool {
	if o == nil || o.WindowsUpdateForBusiness == nil {
		var ret bool
		return ret
	}
	return *o.WindowsUpdateForBusiness
}

// GetWindowsUpdateForBusinessOk returns a tuple with the WindowsUpdateForBusiness field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) GetWindowsUpdateForBusinessOk() (bool, bool) {
	if o == nil || o.WindowsUpdateForBusiness == nil {
		var ret bool
		return ret, false
	}
	return *o.WindowsUpdateForBusiness, true
}

// HasWindowsUpdateForBusiness returns a boolean if a field has been set.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) HasWindowsUpdateForBusiness() bool {
	if o != nil && o.WindowsUpdateForBusiness != nil {
		return true
	}

	return false
}

// SetWindowsUpdateForBusiness gets a reference to the given bool and assigns it to the WindowsUpdateForBusiness field.
func (o *MicrosoftGraphConfigurationManagerClientEnabledFeatures) SetWindowsUpdateForBusiness(v bool) {
	o.WindowsUpdateForBusiness = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphConfigurationManagerClientEnabledFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Inventory != nil {
		toSerialize["inventory"] = o.Inventory
	}
	if o.ModernApps != nil {
		toSerialize["modernApps"] = o.ModernApps
	}
	if o.ResourceAccess != nil {
		toSerialize["resourceAccess"] = o.ResourceAccess
	}
	if o.DeviceConfiguration != nil {
		toSerialize["deviceConfiguration"] = o.DeviceConfiguration
	}
	if o.CompliancePolicy != nil {
		toSerialize["compliancePolicy"] = o.CompliancePolicy
	}
	if o.WindowsUpdateForBusiness != nil {
		toSerialize["windowsUpdateForBusiness"] = o.WindowsUpdateForBusiness
	}
	return json.Marshal(toSerialize)
}


