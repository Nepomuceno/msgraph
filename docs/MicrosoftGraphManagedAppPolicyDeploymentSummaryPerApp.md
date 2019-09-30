# MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileAppIdentifier** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Deployment of an app. | [optional] 
**ConfigurationAppliedUserCount** | Pointer to **int32** | Number of users the policy is applied. | [optional] 

## Methods

### GetMobileAppIdentifier

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) GetMobileAppIdentifier() AnyOfobject`

GetMobileAppIdentifier returns the MobileAppIdentifier field if non-nil, zero value otherwise.

### GetMobileAppIdentifierOk

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) GetMobileAppIdentifierOk() (AnyOfobject, bool)`

GetMobileAppIdentifierOk returns a tuple with the MobileAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobileAppIdentifier

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) HasMobileAppIdentifier() bool`

HasMobileAppIdentifier returns a boolean if a field has been set.

### SetMobileAppIdentifier

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) SetMobileAppIdentifier(v AnyOfobject)`

SetMobileAppIdentifier gets a reference to the given AnyOfobject and assigns it to the MobileAppIdentifier field.

### SetMobileAppIdentifierExplicitNull

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) SetMobileAppIdentifierExplicitNull(b bool)`

SetMobileAppIdentifierExplicitNull (un)sets MobileAppIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MobileAppIdentifier value is set to nil even if false is passed
### GetConfigurationAppliedUserCount

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) GetConfigurationAppliedUserCount() int32`

GetConfigurationAppliedUserCount returns the ConfigurationAppliedUserCount field if non-nil, zero value otherwise.

### GetConfigurationAppliedUserCountOk

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) GetConfigurationAppliedUserCountOk() (int32, bool)`

GetConfigurationAppliedUserCountOk returns a tuple with the ConfigurationAppliedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfigurationAppliedUserCount

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) HasConfigurationAppliedUserCount() bool`

HasConfigurationAppliedUserCount returns a boolean if a field has been set.

### SetConfigurationAppliedUserCount

`func (o *MicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp) SetConfigurationAppliedUserCount(v int32)`

SetConfigurationAppliedUserCount gets a reference to the given int32 and assigns it to the ConfigurationAppliedUserCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


