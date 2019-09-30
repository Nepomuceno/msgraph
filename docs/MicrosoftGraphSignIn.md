# MicrosoftGraphSignIn

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

## Methods

### GetId

`func (o *MicrosoftGraphSignIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSignIn) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphSignIn) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphSignIn) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphSignIn) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphSignIn) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphSignIn) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphSignIn) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetUserDisplayName

`func (o *MicrosoftGraphSignIn) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *MicrosoftGraphSignIn) GetUserDisplayNameOk() (string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserDisplayName

`func (o *MicrosoftGraphSignIn) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### SetUserDisplayName

`func (o *MicrosoftGraphSignIn) SetUserDisplayName(v string)`

SetUserDisplayName gets a reference to the given string and assigns it to the UserDisplayName field.

### SetUserDisplayNameExplicitNull

`func (o *MicrosoftGraphSignIn) SetUserDisplayNameExplicitNull(b bool)`

SetUserDisplayNameExplicitNull (un)sets UserDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserDisplayName value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *MicrosoftGraphSignIn) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphSignIn) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *MicrosoftGraphSignIn) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphSignIn) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphSignIn) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetUserId

`func (o *MicrosoftGraphSignIn) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphSignIn) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *MicrosoftGraphSignIn) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *MicrosoftGraphSignIn) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### GetAppId

`func (o *MicrosoftGraphSignIn) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MicrosoftGraphSignIn) GetAppIdOk() (string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppId

`func (o *MicrosoftGraphSignIn) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppId

`func (o *MicrosoftGraphSignIn) SetAppId(v string)`

SetAppId gets a reference to the given string and assigns it to the AppId field.

### SetAppIdExplicitNull

`func (o *MicrosoftGraphSignIn) SetAppIdExplicitNull(b bool)`

SetAppIdExplicitNull (un)sets AppId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppId value is set to nil even if false is passed
### GetAppDisplayName

`func (o *MicrosoftGraphSignIn) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *MicrosoftGraphSignIn) GetAppDisplayNameOk() (string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppDisplayName

`func (o *MicrosoftGraphSignIn) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayName

`func (o *MicrosoftGraphSignIn) SetAppDisplayName(v string)`

SetAppDisplayName gets a reference to the given string and assigns it to the AppDisplayName field.

### SetAppDisplayNameExplicitNull

`func (o *MicrosoftGraphSignIn) SetAppDisplayNameExplicitNull(b bool)`

SetAppDisplayNameExplicitNull (un)sets AppDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppDisplayName value is set to nil even if false is passed
### GetIpAddress

`func (o *MicrosoftGraphSignIn) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MicrosoftGraphSignIn) GetIpAddressOk() (string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpAddress

`func (o *MicrosoftGraphSignIn) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddress

`func (o *MicrosoftGraphSignIn) SetIpAddress(v string)`

SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.

### SetIpAddressExplicitNull

`func (o *MicrosoftGraphSignIn) SetIpAddressExplicitNull(b bool)`

SetIpAddressExplicitNull (un)sets IpAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IpAddress value is set to nil even if false is passed
### GetStatus

`func (o *MicrosoftGraphSignIn) GetStatus() AnyOfmicrosoftGraphSignInStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphSignIn) GetStatusOk() (AnyOfmicrosoftGraphSignInStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MicrosoftGraphSignIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MicrosoftGraphSignIn) SetStatus(v AnyOfmicrosoftGraphSignInStatus)`

SetStatus gets a reference to the given AnyOfmicrosoftGraphSignInStatus and assigns it to the Status field.

### SetStatusExplicitNull

`func (o *MicrosoftGraphSignIn) SetStatusExplicitNull(b bool)`

SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Status value is set to nil even if false is passed
### GetClientAppUsed

`func (o *MicrosoftGraphSignIn) GetClientAppUsed() string`

GetClientAppUsed returns the ClientAppUsed field if non-nil, zero value otherwise.

### GetClientAppUsedOk

`func (o *MicrosoftGraphSignIn) GetClientAppUsedOk() (string, bool)`

GetClientAppUsedOk returns a tuple with the ClientAppUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientAppUsed

`func (o *MicrosoftGraphSignIn) HasClientAppUsed() bool`

HasClientAppUsed returns a boolean if a field has been set.

### SetClientAppUsed

`func (o *MicrosoftGraphSignIn) SetClientAppUsed(v string)`

SetClientAppUsed gets a reference to the given string and assigns it to the ClientAppUsed field.

### SetClientAppUsedExplicitNull

`func (o *MicrosoftGraphSignIn) SetClientAppUsedExplicitNull(b bool)`

