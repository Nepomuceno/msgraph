# MicrosoftGraphAudio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Album** | Pointer to **string** |  | [optional] 
**AlbumArtist** | Pointer to **string** |  | [optional] 
**Artist** | Pointer to **string** |  | [optional] 
**Bitrate** | Pointer to **int64** |  | [optional] 
**Composers** | Pointer to **string** |  | [optional] 
**Copyright** | Pointer to **string** |  | [optional] 
**Disc** | Pointer to **int32** |  | [optional] 
**DiscCount** | Pointer to **int32** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**Genre** | Pointer to **string** |  | [optional] 
**HasDrm** | Pointer to **bool** |  | [optional] 
**IsVariableBitrate** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Track** | Pointer to **int32** |  | [optional] 
**TrackCount** | Pointer to **int32** |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 

## Methods

### GetAlbum

`func (o *MicrosoftGraphAudio) GetAlbum() string`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *MicrosoftGraphAudio) GetAlbumOk() (string, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlbum

`func (o *MicrosoftGraphAudio) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### SetAlbum

`func (o *MicrosoftGraphAudio) SetAlbum(v string)`

SetAlbum gets a reference to the given string and assigns it to the Album field.

### SetAlbumExplicitNull

`func (o *MicrosoftGraphAudio) SetAlbumExplicitNull(b bool)`

SetAlbumExplicitNull (un)sets Album to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Album value is set to nil even if false is passed
### GetAlbumArtist

`func (o *MicrosoftGraphAudio) GetAlbumArtist() string`

GetAlbumArtist returns the AlbumArtist field if non-nil, zero value otherwise.

### GetAlbumArtistOk

`func (o *MicrosoftGraphAudio) GetAlbumArtistOk() (string, bool)`

GetAlbumArtistOk returns a tuple with the AlbumArtist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlbumArtist

`func (o *MicrosoftGraphAudio) HasAlbumArtist() bool`

HasAlbumArtist returns a boolean if a field has been set.

### SetAlbumArtist

`func (o *MicrosoftGraphAudio) SetAlbumArtist(v string)`

SetAlbumArtist gets a reference to the given string and assigns it to the AlbumArtist field.

### SetAlbumArtistExplicitNull

`func (o *MicrosoftGraphAudio) SetAlbumArtistExplicitNull(b bool)`

SetAlbumArtistExplicitNull (un)sets AlbumArtist to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AlbumArtist value is set to nil even if false is passed
### GetArtist

`func (o *MicrosoftGraphAudio) GetArtist() string`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *MicrosoftGraphAudio) GetArtistOk() (string, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasArtist

`func (o *MicrosoftGraphAudio) HasArtist() bool`

HasArtist returns a boolean if a field has been set.

### SetArtist

`func (o *MicrosoftGraphAudio) SetArtist(v string)`

SetArtist gets a reference to the given string and assigns it to the Artist field.

### SetArtistExplicitNull

`func (o *MicrosoftGraphAudio) SetArtistExplicitNull(b bool)`

SetArtistExplicitNull (un)sets Artist to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Artist value is set to nil even if false is passed
### GetBitrate

`func (o *MicrosoftGraphAudio) GetBitrate() int64`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *MicrosoftGraphAudio) GetBitrateOk() (int64, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBitrate

`func (o *MicrosoftGraphAudio) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrate

`func (o *MicrosoftGraphAudio) SetBitrate(v int64)`

SetBitrate gets a reference to the given int64 and assigns it to the Bitrate field.

### SetBitrateExplicitNull

`func (o *MicrosoftGraphAudio) SetBitrateExplicitNull(b bool)`

SetBitrateExplicitNull (un)sets Bitrate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Bitrate value is set to nil even if false is passed
### GetComposers

`func (o *MicrosoftGraphAudio) GetComposers() string`

GetComposers returns the Composers field if non-nil, zero value otherwise.

### GetComposersOk

`func (o *MicrosoftGraphAudio) GetComposersOk() (string, bool)`

GetComposersOk returns a tuple with the Composers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComposers

`func (o *MicrosoftGraphAudio) HasComposers() bool`

HasComposers returns a boolean if a field has been set.

### SetComposers

`func (o *MicrosoftGraphAudio) SetComposers(v string)`

SetComposers gets a reference to the given string and assigns it to the Composers field.

### SetComposersExplicitNull

`func (o *MicrosoftGraphAudio) SetComposersExplicitNull(b bool)`

SetComposersExplicitNull (un)sets Composers to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Composers value is set to nil even if false is passed
### GetCopyright

`func (o *MicrosoftGraphAudio) GetCopyright() string`

GetCopyright returns the Copyright field if non-nil, zero value otherwise.

### GetCopyrightOk

`func (o *MicrosoftGraphAudio) GetCopyrightOk() (string, bool)`

GetCopyrightOk returns a tuple with the Copyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCopyright

`func (o *MicrosoftGraphAudio) HasCopyright() bool`

HasCopyright returns a boolean if a field has been set.

### SetCopyright

`func (o *MicrosoftGraphAudio) SetCopyright(v string)`

SetCopyright gets a reference to the given string and assigns it to the Copyright field.

### SetCopyrightExplicitNull

`func (o *MicrosoftGraphAudio) SetCopyrightExplicitNull(b bool)`

SetCopyrightExplicitNull (un)sets Copyright to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Copyright value is set to nil even if false is passed
### GetDisc

`func (o *MicrosoftGraphAudio) GetDisc() int32`

GetDisc returns the Disc field if non-nil, zero value otherwise.

### GetDiscOk

`func (o *MicrosoftGraphAudio) GetDiscOk() (int32, bool)`

GetDiscOk returns a tuple with the Disc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisc

`func (o *MicrosoftGraphAudio) HasDisc() bool`

HasDisc returns a boolean if a field has been set.

### SetDisc

`func (o *MicrosoftGraphAudio) SetDisc(v int32)`

SetDisc gets a reference to the given int32 and assigns it to the Disc field.

### SetDiscExplicitNull

`func (o *MicrosoftGraphAudio) SetDiscExplicitNull(b bool)`

SetDiscExplicitNull (un)sets Disc to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Disc value is set to nil even if false is passed
### GetDiscCount

`func (o *MicrosoftGraphAudio) GetDiscCount() int32`

GetDiscCount returns the DiscCount field if non-nil, zero value otherwise.

### GetDiscCountOk

`func (o *MicrosoftGraphAudio) GetDiscCountOk() (int32, bool)`

GetDiscCountOk returns a tuple with the DiscCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscCount

`func (o *MicrosoftGraphAudio) HasDiscCount() bool`

HasDiscCount returns a boolean if a field has been set.

### SetDiscCount

`func (o *MicrosoftGraphAudio) SetDiscCount(v int32)`

SetDiscCount gets a reference to the given int32 and assigns it to the DiscCount field.

### SetDiscCountExplicitNull

`func (o *MicrosoftGraphAudio) SetDiscCountExplicitNull(b bool)`

SetDiscCountExplicitNull (un)sets DiscCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DiscCount value is set to nil even if false is passed
### GetDuration

`func (o *MicrosoftGraphAudio) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MicrosoftGraphAudio) GetDurationOk() (int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDuration

`func (o *MicrosoftGraphAudio) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDuration

`func (o *MicrosoftGraphAudio) SetDuration(v int64)`

SetDuration gets a reference to the given int64 and assigns it to the Duration field.

### SetDurationExplicitNull

`func (o *MicrosoftGraphAudio) SetDurationExplicitNull(b bool)`

SetDurationExplicitNull (un)sets Duration to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Duration value is set to nil even if false is passed
### GetGenre

`func (o *MicrosoftGraphAudio) GetGenre() string`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *MicrosoftGraphAudio) GetGenreOk() (string, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGenre

`func (o *MicrosoftGraphAudio) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### SetGenre

`func (o *MicrosoftGraphAudio) SetGenre(v string)`

SetGenre gets a reference to the given string and assigns it to the Genre field.

### SetGenreExplicitNull

`func (o *MicrosoftGraphAudio) SetGenreExplicitNull(b bool)`

SetGenreExplicitNull (un)sets Genre to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Genre value is set to nil even if false is passed
### GetHasDrm

`func (o *MicrosoftGraphAudio) GetHasDrm() bool`

GetHasDrm returns the HasDrm field if non-nil, zero value otherwise.

### GetHasDrmOk

`func (o *MicrosoftGraphAudio) GetHasDrmOk() (bool, bool)`

GetHasDrmOk returns a tuple with the HasDrm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasDrm

`func (o *MicrosoftGraphAudio) HasHasDrm() bool`

HasHasDrm returns a boolean if a field has been set.

### SetHasDrm

`func (o *MicrosoftGraphAudio) SetHasDrm(v bool)`

SetHasDrm gets a reference to the given bool and assigns it to the HasDrm field.

### SetHasDrmExplicitNull

`func (o *MicrosoftGraphAudio) SetHasDrmExplicitNull(b bool)`

SetHasDrmExplicitNull (un)sets HasDrm to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HasDrm value is set to nil even if false is passed
### GetIsVariableBitrate

`func (o *MicrosoftGraphAudio) GetIsVariableBitrate() bool`

GetIsVariableBitrate returns the IsVariableBitrate field if non-nil, zero value otherwise.

### GetIsVariableBitrateOk

`func (o *MicrosoftGraphAudio) GetIsVariableBitrateOk() (bool, bool)`

GetIsVariableBitrateOk returns a tuple with the IsVariableBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsVariableBitrate

`func (o *MicrosoftGraphAudio) HasIsVariableBitrate() bool`

HasIsVariableBitrate returns a boolean if a field has been set.

### SetIsVariableBitrate

`func (o *MicrosoftGraphAudio) SetIsVariableBitrate(v bool)`

SetIsVariableBitrate gets a reference to the given bool and assigns it to the IsVariableBitrate field.

### SetIsVariableBitrateExplicitNull

`func (o *MicrosoftGraphAudio) SetIsVariableBitrateExplicitNull(b bool)`

SetIsVariableBitrateExplicitNull (un)sets IsVariableBitrate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsVariableBitrate value is set to nil even if false is passed
### GetTitle

`func (o *MicrosoftGraphAudio) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphAudio) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *MicrosoftGraphAudio) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *MicrosoftGraphAudio) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### SetTitleExplicitNull

