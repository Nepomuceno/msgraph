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
	"time"
	"encoding/json"
)
// MicrosoftGraphSharingDetail struct for MicrosoftGraphSharingDetail
type MicrosoftGraphSharingDetail struct {
	SharedBy *AnyOfmicrosoftGraphInsightIdentity `json:"sharedBy,omitempty"`
	isExplicitNullSharedBy bool `json:"-"`
	SharedDateTime *time.Time `json:"sharedDateTime,omitempty"`
	isExplicitNullSharedDateTime bool `json:"-"`
	SharingSubject *string `json:"sharingSubject,omitempty"`
	isExplicitNullSharingSubject bool `json:"-"`
	SharingType *string `json:"sharingType,omitempty"`
	isExplicitNullSharingType bool `json:"-"`
	SharingReference *AnyOfmicrosoftGraphResourceReference `json:"sharingReference,omitempty"`
	isExplicitNullSharingReference bool `json:"-"`
}

// GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingDetail) GetSharedBy() AnyOfmicrosoftGraphInsightIdentity {
	if o == nil || o.SharedBy == nil {
		var ret AnyOfmicrosoftGraphInsightIdentity
		return ret
	}
	return *o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingDetail) GetSharedByOk() (AnyOfmicrosoftGraphInsightIdentity, bool) {
	if o == nil || o.SharedBy == nil {
		var ret AnyOfmicrosoftGraphInsightIdentity
		return ret, false
	}
	return *o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingDetail) HasSharedBy() bool {
	if o != nil && o.SharedBy != nil {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given AnyOfmicrosoftGraphInsightIdentity and assigns it to the SharedBy field.
func (o *MicrosoftGraphSharingDetail) SetSharedBy(v AnyOfmicrosoftGraphInsightIdentity) {
	o.SharedBy = &v
}

// SetSharedByExplicitNull (un)sets SharedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SharedBy value is set to nil even if false is passed
func (o *MicrosoftGraphSharingDetail) SetSharedByExplicitNull(b bool) {
	o.SharedBy = nil
	o.isExplicitNullSharedBy = b
}
// GetSharedDateTime returns the SharedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingDetail) GetSharedDateTime() time.Time {
	if o == nil || o.SharedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.SharedDateTime
}

// GetSharedDateTimeOk returns a tuple with the SharedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingDetail) GetSharedDateTimeOk() (time.Time, bool) {
	if o == nil || o.SharedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.SharedDateTime, true
}

// HasSharedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingDetail) HasSharedDateTime() bool {
	if o != nil && o.SharedDateTime != nil {
		return true
	}

	return false
}

// SetSharedDateTime gets a reference to the given time.Time and assigns it to the SharedDateTime field.
func (o *MicrosoftGraphSharingDetail) SetSharedDateTime(v time.Time) {
	o.SharedDateTime = &v
}

// SetSharedDateTimeExplicitNull (un)sets SharedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SharedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphSharingDetail) SetSharedDateTimeExplicitNull(b bool) {
	o.SharedDateTime = nil
	o.isExplicitNullSharedDateTime = b
}
// GetSharingSubject returns the SharingSubject field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingDetail) GetSharingSubject() string {
	if o == nil || o.SharingSubject == nil {
		var ret string
		return ret
	}
	return *o.SharingSubject
}

// GetSharingSubjectOk returns a tuple with the SharingSubject field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingDetail) GetSharingSubjectOk() (string, bool) {
	if o == nil || o.SharingSubject == nil {
		var ret string
		return ret, false
	}
	return *o.SharingSubject, true
}

// HasSharingSubject returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingDetail) HasSharingSubject() bool {
	if o != nil && o.SharingSubject != nil {
		return true
	}

	return false
}

// SetSharingSubject gets a reference to the given string and assigns it to the SharingSubject field.
func (o *MicrosoftGraphSharingDetail) SetSharingSubject(v string) {
	o.SharingSubject = &v
}

