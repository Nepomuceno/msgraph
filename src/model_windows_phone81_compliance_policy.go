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
// WindowsPhone81CompliancePolicy This class contains compliance settings for Windows 8.1 Mobile.
type WindowsPhone81CompliancePolicy struct {
	// Whether or not to block syncing the calendar.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`

	// Number of days before the password expires.
	PasswordExpirationDays *int32 `json:"passwordExpirationDays,omitempty"`
	isExplicitNullPasswordExpirationDays bool `json:"-"`
	// Minimum length of passwords.
	PasswordMinimumLength *int32 `json:"passwordMinimumLength,omitempty"`
	isExplicitNullPasswordMinimumLength bool `json:"-"`
	// Minutes of inactivity before a password is required.
	PasswordMinutesOfInactivityBeforeLock *int32 `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	isExplicitNullPasswordMinutesOfInactivityBeforeLock bool `json:"-"`
	// The number of character sets required in the password.
	PasswordMinimumCharacterSetCount *int32 `json:"passwordMinimumCharacterSetCount,omitempty"`
	isExplicitNullPasswordMinimumCharacterSetCount bool `json:"-"`
	// The required password type.
	PasswordRequiredType *AnyOfmicrosoftGraphRequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Number of previous passwords to block. Valid values 0 to 24
	PasswordPreviousPasswordBlockCount *int32 `json:"passwordPreviousPasswordBlockCount,omitempty"`
	isExplicitNullPasswordPreviousPasswordBlockCount bool `json:"-"`
	// Whether or not to require a password.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`

	// Minimum Windows Phone version.
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	isExplicitNullOsMinimumVersion bool `json:"-"`
	// Maximum Windows Phone version.
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	isExplicitNullOsMaximumVersion bool `json:"-"`
	// Require encryption on windows phone devices.
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`

}

// GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.
func (o *WindowsPhone81CompliancePolicy) GetPasswordBlockSimple() bool {
	if o == nil || o.PasswordBlockSimple == nil {
		var ret bool
		return ret
	}
	return *o.PasswordBlockSimple
}

// GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsPhone81CompliancePolicy) GetPasswordBlockSimpleOk() (bool, bool) {
	if o == nil || o.PasswordBlockSimple == nil {
		var ret bool
		return ret, false
	}
	return *o.PasswordBlockSimple, true
}

// HasPasswordBlockSimple returns a boolean if a field has been set.
func (o *WindowsPhone81CompliancePolicy) HasPasswordBlockSimple() bool {
	if o != nil && o.PasswordBlockSimple != nil {
		return true
	}

	return false
}

// SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.
func (o *WindowsPhone81CompliancePolicy) SetPasswordBlockSimple(v bool) {
	o.PasswordBlockSimple = &v
}

// GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.
func (o *WindowsPhone81CompliancePolicy) GetPasswordExpirationDays() int32 {
	if o == nil || o.PasswordExpirationDays == nil {
		var ret int32
		return ret
	}
	return *o.PasswordExpirationDays
}

// GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsPhone81CompliancePolicy) GetPasswordExpirationDaysOk() (int32, bool) {
	if o == nil || o.PasswordExpirationDays == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordExpirationDays, true
}

// HasPasswordExpirationDays returns a boolean if a field has been set.
func (o *WindowsPhone81CompliancePolicy) HasPasswordExpirationDays() bool {
	if o != nil && o.PasswordExpirationDays != nil {
		return true
	}

	return false
}

// SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.
func (o *WindowsPhone81CompliancePolicy) SetPasswordExpirationDays(v int32) {
	o.PasswordExpirationDays = &v
}

// SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordExpirationDays value is set to nil even if false is passed
func (o *WindowsPhone81CompliancePolicy) SetPasswordExpirationDaysExplicitNull(b bool) {
	o.PasswordExpirationDays = nil
	o.isExplicitNullPasswordExpirationDays = b
}
// GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.
func (o *WindowsPhone81CompliancePolicy) GetPasswordMinimumLength() int32 {
	if o == nil || o.PasswordMinimumLength == nil {
		var ret int32
		return ret
	}
	return *o.PasswordMinimumLength
}

// GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsPhone81CompliancePolicy) GetPasswordMinimumLengthOk() (int32, bool) {
	if o == nil || o.PasswordMinimumLength == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordMinimumLength, true
}

// HasPasswordMinimumLength returns a boolean if a field has been set.
func (o *WindowsPhone81CompliancePolicy) HasPasswordMinimumLength() bool {
	if o != nil && o.PasswordMinimumLength != nil {
		return true
	}

	return false
}

// SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.
func (o *WindowsPhone81CompliancePolicy) SetPasswordMinimumLength(v int32) {
	o.PasswordMinimumLength = &v
}

// SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordMinimumLength value is set to nil even if false is passed
func (o *WindowsPhone81CompliancePolicy) SetPasswordMinimumLengthExplicitNull(b bool) {
	o.PasswordMinimumLength = nil
	o.isExplicitNullPasswordMinimumLength = b
}
// GetPasswordMinutesOfInactivityBeforeLock returns the PasswordMinutesOfInactivityBeforeLock field if non-nil, zero value otherwise.
func (o *WindowsPhone81CompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock() int32 {
	if o == nil || o.PasswordMinutesOfInactivityBeforeLock == nil {
		var ret int32
		return ret
	}
	return *o.PasswordMinutesOfInactivityBeforeLock
}

// GetPasswordMinutesOfInactivityBeforeLockOk returns a tuple with the PasswordMinutesOfInactivityBeforeLock field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsPhone81CompliancePolicy) GetPasswordMinutesOfInactivityBeforeLockOk() (int32, bool) {
	if o == nil || o.PasswordMinutesOfInactivityBeforeLock == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordMinutesOfInactivityBeforeLock, true
}

// HasPasswordMinutesOfInactivityBeforeLock returns a boolean if a field has been set.
func (o *WindowsPhone81CompliancePolicy) HasPasswordMinutesOfInactivityBeforeLock() bool {
	if o != nil && o.PasswordMinutesOfInactivityBeforeLock != nil {
		return true
	}

	return false
}

// SetPasswordMinutesOfInactivityBeforeLock gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeLock field.
func (o *WindowsPhone81CompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(v int32) {
	o.PasswordMinutesOfInactivityBeforeLock = &v
}

// SetPasswordMinutesOfInactivityBeforeLockExplicitNull (un)sets PasswordMinutesOfInactivityBeforeLock to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordMinutesOfInactivityBeforeLock value is set to nil even if false is passed
func (o *WindowsPhone81CompliancePolicy) SetPasswordMinutesOfInactivityBeforeLockExplicitNull(b bool) {
	o.PasswordMinutesOfInactivityBeforeLock = nil
	o.isExplicitNullPasswordMinutesOfInactivityBeforeLock = b
}
// GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.
func (o *WindowsPhone81CompliancePolicy) GetPasswordMinimumCharacterSetCount() int32 {
	if o == nil || o.PasswordMinimumCharacterSetCount == nil {
		var ret int32
		return ret
	}
	return *o.PasswordMinimumCharacterSetCount
}

// GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsPhone81CompliancePolicy) GetPasswordMinimumCharacterSetCountOk() (int32, bool) {
	if o == nil || o.PasswordMinimumCharacterSetCount == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordMinimumCharacterSetCount, true
}

// HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.
func (o *WindowsPhone81CompliancePolicy) HasPasswordMinimumCharacterSetCount() bool {
	if o != nil && o.PasswordMinimumCharacterSetCount != nil {
		return true
	}

	return false
}

// SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.
func (o *WindowsPhone81CompliancePolicy) SetPasswordMinimumCharacterSetCount(v int32) {
	o.PasswordMinimumCharacterSetCount = &v
}

// SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
func (o *WindowsPhone81CompliancePolicy) SetPasswordMinimumCharacterSetCountExplicitNull(b bool) {
	o.PasswordMinimumCharacterSetCount = nil
	o.isExplicitNullPasswordMinimumCharacterSetCount = b
}
// GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.
func (o *WindowsPhone81CompliancePolicy) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType {
	if o == nil || o.PasswordRequiredType == nil {
		var ret AnyOfmicrosoftGraphRequiredPasswordType
		return ret
	}
	return *o.PasswordRequiredType
}

// GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsPhone81CompliancePolicy) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool) {
	if o == nil || o.PasswordRequiredType == nil {
		var ret AnyOfmicrosoftGraphRequiredPasswordType
		return ret, false
	}
	return *o.PasswordRequiredType, true
}

// HasPasswordRequiredType returns a boolean if a field has been set.
func (o *WindowsPhone81CompliancePolicy) HasPasswordRequiredType() bool {
	if o != nil && o.PasswordRequiredType != nil {
		return true
	}

	return false
}

// SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.
func (o *WindowsPhone81CompliancePolicy) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType) {
	o.PasswordRequiredType = &v
}

// GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.
func (o *WindowsPhone81CompliancePolicy) GetPasswordPreviousPasswordBlockCount() int32 {
	if o == nil || o.PasswordPreviousPasswordBlockCount == nil {
		var ret int32
		return ret
	}
	return *o.PasswordPreviousPasswordBlockCount
}

// GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsPhone81CompliancePolicy) GetPasswordPreviousPasswordBlockCountOk() (int32, bool) {
	if o == nil || o.PasswordPreviousPasswordBlockCount == nil {
		var ret int32
		return ret, false
	}
	return *o.PasswordPreviousPasswordBlockCount, true
}

// HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.
func (o *WindowsPhone81CompliancePolicy) HasPasswordPreviousPasswordBlockCount() bool {
	if o != nil && o.PasswordPreviousPasswordBlockCount != nil {
		return true
	}

	return false
}

// SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.
func (o *WindowsPhone81CompliancePolicy) SetPasswordPreviousPasswordBlockCount(v int32) {
	o.PasswordPreviousPasswordBlockCount = &v
}

// SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
func (o *WindowsPhone81CompliancePolicy) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool) {
	o.PasswordPreviousPasswordBlockCount = nil
	o.isExplicitNullPasswordPreviousPasswordBlockCount = b
}
// GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.
func (o *WindowsPhone81CompliancePolicy) GetPasswordRequired() bool {
	if o == nil || o.PasswordRequired == nil {
		var ret bool
		return ret
	}
	return *o.PasswordRequired
}

// GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsPhone81CompliancePolicy) GetPasswordRequiredOk() (bool, bool) {
	if o == nil || o.PasswordRequired == nil {
		var ret bool
		return ret, false
	}
	return *o.PasswordRequired, true
}

// HasPasswordRequired returns a boolean if a field has been set.
func (o *WindowsPhone81CompliancePolicy) HasPasswordRequired() bool {
	if o != nil && o.PasswordRequired != nil {
		return true
	}

	return false
}

// SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.
func (o *WindowsPhone81CompliancePolicy) SetPasswordRequired(v bool) {
	o.PasswordRequired = &v
}

// GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.
func (o *WindowsPhone81CompliancePolicy) GetOsMinimumVersion() string {
	if o == nil || o.OsMinimumVersion == nil {
		var ret string
		return ret
	}
	return *o.OsMinimumVersion
}

// GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsPhone81CompliancePolicy) GetOsMinimumVersionOk() (string, bool) {
	if o == nil || o.OsMinimumVersion == nil {
		var ret string
		return ret, false
	}
	return *o.OsMinimumVersion, true
}

// HasOsMinimumVersion returns a boolean if a field has been set.
func (o *WindowsPhone81CompliancePolicy) HasOsMinimumVersion() bool {
	if o != nil && o.OsMinimumVersion != nil {
		return true
	}

	return false
}

// SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.
func (o *WindowsPhone81CompliancePolicy) SetOsMinimumVersion(v string) {
	o.OsMinimumVersion = &v
}

// SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OsMinimumVersion value is set to nil even if false is passed
func (o *WindowsPhone81CompliancePolicy) SetOsMinimumVersionExplicitNull(b bool) {
	o.OsMinimumVersion = nil
	o.isExplicitNullOsMinimumVersion = b
}
// GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.
func (o *WindowsPhone81CompliancePolicy) GetOsMaximumVersion() string {
	if o == nil || o.OsMaximumVersion == nil {
		var ret string
		return ret
	}
	return *o.OsMaximumVersion
}

// GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsPhone81CompliancePolicy) GetOsMaximumVersionOk() (string, bool) {
	if o == nil || o.OsMaximumVersion == nil {
		var ret string
		return ret, false
	}
	return *o.OsMaximumVersion, true
}

// HasOsMaximumVersion returns a boolean if a field has been set.
func (o *WindowsPhone81CompliancePolicy) HasOsMaximumVersion() bool {
	if o != nil && o.OsMaximumVersion != nil {
		return true
	}

	return false
}

// SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.
func (o *WindowsPhone81CompliancePolicy) SetOsMaximumVersion(v string) {
	o.OsMaximumVersion = &v
}

// SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OsMaximumVersion value is set to nil even if false is passed
func (o *WindowsPhone81CompliancePolicy) SetOsMaximumVersionExplicitNull(b bool) {
	o.OsMaximumVersion = nil
	o.isExplicitNullOsMaximumVersion = b
}
// GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.
func (o *WindowsPhone81CompliancePolicy) GetStorageRequireEncryption() bool {
	if o == nil || o.StorageRequireEncryption == nil {
		var ret bool
		return ret
	}
	return *o.StorageRequireEncryption
}

// GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsPhone81CompliancePolicy) GetStorageRequireEncryptionOk() (bool, bool) {
	if o == nil || o.StorageRequireEncryption == nil {
		var ret bool
		return ret, false
	}
	return *o.StorageRequireEncryption, true
}

// HasStorageRequireEncryption returns a boolean if a field has been set.
func (o *WindowsPhone81CompliancePolicy) HasStorageRequireEncryption() bool {
	if o != nil && o.StorageRequireEncryption != nil {
		return true
	}

	return false
}

// SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.
func (o *WindowsPhone81CompliancePolicy) SetStorageRequireEncryption(v bool) {
	o.StorageRequireEncryption = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o WindowsPhone81CompliancePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.PasswordPreviousPasswordBlockCount == nil {
		if o.isExplicitNullPasswordPreviousPasswordBlockCount {
			toSerialize["passwordPreviousPasswordBlockCount"] = o.PasswordPreviousPasswordBlockCount
		}
	} else {
		toSerialize["passwordPreviousPasswordBlockCount"] = o.PasswordPreviousPasswordBlockCount
	}
	if o.PasswordRequired != nil {
		toSerialize["passwordRequired"] = o.PasswordRequired
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
	if o.StorageRequireEncryption != nil {
		toSerialize["storageRequireEncryption"] = o.StorageRequireEncryption
	}
	return json.Marshal(toSerialize)
}

