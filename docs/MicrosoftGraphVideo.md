# MicrosoftGraphVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioBitsPerSample** | Pointer to **int32** |  | [optional] 
**AudioChannels** | Pointer to **int32** |  | [optional] 
**AudioFormat** | Pointer to **string** |  | [optional] 
**AudioSamplesPerSecond** | Pointer to **int32** |  | [optional] 
**Bitrate** | Pointer to **int32** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**FourCC** | Pointer to **string** |  | [optional] 
**FrameRate** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Width** | Pointer to **int32** |  | [optional] 

## Methods

### GetAudioBitsPerSample

`func (o *MicrosoftGraphVideo) GetAudioBitsPerSample() int32`

GetAudioBitsPerSample returns the AudioBitsPerSample field if non-nil, zero value otherwise.

### GetAudioBitsPerSampleOk

`func (o *MicrosoftGraphVideo) GetAudioBitsPerSampleOk() (int32, bool)`

GetAudioBitsPerSampleOk returns a tuple with the AudioBitsPerSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudioBitsPerSample

`func (o *MicrosoftGraphVideo) HasAudioBitsPerSample() bool`

HasAudioBitsPerSample returns a boolean if a field has been set.

### SetAudioBitsPerSample

`func (o *MicrosoftGraphVideo) SetAudioBitsPerSample(v int32)`

SetAudioBitsPerSample gets a reference to the given int32 and assigns it to the AudioBitsPerSample field.

### SetAudioBitsPerSampleExplicitNull

`func (o *MicrosoftGraphVideo) SetAudioBitsPerSampleExplicitNull(b bool)`

SetAudioBitsPerSampleExplicitNull (un)sets AudioBitsPerSample to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AudioBitsPerSample value is set to nil even if false is passed
### GetAudioChannels

`func (o *MicrosoftGraphVideo) GetAudioChannels() int32`

GetAudioChannels returns the AudioChannels field if non-nil, zero value otherwise.

### GetAudioChannelsOk

`func (o *MicrosoftGraphVideo) GetAudioChannelsOk() (int32, bool)`

GetAudioChannelsOk returns a tuple with the AudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudioChannels

`func (o *MicrosoftGraphVideo) HasAudioChannels() bool`

HasAudioChannels returns a boolean if a field has been set.

### SetAudioChannels

`func (o *MicrosoftGraphVideo) SetAudioChannels(v int32)`

SetAudioChannels gets a reference to the given int32 and assigns it to the AudioChannels field.

### SetAudioChannelsExplicitNull

`func (o *MicrosoftGraphVideo) SetAudioChannelsExplicitNull(b bool)`

SetAudioChannelsExplicitNull (un)sets AudioChannels to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AudioChannels value is set to nil even if false is passed
### GetAudioFormat

`func (o *MicrosoftGraphVideo) GetAudioFormat() string`

GetAudioFormat returns the AudioFormat field if non-nil, zero value otherwise.

### GetAudioFormatOk

`func (o *MicrosoftGraphVideo) GetAudioFormatOk() (string, bool)`

GetAudioFormatOk returns a tuple with the AudioFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudioFormat

`func (o *MicrosoftGraphVideo) HasAudioFormat() bool`

HasAudioFormat returns a boolean if a field has been set.

### SetAudioFormat

`func (o *MicrosoftGraphVideo) SetAudioFormat(v string)`

SetAudioFormat gets a reference to the given string and assigns it to the AudioFormat field.

### SetAudioFormatExplicitNull

`func (o *MicrosoftGraphVideo) SetAudioFormatExplicitNull(b bool)`

SetAudioFormatExplicitNull (un)sets AudioFormat to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AudioFormat value is set to nil even if false is passed
### GetAudioSamplesPerSecond

`func (o *MicrosoftGraphVideo) GetAudioSamplesPerSecond() int32`

GetAudioSamplesPerSecond returns the AudioSamplesPerSecond field if non-nil, zero value otherwise.

### GetAudioSamplesPerSecondOk

`func (o *MicrosoftGraphVideo) GetAudioSamplesPerSecondOk() (int32, bool)`

GetAudioSamplesPerSecondOk returns a tuple with the AudioSamplesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudioSamplesPerSecond

`func (o *MicrosoftGraphVideo) HasAudioSamplesPerSecond() bool`

HasAudioSamplesPerSecond returns a boolean if a field has been set.

### SetAudioSamplesPerSecond

`func (o *MicrosoftGraphVideo) SetAudioSamplesPerSecond(v int32)`

