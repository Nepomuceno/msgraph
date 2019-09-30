# MicrosoftGraphManagedAppDiagnosticStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationName** | Pointer to **string** | The validation friendly name | [optional] 
**State** | Pointer to **string** | The state of the operation | [optional] 
**MitigationInstruction** | Pointer to **string** | Instruction on how to mitigate a failed validation | [optional] 

## Methods

### GetValidationName

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) GetValidationName() string`

GetValidationName returns the ValidationName field if non-nil, zero value otherwise.

### GetValidationNameOk

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) GetValidationNameOk() (string, bool)`

GetValidationNameOk returns a tuple with the ValidationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidationName

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) HasValidationName() bool`

HasValidationName returns a boolean if a field has been set.

### SetValidationName

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) SetValidationName(v string)`

SetValidationName gets a reference to the given string and assigns it to the ValidationName field.

### SetValidationNameExplicitNull

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) SetValidationNameExplicitNull(b bool)`

SetValidationNameExplicitNull (un)sets ValidationName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ValidationName value is set to nil even if false is passed
### GetState

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### SetStateExplicitNull

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) SetStateExplicitNull(b bool)`

SetStateExplicitNull (un)sets State to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The State value is set to nil even if false is passed
### GetMitigationInstruction

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) GetMitigationInstruction() string`

GetMitigationInstruction returns the MitigationInstruction field if non-nil, zero value otherwise.

### GetMitigationInstructionOk

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) GetMitigationInstructionOk() (string, bool)`

GetMitigationInstructionOk returns a tuple with the MitigationInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMitigationInstruction

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) HasMitigationInstruction() bool`

HasMitigationInstruction returns a boolean if a field has been set.

### SetMitigationInstruction

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) SetMitigationInstruction(v string)`

SetMitigationInstruction gets a reference to the given string and assigns it to the MitigationInstruction field.

### SetMitigationInstructionExplicitNull

`func (o *MicrosoftGraphManagedAppDiagnosticStatus) SetMitigationInstructionExplicitNull(b bool)`

SetMitigationInstructionExplicitNull (un)sets MitigationInstruction to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MitigationInstruction value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


