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
// WorkbookChart struct for WorkbookChart
type WorkbookChart struct {
	Height *AnyOfnumberstringstring `json:"height,omitempty"`

	Left *AnyOfnumberstringstring `json:"left,omitempty"`

	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
	Top *AnyOfnumberstringstring `json:"top,omitempty"`

	Width *AnyOfnumberstringstring `json:"width,omitempty"`

	Axes *AnyOfmicrosoftGraphWorkbookChartAxes `json:"axes,omitempty"`
	isExplicitNullAxes bool `json:"-"`
	DataLabels *AnyOfmicrosoftGraphWorkbookChartDataLabels `json:"dataLabels,omitempty"`
	isExplicitNullDataLabels bool `json:"-"`
	Format *AnyOfmicrosoftGraphWorkbookChartAreaFormat `json:"format,omitempty"`
	isExplicitNullFormat bool `json:"-"`
	Legend *AnyOfmicrosoftGraphWorkbookChartLegend `json:"legend,omitempty"`
	isExplicitNullLegend bool `json:"-"`
	Series *[]MicrosoftGraphWorkbookChartSeries `json:"series,omitempty"`

	Title *AnyOfmicrosoftGraphWorkbookChartTitle `json:"title,omitempty"`
	isExplicitNullTitle bool `json:"-"`
	Worksheet *AnyOfmicrosoftGraphWorkbookWorksheet `json:"worksheet,omitempty"`
	isExplicitNullWorksheet bool `json:"-"`
}

// GetHeight returns the Height field if non-nil, zero value otherwise.
func (o *WorkbookChart) GetHeight() AnyOfnumberstringstring {
	if o == nil || o.Height == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetHeightOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Height == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *WorkbookChart) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given AnyOfnumberstringstring and assigns it to the Height field.
func (o *WorkbookChart) SetHeight(v AnyOfnumberstringstring) {
	o.Height = &v
}

// GetLeft returns the Left field if non-nil, zero value otherwise.
func (o *WorkbookChart) GetLeft() AnyOfnumberstringstring {
	if o == nil || o.Left == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Left
}

// GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetLeftOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Left == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Left, true
}

// HasLeft returns a boolean if a field has been set.
func (o *WorkbookChart) HasLeft() bool {
	if o != nil && o.Left != nil {
		return true
	}

	return false
}

// SetLeft gets a reference to the given AnyOfnumberstringstring and assigns it to the Left field.
func (o *WorkbookChart) SetLeft(v AnyOfnumberstringstring) {
	o.Left = &v
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *WorkbookChart) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkbookChart) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkbookChart) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *WorkbookChart) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetTop returns the Top field if non-nil, zero value otherwise.
func (o *WorkbookChart) GetTop() AnyOfnumberstringstring {
	if o == nil || o.Top == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Top
}

// GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetTopOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Top == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Top, true
}

// HasTop returns a boolean if a field has been set.
func (o *WorkbookChart) HasTop() bool {
	if o != nil && o.Top != nil {
		return true
	}

	return false
}

// SetTop gets a reference to the given AnyOfnumberstringstring and assigns it to the Top field.
func (o *WorkbookChart) SetTop(v AnyOfnumberstringstring) {
	o.Top = &v
}

// GetWidth returns the Width field if non-nil, zero value otherwise.
func (o *WorkbookChart) GetWidth() AnyOfnumberstringstring {
	if o == nil || o.Width == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetWidthOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.Width == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *WorkbookChart) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given AnyOfnumberstringstring and assigns it to the Width field.
func (o *WorkbookChart) SetWidth(v AnyOfnumberstringstring) {
	o.Width = &v
}

// GetAxes returns the Axes field if non-nil, zero value otherwise.
func (o *WorkbookChart) GetAxes() AnyOfmicrosoftGraphWorkbookChartAxes {
	if o == nil || o.Axes == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartAxes
		return ret
	}
	return *o.Axes
}

// GetAxesOk returns a tuple with the Axes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetAxesOk() (AnyOfmicrosoftGraphWorkbookChartAxes, bool) {
	if o == nil || o.Axes == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartAxes
		return ret, false
	}
	return *o.Axes, true
}

