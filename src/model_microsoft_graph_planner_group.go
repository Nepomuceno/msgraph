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
// MicrosoftGraphPlannerGroup struct for MicrosoftGraphPlannerGroup
type MicrosoftGraphPlannerGroup struct {
	Id *string `json:"id,omitempty"`

	Plans *[]MicrosoftGraphPlannerPlan `json:"plans,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPlannerGroup) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerGroup) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphPlannerGroup) SetId(v string) {
	o.Id = &v
}

// GetPlans returns the Plans field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPlannerGroup) GetPlans() []MicrosoftGraphPlannerPlan {
	if o == nil || o.Plans == nil {
		var ret []MicrosoftGraphPlannerPlan
		return ret
	}
	return *o.Plans
}

// GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerGroup) GetPlansOk() ([]MicrosoftGraphPlannerPlan, bool) {
	if o == nil || o.Plans == nil {
		var ret []MicrosoftGraphPlannerPlan
		return ret, false
	}
	return *o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerGroup) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []MicrosoftGraphPlannerPlan and assigns it to the Plans field.
func (o *MicrosoftGraphPlannerGroup) SetPlans(v []MicrosoftGraphPlannerPlan) {
	o.Plans = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphPlannerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Plans != nil {
		toSerialize["plans"] = o.Plans
	}
	return json.Marshal(toSerialize)
}


