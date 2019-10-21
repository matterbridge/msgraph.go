// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// PlannerPlanRequestBuilder is request builder for PlannerPlan
type PlannerPlanRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlannerPlanRequest
func (b *PlannerPlanRequestBuilder) Request() *PlannerPlanRequest {
	return &PlannerPlanRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlannerPlanRequest is request for PlannerPlan
type PlannerPlanRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerPlan
func (r *PlannerPlanRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerPlan, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for PlannerPlan
func (r *PlannerPlanRequest) Get() (*PlannerPlan, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for PlannerPlan
func (r *PlannerPlanRequest) Update(reqObj *PlannerPlan) (*PlannerPlan, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for PlannerPlan
func (r *PlannerPlanRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Buckets returns request builder for PlannerBucket collection
func (b *PlannerPlanRequestBuilder) Buckets() *PlannerPlanBucketsCollectionRequestBuilder {
	bb := &PlannerPlanBucketsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/buckets"
	return bb
}

// PlannerPlanBucketsCollectionRequestBuilder is request builder for PlannerBucket collection
type PlannerPlanBucketsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerBucket collection
func (b *PlannerPlanBucketsCollectionRequestBuilder) Request() *PlannerPlanBucketsCollectionRequest {
	return &PlannerPlanBucketsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerBucket item
func (b *PlannerPlanBucketsCollectionRequestBuilder) ID(id string) *PlannerBucketRequestBuilder {
	bb := &PlannerBucketRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerPlanBucketsCollectionRequest is request for PlannerBucket collection
type PlannerPlanBucketsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerBucket collection
func (r *PlannerPlanBucketsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerBucket, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for PlannerBucket collection
func (r *PlannerPlanBucketsCollectionRequest) Paging(method, path string, obj interface{}) ([]PlannerBucket, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PlannerBucket
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []PlannerBucket
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

// Get performs GET request for PlannerBucket collection
func (r *PlannerPlanBucketsCollectionRequest) Get() ([]PlannerBucket, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for PlannerBucket collection
func (r *PlannerPlanBucketsCollectionRequest) Add(reqObj *PlannerBucket) (*PlannerBucket, error) {
	return r.Do("POST", "", reqObj)
}

// Details is navigation property
func (b *PlannerPlanRequestBuilder) Details() *PlannerPlanDetailsRequestBuilder {
	bb := &PlannerPlanDetailsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/details"
	return bb
}

// Tasks returns request builder for PlannerTask collection
func (b *PlannerPlanRequestBuilder) Tasks() *PlannerPlanTasksCollectionRequestBuilder {
	bb := &PlannerPlanTasksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tasks"
	return bb
}

// PlannerPlanTasksCollectionRequestBuilder is request builder for PlannerTask collection
type PlannerPlanTasksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerTask collection
func (b *PlannerPlanTasksCollectionRequestBuilder) Request() *PlannerPlanTasksCollectionRequest {
	return &PlannerPlanTasksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerTask item
func (b *PlannerPlanTasksCollectionRequestBuilder) ID(id string) *PlannerTaskRequestBuilder {
	bb := &PlannerTaskRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerPlanTasksCollectionRequest is request for PlannerTask collection
type PlannerPlanTasksCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerTask collection
func (r *PlannerPlanTasksCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerTask, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for PlannerTask collection
func (r *PlannerPlanTasksCollectionRequest) Paging(method, path string, obj interface{}) ([]PlannerTask, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PlannerTask
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []PlannerTask
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

// Get performs GET request for PlannerTask collection
func (r *PlannerPlanTasksCollectionRequest) Get() ([]PlannerTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for PlannerTask collection
func (r *PlannerPlanTasksCollectionRequest) Add(reqObj *PlannerTask) (*PlannerTask, error) {
	return r.Do("POST", "", reqObj)
}