# MicrosoftGraphMessageRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Sequence** | Pointer to **int32** |  | [optional] 
**Conditions** | Pointer to [**AnyOfmicrosoftGraphMessageRulePredicates**](anyOf&lt;microsoft.graph.messageRulePredicates&gt;.md) |  | [optional] 
**Actions** | Pointer to [**AnyOfmicrosoftGraphMessageRuleActions**](anyOf&lt;microsoft.graph.messageRuleActions&gt;.md) |  | [optional] 
**Exceptions** | Pointer to [**AnyOfmicrosoftGraphMessageRulePredicates**](anyOf&lt;microsoft.graph.messageRulePredicates&gt;.md) |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**HasError** | Pointer to **bool** |  | [optional] 
**IsReadOnly** | Pointer to **bool** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphMessageRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMessageRule) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphMessageRule) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphMessageRule) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphMessageRule) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphMessageRule) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphMessageRule) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphMessageRule) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphMessageRule) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetSequence

`func (o *MicrosoftGraphMessageRule) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *MicrosoftGraphMessageRule) GetSequenceOk() (int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSequence

`func (o *MicrosoftGraphMessageRule) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequence

`func (o *MicrosoftGraphMessageRule) SetSequence(v int32)`

SetSequence gets a reference to the given int32 and assigns it to the Sequence field.

### SetSequenceExplicitNull

`func (o *MicrosoftGraphMessageRule) SetSequenceExplicitNull(b bool)`

SetSequenceExplicitNull (un)sets Sequence to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sequence value is set to nil even if false is passed
### GetConditions

`func (o *MicrosoftGraphMessageRule) GetConditions() AnyOfmicrosoftGraphMessageRulePredicates`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MicrosoftGraphMessageRule) GetConditionsOk() (AnyOfmicrosoftGraphMessageRulePredicates, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConditions

`func (o *MicrosoftGraphMessageRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditions

`func (o *MicrosoftGraphMessageRule) SetConditions(v AnyOfmicrosoftGraphMessageRulePredicates)`

SetConditions gets a reference to the given AnyOfmicrosoftGraphMessageRulePredicates and assigns it to the Conditions field.

### SetConditionsExplicitNull

`func (o *MicrosoftGraphMessageRule) SetConditionsExplicitNull(b bool)`

SetConditionsExplicitNull (un)sets Conditions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Conditions value is set to nil even if false is passed
### GetActions

`func (o *MicrosoftGraphMessageRule) GetActions() AnyOfmicrosoftGraphMessageRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MicrosoftGraphMessageRule) GetActionsOk() (AnyOfmicrosoftGraphMessageRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActions

`func (o *MicrosoftGraphMessageRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActions

`func (o *MicrosoftGraphMessageRule) SetActions(v AnyOfmicrosoftGraphMessageRuleActions)`

SetActions gets a reference to the given AnyOfmicrosoftGraphMessageRuleActions and assigns it to the Actions field.

### SetActionsExplicitNull

`func (o *MicrosoftGraphMessageRule) SetActionsExplicitNull(b bool)`

SetActionsExplicitNull (un)sets Actions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Actions value is set to nil even if false is passed
### GetExceptions

`func (o *MicrosoftGraphMessageRule) GetExceptions() AnyOfmicrosoftGraphMessageRulePredicates`

GetExceptions returns the Exceptions field if non-nil, zero value otherwise.

### GetExceptionsOk

`func (o *MicrosoftGraphMessageRule) GetExceptionsOk() (AnyOfmicrosoftGraphMessageRulePredicates, bool)`

GetExceptionsOk returns a tuple with the Exceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExceptions

`func (o *MicrosoftGraphMessageRule) HasExceptions() bool`

HasExceptions returns a boolean if a field has been set.

### SetExceptions

`func (o *MicrosoftGraphMessageRule) SetExceptions(v AnyOfmicrosoftGraphMessageRulePredicates)`

SetExceptions gets a reference to the given AnyOfmicrosoftGraphMessageRulePredicates and assigns it to the Exceptions field.

### SetExceptionsExplicitNull

`func (o *MicrosoftGraphMessageRule) SetExceptionsExplicitNull(b bool)`

SetExceptionsExplicitNull (un)sets Exceptions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Exceptions value is set to nil even if false is passed
### GetIsEnabled

`func (o *MicrosoftGraphMessageRule) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MicrosoftGraphMessageRule) GetIsEnabledOk() (bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabled

`func (o *MicrosoftGraphMessageRule) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabled

`func (o *MicrosoftGraphMessageRule) SetIsEnabled(v bool)`

SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.

### SetIsEnabledExplicitNull

`func (o *MicrosoftGraphMessageRule) SetIsEnabledExplicitNull(b bool)`

SetIsEnabledExplicitNull (un)sets IsEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsEnabled value is set to nil even if false is passed
### GetHasError

`func (o *MicrosoftGraphMessageRule) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *MicrosoftGraphMessageRule) GetHasErrorOk() (bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasError

`func (o *MicrosoftGraphMessageRule) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### SetHasError

`func (o *MicrosoftGraphMessageRule) SetHasError(v bool)`

SetHasError gets a reference to the given bool and assigns it to the HasError field.

### SetHasErrorExplicitNull

`func (o *MicrosoftGraphMessageRule) SetHasErrorExplicitNull(b bool)`

SetHasErrorExplicitNull (un)sets HasError to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HasError value is set to nil even if false is passed
### GetIsReadOnly

`func (o *MicrosoftGraphMessageRule) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *MicrosoftGraphMessageRule) GetIsReadOnlyOk() (bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReadOnly

`func (o *MicrosoftGraphMessageRule) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnly

`func (o *MicrosoftGraphMessageRule) SetIsReadOnly(v bool)`

SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.

### SetIsReadOnlyExplicitNull

`func (o *MicrosoftGraphMessageRule) SetIsReadOnlyExplicitNull(b bool)`

SetIsReadOnlyExplicitNull (un)sets IsReadOnly to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsReadOnly value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


