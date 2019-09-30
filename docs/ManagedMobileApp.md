# ManagedMobileApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileAppIdentifier** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The identifier for an app with it&#39;s operating system type. | [optional] 
**Version** | Pointer to **string** | Version of the entity. | [optional] 

## Methods

### GetMobileAppIdentifier

`func (o *ManagedMobileApp) GetMobileAppIdentifier() AnyOfobject`

GetMobileAppIdentifier returns the MobileAppIdentifier field if non-nil, zero value otherwise.

### GetMobileAppIdentifierOk

`func (o *ManagedMobileApp) GetMobileAppIdentifierOk() (AnyOfobject, bool)`

GetMobileAppIdentifierOk returns a tuple with the MobileAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobileAppIdentifier

`func (o *ManagedMobileApp) HasMobileAppIdentifier() bool`

HasMobileAppIdentifier returns a boolean if a field has been set.

### SetMobileAppIdentifier

`func (o *ManagedMobileApp) SetMobileAppIdentifier(v AnyOfobject)`

SetMobileAppIdentifier gets a reference to the given AnyOfobject and assigns it to the MobileAppIdentifier field.

### SetMobileAppIdentifierExplicitNull

`func (o *ManagedMobileApp) SetMobileAppIdentifierExplicitNull(b bool)`

SetMobileAppIdentifierExplicitNull (un)sets MobileAppIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MobileAppIdentifier value is set to nil even if false is passed
### GetVersion

`func (o *ManagedMobileApp) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagedMobileApp) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *ManagedMobileApp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *ManagedMobileApp) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *ManagedMobileApp) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


