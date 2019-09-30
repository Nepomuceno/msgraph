# MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectName** | Pointer to **string** | Data recovery Certificate subject name | [optional] 
**Description** | Pointer to **string** | Data recovery Certificate description | [optional] 
**ExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) | Data recovery Certificate expiration datetime | [optional] 
**Certificate** | Pointer to **string** | Data recovery Certificate | [optional] 

## Methods

### GetSubjectName

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetSubjectNameOk() (string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubjectName

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### SetSubjectName

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetSubjectName(v string)`

SetSubjectName gets a reference to the given string and assigns it to the SubjectName field.

### SetSubjectNameExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetSubjectNameExplicitNull(b bool)`

SetSubjectNameExplicitNull (un)sets SubjectName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SubjectName value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetExpirationDateTime

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetExpirationDateTimeOk() (time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpirationDateTime

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.

### GetCertificate

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) GetCertificateOk() (string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCertificate

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificate

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetCertificate(v string)`

SetCertificate gets a reference to the given string and assigns it to the Certificate field.

### SetCertificateExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate) SetCertificateExplicitNull(b bool)`

SetCertificateExplicitNull (un)sets Certificate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Certificate value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


