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
// InlineObject421 struct for InlineObject421
type InlineObject421 struct {
	Commands *[]AnyOfmicrosoftGraphOnenotePatchContentCommand `json:"commands,omitempty"`

}

// GetCommands returns the Commands field if non-nil, zero value otherwise.
func (o *InlineObject421) GetCommands() []AnyOfmicrosoftGraphOnenotePatchContentCommand {
	if o == nil || o.Commands == nil {
		var ret []AnyOfmicrosoftGraphOnenotePatchContentCommand
		return ret
	}
	return *o.Commands
}

// GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject421) GetCommandsOk() ([]AnyOfmicrosoftGraphOnenotePatchContentCommand, bool) {
	if o == nil || o.Commands == nil {
		var ret []AnyOfmicrosoftGraphOnenotePatchContentCommand
		return ret, false
	}
	return *o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *InlineObject421) HasCommands() bool {
	if o != nil && o.Commands != nil {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []AnyOfmicrosoftGraphOnenotePatchContentCommand and assigns it to the Commands field.
func (o *InlineObject421) SetCommands(v []AnyOfmicrosoftGraphOnenotePatchContentCommand) {
	o.Commands = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject421) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Commands != nil {
		toSerialize["commands"] = o.Commands
	}
	return json.Marshal(toSerialize)
}


