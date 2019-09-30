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
// MicrosoftGraphIosUpdateDeviceStatus struct for MicrosoftGraphIosUpdateDeviceStatus
type MicrosoftGraphIosUpdateDeviceStatus struct {
	Id *string `json:"id,omitempty"`

	// The installation status of the policy report.
	InstallStatus *AnyOfmicrosoftGraphIosUpdatesInstallStatus `json:"installStatus,omitempty"`

	// The device version that is being reported.
	OsVersion *string `json:"osVersion,omitempty"`
	isExplicitNullOsVersion bool `json:"-"`
	// The device id that is being reported.
	DeviceId *string `json:"deviceId,omitempty"`
	isExplicitNullDeviceId bool `json:"-"`
	// The User id that is being reported.
	UserId *string `json:"userId,omitempty"`
	isExplicitNullUserId bool `json:"-"`
	// Device name of the DevicePolicyStatus.
	DeviceDisplayName *string `json:"deviceDisplayName,omitempty"`
	isExplicitNullDeviceDisplayName bool `json:"-"`
	// The User Name that is being reported
	UserName *string `json:"userName,omitempty"`
	isExplicitNullUserName bool `json:"-"`
	// The device model that is being reported
	DeviceModel *string `json:"deviceModel,omitempty"`
	isExplicitNullDeviceModel bool `json:"-"`
	// The DateTime when device compliance grace period expires
	ComplianceGracePeriodExpirationDateTime *time.Time `json:"complianceGracePeriodExpirationDateTime,omitempty"`

	// Compliance status of the policy report.
	Status *AnyOfmicrosoftGraphComplianceStatus `json:"status,omitempty"`

	// Last modified date time of the policy report.
	LastReportedDateTime *time.Time `json:"lastReportedDateTime,omitempty"`

	// UserPrincipalName.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	isExplicitNullUserPrincipalName bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetId(v string) {
	o.Id = &v
}

// GetInstallStatus returns the InstallStatus field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetInstallStatus() AnyOfmicrosoftGraphIosUpdatesInstallStatus {
	if o == nil || o.InstallStatus == nil {
		var ret AnyOfmicrosoftGraphIosUpdatesInstallStatus
		return ret
	}
	return *o.InstallStatus
}

// GetInstallStatusOk returns a tuple with the InstallStatus field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetInstallStatusOk() (AnyOfmicrosoftGraphIosUpdatesInstallStatus, bool) {
	if o == nil || o.InstallStatus == nil {
		var ret AnyOfmicrosoftGraphIosUpdatesInstallStatus
		return ret, false
	}
	return *o.InstallStatus, true
}

// HasInstallStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) HasInstallStatus() bool {
	if o != nil && o.InstallStatus != nil {
		return true
	}

	return false
}

// SetInstallStatus gets a reference to the given AnyOfmicrosoftGraphIosUpdatesInstallStatus and assigns it to the InstallStatus field.
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetInstallStatus(v AnyOfmicrosoftGraphIosUpdatesInstallStatus) {
	o.InstallStatus = &v
}

// GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetOsVersionOk() (string, bool) {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret, false
	}
	return *o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetOsVersion(v string) {
	o.OsVersion = &v
}

// SetOsVersionExplicitNull (un)sets OsVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OsVersion value is set to nil even if false is passed
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetOsVersionExplicitNull(b bool) {
	o.OsVersion = nil
	o.isExplicitNullOsVersion = b
}
// GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetDeviceIdOk() (string, bool) {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret, false
	}
	return *o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetDeviceId(v string) {
	o.DeviceId = &v
}

// SetDeviceIdExplicitNull (un)sets DeviceId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DeviceId value is set to nil even if false is passed
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetDeviceIdExplicitNull(b bool) {
	o.DeviceId = nil
	o.isExplicitNullDeviceId = b
}
// GetUserId returns the UserId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetUserIdOk() (string, bool) {
	if o == nil || o.UserId == nil {
		var ret string
		return ret, false
	}
	return *o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetUserId(v string) {
	o.UserId = &v
}

// SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UserId value is set to nil even if false is passed
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetUserIdExplicitNull(b bool) {
	o.UserId = nil
	o.isExplicitNullUserId = b
}
// GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetDeviceDisplayName() string {
	if o == nil || o.DeviceDisplayName == nil {
		var ret string
		return ret
	}
	return *o.DeviceDisplayName
}

// GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetDeviceDisplayNameOk() (string, bool) {
	if o == nil || o.DeviceDisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DeviceDisplayName, true
}

// HasDeviceDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) HasDeviceDisplayName() bool {
	if o != nil && o.DeviceDisplayName != nil {
		return true
	}

	return false
}

// SetDeviceDisplayName gets a reference to the given string and assigns it to the DeviceDisplayName field.
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetDeviceDisplayName(v string) {
	o.DeviceDisplayName = &v
}

// SetDeviceDisplayNameExplicitNull (un)sets DeviceDisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DeviceDisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetDeviceDisplayNameExplicitNull(b bool) {
	o.DeviceDisplayName = nil
	o.isExplicitNullDeviceDisplayName = b
}
// GetUserName returns the UserName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetUserNameOk() (string, bool) {
	if o == nil || o.UserName == nil {
		var ret string
		return ret, false
	}
	return *o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetUserName(v string) {
	o.UserName = &v
}

