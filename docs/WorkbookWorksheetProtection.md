# WorkbookWorksheetProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions**](anyOf&lt;microsoft.graph.workbookWorksheetProtectionOptions&gt;.md) |  | [optional] 
**Protected** | Pointer to **bool** |  | [optional] 

## Methods

### GetOptions

`func (o *WorkbookWorksheetProtection) GetOptions() AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *WorkbookWorksheetProtection) GetOptionsOk() (AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOptions

`func (o *WorkbookWorksheetProtection) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptions

`func (o *WorkbookWorksheetProtection) SetOptions(v AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions)`

SetOptions gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions and assigns it to the Options field.

### SetOptionsExplicitNull

`func (o *WorkbookWorksheetProtection) SetOptionsExplicitNull(b bool)`

SetOptionsExplicitNull (un)sets Options to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Options value is set to nil even if false is passed
### GetProtected

`func (o *WorkbookWorksheetProtection) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *WorkbookWorksheetProtection) GetProtectedOk() (bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtected

`func (o *WorkbookWorksheetProtection) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### SetProtected

`func (o *WorkbookWorksheetProtection) SetProtected(v bool)`

SetProtected gets a reference to the given bool and assigns it to the Protected field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


