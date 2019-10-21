// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// TermsAndConditionsRequestBuilder is request builder for TermsAndConditions
type TermsAndConditionsRequestBuilder struct{ BaseRequestBuilder }

// Request returns TermsAndConditionsRequest
func (b *TermsAndConditionsRequestBuilder) Request() *TermsAndConditionsRequest {
	return &TermsAndConditionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TermsAndConditionsRequest is request for TermsAndConditions
type TermsAndConditionsRequest struct{ BaseRequest }

// Do performs HTTP request for TermsAndConditions
func (r *TermsAndConditionsRequest) Do(method, path string, reqObj interface{}) (resObj *TermsAndConditions, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for TermsAndConditions
func (r *TermsAndConditionsRequest) Get() (*TermsAndConditions, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for TermsAndConditions
func (r *TermsAndConditionsRequest) Update(reqObj *TermsAndConditions) (*TermsAndConditions, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for TermsAndConditions
func (r *TermsAndConditionsRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AcceptanceStatuses returns request builder for TermsAndConditionsAcceptanceStatus collection
func (b *TermsAndConditionsRequestBuilder) AcceptanceStatuses() *TermsAndConditionsAcceptanceStatusesCollectionRequestBuilder {
	bb := &TermsAndConditionsAcceptanceStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/acceptanceStatuses"
	return bb
}

// TermsAndConditionsAcceptanceStatusesCollectionRequestBuilder is request builder for TermsAndConditionsAcceptanceStatus collection
type TermsAndConditionsAcceptanceStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TermsAndConditionsAcceptanceStatus collection
func (b *TermsAndConditionsAcceptanceStatusesCollectionRequestBuilder) Request() *TermsAndConditionsAcceptanceStatusesCollectionRequest {
	return &TermsAndConditionsAcceptanceStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TermsAndConditionsAcceptanceStatus item
func (b *TermsAndConditionsAcceptanceStatusesCollectionRequestBuilder) ID(id string) *TermsAndConditionsAcceptanceStatusRequestBuilder {
	bb := &TermsAndConditionsAcceptanceStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TermsAndConditionsAcceptanceStatusesCollectionRequest is request for TermsAndConditionsAcceptanceStatus collection
type TermsAndConditionsAcceptanceStatusesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for TermsAndConditionsAcceptanceStatus collection
func (r *TermsAndConditionsAcceptanceStatusesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *TermsAndConditionsAcceptanceStatus, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for TermsAndConditionsAcceptanceStatus collection
func (r *TermsAndConditionsAcceptanceStatusesCollectionRequest) Paging(method, path string, obj interface{}) ([]TermsAndConditionsAcceptanceStatus, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TermsAndConditionsAcceptanceStatus
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []TermsAndConditionsAcceptanceStatus
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

// Get performs GET request for TermsAndConditionsAcceptanceStatus collection
func (r *TermsAndConditionsAcceptanceStatusesCollectionRequest) Get() ([]TermsAndConditionsAcceptanceStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for TermsAndConditionsAcceptanceStatus collection
func (r *TermsAndConditionsAcceptanceStatusesCollectionRequest) Add(reqObj *TermsAndConditionsAcceptanceStatus) (*TermsAndConditionsAcceptanceStatus, error) {
	return r.Do("POST", "", reqObj)
}

// Assignments returns request builder for TermsAndConditionsAssignment collection
func (b *TermsAndConditionsRequestBuilder) Assignments() *TermsAndConditionsAssignmentsCollectionRequestBuilder {
	bb := &TermsAndConditionsAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// TermsAndConditionsAssignmentsCollectionRequestBuilder is request builder for TermsAndConditionsAssignment collection
type TermsAndConditionsAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TermsAndConditionsAssignment collection
func (b *TermsAndConditionsAssignmentsCollectionRequestBuilder) Request() *TermsAndConditionsAssignmentsCollectionRequest {
	return &TermsAndConditionsAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TermsAndConditionsAssignment item
func (b *TermsAndConditionsAssignmentsCollectionRequestBuilder) ID(id string) *TermsAndConditionsAssignmentRequestBuilder {
	bb := &TermsAndConditionsAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TermsAndConditionsAssignmentsCollectionRequest is request for TermsAndConditionsAssignment collection
type TermsAndConditionsAssignmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for TermsAndConditionsAssignment collection
func (r *TermsAndConditionsAssignmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *TermsAndConditionsAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for TermsAndConditionsAssignment collection
func (r *TermsAndConditionsAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]TermsAndConditionsAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TermsAndConditionsAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []TermsAndConditionsAssignment
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

// Get performs GET request for TermsAndConditionsAssignment collection
func (r *TermsAndConditionsAssignmentsCollectionRequest) Get() ([]TermsAndConditionsAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for TermsAndConditionsAssignment collection
func (r *TermsAndConditionsAssignmentsCollectionRequest) Add(reqObj *TermsAndConditionsAssignment) (*TermsAndConditionsAssignment, error) {
	return r.Do("POST", "", reqObj)
}

// GroupAssignments returns request builder for TermsAndConditionsGroupAssignment collection
func (b *TermsAndConditionsRequestBuilder) GroupAssignments() *TermsAndConditionsGroupAssignmentsCollectionRequestBuilder {
	bb := &TermsAndConditionsGroupAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/groupAssignments"
	return bb
}

// TermsAndConditionsGroupAssignmentsCollectionRequestBuilder is request builder for TermsAndConditionsGroupAssignment collection
type TermsAndConditionsGroupAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TermsAndConditionsGroupAssignment collection
func (b *TermsAndConditionsGroupAssignmentsCollectionRequestBuilder) Request() *TermsAndConditionsGroupAssignmentsCollectionRequest {
	return &TermsAndConditionsGroupAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TermsAndConditionsGroupAssignment item
func (b *TermsAndConditionsGroupAssignmentsCollectionRequestBuilder) ID(id string) *TermsAndConditionsGroupAssignmentRequestBuilder {
	bb := &TermsAndConditionsGroupAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TermsAndConditionsGroupAssignmentsCollectionRequest is request for TermsAndConditionsGroupAssignment collection
type TermsAndConditionsGroupAssignmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for TermsAndConditionsGroupAssignment collection
func (r *TermsAndConditionsGroupAssignmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *TermsAndConditionsGroupAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for TermsAndConditionsGroupAssignment collection
func (r *TermsAndConditionsGroupAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]TermsAndConditionsGroupAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TermsAndConditionsGroupAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []TermsAndConditionsGroupAssignment
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

// Get performs GET request for TermsAndConditionsGroupAssignment collection
func (r *TermsAndConditionsGroupAssignmentsCollectionRequest) Get() ([]TermsAndConditionsGroupAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for TermsAndConditionsGroupAssignment collection
func (r *TermsAndConditionsGroupAssignmentsCollectionRequest) Add(reqObj *TermsAndConditionsGroupAssignment) (*TermsAndConditionsGroupAssignment, error) {
	return r.Do("POST", "", reqObj)
}
