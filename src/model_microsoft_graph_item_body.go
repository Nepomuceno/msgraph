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
// MicrosoftGraphItemBody struct for MicrosoftGraphItemBody
type MicrosoftGraphItemBody struct {
	ContentType *AnyOfmicrosoftGraphBodyType `json:"contentType,omitempty"`
	isExplicitNullContentType bool `json:"-"`
	Content *string `json:"content,omitempty"`
	isExplicitNullContent bool `json:"-"`
}

// GetContentType returns the ContentType field if non-nil, zero value otherwise.
func (o *MicrosoftGraphItemBody) GetContentType() AnyOfmicrosoftGraphBodyType {
	if o == nil || o.ContentType == nil {
		var ret AnyOfmicrosoftGraphBodyType
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphItemBody) GetContentTypeOk() (AnyOfmicrosoftGraphBodyType, bool) {
	if o == nil || o.ContentType == nil {
		var ret AnyOfmicrosoftGraphBodyType
		return ret, false
	}
	return *o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *MicrosoftGraphItemBody) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given AnyOfmicrosoftGraphBodyType and assigns it to the ContentType field.
func (o *MicrosoftGraphItemBody) SetContentType(v AnyOfmicrosoftGraphBodyType) {
	o.ContentType = &v
}

// SetContentTypeExplicitNull (un)sets ContentType to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ContentType value is set to nil even if false is passed
func (o *MicrosoftGraphItemBody) SetContentTypeExplicitNull(b bool) {
	o.ContentType = nil
	o.isExplicitNullContentType = b
}
// GetContent returns the Content field if non-nil, zero value otherwise.
func (o *MicrosoftGraphItemBody) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphItemBody) GetContentOk() (string, bool) {
	if o == nil || o.Content == nil {
		var ret string
		return ret, false
	}
	return *o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *MicrosoftGraphItemBody) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *MicrosoftGraphItemBody) SetContent(v string) {
	o.Content = &v
}

// SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Content value is set to nil even if false is passed
func (o *MicrosoftGraphItemBody) SetContentExplicitNull(b bool) {
	o.Content = nil
	o.isExplicitNullContent = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphItemBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentType == nil {
		if o.isExplicitNullContentType {
			toSerialize["contentType"] = o.ContentType
		}
	} else {
		toSerialize["contentType"] = o.ContentType
	}
	if o.Content == nil {
		if o.isExplicitNullContent {
			toSerialize["content"] = o.Content
		}
	} else {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}


