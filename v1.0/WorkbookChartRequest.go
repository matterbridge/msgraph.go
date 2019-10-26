// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// WorkbookChartRequestBuilder is request builder for WorkbookChart
type WorkbookChartRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookChartRequest
func (b *WorkbookChartRequestBuilder) Request() *WorkbookChartRequest {
	return &WorkbookChartRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookChartRequest is request for WorkbookChart
type WorkbookChartRequest struct{ BaseRequest }

// Get performs GET request for WorkbookChart
func (r *WorkbookChartRequest) Get() (resObj *WorkbookChart, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookChart
func (r *WorkbookChartRequest) Update(reqObj *WorkbookChart) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookChart
func (r *WorkbookChartRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Axes is navigation property
func (b *WorkbookChartRequestBuilder) Axes() *WorkbookChartAxesRequestBuilder {
	bb := &WorkbookChartAxesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/axes"
	return bb
}

// DataLabels is navigation property
func (b *WorkbookChartRequestBuilder) DataLabels() *WorkbookChartDataLabelsRequestBuilder {
	bb := &WorkbookChartDataLabelsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/dataLabels"
	return bb
}

// Format is navigation property
func (b *WorkbookChartRequestBuilder) Format() *WorkbookChartAreaFormatRequestBuilder {
	bb := &WorkbookChartAreaFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/format"
	return bb
}

// Legend is navigation property
func (b *WorkbookChartRequestBuilder) Legend() *WorkbookChartLegendRequestBuilder {
	bb := &WorkbookChartLegendRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/legend"
	return bb
}

// Series returns request builder for WorkbookChartSeries collection
func (b *WorkbookChartRequestBuilder) Series() *WorkbookChartSeriesCollectionRequestBuilder {
	bb := &WorkbookChartSeriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/series"
	return bb
}

// WorkbookChartSeriesCollectionRequestBuilder is request builder for WorkbookChartSeries collection
type WorkbookChartSeriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookChartSeries collection
func (b *WorkbookChartSeriesCollectionRequestBuilder) Request() *WorkbookChartSeriesCollectionRequest {
	return &WorkbookChartSeriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookChartSeries item
func (b *WorkbookChartSeriesCollectionRequestBuilder) ID(id string) *WorkbookChartSeriesRequestBuilder {
	bb := &WorkbookChartSeriesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookChartSeriesCollectionRequest is request for WorkbookChartSeries collection
type WorkbookChartSeriesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WorkbookChartSeries collection
func (r *WorkbookChartSeriesCollectionRequest) Paging(method, path string, obj interface{}) ([]WorkbookChartSeries, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WorkbookChartSeries
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []WorkbookChartSeries
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for WorkbookChartSeries collection
func (r *WorkbookChartSeriesCollectionRequest) Get() ([]WorkbookChartSeries, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WorkbookChartSeries collection
func (r *WorkbookChartSeriesCollectionRequest) Add(reqObj *WorkbookChartSeries) (resObj *WorkbookChartSeries, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Title is navigation property
func (b *WorkbookChartRequestBuilder) Title() *WorkbookChartTitleRequestBuilder {
	bb := &WorkbookChartTitleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/title"
	return bb
}

// Worksheet is navigation property
func (b *WorkbookChartRequestBuilder) Worksheet() *WorkbookWorksheetRequestBuilder {
	bb := &WorkbookWorksheetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/worksheet"
	return bb
}
