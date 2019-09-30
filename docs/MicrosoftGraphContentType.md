# MicrosoftGraphContentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**InheritedFrom** | Pointer to [**AnyOfmicrosoftGraphItemReference**](anyOf&lt;microsoft.graph.itemReference&gt;.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Order** | Pointer to [**AnyOfmicrosoftGraphContentTypeOrder**](anyOf&lt;microsoft.graph.contentTypeOrder&gt;.md) |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Sealed** | Pointer to **bool** |  | [optional] 
**ColumnLinks** | Pointer to [**[]MicrosoftGraphColumnLink**](microsoft.graph.columnLink.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphContentType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphContentType) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphContentType) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphContentType) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDescription

`func (o *MicrosoftGraphContentType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphContentType) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphContentType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphContentType) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphContentType) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetGroup

`func (o *MicrosoftGraphContentType) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *MicrosoftGraphContentType) GetGroupOk() (string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGroup

`func (o *MicrosoftGraphContentType) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroup

`func (o *MicrosoftGraphContentType) SetGroup(v string)`

SetGroup gets a reference to the given string and assigns it to the Group field.

### SetGroupExplicitNull

`func (o *MicrosoftGraphContentType) SetGroupExplicitNull(b bool)`

SetGroupExplicitNull (un)sets Group to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Group value is set to nil even if false is passed
### GetHidden

`func (o *MicrosoftGraphContentType) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *MicrosoftGraphContentType) GetHiddenOk() (bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHidden

`func (o *MicrosoftGraphContentType) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHidden

`func (o *MicrosoftGraphContentType) SetHidden(v bool)`

SetHidden gets a reference to the given bool and assigns it to the Hidden field.

### SetHiddenExplicitNull

`func (o *MicrosoftGraphContentType) SetHiddenExplicitNull(b bool)`

SetHiddenExplicitNull (un)sets Hidden to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Hidden value is set to nil even if false is passed
### GetInheritedFrom

`func (o *MicrosoftGraphContentType) GetInheritedFrom() AnyOfmicrosoftGraphItemReference`

GetInheritedFrom returns the InheritedFrom field if non-nil, zero value otherwise.

### GetInheritedFromOk

`func (o *MicrosoftGraphContentType) GetInheritedFromOk() (AnyOfmicrosoftGraphItemReference, bool)`

GetInheritedFromOk returns a tuple with the InheritedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInheritedFrom

`func (o *MicrosoftGraphContentType) HasInheritedFrom() bool`

HasInheritedFrom returns a boolean if a field has been set.

### SetInheritedFrom

`func (o *MicrosoftGraphContentType) SetInheritedFrom(v AnyOfmicrosoftGraphItemReference)`

SetInheritedFrom gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the InheritedFrom field.

### SetInheritedFromExplicitNull

`func (o *MicrosoftGraphContentType) SetInheritedFromExplicitNull(b bool)`

SetInheritedFromExplicitNull (un)sets InheritedFrom to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InheritedFrom value is set to nil even if false is passed
### GetName

`func (o *MicrosoftGraphContentType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphContentType) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphContentType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphContentType) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphContentType) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetOrder

`func (o *MicrosoftGraphContentType) GetOrder() AnyOfmicrosoftGraphContentTypeOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MicrosoftGraphContentType) GetOrderOk() (AnyOfmicrosoftGraphContentTypeOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrder

`func (o *MicrosoftGraphContentType) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrder

`func (o *MicrosoftGraphContentType) SetOrder(v AnyOfmicrosoftGraphContentTypeOrder)`

SetOrder gets a reference to the given AnyOfmicrosoftGraphContentTypeOrder and assigns it to the Order field.

### SetOrderExplicitNull

`func (o *MicrosoftGraphContentType) SetOrderExplicitNull(b bool)`

SetOrderExplicitNull (un)sets Order to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Order value is set to nil even if false is passed
### GetParentId

`func (o *MicrosoftGraphContentType) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *MicrosoftGraphContentType) GetParentIdOk() (string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentId

`func (o *MicrosoftGraphContentType) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentId

`func (o *MicrosoftGraphContentType) SetParentId(v string)`

SetParentId gets a reference to the given string and assigns it to the ParentId field.

### SetParentIdExplicitNull

`func (o *MicrosoftGraphContentType) SetParentIdExplicitNull(b bool)`

SetParentIdExplicitNull (un)sets ParentId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentId value is set to nil even if false is passed
### GetReadOnly

`func (o *MicrosoftGraphContentType) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *MicrosoftGraphContentType) GetReadOnlyOk() (bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReadOnly

`func (o *MicrosoftGraphContentType) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnly

`func (o *MicrosoftGraphContentType) SetReadOnly(v bool)`

SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.

### SetReadOnlyExplicitNull

`func (o *MicrosoftGraphContentType) SetReadOnlyExplicitNull(b bool)`

SetReadOnlyExplicitNull (un)sets ReadOnly to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReadOnly value is set to nil even if false is passed
### GetSealed

`func (o *MicrosoftGraphContentType) GetSealed() bool`

GetSealed returns the Sealed field if non-nil, zero value otherwise.

### GetSealedOk

`func (o *MicrosoftGraphContentType) GetSealedOk() (bool, bool)`

GetSealedOk returns a tuple with the Sealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSealed

`func (o *MicrosoftGraphContentType) HasSealed() bool`

HasSealed returns a boolean if a field has been set.

### SetSealed

`func (o *MicrosoftGraphContentType) SetSealed(v bool)`

SetSealed gets a reference to the given bool and assigns it to the Sealed field.

### SetSealedExplicitNull

`func (o *MicrosoftGraphContentType) SetSealedExplicitNull(b bool)`

SetSealedExplicitNull (un)sets Sealed to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sealed value is set to nil even if false is passed
### GetColumnLinks

`func (o *MicrosoftGraphContentType) GetColumnLinks() []MicrosoftGraphColumnLink`

GetColumnLinks returns the ColumnLinks field if non-nil, zero value otherwise.

### GetColumnLinksOk

`func (o *MicrosoftGraphContentType) GetColumnLinksOk() ([]MicrosoftGraphColumnLink, bool)`

GetColumnLinksOk returns a tuple with the ColumnLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnLinks

`func (o *MicrosoftGraphContentType) HasColumnLinks() bool`

HasColumnLinks returns a boolean if a field has been set.

### SetColumnLinks

`func (o *MicrosoftGraphContentType) SetColumnLinks(v []MicrosoftGraphColumnLink)`

SetColumnLinks gets a reference to the given []MicrosoftGraphColumnLink and assigns it to the ColumnLinks field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