SetClientAppUsedExplicitNull (un)sets ClientAppUsed to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ClientAppUsed value is set to nil even if false is passed
### GetDeviceDetail

`func (o *MicrosoftGraphSignIn) GetDeviceDetail() AnyOfmicrosoftGraphDeviceDetail`

GetDeviceDetail returns the DeviceDetail field if non-nil, zero value otherwise.

### GetDeviceDetailOk

`func (o *MicrosoftGraphSignIn) GetDeviceDetailOk() (AnyOfmicrosoftGraphDeviceDetail, bool)`

GetDeviceDetailOk returns a tuple with the DeviceDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceDetail

`func (o *MicrosoftGraphSignIn) HasDeviceDetail() bool`

HasDeviceDetail returns a boolean if a field has been set.

### SetDeviceDetail

`func (o *MicrosoftGraphSignIn) SetDeviceDetail(v AnyOfmicrosoftGraphDeviceDetail)`

SetDeviceDetail gets a reference to the given AnyOfmicrosoftGraphDeviceDetail and assigns it to the DeviceDetail field.

### SetDeviceDetailExplicitNull

`func (o *MicrosoftGraphSignIn) SetDeviceDetailExplicitNull(b bool)`

SetDeviceDetailExplicitNull (un)sets DeviceDetail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceDetail value is set to nil even if false is passed
### GetLocation

`func (o *MicrosoftGraphSignIn) GetLocation() AnyOfmicrosoftGraphSignInLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphSignIn) GetLocationOk() (AnyOfmicrosoftGraphSignInLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocation

`func (o *MicrosoftGraphSignIn) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocation

`func (o *MicrosoftGraphSignIn) SetLocation(v AnyOfmicrosoftGraphSignInLocation)`

SetLocation gets a reference to the given AnyOfmicrosoftGraphSignInLocation and assigns it to the Location field.

### SetLocationExplicitNull

`func (o *MicrosoftGraphSignIn) SetLocationExplicitNull(b bool)`

SetLocationExplicitNull (un)sets Location to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Location value is set to nil even if false is passed
### GetCorrelationId

`func (o *MicrosoftGraphSignIn) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MicrosoftGraphSignIn) GetCorrelationIdOk() (string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCorrelationId

`func (o *MicrosoftGraphSignIn) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationId

`func (o *MicrosoftGraphSignIn) SetCorrelationId(v string)`

SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.

### SetCorrelationIdExplicitNull

`func (o *MicrosoftGraphSignIn) SetCorrelationIdExplicitNull(b bool)`

SetCorrelationIdExplicitNull (un)sets CorrelationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CorrelationId value is set to nil even if false is passed
### GetConditionalAccessStatus

`func (o *MicrosoftGraphSignIn) GetConditionalAccessStatus() AnyOfmicrosoftGraphConditionalAccessStatus`

GetConditionalAccessStatus returns the ConditionalAccessStatus field if non-nil, zero value otherwise.

### GetConditionalAccessStatusOk

`func (o *MicrosoftGraphSignIn) GetConditionalAccessStatusOk() (AnyOfmicrosoftGraphConditionalAccessStatus, bool)`

GetConditionalAccessStatusOk returns a tuple with the ConditionalAccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConditionalAccessStatus

`func (o *MicrosoftGraphSignIn) HasConditionalAccessStatus() bool`

HasConditionalAccessStatus returns a boolean if a field has been set.

### SetConditionalAccessStatus

`func (o *MicrosoftGraphSignIn) SetConditionalAccessStatus(v AnyOfmicrosoftGraphConditionalAccessStatus)`

SetConditionalAccessStatus gets a reference to the given AnyOfmicrosoftGraphConditionalAccessStatus and assigns it to the ConditionalAccessStatus field.

### SetConditionalAccessStatusExplicitNull

`func (o *MicrosoftGraphSignIn) SetConditionalAccessStatusExplicitNull(b bool)`

SetConditionalAccessStatusExplicitNull (un)sets ConditionalAccessStatus to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConditionalAccessStatus value is set to nil even if false is passed
### GetAppliedConditionalAccessPolicies

`func (o *MicrosoftGraphSignIn) GetAppliedConditionalAccessPolicies() []AnyOfmicrosoftGraphAppliedConditionalAccessPolicy`

GetAppliedConditionalAccessPolicies returns the AppliedConditionalAccessPolicies field if non-nil, zero value otherwise.

### GetAppliedConditionalAccessPoliciesOk

`func (o *MicrosoftGraphSignIn) GetAppliedConditionalAccessPoliciesOk() ([]AnyOfmicrosoftGraphAppliedConditionalAccessPolicy, bool)`

