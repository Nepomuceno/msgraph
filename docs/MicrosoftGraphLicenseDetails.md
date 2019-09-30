# MicrosoftGraphLicenseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ServicePlans** | Pointer to [**[]MicrosoftGraphServicePlanInfo**](microsoft.graph.servicePlanInfo.md) |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 
**SkuPartNumber** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphLicenseDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphLicenseDetails) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphLicenseDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphLicenseDetails) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetServicePlans

`func (o *MicrosoftGraphLicenseDetails) GetServicePlans() []MicrosoftGraphServicePlanInfo`

GetServicePlans returns the ServicePlans field if non-nil, zero value otherwise.

### GetServicePlansOk

`func (o *MicrosoftGraphLicenseDetails) GetServicePlansOk() ([]MicrosoftGraphServicePlanInfo, bool)`

GetServicePlansOk returns a tuple with the ServicePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServicePlans

`func (o *MicrosoftGraphLicenseDetails) HasServicePlans() bool`

HasServicePlans returns a boolean if a field has been set.

### SetServicePlans

`func (o *MicrosoftGraphLicenseDetails) SetServicePlans(v []MicrosoftGraphServicePlanInfo)`

SetServicePlans gets a reference to the given []MicrosoftGraphServicePlanInfo and assigns it to the ServicePlans field.

### GetSkuId

`func (o *MicrosoftGraphLicenseDetails) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *MicrosoftGraphLicenseDetails) GetSkuIdOk() (string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSkuId

`func (o *MicrosoftGraphLicenseDetails) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### SetSkuId

`func (o *MicrosoftGraphLicenseDetails) SetSkuId(v string)`

SetSkuId gets a reference to the given string and assigns it to the SkuId field.

### SetSkuIdExplicitNull

`func (o *MicrosoftGraphLicenseDetails) SetSkuIdExplicitNull(b bool)`

SetSkuIdExplicitNull (un)sets SkuId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SkuId value is set to nil even if false is passed
### GetSkuPartNumber

`func (o *MicrosoftGraphLicenseDetails) GetSkuPartNumber() string`

GetSkuPartNumber returns the SkuPartNumber field if non-nil, zero value otherwise.

### GetSkuPartNumberOk

`func (o *MicrosoftGraphLicenseDetails) GetSkuPartNumberOk() (string, bool)`

GetSkuPartNumberOk returns a tuple with the SkuPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSkuPartNumber

`func (o *MicrosoftGraphLicenseDetails) HasSkuPartNumber() bool`

HasSkuPartNumber returns a boolean if a field has been set.

### SetSkuPartNumber

`func (o *MicrosoftGraphLicenseDetails) SetSkuPartNumber(v string)`

SetSkuPartNumber gets a reference to the given string and assigns it to the SkuPartNumber field.

### SetSkuPartNumberExplicitNull

`func (o *MicrosoftGraphLicenseDetails) SetSkuPartNumberExplicitNull(b bool)`

SetSkuPartNumberExplicitNull (un)sets SkuPartNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SkuPartNumber value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


