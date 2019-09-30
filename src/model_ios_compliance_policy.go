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
// IosCompliancePolicy This class contains compliance settings for IOS.
type IosCompliancePolicy struct {
	// Indicates whether or not to block simple passcodes.
	PasscodeBlockSimple *bool `json:"passcodeBlockSimple,omitempty"`

	// Number of days before the passcode expires. Valid values 1 to 65535
	PasscodeExpirationDays *int32 `json:"passcodeExpirationDays,omitempty"`
	isExplicitNullPasscodeExpirationDays bool `json:"-"`
	// Minimum length of passcode. Valid values 4 to 14
	PasscodeMinimumLength *int32 `json:"passcodeMinimumLength,omitempty"`
	isExplicitNullPasscodeMinimumLength bool `json:"-"`
	// Minutes of inactivity before a passcode is required.
	PasscodeMinutesOfInactivityBeforeLock *int32 `json:"passcodeMinutesOfInactivityBeforeLock,omitempty"`
	isExplicitNullPasscodeMinutesOfInactivityBeforeLock bool `json:"-"`
	// Number of previous passcodes to block. Valid values 1 to 24
	PasscodePreviousPasscodeBlockCount *int32 `json:"passcodePreviousPasscodeBlockCount,omitempty"`
	isExplicitNullPasscodePreviousPasscodeBlockCount bool `json:"-"`
	// The number of character sets required in the password.
	PasscodeMinimumCharacterSetCount *int32 `json:"passcodeMinimumCharacterSetCount,omitempty"`
	isExplicitNullPasscodeMinimumCharacterSetCount bool `json:"-"`
	// The required passcode type.
	PasscodeRequiredType *AnyOfmicrosoftGraphRequiredPasswordType `json:"passcodeRequiredType,omitempty"`

	// Indicates whether or not to require a passcode.
	PasscodeRequired *bool `json:"passcodeRequired,omitempty"`

	// Minimum IOS version.
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	isExplicitNullOsMinimumVersion bool `json:"-"`
	// Maximum IOS version.
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	isExplicitNullOsMaximumVersion bool `json:"-"`
	// Devices must not be jailbroken or rooted.
	SecurityBlockJailbrokenDevices *bool `json:"securityBlockJailbrokenDevices,omitempty"`

	// Require that devices have enabled device threat protection .
	DeviceThreatProtectionEnabled *bool `json:"deviceThreatProtectionEnabled,omitempty"`

	// Require Mobile Threat Protection minimum risk level to report noncompliance.
	DeviceThreatProtectionRequiredSecurityLevel *AnyOfmicrosoftGraphDeviceThreatProtectionLevel `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`

	// Indicates whether or not to require a managed email profile.
	ManagedEmailProfileRequired *bool `json:"managedEmailProfileRequired,omitempty"`

}

// GetPasscodeBlockSimple returns the PasscodeBlockSimple field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetPasscodeBlockSimple() bool {
	if o == nil || o.PasscodeBlockSimple == nil {
		var ret bool
		return ret
	}
	return *o.PasscodeBlockSimple
}

// GetPasscodeBlockSimpleOk returns a tuple with the PasscodeBlockSimple field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetPasscodeBlockSimpleOk() (bool, bool) {
	if o == nil || o.PasscodeBlockSimple == nil {
		var ret bool
		return ret, false
	}
	return *o.PasscodeBlockSimple, true
}

// HasPasscodeBlockSimple returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasPasscodeBlockSimple() bool {
	if o != nil && o.PasscodeBlockSimple != nil {
		return true
	}

	return false
}

// SetPasscodeBlockSimple gets a reference to the given bool and assigns it to the PasscodeBlockSimple field.
func (o *IosCompliancePolicy) SetPasscodeBlockSimple(v bool) {
	o.PasscodeBlockSimple = &v
}

// GetPasscodeExpirationDays returns the PasscodeExpirationDays field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetPasscodeExpirationDays() int32 {
	if o == nil || o.PasscodeExpirationDays == nil {
		var ret int32
		return ret
	}
	return *o.PasscodeExpirationDays
}