`func (o *MicrosoftGraphAudio) SetTitleExplicitNull(b bool)`

SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Title value is set to nil even if false is passed
### GetTrack

`func (o *MicrosoftGraphAudio) GetTrack() int32`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *MicrosoftGraphAudio) GetTrackOk() (int32, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTrack

`func (o *MicrosoftGraphAudio) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### SetTrack

`func (o *MicrosoftGraphAudio) SetTrack(v int32)`

SetTrack gets a reference to the given int32 and assigns it to the Track field.

### SetTrackExplicitNull

`func (o *MicrosoftGraphAudio) SetTrackExplicitNull(b bool)`

SetTrackExplicitNull (un)sets Track to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Track value is set to nil even if false is passed
### GetTrackCount

`func (o *MicrosoftGraphAudio) GetTrackCount() int32`

GetTrackCount returns the TrackCount field if non-nil, zero value otherwise.

### GetTrackCountOk

`func (o *MicrosoftGraphAudio) GetTrackCountOk() (int32, bool)`

GetTrackCountOk returns a tuple with the TrackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTrackCount

`func (o *MicrosoftGraphAudio) HasTrackCount() bool`

HasTrackCount returns a boolean if a field has been set.

### SetTrackCount

`func (o *MicrosoftGraphAudio) SetTrackCount(v int32)`

SetTrackCount gets a reference to the given int32 and assigns it to the TrackCount field.

### SetTrackCountExplicitNull

`func (o *MicrosoftGraphAudio) SetTrackCountExplicitNull(b bool)`

SetTrackCountExplicitNull (un)sets TrackCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TrackCount value is set to nil even if false is passed
### GetYear

`func (o *MicrosoftGraphAudio) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MicrosoftGraphAudio) GetYearOk() (int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasYear

`func (o *MicrosoftGraphAudio) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYear

`func (o *MicrosoftGraphAudio) SetYear(v int32)`

SetYear gets a reference to the given int32 and assigns it to the Year field.

### SetYearExplicitNull

`func (o *MicrosoftGraphAudio) SetYearExplicitNull(b bool)`

SetYearExplicitNull (un)sets Year to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Year value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


