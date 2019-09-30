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
// MicrosoftGraphDriveItemUploadableProperties struct for MicrosoftGraphDriveItemUploadableProperties
type MicrosoftGraphDriveItemUploadableProperties struct {
	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
	FileSystemInfo *AnyOfmicrosoftGraphFileSystemInfo `json:"fileSystemInfo,omitempty"`
	isExplicitNullFileSystemInfo bool `json:"-"`
	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDriveItemUploadableProperties) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDriveItemUploadableProperties) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphDriveItemUploadableProperties) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphDriveItemUploadableProperties) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphDriveItemUploadableProperties) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetFileSystemInfo returns the FileSystemInfo field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDriveItemUploadableProperties) GetFileSystemInfo() AnyOfmicrosoftGraphFileSystemInfo {
	if o == nil || o.FileSystemInfo == nil {
		var ret AnyOfmicrosoftGraphFileSystemInfo
		return ret
	}
	return *o.FileSystemInfo
}

// GetFileSystemInfoOk returns a tuple with the FileSystemInfo field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDriveItemUploadableProperties) GetFileSystemInfoOk() (AnyOfmicrosoftGraphFileSystemInfo, bool) {
	if o == nil || o.FileSystemInfo == nil {
		var ret AnyOfmicrosoftGraphFileSystemInfo
		return ret, false
	}
	return *o.FileSystemInfo, true
}

// HasFileSystemInfo returns a boolean if a field has been set.
func (o *MicrosoftGraphDriveItemUploadableProperties) HasFileSystemInfo() bool {
	if o != nil && o.FileSystemInfo != nil {
		return true
	}

	return false
}

// SetFileSystemInfo gets a reference to the given AnyOfmicrosoftGraphFileSystemInfo and assigns it to the FileSystemInfo field.
func (o *MicrosoftGraphDriveItemUploadableProperties) SetFileSystemInfo(v AnyOfmicrosoftGraphFileSystemInfo) {
	o.FileSystemInfo = &v
}

// SetFileSystemInfoExplicitNull (un)sets FileSystemInfo to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The FileSystemInfo value is set to nil even if false is passed
func (o *MicrosoftGraphDriveItemUploadableProperties) SetFileSystemInfoExplicitNull(b bool) {
	o.FileSystemInfo = nil
	o.isExplicitNullFileSystemInfo = b
}
// GetName returns the Name field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDriveItemUploadableProperties) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDriveItemUploadableProperties) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphDriveItemUploadableProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftGraphDriveItemUploadableProperties) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *MicrosoftGraphDriveItemUploadableProperties) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphDriveItemUploadableProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description == nil {
		if o.isExplicitNullDescription {
			toSerialize["description"] = o.Description
		}
	} else {
		toSerialize["description"] = o.Description
	}
	if o.FileSystemInfo == nil {
		if o.isExplicitNullFileSystemInfo {
			toSerialize["fileSystemInfo"] = o.FileSystemInfo
		}
	} else {
		toSerialize["fileSystemInfo"] = o.FileSystemInfo
	}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}


