# MicrosoftGraphPlannerChecklistItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsChecked** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**OrderHint** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### GetIsChecked

`func (o *MicrosoftGraphPlannerChecklistItem) GetIsChecked() bool`

GetIsChecked returns the IsChecked field if non-nil, zero value otherwise.

### GetIsCheckedOk

`func (o *MicrosoftGraphPlannerChecklistItem) GetIsCheckedOk() (bool, bool)`

GetIsCheckedOk returns a tuple with the IsChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsChecked

`func (o *MicrosoftGraphPlannerChecklistItem) HasIsChecked() bool`

HasIsChecked returns a boolean if a field has been set.

### SetIsChecked

`func (o *MicrosoftGraphPlannerChecklistItem) SetIsChecked(v bool)`

SetIsChecked gets a reference to the given bool and assigns it to the IsChecked field.

### SetIsCheckedExplicitNull

`func (o *MicrosoftGraphPlannerChecklistItem) SetIsCheckedExplicitNull(b bool)`

SetIsCheckedExplicitNull (un)sets IsChecked to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsChecked value is set to nil even if false is passed
### GetTitle

`func (o *MicrosoftGraphPlannerChecklistItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphPlannerChecklistItem) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *MicrosoftGraphPlannerChecklistItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *MicrosoftGraphPlannerChecklistItem) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### SetTitleExplicitNull

`func (o *MicrosoftGraphPlannerChecklistItem) SetTitleExplicitNull(b bool)`

SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Title value is set to nil even if false is passed
### GetOrderHint

`func (o *MicrosoftGraphPlannerChecklistItem) GetOrderHint() string`

GetOrderHint returns the OrderHint field if non-nil, zero value otherwise.

### GetOrderHintOk

`func (o *MicrosoftGraphPlannerChecklistItem) GetOrderHintOk() (string, bool)`

GetOrderHintOk returns a tuple with the OrderHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrderHint

`func (o *MicrosoftGraphPlannerChecklistItem) HasOrderHint() bool`

HasOrderHint returns a boolean if a field has been set.

### SetOrderHint

`func (o *MicrosoftGraphPlannerChecklistItem) SetOrderHint(v string)`

SetOrderHint gets a reference to the given string and assigns it to the OrderHint field.

### SetOrderHintExplicitNull

`func (o *MicrosoftGraphPlannerChecklistItem) SetOrderHintExplicitNull(b bool)`

SetOrderHintExplicitNull (un)sets OrderHint to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OrderHint value is set to nil even if false is passed
### GetLastModifiedBy

`func (o *MicrosoftGraphPlannerChecklistItem) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphPlannerChecklistItem) GetLastModifiedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphPlannerChecklistItem) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphPlannerChecklistItem) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphPlannerChecklistItem) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphPlannerChecklistItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphPlannerChecklistItem) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphPlannerChecklistItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphPlannerChecklistItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphPlannerChecklistItem) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


