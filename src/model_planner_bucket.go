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
// PlannerBucket struct for PlannerBucket
type PlannerBucket struct {
	Name *string `json:"name,omitempty"`

	PlanId *string `json:"planId,omitempty"`
	isExplicitNullPlanId bool `json:"-"`
	OrderHint *string `json:"orderHint,omitempty"`
	isExplicitNullOrderHint bool `json:"-"`
	Tasks *[]MicrosoftGraphPlannerTask `json:"tasks,omitempty"`

}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *PlannerBucket) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PlannerBucket) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlannerBucket) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlannerBucket) SetName(v string) {
	o.Name = &v
}

// GetPlanId returns the PlanId field if non-nil, zero value otherwise.
func (o *PlannerBucket) GetPlanId() string {
	if o == nil || o.PlanId == nil {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PlannerBucket) GetPlanIdOk() (string, bool) {
	if o == nil || o.PlanId == nil {
		var ret string
		return ret, false
	}
	return *o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *PlannerBucket) HasPlanId() bool {
	if o != nil && o.PlanId != nil {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *PlannerBucket) SetPlanId(v string) {
	o.PlanId = &v
}

// SetPlanIdExplicitNull (un)sets PlanId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PlanId value is set to nil even if false is passed
func (o *PlannerBucket) SetPlanIdExplicitNull(b bool) {
	o.PlanId = nil
	o.isExplicitNullPlanId = b
}
// GetOrderHint returns the OrderHint field if non-nil, zero value otherwise.
func (o *PlannerBucket) GetOrderHint() string {
	if o == nil || o.OrderHint == nil {
		var ret string
		return ret
	}
	return *o.OrderHint
}

// GetOrderHintOk returns a tuple with the OrderHint field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PlannerBucket) GetOrderHintOk() (string, bool) {
	if o == nil || o.OrderHint == nil {
		var ret string
		return ret, false
	}
	return *o.OrderHint, true
}

// HasOrderHint returns a boolean if a field has been set.
func (o *PlannerBucket) HasOrderHint() bool {
	if o != nil && o.OrderHint != nil {
		return true
	}

	return false
}

// SetOrderHint gets a reference to the given string and assigns it to the OrderHint field.
func (o *PlannerBucket) SetOrderHint(v string) {
	o.OrderHint = &v
}

// SetOrderHintExplicitNull (un)sets OrderHint to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OrderHint value is set to nil even if false is passed
func (o *PlannerBucket) SetOrderHintExplicitNull(b bool) {
	o.OrderHint = nil
	o.isExplicitNullOrderHint = b
}
// GetTasks returns the Tasks field if non-nil, zero value otherwise.
func (o *PlannerBucket) GetTasks() []MicrosoftGraphPlannerTask {
	if o == nil || o.Tasks == nil {
		var ret []MicrosoftGraphPlannerTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *PlannerBucket) GetTasksOk() ([]MicrosoftGraphPlannerTask, bool) {
	if o == nil || o.Tasks == nil {
		var ret []MicrosoftGraphPlannerTask
		return ret, false
	}
	return *o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *PlannerBucket) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []MicrosoftGraphPlannerTask and assigns it to the Tasks field.
func (o *PlannerBucket) SetTasks(v []MicrosoftGraphPlannerTask) {
	o.Tasks = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o PlannerBucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PlanId == nil {
		if o.isExplicitNullPlanId {
			toSerialize["planId"] = o.PlanId
		}
	} else {
		toSerialize["planId"] = o.PlanId
	}
	if o.OrderHint == nil {
		if o.isExplicitNullOrderHint {
			toSerialize["orderHint"] = o.OrderHint
		}
	} else {
		toSerialize["orderHint"] = o.OrderHint
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}


