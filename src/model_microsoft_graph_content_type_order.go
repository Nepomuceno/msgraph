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
// MicrosoftGraphContentTypeOrder struct for MicrosoftGraphContentTypeOrder
type MicrosoftGraphContentTypeOrder struct {
	Default *bool `json:"default,omitempty"`
	isExplicitNullDefault bool `json:"-"`
	Position *int32 `json:"position,omitempty"`
	isExplicitNullPosition bool `json:"-"`
}

// GetDefault returns the Default field if non-nil, zero value otherwise.
func (o *MicrosoftGraphContentTypeOrder) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphContentTypeOrder) GetDefaultOk() (bool, bool) {
	if o == nil || o.Default == nil {
		var ret bool
		return ret, false
	}
	return *o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *MicrosoftGraphContentTypeOrder) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *MicrosoftGraphContentTypeOrder) SetDefault(v bool) {
	o.Default = &v
}

// SetDefaultExplicitNull (un)sets Default to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Default value is set to nil even if false is passed
func (o *MicrosoftGraphContentTypeOrder) SetDefaultExplicitNull(b bool) {
	o.Default = nil
	o.isExplicitNullDefault = b
}
// GetPosition returns the Position field if non-nil, zero value otherwise.
func (o *MicrosoftGraphContentTypeOrder) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphContentTypeOrder) GetPositionOk() (int32, bool) {
	if o == nil || o.Position == nil {
		var ret int32
		return ret, false
	}
	return *o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *MicrosoftGraphContentTypeOrder) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *MicrosoftGraphContentTypeOrder) SetPosition(v int32) {
	o.Position = &v
}

// SetPositionExplicitNull (un)sets Position to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Position value is set to nil even if false is passed
func (o *MicrosoftGraphContentTypeOrder) SetPositionExplicitNull(b bool) {
	o.Position = nil
	o.isExplicitNullPosition = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphContentTypeOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Default == nil {
		if o.isExplicitNullDefault {
			toSerialize["default"] = o.Default
		}
	} else {
		toSerialize["default"] = o.Default
	}
	if o.Position == nil {
		if o.isExplicitNullPosition {
			toSerialize["position"] = o.Position
		}
	} else {
		toSerialize["position"] = o.Position
	}
	return json.Marshal(toSerialize)
}

