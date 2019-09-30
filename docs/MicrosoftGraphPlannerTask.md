# MicrosoftGraphPlannerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**BucketId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**OrderHint** | Pointer to **string** |  | [optional] 
**AssigneePriority** | Pointer to **string** |  | [optional] 
**PercentComplete** | Pointer to **int32** |  | [optional] 
**StartDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**DueDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**HasDescription** | Pointer to **bool** |  | [optional] 
**PreviewType** | Pointer to [**AnyOfmicrosoftGraphPlannerPreviewType**](anyOf&lt;microsoft.graph.plannerPreviewType&gt;.md) |  | [optional] 
**CompletedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**CompletedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**ReferenceCount** | Pointer to **int32** |  | [optional] 
**ChecklistItemCount** | Pointer to **int32** |  | [optional] 
**ActiveChecklistItemCount** | Pointer to **int32** |  | [optional] 
**AppliedCategories** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Assignments** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**ConversationThreadId** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**AnyOfmicrosoftGraphPlannerTaskDetails**](anyOf&lt;microsoft.graph.plannerTaskDetails&gt;.md) |  | [optional] 
**AssignedToTaskBoardFormat** | Pointer to [**AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat**](anyOf&lt;microsoft.graph.plannerAssignedToTaskBoardTaskFormat&gt;.md) |  | [optional] 
**ProgressTaskBoardFormat** | Pointer to [**AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat**](anyOf&lt;microsoft.graph.plannerProgressTaskBoardTaskFormat&gt;.md) |  | [optional] 
**BucketTaskBoardFormat** | Pointer to [**AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat**](anyOf&lt;microsoft.graph.plannerBucketTaskBoardTaskFormat&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphPlannerTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPlannerTask) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphPlannerTask) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphPlannerTask) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedBy

`func (o *MicrosoftGraphPlannerTask) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphPlannerTask) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *MicrosoftGraphPlannerTask) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *MicrosoftGraphPlannerTask) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.

### SetCreatedByExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetCreatedByExplicitNull(b bool)`

SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedBy value is set to nil even if false is passed
### GetPlanId

`func (o *MicrosoftGraphPlannerTask) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *MicrosoftGraphPlannerTask) GetPlanIdOk() (string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlanId

`func (o *MicrosoftGraphPlannerTask) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanId

`func (o *MicrosoftGraphPlannerTask) SetPlanId(v string)`

SetPlanId gets a reference to the given string and assigns it to the PlanId field.

### SetPlanIdExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetPlanIdExplicitNull(b bool)`

SetPlanIdExplicitNull (un)sets PlanId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PlanId value is set to nil even if false is passed
### GetBucketId

`func (o *MicrosoftGraphPlannerTask) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *MicrosoftGraphPlannerTask) GetBucketIdOk() (string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBucketId

`func (o *MicrosoftGraphPlannerTask) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### SetBucketId

`func (o *MicrosoftGraphPlannerTask) SetBucketId(v string)`

SetBucketId gets a reference to the given string and assigns it to the BucketId field.

### SetBucketIdExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetBucketIdExplicitNull(b bool)`

SetBucketIdExplicitNull (un)sets BucketId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BucketId value is set to nil even if false is passed
### GetTitle

`func (o *MicrosoftGraphPlannerTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphPlannerTask) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *MicrosoftGraphPlannerTask) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *MicrosoftGraphPlannerTask) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetOrderHint

`func (o *MicrosoftGraphPlannerTask) GetOrderHint() string`

GetOrderHint returns the OrderHint field if non-nil, zero value otherwise.

### GetOrderHintOk

`func (o *MicrosoftGraphPlannerTask) GetOrderHintOk() (string, bool)`

GetOrderHintOk returns a tuple with the OrderHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrderHint

`func (o *MicrosoftGraphPlannerTask) HasOrderHint() bool`

HasOrderHint returns a boolean if a field has been set.

### SetOrderHint

`func (o *MicrosoftGraphPlannerTask) SetOrderHint(v string)`

SetOrderHint gets a reference to the given string and assigns it to the OrderHint field.

### SetOrderHintExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetOrderHintExplicitNull(b bool)`

SetOrderHintExplicitNull (un)sets OrderHint to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OrderHint value is set to nil even if false is passed
### GetAssigneePriority

