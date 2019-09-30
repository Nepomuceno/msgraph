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
// UsedInsight struct for UsedInsight
type UsedInsight struct {
	LastUsed *AnyOfmicrosoftGraphUsageDetails `json:"lastUsed,omitempty"`
	isExplicitNullLastUsed bool `json:"-"`
	ResourceVisualization *AnyOfmicrosoftGraphResourceVisualization `json:"resourceVisualization,omitempty"`
	isExplicitNullResourceVisualization bool `json:"-"`
	ResourceReference *AnyOfmicrosoftGraphResourceReference `json:"resourceReference,omitempty"`
	isExplicitNullResourceReference bool `json:"-"`
	Resource *AnyOfmicrosoftGraphEntity `json:"resource,omitempty"`
	isExplicitNullResource bool `json:"-"`
}

// GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.
func (o *UsedInsight) GetLastUsed() AnyOfmicrosoftGraphUsageDetails {
	if o == nil || o.LastUsed == nil {
		var ret AnyOfmicrosoftGraphUsageDetails
		return ret
	}
	return *o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UsedInsight) GetLastUsedOk() (AnyOfmicrosoftGraphUsageDetails, bool) {
	if o == nil || o.LastUsed == nil {
		var ret AnyOfmicrosoftGraphUsageDetails
		return ret, false
	}
	return *o.LastUsed, true
}

// HasLastUsed returns a boolean if a field has been set.
func (o *UsedInsight) HasLastUsed() bool {
	if o != nil && o.LastUsed != nil {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given AnyOfmicrosoftGraphUsageDetails and assigns it to the LastUsed field.
func (o *UsedInsight) SetLastUsed(v AnyOfmicrosoftGraphUsageDetails) {
	o.LastUsed = &v
}

// SetLastUsedExplicitNull (un)sets LastUsed to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastUsed value is set to nil even if false is passed
func (o *UsedInsight) SetLastUsedExplicitNull(b bool) {
	o.LastUsed = nil
	o.isExplicitNullLastUsed = b
}
// GetResourceVisualization returns the ResourceVisualization field if non-nil, zero value otherwise.
func (o *UsedInsight) GetResourceVisualization() AnyOfmicrosoftGraphResourceVisualization {
	if o == nil || o.ResourceVisualization == nil {
		var ret AnyOfmicrosoftGraphResourceVisualization
		return ret
	}
	return *o.ResourceVisualization
}

// GetResourceVisualizationOk returns a tuple with the ResourceVisualization field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UsedInsight) GetResourceVisualizationOk() (AnyOfmicrosoftGraphResourceVisualization, bool) {
	if o == nil || o.ResourceVisualization == nil {
		var ret AnyOfmicrosoftGraphResourceVisualization
		return ret, false
	}
	return *o.ResourceVisualization, true
}

// HasResourceVisualization returns a boolean if a field has been set.
func (o *UsedInsight) HasResourceVisualization() bool {
	if o != nil && o.ResourceVisualization != nil {
		return true
	}

	return false
}

// SetResourceVisualization gets a reference to the given AnyOfmicrosoftGraphResourceVisualization and assigns it to the ResourceVisualization field.
func (o *UsedInsight) SetResourceVisualization(v AnyOfmicrosoftGraphResourceVisualization) {
	o.ResourceVisualization = &v
}

// SetResourceVisualizationExplicitNull (un)sets ResourceVisualization to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ResourceVisualization value is set to nil even if false is passed
func (o *UsedInsight) SetResourceVisualizationExplicitNull(b bool) {
	o.ResourceVisualization = nil
	o.isExplicitNullResourceVisualization = b
}
// GetResourceReference returns the ResourceReference field if non-nil, zero value otherwise.
func (o *UsedInsight) GetResourceReference() AnyOfmicrosoftGraphResourceReference {
	if o == nil || o.ResourceReference == nil {
		var ret AnyOfmicrosoftGraphResourceReference
		return ret
	}
	return *o.ResourceReference
}

// GetResourceReferenceOk returns a tuple with the ResourceReference field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UsedInsight) GetResourceReferenceOk() (AnyOfmicrosoftGraphResourceReference, bool) {
	if o == nil || o.ResourceReference == nil {
		var ret AnyOfmicrosoftGraphResourceReference
		return ret, false
	}
	return *o.ResourceReference, true
}

// HasResourceReference returns a boolean if a field has been set.
func (o *UsedInsight) HasResourceReference() bool {
	if o != nil && o.ResourceReference != nil {
		return true
	}

	return false
}

// SetResourceReference gets a reference to the given AnyOfmicrosoftGraphResourceReference and assigns it to the ResourceReference field.
func (o *UsedInsight) SetResourceReference(v AnyOfmicrosoftGraphResourceReference) {
	o.ResourceReference = &v
}

// SetResourceReferenceExplicitNull (un)sets ResourceReference to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ResourceReference value is set to nil even if false is passed
func (o *UsedInsight) SetResourceReferenceExplicitNull(b bool) {
	o.ResourceReference = nil
	o.isExplicitNullResourceReference = b
}
// GetResource returns the Resource field if non-nil, zero value otherwise.
func (o *UsedInsight) GetResource() AnyOfmicrosoftGraphEntity {
	if o == nil || o.Resource == nil {
		var ret AnyOfmicrosoftGraphEntity
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UsedInsight) GetResourceOk() (AnyOfmicrosoftGraphEntity, bool) {
	if o == nil || o.Resource == nil {
		var ret AnyOfmicrosoftGraphEntity
		return ret, false
	}
	return *o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *UsedInsight) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given AnyOfmicrosoftGraphEntity and assigns it to the Resource field.
func (o *UsedInsight) SetResource(v AnyOfmicrosoftGraphEntity) {
	o.Resource = &v
}

// SetResourceExplicitNull (un)sets Resource to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Resource value is set to nil even if false is passed
func (o *UsedInsight) SetResourceExplicitNull(b bool) {
	o.Resource = nil
	o.isExplicitNullResource = b
}

// MarshalJSON returns the JSON representation of the model.
func (o UsedInsight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastUsed == nil {
		if o.isExplicitNullLastUsed {
			toSerialize["lastUsed"] = o.LastUsed
		}
	} else {
		toSerialize["lastUsed"] = o.LastUsed
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
	if o.Resource == nil {
		if o.isExplicitNullResource {
			toSerialize["resource"] = o.Resource
		}
	} else {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}


