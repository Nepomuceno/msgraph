/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
import (
	"encoding/json"
)
// MicrosoftGraphColumnDefinition struct for MicrosoftGraphColumnDefinition
type MicrosoftGraphColumnDefinition struct {
	Id *string `json:"id,omitempty"`

	Boolean *AnyOfobject `json:"boolean,omitempty"`
	isExplicitNullBoolean bool `json:"-"`
	Calculated *AnyOfmicrosoftGraphCalculatedColumn `json:"calculated,omitempty"`
	isExplicitNullCalculated bool `json:"-"`
	Choice *AnyOfmicrosoftGraphChoiceColumn `json:"choice,omitempty"`
	isExplicitNullChoice bool `json:"-"`
	ColumnGroup *string `json:"columnGroup,omitempty"`
	isExplicitNullColumnGroup bool `json:"-"`
	Currency *AnyOfmicrosoftGraphCurrencyColumn `json:"currency,omitempty"`
	isExplicitNullCurrency bool `json:"-"`
	DateTime *AnyOfmicrosoftGraphDateTimeColumn `json:"dateTime,omitempty"`
	isExplicitNullDateTime bool `json:"-"`
	DefaultValue *AnyOfmicrosoftGraphDefaultColumnValue `json:"defaultValue,omitempty"`
	isExplicitNullDefaultValue bool `json:"-"`
	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	EnforceUniqueValues *bool `json:"enforceUniqueValues,omitempty"`
	isExplicitNullEnforceUniqueValues bool `json:"-"`
	Hidden *bool `json:"hidden,omitempty"`
	isExplicitNullHidden bool `json:"-"`
	Indexed *bool `json:"indexed,omitempty"`
	isExplicitNullIndexed bool `json:"-"`
	Lookup *AnyOfmicrosoftGraphLookupColumn `json:"lookup,omitempty"`
	isExplicitNullLookup bool `json:"-"`
	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
	Number *AnyOfmicrosoftGraphNumberColumn `json:"number,omitempty"`
	isExplicitNullNumber bool `json:"-"`
	PersonOrGroup *AnyOfmicrosoftGraphPersonOrGroupColumn `json:"personOrGroup,omitempty"`
	isExplicitNullPersonOrGroup bool `json:"-"`
	ReadOnly *bool `json:"readOnly,omitempty"`
	isExplicitNullReadOnly bool `json:"-"`
	Required *bool `json:"required,omitempty"`
	isExplicitNullRequired bool `json:"-"`
	Text *AnyOfmicrosoftGraphTextColumn `json:"text,omitempty"`
	isExplicitNullText bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphColumnDefinition) SetId(v string) {
	o.Id = &v
}

// GetBoolean returns the Boolean field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetBoolean() AnyOfobject {
	if o == nil || o.Boolean == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Boolean
}

// GetBooleanOk returns a tuple with the Boolean field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetBooleanOk() (AnyOfobject, bool) {
	if o == nil || o.Boolean == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Boolean, true
}

// HasBoolean returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasBoolean() bool {
	if o != nil && o.Boolean != nil {
		return true
	}

	return false
}

// SetBoolean gets a reference to the given AnyOfobject and assigns it to the Boolean field.
func (o *MicrosoftGraphColumnDefinition) SetBoolean(v AnyOfobject) {
	o.Boolean = &v
}

// SetBooleanExplicitNull (un)sets Boolean to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Boolean value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetBooleanExplicitNull(b bool) {
	o.Boolean = nil
	o.isExplicitNullBoolean = b
}
// GetCalculated returns the Calculated field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetCalculated() AnyOfmicrosoftGraphCalculatedColumn {
	if o == nil || o.Calculated == nil {
		var ret AnyOfmicrosoftGraphCalculatedColumn
		return ret
	}
	return *o.Calculated
}

// GetCalculatedOk returns a tuple with the Calculated field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetCalculatedOk() (AnyOfmicrosoftGraphCalculatedColumn, bool) {
	if o == nil || o.Calculated == nil {
		var ret AnyOfmicrosoftGraphCalculatedColumn
		return ret, false
	}
	return *o.Calculated, true
}

// HasCalculated returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasCalculated() bool {
	if o != nil && o.Calculated != nil {
		return true
	}

	return false
}