`func (o *MicrosoftGraphPlannerTask) GetAssigneePriority() string`

GetAssigneePriority returns the AssigneePriority field if non-nil, zero value otherwise.

### GetAssigneePriorityOk

`func (o *MicrosoftGraphPlannerTask) GetAssigneePriorityOk() (string, bool)`

GetAssigneePriorityOk returns a tuple with the AssigneePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssigneePriority

`func (o *MicrosoftGraphPlannerTask) HasAssigneePriority() bool`

HasAssigneePriority returns a boolean if a field has been set.

### SetAssigneePriority

`func (o *MicrosoftGraphPlannerTask) SetAssigneePriority(v string)`

SetAssigneePriority gets a reference to the given string and assigns it to the AssigneePriority field.

### SetAssigneePriorityExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetAssigneePriorityExplicitNull(b bool)`

SetAssigneePriorityExplicitNull (un)sets AssigneePriority to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AssigneePriority value is set to nil even if false is passed
### GetPercentComplete

`func (o *MicrosoftGraphPlannerTask) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *MicrosoftGraphPlannerTask) GetPercentCompleteOk() (int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPercentComplete

`func (o *MicrosoftGraphPlannerTask) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentComplete

`func (o *MicrosoftGraphPlannerTask) SetPercentComplete(v int32)`

SetPercentComplete gets a reference to the given int32 and assigns it to the PercentComplete field.

### SetPercentCompleteExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetPercentCompleteExplicitNull(b bool)`

SetPercentCompleteExplicitNull (un)sets PercentComplete to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PercentComplete value is set to nil even if false is passed
### GetStartDateTime

`func (o *MicrosoftGraphPlannerTask) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphPlannerTask) GetStartDateTimeOk() (time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDateTime

`func (o *MicrosoftGraphPlannerTask) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTime

`func (o *MicrosoftGraphPlannerTask) SetStartDateTime(v time.Time)`

SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.

### SetStartDateTimeExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetStartDateTimeExplicitNull(b bool)`

SetStartDateTimeExplicitNull (un)sets StartDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StartDateTime value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphPlannerTask) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphPlannerTask) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphPlannerTask) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphPlannerTask) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetDueDateTime

`func (o *MicrosoftGraphPlannerTask) GetDueDateTime() time.Time`

GetDueDateTime returns the DueDateTime field if non-nil, zero value otherwise.

### GetDueDateTimeOk

`func (o *MicrosoftGraphPlannerTask) GetDueDateTimeOk() (time.Time, bool)`

GetDueDateTimeOk returns a tuple with the DueDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDueDateTime

`func (o *MicrosoftGraphPlannerTask) HasDueDateTime() bool`

HasDueDateTime returns a boolean if a field has been set.

### SetDueDateTime

`func (o *MicrosoftGraphPlannerTask) SetDueDateTime(v time.Time)`

SetDueDateTime gets a reference to the given time.Time and assigns it to the DueDateTime field.

### SetDueDateTimeExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetDueDateTimeExplicitNull(b bool)`

SetDueDateTimeExplicitNull (un)sets DueDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DueDateTime value is set to nil even if false is passed
### GetHasDescription

`func (o *MicrosoftGraphPlannerTask) GetHasDescription() bool`

GetHasDescription returns the HasDescription field if non-nil, zero value otherwise.

### GetHasDescriptionOk

`func (o *MicrosoftGraphPlannerTask) GetHasDescriptionOk() (bool, bool)`

GetHasDescriptionOk returns a tuple with the HasDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasDescription

`func (o *MicrosoftGraphPlannerTask) HasHasDescription() bool`

HasHasDescription returns a boolean if a field has been set.

### SetHasDescription

`func (o *MicrosoftGraphPlannerTask) SetHasDescription(v bool)`

SetHasDescription gets a reference to the given bool and assigns it to the HasDescription field.

### SetHasDescriptionExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetHasDescriptionExplicitNull(b bool)`

SetHasDescriptionExplicitNull (un)sets HasDescription to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HasDescription value is set to nil even if false is passed
### GetPreviewType

`func (o *MicrosoftGraphPlannerTask) GetPreviewType() AnyOfmicrosoftGraphPlannerPreviewType`

GetPreviewType returns the PreviewType field if non-nil, zero value otherwise.

### GetPreviewTypeOk

