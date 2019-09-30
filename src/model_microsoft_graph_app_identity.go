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
// MicrosoftGraphAppIdentity struct for MicrosoftGraphAppIdentity
type MicrosoftGraphAppIdentity struct {
	AppId *string `json:"appId,omitempty"`
	isExplicitNullAppId bool `json:"-"`
	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	ServicePrincipalId *string `json:"servicePrincipalId,omitempty"`
	isExplicitNullServicePrincipalId bool `json:"-"`
	ServicePrincipalName *string `json:"servicePrincipalName,omitempty"`
	isExplicitNullServicePrincipalName bool `json:"-"`
}

// GetAppId returns the AppId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAppIdentity) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppIdentity) GetAppIdOk() (string, bool) {
	if o == nil || o.AppId == nil {
		var ret string
		return ret, false
	}
	return *o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *MicrosoftGraphAppIdentity) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *MicrosoftGraphAppIdentity) SetAppId(v string) {
	o.AppId = &v
}

// SetAppIdExplicitNull (un)sets AppId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AppId value is set to nil even if false is passed
func (o *MicrosoftGraphAppIdentity) SetAppIdExplicitNull(b bool) {
	o.AppId = nil
	o.isExplicitNullAppId = b
}
// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAppIdentity) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppIdentity) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphAppIdentity) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphAppIdentity) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphAppIdentity) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetServicePrincipalId returns the ServicePrincipalId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAppIdentity) GetServicePrincipalId() string {
	if o == nil || o.ServicePrincipalId == nil {
		var ret string
		return ret
	}
	return *o.ServicePrincipalId
}

// GetServicePrincipalIdOk returns a tuple with the ServicePrincipalId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppIdentity) GetServicePrincipalIdOk() (string, bool) {
	if o == nil || o.ServicePrincipalId == nil {
		var ret string
		return ret, false
	}
	return *o.ServicePrincipalId, true
}

// HasServicePrincipalId returns a boolean if a field has been set.
func (o *MicrosoftGraphAppIdentity) HasServicePrincipalId() bool {
	if o != nil && o.ServicePrincipalId != nil {
		return true
	}

	return false
}

// SetServicePrincipalId gets a reference to the given string and assigns it to the ServicePrincipalId field.
func (o *MicrosoftGraphAppIdentity) SetServicePrincipalId(v string) {
	o.ServicePrincipalId = &v
}

// SetServicePrincipalIdExplicitNull (un)sets ServicePrincipalId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ServicePrincipalId value is set to nil even if false is passed
func (o *MicrosoftGraphAppIdentity) SetServicePrincipalIdExplicitNull(b bool) {
	o.ServicePrincipalId = nil
	o.isExplicitNullServicePrincipalId = b
}
// GetServicePrincipalName returns the ServicePrincipalName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAppIdentity) GetServicePrincipalName() string {
	if o == nil || o.ServicePrincipalName == nil {
		var ret string
		return ret
	}
	return *o.ServicePrincipalName
}

// GetServicePrincipalNameOk returns a tuple with the ServicePrincipalName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppIdentity) GetServicePrincipalNameOk() (string, bool) {
	if o == nil || o.ServicePrincipalName == nil {
		var ret string
		return ret, false
	}
	return *o.ServicePrincipalName, true
}

// HasServicePrincipalName returns a boolean if a field has been set.
func (o *MicrosoftGraphAppIdentity) HasServicePrincipalName() bool {
	if o != nil && o.ServicePrincipalName != nil {
		return true
	}

	return false
}

// SetServicePrincipalName gets a reference to the given string and assigns it to the ServicePrincipalName field.
func (o *MicrosoftGraphAppIdentity) SetServicePrincipalName(v string) {
	o.ServicePrincipalName = &v
}

// SetServicePrincipalNameExplicitNull (un)sets ServicePrincipalName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ServicePrincipalName value is set to nil even if false is passed
func (o *MicrosoftGraphAppIdentity) SetServicePrincipalNameExplicitNull(b bool) {
	o.ServicePrincipalName = nil
	o.isExplicitNullServicePrincipalName = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphAppIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId == nil {
		if o.isExplicitNullAppId {
			toSerialize["appId"] = o.AppId
		}
	} else {
		toSerialize["appId"] = o.AppId
	}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ServicePrincipalId == nil {
		if o.isExplicitNullServicePrincipalId {
			toSerialize["servicePrincipalId"] = o.ServicePrincipalId
		}
	} else {
		toSerialize["servicePrincipalId"] = o.ServicePrincipalId
	}
	if o.ServicePrincipalName == nil {
		if o.isExplicitNullServicePrincipalName {
			toSerialize["servicePrincipalName"] = o.ServicePrincipalName
		}
	} else {
		toSerialize["servicePrincipalName"] = o.ServicePrincipalName
	}
	return json.Marshal(toSerialize)
}


