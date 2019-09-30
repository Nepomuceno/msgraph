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
// MicrosoftGraphProfilePhoto struct for MicrosoftGraphProfilePhoto
type MicrosoftGraphProfilePhoto struct {
	Id *string `json:"id,omitempty"`

	Height *int32 `json:"height,omitempty"`
	isExplicitNullHeight bool `json:"-"`
	Width *int32 `json:"width,omitempty"`
	isExplicitNullWidth bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphProfilePhoto) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphProfilePhoto) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphProfilePhoto) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphProfilePhoto) SetId(v string) {
	o.Id = &v
}

// GetHeight returns the Height field if non-nil, zero value otherwise.
func (o *MicrosoftGraphProfilePhoto) GetHeight() int32 {
	if o == nil || o.Height == nil {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphProfilePhoto) GetHeightOk() (int32, bool) {
	if o == nil || o.Height == nil {
		var ret int32
		return ret, false
	}
	return *o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *MicrosoftGraphProfilePhoto) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *MicrosoftGraphProfilePhoto) SetHeight(v int32) {
	o.Height = &v
}

// SetHeightExplicitNull (un)sets Height to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Height value is set to nil even if false is passed
func (o *MicrosoftGraphProfilePhoto) SetHeightExplicitNull(b bool) {
	o.Height = nil
	o.isExplicitNullHeight = b
}
// GetWidth returns the Width field if non-nil, zero value otherwise.
func (o *MicrosoftGraphProfilePhoto) GetWidth() int32 {
	if o == nil || o.Width == nil {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphProfilePhoto) GetWidthOk() (int32, bool) {
	if o == nil || o.Width == nil {
		var ret int32
		return ret, false
	}
	return *o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *MicrosoftGraphProfilePhoto) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *MicrosoftGraphProfilePhoto) SetWidth(v int32) {
	o.Width = &v
}

// SetWidthExplicitNull (un)sets Width to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Width value is set to nil even if false is passed
func (o *MicrosoftGraphProfilePhoto) SetWidthExplicitNull(b bool) {
	o.Width = nil
	o.isExplicitNullWidth = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphProfilePhoto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Height == nil {
		if o.isExplicitNullHeight {
			toSerialize["height"] = o.Height
		}
	} else {
		toSerialize["height"] = o.Height
	}
	if o.Width == nil {
		if o.isExplicitNullWidth {
			toSerialize["width"] = o.Width
		}
	} else {
		toSerialize["width"] = o.Width
	}
	return json.Marshal(toSerialize)
}