// SetCalculated gets a reference to the given AnyOfmicrosoftGraphCalculatedColumn and assigns it to the Calculated field.
func (o *MicrosoftGraphColumnDefinition) SetCalculated(v AnyOfmicrosoftGraphCalculatedColumn) {
	o.Calculated = &v
}

// SetCalculatedExplicitNull (un)sets Calculated to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Calculated value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetCalculatedExplicitNull(b bool) {
	o.Calculated = nil
	o.isExplicitNullCalculated = b
}
// GetChoice returns the Choice field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetChoice() AnyOfmicrosoftGraphChoiceColumn {
	if o == nil || o.Choice == nil {
		var ret AnyOfmicrosoftGraphChoiceColumn
		return ret
	}
	return *o.Choice
}

// GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetChoiceOk() (AnyOfmicrosoftGraphChoiceColumn, bool) {
	if o == nil || o.Choice == nil {
		var ret AnyOfmicrosoftGraphChoiceColumn
		return ret, false
	}
	return *o.Choice, true
}

// HasChoice returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasChoice() bool {
	if o != nil && o.Choice != nil {
		return true
	}

	return false
}

// SetChoice gets a reference to the given AnyOfmicrosoftGraphChoiceColumn and assigns it to the Choice field.
func (o *MicrosoftGraphColumnDefinition) SetChoice(v AnyOfmicrosoftGraphChoiceColumn) {
	o.Choice = &v
}

// SetChoiceExplicitNull (un)sets Choice to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Choice value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetChoiceExplicitNull(b bool) {
	o.Choice = nil
	o.isExplicitNullChoice = b
}
// GetColumnGroup returns the ColumnGroup field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetColumnGroup() string {
	if o == nil || o.ColumnGroup == nil {
		var ret string
		return ret
	}
	return *o.ColumnGroup
}

// GetColumnGroupOk returns a tuple with the ColumnGroup field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetColumnGroupOk() (string, bool) {
	if o == nil || o.ColumnGroup == nil {
		var ret string
		return ret, false
	}
	return *o.ColumnGroup, true
}

// HasColumnGroup returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasColumnGroup() bool {
	if o != nil && o.ColumnGroup != nil {
		return true
	}

	return false
}

// SetColumnGroup gets a reference to the given string and assigns it to the ColumnGroup field.
func (o *MicrosoftGraphColumnDefinition) SetColumnGroup(v string) {
	o.ColumnGroup = &v
}

