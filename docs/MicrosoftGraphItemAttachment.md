# MicrosoftGraphItemAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**IsInline** | Pointer to **bool** |  | [optional] 
**Item** | Pointer to [**AnyOfmicrosoftGraphOutlookItem**](anyOf&lt;microsoft.graph.outlookItem&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphItemAttachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphItemAttachment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphItemAttachment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphItemAttachment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphItemAttachment) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphItemAttachment) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphItemAttachment) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphItemAttachment) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphItemAttachment) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetName

`func (o *MicrosoftGraphItemAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphItemAttachment) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphItemAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphItemAttachment) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphItemAttachment) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetContentType

`func (o *MicrosoftGraphItemAttachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MicrosoftGraphItemAttachment) GetContentTypeOk() (string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentType

`func (o *MicrosoftGraphItemAttachment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentType

`func (o *MicrosoftGraphItemAttachment) SetContentType(v string)`

SetContentType gets a reference to the given string and assigns it to the ContentType field.

### SetContentTypeExplicitNull

`func (o *MicrosoftGraphItemAttachment) SetContentTypeExplicitNull(b bool)`

SetContentTypeExplicitNull (un)sets ContentType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContentType value is set to nil even if false is passed
### GetSize

`func (o *MicrosoftGraphItemAttachment) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphItemAttachment) GetSizeOk() (int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *MicrosoftGraphItemAttachment) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *MicrosoftGraphItemAttachment) SetSize(v int32)`

SetSize gets a reference to the given int32 and assigns it to the Size field.

### GetIsInline

`func (o *MicrosoftGraphItemAttachment) GetIsInline() bool`

GetIsInline returns the IsInline field if non-nil, zero value otherwise.

### GetIsInlineOk

`func (o *MicrosoftGraphItemAttachment) GetIsInlineOk() (bool, bool)`

GetIsInlineOk returns a tuple with the IsInline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsInline

`func (o *MicrosoftGraphItemAttachment) HasIsInline() bool`

HasIsInline returns a boolean if a field has been set.

### SetIsInline

`func (o *MicrosoftGraphItemAttachment) SetIsInline(v bool)`

SetIsInline gets a reference to the given bool and assigns it to the IsInline field.

### GetItem

`func (o *MicrosoftGraphItemAttachment) GetItem() AnyOfmicrosoftGraphOutlookItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *MicrosoftGraphItemAttachment) GetItemOk() (AnyOfmicrosoftGraphOutlookItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItem

`func (o *MicrosoftGraphItemAttachment) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItem

`func (o *MicrosoftGraphItemAttachment) SetItem(v AnyOfmicrosoftGraphOutlookItem)`

SetItem gets a reference to the given AnyOfmicrosoftGraphOutlookItem and assigns it to the Item field.

### SetItemExplicitNull

`func (o *MicrosoftGraphItemAttachment) SetItemExplicitNull(b bool)`

SetItemExplicitNull (un)sets Item to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Item value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


