# ManagedAppStatusRaw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Status report content. | [optional] 

## Methods

### GetContent

`func (o *ManagedAppStatusRaw) GetContent() AnyOfobject`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ManagedAppStatusRaw) GetContentOk() (AnyOfobject, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContent

`func (o *ManagedAppStatusRaw) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContent

`func (o *ManagedAppStatusRaw) SetContent(v AnyOfobject)`

SetContent gets a reference to the given AnyOfobject and assigns it to the Content field.

### SetContentExplicitNull

`func (o *ManagedAppStatusRaw) SetContentExplicitNull(b bool)`

SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Content value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


