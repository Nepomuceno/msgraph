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
// MicrosoftGraphGeoCoordinates struct for MicrosoftGraphGeoCoordinates
type MicrosoftGraphGeoCoordinates struct {
	Altitude *AnyOfnumberstringstring `json:"altitude,omitempty"`
	isExplicitNullAltitude bool `json:"-"`
	Latitude *AnyOfnumberstringstring `json:"latitude,omitempty"`
	isExplicitNullLatitude bool `json:"-"`
	Longitude *AnyOfnumberstringstring `json:"longitude,omitempty"`
	isExplicitNullLongitude bool `json:"-"`
}

// GetAltitude returns the Altitude field if non-nil, zero value otherwise.
func (o *MicrosoftGraphGeoCoordinates) GetAltitude() AnyOfnumberstringstring {
	if o == nil || o.Altitude == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphGeoCoordinates) GetAltitudeOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Altitude == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Altitude, true
}

// HasAltitude returns a boolean if a field has been set.
func (o *MicrosoftGraphGeoCoordinates) HasAltitude() bool {
	if o != nil && o.Altitude != nil {
		return true
	}

	return false
}

// SetAltitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Altitude field.
func (o *MicrosoftGraphGeoCoordinates) SetAltitude(v AnyOfnumberstringstring) {
	o.Altitude = &v
}

// SetAltitudeExplicitNull (un)sets Altitude to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Altitude value is set to nil even if false is passed
func (o *MicrosoftGraphGeoCoordinates) SetAltitudeExplicitNull(b bool) {
	o.Altitude = nil
	o.isExplicitNullAltitude = b
}
// GetLatitude returns the Latitude field if non-nil, zero value otherwise.
func (o *MicrosoftGraphGeoCoordinates) GetLatitude() AnyOfnumberstringstring {
	if o == nil || o.Latitude == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphGeoCoordinates) GetLatitudeOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Latitude == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *MicrosoftGraphGeoCoordinates) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Latitude field.
func (o *MicrosoftGraphGeoCoordinates) SetLatitude(v AnyOfnumberstringstring) {
	o.Latitude = &v
}

// SetLatitudeExplicitNull (un)sets Latitude to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Latitude value is set to nil even if false is passed
func (o *MicrosoftGraphGeoCoordinates) SetLatitudeExplicitNull(b bool) {
	o.Latitude = nil
	o.isExplicitNullLatitude = b
}
// GetLongitude returns the Longitude field if non-nil, zero value otherwise.
func (o *MicrosoftGraphGeoCoordinates) GetLongitude() AnyOfnumberstringstring {
	if o == nil || o.Longitude == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphGeoCoordinates) GetLongitudeOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Longitude == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *MicrosoftGraphGeoCoordinates) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Longitude field.
func (o *MicrosoftGraphGeoCoordinates) SetLongitude(v AnyOfnumberstringstring) {
	o.Longitude = &v
}

// SetLongitudeExplicitNull (un)sets Longitude to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Longitude value is set to nil even if false is passed
func (o *MicrosoftGraphGeoCoordinates) SetLongitudeExplicitNull(b bool) {
	o.Longitude = nil
	o.isExplicitNullLongitude = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphGeoCoordinates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Altitude == nil {
		if o.isExplicitNullAltitude {
			toSerialize["altitude"] = o.Altitude
		}
	} else {
		toSerialize["altitude"] = o.Altitude
	}
	if o.Latitude == nil {
		if o.isExplicitNullLatitude {
			toSerialize["latitude"] = o.Latitude
		}
	} else {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude == nil {
		if o.isExplicitNullLongitude {
			toSerialize["longitude"] = o.Longitude
		}
	} else {
		toSerialize["longitude"] = o.Longitude
	}
	return json.Marshal(toSerialize)
}

