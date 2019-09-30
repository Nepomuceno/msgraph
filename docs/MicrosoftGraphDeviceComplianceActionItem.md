# MicrosoftGraphDeviceComplianceActionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GracePeriodHours** | Pointer to **int32** | Number of hours to wait till the action will be enforced. Valid values 0 to 8760 | [optional] 
**ActionType** | Pointer to [**AnyOfmicrosoftGraphDeviceComplianceActionType**](anyOf&lt;microsoft.graph.deviceComplianceActionType&gt;.md) | What action to take | [optional] 
**NotificationTemplateId** | Pointer to **string** | What notification Message template to use | [optional] 
**NotificationMessageCCList** | Pointer to **[]string** | A list of group IDs to speicify who to CC this notification message to. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceComplianceActionItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetGracePeriodHours

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetGracePeriodHours() int32`

GetGracePeriodHours returns the GracePeriodHours field if non-nil, zero value otherwise.

### GetGracePeriodHoursOk

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetGracePeriodHoursOk() (int32, bool)`

GetGracePeriodHoursOk returns a tuple with the GracePeriodHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGracePeriodHours

`func (o *MicrosoftGraphDeviceComplianceActionItem) HasGracePeriodHours() bool`

HasGracePeriodHours returns a boolean if a field has been set.

### SetGracePeriodHours

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetGracePeriodHours(v int32)`

SetGracePeriodHours gets a reference to the given int32 and assigns it to the GracePeriodHours field.

### GetActionType

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetActionType() AnyOfmicrosoftGraphDeviceComplianceActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetActionTypeOk() (AnyOfmicrosoftGraphDeviceComplianceActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActionType

`func (o *MicrosoftGraphDeviceComplianceActionItem) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### SetActionType

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetActionType(v AnyOfmicrosoftGraphDeviceComplianceActionType)`

SetActionType gets a reference to the given AnyOfmicrosoftGraphDeviceComplianceActionType and assigns it to the ActionType field.

### GetNotificationTemplateId

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetNotificationTemplateId() string`

GetNotificationTemplateId returns the NotificationTemplateId field if non-nil, zero value otherwise.

### GetNotificationTemplateIdOk

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetNotificationTemplateIdOk() (string, bool)`

GetNotificationTemplateIdOk returns a tuple with the NotificationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationTemplateId

`func (o *MicrosoftGraphDeviceComplianceActionItem) HasNotificationTemplateId() bool`

HasNotificationTemplateId returns a boolean if a field has been set.

### SetNotificationTemplateId

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetNotificationTemplateId(v string)`

SetNotificationTemplateId gets a reference to the given string and assigns it to the NotificationTemplateId field.

### SetNotificationTemplateIdExplicitNull

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetNotificationTemplateIdExplicitNull(b bool)`

SetNotificationTemplateIdExplicitNull (un)sets NotificationTemplateId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NotificationTemplateId value is set to nil even if false is passed
### GetNotificationMessageCCList

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetNotificationMessageCCList() []string`

GetNotificationMessageCCList returns the NotificationMessageCCList field if non-nil, zero value otherwise.

### GetNotificationMessageCCListOk

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetNotificationMessageCCListOk() ([]string, bool)`

GetNotificationMessageCCListOk returns a tuple with the NotificationMessageCCList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationMessageCCList

`func (o *MicrosoftGraphDeviceComplianceActionItem) HasNotificationMessageCCList() bool`

HasNotificationMessageCCList returns a boolean if a field has been set.

### SetNotificationMessageCCList

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetNotificationMessageCCList(v []string)`

SetNotificationMessageCCList gets a reference to the given []string and assigns it to the NotificationMessageCCList field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


