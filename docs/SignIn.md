# SignIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**UserDisplayName** | Pointer to **string** |  | [optional] 
**UserPrincipalName** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**AppDisplayName** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphSignInStatus**](anyOf&lt;microsoft.graph.signInStatus&gt;.md) |  | [optional] 
**ClientAppUsed** | Pointer to **string** |  | [optional] 
**DeviceDetail** | Pointer to [**AnyOfmicrosoftGraphDeviceDetail**](anyOf&lt;microsoft.graph.deviceDetail&gt;.md) |  | [optional] 
**Location** | Pointer to [**AnyOfmicrosoftGraphSignInLocation**](anyOf&lt;microsoft.graph.signInLocation&gt;.md) |  | [optional] 
**CorrelationId** | Pointer to **string** |  | [optional] 
**ConditionalAccessStatus** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessStatus**](anyOf&lt;microsoft.graph.conditionalAccessStatus&gt;.md) |  | [optional] 
**AppliedConditionalAccessPolicies** | Pointer to [**[]AnyOfmicrosoftGraphAppliedConditionalAccessPolicy**](anyOf&lt;microsoft.graph.appliedConditionalAccessPolicy&gt;.md) |  | [optional] 
**IsInteractive** | Pointer to **bool** |  | [optional] 
**RiskDetail** | Pointer to [**AnyOfmicrosoftGraphRiskDetail**](anyOf&lt;microsoft.graph.riskDetail&gt;.md) |  | [optional] 
**RiskLevelAggregated** | Pointer to [**AnyOfmicrosoftGraphRiskLevel**](anyOf&lt;microsoft.graph.riskLevel&gt;.md) |  | [optional] 
**RiskLevelDuringSignIn** | Pointer to [**AnyOfmicrosoftGraphRiskLevel**](anyOf&lt;microsoft.graph.riskLevel&gt;.md) |  | [optional] 
**RiskState** | Pointer to [**AnyOfmicrosoftGraphRiskState**](anyOf&lt;microsoft.graph.riskState&gt;.md) |  | [optional] 
**RiskEventTypes** | Pointer to [**[]AnyOfmicrosoftGraphRiskEventType**](anyOf&lt;microsoft.graph.riskEventType&gt;.md) |  | [optional] 
**ResourceDisplayName** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 

## Methods

### GetCreatedDateTime

`func (o *SignIn) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *SignIn) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *SignIn) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *SignIn) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetUserDisplayName

`func (o *SignIn) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *SignIn) GetUserDisplayNameOk() (string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserDisplayName

`func (o *SignIn) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayName

`func (o *SignIn) SetUserDisplayName(v string)`

SetUserDisplayName gets a reference to the given string and assigns it to the UserDisplayName field.

### SetUserDisplayNameExplicitNull

`func (o *SignIn) SetUserDisplayNameExplicitNull(b bool)`

SetUserDisplayNameExplicitNull (un)sets UserDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserDisplayName value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *SignIn) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *SignIn) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *SignIn) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *SignIn) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *SignIn) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetUserId

`func (o *SignIn) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SignIn) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *SignIn) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *SignIn) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### GetAppId

`func (o *SignIn) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *SignIn) GetAppIdOk() (string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppId

`func (o *SignIn) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppId

`func (o *SignIn) SetAppId(v string)`

SetAppId gets a reference to the given string and assigns it to the AppId field.

### SetAppIdExplicitNull

`func (o *SignIn) SetAppIdExplicitNull(b bool)`

SetAppIdExplicitNull (un)sets AppId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppId value is set to nil even if false is passed
### GetAppDisplayName

`func (o *SignIn) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *SignIn) GetAppDisplayNameOk() (string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppDisplayName

`func (o *SignIn) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayName

`func (o *SignIn) SetAppDisplayName(v string)`

SetAppDisplayName gets a reference to the given string and assigns it to the AppDisplayName field.

### SetAppDisplayNameExplicitNull

`func (o *SignIn) SetAppDisplayNameExplicitNull(b bool)`

SetAppDisplayNameExplicitNull (un)sets AppDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppDisplayName value is set to nil even if false is passed
### GetIpAddress

`func (o *SignIn) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SignIn) GetIpAddressOk() (string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpAddress

`func (o *SignIn) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddress

`func (o *SignIn) SetIpAddress(v string)`

SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.

### SetIpAddressExplicitNull

`func (o *SignIn) SetIpAddressExplicitNull(b bool)`

SetIpAddressExplicitNull (un)sets IpAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IpAddress value is set to nil even if false is passed
### GetStatus