// SetUserNameExplicitNull (un)sets UserName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UserName value is set to nil even if false is passed
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetUserNameExplicitNull(b bool) {
	o.UserName = nil
	o.isExplicitNullUserName = b
}
// GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetDeviceModel() string {
	if o == nil || o.DeviceModel == nil {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetDeviceModelOk() (string, bool) {
	if o == nil || o.DeviceModel == nil {
		var ret string
		return ret, false
	}
	return *o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) HasDeviceModel() bool {
	if o != nil && o.DeviceModel != nil {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// SetDeviceModelExplicitNull (un)sets DeviceModel to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DeviceModel value is set to nil even if false is passed
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetDeviceModelExplicitNull(b bool) {
	o.DeviceModel = nil
	o.isExplicitNullDeviceModel = b
}
// GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetComplianceGracePeriodExpirationDateTime() time.Time {
	if o == nil || o.ComplianceGracePeriodExpirationDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ComplianceGracePeriodExpirationDateTime
}

// GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetComplianceGracePeriodExpirationDateTimeOk() (time.Time, bool) {
	if o == nil || o.ComplianceGracePeriodExpirationDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ComplianceGracePeriodExpirationDateTime, true
}

// HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) HasComplianceGracePeriodExpirationDateTime() bool {
	if o != nil && o.ComplianceGracePeriodExpirationDateTime != nil {
		return true
	}

	return false
}

// SetComplianceGracePeriodExpirationDateTime gets a reference to the given time.Time and assigns it to the ComplianceGracePeriodExpirationDateTime field.
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetComplianceGracePeriodExpirationDateTime(v time.Time) {
	o.ComplianceGracePeriodExpirationDateTime = &v
}

// GetStatus returns the Status field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetStatus() AnyOfmicrosoftGraphComplianceStatus {
	if o == nil || o.Status == nil {
		var ret AnyOfmicrosoftGraphComplianceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetStatusOk() (AnyOfmicrosoftGraphComplianceStatus, bool) {
	if o == nil || o.Status == nil {
		var ret AnyOfmicrosoftGraphComplianceStatus
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the Status field.
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetStatus(v AnyOfmicrosoftGraphComplianceStatus) {
	o.Status = &v
}

// GetLastReportedDateTime returns the LastReportedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetLastReportedDateTime() time.Time {
	if o == nil || o.LastReportedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReportedDateTime
}

// GetLastReportedDateTimeOk returns a tuple with the LastReportedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetLastReportedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastReportedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastReportedDateTime, true
}

// HasLastReportedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) HasLastReportedDateTime() bool {
	if o != nil && o.LastReportedDateTime != nil {
		return true
	}

	return false
}

// SetLastReportedDateTime gets a reference to the given time.Time and assigns it to the LastReportedDateTime field.
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetLastReportedDateTime(v time.Time) {
	o.LastReportedDateTime = &v
}

// GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) GetUserPrincipalNameOk() (string, bool) {
	if o == nil || o.UserPrincipalName == nil {
		var ret string
		return ret, false
	}
	return *o.UserPrincipalName, true
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *MicrosoftGraphIosUpdateDeviceStatus) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName != nil {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetUserPrincipalName(v string) {
	o.UserPrincipalName = &v
}

// SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UserPrincipalName value is set to nil even if false is passed
func (o *MicrosoftGraphIosUpdateDeviceStatus) SetUserPrincipalNameExplicitNull(b bool) {
	o.UserPrincipalName = nil
	o.isExplicitNullUserPrincipalName = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphIosUpdateDeviceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.InstallStatus != nil {
		toSerialize["installStatus"] = o.InstallStatus
	}
	if o.OsVersion == nil {
		if o.isExplicitNullOsVersion {
			toSerialize["osVersion"] = o.OsVersion
		}
	} else {
		toSerialize["osVersion"] = o.OsVersion
	}
	if o.DeviceId == nil {
		if o.isExplicitNullDeviceId {
			toSerialize["deviceId"] = o.DeviceId
		}
	} else {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.UserId == nil {
		if o.isExplicitNullUserId {
			toSerialize["userId"] = o.UserId
		}
	} else {
		toSerialize["userId"] = o.UserId
	}
	if o.DeviceDisplayName == nil {
		if o.isExplicitNullDeviceDisplayName {
			toSerialize["deviceDisplayName"] = o.DeviceDisplayName
		}
	} else {
		toSerialize["deviceDisplayName"] = o.DeviceDisplayName
	}
	if o.UserName == nil {
		if o.isExplicitNullUserName {
			toSerialize["userName"] = o.UserName
		}
	} else {
		toSerialize["userName"] = o.UserName
	}
	if o.DeviceModel == nil {
		if o.isExplicitNullDeviceModel {
			toSerialize["deviceModel"] = o.DeviceModel
		}
	} else {
		toSerialize["deviceModel"] = o.DeviceModel
	}
	if o.ComplianceGracePeriodExpirationDateTime != nil {
		toSerialize["complianceGracePeriodExpirationDateTime"] = o.ComplianceGracePeriodExpirationDateTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.LastReportedDateTime != nil {
		toSerialize["lastReportedDateTime"] = o.LastReportedDateTime
	}
	if o.UserPrincipalName == nil {
		if o.isExplicitNullUserPrincipalName {
			toSerialize["userPrincipalName"] = o.UserPrincipalName
		}
	} else {
		toSerialize["userPrincipalName"] = o.UserPrincipalName
	}
	return json.Marshal(toSerialize)
}


