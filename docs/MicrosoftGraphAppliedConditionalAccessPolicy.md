# MicrosoftGraphAppliedConditionalAccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EnforcedGrantControls** | Pointer to **[]string** |  | [optional] 
**EnforcedSessionControls** | Pointer to **[]string** |  | [optional] 
**Result** | Pointer to [**AnyOfmicrosoftGraphAppliedConditionalAccessPolicyResult**](anyOf&lt;microsoft.graph.appliedConditionalAccessPolicyResult&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### SetIdExplicitNull

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetIdExplicitNull(b bool)`

SetIdExplicitNull (un)sets Id to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Id value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetEnforcedGrantControls

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetEnforcedGrantControls() []string`

GetEnforcedGrantControls returns the EnforcedGrantControls field if non-nil, zero value otherwise.

### GetEnforcedGrantControlsOk

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetEnforcedGrantControlsOk() ([]string, bool)`

GetEnforcedGrantControlsOk returns a tuple with the EnforcedGrantControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnforcedGrantControls

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) HasEnforcedGrantControls() bool`

HasEnforcedGrantControls returns a boolean if a field has been set.

### SetEnforcedGrantControls

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetEnforcedGrantControls(v []string)`

SetEnforcedGrantControls gets a reference to the given []string and assigns it to the EnforcedGrantControls field.

### GetEnforcedSessionControls

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetEnforcedSessionControls() []string`

GetEnforcedSessionControls returns the EnforcedSessionControls field if non-nil, zero value otherwise.

### GetEnforcedSessionControlsOk

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetEnforcedSessionControlsOk() ([]string, bool)`

GetEnforcedSessionControlsOk returns a tuple with the EnforcedSessionControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnforcedSessionControls

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) HasEnforcedSessionControls() bool`

HasEnforcedSessionControls returns a boolean if a field has been set.

### SetEnforcedSessionControls

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetEnforcedSessionControls(v []string)`

SetEnforcedSessionControls gets a reference to the given []string and assigns it to the EnforcedSessionControls field.

### GetResult

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetResult() AnyOfmicrosoftGraphAppliedConditionalAccessPolicyResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) GetResultOk() (AnyOfmicrosoftGraphAppliedConditionalAccessPolicyResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResult

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResult

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetResult(v AnyOfmicrosoftGraphAppliedConditionalAccessPolicyResult)`

SetResult gets a reference to the given AnyOfmicrosoftGraphAppliedConditionalAccessPolicyResult and assigns it to the Result field.

### SetResultExplicitNull

`func (o *MicrosoftGraphAppliedConditionalAccessPolicy) SetResultExplicitNull(b bool)`

SetResultExplicitNull (un)sets Result to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Result value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


