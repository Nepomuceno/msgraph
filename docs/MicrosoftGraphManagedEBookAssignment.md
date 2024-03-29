# MicrosoftGraphManagedEBookAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Target** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The assignment target for eBook. | [optional] 
**InstallIntent** | Pointer to [**AnyOfmicrosoftGraphInstallIntent**](anyOf&lt;microsoft.graph.installIntent&gt;.md) | The install intent for eBook. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphManagedEBookAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedEBookAssignment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphManagedEBookAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphManagedEBookAssignment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetTarget

`func (o *MicrosoftGraphManagedEBookAssignment) GetTarget() AnyOfobject`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MicrosoftGraphManagedEBookAssignment) GetTargetOk() (AnyOfobject, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *MicrosoftGraphManagedEBookAssignment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *MicrosoftGraphManagedEBookAssignment) SetTarget(v AnyOfobject)`

SetTarget gets a reference to the given AnyOfobject and assigns it to the Target field.

### SetTargetExplicitNull

`func (o *MicrosoftGraphManagedEBookAssignment) SetTargetExplicitNull(b bool)`

SetTargetExplicitNull (un)sets Target to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Target value is set to nil even if false is passed
### GetInstallIntent

`func (o *MicrosoftGraphManagedEBookAssignment) GetInstallIntent() AnyOfmicrosoftGraphInstallIntent`

GetInstallIntent returns the InstallIntent field if non-nil, zero value otherwise.

### GetInstallIntentOk

`func (o *MicrosoftGraphManagedEBookAssignment) GetInstallIntentOk() (AnyOfmicrosoftGraphInstallIntent, bool)`

GetInstallIntentOk returns a tuple with the InstallIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstallIntent

`func (o *MicrosoftGraphManagedEBookAssignment) HasInstallIntent() bool`

HasInstallIntent returns a boolean if a field has been set.

### SetInstallIntent

`func (o *MicrosoftGraphManagedEBookAssignment) SetInstallIntent(v AnyOfmicrosoftGraphInstallIntent)`

SetInstallIntent gets a reference to the given AnyOfmicrosoftGraphInstallIntent and assigns it to the InstallIntent field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


