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
// NotificationMessageTemplate Notification messages are messages that are sent to end users who are determined to be not-compliant with the compliance policies defined by the administrator. Administrators choose notifications and configure them in the Intune Admin Console using the compliance policy creation page under the “Actions for non-compliance” section. Use the notificationMessageTemplate object to create your own custom notifications for administrators to choose while configuring actions for non-compliance.
type NotificationMessageTemplate struct {
	// DateTime the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`

	// Display name for the Notification Message Template.
	DisplayName *string `json:"displayName,omitempty"`

	// The default locale to fallback onto when the requested locale is not available.
	DefaultLocale *string `json:"defaultLocale,omitempty"`
	isExplicitNullDefaultLocale bool `json:"-"`
	// The Message Template Branding Options. Branding is defined in the Intune Admin Console.
	BrandingOptions *AnyOfmicrosoftGraphNotificationTemplateBrandingOptions `json:"brandingOptions,omitempty"`

	LocalizedNotificationMessages *[]MicrosoftGraphLocalizedNotificationMessage `json:"localizedNotificationMessages,omitempty"`

}

// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *NotificationMessageTemplate) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageTemplate) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *NotificationMessageTemplate) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *NotificationMessageTemplate) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *NotificationMessageTemplate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageTemplate) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NotificationMessageTemplate) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NotificationMessageTemplate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDefaultLocale returns the DefaultLocale field if non-nil, zero value otherwise.
func (o *NotificationMessageTemplate) GetDefaultLocale() string {
	if o == nil || o.DefaultLocale == nil {
		var ret string
		return ret
	}
	return *o.DefaultLocale
}

// GetDefaultLocaleOk returns a tuple with the DefaultLocale field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageTemplate) GetDefaultLocaleOk() (string, bool) {
	if o == nil || o.DefaultLocale == nil {
		var ret string
		return ret, false
	}
	return *o.DefaultLocale, true
}

// HasDefaultLocale returns a boolean if a field has been set.
func (o *NotificationMessageTemplate) HasDefaultLocale() bool {
	if o != nil && o.DefaultLocale != nil {
		return true
	}

	return false
}

// SetDefaultLocale gets a reference to the given string and assigns it to the DefaultLocale field.
func (o *NotificationMessageTemplate) SetDefaultLocale(v string) {
	o.DefaultLocale = &v
}

// SetDefaultLocaleExplicitNull (un)sets DefaultLocale to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DefaultLocale value is set to nil even if false is passed
func (o *NotificationMessageTemplate) SetDefaultLocaleExplicitNull(b bool) {
	o.DefaultLocale = nil
	o.isExplicitNullDefaultLocale = b
}
// GetBrandingOptions returns the BrandingOptions field if non-nil, zero value otherwise.
func (o *NotificationMessageTemplate) GetBrandingOptions() AnyOfmicrosoftGraphNotificationTemplateBrandingOptions {
	if o == nil || o.BrandingOptions == nil {
		var ret AnyOfmicrosoftGraphNotificationTemplateBrandingOptions
		return ret
	}
	return *o.BrandingOptions
}

// GetBrandingOptionsOk returns a tuple with the BrandingOptions field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageTemplate) GetBrandingOptionsOk() (AnyOfmicrosoftGraphNotificationTemplateBrandingOptions, bool) {
	if o == nil || o.BrandingOptions == nil {
		var ret AnyOfmicrosoftGraphNotificationTemplateBrandingOptions
		return ret, false
	}
	return *o.BrandingOptions, true
}

// HasBrandingOptions returns a boolean if a field has been set.
func (o *NotificationMessageTemplate) HasBrandingOptions() bool {
	if o != nil && o.BrandingOptions != nil {
		return true
	}

	return false
}

// SetBrandingOptions gets a reference to the given AnyOfmicrosoftGraphNotificationTemplateBrandingOptions and assigns it to the BrandingOptions field.
func (o *NotificationMessageTemplate) SetBrandingOptions(v AnyOfmicrosoftGraphNotificationTemplateBrandingOptions) {
	o.BrandingOptions = &v
}

// GetLocalizedNotificationMessages returns the LocalizedNotificationMessages field if non-nil, zero value otherwise.
func (o *NotificationMessageTemplate) GetLocalizedNotificationMessages() []MicrosoftGraphLocalizedNotificationMessage {
	if o == nil || o.LocalizedNotificationMessages == nil {
		var ret []MicrosoftGraphLocalizedNotificationMessage
		return ret
	}
	return *o.LocalizedNotificationMessages
}

// GetLocalizedNotificationMessagesOk returns a tuple with the LocalizedNotificationMessages field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NotificationMessageTemplate) GetLocalizedNotificationMessagesOk() ([]MicrosoftGraphLocalizedNotificationMessage, bool) {
	if o == nil || o.LocalizedNotificationMessages == nil {
		var ret []MicrosoftGraphLocalizedNotificationMessage
		return ret, false
	}
	return *o.LocalizedNotificationMessages, true
}

// HasLocalizedNotificationMessages returns a boolean if a field has been set.
func (o *NotificationMessageTemplate) HasLocalizedNotificationMessages() bool {
	if o != nil && o.LocalizedNotificationMessages != nil {
		return true
	}

	return false
}

// SetLocalizedNotificationMessages gets a reference to the given []MicrosoftGraphLocalizedNotificationMessage and assigns it to the LocalizedNotificationMessages field.
func (o *NotificationMessageTemplate) SetLocalizedNotificationMessages(v []MicrosoftGraphLocalizedNotificationMessage) {
	o.LocalizedNotificationMessages = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o NotificationMessageTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.DefaultLocale == nil {
		if o.isExplicitNullDefaultLocale {
			toSerialize["defaultLocale"] = o.DefaultLocale
		}
	} else {
		toSerialize["defaultLocale"] = o.DefaultLocale
	}
	if o.BrandingOptions != nil {
		toSerialize["brandingOptions"] = o.BrandingOptions
	}
	if o.LocalizedNotificationMessages != nil {
		toSerialize["localizedNotificationMessages"] = o.LocalizedNotificationMessages
	}
	return json.Marshal(toSerialize)
}