// HasAxes returns a boolean if a field has been set.
func (o *WorkbookChart) HasAxes() bool {
	if o != nil && o.Axes != nil {
		return true
	}

	return false
}

// SetAxes gets a reference to the given AnyOfmicrosoftGraphWorkbookChartAxes and assigns it to the Axes field.
func (o *WorkbookChart) SetAxes(v AnyOfmicrosoftGraphWorkbookChartAxes) {
	o.Axes = &v
}

// SetAxesExplicitNull (un)sets Axes to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Axes value is set to nil even if false is passed
func (o *WorkbookChart) SetAxesExplicitNull(b bool) {
	o.Axes = nil
	o.isExplicitNullAxes = b
}
// GetDataLabels returns the DataLabels field if non-nil, zero value otherwise.
func (o *WorkbookChart) GetDataLabels() AnyOfmicrosoftGraphWorkbookChartDataLabels {
	if o == nil || o.DataLabels == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartDataLabels
		return ret
	}
	return *o.DataLabels
}

// GetDataLabelsOk returns a tuple with the DataLabels field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetDataLabelsOk() (AnyOfmicrosoftGraphWorkbookChartDataLabels, bool) {
	if o == nil || o.DataLabels == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartDataLabels
		return ret, false
	}
	return *o.DataLabels, true
}

// HasDataLabels returns a boolean if a field has been set.
func (o *WorkbookChart) HasDataLabels() bool {
	if o != nil && o.DataLabels != nil {
		return true
	}

	return false
}

// SetDataLabels gets a reference to the given AnyOfmicrosoftGraphWorkbookChartDataLabels and assigns it to the DataLabels field.
func (o *WorkbookChart) SetDataLabels(v AnyOfmicrosoftGraphWorkbookChartDataLabels) {
	o.DataLabels = &v
}

// SetDataLabelsExplicitNull (un)sets DataLabels to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DataLabels value is set to nil even if false is passed
func (o *WorkbookChart) SetDataLabelsExplicitNull(b bool) {
	o.DataLabels = nil
	o.isExplicitNullDataLabels = b
}
// GetFormat returns the Format field if non-nil, zero value otherwise.
func (o *WorkbookChart) GetFormat() AnyOfmicrosoftGraphWorkbookChartAreaFormat {
	if o == nil || o.Format == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartAreaFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetFormatOk() (AnyOfmicrosoftGraphWorkbookChartAreaFormat, bool) {
	if o == nil || o.Format == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartAreaFormat
		return ret, false
	}
	return *o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *WorkbookChart) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartAreaFormat and assigns it to the Format field.
func (o *WorkbookChart) SetFormat(v AnyOfmicrosoftGraphWorkbookChartAreaFormat) {
	o.Format = &v
}

// SetFormatExplicitNull (un)sets Format to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Format value is set to nil even if false is passed
func (o *WorkbookChart) SetFormatExplicitNull(b bool) {
	o.Format = nil
	o.isExplicitNullFormat = b
}
// GetLegend returns the Legend field if non-nil, zero value otherwise.
func (o *WorkbookChart) GetLegend() AnyOfmicrosoftGraphWorkbookChartLegend {
	if o == nil || o.Legend == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartLegend
		return ret
	}
	return *o.Legend
}

// GetLegendOk returns a tuple with the Legend field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetLegendOk() (AnyOfmicrosoftGraphWorkbookChartLegend, bool) {
	if o == nil || o.Legend == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartLegend
		return ret, false
	}
	return *o.Legend, true
}

// HasLegend returns a boolean if a field has been set.
func (o *WorkbookChart) HasLegend() bool {
	if o != nil && o.Legend != nil {
		return true
	}

	return false
}

// SetLegend gets a reference to the given AnyOfmicrosoftGraphWorkbookChartLegend and assigns it to the Legend field.
func (o *WorkbookChart) SetLegend(v AnyOfmicrosoftGraphWorkbookChartLegend) {
	o.Legend = &v
}

