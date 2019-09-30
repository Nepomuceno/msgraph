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
// MacOsCompliancePolicy This class contains compliance settings for Mac OS.
type MacOsCompliancePolicy struct {
	// Whether or not to require a password.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`

	// Indicates whether or not to block simple passwords.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`

	// Number of days before the password expires. Valid values 1 to 65535
	PasswordExpirationDays *int32 `json:"passwordExpirationDays,omitempty"`
	isExplicitNullPasswordExpirationDays bool `json:"-"`
	// Minimum length of password. Valid values 4 to 14
	PasswordMinimumLength *int32 `json:"passwordMinimumLength,omitempty"`
	isExplicitNullPasswordMinimumLength bool `json:"-"`
	// Minutes of inactivity before a password is required.
	PasswordMinutesOfInactivityBeforeLock *int32 `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	isExplicitNullPasswordMinutesOfInactivityBeforeLock bool `json:"-"`
	// Number of previous passwords to block. Valid values 1 to 24
	PasswordPreviousPasswordBlockCount *int32 `json:"passwordPreviousPasswordBlockCount,omitempty"`
	isExplicitNullPasswordPreviousPasswordBlockCount bool `json:"-"`
	// The number of character sets required in the password.
	PasswordMinimumCharacterSetCount *int32 `json:"passwordMinimumCharacterSetCount,omitempty"`
	isExplicitNullPasswordMinimumCharacterSetCount bool `json:"-"`
	// The required password type.
	PasswordRequiredType *AnyOfmicrosoftGraphRequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Minimum MacOS version.
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	isExplicitNullOsMinimumVersion bool `json:"-"`
	// Maximum MacOS version.
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	isExplicitNullOsMaximumVersion bool `json:"-"`
	// Require that devices have enabled system integrity protection.
	SystemIntegrityProtectionEnabled *bool `json:"systemIntegrityProtectionEnabled,omitempty"`

	// Require that devices have enabled device threat protection.
	DeviceThreatProtectionEnabled *bool `json:"deviceThreatProtectionEnabled,omitempty"`

	// Require Mobile Threat Protection minimum risk level to report noncompliance.
	DeviceThreatProtectionRequiredSecurityLevel *AnyOfmicrosoftGraphDeviceThreatProtectionLevel `json:"deviceThreatProtectionRequiredSecurityLevel,omitempty"`

	// Require encryption on Mac OS devices.
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`

	// Whether the firewall should be enabled or not.
	FirewallEnabled *bool `json:"firewallEnabled,omitempty"`

	// Corresponds to the “Block all incoming connections” option.
	FirewallBlockAllIncoming *bool `json:"firewallBlockAllIncoming,omitempty"`

	// Corresponds to “Enable stealth mode.”
	FirewallEnableStealthMode *bool `json:"firewallEnableStealthMode,omitempty"`

}

// GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetPasswordRequired() bool {
	if o == nil || o.PasswordRequired == nil {
		var ret bool
		return ret
	}
	return *o.PasswordRequired
}

// GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetPasswordRequiredOk() (bool, bool) {
	if o == nil || o.PasswordRequired == nil {
		var ret bool
		return ret, false
	}
	return *o.PasswordRequired, true
}

// HasPasswordRequired returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasPasswordRequired() bool {
	if o != nil && o.PasswordRequired != nil {
		return true
	}

	return false
}

// SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.
func (o *MacOsCompliancePolicy) SetPasswordRequired(v bool) {
	o.PasswordRequired = &v
}

// GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetPasswordBlockSimple() bool {
	if o == nil || o.PasswordBlockSimple == nil {
		var ret bool
		return ret
	}
	return *o.PasswordBlockSimple
}

// GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetPasswordBlockSimpleOk() (bool, bool) {
	if o == nil || o.PasswordBlockSimple == nil {
		var ret bool
		return ret, false
	}
	return *o.PasswordBlockSimple, true
}

// HasPasswordBlockSimple returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasPasswordBlockSimple() bool {
	if o != nil && o.PasswordBlockSimple != nil {
		return true
	}

	return false
}

// SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.
func (o *MacOsCompliancePolicy) SetPasswordBlockSimple(v bool) {
	o.PasswordBlockSimple = &v
}

// GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetPasswordExpirationDays() int32 {
	if o == nil || o.PasswordExpirationDays == nil {
		var ret int32
		return ret
	}
	return *o.PasswordExpirationDays
}

// GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetPasswordExpirationDaysOk() (int32, bool) {
	if o == nil || o.PasswordExpirationDays == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordExpirationDays, true
}

