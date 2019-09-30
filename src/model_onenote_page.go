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
// OnenotePage struct for OnenotePage
type OnenotePage struct {
	Title *string `json:"title,omitempty"`
	isExplicitNullTitle bool `json:"-"`
	CreatedByAppId *string `json:"createdByAppId,omitempty"`
	isExplicitNullCreatedByAppId bool `json:"-"`
	Links *AnyOfmicrosoftGraphPageLinks `json:"links,omitempty"`
	isExplicitNullLinks bool `json:"-"`
	ContentUrl *string `json:"contentUrl,omitempty"`
	isExplicitNullContentUrl bool `json:"-"`
	Content *string `json:"content,omitempty"`
	isExplicitNullContent bool `json:"-"`
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	isExplicitNullLastModifiedDateTime bool `json:"-"`
	Level *int32 `json:"level,omitempty"`
	isExplicitNullLevel bool `json:"-"`
	Order *int32 `json:"order,omitempty"`
	isExplicitNullOrder bool `json:"-"`
	UserTags *[]string `json:"userTags,omitempty"`

	ParentSection *AnyOfmicrosoftGraphOnenoteSection `json:"parentSection,omitempty"`
	isExplicitNullParentSection bool `json:"-"`
	ParentNotebook *AnyOfmicrosoftGraphNotebook `json:"parentNotebook,omitempty"`
	isExplicitNullParentNotebook bool `json:"-"`
}

// GetTitle returns the Title field if non-nil, zero value otherwise.
func (o *OnenotePage) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OnenotePage) GetTitleOk() (string, bool) {
	if o == nil || o.Title == nil {
		var ret string
		return ret, false
	}
	return *o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *OnenotePage) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *OnenotePage) SetTitle(v string) {
	o.Title = &v
}

// SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Title value is set to nil even if false is passed
func (o *OnenotePage) SetTitleExplicitNull(b bool) {
	o.Title = nil
	o.isExplicitNullTitle = b
}
// GetCreatedByAppId returns the CreatedByAppId field if non-nil, zero value otherwise.
func (o *OnenotePage) GetCreatedByAppId() string {
	if o == nil || o.CreatedByAppId == nil {
		var ret string
		return ret
	}
	return *o.CreatedByAppId
}

// GetCreatedByAppIdOk returns a tuple with the CreatedByAppId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OnenotePage) GetCreatedByAppIdOk() (string, bool) {
	if o == nil || o.CreatedByAppId == nil {
		var ret string
		return ret, false
	}
	return *o.CreatedByAppId, true
}

// HasCreatedByAppId returns a boolean if a field has been set.
func (o *OnenotePage) HasCreatedByAppId() bool {
	if o != nil && o.CreatedByAppId != nil {
		return true
	}

	return false
}

// SetCreatedByAppId gets a reference to the given string and assigns it to the CreatedByAppId field.
func (o *OnenotePage) SetCreatedByAppId(v string) {
	o.CreatedByAppId = &v
}

// SetCreatedByAppIdExplicitNull (un)sets CreatedByAppId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedByAppId value is set to nil even if false is passed
func (o *OnenotePage) SetCreatedByAppIdExplicitNull(b bool) {
	o.CreatedByAppId = nil
	o.isExplicitNullCreatedByAppId = b
}
// GetLinks returns the Links field if non-nil, zero value otherwise.
func (o *OnenotePage) GetLinks() AnyOfmicrosoftGraphPageLinks {
	if o == nil || o.Links == nil {
		var ret AnyOfmicrosoftGraphPageLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OnenotePage) GetLinksOk() (AnyOfmicrosoftGraphPageLinks, bool) {
	if o == nil || o.Links == nil {
		var ret AnyOfmicrosoftGraphPageLinks
		return ret, false
	}
	return *o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OnenotePage) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AnyOfmicrosoftGraphPageLinks and assigns it to the Links field.
func (o *OnenotePage) SetLinks(v AnyOfmicrosoftGraphPageLinks) {
	o.Links = &v
}

// SetLinksExplicitNull (un)sets Links to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Links value is set to nil even if false is passed
func (o *OnenotePage) SetLinksExplicitNull(b bool) {
	o.Links = nil
	o.isExplicitNullLinks = b
}
// GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.
func (o *OnenotePage) GetContentUrl() string {
	if o == nil || o.ContentUrl == nil {
		var ret string
		return ret
	}
	return *o.ContentUrl
}

// GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OnenotePage) GetContentUrlOk() (string, bool) {
	if o == nil || o.ContentUrl == nil {
		var ret string
		return ret, false
	}
	return *o.ContentUrl, true
}