// SetColumnGroupExplicitNull (un)sets ColumnGroup to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ColumnGroup value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetColumnGroupExplicitNull(b bool) {
	o.ColumnGroup = nil
	o.isExplicitNullColumnGroup = b
}
// GetCurrency returns the Currency field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetCurrency() AnyOfmicrosoftGraphCurrencyColumn {
	if o == nil || o.Currency == nil {
		var ret AnyOfmicrosoftGraphCurrencyColumn
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetCurrencyOk() (AnyOfmicrosoftGraphCurrencyColumn, bool) {
	if o == nil || o.Currency == nil {
		var ret AnyOfmicrosoftGraphCurrencyColumn
		return ret, false
	}
	return *o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given AnyOfmicrosoftGraphCurrencyColumn and assigns it to the Currency field.
func (o *MicrosoftGraphColumnDefinition) SetCurrency(v AnyOfmicrosoftGraphCurrencyColumn) {
	o.Currency = &v
}

// SetCurrencyExplicitNull (un)sets Currency to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Currency value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetCurrencyExplicitNull(b bool) {
	o.Currency = nil
	o.isExplicitNullCurrency = b
}
// GetDateTime returns the DateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetDateTime() AnyOfmicrosoftGraphDateTimeColumn {
	if o == nil || o.DateTime == nil {
		var ret AnyOfmicrosoftGraphDateTimeColumn
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetDateTimeOk() (AnyOfmicrosoftGraphDateTimeColumn, bool) {
	if o == nil || o.DateTime == nil {
		var ret AnyOfmicrosoftGraphDateTimeColumn
		return ret, false
	}
	return *o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasDateTime() bool {
	if o != nil && o.DateTime != nil {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given AnyOfmicrosoftGraphDateTimeColumn and assigns it to the DateTime field.
func (o *MicrosoftGraphColumnDefinition) SetDateTime(v AnyOfmicrosoftGraphDateTimeColumn) {
	o.DateTime = &v
}

// SetDateTimeExplicitNull (un)sets DateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DateTime value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetDateTimeExplicitNull(b bool) {
	o.DateTime = nil
	o.isExplicitNullDateTime = b
}
// GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetDefaultValue() AnyOfmicrosoftGraphDefaultColumnValue {
	if o == nil || o.DefaultValue == nil {
		var ret AnyOfmicrosoftGraphDefaultColumnValue
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetDefaultValueOk() (AnyOfmicrosoftGraphDefaultColumnValue, bool) {
	if o == nil || o.DefaultValue == nil {
		var ret AnyOfmicrosoftGraphDefaultColumnValue
		return ret, false
	}
	return *o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given AnyOfmicrosoftGraphDefaultColumnValue and assigns it to the DefaultValue field.
func (o *MicrosoftGraphColumnDefinition) SetDefaultValue(v AnyOfmicrosoftGraphDefaultColumnValue) {
	o.DefaultValue = &v
}

// SetDefaultValueExplicitNull (un)sets DefaultValue to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DefaultValue value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetDefaultValueExplicitNull(b bool) {
	o.DefaultValue = nil
	o.isExplicitNullDefaultValue = b
}
// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphColumnDefinition) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphColumnDefinition) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetEnforceUniqueValues returns the EnforceUniqueValues field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetEnforceUniqueValues() bool {
	if o == nil || o.EnforceUniqueValues == nil {
		var ret bool
		return ret
	}
	return *o.EnforceUniqueValues
}

// GetEnforceUniqueValuesOk returns a tuple with the EnforceUniqueValues field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetEnforceUniqueValuesOk() (bool, bool) {
	if o == nil || o.EnforceUniqueValues == nil {
		var ret bool
		return ret, false
	}
	return *o.EnforceUniqueValues, true
}

// HasEnforceUniqueValues returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasEnforceUniqueValues() bool {
	if o != nil && o.EnforceUniqueValues != nil {
		return true
	}

	return false
}

// SetEnforceUniqueValues gets a reference to the given bool and assigns it to the EnforceUniqueValues field.
func (o *MicrosoftGraphColumnDefinition) SetEnforceUniqueValues(v bool) {
	o.EnforceUniqueValues = &v
}

// SetEnforceUniqueValuesExplicitNull (un)sets EnforceUniqueValues to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The EnforceUniqueValues value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetEnforceUniqueValuesExplicitNull(b bool) {
	o.EnforceUniqueValues = nil
	o.isExplicitNullEnforceUniqueValues = b
}
// GetHidden returns the Hidden field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetHiddenOk() (bool, bool) {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret, false
	}
	return *o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *MicrosoftGraphColumnDefinition) SetHidden(v bool) {
	o.Hidden = &v
}

// SetHiddenExplicitNull (un)sets Hidden to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Hidden value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetHiddenExplicitNull(b bool) {
	o.Hidden = nil
	o.isExplicitNullHidden = b
}
// GetIndexed returns the Indexed field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetIndexed() bool {
	if o == nil || o.Indexed == nil {
		var ret bool
		return ret
	}
	return *o.Indexed
}

// GetIndexedOk returns a tuple with the Indexed field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetIndexedOk() (bool, bool) {
	if o == nil || o.Indexed == nil {
		var ret bool
		return ret, false
	}
	return *o.Indexed, true
}

// HasIndexed returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasIndexed() bool {
	if o != nil && o.Indexed != nil {
		return true
	}

	return false
}

// SetIndexed gets a reference to the given bool and assigns it to the Indexed field.
func (o *MicrosoftGraphColumnDefinition) SetIndexed(v bool) {
	o.Indexed = &v
}

// SetIndexedExplicitNull (un)sets Indexed to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Indexed value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetIndexedExplicitNull(b bool) {
	o.Indexed = nil
	o.isExplicitNullIndexed = b
}
// GetLookup returns the Lookup field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetLookup() AnyOfmicrosoftGraphLookupColumn {
	if o == nil || o.Lookup == nil {
		var ret AnyOfmicrosoftGraphLookupColumn
		return ret
	}
	return *o.Lookup
}

// GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetLookupOk() (AnyOfmicrosoftGraphLookupColumn, bool) {
	if o == nil || o.Lookup == nil {
		var ret AnyOfmicrosoftGraphLookupColumn
		return ret, false
	}
	return *o.Lookup, true
}

// HasLookup returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasLookup() bool {
	if o != nil && o.Lookup != nil {
		return true
	}

	return false
}

// SetLookup gets a reference to the given AnyOfmicrosoftGraphLookupColumn and assigns it to the Lookup field.
func (o *MicrosoftGraphColumnDefinition) SetLookup(v AnyOfmicrosoftGraphLookupColumn) {
	o.Lookup = &v
}

// SetLookupExplicitNull (un)sets Lookup to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Lookup value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetLookupExplicitNull(b bool) {
	o.Lookup = nil
	o.isExplicitNullLookup = b
}
// GetName returns the Name field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftGraphColumnDefinition) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetNumber returns the Number field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetNumber() AnyOfmicrosoftGraphNumberColumn {
	if o == nil || o.Number == nil {
		var ret AnyOfmicrosoftGraphNumberColumn
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetNumberOk() (AnyOfmicrosoftGraphNumberColumn, bool) {
	if o == nil || o.Number == nil {
		var ret AnyOfmicrosoftGraphNumberColumn
		return ret, false
	}
	return *o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfmicrosoftGraphNumberColumn and assigns it to the Number field.
func (o *MicrosoftGraphColumnDefinition) SetNumber(v AnyOfmicrosoftGraphNumberColumn) {
	o.Number = &v
}

// SetNumberExplicitNull (un)sets Number to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Number value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetNumberExplicitNull(b bool) {
	o.Number = nil
	o.isExplicitNullNumber = b
}
// GetPersonOrGroup returns the PersonOrGroup field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetPersonOrGroup() AnyOfmicrosoftGraphPersonOrGroupColumn {
	if o == nil || o.PersonOrGroup == nil {
		var ret AnyOfmicrosoftGraphPersonOrGroupColumn
		return ret
	}
	return *o.PersonOrGroup
}

// GetPersonOrGroupOk returns a tuple with the PersonOrGroup field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetPersonOrGroupOk() (AnyOfmicrosoftGraphPersonOrGroupColumn, bool) {
	if o == nil || o.PersonOrGroup == nil {
		var ret AnyOfmicrosoftGraphPersonOrGroupColumn
		return ret, false
	}
	return *o.PersonOrGroup, true
}

// HasPersonOrGroup returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasPersonOrGroup() bool {
	if o != nil && o.PersonOrGroup != nil {
		return true
	}

	return false
}

// SetPersonOrGroup gets a reference to the given AnyOfmicrosoftGraphPersonOrGroupColumn and assigns it to the PersonOrGroup field.
func (o *MicrosoftGraphColumnDefinition) SetPersonOrGroup(v AnyOfmicrosoftGraphPersonOrGroupColumn) {
	o.PersonOrGroup = &v
}

// SetPersonOrGroupExplicitNull (un)sets PersonOrGroup to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PersonOrGroup value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetPersonOrGroupExplicitNull(b bool) {
	o.PersonOrGroup = nil
	o.isExplicitNullPersonOrGroup = b
}
// GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetReadOnlyOk() (bool, bool) {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret, false
	}
	return *o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *MicrosoftGraphColumnDefinition) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// SetReadOnlyExplicitNull (un)sets ReadOnly to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ReadOnly value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetReadOnlyExplicitNull(b bool) {
	o.ReadOnly = nil
	o.isExplicitNullReadOnly = b
}
// GetRequired returns the Required field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetRequiredOk() (bool, bool) {
	if o == nil || o.Required == nil {
		var ret bool
		return ret, false
	}
	return *o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *MicrosoftGraphColumnDefinition) SetRequired(v bool) {
	o.Required = &v
}

