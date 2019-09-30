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
// MicrosoftGraphRecentNotebookLinks struct for MicrosoftGraphRecentNotebookLinks
type MicrosoftGraphRecentNotebookLinks struct {
	OneNoteClientUrl *AnyOfmicrosoftGraphExternalLink `json:"oneNoteClientUrl,omitempty"`
	isExplicitNullOneNoteClientUrl bool `json:"-"`
	OneNoteWebUrl *AnyOfmicrosoftGraphExternalLink `json:"oneNoteWebUrl,omitempty"`
	isExplicitNullOneNoteWebUrl bool `json:"-"`
}

// GetOneNoteClientUrl returns the OneNoteClientUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphRecentNotebookLinks) GetOneNoteClientUrl() AnyOfmicrosoftGraphExternalLink {
	if o == nil || o.OneNoteClientUrl == nil {
		var ret AnyOfmicrosoftGraphExternalLink
		return ret
	}
	return *o.OneNoteClientUrl
}

// GetOneNoteClientUrlOk returns a tuple with the OneNoteClientUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRecentNotebookLinks) GetOneNoteClientUrlOk() (AnyOfmicrosoftGraphExternalLink, bool) {
	if o == nil || o.OneNoteClientUrl == nil {
		var ret AnyOfmicrosoftGraphExternalLink
		return ret, false
	}
	return *o.OneNoteClientUrl, true
}

// HasOneNoteClientUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphRecentNotebookLinks) HasOneNoteClientUrl() bool {
	if o != nil && o.OneNoteClientUrl != nil {
		return true
	}

	return false
}

// SetOneNoteClientUrl gets a reference to the given AnyOfmicrosoftGraphExternalLink and assigns it to the OneNoteClientUrl field.
func (o *MicrosoftGraphRecentNotebookLinks) SetOneNoteClientUrl(v AnyOfmicrosoftGraphExternalLink) {
	o.OneNoteClientUrl = &v
}

// SetOneNoteClientUrlExplicitNull (un)sets OneNoteClientUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OneNoteClientUrl value is set to nil even if false is passed
func (o *MicrosoftGraphRecentNotebookLinks) SetOneNoteClientUrlExplicitNull(b bool) {
	o.OneNoteClientUrl = nil
	o.isExplicitNullOneNoteClientUrl = b
}
// GetOneNoteWebUrl returns the OneNoteWebUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphRecentNotebookLinks) GetOneNoteWebUrl() AnyOfmicrosoftGraphExternalLink {
	if o == nil || o.OneNoteWebUrl == nil {
		var ret AnyOfmicrosoftGraphExternalLink
		return ret
	}
	return *o.OneNoteWebUrl
}

// GetOneNoteWebUrlOk returns a tuple with the OneNoteWebUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRecentNotebookLinks) GetOneNoteWebUrlOk() (AnyOfmicrosoftGraphExternalLink, bool) {
	if o == nil || o.OneNoteWebUrl == nil {
		var ret AnyOfmicrosoftGraphExternalLink
		return ret, false
	}
	return *o.OneNoteWebUrl, true
}

// HasOneNoteWebUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphRecentNotebookLinks) HasOneNoteWebUrl() bool {
	if o != nil && o.OneNoteWebUrl != nil {
		return true
	}

	return false
}

// SetOneNoteWebUrl gets a reference to the given AnyOfmicrosoftGraphExternalLink and assigns it to the OneNoteWebUrl field.
func (o *MicrosoftGraphRecentNotebookLinks) SetOneNoteWebUrl(v AnyOfmicrosoftGraphExternalLink) {
	o.OneNoteWebUrl = &v
}

// SetOneNoteWebUrlExplicitNull (un)sets OneNoteWebUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OneNoteWebUrl value is set to nil even if false is passed
func (o *MicrosoftGraphRecentNotebookLinks) SetOneNoteWebUrlExplicitNull(b bool) {
	o.OneNoteWebUrl = nil
	o.isExplicitNullOneNoteWebUrl = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphRecentNotebookLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OneNoteClientUrl == nil {
		if o.isExplicitNullOneNoteClientUrl {
			toSerialize["oneNoteClientUrl"] = o.OneNoteClientUrl
		}
	} else {
		toSerialize["oneNoteClientUrl"] = o.OneNoteClientUrl
	}
	if o.OneNoteWebUrl == nil {
		if o.isExplicitNullOneNoteWebUrl {
			toSerialize["oneNoteWebUrl"] = o.OneNoteWebUrl
		}
	} else {
		toSerialize["oneNoteWebUrl"] = o.OneNoteWebUrl
	}
	return json.Marshal(toSerialize)
}


