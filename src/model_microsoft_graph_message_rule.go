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
// MicrosoftGraphMessageRule struct for MicrosoftGraphMessageRule
type MicrosoftGraphMessageRule struct {
	Id *string `json:"id,omitempty"`

	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	Sequence *int32 `json:"sequence,omitempty"`
	isExplicitNullSequence bool `json:"-"`
	Conditions *AnyOfmicrosoftGraphMessageRulePredicates `json:"conditions,omitempty"`
	isExplicitNullConditions bool `json:"-"`
	Actions *AnyOfmicrosoftGraphMessageRuleActions `json:"actions,omitempty"`
	isExplicitNullActions bool `json:"-"`
	Exceptions *AnyOfmicrosoftGraphMessageRulePredicates `json:"exceptions,omitempty"`
	isExplicitNullExceptions bool `json:"-"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	isExplicitNullIsEnabled bool `json:"-"`
	HasError *bool `json:"hasError,omitempty"`
	isExplicitNullHasError bool `json:"-"`
	IsReadOnly *bool `json:"isReadOnly,omitempty"`
	isExplicitNullIsReadOnly bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphMessageRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMessageRule) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphMessageRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphMessageRule) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphMessageRule) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMessageRule) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphMessageRule) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphMessageRule) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphMessageRule) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetSequence returns the Sequence field if non-nil, zero value otherwise.
func (o *MicrosoftGraphMessageRule) GetSequence() int32 {
	if o == nil || o.Sequence == nil {
		var ret int32
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMessageRule) GetSequenceOk() (int32, bool) {
	if o == nil || o.Sequence == nil {
		var ret int32
		return ret, false
	}
	return *o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *MicrosoftGraphMessageRule) HasSequence() bool {
	if o != nil && o.Sequence != nil {
		return true
	}

	return false
}

// SetSequence gets a reference to the given int32 and assigns it to the Sequence field.
func (o *MicrosoftGraphMessageRule) SetSequence(v int32) {
	o.Sequence = &v
}

// SetSequenceExplicitNull (un)sets Sequence to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Sequence value is set to nil even if false is passed
func (o *MicrosoftGraphMessageRule) SetSequenceExplicitNull(b bool) {
	o.Sequence = nil
	o.isExplicitNullSequence = b
}
// GetConditions returns the Conditions field if non-nil, zero value otherwise.
func (o *MicrosoftGraphMessageRule) GetConditions() AnyOfmicrosoftGraphMessageRulePredicates {
	if o == nil || o.Conditions == nil {
		var ret AnyOfmicrosoftGraphMessageRulePredicates
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMessageRule) GetConditionsOk() (AnyOfmicrosoftGraphMessageRulePredicates, bool) {
	if o == nil || o.Conditions == nil {
		var ret AnyOfmicrosoftGraphMessageRulePredicates
		return ret, false
	}
	return *o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *MicrosoftGraphMessageRule) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given AnyOfmicrosoftGraphMessageRulePredicates and assigns it to the Conditions field.
func (o *MicrosoftGraphMessageRule) SetConditions(v AnyOfmicrosoftGraphMessageRulePredicates) {
	o.Conditions = &v
}

// SetConditionsExplicitNull (un)sets Conditions to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Conditions value is set to nil even if false is passed
func (o *MicrosoftGraphMessageRule) SetConditionsExplicitNull(b bool) {
	o.Conditions = nil
	o.isExplicitNullConditions = b
}
// GetActions returns the Actions field if non-nil, zero value otherwise.
func (o *MicrosoftGraphMessageRule) GetActions() AnyOfmicrosoftGraphMessageRuleActions {
	if o == nil || o.Actions == nil {
		var ret AnyOfmicrosoftGraphMessageRuleActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMessageRule) GetActionsOk() (AnyOfmicrosoftGraphMessageRuleActions, bool) {
	if o == nil || o.Actions == nil {
		var ret AnyOfmicrosoftGraphMessageRuleActions
		return ret, false
	}
	return *o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *MicrosoftGraphMessageRule) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given AnyOfmicrosoftGraphMessageRuleActions and assigns it to the Actions field.
func (o *MicrosoftGraphMessageRule) SetActions(v AnyOfmicrosoftGraphMessageRuleActions) {
	o.Actions = &v
}

// SetActionsExplicitNull (un)sets Actions to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Actions value is set to nil even if false is passed
func (o *MicrosoftGraphMessageRule) SetActionsExplicitNull(b bool) {
	o.Actions = nil
	o.isExplicitNullActions = b
}
// GetExceptions returns the Exceptions field if non-nil, zero value otherwise.
func (o *MicrosoftGraphMessageRule) GetExceptions() AnyOfmicrosoftGraphMessageRulePredicates {
	if o == nil || o.Exceptions == nil {
		var ret AnyOfmicrosoftGraphMessageRulePredicates
		return ret
	}
	return *o.Exceptions
}

// GetExceptionsOk returns a tuple with the Exceptions field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMessageRule) GetExceptionsOk() (AnyOfmicrosoftGraphMessageRulePredicates, bool) {
	if o == nil || o.Exceptions == nil {
		var ret AnyOfmicrosoftGraphMessageRulePredicates
		return ret, false
	}
	return *o.Exceptions, true
}

// HasExceptions returns a boolean if a field has been set.
func (o *MicrosoftGraphMessageRule) HasExceptions() bool {
	if o != nil && o.Exceptions != nil {
		return true
	}

	return false
}

// SetExceptions gets a reference to the given AnyOfmicrosoftGraphMessageRulePredicates and assigns it to the Exceptions field.
func (o *MicrosoftGraphMessageRule) SetExceptions(v AnyOfmicrosoftGraphMessageRulePredicates) {
	o.Exceptions = &v
}

// SetExceptionsExplicitNull (un)sets Exceptions to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Exceptions value is set to nil even if false is passed
func (o *MicrosoftGraphMessageRule) SetExceptionsExplicitNull(b bool) {
	o.Exceptions = nil
	o.isExplicitNullExceptions = b
}
// GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.
func (o *MicrosoftGraphMessageRule) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMessageRule) GetIsEnabledOk() (bool, bool) {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret, false
	}
	return *o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *MicrosoftGraphMessageRule) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *MicrosoftGraphMessageRule) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// SetIsEnabledExplicitNull (un)sets IsEnabled to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The IsEnabled value is set to nil even if false is passed
func (o *MicrosoftGraphMessageRule) SetIsEnabledExplicitNull(b bool) {
	o.IsEnabled = nil
	o.isExplicitNullIsEnabled = b
}
// GetHasError returns the HasError field if non-nil, zero value otherwise.
func (o *MicrosoftGraphMessageRule) GetHasError() bool {
	if o == nil || o.HasError == nil {
		var ret bool
		return ret
	}
	return *o.HasError
}

// GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMessageRule) GetHasErrorOk() (bool, bool) {
	if o == nil || o.HasError == nil {
		var ret bool
		return ret, false
	}
	return *o.HasError, true
}

// HasHasError returns a boolean if a field has been set.
func (o *MicrosoftGraphMessageRule) HasHasError() bool {
	if o != nil && o.HasError != nil {
		return true
	}

	return false
}

// SetHasError gets a reference to the given bool and assigns it to the HasError field.
func (o *MicrosoftGraphMessageRule) SetHasError(v bool) {
	o.HasError = &v
}

// SetHasErrorExplicitNull (un)sets HasError to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The HasError value is set to nil even if false is passed
func (o *MicrosoftGraphMessageRule) SetHasErrorExplicitNull(b bool) {
	o.HasError = nil
	o.isExplicitNullHasError = b
}
// GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.
func (o *MicrosoftGraphMessageRule) GetIsReadOnly() bool {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMessageRule) GetIsReadOnlyOk() (bool, bool) {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret, false
	}
	return *o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *MicrosoftGraphMessageRule) HasIsReadOnly() bool {
	if o != nil && o.IsReadOnly != nil {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *MicrosoftGraphMessageRule) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// SetIsReadOnlyExplicitNull (un)sets IsReadOnly to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The IsReadOnly value is set to nil even if false is passed
func (o *MicrosoftGraphMessageRule) SetIsReadOnlyExplicitNull(b bool) {
	o.IsReadOnly = nil
	o.isExplicitNullIsReadOnly = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphMessageRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Sequence == nil {
		if o.isExplicitNullSequence {
			toSerialize["sequence"] = o.Sequence
		}
	} else {
		toSerialize["sequence"] = o.Sequence
	}
	if o.Conditions == nil {
		if o.isExplicitNullConditions {
			toSerialize["conditions"] = o.Conditions
		}
	} else {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Actions == nil {
		if o.isExplicitNullActions {
			toSerialize["actions"] = o.Actions
		}
	} else {
		toSerialize["actions"] = o.Actions
	}
	if o.Exceptions == nil {
		if o.isExplicitNullExceptions {
			toSerialize["exceptions"] = o.Exceptions
		}
	} else {
		toSerialize["exceptions"] = o.Exceptions
	}
	if o.IsEnabled == nil {
		if o.isExplicitNullIsEnabled {
			toSerialize["isEnabled"] = o.IsEnabled
		}
	} else {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.HasError == nil {
		if o.isExplicitNullHasError {
			toSerialize["hasError"] = o.HasError
		}
	} else {
		toSerialize["hasError"] = o.HasError
	}
	if o.IsReadOnly == nil {
		if o.isExplicitNullIsReadOnly {
			toSerialize["isReadOnly"] = o.IsReadOnly
		}
	} else {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	return json.Marshal(toSerialize)
}


