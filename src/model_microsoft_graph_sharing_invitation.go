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
// MicrosoftGraphSharingInvitation struct for MicrosoftGraphSharingInvitation
type MicrosoftGraphSharingInvitation struct {
	Email *string `json:"email,omitempty"`
	isExplicitNullEmail bool `json:"-"`
	InvitedBy *AnyOfmicrosoftGraphIdentitySet `json:"invitedBy,omitempty"`
	isExplicitNullInvitedBy bool `json:"-"`
	RedeemedBy *string `json:"redeemedBy,omitempty"`
	isExplicitNullRedeemedBy bool `json:"-"`
	SignInRequired *bool `json:"signInRequired,omitempty"`
	isExplicitNullSignInRequired bool `json:"-"`
}

// GetEmail returns the Email field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingInvitation) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingInvitation) GetEmailOk() (string, bool) {
	if o == nil || o.Email == nil {
		var ret string
		return ret, false
	}
	return *o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingInvitation) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MicrosoftGraphSharingInvitation) SetEmail(v string) {
	o.Email = &v
}

// SetEmailExplicitNull (un)sets Email to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Email value is set to nil even if false is passed
func (o *MicrosoftGraphSharingInvitation) SetEmailExplicitNull(b bool) {
	o.Email = nil
	o.isExplicitNullEmail = b
}
// GetInvitedBy returns the InvitedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingInvitation) GetInvitedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.InvitedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.InvitedBy
}

// GetInvitedByOk returns a tuple with the InvitedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingInvitation) GetInvitedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.InvitedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.InvitedBy, true
}

// HasInvitedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingInvitation) HasInvitedBy() bool {
	if o != nil && o.InvitedBy != nil {
		return true
	}

	return false
}

// SetInvitedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the InvitedBy field.
func (o *MicrosoftGraphSharingInvitation) SetInvitedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.InvitedBy = &v
}

// SetInvitedByExplicitNull (un)sets InvitedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The InvitedBy value is set to nil even if false is passed
func (o *MicrosoftGraphSharingInvitation) SetInvitedByExplicitNull(b bool) {
	o.InvitedBy = nil
	o.isExplicitNullInvitedBy = b
}
// GetRedeemedBy returns the RedeemedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingInvitation) GetRedeemedBy() string {
	if o == nil || o.RedeemedBy == nil {
		var ret string
		return ret
	}
	return *o.RedeemedBy
}

// GetRedeemedByOk returns a tuple with the RedeemedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingInvitation) GetRedeemedByOk() (string, bool) {
	if o == nil || o.RedeemedBy == nil {
		var ret string
		return ret, false
	}
	return *o.RedeemedBy, true
}

// HasRedeemedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingInvitation) HasRedeemedBy() bool {
	if o != nil && o.RedeemedBy != nil {
		return true
	}

	return false
}

// SetRedeemedBy gets a reference to the given string and assigns it to the RedeemedBy field.
func (o *MicrosoftGraphSharingInvitation) SetRedeemedBy(v string) {
	o.RedeemedBy = &v
}

// SetRedeemedByExplicitNull (un)sets RedeemedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The RedeemedBy value is set to nil even if false is passed
func (o *MicrosoftGraphSharingInvitation) SetRedeemedByExplicitNull(b bool) {
	o.RedeemedBy = nil
	o.isExplicitNullRedeemedBy = b
}
// GetSignInRequired returns the SignInRequired field if non-nil, zero value otherwise.
func (o *MicrosoftGraphSharingInvitation) GetSignInRequired() bool {
	if o == nil || o.SignInRequired == nil {
		var ret bool
		return ret
	}
	return *o.SignInRequired
}

// GetSignInRequiredOk returns a tuple with the SignInRequired field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphSharingInvitation) GetSignInRequiredOk() (bool, bool) {
	if o == nil || o.SignInRequired == nil {
		var ret bool
		return ret, false
	}
	return *o.SignInRequired, true
}

// HasSignInRequired returns a boolean if a field has been set.
func (o *MicrosoftGraphSharingInvitation) HasSignInRequired() bool {
	if o != nil && o.SignInRequired != nil {
		return true
	}

	return false
}

// SetSignInRequired gets a reference to the given bool and assigns it to the SignInRequired field.
func (o *MicrosoftGraphSharingInvitation) SetSignInRequired(v bool) {
	o.SignInRequired = &v
}

// SetSignInRequiredExplicitNull (un)sets SignInRequired to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SignInRequired value is set to nil even if false is passed
func (o *MicrosoftGraphSharingInvitation) SetSignInRequiredExplicitNull(b bool) {
	o.SignInRequired = nil
	o.isExplicitNullSignInRequired = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphSharingInvitation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email == nil {
		if o.isExplicitNullEmail {
			toSerialize["email"] = o.Email
		}
	} else {
		toSerialize["email"] = o.Email
	}
	if o.InvitedBy == nil {
		if o.isExplicitNullInvitedBy {
			toSerialize["invitedBy"] = o.InvitedBy
		}
	} else {
		toSerialize["invitedBy"] = o.InvitedBy
	}
	if o.RedeemedBy == nil {
		if o.isExplicitNullRedeemedBy {
			toSerialize["redeemedBy"] = o.RedeemedBy
		}
	} else {
		toSerialize["redeemedBy"] = o.RedeemedBy
	}
	if o.SignInRequired == nil {
		if o.isExplicitNullSignInRequired {
			toSerialize["signInRequired"] = o.SignInRequired
		}
	} else {
		toSerialize["signInRequired"] = o.SignInRequired
	}
	return json.Marshal(toSerialize)
}

