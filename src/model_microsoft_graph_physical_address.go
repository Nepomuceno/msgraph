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
// MicrosoftGraphPhysicalAddress struct for MicrosoftGraphPhysicalAddress
type MicrosoftGraphPhysicalAddress struct {
	Street *string `json:"street,omitempty"`
	isExplicitNullStreet bool `json:"-"`
	City *string `json:"city,omitempty"`
	isExplicitNullCity bool `json:"-"`
	State *string `json:"state,omitempty"`
	isExplicitNullState bool `json:"-"`
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	isExplicitNullCountryOrRegion bool `json:"-"`
	PostalCode *string `json:"postalCode,omitempty"`
	isExplicitNullPostalCode bool `json:"-"`
}

// GetStreet returns the Street field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPhysicalAddress) GetStreet() string {
	if o == nil || o.Street == nil {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPhysicalAddress) GetStreetOk() (string, bool) {
	if o == nil || o.Street == nil {
		var ret string
		return ret, false
	}
	return *o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *MicrosoftGraphPhysicalAddress) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *MicrosoftGraphPhysicalAddress) SetStreet(v string) {
	o.Street = &v
}

// SetStreetExplicitNull (un)sets Street to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Street value is set to nil even if false is passed
func (o *MicrosoftGraphPhysicalAddress) SetStreetExplicitNull(b bool) {
	o.Street = nil
	o.isExplicitNullStreet = b
}
// GetCity returns the City field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPhysicalAddress) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPhysicalAddress) GetCityOk() (string, bool) {
	if o == nil || o.City == nil {
		var ret string
		return ret, false
	}
	return *o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *MicrosoftGraphPhysicalAddress) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *MicrosoftGraphPhysicalAddress) SetCity(v string) {
	o.City = &v
}

// SetCityExplicitNull (un)sets City to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The City value is set to nil even if false is passed
func (o *MicrosoftGraphPhysicalAddress) SetCityExplicitNull(b bool) {
	o.City = nil
	o.isExplicitNullCity = b
}
// GetState returns the State field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPhysicalAddress) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPhysicalAddress) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphPhysicalAddress) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MicrosoftGraphPhysicalAddress) SetState(v string) {
	o.State = &v
}

// SetStateExplicitNull (un)sets State to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The State value is set to nil even if false is passed
func (o *MicrosoftGraphPhysicalAddress) SetStateExplicitNull(b bool) {
	o.State = nil
	o.isExplicitNullState = b
}
// GetCountryOrRegion returns the CountryOrRegion field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPhysicalAddress) GetCountryOrRegion() string {
	if o == nil || o.CountryOrRegion == nil {
		var ret string
		return ret
	}
	return *o.CountryOrRegion
}

// GetCountryOrRegionOk returns a tuple with the CountryOrRegion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPhysicalAddress) GetCountryOrRegionOk() (string, bool) {
	if o == nil || o.CountryOrRegion == nil {
		var ret string
		return ret, false
	}
	return *o.CountryOrRegion, true
}

// HasCountryOrRegion returns a boolean if a field has been set.
func (o *MicrosoftGraphPhysicalAddress) HasCountryOrRegion() bool {
	if o != nil && o.CountryOrRegion != nil {
		return true
	}

	return false
}

// SetCountryOrRegion gets a reference to the given string and assigns it to the CountryOrRegion field.
func (o *MicrosoftGraphPhysicalAddress) SetCountryOrRegion(v string) {
	o.CountryOrRegion = &v
}

// SetCountryOrRegionExplicitNull (un)sets CountryOrRegion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CountryOrRegion value is set to nil even if false is passed
func (o *MicrosoftGraphPhysicalAddress) SetCountryOrRegionExplicitNull(b bool) {
	o.CountryOrRegion = nil
	o.isExplicitNullCountryOrRegion = b
}
// GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.
func (o *MicrosoftGraphPhysicalAddress) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPhysicalAddress) GetPostalCodeOk() (string, bool) {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret, false
	}
	return *o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *MicrosoftGraphPhysicalAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *MicrosoftGraphPhysicalAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// SetPostalCodeExplicitNull (un)sets PostalCode to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PostalCode value is set to nil even if false is passed
func (o *MicrosoftGraphPhysicalAddress) SetPostalCodeExplicitNull(b bool) {
	o.PostalCode = nil
	o.isExplicitNullPostalCode = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphPhysicalAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Street == nil {
		if o.isExplicitNullStreet {
			toSerialize["street"] = o.Street
		}
	} else {
		toSerialize["street"] = o.Street
	}
	if o.City == nil {
		if o.isExplicitNullCity {
			toSerialize["city"] = o.City
		}
	} else {
		toSerialize["city"] = o.City
	}
	if o.State == nil {
		if o.isExplicitNullState {
			toSerialize["state"] = o.State
		}
	} else {
		toSerialize["state"] = o.State
	}
	if o.CountryOrRegion == nil {
		if o.isExplicitNullCountryOrRegion {
			toSerialize["countryOrRegion"] = o.CountryOrRegion
		}
	} else {
		toSerialize["countryOrRegion"] = o.CountryOrRegion
	}
	if o.PostalCode == nil {
		if o.isExplicitNullPostalCode {
			toSerialize["postalCode"] = o.PostalCode
		}
	} else {
		toSerialize["postalCode"] = o.PostalCode
	}
	return json.Marshal(toSerialize)
}


