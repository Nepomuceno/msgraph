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
// MicrosoftGraphOnenoteEntityBaseModel struct for MicrosoftGraphOnenoteEntityBaseModel
type MicrosoftGraphOnenoteEntityBaseModel struct {
	Id *string `json:"id,omitempty"`

	Self *string `json:"self,omitempty"`
	isExplicitNullSelf bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteEntityBaseModel) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteEntityBaseModel) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteEntityBaseModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOnenoteEntityBaseModel) SetId(v string) {
	o.Id = &v
}

// GetSelf returns the Self field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteEntityBaseModel) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteEntityBaseModel) GetSelfOk() (string, bool) {
	if o == nil || o.Self == nil {
		var ret string
		return ret, false
	}
	return *o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteEntityBaseModel) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *MicrosoftGraphOnenoteEntityBaseModel) SetSelf(v string) {
	o.Self = &v
}

// SetSelfExplicitNull (un)sets Self to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Self value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteEntityBaseModel) SetSelfExplicitNull(b bool) {
	o.Self = nil
	o.isExplicitNullSelf = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphOnenoteEntityBaseModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Self == nil {
		if o.isExplicitNullSelf {
			toSerialize["self"] = o.Self
		}
	} else {
		toSerialize["self"] = o.Self
	}
	return json.Marshal(toSerialize)
}


