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
// MicrosoftGraphPlannerCategoryDescriptions struct for MicrosoftGraphPlannerCategoryDescriptions
type MicrosoftGraphPlannerCategoryDescriptions struct {
	Category1 *string `json:"category1,omitempty"`
	isExplicitNullCategory1 bool `json:"-"`
	Category2 *string `json:"category2,omitempty"`
	isExplicitNullCategory2 bool `json:"-"`
	Category3 *string `json:"category3,omitempty"`
	isExplicitNullCategory3 bool `json:"-"`
	Category4 *string `json:"category4,omitempty"`
	isExplicitNullCategory4 bool `json:"-"`
	Category5 *string `json:"category5,omitempty"`
	isExplicitNullCategory5 bool `json:"-"`
	Category6 *string `json:"category6,omitempty"`
	isExplicitNullCategory6 bool `json:"-"`
}

// GetCategory1 returns the Category1 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPlannerCategoryDescriptions) GetCategory1() string {
	if o == nil || o.Category1 == nil {
		var ret string
		return ret
	}
	return *o.Category1
}

// GetCategory1Ok returns a tuple with the Category1 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerCategoryDescriptions) GetCategory1Ok() (string, bool) {
	if o == nil || o.Category1 == nil {
		var ret string
		return ret, false
	}
	return *o.Category1, true
}

// HasCategory1 returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerCategoryDescriptions) HasCategory1() bool {
	if o != nil && o.Category1 != nil {
		return true
	}

	return false
}

// SetCategory1 gets a reference to the given string and assigns it to the Category1 field.
func (o *MicrosoftGraphPlannerCategoryDescriptions) SetCategory1(v string) {
	o.Category1 = &v
}

// SetCategory1ExplicitNull (un)sets Category1 to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Category1 value is set to nil even if false is passed
func (o *MicrosoftGraphPlannerCategoryDescriptions) SetCategory1ExplicitNull(b bool) {
	o.Category1 = nil
	o.isExplicitNullCategory1 = b
}
// GetCategory2 returns the Category2 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPlannerCategoryDescriptions) GetCategory2() string {
	if o == nil || o.Category2 == nil {
		var ret string
		return ret
	}
	return *o.Category2
}

// GetCategory2Ok returns a tuple with the Category2 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerCategoryDescriptions) GetCategory2Ok() (string, bool) {
	if o == nil || o.Category2 == nil {
		var ret string
		return ret, false
	}
	return *o.Category2, true
}

// HasCategory2 returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerCategoryDescriptions) HasCategory2() bool {
	if o != nil && o.Category2 != nil {
		return true
	}

	return false
}

// SetCategory2 gets a reference to the given string and assigns it to the Category2 field.
func (o *MicrosoftGraphPlannerCategoryDescriptions) SetCategory2(v string) {
	o.Category2 = &v
}

// SetCategory2ExplicitNull (un)sets Category2 to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Category2 value is set to nil even if false is passed
func (o *MicrosoftGraphPlannerCategoryDescriptions) SetCategory2ExplicitNull(b bool) {
	o.Category2 = nil
	o.isExplicitNullCategory2 = b
}
// GetCategory3 returns the Category3 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPlannerCategoryDescriptions) GetCategory3() string {
	if o == nil || o.Category3 == nil {
		var ret string
		return ret
	}
	return *o.Category3
}

// GetCategory3Ok returns a tuple with the Category3 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerCategoryDescriptions) GetCategory3Ok() (string, bool) {
	if o == nil || o.Category3 == nil {
		var ret string
		return ret, false
	}
	return *o.Category3, true
}

// HasCategory3 returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerCategoryDescriptions) HasCategory3() bool {
	if o != nil && o.Category3 != nil {
		return true
	}

	return false
}

// SetCategory3 gets a reference to the given string and assigns it to the Category3 field.
func (o *MicrosoftGraphPlannerCategoryDescriptions) SetCategory3(v string) {
	o.Category3 = &v
}

// SetCategory3ExplicitNull (un)sets Category3 to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Category3 value is set to nil even if false is passed
func (o *MicrosoftGraphPlannerCategoryDescriptions) SetCategory3ExplicitNull(b bool) {
	o.Category3 = nil
	o.isExplicitNullCategory3 = b
}
// GetCategory4 returns the Category4 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPlannerCategoryDescriptions) GetCategory4() string {
	if o == nil || o.Category4 == nil {
		var ret string
		return ret
	}
	return *o.Category4
}

