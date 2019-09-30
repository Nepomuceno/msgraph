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
// MicrosoftGraphBitLockerRemovableDrivePolicy struct for MicrosoftGraphBitLockerRemovableDrivePolicy
type MicrosoftGraphBitLockerRemovableDrivePolicy struct {
	// Select the encryption method for removable  drives.
	EncryptionMethod *AnyOfmicrosoftGraphBitLockerEncryptionMethod `json:"encryptionMethod,omitempty"`
	isExplicitNullEncryptionMethod bool `json:"-"`
	// Indicates whether to block write access to devices configured in another organization.  If requireEncryptionForWriteAccess is false, this value does not affect.
	RequireEncryptionForWriteAccess *bool `json:"requireEncryptionForWriteAccess,omitempty"`

	// This policy setting determines whether BitLocker protection is required for removable data drives to be writable on a computer.
	BlockCrossOrganizationWriteAccess *bool `json:"blockCrossOrganizationWriteAccess,omitempty"`

}

// GetEncryptionMethod returns the EncryptionMethod field if non-nil, zero value otherwise.
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) GetEncryptionMethod() AnyOfmicrosoftGraphBitLockerEncryptionMethod {
	if o == nil || o.EncryptionMethod == nil {
		var ret AnyOfmicrosoftGraphBitLockerEncryptionMethod
		return ret
	}
	return *o.EncryptionMethod
}

// GetEncryptionMethodOk returns a tuple with the EncryptionMethod field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) GetEncryptionMethodOk() (AnyOfmicrosoftGraphBitLockerEncryptionMethod, bool) {
	if o == nil || o.EncryptionMethod == nil {
		var ret AnyOfmicrosoftGraphBitLockerEncryptionMethod
		return ret, false
	}
	return *o.EncryptionMethod, true
}

// HasEncryptionMethod returns a boolean if a field has been set.
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) HasEncryptionMethod() bool {
	if o != nil && o.EncryptionMethod != nil {
		return true
	}

	return false
}

// SetEncryptionMethod gets a reference to the given AnyOfmicrosoftGraphBitLockerEncryptionMethod and assigns it to the EncryptionMethod field.
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) SetEncryptionMethod(v AnyOfmicrosoftGraphBitLockerEncryptionMethod) {
	o.EncryptionMethod = &v
}

// SetEncryptionMethodExplicitNull (un)sets EncryptionMethod to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The EncryptionMethod value is set to nil even if false is passed
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) SetEncryptionMethodExplicitNull(b bool) {
	o.EncryptionMethod = nil
	o.isExplicitNullEncryptionMethod = b
}
// GetRequireEncryptionForWriteAccess returns the RequireEncryptionForWriteAccess field if non-nil, zero value otherwise.
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) GetRequireEncryptionForWriteAccess() bool {
	if o == nil || o.RequireEncryptionForWriteAccess == nil {
		var ret bool
		return ret
	}
	return *o.RequireEncryptionForWriteAccess
}

// GetRequireEncryptionForWriteAccessOk returns a tuple with the RequireEncryptionForWriteAccess field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) GetRequireEncryptionForWriteAccessOk() (bool, bool) {
	if o == nil || o.RequireEncryptionForWriteAccess == nil {
		var ret bool
		return ret, false
	}
	return *o.RequireEncryptionForWriteAccess, true
}

// HasRequireEncryptionForWriteAccess returns a boolean if a field has been set.
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) HasRequireEncryptionForWriteAccess() bool {
	if o != nil && o.RequireEncryptionForWriteAccess != nil {
		return true
	}

	return false
}

// SetRequireEncryptionForWriteAccess gets a reference to the given bool and assigns it to the RequireEncryptionForWriteAccess field.
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) SetRequireEncryptionForWriteAccess(v bool) {
	o.RequireEncryptionForWriteAccess = &v
}

// GetBlockCrossOrganizationWriteAccess returns the BlockCrossOrganizationWriteAccess field if non-nil, zero value otherwise.
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) GetBlockCrossOrganizationWriteAccess() bool {
	if o == nil || o.BlockCrossOrganizationWriteAccess == nil {
		var ret bool
		return ret
	}
	return *o.BlockCrossOrganizationWriteAccess
}

// GetBlockCrossOrganizationWriteAccessOk returns a tuple with the BlockCrossOrganizationWriteAccess field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) GetBlockCrossOrganizationWriteAccessOk() (bool, bool) {
	if o == nil || o.BlockCrossOrganizationWriteAccess == nil {
		var ret bool
		return ret, false
	}
	return *o.BlockCrossOrganizationWriteAccess, true
}

// HasBlockCrossOrganizationWriteAccess returns a boolean if a field has been set.
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) HasBlockCrossOrganizationWriteAccess() bool {
	if o != nil && o.BlockCrossOrganizationWriteAccess != nil {
		return true
	}

	return false
}

// SetBlockCrossOrganizationWriteAccess gets a reference to the given bool and assigns it to the BlockCrossOrganizationWriteAccess field.
func (o *MicrosoftGraphBitLockerRemovableDrivePolicy) SetBlockCrossOrganizationWriteAccess(v bool) {
	o.BlockCrossOrganizationWriteAccess = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphBitLockerRemovableDrivePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EncryptionMethod == nil {
		if o.isExplicitNullEncryptionMethod {
			toSerialize["encryptionMethod"] = o.EncryptionMethod
		}
	} else {
		toSerialize["encryptionMethod"] = o.EncryptionMethod
	}
	if o.RequireEncryptionForWriteAccess != nil {
		toSerialize["requireEncryptionForWriteAccess"] = o.RequireEncryptionForWriteAccess
	}
	if o.BlockCrossOrganizationWriteAccess != nil {
		toSerialize["blockCrossOrganizationWriteAccess"] = o.BlockCrossOrganizationWriteAccess
	}
	return json.Marshal(toSerialize)
}


