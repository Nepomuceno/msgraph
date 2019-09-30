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
// MicrosoftGraphPatternedRecurrence struct for MicrosoftGraphPatternedRecurrence
type MicrosoftGraphPatternedRecurrence struct {
	Pattern *AnyOfmicrosoftGraphRecurrencePattern `json:"pattern,omitempty"`
	isExplicitNullPattern bool `json:"-"`
	Range *AnyOfmicrosoftGraphRecurrenceRange `json:"range,omitempty"`
	isExplicitNullRange bool `json:"-"`
}

// GetPattern returns the Pattern field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPatternedRecurrence) GetPattern() AnyOfmicrosoftGraphRecurrencePattern {
	if o == nil || o.Pattern == nil {
		var ret AnyOfmicrosoftGraphRecurrencePattern
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPatternedRecurrence) GetPatternOk() (AnyOfmicrosoftGraphRecurrencePattern, bool) {
	if o == nil || o.Pattern == nil {
		var ret AnyOfmicrosoftGraphRecurrencePattern
		return ret, false
	}
	return *o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *MicrosoftGraphPatternedRecurrence) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given AnyOfmicrosoftGraphRecurrencePattern and assigns it to the Pattern field.
func (o *MicrosoftGraphPatternedRecurrence) SetPattern(v AnyOfmicrosoftGraphRecurrencePattern) {
	o.Pattern = &v
}

// SetPatternExplicitNull (un)sets Pattern to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Pattern value is set to nil even if false is passed
func (o *MicrosoftGraphPatternedRecurrence) SetPatternExplicitNull(b bool) {
	o.Pattern = nil
	o.isExplicitNullPattern = b
}
// GetRange returns the Range field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPatternedRecurrence) GetRange() AnyOfmicrosoftGraphRecurrenceRange {
	if o == nil || o.Range == nil {
		var ret AnyOfmicrosoftGraphRecurrenceRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPatternedRecurrence) GetRangeOk() (AnyOfmicrosoftGraphRecurrenceRange, bool) {
	if o == nil || o.Range == nil {
		var ret AnyOfmicrosoftGraphRecurrenceRange
		return ret, false
	}
	return *o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *MicrosoftGraphPatternedRecurrence) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given AnyOfmicrosoftGraphRecurrenceRange and assigns it to the Range field.
func (o *MicrosoftGraphPatternedRecurrence) SetRange(v AnyOfmicrosoftGraphRecurrenceRange) {
	o.Range = &v
}

// SetRangeExplicitNull (un)sets Range to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Range value is set to nil even if false is passed
func (o *MicrosoftGraphPatternedRecurrence) SetRangeExplicitNull(b bool) {
	o.Range = nil
	o.isExplicitNullRange = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphPatternedRecurrence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pattern == nil {
		if o.isExplicitNullPattern {
			toSerialize["pattern"] = o.Pattern
		}
	} else {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Range == nil {
		if o.isExplicitNullRange {
			toSerialize["range"] = o.Range
		}
	} else {
		toSerialize["range"] = o.Range
	}
	return json.Marshal(toSerialize)
}


