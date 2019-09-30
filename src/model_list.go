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
// List struct for List
type List struct {
	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	List *AnyOfmicrosoftGraphListInfo `json:"list,omitempty"`
	isExplicitNullList bool `json:"-"`
	SharepointIds *AnyOfmicrosoftGraphSharepointIds `json:"sharepointIds,omitempty"`
	isExplicitNullSharepointIds bool `json:"-"`
	System *AnyOfobject `json:"system,omitempty"`
	isExplicitNullSystem bool `json:"-"`
	Columns *[]MicrosoftGraphColumnDefinition `json:"columns,omitempty"`

	ContentTypes *[]MicrosoftGraphContentType `json:"contentTypes,omitempty"`

	Drive *AnyOfmicrosoftGraphDrive `json:"drive,omitempty"`
	isExplicitNullDrive bool `json:"-"`
	Items *[]MicrosoftGraphListItem `json:"items,omitempty"`

}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *List) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *List) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *List) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *List) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *List) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetList returns the List field if non-nil, zero value otherwise.
func (o *List) GetList() AnyOfmicrosoftGraphListInfo {
	if o == nil || o.List == nil {
		var ret AnyOfmicrosoftGraphListInfo
		return ret
	}
	return *o.List
}

// GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *List) GetListOk() (AnyOfmicrosoftGraphListInfo, bool) {
	if o == nil || o.List == nil {
		var ret AnyOfmicrosoftGraphListInfo
		return ret, false
	}
	return *o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *List) HasList() bool {
	if o != nil && o.List != nil {
		return true
	}

	return false
}

// SetList gets a reference to the given AnyOfmicrosoftGraphListInfo and assigns it to the List field.
func (o *List) SetList(v AnyOfmicrosoftGraphListInfo) {
	o.List = &v
}

// SetListExplicitNull (un)sets List to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The List value is set to nil even if false is passed
func (o *List) SetListExplicitNull(b bool) {
	o.List = nil
	o.isExplicitNullList = b
}
// GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.
func (o *List) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds {
	if o == nil || o.SharepointIds == nil {
		var ret AnyOfmicrosoftGraphSharepointIds
		return ret
	}
	return *o.SharepointIds
}

// GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *List) GetSharepointIdsOk() (AnyOfmicrosoftGraphSharepointIds, bool) {
	if o == nil || o.SharepointIds == nil {
		var ret AnyOfmicrosoftGraphSharepointIds
		return ret, false
	}
	return *o.SharepointIds, true
}

// HasSharepointIds returns a boolean if a field has been set.
func (o *List) HasSharepointIds() bool {
	if o != nil && o.SharepointIds != nil {
		return true
	}

	return false
}

// SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.
func (o *List) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds) {
	o.SharepointIds = &v
}

// SetSharepointIdsExplicitNull (un)sets SharepointIds to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SharepointIds value is set to nil even if false is passed
func (o *List) SetSharepointIdsExplicitNull(b bool) {
	o.SharepointIds = nil
	o.isExplicitNullSharepointIds = b
}
// GetSystem returns the System field if non-nil, zero value otherwise.
func (o *List) GetSystem() AnyOfobject {
	if o == nil || o.System == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *List) GetSystemOk() (AnyOfobject, bool) {
	if o == nil || o.System == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *List) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given AnyOfobject and assigns it to the System field.
func (o *List) SetSystem(v AnyOfobject) {
	o.System = &v
}