`func (o *MicrosoftGraphPlannerTask) GetPreviewTypeOk() (AnyOfmicrosoftGraphPlannerPreviewType, bool)`

GetPreviewTypeOk returns a tuple with the PreviewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreviewType

`func (o *MicrosoftGraphPlannerTask) HasPreviewType() bool`

HasPreviewType returns a boolean if a field has been set.

### SetPreviewType

`func (o *MicrosoftGraphPlannerTask) SetPreviewType(v AnyOfmicrosoftGraphPlannerPreviewType)`

SetPreviewType gets a reference to the given AnyOfmicrosoftGraphPlannerPreviewType and assigns it to the PreviewType field.

### SetPreviewTypeExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetPreviewTypeExplicitNull(b bool)`

SetPreviewTypeExplicitNull (un)sets PreviewType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PreviewType value is set to nil even if false is passed
### GetCompletedDateTime

`func (o *MicrosoftGraphPlannerTask) GetCompletedDateTime() time.Time`

GetCompletedDateTime returns the CompletedDateTime field if non-nil, zero value otherwise.

### GetCompletedDateTimeOk

`func (o *MicrosoftGraphPlannerTask) GetCompletedDateTimeOk() (time.Time, bool)`

GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompletedDateTime

`func (o *MicrosoftGraphPlannerTask) HasCompletedDateTime() bool`

HasCompletedDateTime returns a boolean if a field has been set.

### SetCompletedDateTime

`func (o *MicrosoftGraphPlannerTask) SetCompletedDateTime(v time.Time)`

SetCompletedDateTime gets a reference to the given time.Time and assigns it to the CompletedDateTime field.

### SetCompletedDateTimeExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetCompletedDateTimeExplicitNull(b bool)`

SetCompletedDateTimeExplicitNull (un)sets CompletedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CompletedDateTime value is set to nil even if false is passed
### GetCompletedBy

`func (o *MicrosoftGraphPlannerTask) GetCompletedBy() AnyOfmicrosoftGraphIdentitySet`

GetCompletedBy returns the CompletedBy field if non-nil, zero value otherwise.

### GetCompletedByOk

`func (o *MicrosoftGraphPlannerTask) GetCompletedByOk() (AnyOfmicrosoftGraphIdentitySet, bool)`

GetCompletedByOk returns a tuple with the CompletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompletedBy

`func (o *MicrosoftGraphPlannerTask) HasCompletedBy() bool`

HasCompletedBy returns a boolean if a field has been set.

### SetCompletedBy

`func (o *MicrosoftGraphPlannerTask) SetCompletedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCompletedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CompletedBy field.

### SetCompletedByExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetCompletedByExplicitNull(b bool)`

SetCompletedByExplicitNull (un)sets CompletedBy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CompletedBy value is set to nil even if false is passed
### GetReferenceCount

`func (o *MicrosoftGraphPlannerTask) GetReferenceCount() int32`

GetReferenceCount returns the ReferenceCount field if non-nil, zero value otherwise.

### GetReferenceCountOk

`func (o *MicrosoftGraphPlannerTask) GetReferenceCountOk() (int32, bool)`

GetReferenceCountOk returns a tuple with the ReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferenceCount

`func (o *MicrosoftGraphPlannerTask) HasReferenceCount() bool`

HasReferenceCount returns a boolean if a field has been set.

### SetReferenceCount

`func (o *MicrosoftGraphPlannerTask) SetReferenceCount(v int32)`

SetReferenceCount gets a reference to the given int32 and assigns it to the ReferenceCount field.

### SetReferenceCountExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetReferenceCountExplicitNull(b bool)`

SetReferenceCountExplicitNull (un)sets ReferenceCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReferenceCount value is set to nil even if false is passed
### GetChecklistItemCount

`func (o *MicrosoftGraphPlannerTask) GetChecklistItemCount() int32`

GetChecklistItemCount returns the ChecklistItemCount field if non-nil, zero value otherwise.

### GetChecklistItemCountOk

`func (o *MicrosoftGraphPlannerTask) GetChecklistItemCountOk() (int32, bool)`

GetChecklistItemCountOk returns a tuple with the ChecklistItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChecklistItemCount

`func (o *MicrosoftGraphPlannerTask) HasChecklistItemCount() bool`

HasChecklistItemCount returns a boolean if a field has been set.

