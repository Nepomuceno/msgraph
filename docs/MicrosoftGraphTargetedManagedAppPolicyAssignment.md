# MicrosoftGraphTargetedManagedAppPolicyAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Identifier for deployment of a group or app | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetTarget

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) GetTarget() AnyOfobject`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) GetTargetOk() (AnyOfobject, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) SetTarget(v AnyOfobject)`

SetTarget gets a reference to the given AnyOfobject and assigns it to the Target field.

### SetTargetExplicitNull

`func (o *MicrosoftGraphTargetedManagedAppPolicyAssignment) SetTargetExplicitNull(b bool)`

SetTargetExplicitNull (un)sets Target to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Target value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