// SetSharingSubjectExplicitNull (un)sets SharingSubject to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SharingSubject value is set to nil even if false is passed
func (o *MicrosoftGraphSharingDetail) SetSharingSubjectExplicitNull(b bool) {
	o.SharingSubject = nil
	o.isExplicitNullSharingSubject = b
}
// GetSharingType returns the SharingType field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingDetail) GetSharingType() string {
	if o == nil || o.SharingType == nil {
		var ret string
		return ret
	}
	return *o.SharingType
}

// GetSharingTypeOk returns a tuple with the SharingType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingDetail) GetSharingTypeOk() (string, bool) {
	if o == nil || o.SharingType == nil {
		var ret string
		return ret, false
	}
	return *o.SharingType, true
}

// HasSharingType returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingDetail) HasSharingType() bool {
	if o != nil && o.SharingType != nil {
		return true
	}

	return false
}

// SetSharingType gets a reference to the given string and assigns it to the SharingType field.
func (o *MicrosoftGraphSharingDetail) SetSharingType(v string) {
	o.SharingType = &v
}

// SetSharingTypeExplicitNull (un)sets SharingType to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SharingType value is set to nil even if false is passed
func (o *MicrosoftGraphSharingDetail) SetSharingTypeExplicitNull(b bool) {
	o.SharingType = nil
	o.isExplicitNullSharingType = b
}
// GetSharingReference returns the SharingReference field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingDetail) GetSharingReference() AnyOfmicrosoftGraphResourceReference {
	if o == nil || o.SharingReference == nil {
		var ret AnyOfmicrosoftGraphResourceReference
		return ret
	}
	return *o.SharingReference
}

// GetSharingReferenceOk returns a tuple with the SharingReference field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingDetail) GetSharingReferenceOk() (AnyOfmicrosoftGraphResourceReference, bool) {
	if o == nil || o.SharingReference == nil {
		var ret AnyOfmicrosoftGraphResourceReference
		return ret, false
	}
	return *o.SharingReference, true
}

// HasSharingReference returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingDetail) HasSharingReference() bool {
	if o != nil && o.SharingReference != nil {
		return true
	}

	return false
}

// SetSharingReference gets a reference to the given AnyOfmicrosoftGraphResourceReference and assigns it to the SharingReference field.
func (o *MicrosoftGraphSharingDetail) SetSharingReference(v AnyOfmicrosoftGraphResourceReference) {
	o.SharingReference = &v
}

// SetSharingReferenceExplicitNull (un)sets SharingReference to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SharingReference value is set to nil even if false is passed
func (o *MicrosoftGraphSharingDetail) SetSharingReferenceExplicitNull(b bool) {
	o.SharingReference = nil
	o.isExplicitNullSharingReference = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphSharingDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SharedBy == nil {
		if o.isExplicitNullSharedBy {
			toSerialize["sharedBy"] = o.SharedBy
		}
	} else {
		toSerialize["sharedBy"] = o.SharedBy
	}
	if o.SharedDateTime == nil {
		if o.isExplicitNullSharedDateTime {
			toSerialize["sharedDateTime"] = o.SharedDateTime
		}
	} else {
		toSerialize["sharedDateTime"] = o.SharedDateTime
	}
	if o.SharingSubject == nil {
		if o.isExplicitNullSharingSubject {
			toSerialize["sharingSubject"] = o.SharingSubject
		}
	} else {
		toSerialize["sharingSubject"] = o.SharingSubject
	}
	if o.SharingType == nil {
		if o.isExplicitNullSharingType {
			toSerialize["sharingType"] = o.SharingType
		}
	} else {
		toSerialize["sharingType"] = o.SharingType
	}
	if o.SharingReference == nil {
		if o.isExplicitNullSharingReference {
			toSerialize["sharingReference"] = o.SharingReference
		}
	} else {
		toSerialize["sharingReference"] = o.SharingReference
	}
	return json.Marshal(toSerialize)
}