### SetChecklistItemCount

`func (o *MicrosoftGraphPlannerTask) SetChecklistItemCount(v int32)`

SetChecklistItemCount gets a reference to the given int32 and assigns it to the ChecklistItemCount field.

### SetChecklistItemCountExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetChecklistItemCountExplicitNull(b bool)`

SetChecklistItemCountExplicitNull (un)sets ChecklistItemCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChecklistItemCount value is set to nil even if false is passed
### GetActiveChecklistItemCount

`func (o *MicrosoftGraphPlannerTask) GetActiveChecklistItemCount() int32`

GetActiveChecklistItemCount returns the ActiveChecklistItemCount field if non-nil, zero value otherwise.

### GetActiveChecklistItemCountOk

`func (o *MicrosoftGraphPlannerTask) GetActiveChecklistItemCountOk() (int32, bool)`

GetActiveChecklistItemCountOk returns a tuple with the ActiveChecklistItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveChecklistItemCount

`func (o *MicrosoftGraphPlannerTask) HasActiveChecklistItemCount() bool`

HasActiveChecklistItemCount returns a boolean if a field has been set.

### SetActiveChecklistItemCount

`func (o *MicrosoftGraphPlannerTask) SetActiveChecklistItemCount(v int32)`

SetActiveChecklistItemCount gets a reference to the given int32 and assigns it to the ActiveChecklistItemCount field.

### SetActiveChecklistItemCountExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetActiveChecklistItemCountExplicitNull(b bool)`

SetActiveChecklistItemCountExplicitNull (un)sets ActiveChecklistItemCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ActiveChecklistItemCount value is set to nil even if false is passed
### GetAppliedCategories

`func (o *MicrosoftGraphPlannerTask) GetAppliedCategories() AnyOfobject`

GetAppliedCategories returns the AppliedCategories field if non-nil, zero value otherwise.

### GetAppliedCategoriesOk

`func (o *MicrosoftGraphPlannerTask) GetAppliedCategoriesOk() (AnyOfobject, bool)`

GetAppliedCategoriesOk returns a tuple with the AppliedCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppliedCategories

`func (o *MicrosoftGraphPlannerTask) HasAppliedCategories() bool`

HasAppliedCategories returns a boolean if a field has been set.

### SetAppliedCategories

`func (o *MicrosoftGraphPlannerTask) SetAppliedCategories(v AnyOfobject)`

SetAppliedCategories gets a reference to the given AnyOfobject and assigns it to the AppliedCategories field.

### SetAppliedCategoriesExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetAppliedCategoriesExplicitNull(b bool)`

SetAppliedCategoriesExplicitNull (un)sets AppliedCategories to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppliedCategories value is set to nil even if false is passed
### GetAssignments

`func (o *MicrosoftGraphPlannerTask) GetAssignments() AnyOfobject`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphPlannerTask) GetAssignmentsOk() (AnyOfobject, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphPlannerTask) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphPlannerTask) SetAssignments(v AnyOfobject)`

SetAssignments gets a reference to the given AnyOfobject and assigns it to the Assignments field.

### SetAssignmentsExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetAssignmentsExplicitNull(b bool)`

SetAssignmentsExplicitNull (un)sets Assignments to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Assignments value is set to nil even if false is passed
### GetConversationThreadId

`func (o *MicrosoftGraphPlannerTask) GetConversationThreadId() string`

GetConversationThreadId returns the ConversationThreadId field if non-nil, zero value otherwise.

### GetConversationThreadIdOk

`func (o *MicrosoftGraphPlannerTask) GetConversationThreadIdOk() (string, bool)`

GetConversationThreadIdOk returns a tuple with the ConversationThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConversationThreadId

`func (o *MicrosoftGraphPlannerTask) HasConversationThreadId() bool`

HasConversationThreadId returns a boolean if a field has been set.

### SetConversationThreadId

`func (o *MicrosoftGraphPlannerTask) SetConversationThreadId(v string)`

SetConversationThreadId gets a reference to the given string and assigns it to the ConversationThreadId field.

### SetConversationThreadIdExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetConversationThreadIdExplicitNull(b bool)`

SetConversationThreadIdExplicitNull (un)sets ConversationThreadId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConversationThreadId value is set to nil even if false is passed
### GetDetails

