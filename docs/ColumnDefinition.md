# ColumnDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boolean** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Calculated** | Pointer to [**AnyOfmicrosoftGraphCalculatedColumn**](anyOf&lt;microsoft.graph.calculatedColumn&gt;.md) |  | [optional] 
**Choice** | Pointer to [**AnyOfmicrosoftGraphChoiceColumn**](anyOf&lt;microsoft.graph.choiceColumn&gt;.md) |  | [optional] 
**ColumnGroup** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to [**AnyOfmicrosoftGraphCurrencyColumn**](anyOf&lt;microsoft.graph.currencyColumn&gt;.md) |  | [optional] 
**DateTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeColumn**](anyOf&lt;microsoft.graph.dateTimeColumn&gt;.md) |  | [optional] 
**DefaultValue** | Pointer to [**AnyOfmicrosoftGraphDefaultColumnValue**](anyOf&lt;microsoft.graph.defaultColumnValue&gt;.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EnforceUniqueValues** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Indexed** | Pointer to **bool** |  | [optional] 
**Lookup** | Pointer to [**AnyOfmicrosoftGraphLookupColumn**](anyOf&lt;microsoft.graph.lookupColumn&gt;.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Number** | Pointer to [**AnyOfmicrosoftGraphNumberColumn**](anyOf&lt;microsoft.graph.numberColumn&gt;.md) |  | [optional] 
**PersonOrGroup** | Pointer to [**AnyOfmicrosoftGraphPersonOrGroupColumn**](anyOf&lt;microsoft.graph.personOrGroupColumn&gt;.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Text** | Pointer to [**AnyOfmicrosoftGraphTextColumn**](anyOf&lt;microsoft.graph.textColumn&gt;.md) |  | [optional] 

## Methods

### GetBoolean

`func (o *ColumnDefinition) GetBoolean() AnyOfobject`

GetBoolean returns the Boolean field if non-nil, zero value otherwise.

### GetBooleanOk

`func (o *ColumnDefinition) GetBooleanOk() (AnyOfobject, bool)`

GetBooleanOk returns a tuple with the Boolean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBoolean

`func (o *ColumnDefinition) HasBoolean() bool`

HasBoolean returns a boolean if a field has been set.

### SetBoolean

`func (o *ColumnDefinition) SetBoolean(v AnyOfobject)`

SetBoolean gets a reference to the given AnyOfobject and assigns it to the Boolean field.

### SetBooleanExplicitNull

`func (o *ColumnDefinition) SetBooleanExplicitNull(b bool)`

SetBooleanExplicitNull (un)sets Boolean to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Boolean value is set to nil even if false is passed
### GetCalculated

`func (o *ColumnDefinition) GetCalculated() AnyOfmicrosoftGraphCalculatedColumn`

GetCalculated returns the Calculated field if non-nil, zero value otherwise.

### GetCalculatedOk

`func (o *ColumnDefinition) GetCalculatedOk() (AnyOfmicrosoftGraphCalculatedColumn, bool)`

GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalculated

`func (o *ColumnDefinition) HasCalculated() bool`

HasCalculated returns a boolean if a field has been set.

### SetCalculated

`func (o *ColumnDefinition) SetCalculated(v AnyOfmicrosoftGraphCalculatedColumn)`

SetCalculated gets a reference to the given AnyOfmicrosoftGraphCalculatedColumn and assigns it to the Calculated field.

### SetCalculatedExplicitNull

`func (o *ColumnDefinition) SetCalculatedExplicitNull(b bool)`

SetCalculatedExplicitNull (un)sets Calculated to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Calculated value is set to nil even if false is passed
### GetChoice

`func (o *ColumnDefinition) GetChoice() AnyOfmicrosoftGraphChoiceColumn`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *ColumnDefinition) GetChoiceOk() (AnyOfmicrosoftGraphChoiceColumn, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChoice

`func (o *ColumnDefinition) HasChoice() bool`

HasChoice returns a boolean if a field has been set.

### SetChoice

`func (o *ColumnDefinition) SetChoice(v AnyOfmicrosoftGraphChoiceColumn)`

SetChoice gets a reference to the given AnyOfmicrosoftGraphChoiceColumn and assigns it to the Choice field.

### SetChoiceExplicitNull

`func (o *ColumnDefinition) SetChoiceExplicitNull(b bool)`

SetChoiceExplicitNull (un)sets Choice to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Choice value is set to nil even if false is passed
### GetColumnGroup

`func (o *ColumnDefinition) GetColumnGroup() string`

GetColumnGroup returns the ColumnGroup field if non-nil, zero value otherwise.

### GetColumnGroupOk

`func (o *ColumnDefinition) GetColumnGroupOk() (string, bool)`

GetColumnGroupOk returns a tuple with the ColumnGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnGroup

`func (o *ColumnDefinition) HasColumnGroup() bool`

HasColumnGroup returns a boolean if a field has been set.

### SetColumnGroup

`func (o *ColumnDefinition) SetColumnGroup(v string)`

SetColumnGroup gets a reference to the given string and assigns it to the ColumnGroup field.

### SetColumnGroupExplicitNull

`func (o *ColumnDefinition) SetColumnGroupExplicitNull(b bool)`

SetColumnGroupExplicitNull (un)sets ColumnGroup to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ColumnGroup value is set to nil even if false is passed
### GetCurrency

`func (o *ColumnDefinition) GetCurrency() AnyOfmicrosoftGraphCurrencyColumn`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ColumnDefinition) GetCurrencyOk() (AnyOfmicrosoftGraphCurrencyColumn, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrency

`func (o *ColumnDefinition) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrency

`func (o *ColumnDefinition) SetCurrency(v AnyOfmicrosoftGraphCurrencyColumn)`

SetCurrency gets a reference to the given AnyOfmicrosoftGraphCurrencyColumn and assigns it to the Currency field.

### SetCurrencyExplicitNull

`func (o *ColumnDefinition) SetCurrencyExplicitNull(b bool)`

SetCurrencyExplicitNull (un)sets Currency to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Currency value is set to nil even if false is passed
### GetDateTime

`func (o *ColumnDefinition) GetDateTime() AnyOfmicrosoftGraphDateTimeColumn`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *ColumnDefinition) GetDateTimeOk() (AnyOfmicrosoftGraphDateTimeColumn, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateTime

`func (o *ColumnDefinition) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTime

`func (o *ColumnDefinition) SetDateTime(v AnyOfmicrosoftGraphDateTimeColumn)`

SetDateTime gets a reference to the given AnyOfmicrosoftGraphDateTimeColumn and assigns it to the DateTime field.

### SetDateTimeExplicitNull

`func (o *ColumnDefinition) SetDateTimeExplicitNull(b bool)`

SetDateTimeExplicitNull (un)sets DateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DateTime value is set to nil even if false is passed
### GetDefaultValue

`func (o *ColumnDefinition) GetDefaultValue() AnyOfmicrosoftGraphDefaultColumnValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ColumnDefinition) GetDefaultValueOk() (AnyOfmicrosoftGraphDefaultColumnValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultValue

`func (o *ColumnDefinition) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValue

`func (o *ColumnDefinition) SetDefaultValue(v AnyOfmicrosoftGraphDefaultColumnValue)`

SetDefaultValue gets a reference to the given AnyOfmicrosoftGraphDefaultColumnValue and assigns it to the DefaultValue field.

### SetDefaultValueExplicitNull

`func (o *ColumnDefinition) SetDefaultValueExplicitNull(b bool)`

SetDefaultValueExplicitNull (un)sets DefaultValue to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefaultValue value is set to nil even if false is passed
### GetDescription

`func (o *ColumnDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ColumnDefinition) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *ColumnDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *ColumnDefinition) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *ColumnDefinition) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *ColumnDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ColumnDefinition) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *ColumnDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *ColumnDefinition) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *ColumnDefinition) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetEnforceUniqueValues

`func (o *ColumnDefinition) GetEnforceUniqueValues() bool`

GetEnforceUniqueValues returns the EnforceUniqueValues field if non-nil, zero value otherwise.

### GetEnforceUniqueValuesOk

`func (o *ColumnDefinition) GetEnforceUniqueValuesOk() (bool, bool)`

GetEnforceUniqueValuesOk returns a tuple with the EnforceUniqueValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnforceUniqueValues

`func (o *ColumnDefinition) HasEnforceUniqueValues() bool`

HasEnforceUniqueValues returns a boolean if a field has been set.

### SetEnforceUniqueValues

`func (o *ColumnDefinition) SetEnforceUniqueValues(v bool)`

SetEnforceUniqueValues gets a reference to the given bool and assigns it to the EnforceUniqueValues field.

### SetEnforceUniqueValuesExplicitNull

`func (o *ColumnDefinition) SetEnforceUniqueValuesExplicitNull(b bool)`

SetEnforceUniqueValuesExplicitNull (un)sets EnforceUniqueValues to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnforceUniqueValues value is set to nil even if false is passed
### GetHidden

`func (o *ColumnDefinition) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ColumnDefinition) GetHiddenOk() (bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHidden

`func (o *ColumnDefinition) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHidden

`func (o *ColumnDefinition) SetHidden(v bool)`

SetHidden gets a reference to the given bool and assigns it to the Hidden field.

### SetHiddenExplicitNull

`func (o *ColumnDefinition) SetHiddenExplicitNull(b bool)`

SetHiddenExplicitNull (un)sets Hidden to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Hidden value is set to nil even if false is passed
### GetIndexed

`func (o *ColumnDefinition) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *ColumnDefinition) GetIndexedOk() (bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndexed

`func (o *ColumnDefinition) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### SetIndexed

`func (o *ColumnDefinition) SetIndexed(v bool)`

SetIndexed gets a reference to the given bool and assigns it to the Indexed field.

### SetIndexedExplicitNull

`func (o *ColumnDefinition) SetIndexedExplicitNull(b bool)`

SetIndexedExplicitNull (un)sets Indexed to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Indexed value is set to nil even if false is passed
### GetLookup

`func (o *ColumnDefinition) GetLookup() AnyOfmicrosoftGraphLookupColumn`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *ColumnDefinition) GetLookupOk() (AnyOfmicrosoftGraphLookupColumn, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLookup

`func (o *ColumnDefinition) HasLookup() bool`

HasLookup returns a boolean if a field has been set.

### SetLookup

`func (o *ColumnDefinition) SetLookup(v AnyOfmicrosoftGraphLookupColumn)`

SetLookup gets a reference to the given AnyOfmicrosoftGraphLookupColumn and assigns it to the Lookup field.

### SetLookupExplicitNull

`func (o *ColumnDefinition) SetLookupExplicitNull(b bool)`

SetLookupExplicitNull (un)sets Lookup to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Lookup value is set to nil even if false is passed
### GetName

`func (o *ColumnDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ColumnDefinition) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *ColumnDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *ColumnDefinition) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *ColumnDefinition) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetNumber

`func (o *ColumnDefinition) GetNumber() AnyOfmicrosoftGraphNumberColumn`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ColumnDefinition) GetNumberOk() (AnyOfmicrosoftGraphNumberColumn, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumber

`func (o *ColumnDefinition) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumber

`func (o *ColumnDefinition) SetNumber(v AnyOfmicrosoftGraphNumberColumn)`

SetNumber gets a reference to the given AnyOfmicrosoftGraphNumberColumn and assigns it to the Number field.

### SetNumberExplicitNull

`func (o *ColumnDefinition) SetNumberExplicitNull(b bool)`

SetNumberExplicitNull (un)sets Number to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Number value is set to nil even if false is passed
### GetPersonOrGroup

`func (o *ColumnDefinition) GetPersonOrGroup() AnyOfmicrosoftGraphPersonOrGroupColumn`

GetPersonOrGroup returns the PersonOrGroup field if non-nil, zero value otherwise.

### GetPersonOrGroupOk

`func (o *ColumnDefinition) GetPersonOrGroupOk() (AnyOfmicrosoftGraphPersonOrGroupColumn, bool)`

GetPersonOrGroupOk returns a tuple with the PersonOrGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPersonOrGroup

`func (o *ColumnDefinition) HasPersonOrGroup() bool`

HasPersonOrGroup returns a boolean if a field has been set.

### SetPersonOrGroup

`func (o *ColumnDefinition) SetPersonOrGroup(v AnyOfmicrosoftGraphPersonOrGroupColumn)`

SetPersonOrGroup gets a reference to the given AnyOfmicrosoftGraphPersonOrGroupColumn and assigns it to the PersonOrGroup field.

### SetPersonOrGroupExplicitNull

`func (o *ColumnDefinition) SetPersonOrGroupExplicitNull(b bool)`

SetPersonOrGroupExplicitNull (un)sets PersonOrGroup to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PersonOrGroup value is set to nil even if false is passed
### GetReadOnly

`func (o *ColumnDefinition) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ColumnDefinition) GetReadOnlyOk() (bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReadOnly

`func (o *ColumnDefinition) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnly

`func (o *ColumnDefinition) SetReadOnly(v bool)`

SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.

### SetReadOnlyExplicitNull

`func (o *ColumnDefinition) SetReadOnlyExplicitNull(b bool)`

SetReadOnlyExplicitNull (un)sets ReadOnly to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReadOnly value is set to nil even if false is passed
### GetRequired

`func (o *ColumnDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ColumnDefinition) GetRequiredOk() (bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequired

`func (o *ColumnDefinition) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequired

`func (o *ColumnDefinition) SetRequired(v bool)`

SetRequired gets a reference to the given bool and assigns it to the Required field.

### SetRequiredExplicitNull

`func (o *ColumnDefinition) SetRequiredExplicitNull(b bool)`

SetRequiredExplicitNull (un)sets Required to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Required value is set to nil even if false is passed
### GetText

`func (o *ColumnDefinition) GetText() AnyOfmicrosoftGraphTextColumn`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ColumnDefinition) GetTextOk() (AnyOfmicrosoftGraphTextColumn, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasText

`func (o *ColumnDefinition) HasText() bool`

HasText returns a boolean if a field has been set.

### SetText

`func (o *ColumnDefinition) SetText(v AnyOfmicrosoftGraphTextColumn)`

SetText gets a reference to the given AnyOfmicrosoftGraphTextColumn and assigns it to the Text field.

### SetTextExplicitNull

`func (o *ColumnDefinition) SetTextExplicitNull(b bool)`

SetTextExplicitNull (un)sets Text to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Text value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


