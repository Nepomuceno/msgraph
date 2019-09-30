# VppToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationName** | Pointer to **string** | The organization associated with the Apple Volume Purchase Program Token | [optional] 
**VppTokenAccountType** | Pointer to [**AnyOfmicrosoftGraphVppTokenAccountType**](anyOf&lt;microsoft.graph.vppTokenAccountType&gt;.md) | The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: &#x60;business&#x60;, &#x60;education&#x60;. | [optional] 
**AppleId** | Pointer to **string** | The apple Id associated with the given Apple Volume Purchase Program Token. | [optional] 
**ExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) | The expiration date time of the Apple Volume Purchase Program Token. | [optional] 
**LastSyncDateTime** | Pointer to [**time.Time**](time.Time.md) | The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token. | [optional] 
**Token** | Pointer to **string** | The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program. | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | Last modification date time associated with the Apple Volume Purchase Program Token. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphVppTokenState**](anyOf&lt;microsoft.graph.vppTokenState&gt;.md) | Current state of the Apple Volume Purchase Program Token. Possible values are: &#x60;unknown&#x60;, &#x60;valid&#x60;, &#x60;expired&#x60;, &#x60;invalid&#x60;, &#x60;assignedToExternalMDM&#x60;. | [optional] 
**LastSyncStatus** | Pointer to [**AnyOfmicrosoftGraphVppTokenSyncStatus**](anyOf&lt;microsoft.graph.vppTokenSyncStatus&gt;.md) | Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: &#x60;none&#x60;, &#x60;inProgress&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;. | [optional] 
**AutomaticallyUpdateApps** | Pointer to **bool** | Whether or not apps for the VPP token will be automatically updated. | [optional] 
**CountryOrRegion** | Pointer to **string** | Whether or not apps for the VPP token will be automatically updated. | [optional] 

## Methods

### GetOrganizationName

`func (o *VppToken) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *VppToken) GetOrganizationNameOk() (string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizationName

`func (o *VppToken) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationName

`func (o *VppToken) SetOrganizationName(v string)`

SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.

### SetOrganizationNameExplicitNull

`func (o *VppToken) SetOrganizationNameExplicitNull(b bool)`

SetOrganizationNameExplicitNull (un)sets OrganizationName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OrganizationName value is set to nil even if false is passed
### GetVppTokenAccountType

`func (o *VppToken) GetVppTokenAccountType() AnyOfmicrosoftGraphVppTokenAccountType`

GetVppTokenAccountType returns the VppTokenAccountType field if non-nil, zero value otherwise.

### GetVppTokenAccountTypeOk

`func (o *VppToken) GetVppTokenAccountTypeOk() (AnyOfmicrosoftGraphVppTokenAccountType, bool)`

GetVppTokenAccountTypeOk returns a tuple with the VppTokenAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVppTokenAccountType

`func (o *VppToken) HasVppTokenAccountType() bool`

HasVppTokenAccountType returns a boolean if a field has been set.

### SetVppTokenAccountType

`func (o *VppToken) SetVppTokenAccountType(v AnyOfmicrosoftGraphVppTokenAccountType)`

SetVppTokenAccountType gets a reference to the given AnyOfmicrosoftGraphVppTokenAccountType and assigns it to the VppTokenAccountType field.

### GetAppleId

`func (o *VppToken) GetAppleId() string`

GetAppleId returns the AppleId field if non-nil, zero value otherwise.

### GetAppleIdOk

`func (o *VppToken) GetAppleIdOk() (string, bool)`

GetAppleIdOk returns a tuple with the AppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppleId

`func (o *VppToken) HasAppleId() bool`

HasAppleId returns a boolean if a field has been set.

### SetAppleId

`func (o *VppToken) SetAppleId(v string)`

SetAppleId gets a reference to the given string and assigns it to the AppleId field.

### SetAppleIdExplicitNull

`func (o *VppToken) SetAppleIdExplicitNull(b bool)`

SetAppleIdExplicitNull (un)sets AppleId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppleId value is set to nil even if false is passed
### GetExpirationDateTime