GetAppliedConditionalAccessPoliciesOk returns a tuple with the AppliedConditionalAccessPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppliedConditionalAccessPolicies

`func (o *MicrosoftGraphSignIn) HasAppliedConditionalAccessPolicies() bool`

HasAppliedConditionalAccessPolicies returns a boolean if a field has been set.

### SetAppliedConditionalAccessPolicies

`func (o *MicrosoftGraphSignIn) SetAppliedConditionalAccessPolicies(v []AnyOfmicrosoftGraphAppliedConditionalAccessPolicy)`

SetAppliedConditionalAccessPolicies gets a reference to the given []AnyOfmicrosoftGraphAppliedConditionalAccessPolicy and assigns it to the AppliedConditionalAccessPolicies field.

### GetIsInteractive

`func (o *MicrosoftGraphSignIn) GetIsInteractive() bool`

GetIsInteractive returns the IsInteractive field if non-nil, zero value otherwise.

### GetIsInteractiveOk

`func (o *MicrosoftGraphSignIn) GetIsInteractiveOk() (bool, bool)`

GetIsInteractiveOk returns a tuple with the IsInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsInteractive

`func (o *MicrosoftGraphSignIn) HasIsInteractive() bool`

HasIsInteractive returns a boolean if a field has been set.

### SetIsInteractive

`func (o *MicrosoftGraphSignIn) SetIsInteractive(v bool)`

SetIsInteractive gets a reference to the given bool and assigns it to the IsInteractive field.

### SetIsInteractiveExplicitNull

`func (o *MicrosoftGraphSignIn) SetIsInteractiveExplicitNull(b bool)`

SetIsInteractiveExplicitNull (un)sets IsInteractive to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsInteractive value is set to nil even if false is passed
### GetRiskDetail

`func (o *MicrosoftGraphSignIn) GetRiskDetail() AnyOfmicrosoftGraphRiskDetail`

GetRiskDetail returns the RiskDetail field if non-nil, zero value otherwise.

### GetRiskDetailOk

`func (o *MicrosoftGraphSignIn) GetRiskDetailOk() (AnyOfmicrosoftGraphRiskDetail, bool)`

GetRiskDetailOk returns a tuple with the RiskDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskDetail

`func (o *MicrosoftGraphSignIn) HasRiskDetail() bool`

HasRiskDetail returns a boolean if a field has been set.

### SetRiskDetail

`func (o *MicrosoftGraphSignIn) SetRiskDetail(v AnyOfmicrosoftGraphRiskDetail)`

SetRiskDetail gets a reference to the given AnyOfmicrosoftGraphRiskDetail and assigns it to the RiskDetail field.

### SetRiskDetailExplicitNull

`func (o *MicrosoftGraphSignIn) SetRiskDetailExplicitNull(b bool)`

SetRiskDetailExplicitNull (un)sets RiskDetail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskDetail value is set to nil even if false is passed
### GetRiskLevelAggregated

`func (o *MicrosoftGraphSignIn) GetRiskLevelAggregated() AnyOfmicrosoftGraphRiskLevel`

GetRiskLevelAggregated returns the RiskLevelAggregated field if non-nil, zero value otherwise.

### GetRiskLevelAggregatedOk

`func (o *MicrosoftGraphSignIn) GetRiskLevelAggregatedOk() (AnyOfmicrosoftGraphRiskLevel, bool)`

GetRiskLevelAggregatedOk returns a tuple with the RiskLevelAggregated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskLevelAggregated

`func (o *MicrosoftGraphSignIn) HasRiskLevelAggregated() bool`

HasRiskLevelAggregated returns a boolean if a field has been set.

### SetRiskLevelAggregated

`func (o *MicrosoftGraphSignIn) SetRiskLevelAggregated(v AnyOfmicrosoftGraphRiskLevel)`

SetRiskLevelAggregated gets a reference to the given AnyOfmicrosoftGraphRiskLevel and assigns it to the RiskLevelAggregated field.

### SetRiskLevelAggregatedExplicitNull

`func (o *MicrosoftGraphSignIn) SetRiskLevelAggregatedExplicitNull(b bool)`

SetRiskLevelAggregatedExplicitNull (un)sets RiskLevelAggregated to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskLevelAggregated value is set to nil even if false is passed
### GetRiskLevelDuringSignIn

`func (o *MicrosoftGraphSignIn) GetRiskLevelDuringSignIn() AnyOfmicrosoftGraphRiskLevel`

GetRiskLevelDuringSignIn returns the RiskLevelDuringSignIn field if non-nil, zero value otherwise.