// GetCategory4Ok returns a tuple with the Category4 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerCategoryDescriptions) GetCategory4Ok() (string, bool) {
	if o == nil || o.Category4 == nil {
		var ret string
		return ret, false
	}
	return *o.Category4, true
}

// HasCategory4 returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerCategoryDescriptions) HasCategory4() bool {
	if o != nil && o.Category4 != nil {
		return true
	}

	return false
}

// SetCategory4 gets a reference to the given string and assigns it to the Category4 field.
func (o *MicrosoftGraphPlannerCategoryDescriptions) SetCategory4(v string) {
	o.Category4 = &v
}

// SetCategory4ExplicitNull (un)sets Category4 to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Category4 value is set to nil even if false is passed
func (o *MicrosoftGraphPlannerCategoryDescriptions) SetCategory4ExplicitNull(b bool) {
	o.Category4 = nil
	o.isExplicitNullCategory4 = b
}
// GetCategory5 returns the Category5 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPlannerCategoryDescriptions) GetCategory5() string {
	if o == nil || o.Category5 == nil {
		var ret string
		return ret
	}
	return *o.Category5
}

// GetCategory5Ok returns a tuple with the Category5 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerCategoryDescriptions) GetCategory5Ok() (string, bool) {
	if o == nil || o.Category5 == nil {
		var ret string
		return ret, false
	}
	return *o.Category5, true
}

// HasCategory5 returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerCategoryDescriptions) HasCategory5() bool {
	if o != nil && o.Category5 != nil {
		return true
	}

	return false
}

// SetCategory5 gets a reference to the given string and assigns it to the Category5 field.
func (o *MicrosoftGraphPlannerCategoryDescriptions) SetCategory5(v string) {
	o.Category5 = &v
}

// SetCategory5ExplicitNull (un)sets Category5 to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Category5 value is set to nil even if false is passed
func (o *MicrosoftGraphPlannerCategoryDescriptions) SetCategory5ExplicitNull(b bool) {
	o.Category5 = nil
	o.isExplicitNullCategory5 = b
}
// GetCategory6 returns the Category6 field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPlannerCategoryDescriptions) GetCategory6() string {
	if o == nil || o.Category6 == nil {
		var ret string
		return ret
	}
	return *o.Category6
}

// GetCategory6Ok returns a tuple with the Category6 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerCategoryDescriptions) GetCategory6Ok() (string, bool) {
	if o == nil || o.Category6 == nil {
		var ret string
		return ret, false
	}
	return *o.Category6, true
}

// HasCategory6 returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerCategoryDescriptions) HasCategory6() bool {
	if o != nil && o.Category6 != nil {
		return true
	}

	return false
}

// SetCategory6 gets a reference to the given string and assigns it to the Category6 field.
func (o *MicrosoftGraphPlannerCategoryDescriptions) SetCategory6(v string) {
	o.Category6 = &v
}

// SetCategory6ExplicitNull (un)sets Category6 to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Category6 value is set to nil even if false is passed
func (o *MicrosoftGraphPlannerCategoryDescriptions) SetCategory6ExplicitNull(b bool) {
	o.Category6 = nil
	o.isExplicitNullCategory6 = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphPlannerCategoryDescriptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category1 == nil {
		if o.isExplicitNullCategory1 {
			toSerialize["category1"] = o.Category1
		}
	} else {
		toSerialize["category1"] = o.Category1
	}
	if o.Category2 == nil {
		if o.isExplicitNullCategory2 {
			toSerialize["category2"] = o.Category2
		}
	} else {
		toSerialize["category2"] = o.Category2
	}
	if o.Category3 == nil {
		if o.isExplicitNullCategory3 {
			toSerialize["category3"] = o.Category3
		}
	} else {
		toSerialize["category3"] = o.Category3
	}
	if o.Category4 == nil {
		if o.isExplicitNullCategory4 {
			toSerialize["category4"] = o.Category4
		}
	} else {
		toSerialize["category4"] = o.Category4
	}
	if o.Category5 == nil {
		if o.isExplicitNullCategory5 {
			toSerialize["category5"] = o.Category5
		}
	} else {
		toSerialize["category5"] = o.Category5
	}
	if o.Category6 == nil {
		if o.isExplicitNullCategory6 {
			toSerialize["category6"] = o.Category6
		}
	} else {
		toSerialize["category6"] = o.Category6
	}
	return json.Marshal(toSerialize)
}

