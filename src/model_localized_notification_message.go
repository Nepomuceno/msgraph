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
// LocalizedNotificationMessage The text content of a Notification Message Template for the specified locale.
type LocalizedNotificationMessage struct {
	// DateTime the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`

	// The Locale for which this message is destined.
	Locale *string `json:"locale,omitempty"`

	// The Message Template Subject.
	Subject *string `json:"subject,omitempty"`

	// The Message Template content.
	MessageTemplate *string `json:"messageTemplate,omitempty"`

	// Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message.
	IsDefault *bool `json:"isDefault,omitempty"`

}

// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *LocalizedNotificationMessage) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LocalizedNotificationMessage) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *LocalizedNotificationMessage) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *LocalizedNotificationMessage) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetLocale returns the Locale field if non-nil, zero value otherwise.
func (o *LocalizedNotificationMessage) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LocalizedNotificationMessage) GetLocaleOk() (string, bool) {
	if o == nil || o.Locale == nil {
		var ret string
		return ret, false
	}
	return *o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *LocalizedNotificationMessage) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *LocalizedNotificationMessage) SetLocale(v string) {
	o.Locale = &v
}

// GetSubject returns the Subject field if non-nil, zero value otherwise.
func (o *LocalizedNotificationMessage) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LocalizedNotificationMessage) GetSubjectOk() (string, bool) {
	if o == nil || o.Subject == nil {
		var ret string
		return ret, false
	}
	return *o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *LocalizedNotificationMessage) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *LocalizedNotificationMessage) SetSubject(v string) {
	o.Subject = &v
}

// GetMessageTemplate returns the MessageTemplate field if non-nil, zero value otherwise.
func (o *LocalizedNotificationMessage) GetMessageTemplate() string {
	if o == nil || o.MessageTemplate == nil {
		var ret string
		return ret
	}
	return *o.MessageTemplate
}

// GetMessageTemplateOk returns a tuple with the MessageTemplate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LocalizedNotificationMessage) GetMessageTemplateOk() (string, bool) {
	if o == nil || o.MessageTemplate == nil {
		var ret string
		return ret, false
	}
	return *o.MessageTemplate, true
}

// HasMessageTemplate returns a boolean if a field has been set.
func (o *LocalizedNotificationMessage) HasMessageTemplate() bool {
	if o != nil && o.MessageTemplate != nil {
		return true
	}

	return false
}

// SetMessageTemplate gets a reference to the given string and assigns it to the MessageTemplate field.
func (o *LocalizedNotificationMessage) SetMessageTemplate(v string) {
	o.MessageTemplate = &v
}

// GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.
func (o *LocalizedNotificationMessage) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LocalizedNotificationMessage) GetIsDefaultOk() (bool, bool) {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret, false
	}
	return *o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *LocalizedNotificationMessage) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *LocalizedNotificationMessage) SetIsDefault(v bool) {
	o.IsDefault = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o LocalizedNotificationMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.MessageTemplate != nil {
		toSerialize["messageTemplate"] = o.MessageTemplate
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	return json.Marshal(toSerialize)
}


