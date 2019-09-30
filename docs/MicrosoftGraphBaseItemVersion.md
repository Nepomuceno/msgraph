# MicrosoftGraphBaseItemVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Publication** | Pointer to [**AnyOfmicrosoftGraphPublicationFacet**](anyOf&lt;microsoft.graph.publicationFacet&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphBaseItemVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphBaseItemVersion) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphBaseItemVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphBaseItemVersion) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedBy

`func (o *MicrosoftGraphBaseItemVersion) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphBaseItemVersion) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphBaseItemVersion) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphBaseItemVersion) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphBaseItemVersion) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphBaseItemVersion) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphBaseItemVersion) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphBaseItemVersion) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphBaseItemVersion) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphBaseItemVersion) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetPublication

`func (o *MicrosoftGraphBaseItemVersion) GetPublication() AnyOfmicrosoftGraphPublicationFacet`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *MicrosoftGraphBaseItemVersion) GetPublicationOk() (AnyOfmicrosoftGraphPublicationFacet, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublication

`func (o *MicrosoftGraphBaseItemVersion) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublication

`func (o *MicrosoftGraphBaseItemVersion) SetPublication(v AnyOfmicrosoftGraphPublicationFacet)`

SetPublication gets a reference to the given AnyOfmicrosoftGraphPublicationFacet and assigns it to the Publication field.

### SetPublicationExplicitNull

`func (o *MicrosoftGraphBaseItemVersion) SetPublicationExplicitNull(b bool)`

SetPublicationExplicitNull (un)sets Publication to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publication value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


