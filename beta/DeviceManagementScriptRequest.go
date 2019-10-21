// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DeviceManagementScriptRequestBuilder is request builder for DeviceManagementScript
type DeviceManagementScriptRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementScriptRequest
func (b *DeviceManagementScriptRequestBuilder) Request() *DeviceManagementScriptRequest {
	return &DeviceManagementScriptRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementScriptRequest is request for DeviceManagementScript
type DeviceManagementScriptRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementScript
func (r *DeviceManagementScriptRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementScript, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceManagementScript
func (r *DeviceManagementScriptRequest) Get() (*DeviceManagementScript, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceManagementScript
func (r *DeviceManagementScriptRequest) Update(reqObj *DeviceManagementScript) (*DeviceManagementScript, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceManagementScript
func (r *DeviceManagementScriptRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Assignments returns request builder for DeviceManagementScriptAssignment collection
func (b *DeviceManagementScriptRequestBuilder) Assignments() *DeviceManagementScriptAssignmentsCollectionRequestBuilder {
	bb := &DeviceManagementScriptAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// DeviceManagementScriptAssignmentsCollectionRequestBuilder is request builder for DeviceManagementScriptAssignment collection
type DeviceManagementScriptAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementScriptAssignment collection
func (b *DeviceManagementScriptAssignmentsCollectionRequestBuilder) Request() *DeviceManagementScriptAssignmentsCollectionRequest {
	return &DeviceManagementScriptAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementScriptAssignment item
func (b *DeviceManagementScriptAssignmentsCollectionRequestBuilder) ID(id string) *DeviceManagementScriptAssignmentRequestBuilder {
	bb := &DeviceManagementScriptAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementScriptAssignmentsCollectionRequest is request for DeviceManagementScriptAssignment collection
type DeviceManagementScriptAssignmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementScriptAssignment collection
func (r *DeviceManagementScriptAssignmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementScriptAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DeviceManagementScriptAssignment collection
func (r *DeviceManagementScriptAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]DeviceManagementScriptAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceManagementScriptAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DeviceManagementScriptAssignment
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

// Get performs GET request for DeviceManagementScriptAssignment collection
func (r *DeviceManagementScriptAssignmentsCollectionRequest) Get() ([]DeviceManagementScriptAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DeviceManagementScriptAssignment collection
func (r *DeviceManagementScriptAssignmentsCollectionRequest) Add(reqObj *DeviceManagementScriptAssignment) (*DeviceManagementScriptAssignment, error) {
	return r.Do("POST", "", reqObj)
}

// DeviceRunStates returns request builder for DeviceManagementScriptDeviceState collection
func (b *DeviceManagementScriptRequestBuilder) DeviceRunStates() *DeviceManagementScriptDeviceRunStatesCollectionRequestBuilder {
	bb := &DeviceManagementScriptDeviceRunStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceRunStates"
	return bb
}

// DeviceManagementScriptDeviceRunStatesCollectionRequestBuilder is request builder for DeviceManagementScriptDeviceState collection
type DeviceManagementScriptDeviceRunStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementScriptDeviceState collection
func (b *DeviceManagementScriptDeviceRunStatesCollectionRequestBuilder) Request() *DeviceManagementScriptDeviceRunStatesCollectionRequest {
	return &DeviceManagementScriptDeviceRunStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementScriptDeviceState item
func (b *DeviceManagementScriptDeviceRunStatesCollectionRequestBuilder) ID(id string) *DeviceManagementScriptDeviceStateRequestBuilder {
	bb := &DeviceManagementScriptDeviceStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementScriptDeviceRunStatesCollectionRequest is request for DeviceManagementScriptDeviceState collection
type DeviceManagementScriptDeviceRunStatesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptDeviceRunStatesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementScriptDeviceState, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptDeviceRunStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]DeviceManagementScriptDeviceState, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceManagementScriptDeviceState
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DeviceManagementScriptDeviceState
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

// Get performs GET request for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptDeviceRunStatesCollectionRequest) Get() ([]DeviceManagementScriptDeviceState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptDeviceRunStatesCollectionRequest) Add(reqObj *DeviceManagementScriptDeviceState) (*DeviceManagementScriptDeviceState, error) {
	return r.Do("POST", "", reqObj)
}

// GroupAssignments returns request builder for DeviceManagementScriptGroupAssignment collection
func (b *DeviceManagementScriptRequestBuilder) GroupAssignments() *DeviceManagementScriptGroupAssignmentsCollectionRequestBuilder {
	bb := &DeviceManagementScriptGroupAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/groupAssignments"
	return bb
}

// DeviceManagementScriptGroupAssignmentsCollectionRequestBuilder is request builder for DeviceManagementScriptGroupAssignment collection
type DeviceManagementScriptGroupAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementScriptGroupAssignment collection
func (b *DeviceManagementScriptGroupAssignmentsCollectionRequestBuilder) Request() *DeviceManagementScriptGroupAssignmentsCollectionRequest {
	return &DeviceManagementScriptGroupAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementScriptGroupAssignment item
func (b *DeviceManagementScriptGroupAssignmentsCollectionRequestBuilder) ID(id string) *DeviceManagementScriptGroupAssignmentRequestBuilder {
	bb := &DeviceManagementScriptGroupAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementScriptGroupAssignmentsCollectionRequest is request for DeviceManagementScriptGroupAssignment collection
type DeviceManagementScriptGroupAssignmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementScriptGroupAssignment collection
func (r *DeviceManagementScriptGroupAssignmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementScriptGroupAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DeviceManagementScriptGroupAssignment collection
func (r *DeviceManagementScriptGroupAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]DeviceManagementScriptGroupAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceManagementScriptGroupAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DeviceManagementScriptGroupAssignment
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

// Get performs GET request for DeviceManagementScriptGroupAssignment collection
func (r *DeviceManagementScriptGroupAssignmentsCollectionRequest) Get() ([]DeviceManagementScriptGroupAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DeviceManagementScriptGroupAssignment collection
func (r *DeviceManagementScriptGroupAssignmentsCollectionRequest) Add(reqObj *DeviceManagementScriptGroupAssignment) (*DeviceManagementScriptGroupAssignment, error) {
	return r.Do("POST", "", reqObj)
}

// RunSummary is navigation property
func (b *DeviceManagementScriptRequestBuilder) RunSummary() *DeviceManagementScriptRunSummaryRequestBuilder {
	bb := &DeviceManagementScriptRunSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/runSummary"
	return bb
}

// UserRunStates returns request builder for DeviceManagementScriptUserState collection
func (b *DeviceManagementScriptRequestBuilder) UserRunStates() *DeviceManagementScriptUserRunStatesCollectionRequestBuilder {
	bb := &DeviceManagementScriptUserRunStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userRunStates"
	return bb
}

// DeviceManagementScriptUserRunStatesCollectionRequestBuilder is request builder for DeviceManagementScriptUserState collection
type DeviceManagementScriptUserRunStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementScriptUserState collection
func (b *DeviceManagementScriptUserRunStatesCollectionRequestBuilder) Request() *DeviceManagementScriptUserRunStatesCollectionRequest {
	return &DeviceManagementScriptUserRunStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementScriptUserState item
func (b *DeviceManagementScriptUserRunStatesCollectionRequestBuilder) ID(id string) *DeviceManagementScriptUserStateRequestBuilder {
	bb := &DeviceManagementScriptUserStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementScriptUserRunStatesCollectionRequest is request for DeviceManagementScriptUserState collection
type DeviceManagementScriptUserRunStatesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementScriptUserState collection
func (r *DeviceManagementScriptUserRunStatesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementScriptUserState, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DeviceManagementScriptUserState collection
func (r *DeviceManagementScriptUserRunStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]DeviceManagementScriptUserState, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceManagementScriptUserState
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DeviceManagementScriptUserState
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

// Get performs GET request for DeviceManagementScriptUserState collection
func (r *DeviceManagementScriptUserRunStatesCollectionRequest) Get() ([]DeviceManagementScriptUserState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DeviceManagementScriptUserState collection
func (r *DeviceManagementScriptUserRunStatesCollectionRequest) Add(reqObj *DeviceManagementScriptUserState) (*DeviceManagementScriptUserState, error) {
	return r.Do("POST", "", reqObj)
}
