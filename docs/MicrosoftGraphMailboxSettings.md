# MicrosoftGraphMailboxSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticRepliesSetting** | Pointer to [**AnyOfmicrosoftGraphAutomaticRepliesSetting**](anyOf&lt;microsoft.graph.automaticRepliesSetting&gt;.md) |  | [optional] 
**ArchiveFolder** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**Language** | Pointer to [**AnyOfmicrosoftGraphLocaleInfo**](anyOf&lt;microsoft.graph.localeInfo&gt;.md) |  | [optional] 
**WorkingHours** | Pointer to [**AnyOfmicrosoftGraphWorkingHours**](anyOf&lt;microsoft.graph.workingHours&gt;.md) |  | [optional] 
**DateFormat** | Pointer to **string** |  | [optional] 
**TimeFormat** | Pointer to **string** |  | [optional] 

## Methods

### GetAutomaticRepliesSetting

`func (o *MicrosoftGraphMailboxSettings) GetAutomaticRepliesSetting() AnyOfmicrosoftGraphAutomaticRepliesSetting`

GetAutomaticRepliesSetting returns the AutomaticRepliesSetting field if non-nil, zero value otherwise.

### GetAutomaticRepliesSettingOk

`func (o *MicrosoftGraphMailboxSettings) GetAutomaticRepliesSettingOk() (AnyOfmicrosoftGraphAutomaticRepliesSetting, bool)`

GetAutomaticRepliesSettingOk returns a tuple with the AutomaticRepliesSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAutomaticRepliesSetting

`func (o *MicrosoftGraphMailboxSettings) HasAutomaticRepliesSetting() bool`

HasAutomaticRepliesSetting returns a boolean if a field has been set.

### SetAutomaticRepliesSetting

`func (o *MicrosoftGraphMailboxSettings) SetAutomaticRepliesSetting(v AnyOfmicrosoftGraphAutomaticRepliesSetting)`

SetAutomaticRepliesSetting gets a reference to the given AnyOfmicrosoftGraphAutomaticRepliesSetting and assigns it to the AutomaticRepliesSetting field.

### SetAutomaticRepliesSettingExplicitNull

`func (o *MicrosoftGraphMailboxSettings) SetAutomaticRepliesSettingExplicitNull(b bool)`

SetAutomaticRepliesSettingExplicitNull (un)sets AutomaticRepliesSetting to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AutomaticRepliesSetting value is set to nil even if false is passed
### GetArchiveFolder

`func (o *MicrosoftGraphMailboxSettings) GetArchiveFolder() string`

GetArchiveFolder returns the ArchiveFolder field if non-nil, zero value otherwise.

### GetArchiveFolderOk

`func (o *MicrosoftGraphMailboxSettings) GetArchiveFolderOk() (string, bool)`

GetArchiveFolderOk returns a tuple with the ArchiveFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasArchiveFolder

`func (o *MicrosoftGraphMailboxSettings) HasArchiveFolder() bool`

HasArchiveFolder returns a boolean if a field has been set.

### SetArchiveFolder

`func (o *MicrosoftGraphMailboxSettings) SetArchiveFolder(v string)`

SetArchiveFolder gets a reference to the given string and assigns it to the ArchiveFolder field.

### SetArchiveFolderExplicitNull

`func (o *MicrosoftGraphMailboxSettings) SetArchiveFolderExplicitNull(b bool)`

SetArchiveFolderExplicitNull (un)sets ArchiveFolder to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ArchiveFolder value is set to nil even if false is passed
### GetTimeZone

`func (o *MicrosoftGraphMailboxSettings) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MicrosoftGraphMailboxSettings) GetTimeZoneOk() (string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeZone

`func (o *MicrosoftGraphMailboxSettings) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZone

`func (o *MicrosoftGraphMailboxSettings) SetTimeZone(v string)`

SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.

### SetTimeZoneExplicitNull

