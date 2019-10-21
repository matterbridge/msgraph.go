// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookTableCollectionAddRequestParameter undocumented
type WorkbookTableCollectionAddRequestParameter struct {
	// Address undocumented
	Address *string `json:"address,omitempty"`
	// HasHeaders undocumented
	HasHeaders *bool `json:"hasHeaders,omitempty"`
}

// WorkbookTableClearFiltersRequestParameter undocumented
type WorkbookTableClearFiltersRequestParameter struct {
}

// WorkbookTableConvertToRangeRequestParameter undocumented
type WorkbookTableConvertToRangeRequestParameter struct {
}

// WorkbookTableReapplyFiltersRequestParameter undocumented
type WorkbookTableReapplyFiltersRequestParameter struct {
}

//
type WorkbookTableCollectionAddRequestBuilder struct{ BaseRequestBuilder }

// Add action undocumented
func (b *WorkbookTablesCollectionRequestBuilder) Add(reqObj *WorkbookTableCollectionAddRequestParameter) *WorkbookTableCollectionAddRequestBuilder {
	bb := &WorkbookTableCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Add action undocumented
func (b *WorkbookWorksheetTablesCollectionRequestBuilder) Add(reqObj *WorkbookTableCollectionAddRequestParameter) *WorkbookTableCollectionAddRequestBuilder {
	bb := &WorkbookTableCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookTableCollectionAddRequest struct{ BaseRequest }

//
func (b *WorkbookTableCollectionAddRequestBuilder) Request() *WorkbookTableCollectionAddRequest {
	return &WorkbookTableCollectionAddRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookTableCollectionAddRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookTable, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *WorkbookTableCollectionAddRequest) Post() (*WorkbookTable, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type WorkbookTableClearFiltersRequestBuilder struct{ BaseRequestBuilder }

// ClearFilters action undocumented
func (b *WorkbookTableRequestBuilder) ClearFilters(reqObj *WorkbookTableClearFiltersRequestParameter) *WorkbookTableClearFiltersRequestBuilder {
	bb := &WorkbookTableClearFiltersRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/clearFilters"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookTableClearFiltersRequest struct{ BaseRequest }

//
func (b *WorkbookTableClearFiltersRequestBuilder) Request() *WorkbookTableClearFiltersRequest {
	return &WorkbookTableClearFiltersRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookTableClearFiltersRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *WorkbookTableClearFiltersRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type WorkbookTableConvertToRangeRequestBuilder struct{ BaseRequestBuilder }

// ConvertToRange action undocumented
func (b *WorkbookTableRequestBuilder) ConvertToRange(reqObj *WorkbookTableConvertToRangeRequestParameter) *WorkbookTableConvertToRangeRequestBuilder {
	bb := &WorkbookTableConvertToRangeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/convertToRange"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookTableConvertToRangeRequest struct{ BaseRequest }

//
func (b *WorkbookTableConvertToRangeRequestBuilder) Request() *WorkbookTableConvertToRangeRequest {
	return &WorkbookTableConvertToRangeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookTableConvertToRangeRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookRange, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

//
func (r *WorkbookTableConvertToRangeRequest) Post() (*WorkbookRange, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type WorkbookTableReapplyFiltersRequestBuilder struct{ BaseRequestBuilder }

// ReapplyFilters action undocumented
func (b *WorkbookTableRequestBuilder) ReapplyFilters(reqObj *WorkbookTableReapplyFiltersRequestParameter) *WorkbookTableReapplyFiltersRequestBuilder {
	bb := &WorkbookTableReapplyFiltersRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/reapplyFilters"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookTableReapplyFiltersRequest struct{ BaseRequest }

//
func (b *WorkbookTableReapplyFiltersRequestBuilder) Request() *WorkbookTableReapplyFiltersRequest {
	return &WorkbookTableReapplyFiltersRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookTableReapplyFiltersRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *WorkbookTableReapplyFiltersRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}
