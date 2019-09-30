# MicrosoftGraphDeviceGeoLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastCollectedDateTime** | Pointer to [**time.Time**](time.Time.md) | Time at which location was recorded, relative to UTC | [optional] 
**Longitude** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Longitude coordinate of the device&#39;s location | [optional] 
**Latitude** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Latitude coordinate of the device&#39;s location | [optional] 
**Altitude** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Altitude, given in meters above sea level | [optional] 
**HorizontalAccuracy** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Accuracy of longitude and latitude in meters | [optional] 
**VerticalAccuracy** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Accuracy of altitude in meters | [optional] 
**Heading** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Heading in degrees from true north | [optional] 
**Speed** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Speed the device is traveling in meters per second | [optional] 

## Methods

### GetLastCollectedDateTime

`func (o *MicrosoftGraphDeviceGeoLocation) GetLastCollectedDateTime() time.Time`

GetLastCollectedDateTime returns the LastCollectedDateTime field if non-nil, zero value otherwise.

### GetLastCollectedDateTimeOk

`func (o *MicrosoftGraphDeviceGeoLocation) GetLastCollectedDateTimeOk() (time.Time, bool)`

GetLastCollectedDateTimeOk returns a tuple with the LastCollectedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastCollectedDateTime

`func (o *MicrosoftGraphDeviceGeoLocation) HasLastCollectedDateTime() bool`

HasLastCollectedDateTime returns a boolean if a field has been set.

### SetLastCollectedDateTime

`func (o *MicrosoftGraphDeviceGeoLocation) SetLastCollectedDateTime(v time.Time)`

SetLastCollectedDateTime gets a reference to the given time.Time and assigns it to the LastCollectedDateTime field.

### GetLongitude

`func (o *MicrosoftGraphDeviceGeoLocation) GetLongitude() AnyOfnumberstringstring`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *MicrosoftGraphDeviceGeoLocation) GetLongitudeOk() (AnyOfnumberstringstring, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLongitude

`func (o *MicrosoftGraphDeviceGeoLocation) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitude

`func (o *MicrosoftGraphDeviceGeoLocation) SetLongitude(v AnyOfnumberstringstring)`

SetLongitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Longitude field.

### GetLatitude

`func (o *MicrosoftGraphDeviceGeoLocation) GetLatitude() AnyOfnumberstringstring`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *MicrosoftGraphDeviceGeoLocation) GetLatitudeOk() (AnyOfnumberstringstring, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLatitude

`func (o *MicrosoftGraphDeviceGeoLocation) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitude

`func (o *MicrosoftGraphDeviceGeoLocation) SetLatitude(v AnyOfnumberstringstring)`

SetLatitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Latitude field.

### GetAltitude

`func (o *MicrosoftGraphDeviceGeoLocation) GetAltitude() AnyOfnumberstringstring`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *MicrosoftGraphDeviceGeoLocation) GetAltitudeOk() (AnyOfnumberstringstring, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAltitude

`func (o *MicrosoftGraphDeviceGeoLocation) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### SetAltitude

`func (o *MicrosoftGraphDeviceGeoLocation) SetAltitude(v AnyOfnumberstringstring)`

SetAltitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Altitude field.

### GetHorizontalAccuracy

`func (o *MicrosoftGraphDeviceGeoLocation) GetHorizontalAccuracy() AnyOfnumberstringstring`

GetHorizontalAccuracy returns the HorizontalAccuracy field if non-nil, zero value otherwise.

### GetHorizontalAccuracyOk

`func (o *MicrosoftGraphDeviceGeoLocation) GetHorizontalAccuracyOk() (AnyOfnumberstringstring, bool)`

GetHorizontalAccuracyOk returns a tuple with the HorizontalAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHorizontalAccuracy

`func (o *MicrosoftGraphDeviceGeoLocation) HasHorizontalAccuracy() bool`

HasHorizontalAccuracy returns a boolean if a field has been set.

### SetHorizontalAccuracy

`func (o *MicrosoftGraphDeviceGeoLocation) SetHorizontalAccuracy(v AnyOfnumberstringstring)`

SetHorizontalAccuracy gets a reference to the given AnyOfnumberstringstring and assigns it to the HorizontalAccuracy field.

### GetVerticalAccuracy

`func (o *MicrosoftGraphDeviceGeoLocation) GetVerticalAccuracy() AnyOfnumberstringstring`

GetVerticalAccuracy returns the VerticalAccuracy field if non-nil, zero value otherwise.

### GetVerticalAccuracyOk

`func (o *MicrosoftGraphDeviceGeoLocation) GetVerticalAccuracyOk() (AnyOfnumberstringstring, bool)`

GetVerticalAccuracyOk returns a tuple with the VerticalAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVerticalAccuracy

`func (o *MicrosoftGraphDeviceGeoLocation) HasVerticalAccuracy() bool`

HasVerticalAccuracy returns a boolean if a field has been set.

### SetVerticalAccuracy

`func (o *MicrosoftGraphDeviceGeoLocation) SetVerticalAccuracy(v AnyOfnumberstringstring)`

SetVerticalAccuracy gets a reference to the given AnyOfnumberstringstring and assigns it to the VerticalAccuracy field.

### GetHeading

`func (o *MicrosoftGraphDeviceGeoLocation) GetHeading() AnyOfnumberstringstring`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *MicrosoftGraphDeviceGeoLocation) GetHeadingOk() (AnyOfnumberstringstring, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeading

`func (o *MicrosoftGraphDeviceGeoLocation) HasHeading() bool`

HasHeading returns a boolean if a field has been set.

### SetHeading

`func (o *MicrosoftGraphDeviceGeoLocation) SetHeading(v AnyOfnumberstringstring)`

SetHeading gets a reference to the given AnyOfnumberstringstring and assigns it to the Heading field.

### GetSpeed

`func (o *MicrosoftGraphDeviceGeoLocation) GetSpeed() AnyOfnumberstringstring`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *MicrosoftGraphDeviceGeoLocation) GetSpeedOk() (AnyOfnumberstringstring, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSpeed

`func (o *MicrosoftGraphDeviceGeoLocation) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeed

`func (o *MicrosoftGraphDeviceGeoLocation) SetSpeed(v AnyOfnumberstringstring)`

SetSpeed gets a reference to the given AnyOfnumberstringstring and assigns it to the Speed field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


