# MicrosoftGraphRestrictedSignIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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
**TargetTenantId** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphRestrictedSignIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphRestrictedSignIn) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphRestrictedSignIn) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphRestrictedSignIn) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphRestrictedSignIn) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphRestrictedSignIn) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphRestrictedSignIn) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphRestrictedSignIn) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetUserDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *MicrosoftGraphRestrictedSignIn) GetUserDisplayNameOk() (string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) SetUserDisplayName(v string)`

SetUserDisplayName gets a reference to the given string and assigns it to the UserDisplayName field.

### SetUserDisplayNameExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetUserDisplayNameExplicitNull(b bool)`

SetUserDisplayNameExplicitNull (un)sets UserDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserDisplayName value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *MicrosoftGraphRestrictedSignIn) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphRestrictedSignIn) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *MicrosoftGraphRestrictedSignIn) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphRestrictedSignIn) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetUserId

`func (o *MicrosoftGraphRestrictedSignIn) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphRestrictedSignIn) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *MicrosoftGraphRestrictedSignIn) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *MicrosoftGraphRestrictedSignIn) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### GetAppId

`func (o *MicrosoftGraphRestrictedSignIn) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MicrosoftGraphRestrictedSignIn) GetAppIdOk() (string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppId

`func (o *MicrosoftGraphRestrictedSignIn) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppId

`func (o *MicrosoftGraphRestrictedSignIn) SetAppId(v string)`

SetAppId gets a reference to the given string and assigns it to the AppId field.

### SetAppIdExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetAppIdExplicitNull(b bool)`

SetAppIdExplicitNull (un)sets AppId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppId value is set to nil even if false is passed
### GetAppDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *MicrosoftGraphRestrictedSignIn) GetAppDisplayNameOk() (string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) SetAppDisplayName(v string)`

SetAppDisplayName gets a reference to the given string and assigns it to the AppDisplayName field.

### SetAppDisplayNameExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetAppDisplayNameExplicitNull(b bool)`

SetAppDisplayNameExplicitNull (un)sets AppDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppDisplayName value is set to nil even if false is passed
### GetIpAddress

`func (o *MicrosoftGraphRestrictedSignIn) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MicrosoftGraphRestrictedSignIn) GetIpAddressOk() (string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpAddress

`func (o *MicrosoftGraphRestrictedSignIn) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddress

`func (o *MicrosoftGraphRestrictedSignIn) SetIpAddress(v string)`

SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.

### SetIpAddressExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetIpAddressExplicitNull(b bool)`

SetIpAddressExplicitNull (un)sets IpAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IpAddress value is set to nil even if false is passed
### GetStatus

`func (o *MicrosoftGraphRestrictedSignIn) GetStatus() AnyOfmicrosoftGraphSignInStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphRestrictedSignIn) GetStatusOk() (AnyOfmicrosoftGraphSignInStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphRestrictedSignIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphRestrictedSignIn) SetStatus(v AnyOfmicrosoftGraphSignInStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphSignInStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetClientAppUsed

`func (o *MicrosoftGraphRestrictedSignIn) GetClientAppUsed() string`

GetClientAppUsed returns the ClientAppUsed field if non-nil, zero value otherwise.

### GetClientAppUsedOk

`func (o *MicrosoftGraphRestrictedSignIn) GetClientAppUsedOk() (string, bool)`

GetClientAppUsedOk returns a tuple with the ClientAppUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientAppUsed

`func (o *MicrosoftGraphRestrictedSignIn) HasClientAppUsed() bool`

HasClientAppUsed returns a boolean if a field has been set.

### SetClientAppUsed

`func (o *MicrosoftGraphRestrictedSignIn) SetClientAppUsed(v string)`

SetClientAppUsed gets a reference to the given string and assigns it to the ClientAppUsed field.

### SetClientAppUsedExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetClientAppUsedExplicitNull(b bool)`

SetClientAppUsedExplicitNull (un)sets ClientAppUsed to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ClientAppUsed value is set to nil even if false is passed
### GetDeviceDetail

`func (o *MicrosoftGraphRestrictedSignIn) GetDeviceDetail() AnyOfmicrosoftGraphDeviceDetail`

GetDeviceDetail returns the DeviceDetail field if non-nil, zero value otherwise.

### GetDeviceDetailOk

`func (o *MicrosoftGraphRestrictedSignIn) GetDeviceDetailOk() (AnyOfmicrosoftGraphDeviceDetail, bool)`

GetDeviceDetailOk returns a tuple with the DeviceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceDetail

`func (o *MicrosoftGraphRestrictedSignIn) HasDeviceDetail() bool`

HasDeviceDetail returns a boolean if a field has been set.

### SetDeviceDetail

`func (o *MicrosoftGraphRestrictedSignIn) SetDeviceDetail(v AnyOfmicrosoftGraphDeviceDetail)`

SetDeviceDetail gets a reference to the given AnyOfmicrosoftGraphDeviceDetail and assigns it to the DeviceDetail field.

### SetDeviceDetailExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetDeviceDetailExplicitNull(b bool)`

SetDeviceDetailExplicitNull (un)sets DeviceDetail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceDetail value is set to nil even if false is passed
### GetLocation

`func (o *MicrosoftGraphRestrictedSignIn) GetLocation() AnyOfmicrosoftGraphSignInLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphRestrictedSignIn) GetLocationOk() (AnyOfmicrosoftGraphSignInLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocation

`func (o *MicrosoftGraphRestrictedSignIn) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocation

`func (o *MicrosoftGraphRestrictedSignIn) SetLocation(v AnyOfmicrosoftGraphSignInLocation)`

SetLocation gets a reference to the given AnyOfmicrosoftGraphSignInLocation and assigns it to the Location field.

### SetLocationExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetLocationExplicitNull(b bool)`

SetLocationExplicitNull (un)sets Location to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Location value is set to nil even if false is passed
### GetCorrelationId

`func (o *MicrosoftGraphRestrictedSignIn) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MicrosoftGraphRestrictedSignIn) GetCorrelationIdOk() (string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCorrelationId

`func (o *MicrosoftGraphRestrictedSignIn) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationId

`func (o *MicrosoftGraphRestrictedSignIn) SetCorrelationId(v string)`

SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.

### SetCorrelationIdExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetCorrelationIdExplicitNull(b bool)`

SetCorrelationIdExplicitNull (un)sets CorrelationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CorrelationId value is set to nil even if false is passed
### GetConditionalAccessStatus

`func (o *MicrosoftGraphRestrictedSignIn) GetConditionalAccessStatus() AnyOfmicrosoftGraphConditionalAccessStatus`

GetConditionalAccessStatus returns the ConditionalAccessStatus field if non-nil, zero value otherwise.

### GetConditionalAccessStatusOk

`func (o *MicrosoftGraphRestrictedSignIn) GetConditionalAccessStatusOk() (AnyOfmicrosoftGraphConditionalAccessStatus, bool)`

GetConditionalAccessStatusOk returns a tuple with the ConditionalAccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConditionalAccessStatus

`func (o *MicrosoftGraphRestrictedSignIn) HasConditionalAccessStatus() bool`

HasConditionalAccessStatus returns a boolean if a field has been set.

### SetConditionalAccessStatus

`func (o *MicrosoftGraphRestrictedSignIn) SetConditionalAccessStatus(v AnyOfmicrosoftGraphConditionalAccessStatus)`

SetConditionalAccessStatus gets a reference to the given AnyOfmicrosoftGraphConditionalAccessStatus and assigns it to the ConditionalAccessStatus field.

### SetConditionalAccessStatusExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetConditionalAccessStatusExplicitNull(b bool)`

SetConditionalAccessStatusExplicitNull (un)sets ConditionalAccessStatus to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConditionalAccessStatus value is set to nil even if false is passed
### GetAppliedConditionalAccessPolicies

`func (o *MicrosoftGraphRestrictedSignIn) GetAppliedConditionalAccessPolicies() []AnyOfmicrosoftGraphAppliedConditionalAccessPolicy`

GetAppliedConditionalAccessPolicies returns the AppliedConditionalAccessPolicies field if non-nil, zero value otherwise.

### GetAppliedConditionalAccessPoliciesOk

`func (o *MicrosoftGraphRestrictedSignIn) GetAppliedConditionalAccessPoliciesOk() ([]AnyOfmicrosoftGraphAppliedConditionalAccessPolicy, bool)`

GetAppliedConditionalAccessPoliciesOk returns a tuple with the AppliedConditionalAccessPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppliedConditionalAccessPolicies

`func (o *MicrosoftGraphRestrictedSignIn) HasAppliedConditionalAccessPolicies() bool`

HasAppliedConditionalAccessPolicies returns a boolean if a field has been set.

### SetAppliedConditionalAccessPolicies

`func (o *MicrosoftGraphRestrictedSignIn) SetAppliedConditionalAccessPolicies(v []AnyOfmicrosoftGraphAppliedConditionalAccessPolicy)`

SetAppliedConditionalAccessPolicies gets a reference to the given []AnyOfmicrosoftGraphAppliedConditionalAccessPolicy and assigns it to the AppliedConditionalAccessPolicies field.

### GetIsInteractive

`func (o *MicrosoftGraphRestrictedSignIn) GetIsInteractive() bool`

GetIsInteractive returns the IsInteractive field if non-nil, zero value otherwise.

### GetIsInteractiveOk

`func (o *MicrosoftGraphRestrictedSignIn) GetIsInteractiveOk() (bool, bool)`

GetIsInteractiveOk returns a tuple with the IsInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsInteractive

`func (o *MicrosoftGraphRestrictedSignIn) HasIsInteractive() bool`

HasIsInteractive returns a boolean if a field has been set.

### SetIsInteractive

`func (o *MicrosoftGraphRestrictedSignIn) SetIsInteractive(v bool)`

SetIsInteractive gets a reference to the given bool and assigns it to the IsInteractive field.

### SetIsInteractiveExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetIsInteractiveExplicitNull(b bool)`

SetIsInteractiveExplicitNull (un)sets IsInteractive to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsInteractive value is set to nil even if false is passed
### GetRiskDetail

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskDetail() AnyOfmicrosoftGraphRiskDetail`

GetRiskDetail returns the RiskDetail field if non-nil, zero value otherwise.

### GetRiskDetailOk

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskDetailOk() (AnyOfmicrosoftGraphRiskDetail, bool)`

GetRiskDetailOk returns a tuple with the RiskDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskDetail

`func (o *MicrosoftGraphRestrictedSignIn) HasRiskDetail() bool`

HasRiskDetail returns a boolean if a field has been set.

### SetRiskDetail

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskDetail(v AnyOfmicrosoftGraphRiskDetail)`

SetRiskDetail gets a reference to the given AnyOfmicrosoftGraphRiskDetail and assigns it to the RiskDetail field.

### SetRiskDetailExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskDetailExplicitNull(b bool)`

SetRiskDetailExplicitNull (un)sets RiskDetail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskDetail value is set to nil even if false is passed
### GetRiskLevelAggregated

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskLevelAggregated() AnyOfmicrosoftGraphRiskLevel`

GetRiskLevelAggregated returns the RiskLevelAggregated field if non-nil, zero value otherwise.

### GetRiskLevelAggregatedOk

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskLevelAggregatedOk() (AnyOfmicrosoftGraphRiskLevel, bool)`

GetRiskLevelAggregatedOk returns a tuple with the RiskLevelAggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskLevelAggregated

`func (o *MicrosoftGraphRestrictedSignIn) HasRiskLevelAggregated() bool`

HasRiskLevelAggregated returns a boolean if a field has been set.

### SetRiskLevelAggregated

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskLevelAggregated(v AnyOfmicrosoftGraphRiskLevel)`

SetRiskLevelAggregated gets a reference to the given AnyOfmicrosoftGraphRiskLevel and assigns it to the RiskLevelAggregated field.

### SetRiskLevelAggregatedExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskLevelAggregatedExplicitNull(b bool)`

SetRiskLevelAggregatedExplicitNull (un)sets RiskLevelAggregated to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskLevelAggregated value is set to nil even if false is passed
### GetRiskLevelDuringSignIn

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskLevelDuringSignIn() AnyOfmicrosoftGraphRiskLevel`

GetRiskLevelDuringSignIn returns the RiskLevelDuringSignIn field if non-nil, zero value otherwise.

### GetRiskLevelDuringSignInOk

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskLevelDuringSignInOk() (AnyOfmicrosoftGraphRiskLevel, bool)`

GetRiskLevelDuringSignInOk returns a tuple with the RiskLevelDuringSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskLevelDuringSignIn

`func (o *MicrosoftGraphRestrictedSignIn) HasRiskLevelDuringSignIn() bool`

HasRiskLevelDuringSignIn returns a boolean if a field has been set.

### SetRiskLevelDuringSignIn

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskLevelDuringSignIn(v AnyOfmicrosoftGraphRiskLevel)`

SetRiskLevelDuringSignIn gets a reference to the given AnyOfmicrosoftGraphRiskLevel and assigns it to the RiskLevelDuringSignIn field.

### SetRiskLevelDuringSignInExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskLevelDuringSignInExplicitNull(b bool)`

SetRiskLevelDuringSignInExplicitNull (un)sets RiskLevelDuringSignIn to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskLevelDuringSignIn value is set to nil even if false is passed
### GetRiskState

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskState() AnyOfmicrosoftGraphRiskState`

GetRiskState returns the RiskState field if non-nil, zero value otherwise.

### GetRiskStateOk

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskStateOk() (AnyOfmicrosoftGraphRiskState, bool)`

GetRiskStateOk returns a tuple with the RiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskState

`func (o *MicrosoftGraphRestrictedSignIn) HasRiskState() bool`

HasRiskState returns a boolean if a field has been set.

### SetRiskState

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskState(v AnyOfmicrosoftGraphRiskState)`

SetRiskState gets a reference to the given AnyOfmicrosoftGraphRiskState and assigns it to the RiskState field.

### SetRiskStateExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskStateExplicitNull(b bool)`

SetRiskStateExplicitNull (un)sets RiskState to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskState value is set to nil even if false is passed
### GetRiskEventTypes

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskEventTypes() []AnyOfmicrosoftGraphRiskEventType`

GetRiskEventTypes returns the RiskEventTypes field if non-nil, zero value otherwise.

### GetRiskEventTypesOk

`func (o *MicrosoftGraphRestrictedSignIn) GetRiskEventTypesOk() ([]AnyOfmicrosoftGraphRiskEventType, bool)`

GetRiskEventTypesOk returns a tuple with the RiskEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskEventTypes

`func (o *MicrosoftGraphRestrictedSignIn) HasRiskEventTypes() bool`

HasRiskEventTypes returns a boolean if a field has been set.

### SetRiskEventTypes

`func (o *MicrosoftGraphRestrictedSignIn) SetRiskEventTypes(v []AnyOfmicrosoftGraphRiskEventType)`

SetRiskEventTypes gets a reference to the given []AnyOfmicrosoftGraphRiskEventType and assigns it to the RiskEventTypes field.

### GetResourceDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) GetResourceDisplayName() string`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *MicrosoftGraphRestrictedSignIn) GetResourceDisplayNameOk() (string, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) HasResourceDisplayName() bool`

HasResourceDisplayName returns a boolean if a field has been set.

### SetResourceDisplayName

`func (o *MicrosoftGraphRestrictedSignIn) SetResourceDisplayName(v string)`

SetResourceDisplayName gets a reference to the given string and assigns it to the ResourceDisplayName field.

### SetResourceDisplayNameExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetResourceDisplayNameExplicitNull(b bool)`

