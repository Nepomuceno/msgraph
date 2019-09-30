# SecureScoreControlProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **string** |  | [optional] 
**ActionUrl** | Pointer to **string** |  | [optional] 
**AzureTenantId** | Pointer to **string** |  | [optional] 
**ComplianceInformation** | Pointer to [**[]AnyOfmicrosoftGraphComplianceInformation**](anyOf&lt;microsoft.graph.complianceInformation&gt;.md) |  | [optional] 
**ControlCategory** | Pointer to **string** |  | [optional] 
**ControlStateUpdates** | Pointer to [**[]AnyOfmicrosoftGraphSecureScoreControlStateUpdate**](anyOf&lt;microsoft.graph.secureScoreControlStateUpdate&gt;.md) |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**ImplementationCost** | Pointer to **string** |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**MaxScore** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 
**Remediation** | Pointer to **string** |  | [optional] 
**RemediationImpact** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Threats** | Pointer to **[]string** |  | [optional] 
**Tier** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UserImpact** | Pointer to **string** |  | [optional] 
**VendorInformation** | Pointer to [**AnyOfmicrosoftGraphSecurityVendorInformation**](anyOf&lt;microsoft.graph.securityVendorInformation&gt;.md) |  | [optional] 

## Methods

### GetActionType

`func (o *SecureScoreControlProfile) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *SecureScoreControlProfile) GetActionTypeOk() (string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionType

`func (o *SecureScoreControlProfile) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### SetActionType

`func (o *SecureScoreControlProfile) SetActionType(v string)`

SetActionType gets a reference to the given string and assigns it to the ActionType field.

### SetActionTypeExplicitNull

`func (o *SecureScoreControlProfile) SetActionTypeExplicitNull(b bool)`

SetActionTypeExplicitNull (un)sets ActionType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActionType value is set to nil even if false is passed
### GetActionUrl