`func (o *VppToken) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *VppToken) GetExpirationDateTimeOk() (time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpirationDateTime

`func (o *VppToken) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTime

`func (o *VppToken) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.

### GetLastSyncDateTime

`func (o *VppToken) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *VppToken) GetLastSyncDateTimeOk() (time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSyncDateTime

`func (o *VppToken) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### SetLastSyncDateTime

`func (o *VppToken) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.

### GetToken

`func (o *VppToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VppToken) GetTokenOk() (string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToken

`func (o *VppToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetToken

`func (o *VppToken) SetToken(v string)`

SetToken gets a reference to the given string and assigns it to the Token field.

### SetTokenExplicitNull

`func (o *VppToken) SetTokenExplicitNull(b bool)`

SetTokenExplicitNull (un)sets Token to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Token value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *VppToken) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *VppToken) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *VppToken) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *VppToken) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetState

`func (o *VppToken) GetState() AnyOfmicrosoftGraphVppTokenState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VppToken) GetStateOk() (AnyOfmicrosoftGraphVppTokenState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *VppToken) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *VppToken) SetState(v AnyOfmicrosoftGraphVppTokenState)`

SetState gets a reference to the given AnyOfmicrosoftGraphVppTokenState and assigns it to the State field.

### GetLastSyncStatus

`func (o *VppToken) GetLastSyncStatus() AnyOfmicrosoftGraphVppTokenSyncStatus`

GetLastSyncStatus returns the LastSyncStatus field if non-nil, zero value otherwise.

### GetLastSyncStatusOk

`func (o *VppToken) GetLastSyncStatusOk() (AnyOfmicrosoftGraphVppTokenSyncStatus, bool)`

GetLastSyncStatusOk returns a tuple with the LastSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSyncStatus

`func (o *VppToken) HasLastSyncStatus() bool`

HasLastSyncStatus returns a boolean if a field has been set.

### SetLastSyncStatus

`func (o *VppToken) SetLastSyncStatus(v AnyOfmicrosoftGraphVppTokenSyncStatus)`

SetLastSyncStatus gets a reference to the given AnyOfmicrosoftGraphVppTokenSyncStatus and assigns it to the LastSyncStatus field.

### GetAutomaticallyUpdateApps

`func (o *VppToken) GetAutomaticallyUpdateApps() bool`

GetAutomaticallyUpdateApps returns the AutomaticallyUpdateApps field if non-nil, zero value otherwise.

### GetAutomaticallyUpdateAppsOk

`func (o *VppToken) GetAutomaticallyUpdateAppsOk() (bool, bool)`

GetAutomaticallyUpdateAppsOk returns a tuple with the AutomaticallyUpdateApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAutomaticallyUpdateApps

`func (o *VppToken) HasAutomaticallyUpdateApps() bool`

HasAutomaticallyUpdateApps returns a boolean if a field has been set.

### SetAutomaticallyUpdateApps

`func (o *VppToken) SetAutomaticallyUpdateApps(v bool)`

SetAutomaticallyUpdateApps gets a reference to the given bool and assigns it to the AutomaticallyUpdateApps field.

### GetCountryOrRegion

`func (o *VppToken) GetCountryOrRegion() string`

GetCountryOrRegion returns the CountryOrRegion field if non-nil, zero value otherwise.

### GetCountryOrRegionOk

`func (o *VppToken) GetCountryOrRegionOk() (string, bool)`

GetCountryOrRegionOk returns a tuple with the CountryOrRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCountryOrRegion

`func (o *VppToken) HasCountryOrRegion() bool`

HasCountryOrRegion returns a boolean if a field has been set.

### SetCountryOrRegion

`func (o *VppToken) SetCountryOrRegion(v string)`

SetCountryOrRegion gets a reference to the given string and assigns it to the CountryOrRegion field.

### SetCountryOrRegionExplicitNull

`func (o *VppToken) SetCountryOrRegionExplicitNull(b bool)`

SetCountryOrRegionExplicitNull (un)sets CountryOrRegion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CountryOrRegion value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


