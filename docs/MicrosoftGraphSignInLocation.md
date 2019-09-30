# MicrosoftGraphSignInLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**CountryOrRegion** | Pointer to **string** |  | [optional] 
**GeoCoordinates** | Pointer to [**AnyOfmicrosoftGraphGeoCoordinates**](anyOf&lt;microsoft.graph.geoCoordinates&gt;.md) |  | [optional] 

## Methods

### GetCity

`func (o *MicrosoftGraphSignInLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MicrosoftGraphSignInLocation) GetCityOk() (string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCity

`func (o *MicrosoftGraphSignInLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCity

`func (o *MicrosoftGraphSignInLocation) SetCity(v string)`

SetCity gets a reference to the given string and assigns it to the City field.

### SetCityExplicitNull

`func (o *MicrosoftGraphSignInLocation) SetCityExplicitNull(b bool)`

SetCityExplicitNull (un)sets City to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The City value is set to nil even if false is passed
### GetState

`func (o *MicrosoftGraphSignInLocation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphSignInLocation) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *MicrosoftGraphSignInLocation) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *MicrosoftGraphSignInLocation) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### SetStateExplicitNull

`func (o *MicrosoftGraphSignInLocation) SetStateExplicitNull(b bool)`

SetStateExplicitNull (un)sets State to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The State value is set to nil even if false is passed
### GetCountryOrRegion

`func (o *MicrosoftGraphSignInLocation) GetCountryOrRegion() string`

GetCountryOrRegion returns the CountryOrRegion field if non-nil, zero value otherwise.

### GetCountryOrRegionOk

`func (o *MicrosoftGraphSignInLocation) GetCountryOrRegionOk() (string, bool)`

GetCountryOrRegionOk returns a tuple with the CountryOrRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCountryOrRegion

`func (o *MicrosoftGraphSignInLocation) HasCountryOrRegion() bool`

HasCountryOrRegion returns a boolean if a field has been set.

### SetCountryOrRegion

`func (o *MicrosoftGraphSignInLocation) SetCountryOrRegion(v string)`

SetCountryOrRegion gets a reference to the given string and assigns it to the CountryOrRegion field.

### SetCountryOrRegionExplicitNull

`func (o *MicrosoftGraphSignInLocation) SetCountryOrRegionExplicitNull(b bool)`

SetCountryOrRegionExplicitNull (un)sets CountryOrRegion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CountryOrRegion value is set to nil even if false is passed
### GetGeoCoordinates

`func (o *MicrosoftGraphSignInLocation) GetGeoCoordinates() AnyOfmicrosoftGraphGeoCoordinates`

GetGeoCoordinates returns the GeoCoordinates field if non-nil, zero value otherwise.

### GetGeoCoordinatesOk

`func (o *MicrosoftGraphSignInLocation) GetGeoCoordinatesOk() (AnyOfmicrosoftGraphGeoCoordinates, bool)`

GetGeoCoordinatesOk returns a tuple with the GeoCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGeoCoordinates

`func (o *MicrosoftGraphSignInLocation) HasGeoCoordinates() bool`

HasGeoCoordinates returns a boolean if a field has been set.

### SetGeoCoordinates

`func (o *MicrosoftGraphSignInLocation) SetGeoCoordinates(v AnyOfmicrosoftGraphGeoCoordinates)`

SetGeoCoordinates gets a reference to the given AnyOfmicrosoftGraphGeoCoordinates and assigns it to the GeoCoordinates field.

### SetGeoCoordinatesExplicitNull

`func (o *MicrosoftGraphSignInLocation) SetGeoCoordinatesExplicitNull(b bool)`

SetGeoCoordinatesExplicitNull (un)sets GeoCoordinates to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The GeoCoordinates value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


