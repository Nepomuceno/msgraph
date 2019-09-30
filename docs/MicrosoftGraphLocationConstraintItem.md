# MicrosoftGraphLocationConstraintItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**LocationEmailAddress** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) |  | [optional] 
**LocationUri** | Pointer to **string** |  | [optional] 
**Coordinates** | Pointer to [**AnyOfmicrosoftGraphOutlookGeoCoordinates**](anyOf&lt;microsoft.graph.outlookGeoCoordinates&gt;.md) |  | [optional] 
**LocationType** | Pointer to [**AnyOfmicrosoftGraphLocationType**](anyOf&lt;microsoft.graph.locationType&gt;.md) |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**UniqueIdType** | Pointer to [**AnyOfmicrosoftGraphLocationUniqueIdType**](anyOf&lt;microsoft.graph.locationUniqueIdType&gt;.md) |  | [optional] 
**ResolveAvailability** | Pointer to **bool** |  | [optional] 

## Methods

### GetDisplayName

`func (o *MicrosoftGraphLocationConstraintItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphLocationConstraintItem) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphLocationConstraintItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphLocationConstraintItem) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphLocationConstraintItem) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetLocationEmailAddress

`func (o *MicrosoftGraphLocationConstraintItem) GetLocationEmailAddress() string`

GetLocationEmailAddress returns the LocationEmailAddress field if non-nil, zero value otherwise.

### GetLocationEmailAddressOk

`func (o *MicrosoftGraphLocationConstraintItem) GetLocationEmailAddressOk() (string, bool)`

GetLocationEmailAddressOk returns a tuple with the LocationEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocationEmailAddress

`func (o *MicrosoftGraphLocationConstraintItem) HasLocationEmailAddress() bool`

HasLocationEmailAddress returns a boolean if a field has been set.

### SetLocationEmailAddress

`func (o *MicrosoftGraphLocationConstraintItem) SetLocationEmailAddress(v string)`

SetLocationEmailAddress gets a reference to the given string and assigns it to the LocationEmailAddress field.

### SetLocationEmailAddressExplicitNull

`func (o *MicrosoftGraphLocationConstraintItem) SetLocationEmailAddressExplicitNull(b bool)`

SetLocationEmailAddressExplicitNull (un)sets LocationEmailAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LocationEmailAddress value is set to nil even if false is passed
### GetAddress

`func (o *MicrosoftGraphLocationConstraintItem) GetAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphLocationConstraintItem) GetAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddress

`func (o *MicrosoftGraphLocationConstraintItem) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddress

`func (o *MicrosoftGraphLocationConstraintItem) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the Address field.

### SetAddressExplicitNull

`func (o *MicrosoftGraphLocationConstraintItem) SetAddressExplicitNull(b bool)`

SetAddressExplicitNull (un)sets Address to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Address value is set to nil even if false is passed
### GetLocationUri

`func (o *MicrosoftGraphLocationConstraintItem) GetLocationUri() string`

GetLocationUri returns the LocationUri field if non-nil, zero value otherwise.

### GetLocationUriOk

`func (o *MicrosoftGraphLocationConstraintItem) GetLocationUriOk() (string, bool)`

GetLocationUriOk returns a tuple with the LocationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocationUri

`func (o *MicrosoftGraphLocationConstraintItem) HasLocationUri() bool`

HasLocationUri returns a boolean if a field has been set.

### SetLocationUri

`func (o *MicrosoftGraphLocationConstraintItem) SetLocationUri(v string)`

SetLocationUri gets a reference to the given string and assigns it to the LocationUri field.

### SetLocationUriExplicitNull

`func (o *MicrosoftGraphLocationConstraintItem) SetLocationUriExplicitNull(b bool)`

SetLocationUriExplicitNull (un)sets LocationUri to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LocationUri value is set to nil even if false is passed
### GetCoordinates

`func (o *MicrosoftGraphLocationConstraintItem) GetCoordinates() AnyOfmicrosoftGraphOutlookGeoCoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *MicrosoftGraphLocationConstraintItem) GetCoordinatesOk() (AnyOfmicrosoftGraphOutlookGeoCoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCoordinates

`func (o *MicrosoftGraphLocationConstraintItem) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### SetCoordinates

`func (o *MicrosoftGraphLocationConstraintItem) SetCoordinates(v AnyOfmicrosoftGraphOutlookGeoCoordinates)`

