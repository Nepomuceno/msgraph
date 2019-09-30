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
// MicrosoftGraphDriveRecipient struct for MicrosoftGraphDriveRecipient
type MicrosoftGraphDriveRecipient struct {
	Alias *string `json:"alias,omitempty"`
	isExplicitNullAlias bool `json:"-"`
	Email *string `json:"email,omitempty"`
	isExplicitNullEmail bool `json:"-"`
	ObjectId *string `json:"objectId,omitempty"`
	isExplicitNullObjectId bool `json:"-"`
}

// GetAlias returns the Alias field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDriveRecipient) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDriveRecipient) GetAliasOk() (string, bool) {
	if o == nil || o.Alias == nil {
		var ret string
		return ret, false
	}
	return *o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *MicrosoftGraphDriveRecipient) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *MicrosoftGraphDriveRecipient) SetAlias(v string) {
	o.Alias = &v
}

// SetAliasExplicitNull (un)sets Alias to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Alias value is set to nil even if false is passed
func (o *MicrosoftGraphDriveRecipient) SetAliasExplicitNull(b bool) {
	o.Alias = nil
	o.isExplicitNullAlias = b
}
// GetEmail returns the Email field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDriveRecipient) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDriveRecipient) GetEmailOk() (string, bool) {
	if o == nil || o.Email == nil {
		var ret string
		return ret, false
	}
	return *o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MicrosoftGraphDriveRecipient) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MicrosoftGraphDriveRecipient) SetEmail(v string) {
	o.Email = &v
}

// SetEmailExplicitNull (un)sets Email to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Email value is set to nil even if false is passed
func (o *MicrosoftGraphDriveRecipient) SetEmailExplicitNull(b bool) {
	o.Email = nil
	o.isExplicitNullEmail = b
}
// GetObjectId returns the ObjectId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDriveRecipient) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDriveRecipient) GetObjectIdOk() (string, bool) {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret, false
	}
	return *o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *MicrosoftGraphDriveRecipient) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *MicrosoftGraphDriveRecipient) SetObjectId(v string) {
	o.ObjectId = &v
}

// SetObjectIdExplicitNull (un)sets ObjectId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ObjectId value is set to nil even if false is passed
func (o *MicrosoftGraphDriveRecipient) SetObjectIdExplicitNull(b bool) {
	o.ObjectId = nil
	o.isExplicitNullObjectId = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphDriveRecipient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alias == nil {
		if o.isExplicitNullAlias {
			toSerialize["alias"] = o.Alias
		}
	} else {
		toSerialize["alias"] = o.Alias
	}
	if o.Email == nil {
		if o.isExplicitNullEmail {
			toSerialize["email"] = o.Email
		}
	} else {
		toSerialize["email"] = o.Email
	}
	if o.ObjectId == nil {
		if o.isExplicitNullObjectId {
			toSerialize["objectId"] = o.ObjectId
		}
	} else {
		toSerialize["objectId"] = o.ObjectId
	}
	return json.Marshal(toSerialize)
}

