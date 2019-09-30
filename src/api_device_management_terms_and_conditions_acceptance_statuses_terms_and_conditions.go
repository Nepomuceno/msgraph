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

// DeviceManagementTermsAndConditionsAcceptanceStatusesTermsAndConditionsApiService DeviceManagementTermsAndConditionsAcceptanceStatusesTermsAndConditionsApi service
type DeviceManagementTermsAndConditionsAcceptanceStatusesTermsAndConditionsApiService service

// DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditionsOpts Optional parameters for the method 'DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions'
type DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditionsOpts struct {
    Select_ optional.Interface
    Expand optional.Interface
}

/*
DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions Get termsAndConditions from deviceManagement
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param termsAndConditionsId key: termsAndConditions-id of termsAndConditions
 * @param termsAndConditionsAcceptanceStatusId key: termsAndConditionsAcceptanceStatus-id of termsAndConditionsAcceptanceStatus
 * @param optional nil or *DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditionsOpts - Optional Parameters:
 * @param "Select_" (optional.Interface of []string) -  Select properties to be returned
 * @param "Expand" (optional.Interface of []string) -  Expand related entities
@return MicrosoftGraphTermsAndConditions
*/
func (a *DeviceManagementTermsAndConditionsAcceptanceStatusesTermsAndConditionsApiService) DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditions(ctx _context.Context, termsAndConditionsId string, termsAndConditionsAcceptanceStatusId string, localVarOptionals *DeviceManagementTermsAndConditionsAcceptanceStatusesGetTermsAndConditionsOpts) (MicrosoftGraphTermsAndConditions, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphTermsAndConditions
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/deviceManagement/termsAndConditions({termsAndConditions-id})/acceptanceStatuses({termsAndConditionsAcceptanceStatus-id})/termsAndConditions"
	localVarPath = strings.Replace(localVarPath, "{"+"termsAndConditions-id"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", termsAndConditionsId)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"termsAndConditionsAcceptanceStatus-id"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", termsAndConditionsAcceptanceStatusId)), -1)

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
			var v MicrosoftGraphTermsAndConditions
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
