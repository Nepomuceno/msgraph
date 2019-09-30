# LicenseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePlans** | Pointer to [**[]MicrosoftGraphServicePlanInfo**](microsoft.graph.servicePlanInfo.md) |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 
**SkuPartNumber** | Pointer to **string** |  | [optional] 

## Methods

### GetServicePlans

`func (o *LicenseDetails) GetServicePlans() []MicrosoftGraphServicePlanInfo`

GetServicePlans returns the ServicePlans field if non-nil, zero value otherwise.

### GetServicePlansOk

`func (o *LicenseDetails) GetServicePlansOk() ([]MicrosoftGraphServicePlanInfo, bool)`

GetServicePlansOk returns a tuple with the ServicePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServicePlans

`func (o *LicenseDetails) HasServicePlans() bool`

HasServicePlans returns a boolean if a field has been set.

### SetServicePlans

`func (o *LicenseDetails) SetServicePlans(v []MicrosoftGraphServicePlanInfo)`

SetServicePlans gets a reference to the given []MicrosoftGraphServicePlanInfo and assigns it to the ServicePlans field.

### GetSkuId

`func (o *LicenseDetails) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *LicenseDetails) GetSkuIdOk() (string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSkuId

`func (o *LicenseDetails) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### SetSkuId

`func (o *LicenseDetails) SetSkuId(v string)`

SetSkuId gets a reference to the given string and assigns it to the SkuId field.

### SetSkuIdExplicitNull

`func (o *LicenseDetails) SetSkuIdExplicitNull(b bool)`

SetSkuIdExplicitNull (un)sets SkuId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SkuId value is set to nil even if false is passed
### GetSkuPartNumber

`func (o *LicenseDetails) GetSkuPartNumber() string`

GetSkuPartNumber returns the SkuPartNumber field if non-nil, zero value otherwise.

### GetSkuPartNumberOk

`func (o *LicenseDetails) GetSkuPartNumberOk() (string, bool)`

GetSkuPartNumberOk returns a tuple with the SkuPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSkuPartNumber

`func (o *LicenseDetails) HasSkuPartNumber() bool`

HasSkuPartNumber returns a boolean if a field has been set.

### SetSkuPartNumber

`func (o *LicenseDetails) SetSkuPartNumber(v string)`

SetSkuPartNumber gets a reference to the given string and assigns it to the SkuPartNumber field.

### SetSkuPartNumberExplicitNull

`func (o *LicenseDetails) SetSkuPartNumberExplicitNull(b bool)`

SetSkuPartNumberExplicitNull (un)sets SkuPartNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SkuPartNumber value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


