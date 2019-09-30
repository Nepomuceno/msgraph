# MicrosoftGraphAuditLogRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SignIns** | Pointer to [**[]MicrosoftGraphSignIn**](microsoft.graph.signIn.md) |  | [optional] 
**DirectoryAudits** | Pointer to [**[]MicrosoftGraphDirectoryAudit**](microsoft.graph.directoryAudit.md) |  | [optional] 
**RestrictedSignIns** | Pointer to [**[]MicrosoftGraphRestrictedSignIn**](microsoft.graph.restrictedSignIn.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphAuditLogRoot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAuditLogRoot) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphAuditLogRoot) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphAuditLogRoot) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetSignIns

`func (o *MicrosoftGraphAuditLogRoot) GetSignIns() []MicrosoftGraphSignIn`

GetSignIns returns the SignIns field if non-nil, zero value otherwise.

### GetSignInsOk

`func (o *MicrosoftGraphAuditLogRoot) GetSignInsOk() ([]MicrosoftGraphSignIn, bool)`

GetSignInsOk returns a tuple with the SignIns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSignIns

`func (o *MicrosoftGraphAuditLogRoot) HasSignIns() bool`

HasSignIns returns a boolean if a field has been set.

### SetSignIns

`func (o *MicrosoftGraphAuditLogRoot) SetSignIns(v []MicrosoftGraphSignIn)`

SetSignIns gets a reference to the given []MicrosoftGraphSignIn and assigns it to the SignIns field.

### GetDirectoryAudits

`func (o *MicrosoftGraphAuditLogRoot) GetDirectoryAudits() []MicrosoftGraphDirectoryAudit`

GetDirectoryAudits returns the DirectoryAudits field if non-nil, zero value otherwise.

### GetDirectoryAuditsOk

`func (o *MicrosoftGraphAuditLogRoot) GetDirectoryAuditsOk() ([]MicrosoftGraphDirectoryAudit, bool)`

GetDirectoryAuditsOk returns a tuple with the DirectoryAudits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirectoryAudits

`func (o *MicrosoftGraphAuditLogRoot) HasDirectoryAudits() bool`

HasDirectoryAudits returns a boolean if a field has been set.

### SetDirectoryAudits

`func (o *MicrosoftGraphAuditLogRoot) SetDirectoryAudits(v []MicrosoftGraphDirectoryAudit)`

SetDirectoryAudits gets a reference to the given []MicrosoftGraphDirectoryAudit and assigns it to the DirectoryAudits field.

### GetRestrictedSignIns

`func (o *MicrosoftGraphAuditLogRoot) GetRestrictedSignIns() []MicrosoftGraphRestrictedSignIn`

GetRestrictedSignIns returns the RestrictedSignIns field if non-nil, zero value otherwise.

### GetRestrictedSignInsOk

`func (o *MicrosoftGraphAuditLogRoot) GetRestrictedSignInsOk() ([]MicrosoftGraphRestrictedSignIn, bool)`

GetRestrictedSignInsOk returns a tuple with the RestrictedSignIns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRestrictedSignIns

`func (o *MicrosoftGraphAuditLogRoot) HasRestrictedSignIns() bool`

HasRestrictedSignIns returns a boolean if a field has been set.

### SetRestrictedSignIns

`func (o *MicrosoftGraphAuditLogRoot) SetRestrictedSignIns(v []MicrosoftGraphRestrictedSignIn)`

SetRestrictedSignIns gets a reference to the given []MicrosoftGraphRestrictedSignIn and assigns it to the RestrictedSignIns field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