// GetPasscodeExpirationDaysOk returns a tuple with the PasscodeExpirationDays field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetPasscodeExpirationDaysOk() (int32, bool) {
	if o == nil || o.PasscodeExpirationDays == nil {
		var ret int32
		return ret, false
	}
	return *o.PasscodeExpirationDays, true
}

// HasPasscodeExpirationDays returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasPasscodeExpirationDays() bool {
	if o != nil && o.PasscodeExpirationDays != nil {
		return true
	}

	return false
}

// SetPasscodeExpirationDays gets a reference to the given int32 and assigns it to the PasscodeExpirationDays field.
func (o *IosCompliancePolicy) SetPasscodeExpirationDays(v int32) {
	o.PasscodeExpirationDays = &v
}

// SetPasscodeExpirationDaysExplicitNull (un)sets PasscodeExpirationDays to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasscodeExpirationDays value is set to nil even if false is passed
func (o *IosCompliancePolicy) SetPasscodeExpirationDaysExplicitNull(b bool) {
	o.PasscodeExpirationDays = nil
	o.isExplicitNullPasscodeExpirationDays = b
}
// GetPasscodeMinimumLength returns the PasscodeMinimumLength field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetPasscodeMinimumLength() int32 {
	if o == nil || o.PasscodeMinimumLength == nil {
		var ret int32
		return ret
	}
	return *o.PasscodeMinimumLength
}

// GetPasscodeMinimumLengthOk returns a tuple with the PasscodeMinimumLength field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetPasscodeMinimumLengthOk() (int32, bool) {
	if o == nil || o.PasscodeMinimumLength == nil {
		var ret int32
		return ret, false
	}
	return *o.PasscodeMinimumLength, true
}

// HasPasscodeMinimumLength returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasPasscodeMinimumLength() bool {
	if o != nil && o.PasscodeMinimumLength != nil {
		return true
	}

	return false
}

// SetPasscodeMinimumLength gets a reference to the given int32 and assigns it to the PasscodeMinimumLength field.
func (o *IosCompliancePolicy) SetPasscodeMinimumLength(v int32) {
	o.PasscodeMinimumLength = &v
}

// SetPasscodeMinimumLengthExplicitNull (un)sets PasscodeMinimumLength to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasscodeMinimumLength value is set to nil even if false is passed
func (o *IosCompliancePolicy) SetPasscodeMinimumLengthExplicitNull(b bool) {
	o.PasscodeMinimumLength = nil
	o.isExplicitNullPasscodeMinimumLength = b
}
// GetPasscodeMinutesOfInactivityBeforeLock returns the PasscodeMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetPasscodeMinutesOfInactivityBeforeLock() int32 {
	if o == nil || o.PasscodeMinutesOfInactivityBeforeLock == nil {
		var ret int32
		return ret
	}
	return *o.PasscodeMinutesOfInactivityBeforeLock
}

// GetPasscodeMinutesOfInactivityBeforeLockOk returns a tuple with the PasscodeMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetPasscodeMinutesOfInactivityBeforeLockOk() (int32, bool) {
	if o == nil || o.PasscodeMinutesOfInactivityBeforeLock == nil {
		var ret int32
		return ret, false
	}
	return *o.PasscodeMinutesOfInactivityBeforeLock, true
}

// HasPasscodeMinutesOfInactivityBeforeLock returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasPasscodeMinutesOfInactivityBeforeLock() bool {
	if o != nil && o.PasscodeMinutesOfInactivityBeforeLock != nil {
		return true
	}

	return false
}

// SetPasscodeMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasscodeMinutesOfInactivityBeforeLock field.
func (o *IosCompliancePolicy) SetPasscodeMinutesOfInactivityBeforeLock(v int32) {
	o.PasscodeMinutesOfInactivityBeforeLock = &v
}

