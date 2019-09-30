# MicrosoftGraphBaseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentReference** | Pointer to [**AnyOfmicrosoftGraphItemReference**](anyOf&lt;microsoft.graph.itemReference&gt;.md) |  | [optional] 
**WebUrl** | Pointer to **string** |  | [optional] 
**CreatedByUser** | Pointer to [**AnyOfmicrosoftGraphUser**](anyOf&lt;microsoft.graph.user&gt;.md) |  | [optional] 
**LastModifiedByUser** | Pointer to [**AnyOfmicrosoftGraphUser**](anyOf&lt;microsoft.graph.user&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphBaseItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphBaseItem) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphBaseItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphBaseItem) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedBy

`func (o *MicrosoftGraphBaseItem) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphBaseItem) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphBaseItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphBaseItem) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphBaseItem) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphBaseItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphBaseItem) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphBaseItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphBaseItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphBaseItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphBaseItem) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphBaseItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphBaseItem) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphBaseItem) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetETag

`func (o *MicrosoftGraphBaseItem) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MicrosoftGraphBaseItem) GetETagOk() (string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasETag

`func (o *MicrosoftGraphBaseItem) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETag

`func (o *MicrosoftGraphBaseItem) SetETag(v string)`

SetETag gets a reference to the given string and assigns it to the ETag field.

### SetETagExplicitNull

`func (o *MicrosoftGraphBaseItem) SetETagExplicitNull(b bool)`

SetETagExplicitNull (un)sets ETag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ETag value is set to nil even if false is passed
### GetLastModifiedBy

`func (o *MicrosoftGraphBaseItem) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphBaseItem) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphBaseItem) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphBaseItem) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphBaseItem) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphBaseItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphBaseItem) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphBaseItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphBaseItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetName

`func (o *MicrosoftGraphBaseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphBaseItem) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphBaseItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphBaseItem) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphBaseItem) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetParentReference

`func (o *MicrosoftGraphBaseItem) GetParentReference() AnyOfmicrosoftGraphItemReference`

GetParentReference returns the ParentReference field if non-nil, zero value otherwise.

### GetParentReferenceOk

`func (o *MicrosoftGraphBaseItem) GetParentReferenceOk() (AnyOfmicrosoftGraphItemReference, bool)`

GetParentReferenceOk returns a tuple with the ParentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentReference

`func (o *MicrosoftGraphBaseItem) HasParentReference() bool`

HasParentReference returns a boolean if a field has been set.

### SetParentReference

`func (o *MicrosoftGraphBaseItem) SetParentReference(v AnyOfmicrosoftGraphItemReference)`

SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.

### SetParentReferenceExplicitNull

`func (o *MicrosoftGraphBaseItem) SetParentReferenceExplicitNull(b bool)`

SetParentReferenceExplicitNull (un)sets ParentReference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentReference value is set to nil even if false is passed
### GetWebUrl

`func (o *MicrosoftGraphBaseItem) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphBaseItem) GetWebUrlOk() (string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebUrl

`func (o *MicrosoftGraphBaseItem) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrl

`func (o *MicrosoftGraphBaseItem) SetWebUrl(v string)`

SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.

### SetWebUrlExplicitNull

`func (o *MicrosoftGraphBaseItem) SetWebUrlExplicitNull(b bool)`

SetWebUrlExplicitNull (un)sets WebUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebUrl value is set to nil even if false is passed
### GetCreatedByUser

`func (o *MicrosoftGraphBaseItem) GetCreatedByUser() AnyOfmicrosoftGraphUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *MicrosoftGraphBaseItem) GetCreatedByUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedByUser

`func (o *MicrosoftGraphBaseItem) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### SetCreatedByUser

`func (o *MicrosoftGraphBaseItem) SetCreatedByUser(v AnyOfmicrosoftGraphUser)`

SetCreatedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the CreatedByUser field.

### SetCreatedByUserExplicitNull

`func (o *MicrosoftGraphBaseItem) SetCreatedByUserExplicitNull(b bool)`

SetCreatedByUserExplicitNull (un)sets CreatedByUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedByUser value is set to nil even if false is passed
### GetLastModifiedByUser

`func (o *MicrosoftGraphBaseItem) GetLastModifiedByUser() AnyOfmicrosoftGraphUser`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *MicrosoftGraphBaseItem) GetLastModifiedByUserOk() (AnyOfmicrosoftGraphUser, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedByUser

`func (o *MicrosoftGraphBaseItem) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.

### SetLastModifiedByUser

`func (o *MicrosoftGraphBaseItem) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser)`

SetLastModifiedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the LastModifiedByUser field.

### SetLastModifiedByUserExplicitNull

`func (o *MicrosoftGraphBaseItem) SetLastModifiedByUserExplicitNull(b bool)`

SetLastModifiedByUserExplicitNull (un)sets LastModifiedByUser to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedByUser value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


