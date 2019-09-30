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
// MicrosoftGraphPasswordProfile struct for MicrosoftGraphPasswordProfile
type MicrosoftGraphPasswordProfile struct {
	Password *string `json:"password,omitempty"`
	isExplicitNullPassword bool `json:"-"`
	ForceChangePasswordNextSignIn *bool `json:"forceChangePasswordNextSignIn,omitempty"`
	isExplicitNullForceChangePasswordNextSignIn bool `json:"-"`
	ForceChangePasswordNextSignInWithMfa *bool `json:"forceChangePasswordNextSignInWithMfa,omitempty"`
	isExplicitNullForceChangePasswordNextSignInWithMfa bool `json:"-"`
}

// GetPassword returns the Password field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPasswordProfile) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPasswordProfile) GetPasswordOk() (string, bool) {
	if o == nil || o.Password == nil {
		var ret string
		return ret, false
	}
	return *o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *MicrosoftGraphPasswordProfile) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *MicrosoftGraphPasswordProfile) SetPassword(v string) {
	o.Password = &v
}

// SetPasswordExplicitNull (un)sets Password to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Password value is set to nil even if false is passed
func (o *MicrosoftGraphPasswordProfile) SetPasswordExplicitNull(b bool) {
	o.Password = nil
	o.isExplicitNullPassword = b
}
// GetForceChangePasswordNextSignIn returns the ForceChangePasswordNextSignIn field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPasswordProfile) GetForceChangePasswordNextSignIn() bool {
	if o == nil || o.ForceChangePasswordNextSignIn == nil {
		var ret bool
		return ret
	}
	return *o.ForceChangePasswordNextSignIn
}

// GetForceChangePasswordNextSignInOk returns a tuple with the ForceChangePasswordNextSignIn field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPasswordProfile) GetForceChangePasswordNextSignInOk() (bool, bool) {
	if o == nil || o.ForceChangePasswordNextSignIn == nil {
		var ret bool
		return ret, false
	}
	return *o.ForceChangePasswordNextSignIn, true
}

// HasForceChangePasswordNextSignIn returns a boolean if a field has been set.
func (o *MicrosoftGraphPasswordProfile) HasForceChangePasswordNextSignIn() bool {
	if o != nil && o.ForceChangePasswordNextSignIn != nil {
		return true
	}

	return false
}

// SetForceChangePasswordNextSignIn gets a reference to the given bool and assigns it to the ForceChangePasswordNextSignIn field.
func (o *MicrosoftGraphPasswordProfile) SetForceChangePasswordNextSignIn(v bool) {
	o.ForceChangePasswordNextSignIn = &v
}

// SetForceChangePasswordNextSignInExplicitNull (un)sets ForceChangePasswordNextSignIn to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ForceChangePasswordNextSignIn value is set to nil even if false is passed
func (o *MicrosoftGraphPasswordProfile) SetForceChangePasswordNextSignInExplicitNull(b bool) {
	o.ForceChangePasswordNextSignIn = nil
	o.isExplicitNullForceChangePasswordNextSignIn = b
}
// GetForceChangePasswordNextSignInWithMfa returns the ForceChangePasswordNextSignInWithMfa field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPasswordProfile) GetForceChangePasswordNextSignInWithMfa() bool {
	if o == nil || o.ForceChangePasswordNextSignInWithMfa == nil {
		var ret bool
		return ret
	}
	return *o.ForceChangePasswordNextSignInWithMfa
}

// GetForceChangePasswordNextSignInWithMfaOk returns a tuple with the ForceChangePasswordNextSignInWithMfa field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPasswordProfile) GetForceChangePasswordNextSignInWithMfaOk() (bool, bool) {
	if o == nil || o.ForceChangePasswordNextSignInWithMfa == nil {
		var ret bool
		return ret, false
	}
	return *o.ForceChangePasswordNextSignInWithMfa, true
}

// HasForceChangePasswordNextSignInWithMfa returns a boolean if a field has been set.
func (o *MicrosoftGraphPasswordProfile) HasForceChangePasswordNextSignInWithMfa() bool {
	if o != nil && o.ForceChangePasswordNextSignInWithMfa != nil {
		return true
	}

	return false
}

// SetForceChangePasswordNextSignInWithMfa gets a reference to the given bool and assigns it to the ForceChangePasswordNextSignInWithMfa field.
func (o *MicrosoftGraphPasswordProfile) SetForceChangePasswordNextSignInWithMfa(v bool) {
	o.ForceChangePasswordNextSignInWithMfa = &v
}

// SetForceChangePasswordNextSignInWithMfaExplicitNull (un)sets ForceChangePasswordNextSignInWithMfa to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ForceChangePasswordNextSignInWithMfa value is set to nil even if false is passed
func (o *MicrosoftGraphPasswordProfile) SetForceChangePasswordNextSignInWithMfaExplicitNull(b bool) {
	o.ForceChangePasswordNextSignInWithMfa = nil
	o.isExplicitNullForceChangePasswordNextSignInWithMfa = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphPasswordProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password == nil {
		if o.isExplicitNullPassword {
			toSerialize["password"] = o.Password
		}
	} else {
		toSerialize["password"] = o.Password
	}
	if o.ForceChangePasswordNextSignIn == nil {
		if o.isExplicitNullForceChangePasswordNextSignIn {
			toSerialize["forceChangePasswordNextSignIn"] = o.ForceChangePasswordNextSignIn
		}
	} else {
		toSerialize["forceChangePasswordNextSignIn"] = o.ForceChangePasswordNextSignIn
	}
	if o.ForceChangePasswordNextSignInWithMfa == nil {
		if o.isExplicitNullForceChangePasswordNextSignInWithMfa {
			toSerialize["forceChangePasswordNextSignInWithMfa"] = o.ForceChangePasswordNextSignInWithMfa
		}
	} else {
		toSerialize["forceChangePasswordNextSignInWithMfa"] = o.ForceChangePasswordNextSignInWithMfa
	}
	return json.Marshal(toSerialize)
}