// SetPasscodeMinutesOfInactivityBeforeLockExplicitNull (un)sets PasscodeMinutesOfInactivityBeforeLock to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasscodeMinutesOfInactivityBeforeLock value is set to nil even if false is passed
func (o *IosCompliancePolicy) SetPasscodeMinutesOfInactivityBeforeLockExplicitNull(b bool) {
	o.PasscodeMinutesOfInactivityBeforeLock = nil
	o.isExplicitNullPasscodeMinutesOfInactivityBeforeLock = b
}
// GetPasscodePreviousPasscodeBlockCount returns the PasscodePreviousPasscodeBlockCount field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetPasscodePreviousPasscodeBlockCount() int32 {
	if o == nil || o.PasscodePreviousPasscodeBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.PasscodePreviousPasscodeBlockCount
}

// GetPasscodePreviousPasscodeBlockCountOk returns a tuple with the PasscodePreviousPasscodeBlockCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetPasscodePreviousPasscodeBlockCountOk() (int32, bool) {
	if o == nil || o.PasscodePreviousPasscodeBlockCount == nil {
		var ret int32
		return ret, false
	}
	return *o.PasscodePreviousPasscodeBlockCount, true
}

// HasPasscodePreviousPasscodeBlockCount returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasPasscodePreviousPasscodeBlockCount() bool {
	if o != nil && o.PasscodePreviousPasscodeBlockCount != nil {
		return true
	}

	return false
}

// SetPasscodePreviousPasscodeBlockCount gets a reference to the given int32 and assigns it to the PasscodePreviousPasscodeBlockCount field.
func (o *IosCompliancePolicy) SetPasscodePreviousPasscodeBlockCount(v int32) {
	o.PasscodePreviousPasscodeBlockCount = &v
}

// SetPasscodePreviousPasscodeBlockCountExplicitNull (un)sets PasscodePreviousPasscodeBlockCount to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasscodePreviousPasscodeBlockCount value is set to nil even if false is passed
func (o *IosCompliancePolicy) SetPasscodePreviousPasscodeBlockCountExplicitNull(b bool) {
	o.PasscodePreviousPasscodeBlockCount = nil
	o.isExplicitNullPasscodePreviousPasscodeBlockCount = b
}
// GetPasscodeMinimumCharacterSetCount returns the PasscodeMinimumCharacterSetCount field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetPasscodeMinimumCharacterSetCount() int32 {
	if o == nil || o.PasscodeMinimumCharacterSetCount == nil {
		var ret int32
		return ret
	}
	return *o.PasscodeMinimumCharacterSetCount
}

// GetPasscodeMinimumCharacterSetCountOk returns a tuple with the PasscodeMinimumCharacterSetCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetPasscodeMinimumCharacterSetCountOk() (int32, bool) {
	if o == nil || o.PasscodeMinimumCharacterSetCount == nil {
		var ret int32
		return ret, false
	}
	return *o.PasscodeMinimumCharacterSetCount, true
}

// HasPasscodeMinimumCharacterSetCount returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasPasscodeMinimumCharacterSetCount() bool {
	if o != nil && o.PasscodeMinimumCharacterSetCount != nil {
		return true
	}

	return false
}

// SetPasscodeMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasscodeMinimumCharacterSetCount field.
func (o *IosCompliancePolicy) SetPasscodeMinimumCharacterSetCount(v int32) {
	o.PasscodeMinimumCharacterSetCount = &v
}

// SetPasscodeMinimumCharacterSetCountExplicitNull (un)sets PasscodeMinimumCharacterSetCount to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasscodeMinimumCharacterSetCount value is set to nil even if false is passed
func (o *IosCompliancePolicy) SetPasscodeMinimumCharacterSetCountExplicitNull(b bool) {
	o.PasscodeMinimumCharacterSetCount = nil
	o.isExplicitNullPasscodeMinimumCharacterSetCount = b
}
// GetPasscodeRequiredType returns the PasscodeRequiredType field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetPasscodeRequiredType() AnyOfmicrosoftGraphRequiredPasswordType {
	if o == nil || o.PasscodeRequiredType == nil {
		var ret AnyOfmicrosoftGraphRequiredPasswordType
		return ret
	}
	return *o.PasscodeRequiredType
}

