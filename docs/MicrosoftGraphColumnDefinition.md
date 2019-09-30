# MicrosoftGraphColumnDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphColumnDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphColumnDefinition) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphColumnDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphColumnDefinition) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetBoolean

`func (o *MicrosoftGraphColumnDefinition) GetBoolean() AnyOfobject`

GetBoolean returns the Boolean field if non-nil, zero value otherwise.

### GetBooleanOk

`func (o *MicrosoftGraphColumnDefinition) GetBooleanOk() (AnyOfobject, bool)`

GetBooleanOk returns a tuple with the Boolean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBoolean

`func (o *MicrosoftGraphColumnDefinition) HasBoolean() bool`

HasBoolean returns a boolean if a field has been set.

### SetBoolean

`func (o *MicrosoftGraphColumnDefinition) SetBoolean(v AnyOfobject)`

SetBoolean gets a reference to the given AnyOfobject and assigns it to the Boolean field.

### SetBooleanExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetBooleanExplicitNull(b bool)`

SetBooleanExplicitNull (un)sets Boolean to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Boolean value is set to nil even if false is passed
### GetCalculated

`func (o *MicrosoftGraphColumnDefinition) GetCalculated() AnyOfmicrosoftGraphCalculatedColumn`

GetCalculated returns the Calculated field if non-nil, zero value otherwise.

### GetCalculatedOk

`func (o *MicrosoftGraphColumnDefinition) GetCalculatedOk() (AnyOfmicrosoftGraphCalculatedColumn, bool)`

GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalculated

`func (o *MicrosoftGraphColumnDefinition) HasCalculated() bool`

HasCalculated returns a boolean if a field has been set.

### SetCalculated

`func (o *MicrosoftGraphColumnDefinition) SetCalculated(v AnyOfmicrosoftGraphCalculatedColumn)`

SetCalculated gets a reference to the given AnyOfmicrosoftGraphCalculatedColumn and assigns it to the Calculated field.

### SetCalculatedExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetCalculatedExplicitNull(b bool)`

SetCalculatedExplicitNull (un)sets Calculated to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Calculated value is set to nil even if false is passed
### GetChoice

