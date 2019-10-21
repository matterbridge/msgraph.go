// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// CommsApplicationRequestBuilder is request builder for CommsApplication
type CommsApplicationRequestBuilder struct{ BaseRequestBuilder }

// Request returns CommsApplicationRequest
func (b *CommsApplicationRequestBuilder) Request() *CommsApplicationRequest {
	return &CommsApplicationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CommsApplicationRequest is request for CommsApplication
type CommsApplicationRequest struct{ BaseRequest }

// Do performs HTTP request for CommsApplication
func (r *CommsApplicationRequest) Do(method, path string, reqObj interface{}) (resObj *CommsApplication, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for CommsApplication
func (r *CommsApplicationRequest) Get() (*CommsApplication, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for CommsApplication
func (r *CommsApplicationRequest) Update(reqObj *CommsApplication) (*CommsApplication, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for CommsApplication
func (r *CommsApplicationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Calls returns request builder for Call collection
func (b *CommsApplicationRequestBuilder) Calls() *CommsApplicationCallsCollectionRequestBuilder {
	bb := &CommsApplicationCallsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calls"
	return bb
}

// CommsApplicationCallsCollectionRequestBuilder is request builder for Call collection
type CommsApplicationCallsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Call collection
func (b *CommsApplicationCallsCollectionRequestBuilder) Request() *CommsApplicationCallsCollectionRequest {
	return &CommsApplicationCallsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Call item
func (b *CommsApplicationCallsCollectionRequestBuilder) ID(id string) *CallRequestBuilder {
	bb := &CallRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CommsApplicationCallsCollectionRequest is request for Call collection
type CommsApplicationCallsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Call collection
func (r *CommsApplicationCallsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Call, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Call collection
func (r *CommsApplicationCallsCollectionRequest) Paging(method, path string, obj interface{}) ([]Call, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Call
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Call
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

// Get performs GET request for Call collection
func (r *CommsApplicationCallsCollectionRequest) Get() ([]Call, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Call collection
func (r *CommsApplicationCallsCollectionRequest) Add(reqObj *Call) (*Call, error) {
	return r.Do("POST", "", reqObj)
}

// OnlineMeetings returns request builder for OnlineMeeting collection
func (b *CommsApplicationRequestBuilder) OnlineMeetings() *CommsApplicationOnlineMeetingsCollectionRequestBuilder {
	bb := &CommsApplicationOnlineMeetingsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/onlineMeetings"
	return bb
}

// CommsApplicationOnlineMeetingsCollectionRequestBuilder is request builder for OnlineMeeting collection
type CommsApplicationOnlineMeetingsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnlineMeeting collection
func (b *CommsApplicationOnlineMeetingsCollectionRequestBuilder) Request() *CommsApplicationOnlineMeetingsCollectionRequest {
	return &CommsApplicationOnlineMeetingsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnlineMeeting item
func (b *CommsApplicationOnlineMeetingsCollectionRequestBuilder) ID(id string) *OnlineMeetingRequestBuilder {
	bb := &OnlineMeetingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CommsApplicationOnlineMeetingsCollectionRequest is request for OnlineMeeting collection
type CommsApplicationOnlineMeetingsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for OnlineMeeting collection
func (r *CommsApplicationOnlineMeetingsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *OnlineMeeting, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for OnlineMeeting collection
func (r *CommsApplicationOnlineMeetingsCollectionRequest) Paging(method, path string, obj interface{}) ([]OnlineMeeting, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OnlineMeeting
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []OnlineMeeting
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

// Get performs GET request for OnlineMeeting collection
func (r *CommsApplicationOnlineMeetingsCollectionRequest) Get() ([]OnlineMeeting, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OnlineMeeting collection
func (r *CommsApplicationOnlineMeetingsCollectionRequest) Add(reqObj *OnlineMeeting) (*OnlineMeeting, error) {
	return r.Do("POST", "", reqObj)
}
