# MicrosoftGraphSharingDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedBy** | Pointer to [**AnyOfmicrosoftGraphInsightIdentity**](anyOf&lt;microsoft.graph.insightIdentity&gt;.md) |  | [optional] 
**SharedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**SharingSubject** | Pointer to **string** |  | [optional] 
**SharingType** | Pointer to **string** |  | [optional] 
**SharingReference** | Pointer to [**AnyOfmicrosoftGraphResourceReference**](anyOf&lt;microsoft.graph.resourceReference&gt;.md) |  | [optional] 

## Methods

### GetSharedBy

`func (o *MicrosoftGraphSharingDetail) GetSharedBy() AnyOfmicrosoftGraphInsightIdentity`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *MicrosoftGraphSharingDetail) GetSharedByOk() (AnyOfmicrosoftGraphInsightIdentity, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharedBy

`func (o *MicrosoftGraphSharingDetail) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### SetSharedBy

`func (o *MicrosoftGraphSharingDetail) SetSharedBy(v AnyOfmicrosoftGraphInsightIdentity)`

SetSharedBy gets a reference to the given AnyOfmicrosoftGraphInsightIdentity and assigns it to the SharedBy field.

### SetSharedByExplicitNull

`func (o *MicrosoftGraphSharingDetail) SetSharedByExplicitNull(b bool)`

SetSharedByExplicitNull (un)sets SharedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharedBy value is set to nil even if false is passed
### GetSharedDateTime

`func (o *MicrosoftGraphSharingDetail) GetSharedDateTime() time.Time`

GetSharedDateTime returns the SharedDateTime field if non-nil, zero value otherwise.

### GetSharedDateTimeOk

`func (o *MicrosoftGraphSharingDetail) GetSharedDateTimeOk() (time.Time, bool)`

GetSharedDateTimeOk returns a tuple with the SharedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharedDateTime

`func (o *MicrosoftGraphSharingDetail) HasSharedDateTime() bool`

HasSharedDateTime returns a boolean if a field has been set.

### SetSharedDateTime

`func (o *MicrosoftGraphSharingDetail) SetSharedDateTime(v time.Time)`

SetSharedDateTime gets a reference to the given time.Time and assigns it to the SharedDateTime field.

### SetSharedDateTimeExplicitNull

`func (o *MicrosoftGraphSharingDetail) SetSharedDateTimeExplicitNull(b bool)`

SetSharedDateTimeExplicitNull (un)sets SharedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharedDateTime value is set to nil even if false is passed
### GetSharingSubject

`func (o *MicrosoftGraphSharingDetail) GetSharingSubject() string`

GetSharingSubject returns the SharingSubject field if non-nil, zero value otherwise.

### GetSharingSubjectOk

`func (o *MicrosoftGraphSharingDetail) GetSharingSubjectOk() (string, bool)`

GetSharingSubjectOk returns a tuple with the SharingSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharingSubject

`func (o *MicrosoftGraphSharingDetail) HasSharingSubject() bool`

HasSharingSubject returns a boolean if a field has been set.

### SetSharingSubject

`func (o *MicrosoftGraphSharingDetail) SetSharingSubject(v string)`

SetSharingSubject gets a reference to the given string and assigns it to the SharingSubject field.

### SetSharingSubjectExplicitNull

`func (o *MicrosoftGraphSharingDetail) SetSharingSubjectExplicitNull(b bool)`

SetSharingSubjectExplicitNull (un)sets SharingSubject to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharingSubject value is set to nil even if false is passed
### GetSharingType

`func (o *MicrosoftGraphSharingDetail) GetSharingType() string`

GetSharingType returns the SharingType field if non-nil, zero value otherwise.

### GetSharingTypeOk

`func (o *MicrosoftGraphSharingDetail) GetSharingTypeOk() (string, bool)`

GetSharingTypeOk returns a tuple with the SharingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharingType

`func (o *MicrosoftGraphSharingDetail) HasSharingType() bool`

HasSharingType returns a boolean if a field has been set.

### SetSharingType

`func (o *MicrosoftGraphSharingDetail) SetSharingType(v string)`

SetSharingType gets a reference to the given string and assigns it to the SharingType field.

### SetSharingTypeExplicitNull

`func (o *MicrosoftGraphSharingDetail) SetSharingTypeExplicitNull(b bool)`

SetSharingTypeExplicitNull (un)sets SharingType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharingType value is set to nil even if false is passed
### GetSharingReference

`func (o *MicrosoftGraphSharingDetail) GetSharingReference() AnyOfmicrosoftGraphResourceReference`

GetSharingReference returns the SharingReference field if non-nil, zero value otherwise.

### GetSharingReferenceOk

`func (o *MicrosoftGraphSharingDetail) GetSharingReferenceOk() (AnyOfmicrosoftGraphResourceReference, bool)`

GetSharingReferenceOk returns a tuple with the SharingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSharingReference

`func (o *MicrosoftGraphSharingDetail) HasSharingReference() bool`

HasSharingReference returns a boolean if a field has been set.

### SetSharingReference

`func (o *MicrosoftGraphSharingDetail) SetSharingReference(v AnyOfmicrosoftGraphResourceReference)`

SetSharingReference gets a reference to the given AnyOfmicrosoftGraphResourceReference and assigns it to the SharingReference field.

### SetSharingReferenceExplicitNull

`func (o *MicrosoftGraphSharingDetail) SetSharingReferenceExplicitNull(b bool)`

SetSharingReferenceExplicitNull (un)sets SharingReference to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SharingReference value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


