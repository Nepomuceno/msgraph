# MicrosoftGraphListItemVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Publication** | Pointer to [**AnyOfmicrosoftGraphPublicationFacet**](anyOf&lt;microsoft.graph.publicationFacet&gt;.md) |  | [optional] 
**Fields** | Pointer to [**AnyOfmicrosoftGraphFieldValueSet**](anyOf&lt;microsoft.graph.fieldValueSet&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphListItemVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphListItemVersion) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphListItemVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphListItemVersion) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedBy

`func (o *MicrosoftGraphListItemVersion) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphListItemVersion) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphListItemVersion) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphListItemVersion) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphListItemVersion) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphListItemVersion) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphListItemVersion) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphListItemVersion) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphListItemVersion) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphListItemVersion) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetPublication

`func (o *MicrosoftGraphListItemVersion) GetPublication() AnyOfmicrosoftGraphPublicationFacet`

GetPublication returns the Publication field if non-nil, zero value otherwise.

### GetPublicationOk

`func (o *MicrosoftGraphListItemVersion) GetPublicationOk() (AnyOfmicrosoftGraphPublicationFacet, bool)`

GetPublicationOk returns a tuple with the Publication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublication

`func (o *MicrosoftGraphListItemVersion) HasPublication() bool`

HasPublication returns a boolean if a field has been set.

### SetPublication

`func (o *MicrosoftGraphListItemVersion) SetPublication(v AnyOfmicrosoftGraphPublicationFacet)`

SetPublication gets a reference to the given AnyOfmicrosoftGraphPublicationFacet and assigns it to the Publication field.

### SetPublicationExplicitNull

`func (o *MicrosoftGraphListItemVersion) SetPublicationExplicitNull(b bool)`

SetPublicationExplicitNull (un)sets Publication to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publication value is set to nil even if false is passed
### GetFields

`func (o *MicrosoftGraphListItemVersion) GetFields() AnyOfmicrosoftGraphFieldValueSet`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MicrosoftGraphListItemVersion) GetFieldsOk() (AnyOfmicrosoftGraphFieldValueSet, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFields

`func (o *MicrosoftGraphListItemVersion) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFields

`func (o *MicrosoftGraphListItemVersion) SetFields(v AnyOfmicrosoftGraphFieldValueSet)`

SetFields gets a reference to the given AnyOfmicrosoftGraphFieldValueSet and assigns it to the Fields field.

### SetFieldsExplicitNull

`func (o *MicrosoftGraphListItemVersion) SetFieldsExplicitNull(b bool)`

SetFieldsExplicitNull (un)sets Fields to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Fields value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