// SetSystemExplicitNull (un)sets System to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The System value is set to nil even if false is passed
func (o *List) SetSystemExplicitNull(b bool) {
	o.System = nil
	o.isExplicitNullSystem = b
}
// GetColumns returns the Columns field if non-nil, zero value otherwise.
func (o *List) GetColumns() []MicrosoftGraphColumnDefinition {
	if o == nil || o.Columns == nil {
		var ret []MicrosoftGraphColumnDefinition
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *List) GetColumnsOk() ([]MicrosoftGraphColumnDefinition, bool) {
	if o == nil || o.Columns == nil {
		var ret []MicrosoftGraphColumnDefinition
		return ret, false
	}
	return *o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *List) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []MicrosoftGraphColumnDefinition and assigns it to the Columns field.
func (o *List) SetColumns(v []MicrosoftGraphColumnDefinition) {
	o.Columns = &v
}

// GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.
func (o *List) GetContentTypes() []MicrosoftGraphContentType {
	if o == nil || o.ContentTypes == nil {
		var ret []MicrosoftGraphContentType
		return ret
	}
	return *o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *List) GetContentTypesOk() ([]MicrosoftGraphContentType, bool) {
	if o == nil || o.ContentTypes == nil {
		var ret []MicrosoftGraphContentType
		return ret, false
	}
	return *o.ContentTypes, true
}

// HasContentTypes returns a boolean if a field has been set.
func (o *List) HasContentTypes() bool {
	if o != nil && o.ContentTypes != nil {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given []MicrosoftGraphContentType and assigns it to the ContentTypes field.
func (o *List) SetContentTypes(v []MicrosoftGraphContentType) {
	o.ContentTypes = &v
}

// GetDrive returns the Drive field if non-nil, zero value otherwise.
func (o *List) GetDrive() AnyOfmicrosoftGraphDrive {
	if o == nil || o.Drive == nil {
		var ret AnyOfmicrosoftGraphDrive
		return ret
	}
	return *o.Drive
}

// GetDriveOk returns a tuple with the Drive field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *List) GetDriveOk() (AnyOfmicrosoftGraphDrive, bool) {
	if o == nil || o.Drive == nil {
		var ret AnyOfmicrosoftGraphDrive
		return ret, false
	}
	return *o.Drive, true
}

// HasDrive returns a boolean if a field has been set.
func (o *List) HasDrive() bool {
	if o != nil && o.Drive != nil {
		return true
	}

	return false
}

// SetDrive gets a reference to the given AnyOfmicrosoftGraphDrive and assigns it to the Drive field.
func (o *List) SetDrive(v AnyOfmicrosoftGraphDrive) {
	o.Drive = &v
}

// SetDriveExplicitNull (un)sets Drive to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Drive value is set to nil even if false is passed
func (o *List) SetDriveExplicitNull(b bool) {
	o.Drive = nil
	o.isExplicitNullDrive = b
}
// GetItems returns the Items field if non-nil, zero value otherwise.
func (o *List) GetItems() []MicrosoftGraphListItem {
	if o == nil || o.Items == nil {
		var ret []MicrosoftGraphListItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *List) GetItemsOk() ([]MicrosoftGraphListItem, bool) {
	if o == nil || o.Items == nil {
		var ret []MicrosoftGraphListItem
		return ret, false
	}
	return *o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *List) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MicrosoftGraphListItem and assigns it to the Items field.
func (o *List) SetItems(v []MicrosoftGraphListItem) {
	o.Items = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o List) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.List == nil {
		if o.isExplicitNullList {
			toSerialize["list"] = o.List
		}
	} else {
		toSerialize["list"] = o.List
	}
	if o.SharepointIds == nil {
		if o.isExplicitNullSharepointIds {
			toSerialize["sharepointIds"] = o.SharepointIds
		}
	} else {
		toSerialize["sharepointIds"] = o.SharepointIds
	}
	if o.System == nil {
		if o.isExplicitNullSystem {
			toSerialize["system"] = o.System
		}
	} else {
		toSerialize["system"] = o.System
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.ContentTypes != nil {
		toSerialize["contentTypes"] = o.ContentTypes
	}
	if o.Drive == nil {
		if o.isExplicitNullDrive {
			toSerialize["drive"] = o.Drive
		}
	} else {
		toSerialize["drive"] = o.Drive
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}


