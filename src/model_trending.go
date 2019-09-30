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
// Trending struct for Trending
type Trending struct {
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

// GetWeight returns the Weight field if non-nil, zero value otherwise.
func (o *Trending) GetWeight() AnyOfnumberstringstring {
	if o == nil || o.Weight == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Trending) GetWeightOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Weight == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *Trending) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given AnyOfnumberstringstring and assigns it to the Weight field.
func (o *Trending) SetWeight(v AnyOfnumberstringstring) {
	o.Weight = &v
}

// GetResourceVisualization returns the ResourceVisualization field if non-nil, zero value otherwise.
func (o *Trending) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization {
	if o == nil || o.ResourceVisualization == nil {
		var ret AnyOfmicrosoftGraphResourceVisualization
		return ret
	}
	return *o.ResourceVisualization
}

// GetResourceVisualizationOk returns a tuple with the ResourceVisualization field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Trending) GetResourceVisualizationOk() (AnyOfmicrosoftGraphResourceVisualization, bool) {
	if o == nil || o.ResourceVisualization == nil {
		var ret AnyOfmicrosoftGraphResourceVisualization
		return ret, false
	}
	return *o.ResourceVisualization, true
}

// HasResourceVisualization returns a boolean if a field has been set.
func (o *Trending) HasResourceVisualization() bool {
	if o != nil && o.ResourceVisualization != nil {
		return true
	}

	return false
}

// SetResourceVisualization gets a reference to the given AnyOfmicrosoftGraphResourceVisualization and assigns it to the ResourceVisualization field.
func (o *Trending) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization) {
	o.ResourceVisualization = &v
}

// SetResourceVisualizationExplicitNull (un)sets ResourceVisualization to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ResourceVisualization value is set to nil even if false is passed
func (o *Trending) SetResourceVisualizationExplicitNull(b bool) {
	o.ResourceVisualization = nil
	o.isExplicitNullResourceVisualization = b
}
// GetResourceReference returns the ResourceReference field if non-nil, zero value otherwise.
func (o *Trending) GetResourceReference() AnyOfmicrosoftGraphResourceReference {
	if o == nil || o.ResourceReference == nil {
		var ret AnyOfmicrosoftGraphResourceReference
		return ret
	}
	return *o.ResourceReference
}

// GetResourceReferenceOk returns a tuple with the ResourceReference field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Trending) GetResourceReferenceOk() (AnyOfmicrosoftGraphResourceReference, bool) {
	if o == nil || o.ResourceReference == nil {
		var ret AnyOfmicrosoftGraphResourceReference
		return ret, false
	}
	return *o.ResourceReference, true
}

// HasResourceReference returns a boolean if a field has been set.
func (o *Trending) HasResourceReference() bool {
	if o != nil && o.ResourceReference != nil {
		return true
	}

	return false
}

// SetResourceReference gets a reference to the given AnyOfmicrosoftGraphResourceReference and assigns it to the ResourceReference field.
func (o *Trending) SetResourceReference(v AnyOfmicrosoftGraphResourceReference) {
	o.ResourceReference = &v
}

// SetResourceReferenceExplicitNull (un)sets ResourceReference to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ResourceReference value is set to nil even if false is passed
func (o *Trending) SetResourceReferenceExplicitNull(b bool) {
	o.ResourceReference = nil
	o.isExplicitNullResourceReference = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *Trending) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Trending) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *Trending) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *Trending) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedDateTime value is set to nil even if false is passed
func (o *Trending) SetLastModifiedDateTimeExplicitNull(b bool) {
	o.LastModifiedDateTime = nil
	o.isExplicitNullLastModifiedDateTime = b
}
// GetResource returns the Resource field if non-nil, zero value otherwise.
func (o *Trending) GetResource() AnyOfmicrosoftGraphEntity {
	if o == nil || o.Resource == nil {
		var ret AnyOfmicrosoftGraphEntity
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Trending) GetResourceOk() (AnyOfmicrosoftGraphEntity, bool) {
	if o == nil || o.Resource == nil {
		var ret AnyOfmicrosoftGraphEntity
		return ret, false
	}
	return *o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *Trending) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given AnyOfmicrosoftGraphEntity and assigns it to the Resource field.
func (o *Trending) SetResource(v AnyOfmicrosoftGraphEntity) {
	o.Resource = &v
}

// SetResourceExplicitNull (un)sets Resource to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Resource value is set to nil even if false is passed
func (o *Trending) SetResourceExplicitNull(b bool) {
	o.Resource = nil
	o.isExplicitNullResource = b
}

// MarshalJSON returns the JSON representation of the model.
func (o Trending) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

