# MicrosoftGraphShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**SharedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**SharedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### GetOwner

`func (o *MicrosoftGraphShared) GetOwner() AnyOfmicrosoftGraphIdentitySet`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MicrosoftGraphShared) GetOwnerOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOwner

`func (o *MicrosoftGraphShared) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwner

`func (o *MicrosoftGraphShared) SetOwner(v AnyOfmicrosoftGraphIdentitySet)`

SetOwner gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the Owner field.

### SetOwnerExplicitNull

`func (o *MicrosoftGraphShared) SetOwnerExplicitNull(b bool)`

SetOwnerExplicitNull (un)sets Owner to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Owner value is set to nil even if false is passed
### GetScope

`func (o *MicrosoftGraphShared) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MicrosoftGraphShared) GetScopeOk() (string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScope

`func (o *MicrosoftGraphShared) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScope

`func (o *MicrosoftGraphShared) SetScope(v string)`

SetScope gets a reference to the given string and assigns it to the Scope field.

### SetScopeExplicitNull

`func (o *MicrosoftGraphShared) SetScopeExplicitNull(b bool)`

SetScopeExplicitNull (un)sets Scope to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Scope value is set to nil even if false is passed
### GetSharedBy

`func (o *MicrosoftGraphShared) GetSharedBy() AnyOfmicrosoftGraphIdentitySet`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *MicrosoftGraphShared) GetSharedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharedBy

`func (o *MicrosoftGraphShared) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### SetSharedBy

`func (o *MicrosoftGraphShared) SetSharedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetSharedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the SharedBy field.

### SetSharedByExplicitNull

`func (o *MicrosoftGraphShared) SetSharedByExplicitNull(b bool)`

SetSharedByExplicitNull (un)sets SharedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharedBy value is set to nil even if false is passed
### GetSharedDateTime

`func (o *MicrosoftGraphShared) GetSharedDateTime() time.Time`

GetSharedDateTime returns the SharedDateTime field if non-nil, zero value otherwise.

### GetSharedDateTimeOk

`func (o *MicrosoftGraphShared) GetSharedDateTimeOk() (time.Time, bool)`

GetSharedDateTimeOk returns a tuple with the SharedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharedDateTime

`func (o *MicrosoftGraphShared) HasSharedDateTime() bool`

HasSharedDateTime returns a boolean if a field has been set.

### SetSharedDateTime

`func (o *MicrosoftGraphShared) SetSharedDateTime(v time.Time)`

SetSharedDateTime gets a reference to the given time.Time and assigns it to the SharedDateTime field.

### SetSharedDateTimeExplicitNull

`func (o *MicrosoftGraphShared) SetSharedDateTimeExplicitNull(b bool)`

SetSharedDateTimeExplicitNull (un)sets SharedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharedDateTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