// SetLegendExplicitNull (un)sets Legend to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Legend value is set to nil even if false is passed
func (o *WorkbookChart) SetLegendExplicitNull(b bool) {
	o.Legend = nil
	o.isExplicitNullLegend = b
}
// GetSeries returns the Series field if non-nil, zero value otherwise.
func (o *WorkbookChart) GetSeries() []MicrosoftGraphWorkbookChartSeries {
	if o == nil || o.Series == nil {
		var ret []MicrosoftGraphWorkbookChartSeries
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetSeriesOk() ([]MicrosoftGraphWorkbookChartSeries, bool) {
	if o == nil || o.Series == nil {
		var ret []MicrosoftGraphWorkbookChartSeries
		return ret, false
	}
	return *o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *WorkbookChart) HasSeries() bool {
	if o != nil && o.Series != nil {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []MicrosoftGraphWorkbookChartSeries and assigns it to the Series field.
func (o *WorkbookChart) SetSeries(v []MicrosoftGraphWorkbookChartSeries) {
	o.Series = &v
}

// GetTitle returns the Title field if non-nil, zero value otherwise.
func (o *WorkbookChart) GetTitle() AnyOfmicrosoftGraphWorkbookChartTitle {
	if o == nil || o.Title == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartTitle
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetTitleOk() (AnyOfmicrosoftGraphWorkbookChartTitle, bool) {
	if o == nil || o.Title == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartTitle
		return ret, false
	}
	return *o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkbookChart) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given AnyOfmicrosoftGraphWorkbookChartTitle and assigns it to the Title field.
func (o *WorkbookChart) SetTitle(v AnyOfmicrosoftGraphWorkbookChartTitle) {
	o.Title = &v
}

// SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Title value is set to nil even if false is passed
func (o *WorkbookChart) SetTitleExplicitNull(b bool) {
	o.Title = nil
	o.isExplicitNullTitle = b
}
// GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.
func (o *WorkbookChart) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet {
	if o == nil || o.Worksheet == nil {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret
	}
	return *o.Worksheet
}

// GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChart) GetWorksheetOk() (AnyOfmicrosoftGraphWorkbookWorksheet, bool) {
	if o == nil || o.Worksheet == nil {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret, false
	}
	return *o.Worksheet, true
}

// HasWorksheet returns a boolean if a field has been set.
func (o *WorkbookChart) HasWorksheet() bool {
	if o != nil && o.Worksheet != nil {
		return true
	}

	return false
}

// SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.
func (o *WorkbookChart) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet) {
	o.Worksheet = &v
}

// SetWorksheetExplicitNull (un)sets Worksheet to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Worksheet value is set to nil even if false is passed
func (o *WorkbookChart) SetWorksheetExplicitNull(b bool) {
	o.Worksheet = nil
	o.isExplicitNullWorksheet = b
}

// MarshalJSON returns the JSON representation of the model.
func (o WorkbookChart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Left != nil {
		toSerialize["left"] = o.Left
	}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
	}
	if o.Top != nil {
		toSerialize["top"] = o.Top
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.Axes == nil {
		if o.isExplicitNullAxes {
			toSerialize["axes"] = o.Axes
		}
	} else {
		toSerialize["axes"] = o.Axes
	}
	if o.DataLabels == nil {
		if o.isExplicitNullDataLabels {
			toSerialize["dataLabels"] = o.DataLabels
		}
	} else {
		toSerialize["dataLabels"] = o.DataLabels
	}
	if o.Format == nil {
		if o.isExplicitNullFormat {
			toSerialize["format"] = o.Format
		}
	} else {
		toSerialize["format"] = o.Format
	}
	if o.Legend == nil {
		if o.isExplicitNullLegend {
			toSerialize["legend"] = o.Legend
		}
	} else {
		toSerialize["legend"] = o.Legend
	}
	if o.Series != nil {
		toSerialize["series"] = o.Series
	}
	if o.Title == nil {
		if o.isExplicitNullTitle {
			toSerialize["title"] = o.Title
		}
	} else {
		toSerialize["title"] = o.Title
	}
	if o.Worksheet == nil {
		if o.isExplicitNullWorksheet {
			toSerialize["worksheet"] = o.Worksheet
		}
	} else {
		toSerialize["worksheet"] = o.Worksheet
	}
	return json.Marshal(toSerialize)
}