### GetRiskLevelDuringSignInOk

`func (o *MicrosoftGraphSignIn) GetRiskLevelDuringSignInOk() (AnyOfmicrosoftGraphRiskLevel, bool)`

GetRiskLevelDuringSignInOk returns a tuple with the RiskLevelDuringSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskLevelDuringSignIn

`func (o *MicrosoftGraphSignIn) HasRiskLevelDuringSignIn() bool`

HasRiskLevelDuringSignIn returns a boolean if a field has been set.

### SetRiskLevelDuringSignIn

`func (o *MicrosoftGraphSignIn) SetRiskLevelDuringSignIn(v AnyOfmicrosoftGraphRiskLevel)`

SetRiskLevelDuringSignIn gets a reference to the given AnyOfmicrosoftGraphRiskLevel and assigns it to the RiskLevelDuringSignIn field.

### SetRiskLevelDuringSignInExplicitNull

`func (o *MicrosoftGraphSignIn) SetRiskLevelDuringSignInExplicitNull(b bool)`

SetRiskLevelDuringSignInExplicitNull (un)sets RiskLevelDuringSignIn to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskLevelDuringSignIn value is set to nil even if false is passed
### GetRiskState

`func (o *MicrosoftGraphSignIn) GetRiskState() AnyOfmicrosoftGraphRiskState`

GetRiskState returns the RiskState field if non-nil, zero value otherwise.

### GetRiskStateOk

`func (o *MicrosoftGraphSignIn) GetRiskStateOk() (AnyOfmicrosoftGraphRiskState, bool)`

GetRiskStateOk returns a tuple with the RiskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskState

`func (o *MicrosoftGraphSignIn) HasRiskState() bool`

HasRiskState returns a boolean if a field has been set.

### SetRiskState

`func (o *MicrosoftGraphSignIn) SetRiskState(v AnyOfmicrosoftGraphRiskState)`

SetRiskState gets a reference to the given AnyOfmicrosoftGraphRiskState and assigns it to the RiskState field.

### SetRiskStateExplicitNull

`func (o *MicrosoftGraphSignIn) SetRiskStateExplicitNull(b bool)`

SetRiskStateExplicitNull (un)sets RiskState to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RiskState value is set to nil even if false is passed
### GetRiskEventTypes

`func (o *MicrosoftGraphSignIn) GetRiskEventTypes() []AnyOfmicrosoftGraphRiskEventType`

GetRiskEventTypes returns the RiskEventTypes field if non-nil, zero value otherwise.

### GetRiskEventTypesOk

`func (o *MicrosoftGraphSignIn) GetRiskEventTypesOk() ([]AnyOfmicrosoftGraphRiskEventType, bool)`

GetRiskEventTypesOk returns a tuple with the RiskEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRiskEventTypes

`func (o *MicrosoftGraphSignIn) HasRiskEventTypes() bool`

HasRiskEventTypes returns a boolean if a field has been set.

### SetRiskEventTypes

`func (o *MicrosoftGraphSignIn) SetRiskEventTypes(v []AnyOfmicrosoftGraphRiskEventType)`

SetRiskEventTypes gets a reference to the given []AnyOfmicrosoftGraphRiskEventType and assigns it to the RiskEventTypes field.

### GetResourceDisplayName

`func (o *MicrosoftGraphSignIn) GetResourceDisplayName() string`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *MicrosoftGraphSignIn) GetResourceDisplayNameOk() (string, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceDisplayName

`func (o *MicrosoftGraphSignIn) HasResourceDisplayName() bool`

HasResourceDisplayName returns a boolean if a field has been set.

### SetResourceDisplayName

`func (o *MicrosoftGraphSignIn) SetResourceDisplayName(v string)`

SetResourceDisplayName gets a reference to the given string and assigns it to the ResourceDisplayName field.

### SetResourceDisplayNameExplicitNull

`func (o *MicrosoftGraphSignIn) SetResourceDisplayNameExplicitNull(b bool)`

SetResourceDisplayNameExplicitNull (un)sets ResourceDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceDisplayName value is set to nil even if false is passed
### GetResourceId

`func (o *MicrosoftGraphSignIn) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *MicrosoftGraphSignIn) GetResourceIdOk() (string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceId

`func (o *MicrosoftGraphSignIn) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceId

`func (o *MicrosoftGraphSignIn) SetResourceId(v string)`

SetResourceId gets a reference to the given string and assigns it to the ResourceId field.

### SetResourceIdExplicitNull

`func (o *MicrosoftGraphSignIn) SetResourceIdExplicitNull(b bool)`

SetResourceIdExplicitNull (un)sets ResourceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResourceId value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


