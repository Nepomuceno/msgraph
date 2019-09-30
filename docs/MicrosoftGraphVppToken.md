# MicrosoftGraphVppToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphVppToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphVppToken) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphVppToken) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphVppToken) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetOrganizationName

`func (o *MicrosoftGraphVppToken) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *MicrosoftGraphVppToken) GetOrganizationNameOk() (string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizationName

`func (o *MicrosoftGraphVppToken) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationName

`func (o *MicrosoftGraphVppToken) SetOrganizationName(v string)`

SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.

### SetOrganizationNameExplicitNull

`func (o *MicrosoftGraphVppToken) SetOrganizationNameExplicitNull(b bool)`

SetOrganizationNameExplicitNull (un)sets OrganizationName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OrganizationName value is set to nil even if false is passed
### GetVppTokenAccountType

`func (o *MicrosoftGraphVppToken) GetVppTokenAccountType() AnyOfmicrosoftGraphVppTokenAccountType`

GetVppTokenAccountType returns the VppTokenAccountType field if non-nil, zero value otherwise.

### GetVppTokenAccountTypeOk

`func (o *MicrosoftGraphVppToken) GetVppTokenAccountTypeOk() (AnyOfmicrosoftGraphVppTokenAccountType, bool)`

GetVppTokenAccountTypeOk returns a tuple with the VppTokenAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVppTokenAccountType

`func (o *MicrosoftGraphVppToken) HasVppTokenAccountType() bool`

HasVppTokenAccountType returns a boolean if a field has been set.

### SetVppTokenAccountType

`func (o *MicrosoftGraphVppToken) SetVppTokenAccountType(v AnyOfmicrosoftGraphVppTokenAccountType)`

SetVppTokenAccountType gets a reference to the given AnyOfmicrosoftGraphVppTokenAccountType and assigns it to the VppTokenAccountType field.

### GetAppleId

`func (o *MicrosoftGraphVppToken) GetAppleId() string`

GetAppleId returns the AppleId field if non-nil, zero value otherwise.

### GetAppleIdOk

`func (o *MicrosoftGraphVppToken) GetAppleIdOk() (string, bool)`

GetAppleIdOk returns a tuple with the AppleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppleId

`func (o *MicrosoftGraphVppToken) HasAppleId() bool`

HasAppleId returns a boolean if a field has been set.

### SetAppleId

`func (o *MicrosoftGraphVppToken) SetAppleId(v string)`

SetAppleId gets a reference to the given string and assigns it to the AppleId field.

### SetAppleIdExplicitNull

`func (o *MicrosoftGraphVppToken) SetAppleIdExplicitNull(b bool)`

SetAppleIdExplicitNull (un)sets AppleId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppleId value is set to nil even if false is passed
### GetExpirationDateTime

`func (o *MicrosoftGraphVppToken) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphVppToken) GetExpirationDateTimeOk() (time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpirationDateTime

`func (o *MicrosoftGraphVppToken) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphVppToken) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.

### GetLastSyncDateTime

`func (o *MicrosoftGraphVppToken) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *MicrosoftGraphVppToken) GetLastSyncDateTimeOk() (time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSyncDateTime

`func (o *MicrosoftGraphVppToken) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### SetLastSyncDateTime

`func (o *MicrosoftGraphVppToken) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.

### GetToken

`func (o *MicrosoftGraphVppToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MicrosoftGraphVppToken) GetTokenOk() (string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToken

`func (o *MicrosoftGraphVppToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetToken

`func (o *MicrosoftGraphVppToken) SetToken(v string)`

SetToken gets a reference to the given string and assigns it to the Token field.

### SetTokenExplicitNull

`func (o *MicrosoftGraphVppToken) SetTokenExplicitNull(b bool)`

SetTokenExplicitNull (un)sets Token to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Token value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphVppToken) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphVppToken) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphVppToken) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphVppToken) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetState

`func (o *MicrosoftGraphVppToken) GetState() AnyOfmicrosoftGraphVppTokenState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphVppToken) GetStateOk() (AnyOfmicrosoftGraphVppTokenState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *MicrosoftGraphVppToken) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *MicrosoftGraphVppToken) SetState(v AnyOfmicrosoftGraphVppTokenState)`

SetState gets a reference to the given AnyOfmicrosoftGraphVppTokenState and assigns it to the State field.

### GetLastSyncStatus

`func (o *MicrosoftGraphVppToken) GetLastSyncStatus() AnyOfmicrosoftGraphVppTokenSyncStatus`

GetLastSyncStatus returns the LastSyncStatus field if non-nil, zero value otherwise.

### GetLastSyncStatusOk

`func (o *MicrosoftGraphVppToken) GetLastSyncStatusOk() (AnyOfmicrosoftGraphVppTokenSyncStatus, bool)`

GetLastSyncStatusOk returns a tuple with the LastSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSyncStatus

`func (o *MicrosoftGraphVppToken) HasLastSyncStatus() bool`

HasLastSyncStatus returns a boolean if a field has been set.

### SetLastSyncStatus

`func (o *MicrosoftGraphVppToken) SetLastSyncStatus(v AnyOfmicrosoftGraphVppTokenSyncStatus)`

SetLastSyncStatus gets a reference to the given AnyOfmicrosoftGraphVppTokenSyncStatus and assigns it to the LastSyncStatus field.

### GetAutomaticallyUpdateApps

`func (o *MicrosoftGraphVppToken) GetAutomaticallyUpdateApps() bool`

GetAutomaticallyUpdateApps returns the AutomaticallyUpdateApps field if non-nil, zero value otherwise.

### GetAutomaticallyUpdateAppsOk

`func (o *MicrosoftGraphVppToken) GetAutomaticallyUpdateAppsOk() (bool, bool)`

GetAutomaticallyUpdateAppsOk returns a tuple with the AutomaticallyUpdateApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAutomaticallyUpdateApps

`func (o *MicrosoftGraphVppToken) HasAutomaticallyUpdateApps() bool`

HasAutomaticallyUpdateApps returns a boolean if a field has been set.

### SetAutomaticallyUpdateApps

`func (o *MicrosoftGraphVppToken) SetAutomaticallyUpdateApps(v bool)`

SetAutomaticallyUpdateApps gets a reference to the given bool and assigns it to the AutomaticallyUpdateApps field.

### GetCountryOrRegion

`func (o *MicrosoftGraphVppToken) GetCountryOrRegion() string`

GetCountryOrRegion returns the CountryOrRegion field if non-nil, zero value otherwise.

### GetCountryOrRegionOk

`func (o *MicrosoftGraphVppToken) GetCountryOrRegionOk() (string, bool)`

GetCountryOrRegionOk returns a tuple with the CountryOrRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCountryOrRegion

`func (o *MicrosoftGraphVppToken) HasCountryOrRegion() bool`

HasCountryOrRegion returns a boolean if a field has been set.

### SetCountryOrRegion

`func (o *MicrosoftGraphVppToken) SetCountryOrRegion(v string)`

SetCountryOrRegion gets a reference to the given string and assigns it to the CountryOrRegion field.

### SetCountryOrRegionExplicitNull

`func (o *MicrosoftGraphVppToken) SetCountryOrRegionExplicitNull(b bool)`

SetCountryOrRegionExplicitNull (un)sets CountryOrRegion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CountryOrRegion value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


