// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// PlannerRequestBuilder is request builder for Planner
type PlannerRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlannerRequest
func (b *PlannerRequestBuilder) Request() *PlannerRequest {
	return &PlannerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlannerRequest is request for Planner
type PlannerRequest struct{ BaseRequest }

// Do performs HTTP request for Planner
func (r *PlannerRequest) Do(method, path string, reqObj interface{}) (resObj *Planner, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Planner
func (r *PlannerRequest) Get() (*Planner, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Planner
func (r *PlannerRequest) Update(reqObj *Planner) (*Planner, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Planner
func (r *PlannerRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Buckets returns request builder for PlannerBucket collection
func (b *PlannerRequestBuilder) Buckets() *PlannerBucketsCollectionRequestBuilder {
	bb := &PlannerBucketsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/buckets"
	return bb
}

// PlannerBucketsCollectionRequestBuilder is request builder for PlannerBucket collection
type PlannerBucketsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerBucket collection
func (b *PlannerBucketsCollectionRequestBuilder) Request() *PlannerBucketsCollectionRequest {
	return &PlannerBucketsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerBucket item
func (b *PlannerBucketsCollectionRequestBuilder) ID(id string) *PlannerBucketRequestBuilder {
	bb := &PlannerBucketRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerBucketsCollectionRequest is request for PlannerBucket collection
type PlannerBucketsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerBucket collection
func (r *PlannerBucketsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerBucket, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for PlannerBucket collection
func (r *PlannerBucketsCollectionRequest) Paging(method, path string, obj interface{}) ([]PlannerBucket, error) {
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
func (r *PlannerBucketsCollectionRequest) Get() ([]PlannerBucket, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for PlannerBucket collection
func (r *PlannerBucketsCollectionRequest) Add(reqObj *PlannerBucket) (*PlannerBucket, error) {
	return r.Do("POST", "", reqObj)
}

// Plans returns request builder for PlannerPlan collection
func (b *PlannerRequestBuilder) Plans() *PlannerPlansCollectionRequestBuilder {
	bb := &PlannerPlansCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/plans"
	return bb
}

// PlannerPlansCollectionRequestBuilder is request builder for PlannerPlan collection
type PlannerPlansCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerPlan collection
func (b *PlannerPlansCollectionRequestBuilder) Request() *PlannerPlansCollectionRequest {
	return &PlannerPlansCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerPlan item
func (b *PlannerPlansCollectionRequestBuilder) ID(id string) *PlannerPlanRequestBuilder {
	bb := &PlannerPlanRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerPlansCollectionRequest is request for PlannerPlan collection
type PlannerPlansCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerPlan collection
func (r *PlannerPlansCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerPlan, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for PlannerPlan collection
func (r *PlannerPlansCollectionRequest) Paging(method, path string, obj interface{}) ([]PlannerPlan, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PlannerPlan
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []PlannerPlan
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

// Get performs GET request for PlannerPlan collection
func (r *PlannerPlansCollectionRequest) Get() ([]PlannerPlan, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for PlannerPlan collection
func (r *PlannerPlansCollectionRequest) Add(reqObj *PlannerPlan) (*PlannerPlan, error) {
	return r.Do("POST", "", reqObj)
}

// Tasks returns request builder for PlannerTask collection
func (b *PlannerRequestBuilder) Tasks() *PlannerTasksCollectionRequestBuilder {
	bb := &PlannerTasksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tasks"
	return bb
}

// PlannerTasksCollectionRequestBuilder is request builder for PlannerTask collection
type PlannerTasksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerTask collection
func (b *PlannerTasksCollectionRequestBuilder) Request() *PlannerTasksCollectionRequest {
	return &PlannerTasksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerTask item
func (b *PlannerTasksCollectionRequestBuilder) ID(id string) *PlannerTaskRequestBuilder {
	bb := &PlannerTaskRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerTasksCollectionRequest is request for PlannerTask collection
type PlannerTasksCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerTask collection
func (r *PlannerTasksCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerTask, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for PlannerTask collection
func (r *PlannerTasksCollectionRequest) Paging(method, path string, obj interface{}) ([]PlannerTask, error) {
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
func (r *PlannerTasksCollectionRequest) Get() ([]PlannerTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for PlannerTask collection
func (r *PlannerTasksCollectionRequest) Add(reqObj *PlannerTask) (*PlannerTask, error) {
	return r.Do("POST", "", reqObj)
}