SetAudioSamplesPerSecond gets a reference to the given int32 and assigns it to the AudioSamplesPerSecond field.

### SetAudioSamplesPerSecondExplicitNull

`func (o *MicrosoftGraphVideo) SetAudioSamplesPerSecondExplicitNull(b bool)`

SetAudioSamplesPerSecondExplicitNull (un)sets AudioSamplesPerSecond to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AudioSamplesPerSecond value is set to nil even if false is passed
### GetBitrate

`func (o *MicrosoftGraphVideo) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *MicrosoftGraphVideo) GetBitrateOk() (int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBitrate

`func (o *MicrosoftGraphVideo) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrate

`func (o *MicrosoftGraphVideo) SetBitrate(v int32)`

SetBitrate gets a reference to the given int32 and assigns it to the Bitrate field.

### SetBitrateExplicitNull

`func (o *MicrosoftGraphVideo) SetBitrateExplicitNull(b bool)`

SetBitrateExplicitNull (un)sets Bitrate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Bitrate value is set to nil even if false is passed
### GetDuration

`func (o *MicrosoftGraphVideo) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MicrosoftGraphVideo) GetDurationOk() (int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDuration

`func (o *MicrosoftGraphVideo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDuration

`func (o *MicrosoftGraphVideo) SetDuration(v int64)`

SetDuration gets a reference to the given int64 and assigns it to the Duration field.

### SetDurationExplicitNull

`func (o *MicrosoftGraphVideo) SetDurationExplicitNull(b bool)`

SetDurationExplicitNull (un)sets Duration to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Duration value is set to nil even if false is passed
### GetFourCC

`func (o *MicrosoftGraphVideo) GetFourCC() string`

GetFourCC returns the FourCC field if non-nil, zero value otherwise.

### GetFourCCOk

`func (o *MicrosoftGraphVideo) GetFourCCOk() (string, bool)`

GetFourCCOk returns a tuple with the FourCC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFourCC

`func (o *MicrosoftGraphVideo) HasFourCC() bool`

HasFourCC returns a boolean if a field has been set.

### SetFourCC

`func (o *MicrosoftGraphVideo) SetFourCC(v string)`

SetFourCC gets a reference to the given string and assigns it to the FourCC field.

### SetFourCCExplicitNull

`func (o *MicrosoftGraphVideo) SetFourCCExplicitNull(b bool)`

SetFourCCExplicitNull (un)sets FourCC to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FourCC value is set to nil even if false is passed
### GetFrameRate

`func (o *MicrosoftGraphVideo) GetFrameRate() AnyOfnumberstringstring`

GetFrameRate returns the FrameRate field if non-nil, zero value otherwise.

### GetFrameRateOk

`func (o *MicrosoftGraphVideo) GetFrameRateOk() (AnyOfnumberstringstring, bool)`

GetFrameRateOk returns a tuple with the FrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFrameRate

`func (o *MicrosoftGraphVideo) HasFrameRate() bool`

HasFrameRate returns a boolean if a field has been set.

### SetFrameRate

`func (o *MicrosoftGraphVideo) SetFrameRate(v AnyOfnumberstringstring)`

SetFrameRate gets a reference to the given AnyOfnumberstringstring and assigns it to the FrameRate field.

### SetFrameRateExplicitNull

`func (o *MicrosoftGraphVideo) SetFrameRateExplicitNull(b bool)`

SetFrameRateExplicitNull (un)sets FrameRate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FrameRate value is set to nil even if false is passed
### GetHeight

`func (o *MicrosoftGraphVideo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *MicrosoftGraphVideo) GetHeightOk() (int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeight

`func (o *MicrosoftGraphVideo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeight

`func (o *MicrosoftGraphVideo) SetHeight(v int32)`

SetHeight gets a reference to the given int32 and assigns it to the Height field.

### SetHeightExplicitNull

`func (o *MicrosoftGraphVideo) SetHeightExplicitNull(b bool)`

SetHeightExplicitNull (un)sets Height to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Height value is set to nil even if false is passed
### GetWidth

`func (o *MicrosoftGraphVideo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MicrosoftGraphVideo) GetWidthOk() (int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWidth

`func (o *MicrosoftGraphVideo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidth

`func (o *MicrosoftGraphVideo) SetWidth(v int32)`

SetWidth gets a reference to the given int32 and assigns it to the Width field.

### SetWidthExplicitNull

`func (o *MicrosoftGraphVideo) SetWidthExplicitNull(b bool)`

SetWidthExplicitNull (un)sets Width to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Width value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


