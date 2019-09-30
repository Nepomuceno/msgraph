# ManagedApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppAvailability** | Pointer to [**AnyOfmicrosoftGraphManagedAppAvailability**](anyOf&lt;microsoft.graph.managedAppAvailability&gt;.md) | The Application&#39;s availability. | [optional] 
**Version** | Pointer to **string** | The Application&#39;s version. | [optional] 

## Methods

### GetAppAvailability

`func (o *ManagedApp) GetAppAvailability() AnyOfmicrosoftGraphManagedAppAvailability`

GetAppAvailability returns the AppAvailability field if non-nil, zero value otherwise.

### GetAppAvailabilityOk

`func (o *ManagedApp) GetAppAvailabilityOk() (AnyOfmicrosoftGraphManagedAppAvailability, bool)`

GetAppAvailabilityOk returns a tuple with the AppAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppAvailability

`func (o *ManagedApp) HasAppAvailability() bool`

HasAppAvailability returns a boolean if a field has been set.

### SetAppAvailability

`func (o *ManagedApp) SetAppAvailability(v AnyOfmicrosoftGraphManagedAppAvailability)`

SetAppAvailability gets a reference to the given AnyOfmicrosoftGraphManagedAppAvailability and assigns it to the AppAvailability field.

### GetVersion

`func (o *ManagedApp) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagedApp) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *ManagedApp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *ManagedApp) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *ManagedApp) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


