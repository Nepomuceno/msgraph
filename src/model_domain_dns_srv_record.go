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
// DomainDnsSrvRecord struct for DomainDnsSrvRecord
type DomainDnsSrvRecord struct {
	NameTarget *string `json:"nameTarget,omitempty"`
	isExplicitNullNameTarget bool `json:"-"`
	Port *int32 `json:"port,omitempty"`
	isExplicitNullPort bool `json:"-"`
	Priority *int32 `json:"priority,omitempty"`
	isExplicitNullPriority bool `json:"-"`
	Protocol *string `json:"protocol,omitempty"`
	isExplicitNullProtocol bool `json:"-"`
	Service *string `json:"service,omitempty"`
	isExplicitNullService bool `json:"-"`
	Weight *int32 `json:"weight,omitempty"`
	isExplicitNullWeight bool `json:"-"`
}

// GetNameTarget returns the NameTarget field if non-nil, zero value otherwise.
func (o *DomainDnsSrvRecord) GetNameTarget() string {
	if o == nil || o.NameTarget == nil {
		var ret string
		return ret
	}
	return *o.NameTarget
}

// GetNameTargetOk returns a tuple with the NameTarget field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DomainDnsSrvRecord) GetNameTargetOk() (string, bool) {
	if o == nil || o.NameTarget == nil {
		var ret string
		return ret, false
	}
	return *o.NameTarget, true
}

// HasNameTarget returns a boolean if a field has been set.
func (o *DomainDnsSrvRecord) HasNameTarget() bool {
	if o != nil && o.NameTarget != nil {
		return true
	}

	return false
}

// SetNameTarget gets a reference to the given string and assigns it to the NameTarget field.
func (o *DomainDnsSrvRecord) SetNameTarget(v string) {
	o.NameTarget = &v
}

// SetNameTargetExplicitNull (un)sets NameTarget to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The NameTarget value is set to nil even if false is passed
func (o *DomainDnsSrvRecord) SetNameTargetExplicitNull(b bool) {
	o.NameTarget = nil
	o.isExplicitNullNameTarget = b
}
// GetPort returns the Port field if non-nil, zero value otherwise.
func (o *DomainDnsSrvRecord) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DomainDnsSrvRecord) GetPortOk() (int32, bool) {
	if o == nil || o.Port == nil {
		var ret int32
		return ret, false
	}
	return *o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DomainDnsSrvRecord) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *DomainDnsSrvRecord) SetPort(v int32) {
	o.Port = &v
}

// SetPortExplicitNull (un)sets Port to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Port value is set to nil even if false is passed
func (o *DomainDnsSrvRecord) SetPortExplicitNull(b bool) {
	o.Port = nil
	o.isExplicitNullPort = b
}
// GetPriority returns the Priority field if non-nil, zero value otherwise.
func (o *DomainDnsSrvRecord) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DomainDnsSrvRecord) GetPriorityOk() (int32, bool) {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret, false
	}
	return *o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *DomainDnsSrvRecord) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *DomainDnsSrvRecord) SetPriority(v int32) {
	o.Priority = &v
}

// SetPriorityExplicitNull (un)sets Priority to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Priority value is set to nil even if false is passed
func (o *DomainDnsSrvRecord) SetPriorityExplicitNull(b bool) {
	o.Priority = nil
	o.isExplicitNullPriority = b
}
// GetProtocol returns the Protocol field if non-nil, zero value otherwise.
func (o *DomainDnsSrvRecord) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DomainDnsSrvRecord) GetProtocolOk() (string, bool) {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret, false
	}
	return *o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *DomainDnsSrvRecord) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *DomainDnsSrvRecord) SetProtocol(v string) {
	o.Protocol = &v
}

// SetProtocolExplicitNull (un)sets Protocol to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Protocol value is set to nil even if false is passed
func (o *DomainDnsSrvRecord) SetProtocolExplicitNull(b bool) {
	o.Protocol = nil
	o.isExplicitNullProtocol = b
}
// GetService returns the Service field if non-nil, zero value otherwise.
func (o *DomainDnsSrvRecord) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DomainDnsSrvRecord) GetServiceOk() (string, bool) {
	if o == nil || o.Service == nil {
		var ret string
		return ret, false
	}
	return *o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *DomainDnsSrvRecord) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *DomainDnsSrvRecord) SetService(v string) {
	o.Service = &v
}

// SetServiceExplicitNull (un)sets Service to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Service value is set to nil even if false is passed
func (o *DomainDnsSrvRecord) SetServiceExplicitNull(b bool) {
	o.Service = nil
	o.isExplicitNullService = b
}
// GetWeight returns the Weight field if non-nil, zero value otherwise.
func (o *DomainDnsSrvRecord) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DomainDnsSrvRecord) GetWeightOk() (int32, bool) {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret, false
	}
	return *o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *DomainDnsSrvRecord) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *DomainDnsSrvRecord) SetWeight(v int32) {
	o.Weight = &v
}

// SetWeightExplicitNull (un)sets Weight to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Weight value is set to nil even if false is passed
func (o *DomainDnsSrvRecord) SetWeightExplicitNull(b bool) {
	o.Weight = nil
	o.isExplicitNullWeight = b
}

// MarshalJSON returns the JSON representation of the model.
func (o DomainDnsSrvRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NameTarget == nil {
		if o.isExplicitNullNameTarget {
			toSerialize["nameTarget"] = o.NameTarget
		}
	} else {
		toSerialize["nameTarget"] = o.NameTarget
	}
	if o.Port == nil {
		if o.isExplicitNullPort {
			toSerialize["port"] = o.Port
		}
	} else {
		toSerialize["port"] = o.Port
	}
	if o.Priority == nil {
		if o.isExplicitNullPriority {
			toSerialize["priority"] = o.Priority
		}
	} else {
		toSerialize["priority"] = o.Priority
	}
	if o.Protocol == nil {
		if o.isExplicitNullProtocol {
			toSerialize["protocol"] = o.Protocol
		}
	} else {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Service == nil {
		if o.isExplicitNullService {
			toSerialize["service"] = o.Service
		}
	} else {
		toSerialize["service"] = o.Service
	}
	if o.Weight == nil {
		if o.isExplicitNullWeight {
			toSerialize["weight"] = o.Weight
		}
	} else {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

