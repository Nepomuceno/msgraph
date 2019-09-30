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
	"time"
	"encoding/json"
)
// MicrosoftGraphFileAttachment struct for MicrosoftGraphFileAttachment
type MicrosoftGraphFileAttachment struct {
	Id *string `json:"id,omitempty"`

	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	isExplicitNullLastModifiedDateTime bool `json:"-"`
	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
	ContentType *string `json:"contentType,omitempty"`
	isExplicitNullContentType bool `json:"-"`
	Size *int32 `json:"size,omitempty"`

	IsInline *bool `json:"isInline,omitempty"`

	ContentId *string `json:"contentId,omitempty"`
	isExplicitNullContentId bool `json:"-"`
	ContentLocation *string `json:"contentLocation,omitempty"`
	isExplicitNullContentLocation bool `json:"-"`
	ContentBytes *string `json:"contentBytes,omitempty"`
	isExplicitNullContentBytes bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphFileAttachment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphFileAttachment) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphFileAttachment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphFileAttachment) SetId(v string) {
	o.Id = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphFileAttachment) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphFileAttachment) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphFileAttachment) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphFileAttachment) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphFileAttachment) SetLastModifiedDateTimeExplicitNull(b bool) {
	o.LastModifiedDateTime = nil
	o.isExplicitNullLastModifiedDateTime = b
}
// GetName returns the Name field if non-nil, zero value otherwise.
func (o *MicrosoftGraphFileAttachment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphFileAttachment) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphFileAttachment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftGraphFileAttachment) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *MicrosoftGraphFileAttachment) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetContentType returns the ContentType field if non-nil, zero value otherwise.
func (o *MicrosoftGraphFileAttachment) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphFileAttachment) GetContentTypeOk() (string, bool) {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret, false
	}
	return *o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *MicrosoftGraphFileAttachment) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *MicrosoftGraphFileAttachment) SetContentType(v string) {
	o.ContentType = &v
}

// SetContentTypeExplicitNull (un)sets ContentType to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ContentType value is set to nil even if false is passed
func (o *MicrosoftGraphFileAttachment) SetContentTypeExplicitNull(b bool) {
	o.ContentType = nil
	o.isExplicitNullContentType = b
}
// GetSize returns the Size field if non-nil, zero value otherwise.
func (o *MicrosoftGraphFileAttachment) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphFileAttachment) GetSizeOk() (int32, bool) {
	if o == nil || o.Size == nil {
		var ret int32
		return ret, false
	}
	return *o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MicrosoftGraphFileAttachment) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *MicrosoftGraphFileAttachment) SetSize(v int32) {
	o.Size = &v
}

// GetIsInline returns the IsInline field if non-nil, zero value otherwise.
func (o *MicrosoftGraphFileAttachment) GetIsInline() bool {
	if o == nil || o.IsInline == nil {
		var ret bool
		return ret
	}
	return *o.IsInline
}

// GetIsInlineOk returns a tuple with the IsInline field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphFileAttachment) GetIsInlineOk() (bool, bool) {
	if o == nil || o.IsInline == nil {
		var ret bool
		return ret, false
	}
	return *o.IsInline, true
}

// HasIsInline returns a boolean if a field has been set.
func (o *MicrosoftGraphFileAttachment) HasIsInline() bool {
	if o != nil && o.IsInline != nil {
		return true
	}

	return false
}

// SetIsInline gets a reference to the given bool and assigns it to the IsInline field.
func (o *MicrosoftGraphFileAttachment) SetIsInline(v bool) {
	o.IsInline = &v
}

// GetContentId returns the ContentId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphFileAttachment) GetContentId() string {
	if o == nil || o.ContentId == nil {
		var ret string
		return ret
	}
	return *o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphFileAttachment) GetContentIdOk() (string, bool) {
	if o == nil || o.ContentId == nil {
		var ret string
		return ret, false
	}
	return *o.ContentId, true
}

// HasContentId returns a boolean if a field has been set.
func (o *MicrosoftGraphFileAttachment) HasContentId() bool {
	if o != nil && o.ContentId != nil {
		return true
	}

	return false
}