`func (o *MicrosoftGraphMailboxSettings) SetTimeZoneExplicitNull(b bool)`

SetTimeZoneExplicitNull (un)sets TimeZone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TimeZone value is set to nil even if false is passed
### GetLanguage

`func (o *MicrosoftGraphMailboxSettings) GetLanguage() AnyOfmicrosoftGraphLocaleInfo`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MicrosoftGraphMailboxSettings) GetLanguageOk() (AnyOfmicrosoftGraphLocaleInfo, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLanguage

`func (o *MicrosoftGraphMailboxSettings) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguage

`func (o *MicrosoftGraphMailboxSettings) SetLanguage(v AnyOfmicrosoftGraphLocaleInfo)`

SetLanguage gets a reference to the given AnyOfmicrosoftGraphLocaleInfo and assigns it to the Language field.

### SetLanguageExplicitNull

`func (o *MicrosoftGraphMailboxSettings) SetLanguageExplicitNull(b bool)`

SetLanguageExplicitNull (un)sets Language to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Language value is set to nil even if false is passed
### GetWorkingHours

`func (o *MicrosoftGraphMailboxSettings) GetWorkingHours() AnyOfmicrosoftGraphWorkingHours`

GetWorkingHours returns the WorkingHours field if non-nil, zero value otherwise.

### GetWorkingHoursOk

`func (o *MicrosoftGraphMailboxSettings) GetWorkingHoursOk() (AnyOfmicrosoftGraphWorkingHours, bool)`

GetWorkingHoursOk returns a tuple with the WorkingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorkingHours

`func (o *MicrosoftGraphMailboxSettings) HasWorkingHours() bool`

HasWorkingHours returns a boolean if a field has been set.

### SetWorkingHours

`func (o *MicrosoftGraphMailboxSettings) SetWorkingHours(v AnyOfmicrosoftGraphWorkingHours)`

SetWorkingHours gets a reference to the given AnyOfmicrosoftGraphWorkingHours and assigns it to the WorkingHours field.

### SetWorkingHoursExplicitNull

`func (o *MicrosoftGraphMailboxSettings) SetWorkingHoursExplicitNull(b bool)`

SetWorkingHoursExplicitNull (un)sets WorkingHours to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WorkingHours value is set to nil even if false is passed
### GetDateFormat

`func (o *MicrosoftGraphMailboxSettings) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *MicrosoftGraphMailboxSettings) GetDateFormatOk() (string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDateFormat

`func (o *MicrosoftGraphMailboxSettings) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### SetDateFormat

`func (o *MicrosoftGraphMailboxSettings) SetDateFormat(v string)`

SetDateFormat gets a reference to the given string and assigns it to the DateFormat field.

### SetDateFormatExplicitNull

`func (o *MicrosoftGraphMailboxSettings) SetDateFormatExplicitNull(b bool)`

SetDateFormatExplicitNull (un)sets DateFormat to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DateFormat value is set to nil even if false is passed
### GetTimeFormat

`func (o *MicrosoftGraphMailboxSettings) GetTimeFormat() string`

GetTimeFormat returns the TimeFormat field if non-nil, zero value otherwise.

### GetTimeFormatOk

`func (o *MicrosoftGraphMailboxSettings) GetTimeFormatOk() (string, bool)`

GetTimeFormatOk returns a tuple with the TimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeFormat

`func (o *MicrosoftGraphMailboxSettings) HasTimeFormat() bool`

HasTimeFormat returns a boolean if a field has been set.

### SetTimeFormat

`func (o *MicrosoftGraphMailboxSettings) SetTimeFormat(v string)`

SetTimeFormat gets a reference to the given string and assigns it to the TimeFormat field.

### SetTimeFormatExplicitNull

`func (o *MicrosoftGraphMailboxSettings) SetTimeFormatExplicitNull(b bool)`

SetTimeFormatExplicitNull (un)sets TimeFormat to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TimeFormat value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


