# NotificationMessageTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was last modified. | [optional] 
**DisplayName** | Pointer to **string** | Display name for the Notification Message Template. | [optional] 
**DefaultLocale** | Pointer to **string** | The default locale to fallback onto when the requested locale is not available. | [optional] 
**BrandingOptions** | Pointer to [**AnyOfmicrosoftGraphNotificationTemplateBrandingOptions**](anyOf&lt;microsoft.graph.notificationTemplateBrandingOptions&gt;.md) | The Message Template Branding Options. Branding is defined in the Intune Admin Console. | [optional] 
**LocalizedNotificationMessages** | Pointer to [**[]MicrosoftGraphLocalizedNotificationMessage**](microsoft.graph.localizedNotificationMessage.md) |  | [optional] 

## Methods

### GetLastModifiedDateTime

`func (o *NotificationMessageTemplate) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *NotificationMessageTemplate) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *NotificationMessageTemplate) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *NotificationMessageTemplate) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetDisplayName

`func (o *NotificationMessageTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NotificationMessageTemplate) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *NotificationMessageTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *NotificationMessageTemplate) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDefaultLocale

`func (o *NotificationMessageTemplate) GetDefaultLocale() string`

GetDefaultLocale returns the DefaultLocale field if non-nil, zero value otherwise.

### GetDefaultLocaleOk

`func (o *NotificationMessageTemplate) GetDefaultLocaleOk() (string, bool)`

GetDefaultLocaleOk returns a tuple with the DefaultLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultLocale

`func (o *NotificationMessageTemplate) HasDefaultLocale() bool`

HasDefaultLocale returns a boolean if a field has been set.

### SetDefaultLocale

`func (o *NotificationMessageTemplate) SetDefaultLocale(v string)`

SetDefaultLocale gets a reference to the given string and assigns it to the DefaultLocale field.

### SetDefaultLocaleExplicitNull

`func (o *NotificationMessageTemplate) SetDefaultLocaleExplicitNull(b bool)`

SetDefaultLocaleExplicitNull (un)sets DefaultLocale to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefaultLocale value is set to nil even if false is passed
### GetBrandingOptions

`func (o *NotificationMessageTemplate) GetBrandingOptions() AnyOfmicrosoftGraphNotificationTemplateBrandingOptions`

GetBrandingOptions returns the BrandingOptions field if non-nil, zero value otherwise.

### GetBrandingOptionsOk

`func (o *NotificationMessageTemplate) GetBrandingOptionsOk() (AnyOfmicrosoftGraphNotificationTemplateBrandingOptions, bool)`

GetBrandingOptionsOk returns a tuple with the BrandingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBrandingOptions

`func (o *NotificationMessageTemplate) HasBrandingOptions() bool`

HasBrandingOptions returns a boolean if a field has been set.

### SetBrandingOptions

`func (o *NotificationMessageTemplate) SetBrandingOptions(v AnyOfmicrosoftGraphNotificationTemplateBrandingOptions)`

SetBrandingOptions gets a reference to the given AnyOfmicrosoftGraphNotificationTemplateBrandingOptions and assigns it to the BrandingOptions field.

### GetLocalizedNotificationMessages

`func (o *NotificationMessageTemplate) GetLocalizedNotificationMessages() []MicrosoftGraphLocalizedNotificationMessage`

GetLocalizedNotificationMessages returns the LocalizedNotificationMessages field if non-nil, zero value otherwise.

### GetLocalizedNotificationMessagesOk

`func (o *NotificationMessageTemplate) GetLocalizedNotificationMessagesOk() ([]MicrosoftGraphLocalizedNotificationMessage, bool)`

GetLocalizedNotificationMessagesOk returns a tuple with the LocalizedNotificationMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocalizedNotificationMessages

`func (o *NotificationMessageTemplate) HasLocalizedNotificationMessages() bool`

HasLocalizedNotificationMessages returns a boolean if a field has been set.

### SetLocalizedNotificationMessages

`func (o *NotificationMessageTemplate) SetLocalizedNotificationMessages(v []MicrosoftGraphLocalizedNotificationMessage)`

SetLocalizedNotificationMessages gets a reference to the given []MicrosoftGraphLocalizedNotificationMessage and assigns it to the LocalizedNotificationMessages field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