// SetContentId gets a reference to the given string and assigns it to the ContentId field.
func (o *MicrosoftGraphFileAttachment) SetContentId(v string) {
	o.ContentId = &v
}

// SetContentIdExplicitNull (un)sets ContentId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ContentId value is set to nil even if false is passed
func (o *MicrosoftGraphFileAttachment) SetContentIdExplicitNull(b bool) {
	o.ContentId = nil
	o.isExplicitNullContentId = b
}
// GetContentLocation returns the ContentLocation field if non-nil, zero value otherwise.
func (o *MicrosoftGraphFileAttachment) GetContentLocation() string {
	if o == nil || o.ContentLocation == nil {
		var ret string
		return ret
	}
	return *o.ContentLocation
}

// GetContentLocationOk returns a tuple with the ContentLocation field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphFileAttachment) GetContentLocationOk() (string, bool) {
	if o == nil || o.ContentLocation == nil {
		var ret string
		return ret, false
	}
	return *o.ContentLocation, true
}

// HasContentLocation returns a boolean if a field has been set.
func (o *MicrosoftGraphFileAttachment) HasContentLocation() bool {
	if o != nil && o.ContentLocation != nil {
		return true
	}

	return false
}

// SetContentLocation gets a reference to the given string and assigns it to the ContentLocation field.
func (o *MicrosoftGraphFileAttachment) SetContentLocation(v string) {
	o.ContentLocation = &v
}

// SetContentLocationExplicitNull (un)sets ContentLocation to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ContentLocation value is set to nil even if false is passed
func (o *MicrosoftGraphFileAttachment) SetContentLocationExplicitNull(b bool) {
	o.ContentLocation = nil
	o.isExplicitNullContentLocation = b
}
// GetContentBytes returns the ContentBytes field if non-nil, zero value otherwise.
func (o *MicrosoftGraphFileAttachment) GetContentBytes() string {
	if o == nil || o.ContentBytes == nil {
		var ret string
		return ret
	}
	return *o.ContentBytes
}

// GetContentBytesOk returns a tuple with the ContentBytes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphFileAttachment) GetContentBytesOk() (string, bool) {
	if o == nil || o.ContentBytes == nil {
		var ret string
		return ret, false
	}
	return *o.ContentBytes, true
}

// HasContentBytes returns a boolean if a field has been set.
func (o *MicrosoftGraphFileAttachment) HasContentBytes() bool {
	if o != nil && o.ContentBytes != nil {
		return true
	}

	return false
}

// SetContentBytes gets a reference to the given string and assigns it to the ContentBytes field.
func (o *MicrosoftGraphFileAttachment) SetContentBytes(v string) {
	o.ContentBytes = &v
}

// SetContentBytesExplicitNull (un)sets ContentBytes to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ContentBytes value is set to nil even if false is passed
func (o *MicrosoftGraphFileAttachment) SetContentBytesExplicitNull(b bool) {
	o.ContentBytes = nil
	o.isExplicitNullContentBytes = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphFileAttachment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastModifiedDateTime == nil {
		if o.isExplicitNullLastModifiedDateTime {
			toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
		}
	} else {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
	}
	if o.ContentType == nil {
		if o.isExplicitNullContentType {
			toSerialize["contentType"] = o.ContentType
		}
	} else {
		toSerialize["contentType"] = o.ContentType
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.IsInline != nil {
		toSerialize["isInline"] = o.IsInline
	}
	if o.ContentId == nil {
		if o.isExplicitNullContentId {
			toSerialize["contentId"] = o.ContentId
		}
	} else {
		toSerialize["contentId"] = o.ContentId
	}
	if o.ContentLocation == nil {
		if o.isExplicitNullContentLocation {
			toSerialize["contentLocation"] = o.ContentLocation
		}
	} else {
		toSerialize["contentLocation"] = o.ContentLocation
	}
	if o.ContentBytes == nil {
		if o.isExplicitNullContentBytes {
			toSerialize["contentBytes"] = o.ContentBytes
		}
	} else {
		toSerialize["contentBytes"] = o.ContentBytes
	}
	return json.Marshal(toSerialize)
}