`func (o *SecureScoreControlProfile) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *SecureScoreControlProfile) GetActionUrlOk() (string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionUrl

`func (o *SecureScoreControlProfile) HasActionUrl() bool`

HasActionUrl returns a boolean if a field has been set.

### SetActionUrl

`func (o *SecureScoreControlProfile) SetActionUrl(v string)`

SetActionUrl gets a reference to the given string and assigns it to the ActionUrl field.

### SetActionUrlExplicitNull

`func (o *SecureScoreControlProfile) SetActionUrlExplicitNull(b bool)`

SetActionUrlExplicitNull (un)sets ActionUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActionUrl value is set to nil even if false is passed
### GetAzureTenantId

`func (o *SecureScoreControlProfile) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *SecureScoreControlProfile) GetAzureTenantIdOk() (string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureTenantId

`func (o *SecureScoreControlProfile) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### SetAzureTenantId

`func (o *SecureScoreControlProfile) SetAzureTenantId(v string)`

SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.

### GetComplianceInformation

`func (o *SecureScoreControlProfile) GetComplianceInformation() []AnyOfmicrosoftGraphComplianceInformation`

GetComplianceInformation returns the ComplianceInformation field if non-nil, zero value otherwise.

### GetComplianceInformationOk

`func (o *SecureScoreControlProfile) GetComplianceInformationOk() ([]AnyOfmicrosoftGraphComplianceInformation, bool)`

GetComplianceInformationOk returns a tuple with the ComplianceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplianceInformation

`func (o *SecureScoreControlProfile) HasComplianceInformation() bool`

HasComplianceInformation returns a boolean if a field has been set.

### SetComplianceInformation

`func (o *SecureScoreControlProfile) SetComplianceInformation(v []AnyOfmicrosoftGraphComplianceInformation)`

SetComplianceInformation gets a reference to the given []AnyOfmicrosoftGraphComplianceInformation and assigns it to the ComplianceInformation field.

### GetControlCategory

`func (o *SecureScoreControlProfile) GetControlCategory() string`

GetControlCategory returns the ControlCategory field if non-nil, zero value otherwise.

### GetControlCategoryOk

`func (o *SecureScoreControlProfile) GetControlCategoryOk() (string, bool)`

GetControlCategoryOk returns a tuple with the ControlCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasControlCategory

`func (o *SecureScoreControlProfile) HasControlCategory() bool`

HasControlCategory returns a boolean if a field has been set.

### SetControlCategory

`func (o *SecureScoreControlProfile) SetControlCategory(v string)`

SetControlCategory gets a reference to the given string and assigns it to the ControlCategory field.

### SetControlCategoryExplicitNull

`func (o *SecureScoreControlProfile) SetControlCategoryExplicitNull(b bool)`

SetControlCategoryExplicitNull (un)sets ControlCategory to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ControlCategory value is set to nil even if false is passed
### GetControlStateUpdates

`func (o *SecureScoreControlProfile) GetControlStateUpdates() []AnyOfmicrosoftGraphSecureScoreControlStateUpdate`

GetControlStateUpdates returns the ControlStateUpdates field if non-nil, zero value otherwise.

### GetControlStateUpdatesOk

`func (o *SecureScoreControlProfile) GetControlStateUpdatesOk() ([]AnyOfmicrosoftGraphSecureScoreControlStateUpdate, bool)`

GetControlStateUpdatesOk returns a tuple with the ControlStateUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasControlStateUpdates

`func (o *SecureScoreControlProfile) HasControlStateUpdates() bool`

HasControlStateUpdates returns a boolean if a field has been set.

### SetControlStateUpdates

`func (o *SecureScoreControlProfile) SetControlStateUpdates(v []AnyOfmicrosoftGraphSecureScoreControlStateUpdate)`

SetControlStateUpdates gets a reference to the given []AnyOfmicrosoftGraphSecureScoreControlStateUpdate and assigns it to the ControlStateUpdates field.

### GetDeprecated

`func (o *SecureScoreControlProfile) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *SecureScoreControlProfile) GetDeprecatedOk() (bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeprecated

`func (o *SecureScoreControlProfile) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### SetDeprecated

`func (o *SecureScoreControlProfile) SetDeprecated(v bool)`

SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.

### SetDeprecatedExplicitNull

`func (o *SecureScoreControlProfile) SetDeprecatedExplicitNull(b bool)`

SetDeprecatedExplicitNull (un)sets Deprecated to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Deprecated value is set to nil even if false is passed
### GetImplementationCost

`func (o *SecureScoreControlProfile) GetImplementationCost() string`

GetImplementationCost returns the ImplementationCost field if non-nil, zero value otherwise.

### GetImplementationCostOk

`func (o *SecureScoreControlProfile) GetImplementationCostOk() (string, bool)`

GetImplementationCostOk returns a tuple with the ImplementationCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImplementationCost

`func (o *SecureScoreControlProfile) HasImplementationCost() bool`

HasImplementationCost returns a boolean if a field has been set.

### SetImplementationCost

`func (o *SecureScoreControlProfile) SetImplementationCost(v string)`

SetImplementationCost gets a reference to the given string and assigns it to the ImplementationCost field.

### SetImplementationCostExplicitNull

`func (o *SecureScoreControlProfile) SetImplementationCostExplicitNull(b bool)`

SetImplementationCostExplicitNull (un)sets ImplementationCost to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ImplementationCost value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *SecureScoreControlProfile) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *SecureScoreControlProfile) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *SecureScoreControlProfile) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *SecureScoreControlProfile) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *SecureScoreControlProfile) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetMaxScore

`func (o *SecureScoreControlProfile) GetMaxScore() AnyOfnumberstringstring`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *SecureScoreControlProfile) GetMaxScoreOk() (AnyOfnumberstringstring, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxScore

`func (o *SecureScoreControlProfile) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### SetMaxScore

`func (o *SecureScoreControlProfile) SetMaxScore(v AnyOfnumberstringstring)`

SetMaxScore gets a reference to the given AnyOfnumberstringstring and assigns it to the MaxScore field.

### SetMaxScoreExplicitNull

`func (o *SecureScoreControlProfile) SetMaxScoreExplicitNull(b bool)`

SetMaxScoreExplicitNull (un)sets MaxScore to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MaxScore value is set to nil even if false is passed
### GetRank

`func (o *SecureScoreControlProfile) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *SecureScoreControlProfile) GetRankOk() (int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRank

`func (o *SecureScoreControlProfile) HasRank() bool`

HasRank returns a boolean if a field has been set.

### SetRank

`func (o *SecureScoreControlProfile) SetRank(v int32)`

SetRank gets a reference to the given int32 and assigns it to the Rank field.

### SetRankExplicitNull

`func (o *SecureScoreControlProfile) SetRankExplicitNull(b bool)`

SetRankExplicitNull (un)sets Rank to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Rank value is set to nil even if false is passed
### GetRemediation

`func (o *SecureScoreControlProfile) GetRemediation() string`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *SecureScoreControlProfile) GetRemediationOk() (string, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemediation

`func (o *SecureScoreControlProfile) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### SetRemediation

`func (o *SecureScoreControlProfile) SetRemediation(v string)`

SetRemediation gets a reference to the given string and assigns it to the Remediation field.

### SetRemediationExplicitNull

`func (o *SecureScoreControlProfile) SetRemediationExplicitNull(b bool)`

SetRemediationExplicitNull (un)sets Remediation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Remediation value is set to nil even if false is passed
### GetRemediationImpact

`func (o *SecureScoreControlProfile) GetRemediationImpact() string`

GetRemediationImpact returns the RemediationImpact field if non-nil, zero value otherwise.

### GetRemediationImpactOk

`func (o *SecureScoreControlProfile) GetRemediationImpactOk() (string, bool)`

GetRemediationImpactOk returns a tuple with the RemediationImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemediationImpact

`func (o *SecureScoreControlProfile) HasRemediationImpact() bool`

HasRemediationImpact returns a boolean if a field has been set.

### SetRemediationImpact

`func (o *SecureScoreControlProfile) SetRemediationImpact(v string)`

SetRemediationImpact gets a reference to the given string and assigns it to the RemediationImpact field.

### SetRemediationImpactExplicitNull

`func (o *SecureScoreControlProfile) SetRemediationImpactExplicitNull(b bool)`

SetRemediationImpactExplicitNull (un)sets RemediationImpact to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RemediationImpact value is set to nil even if false is passed
### GetService

`func (o *SecureScoreControlProfile) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SecureScoreControlProfile) GetServiceOk() (string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasService

`func (o *SecureScoreControlProfile) HasService() bool`

HasService returns a boolean if a field has been set.

### SetService

`func (o *SecureScoreControlProfile) SetService(v string)`

SetService gets a reference to the given string and assigns it to the Service field.

### SetServiceExplicitNull

`func (o *SecureScoreControlProfile) SetServiceExplicitNull(b bool)`

SetServiceExplicitNull (un)sets Service to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Service value is set to nil even if false is passed
### GetThreats

`func (o *SecureScoreControlProfile) GetThreats() []string`

GetThreats returns the Threats field if non-nil, zero value otherwise.

### GetThreatsOk

`func (o *SecureScoreControlProfile) GetThreatsOk() ([]string, bool)`

GetThreatsOk returns a tuple with the Threats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThreats

`func (o *SecureScoreControlProfile) HasThreats() bool`

HasThreats returns a boolean if a field has been set.

### SetThreats

`func (o *SecureScoreControlProfile) SetThreats(v []string)`

SetThreats gets a reference to the given []string and assigns it to the Threats field.

### GetTier

`func (o *SecureScoreControlProfile) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SecureScoreControlProfile) GetTierOk() (string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTier

`func (o *SecureScoreControlProfile) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTier

`func (o *SecureScoreControlProfile) SetTier(v string)`

SetTier gets a reference to the given string and assigns it to the Tier field.

### SetTierExplicitNull

`func (o *SecureScoreControlProfile) SetTierExplicitNull(b bool)`

SetTierExplicitNull (un)sets Tier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Tier value is set to nil even if false is passed
### GetTitle

`func (o *SecureScoreControlProfile) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SecureScoreControlProfile) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *SecureScoreControlProfile) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *SecureScoreControlProfile) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### SetTitleExplicitNull