// HasContentUrl returns a boolean if a field has been set.
func (o *OnenotePage) HasContentUrl() bool {
	if o != nil && o.ContentUrl != nil {
		return true
	}

	return false
}

// SetContentUrl gets a reference to the given string and assigns it to the ContentUrl field.
func (o *OnenotePage) SetContentUrl(v string) {
	o.ContentUrl = &v
}

// SetContentUrlExplicitNull (un)sets ContentUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ContentUrl value is set to nil even if false is passed
func (o *OnenotePage) SetContentUrlExplicitNull(b bool) {
	o.ContentUrl = nil
	o.isExplicitNullContentUrl = b
}
// GetContent returns the Content field if non-nil, zero value otherwise.
func (o *OnenotePage) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OnenotePage) GetContentOk() (string, bool) {
	if o == nil || o.Content == nil {
		var ret string
		return ret, false
	}
	return *o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *OnenotePage) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *OnenotePage) SetContent(v string) {
	o.Content = &v
}

// SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Content value is set to nil even if false is passed
func (o *OnenotePage) SetContentExplicitNull(b bool) {
	o.Content = nil
	o.isExplicitNullContent = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *OnenotePage) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OnenotePage) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *OnenotePage) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *OnenotePage) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedDateTime value is set to nil even if false is passed
func (o *OnenotePage) SetLastModifiedDateTimeExplicitNull(b bool) {
	o.LastModifiedDateTime = nil
	o.isExplicitNullLastModifiedDateTime = b
}
// GetLevel returns the Level field if non-nil, zero value otherwise.
func (o *OnenotePage) GetLevel() int32 {
	if o == nil || o.Level == nil {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OnenotePage) GetLevelOk() (int32, bool) {
	if o == nil || o.Level == nil {
		var ret int32
		return ret, false
	}
	return *o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *OnenotePage) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *OnenotePage) SetLevel(v int32) {
	o.Level = &v
}

// SetLevelExplicitNull (un)sets Level to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Level value is set to nil even if false is passed
func (o *OnenotePage) SetLevelExplicitNull(b bool) {
	o.Level = nil
	o.isExplicitNullLevel = b
}
// GetOrder returns the Order field if non-nil, zero value otherwise.
func (o *OnenotePage) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OnenotePage) GetOrderOk() (int32, bool) {
	if o == nil || o.Order == nil {
		var ret int32
		return ret, false
	}
	return *o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *OnenotePage) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *OnenotePage) SetOrder(v int32) {
	o.Order = &v
}

// SetOrderExplicitNull (un)sets Order to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Order value is set to nil even if false is passed
func (o *OnenotePage) SetOrderExplicitNull(b bool) {
	o.Order = nil
	o.isExplicitNullOrder = b
}
// GetUserTags returns the UserTags field if non-nil, zero value otherwise.
func (o *OnenotePage) GetUserTags() []string {
	if o == nil || o.UserTags == nil {
		var ret []string
		return ret
	}
	return *o.UserTags
}

// GetUserTagsOk returns a tuple with the UserTags field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OnenotePage) GetUserTagsOk() ([]string, bool) {
	if o == nil || o.UserTags == nil {
		var ret []string
		return ret, false
	}
	return *o.UserTags, true
}

// HasUserTags returns a boolean if a field has been set.
func (o *OnenotePage) HasUserTags() bool {
	if o != nil && o.UserTags != nil {
		return true
	}

	return false
}

// SetUserTags gets a reference to the given []string and assigns it to the UserTags field.
func (o *OnenotePage) SetUserTags(v []string) {
	o.UserTags = &v
}

// GetParentSection returns the ParentSection field if non-nil, zero value otherwise.
func (o *OnenotePage) GetParentSection() AnyOfmicrosoftGraphOnenoteSection {
	if o == nil || o.ParentSection == nil {
		var ret AnyOfmicrosoftGraphOnenoteSection
		return ret
	}
	return *o.ParentSection
}

// GetParentSectionOk returns a tuple with the ParentSection field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OnenotePage) GetParentSectionOk() (AnyOfmicrosoftGraphOnenoteSection, bool) {
	if o == nil || o.ParentSection == nil {
		var ret AnyOfmicrosoftGraphOnenoteSection
		return ret, false
	}
	return *o.ParentSection, true
}

