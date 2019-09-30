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
// ManagedAppRegistration The ManagedAppEntity is the base entity type for all other entity types under app management workflow.
type ManagedAppRegistration struct {
	// Date and time of creation
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`

	// Date and time of last the app synced with management service.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`

	// App version
	ApplicationVersion *string `json:"applicationVersion,omitempty"`
	isExplicitNullApplicationVersion bool `json:"-"`
	// App management SDK version
	ManagementSdkVersion *string `json:"managementSdkVersion,omitempty"`
	isExplicitNullManagementSdkVersion bool `json:"-"`
	// Operating System version
	PlatformVersion *string `json:"platformVersion,omitempty"`
	isExplicitNullPlatformVersion bool `json:"-"`
	// Host device type
	DeviceType *string `json:"deviceType,omitempty"`
	isExplicitNullDeviceType bool `json:"-"`
	// App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions.
	DeviceTag *string `json:"deviceTag,omitempty"`
	isExplicitNullDeviceTag bool `json:"-"`
	// Host device name
	DeviceName *string `json:"deviceName,omitempty"`
	isExplicitNullDeviceName bool `json:"-"`
	// Zero or more reasons an app registration is flagged. E.g. app running on rooted device
	FlaggedReasons *[]AnyOfmicrosoftGraphManagedAppFlaggedReason `json:"flaggedReasons,omitempty"`

	// The user Id to who this app registration belongs.
	UserId *string `json:"userId,omitempty"`
	isExplicitNullUserId bool `json:"-"`
	// The app package Identifier
	AppIdentifier *AnyOfobject `json:"appIdentifier,omitempty"`
	isExplicitNullAppIdentifier bool `json:"-"`
	// Version of the entity.
	Version *string `json:"version,omitempty"`
	isExplicitNullVersion bool `json:"-"`
	AppliedPolicies *[]MicrosoftGraphManagedAppPolicy `json:"appliedPolicies,omitempty"`

	IntendedPolicies *[]MicrosoftGraphManagedAppPolicy `json:"intendedPolicies,omitempty"`

	Operations *[]MicrosoftGraphManagedAppOperation `json:"operations,omitempty"`

}

// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *ManagedAppRegistration) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetLastSyncDateTime() time.Time {
	if o == nil || o.LastSyncDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSyncDateTime
}

// GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetLastSyncDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastSyncDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastSyncDateTime, true
}

// HasLastSyncDateTime returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasLastSyncDateTime() bool {
	if o != nil && o.LastSyncDateTime != nil {
		return true
	}

	return false
}

// SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.
func (o *ManagedAppRegistration) SetLastSyncDateTime(v time.Time) {
	o.LastSyncDateTime = &v
}

// GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetApplicationVersion() string {
	if o == nil || o.ApplicationVersion == nil {
		var ret string
		return ret
	}
	return *o.ApplicationVersion
}

// GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetApplicationVersionOk() (string, bool) {
	if o == nil || o.ApplicationVersion == nil {
		var ret string
		return ret, false
	}
	return *o.ApplicationVersion, true
}

// HasApplicationVersion returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasApplicationVersion() bool {
	if o != nil && o.ApplicationVersion != nil {
		return true
	}

	return false
}

// SetApplicationVersion gets a reference to the given string and assigns it to the ApplicationVersion field.
func (o *ManagedAppRegistration) SetApplicationVersion(v string) {
	o.ApplicationVersion = &v
}

// SetApplicationVersionExplicitNull (un)sets ApplicationVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ApplicationVersion value is set to nil even if false is passed
func (o *ManagedAppRegistration) SetApplicationVersionExplicitNull(b bool) {
	o.ApplicationVersion = nil
	o.isExplicitNullApplicationVersion = b
}
// GetManagementSdkVersion returns the ManagementSdkVersion field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetManagementSdkVersion() string {
	if o == nil || o.ManagementSdkVersion == nil {
		var ret string
		return ret
	}
	return *o.ManagementSdkVersion
}

// GetManagementSdkVersionOk returns a tuple with the ManagementSdkVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetManagementSdkVersionOk() (string, bool) {
	if o == nil || o.ManagementSdkVersion == nil {
		var ret string
		return ret, false
	}
	return *o.ManagementSdkVersion, true
}

