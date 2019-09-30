# MicrosoftGraphManagedAppPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | Policy display name. | [optional] 
**Description** | Pointer to **string** | The policy&#39;s description. | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | The date and time the policy was created. | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | Last time the policy was modified. | [optional] 
**Version** | Pointer to **string** | Version of the entity. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphManagedAppPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedAppPolicy) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphManagedAppPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphManagedAppPolicy) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphManagedAppPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphManagedAppPolicy) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphManagedAppPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphManagedAppPolicy) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphManagedAppPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphManagedAppPolicy) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphManagedAppPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphManagedAppPolicy) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphManagedAppPolicy) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphManagedAppPolicy) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphManagedAppPolicy) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphManagedAppPolicy) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphManagedAppPolicy) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphManagedAppPolicy) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphManagedAppPolicy) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphManagedAppPolicy) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphManagedAppPolicy) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetVersion

`func (o *MicrosoftGraphManagedAppPolicy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphManagedAppPolicy) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphManagedAppPolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphManagedAppPolicy) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphManagedAppPolicy) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


