# MicrosoftGraphAuditActivityInitiator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**AnyOfmicrosoftGraphUserIdentity**](anyOf&lt;microsoft.graph.userIdentity&gt;.md) |  | [optional] 
**App** | Pointer to [**AnyOfmicrosoftGraphAppIdentity**](anyOf&lt;microsoft.graph.appIdentity&gt;.md) |  | [optional] 

## Methods

### GetUser

`func (o *MicrosoftGraphAuditActivityInitiator) GetUser() AnyOfmicrosoftGraphUserIdentity`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MicrosoftGraphAuditActivityInitiator) GetUserOk() (AnyOfmicrosoftGraphUserIdentity, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUser

`func (o *MicrosoftGraphAuditActivityInitiator) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUser

`func (o *MicrosoftGraphAuditActivityInitiator) SetUser(v AnyOfmicrosoftGraphUserIdentity)`

SetUser gets a reference to the given AnyOfmicrosoftGraphUserIdentity and assigns it to the User field.

### SetUserExplicitNull

`func (o *MicrosoftGraphAuditActivityInitiator) SetUserExplicitNull(b bool)`

SetUserExplicitNull (un)sets User to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The User value is set to nil even if false is passed
### GetApp

`func (o *MicrosoftGraphAuditActivityInitiator) GetApp() AnyOfmicrosoftGraphAppIdentity`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *MicrosoftGraphAuditActivityInitiator) GetAppOk() (AnyOfmicrosoftGraphAppIdentity, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApp

`func (o *MicrosoftGraphAuditActivityInitiator) HasApp() bool`

HasApp returns a boolean if a field has been set.

### SetApp

`func (o *MicrosoftGraphAuditActivityInitiator) SetApp(v AnyOfmicrosoftGraphAppIdentity)`

SetApp gets a reference to the given AnyOfmicrosoftGraphAppIdentity and assigns it to the App field.

### SetAppExplicitNull

`func (o *MicrosoftGraphAuditActivityInitiator) SetAppExplicitNull(b bool)`

SetAppExplicitNull (un)sets App to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The App value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


