# MessageRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Sequence** | Pointer to **int32** |  | [optional] 
**Conditions** | Pointer to [**AnyOfmicrosoftGraphMessageRulePredicates**](anyOf&lt;microsoft.graph.messageRulePredicates&gt;.md) |  | [optional] 
**Actions** | Pointer to [**AnyOfmicrosoftGraphMessageRuleActions**](anyOf&lt;microsoft.graph.messageRuleActions&gt;.md) |  | [optional] 
**Exceptions** | Pointer to [**AnyOfmicrosoftGraphMessageRulePredicates**](anyOf&lt;microsoft.graph.messageRulePredicates&gt;.md) |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**HasError** | Pointer to **bool** |  | [optional] 
**IsReadOnly** | Pointer to **bool** |  | [optional] 

## Methods

### GetDisplayName

`func (o *MessageRule) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MessageRule) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MessageRule) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MessageRule) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MessageRule) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetSequence

`func (o *MessageRule) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *MessageRule) GetSequenceOk() (int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSequence

`func (o *MessageRule) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequence

`func (o *MessageRule) SetSequence(v int32)`

SetSequence gets a reference to the given int32 and assigns it to the Sequence field.

### SetSequenceExplicitNull

`func (o *MessageRule) SetSequenceExplicitNull(b bool)`

SetSequenceExplicitNull (un)sets Sequence to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sequence value is set to nil even if false is passed
### GetConditions

`func (o *MessageRule) GetConditions() AnyOfmicrosoftGraphMessageRulePredicates`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MessageRule) GetConditionsOk() (AnyOfmicrosoftGraphMessageRulePredicates, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConditions

`func (o *MessageRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditions

`func (o *MessageRule) SetConditions(v AnyOfmicrosoftGraphMessageRulePredicates)`

SetConditions gets a reference to the given AnyOfmicrosoftGraphMessageRulePredicates and assigns it to the Conditions field.

### SetConditionsExplicitNull

`func (o *MessageRule) SetConditionsExplicitNull(b bool)`

SetConditionsExplicitNull (un)sets Conditions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Conditions value is set to nil even if false is passed
### GetActions

`func (o *MessageRule) GetActions() AnyOfmicrosoftGraphMessageRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MessageRule) GetActionsOk() (AnyOfmicrosoftGraphMessageRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActions

`func (o *MessageRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActions

`func (o *MessageRule) SetActions(v AnyOfmicrosoftGraphMessageRuleActions)`

SetActions gets a reference to the given AnyOfmicrosoftGraphMessageRuleActions and assigns it to the Actions field.

### SetActionsExplicitNull

`func (o *MessageRule) SetActionsExplicitNull(b bool)`

SetActionsExplicitNull (un)sets Actions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Actions value is set to nil even if false is passed
### GetExceptions

`func (o *MessageRule) GetExceptions() AnyOfmicrosoftGraphMessageRulePredicates`

GetExceptions returns the Exceptions field if non-nil, zero value otherwise.

### GetExceptionsOk

`func (o *MessageRule) GetExceptionsOk() (AnyOfmicrosoftGraphMessageRulePredicates, bool)`

GetExceptionsOk returns a tuple with the Exceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExceptions

`func (o *MessageRule) HasExceptions() bool`

HasExceptions returns a boolean if a field has been set.

### SetExceptions

`func (o *MessageRule) SetExceptions(v AnyOfmicrosoftGraphMessageRulePredicates)`

SetExceptions gets a reference to the given AnyOfmicrosoftGraphMessageRulePredicates and assigns it to the Exceptions field.

### SetExceptionsExplicitNull

`func (o *MessageRule) SetExceptionsExplicitNull(b bool)`

SetExceptionsExplicitNull (un)sets Exceptions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Exceptions value is set to nil even if false is passed
### GetIsEnabled

`func (o *MessageRule) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MessageRule) GetIsEnabledOk() (bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabled

`func (o *MessageRule) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabled

`func (o *MessageRule) SetIsEnabled(v bool)`

SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.

### SetIsEnabledExplicitNull

`func (o *MessageRule) SetIsEnabledExplicitNull(b bool)`

SetIsEnabledExplicitNull (un)sets IsEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsEnabled value is set to nil even if false is passed
### GetHasError

`func (o *MessageRule) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *MessageRule) GetHasErrorOk() (bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasError

`func (o *MessageRule) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### SetHasError

`func (o *MessageRule) SetHasError(v bool)`

SetHasError gets a reference to the given bool and assigns it to the HasError field.

### SetHasErrorExplicitNull

`func (o *MessageRule) SetHasErrorExplicitNull(b bool)`

SetHasErrorExplicitNull (un)sets HasError to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HasError value is set to nil even if false is passed
### GetIsReadOnly

`func (o *MessageRule) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *MessageRule) GetIsReadOnlyOk() (bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReadOnly

`func (o *MessageRule) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnly

`func (o *MessageRule) SetIsReadOnly(v bool)`

SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.

### SetIsReadOnlyExplicitNull

`func (o *MessageRule) SetIsReadOnlyExplicitNull(b bool)`

SetIsReadOnlyExplicitNull (un)sets IsReadOnly to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsReadOnly value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


