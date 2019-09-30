# MicrosoftGraphIntuneBrand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Company/organization name that is displayed to end users. | [optional] 
**ContactITName** | Pointer to **string** | Name of the person/organization responsible for IT support. | [optional] 
**ContactITPhoneNumber** | Pointer to **string** | Phone number of the person/organization responsible for IT support. | [optional] 
**ContactITEmailAddress** | Pointer to **string** | Email address of the person/organization responsible for IT support. | [optional] 
**ContactITNotes** | Pointer to **string** | Text comments regarding the person/organization responsible for IT support. | [optional] 
**PrivacyUrl** | Pointer to **string** | URL to the company/organization’s privacy policy. | [optional] 
**OnlineSupportSiteUrl** | Pointer to **string** | URL to the company/organization’s IT helpdesk site. | [optional] 
**OnlineSupportSiteName** | Pointer to **string** | Display name of the company/organization’s IT helpdesk site. | [optional] 
**ThemeColor** | Pointer to [**AnyOfmicrosoftGraphRgbColor**](anyOf&lt;microsoft.graph.rgbColor&gt;.md) | Primary theme color used in the Company Portal applications and web portal. | [optional] 
**ShowLogo** | Pointer to **bool** | Boolean that represents whether the administrator-supplied logo images are shown or not shown. | [optional] 
**LightBackgroundLogo** | Pointer to [**AnyOfmicrosoftGraphMimeContent**](anyOf&lt;microsoft.graph.mimeContent&gt;.md) | Logo image displayed in Company Portal apps which have a light background behind the logo. | [optional] 
**DarkBackgroundLogo** | Pointer to [**AnyOfmicrosoftGraphMimeContent**](anyOf&lt;microsoft.graph.mimeContent&gt;.md) | Logo image displayed in Company Portal apps which have a dark background behind the logo. | [optional] 
**ShowNameNextToLogo** | Pointer to **bool** | Boolean that represents whether the administrator-supplied display name will be shown next to the logo image. | [optional] 
**ShowDisplayNameNextToLogo** | Pointer to **bool** | Boolean that represents whether the administrator-supplied display name will be shown next to the logo image. | [optional] 

## Methods

### GetDisplayName

