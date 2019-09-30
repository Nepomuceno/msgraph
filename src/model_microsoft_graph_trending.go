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
// MicrosoftGraphTrending struct for MicrosoftGraphTrending
type MicrosoftGraphTrending struct {
	Id *string `json:"id,omitempty"`

	Weight *AnyOfnumberstringstring `json:"weight,omitempty"`

	ResourceVisualization *AnyOfmicrosoftGraphResourceVisualization `json:"resourceVisualization,omitempty"`
	isExplicitNullResourceVisualization bool `json:"-"`
	ResourceReference *AnyOfmicrosoftGraphResourceReference `json:"resourceReference,omitempty"`
	isExplicitNullResourceReference bool `json:"-"`
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	isExplicitNullLastModifiedDateTime bool `json:"-"`
	Resource *AnyOfmicrosoftGraphEntity `json:"resource,omitempty"`
	isExplicitNullResource bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphTrending) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTrending) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphTrending) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphTrending) SetId(v string) {
	o.Id = &v
}

// GetWeight returns the Weight field if non-nil, zero value otherwise.
func (o *MicrosoftGraphTrending) GetWeight() AnyOfnumberstringstring {
	if o == nil || o.Weight == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTrending) GetWeightOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Weight == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *MicrosoftGraphTrending) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given AnyOfnumberstringstring and assigns it to the Weight field.
func (o *MicrosoftGraphTrending) SetWeight(v AnyOfnumberstringstring) {
	o.Weight = &v
}

// GetResourceVisualization returns the ResourceVisualization field if non-nil, zero value otherwise.
func (o *MicrosoftGraphTrending) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization {
	if o == nil || o.ResourceVisualization == nil {
		var ret AnyOfmicrosoftGraphResourceVisualization
		return ret
	}
	return *o.ResourceVisualization
}

// GetResourceVisualizationOk returns a tuple with the ResourceVisualization field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTrending) GetResourceVisualizationOk() (AnyOfmicrosoftGraphResourceVisualization, bool) {
	if o == nil || o.ResourceVisualization == nil {
		var ret AnyOfmicrosoftGraphResourceVisualization
		return ret, false
	}
	return *o.ResourceVisualization, true
}

// HasResourceVisualization returns a boolean if a field has been set.
func (o *MicrosoftGraphTrending) HasResourceVisualization() bool {
	if o != nil && o.ResourceVisualization != nil {
		return true
	}

	return false
}

// SetResourceVisualization gets a reference to the given AnyOfmicrosoftGraphResourceVisualization and assigns it to the ResourceVisualization field.
func (o *MicrosoftGraphTrending) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization) {
	o.ResourceVisualization = &v
}

// SetResourceVisualizationExplicitNull (un)sets ResourceVisualization to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ResourceVisualization value is set to nil even if false is passed
func (o *MicrosoftGraphTrending) SetResourceVisualizationExplicitNull(b bool) {
	o.ResourceVisualization = nil
	o.isExplicitNullResourceVisualization = b
}
// GetResourceReference returns the ResourceReference field if non-nil, zero value otherwise.
func (o *MicrosoftGraphTrending) GetResourceReference() AnyOfmicrosoftGraphResourceReference {
	if o == nil || o.ResourceReference == nil {
		var ret AnyOfmicrosoftGraphResourceReference
		return ret
	}
	return *o.ResourceReference
}

// GetResourceReferenceOk returns a tuple with the ResourceReference field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTrending) GetResourceReferenceOk() (AnyOfmicrosoftGraphResourceReference, bool) {
	if o == nil || o.ResourceReference == nil {
		var ret AnyOfmicrosoftGraphResourceReference
		return ret, false
	}
	return *o.ResourceReference, true
}

// HasResourceReference returns a boolean if a field has been set.
func (o *MicrosoftGraphTrending) HasResourceReference() bool {
	if o != nil && o.ResourceReference != nil {
		return true
	}

	return false
}

// SetResourceReference gets a reference to the given AnyOfmicrosoftGraphResourceReference and assigns it to the ResourceReference field.
func (o *MicrosoftGraphTrending) SetResourceReference(v AnyOfmicrosoftGraphResourceReference) {
	o.ResourceReference = &v
}

// SetResourceReferenceExplicitNull (un)sets ResourceReference to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ResourceReference value is set to nil even if false is passed
func (o *MicrosoftGraphTrending) SetResourceReferenceExplicitNull(b bool) {
	o.ResourceReference = nil
	o.isExplicitNullResourceReference = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphTrending) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTrending) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphTrending) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphTrending) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphTrending) SetLastModifiedDateTimeExplicitNull(b bool) {
	o.LastModifiedDateTime = nil
	o.isExplicitNullLastModifiedDateTime = b
}
// GetResource returns the Resource field if non-nil, zero value otherwise.
func (o *MicrosoftGraphTrending) GetResource() AnyOfmicrosoftGraphEntity {
	if o == nil || o.Resource == nil {
		var ret AnyOfmicrosoftGraphEntity
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTrending) GetResourceOk() (AnyOfmicrosoftGraphEntity, bool) {
	if o == nil || o.Resource == nil {
		var ret AnyOfmicrosoftGraphEntity
		return ret, false
	}
	return *o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *MicrosoftGraphTrending) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given AnyOfmicrosoftGraphEntity and assigns it to the Resource field.
func (o *MicrosoftGraphTrending) SetResource(v AnyOfmicrosoftGraphEntity) {
	o.Resource = &v
}

// SetResourceExplicitNull (un)sets Resource to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Resource value is set to nil even if false is passed
func (o *MicrosoftGraphTrending) SetResourceExplicitNull(b bool) {
	o.Resource = nil
	o.isExplicitNullResource = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphTrending) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.ResourceVisualization == nil {
		if o.isExplicitNullResourceVisualization {
			toSerialize["resourceVisualization"] = o.ResourceVisualization
		}
	} else {
		toSerialize["resourceVisualization"] = o.ResourceVisualization
	}
	if o.ResourceReference == nil {
		if o.isExplicitNullResourceReference {
			toSerialize["resourceReference"] = o.ResourceReference
		}
	} else {
		toSerialize["resourceReference"] = o.ResourceReference
	}
	if o.LastModifiedDateTime == nil {
		if o.isExplicitNullLastModifiedDateTime {
			toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
		}
	} else {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Resource == nil {
		if o.isExplicitNullResource {
			toSerialize["resource"] = o.Resource
		}
	} else {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}


