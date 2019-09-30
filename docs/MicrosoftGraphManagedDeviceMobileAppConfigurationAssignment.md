# MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Assignment target that the T&amp;C policy is assigned to. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetTarget

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment) GetTarget() AnyOfobject`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment) GetTargetOk() (AnyOfobject, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment) SetTarget(v AnyOfobject)`

SetTarget gets a reference to the given AnyOfobject and assigns it to the Target field.

### SetTargetExplicitNull

`func (o *MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment) SetTargetExplicitNull(b bool)`

SetTargetExplicitNull (un)sets Target to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Target value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


