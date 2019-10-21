// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// RoleScopeTagRequestBuilder is request builder for RoleScopeTag
type RoleScopeTagRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoleScopeTagRequest
func (b *RoleScopeTagRequestBuilder) Request() *RoleScopeTagRequest {
	return &RoleScopeTagRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RoleScopeTagRequest is request for RoleScopeTag
type RoleScopeTagRequest struct{ BaseRequest }

// Do performs HTTP request for RoleScopeTag
func (r *RoleScopeTagRequest) Do(method, path string, reqObj interface{}) (resObj *RoleScopeTag, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for RoleScopeTag
func (r *RoleScopeTagRequest) Get() (*RoleScopeTag, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for RoleScopeTag
func (r *RoleScopeTagRequest) Update(reqObj *RoleScopeTag) (*RoleScopeTag, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for RoleScopeTag
func (r *RoleScopeTagRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Assignments returns request builder for RoleScopeTagAutoAssignment collection
func (b *RoleScopeTagRequestBuilder) Assignments() *RoleScopeTagAssignmentsCollectionRequestBuilder {
	bb := &RoleScopeTagAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// RoleScopeTagAssignmentsCollectionRequestBuilder is request builder for RoleScopeTagAutoAssignment collection
type RoleScopeTagAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RoleScopeTagAutoAssignment collection
func (b *RoleScopeTagAssignmentsCollectionRequestBuilder) Request() *RoleScopeTagAssignmentsCollectionRequest {
	return &RoleScopeTagAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RoleScopeTagAutoAssignment item
func (b *RoleScopeTagAssignmentsCollectionRequestBuilder) ID(id string) *RoleScopeTagAutoAssignmentRequestBuilder {
	bb := &RoleScopeTagAutoAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RoleScopeTagAssignmentsCollectionRequest is request for RoleScopeTagAutoAssignment collection
type RoleScopeTagAssignmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for RoleScopeTagAutoAssignment collection
func (r *RoleScopeTagAssignmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *RoleScopeTagAutoAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for RoleScopeTagAutoAssignment collection
func (r *RoleScopeTagAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]RoleScopeTagAutoAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []RoleScopeTagAutoAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []RoleScopeTagAutoAssignment
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

// Get performs GET request for RoleScopeTagAutoAssignment collection
func (r *RoleScopeTagAssignmentsCollectionRequest) Get() ([]RoleScopeTagAutoAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for RoleScopeTagAutoAssignment collection
func (r *RoleScopeTagAssignmentsCollectionRequest) Add(reqObj *RoleScopeTagAutoAssignment) (*RoleScopeTagAutoAssignment, error) {
	return r.Do("POST", "", reqObj)
}
