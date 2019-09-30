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
	"errors"
)
// OdataErrorDetail struct for OdataErrorDetail
type OdataErrorDetail struct {
	Code *string `json:"code,omitempty"`

	Message *string `json:"message,omitempty"`

	Target *string `json:"target,omitempty"`

}

// GetCode returns the Code field if non-nil, zero value otherwise.
func (o *OdataErrorDetail) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OdataErrorDetail) GetCodeOk() (string, bool) {
	if o == nil || o.Code == nil {
		var ret string
		return ret, false
	}
	return *o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OdataErrorDetail) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OdataErrorDetail) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field if non-nil, zero value otherwise.
func (o *OdataErrorDetail) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OdataErrorDetail) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OdataErrorDetail) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OdataErrorDetail) SetMessage(v string) {
	o.Message = &v
}

// GetTarget returns the Target field if non-nil, zero value otherwise.
func (o *OdataErrorDetail) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OdataErrorDetail) GetTargetOk() (string, bool) {
	if o == nil || o.Target == nil {
		var ret string
		return ret, false
	}
	return *o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *OdataErrorDetail) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *OdataErrorDetail) SetTarget(v string) {
	o.Target = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o OdataErrorDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code == nil {
		return nil, errors.New("Code is required and not nullable, but was not set on OdataErrorDetail")
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message == nil {
		return nil, errors.New("Message is required and not nullable, but was not set on OdataErrorDetail")
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}


