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
// ThumbnailSet struct for ThumbnailSet
type ThumbnailSet struct {
	Large *AnyOfmicrosoftGraphThumbnail `json:"large,omitempty"`
	isExplicitNullLarge bool `json:"-"`
	Medium *AnyOfmicrosoftGraphThumbnail `json:"medium,omitempty"`
	isExplicitNullMedium bool `json:"-"`
	Small *AnyOfmicrosoftGraphThumbnail `json:"small,omitempty"`
	isExplicitNullSmall bool `json:"-"`
	Source *AnyOfmicrosoftGraphThumbnail `json:"source,omitempty"`
	isExplicitNullSource bool `json:"-"`
}

// GetLarge returns the Large field if non-nil, zero value otherwise.
func (o *ThumbnailSet) GetLarge() AnyOfmicrosoftGraphThumbnail {
	if o == nil || o.Large == nil {
		var ret AnyOfmicrosoftGraphThumbnail
		return ret
	}
	return *o.Large
}

// GetLargeOk returns a tuple with the Large field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ThumbnailSet) GetLargeOk() (AnyOfmicrosoftGraphThumbnail, bool) {
	if o == nil || o.Large == nil {
		var ret AnyOfmicrosoftGraphThumbnail
		return ret, false
	}
	return *o.Large, true
}

// HasLarge returns a boolean if a field has been set.
func (o *ThumbnailSet) HasLarge() bool {
	if o != nil && o.Large != nil {
		return true
	}

	return false
}

// SetLarge gets a reference to the given AnyOfmicrosoftGraphThumbnail and assigns it to the Large field.
func (o *ThumbnailSet) SetLarge(v AnyOfmicrosoftGraphThumbnail) {
	o.Large = &v
}

// SetLargeExplicitNull (un)sets Large to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Large value is set to nil even if false is passed
func (o *ThumbnailSet) SetLargeExplicitNull(b bool) {
	o.Large = nil
	o.isExplicitNullLarge = b
}
// GetMedium returns the Medium field if non-nil, zero value otherwise.
func (o *ThumbnailSet) GetMedium() AnyOfmicrosoftGraphThumbnail {
	if o == nil || o.Medium == nil {
		var ret AnyOfmicrosoftGraphThumbnail
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ThumbnailSet) GetMediumOk() (AnyOfmicrosoftGraphThumbnail, bool) {
	if o == nil || o.Medium == nil {
		var ret AnyOfmicrosoftGraphThumbnail
		return ret, false
	}
	return *o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *ThumbnailSet) HasMedium() bool {
	if o != nil && o.Medium != nil {
		return true
	}

	return false
}

// SetMedium gets a reference to the given AnyOfmicrosoftGraphThumbnail and assigns it to the Medium field.
func (o *ThumbnailSet) SetMedium(v AnyOfmicrosoftGraphThumbnail) {
	o.Medium = &v
}

// SetMediumExplicitNull (un)sets Medium to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Medium value is set to nil even if false is passed
func (o *ThumbnailSet) SetMediumExplicitNull(b bool) {
	o.Medium = nil
	o.isExplicitNullMedium = b
}
// GetSmall returns the Small field if non-nil, zero value otherwise.
func (o *ThumbnailSet) GetSmall() AnyOfmicrosoftGraphThumbnail {
	if o == nil || o.Small == nil {
		var ret AnyOfmicrosoftGraphThumbnail
		return ret
	}
	return *o.Small
}

// GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ThumbnailSet) GetSmallOk() (AnyOfmicrosoftGraphThumbnail, bool) {
	if o == nil || o.Small == nil {
		var ret AnyOfmicrosoftGraphThumbnail
		return ret, false
	}
	return *o.Small, true
}

// HasSmall returns a boolean if a field has been set.
func (o *ThumbnailSet) HasSmall() bool {
	if o != nil && o.Small != nil {
		return true
	}

	return false
}

// SetSmall gets a reference to the given AnyOfmicrosoftGraphThumbnail and assigns it to the Small field.
func (o *ThumbnailSet) SetSmall(v AnyOfmicrosoftGraphThumbnail) {
	o.Small = &v
}

// SetSmallExplicitNull (un)sets Small to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Small value is set to nil even if false is passed
func (o *ThumbnailSet) SetSmallExplicitNull(b bool) {
	o.Small = nil
	o.isExplicitNullSmall = b
}
// GetSource returns the Source field if non-nil, zero value otherwise.
func (o *ThumbnailSet) GetSource() AnyOfmicrosoftGraphThumbnail {
	if o == nil || o.Source == nil {
		var ret AnyOfmicrosoftGraphThumbnail
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ThumbnailSet) GetSourceOk() (AnyOfmicrosoftGraphThumbnail, bool) {
	if o == nil || o.Source == nil {
		var ret AnyOfmicrosoftGraphThumbnail
		return ret, false
	}
	return *o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ThumbnailSet) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given AnyOfmicrosoftGraphThumbnail and assigns it to the Source field.
func (o *ThumbnailSet) SetSource(v AnyOfmicrosoftGraphThumbnail) {
	o.Source = &v
}

// SetSourceExplicitNull (un)sets Source to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Source value is set to nil even if false is passed
func (o *ThumbnailSet) SetSourceExplicitNull(b bool) {
	o.Source = nil
	o.isExplicitNullSource = b
}

// MarshalJSON returns the JSON representation of the model.
func (o ThumbnailSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Large == nil {
		if o.isExplicitNullLarge {
			toSerialize["large"] = o.Large
		}
	} else {
		toSerialize["large"] = o.Large
	}
	if o.Medium == nil {
		if o.isExplicitNullMedium {
			toSerialize["medium"] = o.Medium
		}
	} else {
		toSerialize["medium"] = o.Medium
	}
	if o.Small == nil {
		if o.isExplicitNullSmall {
			toSerialize["small"] = o.Small
		}
	} else {
		toSerialize["small"] = o.Small
	}
	if o.Source == nil {
		if o.isExplicitNullSource {
			toSerialize["source"] = o.Source
		}
	} else {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}


