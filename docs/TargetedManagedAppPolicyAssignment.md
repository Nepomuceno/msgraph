# TargetedManagedAppPolicyAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Identifier for deployment of a group or app | [optional] 

## Methods

### GetTarget

`func (o *TargetedManagedAppPolicyAssignment) GetTarget() AnyOfobject`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TargetedManagedAppPolicyAssignment) GetTargetOk() (AnyOfobject, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *TargetedManagedAppPolicyAssignment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *TargetedManagedAppPolicyAssignment) SetTarget(v AnyOfobject)`

SetTarget gets a reference to the given AnyOfobject and assigns it to the Target field.

### SetTargetExplicitNull

`func (o *TargetedManagedAppPolicyAssignment) SetTargetExplicitNull(b bool)`

SetTargetExplicitNull (un)sets Target to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Target value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


