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
// InlineObject883 struct for InlineObject883
type InlineObject883 struct {
	Principal *AnyOfobject `json:"principal,omitempty"`
	isExplicitNullPrincipal bool `json:"-"`
	Schedule *AnyOfobject `json:"schedule,omitempty"`
	isExplicitNullSchedule bool `json:"-"`
}

// GetPrincipal returns the Principal field if non-nil, zero value otherwise.
func (o *InlineObject883) GetPrincipal() AnyOfobject {
	if o == nil || o.Principal == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject883) GetPrincipalOk() (AnyOfobject, bool) {
	if o == nil || o.Principal == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *InlineObject883) HasPrincipal() bool {
	if o != nil && o.Principal != nil {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given AnyOfobject and assigns it to the Principal field.
func (o *InlineObject883) SetPrincipal(v AnyOfobject) {
	o.Principal = &v
}

// SetPrincipalExplicitNull (un)sets Principal to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Principal value is set to nil even if false is passed
func (o *InlineObject883) SetPrincipalExplicitNull(b bool) {
	o.Principal = nil
	o.isExplicitNullPrincipal = b
}
// GetSchedule returns the Schedule field if non-nil, zero value otherwise.
func (o *InlineObject883) GetSchedule() AnyOfobject {
	if o == nil || o.Schedule == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject883) GetScheduleOk() (AnyOfobject, bool) {
	if o == nil || o.Schedule == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *InlineObject883) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given AnyOfobject and assigns it to the Schedule field.
func (o *InlineObject883) SetSchedule(v AnyOfobject) {
	o.Schedule = &v
}

// SetScheduleExplicitNull (un)sets Schedule to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Schedule value is set to nil even if false is passed
func (o *InlineObject883) SetScheduleExplicitNull(b bool) {
	o.Schedule = nil
	o.isExplicitNullSchedule = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject883) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Principal == nil {
		if o.isExplicitNullPrincipal {
			toSerialize["principal"] = o.Principal
		}
	} else {
		toSerialize["principal"] = o.Principal
	}
	if o.Schedule == nil {
		if o.isExplicitNullSchedule {
			toSerialize["schedule"] = o.Schedule
		}
	} else {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}


