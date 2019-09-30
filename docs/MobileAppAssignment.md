# MobileAppAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Intent** | Pointer to [**AnyOfmicrosoftGraphInstallIntent**](anyOf&lt;microsoft.graph.installIntent&gt;.md) | The install intent defined by the admin. | [optional] 
**Target** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The target group assignment defined by the admin. | [optional] 
**Settings** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The settings for target assignment defined by the admin. | [optional] 

## Methods

### GetIntent

`func (o *MobileAppAssignment) GetIntent() AnyOfmicrosoftGraphInstallIntent`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *MobileAppAssignment) GetIntentOk() (AnyOfmicrosoftGraphInstallIntent, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntent

`func (o *MobileAppAssignment) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### SetIntent

`func (o *MobileAppAssignment) SetIntent(v AnyOfmicrosoftGraphInstallIntent)`

SetIntent gets a reference to the given AnyOfmicrosoftGraphInstallIntent and assigns it to the Intent field.

### GetTarget

`func (o *MobileAppAssignment) GetTarget() AnyOfobject`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MobileAppAssignment) GetTargetOk() (AnyOfobject, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *MobileAppAssignment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *MobileAppAssignment) SetTarget(v AnyOfobject)`

SetTarget gets a reference to the given AnyOfobject and assigns it to the Target field.

### SetTargetExplicitNull

`func (o *MobileAppAssignment) SetTargetExplicitNull(b bool)`

SetTargetExplicitNull (un)sets Target to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Target value is set to nil even if false is passed
### GetSettings

`func (o *MobileAppAssignment) GetSettings() AnyOfobject`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MobileAppAssignment) GetSettingsOk() (AnyOfobject, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettings

`func (o *MobileAppAssignment) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettings

`func (o *MobileAppAssignment) SetSettings(v AnyOfobject)`

SetSettings gets a reference to the given AnyOfobject and assigns it to the Settings field.

### SetSettingsExplicitNull

`func (o *MobileAppAssignment) SetSettingsExplicitNull(b bool)`

SetSettingsExplicitNull (un)sets Settings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Settings value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


