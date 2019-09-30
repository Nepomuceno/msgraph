# RemoteAssistancePartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of the partner. | [optional] 
**OnboardingUrl** | Pointer to **string** | URL of the partner&#39;s onboarding portal, where an administrator can configure their Remote Assistance service. | [optional] 
**OnboardingStatus** | Pointer to [**AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus**](anyOf&lt;microsoft.graph.remoteAssistanceOnboardingStatus&gt;.md) | TBD | [optional] 
**LastConnectionDateTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the last request sent to Intune by the TEM partner. | [optional] 

## Methods

### GetDisplayName

`func (o *RemoteAssistancePartner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RemoteAssistancePartner) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *RemoteAssistancePartner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *RemoteAssistancePartner) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *RemoteAssistancePartner) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetOnboardingUrl

`func (o *RemoteAssistancePartner) GetOnboardingUrl() string`

GetOnboardingUrl returns the OnboardingUrl field if non-nil, zero value otherwise.

### GetOnboardingUrlOk

`func (o *RemoteAssistancePartner) GetOnboardingUrlOk() (string, bool)`

GetOnboardingUrlOk returns a tuple with the OnboardingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnboardingUrl

`func (o *RemoteAssistancePartner) HasOnboardingUrl() bool`

HasOnboardingUrl returns a boolean if a field has been set.

### SetOnboardingUrl

`func (o *RemoteAssistancePartner) SetOnboardingUrl(v string)`

SetOnboardingUrl gets a reference to the given string and assigns it to the OnboardingUrl field.

### SetOnboardingUrlExplicitNull

`func (o *RemoteAssistancePartner) SetOnboardingUrlExplicitNull(b bool)`

SetOnboardingUrlExplicitNull (un)sets OnboardingUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnboardingUrl value is set to nil even if false is passed
### GetOnboardingStatus

`func (o *RemoteAssistancePartner) GetOnboardingStatus() AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus`

GetOnboardingStatus returns the OnboardingStatus field if non-nil, zero value otherwise.

### GetOnboardingStatusOk

`func (o *RemoteAssistancePartner) GetOnboardingStatusOk() (AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus, bool)`

GetOnboardingStatusOk returns a tuple with the OnboardingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnboardingStatus

`func (o *RemoteAssistancePartner) HasOnboardingStatus() bool`

HasOnboardingStatus returns a boolean if a field has been set.

### SetOnboardingStatus

`func (o *RemoteAssistancePartner) SetOnboardingStatus(v AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus)`

SetOnboardingStatus gets a reference to the given AnyOfmicrosoftGraphRemoteAssistanceOnboardingStatus and assigns it to the OnboardingStatus field.

### GetLastConnectionDateTime

`func (o *RemoteAssistancePartner) GetLastConnectionDateTime() time.Time`

GetLastConnectionDateTime returns the LastConnectionDateTime field if non-nil, zero value otherwise.

### GetLastConnectionDateTimeOk

`func (o *RemoteAssistancePartner) GetLastConnectionDateTimeOk() (time.Time, bool)`

GetLastConnectionDateTimeOk returns a tuple with the LastConnectionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastConnectionDateTime

`func (o *RemoteAssistancePartner) HasLastConnectionDateTime() bool`

HasLastConnectionDateTime returns a boolean if a field has been set.

### SetLastConnectionDateTime

`func (o *RemoteAssistancePartner) SetLastConnectionDateTime(v time.Time)`

SetLastConnectionDateTime gets a reference to the given time.Time and assigns it to the LastConnectionDateTime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


