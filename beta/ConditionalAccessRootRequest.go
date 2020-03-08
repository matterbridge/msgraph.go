// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// ConditionalAccessRootRequestBuilder is request builder for ConditionalAccessRoot
type ConditionalAccessRootRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessRootRequest
func (b *ConditionalAccessRootRequestBuilder) Request() *ConditionalAccessRootRequest {
	return &ConditionalAccessRootRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessRootRequest is request for ConditionalAccessRoot
type ConditionalAccessRootRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessRoot
func (r *ConditionalAccessRootRequest) Get(ctx context.Context) (resObj *ConditionalAccessRoot, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessRoot
func (r *ConditionalAccessRootRequest) Update(ctx context.Context, reqObj *ConditionalAccessRoot) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessRoot
func (r *ConditionalAccessRootRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// NamedLocations returns request builder for NamedLocation collection
func (b *ConditionalAccessRootRequestBuilder) NamedLocations() *ConditionalAccessRootNamedLocationsCollectionRequestBuilder {
	bb := &ConditionalAccessRootNamedLocationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/namedLocations"
	return bb
}

// ConditionalAccessRootNamedLocationsCollectionRequestBuilder is request builder for NamedLocation collection
type ConditionalAccessRootNamedLocationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for NamedLocation collection
func (b *ConditionalAccessRootNamedLocationsCollectionRequestBuilder) Request() *ConditionalAccessRootNamedLocationsCollectionRequest {
	return &ConditionalAccessRootNamedLocationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for NamedLocation item
func (b *ConditionalAccessRootNamedLocationsCollectionRequestBuilder) ID(id string) *NamedLocationRequestBuilder {
	bb := &NamedLocationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ConditionalAccessRootNamedLocationsCollectionRequest is request for NamedLocation collection
type ConditionalAccessRootNamedLocationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for NamedLocation collection
func (r *ConditionalAccessRootNamedLocationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]NamedLocation, error) {
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
	var values []NamedLocation
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
			value  []NamedLocation
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

// Get performs GET request for NamedLocation collection
func (r *ConditionalAccessRootNamedLocationsCollectionRequest) Get(ctx context.Context) ([]NamedLocation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for NamedLocation collection
func (r *ConditionalAccessRootNamedLocationsCollectionRequest) Add(ctx context.Context, reqObj *NamedLocation) (resObj *NamedLocation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Policies returns request builder for ConditionalAccessPolicy collection
func (b *ConditionalAccessRootRequestBuilder) Policies() *ConditionalAccessRootPoliciesCollectionRequestBuilder {
	bb := &ConditionalAccessRootPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/policies"
	return bb
}

// ConditionalAccessRootPoliciesCollectionRequestBuilder is request builder for ConditionalAccessPolicy collection
type ConditionalAccessRootPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ConditionalAccessPolicy collection
func (b *ConditionalAccessRootPoliciesCollectionRequestBuilder) Request() *ConditionalAccessRootPoliciesCollectionRequest {
	return &ConditionalAccessRootPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ConditionalAccessPolicy item
func (b *ConditionalAccessRootPoliciesCollectionRequestBuilder) ID(id string) *ConditionalAccessPolicyRequestBuilder {
	bb := &ConditionalAccessPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ConditionalAccessRootPoliciesCollectionRequest is request for ConditionalAccessPolicy collection
type ConditionalAccessRootPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ConditionalAccessPolicy collection
func (r *ConditionalAccessRootPoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ConditionalAccessPolicy, error) {
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
	var values []ConditionalAccessPolicy
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
			value  []ConditionalAccessPolicy
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

// Get performs GET request for ConditionalAccessPolicy collection
func (r *ConditionalAccessRootPoliciesCollectionRequest) Get(ctx context.Context) ([]ConditionalAccessPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ConditionalAccessPolicy collection
func (r *ConditionalAccessRootPoliciesCollectionRequest) Add(ctx context.Context, reqObj *ConditionalAccessPolicy) (resObj *ConditionalAccessPolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
