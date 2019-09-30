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
// MicrosoftGraphDeviceConfigurationSettingState struct for MicrosoftGraphDeviceConfigurationSettingState
type MicrosoftGraphDeviceConfigurationSettingState struct {
	// The setting that is being reported
	Setting *string `json:"setting,omitempty"`
	isExplicitNullSetting bool `json:"-"`
	// Localized/user friendly setting name that is being reported
	SettingName *string `json:"settingName,omitempty"`
	isExplicitNullSettingName bool `json:"-"`
	// Name of setting instance that is being reported.
	InstanceDisplayName *string `json:"instanceDisplayName,omitempty"`
	isExplicitNullInstanceDisplayName bool `json:"-"`
	// The compliance state of the setting
	State *AnyOfmicrosoftGraphComplianceStatus `json:"state,omitempty"`

	// Error code for the setting
	ErrorCode *int64 `json:"errorCode,omitempty"`

	// Error description
	ErrorDescription *string `json:"errorDescription,omitempty"`
	isExplicitNullErrorDescription bool `json:"-"`
	// UserId
	UserId *string `json:"userId,omitempty"`
	isExplicitNullUserId bool `json:"-"`
	// UserName
	UserName *string `json:"userName,omitempty"`
	isExplicitNullUserName bool `json:"-"`
	// UserEmail
	UserEmail *string `json:"userEmail,omitempty"`
	isExplicitNullUserEmail bool `json:"-"`
	// UserPrincipalName.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	isExplicitNullUserPrincipalName bool `json:"-"`
	// Contributing policies
	Sources *[]AnyOfmicrosoftGraphSettingSource `json:"sources,omitempty"`

	// Current value of setting on device
	CurrentValue *string `json:"currentValue,omitempty"`
	isExplicitNullCurrentValue bool `json:"-"`
}

// GetSetting returns the Setting field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSetting() string {
	if o == nil || o.Setting == nil {
		var ret string
		return ret
	}
	return *o.Setting
}

// GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSettingOk() (string, bool) {
	if o == nil || o.Setting == nil {
		var ret string
		return ret, false
	}
	return *o.Setting, true
}

// HasSetting returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasSetting() bool {
	if o != nil && o.Setting != nil {
		return true
	}

	return false
}

// SetSetting gets a reference to the given string and assigns it to the Setting field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSetting(v string) {
	o.Setting = &v
}

// SetSettingExplicitNull (un)sets Setting to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Setting value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSettingExplicitNull(b bool) {
	o.Setting = nil
	o.isExplicitNullSetting = b
}
// GetSettingName returns the SettingName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSettingName() string {
	if o == nil || o.SettingName == nil {
		var ret string
		return ret
	}
	return *o.SettingName
}

// GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSettingNameOk() (string, bool) {
	if o == nil || o.SettingName == nil {
		var ret string
		return ret, false
	}
	return *o.SettingName, true
}

// HasSettingName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasSettingName() bool {
	if o != nil && o.SettingName != nil {
		return true
	}

	return false
}

// SetSettingName gets a reference to the given string and assigns it to the SettingName field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSettingName(v string) {
	o.SettingName = &v
}

// SetSettingNameExplicitNull (un)sets SettingName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SettingName value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSettingNameExplicitNull(b bool) {
	o.SettingName = nil
	o.isExplicitNullSettingName = b
}
// GetInstanceDisplayName returns the InstanceDisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetInstanceDisplayName() string {
	if o == nil || o.InstanceDisplayName == nil {
		var ret string
		return ret
	}
	return *o.InstanceDisplayName
}

// GetInstanceDisplayNameOk returns a tuple with the InstanceDisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetInstanceDisplayNameOk() (string, bool) {
	if o == nil || o.InstanceDisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.InstanceDisplayName, true
}

// HasInstanceDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasInstanceDisplayName() bool {
	if o != nil && o.InstanceDisplayName != nil {
		return true
	}

	return false
}

// SetInstanceDisplayName gets a reference to the given string and assigns it to the InstanceDisplayName field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetInstanceDisplayName(v string) {
	o.InstanceDisplayName = &v
}

