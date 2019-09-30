# MicrosoftGraphOnenotePatchContentCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**AnyOfmicrosoftGraphOnenotePatchActionType**](anyOf&lt;microsoft.graph.onenotePatchActionType&gt;.md) |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Position** | Pointer to [**AnyOfmicrosoftGraphOnenotePatchInsertPosition**](anyOf&lt;microsoft.graph.onenotePatchInsertPosition&gt;.md) |  | [optional] 

## Methods

### GetAction

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetAction() AnyOfmicrosoftGraphOnenotePatchActionType`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetActionOk() (AnyOfmicrosoftGraphOnenotePatchActionType, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *MicrosoftGraphOnenotePatchContentCommand) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetAction(v AnyOfmicrosoftGraphOnenotePatchActionType)`

SetAction gets a reference to the given AnyOfmicrosoftGraphOnenotePatchActionType and assigns it to the Action field.

### GetTarget

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetTargetOk() (string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *MicrosoftGraphOnenotePatchContentCommand) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetTarget(v string)`

SetTarget gets a reference to the given string and assigns it to the Target field.

### GetContent

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetContentOk() (string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContent

`func (o *MicrosoftGraphOnenotePatchContentCommand) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContent

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetContent(v string)`

SetContent gets a reference to the given string and assigns it to the Content field.

### SetContentExplicitNull

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetContentExplicitNull(b bool)`

SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Content value is set to nil even if false is passed
### GetPosition

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetPosition() AnyOfmicrosoftGraphOnenotePatchInsertPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetPositionOk() (AnyOfmicrosoftGraphOnenotePatchInsertPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPosition

`func (o *MicrosoftGraphOnenotePatchContentCommand) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPosition

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetPosition(v AnyOfmicrosoftGraphOnenotePatchInsertPosition)`

SetPosition gets a reference to the given AnyOfmicrosoftGraphOnenotePatchInsertPosition and assigns it to the Position field.

### SetPositionExplicitNull

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetPositionExplicitNull(b bool)`

SetPositionExplicitNull (un)sets Position to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Position value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