`func (o *SecureScoreControlProfile) SetTitleExplicitNull(b bool)`

SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Title value is set to nil even if false is passed
### GetUserImpact

`func (o *SecureScoreControlProfile) GetUserImpact() string`

GetUserImpact returns the UserImpact field if non-nil, zero value otherwise.

### GetUserImpactOk

`func (o *SecureScoreControlProfile) GetUserImpactOk() (string, bool)`

GetUserImpactOk returns a tuple with the UserImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserImpact

`func (o *SecureScoreControlProfile) HasUserImpact() bool`

HasUserImpact returns a boolean if a field has been set.

### SetUserImpact

`func (o *SecureScoreControlProfile) SetUserImpact(v string)`

SetUserImpact gets a reference to the given string and assigns it to the UserImpact field.

### SetUserImpactExplicitNull

`func (o *SecureScoreControlProfile) SetUserImpactExplicitNull(b bool)`

SetUserImpactExplicitNull (un)sets UserImpact to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserImpact value is set to nil even if false is passed
### GetVendorInformation

`func (o *SecureScoreControlProfile) GetVendorInformation() AnyOfmicrosoftGraphSecurityVendorInformation`

GetVendorInformation returns the VendorInformation field if non-nil, zero value otherwise.

### GetVendorInformationOk

`func (o *SecureScoreControlProfile) GetVendorInformationOk() (AnyOfmicrosoftGraphSecurityVendorInformation, bool)`

GetVendorInformationOk returns a tuple with the VendorInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVendorInformation

`func (o *SecureScoreControlProfile) HasVendorInformation() bool`

HasVendorInformation returns a boolean if a field has been set.

### SetVendorInformation

`func (o *SecureScoreControlProfile) SetVendorInformation(v AnyOfmicrosoftGraphSecurityVendorInformation)`

SetVendorInformation gets a reference to the given AnyOfmicrosoftGraphSecurityVendorInformation and assigns it to the VendorInformation field.

### SetVendorInformationExplicitNull

`func (o *SecureScoreControlProfile) SetVendorInformationExplicitNull(b bool)`

SetVendorInformationExplicitNull (un)sets VendorInformation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VendorInformation value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


