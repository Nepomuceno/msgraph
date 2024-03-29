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
// MicrosoftGraphDomainDnsRecord struct for MicrosoftGraphDomainDnsRecord
type MicrosoftGraphDomainDnsRecord struct {
	Id *string `json:"id,omitempty"`

	IsOptional *bool `json:"isOptional,omitempty"`

	Label *string `json:"label,omitempty"`

	RecordType *string `json:"recordType,omitempty"`
	isExplicitNullRecordType bool `json:"-"`
	SupportedService *string `json:"supportedService,omitempty"`

	Ttl *int32 `json:"ttl,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDomainDnsRecord) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomainDnsRecord) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDomainDnsRecord) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDomainDnsRecord) SetId(v string) {
	o.Id = &v
}

// GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDomainDnsRecord) GetIsOptional() bool {
	if o == nil || o.IsOptional == nil {
		var ret bool
		return ret
	}
	return *o.IsOptional
}

// GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomainDnsRecord) GetIsOptionalOk() (bool, bool) {
	if o == nil || o.IsOptional == nil {
		var ret bool
		return ret, false
	}
	return *o.IsOptional, true
}

// HasIsOptional returns a boolean if a field has been set.
func (o *MicrosoftGraphDomainDnsRecord) HasIsOptional() bool {
	if o != nil && o.IsOptional != nil {
		return true
	}

	return false
}

// SetIsOptional gets a reference to the given bool and assigns it to the IsOptional field.
func (o *MicrosoftGraphDomainDnsRecord) SetIsOptional(v bool) {
	o.IsOptional = &v
}

// GetLabel returns the Label field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDomainDnsRecord) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomainDnsRecord) GetLabelOk() (string, bool) {
	if o == nil || o.Label == nil {
		var ret string
		return ret, false
	}
	return *o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *MicrosoftGraphDomainDnsRecord) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *MicrosoftGraphDomainDnsRecord) SetLabel(v string) {
	o.Label = &v
}

// GetRecordType returns the RecordType field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDomainDnsRecord) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomainDnsRecord) GetRecordTypeOk() (string, bool) {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret, false
	}
	return *o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *MicrosoftGraphDomainDnsRecord) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *MicrosoftGraphDomainDnsRecord) SetRecordType(v string) {
	o.RecordType = &v
}

// SetRecordTypeExplicitNull (un)sets RecordType to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The RecordType value is set to nil even if false is passed
func (o *MicrosoftGraphDomainDnsRecord) SetRecordTypeExplicitNull(b bool) {
	o.RecordType = nil
	o.isExplicitNullRecordType = b
}
// GetSupportedService returns the SupportedService field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDomainDnsRecord) GetSupportedService() string {
	if o == nil || o.SupportedService == nil {
		var ret string
		return ret
	}
	return *o.SupportedService
}

// GetSupportedServiceOk returns a tuple with the SupportedService field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomainDnsRecord) GetSupportedServiceOk() (string, bool) {
	if o == nil || o.SupportedService == nil {
		var ret string
		return ret, false
	}
	return *o.SupportedService, true
}

// HasSupportedService returns a boolean if a field has been set.
func (o *MicrosoftGraphDomainDnsRecord) HasSupportedService() bool {
	if o != nil && o.SupportedService != nil {
		return true
	}

	return false
}

// SetSupportedService gets a reference to the given string and assigns it to the SupportedService field.
func (o *MicrosoftGraphDomainDnsRecord) SetSupportedService(v string) {
	o.SupportedService = &v
}

// GetTtl returns the Ttl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDomainDnsRecord) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomainDnsRecord) GetTtlOk() (int32, bool) {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret, false
	}
	return *o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *MicrosoftGraphDomainDnsRecord) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *MicrosoftGraphDomainDnsRecord) SetTtl(v int32) {
	o.Ttl = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphDomainDnsRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsOptional != nil {
		toSerialize["isOptional"] = o.IsOptional
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.RecordType == nil {
		if o.isExplicitNullRecordType {
			toSerialize["recordType"] = o.RecordType
		}
	} else {
		toSerialize["recordType"] = o.RecordType
	}
	if o.SupportedService != nil {
		toSerialize["supportedService"] = o.SupportedService
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}