// HasManagementSdkVersion returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasManagementSdkVersion() bool {
	if o != nil && o.ManagementSdkVersion != nil {
		return true
	}

	return false
}

// SetManagementSdkVersion gets a reference to the given string and assigns it to the ManagementSdkVersion field.
func (o *ManagedAppRegistration) SetManagementSdkVersion(v string) {
	o.ManagementSdkVersion = &v
}

// SetManagementSdkVersionExplicitNull (un)sets ManagementSdkVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ManagementSdkVersion value is set to nil even if false is passed
func (o *ManagedAppRegistration) SetManagementSdkVersionExplicitNull(b bool) {
	o.ManagementSdkVersion = nil
	o.isExplicitNullManagementSdkVersion = b
}
// GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetPlatformVersion() string {
	if o == nil || o.PlatformVersion == nil {
		var ret string
		return ret
	}
	return *o.PlatformVersion
}

// GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetPlatformVersionOk() (string, bool) {
	if o == nil || o.PlatformVersion == nil {
		var ret string
		return ret, false
	}
	return *o.PlatformVersion, true
}

// HasPlatformVersion returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasPlatformVersion() bool {
	if o != nil && o.PlatformVersion != nil {
		return true
	}

	return false
}

// SetPlatformVersion gets a reference to the given string and assigns it to the PlatformVersion field.
func (o *ManagedAppRegistration) SetPlatformVersion(v string) {
	o.PlatformVersion = &v
}

// SetPlatformVersionExplicitNull (un)sets PlatformVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PlatformVersion value is set to nil even if false is passed
func (o *ManagedAppRegistration) SetPlatformVersionExplicitNull(b bool) {
	o.PlatformVersion = nil
	o.isExplicitNullPlatformVersion = b
}
// GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetDeviceTypeOk() (string, bool) {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret, false
	}
	return *o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *ManagedAppRegistration) SetDeviceType(v string) {
	o.DeviceType = &v
}

// SetDeviceTypeExplicitNull (un)sets DeviceType to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DeviceType value is set to nil even if false is passed
func (o *ManagedAppRegistration) SetDeviceTypeExplicitNull(b bool) {
	o.DeviceType = nil
	o.isExplicitNullDeviceType = b
}
// GetDeviceTag returns the DeviceTag field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetDeviceTag() string {
	if o == nil || o.DeviceTag == nil {
		var ret string
		return ret
	}
	return *o.DeviceTag
}

// GetDeviceTagOk returns a tuple with the DeviceTag field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetDeviceTagOk() (string, bool) {
	if o == nil || o.DeviceTag == nil {
		var ret string
		return ret, false
	}
	return *o.DeviceTag, true
}

// HasDeviceTag returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasDeviceTag() bool {
	if o != nil && o.DeviceTag != nil {
		return true
	}

	return false
}

// SetDeviceTag gets a reference to the given string and assigns it to the DeviceTag field.
func (o *ManagedAppRegistration) SetDeviceTag(v string) {
	o.DeviceTag = &v
}

// SetDeviceTagExplicitNull (un)sets DeviceTag to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DeviceTag value is set to nil even if false is passed
func (o *ManagedAppRegistration) SetDeviceTagExplicitNull(b bool) {
	o.DeviceTag = nil
	o.isExplicitNullDeviceTag = b
}
// GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetDeviceNameOk() (string, bool) {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret, false
	}
	return *o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *ManagedAppRegistration) SetDeviceName(v string) {
	o.DeviceName = &v
}

// SetDeviceNameExplicitNull (un)sets DeviceName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DeviceName value is set to nil even if false is passed
func (o *ManagedAppRegistration) SetDeviceNameExplicitNull(b bool) {
	o.DeviceName = nil
	o.isExplicitNullDeviceName = b
}
// GetFlaggedReasons returns the FlaggedReasons field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetFlaggedReasons() []AnyOfmicrosoftGraphManagedAppFlaggedReason {
	if o == nil || o.FlaggedReasons == nil {
		var ret []AnyOfmicrosoftGraphManagedAppFlaggedReason
		return ret
	}
	return *o.FlaggedReasons
}

