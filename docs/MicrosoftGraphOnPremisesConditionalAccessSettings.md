# MicrosoftGraphOnPremisesConditionalAccessSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates if on premises conditional access is enabled for this organization | [optional] 
**IncludedGroups** | Pointer to **[]string** | User groups that will be targeted by on premises conditional access. All users in these groups will be required to have mobile device managed and compliant for mail access. | [optional] 
**ExcludedGroups** | Pointer to **[]string** | User groups that will be exempt by on premises conditional access. All users in these groups will be exempt from the conditional access policy. | [optional] 
**OverrideDefaultRule** | Pointer to **bool** | Override the default access rule when allowing a device to ensure access is granted. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetEnabled

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetIncludedGroups

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) GetIncludedGroups() []string`

GetIncludedGroups returns the IncludedGroups field if non-nil, zero value otherwise.

### GetIncludedGroupsOk

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) GetIncludedGroupsOk() ([]string, bool)`

GetIncludedGroupsOk returns a tuple with the IncludedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIncludedGroups

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) HasIncludedGroups() bool`

HasIncludedGroups returns a boolean if a field has been set.

### SetIncludedGroups

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) SetIncludedGroups(v []string)`

SetIncludedGroups gets a reference to the given []string and assigns it to the IncludedGroups field.

### GetExcludedGroups

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) GetExcludedGroups() []string`

GetExcludedGroups returns the ExcludedGroups field if non-nil, zero value otherwise.

### GetExcludedGroupsOk

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) GetExcludedGroupsOk() ([]string, bool)`

GetExcludedGroupsOk returns a tuple with the ExcludedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExcludedGroups

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) HasExcludedGroups() bool`

HasExcludedGroups returns a boolean if a field has been set.

### SetExcludedGroups

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) SetExcludedGroups(v []string)`

SetExcludedGroups gets a reference to the given []string and assigns it to the ExcludedGroups field.

### GetOverrideDefaultRule

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) GetOverrideDefaultRule() bool`

GetOverrideDefaultRule returns the OverrideDefaultRule field if non-nil, zero value otherwise.

### GetOverrideDefaultRuleOk

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) GetOverrideDefaultRuleOk() (bool, bool)`

GetOverrideDefaultRuleOk returns a tuple with the OverrideDefaultRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOverrideDefaultRule

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) HasOverrideDefaultRule() bool`

HasOverrideDefaultRule returns a boolean if a field has been set.

### SetOverrideDefaultRule

`func (o *MicrosoftGraphOnPremisesConditionalAccessSettings) SetOverrideDefaultRule(v bool)`

SetOverrideDefaultRule gets a reference to the given bool and assigns it to the OverrideDefaultRule field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