SetResourceDisplayNameExplicitNull (un)sets ResourceDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceDisplayName value is set to nil even if false is passed
### GetResourceId

`func (o *MicrosoftGraphRestrictedSignIn) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *MicrosoftGraphRestrictedSignIn) GetResourceIdOk() (string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceId

`func (o *MicrosoftGraphRestrictedSignIn) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceId

`func (o *MicrosoftGraphRestrictedSignIn) SetResourceId(v string)`

SetResourceId gets a reference to the given string and assigns it to the ResourceId field.

### SetResourceIdExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetResourceIdExplicitNull(b bool)`

SetResourceIdExplicitNull (un)sets ResourceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceId value is set to nil even if false is passed
### GetTargetTenantId

`func (o *MicrosoftGraphRestrictedSignIn) GetTargetTenantId() string`

GetTargetTenantId returns the TargetTenantId field if non-nil, zero value otherwise.

### GetTargetTenantIdOk

`func (o *MicrosoftGraphRestrictedSignIn) GetTargetTenantIdOk() (string, bool)`

GetTargetTenantIdOk returns a tuple with the TargetTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetTenantId

`func (o *MicrosoftGraphRestrictedSignIn) HasTargetTenantId() bool`

HasTargetTenantId returns a boolean if a field has been set.

### SetTargetTenantId

`func (o *MicrosoftGraphRestrictedSignIn) SetTargetTenantId(v string)`

SetTargetTenantId gets a reference to the given string and assigns it to the TargetTenantId field.

### SetTargetTenantIdExplicitNull

`func (o *MicrosoftGraphRestrictedSignIn) SetTargetTenantIdExplicitNull(b bool)`

SetTargetTenantIdExplicitNull (un)sets TargetTenantId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TargetTenantId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