// SetInstanceDisplayNameExplicitNull (un)sets InstanceDisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The InstanceDisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetInstanceDisplayNameExplicitNull(b bool) {
	o.InstanceDisplayName = nil
	o.isExplicitNullInstanceDisplayName = b
}
// GetState returns the State field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetState() AnyOfmicrosoftGraphComplianceStatus {
	if o == nil || o.State == nil {
		var ret AnyOfmicrosoftGraphComplianceStatus
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetStateOk() (AnyOfmicrosoftGraphComplianceStatus, bool) {
	if o == nil || o.State == nil {
		var ret AnyOfmicrosoftGraphComplianceStatus
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the State field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetState(v AnyOfmicrosoftGraphComplianceStatus) {
	o.State = &v
}

// GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetErrorCode() int64 {
	if o == nil || o.ErrorCode == nil {
		var ret int64
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetErrorCodeOk() (int64, bool) {
	if o == nil || o.ErrorCode == nil {
		var ret int64
		return ret, false
	}
	return *o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int64 and assigns it to the ErrorCode field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetErrorCode(v int64) {
	o.ErrorCode = &v
}

// GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetErrorDescriptionOk() (string, bool) {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret, false
	}
	return *o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// SetErrorDescriptionExplicitNull (un)sets ErrorDescription to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ErrorDescription value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetErrorDescriptionExplicitNull(b bool) {
	o.ErrorDescription = nil
	o.isExplicitNullErrorDescription = b
}
// GetUserId returns the UserId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserIdOk() (string, bool) {
	if o == nil || o.UserId == nil {
		var ret string
		return ret, false
	}
	return *o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserId(v string) {
	o.UserId = &v
}

// SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UserId value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserIdExplicitNull(b bool) {
	o.UserId = nil
	o.isExplicitNullUserId = b
}
// GetUserName returns the UserName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserNameOk() (string, bool) {
	if o == nil || o.UserName == nil {
		var ret string
		return ret, false
	}
	return *o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserName(v string) {
	o.UserName = &v
}

// SetUserNameExplicitNull (un)sets UserName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UserName value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserNameExplicitNull(b bool) {
	o.UserName = nil
	o.isExplicitNullUserName = b
}
// GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserEmail() string {
	if o == nil || o.UserEmail == nil {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserEmailOk() (string, bool) {
	if o == nil || o.UserEmail == nil {
		var ret string
		return ret, false
	}
	return *o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasUserEmail() bool {
	if o != nil && o.UserEmail != nil {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserEmail(v string) {
	o.UserEmail = &v
}

// SetUserEmailExplicitNull (un)sets UserEmail to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UserEmail value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserEmailExplicitNull(b bool) {
	o.UserEmail = nil
	o.isExplicitNullUserEmail = b
}
// GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetUserPrincipalNameOk() (string, bool) {
	if o == nil || o.UserPrincipalName == nil {
		var ret string
		return ret, false
	}
	return *o.UserPrincipalName, true
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName != nil {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserPrincipalName(v string) {
	o.UserPrincipalName = &v
}

// SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UserPrincipalName value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetUserPrincipalNameExplicitNull(b bool) {
	o.UserPrincipalName = nil
	o.isExplicitNullUserPrincipalName = b
}
// GetSources returns the Sources field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSources() []AnyOfmicrosoftGraphSettingSource {
	if o == nil || o.Sources == nil {
		var ret []AnyOfmicrosoftGraphSettingSource
		return ret
	}
	return *o.Sources
}

// GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetSourcesOk() ([]AnyOfmicrosoftGraphSettingSource, bool) {
	if o == nil || o.Sources == nil {
		var ret []AnyOfmicrosoftGraphSettingSource
		return ret, false
	}
	return *o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []AnyOfmicrosoftGraphSettingSource and assigns it to the Sources field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetSources(v []AnyOfmicrosoftGraphSettingSource) {
	o.Sources = &v
}

// GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetCurrentValue() string {
	if o == nil || o.CurrentValue == nil {
		var ret string
		return ret
	}
	return *o.CurrentValue
}

// GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) GetCurrentValueOk() (string, bool) {
	if o == nil || o.CurrentValue == nil {
		var ret string
		return ret, false
	}
	return *o.CurrentValue, true
}

// HasCurrentValue returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfigurationSettingState) HasCurrentValue() bool {
	if o != nil && o.CurrentValue != nil {
		return true
	}

	return false
}

// SetCurrentValue gets a reference to the given string and assigns it to the CurrentValue field.
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetCurrentValue(v string) {
	o.CurrentValue = &v
}

// SetCurrentValueExplicitNull (un)sets CurrentValue to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CurrentValue value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceConfigurationSettingState) SetCurrentValueExplicitNull(b bool) {
	o.CurrentValue = nil
	o.isExplicitNullCurrentValue = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphDeviceConfigurationSettingState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Setting == nil {
		if o.isExplicitNullSetting {
			toSerialize["setting"] = o.Setting
		}
	} else {
		toSerialize["setting"] = o.Setting
	}
	if o.SettingName == nil {
		if o.isExplicitNullSettingName {
			toSerialize["settingName"] = o.SettingName
		}
	} else {
		toSerialize["settingName"] = o.SettingName
	}
	if o.InstanceDisplayName == nil {
		if o.isExplicitNullInstanceDisplayName {
			toSerialize["instanceDisplayName"] = o.InstanceDisplayName
		}
	} else {
		toSerialize["instanceDisplayName"] = o.InstanceDisplayName
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorDescription == nil {
		if o.isExplicitNullErrorDescription {
			toSerialize["errorDescription"] = o.ErrorDescription
		}
	} else {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if o.UserId == nil {
		if o.isExplicitNullUserId {
			toSerialize["userId"] = o.UserId
		}
	} else {
		toSerialize["userId"] = o.UserId
	}
	if o.UserName == nil {
		if o.isExplicitNullUserName {
			toSerialize["userName"] = o.UserName
		}
	} else {
		toSerialize["userName"] = o.UserName
	}
	if o.UserEmail == nil {
		if o.isExplicitNullUserEmail {
			toSerialize["userEmail"] = o.UserEmail
		}
	} else {
		toSerialize["userEmail"] = o.UserEmail
	}
	if o.UserPrincipalName == nil {
		if o.isExplicitNullUserPrincipalName {
			toSerialize["userPrincipalName"] = o.UserPrincipalName
		}
	} else {
		toSerialize["userPrincipalName"] = o.UserPrincipalName
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.CurrentValue == nil {
		if o.isExplicitNullCurrentValue {
			toSerialize["currentValue"] = o.CurrentValue
		}
	} else {
		toSerialize["currentValue"] = o.CurrentValue
	}
	return json.Marshal(toSerialize)
}

