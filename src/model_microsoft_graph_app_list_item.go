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
// MicrosoftGraphAppListItem struct for MicrosoftGraphAppListItem
type MicrosoftGraphAppListItem struct {
	// The application name
	Name *string `json:"name,omitempty"`

	// The publisher of the application
	Publisher *string `json:"publisher,omitempty"`
	isExplicitNullPublisher bool `json:"-"`
	// The Store URL of the application
	AppStoreUrl *string `json:"appStoreUrl,omitempty"`
	isExplicitNullAppStoreUrl bool `json:"-"`
	// The application or bundle identifier of the application
	AppId *string `json:"appId,omitempty"`
	isExplicitNullAppId bool `json:"-"`
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAppListItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppListItem) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphAppListItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftGraphAppListItem) SetName(v string) {
	o.Name = &v
}

// GetPublisher returns the Publisher field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAppListItem) GetPublisher() string {
	if o == nil || o.Publisher == nil {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppListItem) GetPublisherOk() (string, bool) {
	if o == nil || o.Publisher == nil {
		var ret string
		return ret, false
	}
	return *o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *MicrosoftGraphAppListItem) HasPublisher() bool {
	if o != nil && o.Publisher != nil {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *MicrosoftGraphAppListItem) SetPublisher(v string) {
	o.Publisher = &v
}

// SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Publisher value is set to nil even if false is passed
func (o *MicrosoftGraphAppListItem) SetPublisherExplicitNull(b bool) {
	o.Publisher = nil
	o.isExplicitNullPublisher = b
}
// GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAppListItem) GetAppStoreUrl() string {
	if o == nil || o.AppStoreUrl == nil {
		var ret string
		return ret
	}
	return *o.AppStoreUrl
}

// GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppListItem) GetAppStoreUrlOk() (string, bool) {
	if o == nil || o.AppStoreUrl == nil {
		var ret string
		return ret, false
	}
	return *o.AppStoreUrl, true
}

// HasAppStoreUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphAppListItem) HasAppStoreUrl() bool {
	if o != nil && o.AppStoreUrl != nil {
		return true
	}

	return false
}

// SetAppStoreUrl gets a reference to the given string and assigns it to the AppStoreUrl field.
func (o *MicrosoftGraphAppListItem) SetAppStoreUrl(v string) {
	o.AppStoreUrl = &v
}

// SetAppStoreUrlExplicitNull (un)sets AppStoreUrl to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AppStoreUrl value is set to nil even if false is passed
func (o *MicrosoftGraphAppListItem) SetAppStoreUrlExplicitNull(b bool) {
	o.AppStoreUrl = nil
	o.isExplicitNullAppStoreUrl = b
}
// GetAppId returns the AppId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAppListItem) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppListItem) GetAppIdOk() (string, bool) {
	if o == nil || o.AppId == nil {
		var ret string
		return ret, false
	}
	return *o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *MicrosoftGraphAppListItem) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *MicrosoftGraphAppListItem) SetAppId(v string) {
	o.AppId = &v
}

// SetAppIdExplicitNull (un)sets AppId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AppId value is set to nil even if false is passed
func (o *MicrosoftGraphAppListItem) SetAppIdExplicitNull(b bool) {
	o.AppId = nil
	o.isExplicitNullAppId = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphAppListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Publisher == nil {
		if o.isExplicitNullPublisher {
			toSerialize["publisher"] = o.Publisher
		}
	} else {
		toSerialize["publisher"] = o.Publisher
	}
	if o.AppStoreUrl == nil {
		if o.isExplicitNullAppStoreUrl {
			toSerialize["appStoreUrl"] = o.AppStoreUrl
		}
	} else {
		toSerialize["appStoreUrl"] = o.AppStoreUrl
	}
	if o.AppId == nil {
		if o.isExplicitNullAppId {
			toSerialize["appId"] = o.AppId
		}
	} else {
		toSerialize["appId"] = o.AppId
	}
	return json.Marshal(toSerialize)
}