`func (o *SignIn) GetStatus() AnyOfmicrosoftGraphSignInStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SignIn) GetStatusOk() (AnyOfmicrosoftGraphSignInStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *SignIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *SignIn) SetStatus(v AnyOfmicrosoftGraphSignInStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphSignInStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *SignIn) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetClientAppUsed

`func (o *SignIn) GetClientAppUsed() string`

GetClientAppUsed returns the ClientAppUsed field if non-nil, zero value otherwise.

### GetClientAppUsedOk

`func (o *SignIn) GetClientAppUsedOk() (string, bool)`

GetClientAppUsedOk returns a tuple with the ClientAppUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientAppUsed

`func (o *SignIn) HasClientAppUsed() bool`

HasClientAppUsed returns a boolean if a field has been set.

### SetClientAppUsed

`func (o *SignIn) SetClientAppUsed(v string)`

SetClientAppUsed gets a reference to the given string and assigns it to the ClientAppUsed field.

### SetClientAppUsedExplicitNull

`func (o *SignIn) SetClientAppUsedExplicitNull(b bool)`

SetClientAppUsedExplicitNull (un)sets ClientAppUsed to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ClientAppUsed value is set to nil even if false is passed
### GetDeviceDetail

`func (o *SignIn) GetDeviceDetail() AnyOfmicrosoftGraphDeviceDetail`

GetDeviceDetail returns the DeviceDetail field if non-nil, zero value otherwise.

### GetDeviceDetailOk

`func (o *SignIn) GetDeviceDetailOk() (AnyOfmicrosoftGraphDeviceDetail, bool)`

GetDeviceDetailOk returns a tuple with the DeviceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceDetail

`func (o *SignIn) HasDeviceDetail() bool`

HasDeviceDetail returns a boolean if a field has been set.

### SetDeviceDetail

`func (o *SignIn) SetDeviceDetail(v AnyOfmicrosoftGraphDeviceDetail)`

SetDeviceDetail gets a reference to the given AnyOfmicrosoftGraphDeviceDetail and assigns it to the DeviceDetail field.

### SetDeviceDetailExplicitNull

`func (o *SignIn) SetDeviceDetailExplicitNull(b bool)`

SetDeviceDetailExplicitNull (un)sets DeviceDetail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceDetail value is set to nil even if false is passed
### GetLocation

`func (o *SignIn) GetLocation() AnyOfmicrosoftGraphSignInLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SignIn) GetLocationOk() (AnyOfmicrosoftGraphSignInLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocation

`func (o *SignIn) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocation

`func (o *SignIn) SetLocation(v AnyOfmicrosoftGraphSignInLocation)`

SetLocation gets a reference to the given AnyOfmicrosoftGraphSignInLocation and assigns it to the Location field.

### SetLocationExplicitNull

`func (o *SignIn) SetLocationExplicitNull(b bool)`

SetLocationExplicitNull (un)sets Location to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Location value is set to nil even if false is passed
### GetCorrelationId

`func (o *SignIn) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *SignIn) GetCorrelationIdOk() (string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCorrelationId

`func (o *SignIn) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationId

`func (o *SignIn) SetCorrelationId(v string)`

SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.

### SetCorrelationIdExplicitNull

`func (o *SignIn) SetCorrelationIdExplicitNull(b bool)`

SetCorrelationIdExplicitNull (un)sets CorrelationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CorrelationId value is set to nil even if false is passed
### GetConditionalAccessStatus

`func (o *SignIn) GetConditionalAccessStatus() AnyOfmicrosoftGraphConditionalAccessStatus`

GetConditionalAccessStatus returns the ConditionalAccessStatus field if non-nil, zero value otherwise.

### GetConditionalAccessStatusOk

`func (o *SignIn) GetConditionalAccessStatusOk() (AnyOfmicrosoftGraphConditionalAccessStatus, bool)`

GetConditionalAccessStatusOk returns a tuple with the ConditionalAccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConditionalAccessStatus

`func (o *SignIn) HasConditionalAccessStatus() bool`

HasConditionalAccessStatus returns a boolean if a field has been set.

### SetConditionalAccessStatus

`func (o *SignIn) SetConditionalAccessStatus(v AnyOfmicrosoftGraphConditionalAccessStatus)`

SetConditionalAccessStatus gets a reference to the given AnyOfmicrosoftGraphConditionalAccessStatus and assigns it to the ConditionalAccessStatus field.

### SetConditionalAccessStatusExplicitNull

`func (o *SignIn) SetConditionalAccessStatusExplicitNull(b bool)`

SetConditionalAccessStatusExplicitNull (un)sets ConditionalAccessStatus to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConditionalAccessStatus value is set to nil even if false is passed
### GetAppliedConditionalAccessPolicies

`func (o *SignIn) GetAppliedConditionalAccessPolicies() []AnyOfmicrosoftGraphAppliedConditionalAccessPolicy`

GetAppliedConditionalAccessPolicies returns the AppliedConditionalAccessPolicies field if non-nil, zero value otherwise.

### GetAppliedConditionalAccessPoliciesOk

`func (o *SignIn) GetAppliedConditionalAccessPoliciesOk() ([]AnyOfmicrosoftGraphAppliedConditionalAccessPolicy, bool)`

GetAppliedConditionalAccessPoliciesOk returns a tuple with the AppliedConditionalAccessPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppliedConditionalAccessPolicies

`func (o *SignIn) HasAppliedConditionalAccessPolicies() bool`

HasAppliedConditionalAccessPolicies returns a boolean if a field has been set.

### SetAppliedConditionalAccessPolicies

`func (o *SignIn) SetAppliedConditionalAccessPolicies(v []AnyOfmicrosoftGraphAppliedConditionalAccessPolicy)`

SetAppliedConditionalAccessPolicies gets a reference to the given []AnyOfmicrosoftGraphAppliedConditionalAccessPolicy and assigns it to the AppliedConditionalAccessPolicies field.

### GetIsInteractive

`func (o *SignIn) GetIsInteractive() bool`

GetIsInteractive returns the IsInteractive field if non-nil, zero value otherwise.

### GetIsInteractiveOk

`func (o *SignIn) GetIsInteractiveOk() (bool, bool)`

GetIsInteractiveOk returns a tuple with the IsInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsInteractive

`func (o *SignIn) HasIsInteractive() bool`

HasIsInteractive returns a boolean if a field has been set.

### SetIsInteractive

`func (o *SignIn) SetIsInteractive(v bool)`

SetIsInteractive gets a reference to the given bool and assigns it to the IsInteractive field.

### SetIsInteractiveExplicitNull

`func (o *SignIn) SetIsInteractiveExplicitNull(b bool)`

SetIsInteractiveExplicitNull (un)sets IsInteractive to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsInteractive value is set to nil even if false is passed
### GetRiskDetail

`func (o *SignIn) GetRiskDetail() AnyOfmicrosoftGraphRiskDetail`

GetRiskDetail returns the RiskDetail field if non-nil, zero value otherwise.

### GetRiskDetailOk

`func (o *SignIn) GetRiskDetailOk() (AnyOfmicrosoftGraphRiskDetail, bool)`

GetRiskDetailOk returns a tuple with the RiskDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskDetail

`func (o *SignIn) HasRiskDetail() bool`

HasRiskDetail returns a boolean if a field has been set.

### SetRiskDetail

`func (o *SignIn) SetRiskDetail(v AnyOfmicrosoftGraphRiskDetail)`

SetRiskDetail gets a reference to the given AnyOfmicrosoftGraphRiskDetail and assigns it to the RiskDetail field.

### SetRiskDetailExplicitNull

`func (o *SignIn) SetRiskDetailExplicitNull(b bool)`

SetRiskDetailExplicitNull (un)sets RiskDetail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskDetail value is set to nil even if false is passed
### GetRiskLevelAggregated

`func (o *SignIn) GetRiskLevelAggregated() AnyOfmicrosoftGraphRiskLevel`

GetRiskLevelAggregated returns the RiskLevelAggregated field if non-nil, zero value otherwise.

### GetRiskLevelAggregatedOk

`func (o *SignIn) GetRiskLevelAggregatedOk() (AnyOfmicrosoftGraphRiskLevel, bool)`

GetRiskLevelAggregatedOk returns a tuple with the RiskLevelAggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskLevelAggregated

`func (o *SignIn) HasRiskLevelAggregated() bool`

HasRiskLevelAggregated returns a boolean if a field has been set.

### SetRiskLevelAggregated

`func (o *SignIn) SetRiskLevelAggregated(v AnyOfmicrosoftGraphRiskLevel)`

SetRiskLevelAggregated gets a reference to the given AnyOfmicrosoftGraphRiskLevel and assigns it to the RiskLevelAggregated field.

### SetRiskLevelAggregatedExplicitNull

`func (o *SignIn) SetRiskLevelAggregatedExplicitNull(b bool)`

SetRiskLevelAggregatedExplicitNull (un)sets RiskLevelAggregated to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskLevelAggregated value is set to nil even if false is passed
### GetRiskLevelDuringSignIn

`func (o *SignIn) GetRiskLevelDuringSignIn() AnyOfmicrosoftGraphRiskLevel`

GetRiskLevelDuringSignIn returns the RiskLevelDuringSignIn field if non-nil, zero value otherwise.

### GetRiskLevelDuringSignInOk

`func (o *SignIn) GetRiskLevelDuringSignInOk() (AnyOfmicrosoftGraphRiskLevel, bool)`

GetRiskLevelDuringSignInOk returns a tuple with the RiskLevelDuringSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskLevelDuringSignIn

`func (o *SignIn) HasRiskLevelDuringSignIn() bool`

HasRiskLevelDuringSignIn returns a boolean if a field has been set.

### SetRiskLevelDuringSignIn

`func (o *SignIn) SetRiskLevelDuringSignIn(v AnyOfmicrosoftGraphRiskLevel)`

SetRiskLevelDuringSignIn gets a reference to the given AnyOfmicrosoftGraphRiskLevel and assigns it to the RiskLevelDuringSignIn field.

### SetRiskLevelDuringSignInExplicitNull

`func (o *SignIn) SetRiskLevelDuringSignInExplicitNull(b bool)`

SetRiskLevelDuringSignInExplicitNull (un)sets RiskLevelDuringSignIn to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskLevelDuringSignIn value is set to nil even if false is passed
### GetRiskState

`func (o *SignIn) GetRiskState() AnyOfmicrosoftGraphRiskState`

GetRiskState returns the RiskState field if non-nil, zero value otherwise.

### GetRiskStateOk

`func (o *SignIn) GetRiskStateOk() (AnyOfmicrosoftGraphRiskState, bool)`

GetRiskStateOk returns a tuple with the RiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskState

`func (o *SignIn) HasRiskState() bool`

HasRiskState returns a boolean if a field has been set.

### SetRiskState

`func (o *SignIn) SetRiskState(v AnyOfmicrosoftGraphRiskState)`

SetRiskState gets a reference to the given AnyOfmicrosoftGraphRiskState and assigns it to the RiskState field.

### SetRiskStateExplicitNull

`func (o *SignIn) SetRiskStateExplicitNull(b bool)`

SetRiskStateExplicitNull (un)sets RiskState to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskState value is set to nil even if false is passed
### GetRiskEventTypes

`func (o *SignIn) GetRiskEventTypes() []AnyOfmicrosoftGraphRiskEventType`

GetRiskEventTypes returns the RiskEventTypes field if non-nil, zero value otherwise.

### GetRiskEventTypesOk

`func (o *SignIn) GetRiskEventTypesOk() ([]AnyOfmicrosoftGraphRiskEventType, bool)`

GetRiskEventTypesOk returns a tuple with the RiskEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskEventTypes

`func (o *SignIn) HasRiskEventTypes() bool`

HasRiskEventTypes returns a boolean if a field has been set.

### SetRiskEventTypes

`func (o *SignIn) SetRiskEventTypes(v []AnyOfmicrosoftGraphRiskEventType)`

SetRiskEventTypes gets a reference to the given []AnyOfmicrosoftGraphRiskEventType and assigns it to the RiskEventTypes field.

### GetResourceDisplayName

`func (o *SignIn) GetResourceDisplayName() string`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *SignIn) GetResourceDisplayNameOk() (string, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceDisplayName

`func (o *SignIn) HasResourceDisplayName() bool`

HasResourceDisplayName returns a boolean if a field has been set.

### SetResourceDisplayName

`func (o *SignIn) SetResourceDisplayName(v string)`

SetResourceDisplayName gets a reference to the given string and assigns it to the ResourceDisplayName field.

### SetResourceDisplayNameExplicitNull

`func (o *SignIn) SetResourceDisplayNameExplicitNull(b bool)`

SetResourceDisplayNameExplicitNull (un)sets ResourceDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceDisplayName value is set to nil even if false is passed
### GetResourceId

`func (o *SignIn) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *SignIn) GetResourceIdOk() (string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceId

`func (o *SignIn) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceId

`func (o *SignIn) SetResourceId(v string)`

SetResourceId gets a reference to the given string and assigns it to the ResourceId field.

### SetResourceIdExplicitNull

`func (o *SignIn) SetResourceIdExplicitNull(b bool)`

SetResourceIdExplicitNull (un)sets ResourceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