// GetFlaggedReasonsOk returns a tuple with the FlaggedReasons field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetFlaggedReasonsOk() ([]AnyOfmicrosoftGraphManagedAppFlaggedReason, bool) {
	if o == nil || o.FlaggedReasons == nil {
		var ret []AnyOfmicrosoftGraphManagedAppFlaggedReason
		return ret, false
	}
	return *o.FlaggedReasons, true
}

// HasFlaggedReasons returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasFlaggedReasons() bool {
	if o != nil && o.FlaggedReasons != nil {
		return true
	}

	return false
}

// SetFlaggedReasons gets a reference to the given []AnyOfmicrosoftGraphManagedAppFlaggedReason and assigns it to the FlaggedReasons field.
func (o *ManagedAppRegistration) SetFlaggedReasons(v []AnyOfmicrosoftGraphManagedAppFlaggedReason) {
	o.FlaggedReasons = &v
}

// GetUserId returns the UserId field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetUserIdOk() (string, bool) {
	if o == nil || o.UserId == nil {
		var ret string
		return ret, false
	}
	return *o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ManagedAppRegistration) SetUserId(v string) {
	o.UserId = &v
}

// SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UserId value is set to nil even if false is passed
func (o *ManagedAppRegistration) SetUserIdExplicitNull(b bool) {
	o.UserId = nil
	o.isExplicitNullUserId = b
}
// GetAppIdentifier returns the AppIdentifier field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetAppIdentifier() AnyOfobject {
	if o == nil || o.AppIdentifier == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.AppIdentifier
}

// GetAppIdentifierOk returns a tuple with the AppIdentifier field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetAppIdentifierOk() (AnyOfobject, bool) {
	if o == nil || o.AppIdentifier == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.AppIdentifier, true
}

// HasAppIdentifier returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasAppIdentifier() bool {
	if o != nil && o.AppIdentifier != nil {
		return true
	}

	return false
}

// SetAppIdentifier gets a reference to the given AnyOfobject and assigns it to the AppIdentifier field.
func (o *ManagedAppRegistration) SetAppIdentifier(v AnyOfobject) {
	o.AppIdentifier = &v
}

// SetAppIdentifierExplicitNull (un)sets AppIdentifier to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AppIdentifier value is set to nil even if false is passed
func (o *ManagedAppRegistration) SetAppIdentifierExplicitNull(b bool) {
	o.AppIdentifier = nil
	o.isExplicitNullAppIdentifier = b
}
// GetVersion returns the Version field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetVersionOk() (string, bool) {
	if o == nil || o.Version == nil {
		var ret string
		return ret, false
	}
	return *o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ManagedAppRegistration) SetVersion(v string) {
	o.Version = &v
}

// SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Version value is set to nil even if false is passed
func (o *ManagedAppRegistration) SetVersionExplicitNull(b bool) {
	o.Version = nil
	o.isExplicitNullVersion = b
}
// GetAppliedPolicies returns the AppliedPolicies field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetAppliedPolicies() []MicrosoftGraphManagedAppPolicy {
	if o == nil || o.AppliedPolicies == nil {
		var ret []MicrosoftGraphManagedAppPolicy
		return ret
	}
	return *o.AppliedPolicies
}

// GetAppliedPoliciesOk returns a tuple with the AppliedPolicies field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetAppliedPoliciesOk() ([]MicrosoftGraphManagedAppPolicy, bool) {
	if o == nil || o.AppliedPolicies == nil {
		var ret []MicrosoftGraphManagedAppPolicy
		return ret, false
	}
	return *o.AppliedPolicies, true
}

// HasAppliedPolicies returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasAppliedPolicies() bool {
	if o != nil && o.AppliedPolicies != nil {
		return true
	}

	return false
}

// SetAppliedPolicies gets a reference to the given []MicrosoftGraphManagedAppPolicy and assigns it to the AppliedPolicies field.
func (o *ManagedAppRegistration) SetAppliedPolicies(v []MicrosoftGraphManagedAppPolicy) {
	o.AppliedPolicies = &v
}

// GetIntendedPolicies returns the IntendedPolicies field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetIntendedPolicies() []MicrosoftGraphManagedAppPolicy {
	if o == nil || o.IntendedPolicies == nil {
		var ret []MicrosoftGraphManagedAppPolicy
		return ret
	}
	return *o.IntendedPolicies
}

