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

// WorkbooksWorkbookWorksheetsWorkbookWorksheetProtectionApiService WorkbooksWorkbookWorksheetsWorkbookWorksheetProtectionApi service
type WorkbooksWorkbookWorksheetsWorkbookWorksheetProtectionApiService service

// WorkbooksWorkbookWorksheetsGetProtectionOpts Optional parameters for the method 'WorkbooksWorkbookWorksheetsGetProtection'
type WorkbooksWorkbookWorksheetsGetProtectionOpts struct {
    Select_ optional.Interface
    Expand optional.Interface
}

/*
WorkbooksWorkbookWorksheetsGetProtection Get protection from workbooks
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param driveItemId key: driveItem-id of driveItem
 * @param workbookWorksheetId key: workbookWorksheet-id of workbookWorksheet
 * @param optional nil or *WorkbooksWorkbookWorksheetsGetProtectionOpts - Optional Parameters:
 * @param "Select_" (optional.Interface of []string) -  Select properties to be returned
 * @param "Expand" (optional.Interface of []string) -  Expand related entities
@return MicrosoftGraphWorkbookWorksheetProtection
*/
func (a *WorkbooksWorkbookWorksheetsWorkbookWorksheetProtectionApiService) WorkbooksWorkbookWorksheetsGetProtection(ctx _context.Context, driveItemId string, workbookWorksheetId string, localVarOptionals *WorkbooksWorkbookWorksheetsGetProtectionOpts) (MicrosoftGraphWorkbookWorksheetProtection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MicrosoftGraphWorkbookWorksheetProtection
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/protection"
	localVarPath = strings.Replace(localVarPath, "{"+"driveItem-id"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", driveItemId)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workbookWorksheet-id"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", workbookWorksheetId)), -1)

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
			var v MicrosoftGraphWorkbookWorksheetProtection
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
WorkbooksWorkbookWorksheetsUpdateProtection Update the navigation property protection in workbooks
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param driveItemId key: driveItem-id of driveItem
 * @param workbookWorksheetId key: workbookWorksheet-id of workbookWorksheet
 * @param microsoftGraphWorkbookWorksheetProtection New navigation property values
*/
func (a *WorkbooksWorkbookWorksheetsWorkbookWorksheetProtectionApiService) WorkbooksWorkbookWorksheetsUpdateProtection(ctx _context.Context, driveItemId string, workbookWorksheetId string, microsoftGraphWorkbookWorksheetProtection MicrosoftGraphWorkbookWorksheetProtection) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/workbooks({driveItem-id})/workbook/worksheets({workbookWorksheet-id})/protection"
	localVarPath = strings.Replace(localVarPath, "{"+"driveItem-id"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", driveItemId)), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workbookWorksheet-id"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", workbookWorksheetId)), -1)

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
	localVarPostBody = &microsoftGraphWorkbookWorksheetProtection
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