`func (o *MicrosoftGraphIntuneBrand) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIntuneBrand) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphIntuneBrand) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphIntuneBrand) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphIntuneBrand) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetContactITName

`func (o *MicrosoftGraphIntuneBrand) GetContactITName() string`

GetContactITName returns the ContactITName field if non-nil, zero value otherwise.

### GetContactITNameOk

`func (o *MicrosoftGraphIntuneBrand) GetContactITNameOk() (string, bool)`

GetContactITNameOk returns a tuple with the ContactITName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContactITName

`func (o *MicrosoftGraphIntuneBrand) HasContactITName() bool`

HasContactITName returns a boolean if a field has been set.

### SetContactITName

`func (o *MicrosoftGraphIntuneBrand) SetContactITName(v string)`

SetContactITName gets a reference to the given string and assigns it to the ContactITName field.

### SetContactITNameExplicitNull

`func (o *MicrosoftGraphIntuneBrand) SetContactITNameExplicitNull(b bool)`

SetContactITNameExplicitNull (un)sets ContactITName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContactITName value is set to nil even if false is passed
### GetContactITPhoneNumber

`func (o *MicrosoftGraphIntuneBrand) GetContactITPhoneNumber() string`

GetContactITPhoneNumber returns the ContactITPhoneNumber field if non-nil, zero value otherwise.

### GetContactITPhoneNumberOk

`func (o *MicrosoftGraphIntuneBrand) GetContactITPhoneNumberOk() (string, bool)`

GetContactITPhoneNumberOk returns a tuple with the ContactITPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContactITPhoneNumber

`func (o *MicrosoftGraphIntuneBrand) HasContactITPhoneNumber() bool`

HasContactITPhoneNumber returns a boolean if a field has been set.

### SetContactITPhoneNumber

`func (o *MicrosoftGraphIntuneBrand) SetContactITPhoneNumber(v string)`

SetContactITPhoneNumber gets a reference to the given string and assigns it to the ContactITPhoneNumber field.

### SetContactITPhoneNumberExplicitNull

`func (o *MicrosoftGraphIntuneBrand) SetContactITPhoneNumberExplicitNull(b bool)`

SetContactITPhoneNumberExplicitNull (un)sets ContactITPhoneNumber to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContactITPhoneNumber value is set to nil even if false is passed
### GetContactITEmailAddress

`func (o *MicrosoftGraphIntuneBrand) GetContactITEmailAddress() string`

GetContactITEmailAddress returns the ContactITEmailAddress field if non-nil, zero value otherwise.

### GetContactITEmailAddressOk

`func (o *MicrosoftGraphIntuneBrand) GetContactITEmailAddressOk() (string, bool)`

GetContactITEmailAddressOk returns a tuple with the ContactITEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContactITEmailAddress

`func (o *MicrosoftGraphIntuneBrand) HasContactITEmailAddress() bool`

HasContactITEmailAddress returns a boolean if a field has been set.

### SetContactITEmailAddress

`func (o *MicrosoftGraphIntuneBrand) SetContactITEmailAddress(v string)`

SetContactITEmailAddress gets a reference to the given string and assigns it to the ContactITEmailAddress field.

### SetContactITEmailAddressExplicitNull

`func (o *MicrosoftGraphIntuneBrand) SetContactITEmailAddressExplicitNull(b bool)`

SetContactITEmailAddressExplicitNull (un)sets ContactITEmailAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContactITEmailAddress value is set to nil even if false is passed
### GetContactITNotes

`func (o *MicrosoftGraphIntuneBrand) GetContactITNotes() string`

GetContactITNotes returns the ContactITNotes field if non-nil, zero value otherwise.

### GetContactITNotesOk

`func (o *MicrosoftGraphIntuneBrand) GetContactITNotesOk() (string, bool)`

GetContactITNotesOk returns a tuple with the ContactITNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContactITNotes

`func (o *MicrosoftGraphIntuneBrand) HasContactITNotes() bool`

HasContactITNotes returns a boolean if a field has been set.

### SetContactITNotes

`func (o *MicrosoftGraphIntuneBrand) SetContactITNotes(v string)`

SetContactITNotes gets a reference to the given string and assigns it to the ContactITNotes field.

### SetContactITNotesExplicitNull

`func (o *MicrosoftGraphIntuneBrand) SetContactITNotesExplicitNull(b bool)`

SetContactITNotesExplicitNull (un)sets ContactITNotes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ContactITNotes value is set to nil even if false is passed
### GetPrivacyUrl

`func (o *MicrosoftGraphIntuneBrand) GetPrivacyUrl() string`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *MicrosoftGraphIntuneBrand) GetPrivacyUrlOk() (string, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivacyUrl

`func (o *MicrosoftGraphIntuneBrand) HasPrivacyUrl() bool`

HasPrivacyUrl returns a boolean if a field has been set.

### SetPrivacyUrl

`func (o *MicrosoftGraphIntuneBrand) SetPrivacyUrl(v string)`

SetPrivacyUrl gets a reference to the given string and assigns it to the PrivacyUrl field.

### SetPrivacyUrlExplicitNull

`func (o *MicrosoftGraphIntuneBrand) SetPrivacyUrlExplicitNull(b bool)`

SetPrivacyUrlExplicitNull (un)sets PrivacyUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PrivacyUrl value is set to nil even if false is passed
### GetOnlineSupportSiteUrl

`func (o *MicrosoftGraphIntuneBrand) GetOnlineSupportSiteUrl() string`

GetOnlineSupportSiteUrl returns the OnlineSupportSiteUrl field if non-nil, zero value otherwise.

### GetOnlineSupportSiteUrlOk

`func (o *MicrosoftGraphIntuneBrand) GetOnlineSupportSiteUrlOk() (string, bool)`

GetOnlineSupportSiteUrlOk returns a tuple with the OnlineSupportSiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnlineSupportSiteUrl

`func (o *MicrosoftGraphIntuneBrand) HasOnlineSupportSiteUrl() bool`

HasOnlineSupportSiteUrl returns a boolean if a field has been set.

### SetOnlineSupportSiteUrl

`func (o *MicrosoftGraphIntuneBrand) SetOnlineSupportSiteUrl(v string)`

SetOnlineSupportSiteUrl gets a reference to the given string and assigns it to the OnlineSupportSiteUrl field.

### SetOnlineSupportSiteUrlExplicitNull

`func (o *MicrosoftGraphIntuneBrand) SetOnlineSupportSiteUrlExplicitNull(b bool)`

SetOnlineSupportSiteUrlExplicitNull (un)sets OnlineSupportSiteUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnlineSupportSiteUrl value is set to nil even if false is passed
### GetOnlineSupportSiteName

`func (o *MicrosoftGraphIntuneBrand) GetOnlineSupportSiteName() string`

GetOnlineSupportSiteName returns the OnlineSupportSiteName field if non-nil, zero value otherwise.

### GetOnlineSupportSiteNameOk

`func (o *MicrosoftGraphIntuneBrand) GetOnlineSupportSiteNameOk() (string, bool)`

GetOnlineSupportSiteNameOk returns a tuple with the OnlineSupportSiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnlineSupportSiteName

`func (o *MicrosoftGraphIntuneBrand) HasOnlineSupportSiteName() bool`

HasOnlineSupportSiteName returns a boolean if a field has been set.

### SetOnlineSupportSiteName

`func (o *MicrosoftGraphIntuneBrand) SetOnlineSupportSiteName(v string)`

SetOnlineSupportSiteName gets a reference to the given string and assigns it to the OnlineSupportSiteName field.

### SetOnlineSupportSiteNameExplicitNull

`func (o *MicrosoftGraphIntuneBrand) SetOnlineSupportSiteNameExplicitNull(b bool)`

SetOnlineSupportSiteNameExplicitNull (un)sets OnlineSupportSiteName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnlineSupportSiteName value is set to nil even if false is passed
### GetThemeColor

`func (o *MicrosoftGraphIntuneBrand) GetThemeColor() AnyOfmicrosoftGraphRgbColor`

GetThemeColor returns the ThemeColor field if non-nil, zero value otherwise.

### GetThemeColorOk

`func (o *MicrosoftGraphIntuneBrand) GetThemeColorOk() (AnyOfmicrosoftGraphRgbColor, bool)`

GetThemeColorOk returns a tuple with the ThemeColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThemeColor

`func (o *MicrosoftGraphIntuneBrand) HasThemeColor() bool`

HasThemeColor returns a boolean if a field has been set.

### SetThemeColor

`func (o *MicrosoftGraphIntuneBrand) SetThemeColor(v AnyOfmicrosoftGraphRgbColor)`

SetThemeColor gets a reference to the given AnyOfmicrosoftGraphRgbColor and assigns it to the ThemeColor field.

### SetThemeColorExplicitNull

`func (o *MicrosoftGraphIntuneBrand) SetThemeColorExplicitNull(b bool)`

SetThemeColorExplicitNull (un)sets ThemeColor to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ThemeColor value is set to nil even if false is passed
### GetShowLogo

`func (o *MicrosoftGraphIntuneBrand) GetShowLogo() bool`

GetShowLogo returns the ShowLogo field if non-nil, zero value otherwise.

### GetShowLogoOk

`func (o *MicrosoftGraphIntuneBrand) GetShowLogoOk() (bool, bool)`

GetShowLogoOk returns a tuple with the ShowLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowLogo

`func (o *MicrosoftGraphIntuneBrand) HasShowLogo() bool`

HasShowLogo returns a boolean if a field has been set.

### SetShowLogo

`func (o *MicrosoftGraphIntuneBrand) SetShowLogo(v bool)`

SetShowLogo gets a reference to the given bool and assigns it to the ShowLogo field.

### GetLightBackgroundLogo

`func (o *MicrosoftGraphIntuneBrand) GetLightBackgroundLogo() AnyOfmicrosoftGraphMimeContent`

GetLightBackgroundLogo returns the LightBackgroundLogo field if non-nil, zero value otherwise.

### GetLightBackgroundLogoOk

`func (o *MicrosoftGraphIntuneBrand) GetLightBackgroundLogoOk() (AnyOfmicrosoftGraphMimeContent, bool)`

GetLightBackgroundLogoOk returns a tuple with the LightBackgroundLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLightBackgroundLogo

`func (o *MicrosoftGraphIntuneBrand) HasLightBackgroundLogo() bool`

HasLightBackgroundLogo returns a boolean if a field has been set.

### SetLightBackgroundLogo

`func (o *MicrosoftGraphIntuneBrand) SetLightBackgroundLogo(v AnyOfmicrosoftGraphMimeContent)`

SetLightBackgroundLogo gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the LightBackgroundLogo field.

### SetLightBackgroundLogoExplicitNull

`func (o *MicrosoftGraphIntuneBrand) SetLightBackgroundLogoExplicitNull(b bool)`

SetLightBackgroundLogoExplicitNull (un)sets LightBackgroundLogo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LightBackgroundLogo value is set to nil even if false is passed
### GetDarkBackgroundLogo

`func (o *MicrosoftGraphIntuneBrand) GetDarkBackgroundLogo() AnyOfmicrosoftGraphMimeContent`

GetDarkBackgroundLogo returns the DarkBackgroundLogo field if non-nil, zero value otherwise.

### GetDarkBackgroundLogoOk

`func (o *MicrosoftGraphIntuneBrand) GetDarkBackgroundLogoOk() (AnyOfmicrosoftGraphMimeContent, bool)`

GetDarkBackgroundLogoOk returns a tuple with the DarkBackgroundLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDarkBackgroundLogo

`func (o *MicrosoftGraphIntuneBrand) HasDarkBackgroundLogo() bool`

HasDarkBackgroundLogo returns a boolean if a field has been set.

### SetDarkBackgroundLogo

`func (o *MicrosoftGraphIntuneBrand) SetDarkBackgroundLogo(v AnyOfmicrosoftGraphMimeContent)`

SetDarkBackgroundLogo gets a reference to the given AnyOfmicrosoftGraphMimeContent and assigns it to the DarkBackgroundLogo field.

### SetDarkBackgroundLogoExplicitNull

`func (o *MicrosoftGraphIntuneBrand) SetDarkBackgroundLogoExplicitNull(b bool)`

SetDarkBackgroundLogoExplicitNull (un)sets DarkBackgroundLogo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DarkBackgroundLogo value is set to nil even if false is passed
### GetShowNameNextToLogo

`func (o *MicrosoftGraphIntuneBrand) GetShowNameNextToLogo() bool`

GetShowNameNextToLogo returns the ShowNameNextToLogo field if non-nil, zero value otherwise.

### GetShowNameNextToLogoOk

`func (o *MicrosoftGraphIntuneBrand) GetShowNameNextToLogoOk() (bool, bool)`

GetShowNameNextToLogoOk returns a tuple with the ShowNameNextToLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowNameNextToLogo

`func (o *MicrosoftGraphIntuneBrand) HasShowNameNextToLogo() bool`

HasShowNameNextToLogo returns a boolean if a field has been set.

### SetShowNameNextToLogo

`func (o *MicrosoftGraphIntuneBrand) SetShowNameNextToLogo(v bool)`

SetShowNameNextToLogo gets a reference to the given bool and assigns it to the ShowNameNextToLogo field.

### GetShowDisplayNameNextToLogo

`func (o *MicrosoftGraphIntuneBrand) GetShowDisplayNameNextToLogo() bool`

GetShowDisplayNameNextToLogo returns the ShowDisplayNameNextToLogo field if non-nil, zero value otherwise.

### GetShowDisplayNameNextToLogoOk

`func (o *MicrosoftGraphIntuneBrand) GetShowDisplayNameNextToLogoOk() (bool, bool)`

GetShowDisplayNameNextToLogoOk returns a tuple with the ShowDisplayNameNextToLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowDisplayNameNextToLogo

`func (o *MicrosoftGraphIntuneBrand) HasShowDisplayNameNextToLogo() bool`

HasShowDisplayNameNextToLogo returns a boolean if a field has been set.

### SetShowDisplayNameNextToLogo

`func (o *MicrosoftGraphIntuneBrand) SetShowDisplayNameNextToLogo(v bool)`

SetShowDisplayNameNextToLogo gets a reference to the given bool and assigns it to the ShowDisplayNameNextToLogo field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


