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
	"time"
	"encoding/json"
)
// ManagedAppPolicyDeploymentSummary The ManagedAppEntity is the base entity type for all other entity types under app management workflow.
type ManagedAppPolicyDeploymentSummary struct {
	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	ConfigurationDeployedUserCount *int32 `json:"configurationDeployedUserCount,omitempty"`

	LastRefreshTime *time.Time `json:"lastRefreshTime,omitempty"`

	ConfigurationDeploymentSummaryPerApp *[]AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp `json:"configurationDeploymentSummaryPerApp,omitempty"`

	// Version of the entity.
	Version *string `json:"version,omitempty"`
	isExplicitNullVersion bool `json:"-"`
}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *ManagedAppPolicyDeploymentSummary) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppPolicyDeploymentSummary) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ManagedAppPolicyDeploymentSummary) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ManagedAppPolicyDeploymentSummary) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *ManagedAppPolicyDeploymentSummary) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetConfigurationDeployedUserCount returns the ConfigurationDeployedUserCount field if non-nil, zero value otherwise.
func (o *ManagedAppPolicyDeploymentSummary) GetConfigurationDeployedUserCount() int32 {
	if o == nil || o.ConfigurationDeployedUserCount == nil {
		var ret int32
		return ret
	}
	return *o.ConfigurationDeployedUserCount
}

// GetConfigurationDeployedUserCountOk returns a tuple with the ConfigurationDeployedUserCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppPolicyDeploymentSummary) GetConfigurationDeployedUserCountOk() (int32, bool) {
	if o == nil || o.ConfigurationDeployedUserCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ConfigurationDeployedUserCount, true
}

// HasConfigurationDeployedUserCount returns a boolean if a field has been set.
func (o *ManagedAppPolicyDeploymentSummary) HasConfigurationDeployedUserCount() bool {
	if o != nil && o.ConfigurationDeployedUserCount != nil {
		return true
	}

	return false
}

// SetConfigurationDeployedUserCount gets a reference to the given int32 and assigns it to the ConfigurationDeployedUserCount field.
func (o *ManagedAppPolicyDeploymentSummary) SetConfigurationDeployedUserCount(v int32) {
	o.ConfigurationDeployedUserCount = &v
}

// GetLastRefreshTime returns the LastRefreshTime field if non-nil, zero value otherwise.
func (o *ManagedAppPolicyDeploymentSummary) GetLastRefreshTime() time.Time {
	if o == nil || o.LastRefreshTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRefreshTime
}

// GetLastRefreshTimeOk returns a tuple with the LastRefreshTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppPolicyDeploymentSummary) GetLastRefreshTimeOk() (time.Time, bool) {
	if o == nil || o.LastRefreshTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastRefreshTime, true
}

// HasLastRefreshTime returns a boolean if a field has been set.
func (o *ManagedAppPolicyDeploymentSummary) HasLastRefreshTime() bool {
	if o != nil && o.LastRefreshTime != nil {
		return true
	}

	return false
}

// SetLastRefreshTime gets a reference to the given time.Time and assigns it to the LastRefreshTime field.
func (o *ManagedAppPolicyDeploymentSummary) SetLastRefreshTime(v time.Time) {
	o.LastRefreshTime = &v
}

// GetConfigurationDeploymentSummaryPerApp returns the ConfigurationDeploymentSummaryPerApp field if non-nil, zero value otherwise.
func (o *ManagedAppPolicyDeploymentSummary) GetConfigurationDeploymentSummaryPerApp() []AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp {
	if o == nil || o.ConfigurationDeploymentSummaryPerApp == nil {
		var ret []AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp
		return ret
	}
	return *o.ConfigurationDeploymentSummaryPerApp
}

// GetConfigurationDeploymentSummaryPerAppOk returns a tuple with the ConfigurationDeploymentSummaryPerApp field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppPolicyDeploymentSummary) GetConfigurationDeploymentSummaryPerAppOk() ([]AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp, bool) {
	if o == nil || o.ConfigurationDeploymentSummaryPerApp == nil {
		var ret []AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp
		return ret, false
	}
	return *o.ConfigurationDeploymentSummaryPerApp, true
}

// HasConfigurationDeploymentSummaryPerApp returns a boolean if a field has been set.
func (o *ManagedAppPolicyDeploymentSummary) HasConfigurationDeploymentSummaryPerApp() bool {
	if o != nil && o.ConfigurationDeploymentSummaryPerApp != nil {
		return true
	}

	return false
}

// SetConfigurationDeploymentSummaryPerApp gets a reference to the given []AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp and assigns it to the ConfigurationDeploymentSummaryPerApp field.
func (o *ManagedAppPolicyDeploymentSummary) SetConfigurationDeploymentSummaryPerApp(v []AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) {
	o.ConfigurationDeploymentSummaryPerApp = &v
}

// GetVersion returns the Version field if non-nil, zero value otherwise.
func (o *ManagedAppPolicyDeploymentSummary) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppPolicyDeploymentSummary) GetVersionOk() (string, bool) {
	if o == nil || o.Version == nil {
		var ret string
		return ret, false
	}
	return *o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ManagedAppPolicyDeploymentSummary) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ManagedAppPolicyDeploymentSummary) SetVersion(v string) {
	o.Version = &v
}

// SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Version value is set to nil even if false is passed
func (o *ManagedAppPolicyDeploymentSummary) SetVersionExplicitNull(b bool) {
	o.Version = nil
	o.isExplicitNullVersion = b
}

// MarshalJSON returns the JSON representation of the model.
func (o ManagedAppPolicyDeploymentSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ConfigurationDeployedUserCount != nil {
		toSerialize["configurationDeployedUserCount"] = o.ConfigurationDeployedUserCount
	}
	if o.LastRefreshTime != nil {
		toSerialize["lastRefreshTime"] = o.LastRefreshTime
	}
	if o.ConfigurationDeploymentSummaryPerApp != nil {
		toSerialize["configurationDeploymentSummaryPerApp"] = o.ConfigurationDeploymentSummaryPerApp
	}
	if o.Version == nil {
		if o.isExplicitNullVersion {
			toSerialize["version"] = o.Version
		}
	} else {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

