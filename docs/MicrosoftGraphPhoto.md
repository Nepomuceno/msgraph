# MicrosoftGraphPhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CameraMake** | Pointer to **string** |  | [optional] 
**CameraModel** | Pointer to **string** |  | [optional] 
**ExposureDenominator** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**ExposureNumerator** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**FNumber** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**FocalLength** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**Iso** | Pointer to **int32** |  | [optional] 
**TakenDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### GetCameraMake

`func (o *MicrosoftGraphPhoto) GetCameraMake() string`

GetCameraMake returns the CameraMake field if non-nil, zero value otherwise.

### GetCameraMakeOk

`func (o *MicrosoftGraphPhoto) GetCameraMakeOk() (string, bool)`

GetCameraMakeOk returns a tuple with the CameraMake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCameraMake

`func (o *MicrosoftGraphPhoto) HasCameraMake() bool`

HasCameraMake returns a boolean if a field has been set.

### SetCameraMake

`func (o *MicrosoftGraphPhoto) SetCameraMake(v string)`

SetCameraMake gets a reference to the given string and assigns it to the CameraMake field.

### SetCameraMakeExplicitNull

`func (o *MicrosoftGraphPhoto) SetCameraMakeExplicitNull(b bool)`

SetCameraMakeExplicitNull (un)sets CameraMake to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CameraMake value is set to nil even if false is passed
### GetCameraModel

`func (o *MicrosoftGraphPhoto) GetCameraModel() string`

GetCameraModel returns the CameraModel field if non-nil, zero value otherwise.

### GetCameraModelOk

`func (o *MicrosoftGraphPhoto) GetCameraModelOk() (string, bool)`

GetCameraModelOk returns a tuple with the CameraModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCameraModel

`func (o *MicrosoftGraphPhoto) HasCameraModel() bool`

HasCameraModel returns a boolean if a field has been set.

### SetCameraModel

`func (o *MicrosoftGraphPhoto) SetCameraModel(v string)`

SetCameraModel gets a reference to the given string and assigns it to the CameraModel field.

### SetCameraModelExplicitNull

`func (o *MicrosoftGraphPhoto) SetCameraModelExplicitNull(b bool)`

SetCameraModelExplicitNull (un)sets CameraModel to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CameraModel value is set to nil even if false is passed
### GetExposureDenominator

`func (o *MicrosoftGraphPhoto) GetExposureDenominator() AnyOfnumberstringstring`

GetExposureDenominator returns the ExposureDenominator field if non-nil, zero value otherwise.

### GetExposureDenominatorOk

`func (o *MicrosoftGraphPhoto) GetExposureDenominatorOk() (AnyOfnumberstringstring, bool)`

GetExposureDenominatorOk returns a tuple with the ExposureDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExposureDenominator

`func (o *MicrosoftGraphPhoto) HasExposureDenominator() bool`

HasExposureDenominator returns a boolean if a field has been set.

### SetExposureDenominator

`func (o *MicrosoftGraphPhoto) SetExposureDenominator(v AnyOfnumberstringstring)`

SetExposureDenominator gets a reference to the given AnyOfnumberstringstring and assigns it to the ExposureDenominator field.

### SetExposureDenominatorExplicitNull

`func (o *MicrosoftGraphPhoto) SetExposureDenominatorExplicitNull(b bool)`

SetExposureDenominatorExplicitNull (un)sets ExposureDenominator to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExposureDenominator value is set to nil even if false is passed
### GetExposureNumerator

`func (o *MicrosoftGraphPhoto) GetExposureNumerator() AnyOfnumberstringstring`

GetExposureNumerator returns the ExposureNumerator field if non-nil, zero value otherwise.

### GetExposureNumeratorOk

`func (o *MicrosoftGraphPhoto) GetExposureNumeratorOk() (AnyOfnumberstringstring, bool)`

GetExposureNumeratorOk returns a tuple with the ExposureNumerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExposureNumerator

`func (o *MicrosoftGraphPhoto) HasExposureNumerator() bool`

HasExposureNumerator returns a boolean if a field has been set.

### SetExposureNumerator

`func (o *MicrosoftGraphPhoto) SetExposureNumerator(v AnyOfnumberstringstring)`

