# SubscribedSku

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapabilityStatus** | Pointer to **string** |  | [optional] 
**ConsumedUnits** | Pointer to **int32** |  | [optional] 
**PrepaidUnits** | Pointer to [**AnyOfmicrosoftGraphLicenseUnitsDetail**](anyOf&lt;microsoft.graph.licenseUnitsDetail&gt;.md) |  | [optional] 
**ServicePlans** | Pointer to [**[]MicrosoftGraphServicePlanInfo**](microsoft.graph.servicePlanInfo.md) |  | [optional] 
**SkuId** | Pointer to **string** |  | [optional] 
**SkuPartNumber** | Pointer to **string** |  | [optional] 
**AppliesTo** | Pointer to **string** |  | [optional] 

## Methods

### GetCapabilityStatus

`func (o *SubscribedSku) GetCapabilityStatus() string`

GetCapabilityStatus returns the CapabilityStatus field if non-nil, zero value otherwise.

### GetCapabilityStatusOk

`func (o *SubscribedSku) GetCapabilityStatusOk() (string, bool)`

GetCapabilityStatusOk returns a tuple with the CapabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCapabilityStatus

`func (o *SubscribedSku) HasCapabilityStatus() bool`

HasCapabilityStatus returns a boolean if a field has been set.

### SetCapabilityStatus

`func (o *SubscribedSku) SetCapabilityStatus(v string)`

SetCapabilityStatus gets a reference to the given string and assigns it to the CapabilityStatus field.

### SetCapabilityStatusExplicitNull

`func (o *SubscribedSku) SetCapabilityStatusExplicitNull(b bool)`

SetCapabilityStatusExplicitNull (un)sets CapabilityStatus to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CapabilityStatus value is set to nil even if false is passed
### GetConsumedUnits

`func (o *SubscribedSku) GetConsumedUnits() int32`

GetConsumedUnits returns the ConsumedUnits field if non-nil, zero value otherwise.

### GetConsumedUnitsOk

`func (o *SubscribedSku) GetConsumedUnitsOk() (int32, bool)`

GetConsumedUnitsOk returns a tuple with the ConsumedUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConsumedUnits

`func (o *SubscribedSku) HasConsumedUnits() bool`

HasConsumedUnits returns a boolean if a field has been set.

### SetConsumedUnits

`func (o *SubscribedSku) SetConsumedUnits(v int32)`

SetConsumedUnits gets a reference to the given int32 and assigns it to the ConsumedUnits field.

### SetConsumedUnitsExplicitNull

`func (o *SubscribedSku) SetConsumedUnitsExplicitNull(b bool)`

SetConsumedUnitsExplicitNull (un)sets ConsumedUnits to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConsumedUnits value is set to nil even if false is passed
### GetPrepaidUnits

`func (o *SubscribedSku) GetPrepaidUnits() AnyOfmicrosoftGraphLicenseUnitsDetail`

GetPrepaidUnits returns the PrepaidUnits field if non-nil, zero value otherwise.

### GetPrepaidUnitsOk

`func (o *SubscribedSku) GetPrepaidUnitsOk() (AnyOfmicrosoftGraphLicenseUnitsDetail, bool)`

GetPrepaidUnitsOk returns a tuple with the PrepaidUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrepaidUnits

`func (o *SubscribedSku) HasPrepaidUnits() bool`

HasPrepaidUnits returns a boolean if a field has been set.

### SetPrepaidUnits

`func (o *SubscribedSku) SetPrepaidUnits(v AnyOfmicrosoftGraphLicenseUnitsDetail)`

SetPrepaidUnits gets a reference to the given AnyOfmicrosoftGraphLicenseUnitsDetail and assigns it to the PrepaidUnits field.

### SetPrepaidUnitsExplicitNull

`func (o *SubscribedSku) SetPrepaidUnitsExplicitNull(b bool)`

SetPrepaidUnitsExplicitNull (un)sets PrepaidUnits to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrepaidUnits value is set to nil even if false is passed
### GetServicePlans

`func (o *SubscribedSku) GetServicePlans() []MicrosoftGraphServicePlanInfo`

GetServicePlans returns the ServicePlans field if non-nil, zero value otherwise.

### GetServicePlansOk

`func (o *SubscribedSku) GetServicePlansOk() ([]MicrosoftGraphServicePlanInfo, bool)`

GetServicePlansOk returns a tuple with the ServicePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServicePlans

`func (o *SubscribedSku) HasServicePlans() bool`

HasServicePlans returns a boolean if a field has been set.

### SetServicePlans

`func (o *SubscribedSku) SetServicePlans(v []MicrosoftGraphServicePlanInfo)`

SetServicePlans gets a reference to the given []MicrosoftGraphServicePlanInfo and assigns it to the ServicePlans field.

### GetSkuId

`func (o *SubscribedSku) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *SubscribedSku) GetSkuIdOk() (string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSkuId

`func (o *SubscribedSku) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### SetSkuId

`func (o *SubscribedSku) SetSkuId(v string)`

SetSkuId gets a reference to the given string and assigns it to the SkuId field.

### SetSkuIdExplicitNull

`func (o *SubscribedSku) SetSkuIdExplicitNull(b bool)`

SetSkuIdExplicitNull (un)sets SkuId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SkuId value is set to nil even if false is passed
### GetSkuPartNumber

`func (o *SubscribedSku) GetSkuPartNumber() string`

GetSkuPartNumber returns the SkuPartNumber field if non-nil, zero value otherwise.

### GetSkuPartNumberOk

`func (o *SubscribedSku) GetSkuPartNumberOk() (string, bool)`

GetSkuPartNumberOk returns a tuple with the SkuPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSkuPartNumber

`func (o *SubscribedSku) HasSkuPartNumber() bool`

HasSkuPartNumber returns a boolean if a field has been set.

### SetSkuPartNumber

`func (o *SubscribedSku) SetSkuPartNumber(v string)`

SetSkuPartNumber gets a reference to the given string and assigns it to the SkuPartNumber field.

### SetSkuPartNumberExplicitNull

`func (o *SubscribedSku) SetSkuPartNumberExplicitNull(b bool)`

SetSkuPartNumberExplicitNull (un)sets SkuPartNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SkuPartNumber value is set to nil even if false is passed
### GetAppliesTo

`func (o *SubscribedSku) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *SubscribedSku) GetAppliesToOk() (string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppliesTo

`func (o *SubscribedSku) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### SetAppliesTo

`func (o *SubscribedSku) SetAppliesTo(v string)`

SetAppliesTo gets a reference to the given string and assigns it to the AppliesTo field.

### SetAppliesToExplicitNull

`func (o *SubscribedSku) SetAppliesToExplicitNull(b bool)`

SetAppliesToExplicitNull (un)sets AppliesTo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppliesTo value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


