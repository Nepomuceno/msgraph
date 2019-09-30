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
// MicrosoftGraphServicePlanInfo struct for MicrosoftGraphServicePlanInfo
type MicrosoftGraphServicePlanInfo struct {
	ServicePlanId *string `json:"servicePlanId,omitempty"`
	isExplicitNullServicePlanId bool `json:"-"`
	ServicePlanName *string `json:"servicePlanName,omitempty"`
	isExplicitNullServicePlanName bool `json:"-"`
	ProvisioningStatus *string `json:"provisioningStatus,omitempty"`
	isExplicitNullProvisioningStatus bool `json:"-"`
	AppliesTo *string `json:"appliesTo,omitempty"`
	isExplicitNullAppliesTo bool `json:"-"`
}

// GetServicePlanId returns the ServicePlanId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphServicePlanInfo) GetServicePlanId() string {
	if o == nil || o.ServicePlanId == nil {
		var ret string
		return ret
	}
	return *o.ServicePlanId
}

// GetServicePlanIdOk returns a tuple with the ServicePlanId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServicePlanInfo) GetServicePlanIdOk() (string, bool) {
	if o == nil || o.ServicePlanId == nil {
		var ret string
		return ret, false
	}
	return *o.ServicePlanId, true
}

// HasServicePlanId returns a boolean if a field has been set.
func (o *MicrosoftGraphServicePlanInfo) HasServicePlanId() bool {
	if o != nil && o.ServicePlanId != nil {
		return true
	}

	return false
}

// SetServicePlanId gets a reference to the given string and assigns it to the ServicePlanId field.
func (o *MicrosoftGraphServicePlanInfo) SetServicePlanId(v string) {
	o.ServicePlanId = &v
}

// SetServicePlanIdExplicitNull (un)sets ServicePlanId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ServicePlanId value is set to nil even if false is passed
func (o *MicrosoftGraphServicePlanInfo) SetServicePlanIdExplicitNull(b bool) {
	o.ServicePlanId = nil
	o.isExplicitNullServicePlanId = b
}
// GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphServicePlanInfo) GetServicePlanName() string {
	if o == nil || o.ServicePlanName == nil {
		var ret string
		return ret
	}
	return *o.ServicePlanName
}

// GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServicePlanInfo) GetServicePlanNameOk() (string, bool) {
	if o == nil || o.ServicePlanName == nil {
		var ret string
		return ret, false
	}
	return *o.ServicePlanName, true
}

// HasServicePlanName returns a boolean if a field has been set.
func (o *MicrosoftGraphServicePlanInfo) HasServicePlanName() bool {
	if o != nil && o.ServicePlanName != nil {
		return true
	}

	return false
}

// SetServicePlanName gets a reference to the given string and assigns it to the ServicePlanName field.
func (o *MicrosoftGraphServicePlanInfo) SetServicePlanName(v string) {
	o.ServicePlanName = &v
}

// SetServicePlanNameExplicitNull (un)sets ServicePlanName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ServicePlanName value is set to nil even if false is passed
func (o *MicrosoftGraphServicePlanInfo) SetServicePlanNameExplicitNull(b bool) {
	o.ServicePlanName = nil
	o.isExplicitNullServicePlanName = b
}
// GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.
func (o *MicrosoftGraphServicePlanInfo) GetProvisioningStatus() string {
	if o == nil || o.ProvisioningStatus == nil {
		var ret string
		return ret
	}
	return *o.ProvisioningStatus
}

// GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServicePlanInfo) GetProvisioningStatusOk() (string, bool) {
	if o == nil || o.ProvisioningStatus == nil {
		var ret string
		return ret, false
	}
	return *o.ProvisioningStatus, true
}

// HasProvisioningStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphServicePlanInfo) HasProvisioningStatus() bool {
	if o != nil && o.ProvisioningStatus != nil {
		return true
	}

	return false
}

// SetProvisioningStatus gets a reference to the given string and assigns it to the ProvisioningStatus field.
func (o *MicrosoftGraphServicePlanInfo) SetProvisioningStatus(v string) {
	o.ProvisioningStatus = &v
}

// SetProvisioningStatusExplicitNull (un)sets ProvisioningStatus to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ProvisioningStatus value is set to nil even if false is passed
func (o *MicrosoftGraphServicePlanInfo) SetProvisioningStatusExplicitNull(b bool) {
	o.ProvisioningStatus = nil
	o.isExplicitNullProvisioningStatus = b
}
// GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.
func (o *MicrosoftGraphServicePlanInfo) GetAppliesTo() string {
	if o == nil || o.AppliesTo == nil {
		var ret string
		return ret
	}
	return *o.AppliesTo
}

// GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServicePlanInfo) GetAppliesToOk() (string, bool) {
	if o == nil || o.AppliesTo == nil {
		var ret string
		return ret, false
	}
	return *o.AppliesTo, true
}

// HasAppliesTo returns a boolean if a field has been set.
func (o *MicrosoftGraphServicePlanInfo) HasAppliesTo() bool {
	if o != nil && o.AppliesTo != nil {
		return true
	}

	return false
}

// SetAppliesTo gets a reference to the given string and assigns it to the AppliesTo field.
func (o *MicrosoftGraphServicePlanInfo) SetAppliesTo(v string) {
	o.AppliesTo = &v
}

// SetAppliesToExplicitNull (un)sets AppliesTo to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AppliesTo value is set to nil even if false is passed
func (o *MicrosoftGraphServicePlanInfo) SetAppliesToExplicitNull(b bool) {
	o.AppliesTo = nil
	o.isExplicitNullAppliesTo = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphServicePlanInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServicePlanId == nil {
		if o.isExplicitNullServicePlanId {
			toSerialize["servicePlanId"] = o.ServicePlanId
		}
	} else {
		toSerialize["servicePlanId"] = o.ServicePlanId
	}
	if o.ServicePlanName == nil {
		if o.isExplicitNullServicePlanName {
			toSerialize["servicePlanName"] = o.ServicePlanName
		}
	} else {
		toSerialize["servicePlanName"] = o.ServicePlanName
	}
	if o.ProvisioningStatus == nil {
		if o.isExplicitNullProvisioningStatus {
			toSerialize["provisioningStatus"] = o.ProvisioningStatus
		}
	} else {
		toSerialize["provisioningStatus"] = o.ProvisioningStatus
	}
	if o.AppliesTo == nil {
		if o.isExplicitNullAppliesTo {
			toSerialize["appliesTo"] = o.AppliesTo
		}
	} else {
		toSerialize["appliesTo"] = o.AppliesTo
	}
	return json.Marshal(toSerialize)
}

