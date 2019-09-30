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
// MobileAppContent Contains content properties for a specific app version. Each mobileAppContent can have multiple mobileAppContentFile.
type MobileAppContent struct {
	Files *[]MicrosoftGraphMobileAppContentFile `json:"files,omitempty"`

}

// GetFiles returns the Files field if non-nil, zero value otherwise.
func (o *MobileAppContent) GetFiles() []MicrosoftGraphMobileAppContentFile {
	if o == nil || o.Files == nil {
		var ret []MicrosoftGraphMobileAppContentFile
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MobileAppContent) GetFilesOk() ([]MicrosoftGraphMobileAppContentFile, bool) {
	if o == nil || o.Files == nil {
		var ret []MicrosoftGraphMobileAppContentFile
		return ret, false
	}
	return *o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *MobileAppContent) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []MicrosoftGraphMobileAppContentFile and assigns it to the Files field.
func (o *MobileAppContent) SetFiles(v []MicrosoftGraphMobileAppContentFile) {
	o.Files = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MobileAppContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	return json.Marshal(toSerialize)
}

