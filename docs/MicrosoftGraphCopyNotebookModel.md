# MicrosoftGraphCopyNotebookModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **bool** |  | [optional] 
**UserRole** | Pointer to [**AnyOfmicrosoftGraphOnenoteUserRole**](anyOf&lt;microsoft.graph.onenoteUserRole&gt;.md) |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**SectionsUrl** | Pointer to **string** |  | [optional] 
**SectionGroupsUrl** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphNotebookLinks**](anyOf&lt;microsoft.graph.notebookLinks&gt;.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedByIdentity** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 
**LastModifiedByIdentity** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**LastModifiedTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Self** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### GetIsDefault

`func (o *MicrosoftGraphCopyNotebookModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphCopyNotebookModel) GetIsDefaultOk() (bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDefault

`func (o *MicrosoftGraphCopyNotebookModel) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefault

`func (o *MicrosoftGraphCopyNotebookModel) SetIsDefault(v bool)`

SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.

### SetIsDefaultExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetIsDefaultExplicitNull(b bool)`

SetIsDefaultExplicitNull (un)sets IsDefault to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsDefault value is set to nil even if false is passed
### GetUserRole

`func (o *MicrosoftGraphCopyNotebookModel) GetUserRole() AnyOfmicrosoftGraphOnenoteUserRole`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *MicrosoftGraphCopyNotebookModel) GetUserRoleOk() (AnyOfmicrosoftGraphOnenoteUserRole, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserRole

`func (o *MicrosoftGraphCopyNotebookModel) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### SetUserRole

`func (o *MicrosoftGraphCopyNotebookModel) SetUserRole(v AnyOfmicrosoftGraphOnenoteUserRole)`

SetUserRole gets a reference to the given AnyOfmicrosoftGraphOnenoteUserRole and assigns it to the UserRole field.

### SetUserRoleExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetUserRoleExplicitNull(b bool)`

SetUserRoleExplicitNull (un)sets UserRole to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserRole value is set to nil even if false is passed
### GetIsShared

`func (o *MicrosoftGraphCopyNotebookModel) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *MicrosoftGraphCopyNotebookModel) GetIsSharedOk() (bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsShared

`func (o *MicrosoftGraphCopyNotebookModel) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### SetIsShared

`func (o *MicrosoftGraphCopyNotebookModel) SetIsShared(v bool)`

SetIsShared gets a reference to the given bool and assigns it to the IsShared field.

### SetIsSharedExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetIsSharedExplicitNull(b bool)`

SetIsSharedExplicitNull (un)sets IsShared to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsShared value is set to nil even if false is passed
### GetSectionsUrl

`func (o *MicrosoftGraphCopyNotebookModel) GetSectionsUrl() string`

GetSectionsUrl returns the SectionsUrl field if non-nil, zero value otherwise.

### GetSectionsUrlOk

`func (o *MicrosoftGraphCopyNotebookModel) GetSectionsUrlOk() (string, bool)`

GetSectionsUrlOk returns a tuple with the SectionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionsUrl

`func (o *MicrosoftGraphCopyNotebookModel) HasSectionsUrl() bool`

HasSectionsUrl returns a boolean if a field has been set.

### SetSectionsUrl

`func (o *MicrosoftGraphCopyNotebookModel) SetSectionsUrl(v string)`

SetSectionsUrl gets a reference to the given string and assigns it to the SectionsUrl field.

### SetSectionsUrlExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetSectionsUrlExplicitNull(b bool)`

SetSectionsUrlExplicitNull (un)sets SectionsUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SectionsUrl value is set to nil even if false is passed
### GetSectionGroupsUrl

`func (o *MicrosoftGraphCopyNotebookModel) GetSectionGroupsUrl() string`

GetSectionGroupsUrl returns the SectionGroupsUrl field if non-nil, zero value otherwise.

### GetSectionGroupsUrlOk

`func (o *MicrosoftGraphCopyNotebookModel) GetSectionGroupsUrlOk() (string, bool)`

GetSectionGroupsUrlOk returns a tuple with the SectionGroupsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSectionGroupsUrl

`func (o *MicrosoftGraphCopyNotebookModel) HasSectionGroupsUrl() bool`

HasSectionGroupsUrl returns a boolean if a field has been set.

### SetSectionGroupsUrl

`func (o *MicrosoftGraphCopyNotebookModel) SetSectionGroupsUrl(v string)`

SetSectionGroupsUrl gets a reference to the given string and assigns it to the SectionGroupsUrl field.

### SetSectionGroupsUrlExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetSectionGroupsUrlExplicitNull(b bool)`

SetSectionGroupsUrlExplicitNull (un)sets SectionGroupsUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SectionGroupsUrl value is set to nil even if false is passed
### GetLinks

`func (o *MicrosoftGraphCopyNotebookModel) GetLinks() AnyOfmicrosoftGraphNotebookLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MicrosoftGraphCopyNotebookModel) GetLinksOk() (AnyOfmicrosoftGraphNotebookLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinks

`func (o *MicrosoftGraphCopyNotebookModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinks

`func (o *MicrosoftGraphCopyNotebookModel) SetLinks(v AnyOfmicrosoftGraphNotebookLinks)`

SetLinks gets a reference to the given AnyOfmicrosoftGraphNotebookLinks and assigns it to the Links field.

### SetLinksExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetLinksExplicitNull(b bool)`

SetLinksExplicitNull (un)sets Links to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Links value is set to nil even if false is passed
### GetName

`func (o *MicrosoftGraphCopyNotebookModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphCopyNotebookModel) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphCopyNotebookModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphCopyNotebookModel) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetCreatedBy

`func (o *MicrosoftGraphCopyNotebookModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphCopyNotebookModel) GetCreatedByOk() (string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphCopyNotebookModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphCopyNotebookModel) SetCreatedBy(v string)`

SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetCreatedByIdentity

`func (o *MicrosoftGraphCopyNotebookModel) GetCreatedByIdentity() AnyOfmicrosoftGraphIdentitySet`

GetCreatedByIdentity returns the CreatedByIdentity field if non-nil, zero value otherwise.

### GetCreatedByIdentityOk

`func (o *MicrosoftGraphCopyNotebookModel) GetCreatedByIdentityOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByIdentityOk returns a tuple with the CreatedByIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedByIdentity

`func (o *MicrosoftGraphCopyNotebookModel) HasCreatedByIdentity() bool`

HasCreatedByIdentity returns a boolean if a field has been set.

### SetCreatedByIdentity

`func (o *MicrosoftGraphCopyNotebookModel) SetCreatedByIdentity(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedByIdentity gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedByIdentity field.

### SetCreatedByIdentityExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetCreatedByIdentityExplicitNull(b bool)`

SetCreatedByIdentityExplicitNull (un)sets CreatedByIdentity to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedByIdentity value is set to nil even if false is passed
### GetLastModifiedBy

`func (o *MicrosoftGraphCopyNotebookModel) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphCopyNotebookModel) GetLastModifiedByOk() (string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedBy

`func (o *MicrosoftGraphCopyNotebookModel) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphCopyNotebookModel) SetLastModifiedBy(v string)`

SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.

### SetLastModifiedByExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetLastModifiedByExplicitNull(b bool)`

SetLastModifiedByExplicitNull (un)sets LastModifiedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedBy value is set to nil even if false is passed
### GetLastModifiedByIdentity

`func (o *MicrosoftGraphCopyNotebookModel) GetLastModifiedByIdentity() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedByIdentity returns the LastModifiedByIdentity field if non-nil, zero value otherwise.

### GetLastModifiedByIdentityOk

`func (o *MicrosoftGraphCopyNotebookModel) GetLastModifiedByIdentityOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByIdentityOk returns a tuple with the LastModifiedByIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedByIdentity

`func (o *MicrosoftGraphCopyNotebookModel) HasLastModifiedByIdentity() bool`

HasLastModifiedByIdentity returns a boolean if a field has been set.

### SetLastModifiedByIdentity

`func (o *MicrosoftGraphCopyNotebookModel) SetLastModifiedByIdentity(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedByIdentity gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedByIdentity field.

### SetLastModifiedByIdentityExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetLastModifiedByIdentityExplicitNull(b bool)`

SetLastModifiedByIdentityExplicitNull (un)sets LastModifiedByIdentity to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedByIdentity value is set to nil even if false is passed
### GetLastModifiedTime

`func (o *MicrosoftGraphCopyNotebookModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *MicrosoftGraphCopyNotebookModel) GetLastModifiedTimeOk() (time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedTime

`func (o *MicrosoftGraphCopyNotebookModel) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### SetLastModifiedTime

`func (o *MicrosoftGraphCopyNotebookModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.

### SetLastModifiedTimeExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetLastModifiedTimeExplicitNull(b bool)`

SetLastModifiedTimeExplicitNull (un)sets LastModifiedTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedTime value is set to nil even if false is passed
### GetId

`func (o *MicrosoftGraphCopyNotebookModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphCopyNotebookModel) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphCopyNotebookModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphCopyNotebookModel) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### SetIdExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetIdExplicitNull(b bool)`

SetIdExplicitNull (un)sets Id to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Id value is set to nil even if false is passed
### GetSelf

`func (o *MicrosoftGraphCopyNotebookModel) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MicrosoftGraphCopyNotebookModel) GetSelfOk() (string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSelf

`func (o *MicrosoftGraphCopyNotebookModel) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### SetSelf

`func (o *MicrosoftGraphCopyNotebookModel) SetSelf(v string)`

SetSelf gets a reference to the given string and assigns it to the Self field.

### SetSelfExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetSelfExplicitNull(b bool)`

SetSelfExplicitNull (un)sets Self to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Self value is set to nil even if false is passed
### GetCreatedTime

`func (o *MicrosoftGraphCopyNotebookModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *MicrosoftGraphCopyNotebookModel) GetCreatedTimeOk() (time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedTime

`func (o *MicrosoftGraphCopyNotebookModel) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTime

`func (o *MicrosoftGraphCopyNotebookModel) SetCreatedTime(v time.Time)`

SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.

### SetCreatedTimeExplicitNull

`func (o *MicrosoftGraphCopyNotebookModel) SetCreatedTimeExplicitNull(b bool)`

SetCreatedTimeExplicitNull (un)sets CreatedTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


