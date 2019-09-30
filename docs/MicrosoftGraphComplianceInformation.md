# MicrosoftGraphComplianceInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificationControls** | Pointer to [**[]AnyOfmicrosoftGraphCertificationControl**](anyOf&lt;microsoft.graph.certificationControl&gt;.md) |  | [optional] 
**CertificationName** | Pointer to **string** |  | [optional] 

## Methods

### GetCertificationControls

`func (o *MicrosoftGraphComplianceInformation) GetCertificationControls() []AnyOfmicrosoftGraphCertificationControl`

GetCertificationControls returns the CertificationControls field if non-nil, zero value otherwise.

### GetCertificationControlsOk

`func (o *MicrosoftGraphComplianceInformation) GetCertificationControlsOk() ([]AnyOfmicrosoftGraphCertificationControl, bool)`

GetCertificationControlsOk returns a tuple with the CertificationControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCertificationControls

`func (o *MicrosoftGraphComplianceInformation) HasCertificationControls() bool`

HasCertificationControls returns a boolean if a field has been set.

### SetCertificationControls

`func (o *MicrosoftGraphComplianceInformation) SetCertificationControls(v []AnyOfmicrosoftGraphCertificationControl)`

SetCertificationControls gets a reference to the given []AnyOfmicrosoftGraphCertificationControl and assigns it to the CertificationControls field.

### GetCertificationName

`func (o *MicrosoftGraphComplianceInformation) GetCertificationName() string`

GetCertificationName returns the CertificationName field if non-nil, zero value otherwise.

### GetCertificationNameOk

`func (o *MicrosoftGraphComplianceInformation) GetCertificationNameOk() (string, bool)`

GetCertificationNameOk returns a tuple with the CertificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCertificationName

`func (o *MicrosoftGraphComplianceInformation) HasCertificationName() bool`

HasCertificationName returns a boolean if a field has been set.

### SetCertificationName

`func (o *MicrosoftGraphComplianceInformation) SetCertificationName(v string)`

SetCertificationName gets a reference to the given string and assigns it to the CertificationName field.

### SetCertificationNameExplicitNull

`func (o *MicrosoftGraphComplianceInformation) SetCertificationNameExplicitNull(b bool)`

SetCertificationNameExplicitNull (un)sets CertificationName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CertificationName value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


