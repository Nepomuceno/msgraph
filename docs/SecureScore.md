# SecureScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveUserCount** | Pointer to **int32** |  | [optional] 
**AverageComparativeScores** | Pointer to [**[]AnyOfmicrosoftGraphAverageComparativeScore**](anyOf&lt;microsoft.graph.averageComparativeScore&gt;.md) |  | [optional] 
**AzureTenantId** | Pointer to **string** |  | [optional] 
**ControlScores** | Pointer to [**[]AnyOfmicrosoftGraphControlScore**](anyOf&lt;microsoft.graph.controlScore&gt;.md) |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**CurrentScore** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**EnabledServices** | Pointer to **[]string** |  | [optional] 
**LicensedUserCount** | Pointer to **int32** |  | [optional] 
**MaxScore** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**VendorInformation** | Pointer to [**AnyOfmicrosoftGraphSecurityVendorInformation**](anyOf&lt;microsoft.graph.securityVendorInformation&gt;.md) |  | [optional] 

## Methods

### GetActiveUserCount

`func (o *SecureScore) GetActiveUserCount() int32`

GetActiveUserCount returns the ActiveUserCount field if non-nil, zero value otherwise.

### GetActiveUserCountOk

`func (o *SecureScore) GetActiveUserCountOk() (int32, bool)`

GetActiveUserCountOk returns a tuple with the ActiveUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveUserCount

`func (o *SecureScore) HasActiveUserCount() bool`

HasActiveUserCount returns a boolean if a field has been set.

### SetActiveUserCount

`func (o *SecureScore) SetActiveUserCount(v int32)`

SetActiveUserCount gets a reference to the given int32 and assigns it to the ActiveUserCount field.

### SetActiveUserCountExplicitNull

`func (o *SecureScore) SetActiveUserCountExplicitNull(b bool)`

SetActiveUserCountExplicitNull (un)sets ActiveUserCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActiveUserCount value is set to nil even if false is passed
### GetAverageComparativeScores

`func (o *SecureScore) GetAverageComparativeScores() []AnyOfmicrosoftGraphAverageComparativeScore`

GetAverageComparativeScores returns the AverageComparativeScores field if non-nil, zero value otherwise.

### GetAverageComparativeScoresOk

`func (o *SecureScore) GetAverageComparativeScoresOk() ([]AnyOfmicrosoftGraphAverageComparativeScore, bool)`

GetAverageComparativeScoresOk returns a tuple with the AverageComparativeScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAverageComparativeScores

`func (o *SecureScore) HasAverageComparativeScores() bool`

HasAverageComparativeScores returns a boolean if a field has been set.

### SetAverageComparativeScores

`func (o *SecureScore) SetAverageComparativeScores(v []AnyOfmicrosoftGraphAverageComparativeScore)`

SetAverageComparativeScores gets a reference to the given []AnyOfmicrosoftGraphAverageComparativeScore and assigns it to the AverageComparativeScores field.

### GetAzureTenantId

`func (o *SecureScore) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *SecureScore) GetAzureTenantIdOk() (string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureTenantId

`func (o *SecureScore) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### SetAzureTenantId

`func (o *SecureScore) SetAzureTenantId(v string)`

SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.

### GetControlScores

`func (o *SecureScore) GetControlScores() []AnyOfmicrosoftGraphControlScore`

GetControlScores returns the ControlScores field if non-nil, zero value otherwise.

### GetControlScoresOk

`func (o *SecureScore) GetControlScoresOk() ([]AnyOfmicrosoftGraphControlScore, bool)`

GetControlScoresOk returns a tuple with the ControlScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasControlScores

`func (o *SecureScore) HasControlScores() bool`

HasControlScores returns a boolean if a field has been set.

### SetControlScores

`func (o *SecureScore) SetControlScores(v []AnyOfmicrosoftGraphControlScore)`

SetControlScores gets a reference to the given []AnyOfmicrosoftGraphControlScore and assigns it to the ControlScores field.

### GetCreatedDateTime

