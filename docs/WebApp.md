# WebApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | Pointer to **string** | The web app URL. | [optional] 
**UseManagedBrowser** | Pointer to **bool** | Whether or not to use managed browser. This property is only applicable for Android and IOS. | [optional] 

## Methods

### GetAppUrl

`func (o *WebApp) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *WebApp) GetAppUrlOk() (string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppUrl

`func (o *WebApp) HasAppUrl() bool`

HasAppUrl returns a boolean if a field has been set.

### SetAppUrl

`func (o *WebApp) SetAppUrl(v string)`

SetAppUrl gets a reference to the given string and assigns it to the AppUrl field.

### SetAppUrlExplicitNull

`func (o *WebApp) SetAppUrlExplicitNull(b bool)`

SetAppUrlExplicitNull (un)sets AppUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppUrl value is set to nil even if false is passed
### GetUseManagedBrowser

`func (o *WebApp) GetUseManagedBrowser() bool`

GetUseManagedBrowser returns the UseManagedBrowser field if non-nil, zero value otherwise.

### GetUseManagedBrowserOk

`func (o *WebApp) GetUseManagedBrowserOk() (bool, bool)`

GetUseManagedBrowserOk returns a tuple with the UseManagedBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUseManagedBrowser

`func (o *WebApp) HasUseManagedBrowser() bool`

HasUseManagedBrowser returns a boolean if a field has been set.

### SetUseManagedBrowser

`func (o *WebApp) SetUseManagedBrowser(v bool)`

SetUseManagedBrowser gets a reference to the given bool and assigns it to the UseManagedBrowser field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