// HasPasswordExpirationDays returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasPasswordExpirationDays() bool {
	if o != nil && o.PasswordExpirationDays != nil {
		return true
	}

	return false
}

// SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.
func (o *MacOsCompliancePolicy) SetPasswordExpirationDays(v int32) {
	o.PasswordExpirationDays = &v
}

// SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordExpirationDays value is set to nil even if false is passed
func (o *MacOsCompliancePolicy) SetPasswordExpirationDaysExplicitNull(b bool) {
	o.PasswordExpirationDays = nil
	o.isExplicitNullPasswordExpirationDays = b
}
// GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetPasswordMinimumLength() int32 {
	if o == nil || o.PasswordMinimumLength == nil {
		var ret int32
		return ret
	}
	return *o.PasswordMinimumLength
}

// GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetPasswordMinimumLengthOk() (int32, bool) {
	if o == nil || o.PasswordMinimumLength == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordMinimumLength, true
}

// HasPasswordMinimumLength returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasPasswordMinimumLength() bool {
	if o != nil && o.PasswordMinimumLength != nil {
		return true
	}

	return false
}

// SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.
func (o *MacOsCompliancePolicy) SetPasswordMinimumLength(v int32) {
	o.PasswordMinimumLength = &v
}

// SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordMinimumLength value is set to nil even if false is passed
func (o *MacOsCompliancePolicy) SetPasswordMinimumLengthExplicitNull(b bool) {
	o.PasswordMinimumLength = nil
	o.isExplicitNullPasswordMinimumLength = b
}
// GetPasswordMinutesOfInactivityBeforeLock returns the PasswordMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock() int32 {
	if o == nil || o.PasswordMinutesOfInactivityBeforeLock == nil {
		var ret int32
		return ret
	}
	return *o.PasswordMinutesOfInactivityBeforeLock
}

// GetPasswordMinutesOfInactivityBeforeLockOk returns a tuple with the PasswordMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLockOk() (int32, bool) {
	if o == nil || o.PasswordMinutesOfInactivityBeforeLock == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordMinutesOfInactivityBeforeLock, true
}

// HasPasswordMinutesOfInactivityBeforeLock returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasPasswordMinutesOfInactivityBeforeLock() bool {
	if o != nil && o.PasswordMinutesOfInactivityBeforeLock != nil {
		return true
	}

	return false
}

// SetPasswordMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeLock field.
func (o *MacOsCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(v int32) {
	o.PasswordMinutesOfInactivityBeforeLock = &v
}

// SetPasswordMinutesOfInactivityBeforeLockExplicitNull (un)sets PasswordMinutesOfInactivityBeforeLock to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordMinutesOfInactivityBeforeLock value is set to nil even if false is passed
func (o *MacOsCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLockExplicitNull(b bool) {
	o.PasswordMinutesOfInactivityBeforeLock = nil
	o.isExplicitNullPasswordMinutesOfInactivityBeforeLock = b
}
// GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetPasswordPreviousPasswordBlockCount() int32 {
	if o == nil || o.PasswordPreviousPasswordBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.PasswordPreviousPasswordBlockCount
}

// GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetPasswordPreviousPasswordBlockCountOk() (int32, bool) {
	if o == nil || o.PasswordPreviousPasswordBlockCount == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordPreviousPasswordBlockCount, true
}

// HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasPasswordPreviousPasswordBlockCount() bool {
	if o != nil && o.PasswordPreviousPasswordBlockCount != nil {
		return true
	}

	return false
}

// SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.
func (o *MacOsCompliancePolicy) SetPasswordPreviousPasswordBlockCount(v int32) {
	o.PasswordPreviousPasswordBlockCount = &v
}

// SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
func (o *MacOsCompliancePolicy) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool) {
	o.PasswordPreviousPasswordBlockCount = nil
	o.isExplicitNullPasswordPreviousPasswordBlockCount = b
}
// GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetPasswordMinimumCharacterSetCount() int32 {
	if o == nil || o.PasswordMinimumCharacterSetCount == nil {
		var ret int32
		return ret
	}
	return *o.PasswordMinimumCharacterSetCount
}

// GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetPasswordMinimumCharacterSetCountOk() (int32, bool) {
	if o == nil || o.PasswordMinimumCharacterSetCount == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordMinimumCharacterSetCount, true
}

// HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasPasswordMinimumCharacterSetCount() bool {
	if o != nil && o.PasswordMinimumCharacterSetCount != nil {
		return true
	}

	return false
}

// SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.
func (o *MacOsCompliancePolicy) SetPasswordMinimumCharacterSetCount(v int32) {
	o.PasswordMinimumCharacterSetCount = &v
}

// SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
func (o *MacOsCompliancePolicy) SetPasswordMinimumCharacterSetCountExplicitNull(b bool) {
	o.PasswordMinimumCharacterSetCount = nil
	o.isExplicitNullPasswordMinimumCharacterSetCount = b
}
// GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType {
	if o == nil || o.PasswordRequiredType == nil {
		var ret AnyOfmicrosoftGraphRequiredPasswordType
		return ret
	}
	return *o.PasswordRequiredType
}

// GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool) {
	if o == nil || o.PasswordRequiredType == nil {
		var ret AnyOfmicrosoftGraphRequiredPasswordType
		return ret, false
	}
	return *o.PasswordRequiredType, true
}

// HasPasswordRequiredType returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasPasswordRequiredType() bool {
	if o != nil && o.PasswordRequiredType != nil {
		return true
	}

	return false
}

// SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.
func (o *MacOsCompliancePolicy) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType) {
	o.PasswordRequiredType = &v
}

// GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetOsMinimumVersion() string {
	if o == nil || o.OsMinimumVersion == nil {
		var ret string
		return ret
	}
	return *o.OsMinimumVersion
}

// GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetOsMinimumVersionOk() (string, bool) {
	if o == nil || o.OsMinimumVersion == nil {
		var ret string
		return ret, false
	}
	return *o.OsMinimumVersion, true
}

// HasOsMinimumVersion returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasOsMinimumVersion() bool {
	if o != nil && o.OsMinimumVersion != nil {
		return true
	}

	return false
}

// SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.
func (o *MacOsCompliancePolicy) SetOsMinimumVersion(v string) {
	o.OsMinimumVersion = &v
}

// SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OsMinimumVersion value is set to nil even if false is passed
func (o *MacOsCompliancePolicy) SetOsMinimumVersionExplicitNull(b bool) {
	o.OsMinimumVersion = nil
	o.isExplicitNullOsMinimumVersion = b
}
// GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetOsMaximumVersion() string {
	if o == nil || o.OsMaximumVersion == nil {
		var ret string
		return ret
	}
	return *o.OsMaximumVersion
}

// GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetOsMaximumVersionOk() (string, bool) {
	if o == nil || o.OsMaximumVersion == nil {
		var ret string
		return ret, false
	}
	return *o.OsMaximumVersion, true
}

// HasOsMaximumVersion returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasOsMaximumVersion() bool {
	if o != nil && o.OsMaximumVersion != nil {
		return true
	}

	return false
}

// SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.
func (o *MacOsCompliancePolicy) SetOsMaximumVersion(v string) {
	o.OsMaximumVersion = &v
}

// SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OsMaximumVersion value is set to nil even if false is passed
func (o *MacOsCompliancePolicy) SetOsMaximumVersionExplicitNull(b bool) {
	o.OsMaximumVersion = nil
	o.isExplicitNullOsMaximumVersion = b
}
// GetSystemIntegrityProtectionEnabled returns the SystemIntegrityProtectionEnabled field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetSystemIntegrityProtectionEnabled() bool {
	if o == nil || o.SystemIntegrityProtectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SystemIntegrityProtectionEnabled
}

// GetSystemIntegrityProtectionEnabledOk returns a tuple with the SystemIntegrityProtectionEnabled field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetSystemIntegrityProtectionEnabledOk() (bool, bool) {
	if o == nil || o.SystemIntegrityProtectionEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.SystemIntegrityProtectionEnabled, true
}

// HasSystemIntegrityProtectionEnabled returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasSystemIntegrityProtectionEnabled() bool {
	if o != nil && o.SystemIntegrityProtectionEnabled != nil {
		return true
	}

	return false
}

// SetSystemIntegrityProtectionEnabled gets a reference to the given bool and assigns it to the SystemIntegrityProtectionEnabled field.
func (o *MacOsCompliancePolicy) SetSystemIntegrityProtectionEnabled(v bool) {
	o.SystemIntegrityProtectionEnabled = &v
}

// GetDeviceThreatProtectionEnabled returns the DeviceThreatProtectionEnabled field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetDeviceThreatProtectionEnabled() bool {
	if o == nil || o.DeviceThreatProtectionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DeviceThreatProtectionEnabled
}

// GetDeviceThreatProtectionEnabledOk returns a tuple with the DeviceThreatProtectionEnabled field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetDeviceThreatProtectionEnabledOk() (bool, bool) {
	if o == nil || o.DeviceThreatProtectionEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.DeviceThreatProtectionEnabled, true
}

// HasDeviceThreatProtectionEnabled returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasDeviceThreatProtectionEnabled() bool {
	if o != nil && o.DeviceThreatProtectionEnabled != nil {
		return true
	}

	return false
}

