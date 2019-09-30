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
// MicrosoftGraphAppCatalogs struct for MicrosoftGraphAppCatalogs
type MicrosoftGraphAppCatalogs struct {
	Id *string `json:"id,omitempty"`

	TeamsApps *[]MicrosoftGraphTeamsApp `json:"teamsApps,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAppCatalogs) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppCatalogs) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphAppCatalogs) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphAppCatalogs) SetId(v string) {
	o.Id = &v
}

// GetTeamsApps returns the TeamsApps field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAppCatalogs) GetTeamsApps() []MicrosoftGraphTeamsApp {
	if o == nil || o.TeamsApps == nil {
		var ret []MicrosoftGraphTeamsApp
		return ret
	}
	return *o.TeamsApps
}

// GetTeamsAppsOk returns a tuple with the TeamsApps field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppCatalogs) GetTeamsAppsOk() ([]MicrosoftGraphTeamsApp, bool) {
	if o == nil || o.TeamsApps == nil {
		var ret []MicrosoftGraphTeamsApp
		return ret, false
	}
	return *o.TeamsApps, true
}

// HasTeamsApps returns a boolean if a field has been set.
func (o *MicrosoftGraphAppCatalogs) HasTeamsApps() bool {
	if o != nil && o.TeamsApps != nil {
		return true
	}

	return false
}

// SetTeamsApps gets a reference to the given []MicrosoftGraphTeamsApp and assigns it to the TeamsApps field.
func (o *MicrosoftGraphAppCatalogs) SetTeamsApps(v []MicrosoftGraphTeamsApp) {
	o.TeamsApps = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphAppCatalogs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TeamsApps != nil {
		toSerialize["teamsApps"] = o.TeamsApps
	}
	return json.Marshal(toSerialize)
}