// HasParentSection returns a boolean if a field has been set.
func (o *OnenotePage) HasParentSection() bool {
	if o != nil && o.ParentSection != nil {
		return true
	}

	return false
}

// SetParentSection gets a reference to the given AnyOfmicrosoftGraphOnenoteSection and assigns it to the ParentSection field.
func (o *OnenotePage) SetParentSection(v AnyOfmicrosoftGraphOnenoteSection) {
	o.ParentSection = &v
}

// SetParentSectionExplicitNull (un)sets ParentSection to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ParentSection value is set to nil even if false is passed
func (o *OnenotePage) SetParentSectionExplicitNull(b bool) {
	o.ParentSection = nil
	o.isExplicitNullParentSection = b
}
// GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.
func (o *OnenotePage) GetParentNotebook() AnyOfmicrosoftGraphNotebook {
	if o == nil || o.ParentNotebook == nil {
		var ret AnyOfmicrosoftGraphNotebook
		return ret
	}
	return *o.ParentNotebook
}

// GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OnenotePage) GetParentNotebookOk() (AnyOfmicrosoftGraphNotebook, bool) {
	if o == nil || o.ParentNotebook == nil {
		var ret AnyOfmicrosoftGraphNotebook
		return ret, false
	}
	return *o.ParentNotebook, true
}

// HasParentNotebook returns a boolean if a field has been set.
func (o *OnenotePage) HasParentNotebook() bool {
	if o != nil && o.ParentNotebook != nil {
		return true
	}

	return false
}

// SetParentNotebook gets a reference to the given AnyOfmicrosoftGraphNotebook and assigns it to the ParentNotebook field.
func (o *OnenotePage) SetParentNotebook(v AnyOfmicrosoftGraphNotebook) {
	o.ParentNotebook = &v
}

// SetParentNotebookExplicitNull (un)sets ParentNotebook to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ParentNotebook value is set to nil even if false is passed
func (o *OnenotePage) SetParentNotebookExplicitNull(b bool) {
	o.ParentNotebook = nil
	o.isExplicitNullParentNotebook = b
}

// MarshalJSON returns the JSON representation of the model.
func (o OnenotePage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title == nil {
		if o.isExplicitNullTitle {
			toSerialize["title"] = o.Title
		}
	} else {
		toSerialize["title"] = o.Title
	}
	if o.CreatedByAppId == nil {
		if o.isExplicitNullCreatedByAppId {
			toSerialize["createdByAppId"] = o.CreatedByAppId
		}
	} else {
		toSerialize["createdByAppId"] = o.CreatedByAppId
	}
	if o.Links == nil {
		if o.isExplicitNullLinks {
			toSerialize["links"] = o.Links
		}
	} else {
		toSerialize["links"] = o.Links
	}
	if o.ContentUrl == nil {
		if o.isExplicitNullContentUrl {
			toSerialize["contentUrl"] = o.ContentUrl
		}
	} else {
		toSerialize["contentUrl"] = o.ContentUrl
	}
	if o.Content == nil {
		if o.isExplicitNullContent {
			toSerialize["content"] = o.Content
		}
	} else {
		toSerialize["content"] = o.Content
	}
	if o.LastModifiedDateTime == nil {
		if o.isExplicitNullLastModifiedDateTime {
			toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
		}
	} else {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Level == nil {
		if o.isExplicitNullLevel {
			toSerialize["level"] = o.Level
		}
	} else {
		toSerialize["level"] = o.Level
	}
	if o.Order == nil {
		if o.isExplicitNullOrder {
			toSerialize["order"] = o.Order
		}
	} else {
		toSerialize["order"] = o.Order
	}
	if o.UserTags != nil {
		toSerialize["userTags"] = o.UserTags
	}
	if o.ParentSection == nil {
		if o.isExplicitNullParentSection {
			toSerialize["parentSection"] = o.ParentSection
		}
	} else {
		toSerialize["parentSection"] = o.ParentSection
	}
	if o.ParentNotebook == nil {
		if o.isExplicitNullParentNotebook {
			toSerialize["parentNotebook"] = o.ParentNotebook
		}
	} else {
		toSerialize["parentNotebook"] = o.ParentNotebook
	}
	return json.Marshal(toSerialize)
}