SetCoordinates gets a reference to the given AnyOfmicrosoftGraphOutlookGeoCoordinates and assigns it to the Coordinates field.

### SetCoordinatesExplicitNull

`func (o *MicrosoftGraphLocationConstraintItem) SetCoordinatesExplicitNull(b bool)`

SetCoordinatesExplicitNull (un)sets Coordinates to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Coordinates value is set to nil even if false is passed
### GetLocationType

`func (o *MicrosoftGraphLocationConstraintItem) GetLocationType() AnyOfmicrosoftGraphLocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *MicrosoftGraphLocationConstraintItem) GetLocationTypeOk() (AnyOfmicrosoftGraphLocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocationType

`func (o *MicrosoftGraphLocationConstraintItem) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### SetLocationType

`func (o *MicrosoftGraphLocationConstraintItem) SetLocationType(v AnyOfmicrosoftGraphLocationType)`

SetLocationType gets a reference to the given AnyOfmicrosoftGraphLocationType and assigns it to the LocationType field.

### SetLocationTypeExplicitNull

`func (o *MicrosoftGraphLocationConstraintItem) SetLocationTypeExplicitNull(b bool)`

SetLocationTypeExplicitNull (un)sets LocationType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LocationType value is set to nil even if false is passed
### GetUniqueId

`func (o *MicrosoftGraphLocationConstraintItem) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *MicrosoftGraphLocationConstraintItem) GetUniqueIdOk() (string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueId

`func (o *MicrosoftGraphLocationConstraintItem) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueId

`func (o *MicrosoftGraphLocationConstraintItem) SetUniqueId(v string)`

SetUniqueId gets a reference to the given string and assigns it to the UniqueId field.

### SetUniqueIdExplicitNull

`func (o *MicrosoftGraphLocationConstraintItem) SetUniqueIdExplicitNull(b bool)`

SetUniqueIdExplicitNull (un)sets UniqueId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UniqueId value is set to nil even if false is passed
### GetUniqueIdType

`func (o *MicrosoftGraphLocationConstraintItem) GetUniqueIdType() AnyOfmicrosoftGraphLocationUniqueIdType`

GetUniqueIdType returns the UniqueIdType field if non-nil, zero value otherwise.

### GetUniqueIdTypeOk

`func (o *MicrosoftGraphLocationConstraintItem) GetUniqueIdTypeOk() (AnyOfmicrosoftGraphLocationUniqueIdType, bool)`

GetUniqueIdTypeOk returns a tuple with the UniqueIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueIdType

`func (o *MicrosoftGraphLocationConstraintItem) HasUniqueIdType() bool`

HasUniqueIdType returns a boolean if a field has been set.

### SetUniqueIdType

`func (o *MicrosoftGraphLocationConstraintItem) SetUniqueIdType(v AnyOfmicrosoftGraphLocationUniqueIdType)`

SetUniqueIdType gets a reference to the given AnyOfmicrosoftGraphLocationUniqueIdType and assigns it to the UniqueIdType field.

### SetUniqueIdTypeExplicitNull

`func (o *MicrosoftGraphLocationConstraintItem) SetUniqueIdTypeExplicitNull(b bool)`

SetUniqueIdTypeExplicitNull (un)sets UniqueIdType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UniqueIdType value is set to nil even if false is passed
### GetResolveAvailability

`func (o *MicrosoftGraphLocationConstraintItem) GetResolveAvailability() bool`

GetResolveAvailability returns the ResolveAvailability field if non-nil, zero value otherwise.

### GetResolveAvailabilityOk

`func (o *MicrosoftGraphLocationConstraintItem) GetResolveAvailabilityOk() (bool, bool)`

GetResolveAvailabilityOk returns a tuple with the ResolveAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResolveAvailability

`func (o *MicrosoftGraphLocationConstraintItem) HasResolveAvailability() bool`

HasResolveAvailability returns a boolean if a field has been set.

### SetResolveAvailability

`func (o *MicrosoftGraphLocationConstraintItem) SetResolveAvailability(v bool)`

SetResolveAvailability gets a reference to the given bool and assigns it to the ResolveAvailability field.

### SetResolveAvailabilityExplicitNull

`func (o *MicrosoftGraphLocationConstraintItem) SetResolveAvailabilityExplicitNull(b bool)`

SetResolveAvailabilityExplicitNull (un)sets ResolveAvailability to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResolveAvailability value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


