// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// WindowsManagementAppRequestBuilder is request builder for WindowsManagementApp
type WindowsManagementAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsManagementAppRequest
func (b *WindowsManagementAppRequestBuilder) Request() *WindowsManagementAppRequest {
	return &WindowsManagementAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsManagementAppRequest is request for WindowsManagementApp
type WindowsManagementAppRequest struct{ BaseRequest }

// Get performs GET request for WindowsManagementApp
func (r *WindowsManagementAppRequest) Get() (resObj *WindowsManagementApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsManagementApp
func (r *WindowsManagementAppRequest) Update(reqObj *WindowsManagementApp) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsManagementApp
func (r *WindowsManagementAppRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// HealthStates returns request builder for WindowsManagementAppHealthState collection
func (b *WindowsManagementAppRequestBuilder) HealthStates() *WindowsManagementAppHealthStatesCollectionRequestBuilder {
	bb := &WindowsManagementAppHealthStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/healthStates"
	return bb
}

// WindowsManagementAppHealthStatesCollectionRequestBuilder is request builder for WindowsManagementAppHealthState collection
type WindowsManagementAppHealthStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WindowsManagementAppHealthState collection
func (b *WindowsManagementAppHealthStatesCollectionRequestBuilder) Request() *WindowsManagementAppHealthStatesCollectionRequest {
	return &WindowsManagementAppHealthStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WindowsManagementAppHealthState item
func (b *WindowsManagementAppHealthStatesCollectionRequestBuilder) ID(id string) *WindowsManagementAppHealthStateRequestBuilder {
	bb := &WindowsManagementAppHealthStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsManagementAppHealthStatesCollectionRequest is request for WindowsManagementAppHealthState collection
type WindowsManagementAppHealthStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WindowsManagementAppHealthState collection
func (r *WindowsManagementAppHealthStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]WindowsManagementAppHealthState, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WindowsManagementAppHealthState
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
			value  []WindowsManagementAppHealthState
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

// Get performs GET request for WindowsManagementAppHealthState collection
func (r *WindowsManagementAppHealthStatesCollectionRequest) Get() ([]WindowsManagementAppHealthState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WindowsManagementAppHealthState collection
func (r *WindowsManagementAppHealthStatesCollectionRequest) Add(reqObj *WindowsManagementAppHealthState) (resObj *WindowsManagementAppHealthState, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
