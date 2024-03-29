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
// MicrosoftGraphSharingLink struct for MicrosoftGraphSharingLink
type MicrosoftGraphSharingLink struct {
	Application *AnyOfmicrosoftGraphIdentity `json:"application,omitempty"`
	isExplicitNullApplication bool `json:"-"`
	Scope *string `json:"scope,omitempty"`
	isExplicitNullScope bool `json:"-"`
	Type *string `json:"type,omitempty"`
	isExplicitNullType bool `json:"-"`
	WebUrl *string `json:"webUrl,omitempty"`
	isExplicitNullWebUrl bool `json:"-"`
}

// GetApplication returns the Application field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingLink) GetApplication() AnyOfmicrosoftGraphIdentity {
	if o == nil || o.Application == nil {
		var ret AnyOfmicrosoftGraphIdentity
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingLink) GetApplicationOk() (AnyOfmicrosoftGraphIdentity, bool) {
	if o == nil || o.Application == nil {
		var ret AnyOfmicrosoftGraphIdentity
		return ret, false
	}
	return *o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingLink) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given AnyOfmicrosoftGraphIdentity and assigns it to the Application field.
func (o *MicrosoftGraphSharingLink) SetApplication(v AnyOfmicrosoftGraphIdentity) {
	o.Application = &v
}

// SetApplicationExplicitNull (un)sets Application to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Application value is set to nil even if false is passed
func (o *MicrosoftGraphSharingLink) SetApplicationExplicitNull(b bool) {
	o.Application = nil
	o.isExplicitNullApplication = b
}
// GetScope returns the Scope field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingLink) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingLink) GetScopeOk() (string, bool) {
	if o == nil || o.Scope == nil {
		var ret string
		return ret, false
	}
	return *o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingLink) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *MicrosoftGraphSharingLink) SetScope(v string) {
	o.Scope = &v
}

// SetScopeExplicitNull (un)sets Scope to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Scope value is set to nil even if false is passed
func (o *MicrosoftGraphSharingLink) SetScopeExplicitNull(b bool) {
	o.Scope = nil
	o.isExplicitNullScope = b
}
// GetType returns the Type field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingLink) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingLink) GetTypeOk() (string, bool) {
	if o == nil || o.Type == nil {
		var ret string
		return ret, false
	}
	return *o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingLink) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MicrosoftGraphSharingLink) SetType(v string) {
	o.Type = &v
}

// SetTypeExplicitNull (un)sets Type to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Type value is set to nil even if false is passed
func (o *MicrosoftGraphSharingLink) SetTypeExplicitNull(b bool) {
	o.Type = nil
	o.isExplicitNullType = b
}
// GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingLink) GetWebUrl() string {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingLink) GetWebUrlOk() (string, bool) {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret, false
	}
	return *o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingLink) HasWebUrl() bool {
	if o != nil && o.WebUrl != nil {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *MicrosoftGraphSharingLink) SetWebUrl(v string) {
	o.WebUrl = &v
}

// SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The WebUrl value is set to nil even if false is passed
func (o *MicrosoftGraphSharingLink) SetWebUrlExplicitNull(b bool) {
	o.WebUrl = nil
	o.isExplicitNullWebUrl = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphSharingLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Application == nil {
		if o.isExplicitNullApplication {
			toSerialize["application"] = o.Application
		}
	} else {
		toSerialize["application"] = o.Application
	}
	if o.Scope == nil {
		if o.isExplicitNullScope {
			toSerialize["scope"] = o.Scope
		}
	} else {
		toSerialize["scope"] = o.Scope
	}
	if o.Type == nil {
		if o.isExplicitNullType {
			toSerialize["type"] = o.Type
		}
	} else {
		toSerialize["type"] = o.Type
	}
	if o.WebUrl == nil {
		if o.isExplicitNullWebUrl {
			toSerialize["webUrl"] = o.WebUrl
		}
	} else {
		toSerialize["webUrl"] = o.WebUrl
	}
	return json.Marshal(toSerialize)
}


