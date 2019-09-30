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
// ManagedMobileLobApp An abstract base class containing properties for all managed mobile line of business apps.
type ManagedMobileLobApp struct {
	// The internal committed content version.
	CommittedContentVersion *string `json:"committedContentVersion,omitempty"`
	isExplicitNullCommittedContentVersion bool `json:"-"`
	// The name of the main Lob application file.
	FileName *string `json:"fileName,omitempty"`
	isExplicitNullFileName bool `json:"-"`
	// The total size, including all uploaded files.
	Size *int64 `json:"size,omitempty"`

	ContentVersions *[]MicrosoftGraphMobileAppContent `json:"contentVersions,omitempty"`

}

// GetCommittedContentVersion returns the CommittedContentVersion field if non-nil, zero value otherwise.
func (o *ManagedMobileLobApp) GetCommittedContentVersion() string {
	if o == nil || o.CommittedContentVersion == nil {
		var ret string
		return ret
	}
	return *o.CommittedContentVersion
}

// GetCommittedContentVersionOk returns a tuple with the CommittedContentVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMobileLobApp) GetCommittedContentVersionOk() (string, bool) {
	if o == nil || o.CommittedContentVersion == nil {
		var ret string
		return ret, false
	}
	return *o.CommittedContentVersion, true
}

// HasCommittedContentVersion returns a boolean if a field has been set.
func (o *ManagedMobileLobApp) HasCommittedContentVersion() bool {
	if o != nil && o.CommittedContentVersion != nil {
		return true
	}

	return false
}

// SetCommittedContentVersion gets a reference to the given string and assigns it to the CommittedContentVersion field.
func (o *ManagedMobileLobApp) SetCommittedContentVersion(v string) {
	o.CommittedContentVersion = &v
}

// SetCommittedContentVersionExplicitNull (un)sets CommittedContentVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CommittedContentVersion value is set to nil even if false is passed
func (o *ManagedMobileLobApp) SetCommittedContentVersionExplicitNull(b bool) {
	o.CommittedContentVersion = nil
	o.isExplicitNullCommittedContentVersion = b
}
// GetFileName returns the FileName field if non-nil, zero value otherwise.
func (o *ManagedMobileLobApp) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMobileLobApp) GetFileNameOk() (string, bool) {
	if o == nil || o.FileName == nil {
		var ret string
		return ret, false
	}
	return *o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ManagedMobileLobApp) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ManagedMobileLobApp) SetFileName(v string) {
	o.FileName = &v
}

// SetFileNameExplicitNull (un)sets FileName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The FileName value is set to nil even if false is passed
func (o *ManagedMobileLobApp) SetFileNameExplicitNull(b bool) {
	o.FileName = nil
	o.isExplicitNullFileName = b
}
// GetSize returns the Size field if non-nil, zero value otherwise.
func (o *ManagedMobileLobApp) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMobileLobApp) GetSizeOk() (int64, bool) {
	if o == nil || o.Size == nil {
		var ret int64
		return ret, false
	}
	return *o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ManagedMobileLobApp) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ManagedMobileLobApp) SetSize(v int64) {
	o.Size = &v
}

// GetContentVersions returns the ContentVersions field if non-nil, zero value otherwise.
func (o *ManagedMobileLobApp) GetContentVersions() []MicrosoftGraphMobileAppContent {
	if o == nil || o.ContentVersions == nil {
		var ret []MicrosoftGraphMobileAppContent
		return ret
	}
	return *o.ContentVersions
}

// GetContentVersionsOk returns a tuple with the ContentVersions field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMobileLobApp) GetContentVersionsOk() ([]MicrosoftGraphMobileAppContent, bool) {
	if o == nil || o.ContentVersions == nil {
		var ret []MicrosoftGraphMobileAppContent
		return ret, false
	}
	return *o.ContentVersions, true
}

// HasContentVersions returns a boolean if a field has been set.
func (o *ManagedMobileLobApp) HasContentVersions() bool {
	if o != nil && o.ContentVersions != nil {
		return true
	}

	return false
}

// SetContentVersions gets a reference to the given []MicrosoftGraphMobileAppContent and assigns it to the ContentVersions field.
func (o *ManagedMobileLobApp) SetContentVersions(v []MicrosoftGraphMobileAppContent) {
	o.ContentVersions = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o ManagedMobileLobApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommittedContentVersion == nil {
		if o.isExplicitNullCommittedContentVersion {
			toSerialize["committedContentVersion"] = o.CommittedContentVersion
		}
	} else {
		toSerialize["committedContentVersion"] = o.CommittedContentVersion
	}
	if o.FileName == nil {
		if o.isExplicitNullFileName {
			toSerialize["fileName"] = o.FileName
		}
	} else {
		toSerialize["fileName"] = o.FileName
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.ContentVersions != nil {
		toSerialize["contentVersions"] = o.ContentVersions
	}
	return json.Marshal(toSerialize)
}