`func (o *MicrosoftGraphColumnDefinition) GetChoice() AnyOfmicrosoftGraphChoiceColumn`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *MicrosoftGraphColumnDefinition) GetChoiceOk() (AnyOfmicrosoftGraphChoiceColumn, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChoice

`func (o *MicrosoftGraphColumnDefinition) HasChoice() bool`

HasChoice returns a boolean if a field has been set.

### SetChoice

`func (o *MicrosoftGraphColumnDefinition) SetChoice(v AnyOfmicrosoftGraphChoiceColumn)`

SetChoice gets a reference to the given AnyOfmicrosoftGraphChoiceColumn and assigns it to the Choice field.

### SetChoiceExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetChoiceExplicitNull(b bool)`

SetChoiceExplicitNull (un)sets Choice to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Choice value is set to nil even if false is passed
### GetColumnGroup

`func (o *MicrosoftGraphColumnDefinition) GetColumnGroup() string`

GetColumnGroup returns the ColumnGroup field if non-nil, zero value otherwise.

### GetColumnGroupOk

`func (o *MicrosoftGraphColumnDefinition) GetColumnGroupOk() (string, bool)`

GetColumnGroupOk returns a tuple with the ColumnGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnGroup

`func (o *MicrosoftGraphColumnDefinition) HasColumnGroup() bool`

HasColumnGroup returns a boolean if a field has been set.

### SetColumnGroup

`func (o *MicrosoftGraphColumnDefinition) SetColumnGroup(v string)`

SetColumnGroup gets a reference to the given string and assigns it to the ColumnGroup field.

### SetColumnGroupExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetColumnGroupExplicitNull(b bool)`

SetColumnGroupExplicitNull (un)sets ColumnGroup to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ColumnGroup value is set to nil even if false is passed
### GetCurrency

`func (o *MicrosoftGraphColumnDefinition) GetCurrency() AnyOfmicrosoftGraphCurrencyColumn`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MicrosoftGraphColumnDefinition) GetCurrencyOk() (AnyOfmicrosoftGraphCurrencyColumn, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrency

`func (o *MicrosoftGraphColumnDefinition) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrency

`func (o *MicrosoftGraphColumnDefinition) SetCurrency(v AnyOfmicrosoftGraphCurrencyColumn)`

SetCurrency gets a reference to the given AnyOfmicrosoftGraphCurrencyColumn and assigns it to the Currency field.

### SetCurrencyExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetCurrencyExplicitNull(b bool)`

SetCurrencyExplicitNull (un)sets Currency to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Currency value is set to nil even if false is passed
### GetDateTime

`func (o *MicrosoftGraphColumnDefinition) GetDateTime() AnyOfmicrosoftGraphDateTimeColumn`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *MicrosoftGraphColumnDefinition) GetDateTimeOk() (AnyOfmicrosoftGraphDateTimeColumn, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateTime

`func (o *MicrosoftGraphColumnDefinition) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### SetDateTime

`func (o *MicrosoftGraphColumnDefinition) SetDateTime(v AnyOfmicrosoftGraphDateTimeColumn)`

SetDateTime gets a reference to the given AnyOfmicrosoftGraphDateTimeColumn and assigns it to the DateTime field.

### SetDateTimeExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetDateTimeExplicitNull(b bool)`

SetDateTimeExplicitNull (un)sets DateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DateTime value is set to nil even if false is passed
### GetDefaultValue

`func (o *MicrosoftGraphColumnDefinition) GetDefaultValue() AnyOfmicrosoftGraphDefaultColumnValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *MicrosoftGraphColumnDefinition) GetDefaultValueOk() (AnyOfmicrosoftGraphDefaultColumnValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultValue

`func (o *MicrosoftGraphColumnDefinition) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValue

`func (o *MicrosoftGraphColumnDefinition) SetDefaultValue(v AnyOfmicrosoftGraphDefaultColumnValue)`

SetDefaultValue gets a reference to the given AnyOfmicrosoftGraphDefaultColumnValue and assigns it to the DefaultValue field.

### SetDefaultValueExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetDefaultValueExplicitNull(b bool)`

SetDefaultValueExplicitNull (un)sets DefaultValue to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefaultValue value is set to nil even if false is passed
### GetDescription

`func (o *MicrosoftGraphColumnDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphColumnDefinition) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphColumnDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphColumnDefinition) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphColumnDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphColumnDefinition) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphColumnDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphColumnDefinition) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetEnforceUniqueValues

`func (o *MicrosoftGraphColumnDefinition) GetEnforceUniqueValues() bool`

GetEnforceUniqueValues returns the EnforceUniqueValues field if non-nil, zero value otherwise.

### GetEnforceUniqueValuesOk

`func (o *MicrosoftGraphColumnDefinition) GetEnforceUniqueValuesOk() (bool, bool)`

GetEnforceUniqueValuesOk returns a tuple with the EnforceUniqueValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnforceUniqueValues

`func (o *MicrosoftGraphColumnDefinition) HasEnforceUniqueValues() bool`

HasEnforceUniqueValues returns a boolean if a field has been set.

### SetEnforceUniqueValues

`func (o *MicrosoftGraphColumnDefinition) SetEnforceUniqueValues(v bool)`

SetEnforceUniqueValues gets a reference to the given bool and assigns it to the EnforceUniqueValues field.

### SetEnforceUniqueValuesExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetEnforceUniqueValuesExplicitNull(b bool)`

SetEnforceUniqueValuesExplicitNull (un)sets EnforceUniqueValues to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnforceUniqueValues value is set to nil even if false is passed
### GetHidden

`func (o *MicrosoftGraphColumnDefinition) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *MicrosoftGraphColumnDefinition) GetHiddenOk() (bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHidden

`func (o *MicrosoftGraphColumnDefinition) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHidden

`func (o *MicrosoftGraphColumnDefinition) SetHidden(v bool)`

SetHidden gets a reference to the given bool and assigns it to the Hidden field.

### SetHiddenExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetHiddenExplicitNull(b bool)`

SetHiddenExplicitNull (un)sets Hidden to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Hidden value is set to nil even if false is passed
### GetIndexed

`func (o *MicrosoftGraphColumnDefinition) GetIndexed() bool`

GetIndexed returns the Indexed field if non-nil, zero value otherwise.

### GetIndexedOk

`func (o *MicrosoftGraphColumnDefinition) GetIndexedOk() (bool, bool)`

GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndexed

`func (o *MicrosoftGraphColumnDefinition) HasIndexed() bool`

HasIndexed returns a boolean if a field has been set.

### SetIndexed

`func (o *MicrosoftGraphColumnDefinition) SetIndexed(v bool)`

SetIndexed gets a reference to the given bool and assigns it to the Indexed field.

### SetIndexedExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetIndexedExplicitNull(b bool)`

SetIndexedExplicitNull (un)sets Indexed to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Indexed value is set to nil even if false is passed
### GetLookup

`func (o *MicrosoftGraphColumnDefinition) GetLookup() AnyOfmicrosoftGraphLookupColumn`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *MicrosoftGraphColumnDefinition) GetLookupOk() (AnyOfmicrosoftGraphLookupColumn, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLookup

`func (o *MicrosoftGraphColumnDefinition) HasLookup() bool`

HasLookup returns a boolean if a field has been set.

### SetLookup

`func (o *MicrosoftGraphColumnDefinition) SetLookup(v AnyOfmicrosoftGraphLookupColumn)`

SetLookup gets a reference to the given AnyOfmicrosoftGraphLookupColumn and assigns it to the Lookup field.

### SetLookupExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetLookupExplicitNull(b bool)`

SetLookupExplicitNull (un)sets Lookup to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Lookup value is set to nil even if false is passed
### GetName

`func (o *MicrosoftGraphColumnDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphColumnDefinition) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphColumnDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphColumnDefinition) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetNumber

`func (o *MicrosoftGraphColumnDefinition) GetNumber() AnyOfmicrosoftGraphNumberColumn`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *MicrosoftGraphColumnDefinition) GetNumberOk() (AnyOfmicrosoftGraphNumberColumn, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumber

`func (o *MicrosoftGraphColumnDefinition) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumber

`func (o *MicrosoftGraphColumnDefinition) SetNumber(v AnyOfmicrosoftGraphNumberColumn)`

SetNumber gets a reference to the given AnyOfmicrosoftGraphNumberColumn and assigns it to the Number field.

### SetNumberExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetNumberExplicitNull(b bool)`

SetNumberExplicitNull (un)sets Number to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Number value is set to nil even if false is passed
### GetPersonOrGroup

`func (o *MicrosoftGraphColumnDefinition) GetPersonOrGroup() AnyOfmicrosoftGraphPersonOrGroupColumn`

GetPersonOrGroup returns the PersonOrGroup field if non-nil, zero value otherwise.

### GetPersonOrGroupOk

`func (o *MicrosoftGraphColumnDefinition) GetPersonOrGroupOk() (AnyOfmicrosoftGraphPersonOrGroupColumn, bool)`

GetPersonOrGroupOk returns a tuple with the PersonOrGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPersonOrGroup

`func (o *MicrosoftGraphColumnDefinition) HasPersonOrGroup() bool`

HasPersonOrGroup returns a boolean if a field has been set.

### SetPersonOrGroup

`func (o *MicrosoftGraphColumnDefinition) SetPersonOrGroup(v AnyOfmicrosoftGraphPersonOrGroupColumn)`

SetPersonOrGroup gets a reference to the given AnyOfmicrosoftGraphPersonOrGroupColumn and assigns it to the PersonOrGroup field.

### SetPersonOrGroupExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetPersonOrGroupExplicitNull(b bool)`

SetPersonOrGroupExplicitNull (un)sets PersonOrGroup to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PersonOrGroup value is set to nil even if false is passed
### GetReadOnly

`func (o *MicrosoftGraphColumnDefinition) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *MicrosoftGraphColumnDefinition) GetReadOnlyOk() (bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReadOnly

`func (o *MicrosoftGraphColumnDefinition) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### SetReadOnly

`func (o *MicrosoftGraphColumnDefinition) SetReadOnly(v bool)`

SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.

### SetReadOnlyExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetReadOnlyExplicitNull(b bool)`

SetReadOnlyExplicitNull (un)sets ReadOnly to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReadOnly value is set to nil even if false is passed
### GetRequired

`func (o *MicrosoftGraphColumnDefinition) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *MicrosoftGraphColumnDefinition) GetRequiredOk() (bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequired

`func (o *MicrosoftGraphColumnDefinition) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequired

`func (o *MicrosoftGraphColumnDefinition) SetRequired(v bool)`

SetRequired gets a reference to the given bool and assigns it to the Required field.

### SetRequiredExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetRequiredExplicitNull(b bool)`

SetRequiredExplicitNull (un)sets Required to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Required value is set to nil even if false is passed
### GetText

`func (o *MicrosoftGraphColumnDefinition) GetText() AnyOfmicrosoftGraphTextColumn`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MicrosoftGraphColumnDefinition) GetTextOk() (AnyOfmicrosoftGraphTextColumn, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasText

`func (o *MicrosoftGraphColumnDefinition) HasText() bool`

HasText returns a boolean if a field has been set.

### SetText

`func (o *MicrosoftGraphColumnDefinition) SetText(v AnyOfmicrosoftGraphTextColumn)`

SetText gets a reference to the given AnyOfmicrosoftGraphTextColumn and assigns it to the Text field.

### SetTextExplicitNull

`func (o *MicrosoftGraphColumnDefinition) SetTextExplicitNull(b bool)`

SetTextExplicitNull (un)sets Text to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Text value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


