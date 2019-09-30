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
// MicrosoftGraphSiteCollection struct for MicrosoftGraphSiteCollection
type MicrosoftGraphSiteCollection struct {
	Hostname *string `json:"hostname,omitempty"`
	isExplicitNullHostname bool `json:"-"`
	Root *AnyOfobject `json:"root,omitempty"`
	isExplicitNullRoot bool `json:"-"`
}

// GetHostname returns the Hostname field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSiteCollection) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSiteCollection) GetHostnameOk() (string, bool) {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret, false
	}
	return *o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *MicrosoftGraphSiteCollection) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *MicrosoftGraphSiteCollection) SetHostname(v string) {
	o.Hostname = &v
}

// SetHostnameExplicitNull (un)sets Hostname to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Hostname value is set to nil even if false is passed
func (o *MicrosoftGraphSiteCollection) SetHostnameExplicitNull(b bool) {
	o.Hostname = nil
	o.isExplicitNullHostname = b
}
// GetRoot returns the Root field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSiteCollection) GetRoot() AnyOfobject {
	if o == nil || o.Root == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSiteCollection) GetRootOk() (AnyOfobject, bool) {
	if o == nil || o.Root == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *MicrosoftGraphSiteCollection) HasRoot() bool {
	if o != nil && o.Root != nil {
		return true
	}

	return false
}

// SetRoot gets a reference to the given AnyOfobject and assigns it to the Root field.
func (o *MicrosoftGraphSiteCollection) SetRoot(v AnyOfobject) {
	o.Root = &v
}

// SetRootExplicitNull (un)sets Root to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Root value is set to nil even if false is passed
func (o *MicrosoftGraphSiteCollection) SetRootExplicitNull(b bool) {
	o.Root = nil
	o.isExplicitNullRoot = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphSiteCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hostname == nil {
		if o.isExplicitNullHostname {
			toSerialize["hostname"] = o.Hostname
		}
	} else {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Root == nil {
		if o.isExplicitNullRoot {
			toSerialize["root"] = o.Root
		}
	} else {
		toSerialize["root"] = o.Root
	}
	return json.Marshal(toSerialize)
}