// GetPasscodeRequiredTypeOk returns a tuple with the PasscodeRequiredType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetPasscodeRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool) {
	if o == nil || o.PasscodeRequiredType == nil {
		var ret AnyOfmicrosoftGraphRequiredPasswordType
		return ret, false
	}
	return *o.PasscodeRequiredType, true
}

// HasPasscodeRequiredType returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasPasscodeRequiredType() bool {
	if o != nil && o.PasscodeRequiredType != nil {
		return true
	}

	return false
}

// SetPasscodeRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasscodeRequiredType field.
func (o *IosCompliancePolicy) SetPasscodeRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType) {
	o.PasscodeRequiredType = &v
}

// GetPasscodeRequired returns the PasscodeRequired field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetPasscodeRequired() bool {
	if o == nil || o.PasscodeRequired == nil {
		var ret bool
		return ret
	}
	return *o.PasscodeRequired
}

// GetPasscodeRequiredOk returns a tuple with the PasscodeRequired field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetPasscodeRequiredOk() (bool, bool) {
	if o == nil || o.PasscodeRequired == nil {
		var ret bool
		return ret, false
	}
	return *o.PasscodeRequired, true
}

// HasPasscodeRequired returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasPasscodeRequired() bool {
	if o != nil && o.PasscodeRequired != nil {
		return true
	}

	return false
}

// SetPasscodeRequired gets a reference to the given bool and assigns it to the PasscodeRequired field.
func (o *IosCompliancePolicy) SetPasscodeRequired(v bool) {
	o.PasscodeRequired = &v
}

// GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetOsMinimumVersion() string {
	if o == nil || o.OsMinimumVersion == nil {
		var ret string
		return ret
	}
	return *o.OsMinimumVersion
}

// GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetOsMinimumVersionOk() (string, bool) {
	if o == nil || o.OsMinimumVersion == nil {
		var ret string
		return ret, false
	}
	return *o.OsMinimumVersion, true
}

// HasOsMinimumVersion returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasOsMinimumVersion() bool {
	if o != nil && o.OsMinimumVersion != nil {
		return true
	}

	return false
}

// SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.
func (o *IosCompliancePolicy) SetOsMinimumVersion(v string) {
	o.OsMinimumVersion = &v
}

// SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OsMinimumVersion value is set to nil even if false is passed
func (o *IosCompliancePolicy) SetOsMinimumVersionExplicitNull(b bool) {
	o.OsMinimumVersion = nil
	o.isExplicitNullOsMinimumVersion = b
}
// GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetOsMaximumVersion() string {
	if o == nil || o.OsMaximumVersion == nil {
		var ret string
		return ret
	}
	return *o.OsMaximumVersion
}

// GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetOsMaximumVersionOk() (string, bool) {
	if o == nil || o.OsMaximumVersion == nil {
		var ret string
		return ret, false
	}
	return *o.OsMaximumVersion, true
}

// HasOsMaximumVersion returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasOsMaximumVersion() bool {
	if o != nil && o.OsMaximumVersion != nil {
		return true
	}

	return false
}

// SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.
func (o *IosCompliancePolicy) SetOsMaximumVersion(v string) {
	o.OsMaximumVersion = &v
}

// SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OsMaximumVersion value is set to nil even if false is passed
func (o *IosCompliancePolicy) SetOsMaximumVersionExplicitNull(b bool) {
	o.OsMaximumVersion = nil
	o.isExplicitNullOsMaximumVersion = b
}
// GetSecurityBlockJailbrokenDevices returns the SecurityBlockJailbrokenDevices field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetSecurityBlockJailbrokenDevices() bool {
	if o == nil || o.SecurityBlockJailbrokenDevices == nil {
		var ret bool
		return ret
	}
	return *o.SecurityBlockJailbrokenDevices
}

// GetSecurityBlockJailbrokenDevicesOk returns a tuple with the SecurityBlockJailbrokenDevices field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetSecurityBlockJailbrokenDevicesOk() (bool, bool) {
	if o == nil || o.SecurityBlockJailbrokenDevices == nil {
		var ret bool
		return ret, false
	}
	return *o.SecurityBlockJailbrokenDevices, true
}

