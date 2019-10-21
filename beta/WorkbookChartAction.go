// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "encoding/json"

// WorkbookChartSetDataRequestParameter undocumented
type WorkbookChartSetDataRequestParameter struct {
	// SourceData undocumented
	SourceData json.RawMessage `json:"sourceData,omitempty"`
	// SeriesBy undocumented
	SeriesBy *string `json:"seriesBy,omitempty"`
}

// WorkbookChartSetPositionRequestParameter undocumented
type WorkbookChartSetPositionRequestParameter struct {
	// StartCell undocumented
	StartCell json.RawMessage `json:"startCell,omitempty"`
	// EndCell undocumented
	EndCell json.RawMessage `json:"endCell,omitempty"`
}

// WorkbookChartCollectionAddRequestParameter undocumented
type WorkbookChartCollectionAddRequestParameter struct {
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// SourceData undocumented
	SourceData json.RawMessage `json:"sourceData,omitempty"`
	// SeriesBy undocumented
	SeriesBy *string `json:"seriesBy,omitempty"`
}

//
type WorkbookChartSetDataRequestBuilder struct{ BaseRequestBuilder }

// SetData action undocumented
func (b *WorkbookChartRequestBuilder) SetData(reqObj *WorkbookChartSetDataRequestParameter) *WorkbookChartSetDataRequestBuilder {
	bb := &WorkbookChartSetDataRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/setData"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookChartSetDataRequest struct{ BaseRequest }

//
func (b *WorkbookChartSetDataRequestBuilder) Request() *WorkbookChartSetDataRequest {
	return &WorkbookChartSetDataRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookChartSetDataRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *WorkbookChartSetDataRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type WorkbookChartSetPositionRequestBuilder struct{ BaseRequestBuilder }

// SetPosition action undocumented
func (b *WorkbookChartRequestBuilder) SetPosition(reqObj *WorkbookChartSetPositionRequestParameter) *WorkbookChartSetPositionRequestBuilder {
	bb := &WorkbookChartSetPositionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/setPosition"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookChartSetPositionRequest struct{ BaseRequest }

//
func (b *WorkbookChartSetPositionRequestBuilder) Request() *WorkbookChartSetPositionRequest {
	return &WorkbookChartSetPositionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookChartSetPositionRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *WorkbookChartSetPositionRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type WorkbookChartCollectionAddRequestBuilder struct{ BaseRequestBuilder }

// Add action undocumented
func (b *WorkbookWorksheetChartsCollectionRequestBuilder) Add(reqObj *WorkbookChartCollectionAddRequestParameter) *WorkbookChartCollectionAddRequestBuilder {
	bb := &WorkbookChartCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookChartCollectionAddRequest struct{ BaseRequest }

//
func (b *WorkbookChartCollectionAddRequestBuilder) Request() *WorkbookChartCollectionAddRequest {
	return &WorkbookChartCollectionAddRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookChartCollectionAddRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookChart, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *WorkbookChartCollectionAddRequest) Post() (*WorkbookChart, error) {
	return r.Do("POST", "", r.requestObject)
}
