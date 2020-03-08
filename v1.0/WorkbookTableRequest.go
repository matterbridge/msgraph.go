// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// WorkbookTableRequestBuilder is request builder for WorkbookTable
type WorkbookTableRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookTableRequest
func (b *WorkbookTableRequestBuilder) Request() *WorkbookTableRequest {
	return &WorkbookTableRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookTableRequest is request for WorkbookTable
type WorkbookTableRequest struct{ BaseRequest }

// Get performs GET request for WorkbookTable
func (r *WorkbookTableRequest) Get(ctx context.Context) (resObj *WorkbookTable, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookTable
func (r *WorkbookTableRequest) Update(ctx context.Context, reqObj *WorkbookTable) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookTable
func (r *WorkbookTableRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Columns returns request builder for WorkbookTableColumn collection
func (b *WorkbookTableRequestBuilder) Columns() *WorkbookTableColumnsCollectionRequestBuilder {
	bb := &WorkbookTableColumnsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/columns"
	return bb
}

// WorkbookTableColumnsCollectionRequestBuilder is request builder for WorkbookTableColumn collection
type WorkbookTableColumnsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookTableColumn collection
func (b *WorkbookTableColumnsCollectionRequestBuilder) Request() *WorkbookTableColumnsCollectionRequest {
	return &WorkbookTableColumnsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookTableColumn item
func (b *WorkbookTableColumnsCollectionRequestBuilder) ID(id string) *WorkbookTableColumnRequestBuilder {
	bb := &WorkbookTableColumnRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookTableColumnsCollectionRequest is request for WorkbookTableColumn collection
type WorkbookTableColumnsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WorkbookTableColumn collection
func (r *WorkbookTableColumnsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]WorkbookTableColumn, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WorkbookTableColumn
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
			value  []WorkbookTableColumn
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
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for WorkbookTableColumn collection
func (r *WorkbookTableColumnsCollectionRequest) Get(ctx context.Context) ([]WorkbookTableColumn, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for WorkbookTableColumn collection
func (r *WorkbookTableColumnsCollectionRequest) Add(ctx context.Context, reqObj *WorkbookTableColumn) (resObj *WorkbookTableColumn, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Rows returns request builder for WorkbookTableRow collection
func (b *WorkbookTableRequestBuilder) Rows() *WorkbookTableRowsCollectionRequestBuilder {
	bb := &WorkbookTableRowsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rows"
	return bb
}

// WorkbookTableRowsCollectionRequestBuilder is request builder for WorkbookTableRow collection
type WorkbookTableRowsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkbookTableRow collection
func (b *WorkbookTableRowsCollectionRequestBuilder) Request() *WorkbookTableRowsCollectionRequest {
	return &WorkbookTableRowsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkbookTableRow item
func (b *WorkbookTableRowsCollectionRequestBuilder) ID(id string) *WorkbookTableRowRequestBuilder {
	bb := &WorkbookTableRowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WorkbookTableRowsCollectionRequest is request for WorkbookTableRow collection
type WorkbookTableRowsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WorkbookTableRow collection
func (r *WorkbookTableRowsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]WorkbookTableRow, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WorkbookTableRow
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
			value  []WorkbookTableRow
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
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for WorkbookTableRow collection
func (r *WorkbookTableRowsCollectionRequest) Get(ctx context.Context) ([]WorkbookTableRow, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for WorkbookTableRow collection
func (r *WorkbookTableRowsCollectionRequest) Add(ctx context.Context, reqObj *WorkbookTableRow) (resObj *WorkbookTableRow, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Sort is navigation property
func (b *WorkbookTableRequestBuilder) Sort() *WorkbookTableSortRequestBuilder {
	bb := &WorkbookTableSortRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sort"
	return bb
}

// Worksheet is navigation property
func (b *WorkbookTableRequestBuilder) Worksheet() *WorkbookWorksheetRequestBuilder {
	bb := &WorkbookWorksheetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/worksheet"
	return bb
}
