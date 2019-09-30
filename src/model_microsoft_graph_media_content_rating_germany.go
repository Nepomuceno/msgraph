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
// MicrosoftGraphMediaContentRatingGermany struct for MicrosoftGraphMediaContentRatingGermany
type MicrosoftGraphMediaContentRatingGermany struct {
	// Movies rating selected for Germany
	MovieRating *AnyOfmicrosoftGraphRatingGermanyMoviesType `json:"movieRating,omitempty"`

	// TV rating selected for Germany
	TvRating *AnyOfmicrosoftGraphRatingGermanyTelevisionType `json:"tvRating,omitempty"`

}

// GetMovieRating returns the MovieRating field if non-nil, zero value otherwise.
func (o *MicrosoftGraphMediaContentRatingGermany) GetMovieRating() AnyOfmicrosoftGraphRatingGermanyMoviesType {
	if o == nil || o.MovieRating == nil {
		var ret AnyOfmicrosoftGraphRatingGermanyMoviesType
		return ret
	}
	return *o.MovieRating
}

// GetMovieRatingOk returns a tuple with the MovieRating field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMediaContentRatingGermany) GetMovieRatingOk() (AnyOfmicrosoftGraphRatingGermanyMoviesType, bool) {
	if o == nil || o.MovieRating == nil {
		var ret AnyOfmicrosoftGraphRatingGermanyMoviesType
		return ret, false
	}
	return *o.MovieRating, true
}

// HasMovieRating returns a boolean if a field has been set.
func (o *MicrosoftGraphMediaContentRatingGermany) HasMovieRating() bool {
	if o != nil && o.MovieRating != nil {
		return true
	}

	return false
}

// SetMovieRating gets a reference to the given AnyOfmicrosoftGraphRatingGermanyMoviesType and assigns it to the MovieRating field.
func (o *MicrosoftGraphMediaContentRatingGermany) SetMovieRating(v AnyOfmicrosoftGraphRatingGermanyMoviesType) {
	o.MovieRating = &v
}

// GetTvRating returns the TvRating field if non-nil, zero value otherwise.
func (o *MicrosoftGraphMediaContentRatingGermany) GetTvRating() AnyOfmicrosoftGraphRatingGermanyTelevisionType {
	if o == nil || o.TvRating == nil {
		var ret AnyOfmicrosoftGraphRatingGermanyTelevisionType
		return ret
	}
	return *o.TvRating
}

// GetTvRatingOk returns a tuple with the TvRating field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphMediaContentRatingGermany) GetTvRatingOk() (AnyOfmicrosoftGraphRatingGermanyTelevisionType, bool) {
	if o == nil || o.TvRating == nil {
		var ret AnyOfmicrosoftGraphRatingGermanyTelevisionType
		return ret, false
	}
	return *o.TvRating, true
}

// HasTvRating returns a boolean if a field has been set.
func (o *MicrosoftGraphMediaContentRatingGermany) HasTvRating() bool {
	if o != nil && o.TvRating != nil {
		return true
	}

	return false
}

// SetTvRating gets a reference to the given AnyOfmicrosoftGraphRatingGermanyTelevisionType and assigns it to the TvRating field.
func (o *MicrosoftGraphMediaContentRatingGermany) SetTvRating(v AnyOfmicrosoftGraphRatingGermanyTelevisionType) {
	o.TvRating = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphMediaContentRatingGermany) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MovieRating != nil {
		toSerialize["movieRating"] = o.MovieRating
	}
	if o.TvRating != nil {
		toSerialize["tvRating"] = o.TvRating
	}
	return json.Marshal(toSerialize)
}


