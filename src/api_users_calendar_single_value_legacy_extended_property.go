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
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"fmt"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// UsersCalendarSingleValueLegacyExtendedPropertyApiService UsersCalendarSingleValueLegacyExtendedPropertyApi service
type UsersCalendarSingleValueLegacyExtendedPropertyApiService service

/*
UsersCalendarCreateSingleValueExtendedProperties Create new navigation property to singleValueExtendedProperties for users
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId key: user-id of user
 * @param microsoftGraphSingleValueLegacyExtendedProperty New navigation property
@return MicrosoftGraphSingleValueLegacyExtendedProperty
*/
func (a *UsersCalendarSingleValueLegacyExtendedPropertyApiService) UsersCalendarCreateSingleValueExtendedProperties(ctx _context.Context, userId string, microsoftGraphSingleValueLegacyExtendedProperty MicrosoftGraphSingleValueLegacyExtendedProperty) (MicrosoftGraphSingleValueLegacyExtendedProperty, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphSingleValueLegacyExtendedProperty
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users({user-id})/calendar/singleValueExtendedProperties"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", userId)), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &microsoftGraphSingleValueLegacyExtendedProperty
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 201 {
			var v MicrosoftGraphSingleValueLegacyExtendedProperty
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 0 {
			var v OdataError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// UsersCalendarGetSingleValueExtendedPropertiesOpts Optional parameters for the method 'UsersCalendarGetSingleValueExtendedProperties'
type UsersCalendarGetSingleValueExtendedPropertiesOpts struct {
    Select_ optional.Interface
    Expand optional.Interface
}

/*
UsersCalendarGetSingleValueExtendedProperties Get singleValueExtendedProperties from users
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId key: user-id of user
 * @param singleValueLegacyExtendedPropertyId key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty
 * @param optional nil or *UsersCalendarGetSingleValueExtendedPropertiesOpts - Optional Parameters:
 * @param "Select_" (optional.Interface of []string) -  Select properties to be returned
 * @param "Expand" (optional.Interface of []string) -  Expand related entities
@return MicrosoftGraphSingleValueLegacyExtendedProperty
*/
func (a *UsersCalendarSingleValueLegacyExtendedPropertyApiService) UsersCalendarGetSingleValueExtendedProperties(ctx _context.Context, userId string, singleValueLegacyExtendedPropertyId string, localVarOptionals *UsersCalendarGetSingleValueExtendedPropertiesOpts) (MicrosoftGraphSingleValueLegacyExtendedProperty, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphSingleValueLegacyExtendedProperty
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users({user-id})/calendar/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id})"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", userId)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"singleValueLegacyExtendedProperty-id"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", singleValueLegacyExtendedPropertyId)), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("$select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("$expand", parameterToString(localVarOptionals.Expand.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v MicrosoftGraphSingleValueLegacyExtendedProperty
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 0 {
			var v OdataError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// UsersCalendarListSingleValueExtendedPropertiesOpts Optional parameters for the method 'UsersCalendarListSingleValueExtendedProperties'
type UsersCalendarListSingleValueExtendedPropertiesOpts struct {
    Top optional.Int32
    Skip optional.Int32
    Search optional.String
    Filter optional.String
    Count optional.Bool
    Orderby optional.Interface
    Select_ optional.Interface
    Expand optional.Interface
}

/*
UsersCalendarListSingleValueExtendedProperties Get singleValueExtendedProperties from users
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId key: user-id of user
 * @param optional nil or *UsersCalendarListSingleValueExtendedPropertiesOpts - Optional Parameters:
 * @param "Top" (optional.Int32) -  Show only the first n items
 * @param "Skip" (optional.Int32) -  Skip the first n items
 * @param "Search" (optional.String) -  Search items by search phrases
 * @param "Filter" (optional.String) -  Filter items by property values
 * @param "Count" (optional.Bool) -  Include count of items
 * @param "Orderby" (optional.Interface of []string) -  Order items by property values
 * @param "Select_" (optional.Interface of []string) -  Select properties to be returned
 * @param "Expand" (optional.Interface of []string) -  Expand related entities
@return CollectionOfSingleValueLegacyExtendedProperty
*/
func (a *UsersCalendarSingleValueLegacyExtendedPropertyApiService) UsersCalendarListSingleValueExtendedProperties(ctx _context.Context, userId string, localVarOptionals *UsersCalendarListSingleValueExtendedPropertiesOpts) (CollectionOfSingleValueLegacyExtendedProperty, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CollectionOfSingleValueLegacyExtendedProperty
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users({user-id})/calendar/singleValueExtendedProperties"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", userId)), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Top.IsSet() {
		localVarQueryParams.Add("$top", parameterToString(localVarOptionals.Top.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Skip.IsSet() {
		localVarQueryParams.Add("$skip", parameterToString(localVarOptionals.Skip.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Search.IsSet() {
		localVarQueryParams.Add("$search", parameterToString(localVarOptionals.Search.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("$filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("$count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Orderby.IsSet() {
		localVarQueryParams.Add("$orderby", parameterToString(localVarOptionals.Orderby.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("$select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("$expand", parameterToString(localVarOptionals.Expand.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v CollectionOfSingleValueLegacyExtendedProperty
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 0 {
			var v OdataError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
UsersCalendarUpdateSingleValueExtendedProperties Update the navigation property singleValueExtendedProperties in users
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId key: user-id of user
 * @param singleValueLegacyExtendedPropertyId key: singleValueLegacyExtendedProperty-id of singleValueLegacyExtendedProperty
 * @param microsoftGraphSingleValueLegacyExtendedProperty New navigation property values
*/
func (a *UsersCalendarSingleValueLegacyExtendedPropertyApiService) UsersCalendarUpdateSingleValueExtendedProperties(ctx _context.Context, userId string, singleValueLegacyExtendedPropertyId string, microsoftGraphSingleValueLegacyExtendedProperty MicrosoftGraphSingleValueLegacyExtendedProperty) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users({user-id})/calendar/singleValueExtendedProperties({singleValueLegacyExtendedProperty-id})"
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", userId)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"singleValueLegacyExtendedProperty-id"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", singleValueLegacyExtendedPropertyId)), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &microsoftGraphSingleValueLegacyExtendedProperty
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 0 {
			var v OdataError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