`func (o *MicrosoftGraphPlannerTask) GetDetails() AnyOfmicrosoftGraphPlannerTaskDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphPlannerTask) GetDetailsOk() (AnyOfmicrosoftGraphPlannerTaskDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetails

`func (o *MicrosoftGraphPlannerTask) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetails

`func (o *MicrosoftGraphPlannerTask) SetDetails(v AnyOfmicrosoftGraphPlannerTaskDetails)`

SetDetails gets a reference to the given AnyOfmicrosoftGraphPlannerTaskDetails and assigns it to the Details field.

### SetDetailsExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetDetailsExplicitNull(b bool)`

SetDetailsExplicitNull (un)sets Details to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Details value is set to nil even if false is passed
### GetAssignedToTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) GetAssignedToTaskBoardFormat() AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat`

GetAssignedToTaskBoardFormat returns the AssignedToTaskBoardFormat field if non-nil, zero value otherwise.

### GetAssignedToTaskBoardFormatOk

`func (o *MicrosoftGraphPlannerTask) GetAssignedToTaskBoardFormatOk() (AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat, bool)`

GetAssignedToTaskBoardFormatOk returns a tuple with the AssignedToTaskBoardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignedToTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) HasAssignedToTaskBoardFormat() bool`

HasAssignedToTaskBoardFormat returns a boolean if a field has been set.

### SetAssignedToTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) SetAssignedToTaskBoardFormat(v AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat)`

SetAssignedToTaskBoardFormat gets a reference to the given AnyOfmicrosoftGraphPlannerAssignedToTaskBoardTaskFormat and assigns it to the AssignedToTaskBoardFormat field.

### SetAssignedToTaskBoardFormatExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetAssignedToTaskBoardFormatExplicitNull(b bool)`

SetAssignedToTaskBoardFormatExplicitNull (un)sets AssignedToTaskBoardFormat to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AssignedToTaskBoardFormat value is set to nil even if false is passed
### GetProgressTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) GetProgressTaskBoardFormat() AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat`

GetProgressTaskBoardFormat returns the ProgressTaskBoardFormat field if non-nil, zero value otherwise.

### GetProgressTaskBoardFormatOk

`func (o *MicrosoftGraphPlannerTask) GetProgressTaskBoardFormatOk() (AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat, bool)`

GetProgressTaskBoardFormatOk returns a tuple with the ProgressTaskBoardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgressTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) HasProgressTaskBoardFormat() bool`

HasProgressTaskBoardFormat returns a boolean if a field has been set.

### SetProgressTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) SetProgressTaskBoardFormat(v AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat)`

SetProgressTaskBoardFormat gets a reference to the given AnyOfmicrosoftGraphPlannerProgressTaskBoardTaskFormat and assigns it to the ProgressTaskBoardFormat field.

### SetProgressTaskBoardFormatExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetProgressTaskBoardFormatExplicitNull(b bool)`

SetProgressTaskBoardFormatExplicitNull (un)sets ProgressTaskBoardFormat to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProgressTaskBoardFormat value is set to nil even if false is passed
### GetBucketTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) GetBucketTaskBoardFormat() AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat`

GetBucketTaskBoardFormat returns the BucketTaskBoardFormat field if non-nil, zero value otherwise.

### GetBucketTaskBoardFormatOk

`func (o *MicrosoftGraphPlannerTask) GetBucketTaskBoardFormatOk() (AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat, bool)`

GetBucketTaskBoardFormatOk returns a tuple with the BucketTaskBoardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBucketTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) HasBucketTaskBoardFormat() bool`

HasBucketTaskBoardFormat returns a boolean if a field has been set.

### SetBucketTaskBoardFormat

`func (o *MicrosoftGraphPlannerTask) SetBucketTaskBoardFormat(v AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat)`

SetBucketTaskBoardFormat gets a reference to the given AnyOfmicrosoftGraphPlannerBucketTaskBoardTaskFormat and assigns it to the BucketTaskBoardFormat field.

### SetBucketTaskBoardFormatExplicitNull

`func (o *MicrosoftGraphPlannerTask) SetBucketTaskBoardFormatExplicitNull(b bool)`

SetBucketTaskBoardFormatExplicitNull (un)sets BucketTaskBoardFormat to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BucketTaskBoardFormat value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


