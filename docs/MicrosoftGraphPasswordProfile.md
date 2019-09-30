# MicrosoftGraphPasswordProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 
**ForceChangePasswordNextSignIn** | Pointer to **bool** |  | [optional] 
**ForceChangePasswordNextSignInWithMfa** | Pointer to **bool** |  | [optional] 

## Methods

### GetPassword

`func (o *MicrosoftGraphPasswordProfile) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MicrosoftGraphPasswordProfile) GetPasswordOk() (string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPassword

`func (o *MicrosoftGraphPasswordProfile) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPassword

`func (o *MicrosoftGraphPasswordProfile) SetPassword(v string)`

SetPassword gets a reference to the given string and assigns it to the Password field.

### SetPasswordExplicitNull

`func (o *MicrosoftGraphPasswordProfile) SetPasswordExplicitNull(b bool)`

SetPasswordExplicitNull (un)sets Password to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Password value is set to nil even if false is passed
### GetForceChangePasswordNextSignIn

`func (o *MicrosoftGraphPasswordProfile) GetForceChangePasswordNextSignIn() bool`

GetForceChangePasswordNextSignIn returns the ForceChangePasswordNextSignIn field if non-nil, zero value otherwise.

### GetForceChangePasswordNextSignInOk

`func (o *MicrosoftGraphPasswordProfile) GetForceChangePasswordNextSignInOk() (bool, bool)`

GetForceChangePasswordNextSignInOk returns a tuple with the ForceChangePasswordNextSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasForceChangePasswordNextSignIn

`func (o *MicrosoftGraphPasswordProfile) HasForceChangePasswordNextSignIn() bool`

HasForceChangePasswordNextSignIn returns a boolean if a field has been set.

### SetForceChangePasswordNextSignIn

`func (o *MicrosoftGraphPasswordProfile) SetForceChangePasswordNextSignIn(v bool)`

SetForceChangePasswordNextSignIn gets a reference to the given bool and assigns it to the ForceChangePasswordNextSignIn field.

### SetForceChangePasswordNextSignInExplicitNull

`func (o *MicrosoftGraphPasswordProfile) SetForceChangePasswordNextSignInExplicitNull(b bool)`

SetForceChangePasswordNextSignInExplicitNull (un)sets ForceChangePasswordNextSignIn to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ForceChangePasswordNextSignIn value is set to nil even if false is passed
### GetForceChangePasswordNextSignInWithMfa

`func (o *MicrosoftGraphPasswordProfile) GetForceChangePasswordNextSignInWithMfa() bool`

GetForceChangePasswordNextSignInWithMfa returns the ForceChangePasswordNextSignInWithMfa field if non-nil, zero value otherwise.

### GetForceChangePasswordNextSignInWithMfaOk

`func (o *MicrosoftGraphPasswordProfile) GetForceChangePasswordNextSignInWithMfaOk() (bool, bool)`

GetForceChangePasswordNextSignInWithMfaOk returns a tuple with the ForceChangePasswordNextSignInWithMfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasForceChangePasswordNextSignInWithMfa

`func (o *MicrosoftGraphPasswordProfile) HasForceChangePasswordNextSignInWithMfa() bool`

HasForceChangePasswordNextSignInWithMfa returns a boolean if a field has been set.

### SetForceChangePasswordNextSignInWithMfa

`func (o *MicrosoftGraphPasswordProfile) SetForceChangePasswordNextSignInWithMfa(v bool)`

SetForceChangePasswordNextSignInWithMfa gets a reference to the given bool and assigns it to the ForceChangePasswordNextSignInWithMfa field.

### SetForceChangePasswordNextSignInWithMfaExplicitNull

`func (o *MicrosoftGraphPasswordProfile) SetForceChangePasswordNextSignInWithMfaExplicitNull(b bool)`

SetForceChangePasswordNextSignInWithMfaExplicitNull (un)sets ForceChangePasswordNextSignInWithMfa to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ForceChangePasswordNextSignInWithMfa value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


