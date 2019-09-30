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
// MicrosoftGraphDiagnostic struct for MicrosoftGraphDiagnostic
type MicrosoftGraphDiagnostic struct {
	Message *string `json:"message,omitempty"`
	isExplicitNullMessage bool `json:"-"`
	Url *string `json:"url,omitempty"`
	isExplicitNullUrl bool `json:"-"`
}

// GetMessage returns the Message field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDiagnostic) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDiagnostic) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MicrosoftGraphDiagnostic) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MicrosoftGraphDiagnostic) SetMessage(v string) {
	o.Message = &v
}

// SetMessageExplicitNull (un)sets Message to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Message value is set to nil even if false is passed
func (o *MicrosoftGraphDiagnostic) SetMessageExplicitNull(b bool) {
	o.Message = nil
	o.isExplicitNullMessage = b
}
// GetUrl returns the Url field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDiagnostic) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDiagnostic) GetUrlOk() (string, bool) {
	if o == nil || o.Url == nil {
		var ret string
		return ret, false
	}
	return *o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphDiagnostic) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MicrosoftGraphDiagnostic) SetUrl(v string) {
	o.Url = &v
}

// SetUrlExplicitNull (un)sets Url to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Url value is set to nil even if false is passed
func (o *MicrosoftGraphDiagnostic) SetUrlExplicitNull(b bool) {
	o.Url = nil
	o.isExplicitNullUrl = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphDiagnostic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message == nil {
		if o.isExplicitNullMessage {
			toSerialize["message"] = o.Message
		}
	} else {
		toSerialize["message"] = o.Message
	}
	if o.Url == nil {
		if o.isExplicitNullUrl {
			toSerialize["url"] = o.Url
		}
	} else {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