// SetDeviceThreatProtectionEnabled gets a reference to the given bool and assigns it to the DeviceThreatProtectionEnabled field.
func (o *MacOsCompliancePolicy) SetDeviceThreatProtectionEnabled(v bool) {
	o.DeviceThreatProtectionEnabled = &v
}

// GetDeviceThreatProtectionRequiredSecurityLevel returns the DeviceThreatProtectionRequiredSecurityLevel field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel() AnyOfmicrosoftGraphDeviceThreatProtectionLevel {
	if o == nil || o.DeviceThreatProtectionRequiredSecurityLevel == nil {
		var ret AnyOfmicrosoftGraphDeviceThreatProtectionLevel
		return ret
	}
	return *o.DeviceThreatProtectionRequiredSecurityLevel
}

// GetDeviceThreatProtectionRequiredSecurityLevelOk returns a tuple with the DeviceThreatProtectionRequiredSecurityLevel field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevelOk() (AnyOfmicrosoftGraphDeviceThreatProtectionLevel, bool) {
	if o == nil || o.DeviceThreatProtectionRequiredSecurityLevel == nil {
		var ret AnyOfmicrosoftGraphDeviceThreatProtectionLevel
		return ret, false
	}
	return *o.DeviceThreatProtectionRequiredSecurityLevel, true
}

// HasDeviceThreatProtectionRequiredSecurityLevel returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasDeviceThreatProtectionRequiredSecurityLevel() bool {
	if o != nil && o.DeviceThreatProtectionRequiredSecurityLevel != nil {
		return true
	}

	return false
}

// SetDeviceThreatProtectionRequiredSecurityLevel gets a reference to the given AnyOfmicrosoftGraphDeviceThreatProtectionLevel and assigns it to the DeviceThreatProtectionRequiredSecurityLevel field.
func (o *MacOsCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(v AnyOfmicrosoftGraphDeviceThreatProtectionLevel) {
	o.DeviceThreatProtectionRequiredSecurityLevel = &v
}

// GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetStorageRequireEncryption() bool {
	if o == nil || o.StorageRequireEncryption == nil {
		var ret bool
		return ret
	}
	return *o.StorageRequireEncryption
}

// GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetStorageRequireEncryptionOk() (bool, bool) {
	if o == nil || o.StorageRequireEncryption == nil {
		var ret bool
		return ret, false
	}
	return *o.StorageRequireEncryption, true
}

// HasStorageRequireEncryption returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasStorageRequireEncryption() bool {
	if o != nil && o.StorageRequireEncryption != nil {
		return true
	}

	return false
}

// SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.
func (o *MacOsCompliancePolicy) SetStorageRequireEncryption(v bool) {
	o.StorageRequireEncryption = &v
}

// GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetFirewallEnabled() bool {
	if o == nil || o.FirewallEnabled == nil {
		var ret bool
		return ret
	}
	return *o.FirewallEnabled
}

// GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetFirewallEnabledOk() (bool, bool) {
	if o == nil || o.FirewallEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.FirewallEnabled, true
}

// HasFirewallEnabled returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasFirewallEnabled() bool {
	if o != nil && o.FirewallEnabled != nil {
		return true
	}

	return false
}

// SetFirewallEnabled gets a reference to the given bool and assigns it to the FirewallEnabled field.
func (o *MacOsCompliancePolicy) SetFirewallEnabled(v bool) {
	o.FirewallEnabled = &v
}

// GetFirewallBlockAllIncoming returns the FirewallBlockAllIncoming field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetFirewallBlockAllIncoming() bool {
	if o == nil || o.FirewallBlockAllIncoming == nil {
		var ret bool
		return ret
	}
	return *o.FirewallBlockAllIncoming
}

// GetFirewallBlockAllIncomingOk returns a tuple with the FirewallBlockAllIncoming field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetFirewallBlockAllIncomingOk() (bool, bool) {
	if o == nil || o.FirewallBlockAllIncoming == nil {
		var ret bool
		return ret, false
	}
	return *o.FirewallBlockAllIncoming, true
}

// HasFirewallBlockAllIncoming returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasFirewallBlockAllIncoming() bool {
	if o != nil && o.FirewallBlockAllIncoming != nil {
		return true
	}

	return false
}

// SetFirewallBlockAllIncoming gets a reference to the given bool and assigns it to the FirewallBlockAllIncoming field.
func (o *MacOsCompliancePolicy) SetFirewallBlockAllIncoming(v bool) {
	o.FirewallBlockAllIncoming = &v
}