// HasSecurityBlockJailbrokenDevices returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasSecurityBlockJailbrokenDevices() bool {
	if o != nil && o.SecurityBlockJailbrokenDevices != nil {
		return true
	}

	return false
}

// SetSecurityBlockJailbrokenDevices gets a reference to the given bool and assigns it to the SecurityBlockJailbrokenDevices field.
func (o *IosCompliancePolicy) SetSecurityBlockJailbrokenDevices(v bool) {
	o.SecurityBlockJailbrokenDevices = &v
}

// GetDeviceThreatProtectionEnabled returns the DeviceThreatProtectionEnabled field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetDeviceThreatProtectionEnabled() bool {
	if o == nil || o.DeviceThreatProtectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DeviceThreatProtectionEnabled
}

// GetDeviceThreatProtectionEnabledOk returns a tuple with the DeviceThreatProtectionEnabled field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetDeviceThreatProtectionEnabledOk() (bool, bool) {
	if o == nil || o.DeviceThreatProtectionEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.DeviceThreatProtectionEnabled, true
}

// HasDeviceThreatProtectionEnabled returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasDeviceThreatProtectionEnabled() bool {
	if o != nil && o.DeviceThreatProtectionEnabled != nil {
		return true
	}

	return false
}

// SetDeviceThreatProtectionEnabled gets a reference to the given bool and assigns it to the DeviceThreatProtectionEnabled field.
func (o *IosCompliancePolicy) SetDeviceThreatProtectionEnabled(v bool) {
	o.DeviceThreatProtectionEnabled = &v
}

// GetDeviceThreatProtectionRequiredSecurityLevel returns the DeviceThreatProtectionRequiredSecurityLevel field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel() AnyOfmicrosoftGraphDeviceThreatProtectionLevel {
	if o == nil || o.DeviceThreatProtectionRequiredSecurityLevel == nil {
		var ret AnyOfmicrosoftGraphDeviceThreatProtectionLevel
		return ret
	}
	return *o.DeviceThreatProtectionRequiredSecurityLevel
}

// GetDeviceThreatProtectionRequiredSecurityLevelOk returns a tuple with the DeviceThreatProtectionRequiredSecurityLevel field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevelOk() (AnyOfmicrosoftGraphDeviceThreatProtectionLevel, bool) {
	if o == nil || o.DeviceThreatProtectionRequiredSecurityLevel == nil {
		var ret AnyOfmicrosoftGraphDeviceThreatProtectionLevel
		return ret, false
	}
	return *o.DeviceThreatProtectionRequiredSecurityLevel, true
}

// HasDeviceThreatProtectionRequiredSecurityLevel returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasDeviceThreatProtectionRequiredSecurityLevel() bool {
	if o != nil && o.DeviceThreatProtectionRequiredSecurityLevel != nil {
		return true
	}

	return false
}

// SetDeviceThreatProtectionRequiredSecurityLevel gets a reference to the given AnyOfmicrosoftGraphDeviceThreatProtectionLevel and assigns it to the DeviceThreatProtectionRequiredSecurityLevel field.
func (o *IosCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(v AnyOfmicrosoftGraphDeviceThreatProtectionLevel) {
	o.DeviceThreatProtectionRequiredSecurityLevel = &v
}

// GetManagedEmailProfileRequired returns the ManagedEmailProfileRequired field if non-nil, zero value otherwise.
func (o *IosCompliancePolicy) GetManagedEmailProfileRequired() bool {
	if o == nil || o.ManagedEmailProfileRequired == nil {
		var ret bool
		return ret
	}
	return *o.ManagedEmailProfileRequired
}

// GetManagedEmailProfileRequiredOk returns a tuple with the ManagedEmailProfileRequired field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *IosCompliancePolicy) GetManagedEmailProfileRequiredOk() (bool, bool) {
	if o == nil || o.ManagedEmailProfileRequired == nil {
		var ret bool
		return ret, false
	}
	return *o.ManagedEmailProfileRequired, true
}