// SetRequiredExplicitNull (un)sets Required to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Required value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetRequiredExplicitNull(b bool) {
	o.Required = nil
	o.isExplicitNullRequired = b
}
// GetText returns the Text field if non-nil, zero value otherwise.
func (o *MicrosoftGraphColumnDefinition) GetText() AnyOfmicrosoftGraphTextColumn {
	if o == nil || o.Text == nil {
		var ret AnyOfmicrosoftGraphTextColumn
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphColumnDefinition) GetTextOk() (AnyOfmicrosoftGraphTextColumn, bool) {
	if o == nil || o.Text == nil {
		var ret AnyOfmicrosoftGraphTextColumn
		return ret, false
	}
	return *o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *MicrosoftGraphColumnDefinition) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given AnyOfmicrosoftGraphTextColumn and assigns it to the Text field.
func (o *MicrosoftGraphColumnDefinition) SetText(v AnyOfmicrosoftGraphTextColumn) {
	o.Text = &v
}

// SetTextExplicitNull (un)sets Text to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Text value is set to nil even if false is passed
func (o *MicrosoftGraphColumnDefinition) SetTextExplicitNull(b bool) {
	o.Text = nil
	o.isExplicitNullText = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphColumnDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Boolean == nil {
		if o.isExplicitNullBoolean {
			toSerialize["boolean"] = o.Boolean
		}
	} else {
		toSerialize["boolean"] = o.Boolean
	}
	if o.Calculated == nil {
		if o.isExplicitNullCalculated {
			toSerialize["calculated"] = o.Calculated
		}
	} else {
		toSerialize["calculated"] = o.Calculated
	}
	if o.Choice == nil {
		if o.isExplicitNullChoice {
			toSerialize["choice"] = o.Choice
		}
	} else {
		toSerialize["choice"] = o.Choice
	}
	if o.ColumnGroup == nil {
		if o.isExplicitNullColumnGroup {
			toSerialize["columnGroup"] = o.ColumnGroup
		}
	} else {
		toSerialize["columnGroup"] = o.ColumnGroup
	}
	if o.Currency == nil {
		if o.isExplicitNullCurrency {
			toSerialize["currency"] = o.Currency
		}
	} else {
		toSerialize["currency"] = o.Currency
	}
	if o.DateTime == nil {
		if o.isExplicitNullDateTime {
			toSerialize["dateTime"] = o.DateTime
		}
	} else {
		toSerialize["dateTime"] = o.DateTime
	}
	if o.DefaultValue == nil {
		if o.isExplicitNullDefaultValue {
			toSerialize["defaultValue"] = o.DefaultValue
		}
	} else {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.Description == nil {
		if o.isExplicitNullDescription {
			toSerialize["description"] = o.Description
		}
	} else {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.EnforceUniqueValues == nil {
		if o.isExplicitNullEnforceUniqueValues {
			toSerialize["enforceUniqueValues"] = o.EnforceUniqueValues
		}
	} else {
		toSerialize["enforceUniqueValues"] = o.EnforceUniqueValues
	}
	if o.Hidden == nil {
		if o.isExplicitNullHidden {
			toSerialize["hidden"] = o.Hidden
		}
	} else {
		toSerialize["hidden"] = o.Hidden
	}
	if o.Indexed == nil {
		if o.isExplicitNullIndexed {
			toSerialize["indexed"] = o.Indexed
		}
	} else {
		toSerialize["indexed"] = o.Indexed
	}
	if o.Lookup == nil {
		if o.isExplicitNullLookup {
			toSerialize["lookup"] = o.Lookup
		}
	} else {
		toSerialize["lookup"] = o.Lookup
	}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
	}
	if o.Number == nil {
		if o.isExplicitNullNumber {
			toSerialize["number"] = o.Number
		}
	} else {
		toSerialize["number"] = o.Number
	}
	if o.PersonOrGroup == nil {
		if o.isExplicitNullPersonOrGroup {
			toSerialize["personOrGroup"] = o.PersonOrGroup
		}
	} else {
		toSerialize["personOrGroup"] = o.PersonOrGroup
	}
	if o.ReadOnly == nil {
		if o.isExplicitNullReadOnly {
			toSerialize["readOnly"] = o.ReadOnly
		}
	} else {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.Required == nil {
		if o.isExplicitNullRequired {
			toSerialize["required"] = o.Required
		}
	} else {
		toSerialize["required"] = o.Required
	}
	if o.Text == nil {
		if o.isExplicitNullText {
			toSerialize["text"] = o.Text
		}
	} else {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}


