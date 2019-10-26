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

// Get performs GET request for TermsAndConditions
func (r *TermsAndConditionsRequest) Get() (resObj *TermsAndConditions, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TermsAndConditions
func (r *TermsAndConditionsRequest) Update(reqObj *TermsAndConditions) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
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
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
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
func (r *TermsAndConditionsAcceptanceStatusesCollectionRequest) Add(reqObj *TermsAndConditionsAcceptanceStatus) (resObj *TermsAndConditionsAcceptanceStatus, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
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
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
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
func (r *TermsAndConditionsAssignmentsCollectionRequest) Add(reqObj *TermsAndConditionsAssignment) (resObj *TermsAndConditionsAssignment, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
