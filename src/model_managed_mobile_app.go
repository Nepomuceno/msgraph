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
// ManagedMobileApp The identifier for the deployment an app.
type ManagedMobileApp struct {
	// The identifier for an app with it's operating system type.
	MobileAppIdentifier *AnyOfobject `json:"mobileAppIdentifier,omitempty"`
	isExplicitNullMobileAppIdentifier bool `json:"-"`
	// Version of the entity.
	Version *string `json:"version,omitempty"`
	isExplicitNullVersion bool `json:"-"`
}

// GetMobileAppIdentifier returns the MobileAppIdentifier field if non-nil, zero value otherwise.
func (o *ManagedMobileApp) GetMobileAppIdentifier() AnyOfobject {
	if o == nil || o.MobileAppIdentifier == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.MobileAppIdentifier
}

// GetMobileAppIdentifierOk returns a tuple with the MobileAppIdentifier field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMobileApp) GetMobileAppIdentifierOk() (AnyOfobject, bool) {
	if o == nil || o.MobileAppIdentifier == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.MobileAppIdentifier, true
}

// HasMobileAppIdentifier returns a boolean if a field has been set.
func (o *ManagedMobileApp) HasMobileAppIdentifier() bool {
	if o != nil && o.MobileAppIdentifier != nil {
		return true
	}

	return false
}

// SetMobileAppIdentifier gets a reference to the given AnyOfobject and assigns it to the MobileAppIdentifier field.
func (o *ManagedMobileApp) SetMobileAppIdentifier(v AnyOfobject) {
	o.MobileAppIdentifier = &v
}

// SetMobileAppIdentifierExplicitNull (un)sets MobileAppIdentifier to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The MobileAppIdentifier value is set to nil even if false is passed
func (o *ManagedMobileApp) SetMobileAppIdentifierExplicitNull(b bool) {
	o.MobileAppIdentifier = nil
	o.isExplicitNullMobileAppIdentifier = b
}
// GetVersion returns the Version field if non-nil, zero value otherwise.
func (o *ManagedMobileApp) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMobileApp) GetVersionOk() (string, bool) {
	if o == nil || o.Version == nil {
		var ret string
		return ret, false
	}
	return *o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ManagedMobileApp) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ManagedMobileApp) SetVersion(v string) {
	o.Version = &v
}

// SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Version value is set to nil even if false is passed
func (o *ManagedMobileApp) SetVersionExplicitNull(b bool) {
	o.Version = nil
	o.isExplicitNullVersion = b
}

// MarshalJSON returns the JSON representation of the model.
func (o ManagedMobileApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MobileAppIdentifier == nil {
		if o.isExplicitNullMobileAppIdentifier {
			toSerialize["mobileAppIdentifier"] = o.MobileAppIdentifier
		}
	} else {
		toSerialize["mobileAppIdentifier"] = o.MobileAppIdentifier
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

