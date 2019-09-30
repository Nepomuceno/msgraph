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
// PlannerGroup struct for PlannerGroup
type PlannerGroup struct {
	Plans *[]MicrosoftGraphPlannerPlan `json:"plans,omitempty"`

}

// GetPlans returns the Plans field if non-nil, zero value otherwise.
func (o *PlannerGroup) GetPlans() []MicrosoftGraphPlannerPlan {
	if o == nil || o.Plans == nil {
		var ret []MicrosoftGraphPlannerPlan
		return ret
	}
	return *o.Plans
}

// GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PlannerGroup) GetPlansOk() ([]MicrosoftGraphPlannerPlan, bool) {
	if o == nil || o.Plans == nil {
		var ret []MicrosoftGraphPlannerPlan
		return ret, false
	}
	return *o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *PlannerGroup) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []MicrosoftGraphPlannerPlan and assigns it to the Plans field.
func (o *PlannerGroup) SetPlans(v []MicrosoftGraphPlannerPlan) {
	o.Plans = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o PlannerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Plans != nil {
		toSerialize["plans"] = o.Plans
	}
	return json.Marshal(toSerialize)
}

