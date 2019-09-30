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
// MicrosoftGraphOutlookGeoCoordinates struct for MicrosoftGraphOutlookGeoCoordinates
type MicrosoftGraphOutlookGeoCoordinates struct {
	Latitude *AnyOfnumberstringstring `json:"latitude,omitempty"`
	isExplicitNullLatitude bool `json:"-"`
	Longitude *AnyOfnumberstringstring `json:"longitude,omitempty"`
	isExplicitNullLongitude bool `json:"-"`
	Accuracy *AnyOfnumberstringstring `json:"accuracy,omitempty"`
	isExplicitNullAccuracy bool `json:"-"`
	Altitude *AnyOfnumberstringstring `json:"altitude,omitempty"`
	isExplicitNullAltitude bool `json:"-"`
	AltitudeAccuracy *AnyOfnumberstringstring `json:"altitudeAccuracy,omitempty"`
	isExplicitNullAltitudeAccuracy bool `json:"-"`
}

// GetLatitude returns the Latitude field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookGeoCoordinates) GetLatitude() AnyOfnumberstringstring {
	if o == nil || o.Latitude == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookGeoCoordinates) GetLatitudeOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Latitude == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookGeoCoordinates) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Latitude field.
func (o *MicrosoftGraphOutlookGeoCoordinates) SetLatitude(v AnyOfnumberstringstring) {
	o.Latitude = &v
}

// SetLatitudeExplicitNull (un)sets Latitude to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Latitude value is set to nil even if false is passed
func (o *MicrosoftGraphOutlookGeoCoordinates) SetLatitudeExplicitNull(b bool) {
	o.Latitude = nil
	o.isExplicitNullLatitude = b
}
// GetLongitude returns the Longitude field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookGeoCoordinates) GetLongitude() AnyOfnumberstringstring {
	if o == nil || o.Longitude == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookGeoCoordinates) GetLongitudeOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Longitude == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookGeoCoordinates) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Longitude field.
func (o *MicrosoftGraphOutlookGeoCoordinates) SetLongitude(v AnyOfnumberstringstring) {
	o.Longitude = &v
}

// SetLongitudeExplicitNull (un)sets Longitude to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Longitude value is set to nil even if false is passed
func (o *MicrosoftGraphOutlookGeoCoordinates) SetLongitudeExplicitNull(b bool) {
	o.Longitude = nil
	o.isExplicitNullLongitude = b
}
// GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookGeoCoordinates) GetAccuracy() AnyOfnumberstringstring {
	if o == nil || o.Accuracy == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookGeoCoordinates) GetAccuracyOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Accuracy == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookGeoCoordinates) HasAccuracy() bool {
	if o != nil && o.Accuracy != nil {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given AnyOfnumberstringstring and assigns it to the Accuracy field.
func (o *MicrosoftGraphOutlookGeoCoordinates) SetAccuracy(v AnyOfnumberstringstring) {
	o.Accuracy = &v
}

// SetAccuracyExplicitNull (un)sets Accuracy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Accuracy value is set to nil even if false is passed
func (o *MicrosoftGraphOutlookGeoCoordinates) SetAccuracyExplicitNull(b bool) {
	o.Accuracy = nil
	o.isExplicitNullAccuracy = b
}
// GetAltitude returns the Altitude field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookGeoCoordinates) GetAltitude() AnyOfnumberstringstring {
	if o == nil || o.Altitude == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookGeoCoordinates) GetAltitudeOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Altitude == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Altitude, true
}

// HasAltitude returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookGeoCoordinates) HasAltitude() bool {
	if o != nil && o.Altitude != nil {
		return true
	}

	return false
}

// SetAltitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Altitude field.
func (o *MicrosoftGraphOutlookGeoCoordinates) SetAltitude(v AnyOfnumberstringstring) {
	o.Altitude = &v
}

// SetAltitudeExplicitNull (un)sets Altitude to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Altitude value is set to nil even if false is passed
func (o *MicrosoftGraphOutlookGeoCoordinates) SetAltitudeExplicitNull(b bool) {
	o.Altitude = nil
	o.isExplicitNullAltitude = b
}
// GetAltitudeAccuracy returns the AltitudeAccuracy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOutlookGeoCoordinates) GetAltitudeAccuracy() AnyOfnumberstringstring {
	if o == nil || o.AltitudeAccuracy == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.AltitudeAccuracy
}

// GetAltitudeAccuracyOk returns a tuple with the AltitudeAccuracy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOutlookGeoCoordinates) GetAltitudeAccuracyOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.AltitudeAccuracy == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.AltitudeAccuracy, true
}

// HasAltitudeAccuracy returns a boolean if a field has been set.
func (o *MicrosoftGraphOutlookGeoCoordinates) HasAltitudeAccuracy() bool {
	if o != nil && o.AltitudeAccuracy != nil {
		return true
	}

	return false
}

// SetAltitudeAccuracy gets a reference to the given AnyOfnumberstringstring and assigns it to the AltitudeAccuracy field.
func (o *MicrosoftGraphOutlookGeoCoordinates) SetAltitudeAccuracy(v AnyOfnumberstringstring) {
	o.AltitudeAccuracy = &v
}

// SetAltitudeAccuracyExplicitNull (un)sets AltitudeAccuracy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AltitudeAccuracy value is set to nil even if false is passed
func (o *MicrosoftGraphOutlookGeoCoordinates) SetAltitudeAccuracyExplicitNull(b bool) {
	o.AltitudeAccuracy = nil
	o.isExplicitNullAltitudeAccuracy = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphOutlookGeoCoordinates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.Accuracy == nil {
		if o.isExplicitNullAccuracy {
			toSerialize["accuracy"] = o.Accuracy
		}
	} else {
		toSerialize["accuracy"] = o.Accuracy
	}
	if o.Altitude == nil {
		if o.isExplicitNullAltitude {
			toSerialize["altitude"] = o.Altitude
		}
	} else {
		toSerialize["altitude"] = o.Altitude
	}
	if o.AltitudeAccuracy == nil {
		if o.isExplicitNullAltitudeAccuracy {
			toSerialize["altitudeAccuracy"] = o.AltitudeAccuracy
		}
	} else {
		toSerialize["altitudeAccuracy"] = o.AltitudeAccuracy
	}
	return json.Marshal(toSerialize)
}