// GetFirewallEnableStealthMode returns the FirewallEnableStealthMode field if non-nil, zero value otherwise.
func (o *MacOsCompliancePolicy) GetFirewallEnableStealthMode() bool {
	if o == nil || o.FirewallEnableStealthMode == nil {
		var ret bool
		return ret
	}
	return *o.FirewallEnableStealthMode
}

// GetFirewallEnableStealthModeOk returns a tuple with the FirewallEnableStealthMode field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MacOsCompliancePolicy) GetFirewallEnableStealthModeOk() (bool, bool) {
	if o == nil || o.FirewallEnableStealthMode == nil {
		var ret bool
		return ret, false
	}
	return *o.FirewallEnableStealthMode, true
}

// HasFirewallEnableStealthMode returns a boolean if a field has been set.
func (o *MacOsCompliancePolicy) HasFirewallEnableStealthMode() bool {
	if o != nil && o.FirewallEnableStealthMode != nil {
		return true
	}

	return false
}

// SetFirewallEnableStealthMode gets a reference to the given bool and assigns it to the FirewallEnableStealthMode field.
func (o *MacOsCompliancePolicy) SetFirewallEnableStealthMode(v bool) {
	o.FirewallEnableStealthMode = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MacOsCompliancePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PasswordRequired != nil {
		toSerialize["passwordRequired"] = o.PasswordRequired
	}
	if o.PasswordBlockSimple != nil {
		toSerialize["passwordBlockSimple"] = o.PasswordBlockSimple
	}
	if o.PasswordExpirationDays == nil {
		if o.isExplicitNullPasswordExpirationDays {
			toSerialize["passwordExpirationDays"] = o.PasswordExpirationDays
		}
	} else {
		toSerialize["passwordExpirationDays"] = o.PasswordExpirationDays
	}
	if o.PasswordMinimumLength == nil {
		if o.isExplicitNullPasswordMinimumLength {
			toSerialize["passwordMinimumLength"] = o.PasswordMinimumLength
		}
	} else {
		toSerialize["passwordMinimumLength"] = o.PasswordMinimumLength
	}
	if o.PasswordMinutesOfInactivityBeforeLock == nil {
		if o.isExplicitNullPasswordMinutesOfInactivityBeforeLock {
			toSerialize["passwordMinutesOfInactivityBeforeLock"] = o.PasswordMinutesOfInactivityBeforeLock
		}
	} else {
		toSerialize["passwordMinutesOfInactivityBeforeLock"] = o.PasswordMinutesOfInactivityBeforeLock
	}
	if o.PasswordPreviousPasswordBlockCount == nil {
		if o.isExplicitNullPasswordPreviousPasswordBlockCount {
			toSerialize["passwordPreviousPasswordBlockCount"] = o.PasswordPreviousPasswordBlockCount
		}
	} else {
		toSerialize["passwordPreviousPasswordBlockCount"] = o.PasswordPreviousPasswordBlockCount
	}
	if o.PasswordMinimumCharacterSetCount == nil {
		if o.isExplicitNullPasswordMinimumCharacterSetCount {
			toSerialize["passwordMinimumCharacterSetCount"] = o.PasswordMinimumCharacterSetCount
		}
	} else {
		toSerialize["passwordMinimumCharacterSetCount"] = o.PasswordMinimumCharacterSetCount
	}
	if o.PasswordRequiredType != nil {
		toSerialize["passwordRequiredType"] = o.PasswordRequiredType
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
	if o.SystemIntegrityProtectionEnabled != nil {
		toSerialize["systemIntegrityProtectionEnabled"] = o.SystemIntegrityProtectionEnabled
	}
	if o.DeviceThreatProtectionEnabled != nil {
		toSerialize["deviceThreatProtectionEnabled"] = o.DeviceThreatProtectionEnabled
	}
	if o.DeviceThreatProtectionRequiredSecurityLevel != nil {
		toSerialize["deviceThreatProtectionRequiredSecurityLevel"] = o.DeviceThreatProtectionRequiredSecurityLevel
	}
	if o.StorageRequireEncryption != nil {
		toSerialize["storageRequireEncryption"] = o.StorageRequireEncryption
	}
	if o.FirewallEnabled != nil {
		toSerialize["firewallEnabled"] = o.FirewallEnabled
	}
	if o.FirewallBlockAllIncoming != nil {
		toSerialize["firewallBlockAllIncoming"] = o.FirewallBlockAllIncoming
	}
	if o.FirewallEnableStealthMode != nil {
		toSerialize["firewallEnableStealthMode"] = o.FirewallEnableStealthMode
	}
	return json.Marshal(toSerialize)
}