SetExposureNumerator gets a reference to the given AnyOfnumberstringstring and assigns it to the ExposureNumerator field.

### SetExposureNumeratorExplicitNull

`func (o *MicrosoftGraphPhoto) SetExposureNumeratorExplicitNull(b bool)`

SetExposureNumeratorExplicitNull (un)sets ExposureNumerator to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExposureNumerator value is set to nil even if false is passed
### GetFNumber

`func (o *MicrosoftGraphPhoto) GetFNumber() AnyOfnumberstringstring`

GetFNumber returns the FNumber field if non-nil, zero value otherwise.

### GetFNumberOk

`func (o *MicrosoftGraphPhoto) GetFNumberOk() (AnyOfnumberstringstring, bool)`

GetFNumberOk returns a tuple with the FNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFNumber

`func (o *MicrosoftGraphPhoto) HasFNumber() bool`

HasFNumber returns a boolean if a field has been set.

### SetFNumber

`func (o *MicrosoftGraphPhoto) SetFNumber(v AnyOfnumberstringstring)`

SetFNumber gets a reference to the given AnyOfnumberstringstring and assigns it to the FNumber field.

### SetFNumberExplicitNull

`func (o *MicrosoftGraphPhoto) SetFNumberExplicitNull(b bool)`

SetFNumberExplicitNull (un)sets FNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FNumber value is set to nil even if false is passed
### GetFocalLength

`func (o *MicrosoftGraphPhoto) GetFocalLength() AnyOfnumberstringstring`

GetFocalLength returns the FocalLength field if non-nil, zero value otherwise.

### GetFocalLengthOk

`func (o *MicrosoftGraphPhoto) GetFocalLengthOk() (AnyOfnumberstringstring, bool)`

GetFocalLengthOk returns a tuple with the FocalLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFocalLength

`func (o *MicrosoftGraphPhoto) HasFocalLength() bool`

HasFocalLength returns a boolean if a field has been set.

### SetFocalLength

`func (o *MicrosoftGraphPhoto) SetFocalLength(v AnyOfnumberstringstring)`

SetFocalLength gets a reference to the given AnyOfnumberstringstring and assigns it to the FocalLength field.

### SetFocalLengthExplicitNull

`func (o *MicrosoftGraphPhoto) SetFocalLengthExplicitNull(b bool)`

SetFocalLengthExplicitNull (un)sets FocalLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FocalLength value is set to nil even if false is passed
### GetIso

`func (o *MicrosoftGraphPhoto) GetIso() int32`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *MicrosoftGraphPhoto) GetIsoOk() (int32, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIso

`func (o *MicrosoftGraphPhoto) HasIso() bool`

HasIso returns a boolean if a field has been set.

### SetIso

`func (o *MicrosoftGraphPhoto) SetIso(v int32)`

SetIso gets a reference to the given int32 and assigns it to the Iso field.

### SetIsoExplicitNull

`func (o *MicrosoftGraphPhoto) SetIsoExplicitNull(b bool)`

SetIsoExplicitNull (un)sets Iso to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Iso value is set to nil even if false is passed
### GetTakenDateTime

`func (o *MicrosoftGraphPhoto) GetTakenDateTime() time.Time`

GetTakenDateTime returns the TakenDateTime field if non-nil, zero value otherwise.

### GetTakenDateTimeOk

`func (o *MicrosoftGraphPhoto) GetTakenDateTimeOk() (time.Time, bool)`

GetTakenDateTimeOk returns a tuple with the TakenDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTakenDateTime

`func (o *MicrosoftGraphPhoto) HasTakenDateTime() bool`

HasTakenDateTime returns a boolean if a field has been set.

### SetTakenDateTime

`func (o *MicrosoftGraphPhoto) SetTakenDateTime(v time.Time)`

SetTakenDateTime gets a reference to the given time.Time and assigns it to the TakenDateTime field.

### SetTakenDateTimeExplicitNull

`func (o *MicrosoftGraphPhoto) SetTakenDateTimeExplicitNull(b bool)`

SetTakenDateTimeExplicitNull (un)sets TakenDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TakenDateTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