`func (o *SecureScore) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *SecureScore) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *SecureScore) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *SecureScore) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *SecureScore) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetCurrentScore

`func (o *SecureScore) GetCurrentScore() AnyOfnumberstringstring`

GetCurrentScore returns the CurrentScore field if non-nil, zero value otherwise.

### GetCurrentScoreOk

`func (o *SecureScore) GetCurrentScoreOk() (AnyOfnumberstringstring, bool)`

GetCurrentScoreOk returns a tuple with the CurrentScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentScore

`func (o *SecureScore) HasCurrentScore() bool`

HasCurrentScore returns a boolean if a field has been set.

### SetCurrentScore

`func (o *SecureScore) SetCurrentScore(v AnyOfnumberstringstring)`

SetCurrentScore gets a reference to the given AnyOfnumberstringstring and assigns it to the CurrentScore field.

### SetCurrentScoreExplicitNull

`func (o *SecureScore) SetCurrentScoreExplicitNull(b bool)`

SetCurrentScoreExplicitNull (un)sets CurrentScore to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CurrentScore value is set to nil even if false is passed
### GetEnabledServices

`func (o *SecureScore) GetEnabledServices() []string`

GetEnabledServices returns the EnabledServices field if non-nil, zero value otherwise.

### GetEnabledServicesOk

`func (o *SecureScore) GetEnabledServicesOk() ([]string, bool)`

GetEnabledServicesOk returns a tuple with the EnabledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabledServices

`func (o *SecureScore) HasEnabledServices() bool`

HasEnabledServices returns a boolean if a field has been set.

### SetEnabledServices

`func (o *SecureScore) SetEnabledServices(v []string)`

SetEnabledServices gets a reference to the given []string and assigns it to the EnabledServices field.

### GetLicensedUserCount

`func (o *SecureScore) GetLicensedUserCount() int32`

GetLicensedUserCount returns the LicensedUserCount field if non-nil, zero value otherwise.

### GetLicensedUserCountOk

`func (o *SecureScore) GetLicensedUserCountOk() (int32, bool)`

GetLicensedUserCountOk returns a tuple with the LicensedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLicensedUserCount

`func (o *SecureScore) HasLicensedUserCount() bool`

HasLicensedUserCount returns a boolean if a field has been set.

### SetLicensedUserCount

`func (o *SecureScore) SetLicensedUserCount(v int32)`

SetLicensedUserCount gets a reference to the given int32 and assigns it to the LicensedUserCount field.

### SetLicensedUserCountExplicitNull

`func (o *SecureScore) SetLicensedUserCountExplicitNull(b bool)`

SetLicensedUserCountExplicitNull (un)sets LicensedUserCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LicensedUserCount value is set to nil even if false is passed
### GetMaxScore

`func (o *SecureScore) GetMaxScore() AnyOfnumberstringstring`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *SecureScore) GetMaxScoreOk() (AnyOfnumberstringstring, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxScore

`func (o *SecureScore) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### SetMaxScore

`func (o *SecureScore) SetMaxScore(v AnyOfnumberstringstring)`

SetMaxScore gets a reference to the given AnyOfnumberstringstring and assigns it to the MaxScore field.

### SetMaxScoreExplicitNull

`func (o *SecureScore) SetMaxScoreExplicitNull(b bool)`

SetMaxScoreExplicitNull (un)sets MaxScore to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MaxScore value is set to nil even if false is passed
### GetVendorInformation

`func (o *SecureScore) GetVendorInformation() AnyOfmicrosoftGraphSecurityVendorInformation`

GetVendorInformation returns the VendorInformation field if non-nil, zero value otherwise.

### GetVendorInformationOk

`func (o *SecureScore) GetVendorInformationOk() (AnyOfmicrosoftGraphSecurityVendorInformation, bool)`

GetVendorInformationOk returns a tuple with the VendorInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVendorInformation

`func (o *SecureScore) HasVendorInformation() bool`

HasVendorInformation returns a boolean if a field has been set.

### SetVendorInformation

`func (o *SecureScore) SetVendorInformation(v AnyOfmicrosoftGraphSecurityVendorInformation)`

SetVendorInformation gets a reference to the given AnyOfmicrosoftGraphSecurityVendorInformation and assigns it to the VendorInformation field.

### SetVendorInformationExplicitNull

`func (o *SecureScore) SetVendorInformationExplicitNull(b bool)`

SetVendorInformationExplicitNull (un)sets VendorInformation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The VendorInformation value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