// GetIntendedPoliciesOk returns a tuple with the IntendedPolicies field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetIntendedPoliciesOk() ([]MicrosoftGraphManagedAppPolicy, bool) {
	if o == nil || o.IntendedPolicies == nil {
		var ret []MicrosoftGraphManagedAppPolicy
		return ret, false
	}
	return *o.IntendedPolicies, true
}

// HasIntendedPolicies returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasIntendedPolicies() bool {
	if o != nil && o.IntendedPolicies != nil {
		return true
	}

	return false
}

// SetIntendedPolicies gets a reference to the given []MicrosoftGraphManagedAppPolicy and assigns it to the IntendedPolicies field.
func (o *ManagedAppRegistration) SetIntendedPolicies(v []MicrosoftGraphManagedAppPolicy) {
	o.IntendedPolicies = &v
}

// GetOperations returns the Operations field if non-nil, zero value otherwise.
func (o *ManagedAppRegistration) GetOperations() []MicrosoftGraphManagedAppOperation {
	if o == nil || o.Operations == nil {
		var ret []MicrosoftGraphManagedAppOperation
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetOperationsOk() ([]MicrosoftGraphManagedAppOperation, bool) {
	if o == nil || o.Operations == nil {
		var ret []MicrosoftGraphManagedAppOperation
		return ret, false
	}
	return *o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []MicrosoftGraphManagedAppOperation and assigns it to the Operations field.
func (o *ManagedAppRegistration) SetOperations(v []MicrosoftGraphManagedAppOperation) {
	o.Operations = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o ManagedAppRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.LastSyncDateTime != nil {
		toSerialize["lastSyncDateTime"] = o.LastSyncDateTime
	}
	if o.ApplicationVersion == nil {
		if o.isExplicitNullApplicationVersion {
			toSerialize["applicationVersion"] = o.ApplicationVersion
		}
	} else {
		toSerialize["applicationVersion"] = o.ApplicationVersion
	}
	if o.ManagementSdkVersion == nil {
		if o.isExplicitNullManagementSdkVersion {
			toSerialize["managementSdkVersion"] = o.ManagementSdkVersion
		}
	} else {
		toSerialize["managementSdkVersion"] = o.ManagementSdkVersion
	}
	if o.PlatformVersion == nil {
		if o.isExplicitNullPlatformVersion {
			toSerialize["platformVersion"] = o.PlatformVersion
		}
	} else {
		toSerialize["platformVersion"] = o.PlatformVersion
	}
	if o.DeviceType == nil {
		if o.isExplicitNullDeviceType {
			toSerialize["deviceType"] = o.DeviceType
		}
	} else {
		toSerialize["deviceType"] = o.DeviceType
	}
	if o.DeviceTag == nil {
		if o.isExplicitNullDeviceTag {
			toSerialize["deviceTag"] = o.DeviceTag
		}
	} else {
		toSerialize["deviceTag"] = o.DeviceTag
	}
	if o.DeviceName == nil {
		if o.isExplicitNullDeviceName {
			toSerialize["deviceName"] = o.DeviceName
		}
	} else {
		toSerialize["deviceName"] = o.DeviceName
	}
	if o.FlaggedReasons != nil {
		toSerialize["flaggedReasons"] = o.FlaggedReasons
	}
	if o.UserId == nil {
		if o.isExplicitNullUserId {
			toSerialize["userId"] = o.UserId
		}
	} else {
		toSerialize["userId"] = o.UserId
	}
	if o.AppIdentifier == nil {
		if o.isExplicitNullAppIdentifier {
			toSerialize["appIdentifier"] = o.AppIdentifier
		}
	} else {
		toSerialize["appIdentifier"] = o.AppIdentifier
	}
	if o.Version == nil {
		if o.isExplicitNullVersion {
			toSerialize["version"] = o.Version
		}
	} else {
		toSerialize["version"] = o.Version
	}
	if o.AppliedPolicies != nil {
		toSerialize["appliedPolicies"] = o.AppliedPolicies
	}
	if o.IntendedPolicies != nil {
		toSerialize["intendedPolicies"] = o.IntendedPolicies
	}
	if o.Operations != nil {
		toSerialize["operations"] = o.Operations
	}
	return json.Marshal(toSerialize)
}

