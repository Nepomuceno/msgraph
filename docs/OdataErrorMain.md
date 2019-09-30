# OdataErrorMain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | 
**Message** | Pointer to **string** |  | 
**Target** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]OdataErrorDetail**](odata.error.detail.md) |  | [optional] 
**Innererror** | Pointer to [**map[string]interface{}**](.md) | The structure of this object is service-specific | [optional] 

## Methods

### GetCode

`func (o *OdataErrorMain) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OdataErrorMain) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *OdataErrorMain) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *OdataErrorMain) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### GetMessage

`func (o *OdataErrorMain) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OdataErrorMain) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *OdataErrorMain) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *OdataErrorMain) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetTarget

`func (o *OdataErrorMain) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *OdataErrorMain) GetTargetOk() (string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *OdataErrorMain) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *OdataErrorMain) SetTarget(v string)`

SetTarget gets a reference to the given string and assigns it to the Target field.

### GetDetails

`func (o *OdataErrorMain) GetDetails() []OdataErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *OdataErrorMain) GetDetailsOk() ([]OdataErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetails

`func (o *OdataErrorMain) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetails

`func (o *OdataErrorMain) SetDetails(v []OdataErrorDetail)`

SetDetails gets a reference to the given []OdataErrorDetail and assigns it to the Details field.

### GetInnererror

`func (o *OdataErrorMain) GetInnererror() map[string]interface{}`

GetInnererror returns the Innererror field if non-nil, zero value otherwise.

### GetInnererrorOk

`func (o *OdataErrorMain) GetInnererrorOk() (map[string]interface{}, bool)`

GetInnererrorOk returns a tuple with the Innererror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInnererror

`func (o *OdataErrorMain) HasInnererror() bool`

HasInnererror returns a boolean if a field has been set.

### SetInnererror

`func (o *OdataErrorMain) SetInnererror(v map[string]interface{})`

SetInnererror gets a reference to the given map[string]interface{} and assigns it to the Innererror field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


