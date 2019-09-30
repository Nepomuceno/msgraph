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
// MicrosoftGraphOnenoteOperationError struct for MicrosoftGraphOnenoteOperationError
type MicrosoftGraphOnenoteOperationError struct {
	Code *string `json:"code,omitempty"`
	isExplicitNullCode bool `json:"-"`
	Message *string `json:"message,omitempty"`
	isExplicitNullMessage bool `json:"-"`
}

// GetCode returns the Code field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteOperationError) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteOperationError) GetCodeOk() (string, bool) {
	if o == nil || o.Code == nil {
		var ret string
		return ret, false
	}
	return *o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperationError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *MicrosoftGraphOnenoteOperationError) SetCode(v string) {
	o.Code = &v
}

// SetCodeExplicitNull (un)sets Code to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Code value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteOperationError) SetCodeExplicitNull(b bool) {
	o.Code = nil
	o.isExplicitNullCode = b
}
// GetMessage returns the Message field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteOperationError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteOperationError) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperationError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MicrosoftGraphOnenoteOperationError) SetMessage(v string) {
	o.Message = &v
}

// SetMessageExplicitNull (un)sets Message to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Message value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteOperationError) SetMessageExplicitNull(b bool) {
	o.Message = nil
	o.isExplicitNullMessage = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphOnenoteOperationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code == nil {
		if o.isExplicitNullCode {
			toSerialize["code"] = o.Code
		}
	} else {
		toSerialize["code"] = o.Code
	}
	if o.Message == nil {
		if o.isExplicitNullMessage {
			toSerialize["message"] = o.Message
		}
	} else {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}


