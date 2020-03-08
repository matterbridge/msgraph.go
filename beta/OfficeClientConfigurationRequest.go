// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// OfficeClientConfigurationRequestBuilder is request builder for OfficeClientConfiguration
type OfficeClientConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OfficeClientConfigurationRequest
func (b *OfficeClientConfigurationRequestBuilder) Request() *OfficeClientConfigurationRequest {
	return &OfficeClientConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OfficeClientConfigurationRequest is request for OfficeClientConfiguration
type OfficeClientConfigurationRequest struct{ BaseRequest }

// Get performs GET request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Get(ctx context.Context) (resObj *OfficeClientConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Update(ctx context.Context, reqObj *OfficeClientConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OfficeClientConfiguration
func (r *OfficeClientConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for OfficeClientConfigurationAssignment collection
func (b *OfficeClientConfigurationRequestBuilder) Assignments() *OfficeClientConfigurationAssignmentsCollectionRequestBuilder {
	bb := &OfficeClientConfigurationAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// OfficeClientConfigurationAssignmentsCollectionRequestBuilder is request builder for OfficeClientConfigurationAssignment collection
type OfficeClientConfigurationAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OfficeClientConfigurationAssignment collection
func (b *OfficeClientConfigurationAssignmentsCollectionRequestBuilder) Request() *OfficeClientConfigurationAssignmentsCollectionRequest {
	return &OfficeClientConfigurationAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OfficeClientConfigurationAssignment item
func (b *OfficeClientConfigurationAssignmentsCollectionRequestBuilder) ID(id string) *OfficeClientConfigurationAssignmentRequestBuilder {
	bb := &OfficeClientConfigurationAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OfficeClientConfigurationAssignmentsCollectionRequest is request for OfficeClientConfigurationAssignment collection
type OfficeClientConfigurationAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OfficeClientConfigurationAssignment collection
func (r *OfficeClientConfigurationAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OfficeClientConfigurationAssignment, error) {
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
	var values []OfficeClientConfigurationAssignment
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
			value  []OfficeClientConfigurationAssignment
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

// Get performs GET request for OfficeClientConfigurationAssignment collection
func (r *OfficeClientConfigurationAssignmentsCollectionRequest) Get(ctx context.Context) ([]OfficeClientConfigurationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OfficeClientConfigurationAssignment collection
func (r *OfficeClientConfigurationAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *OfficeClientConfigurationAssignment) (resObj *OfficeClientConfigurationAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
