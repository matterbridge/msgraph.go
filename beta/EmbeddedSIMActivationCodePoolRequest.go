// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// EmbeddedSIMActivationCodePoolRequestBuilder is request builder for EmbeddedSIMActivationCodePool
type EmbeddedSIMActivationCodePoolRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmbeddedSIMActivationCodePoolRequest
func (b *EmbeddedSIMActivationCodePoolRequestBuilder) Request() *EmbeddedSIMActivationCodePoolRequest {
	return &EmbeddedSIMActivationCodePoolRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmbeddedSIMActivationCodePoolRequest is request for EmbeddedSIMActivationCodePool
type EmbeddedSIMActivationCodePoolRequest struct{ BaseRequest }

// Get performs GET request for EmbeddedSIMActivationCodePool
func (r *EmbeddedSIMActivationCodePoolRequest) Get() (resObj *EmbeddedSIMActivationCodePool, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmbeddedSIMActivationCodePool
func (r *EmbeddedSIMActivationCodePoolRequest) Update(reqObj *EmbeddedSIMActivationCodePool) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmbeddedSIMActivationCodePool
func (r *EmbeddedSIMActivationCodePoolRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Assignments returns request builder for EmbeddedSIMActivationCodePoolAssignment collection
func (b *EmbeddedSIMActivationCodePoolRequestBuilder) Assignments() *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder {
	bb := &EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder is request builder for EmbeddedSIMActivationCodePoolAssignment collection
type EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EmbeddedSIMActivationCodePoolAssignment collection
func (b *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder) Request() *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest {
	return &EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EmbeddedSIMActivationCodePoolAssignment item
func (b *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder) ID(id string) *EmbeddedSIMActivationCodePoolAssignmentRequestBuilder {
	bb := &EmbeddedSIMActivationCodePoolAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest is request for EmbeddedSIMActivationCodePoolAssignment collection
type EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EmbeddedSIMActivationCodePoolAssignment collection
func (r *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]EmbeddedSIMActivationCodePoolAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EmbeddedSIMActivationCodePoolAssignment
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
			value  []EmbeddedSIMActivationCodePoolAssignment
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

// Get performs GET request for EmbeddedSIMActivationCodePoolAssignment collection
func (r *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest) Get() ([]EmbeddedSIMActivationCodePoolAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EmbeddedSIMActivationCodePoolAssignment collection
func (r *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest) Add(reqObj *EmbeddedSIMActivationCodePoolAssignment) (resObj *EmbeddedSIMActivationCodePoolAssignment, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// DeviceStates returns request builder for EmbeddedSIMDeviceState collection
func (b *EmbeddedSIMActivationCodePoolRequestBuilder) DeviceStates() *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder {
	bb := &EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStates"
	return bb
}

// EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder is request builder for EmbeddedSIMDeviceState collection
type EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EmbeddedSIMDeviceState collection
func (b *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder) Request() *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest {
	return &EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EmbeddedSIMDeviceState item
func (b *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder) ID(id string) *EmbeddedSIMDeviceStateRequestBuilder {
	bb := &EmbeddedSIMDeviceStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest is request for EmbeddedSIMDeviceState collection
type EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EmbeddedSIMDeviceState collection
func (r *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]EmbeddedSIMDeviceState, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EmbeddedSIMDeviceState
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
			value  []EmbeddedSIMDeviceState
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

// Get performs GET request for EmbeddedSIMDeviceState collection
func (r *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest) Get() ([]EmbeddedSIMDeviceState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EmbeddedSIMDeviceState collection
func (r *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest) Add(reqObj *EmbeddedSIMDeviceState) (resObj *EmbeddedSIMDeviceState, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