// HasManagedEmailProfileRequired returns a boolean if a field has been set.
func (o *IosCompliancePolicy) HasManagedEmailProfileRequired() bool {
	if o != nil && o.ManagedEmailProfileRequired != nil {
		return true
	}

	return false
}

// SetManagedEmailProfileRequired gets a reference to the given bool and assigns it to the ManagedEmailProfileRequired field.
func (o *IosCompliancePolicy) SetManagedEmailProfileRequired(v bool) {
	o.ManagedEmailProfileRequired = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o IosCompliancePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PasscodeBlockSimple != nil {
		toSerialize["passcodeBlockSimple"] = o.PasscodeBlockSimple
	}
	if o.PasscodeExpirationDays == nil {
		if o.isExplicitNullPasscodeExpirationDays {
			toSerialize["passcodeExpirationDays"] = o.PasscodeExpirationDays
		}
	} else {
		toSerialize["passcodeExpirationDays"] = o.PasscodeExpirationDays
	}
	if o.PasscodeMinimumLength == nil {
		if o.isExplicitNullPasscodeMinimumLength {
			toSerialize["passcodeMinimumLength"] = o.PasscodeMinimumLength
		}
	} else {
		toSerialize["passcodeMinimumLength"] = o.PasscodeMinimumLength
	}
	if o.PasscodeMinutesOfInactivityBeforeLock == nil {
		if o.isExplicitNullPasscodeMinutesOfInactivityBeforeLock {
			toSerialize["passcodeMinutesOfInactivityBeforeLock"] = o.PasscodeMinutesOfInactivityBeforeLock
		}
	} else {
		toSerialize["passcodeMinutesOfInactivityBeforeLock"] = o.PasscodeMinutesOfInactivityBeforeLock
	}
	if o.PasscodePreviousPasscodeBlockCount == nil {
		if o.isExplicitNullPasscodePreviousPasscodeBlockCount {
			toSerialize["passcodePreviousPasscodeBlockCount"] = o.PasscodePreviousPasscodeBlockCount
		}
	} else {
		toSerialize["passcodePreviousPasscodeBlockCount"] = o.PasscodePreviousPasscodeBlockCount
	}
	if o.PasscodeMinimumCharacterSetCount == nil {
		if o.isExplicitNullPasscodeMinimumCharacterSetCount {
			toSerialize["passcodeMinimumCharacterSetCount"] = o.PasscodeMinimumCharacterSetCount
		}
	} else {
		toSerialize["passcodeMinimumCharacterSetCount"] = o.PasscodeMinimumCharacterSetCount
	}
	if o.PasscodeRequiredType != nil {
		toSerialize["passcodeRequiredType"] = o.PasscodeRequiredType
	}
	if o.PasscodeRequired != nil {
		toSerialize["passcodeRequired"] = o.PasscodeRequired
	}
	if o.OsMinimumVersion == nil {
		if o.isExplicitNullOsMinimumVersion {
			toSerialize["osMinimumVersion"] = o.OsMinimumVersion
		}
	} else {
		toSerialize["osMinimumVersion"] = o.OsMinimumVersion
	}
	if o.OsMaximumVersion == nil {
		if o.isExplicitNullOsMaximumVersion {
			toSerialize["osMaximumVersion"] = o.OsMaximumVersion
		}
	} else {
		toSerialize["osMaximumVersion"] = o.OsMaximumVersion
	}
	if o.SecurityBlockJailbrokenDevices != nil {
		toSerialize["securityBlockJailbrokenDevices"] = o.SecurityBlockJailbrokenDevices
	}
	if o.DeviceThreatProtectionEnabled != nil {
		toSerialize["deviceThreatProtectionEnabled"] = o.DeviceThreatProtectionEnabled
	}
	if o.DeviceThreatProtectionRequiredSecurityLevel != nil {
		toSerialize["deviceThreatProtectionRequiredSecurityLevel"] = o.DeviceThreatProtectionRequiredSecurityLevel
	}
	if o.ManagedEmailProfileRequired != nil {
		toSerialize["managedEmailProfileRequired"] = o.ManagedEmailProfileRequired
	}
	return json.Marshal(toSerialize)
}


