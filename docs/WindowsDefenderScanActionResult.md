# WindowsDefenderScanActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanType** | Pointer to **string** | Scan type either full scan or quick scan | [optional] 

## Methods

### GetScanType

`func (o *WindowsDefenderScanActionResult) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *WindowsDefenderScanActionResult) GetScanTypeOk() (string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScanType

`func (o *WindowsDefenderScanActionResult) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### SetScanType

`func (o *WindowsDefenderScanActionResult) SetScanType(v string)`

SetScanType gets a reference to the given string and assigns it to the ScanType field.

### SetScanTypeExplicitNull

`func (o *WindowsDefenderScanActionResult) SetScanTypeExplicitNull(b bool)`

SetScanTypeExplicitNull (un)sets ScanType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ScanType value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


