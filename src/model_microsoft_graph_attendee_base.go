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
// MicrosoftGraphAttendeeBase struct for MicrosoftGraphAttendeeBase
type MicrosoftGraphAttendeeBase struct {
	EmailAddress *AnyOfmicrosoftGraphEmailAddress `json:"emailAddress,omitempty"`
	isExplicitNullEmailAddress bool `json:"-"`
	Type *AnyOfmicrosoftGraphAttendeeType `json:"type,omitempty"`
	isExplicitNullType bool `json:"-"`
}

// GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAttendeeBase) GetEmailAddress() AnyOfmicrosoftGraphEmailAddress {
	if o == nil || o.EmailAddress == nil {
		var ret AnyOfmicrosoftGraphEmailAddress
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAttendeeBase) GetEmailAddressOk() (AnyOfmicrosoftGraphEmailAddress, bool) {
	if o == nil || o.EmailAddress == nil {
		var ret AnyOfmicrosoftGraphEmailAddress
		return ret, false
	}
	return *o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphAttendeeBase) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given AnyOfmicrosoftGraphEmailAddress and assigns it to the EmailAddress field.
func (o *MicrosoftGraphAttendeeBase) SetEmailAddress(v AnyOfmicrosoftGraphEmailAddress) {
	o.EmailAddress = &v
}

// SetEmailAddressExplicitNull (un)sets EmailAddress to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The EmailAddress value is set to nil even if false is passed
func (o *MicrosoftGraphAttendeeBase) SetEmailAddressExplicitNull(b bool) {
	o.EmailAddress = nil
	o.isExplicitNullEmailAddress = b
}
// GetType returns the Type field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAttendeeBase) GetType() AnyOfmicrosoftGraphAttendeeType {
	if o == nil || o.Type == nil {
		var ret AnyOfmicrosoftGraphAttendeeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAttendeeBase) GetTypeOk() (AnyOfmicrosoftGraphAttendeeType, bool) {
	if o == nil || o.Type == nil {
		var ret AnyOfmicrosoftGraphAttendeeType
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphAttendeeBase) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given AnyOfmicrosoftGraphAttendeeType and assigns it to the Type field.
func (o *MicrosoftGraphAttendeeBase) SetType(v AnyOfmicrosoftGraphAttendeeType) {
	o.Type = &v
}

// SetTypeExplicitNull (un)sets Type to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Type value is set to nil even if false is passed
func (o *MicrosoftGraphAttendeeBase) SetTypeExplicitNull(b bool) {
	o.Type = nil
	o.isExplicitNullType = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphAttendeeBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailAddress == nil {
		if o.isExplicitNullEmailAddress {
			toSerialize["emailAddress"] = o.EmailAddress
		}
	} else {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if o.Type == nil {
		if o.isExplicitNullType {
			toSerialize["type"] = o.Type
		}
	} else {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}


