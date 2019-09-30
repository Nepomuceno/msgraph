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
// MicrosoftGraphAuditLogRoot struct for MicrosoftGraphAuditLogRoot
type MicrosoftGraphAuditLogRoot struct {
	Id *string `json:"id,omitempty"`

	SignIns *[]MicrosoftGraphSignIn `json:"signIns,omitempty"`

	DirectoryAudits *[]MicrosoftGraphDirectoryAudit `json:"directoryAudits,omitempty"`

	RestrictedSignIns *[]MicrosoftGraphRestrictedSignIn `json:"restrictedSignIns,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAuditLogRoot) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAuditLogRoot) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphAuditLogRoot) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphAuditLogRoot) SetId(v string) {
	o.Id = &v
}

// GetSignIns returns the SignIns field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAuditLogRoot) GetSignIns() []MicrosoftGraphSignIn {
	if o == nil || o.SignIns == nil {
		var ret []MicrosoftGraphSignIn
		return ret
	}
	return *o.SignIns
}

// GetSignInsOk returns a tuple with the SignIns field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAuditLogRoot) GetSignInsOk() ([]MicrosoftGraphSignIn, bool) {
	if o == nil || o.SignIns == nil {
		var ret []MicrosoftGraphSignIn
		return ret, false
	}
	return *o.SignIns, true
}

// HasSignIns returns a boolean if a field has been set.
func (o *MicrosoftGraphAuditLogRoot) HasSignIns() bool {
	if o != nil && o.SignIns != nil {
		return true
	}

	return false
}

// SetSignIns gets a reference to the given []MicrosoftGraphSignIn and assigns it to the SignIns field.
func (o *MicrosoftGraphAuditLogRoot) SetSignIns(v []MicrosoftGraphSignIn) {
	o.SignIns = &v
}

// GetDirectoryAudits returns the DirectoryAudits field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAuditLogRoot) GetDirectoryAudits() []MicrosoftGraphDirectoryAudit {
	if o == nil || o.DirectoryAudits == nil {
		var ret []MicrosoftGraphDirectoryAudit
		return ret
	}
	return *o.DirectoryAudits
}

// GetDirectoryAuditsOk returns a tuple with the DirectoryAudits field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAuditLogRoot) GetDirectoryAuditsOk() ([]MicrosoftGraphDirectoryAudit, bool) {
	if o == nil || o.DirectoryAudits == nil {
		var ret []MicrosoftGraphDirectoryAudit
		return ret, false
	}
	return *o.DirectoryAudits, true
}

// HasDirectoryAudits returns a boolean if a field has been set.
func (o *MicrosoftGraphAuditLogRoot) HasDirectoryAudits() bool {
	if o != nil && o.DirectoryAudits != nil {
		return true
	}

	return false
}

// SetDirectoryAudits gets a reference to the given []MicrosoftGraphDirectoryAudit and assigns it to the DirectoryAudits field.
func (o *MicrosoftGraphAuditLogRoot) SetDirectoryAudits(v []MicrosoftGraphDirectoryAudit) {
	o.DirectoryAudits = &v
}

// GetRestrictedSignIns returns the RestrictedSignIns field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAuditLogRoot) GetRestrictedSignIns() []MicrosoftGraphRestrictedSignIn {
	if o == nil || o.RestrictedSignIns == nil {
		var ret []MicrosoftGraphRestrictedSignIn
		return ret
	}
	return *o.RestrictedSignIns
}

// GetRestrictedSignInsOk returns a tuple with the RestrictedSignIns field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAuditLogRoot) GetRestrictedSignInsOk() ([]MicrosoftGraphRestrictedSignIn, bool) {
	if o == nil || o.RestrictedSignIns == nil {
		var ret []MicrosoftGraphRestrictedSignIn
		return ret, false
	}
	return *o.RestrictedSignIns, true
}

// HasRestrictedSignIns returns a boolean if a field has been set.
func (o *MicrosoftGraphAuditLogRoot) HasRestrictedSignIns() bool {
	if o != nil && o.RestrictedSignIns != nil {
		return true
	}

	return false
}

// SetRestrictedSignIns gets a reference to the given []MicrosoftGraphRestrictedSignIn and assigns it to the RestrictedSignIns field.
func (o *MicrosoftGraphAuditLogRoot) SetRestrictedSignIns(v []MicrosoftGraphRestrictedSignIn) {
	o.RestrictedSignIns = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphAuditLogRoot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SignIns != nil {
		toSerialize["signIns"] = o.SignIns
	}
	if o.DirectoryAudits != nil {
		toSerialize["directoryAudits"] = o.DirectoryAudits
	}
	if o.RestrictedSignIns != nil {
		toSerialize["restrictedSignIns"] = o.RestrictedSignIns
	}
	return json.Marshal(toSerialize)
}


